package ngap

import (
	"encoding/hex"
	"io"
	"net"
	"sync"
	"syscall"

	"etrib5gc/nfs/amf"
	"etrib5gc/nfs/amf/nas"

	"git.cs.nctu.edu.tw/calee/sctp"

	libngap "github.com/free5gc/ngap"
)

type Server struct {
	connections sync.Map
	listener    *sctp.SCTPListener
	ngap        *Ngap
}

const readBufSize uint32 = 8192

// set default read timeout to 2 seconds
var readTimeout syscall.Timeval = syscall.Timeval{Sec: 2, Usec: 0}

var sctpConfig sctp.SocketConfig = sctp.SocketConfig{
	InitMsg:   sctp.InitMsg{NumOstreams: 3, MaxInstreams: 5, MaxAttempts: 2, MaxInitTimeout: 2},
	RtoInfo:   &sctp.RtoInfo{SrtoAssocID: 0, SrtoInitial: 500, SrtoMax: 1500, StroMin: 100},
	AssocInfo: &sctp.AssocInfo{AsocMaxRxt: 4},
}

func NewServer(b amf.Backend) *Server {
	return &Server{
		ngap: NewNgap(b),
	}
}

// expose ngap message sender
func (s *Server) Sender() *ngapSender {
	return &s.ngap.sender
}

// expose Nas
func (s *Server) Nas() *nas.Nas {
	return s.ngap.nas
}

func (s *Server) Run(addresses []string, port int) {

	if s.listener != nil {
		log.Errorf("Server is running")
		return
	}

	ips := []net.IPAddr{}

	for _, addr := range addresses {
		if netAddr, err := net.ResolveIPAddr("ip", addr); err != nil {
			log.Errorf("Error resolving address '%s': %v\n", addr, err)
		} else {
			log.Debugf("Resolved address '%s' to %s\n", addr, netAddr)
			ips = append(ips, *netAddr)
		}
	}

	addr := &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    port,
	}

	go s.serve(addr)
}

func (s *Server) serve(addr *sctp.SCTPAddr) {
	if listener, err := sctpConfig.Listen("sctp", addr); err != nil {
		log.Errorf("Failed to listen: %+v", err)
		return
	} else {
		s.listener = listener
	}

	log.Infof("Listen on %s", s.listener.Addr())

	for {
		newConn, err := s.listener.AcceptSCTP()
		if err != nil {
			switch err {
			case syscall.EINTR, syscall.EAGAIN:
				log.Debugf("AcceptSCTP: %+v", err)
			default:
				log.Errorf("Failed to accept: %+v", err)
			}
			continue
		}

		var info *sctp.SndRcvInfo
		if infoTmp, err := newConn.GetDefaultSentParam(); err != nil {
			log.Errorf("Get default sent param error: %+v, accept failed", err)
			if err = newConn.Close(); err != nil {
				log.Errorf("Close error: %+v", err)
			}
			continue
		} else {
			info = infoTmp
			log.Debugf("Get default sent param[value: %+v]", info)
		}

		info.PPID = libngap.PPID
		if err := newConn.SetDefaultSentParam(info); err != nil {
			log.Errorf("Set default sent param error: %+v, accept failed", err)
			if err = newConn.Close(); err != nil {
				log.Errorf("Close error: %+v", err)
			}
			continue
		} else {
			log.Debugf("Set default sent param[value: %+v]", info)
		}

		events := sctp.SCTP_EVENT_DATA_IO | sctp.SCTP_EVENT_SHUTDOWN | sctp.SCTP_EVENT_ASSOCIATION
		if err := newConn.SubscribeEvents(events); err != nil {
			log.Errorf("Failed to accept: %+v", err)
			if err = newConn.Close(); err != nil {
				log.Errorf("Close error: %+v", err)
			}
			continue
		} else {
			log.Debugln("Subscribe SCTP event[DATA_IO, SHUTDOWN_EVENT, ASSOCIATION_CHANGE]")
		}

		if err := newConn.SetReadBuffer(int(readBufSize)); err != nil {
			log.Errorf("Set read buffer error: %+v, accept failed", err)
			if err = newConn.Close(); err != nil {
				log.Errorf("Close error: %+v", err)
			}
			continue
		} else {
			log.Debugf("Set read buffer to %d bytes", readBufSize)
		}

		if err := newConn.SetReadTimeout(readTimeout); err != nil {
			log.Errorf("Set read timeout error: %+v, accept failed", err)
			if err = newConn.Close(); err != nil {
				log.Errorf("Close error: %+v", err)
			}
			continue
		} else {
			log.Debugf("Set read timeout: %+v", readTimeout)
		}

		log.Infof("[AMF] SCTP Accept from: %s", newConn.RemoteAddr().String())
		s.connections.Store(newConn, newConn)

		go s.handleConnection(newConn, readBufSize)
	}
}

func (s *Server) Stop() {
	if s.listener == nil {
		log.Infof("No SCTP server to stop")
		return
	}

	log.Infof("Close SCTP server...")
	if err := s.listener.Close(); err != nil {
		log.Error(err)
		log.Infof("SCTP server may not close normally.")
	}

	s.connections.Range(func(key, value interface{}) bool {
		conn := value.(net.Conn)
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
		return true
	})

	log.Infof("SCTP server closed")
}

func (s *Server) handleConnection(conn *sctp.SCTPConn, bufsize uint32) {
	defer func() {
		// if AMF call Stop(), then conn.Close() will return EBADF because conn has been closed inside Stop()
		if err := conn.Close(); err != nil && err != syscall.EBADF {
			log.Errorf("close connection error: %+v", err)
		}
		s.connections.Delete(conn)
	}()

	for {
		buf := make([]byte, bufsize)

		n, info, notification, err := conn.SCTPRead(buf)
		if err != nil {
			switch err {
			case io.EOF, io.ErrUnexpectedEOF:
				log.Debugln("Read EOF from client")
				return
			case syscall.EAGAIN:
				log.Debugln("SCTP read timeout")
				continue
			case syscall.EINTR:
				log.Debugf("SCTPRead: %+v", err)
				continue
			default:
				log.Errorf("Handle connection[addr: %+v] error: %+v", conn.RemoteAddr(), err)
				return
			}
		}

		if notification != nil {
			s.ngap.HandleSCTPNotification(conn, notification)
		} else {
			if info == nil || info.PPID != libngap.PPID {
				log.Warnln("Received SCTP PPID != 60, discard this packet")
				continue
			}

			log.Tracef("Read %d bytes", n)
			log.Tracef("Packet content:\n%+v", hex.Dump(buf[:n]))

			// TODO: concurrent on per-UE message
			s.ngap.HandleMessage(conn, buf[:n])
		}
	}
}

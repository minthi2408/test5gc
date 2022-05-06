package ngap

import (
	"net"
	"etri5gc/nfs/amf"
	"etri5gc/nfs/amf/nas"


	"github.com/sirupsen/logrus"
	"github.com/free5gc/ngap/ngapType"
	libngap "github.com/free5gc/ngap"
	"git.cs.nctu.edu.tw/calee/sctp"
)	

var log	*logrus.Entry
func init() {
	log = logrus.WithFields(logrus.Fields{"tag": "ngap"})
}

type ngapHandler struct {
	backend amf.Backend
	sender	*ngapSender
	nas		*nas.Nas
}

type Ngap struct {
	backend		amf.Backend
	sender		ngapSender
	handler 	ngapHandler
	nas			*nas.Nas
}


func NewNgap(b amf.Backend) *Ngap {
	ret := &Ngap{
		backend: b,
		sender: ngapSender{
			backend: b,
		},
		handler: ngapHandler{
			backend: b,
		},
	}
	//create Nas
	ret.nas = nas.NewNas(b, &ret.sender)
	ret.handler.sender = &ret.sender
	ret.handler.nas = ret.nas
	return ret
}

//expose ngap message sender
func (n *Ngap) Sender() *ngapSender {
	return &n.sender
}

//expose Nas
func (n *Ngap) Nas() *nas.Nas {
	return n.nas
}

func (h *Ngap) HandleMessage(conn net.Conn, pdu []byte) {
	amf := h.backend.Context()
	ran, ok := amf.AmfRanFindByConn(conn)
	if !ok {
		log.Infof("Create a new NG connection for: %s", conn.RemoteAddr().String())
		ran = amf.NewAmfRan(conn)
	}

	if len(pdu) == 0 {
		log.Infof("RAN close the connection.")
		amf.RemoveRan(ran)
		return
	}

	ngapmsg, err := libngap.Decoder(pdu)
	if err != nil {
		log.Errorf("NGAP decode error : %+v", err)
		return
	}

	switch ngapmsg.Present {
	case ngapType.NGAPPDUPresentInitiatingMessage:
		iMsg := ngapmsg.InitiatingMessage
		if iMsg == nil {
			log.Errorln("Initiating Message is nil")
			return
		}
		switch iMsg.ProcedureCode.Value {
		case ngapType.ProcedureCodeNGSetup:
			if iMsg.Value.NGSetupRequest != nil {
				h.handler.handleNGSetupRequest(ran, iMsg.Value.NGSetupRequest)
				return
			}
		case ngapType.ProcedureCodeInitialUEMessage:
			if iMsg.Value.InitialUEMessage != nil {
				h.handler.handleInitialUEMessage(ran, iMsg.Value.InitialUEMessage, pdu)
			}
		case ngapType.ProcedureCodeUplinkNASTransport:
			if iMsg.Value.UplinkNASTransport != nil {
				h.handler.handleUplinkNasTransport(ran, iMsg.Value.UplinkNASTransport)
				return
			}
		case ngapType.ProcedureCodeNGReset:
			if iMsg.Value.NGReset != nil {
				h.handler.handleNGReset(ran, iMsg.Value.NGReset)
				return
			}
		case ngapType.ProcedureCodeHandoverCancel:
			if iMsg.Value.HandoverCancel != nil {
				h.handler.handleHandoverCancel(ran, iMsg.Value.HandoverCancel)
				return
			}
		case ngapType.ProcedureCodeUEContextReleaseRequest:
			if iMsg.Value.UEContextReleaseRequest != nil {
				h.handler.handleUEContextReleaseRequest(ran, iMsg.Value.UEContextReleaseRequest)
				return
			}
		case ngapType.ProcedureCodeNASNonDeliveryIndication:
			if iMsg.Value.NASNonDeliveryIndication != nil {
				h.handler.handleNasNonDeliveryIndication(ran, iMsg.Value.NASNonDeliveryIndication)
				return
			}
		case ngapType.ProcedureCodeLocationReportingFailureIndication:
			if iMsg.Value.LocationReportingFailureIndication != nil {
				h.handler.handleLocationReportingFailureIndication(ran, iMsg.Value.LocationReportingFailureIndication)
				return
			}

		case ngapType.ProcedureCodeErrorIndication:
			if iMsg.Value.ErrorIndication != nil {
				h.handler.handleErrorIndication(ran, iMsg.Value.ErrorIndication)
				return
			}
		case ngapType.ProcedureCodeUERadioCapabilityInfoIndication:
			if iMsg.Value.UERadioCapabilityInfoIndication != nil {
				h.handler.handleUERadioCapabilityInfoIndication(ran, iMsg.Value.UERadioCapabilityInfoIndication)
				return
			}
		case ngapType.ProcedureCodeHandoverNotification:
			if iMsg.Value.HandoverNotify != nil {
				h.handler.handleHandoverNotify(ran, iMsg.Value.HandoverNotify)
				return
			}
		case ngapType.ProcedureCodeHandoverPreparation:
			if iMsg.Value.HandoverRequired != nil {
				h.handler.handleHandoverRequired(ran, iMsg.Value.HandoverRequired)
				return
			}
		case ngapType.ProcedureCodeRANConfigurationUpdate:
			if iMsg.Value.RANConfigurationUpdate != nil {
				h.handler.handleRanConfigurationUpdate(ran, iMsg.Value.RANConfigurationUpdate)
				return
			}
		case ngapType.ProcedureCodeRRCInactiveTransitionReport:
			if iMsg.Value.RRCInactiveTransitionReport != nil {
				h.handler.handleRRCInactiveTransitionReport(ran, iMsg.Value.RRCInactiveTransitionReport)
				return
			}
		case ngapType.ProcedureCodePDUSessionResourceNotify:
			if iMsg.Value.PDUSessionResourceNotify != nil {
				h.handler.handlePDUSessionResourceNotify(ran, iMsg.Value.PDUSessionResourceNotify)
				return
			}
		case ngapType.ProcedureCodePathSwitchRequest:
			if iMsg.Value.PathSwitchRequest != nil {
				h.handler.handlePathSwitchRequest(ran, iMsg.Value.PathSwitchRequest)
				return
			}
		case ngapType.ProcedureCodeLocationReport:
			if iMsg.Value.LocationReport != nil {
				h.handler.handleLocationReport(ran, iMsg.Value.LocationReport)
				return
			}
		case ngapType.ProcedureCodeUplinkUEAssociatedNRPPaTransport:
			if iMsg.Value.UplinkUEAssociatedNRPPaTransport != nil {
				h.handler.handleUplinkUEAssociatedNRPPATransport(ran, iMsg.Value.UplinkUEAssociatedNRPPaTransport)
				return
			}
		case ngapType.ProcedureCodeUplinkRANConfigurationTransfer:
			if iMsg.Value.UplinkRANConfigurationTransfer != nil {
				h.handler.handleUplinkRanConfigurationTransfer(ran, iMsg.Value.UplinkRANConfigurationTransfer)
				return
			}
		case ngapType.ProcedureCodePDUSessionResourceModifyIndication:
			if iMsg.Value.PDUSessionResourceModifyIndication != nil {
				h.handler.handlePDUSessionResourceModifyIndication(ran, iMsg.Value.PDUSessionResourceModifyIndication)
				return
			}
		case ngapType.ProcedureCodeCellTrafficTrace:
			if iMsg.Value.CellTrafficTrace != nil {
					h.handler.handleCellTrafficTrace(ran, iMsg.Value.CellTrafficTrace)
				return
			}
		case ngapType.ProcedureCodeUplinkRANStatusTransfer:
			if iMsg.Value.UplinkRANStatusTransfer != nil {
				h.handler.handleUplinkRanStatusTransfer(ran, iMsg.Value.UplinkRANStatusTransfer)
				return
			}
		case ngapType.ProcedureCodeUplinkNonUEAssociatedNRPPaTransport:
			if iMsg.Value.UplinkNonUEAssociatedNRPPaTransport != nil {
				h.handler.handleUplinkNonUEAssociatedNRPPATransport(ran, iMsg.Value.UplinkNonUEAssociatedNRPPaTransport) 
				return
			}
		default:
			log.Warnf("Not implemented(choice:%d, procedureCode:%d)\n", ngapmsg.Present, iMsg.ProcedureCode.Value)
		}
	case ngapType.NGAPPDUPresentSuccessfulOutcome:
		sMsg := ngapmsg.SuccessfulOutcome
		if sMsg == nil {
			log.Errorln("successful Outcome is nil")
			return
		}
		switch sMsg.ProcedureCode.Value {
		case ngapType.ProcedureCodeNGReset:
		//	h.handler.handleNGResetAcknowledge(ran, pdu)
		case ngapType.ProcedureCodeUEContextRelease:
		//	h.handler.handleUEContextReleaseComplete(ran, pdu)
		case ngapType.ProcedureCodePDUSessionResourceRelease:
		//	h.handler.handlePDUSessionResourceReleaseResponse(ran, pdu)
		case ngapType.ProcedureCodeUERadioCapabilityCheck:
		//	h.handler.handleUERadioCapabilityCheckResponse(ran, pdu)
		case ngapType.ProcedureCodeAMFConfigurationUpdate:
		//	h.handler.handleAMFconfigurationUpdateAcknowledge(ran, pdu)
		case ngapType.ProcedureCodeInitialContextSetup:
		//	h.handler.handleInitialContextSetupResponse(ran, pdu)
		case ngapType.ProcedureCodeUEContextModification:
		//	h.handler.handleUEContextModificationResponse(ran, pdu)
		case ngapType.ProcedureCodePDUSessionResourceSetup:
		//	h.handler.handlePDUSessionResourceSetupResponse(ran, pdu)
		case ngapType.ProcedureCodePDUSessionResourceModify:
		//	h.handler.handlePDUSessionResourceModifyResponse(ran, pdu)
		case ngapType.ProcedureCodeHandoverResourceAllocation:
			if content := sMsg.Value.HandoverRequestAcknowledge; content != nil {
				h.handler.handleHandoverRequestAcknowledge(ran, content)
				return
			}
		default:
			log.Warnf("Not implemented(choice:%d, procedureCode:%d)\n", ngapmsg.Present, sMsg.ProcedureCode.Value)
		}
	case ngapType.NGAPPDUPresentUnsuccessfulOutcome:
		uMsg := ngapmsg.UnsuccessfulOutcome
		if uMsg == nil {
			log.Errorln("unsuccessful Outcome is nil")
			return
		}
		switch uMsg.ProcedureCode.Value {
		case ngapType.ProcedureCodeAMFConfigurationUpdate:
		//	h.handler.handleAMFconfigurationUpdateFailure(ran, pdu)
		case ngapType.ProcedureCodeInitialContextSetup:
		//	h.handler.handleInitialContextSetupFailure(ran, pdu)
		case ngapType.ProcedureCodeUEContextModification:
		//	h.handler.handleUEContextModificationFailure(ran, pdu)
		case ngapType.ProcedureCodeHandoverResourceAllocation:
		//	h.handler.handleHandoverFailure(ran, pdu)
		default:
			log.Warnf("Not implemented(choice:%d, procedureCode:%d)\n", ngapmsg.Present, uMsg.ProcedureCode.Value)
		}
	}
}

func (h *Ngap) HandleSCTPNotification(conn net.Conn, notification sctp.Notification) {
	amf := h.backend.Context()

	log.Infof("Handle SCTP Notification[addr: %+v]", conn.RemoteAddr())

	ran, ok := h.backend.Context().AmfRanFindByConn(conn)
	if !ok {
		log.Warnf("RAN context has been removed[addr: %+v]", conn.RemoteAddr())
		return
	}

	switch notification.Type() {
	case sctp.SCTP_ASSOC_CHANGE:
		log.Infof("SCTP_ASSOC_CHANGE notification")
		event := notification.(*sctp.SCTPAssocChangeEvent)
		switch event.State() {
		case sctp.SCTP_COMM_LOST:
			log.Infof("SCTP state is SCTP_COMM_LOST, close the connection")
			amf.RemoveRan(ran)
		case sctp.SCTP_SHUTDOWN_COMP:
			log.Infof("SCTP state is SCTP_SHUTDOWN_COMP, close the connection")
			amf.RemoveRan(ran)
		default:
			log.Warnf("SCTP state[%+v] is not handled", event.State())
		}
	case sctp.SCTP_SHUTDOWN_EVENT:
		log.Infof("SCTP_SHUTDOWN_EVENT notification, close the connection")
		amf.RemoveRan(ran)
	default:
		log.Warnf("Non handled notification type: 0x%x", notification.Type())
	}
}

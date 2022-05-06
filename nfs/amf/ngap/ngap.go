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

func (h *Ngap) HandleMessage(conn net.Conn, msg []byte) {
	amf := h.backend.Context()
	ran, ok := amf.AmfRanFindByConn(conn)
	if !ok {
		log.Infof("Create a new NG connection for: %s", conn.RemoteAddr().String())
		ran = amf.NewAmfRan(conn)
	}

	if len(msg) == 0 {
		log.Infof("RAN close the connection.")
		amf.RemoveRan(ran)
		return
	}

	pdu, err := libngap.Decoder(msg)
	if err != nil {
		log.Errorf("NGAP decode error : %+v", err)
		return
	}

	switch pdu.Present {
	case ngapType.NGAPPDUPresentInitiatingMessage:
		initiatingMessage := pdu.InitiatingMessage
		if initiatingMessage == nil {
			log.Errorln("Initiating Message is nil")
			return
		}
		switch initiatingMessage.ProcedureCode.Value {
		case ngapType.ProcedureCodeNGSetup:
			h.handler.handleNGSetupRequest(ran, pdu)
		case ngapType.ProcedureCodeInitialUEMessage:
			h.handler.handleInitialUEMessage(ran, pdu)
		case ngapType.ProcedureCodeUplinkNASTransport:
			h.handler.handleUplinkNasTransport(ran, pdu)
		case ngapType.ProcedureCodeNGReset:
			h.handler.handleNGReset(ran, pdu)
		case ngapType.ProcedureCodeHandoverCancel:
			h.handler.handleHandoverCancel(ran, pdu)
		case ngapType.ProcedureCodeUEContextReleaseRequest:
			h.handler.handleUEContextReleaseRequest(ran, pdu)
		case ngapType.ProcedureCodeNASNonDeliveryIndication:
			h.handler.handleNasNonDeliveryIndication(ran, pdu)
		case ngapType.ProcedureCodeLocationReportingFailureIndication:
			h.handler.handleLocationReportingFailureIndication(ran, pdu)
		case ngapType.ProcedureCodeErrorIndication:
			h.handler.handleErrorIndication(ran, pdu)
		case ngapType.ProcedureCodeUERadioCapabilityInfoIndication:
			h.handler.handleUERadioCapabilityInfoIndication(ran, pdu)
		case ngapType.ProcedureCodeHandoverNotification:
			h.handler.handleHandoverNotify(ran, pdu)
		case ngapType.ProcedureCodeHandoverPreparation:
			h.handler.handleHandoverRequired(ran, pdu)
		case ngapType.ProcedureCodeRANConfigurationUpdate:
			h.handler.handleRanConfigurationUpdate(ran, pdu)
		case ngapType.ProcedureCodeRRCInactiveTransitionReport:
			h.handler.handleRRCInactiveTransitionReport(ran, pdu)
		case ngapType.ProcedureCodePDUSessionResourceNotify:
			h.handler.handlePDUSessionResourceNotify(ran, pdu)
		case ngapType.ProcedureCodePathSwitchRequest:
			h.handler.handlePathSwitchRequest(ran, pdu)
		case ngapType.ProcedureCodeLocationReport:
			h.handler.handleLocationReport(ran, pdu)
		case ngapType.ProcedureCodeUplinkUEAssociatedNRPPaTransport:
			h.handler.handleUplinkUEAssociatedNRPPATransport(ran, pdu)
		case ngapType.ProcedureCodeUplinkRANConfigurationTransfer:
			h.handler.handleUplinkRanConfigurationTransfer(ran, pdu)
		case ngapType.ProcedureCodePDUSessionResourceModifyIndication:
			h.handler.handlePDUSessionResourceModifyIndication(ran, pdu)
		case ngapType.ProcedureCodeCellTrafficTrace:
			h.handler.handleCellTrafficTrace(ran, pdu)
		case ngapType.ProcedureCodeUplinkRANStatusTransfer:
			h.handler.handleUplinkRanStatusTransfer(ran, pdu)
		case ngapType.ProcedureCodeUplinkNonUEAssociatedNRPPaTransport:
			h.handler.handleUplinkNonUEAssociatedNRPPATransport(ran, pdu)
		default:
			log.Warnf("Not implemented(choice:%d, procedureCode:%d)\n", pdu.Present, initiatingMessage.ProcedureCode.Value)
		}
	case ngapType.NGAPPDUPresentSuccessfulOutcome:
		successfulOutcome := pdu.SuccessfulOutcome
		if successfulOutcome == nil {
			log.Errorln("successful Outcome is nil")
			return
		}
		switch successfulOutcome.ProcedureCode.Value {
		case ngapType.ProcedureCodeNGReset:
			h.handler.handleNGResetAcknowledge(ran, pdu)
		case ngapType.ProcedureCodeUEContextRelease:
			h.handler.handleUEContextReleaseComplete(ran, pdu)
		case ngapType.ProcedureCodePDUSessionResourceRelease:
			h.handler.handlePDUSessionResourceReleaseResponse(ran, pdu)
		case ngapType.ProcedureCodeUERadioCapabilityCheck:
			h.handler.handleUERadioCapabilityCheckResponse(ran, pdu)
		case ngapType.ProcedureCodeAMFConfigurationUpdate:
			h.handler.handleAMFconfigurationUpdateAcknowledge(ran, pdu)
		case ngapType.ProcedureCodeInitialContextSetup:
			h.handler.handleInitialContextSetupResponse(ran, pdu)
		case ngapType.ProcedureCodeUEContextModification:
			h.handler.handleUEContextModificationResponse(ran, pdu)
		case ngapType.ProcedureCodePDUSessionResourceSetup:
			h.handler.handlePDUSessionResourceSetupResponse(ran, pdu)
		case ngapType.ProcedureCodePDUSessionResourceModify:
			h.handler.handlePDUSessionResourceModifyResponse(ran, pdu)
		case ngapType.ProcedureCodeHandoverResourceAllocation:
			msg := successfulOutcome.Value.HandoverRequestAcknowledge
			h.handler.handleHandoverRequestAcknowledge(ran, msg)
		default:
			log.Warnf("Not implemented(choice:%d, procedureCode:%d)\n", pdu.Present, successfulOutcome.ProcedureCode.Value)
		}
	case ngapType.NGAPPDUPresentUnsuccessfulOutcome:
		unsuccessfulOutcome := pdu.UnsuccessfulOutcome
		if unsuccessfulOutcome == nil {
			log.Errorln("unsuccessful Outcome is nil")
			return
		}
		switch unsuccessfulOutcome.ProcedureCode.Value {
		case ngapType.ProcedureCodeAMFConfigurationUpdate:
			h.handler.handleAMFconfigurationUpdateFailure(ran, pdu)
		case ngapType.ProcedureCodeInitialContextSetup:
			h.handler.handleInitialContextSetupFailure(ran, pdu)
		case ngapType.ProcedureCodeUEContextModification:
			h.handler.handleUEContextModificationFailure(ran, pdu)
		case ngapType.ProcedureCodeHandoverResourceAllocation:
			h.handler.handleHandoverFailure(ran, pdu)
		default:
			log.Warnf("Not implemented(choice:%d, procedureCode:%d)\n", pdu.Present, unsuccessfulOutcome.ProcedureCode.Value)
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

package nas

import (
	"errors"
	"fmt"

	"etri5gc/nfs/amf"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/nas/nas_security"

	"github.com/free5gc/aper"
	libnas "github.com/free5gc/nas"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/fsm"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "nas"})
}

type NgapSender interface {
	SendDownlinkNasTransport(*context.RanUe, []byte, *ngapType.MobilityRestrictionList)
	SendInitialContextSetupRequest(
		*context.AmfUe,
		models.AccessType,
		[]byte,
		*ngapType.PDUSessionResourceSetupListCxtReq,
		*ngapType.RRCInactiveTransitionReportRequest,
		*ngapType.CoreNetworkAssistanceInformation,
		*ngapType.EmergencyFallbackIndicator)

	SendPDUSessionResourceReleaseCommand(*context.RanUe, []byte, ngapType.PDUSessionResourceToReleaseListRelCmd)
	SendPDUSessionResourceModifyRequest(*context.RanUe, ngapType.PDUSessionResourceModifyListModReq)

	SendPDUSessionResourceSetupRequest(*context.RanUe, []byte, ngapType.PDUSessionResourceSetupListSUReq)
	SendRerouteNasRequest(*context.AmfUe, models.AccessType, *int64, []byte, *ngapType.AllowedNSSAI)

	SendUEContextReleaseCommand(*context.RanUe, context.RelAction, int, aper.Enumerated)
}

type Nas struct {
	backend amf.Backend
	ngap    NgapSender
	sm      *fsm.FSM
}

func NewNas(b amf.Backend, ngap NgapSender) *Nas {
	ret := &Nas{
		backend: b,
		ngap:    ngap,
	}
	gmm := newGmmFsm(ret)
	ret.sm = gmm.sm
	return ret
}

func (n *Nas) amf() *context.AMFContext {
	return n.backend.Context()
}

func (n *Nas) Fsm() *fsm.FSM {
	return n.sm
}

func (n *Nas) HandleNAS(ue *context.RanUe, code int64, naspdu []byte) {
	if ue.AmfUe() == nil {
		newAmfUe := n.backend.Context().NewAmfUe("")
		newAmfUe.AttachRanUe(ue)
	}

	msg, err := nas_security.Decode(ue.AmfUe(), ue.Ran().AnType(), naspdu)
	if err != nil {
		log.Errorln(err)
		return
	}

	if err := n.dispatch(ue.AmfUe(), ue.Ran().AnType(), code, msg); err != nil {
		log.Errorf("Handle NAS Error: %v", err)
	}

}

func (n *Nas) dispatch(ue *context.AmfUe, atype models.AccessType, code int64, msg *libnas.Message) error {
	if msg.GmmMessage == nil {
		return errors.New("Gmm Message is nil")
	}

	if msg.GsmMessage != nil {
		return errors.New("GSM Message should include in GMM Message")
	}

	if ue.State[atype] == nil {
		return fmt.Errorf("UE State is empty (accessType=%q). Can't send GSM Message", atype)
	}

	return n.sm.SendEvent(ue.State[atype], GmmMessageEvent, fsm.ArgsType{
		ArgAmfUe:         ue,
		ArgAccessType:    atype,
		ArgNASMessage:    msg.GmmMessage,
		ArgProcedureCode: code,
	})
}

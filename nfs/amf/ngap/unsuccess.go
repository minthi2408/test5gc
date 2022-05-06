package ngap

import (
//	"encoding/hex"
//	"strconv"
	"etri5gc/nfs/amf/context"
//	"etri5gc/nfs/amf/nas"

//	"github.com/free5gc/aper"
//	"github.com/free5gc/nas/nasMessage"
//	libngap "github.com/free5gc/ngap"
//	"github.com/free5gc/ngap/ngapConvert"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
)

func (h *ngapHandler) handleInitialContextSetupFailure(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var pDUSessionResourceFailedToSetupList *ngapType.PDUSessionResourceFailedToSetupListCxtFail
	var cause *ngapType.Cause
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	unsuccessfulOutcome := message.UnsuccessfulOutcome
	if unsuccessfulOutcome == nil {
		log.Error("UnsuccessfulOutcome is nil")
		return
	}
	initialContextSetupFailure := unsuccessfulOutcome.Value.InitialContextSetupFailure
	if initialContextSetupFailure == nil {
		log.Error("InitialContextSetupFailure is nil")
		return
	}

	for _, ie := range initialContextSetupFailure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			aMFUENGAPID = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if aMFUENGAPID == nil {
				log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			rANUENGAPID = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if rANUENGAPID == nil {
				log.Warn("RanUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail:
			pDUSessionResourceFailedToSetupList = ie.Value.PDUSessionResourceFailedToSetupListCxtFail
			log.Trace("Decode IE PDUSessionResourceFailedToSetupList")
			if pDUSessionResourceFailedToSetupList == nil {
				log.Warn("PDUSessionResourceFailedToSetupList is nil")
			}
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Warn("Cause is nil")
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE Criticality Diagnostics")
			if criticalityDiagnostics == nil {
				log.Warn("CriticalityDiagnostics is nil")
			}
		}
	}

	printAndGetCause(ran, cause)

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
	ranUe := ran.RanUeFindByRanUeNgapID(rANUENGAPID.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", rANUENGAPID.Value)
		return
	}

	log.Info("Handle Initial Context Setup Failure")

	amfUe := ranUe.AmfUe()
	if amfUe == nil {
		log.Error("amfUe is nil")
		return
	}

	if pDUSessionResourceFailedToSetupList != nil {
		log.Trace("Send PDUSessionResourceSetupUnsuccessfulTransfer to SMF")

		for _, item := range pDUSessionResourceFailedToSetupList.List {
			pduSessionID := int32(item.PDUSessionID.Value)
			transfer := item.PDUSessionResourceSetupUnsuccessfulTransfer
			smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
			if !ok {
				log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
			}
			_, _, _, err := h.backend.Consumer().Smf().SendUpdateSmContextN2Info(amfUe, smContext,
				models.N2SmInfoType_PDU_RES_SETUP_FAIL, transfer)
			if err != nil {
				log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceSetupUnsuccessfulTransfer] Error: %+v", err)
			}

			// if response != nil && response.BinaryDataN2SmInformation != nil {
			// TODO: n2SmInfo send to RAN
			// } else if response == nil {
			// TODO: error handling
			// }
		}
	}
}


func (h *ngapHandler) handleUEContextModificationFailure(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var cause *ngapType.Cause
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	var ranUe *context.RanUe

	unsuccessfulOutcome := message.UnsuccessfulOutcome
	if unsuccessfulOutcome == nil {
		log.Error("UnsuccessfulOutcome is nil")
		return
	}
	uEContextModificationFailure := unsuccessfulOutcome.Value.UEContextModificationFailure
	if uEContextModificationFailure == nil {
		log.Error("UEContextModificationFailure is nil")
		return
	}

	for _, ie := range uEContextModificationFailure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			aMFUENGAPID = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if aMFUENGAPID == nil {
				log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			rANUENGAPID = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if rANUENGAPID == nil {
				log.Warn("RanUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDCause: // ignore
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Warn("Cause is nil")
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // optional, ignore
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	if rANUENGAPID != nil {
		ranUe = ran.RanUeFindByRanUeNgapID(rANUENGAPID.Value)
		if ranUe == nil {
			log.Warnf("No UE Context[RanUeNgapID: %d]", rANUENGAPID.Value)
		}
	}

	if aMFUENGAPID != nil {
		ranUe = h.backend.Context().RanUeFindByAmfUeNgapID(aMFUENGAPID.Value)
		if ranUe == nil {
			log.Warnf("No UE Context[AmfUeNgapID: %d]", aMFUENGAPID.Value)
		}
	}

	if ranUe != nil {
		log.Tracef("AmfUeNgapID[%d] RanUeNgapID[%d]", ranUe.AmfUeNgapId(), ranUe.RanUeNgapId())
		log.Info("Handle UE Context Modification Failure")
	}

	if cause != nil {
		printAndGetCause(ran, cause)
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}

func (h *ngapHandler) handleHandoverFailure(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var cause *ngapType.Cause
	var targetUe *context.RanUe
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	unsuccessfulOutcome := message.UnsuccessfulOutcome // reject
	if unsuccessfulOutcome == nil {
		log.Error("Unsuccessful Message is nil")
		return
	}

	handoverFailure := unsuccessfulOutcome.Value.HandoverFailure
	if handoverFailure == nil {
		log.Error("HandoverFailure is nil")
		return
	}

	for _, ie := range handoverFailure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			aMFUENGAPID = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDCause: // ignore
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // ignore
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	causePresent := ngapType.CausePresentRadioNetwork
	causeValue := ngapType.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem
	if cause != nil {
		causePresent, causeValue = printAndGetCause(ran, cause)
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}

	targetUe = h.backend.Context().RanUeFindByAmfUeNgapID(aMFUENGAPID.Value)

	if targetUe == nil {
		log.Errorf("No UE Context[AmfUENGAPID: %d]", aMFUENGAPID.Value)
		cause := ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		h.sender.SendErrorIndication(ran, aMFUENGAPID, nil, &cause, nil)
		return
	}

	log.Info("Handle Handover Failure")

	sourceUe := targetUe.GetHandoverInfo().SourceUe
	if sourceUe == nil {
		// TODO: handle N2 Handover between AMF
		log.Error("N2 Handover between AMF has not been implemented yet")
	} else {
		amfUe := targetUe.AmfUe()
		if amfUe != nil {
			amfUe.SmContextList.Range(func(key, value interface{}) bool {
				pduSessionID := key.(int32)
				smContext := value.(*context.SmContext)
				causeAll := context.CauseAll{
					NgapCause: &models.NgApCause{
						Group: int32(causePresent),
						Value: int32(causeValue),
					},
				}
				_, _, _, err := h.backend.Consumer().Smf().SendUpdateSmContextN2HandoverCanceled(amfUe, smContext, causeAll)
				if err != nil {
					log.Errorf("Send UpdateSmContextN2HandoverCanceled Error for PduSessionId[%d]", pduSessionID)
				}
				return true
			})
		}
		h.sender.SendHandoverPreparationFailure(sourceUe, *cause, criticalityDiagnostics)
	}

	h.sender.SendUEContextReleaseCommand(targetUe, context.UeContextReleaseHandover, causePresent, causeValue)
}

func (h *ngapHandler) handleAMFconfigurationUpdateFailure(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var cause *ngapType.Cause
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	unsuccessfulOutcome := message.UnsuccessfulOutcome
	if unsuccessfulOutcome == nil {
		log.Error("Unsuccessful Message is nil")
		return
	}

	AMFconfigurationUpdateFailure := unsuccessfulOutcome.Value.AMFConfigurationUpdateFailure
	if AMFconfigurationUpdateFailure == nil {
		log.Error("AMFConfigurationUpdateFailure is nil")
		return
	}

	log.Info("Handle AMF Confioguration Update Failure")

	for _, ie := range AMFconfigurationUpdateFailure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Error("Cause is nil")
				return
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	//	TODO: Time To Wait

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}



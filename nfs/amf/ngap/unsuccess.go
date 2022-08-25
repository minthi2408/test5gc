package ngap

import (
	"etri5gc/nfs/amf/context"
	"etri5gc/openapi/models"

	"github.com/free5gc/ngap/ngapType"
)

func (h *ngapHandler) handleInitialContextSetupFailure(ran *context.AmfRan, initialContextSetupFailure *ngapType.InitialContextSetupFailure) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var pDUSessionResourceFailedToSetupList *ngapType.PDUSessionResourceFailedToSetupListCxtFail
	var cause *ngapType.Cause
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	for _, ie := range initialContextSetupFailure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
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
	ranUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
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
			_, _, _, err := smContext.SmfClient().SendUpdateSmContextN2Info(models.N2SmInfoType_PDU_RES_SETUP_FAIL, transfer)
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

func (h *ngapHandler) handleUEContextModificationFailure(ran *context.AmfRan, uEContextModificationFailure *ngapType.UEContextModificationFailure) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var cause *ngapType.Cause
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	var ranUe *context.RanUe

	for _, ie := range uEContextModificationFailure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
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

	if ran_ngapid != nil {
		ranUe = ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
		if ranUe == nil {
			log.Warnf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		}
	}

	if amf_ngapid != nil {
		ranUe = h.backend.Context().RanUeFindByAmfUeNgapID(amf_ngapid.Value)
		if ranUe == nil {
			log.Warnf("No UE Context[AmfUeNgapID: %d]", amf_ngapid.Value)
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

func (h *ngapHandler) handleHandoverFailure(ran *context.AmfRan, handoverFailure *ngapType.HandoverFailure) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var cause *ngapType.Cause
	var targetUe *context.RanUe
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	for _, ie := range handoverFailure.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			amf_ngapid = ie.Value.AMFUENGAPID
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

	targetUe = h.backend.Context().RanUeFindByAmfUeNgapID(amf_ngapid.Value)

	if targetUe == nil {
		log.Errorf("No UE Context[AmfUENGAPID: %d]", amf_ngapid.Value)
		cause := ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		h.sender.SendErrorIndication(ran, amf_ngapid, nil, &cause, nil)
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
				_, _, _, err := smContext.SmfClient().SendUpdateSmContextN2HandoverCanceled(causeAll)
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

func (h *ngapHandler) handleAMFconfigurationUpdateFailure(ran *context.AmfRan, AMFconfigurationUpdateFailure *ngapType.AMFConfigurationUpdateFailure) {
	var cause *ngapType.Cause
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

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

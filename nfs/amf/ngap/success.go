package ngap

import (
//	"encoding/hex"
//	"strconv"
	"etri5gc/nfs/amf/context"
//	"etri5gc/nfs/amf/nas"

//	"github.com/free5gc/aper"
//	"github.com/free5gc/nas/nasMessage"
//	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapConvert"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
)

func (h *ngapHandler) HandleUERadioCapabilityCheckResponse(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var iMSVoiceSupportIndicator *ngapType.IMSVoiceSupportIndicator
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics
	var ranUe *context.RanUe

	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
	//	ran.Log.Error("SuccessfulOutcome is nil")
		return
	}

	uERadioCapabilityCheckResponse := successfulOutcome.Value.UERadioCapabilityCheckResponse
	if uERadioCapabilityCheckResponse == nil {
	//	ran.Log.Error("UERadioCapabilityCheckResponse is nil")
		return
	}

	for i := 0; i < len(uERadioCapabilityCheckResponse.ProtocolIEs.List); i++ {
		ie := uERadioCapabilityCheckResponse.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			aMFUENGAPID = ie.Value.AMFUENGAPID
	//		ran.Log.Trace("Decode IE AmfUeNgapID")
			if aMFUENGAPID == nil {
	//			ran.Log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			rANUENGAPID = ie.Value.RANUENGAPID
	//		ran.Log.Trace("Decode IE RanUeNgapID")
			if rANUENGAPID == nil {
	//			ran.Log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDIMSVoiceSupportIndicator:
			iMSVoiceSupportIndicator = ie.Value.IMSVoiceSupportIndicator
	//		ran.Log.Trace("Decode IE IMSVoiceSupportIndicator")
			if iMSVoiceSupportIndicator == nil {
	//			ran.Log.Error("iMSVoiceSupportIndicator is nil")
				return
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
	//		ran.Log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	ranUe = ran.RanUeFindByRanUeNgapID(rANUENGAPID.Value)
	if ranUe == nil {
	//	ran.Log.Errorf("No UE Context[RanUeNgapID: %d]", rANUENGAPID.Value)
		return
	}

	//ranUe.Log.Info("Handle UE Radio Capability Check Response")

	// TODO: handle iMSVoiceSupportIndicator

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}


func (h *ngapHandler) HandleNGResetAcknowledge(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var uEAssociatedLogicalNGConnectionList *ngapType.UEAssociatedLogicalNGConnectionList
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
	//	ran.Log.Error("SuccessfulOutcome is nil")
		return
	}
	nGResetAcknowledge := successfulOutcome.Value.NGResetAcknowledge
	if nGResetAcknowledge == nil {
	//	ran.Log.Error("NGResetAcknowledge is nil")
		return
	}

	//ran.Log.Info("Handle NG Reset Acknowledge")

	for _, ie := range nGResetAcknowledge.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDUEAssociatedLogicalNGConnectionList:
			uEAssociatedLogicalNGConnectionList = ie.Value.UEAssociatedLogicalNGConnectionList
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
		}
	}

	if uEAssociatedLogicalNGConnectionList != nil {
	//	ran.Log.Tracef("%d UE association(s) has been reset", len(uEAssociatedLogicalNGConnectionList.List))
		for _, item := range uEAssociatedLogicalNGConnectionList.List {
			if item.AMFUENGAPID != nil && item.RANUENGAPID != nil {
	//			ran.Log.Tracef("%d: AmfUeNgapID[%d] RanUeNgapID[%d]", i+1, item.AMFUENGAPID.Value, item.RANUENGAPID.Value)
			} else if item.AMFUENGAPID != nil {
	//			ran.Log.Tracef("%d: AmfUeNgapID[%d] RanUeNgapID[-1]", i+1, item.AMFUENGAPID.Value)
			} else if item.RANUENGAPID != nil {
	//			ran.Log.Tracef("%d: AmfUeNgapID[-1] RanUeNgapID[%d]", i+1, item.RANUENGAPID.Value)
			}
		}
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}

func (h *ngapHandler) HandleUEContextReleaseComplete(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var userLocationInformation *ngapType.UserLocationInformation
	var infoOnRecommendedCellsAndRANNodesForPaging *ngapType.InfoOnRecommendedCellsAndRANNodesForPaging
	var pDUSessionResourceList *ngapType.PDUSessionResourceListCxtRelCpl
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
//		ran.Log.Error("SuccessfulOutcome is nil")
		return
	}
	uEContextReleaseComplete := successfulOutcome.Value.UEContextReleaseComplete
	if uEContextReleaseComplete == nil {
//		ran.Log.Error("NGResetAcknowledge is nil")
		return
	}

	for _, ie := range uEContextReleaseComplete.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			aMFUENGAPID = ie.Value.AMFUENGAPID
//			ran.Log.Trace("Decode IE AmfUeNgapID")
			if aMFUENGAPID == nil {
//				ran.Log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			rANUENGAPID = ie.Value.RANUENGAPID
//			ran.Log.Trace("Decode IE RanUeNgapID")
			if rANUENGAPID == nil {
//				ran.Log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDUserLocationInformation:
			userLocationInformation = ie.Value.UserLocationInformation
//			ran.Log.Trace("Decode IE UserLocationInformation")
		case ngapType.ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging:
			infoOnRecommendedCellsAndRANNodesForPaging = ie.Value.InfoOnRecommendedCellsAndRANNodesForPaging
//			ran.Log.Trace("Decode IE InfoOnRecommendedCellsAndRANNodesForPaging")
			if infoOnRecommendedCellsAndRANNodesForPaging != nil {
//				ran.Log.Warn("IE infoOnRecommendedCellsAndRANNodesForPaging is not support")
			}
		case ngapType.ProtocolIEIDPDUSessionResourceListCxtRelCpl:
			pDUSessionResourceList = ie.Value.PDUSessionResourceListCxtRelCpl
//			ran.Log.Trace("Decode IE PDUSessionResourceList")
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
//			ran.Log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	ranUe := h.backend.Context().RanUeFindByAmfUeNgapID(aMFUENGAPID.Value)
	if ranUe == nil {
//		ran.Log.Errorf("No RanUe Context[AmfUeNgapID: %d]", aMFUENGAPID.Value)
		cause := ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		h.sender.SendErrorIndication(ran, aMFUENGAPID, rANUENGAPID, &cause, nil)
		return
	}

//	ranUe.Log.Info("Handle UE Context Release Complete")

	if userLocationInformation != nil {
		ranUe.UpdateLocation(userLocationInformation)
	}
	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}

	amfUe := ranUe.AmfUe
	if amfUe == nil {
//		ran.Log.Infof("Release UE Context : RanUe[AmfUeNgapId: %d]", ranUe.AmfUeNgapId)
		err := ranUe.Remove()
		if err != nil {
//			ran.Log.Errorln(err.Error())
		}
		return
	}
	// TODO: AMF shall, if supported, store it and may use it for subsequent paging
	if infoOnRecommendedCellsAndRANNodesForPaging != nil {
		amfUe.InfoOnRecommendedCellsAndRanNodesForPaging = new(context.InfoOnRecommendedCellsAndRanNodesForPaging)

		recommendedCells := &amfUe.InfoOnRecommendedCellsAndRanNodesForPaging.RecommendedCells
		for _, item := range infoOnRecommendedCellsAndRANNodesForPaging.RecommendedCellsForPaging.RecommendedCellList.List {
			recommendedCell := context.RecommendedCell{}

			switch item.NGRANCGI.Present {
			case ngapType.NGRANCGIPresentNRCGI:
				recommendedCell.NgRanCGI.Present = context.NgRanCgiPresentNRCGI
				recommendedCell.NgRanCGI.NRCGI = new(models.Ncgi)
				plmnID := ngapConvert.PlmnIdToModels(item.NGRANCGI.NRCGI.PLMNIdentity)
				recommendedCell.NgRanCGI.NRCGI.PlmnId = &plmnID
				recommendedCell.NgRanCGI.NRCGI.NrCellId = ngapConvert.BitStringToHex(&item.NGRANCGI.NRCGI.NRCellIdentity.Value)
			case ngapType.NGRANCGIPresentEUTRACGI:
				recommendedCell.NgRanCGI.Present = context.NgRanCgiPresentEUTRACGI
				recommendedCell.NgRanCGI.EUTRACGI = new(models.Ecgi)
				plmnID := ngapConvert.PlmnIdToModels(item.NGRANCGI.EUTRACGI.PLMNIdentity)
				recommendedCell.NgRanCGI.EUTRACGI.PlmnId = &plmnID
				recommendedCell.NgRanCGI.EUTRACGI.EutraCellId = ngapConvert.BitStringToHex(
					&item.NGRANCGI.EUTRACGI.EUTRACellIdentity.Value)
			}

			if item.TimeStayedInCell != nil {
				recommendedCell.TimeStayedInCell = new(int64)
				*recommendedCell.TimeStayedInCell = *item.TimeStayedInCell
			}

			*recommendedCells = append(*recommendedCells, recommendedCell)
		}

		recommendedRanNodes := &amfUe.InfoOnRecommendedCellsAndRanNodesForPaging.RecommendedRanNodes
		ranNodeList := infoOnRecommendedCellsAndRANNodesForPaging.RecommendRANNodesForPaging.RecommendedRANNodeList.List
		for _, item := range ranNodeList {
			recommendedRanNode := context.RecommendRanNode{}

			switch item.AMFPagingTarget.Present {
			case ngapType.AMFPagingTargetPresentGlobalRANNodeID:
				recommendedRanNode.Present = context.RecommendRanNodePresentRanNode
				recommendedRanNode.GlobalRanNodeId = new(models.GlobalRanNodeId)
				// TODO: recommendedRanNode.GlobalRanNodeId = ngapConvert.RanIdToModels(item.AMFPagingTarget.GlobalRANNodeID)
			case ngapType.AMFPagingTargetPresentTAI:
				recommendedRanNode.Present = context.RecommendRanNodePresentTAI
				tai := ngapConvert.TaiToModels(*item.AMFPagingTarget.TAI)
				recommendedRanNode.Tai = &tai
			}
			*recommendedRanNodes = append(*recommendedRanNodes, recommendedRanNode)
		}
	}

	// for each pduSessionID invoke Nsmf_PDUSession_UpdateSMContext Request
	var cause context.CauseAll
	if tmp, exist := amfUe.ReleaseCause[ran.AnType()]; exist {
		cause = *tmp
	}
	if amfUe.State[ran.AnType()].Is(context.Registered) {
//		ranUe.Log.Info("Rel Ue Context in GMM-Registered")
		if cause.NgapCause != nil && pDUSessionResourceList != nil {
			for _, pduSessionReourceItem := range pDUSessionResourceList.List {
				pduSessionID := int32(pduSessionReourceItem.PDUSessionID.Value)
				smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
				if !ok {
//					ranUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
					// TODO: Check if doing error handling here
					continue
				}
				response, _, _, err := h.backend.Consumer().Smf().SendUpdateSmContextDeactivateUpCnxState(amfUe, smContext, cause)
				if err != nil {
//					ran.Log.Errorf("Send Update SmContextDeactivate UpCnxState Error[%s]", err.Error())
				} else if response == nil {
//					ran.Log.Errorln("Send Update SmContextDeactivate UpCnxState Error")
				}
			}
		}
	}

	// Remove UE N2 Connection
	delete(amfUe.ReleaseCause, ran.AnType())
	switch ranUe.ReleaseAction {
	case context.UeContextN2NormalRelease:
//		ranUe.Log.Infof("Release UE[%s] Context : N2 Connection Release", amfUe.Supi)
		// amfUe.DetachRanUe(ran.AnType)
		err := ranUe.Remove()
		if err != nil {
	//		ran.Log.Errorln(err.Error())
		}
	case context.UeContextReleaseUeContext:
	//	ranUe.Log.Infof("Release UE[%s] Context : Release Ue Context", amfUe.Supi)
	//TODO:tungtq
	//	gmm_common.RemoveAmfUe(amfUe)
	case context.UeContextReleaseHandover:
//		ranUe.Log.Infof("Release UE[%s] Context : Release for Handover", amfUe.Supi)
		// TODO: it's a workaround, need to fix it.
		targetRanUe := h.backend.Context().RanUeFindByAmfUeNgapID(ranUe.TargetUe.AmfUeNgapId)

		context.DetachSourceUeTargetUe(ranUe)
		err := ranUe.Remove()
		if err != nil {
//			ran.Log.Errorln(err.Error())
		}
		amfUe.AttachRanUe(targetRanUe)
		// Todo: remove indirect tunnel
	default:
//		ran.Log.Errorf("Invalid Release Action[%d]", ranUe.ReleaseAction)
	}
}

func (h *ngapHandler) HandlePDUSessionResourceReleaseResponse(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var pDUSessionResourceReleasedList *ngapType.PDUSessionResourceReleasedListRelRes
	var userLocationInformation *ngapType.UserLocationInformation
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
//		ran.Log.Error("SuccessfulOutcome is nil")
		return
	}
	pDUSessionResourceReleaseResponse := successfulOutcome.Value.PDUSessionResourceReleaseResponse
	if pDUSessionResourceReleaseResponse == nil {
//		ran.Log.Error("PDUSessionResourceReleaseResponse is nil")
		return
	}

	for _, ie := range pDUSessionResourceReleaseResponse.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			aMFUENGAPID = ie.Value.AMFUENGAPID
//			ran.Log.Trace("Decode IE AmfUeNgapID")
			if aMFUENGAPID == nil {
//				ran.Log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			rANUENGAPID = ie.Value.RANUENGAPID
//			ran.Log.Trace("Decode IE RanUENgapID")
			if rANUENGAPID == nil {
//				ran.Log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDPDUSessionResourceReleasedListRelRes:
			pDUSessionResourceReleasedList = ie.Value.PDUSessionResourceReleasedListRelRes
//			ran.Log.Trace("Decode IE PDUSessionResourceReleasedList")
			if pDUSessionResourceReleasedList == nil {
//				ran.Log.Error("PDUSessionResourceReleasedList is nil")
				return
			}
		case ngapType.ProtocolIEIDUserLocationInformation:
			userLocationInformation = ie.Value.UserLocationInformation
//			ran.Log.Trace("Decode IE UserLocationInformation")
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
//			ran.Log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	ranUe := ran.RanUeFindByRanUeNgapID(rANUENGAPID.Value)
	if ranUe == nil {
//		ran.Log.Errorf("No UE Context[RanUeNgapID: %d]", rANUENGAPID.Value)
		return
	}

//	ranUe.Log.Info("Handle PDU Session Resource Release Response")

	if userLocationInformation != nil {
		ranUe.UpdateLocation(userLocationInformation)
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}

	amfUe := ranUe.AmfUe
	if amfUe == nil {
//		ranUe.Log.Error("amfUe is nil")
		return
	}
	if pDUSessionResourceReleasedList != nil {
//		ranUe.Log.Trace("Send PDUSessionResourceReleaseResponseTransfer to SMF")

		for _, item := range pDUSessionResourceReleasedList.List {
			pduSessionID := int32(item.PDUSessionID.Value)
			transfer := item.PDUSessionResourceReleaseResponseTransfer
			smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
			if !ok {
//				ranUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
				continue
			}
			_, responseErr, problemDetail, err := h.backend.Consumer().Smf().SendUpdateSmContextN2Info(amfUe, smContext,
				models.N2SmInfoType_PDU_RES_REL_RSP, transfer)
			// TODO: error handling
			if err != nil {
//				ranUe.Log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceReleaseResponse] Error: %+v", err)
			} else if responseErr != nil && responseErr.JsonData.Error != nil {
//				ranUe.Log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceReleaseResponse] Error: %+v",
//					responseErr.JsonData.Error.Cause)
			} else if problemDetail != nil {
//				ranUe.Log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceReleaseResponse] Failed: %+v", problemDetail)
			}
		}
	}
}


func (h *ngapHandler) HandlePDUSessionResourceSetupResponse(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var pDUSessionResourceSetupResponseList *ngapType.PDUSessionResourceSetupListSURes
	var pDUSessionResourceFailedToSetupList *ngapType.PDUSessionResourceFailedToSetupListSURes
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	var ranUe *context.RanUe

	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
//		ran.Log.Error("SuccessfulOutcome is nil")
		return
	}
	pDUSessionResourceSetupResponse := successfulOutcome.Value.PDUSessionResourceSetupResponse
	if pDUSessionResourceSetupResponse == nil {
//		ran.Log.Error("PDUSessionResourceSetupResponse is nil")
		return
	}

	for _, ie := range pDUSessionResourceSetupResponse.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			aMFUENGAPID = ie.Value.AMFUENGAPID
//			ran.Log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			rANUENGAPID = ie.Value.RANUENGAPID
//			ran.Log.Trace("Decode IE RanUeNgapID")
		case ngapType.ProtocolIEIDPDUSessionResourceSetupListSURes: // ignore
			pDUSessionResourceSetupResponseList = ie.Value.PDUSessionResourceSetupListSURes
//			ran.Log.Trace("Decode IE PDUSessionResourceSetupListSURes")
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListSURes: // ignore
			pDUSessionResourceFailedToSetupList = ie.Value.PDUSessionResourceFailedToSetupListSURes
//			ran.Log.Trace("Decode IE PDUSessionResourceFailedToSetupListSURes")
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // optional, ignore
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
//			ran.Log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	if rANUENGAPID != nil {
		ranUe = ran.RanUeFindByRanUeNgapID(rANUENGAPID.Value)
		if ranUe == nil {
//			ran.Log.Warnf("No UE Context[RanUeNgapID: %d]", rANUENGAPID.Value)
		}
	}

	if aMFUENGAPID != nil {
		ranUe = h.backend.Context().RanUeFindByAmfUeNgapID(aMFUENGAPID.Value)
		if ranUe == nil {
//			ran.Log.Warnf("No UE Context[AmfUeNgapID: %d]", aMFUENGAPID.Value)
			return
		}
	}

	if ranUe != nil {
//		ranUe.Log.Info("Handle PDU Session Resource Setup Response")
//		ranUe.Log.Tracef("AmfUeNgapID[%d] RanUeNgapID[%d]", ranUe.AmfUeNgapId, ranUe.RanUeNgapId)
		amfUe := ranUe.AmfUe
		if amfUe == nil {
//			ranUe.Log.Error("amfUe is nil")
			return
		}

		if pDUSessionResourceSetupResponseList != nil {
//			ranUe.Log.Trace("Send PDUSessionResourceSetupResponseTransfer to SMF")

			for _, item := range pDUSessionResourceSetupResponseList.List {
				pduSessionID := int32(item.PDUSessionID.Value)
				transfer := item.PDUSessionResourceSetupResponseTransfer
				smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
				if !ok {
//					ranUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
				}
				_, _, _, err := h.backend.Consumer().Smf().SendUpdateSmContextN2Info(amfUe, smContext,
					models.N2SmInfoType_PDU_RES_SETUP_RSP, transfer)
				if err != nil {
					//ranUe.Log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceSetupResponseTransfer] Error: %+v", err)
				}
				// RAN initiated QoS Flow Mobility in subclause 5.2.2.3.7
				// if response != nil && response.BinaryDataN2SmInformation != nil {
				// TODO: n2SmInfo send to RAN
				// } else if response == nil {
				// TODO: error handling
				// }
			}
		}

		if pDUSessionResourceFailedToSetupList != nil {
			//ranUe.Log.Trace("Send PDUSessionResourceSetupUnsuccessfulTransfer to SMF")

			for _, item := range pDUSessionResourceFailedToSetupList.List {
				pduSessionID := int32(item.PDUSessionID.Value)
				transfer := item.PDUSessionResourceSetupUnsuccessfulTransfer
				smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
				if !ok {
			//		ranUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
				}
				_, _, _, err := h.backend.Consumer().Smf().SendUpdateSmContextN2Info(amfUe, smContext,
					models.N2SmInfoType_PDU_RES_SETUP_FAIL, transfer)
				if err != nil {
					//ranUe.Log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceSetupUnsuccessfulTransfer] Error: %+v", err)
				}

				// if response != nil && response.BinaryDataN2SmInformation != nil {
				// TODO: n2SmInfo send to RAN
				// } else if response == nil {
				// TODO: error handling
				// }
			}
		}
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}

func (h *ngapHandler) HandlePDUSessionResourceModifyResponse(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var pduSessionResourceModifyResponseList *ngapType.PDUSessionResourceModifyListModRes
	var pduSessionResourceFailedToModifyList *ngapType.PDUSessionResourceFailedToModifyListModRes
	var userLocationInformation *ngapType.UserLocationInformation
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	var ranUe *context.RanUe

	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
//		ran.Log.Error("SuccessfulOutcome is nil")
		return
	}
	pDUSessionResourceModifyResponse := successfulOutcome.Value.PDUSessionResourceModifyResponse
	if pDUSessionResourceModifyResponse == nil {
//		ran.Log.Error("PDUSessionResourceModifyResponse is nil")
		return
	}

	for _, ie := range pDUSessionResourceModifyResponse.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			aMFUENGAPID = ie.Value.AMFUENGAPID
//			ran.Log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			rANUENGAPID = ie.Value.RANUENGAPID
//			ran.Log.Trace("Decode IE RanUeNgapID")
		case ngapType.ProtocolIEIDPDUSessionResourceModifyListModRes: // ignore
			pduSessionResourceModifyResponseList = ie.Value.PDUSessionResourceModifyListModRes
//			ran.Log.Trace("Decode IE PDUSessionResourceModifyListModRes")
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToModifyListModRes: // ignore
			pduSessionResourceFailedToModifyList = ie.Value.PDUSessionResourceFailedToModifyListModRes
//			ran.Log.Trace("Decode IE PDUSessionResourceFailedToModifyListModRes")
		case ngapType.ProtocolIEIDUserLocationInformation: // optional, ignore
			userLocationInformation = ie.Value.UserLocationInformation
//			ran.Log.Trace("Decode IE UserLocationInformation")
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // optional, ignore
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
//			ran.Log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	if rANUENGAPID != nil {
		ranUe = ran.RanUeFindByRanUeNgapID(rANUENGAPID.Value)
		if ranUe == nil {
//			ran.Log.Warnf("No UE Context[RanUeNgapID: %d]", rANUENGAPID.Value)
		}
	}

	if aMFUENGAPID != nil {
		ranUe = h.backend.Context().RanUeFindByAmfUeNgapID(aMFUENGAPID.Value)
		if ranUe == nil {
//			ran.Log.Warnf("No UE Context[AmfUeNgapID: %d]", aMFUENGAPID.Value)
			return
		}
	}

	if ranUe != nil {
//		ranUe.Log.Info("Handle PDU Session Resource Modify Response")
//		ranUe.Log.Tracef("AmfUeNgapID[%d] RanUeNgapID[%d]", ranUe.AmfUeNgapId, ranUe.RanUeNgapId)
		amfUe := ranUe.AmfUe
		if amfUe == nil {
//			ranUe.Log.Error("amfUe is nil")
			return
		}

		if pduSessionResourceModifyResponseList != nil {
//			ranUe.Log.Trace("Send PDUSessionResourceModifyResponseTransfer to SMF")

			for _, item := range pduSessionResourceModifyResponseList.List {
				pduSessionID := int32(item.PDUSessionID.Value)
				transfer := item.PDUSessionResourceModifyResponseTransfer
				smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
				if !ok {
//					ranUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
				}
				_, _, _, err := h.backend.Consumer().Smf().SendUpdateSmContextN2Info(amfUe, smContext,
					models.N2SmInfoType_PDU_RES_MOD_RSP, transfer)
				if err != nil {
//					ranUe.Log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceModifyResponseTransfer] Error: %+v", err)
				}
				// if response != nil && response.BinaryDataN2SmInformation != nil {
				// TODO: n2SmInfo send to RAN
				// } else if response == nil {
				// TODO: error handling
				// }
			}
		}

		if pduSessionResourceFailedToModifyList != nil {
//			ranUe.Log.Trace("Send PDUSessionResourceModifyUnsuccessfulTransfer to SMF")

			for _, item := range pduSessionResourceFailedToModifyList.List {
				pduSessionID := int32(item.PDUSessionID.Value)
				transfer := item.PDUSessionResourceModifyUnsuccessfulTransfer
				smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
				if !ok {
//					ranUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
				}
				// response, _, _, err := consumer.SendUpdateSmContextN2Info(amfUe, pduSessionID,
				_, _, _, err := h.backend.Consumer().Smf().SendUpdateSmContextN2Info(amfUe, smContext,
					models.N2SmInfoType_PDU_RES_MOD_FAIL, transfer)
				if err != nil {
//					ranUe.Log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceModifyUnsuccessfulTransfer] Error: %+v", err)
				}
				// if response != nil && response.BinaryDataN2SmInformation != nil {
				// TODO: n2SmInfo send to RAN
				// } else if response == nil {
				// TODO: error handling
				// }
			}
		}

		if userLocationInformation != nil {
			ranUe.UpdateLocation(userLocationInformation)
		}
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}


func (h *ngapHandler) HandleInitialContextSetupResponse(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var pDUSessionResourceSetupResponseList *ngapType.PDUSessionResourceSetupListCxtRes
	var pDUSessionResourceFailedToSetupList *ngapType.PDUSessionResourceFailedToSetupListCxtRes
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
//		ran.Log.Error("SuccessfulOutcome is nil")
		return
	}
	initialContextSetupResponse := successfulOutcome.Value.InitialContextSetupResponse
	if initialContextSetupResponse == nil {
//		ran.Log.Error("InitialContextSetupResponse is nil")
		return
	}

	for _, ie := range initialContextSetupResponse.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			aMFUENGAPID = ie.Value.AMFUENGAPID
//			ran.Log.Trace("Decode IE AmfUeNgapID")
			if aMFUENGAPID == nil {
//				ran.Log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			rANUENGAPID = ie.Value.RANUENGAPID
//			ran.Log.Trace("Decode IE RanUeNgapID")
			if rANUENGAPID == nil {
//				ran.Log.Warn("RanUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtRes:
			pDUSessionResourceSetupResponseList = ie.Value.PDUSessionResourceSetupListCxtRes
//			ran.Log.Trace("Decode IE PDUSessionResourceSetupResponseList")
			if pDUSessionResourceSetupResponseList == nil {
//				ran.Log.Warn("PDUSessionResourceSetupResponseList is nil")
			}
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes:
			pDUSessionResourceFailedToSetupList = ie.Value.PDUSessionResourceFailedToSetupListCxtRes
//			ran.Log.Trace("Decode IE PDUSessionResourceFailedToSetupList")
			if pDUSessionResourceFailedToSetupList == nil {
//				ran.Log.Warn("PDUSessionResourceFailedToSetupList is nil")
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
//			ran.Log.Trace("Decode IE Criticality Diagnostics")
			if criticalityDiagnostics == nil {
//				ran.Log.Warn("Criticality Diagnostics is nil")
			}
		}
	}

	ranUe := ran.RanUeFindByRanUeNgapID(rANUENGAPID.Value)
	if ranUe == nil {
//		ran.Log.Errorf("No UE Context[RanUeNgapID: %d]", rANUENGAPID.Value)
		return
	}
	amfUe := ranUe.AmfUe
	if amfUe == nil {
//		ran.Log.Error("amfUe is nil")
		return
	}

//	ran.Log.Tracef("RanUeNgapID[%d] AmfUeNgapID[%d]", ranUe.RanUeNgapId, ranUe.AmfUeNgapId)
//	ranUe.Log.Info("Handle Initial Context Setup Response")

	if pDUSessionResourceSetupResponseList != nil {
//		ranUe.Log.Trace("Send PDUSessionResourceSetupResponseTransfer to SMF")

		for _, item := range pDUSessionResourceSetupResponseList.List {
			pduSessionID := int32(item.PDUSessionID.Value)
			transfer := item.PDUSessionResourceSetupResponseTransfer
			smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
			if !ok {
//				ranUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
			}
			// response, _, _, err := consumer.SendUpdateSmContextN2Info(amfUe, pduSessionID,
			_, _, _, err := h.backend.Consumer().Smf().SendUpdateSmContextN2Info(amfUe, smContext,
				models.N2SmInfoType_PDU_RES_SETUP_RSP, transfer)
			if err != nil {
//				ranUe.Log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceSetupResponseTransfer] Error: %+v", err)
			}
			// RAN initiated QoS Flow Mobility in subclause 5.2.2.3.7
			// if response != nil && response.BinaryDataN2SmInformation != nil {
			// TODO: n2SmInfo send to RAN
			// } else if response == nil {
			// TODO: error handling
			// }
		}
	}

	if pDUSessionResourceFailedToSetupList != nil {
//		ranUe.Log.Trace("Send PDUSessionResourceSetupUnsuccessfulTransfer to SMF")

		for _, item := range pDUSessionResourceFailedToSetupList.List {
			pduSessionID := int32(item.PDUSessionID.Value)
			transfer := item.PDUSessionResourceSetupUnsuccessfulTransfer
			smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
			if !ok {
//				ranUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
			}
			// response, _, _, err := consumer.SendUpdateSmContextN2Info(amfUe, pduSessionID,
			_, _, _, err := h.backend.Consumer().Smf().SendUpdateSmContextN2Info(amfUe, smContext,
				models.N2SmInfoType_PDU_RES_SETUP_FAIL, transfer)
			if err != nil {
//				ranUe.Log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceSetupUnsuccessfulTransfer] Error: %+v", err)
			}

			// if response != nil && response.BinaryDataN2SmInformation != nil {
			// TODO: n2SmInfo send to RAN
			// } else if response == nil {
			// TODO: error handling
			// }
		}
	}

	if ranUe.Ran.AnType() == models.AccessType_NON_3_GPP_ACCESS {
		h.sender.SendDownlinkNasTransport(ranUe, amfUe.RegistrationAcceptForNon3GPPAccess, nil)
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}

func (h *ngapHandler) HandleUEContextModificationResponse(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var rRCState *ngapType.RRCState
	var userLocationInformation *ngapType.UserLocationInformation
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	var ranUe *context.RanUe


	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
//		ran.Log.Error("SuccessfulOutcome is nil")
		return
	}
	uEContextModificationResponse := successfulOutcome.Value.UEContextModificationResponse
	if uEContextModificationResponse == nil {
//		ran.Log.Error("UEContextModificationResponse is nil")
		return
	}

	for _, ie := range uEContextModificationResponse.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			aMFUENGAPID = ie.Value.AMFUENGAPID
//			ran.Log.Trace("Decode IE AmfUeNgapID")
			if aMFUENGAPID == nil {
//				ran.Log.Warn("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			rANUENGAPID = ie.Value.RANUENGAPID
//			ran.Log.Trace("Decode IE RanUeNgapID")
			if rANUENGAPID == nil {
//				ran.Log.Warn("RanUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRRCState: // optional, ignore
			rRCState = ie.Value.RRCState
//			ran.Log.Trace("Decode IE RRCState")
		case ngapType.ProtocolIEIDUserLocationInformation: // optional, ignore
			userLocationInformation = ie.Value.UserLocationInformation
//			ran.Log.Trace("Decode IE UserLocationInformation")
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // optional, ignore
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
//			ran.Log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	if rANUENGAPID != nil {
		ranUe = ran.RanUeFindByRanUeNgapID(rANUENGAPID.Value)
		if ranUe == nil {
//			ran.Log.Warnf("No UE Context[RanUeNgapID: %d]", rANUENGAPID.Value)
		}
	}

	if aMFUENGAPID != nil {
		ranUe = h.backend.Context().RanUeFindByAmfUeNgapID(aMFUENGAPID.Value)
		if ranUe == nil {
//			ran.Log.Warnf("No UE Context[AmfUeNgapID: %d]", aMFUENGAPID.Value)
			return
		}
	}

	if ranUe != nil {
//		ranUe.Log.Tracef("AmfUeNgapID[%d] RanUeNgapID[%d]", ranUe.AmfUeNgapId, ranUe.RanUeNgapId)
//		ranUe.Log.Info("Handle UE Context Modification Response")

		if rRCState != nil {
			switch rRCState.Value {
			case ngapType.RRCStatePresentInactive:
//				ranUe.Log.Trace("UE RRC State: Inactive")
			case ngapType.RRCStatePresentConnected:
//				ranUe.Log.Trace("UE RRC State: Connected")
			}
		}

		if userLocationInformation != nil {
			ranUe.UpdateLocation(userLocationInformation)
		}
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}

func (h *ngapHandler) HandleHandoverRequestAcknowledge(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	var pDUSessionResourceAdmittedList *ngapType.PDUSessionResourceAdmittedList
	var pDUSessionResourceFailedToSetupListHOAck *ngapType.PDUSessionResourceFailedToSetupListHOAck
	var targetToSourceTransparentContainer *ngapType.TargetToSourceTransparentContainer
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
//		ran.Log.Error("SuccessfulOutcome is nil")
		return
	}
	handoverRequestAcknowledge := successfulOutcome.Value.HandoverRequestAcknowledge // reject
	if handoverRequestAcknowledge == nil {
//		ran.Log.Error("HandoverRequestAcknowledge is nil")
		return
	}

	for _, ie := range handoverRequestAcknowledge.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // ignore
			aMFUENGAPID = ie.Value.AMFUENGAPID
//			ran.Log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			rANUENGAPID = ie.Value.RANUENGAPID
//			ran.Log.Trace("Decode IE RanUeNgapID")
		case ngapType.ProtocolIEIDPDUSessionResourceAdmittedList: // ignore
			pDUSessionResourceAdmittedList = ie.Value.PDUSessionResourceAdmittedList
//			ran.Log.Trace("Decode IE PduSessionResourceAdmittedList")
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck: // ignore
			pDUSessionResourceFailedToSetupListHOAck = ie.Value.PDUSessionResourceFailedToSetupListHOAck
//			ran.Log.Trace("Decode IE PduSessionResourceFailedToSetupListHOAck")
		case ngapType.ProtocolIEIDTargetToSourceTransparentContainer: // reject
			targetToSourceTransparentContainer = ie.Value.TargetToSourceTransparentContainer
//			ran.Log.Trace("Decode IE TargetToSourceTransparentContainer")
		case ngapType.ProtocolIEIDCriticalityDiagnostics: // ignore
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
//			ran.Log.Trace("Decode IE CriticalityDiagnostics")
		}
	}
	if targetToSourceTransparentContainer == nil {
//		ran.Log.Error("TargetToSourceTransparentContainer is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
			ngapType.ProtocolIEIDTargetToSourceTransparentContainer, ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}
	if len(iesCriticalityDiagnostics.List) > 0 {
//		ran.Log.Error("Has missing reject IE(s)")

		procedureCode := ngapType.ProcedureCodeHandoverResourceAllocation
		triggeringMessage := ngapType.TriggeringMessagePresentSuccessfulOutcome
		procedureCriticality := ngapType.CriticalityPresentReject
		criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage,
			&procedureCriticality, &iesCriticalityDiagnostics)
		h.sender.SendErrorIndication(ran, aMFUENGAPID, rANUENGAPID, nil, &criticalityDiagnostics)
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}

	targetUe := h.backend.Context().RanUeFindByAmfUeNgapID(aMFUENGAPID.Value)
	if targetUe == nil {
//		ran.Log.Errorf("No UE Context[AMFUENGAPID: %d]", aMFUENGAPID.Value)
		return
	}

//	targetUe.Log.Info("Handle Handover Request Acknowledge")

	if rANUENGAPID != nil {
		targetUe.RanUeNgapId = rANUENGAPID.Value
//		targetUe.UpdateLogFields()
	}
//	ran.Log.Debugf("Target Ue RanUeNgapID[%d] AmfUeNgapID[%d]", targetUe.RanUeNgapId, targetUe.AmfUeNgapId)

	amfUe := targetUe.AmfUe
	if amfUe == nil {
//		targetUe.Log.Error("amfUe is nil")
		return
	}

	var pduSessionResourceHandoverList ngapType.PDUSessionResourceHandoverList
	var pduSessionResourceToReleaseList ngapType.PDUSessionResourceToReleaseListHOCmd

	// describe in 23.502 4.9.1.3.2 step11
	if pDUSessionResourceAdmittedList != nil {
		for _, item := range pDUSessionResourceAdmittedList.List {
			pduSessionID := item.PDUSessionID.Value
			transfer := item.HandoverRequestAcknowledgeTransfer
			pduSessionId := int32(pduSessionID)
			if smContext, exist := amfUe.SmContextFindByPDUSessionID(pduSessionId); exist {
				response, errResponse, problemDetails, err := h.backend.Consumer().Smf().SendUpdateSmContextN2HandoverPrepared(amfUe,
					smContext, models.N2SmInfoType_HANDOVER_REQ_ACK, transfer)
				if err != nil {
//					targetUe.Log.Errorf("Send HandoverRequestAcknowledgeTransfer error: %v", err)
				}
				if problemDetails != nil {
//					targetUe.Log.Warnf("ProblemDetails[status: %d, Cause: %s]", problemDetails.Status, problemDetails.Cause)
				}
				if response != nil && response.BinaryDataN2SmInformation != nil {
					handoverItem := ngapType.PDUSessionResourceHandoverItem{}
					handoverItem.PDUSessionID = item.PDUSessionID
					handoverItem.HandoverCommandTransfer = response.BinaryDataN2SmInformation
					pduSessionResourceHandoverList.List = append(pduSessionResourceHandoverList.List, handoverItem)
					targetUe.SuccessPduSessionId = append(targetUe.SuccessPduSessionId, pduSessionId)
				}
				if errResponse != nil && errResponse.BinaryDataN2SmInformation != nil {
					releaseItem := ngapType.PDUSessionResourceToReleaseItemHOCmd{}
					releaseItem.PDUSessionID = item.PDUSessionID
					releaseItem.HandoverPreparationUnsuccessfulTransfer = errResponse.BinaryDataN2SmInformation
					pduSessionResourceToReleaseList.List = append(pduSessionResourceToReleaseList.List, releaseItem)
				}
			}
		}
	}

	if pDUSessionResourceFailedToSetupListHOAck != nil {
		for _, item := range pDUSessionResourceFailedToSetupListHOAck.List {
			pduSessionID := item.PDUSessionID.Value
			transfer := item.HandoverResourceAllocationUnsuccessfulTransfer
			pduSessionId := int32(pduSessionID)
			if smContext, exist := amfUe.SmContextFindByPDUSessionID(pduSessionId); exist {
				_, _, problemDetails, err := h.backend.Consumer().Smf().SendUpdateSmContextN2HandoverPrepared(amfUe, smContext,
					models.N2SmInfoType_HANDOVER_RES_ALLOC_FAIL, transfer)
				if err != nil {
//					targetUe.Log.Errorf("Send HandoverResourceAllocationUnsuccessfulTransfer error: %v", err)
				}
				if problemDetails != nil {
//					targetUe.Log.Warnf("ProblemDetails[status: %d, Cause: %s]", problemDetails.Status, problemDetails.Cause)
				}
			}
		}
	}

	sourceUe := targetUe.SourceUe
	if sourceUe == nil {
		// TODO: Send Namf_Communication_CreateUEContext Response to S-AMF
//		ran.Log.Error("handover between different Ue has not been implement yet")
	} else {
//		ran.Log.Tracef("Source: RanUeNgapID[%d] AmfUeNgapID[%d]", sourceUe.RanUeNgapId, sourceUe.AmfUeNgapId)
//		ran.Log.Tracef("Target: RanUeNgapID[%d] AmfUeNgapID[%d]", targetUe.RanUeNgapId, targetUe.AmfUeNgapId)
		if len(pduSessionResourceHandoverList.List) == 0 {
//			targetUe.Log.Info("Handle Handover Preparation Failure [HoFailure In Target5GC NgranNode Or TargetSystem]")
			cause := &ngapType.Cause{
				Present: ngapType.CausePresentRadioNetwork,
				RadioNetwork: &ngapType.CauseRadioNetwork{
					Value: ngapType.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem,
				},
			}
			h.sender.SendHandoverPreparationFailure(sourceUe, *cause, nil)
			return
		}
		h.sender.SendHandoverCommand(sourceUe, pduSessionResourceHandoverList, pduSessionResourceToReleaseList,
			*targetToSourceTransparentContainer, nil)
	}
}

func (h *ngapHandler) HandleAMFconfigurationUpdateAcknowledge(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var aMFTNLAssociationSetupList *ngapType.AMFTNLAssociationSetupList
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics
	var aMFTNLAssociationFailedToSetupList *ngapType.TNLAssociationList

	successfulOutcome := message.SuccessfulOutcome
	if successfulOutcome == nil {
//		ran.Log.Error("SuccessfulOutcome is nil")
		return
	}
	aMFConfigurationUpdateAcknowledge := successfulOutcome.Value.AMFConfigurationUpdateAcknowledge
	if aMFConfigurationUpdateAcknowledge == nil {
//		ran.Log.Error("AMFConfigurationUpdateAcknowledge is nil")
		return
	}

//	ran.Log.Info("Handle AMF Configuration Update Acknowledge")

	for i := 0; i < len(aMFConfigurationUpdateAcknowledge.ProtocolIEs.List); i++ {
		ie := aMFConfigurationUpdateAcknowledge.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFTNLAssociationSetupList:
			aMFTNLAssociationSetupList = ie.Value.AMFTNLAssociationSetupList
//			ran.Log.Trace("Decode IE AMFTNLAssociationSetupList")
			if aMFTNLAssociationSetupList == nil {
//				ran.Log.Error("AMFTNLAssociationSetupList is nil")
				return
			}
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
//			ran.Log.Trace("Decode IE Criticality Diagnostics")

		case ngapType.ProtocolIEIDAMFTNLAssociationFailedToSetupList:
			aMFTNLAssociationFailedToSetupList = ie.Value.AMFTNLAssociationFailedToSetupList
//			ran.Log.Trace("Decode IE AMFTNLAssociationFailedToSetupList")
			if aMFTNLAssociationFailedToSetupList == nil {
//				ran.Log.Error("AMFTNLAssociationFailedToSetupList is nil")
				return
			}
		}
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}
}



package ngap

import (
	"encoding/hex"
	"etri5gc/nfs/amf/context"
	"strconv"

	"etri5gc/nfs/amf/ngap/util"
	"etri5gc/sbi/models"
	"etri5gc/sbi/utils/ngapConvert"

	//	"github.com/free5gc/aper"
	"github.com/free5gc/nas/nasMessage"
	//	libngap "github.com/free5gc/ngap"
	"github.com/free5gc/ngap/ngapType"
)

func (h *ngapHandler) handleNGSetupRequest(ran *context.AmfRan, nGSetupRequest *ngapType.NGSetupRequest) {
	var ranId *ngapType.GlobalRANNodeID
	var ranName *ngapType.RANNodeName
	var supportedTAList *ngapType.SupportedTAList
	var pagingDRX *ngapType.PagingDRX

	var cause ngapType.Cause

	log.Info("Handle NG Setup request")

	for i := 0; i < len(nGSetupRequest.ProtocolIEs.List); i++ {
		ie := nGSetupRequest.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDGlobalRANNodeID:
			ranId = ie.Value.GlobalRANNodeID
			log.Trace("Decode IE GlobalRANNodeID")
			if ranId == nil {
				log.Error("GlobalRANNodeID is nil")
				return
			}
		case ngapType.ProtocolIEIDSupportedTAList:
			supportedTAList = ie.Value.SupportedTAList
			log.Trace("Decode IE SupportedTAList")
			if supportedTAList == nil {
				log.Error("SupportedTAList is nil")
				return
			}
		case ngapType.ProtocolIEIDRANNodeName:
			ranName = ie.Value.RANNodeName
			log.Trace("Decode IE RANNodeName")
			if ranName == nil {
				log.Error("RANNodeName is nil")
				return
			}
		case ngapType.ProtocolIEIDDefaultPagingDRX:
			pagingDRX = ie.Value.DefaultPagingDRX
			log.Trace("Decode IE DefaultPagingDRX")
			if pagingDRX == nil {
				log.Error("DefaultPagingDRX is nil")
				return
			}
		}
	}

	ran.SetRanId(ranId)
	if ranName != nil {
		ran.SetName(ranName.Value)
	}
	if pagingDRX != nil {
		log.Tracef("PagingDRX[%d]", pagingDRX.Value)
	}
	//NOTE: here, free5gc implementation is not clean, we may need to clean it up
	tailist := []context.SupportedTAI{}

	for i := 0; i < len(supportedTAList.List); i++ {
		supportedTAItem := supportedTAList.List[i]
		tac := hex.EncodeToString(supportedTAItem.TAC.Value)
		capOfSupportTai := 10 //TungTQ: do it latercap(ran.SupportedTAList)
		for j := 0; j < len(supportedTAItem.BroadcastPLMNList.List); j++ {
			supportedTAI := context.NewSupportedTAI()
			supportedTAI.Tai.Tac = tac
			broadcastPLMNItem := supportedTAItem.BroadcastPLMNList.List[j]
			plmnId := ngapConvert.PlmnIdToModels(broadcastPLMNItem.PLMNIdentity)
			supportedTAI.Tai.PlmnId = plmnId
			capOfSNssaiList := cap(supportedTAI.SNssaiList)
			for k := 0; k < len(broadcastPLMNItem.TAISliceSupportList.List); k++ {
				tAISliceSupportItem := broadcastPLMNItem.TAISliceSupportList.List[k]
				if len(supportedTAI.SNssaiList) < capOfSNssaiList {
					supportedTAI.SNssaiList = append(supportedTAI.SNssaiList, ngapConvert.SNssaiToModels(tAISliceSupportItem.SNSSAI))
				} else {
					break
				}
			}
			log.Tracef("PLMN_ID[MCC:%s MNC:%s] TAC[%s]", plmnId.Mcc, plmnId.Mnc, tac)
			if len(tailist) < capOfSupportTai {
				tailist = append(tailist, supportedTAI)
			} else {
				break
			}
		}
	}
	//TODO: tqtung tailist in ran is created with a cap, need to handle following code carefuly. DO it later
	if len(tailist) == 0 {
		log.Warn("NG-Setup failure: No supported TA exist in NG-Setup request")
		cause.Present = ngapType.CausePresentMisc
		cause.Misc = &ngapType.CauseMisc{
			Value: ngapType.CauseMiscPresentUnspecified,
		}
	} else {
		var found bool
		for i, tai := range tailist {
			if context.InTaiList(tai.Tai, h.backend.Context().SupportTaiList()) {
				log.Tracef("SERVED_TAI_INDEX[%d]", i)
				found = true
				break
			}
		}
		if !found {
			log.Warn("NG-Setup failure: Cannot find Served TAI in AMF")
			cause.Present = ngapType.CausePresentMisc
			cause.Misc = &ngapType.CauseMisc{
				Value: ngapType.CauseMiscPresentUnknownPLMN,
			}
		}
	}

	if cause.Present == ngapType.CausePresentNothing {
		ran.SetSupportedTAList(tailist)
		h.sender.SendNGSetupResponse(ran)
	} else {
		h.sender.SendNGSetupFailure(ran, cause)
	}
}

func (h *ngapHandler) handleUplinkNasTransport(ran *context.AmfRan, uplinkNasTransport *ngapType.UplinkNASTransport) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var nasPdu *ngapType.NASPDU
	var userLocationInformation *ngapType.UserLocationInformation

	for i := 0; i < len(uplinkNasTransport.ProtocolIEs.List); i++ {
		ie := uplinkNasTransport.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDNASPDU:
			nasPdu = ie.Value.NASPDU
			log.Trace("Decode IE NasPdu")
			if nasPdu == nil {
				log.Error("nasPdu is nil")
				return
			}
		case ngapType.ProtocolIEIDUserLocationInformation:
			userLocationInformation = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
			if userLocationInformation == nil {
				log.Error("UserLocationInformation is nil")
				return
			}
		}
	}

	ranUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		return
	}
	amfUe := ranUe.AmfUe()
	if amfUe == nil {
		err := ranUe.Remove()
		if err != nil {
			log.Errorf(err.Error())
		}
		log.Errorf("No UE Context of RanUe with RANUENGAPID[%d] AMFUENGAPID[%d] ",
			ran_ngapid.Value, amf_ngapid.Value)
		return
	}

	log.Infof("Uplink NAS Transport (RAN UE NGAP ID: %d)", ranUe.RanUeNgapId())

	if userLocationInformation != nil {
		ranUe.UpdateLocation(userLocationInformation)
	}

	h.nas.HandleNAS(ranUe, ngapType.ProcedureCodeUplinkNASTransport, nasPdu.Value)
}

func (h *ngapHandler) handleNGReset(ran *context.AmfRan, nGReset *ngapType.NGReset) {
	var cause *ngapType.Cause
	var resetType *ngapType.ResetType

	log.Info("Handle NG Reset")

	for _, ie := range nGReset.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Error("Cause is nil")
				return
			}
		case ngapType.ProtocolIEIDResetType:
			resetType = ie.Value.ResetType
			log.Trace("Decode IE ResetType")
			if resetType == nil {
				log.Error("ResetType is nil")
				return
			}
		}
	}

	printAndGetCause(ran, cause)

	switch resetType.Present {
	case ngapType.ResetTypePresentNGInterface:
		log.Trace("ResetType Present: NG Interface")
		ran.RemoveAllUe()
		h.sender.SendNGResetAcknowledge(ran, nil, nil)
	case ngapType.ResetTypePresentPartOfNGInterface:
		log.Trace("ResetType Present: Part of NG Interface")

		partOfNGInterface := resetType.PartOfNGInterface
		if partOfNGInterface == nil {
			log.Error("PartOfNGInterface is nil")
			return
		}

		var ranUe *context.RanUe

		for _, ueItem := range partOfNGInterface.List {
			if ueItem.AMFUENGAPID != nil {
				log.Tracef("AmfUeNgapID[%d]", ueItem.AMFUENGAPID.Value)
				for _, ue := range ran.RanUeList() {
					if ue.AmfUeNgapId() == ueItem.AMFUENGAPID.Value {
						ranUe = ue
						break
					}
				}
			} else if ueItem.RANUENGAPID != nil {
				log.Tracef("RanUeNgapID[%d]", ueItem.RANUENGAPID.Value)
				ranUe = ran.RanUeFindByRanUeNgapID(ueItem.RANUENGAPID.Value)
			}

			if ranUe == nil {
				log.Warn("Cannot not find UE Context")
				if ueItem.AMFUENGAPID != nil {
					log.Warnf("AmfUeNgapID[%d]", ueItem.AMFUENGAPID.Value)
				}
				if ueItem.RANUENGAPID != nil {
					log.Warnf("RanUeNgapID[%d]", ueItem.RANUENGAPID.Value)
				}
			}

			err := ranUe.Remove()
			if err != nil {
				log.Error(err.Error())
			}
		}
		h.sender.SendNGResetAcknowledge(ran, partOfNGInterface, nil)
	default:
		log.Warnf("Invalid ResetType[%d]", resetType.Present)
	}
}

func (h *ngapHandler) handleLocationReportingFailureIndication(ran *context.AmfRan, locationReportingFailureIndication *ngapType.LocationReportingFailureIndication) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var ranUe *context.RanUe

	var cause *ngapType.Cause

	for i := 0; i < len(locationReportingFailureIndication.ProtocolIEs.List); i++ {
		ie := locationReportingFailureIndication.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Error("Cause is nil")
				return
			}
		}
	}

	printAndGetCause(ran, cause)

	ranUe = ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		return
	}
	log.Info("Handle Location Reporting Failure Indication")
	//TODO: tqtung - it seems free5gc have not implement the handling of this event
}

func (h *ngapHandler) handleInitialUEMessage(ran *context.AmfRan, initialUEMessage *ngapType.InitialUEMessage, pdu []byte) {
	amf := h.backend.Context()

	var ran_ngapid *ngapType.RANUENGAPID
	var nasPdu *ngapType.NASPDU
	var userLocationInformation *ngapType.UserLocationInformation
	var rRCEstablishmentCause *ngapType.RRCEstablishmentCause
	var fiveGSTMSI *ngapType.FiveGSTMSI
	// var aMFSetID *ngapType.AMFSetID
	var uEContextRequest *ngapType.UEContextRequest
	// var allowedNSSAI *ngapType.AllowedNSSAI

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	log.Info("Handle Initial UE Message")

	for _, ie := range initialUEMessage.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDRANUENGAPID, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDNASPDU: // reject
			nasPdu = ie.Value.NASPDU
			log.Trace("Decode IE NasPdu")
			if nasPdu == nil {
				log.Error("NasPdu is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDNASPDU,
					ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDUserLocationInformation: // reject
			userLocationInformation = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
			if userLocationInformation == nil {
				log.Error("UserLocationInformation is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDUserLocationInformation, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDRRCEstablishmentCause: // ignore
			rRCEstablishmentCause = ie.Value.RRCEstablishmentCause
			log.Trace("Decode IE RRCEstablishmentCause")
		case ngapType.ProtocolIEIDFiveGSTMSI: // optional, reject
			fiveGSTMSI = ie.Value.FiveGSTMSI
			log.Trace("Decode IE 5G-S-TMSI")
		case ngapType.ProtocolIEIDAMFSetID: // optional, ignore
			// aMFSetID = ie.Value.AMFSetID
			log.Trace("Decode IE AmfSetID")
		case ngapType.ProtocolIEIDUEContextRequest: // optional, ignore
			uEContextRequest = ie.Value.UEContextRequest
			log.Trace("Decode IE UEContextRequest")
		case ngapType.ProtocolIEIDAllowedNSSAI: // optional, reject
			// allowedNSSAI = ie.Value.AllowedNSSAI
			log.Trace("Decode IE Allowed NSSAI")
		}
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		log.Trace("Has missing reject IE(s)")

		procedureCode := ngapType.ProcedureCodeInitialUEMessage
		triggeringMessage := ngapType.TriggeringMessagePresentInitiatingMessage
		procedureCriticality := ngapType.CriticalityPresentIgnore
		criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage, &procedureCriticality,
			&iesCriticalityDiagnostics)
		h.sender.SendErrorIndication(ran, nil, ran_ngapid, nil, &criticalityDiagnostics)
	}

	ranUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe != nil && ranUe.AmfUe() == nil {
		err := ranUe.Remove()
		if err != nil {
			log.Errorln(err.Error())
		}
		ranUe = nil
	}
	if ranUe == nil {
		var err error
		ranUe, err = ran.NewRanUe(ran_ngapid.Value)
		if err != nil {
			log.Errorf("NewRanUe Error: %+v", err)
		}
		log.Debugf("New RanUe [RanUeNgapID: %d]", ranUe.RanUeNgapId())

		if fiveGSTMSI != nil {
			log.Debug("Receive 5G-S-TMSI")

			servedGuami := amf.ServedGuamiList()[0]

			// <5G-S-TMSI> := <AMF Set ID><AMF Pointer><5G-TMSI>
			// GUAMI := <MCC><MNC><AMF Region ID><AMF Set ID><AMF Pointer>
			// 5G-GUTI := <GUAMI><5G-TMSI>
			tmpReginID, _, _ := ngapConvert.AmfIdToNgap(servedGuami.AmfId)
			amfID := ngapConvert.AmfIdToModels(tmpReginID, fiveGSTMSI.AMFSetID.Value, fiveGSTMSI.AMFPointer.Value)
			tmsi := hex.EncodeToString(fiveGSTMSI.FiveGTMSI.Value)
			guti := servedGuami.PlmnId.Mcc + servedGuami.PlmnId.Mnc + amfID + tmsi

			// TODO: invoke Namf_Communication_UEContextTransfer if serving AMF has changed since
			// last Registration Request procedure
			// Described in TS 23.502 4.2.2.2.2 step 4 (without UDSF deployment)

			if amfUe, ok := amf.AmfUeFindByGuti(guti); !ok {
				log.Warnf("Unknown UE [GUTI: %s]", guti)
			} else {
				log.Tracef("find AmfUe [GUTI: %s]", guti)
				log.Debugf("AmfUe Attach RanUe [RanUeNgapID: %d]", ranUe.RanUeNgapId())
				amfUe.AttachRanUe(ranUe)
			}
		}
	} else {
		ranUe.AmfUe().AttachRanUe(ranUe)
	}

	if userLocationInformation != nil {
		ranUe.UpdateLocation(userLocationInformation)
	}

	if rRCEstablishmentCause != nil {
		log.Tracef("[Initial UE Message] RRC Establishment Cause[%d]", rRCEstablishmentCause.Value)
		ranUe.RRCEstablishmentCause = strconv.Itoa(int(rRCEstablishmentCause.Value))
	}

	if uEContextRequest != nil {
		log.Debug("Trigger initial Context Setup procedure")
		ranUe.SetUeContextRequest(true)
		// TODO: Trigger Initial Context Setup procedure
	} else {
		ranUe.SetUeContextRequest(false)
	}

	// TS 23.502 4.2.2.2.3 step 6a Nnrf_NFDiscovery_Request (NF type, AMF Set)
	// if aMFSetID != nil {
	// TODO: This is a rerouted message
	// TS 38.413: AMF shall, if supported, use the IE as described in TS 23.502
	// }

	// ng-ran propagate allowedNssai in the rerouted initial ue message (TS 38.413 8.6.5)
	// TS 23.502 4.2.2.2.3 step 4a Nnssf_NSSelection_Get
	// if allowedNSSAI != nil {
	// TODO: AMF should use it as defined in TS 23.502
	// }

	/*
		//NOTE: TungTQ remove the encoding and use the original PDU instead
		pdu, err := libngap.Encoder(*message)
		if err != nil {
			log.Errorf("libngap Encoder Error: %+v", err)
		}
	*/
	ranUe.InitialUEMessage = pdu
	h.nas.HandleNAS(ranUe, ngapType.ProcedureCodeInitialUEMessage, nasPdu.Value)
}

func (h *ngapHandler) handlePDUSessionResourceNotify(ran *context.AmfRan, PDUSessionResourceNotify *ngapType.PDUSessionResourceNotify) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var pDUSessionResourceNotifyList *ngapType.PDUSessionResourceNotifyList
	var pDUSessionResourceReleasedListNot *ngapType.PDUSessionResourceReleasedListNot
	var userLocationInformation *ngapType.UserLocationInformation

	var ranUe *context.RanUe

	for _, ie := range PDUSessionResourceNotify.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID // reject
			log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID // reject
			log.Trace("Decode IE RanUeNgapID")
		case ngapType.ProtocolIEIDPDUSessionResourceNotifyList: // reject
			pDUSessionResourceNotifyList = ie.Value.PDUSessionResourceNotifyList
			log.Trace("Decode IE pDUSessionResourceNotifyList")
			if pDUSessionResourceNotifyList == nil {
				log.Error("pDUSessionResourceNotifyList is nil")
			}
		case ngapType.ProtocolIEIDPDUSessionResourceReleasedListNot: // ignore
			pDUSessionResourceReleasedListNot = ie.Value.PDUSessionResourceReleasedListNot
			log.Trace("Decode IE PDUSessionResourceReleasedListNot")
			if pDUSessionResourceReleasedListNot == nil {
				log.Error("PDUSessionResourceReleasedListNot is nil")
			}
		case ngapType.ProtocolIEIDUserLocationInformation: // optional, ignore
			userLocationInformation = ie.Value.UserLocationInformation
			log.Trace("Decode IE userLocationInformation")
			if userLocationInformation == nil {
				log.Warn("userLocationInformation is nil [optional]")
			}
		}
	}

	ranUe = ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Warnf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
	}

	ranUe = h.backend.Context().RanUeFindByAmfUeNgapID(amf_ngapid.Value)
	if ranUe == nil {
		log.Warnf("No UE Context[AmfUeNgapID: %d]", amf_ngapid.Value)
		return
	}

	log.Info("Handle PDU Session Resource Notify")
	log.Tracef("AmfUeNgapID[%d] RanUeNgapID[%d]", ranUe.AmfUeNgapId(), ranUe.RanUeNgapId())
	amfUe := ranUe.AmfUe()
	if amfUe == nil {
		log.Error("amfUe is nil")
		return
	}

	if userLocationInformation != nil {
		ranUe.UpdateLocation(userLocationInformation)
	}

	log.Trace("Send PDUSessionResourceNotifyTransfer to SMF")

	for _, item := range pDUSessionResourceNotifyList.List {
		pduSessionID := int32(item.PDUSessionID.Value)
		transfer := item.PDUSessionResourceNotifyTransfer
		smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
		if !ok {
			log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
		}
		response, errResponse, problemDetail, err := smContext.SmfClient().SendUpdateSmContextN2Info(models.N2SMINFOTYPE_PDU_RES_NTY, transfer)
		if err != nil {
			log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceNotifyTransfer] Error: %+v", err)
		}

		if response != nil {
			responseData := response.JsonData
			n2Info := response.BinaryDataN1SmMessage
			n1Msg := response.BinaryDataN2SmInformation
			if n2Info != nil {
				switch responseData.N2SmInfoType {
				case models.N2SMINFOTYPE_PDU_RES_MOD_REQ:
					log.Debugln("AMF Transfer NGAP PDU Resource Modify Req from SMF")
					var nasPdu []byte
					if n1Msg != nil {
						pduSessionId := uint8(pduSessionID)
						nasPdu, err =
							h.nas.BuildDLNASTransport(amfUe, ran.AnType(), nasMessage.PayloadContainerTypeN1SMInfo,
								n1Msg, pduSessionId, nil, nil, 0)
						if err != nil {
							log.Warnf("GMM Message build DL NAS Transport filaed: %v", err)
						}
					}
					list := ngapType.PDUSessionResourceModifyListModReq{}
					util.AppendPDUSessionResourceModifyListModReq(&list, pduSessionID, nasPdu, n2Info)
					h.sender.SendPDUSessionResourceModifyRequest(ranUe, list)
				}
			}
		} else if errResponse != nil {
			errJSON := errResponse.JsonData
			n1Msg := errResponse.BinaryDataN2SmInformation
			log.Warnf("PDU Session Modification is rejected by SMF[pduSessionId:%d], Error[%s]\n",
				pduSessionID, errJSON.Error.Cause)
			if n1Msg != nil {
				h.nas.SendDLNASTransport(
					ranUe, nasMessage.PayloadContainerTypeN1SMInfo, errResponse.BinaryDataN1SmMessage, pduSessionID, 0, nil, 0)
			}
			// TODO: handle n2 info transfer
		} else if err != nil {
			return
		} else {
			// TODO: error handling
			log.Errorf("Failed to Update smContext[pduSessionID: %d], Error[%v]", pduSessionID, problemDetail)
			return
		}
	}

	if pDUSessionResourceReleasedListNot != nil {
		log.Trace("Send PDUSessionResourceNotifyReleasedTransfer to SMF")
		for _, item := range pDUSessionResourceReleasedListNot.List {
			pduSessionID := int32(item.PDUSessionID.Value)
			transfer := item.PDUSessionResourceNotifyReleasedTransfer
			smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
			if !ok {
				log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
			}
			response, errResponse, problemDetail, err := smContext.SmfClient().SendUpdateSmContextN2Info(models.N2SMINFOTYPE_PDU_RES_NTY_REL, transfer)
			if err != nil {
				log.Errorf("SendUpdateSmContextN2Info[PDUSessionResourceNotifyReleasedTransfer] Error: %+v", err)
			}
			if response != nil {
				responseData := response.JsonData
				n2Info := response.BinaryDataN1SmMessage
				n1Msg := response.BinaryDataN2SmInformation
				if n2Info != nil {
					switch responseData.N2SmInfoType {
					case models.N2SMINFOTYPE_PDU_RES_REL_CMD:
						log.Debugln("AMF Transfer NGAP PDU Session Resource Rel Co from SMF")
						var nasPdu []byte
						if n1Msg != nil {
							pduSessionId := uint8(pduSessionID)
							nasPdu, err = h.nas.BuildDLNASTransport(amfUe, ran.AnType(),
								nasMessage.PayloadContainerTypeN1SMInfo, n1Msg, pduSessionId, nil, nil, 0)
							if err != nil {
								log.Warnf("GMM Message build DL NAS Transport filaed: %v", err)
							}
						}
						list := ngapType.PDUSessionResourceToReleaseListRelCmd{}
						util.AppendPDUSessionResourceToReleaseListRelCmd(&list, pduSessionID, n2Info)
						h.sender.SendPDUSessionResourceReleaseCommand(ranUe, nasPdu, list)
					}
				}
			} else if errResponse != nil {
				errJSON := errResponse.JsonData
				n1Msg := errResponse.BinaryDataN2SmInformation
				log.Warnf("PDU Session Release is rejected by SMF[pduSessionId:%d], Error[%s]\n",
					pduSessionID, errJSON.Error.Cause)
				if n1Msg != nil {
					h.nas.SendDLNASTransport(
						ranUe, nasMessage.PayloadContainerTypeN1SMInfo, errResponse.BinaryDataN1SmMessage, pduSessionID, 0, nil, 0)
				}
			} else if err != nil {
				return
			} else {
				// TODO: error handling
				log.Errorf("Failed to Update smContext[pduSessionID: %d], Error[%v]", pduSessionID, problemDetail)
				return
			}
		}
	}
}

func (h *ngapHandler) handlePDUSessionResourceModifyIndication(ran *context.AmfRan, pDUSessionResourceModifyIndication *ngapType.PDUSessionResourceModifyIndication) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var pduSessionResourceModifyIndicationList *ngapType.PDUSessionResourceModifyListModInd

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	var ranUe *context.RanUe

	/*
		//NOTE: tungtq - need a special treatment
		initiatingMessage := message.InitiatingMessage // reject
		if initiatingMessage == nil {
		log.Error("InitiatingMessage is nil")
			cause := ngapType.Cause{
				Present: ngapType.CausePresentProtocol,
				Protocol: &ngapType.CauseProtocol{
					Value: ngapType.CauseProtocolPresentAbstractSyntaxErrorReject,
				},
			}
			h.sender.SendErrorIndication(ran, nil, nil, &cause, nil)
			return
		}
	*/

	/*
		//NOTE: tungtq - need a special treatment
		pDUSessionResourceModifyIndication := initiatingMessage.Value.PDUSessionResourceModifyIndication
		if pDUSessionResourceModifyIndication == nil {
		log.Error("PDUSessionResourceModifyIndication is nil")
			cause := ngapType.Cause{
				Present: ngapType.CausePresentProtocol,
				Protocol: &ngapType.CauseProtocol{
					Value: ngapType.CauseProtocolPresentAbstractSyntaxErrorReject,
				},
			}
			h.sender.SendErrorIndication(ran, nil, nil, &cause, nil)
			return
		}
	*/
	for _, ie := range pDUSessionResourceModifyIndication.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDAMFUENGAPID, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDRANUENGAPID, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDPDUSessionResourceModifyListModInd: // reject
			pduSessionResourceModifyIndicationList = ie.Value.PDUSessionResourceModifyListModInd
			log.Trace("Decode IE PDUSessionResourceModifyListModInd")
			if pduSessionResourceModifyIndicationList == nil {
				log.Error("PDUSessionResourceModifyListModInd is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
					ngapType.ProtocolIEIDPDUSessionResourceModifyListModInd, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		}
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		log.Error("Has missing reject IE(s)")

		procedureCode := ngapType.ProcedureCodePDUSessionResourceModifyIndication
		triggeringMessage := ngapType.TriggeringMessagePresentInitiatingMessage
		procedureCriticality := ngapType.CriticalityPresentReject
		criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage, &procedureCriticality,
			&iesCriticalityDiagnostics)
		h.sender.SendErrorIndication(ran, amf_ngapid, ran_ngapid, nil, &criticalityDiagnostics)
		return
	}

	ranUe = ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		cause := ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		h.sender.SendErrorIndication(ran, amf_ngapid, ran_ngapid, &cause, nil)
		return
	}

	log.Info("Handle PDU Session Resource Modify Indication")
	log.Tracef("UE Context AmfUeNgapID[%d] RanUeNgapID[%d]", ranUe.AmfUeNgapId(), ranUe.RanUeNgapId())

	amfUe := ranUe.AmfUe()
	if amfUe == nil {
		log.Error("AmfUe is nil")
		return
	}

	pduSessionResourceModifyListModCfm := ngapType.PDUSessionResourceModifyListModCfm{}
	pduSessionResourceFailedToModifyListModCfm := ngapType.PDUSessionResourceFailedToModifyListModCfm{}

	log.Trace("Send PDUSessionResourceModifyIndicationTransfer to SMF")
	for _, item := range pduSessionResourceModifyIndicationList.List {
		pduSessionID := int32(item.PDUSessionID.Value)
		transfer := item.PDUSessionResourceModifyIndicationTransfer
		smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
		if !ok {
			log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
		}
		response, errResponse, _, err := smContext.SmfClient().SendUpdateSmContextN2Info(models.N2SMINFOTYPE_PDU_RES_MOD_IND, transfer)
		if err != nil {
			log.Errorf("SendUpdateSmContextN2Info Error:\n%s", err.Error())
		}

		if response != nil && response.BinaryDataN2SmInformation != nil {
			util.AppendPDUSessionResourceModifyListModCfm(&pduSessionResourceModifyListModCfm, int64(pduSessionID),
				response.BinaryDataN2SmInformation)
		}
		if errResponse != nil && errResponse.BinaryDataN2SmInformation != nil {
			util.AppendPDUSessionResourceFailedToModifyListModCfm(&pduSessionResourceFailedToModifyListModCfm,
				int64(pduSessionID), errResponse.BinaryDataN2SmInformation)
		}
	}

	h.sender.SendPDUSessionResourceModifyConfirm(ranUe, pduSessionResourceModifyListModCfm,
		pduSessionResourceFailedToModifyListModCfm, nil)
}

func (h *ngapHandler) handleUEContextReleaseRequest(ran *context.AmfRan, uEContextReleaseRequest *ngapType.UEContextReleaseRequest) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var pDUSessionResourceList *ngapType.PDUSessionResourceListCxtRelReq
	var cause *ngapType.Cause

	for _, ie := range uEContextReleaseRequest.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDPDUSessionResourceListCxtRelReq:
			pDUSessionResourceList = ie.Value.PDUSessionResourceListCxtRelReq
			log.Trace("Decode IE Pdu Session Resource List")
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Warn("Cause is nil")
			}
		}
	}

	ranUe := h.backend.Context().RanUeFindByAmfUeNgapID(amf_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No RanUe Context[AmfUeNgapID: %d]", amf_ngapid.Value)
		cause = &ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		h.sender.SendErrorIndication(ran, amf_ngapid, ran_ngapid, cause, nil)
		return
	}

	log.Tracef("RanUeNgapID[%d] AmfUeNgapID[%d]", ranUe.RanUeNgapId(), ranUe.AmfUeNgapId())
	log.Info("Handle UE Context Release Request")

	causeGroup := ngapType.CausePresentRadioNetwork
	causeValue := ngapType.CauseRadioNetworkPresentUnspecified
	if cause != nil {
		causeGroup, causeValue = printAndGetCause(ran, cause)
	}

	amfUe := ranUe.AmfUe()
	if amfUe != nil {
		if !h.isLatestAmfUe(amfUe) {
			amfUe.DetachRanUe(ran.AnType())
			ranUe.DetachAmfUe()
			//TODO: comment out by tungtq
			//gmm_common.StopAll5GSMMTimers(amfUe)
			causeValue = ngapType.CauseRadioNetworkPresentReleaseDueToNgranGeneratedReason
			h.sender.SendUEContextReleaseCommand(ranUe, context.UeContextReleaseUeContext, causeGroup, causeValue)
			return
		}
		//TODO: comment out by tungtq
		//	gmm_common.StopAll5GSMMTimers(amfUe)
		causeAll := context.CauseAll{
			NgapCause: &models.NgApCause{
				Group: int32(causeGroup),
				Value: int32(causeValue),
			},
		}
		if amfUe.State[ran.AnType()].Is(context.Registered) {
			log.Info("Ue Context in GMM-Registered")
			if pDUSessionResourceList != nil {
				for _, pduSessionReourceItem := range pDUSessionResourceList.List {
					pduSessionID := int32(pduSessionReourceItem.PDUSessionID.Value)
					smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
					if !ok {
						log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
						// TODO: Check if doing error handling here
						continue
					}
					response, _, _, err := smContext.SmfClient().SendUpdateSmContextDeactivateUpCnxState(causeAll)
					if err != nil {
						log.Errorf("Send Update SmContextDeactivate UpCnxState Error[%s]", err.Error())
					} else if response == nil {
						log.Errorln("Send Update SmContextDeactivate UpCnxState Error")
					}
				}
			}
		} else {
			log.Info("Ue Context in Non GMM-Registered")
			amfUe.SmContextList.Range(func(key, value interface{}) bool {
				smContext := value.(*context.SmContext)
				detail, err := smContext.SmfClient().SendReleaseSmContextRequest(&causeAll, "", nil)
				if err != nil {
					log.Errorf("Send ReleaseSmContextRequest Error[%s]", err.Error())
				} else if detail != nil {
					log.Errorf("Send ReleaseSmContextRequeste Error[%s]", detail.Cause)
				}
				return true
			})
			h.sender.SendUEContextReleaseCommand(ranUe, context.UeContextReleaseUeContext, causeGroup, causeValue)
			return
		}
	}
	h.sender.SendUEContextReleaseCommand(ranUe, context.UeContextN2NormalRelease, causeGroup, causeValue)
}

func (h *ngapHandler) handleRRCInactiveTransitionReport(ran *context.AmfRan, rRCInactiveTransitionReport *ngapType.RRCInactiveTransitionReport) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var rRCState *ngapType.RRCState
	var userLocationInformation *ngapType.UserLocationInformation

	for i := 0; i < len(rRCInactiveTransitionReport.ProtocolIEs.List); i++ {
		ie := rRCInactiveTransitionReport.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRRCState: // ignore
			rRCState = ie.Value.RRCState
			log.Trace("Decode IE RRCState")
			if rRCState == nil {
				log.Error("RRCState is nil")
				return
			}
		case ngapType.ProtocolIEIDUserLocationInformation: // ignore
			userLocationInformation = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
			if userLocationInformation == nil {
				log.Error("UserLocationInformation is nil")
				return
			}
		}
	}

	ranUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Warnf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
	} else {
		log.Tracef("RANUENGAPID[%d] AMFUENGAPID[%d]", ranUe.RanUeNgapId(), ranUe.AmfUeNgapId())
		log.Info("Handle RRC Inactive Transition Report")

		if rRCState != nil {
			switch rRCState.Value {
			case ngapType.RRCStatePresentInactive:
				log.Trace("UE RRC State: Inactive")
			case ngapType.RRCStatePresentConnected:
				log.Trace("UE RRC State: Connected")
			}
		}
		ranUe.UpdateLocation(userLocationInformation)
	}
}

func (h *ngapHandler) handleHandoverNotify(ran *context.AmfRan, HandoverNotify *ngapType.HandoverNotify) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var userLocationInformation *ngapType.UserLocationInformation

	for i := 0; i < len(HandoverNotify.ProtocolIEs.List); i++ {
		ie := HandoverNotify.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AMFUENGAPID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RANUENGAPID is nil")
				return
			}
		case ngapType.ProtocolIEIDUserLocationInformation:
			userLocationInformation = ie.Value.UserLocationInformation
			log.Trace("Decode IE userLocationInformation")
			if userLocationInformation == nil {
				log.Error("userLocationInformation is nil")
				return
			}
		}
	}

	targetUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)

	if targetUe == nil {
		log.Errorf("No RanUe Context[AmfUeNgapID: %d]", amf_ngapid.Value)
		cause := ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		h.sender.SendErrorIndication(ran, amf_ngapid, ran_ngapid, &cause, nil)
		return
	}

	//	targetUe.Log.Info("Handle Handover notification")

	if userLocationInformation != nil {
		targetUe.UpdateLocation(userLocationInformation)
	}
	amfUe := targetUe.AmfUe()
	if amfUe == nil {
		log.Error("AmfUe is nil")
		return
	}
	sourceUe := targetUe.GetHandoverInfo().SourceUe
	if sourceUe == nil {
		// TODO: Send to S-AMF
		// Desciibed in (23.502 4.9.1.3.3) [conditional] 6a.Namf_Communication_N2InfoNotify.
		log.Error("N2 Handover between AMF has not been implemented yet")
	} else {
		//		targetUe.Log.Info("Handle Handover notification Finshed ")
		for _, pduSessionid := range targetUe.GetHandoverInfo().SuccessPduSessionId {
			smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionid)
			if !ok {
				//				sourceUe.Log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionid)
			}
			_, _, _, err := smContext.SmfClient().SendUpdateSmContextN2HandoverComplete("", nil)
			if err != nil {
				log.Errorf("Send UpdateSmContextN2HandoverComplete Error[%s]", err.Error())
			}
		}
		amfUe.AttachRanUe(targetUe)

		h.sender.SendUEContextReleaseCommand(sourceUe, context.UeContextReleaseHandover, ngapType.CausePresentNas,
			ngapType.CauseNasPresentNormalRelease)
	}

	// TODO: The UE initiates Mobility Registration Update procedure as described in clause 4.2.2.2.2.
}

// TS 23.502 4.9.1
func (h *ngapHandler) handlePathSwitchRequest(ran *context.AmfRan, pathSwitchRequest *ngapType.PathSwitchRequest) {
	var ran_ngapid *ngapType.RANUENGAPID
	var sourceAMFUENGAPID *ngapType.AMFUENGAPID
	var userLocationInformation *ngapType.UserLocationInformation
	var uESecurityCapabilities *ngapType.UESecurityCapabilities
	var pduSessionResourceToBeSwitchedInDLList *ngapType.PDUSessionResourceToBeSwitchedDLList
	var pduSessionResourceFailedToSetupList *ngapType.PDUSessionResourceFailedToSetupListPSReq

	var ranUe *context.RanUe

	for _, ie := range pathSwitchRequest.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDSourceAMFUENGAPID: // reject
			sourceAMFUENGAPID = ie.Value.SourceAMFUENGAPID
			log.Trace("Decode IE SourceAmfUeNgapID")
			if sourceAMFUENGAPID == nil {
				log.Error("SourceAmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDUserLocationInformation: // ignore
			userLocationInformation = ie.Value.UserLocationInformation
			log.Trace("Decode IE UserLocationInformation")
		case ngapType.ProtocolIEIDUESecurityCapabilities: // ignore
			uESecurityCapabilities = ie.Value.UESecurityCapabilities
			log.Trace("Decode IE UESecurityCapabilities")
		case ngapType.ProtocolIEIDPDUSessionResourceToBeSwitchedDLList: // reject
			pduSessionResourceToBeSwitchedInDLList = ie.Value.PDUSessionResourceToBeSwitchedDLList
			log.Trace("Decode IE PDUSessionResourceToBeSwitchedDLList")
			if pduSessionResourceToBeSwitchedInDLList == nil {
				log.Error("PDUSessionResourceToBeSwitchedDLList is nil")
				return
			}
		case ngapType.ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq: // ignore
			pduSessionResourceFailedToSetupList = ie.Value.PDUSessionResourceFailedToSetupListPSReq
			log.Trace("Decode IE PDUSessionResourceFailedToSetupListPSReq")
		}
	}

	if sourceAMFUENGAPID == nil {
		log.Error("SourceAmfUeNgapID is nil")
		return
	}
	ranUe = h.backend.Context().RanUeFindByAmfUeNgapID(sourceAMFUENGAPID.Value)
	if ranUe == nil {
		log.Errorf("Cannot find UE from sourceAMfUeNgapID[%d]", sourceAMFUENGAPID.Value)
		h.sender.SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ran_ngapid.Value, nil, nil)
		return
	}

	log.Tracef("AmfUeNgapID[%d] RanUeNgapID[%d]", ranUe.AmfUeNgapId(), ranUe.RanUeNgapId())
	log.Info("Handle Path Switch Request")

	amfUe := ranUe.AmfUe()
	ausf := amfUe.AusfClient()

	if amfUe == nil {
		log.Error("AmfUe is nil")
		h.sender.SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ran_ngapid.Value, nil, nil)
		return
	}

	if ausf.SecurityContextIsValid() {
		// Update NH
		ausf.UpdateNH()
	} else {
		log.Errorf("No Security Context : SUPI[%s]", amfUe.Supi)
		h.sender.SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ran_ngapid.Value, nil, nil)
		return
	}

	ausf.SetUeSecCap(uESecurityCapabilities)

	if ran_ngapid != nil {
		ranUe.SetRanUeNgapId(ran_ngapid.Value)
	}

	ranUe.UpdateLocation(userLocationInformation)

	var pduSessionResourceSwitchedList ngapType.PDUSessionResourceSwitchedList
	var pduSessionResourceReleasedListPSAck ngapType.PDUSessionResourceReleasedListPSAck
	var pduSessionResourceReleasedListPSFail ngapType.PDUSessionResourceReleasedListPSFail

	if pduSessionResourceToBeSwitchedInDLList != nil {
		for _, item := range pduSessionResourceToBeSwitchedInDLList.List {
			pduSessionID := int32(item.PDUSessionID.Value)
			transfer := item.PathSwitchRequestTransfer
			smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
			if !ok {
				log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
			}
			response, errResponse, _, err := smContext.SmfClient().SendUpdateSmContextXnHandover(models.N2SMINFOTYPE_PATH_SWITCH_REQ, transfer)
			if err != nil {
				log.Errorf("SendUpdateSmContextXnHandover[PathSwitchRequestTransfer] Error:\n%s", err.Error())
			}
			if response != nil && response.BinaryDataN2SmInformation != nil {
				pduSessionResourceSwitchedItem := ngapType.PDUSessionResourceSwitchedItem{}
				pduSessionResourceSwitchedItem.PDUSessionID.Value = int64(pduSessionID)
				pduSessionResourceSwitchedItem.PathSwitchRequestAcknowledgeTransfer = response.BinaryDataN2SmInformation
				pduSessionResourceSwitchedList.List = append(pduSessionResourceSwitchedList.List, pduSessionResourceSwitchedItem)
			}
			if errResponse != nil && errResponse.BinaryDataN2SmInformation != nil {
				pduSessionResourceReleasedItem := ngapType.PDUSessionResourceReleasedItemPSFail{}
				pduSessionResourceReleasedItem.PDUSessionID.Value = int64(pduSessionID)
				pduSessionResourceReleasedItem.PathSwitchRequestUnsuccessfulTransfer = errResponse.BinaryDataN2SmInformation
				pduSessionResourceReleasedListPSFail.List = append(pduSessionResourceReleasedListPSFail.List,
					pduSessionResourceReleasedItem)
			}
		}
	}

	if pduSessionResourceFailedToSetupList != nil {
		for _, item := range pduSessionResourceFailedToSetupList.List {
			pduSessionID := int32(item.PDUSessionID.Value)
			transfer := item.PathSwitchRequestSetupFailedTransfer
			smContext, ok := amfUe.SmContextFindByPDUSessionID(pduSessionID)
			if !ok {
				log.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionID)
			}
			response, errResponse, _, err := smContext.SmfClient().SendUpdateSmContextXnHandoverFailed(models.N2SMINFOTYPE_PATH_SWITCH_SETUP_FAIL, transfer)
			if err != nil {
				log.Errorf("SendUpdateSmContextXnHandoverFailed[PathSwitchRequestSetupFailedTransfer] Error: %+v", err)
			}
			if response != nil && response.BinaryDataN2SmInformation != nil {
				pduSessionResourceReleasedItem := ngapType.PDUSessionResourceReleasedItemPSAck{}
				pduSessionResourceReleasedItem.PDUSessionID.Value = int64(pduSessionID)
				pduSessionResourceReleasedItem.PathSwitchRequestUnsuccessfulTransfer = response.BinaryDataN2SmInformation
				pduSessionResourceReleasedListPSAck.List = append(pduSessionResourceReleasedListPSAck.List,
					pduSessionResourceReleasedItem)
			}
			if errResponse != nil && errResponse.BinaryDataN2SmInformation != nil {
				pduSessionResourceReleasedItem := ngapType.PDUSessionResourceReleasedItemPSFail{}
				pduSessionResourceReleasedItem.PDUSessionID.Value = int64(pduSessionID)
				pduSessionResourceReleasedItem.PathSwitchRequestUnsuccessfulTransfer = errResponse.BinaryDataN2SmInformation
				pduSessionResourceReleasedListPSFail.List = append(pduSessionResourceReleasedListPSFail.List,
					pduSessionResourceReleasedItem)
			}
		}
	}

	// TS 23.502 4.9.1.2.2 step 7: send ack to Target NG-RAN. If none of the requested PDU Sessions have been switched
	// successfully, the AMF shall send an N2 Path Switch Request Failure message to the Target NG-RAN
	if len(pduSessionResourceSwitchedList.List) > 0 {
		// TODO: set newSecurityContextIndicator to true if there is a new security context
		err := ranUe.SwitchToRan(ran, ran_ngapid.Value)
		if err != nil {
			log.Error(err.Error())
			return
		}
		h.sender.SendPathSwitchRequestAcknowledge(ranUe, pduSessionResourceSwitchedList,
			pduSessionResourceReleasedListPSAck, false, nil, nil, nil)
	} else if len(pduSessionResourceReleasedListPSFail.List) > 0 {
		h.sender.SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ran_ngapid.Value,
			&pduSessionResourceReleasedListPSFail, nil)
	} else {
		h.sender.SendPathSwitchRequestFailure(ran, sourceAMFUENGAPID.Value, ran_ngapid.Value, nil, nil)
	}
}

func (h *ngapHandler) handleHandoverRequired(ran *context.AmfRan, HandoverRequired *ngapType.HandoverRequired) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var handoverType *ngapType.HandoverType
	var cause *ngapType.Cause
	var targetID *ngapType.TargetID
	var pDUSessionResourceListHORqd *ngapType.PDUSessionResourceListHORqd
	var sourceToTargetTransparentContainer *ngapType.SourceToTargetTransparentContainer
	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	for i := 0; i < len(HandoverRequired.ProtocolIEs.List); i++ {
		ie := HandoverRequired.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID // reject
			log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
		case ngapType.ProtocolIEIDHandoverType: // reject
			handoverType = ie.Value.HandoverType
			log.Trace("Decode IE HandoverType")
		case ngapType.ProtocolIEIDCause: // ignore
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
		case ngapType.ProtocolIEIDTargetID: // reject
			targetID = ie.Value.TargetID
			log.Trace("Decode IE TargetID")
		case ngapType.ProtocolIEIDPDUSessionResourceListHORqd: // reject
			pDUSessionResourceListHORqd = ie.Value.PDUSessionResourceListHORqd
			log.Trace("Decode IE PDUSessionResourceListHORqd")
		case ngapType.ProtocolIEIDSourceToTargetTransparentContainer: // reject
			sourceToTargetTransparentContainer = ie.Value.SourceToTargetTransparentContainer
			log.Trace("Decode IE SourceToTargetTransparentContainer")
		}
	}

	if amf_ngapid == nil {
		log.Error("AmfUeNgapID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDAMFUENGAPID,
			ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}
	if ran_ngapid == nil {
		log.Error("RanUeNgapID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDRANUENGAPID,
			ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}

	if handoverType == nil {
		log.Error("handoverType is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDHandoverType,
			ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}
	if targetID == nil {
		log.Error("targetID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDTargetID,
			ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}
	if pDUSessionResourceListHORqd == nil {
		log.Error("pDUSessionResourceListHORqd is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
			ngapType.ProtocolIEIDPDUSessionResourceListHORqd, ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}
	if sourceToTargetTransparentContainer == nil {
		log.Error("sourceToTargetTransparentContainer is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject,
			ngapType.ProtocolIEIDSourceToTargetTransparentContainer, ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		procedureCode := ngapType.ProcedureCodeHandoverPreparation
		triggeringMessage := ngapType.TriggeringMessagePresentInitiatingMessage
		procedureCriticality := ngapType.CriticalityPresentReject
		criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage,
			&procedureCriticality, &iesCriticalityDiagnostics)
		h.sender.SendErrorIndication(ran, amf_ngapid, ran_ngapid, nil, &criticalityDiagnostics)
		return
	}

	sourceUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if sourceUe == nil {
		log.Errorf("Cannot find UE for RAN_UE_NGAP_ID[%d] ", ran_ngapid.Value)
		cause := ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		h.sender.SendErrorIndication(ran, amf_ngapid, ran_ngapid, &cause, nil)
		return
	}

	//	sourceUe.Log.Info("Handle HandoverRequired")

	amfUe := sourceUe.AmfUe()
	ausf := amfUe.AusfClient()

	if amfUe == nil {
		log.Error("Cannot find amfUE from sourceUE")
		return
	}

	if targetID.Present != ngapType.TargetIDPresentTargetRANNodeID {
		log.Errorf("targetID type[%d] is not supported", targetID.Present)
		return
	}
	amfUe.SetOnGoing(sourceUe.Ran().AnType(), &context.OnGoing{
		Procedure: context.OnGoingProcedureN2Handover,
	})
	if !ausf.SecurityContextIsValid() {
		//		sourceUe.Log.Info("Handle Handover Preparation Failure [Authentication Failure]")
		cause = &ngapType.Cause{
			Present: ngapType.CausePresentNas,
			Nas: &ngapType.CauseNas{
				Value: ngapType.CauseNasPresentAuthenticationFailure,
			},
		}
		h.sender.SendHandoverPreparationFailure(sourceUe, *cause, nil)
		return
	}
	targetRanNodeId := ngapConvert.RanIdToModels(targetID.TargetRANNodeID.GlobalRANNodeID)
	targetRan, ok := h.backend.Context().AmfRanFindByRanID(targetRanNodeId)
	if !ok {
		// handover between different AMF
		//		sourceUe.Log.Warnf("Handover required : cannot find target Ran Node Id[%+v] in this AMF", targetRanNodeId)
		//		sourceUe.Log.Error("Handover between different AMF has not been implemented yet")
		return
		// TODO: Send to T-AMF
		// Described in (23.502 4.9.1.3.2) step 3.Namf_Communication_CreateUEContext Request
	} else {
		// Handover in same AMF
		sourceUe.GetHandoverInfo().HandOverType.Value = handoverType.Value
		tai := ngapConvert.TaiToModels(targetID.TargetRANNodeID.SelectedTAI)
		targetId := models.NgRanTargetId{
			RanNodeId: targetRanNodeId,
			Tai:       tai,
		}
		var pduSessionReqList ngapType.PDUSessionResourceSetupListHOReq
		for _, pDUSessionResourceHoItem := range pDUSessionResourceListHORqd.List {
			pduSessionId := int32(pDUSessionResourceHoItem.PDUSessionID.Value)
			if smContext, exist := amfUe.SmContextFindByPDUSessionID(pduSessionId); exist {
				response, _, _, err := smContext.SmfClient().SendUpdateSmContextN2HandoverPreparing(models.N2SMINFOTYPE_HANDOVER_REQUIRED, pDUSessionResourceHoItem.HandoverRequiredTransfer, "", &targetId)
				if err != nil {
					//					sourceUe.Log.Errorf("consumer.SendUpdateSmContextN2HandoverPreparing Error: %+v", err)
				}
				if response == nil {
					//					sourceUe.Log.Errorf("SendUpdateSmContextN2HandoverPreparing Error for PduSessionId[%d]", pduSessionId)
					continue
				} else if response.BinaryDataN2SmInformation != nil {
					util.AppendPDUSessionResourceSetupListHOReq(&pduSessionReqList, pduSessionId,
						smContext.Snssai(), response.BinaryDataN2SmInformation)
				}
			}
		}
		if len(pduSessionReqList.List) == 0 {
			//			sourceUe.Log.Info("Handle Handover Preparation Failure [HoFailure In Target5GC NgranNode Or TargetSystem]")
			cause = &ngapType.Cause{
				Present: ngapType.CausePresentRadioNetwork,
				RadioNetwork: &ngapType.CauseRadioNetwork{
					Value: ngapType.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem,
				},
			}
			h.sender.SendHandoverPreparationFailure(sourceUe, *cause, nil)
			return
		}
		// Update NH
		ausf.UpdateNH()
		h.sender.SendHandoverRequest(sourceUe, targetRan, *cause, pduSessionReqList,
			*sourceToTargetTransparentContainer, false)
	}
}

func (h *ngapHandler) handleHandoverCancel(ran *context.AmfRan, HandoverCancel *ngapType.HandoverCancel) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var cause *ngapType.Cause

	for i := 0; i < len(HandoverCancel.ProtocolIEs.List); i++ {
		ie := HandoverCancel.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AMFUENGAPID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RANUENGAPID is nil")
				return
			}
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
			if cause == nil {
				log.Error(cause, "cause is nil")
				return
			}
		}
	}

	sourceUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if sourceUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		cause := ngapType.Cause{
			Present: ngapType.CausePresentRadioNetwork,
			RadioNetwork: &ngapType.CauseRadioNetwork{
				Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
			},
		}
		h.sender.SendErrorIndication(ran, amf_ngapid, ran_ngapid, &cause, nil)
		return
	}

	//	sourceUe.Log.Info("Handle Handover Cancel")

	if sourceUe.AmfUeNgapId() != amf_ngapid.Value {
		log.Warnf("Conflict AMF_UE_NGAP_ID : %d != %d", sourceUe.AmfUeNgapId(), amf_ngapid.Value)
	}
	log.Tracef("Source: RAN_UE_NGAP_ID[%d] AMF_UE_NGAP_ID[%d]", sourceUe.RanUeNgapId(), sourceUe.AmfUeNgapId())

	causePresent := ngapType.CausePresentRadioNetwork
	causeValue := ngapType.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem
	if cause != nil {
		causePresent, causeValue = printAndGetCause(ran, cause)
	}
	targetUe := sourceUe.GetHandoverInfo().TargetUe
	if targetUe == nil {
		// Described in (23.502 4.11.1.2.3) step 2
		// Todo : send to T-AMF invoke Namf_UeContextReleaseRequest(targetUe)
		log.Error("N2 Handover between AMF has not been implemented yet")
	} else {
		log.Tracef("Target : RAN_UE_NGAP_ID[%d] AMF_UE_NGAP_ID[%d]", targetUe.RanUeNgapId(), targetUe.AmfUeNgapId())
		amfUe := sourceUe.AmfUe()
		if amfUe != nil {
			amfUe.SmContextList.Range(func(key, value interface{}) bool {
				//pduSessionID := key.(int32)
				smContext := value.(*context.SmContext)
				causeAll := context.CauseAll{
					NgapCause: &models.NgApCause{
						Group: int32(causePresent),
						Value: int32(causeValue),
					},
				}
				_, _, _, err := smContext.SmfClient().SendUpdateSmContextN2HandoverCanceled(causeAll)
				if err != nil {
					//					sourceUe.Log.Errorf("Send UpdateSmContextN2HandoverCanceled Error for PduSessionId[%d]", pduSessionID)
				}
				return true
			})
		}
		h.sender.SendUEContextReleaseCommand(targetUe, context.UeContextReleaseHandover, causePresent, causeValue)
		h.sender.SendHandoverCancelAcknowledge(sourceUe, nil)
	}
}

func (h *ngapHandler) handleUplinkRanStatusTransfer(ran *context.AmfRan, uplinkRanStatusTransfer *ngapType.UplinkRANStatusTransfer) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var rANStatusTransferTransparentContainer *ngapType.RANStatusTransferTransparentContainer
	var ranUe *context.RanUe

	for _, ie := range uplinkRanStatusTransfer.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANStatusTransferTransparentContainer: // reject
			rANStatusTransferTransparentContainer = ie.Value.RANStatusTransferTransparentContainer
			log.Trace("Decode IE RANStatusTransferTransparentContainer")
			if rANStatusTransferTransparentContainer == nil {
				log.Error("RANStatusTransferTransparentContainer is nil")
			}
		}
	}

	ranUe = ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		return
	}

	log.Tracef("UE Context AmfUeNgapID[%d] RanUeNgapID[%d]", ranUe.AmfUeNgapId(), ranUe.RanUeNgapId())
	log.Info("Handle Uplink Ran Status Transfer")

	amfUe := ranUe.AmfUe()
	if amfUe == nil {
		log.Error("AmfUe is nil")
		return
	}
	// TODO: send to T-AMF using N1N2MessageTransfer (R16)
}

func (h *ngapHandler) handleNasNonDeliveryIndication(ran *context.AmfRan, nASNonDeliveryIndication *ngapType.NASNonDeliveryIndication) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var nasPdu *ngapType.NASPDU
	var cause *ngapType.Cause

	for _, ie := range nASNonDeliveryIndication.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDNASPDU:
			nasPdu = ie.Value.NASPDU
			if nasPdu == nil {
				log.Error("NasPdu is nil")
				return
			}
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			if cause == nil {
				log.Error("Cause is nil")
				return
			}
		}
	}

	ranUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		return
	}

	log.Tracef("RanUeNgapID[%d] AmfUeNgapID[%d]", ranUe.RanUeNgapId(), ranUe.AmfUeNgapId())
	log.Info("Handle Nas Non Delivery Indication")

	printAndGetCause(ran, cause)

	h.nas.HandleNAS(ranUe, ngapType.ProcedureCodeNASNonDeliveryIndication, nasPdu.Value)
}

func (h *ngapHandler) handleRanConfigurationUpdate(ran *context.AmfRan, rANConfigurationUpdate *ngapType.RANConfigurationUpdate) {
	/*
		//TODO: TungTQ do it later
		var ranName *ngapType.RANNodeName
		var supportedTAList *ngapType.SupportedTAList
		var pagingDRX *ngapType.PagingDRX

		var cause ngapType.Cause

		log.Info("Handle Ran Configuration Update")
		for i := 0; i < len(rANConfigurationUpdate.ProtocolIEs.List); i++ {
			ie := rANConfigurationUpdate.ProtocolIEs.List[i]
			switch ie.Id.Value {
			case ngapType.ProtocolIEIDRANNodeName:
				ranName = ie.Value.RANNodeName
				if ranName == nil {
		log.Error("RAN Node Name is nil")
					return
				}
		log.Tracef("Decode IE RANNodeName = [%s]", ranName.Value)
			case ngapType.ProtocolIEIDSupportedTAList:
				supportedTAList = ie.Value.SupportedTAList
		log.Trace("Decode IE SupportedTAList")
				if supportedTAList == nil {
		log.Error("Supported TA List is nil")
					return
				}
			case ngapType.ProtocolIEIDDefaultPagingDRX:
				pagingDRX = ie.Value.DefaultPagingDRX
				if pagingDRX == nil {
		log.Error("PagingDRX is nil")
					return
				}
		log.Tracef("Decode IE PagingDRX = [%d]", pagingDRX.Value)
			}
		}

		for i := 0; i < len(supportedTAList.List); i++ {
			supportedTAItem := supportedTAList.List[i]
			tac := hex.EncodeToString(supportedTAItem.TAC.Value)
			capOfSupportTai := cap(ran.SupportedTAList)
			for j := 0; j < len(supportedTAItem.BroadcastPLMNList.List); j++ {
				supportedTAI := context.NewSupportedTAI()
				supportedTAI.Tai.Tac = tac
				broadcastPLMNItem := supportedTAItem.BroadcastPLMNList.List[j]
				plmnId := ngapConvert.PlmnIdToModels(broadcastPLMNItem.PLMNIdentity)
				supportedTAI.Tai.PlmnId = &plmnId
				capOfSNssaiList := cap(supportedTAI.SNssaiList)
				for k := 0; k < len(broadcastPLMNItem.TAISliceSupportList.List); k++ {
					tAISliceSupportItem := broadcastPLMNItem.TAISliceSupportList.List[k]
					if len(supportedTAI.SNssaiList) < capOfSNssaiList {
						supportedTAI.SNssaiList = append(supportedTAI.SNssaiList, ngapConvert.SNssaiToModels(tAISliceSupportItem.SNSSAI))
					} else {
						break
					}
				}
		log.Tracef("PLMN_ID[MCC:%s MNC:%s] TAC[%s]", plmnId.Mcc, plmnId.Mnc, tac)
				if len(ran.SupportedTAList) < capOfSupportTai {
					ran.SupportedTAList = append(ran.SupportedTAList, supportedTAI)
				} else {
					break
				}
			}
		}

		if len(ran.SupportedTAList) == 0 {
				log.Warn("RanConfigurationUpdate failure: No supported TA exist in RanConfigurationUpdate")
			cause.Present = ngapType.CausePresentMisc
			cause.Misc = &ngapType.CauseMisc{
				Value: ngapType.CauseMiscPresentUnspecified,
			}
		} else {
			var found bool
			for i, tai := range ran.SupportedTAList {
				if context.InTaiList(tai.Tai, h.backend.Context().SupportTaiList()) {
		log.Tracef("SERVED_TAI_INDEX[%d]", i)
					found = true
					break
				}
			}
			if !found {
					log.Warn("RanConfigurationUpdate failure: Cannot find Served TAI in AMF")
				cause.Present = ngapType.CausePresentMisc
				cause.Misc = &ngapType.CauseMisc{
					Value: ngapType.CauseMiscPresentUnknownPLMN,
				}
			}
		}

		if cause.Present == ngapType.CausePresentNothing {
				log.Info("Handle RanConfigurationUpdateAcknowledge")
			h.sender.SendRanConfigurationUpdateAcknowledge(ran, nil)
		} else {
				log.Info("Handle RanConfigurationUpdateAcknowledgeFailure")
			h.sender.SendRanConfigurationUpdateFailure(ran, cause, nil)
		}
	*/
}

func (h *ngapHandler) handleUplinkRanConfigurationTransfer(ran *context.AmfRan, uplinkRANConfigurationTransfer *ngapType.UplinkRANConfigurationTransfer) {
	var sONConfigurationTransferUL *ngapType.SONConfigurationTransfer

	for _, ie := range uplinkRANConfigurationTransfer.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDSONConfigurationTransferUL: // optional, ignore
			sONConfigurationTransferUL = ie.Value.SONConfigurationTransferUL
			log.Trace("Decode IE SONConfigurationTransferUL")
			if sONConfigurationTransferUL == nil {
				log.Warn("sONConfigurationTransferUL is nil")
			}
		}
	}

	if sONConfigurationTransferUL != nil {
		targetRanNodeID := ngapConvert.RanIdToModels(sONConfigurationTransferUL.TargetRANNodeID.GlobalRANNodeID)

		if targetRanNodeID.GNbId.GNBValue != "" {
			log.Tracef("targerRanID [%s]", targetRanNodeID.GNbId.GNBValue)
		}

		targetRan, ok := h.backend.Context().AmfRanFindByRanID(targetRanNodeID)
		if !ok {
			log.Warn("targetRan is nil")
		}

		h.sender.SendDownlinkRanConfigurationTransfer(targetRan, sONConfigurationTransferUL)
	}
}

func (h *ngapHandler) handleUplinkUEAssociatedNRPPATransport(ran *context.AmfRan, uplinkUEAssociatedNRPPaTransport *ngapType.UplinkUEAssociatedNRPPaTransport) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var routingID *ngapType.RoutingID
	var nRPPaPDU *ngapType.NRPPaPDU

	for _, ie := range uplinkUEAssociatedNRPPaTransport.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE amf_ngapid")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE ran_ngapid")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRoutingID: // reject
			routingID = ie.Value.RoutingID
			log.Trace("Decode IE routingID")
			if routingID == nil {
				log.Error("routingID is nil")
				return
			}
		case ngapType.ProtocolIEIDNRPPaPDU: // reject
			nRPPaPDU = ie.Value.NRPPaPDU
			log.Trace("Decode IE nRPPaPDU")
			if nRPPaPDU == nil {
				log.Error("nRPPaPDU is nil")
				return
			}
		}
	}

	ranUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		return
	}

	log.Tracef("RanUeNgapId[%d] AmfUeNgapId[%d]", ranUe.RanUeNgapId(), ranUe.AmfUeNgapId())
	log.Info("Handle Uplink UE Associated NRPPA Transpor")

	ranUe.SetRoutingId(hex.EncodeToString(routingID.Value))

	// TODO: Forward NRPPaPDU to LMF
}

func (h *ngapHandler) handleUplinkNonUEAssociatedNRPPATransport(ran *context.AmfRan, uplinkNonUEAssociatedNRPPATransport *ngapType.UplinkNonUEAssociatedNRPPaTransport) {
	var routingID *ngapType.RoutingID
	var nRPPaPDU *ngapType.NRPPaPDU

	log.Info("Handle Uplink Non UE Associated NRPPA Transport")

	for i := 0; i < len(uplinkNonUEAssociatedNRPPATransport.ProtocolIEs.List); i++ {
		ie := uplinkNonUEAssociatedNRPPATransport.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDRoutingID:
			routingID = ie.Value.RoutingID
			log.Trace("Decode IE RoutingID")

		case ngapType.ProtocolIEIDNRPPaPDU:
			nRPPaPDU = ie.Value.NRPPaPDU
			log.Trace("Decode IE NRPPaPDU")
		}
	}

	if routingID == nil {
		log.Error("RoutingID is nil")
		return
	}
	// Forward routingID to LMF
	// Described in (23.502 4.13.5.6)

	if nRPPaPDU == nil {
		log.Error("NRPPaPDU is nil")
		return
	}
	// TODO: Forward NRPPaPDU to LMF
}

func (h *ngapHandler) handleLocationReport(ran *context.AmfRan, locationReport *ngapType.LocationReport) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var userLocationInformation *ngapType.UserLocationInformation
	var uEPresenceInAreaOfInterestList *ngapType.UEPresenceInAreaOfInterestList
	var locationReportingRequestType *ngapType.LocationReportingRequestType

	for _, ie := range locationReport.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDUserLocationInformation: // ignore
			userLocationInformation = ie.Value.UserLocationInformation
			log.Trace("Decode IE userLocationInformation")
			if userLocationInformation == nil {
				log.Warn("userLocationInformation is nil")
			}
		case ngapType.ProtocolIEIDUEPresenceInAreaOfInterestList: // optional, ignore
			uEPresenceInAreaOfInterestList = ie.Value.UEPresenceInAreaOfInterestList
			log.Trace("Decode IE uEPresenceInAreaOfInterestList")
			if uEPresenceInAreaOfInterestList == nil {
				log.Warn("uEPresenceInAreaOfInterestList is nil [optional]")
			}
		case ngapType.ProtocolIEIDLocationReportingRequestType: // ignore
			locationReportingRequestType = ie.Value.LocationReportingRequestType
			log.Trace("Decode IE LocationReportingRequestType")
			if locationReportingRequestType == nil {
				log.Warn("LocationReportingRequestType is nil")
			}
		}
	}

	ranUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		return
	}
	log.Info("Handle Location Report")

	ranUe.UpdateLocation(userLocationInformation)

	log.Tracef("Report Area[%d]", locationReportingRequestType.ReportArea.Value)

	switch locationReportingRequestType.EventType.Value {
	case ngapType.EventTypePresentDirect:
		log.Trace("To report directly")

	case ngapType.EventTypePresentChangeOfServeCell:
		log.Trace("To report upon change of serving cell")

	case ngapType.EventTypePresentUePresenceInAreaOfInterest:
		log.Trace("To report UE presence in the area of interest")
		for _, uEPresenceInAreaOfInterestItem := range uEPresenceInAreaOfInterestList.List {
			uEPresence := uEPresenceInAreaOfInterestItem.UEPresence.Value
			referenceID := uEPresenceInAreaOfInterestItem.LocationReportingReferenceID.Value

			for _, AOIitem := range locationReportingRequestType.AreaOfInterestList.List {
				if referenceID == AOIitem.LocationReportingReferenceID.Value {
					log.Tracef("uEPresence[%d], presence AOI ReferenceID[%d]", uEPresence, referenceID)
				}
			}
		}

	case ngapType.EventTypePresentStopChangeOfServeCell:
		log.Trace("To stop reporting at change of serving cell")
		h.sender.SendLocationReportingControl(ranUe, nil, 0, locationReportingRequestType.EventType)
		// TODO: Clear location report

	case ngapType.EventTypePresentStopUePresenceInAreaOfInterest:
		log.Trace("To stop reporting UE presence in the area of interest")
		log.Tracef("ReferenceID To Be Cancelled[%d]",
			locationReportingRequestType.LocationReportingReferenceIDToBeCancelled.Value)
		// TODO: Clear location report

	case ngapType.EventTypePresentCancelLocationReportingForTheUe:
		log.Trace("To cancel location reporting for the UE")
		// TODO: Clear location report
	}
}

func (h *ngapHandler) handleUERadioCapabilityInfoIndication(ran *context.AmfRan, uERadioCapabilityInfoIndication *ngapType.UERadioCapabilityInfoIndication) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID

	var uERadioCapability *ngapType.UERadioCapability
	var uERadioCapabilityForPaging *ngapType.UERadioCapabilityForPaging

	for i := 0; i < len(uERadioCapabilityInfoIndication.ProtocolIEs.List); i++ {
		ie := uERadioCapabilityInfoIndication.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
				return
			}
		case ngapType.ProtocolIEIDUERadioCapability:
			uERadioCapability = ie.Value.UERadioCapability
			log.Trace("Decode IE UERadioCapability")
			if uERadioCapability == nil {
				log.Error("UERadioCapability is nil")
				return
			}
		case ngapType.ProtocolIEIDUERadioCapabilityForPaging:
			uERadioCapabilityForPaging = ie.Value.UERadioCapabilityForPaging
			log.Trace("Decode IE UERadioCapabilityForPaging")
			if uERadioCapabilityForPaging == nil {
				log.Error("UERadioCapabilityForPaging is nil")
				return
			}
		}
	}

	ranUe := ran.RanUeFindByRanUeNgapID(ran_ngapid.Value)
	if ranUe == nil {
		log.Errorf("No UE Context[RanUeNgapID: %d]", ran_ngapid.Value)
		return
	}
	log.Tracef("RanUeNgapID[%d] AmfUeNgapID[%d]", ranUe.RanUeNgapId(), ranUe.AmfUeNgapId())
	log.Info("Handle UE Radio Capability Info Indication")

	amfUe := ranUe.AmfUe()
	if amfUe == nil {
		log.Errorln("amfUe is nil")
		return
	}

	if uERadioCapability != nil {
		amfUe.UeRadioCapability = hex.EncodeToString(uERadioCapability.Value)
	}

	if uERadioCapabilityForPaging != nil {
		amfUe.UeRadioCapabilityForPaging = &context.UERadioCapabilityForPaging{}
		if uERadioCapabilityForPaging.UERadioCapabilityForPagingOfNR != nil {
			amfUe.UeRadioCapabilityForPaging.NR = hex.EncodeToString(
				uERadioCapabilityForPaging.UERadioCapabilityForPagingOfNR.Value)
		}
		if uERadioCapabilityForPaging.UERadioCapabilityForPagingOfEUTRA != nil {
			amfUe.UeRadioCapabilityForPaging.EUTRA = hex.EncodeToString(
				uERadioCapabilityForPaging.UERadioCapabilityForPagingOfEUTRA.Value)
		}
	}

	// TS 38.413 8.14.1.2/TS 23.502 4.2.8a step5/TS 23.501, clause 5.4.4.1.
	// send its most up to date UE Radio Capability information to the RAN in the N2 REQUEST message.
}

func (h *ngapHandler) handleErrorIndication(ran *context.AmfRan, errorIndication *ngapType.ErrorIndication) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var cause *ngapType.Cause
	var criticalityDiagnostics *ngapType.CriticalityDiagnostics

	for _, ie := range errorIndication.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
			if amf_ngapid == nil {
				log.Error("AmfUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")
			if ran_ngapid == nil {
				log.Error("RanUeNgapID is nil")
			}
		case ngapType.ProtocolIEIDCause:
			cause = ie.Value.Cause
			log.Trace("Decode IE Cause")
		case ngapType.ProtocolIEIDCriticalityDiagnostics:
			criticalityDiagnostics = ie.Value.CriticalityDiagnostics
			log.Trace("Decode IE CriticalityDiagnostics")
		}
	}

	log.Infof("Handle Error Indication: RAN_UE_NGAP_ID:%v AMF_UE_NGAP_ID:%v", ran_ngapid, amf_ngapid)

	if cause == nil && criticalityDiagnostics == nil {
		log.Error("[ErrorIndication] both Cause IE and CriticalityDiagnostics IE are nil, should have at least one")
		return
	}

	if cause != nil {
		printAndGetCause(ran, cause)
	}

	if criticalityDiagnostics != nil {
		printCriticalityDiagnostics(ran, criticalityDiagnostics)
	}

	// TODO: handle error based on cause/criticalityDiagnostics
}

func (h *ngapHandler) handleCellTrafficTrace(ran *context.AmfRan, cellTrafficTrace *ngapType.CellTrafficTrace) {
	var amf_ngapid *ngapType.AMFUENGAPID
	var ran_ngapid *ngapType.RANUENGAPID
	var nGRANTraceID *ngapType.NGRANTraceID
	var nGRANCGI *ngapType.NGRANCGI
	var traceCollectionEntityIPAddress *ngapType.TransportLayerAddress

	var ranUe *context.RanUe

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	log.Info("Handle Cell Traffic Trace")

	for _, ie := range cellTrafficTrace.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID: // reject
			amf_ngapid = ie.Value.AMFUENGAPID
			log.Trace("Decode IE AmfUeNgapID")
		case ngapType.ProtocolIEIDRANUENGAPID: // reject
			ran_ngapid = ie.Value.RANUENGAPID
			log.Trace("Decode IE RanUeNgapID")

		case ngapType.ProtocolIEIDNGRANTraceID: // ignore
			nGRANTraceID = ie.Value.NGRANTraceID
			log.Trace("Decode IE NGRANTraceID")
		case ngapType.ProtocolIEIDNGRANCGI: // ignore
			nGRANCGI = ie.Value.NGRANCGI
			log.Trace("Decode IE NGRANCGI")
		case ngapType.ProtocolIEIDTraceCollectionEntityIPAddress: // ignore
			traceCollectionEntityIPAddress = ie.Value.TraceCollectionEntityIPAddress
			log.Trace("Decode IE TraceCollectionEntityIPAddress")
		}
	}
	if amf_ngapid == nil {
		log.Error("AmfUeNgapID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDAMFUENGAPID,
			ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}
	if ran_ngapid == nil {
		log.Error("RanUeNgapID is nil")
		item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ngapType.ProtocolIEIDRANUENGAPID,
			ngapType.TypeOfErrorPresentMissing)
		iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		procedureCode := ngapType.ProcedureCodeCellTrafficTrace
		triggeringMessage := ngapType.TriggeringMessagePresentInitiatingMessage
		procedureCriticality := ngapType.CriticalityPresentIgnore
		criticalityDiagnostics := buildCriticalityDiagnostics(&procedureCode, &triggeringMessage, &procedureCriticality,
			&iesCriticalityDiagnostics)
		h.sender.SendErrorIndication(ran, amf_ngapid, ran_ngapid, nil, &criticalityDiagnostics)
		return
	}

	if amf_ngapid != nil {
		ranUe = h.backend.Context().RanUeFindByAmfUeNgapID(amf_ngapid.Value)
		if ranUe == nil {
			log.Errorf("No UE Context[AmfUeNgapID: %d]", amf_ngapid.Value)
			cause := ngapType.Cause{
				Present: ngapType.CausePresentRadioNetwork,
				RadioNetwork: &ngapType.CauseRadioNetwork{
					Value: ngapType.CauseRadioNetworkPresentUnknownLocalUENGAPID,
				},
			}
			h.sender.SendErrorIndication(ran, amf_ngapid, ran_ngapid, &cause, nil)
			return
		}
	}

	log.Debugf("UE: AmfUeNgapID[%d], RanUeNgapID[%d]", ranUe.AmfUeNgapId(), ranUe.RanUeNgapId())

	ranUe.SetTrsr(hex.EncodeToString(nGRANTraceID.Value[6:]))

	log.Tracef("TRSR[%s]", ranUe.Trsr())

	switch nGRANCGI.Present {
	case ngapType.NGRANCGIPresentNRCGI:
		plmnID := ngapConvert.PlmnIdToModels(nGRANCGI.NRCGI.PLMNIdentity)
		cellID := ngapConvert.BitStringToHex(&nGRANCGI.NRCGI.NRCellIdentity.Value)
		log.Debugf("NRCGI[plmn: %s, cellID: %s]", plmnID, cellID)
	case ngapType.NGRANCGIPresentEUTRACGI:
		plmnID := ngapConvert.PlmnIdToModels(nGRANCGI.EUTRACGI.PLMNIdentity)
		cellID := ngapConvert.BitStringToHex(&nGRANCGI.EUTRACGI.EUTRACellIdentity.Value)
		log.Debugf("EUTRACGI[plmn: %s, cellID: %s]", plmnID, cellID)
	}

	tceIpv4, tceIpv6 := ngapConvert.IPAddressToString(*traceCollectionEntityIPAddress)
	if tceIpv4 != "" {
		log.Debugf("TCE IP Address[v4: %s]", tceIpv4)
	}
	if tceIpv6 != "" {
		log.Debugf("TCE IP Address[v6: %s]", tceIpv6)
	}

	// TODO: TS 32.422 4.2.2.10
	// When AMF receives this new NG signalling message containing the Trace Recording Session Reference (TRSR)
	// and Trace Reference (TR), the AMF shall look up the SUPI/IMEI(SV) of the given call from its database and
	// shall send the SUPI/IMEI(SV) numbers together with the Trace Recording Session Reference and Trace Reference
	// to the Trace Collection Entity.
}

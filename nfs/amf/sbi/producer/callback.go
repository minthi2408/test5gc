package producer

import (
	"fmt"
	"net/http"
	"strconv"

//	"github.com/mohae/deepcopy"

	"etri5gc/nfs/amf/context"

//	"github.com/free5gc/amf/internal/util"
//	"github.com/free5gc/nas/nasConvert"
//	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)
//NOTE: tungtq: some parts which relate to ngap and nas has been commented out, need more time

func (h *Handler) HandleSmContextStatusNotify(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.ProducerLog.Infoln("[AMF] Handle SmContext Status Notify")

	guti := request.Params["guti"]
	pduIdStr := request.Params["pduSessionId"]
	if pduId, err := strconv.Atoi(pduIdStr); err == nil {
		notification := request.Body.(models.SmContextStatusNotification)
	 	if prob:= h.SmContextStatusNotifyProcedure(guti, int32(pduId), notification); prob != nil {
			return httpwrapper.NewResponse(int(prob.Status), nil, prob)
		}
	}
	return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
}

func (h *Handler) SmContextStatusNotifyProcedure(guti string, pduId int32, notification models.SmContextStatusNotification) *models.ProblemDetails {
	if ue, ok := h.amf().AmfUeFindByGuti(guti); !ok {
		return &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
			Detail: fmt.Sprintf("Guti[%s] Not Found", guti),
		}
	} else {

		if smContext, ok := ue.SmContextFindByPDUSessionID(pduId); ! ok {
			return &models.ProblemDetails{
				Status: http.StatusNotFound,
				Cause:  "CONTEXT_NOT_FOUND",
				Detail: fmt.Sprintf("PDUSessionID[%d] Not Found", pduId),
			}
		} else {
			if notification.StatusInfo.ResourceStatus != models.ResourceStatus_RELEASED {
				return &models.ProblemDetails{
					Status: http.StatusBadRequest,
					Cause:  "INVALID_MSG_FORMAT",
					InvalidParams: []models.InvalidParam{
						{Param: "StatusInfo.ResourceStatus", Reason: "invalid value"},
					},
				}
			}

			//ue.ProducerLog.Debugf("Release PDU Session[%d] (Cause: %s)", pduSessionID,
			//	smContextStatusNotification.StatusInfo.Cause)

			if smContext.PduSessionIDDuplicated() {
				smContext.SetDuplicatedPduSessionID(false)
				go ResumePduSession(ue, smContext)
				
			} else {
				ue.DeleteSmContext(pduId)
			}
		}
	}
	return nil
}

func ResumePduSession(ue *context.AmfUe, sm *context.SmContext) {
	/*
	TungTQ : I do not understand this procedure yet, need more time 

	//ue.ProducerLog.Debugf("Resume establishing PDU Session[%d]", pduSessionID)
	var (
		snssai    models.Snssai
		dnn       string
		smMessage []byte
	)
	smMessage = sm.ULNASTransport().GetPayloadContainerContents()

	//get snssai
	if sm.ULNASTransport().SNSSAI != nil {
		snssai = nasConvert.SnssaiToModels(sm.ULNASTransport().SNSSAI)
	} else {
		if allowedNssai, ok := ue.AllowedNssai[smContext.AccessType()]; ok {
			snssai = *allowedNssai[0].AllowedSnssai
		} else {
			//ue.GmmLog.Error("Ue doesn't have allowedNssai")
			return
		}
	}

	//get dnn
	if sm.ULNASTransport().DNN != nil {
		dnn = sm.ULNASTransport().DNN.GetDNN()
	} else {
		if ue.SmfSelectionData != nil {
			snssaiStr := util.SnssaiModelsToHex(snssai)
			if snssaiInfo, ok := ue.SmfSelectionData.SubscribedSnssaiInfos[snssaiStr]; ok {
				for _, dnnInfo := range snssaiInfo.DnnInfos {
					if dnnInfo.DefaultDnnIndicator {
						dnn = dnnInfo.Dnn
					}
				}
			} else {
				// user's subscription context obtained from UDM does not contain the default DNN for the,
				// S-NSSAI, the AMF shall use a locally configured DNN as the DNN
				dnn = "internet"
			}
		}
	}
	//select smf
	newSmContext, cause, err := consumer.SelectSmf(ue, smContext.AccessType(), pduSessionID, snssai, dnn)
	if err != nil {
		//logger.CallbackLog.Error(err)
		gmm_message.SendDLNASTransport(ue.RanUe[smContext.AccessType()],
			nasMessage.PayloadContainerTypeN1SMInfo,
			smContext.ULNASTransport().GetPayloadContainerContents(), pduSessionID, cause, nil, 0)
		return
	}

	// send CreateSmContextRequest to the selected SMF
	response, smContextRef, errResponse, problemDetail, err := consumer.SendCreateSmContextRequest(
		ue, newSmContext, nil, smMessage)
	if response != nil {
		newSmContext.SetSmContextRef(smContextRef)
		newSmContext.SetUserLocation(deepcopy.Copy(ue.Location).(models.UserLocation))
		ue.GmmLog.Infof("create smContext[pduSessionID: %d] Success", pduSessionID)
		ue.StoreSmContext(pduSessionID, newSmContext)
		// TODO: handle response(response N2SmInfo to RAN if exists)
	} else if errResponse != nil {
		ue.ProducerLog.Warnf("PDU Session Establishment Request is rejected by SMF[pduSessionId:%d]\n", pduSessionID)
		gmm_message.SendDLNASTransport(ue.RanUe[smContext.AccessType()],
			nasMessage.PayloadContainerTypeN1SMInfo, errResponse.BinaryDataN1SmMessage, pduSessionID, 0, nil, 0)
	} else if err != nil {
		ue.ProducerLog.Errorf("Failed to Create smContext[pduSessionID: %d], Error[%s]\n", pduSessionID, err.Error())
	} else {
		ue.ProducerLog.Errorf("Failed to Create smContext[pduSessionID: %d], Error[%v]\n", pduSessionID, problemDetail)
	}
	sm.DeleteULNASTransport()
}
*/
}


func (h *Handler) HandleAmPolicyControlUpdateNotifyUpdate(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.ProducerLog.Infoln("Handle AM Policy Control Update Notify [Policy update notification]")

	polId := request.Params["polAssoId"]
	polUpdate := request.Body.(models.PolicyUpdate)

	if prob:= h.doAmPolicyControlUpdateNotifyUpdate(polId, polUpdate); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
}

func (h *Handler) doAmPolicyControlUpdateNotifyUpdate(polId string, polUpdate models.PolicyUpdate) *models.ProblemDetails {
	ue, ok := h.amf().AmfUeFindByPolicyAssociationID(polId)
	if !ok {
		return &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
			Detail: fmt.Sprintf("Policy Association ID[%s] Not Found", polId),
		}
	}
	pcfinfo := ue.GetPcfInfo()
	pcfinfo.AmPolicyAssociation.Triggers = polUpdate.Triggers
	pcfinfo.RequestTriggerLocationChange = false

	for _, trigger := range polUpdate.Triggers {
		if trigger == models.RequestTrigger_LOC_CH {
			pcfinfo.RequestTriggerLocationChange = true
		}
		//if trigger == models.RequestTrigger_PRA_CH {
		// TODO: Presence Reporting Area handling (TS 23.503 6.1.2.5, TS 23.501 5.6.11)
		//}
	}

	if polUpdate.ServAreaRes != nil {
		pcfinfo.AmPolicyAssociation.ServAreaRes = polUpdate.ServAreaRes
	}

	if polUpdate.Rfsp != 0 {
		pcfinfo.AmPolicyAssociation.Rfsp = polUpdate.Rfsp
	}

	// use go routine to write response first to ensure the order of the procedure
	go func() {
		// UE is CM-Connected State
		if ue.CmConnect(models.AccessType__3_GPP_ACCESS) {
			h.nas.SendConfigurationUpdateCommand(ue, models.AccessType__3_GPP_ACCESS, nil)
			// UE is CM-IDLE => paging
		} else {
			message, err := h.nas.BuildConfigurationUpdateCommand(ue, models.AccessType__3_GPP_ACCESS, nil)
			if err != nil {
			//	logger.GmmLog.Errorf("Build Configuration Update Command Failed : %s", err.Error())
				return
			}
			pcfinfo.ConfigurationUpdateMessage = message
			ue.SetOnGoing(models.AccessType__3_GPP_ACCESS, &context.OnGoing{
				Procedure: context.OnGoingProcedurePaging,
			})

			pkg, err := h.ngap.BuildPaging(ue, nil, false)
			if err != nil {
				//logger.NgapLog.Errorf("Build Paging failed : %s", err.Error())
				return
			}
			h.ngap.SendPaging(ue, pkg)
		}
	}()
	return nil
}

// TS 29.507 4.2.4.3
func (h *Handler) HandleAmPolicyControlUpdateNotifyTerminate(request *httpwrapper.Request) *httpwrapper.Response {
//	logger.ProducerLog.Infoln("Handle AM Policy Control Update Notify [Request for termination of the policy association]")

	polAssoID := request.Params["polAssoId"]
	terminationNotification := request.Body.(models.TerminationNotification)

	if prob := h.doAmPolicyControlUpdateNotifyTerminate(polAssoID, terminationNotification); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
}

func (h *Handler) doAmPolicyControlUpdateNotifyTerminate(polAssoID string,
	terminationNotification models.TerminationNotification) *models.ProblemDetails {

	ue, ok := h.amf().AmfUeFindByPolicyAssociationID(polAssoID)
	if !ok {
		return &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
			Detail: fmt.Sprintf("Policy Association ID[%s] Not Found", polAssoID),
		}
	}

//	logger.CallbackLog.Infof("Cause of AM Policy termination[%+v]", terminationNotification.Cause)

	// use go routine to write response first to ensure the order of the procedure
	go func() {
		if prob, err := h.backend.Consumer().Pcf().AMPolicyControlDelete(ue); prob != nil {
			//logger.ProducerLog.Errorf("AM Policy Control Delete Failed Problem[%+v]", prob)
		} else if err != nil {
			//logger.ProducerLog.Errorf("AM Policy Control Delete Error[%v]", err.Error())
		}
	}()
	return nil
}

// TS 23.502 4.2.2.2.3 Registration with AMF re-allocation
func (h *Handler) HandleN1MessageNotify(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.ProducerLog.Infoln("[AMF] Handle N1 Message Notify")

	n1MessageNotify := request.Body.(models.N1MessageNotify)

	if prob := h.doN1MessageNotify(n1MessageNotify); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
}

func (h *Handler) doN1MessageNotify(n1MessageNotify models.N1MessageNotify) *models.ProblemDetails {
	//logger.ProducerLog.Debugf("n1MessageNotify: %+v", n1MessageNotify)

	registrationCtxtContainer := n1MessageNotify.JsonData.RegistrationCtxtContainer
	if registrationCtxtContainer.UeContext == nil {
		return &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "MANDATORY_IE_MISSING", // Defined in TS 29.500 5.2.7.2
			Detail: "Missing IE [UeContext] in RegistrationCtxtContainer",
		}
	}

	ran, ok := h.amf().AmfRanFindByRanID(*registrationCtxtContainer.RanNodeId)
	if !ok {
		return &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "MANDATORY_IE_INCORRECT",
			Detail: fmt.Sprintf("Can not find RAN[RanId: %+v]", *registrationCtxtContainer.RanNodeId),
		}
	}

	//TODO: tungtq - should move this part to the AmfContext module
	go func() {
		var amfUe *context.AmfUe
		ueContext := registrationCtxtContainer.UeContext
		if ueContext.Supi != "" {
			amfUe = h.amf().NewAmfUe(ueContext.Supi)
		} else {
			amfUe = h.amf().NewAmfUe("")
		}
		amfUe.CopyDataFromUeContextModel(*ueContext)

		ranUe := ran.RanUeFindByRanUeNgapID(int64(registrationCtxtContainer.AnN2ApId))

		ranUe.SetLocation(*registrationCtxtContainer.UserLocation)
		amfUe.GetLocInfo().Location = *registrationCtxtContainer.UserLocation
		ranUe.SetUeContextRequest(registrationCtxtContainer.UeContextRequest)
		ranUe.OldAmfName = registrationCtxtContainer.InitialAmfName

		if registrationCtxtContainer.AllowedNssai != nil {
			allowedNssai := registrationCtxtContainer.AllowedNssai
			amfUe.GetNssfInfo().AllowedNssai[allowedNssai.AccessType] = allowedNssai.AllowedSnssaiList
		}

		if len(registrationCtxtContainer.ConfiguredNssai) > 0 {
			amfUe.GetNssfInfo().ConfiguredNssai = registrationCtxtContainer.ConfiguredNssai
		}

		amfUe.AttachRanUe(ranUe)
		h.nas.HandleNAS(ranUe, ngapType.ProcedureCodeInitialUEMessage, n1MessageNotify.BinaryDataN1Message)
	}()
	return nil
}

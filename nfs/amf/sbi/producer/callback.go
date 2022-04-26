package producer

import (
	"fmt"
	"net/http"
	"strconv"

//	"github.com/mohae/deepcopy"

	"etri5gc/nfs/amf/context"

//	gmm_message "github.com/free5gc/amf/internal/gmm/message"
//	"github.com/free5gc/amf/internal/nas"
//	ngap_message "github.com/free5gc/amf/internal/ngap/message"
//	"github.com/free5gc/amf/internal/sbi/consumer"
//	"github.com/free5gc/amf/internal/util"
//	"github.com/free5gc/nas/nasConvert"
//	"github.com/free5gc/nas/nasMessage"
//	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

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
	if ue, ok := h.backend.Context().AmfUeFindByGuti(guti); !ok {
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
	TungTQ : need more time

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
	/*
	logger.ProducerLog.Infoln("Handle AM Policy Control Update Notify [Policy update notification]")

	polAssoID := request.Params["polAssoId"]
	policyUpdate := request.Body.(models.PolicyUpdate)

	problemDetails := AmPolicyControlUpdateNotifyUpdateProcedure(polAssoID, policyUpdate)

	if problemDetails != nil {
		return httpwrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
	*/
	return nil
}
/*
func AmPolicyControlUpdateNotifyUpdateProcedure(polAssoID string,
	policyUpdate models.PolicyUpdate) *models.ProblemDetails {
	amfSelf := context.AMF_Self()

	ue, ok := amfSelf.AmfUeFindByPolicyAssociationID(polAssoID)
	if !ok {
		problemDetails := &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
			Detail: fmt.Sprintf("Policy Association ID[%s] Not Found", polAssoID),
		}
		return problemDetails
	}

	ue.AmPolicyAssociation.Triggers = policyUpdate.Triggers
	ue.RequestTriggerLocationChange = false

	for _, trigger := range policyUpdate.Triggers {
		if trigger == models.RequestTrigger_LOC_CH {
			ue.RequestTriggerLocationChange = true
		}
		//if trigger == models.RequestTrigger_PRA_CH {
		// TODO: Presence Reporting Area handling (TS 23.503 6.1.2.5, TS 23.501 5.6.11)
		//}
	}

	if policyUpdate.ServAreaRes != nil {
		ue.AmPolicyAssociation.ServAreaRes = policyUpdate.ServAreaRes
	}

	if policyUpdate.Rfsp != 0 {
		ue.AmPolicyAssociation.Rfsp = policyUpdate.Rfsp
	}

	if ue != nil {
		// use go routine to write response first to ensure the order of the procedure
		go func() {
			// UE is CM-Connected State
			if ue.CmConnect(models.AccessType__3_GPP_ACCESS) {
				gmm_message.SendConfigurationUpdateCommand(ue, models.AccessType__3_GPP_ACCESS, nil)
				// UE is CM-IDLE => paging
			} else {
				message, err := gmm_message.BuildConfigurationUpdateCommand(ue, models.AccessType__3_GPP_ACCESS, nil)
				if err != nil {
					logger.GmmLog.Errorf("Build Configuration Update Command Failed : %s", err.Error())
					return
				}

				ue.ConfigurationUpdateMessage = message
				ue.SetOnGoing(models.AccessType__3_GPP_ACCESS, &context.OnGoing{
					Procedure: context.OnGoingProcedurePaging,
				})

				pkg, err := ngap_message.BuildPaging(ue, nil, false)
				if err != nil {
					logger.NgapLog.Errorf("Build Paging failed : %s", err.Error())
					return
				}
				ngap_message.SendPaging(ue, pkg)
			}
		}()
	}
	return nil
}
*/
// TS 29.507 4.2.4.3
func (h *Handler) HandleAmPolicyControlUpdateNotifyTerminate(request *httpwrapper.Request) *httpwrapper.Response {
	/*
	logger.ProducerLog.Infoln("Handle AM Policy Control Update Notify [Request for termination of the policy association]")

	polAssoID := request.Params["polAssoId"]
	terminationNotification := request.Body.(models.TerminationNotification)

	problemDetails := AmPolicyControlUpdateNotifyTerminateProcedure(polAssoID, terminationNotification)
	if problemDetails != nil {
		return httpwrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
	*/
	return nil
}

/*
func AmPolicyControlUpdateNotifyTerminateProcedure(polAssoID string,
	terminationNotification models.TerminationNotification) *models.ProblemDetails {
	amfSelf := context.AMF_Self()

	ue, ok := amfSelf.AmfUeFindByPolicyAssociationID(polAssoID)
	if !ok {
		problemDetails := &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
			Detail: fmt.Sprintf("Policy Association ID[%s] Not Found", polAssoID),
		}
		return problemDetails
	}

	logger.CallbackLog.Infof("Cause of AM Policy termination[%+v]", terminationNotification.Cause)

	// use go routine to write response first to ensure the order of the procedure
	go func() {
		problem, err := consumer.AMPolicyControlDelete(ue)
		if problem != nil {
			logger.ProducerLog.Errorf("AM Policy Control Delete Failed Problem[%+v]", problem)
		} else if err != nil {
			logger.ProducerLog.Errorf("AM Policy Control Delete Error[%v]", err.Error())
		}
	}()
	return nil
}
*/

// TS 23.502 4.2.2.2.3 Registration with AMF re-allocation
func (h *Handler) HandleN1MessageNotify(request *httpwrapper.Request) *httpwrapper.Response {
	/*
	logger.ProducerLog.Infoln("[AMF] Handle N1 Message Notify")

	n1MessageNotify := request.Body.(models.N1MessageNotify)

	problemDetails := N1MessageNotifyProcedure(n1MessageNotify)
	if problemDetails != nil {
		return httpwrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
	*/
	return nil
}

/*
func N1MessageNotifyProcedure(n1MessageNotify models.N1MessageNotify) *models.ProblemDetails {
	logger.ProducerLog.Debugf("n1MessageNotify: %+v", n1MessageNotify)

	amfSelf := context.AMF_Self()

	registrationCtxtContainer := n1MessageNotify.JsonData.RegistrationCtxtContainer
	if registrationCtxtContainer.UeContext == nil {
		problemDetails := &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "MANDATORY_IE_MISSING", // Defined in TS 29.500 5.2.7.2
			Detail: "Missing IE [UeContext] in RegistrationCtxtContainer",
		}
		return problemDetails
	}

	ran, ok := amfSelf.AmfRanFindByRanID(*registrationCtxtContainer.RanNodeId)
	if !ok {
		problemDetails := &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "MANDATORY_IE_INCORRECT",
			Detail: fmt.Sprintf("Can not find RAN[RanId: %+v]", *registrationCtxtContainer.RanNodeId),
		}
		return problemDetails
	}

	go func() {
		var amfUe *context.AmfUe
		ueContext := registrationCtxtContainer.UeContext
		if ueContext.Supi != "" {
			amfUe = amfSelf.NewAmfUe(ueContext.Supi)
		} else {
			amfUe = amfSelf.NewAmfUe("")
		}
		amfUe.CopyDataFromUeContextModel(*ueContext)

		ranUe := ran.RanUeFindByRanUeNgapID(int64(registrationCtxtContainer.AnN2ApId))

		ranUe.Location = *registrationCtxtContainer.UserLocation
		amfUe.Location = *registrationCtxtContainer.UserLocation
		ranUe.UeContextRequest = registrationCtxtContainer.UeContextRequest
		ranUe.OldAmfName = registrationCtxtContainer.InitialAmfName

		if registrationCtxtContainer.AllowedNssai != nil {
			allowedNssai := registrationCtxtContainer.AllowedNssai
			amfUe.AllowedNssai[allowedNssai.AccessType] = allowedNssai.AllowedSnssaiList
		}

		if len(registrationCtxtContainer.ConfiguredNssai) > 0 {
			amfUe.ConfiguredNssai = registrationCtxtContainer.ConfiguredNssai
		}

		amfUe.AttachRanUe(ranUe)

		nas.HandleNAS(ranUe, ngapType.ProcedureCodeInitialUEMessage, n1MessageNotify.BinaryDataN1Message)
	}()
	return nil
}
*/

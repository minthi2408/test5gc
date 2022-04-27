package producer

import (
	"net/http"
	"strings"

	"etri5gc/nfs/amf/context"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)
//NOTE tuntq: removing AmfUe has been commented, need to handle it


// TS 29.518 5.2.2.2.3
func (h *Handler) HandleCreateUEContextRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.CommLog.Infof("Handle Create UE Context Request")

	createUeContextRequest := request.Body.(models.CreateUeContextRequest)
	ueContextID := request.Params["ueContextId"]

	if createUeContextResponse, err := h.doCreateUEContext(ueContextID, createUeContextRequest); err != nil {
		return httpwrapper.NewResponse(int(err.Error.Status), nil, err)
	} else {
		return httpwrapper.NewResponse(http.StatusCreated, nil, createUeContextResponse)
	}
	return nil
}
func (h *Handler) doCreateUEContext(ueContextID string, createUeContextRequest models.CreateUeContextRequest) (
	*models.CreateUeContextResponse, *models.UeContextCreateError) {
	amf := h.amf() 
	ueContextCreateData := createUeContextRequest.JsonData

	if ueContextCreateData.UeContext == nil || ueContextCreateData.TargetId == nil ||
		ueContextCreateData.PduSessionList == nil || ueContextCreateData.SourceToTargetData == nil ||
		ueContextCreateData.N2NotifyUri == "" {
		return nil, &models.UeContextCreateError{
			Error: &models.ProblemDetails{
				Status: http.StatusForbidden,
				Cause:  "HANDOVER_FAILURE",
			},
		}
	}
	// create the UE context in target amf
	amf.NewAmfUeByReq(ueContextID, ueContextCreateData)

	createUeContextResponse := &models.CreateUeContextResponse {
		JsonData: &models.UeContextCreatedData{
				UeContext: &models.UeContext{
					Supi: ueContextCreateData.UeContext.Supi,
				},
			},
		}

	// response.JsonData.TargetToSourceData =
	// ue.N1N2Message[ueContextId].Request.JsonData.N2InfoContainer.SmInfo.N2InfoContent
	createUeContextResponse.JsonData.PduSessionList = ueContextCreateData.PduSessionList
	createUeContextResponse.JsonData.PcfReselectedInd = false
	// TODO: When  Target AMF selects a nw PCF for AM policy, set the flag to true.

	//response.UeContext = ueContextCreateData.UeContext
	//response.TargetToSourceData = ue.N1N2Message[amf.Uri].Request.JsonData.N2InfoContainer.SmInfo.N2InfoContent
	//response.PduSessionList = ueContextCreateData.PduSessionList
	//response.PcfReselectedInd = false // TODO:When  Target AMF selects a nw PCF for AM policy, set the flag to true.
	//

	// return httpwrapper.NewResponse(http.StatusCreated, nil, createUeContextResponse)
	return createUeContextResponse, nil
}

// TS 29.518 5.2.2.2.4
func (h *Handler) HandleReleaseUEContextRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.CommLog.Info("Handle Release UE Context Request")

	ueContextRelease := request.Body.(models.UeContextRelease)
	ueContextID := request.Params["ueContextId"]

	problemDetails := h.doReleaseUEContext(ueContextID, ueContextRelease)
	if problemDetails != nil {
		return httpwrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
	return nil
}

func (h *Handler) doReleaseUEContext(ueContextID string, ueContextRelease models.UeContextRelease) *models.ProblemDetails {
	// TODO: UE is emergency registered and the SUPI is not authenticated
	if ueContextRelease.Supi != "" {
		//logger.GmmLog.Warnf("Emergency registered UE is not supported.")
		return &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "UNSPECIFIED",
		}
	}

	if ueContextRelease.NgapCause == nil {
		return &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "MANDATORY_IE_MISSING",
		}
	}

	//logger.CommLog.Debugf("Release UE Context NGAP cause: %+v", ueContextRelease.NgapCause)

	if /*ue*/_, ok := h.amf().AmfUeFindByUeContextID(ueContextID); ok {
		//TODO: tungtq - a bad design by free5gc - let move to a different place (amf context)
		//gmm_common.RemoveAmfUe(ue)
	} else {
		return &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
	}

	return nil
}
// TS 29.518 5.2.2.2.1
func (h *Handler) HandleUEContextTransferRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.CommLog.Info("Handle UE Context Transfer Request")

	req := request.Body.(models.UeContextTransferRequest)
	ueId := request.Params["ueContextId"]

	if res, prob := h.doUEContextTransfer(ueId, req); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, res)
	}
}

func (h *Handler) doUEContextTransfer(ueId string, req models.UeContextTransferRequest) (
	*models.UeContextTransferResponse, *models.ProblemDetails) {

	amf := h.amf() 

	if req.JsonData == nil {
		return nil, &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "MANDATORY_IE_MISSING",
		}
	}

	reqdat := req.JsonData

	if reqdat.AccessType == "" || reqdat.Reason == "" {
		return nil, &models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "MANDATORY_IE_MISSING",
		}
	}

	ue, ok := amf.AmfUeFindByUeContextID(ueId)
	if !ok {
		return nil, &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
	}

	res := &models.UeContextTransferResponse {
		JsonData: &models.UeContextTransferRspData{},
	}

	if prob := ue.BuildContextTransferResponse(reqdat, res); prob != nil {
		return nil, prob
	}
	return res, nil

	
}



// TS 29.518 5.2.2.6
func (h *Handler) HandleAssignEbiDataRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.CommLog.Info("Handle Assign Ebi Data Request")

	ebidat := request.Body.(models.AssignEbiData)
	ueId := request.Params["ueContextId"]

	if assignedEbiData, assignEbiError, prob := h.doAssignEbiData(ueId, ebidat); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else if assignEbiError != nil {
		return httpwrapper.NewResponse(int(assignEbiError.Error.Status), nil, assignEbiError)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, assignedEbiData)
	}
}

func (h *Handler) doAssignEbiData(ueId string, ebidat models.AssignEbiData) (
	*models.AssignedEbiData, *models.AssignEbiError, *models.ProblemDetails) {

	if ue, ok := h.amf().AmfUeFindByUeContextID(ueId); !ok {
		return nil, nil, &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
	} else {

		// TODO: AssignEbiError not used, check it!
		if _, ok := ue.SmContextFindByPDUSessionID(ebidat.PduSessionId); ok {
			var assignedEbiData models.AssignedEbiData
			assignedEbiData.PduSessionId = ebidat.PduSessionId
			return &assignedEbiData, nil, nil
		} else {
			//logger.ProducerLog.Errorln("ue.SmContextList is nil")
			return nil, nil, nil
		}
	}
}

// TS 29.518 5.2.2.2.2
func (h *Handler) HandleRegistrationStatusUpdateRequest(request *httpwrapper.Request) *httpwrapper.Response {
//	logger.CommLog.Info("Handle Registration Status Update Request")

	reqdat := request.Body.(models.UeRegStatusUpdateReqData)
	ueId := request.Params["ueContextId"]

	if resdat, prob := h.doRegistrationStatusUpdate(ueId, reqdat); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, resdat)
	}
}

func (h *Handler) doRegistrationStatusUpdate(ueId string, reqdat models.UeRegStatusUpdateReqData) (
	*models.UeRegStatusUpdateRspData, *models.ProblemDetails) {

	// ueId must be a 5g GUTI (TS 29.518 6.1.3.2.4.5.1)
	if !strings.HasPrefix(ueId, "5g-guti") {
		return nil, &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "UNSPECIFIED",
		}
	}

	if ue, ok := h.amf().AmfUeFindByUeContextID(ueId); !ok {
		return nil, &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
	} else {

		resdat := new(models.UeRegStatusUpdateRspData)

		if reqdat.TransferStatus == models.UeContextTransferStatus_TRANSFERRED {
			// remove the individual ueContext resource and release any PDU session(s)
			for _, pduSessionId := range reqdat.ToReleaseSessionList {
				cause := models.Cause_REL_DUE_TO_SLICE_NOT_AVAILABLE
				causeAll := &context.CauseAll{
					Cause: &cause,
				}

				smContext, ok := ue.SmContextFindByPDUSessionID(pduSessionId)
				if !ok {
					//ue.ProducerLog.Errorf("SmContext[PDU Session ID:%d] not found", pduSessionId)
				}

				problem, err := h.backend.Consumer().Smf().SendReleaseSmContextRequest(ue, smContext, causeAll, "", nil)
				if problem != nil {
					//logger.GmmLog.Errorf("Release SmContext[pduSessionId: %d] Failed Problem[%+v]", pduSessionId, problem)
				} else if err != nil {
					//logger.GmmLog.Errorf("Release SmContext[pduSessionId: %d] Error[%v]", pduSessionId, err.Error())
				}
			}

			if reqdat.PcfReselectedInd {
				problem, err := h.backend.Consumer().Pcf().AMPolicyControlDelete(ue)
				if problem != nil {
			//		logger.GmmLog.Errorf("AM Policy Control Delete Failed Problem[%+v]", problem)
				} else if err != nil {
			//		logger.GmmLog.Errorf("AM Policy Control Delete Error[%v]", err.Error())
				}
			}
			//TODO: tungtq: it seems it is a bad idea to remove AmfUe inside gmm_comm (a bad design from free5gc)
			//gmm_common.RemoveAmfUe(ue)
		} else {
			// NOT_TRANSFERRED
			//logger.CommLog.Debug("[AMF] RegistrationStatusUpdate: NOT_TRANSFERRED")
		}

		resdat.RegStatusTransferComplete = true
		return resdat, nil
	}
}

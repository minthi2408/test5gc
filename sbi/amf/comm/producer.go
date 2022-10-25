/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package comm

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
)

//sbi producer handler for AMFStatusChangeSubscribeModfy
func OnAMFStatusChangeSubscribeModfy(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	subscriptionId := ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		//subscriptionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subscriptionId is required",
		}))
		return
	}

	var input models.SubscriptionData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SubscriptionData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleAMFStatusChangeSubscribeModfy(subscriptionId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for AMFStatusChangeUnSubscribe
func OnAMFStatusChangeUnSubscribe(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	subscriptionId := ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		//subscriptionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subscriptionId is required",
		}))
		return
	}

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.COMM_HandleAMFStatusChangeUnSubscribe(subscriptionId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}

//sbi producer handler for CancelRelocateUEContext
func OnCancelRelocateUEContext(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.CancelRelocateUEContextRequest

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.COMM_HandleCancelRelocateUEContext(ueContextId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}

//sbi producer handler for CreateUEContext
func OnCreateUEContext(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.CreateUEContextRequest

	var apierr *sbi.ApiError
	var successCode int32
	var result models.UeContextCreatedData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleCreateUEContext(ueContextId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for EBIAssignment
func OnEBIAssignment(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.AssignEbiData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.AssignedEbiData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleEBIAssignment(ueContextId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for RegistrationStatusUpdate
func OnRegistrationStatusUpdate(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.UeRegStatusUpdateReqData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.UeRegStatusUpdateRspData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleRegistrationStatusUpdate(ueContextId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for ReleaseUEContext
func OnReleaseUEContext(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.UEContextRelease

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.COMM_HandleReleaseUEContext(ueContextId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}

//sbi producer handler for RelocateUEContext
func OnRelocateUEContext(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.RelocateUEContextRequest

	var apierr *sbi.ApiError
	var successCode int32
	var result models.UeContextRelocatedData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleRelocateUEContext(ueContextId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for UEContextTransfer
func OnUEContextTransfer(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.UeContextTransferReqData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.UeContextTransferRspData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleUEContextTransfer(ueContextId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for N1N2MessageUnSubscribe
func OnN1N2MessageUnSubscribe(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}
	subscriptionId := ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		//subscriptionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subscriptionId is required",
		}))
		return
	}

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.COMM_HandleN1N2MessageUnSubscribe(ueContextId, subscriptionId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}

//sbi producer handler for N1N2MessageTransfer
func OnN1N2MessageTransfer(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.N1N2MessageTransferReqData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.N1N2MessageTransferRspData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleN1N2MessageTransfer(ueContextId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for N1N2MessageSubscribe
func OnN1N2MessageSubscribe(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.UeN1N2InfoSubscriptionCreateData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.UeN1N2InfoSubscriptionCreatedData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleN1N2MessageSubscribe(ueContextId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for NonUeN2InfoUnSubscribe
func OnNonUeN2InfoUnSubscribe(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	n2NotifySubscriptionId := ctx.Param("n2NotifySubscriptionId")
	if len(n2NotifySubscriptionId) == 0 {
		//n2NotifySubscriptionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "n2NotifySubscriptionId is required",
		}))
		return
	}

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.COMM_HandleNonUeN2InfoUnSubscribe(n2NotifySubscriptionId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}

//sbi producer handler for NonUeN2MessageTransfer
func OnNonUeN2MessageTransfer(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	var input models.N2InformationTransferReqData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.N2InformationTransferRspData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleNonUeN2MessageTransfer(input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for NonUeN2InfoSubscribe
func OnNonUeN2InfoSubscribe(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	var input models.NonUeN2InfoSubscriptionCreateData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.NonUeN2InfoSubscriptionCreatedData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleNonUeN2InfoSubscribe(input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for AMFStatusChangeSubscribe
func OnAMFStatusChangeSubscribe(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	var input models.SubscriptionData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SubscriptionData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleAMFStatusChangeSubscribe(input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


func OnDiscovery(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	log.Infof("ctx: %v", ctx.RequestBody())
	successCode := 200
	resp.SetBody(int(successCode), "retun from OnDiscovery")
	return
}

func OnDiscoveryResult(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	log.Info("ctx: ", ctx)
	successCode := 200
	resp.SetBody(int(successCode), "retun from OnDiscoveryResult")
	return
}

type Producer interface {
	COMM_HandleAMFStatusChangeSubscribeModfy(subscriptionId string, body models.SubscriptionData) (successCode int32, result models.SubscriptionData, err *sbi.ApiError)
	COMM_HandleAMFStatusChangeUnSubscribe(subscriptionId string) (successCode int32, err *sbi.ApiError)
	COMM_HandleCancelRelocateUEContext(ueContextId string, body models.CancelRelocateUEContextRequest) (successCode int32, err *sbi.ApiError)
	COMM_HandleCreateUEContext(ueContextId string, body models.CreateUEContextRequest) (successCode int32, result models.UeContextCreatedData, err *sbi.ApiError)
	COMM_HandleEBIAssignment(ueContextId string, body models.AssignEbiData) (successCode int32, result models.AssignedEbiData, err *sbi.ApiError)
	COMM_HandleRegistrationStatusUpdate(ueContextId string, body models.UeRegStatusUpdateReqData) (successCode int32, result models.UeRegStatusUpdateRspData, err *sbi.ApiError)
	COMM_HandleReleaseUEContext(ueContextId string, body models.UEContextRelease) (successCode int32, err *sbi.ApiError)
	COMM_HandleRelocateUEContext(ueContextId string, body models.RelocateUEContextRequest) (successCode int32, result models.UeContextRelocatedData, err *sbi.ApiError)
	COMM_HandleUEContextTransfer(ueContextId string, body models.UeContextTransferReqData) (successCode int32, result models.UeContextTransferRspData, err *sbi.ApiError)
	COMM_HandleN1N2MessageUnSubscribe(ueContextId string, subscriptionId string) (successCode int32, err *sbi.ApiError)
	COMM_HandleN1N2MessageTransfer(ueContextId string, body models.N1N2MessageTransferReqData) (successCode int32, result models.N1N2MessageTransferRspData, err *sbi.ApiError)
	COMM_HandleN1N2MessageSubscribe(ueContextId string, body models.UeN1N2InfoSubscriptionCreateData) (successCode int32, result models.UeN1N2InfoSubscriptionCreatedData, err *sbi.ApiError)
	COMM_HandleNonUeN2InfoUnSubscribe(n2NotifySubscriptionId string) (successCode int32, err *sbi.ApiError)
	COMM_HandleNonUeN2MessageTransfer(body models.N2InformationTransferReqData) (successCode int32, result models.N2InformationTransferRspData, err *sbi.ApiError)
	COMM_HandleNonUeN2InfoSubscribe(body models.NonUeN2InfoSubscriptionCreateData) (successCode int32, result models.NonUeN2InfoSubscriptionCreatedData, err *sbi.ApiError)
	COMM_HandleAMFStatusChangeSubscribe(body models.SubscriptionData) (successCode int32, result models.SubscriptionData, err *sbi.ApiError)
}

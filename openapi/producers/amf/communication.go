package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func AMFStatusChangeSubscribeModify(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {

	amfproducer := handler.(AmfProducer)

	var input models.SubscriptionData
	var result models.SubscriptionData
	var prob *models.ProblemDetails

	subId := ctx.Param("subscriptionId")

	//decode the request (and body)
	if prob = ctx.DecodeRequest(&input); prob == nil {
		//call the application handler
		result, prob = amfproducer.HandleStatusChangeSubscribeModify(subId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(202, &result)
	}
	return
}

func AMFStatusChangeSubscribe(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)

	var input models.SubscriptionData
	var result models.SubscriptionData
	var prob *models.ProblemDetails

	//decode the request (and body)
	if prob = ctx.DecodeRequest(&input); prob == nil {
		//call the application handler
		result, _, prob = amfproducer.HandleStatusChangeSubscribe(&input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(202, &result)
	}
	return
}

func AMFStatusChangeUnSubscribe(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	subId := ctx.Param("subscriptionId")
	//nothing to decode
	//call the application handler
	if prob := amfproducer.HandleStatusChangeUnSubscribe(subId); prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(204, nil)
	}
	return
}

func CreateUEContext(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {

	amfproducer := handler.(AmfProducer)
	ueContextId := ctx.Param("ueContextId")
	var input models.CreateUeContextRequest
	input.JsonData = new(models.UeContextCreateData)
	var prob *models.ProblemDetails
	var err *openapi.ApiError
	var result models.CreateUeContextResponse
	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, err, prob = amfproducer.HandleCreateUEContext(ueContextId, &input)
	}
	if prob != nil {
		resp.SetProblem(prob)
	} else if err != nil {
		resp.SetApiError(err)
	} else {
		resp.SetBody(201, &result)
	}

	return
}

// EBIAssignment - Namf_Communication EBI Assignment service Operation
func EBIAssignment(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {

	//get producer handler
	amfproducer := handler.(AmfProducer)

	ueContextId := ctx.Param("ueContextId")

	//decode request and call application handler
	var input models.AssignEbiData
	var result models.AssignedEbiData
	var prob *models.ProblemDetails
	var err *openapi.ApiError

	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, err, prob = amfproducer.HandleEBIAssignment(ueContextId, &input)
	}

	// set response body and status
	if prob != nil {
		resp.SetProblem(prob)
	} else if err != nil {
		resp.SetApiError(err)
	} else {
		resp.SetBody(200, &result)
	}
	return
}

func RegistrationStatusUpdate(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {

	amfproducer := handler.(AmfProducer)
	ueContextId := ctx.Param("ueContextId")
	var prob *models.ProblemDetails
	var input models.UeRegStatusUpdateReqData
	var result models.UeRegStatusUpdateRspData
	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, prob = amfproducer.HandleRegistrationStatusUpdate(ueContextId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, &result)
	}
	return
}

func ReleaseUEContext(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {

	amfproducer := handler.(AmfProducer)

	var input models.UeContextRelease
	var prob *models.ProblemDetails
	ueContextId := ctx.Param("ueContextId")
	if prob = ctx.DecodeRequest(&input); prob == nil {
		prob = amfproducer.HandleReleaseUEContext(ueContextId, &input)
	}
	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(204, nil)
	}
	return
}

func UEContextTransfer(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {

	amfproducer := handler.(AmfProducer)
	ueContextId := ctx.Param("ueContextId")

	var input models.UeContextTransferRequest
	var prob *models.ProblemDetails
	var result models.UeContextTransferResponse

	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, prob = amfproducer.HandleUEContextTransfer(ueContextId, &input)
	}
	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, &result)
	}

	return
}

func N1N2MessageUnSubscribe(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	ueContextId := ctx.Param("ueContextId")
	subId := ctx.Param("subscriptionId")

	//no need to decode request

	if prob := amfproducer.HandleN1N2MessageUnSubscribe(ueContextId, subId); prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(204, nil)
	}
	return
}

func N1N2MessageTransfer(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	ueContextId := ctx.Param("ueContextId")
	var input models.N1N2MessageTransferRequest
	var prob *models.ProblemDetails
	var err *openapi.ApiError
	var result models.N1N2MessageTransferRspData
	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, err, prob = amfproducer.HandleN1N2MessageTransfer(ueContextId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else if err != nil {
		resp.SetApiError(err)
	} else {
		resp.SetBody(200, &result) //NOTE: 200 or 202?
	}
	return
}

func N1N2MessageTransferStatus(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	ueContextId := ctx.Param("ueContextId")
	reqUri := ctx.Param("reqUri") //TODO: check how this param is decoded
	var prob *models.ProblemDetails
	prob = amfproducer.HandleN1N2MessageTransferStatus(ueContextId, reqUri)
	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, nil) //TODO: check the success code
	}
	return
}

func N1N2TransferFailureNotification(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	var prob *models.ProblemDetails
	var input models.N1N2MsgTxfrFailureNotification
	if prob = ctx.DecodeRequest(&input); prob == nil {
		prob = amfproducer.HandleN1N2TransferFailureNotification(&input)
	}
	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(204, nil)
	}
	return
}

func N1N2MessageSubscribe(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	ueContextId := ctx.Param("ueContextId")
	var prob *models.ProblemDetails
	var input models.UeN1N2InfoSubscriptionCreateData
	var result models.UeN1N2InfoSubscriptionCreatedData
	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, prob = amfproducer.HandleN1N2MessageSubscribe(ueContextId, &input)
	}
	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(201, &result) //NOTE: 200 or 202?
	}

	return
}

func NonUeN2InfoUnSubscribe(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	subId := ctx.Param("subscriptionId")
	var prob *models.ProblemDetails
	prob = amfproducer.HandleNonUeN2InfoUnSubscribe(subId)

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, nil) //NOTE: need to check the success code
	}
	return
}

func NonUeN2MessageTransfer(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	var prob *models.ProblemDetails
	var err *openapi.ApiError
	var input models.NonUeN2MessageTransferRequest
	var result models.N2InformationTransferRspData
	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, err, prob = amfproducer.HandleNonUeN2MessageTransfer(&input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else if err != nil {
		resp.SetApiError(err)
	} else {
		resp.SetBody(200, &result)
	}
	return
}

func NonUeN2InfoSubscribe(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	var prob *models.ProblemDetails
	var input models.NonUeN2InfoSubscriptionCreateData
	var result models.NonUeN2InfoSubscriptionCreatedData

	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, prob = amfproducer.HandleNonUeN2InfoSubscribe(&input)
	}
	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(201, &result)
	}
	return
}

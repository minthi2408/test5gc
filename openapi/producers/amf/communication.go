package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func HTTPAMFStatusChangeSubscribeModify(ctx openapi.RequestContext) {
	var subData models.SubscriptionData
	//decode the request (and body)
	ctx.DecodeRequest(&subData)
}
func HTTPAMFStatusChangeUnSubscribe(ctx openapi.RequestContext) {
	//do nothing
}

func HTTPCreateUEContext(ctx openapi.RequestContext) {
	/*
		    tqtung: complex

			var createUeContextRequest models.CreateUeContextRequest
			createUeContextRequest.JsonData = new(models.UeContextCreateData)
	*/
}

// EBIAssignment - Namf_Communication EBI Assignment service Operation
func HTTPEBIAssignment(ctx openapi.RequestContext) {
	var body models.AssignEbiData
	ctx.DecodeRequest(&body)
}

func HTTPRegistrationStatusUpdate(ctx openapi.RequestContext) {
	var body models.UeRegStatusUpdateReqData
	ctx.DecodeRequest(&body)
}

func HTTPReleaseUEContext(ctx openapi.RequestContext) {
	var body models.UeContextRelease
	ctx.DecodeRequest(&body)
}

func HTTPUEContextTransfer(ctx openapi.RequestContext) {
	//TODO: complex scenario
}

func HTTPN1N2MessageUnSubscribe(ctx openapi.RequestContext) {
	//do nothing
}

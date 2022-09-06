package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func ProvideLocationInfo(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)

	var input models.RequestLocInfo
	var result models.ProvideLocInfo
	var prob *models.ProblemDetails

	ueContextId := ctx.Param("ueContextId")

	//decode the request (and body)
	if prob = ctx.DecodeRequest(&input); prob == nil {
		//call the application handler
		result, prob = amfproducer.HandleProvideLocationInfo(ueContextId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, &result)
	}

	return
}

func ProvidePositioningInfo(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)

	var input models.RequestPosInfo
	var result models.ProvidePosInfo
	var prob *models.ProblemDetails

	ueContextId := ctx.Param("ueContextId")

	//decode the request (and body)
	if prob = ctx.DecodeRequest(&input); prob == nil {
		//call the application handler
		result, prob = amfproducer.HandleProvidePositioningInfo(ueContextId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, &result)
	}

	return
}

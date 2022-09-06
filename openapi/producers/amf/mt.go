package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func ProvideDomainSelectionInfo(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)

	var result models.UeContextInfo

	var prob *models.ProblemDetails

	ueContextId := ctx.Param("ueContextId")
	infoclass := ctx.Param("info-class")
	supfeatures := ctx.Param("supported-features")
	//TODO: validity check for the parameters and set a right status code

	result, prob = amfproducer.HandleProvideDomainSelectionInfo(ueContextId, infoclass, supfeatures)

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, &result)
	}

	return
}

func EnableUeReachability(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)

	var input models.EnableUeReachabilityReqData
	var result models.EnableUeReachabilityRspData
	var prob *models.ProblemDetails

	ueContextId := ctx.Param("ueContextId")

	//decode the request (and body)
	if prob = ctx.DecodeRequest(&input); prob == nil {
		//call the application handler
		result, prob = amfproducer.HandleEnableUeReachability(ueContextId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, &result)
	}

	return
}

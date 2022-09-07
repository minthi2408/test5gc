package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func SmContextStatusNotify(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	var prob *models.ProblemDetails
	var input models.SmContextStatusNotification
	guti := ctx.Param("guti")
	pduSessionId := ctx.Param("pduSessionId")

	if prob = ctx.DecodeRequest(&input); prob == nil {
		prob = amfproducer.HandleSmContextStatusNotify(guti, pduSessionId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, nil) //NOTE: need to check the success code
	}

	return
}

func AmPolicyControlUpdateNotifyUpdate(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	var prob *models.ProblemDetails
	var input models.PolicyUpdate
	polAssoId := ctx.Param("polAssoId")
	if prob = ctx.DecodeRequest(&input); prob == nil {
		prob = amfproducer.HandleAmPolicyControlUpdateNotifyUpdate(polAssoId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, nil) //NOTE: need to check the success code
	}
	return
}

func AmPolicyControlUpdateNotifyTerminate(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	var prob *models.ProblemDetails
	var input models.TerminationNotification
	polAssoId := ctx.Param("polAssoId")
	if prob = ctx.DecodeRequest(&input); prob == nil {
		prob = amfproducer.HandleAmPolicyControlUpdateNotifyTerminate(polAssoId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, nil) //NOTE: need to check the success code
	}

	return
}

func N1MessageNotify(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	var prob *models.ProblemDetails
	var input models.N1MessageNotify
	if prob = ctx.DecodeRequest(&input); prob == nil {
		prob = amfproducer.HandleN1MessageNotify(&input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, nil) //NOTE: need to check the success code
	}
	return

}

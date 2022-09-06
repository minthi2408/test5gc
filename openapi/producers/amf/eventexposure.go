package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func CreateSubscription(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {

	amfproducer := handler.(AmfProducer)
	var input models.AmfCreateEventSubscription
	var result models.AmfCreatedEventSubscription

	var prob *models.ProblemDetails

	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, prob = amfproducer.HandleCreateSubscription(&input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(201, &result)
	}

	return
}

func DeleteSubscription(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	subId := ctx.Param("subscriptionId")
	//TODO: check if subId is empty

	var prob *models.ProblemDetails

	prob = amfproducer.HandleDeleteSubscription(subId)

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, nil)
	}

	return
}
func ModifySubscription(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)
	var input models.ModifySubscriptionRequest
	var result models.AmfUpdatedEventSubscription

	var prob *models.ProblemDetails
	subId := ctx.Param("subscriptionId")
	//TODO: check if subId is empty

	if prob = ctx.DecodeRequest(&input); prob == nil {
		result, prob = amfproducer.HandleModifySubscription(subId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(200, &result)
	}

	return
}

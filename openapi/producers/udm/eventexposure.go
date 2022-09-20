package udm

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func CreateEeSubscription(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {

	udmproducer := handler.(UdmProducer)

	var input models.EeSubscription
	var result models.CreatedEeSubscription
	var prob *models.ProblemDetails

	ueId := ctx.Param("ueIdentity")

	//decode the request (and body)
	if prob = ctx.DecodeRequest(&input); prob == nil {
		//call the application handler
		result, prob = udmproducer.HandleCreateEeSubscription(ueId, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(201, &result)
	}
	return
}

func DeleteEeSubscription(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	udmproducer := handler.(UdmProducer)

	var prob *models.ProblemDetails

	ueId := ctx.Param("ueIdentity")
	subId := ctx.Param("subscriptionId")
	//TODO: check if any of the identities is empty

	//call the application handler
	prob = udmproducer.HandleDeleteEeSubscription(ueId, subId)

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(204, nil)
	}
	return
}

func UpdateEeSubscription(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	udmproducer := handler.(UdmProducer)

	var input []models.PatchItem
	var prob *models.ProblemDetails

	ueId := ctx.Param("ueIdentity")
	subId := ctx.Param("subscriptionId")
	//TODO: check if any of the identities is empty

	//decode the request (and body)
	if prob = ctx.DecodeRequest(&input); prob == nil {
		//call the application handler
		prob = udmproducer.HandleUpdateEeSubscription(ueId, subId, input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(204, nil)
	}
	return
}

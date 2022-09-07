package producer

import (
	"etri5gc/fabric/httpdp"
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	amfprod "etri5gc/openapi/producers/amf"
	"strings"

	openapi_http "etri5gc/openapi/httpdp"
)

func eventexposureService(p *Producer) (service httpdp.HttpService) {
	fn := openapi_http.CreateGinHandler
	service.Routes = httpdp.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			httpdp.HttpIndexHandler,
		},

		{
			"DeleteSubscription",
			strings.ToUpper("Delete"),
			"/subscriptions/:subscriptionId",
			fn(amfprod.DeleteSubscription, p),
		},

		{
			"ModifySubscription",
			strings.ToUpper("Patch"),
			"/subscriptions/:subscriptionId",
			fn(amfprod.ModifySubscription, p),
		},

		{
			"CreateSubscription",
			strings.ToUpper("Post"),
			"/subscriptions",
			fn(amfprod.CreateSubscription, p),
		},
	}
	service.Group = openapi.AMF_EVENTEXPOSURE
	return
}

func (p *Producer) HandleCreateSubscription(input *models.AmfCreateEventSubscription) (result models.AmfCreatedEventSubscription, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleDeleteSubscription(subId string) (prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleModifySubscription(subId string, input *models.ModifySubscriptionRequest) (result models.AmfUpdatedEventSubscription, prob *models.ProblemDetails) {
	return
}

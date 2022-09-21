package producer

import (
	"etri5gc/fabric/httpdp"
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	udmprod "etri5gc/openapi/producers/udm"
	"fmt"
	"strings"
)

func eventexposureService(p *Producer) (service httpdp.HttpService) {
	fn := ginHandler
	eesubs := openapi.UDM_EE_SUBSCRIPTIONS

	service.Routes = httpdp.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			httpdp.HttpIndexHandler,
		},
		{
			"HTTPCreateEeSubscription",
			strings.ToUpper("Post"),
			fmt.Sprintf("/:ueIdentity/%s", eesubs),
			fn(udmprod.CreateEeSubscription, p),
		},

		{
			"HTTPDeleteEeSubscription",
			strings.ToUpper("Delete"),
			fmt.Sprintf("/:ueIdentity/%s/:subscriptionId", eesubs),
			fn(udmprod.DeleteEeSubscription, p),
		},
		{
			"HTTPUpdateEeSubscription",
			strings.ToUpper("Patch"),
			fmt.Sprintf("/:ueIdentity/%s/:subscriptionId", eesubs),
			fn(udmprod.UpdateEeSubscription, p),
		},
	}
	service.Group = openapi.UDM_EVENTEXPOSURE
	return
}

func (p *Producer) HandleCreateEeSubscription(ueId string, input *models.EeSubscription) (result models.CreatedEeSubscription, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleDeleteEeSubscription(ueId string, subId string) (prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleUpdateEeSubscription(ueId string, subId string, items []models.PatchItem) (prob *models.ProblemDetails) {
	return
}



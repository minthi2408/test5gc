package producer

import (
	"etri5gc/fabric/httpdp"
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	amfprod "etri5gc/openapi/producers/amf"
)

func oamService(p *Producer) (service httpdp.HttpService) {
	fn := ginHandler
	service.Routes = httpdp.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			httpdp.HttpIndexHandler,
		},

		{
			"Registered UE Context",
			"GET",
			"/registered-ue-context",
			fn(amfprod.RegisteredUEContext, p),
		},

		{
			"Individual Registered UE Context",
			"GET",
			"/registered-ue-context/:supi",
			fn(amfprod.RegisteredUEContext, p),
		},
	}
	service.Group = openapi.AMF_OAM
	return
}

func (p *Producer) HandleRegisteredUEContext(supi string) (result []byte, prob *models.ProblemDetails) {
	return
}

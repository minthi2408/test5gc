package producer

import (
	"etri5gc/fabric/httpdp"
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	amfprod "etri5gc/openapi/producers/amf"
	"strings"
)

func locationService(p *Producer) (service httpdp.HttpService) {
	fn := ginHandler

	service.Routes = httpdp.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			httpdp.HttpIndexHandler,
		},

		{
			"ProvideLocationInfo",
			strings.ToUpper("Post"),
			"/:ueContextId/provide-loc-info",
			fn(amfprod.ProvideLocationInfo, p),
		},

		{
			"ProvidePositioningInfo",
			strings.ToUpper("Post"),
			"/:ueContextId/provide-pos-info",
			fn(amfprod.ProvidePositioningInfo, p),
		},
	}
	service.Group = openapi.AMF_LOCATION
	return
}

func (p *Producer) HandleProvideLocationInfo(ueContextId string, input *models.RequestLocInfo) (result models.ProvideLocInfo, prob *models.ProblemDetails) {
	return
}

func (p *Producer) HandleProvidePositioningInfo(ueContextId string, input *models.RequestPosInfo) (result models.ProvidePosInfo, prob *models.ProblemDetails) {
	return
}

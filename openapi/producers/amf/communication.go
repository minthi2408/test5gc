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

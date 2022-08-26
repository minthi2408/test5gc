package amf

import (
	"etri5gc/openapi/models"
	"etri5gc/openapi/prodcontext"
)

func HTTPAMFStatusChangeSubscribeModify(ctx prodcontext.RequestContext) {
	var subData models.SubscriptionData
	//decode the request (and body)
	ctx.DecodeRequest(&subData)
}

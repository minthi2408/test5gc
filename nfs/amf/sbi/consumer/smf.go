package consumer

import (
	"etri5gc/nfs/amf/context"
)

type SmfConsumer interface {
}

type smfConsumer struct {
}


func newSmfConsumer(context *context.AMFContext) SmfConsumer {
	return &smfConsumer{}
}

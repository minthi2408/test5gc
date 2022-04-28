package amf

import (
	"etri5gc/nfs/amf/sbi/consumer"
	"etri5gc/nfs/amf/context"
)


type Backend interface {
	Context() *context.AMFContext
	Consumer() *consumer.Consumer
}

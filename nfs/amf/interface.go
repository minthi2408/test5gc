package amf

import (
	"etri5gc/nfs/amf/sbi/consumer"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/config"
)


type Backend interface {
	Context() *context.AMFContext
	Consumer() *consumer.Consumer
	Config() *config.Config
}

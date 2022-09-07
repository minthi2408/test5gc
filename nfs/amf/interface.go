package amf

import (
	"etri5gc/nfs/amf/config"
	"etri5gc/nfs/amf/context"
)

type Backend interface {
	Context() *context.AmfContext
	Config() *config.AmfConfig
}

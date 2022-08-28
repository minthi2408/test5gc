package amf

import (
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/config"
)


type Backend interface {
	Context() *context.AmfContext
	Config() *config.Config
}

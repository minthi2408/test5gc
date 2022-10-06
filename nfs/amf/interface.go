package amf

import (
	"etrib5gc/nfs/amf/config"
	"etrib5gc/nfs/amf/context"
)

type Backend interface {
	Context() *context.AmfContext
	Config() *config.AmfConfig
}

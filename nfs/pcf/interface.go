package amf

import (
	"etri5gc/nfs/pcf/config"
	"etri5gc/nfs/pcf/context"
)

type Backend interface {
	Context() *context.PcfContext
	Config() *config.PcfConfig
}

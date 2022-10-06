package amf

import (
	"etrib5gc/nfs/pcf/config"
	"etrib5gc/nfs/pcf/context"
)

type Backend interface {
	Context() *context.PcfContext
	Config() *config.PcfConfig
}

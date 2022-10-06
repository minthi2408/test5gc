package amf

import (
	"etrib5gc/nfs/udm/config"
	"etrib5gc/nfs/udm/context"
)

type Backend interface {
	Context() *context.UdmContext
	Config() *config.UdmConfig
}

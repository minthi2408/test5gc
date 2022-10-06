package amf

import (
	"etri5gc/nfs/udm/config"
	"etri5gc/nfs/udm/context"
)

type Backend interface {
	Context() *context.UdmContext
	Config() *config.UdmConfig
}

package amf

import (
	"etri5gc/nfs/smf/config"
	"etri5gc/nfs/smf/context"
)

type Backend interface {
	Context() *context.SmfContext
	Config() *config.SmfConfig
}

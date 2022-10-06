package amf

import (
	"etrib5gc/nfs/smf/config"
	"etrib5gc/nfs/smf/context"
)

type Backend interface {
	Context() *context.SmfContext
	Config() *config.SmfConfig
}

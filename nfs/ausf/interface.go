package amf

import (
	"etrib5gc/nfs/ausf/config"
	"etrib5gc/nfs/ausf/context"
)

type Backend interface {
	Context() *context.AusfContext
	Config() *config.AusfConfig
}

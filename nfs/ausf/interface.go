package amf

import (
	"etri5gc/nfs/ausf/config"
	"etri5gc/nfs/ausf/context"
)

type Backend interface {
	Context() *context.AusfContext
	Config() *config.AusfConfig
}

package context

import (
	"etrib5gc/nfs/ausf/config"
)

type AusfContext struct {
	conf   *config.AusfConfig
	uelist AusfUeList
}

func New(conf *config.AusfConfig) *AusfContext {
	ret := &AusfContext{
		conf: conf,
	}
	return ret
}

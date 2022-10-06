package context

import (
	"etrib5gc/nfs/pcf/config"
)

type PcfContext struct {
	conf   *config.PcfConfig
	uelist PcfUeList
}

func New(conf *config.PcfConfig) *PcfContext {
	ret := &PcfContext{
		conf: conf,
	}
	return ret
}

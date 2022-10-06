package context

import (
	"etri5gc/nfs/smf/config"
)

type SmfContext struct {
	conf   *config.SmfConfig
	uelist SmfUeList
}

func New(conf *config.SmfConfig) *SmfContext {
	ret := &SmfContext{
		conf: conf,
	}
	return ret
}

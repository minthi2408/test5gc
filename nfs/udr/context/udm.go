package context

import (
	"etri5gc/nfs/udr/config"
)

type UdrContext struct {
	conf   *config.UdrConfig
	uelist UdrUeList
}

func New(conf *config.UdrConfig) *UdrContext {
	ret := &UdrContext{
		conf: conf,
	}
	return ret
}

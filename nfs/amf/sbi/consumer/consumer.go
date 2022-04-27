package consumer

import (
//	"fmt"
	"etri5gc/sbi"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/config"
)

type Backend interface {
	Context() *context.AMFContext
	Config() *config.Config
}

type Consumer struct {
	backend		Backend
	amf			AmfConsumer
	smf			SmfConsumer
	pcf			PcfConsumer
	nrf			sbi.NrfConsumer
}


func NewConsumer(b Backend) *Consumer {
	ret := &Consumer{backend: b}
	
	// create an interface to NRF
	ret.nrf = sbi.NewNrfConsumer(b.Context())
	ret.amf = newAmfConsumer(b.Context())
	ret.smf = newSmfConsumer(b.Context())
	ret.pcf = newPcfConsumer(b.Context())
	return ret
}

func (c *Consumer) Nrf() sbi.NrfConsumer  {
	return c.nrf
}
func (c *Consumer) Amf() AmfConsumer {
	return c.amf
}
func (c *Consumer) Pcf() PcfConsumer {
	return c.pcf
}

func (c *Consumer) Smf() SmfConsumer {
	return c.smf
}


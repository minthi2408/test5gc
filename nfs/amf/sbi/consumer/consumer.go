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
	ausf		*ausfConsumer
	amf			*amfConsumer
	smf			*smfConsumer
	udm			*udmConsumer
	pcf			*pcfConsumer
	nssf		*nssfConsumer
	callback	*callback
	nrf			sbi.NrfConsumer
}


func NewConsumer(b Backend) *Consumer {
	ret := &Consumer{backend: b}
	
	// create an interface to NRF
	ret.nrf			= sbi.NewNrfConsumer(b.Context())
	ret.amf			= newAmfConsumer(b.Context())
	ret.ausf		= newAusfConsumer(b.Context())
	ret.smf			= newSmfConsumer(b.Context())
	ret.udm			= newUdmConsumer(b.Context())
	ret.pcf			= newPcfConsumer(b.Context())
	ret.nssf		= newNssfConsumer(b.Context())
	ret.callback	= newCallback(b.Context())
	return ret
}

func (c *Consumer) Nrf() sbi.NrfConsumer  {
	return c.nrf
}
func (c *Consumer) Amf() *amfConsumer {
	return c.amf
}

func (c *Consumer) Ausf() *ausfConsumer {
	return c.ausf
}
func (c *Consumer) Nssf() *nssfConsumer {
	return c.nssf
}
func (c *Consumer) Pcf() *pcfConsumer {
	return c.pcf
}

func (c *Consumer) Smf() *smfConsumer {
	return c.smf
}

func (c *Consumer) Udm() *udmConsumer {
	return c.udm
}
func (c *Consumer) Callback() *callback {
	return c.callback
}

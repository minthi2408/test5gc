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
	nrf			sbi.NrfConsumer
}


func NewConsumer(b Backend) *Consumer {
	ret := &Consumer{backend: b}
	
	// create an interface to NRF
	ret.nrf = sbi.NewNrfConsumer(b.Config().Configuration.NrfUri, b.Context().BuildNfProfile)
	ret.amf = newAmfConsumer(b.Context())
	ret.smf = newSmfConsumer(b.Context())
	return ret
}

func (c *Consumer) NRF() sbi.NrfConsumer  {
	return c.nrf
}
func (c *Consumer) AMF() AmfConsumer {
	return c.amf
}


func (c *Consumer) SMF() SmfConsumer {
	return c.smf
}


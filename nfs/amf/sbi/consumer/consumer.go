package consumer

import (
	"fmt"
	"etri5gc/common"
)

type AmfConsumer interface {
	//all requesting methods to AMF
	CreateUEContextRequest(ue *amf_context.AmfUe, ueContextCreateData models.UeContextCreateData)

	ReleaseUEContextRequest(ue *amf_context.AmfUe, ngapCause models.NgApCause)
	UEContextTransferRequest(ue *amf_context.AmfUe, accessType models.AccessType, transferReason models.TransferReason) (ueContextTransferRspData *models.UeContextTransferRspData, problemDetails *models.ProblemDetails, err error) 
}

type SmfConsumer interface {
	//all requesting methods to SMF
}

type Backend interface {
	Context() context.AmfContext
}

type Consumer struct {
	backend		Backend
	amf			AmfConsumer
	smf			SmfConsumer
	nrf			common.NrfConsumer
}


func NewConsumer(b Backend) *Consumer {
	ret := &Consumer{backend: b}
	
	// create an interface to NRF
	ret.nrf = common.NewNrfConsumer(b.Context().BuildProfile)
	ret.amf = newAmf(b.Context)
	ret.smf = newSmf(b.Context)
	//TODO: Create consumers
	return ret
}

func (c *Consumer) NRF()  {
	return c.nrf
}
func (c *Consumer) AMF() AmfConsumer {
	return c.amf
}


func (c *Consumer) SMF() SmfConsumer {
	return c.smf
}


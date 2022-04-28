package nas

import (
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf"
	"etri5gc/nfs/amf/sbi/consumer"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/openapi/models"
)

type NgapSender interface {
}

type Nas struct {
	backend		amf.Backend
	ngap		NgapSender
}

func NewNas(b amf.Backend, ngap NgapSender) *Nas {
	return &Nas{
		backend:	b,
		ngap:		ngap,
	}
}

func (n *Nas) amf() *context.AMFContext {
	return n.backend.Context()
}
func (n *Nas) callback() consumer.Callback {
	return n.backend.Consumer().Callback()
}

func (n *Nas) Send() {
}

func (n *Nas) HandleNAS(ranue *context.RanUe,code int64,naspdu []byte) {
	//TODO:
}


func (n *Nas) SendConfigurationUpdateCommand(*context.AmfUe, models.AccessType, *nasType.NetworkSlicingIndication) {

}

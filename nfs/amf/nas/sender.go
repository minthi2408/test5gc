package nas

import (
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/openapi/models"
)

type Nas struct {
	backend		amf.Backend
}

func NewNas(b amf.Backend) *Nas {
	return &Nas{
		backend:	b,
	}
}

func (n *Nas) Send() {
}

func (n *Nas) HandleNAS(ranue *context.RanUe,code int64,naspdu []byte) {
	//TODO:
}

func (n *Nas) BuildConfigurationUpdateCommand(*context.AmfUe, models.AccessType, *nasType.NetworkSlicingIndication) ([]byte, error) {
	return []byte{}, nil 
}

func (n *Nas) SendConfigurationUpdateCommand(*context.AmfUe, models.AccessType, *nasType.NetworkSlicingIndication) {

}

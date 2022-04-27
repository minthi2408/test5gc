package nas

import (
	"etri5gc/nfs/amf/sbi/producer"
	"etri5gc/nfs/amf/context"
)

type nasSender struct {
	backend		producer.Backend
}

func NewNas(b producer.Backend) producer.NasSender {
	return &nasSender{
		backend:	b,
	}
}

func (nas *nasSender) Send() {
}

func (nas *nasSender) HandleNAS(ranue *context.RanUe,code int64,naspdu []byte) {
	//TODO:
}

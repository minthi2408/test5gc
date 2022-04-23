package producer

import (
	"fmt"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/sbi/producer"
)

type Backend interface {
	GetConsumer() consumer.Consumer
	GetContext() context.AmfContext
}

type Handler struct {
	backend		Backend
}

func NewHandler(b Backend) *Handler {
	return &Hander{backend: b}
}



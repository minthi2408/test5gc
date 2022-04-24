package producer

import (
//	"fmt"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/sbi/consumer"
)

type Backend interface {
	Consumer() *consumer.Consumer
	Context() *context.AMFContext
}

type Handler struct {
	backend		Backend
}

func NewHandler(b Backend) *Handler {
	return &Handler{backend: b}
}



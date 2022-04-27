package producer

import (
//	"fmt"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/sbi/consumer"
)

type NasSender interface {
	HandleNAS(*context.RanUe, int64, []byte)
}

type NgapSender interface {
}


type Backend interface {
	Consumer() *consumer.Consumer
	Context() *context.AMFContext
	Nas()	NasSender
	Ngap()  NgapSender
}

type Handler struct {
	backend		Backend
}

func NewHandler(b Backend) *Handler {
	return &Handler{backend: b}
}

func (h *Handler) nas() NasSender {
	return h.backend.Nas()
}

func (h *Handler) ngap() NgapSender {
	return h.backend.Ngap()
}

func (h *Handler) amf() *context.AMFContext {
	return h.backend.Context()
}

package ngap

import (
		"etri5gc/nfs/amf/sbi/producer"
	)	
/*
type NgapSender interface {
	SendToRan(*context.AmfRan, []byte)
	SendToRanUe(*context.RanUe, []byte) {
}
*/
/*
type NgapHandler interface {
	Handle()
}
*/

type Ngap struct {
	sender		ngapSender
	handler 	ngapHandler
}

type ngapSender struct {
	backend		producer.Backend
}

type ngapHandler struct {
	backend producer.Backend
}

func NewNgap(b producer.Backend) *Ngap {
	return &Ngap{
		sender: ngapSender{
			backend: b,
		},
		handler: ngapHandler{
			backend: b,
		},
	}
}


func (n *Ngap) Handler() *ngapHandler {
	return &n.handler
}

func (n *Ngap) Sender() *ngapSender {
	return &n.sender
}

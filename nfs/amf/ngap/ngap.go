package ngap

import (
		"etri5gc/nfs/amf"
		"etri5gc/nfs/amf/nas"
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
	nas			*nas.Nas
}

type ngapSender struct {
	backend		amf.Backend
}

type ngapHandler struct {
	backend amf.Backend
	sender	*ngapSender
	nas		*nas.Nas
}

func NewNgap(b amf.Backend) *Ngap {
	ret := &Ngap{
		sender: ngapSender{
			backend: b,
		},
		handler: ngapHandler{
			backend: b,
		},
	}
	//create Nas
	ret.nas = nas.NewNas(b, &ret.sender)
	ret.handler.sender = &ret.sender
	ret.handler.nas = ret.nas
	return ret
}


func (n *Ngap) Handler() *ngapHandler {
	return &n.handler
}

func (n *Ngap) Sender() *ngapSender {
	return &n.sender
}

//expose Nas
func (n *Ngap) Nas() *nas.Nas {
	return n.nas
}

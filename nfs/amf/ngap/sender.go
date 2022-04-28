package ngap

import (
		"etri5gc/nfs/amf/sbi/producer"
	)
	
	
type ngapSender struct {
	backend		producer.Backend
}
func NewNgap(b producer.Backend) producer.NgapSender {
	return &ngapSender{
		backend: b,
	}
}

func (ngap *ngapSender) Send() {
}

package service

import (
	"etri5gc/nfs/amf/sbi/nfselect"
)

type nfSelector struct {
	nf	*AMF
}


func NewNfSelector(nf *AMF) nfselect.NfSelector {
	return &nfSelector{
		nf:	nf,
	}
}


func (s *nfSelector) SelectUdm() {
}

func (s *nfSelector) SelectPcf() {
}

func (s *nfSelector) SelectSmf() {
}

func (s *nfSelector) SelectAusf() {
}

func (s *nfSelector) SelectNssi() {
}

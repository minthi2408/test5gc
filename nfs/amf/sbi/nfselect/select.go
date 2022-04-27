package nfselect

import (
//	"fmt"
)

type NfSelector interface {
	UdmSelector
	PcfSelector
	SmfSelector
	AusfSelector
	NssiSelector
}

type UdmSelector interface {
	SelectUdm()
}


type PcfSelector interface {
	SelectPcf()
}

type SmfSelector interface {
	SelectSmf()
}

type AusfSelector interface {
	SelectAusf()
}

type NssiSelector interface {
	SelectAusf()
}

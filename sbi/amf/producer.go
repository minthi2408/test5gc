package amf

import (
	"etri5gc/sbi/amf/comm"
	"etri5gc/sbi/amf/ee"
	"etri5gc/sbi/amf/loc"
	"etri5gc/sbi/amf/mt"
)

type Producer interface {
	comm.Producer
	ee.Producer
	loc.Producer
	mt.Producer
}

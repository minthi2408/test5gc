package amf

import (
	"etr5gc/sbi/amf/comm"
	"etr5gc/sbi/amf/ee"
	"etr5gc/sbi/amf/loc"
	"etr5gc/sbi/amf/mt"
)

type Producer interface {
	comm.Producer
	ee.Producer
	loc.Producer
	mt.Producer
}

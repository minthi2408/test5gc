package pcf

import (
	"etri5gc/sbi/pcf/ampc"
	"etri5gc/sbi/pcf/btdpc"
	"etri5gc/sbi/pcf/ee"
	"etri5gc/sbi/pcf/pa"
	"etri5gc/sbi/pcf/smpc"
	"etri5gc/sbi/pcf/uepc"
)

type Producer interface {
	pa.Producer
	ee.Producer
	ampc.Producer
	smpc.Producer
	uepc.Producer
	btdpc.Producer
}

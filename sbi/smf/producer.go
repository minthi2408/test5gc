package smf

import (
	"etri5gc/sbi/smf/ee"
	"etri5gc/sbi/smf/nidd"
	"etri5gc/sbi/smf/pdu"
)

type Producer interface {
	pdu.Producer
	ee.Producer
	nidd.Producer
}

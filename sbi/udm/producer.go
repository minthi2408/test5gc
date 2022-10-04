package udm

import (
	"etri5gc/sbi/udm/ee"
	"etri5gc/sbi/udm/mt"
	"etri5gc/sbi/udm/niddau"
	"etri5gc/sbi/udm/pp"
	"etri5gc/sbi/udm/sdm"
	"etri5gc/sbi/udm/ueau"
	"etri5gc/sbi/udm/uecm"
)

type Producer interface {
	ee.Producer
	mt.Producer
	sdm.Producer
	pp.Producer
	ueau.Producer
	uecm.Producer
	niddau.Producer
}

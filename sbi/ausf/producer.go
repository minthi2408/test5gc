package ausf

import (
	"etri5gc/sbi/ausf/sor"
	"etri5gc/sbi/ausf/uea"
	"etri5gc/sbi/ausf/upu"
)

type Producer interface {
	sor.Producer
	uea.Producer
	upu.Producer
}

package producer

import (
	//	"fmt"
	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/nfs/pcf/context"
	sbi_httpdp "etri5gc/sbi/httpdp"
	pcfampc "etri5gc/sbi/pcf/ampc"
	pcfbtdpc "etri5gc/sbi/pcf/btdpc"
	pcfee "etri5gc/sbi/pcf/ee"
	pcfpa "etri5gc/sbi/pcf/pa"
	pcfsmpc "etri5gc/sbi/pcf/smpc"
	pcfuepc "etri5gc/sbi/pcf/uepc"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi.producer"})
}

type Backend interface {
	Context() *context.PcfContext
}

type Producer struct {
	backend Backend
}

func New(b Backend) *Producer {
	return &Producer{
		backend: b,
	}
}

// build services to register to the underlying server (http server in service
// agent)
func (prod *Producer) Services() []common.Service {
	services := make([]common.Service, 6, 6)
	services[0] = httpdp.HttpService{
		Group:  "ee",
		Routes: sbi_httpdp.MakeHttpRoutes(pcfee.Routes, prod),
	}
	services[1] = httpdp.HttpService{
		Group:  "pa",
		Routes: sbi_httpdp.MakeHttpRoutes(pcfpa.Routes, prod),
	}
	services[2] = httpdp.HttpService{
		Group:  "ampc",
		Routes: sbi_httpdp.MakeHttpRoutes(pcfampc.Routes, prod),
	}
	services[3] = httpdp.HttpService{
		Group:  "smpc",
		Routes: sbi_httpdp.MakeHttpRoutes(pcfsmpc.Routes, prod),
	}
	services[4] = httpdp.HttpService{
		Group:  "uepc",
		Routes: sbi_httpdp.MakeHttpRoutes(pcfuepc.Routes, prod),
	}
	services[5] = httpdp.HttpService{
		Group:  "btdpc",
		Routes: sbi_httpdp.MakeHttpRoutes(pcfbtdpc.Routes, prod),
	}

	return services
}

// access to the internal data structures of the AMF
func (prod *Producer) pcf() *context.PcfContext {
	return prod.backend.Context()
}

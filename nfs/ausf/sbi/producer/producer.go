package producer

import (
	//	"fmt"
	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/nfs/ausf/context"
	ausfsor "etri5gc/sbi/ausf/sor"
	ausfuea "etri5gc/sbi/ausf/uea"
	ausfupu "etri5gc/sbi/ausf/upu"
	sbi_httpdp "etri5gc/sbi/httpdp"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi.producer"})
}

type Backend interface {
	Context() *context.AusfContext
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
	services := make([]common.Service, 3, 3)
	services[0] = httpdp.HttpService{
		Group:  "sor",
		Routes: sbi_httpdp.MakeHttpRoutes(ausfsor.Routes, prod),
	}
	services[1] = httpdp.HttpService{
		Group:  "uea",
		Routes: sbi_httpdp.MakeHttpRoutes(ausfuea.Routes, prod),
	}
	services[2] = httpdp.HttpService{
		Group:  "upu",
		Routes: sbi_httpdp.MakeHttpRoutes(ausfupu.Routes, prod),
	}

	return services
}

// access to the internal data structures of the AMF
func (prod *Producer) ausf() *context.AusfContext {
	return prod.backend.Context()
}

package producer

import (
	//	"fmt"
	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/nfs/udr/context"
	sbi_httpdp "etri5gc/sbi/httpdp"
	udrdr "etri5gc/sbi/udr/dr"
	udrgroup "etri5gc/sbi/udr/group"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi.producer"})
}

type Backend interface {
	Context() *context.UdrContext
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
	services := make([]common.Service, 2, 2)
	services[0] = httpdp.HttpService{
		Group:  "dr",
		Routes: sbi_httpdp.MakeHttpRoutes(udrdr.Routes, prod),
	}
	services[1] = httpdp.HttpService{
		Group:  "group",
		Routes: sbi_httpdp.MakeHttpRoutes(udrgroup.Routes, prod),
	}
	return services
}

// access to the internal data structures of the AMF
func (prod *Producer) udr() *context.UdrContext {
	return prod.backend.Context()
}

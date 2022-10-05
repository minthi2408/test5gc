package producer

import (
	//	"fmt"
	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/nfs/udm/context"
	sbi_httpdp "etri5gc/sbi/httpdp"
	udmee "etri5gc/sbi/udm/ee"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi.producer"})
}

type Backend interface {
	Context() *context.UdmContext
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
	services := make([]common.Service, 7, 7)
	services[0] = httpdp.HttpService{
		Group:  "comm",
		Routes: sbi_httpdp.MakeHttpRoutes(udmee.Routes, prod),
	}

	return services
}

// access to the internal data structures of the AMF
func (prod *Producer) udm() *context.UdmContext {
	return prod.backend.Context()
}

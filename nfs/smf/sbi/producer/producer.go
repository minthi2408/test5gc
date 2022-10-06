package producer

import (
	//	"fmt"
	"etrib5gc/fabric/common"
	"etrib5gc/fabric/httpdp"
	"etrib5gc/nfs/smf/context"
	sbi_httpdp "etrib5gc/sbi/httpdp"
	smfee "etrib5gc/sbi/smf/ee"
	smfnidd "etrib5gc/sbi/smf/nidd"
	smfpdu "etrib5gc/sbi/smf/pdu"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi.producer"})
}

type Backend interface {
	Context() *context.SmfContext
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
		Group:  "ee",
		Routes: sbi_httpdp.MakeHttpRoutes(smfee.Routes, prod),
	}
	services[1] = httpdp.HttpService{
		Group:  "pdu",
		Routes: sbi_httpdp.MakeHttpRoutes(smfpdu.Routes, prod),
	}
	services[2] = httpdp.HttpService{
		Group:  "nidd",
		Routes: sbi_httpdp.MakeHttpRoutes(smfnidd.Routes, prod),
	}
	return services
}

// access to the internal data structures of the AMF
func (prod *Producer) smf() *context.SmfContext {
	return prod.backend.Context()
}

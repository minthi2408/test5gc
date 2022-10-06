package producer

import (
	//	"fmt"
	"etrib5gc/fabric/common"
	"etrib5gc/fabric/httpdp"
	"etrib5gc/nfs/udm/context"
	sbi_httpdp "etrib5gc/sbi/httpdp"
	udmee "etrib5gc/sbi/udm/ee"
	udmmt "etrib5gc/sbi/udm/mt"
	udmniddau "etrib5gc/sbi/udm/niddau"
	udmpp "etrib5gc/sbi/udm/pp"
	udmsdm "etrib5gc/sbi/udm/sdm"
	udmueau "etrib5gc/sbi/udm/ueau"
	udmuecm "etrib5gc/sbi/udm/uecm"

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
		Group:  "ee",
		Routes: sbi_httpdp.MakeHttpRoutes(udmee.Routes, prod),
	}
	services[1] = httpdp.HttpService{
		Group:  "pp",
		Routes: sbi_httpdp.MakeHttpRoutes(udmpp.Routes, prod),
	}
	services[2] = httpdp.HttpService{
		Group:  "mt",
		Routes: sbi_httpdp.MakeHttpRoutes(udmmt.Routes, prod),
	}
	services[3] = httpdp.HttpService{
		Group:  "sdm",
		Routes: sbi_httpdp.MakeHttpRoutes(udmsdm.Routes, prod),
	}
	services[4] = httpdp.HttpService{
		Group:  "ueau",
		Routes: sbi_httpdp.MakeHttpRoutes(udmueau.Routes, prod),
	}
	services[5] = httpdp.HttpService{
		Group:  "uecm",
		Routes: sbi_httpdp.MakeHttpRoutes(udmuecm.Routes, prod),
	}
	services[6] = httpdp.HttpService{
		Group:  "niddau",
		Routes: sbi_httpdp.MakeHttpRoutes(udmniddau.Routes, prod),
	}
	return services
}

// access to the internal data structures of the AMF
func (prod *Producer) udm() *context.UdmContext {
	return prod.backend.Context()
}

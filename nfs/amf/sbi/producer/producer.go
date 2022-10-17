package producer

import (
	//	"fmt"
	"etrib5gc/nfs/amf/context"
	amfcomm "etrib5gc/sbi/amf/comm"
	amfee "etrib5gc/sbi/amf/ee"
	amfloc "etrib5gc/sbi/amf/loc"
	amfmt "etrib5gc/sbi/amf/mt"

	"etrib5gc/fabric/common"
	"etrib5gc/sbi/models"

	"etrib5gc/fabric/httpdp"
	sbi_httpdp "etrib5gc/sbi/httpdp"

	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/ngap/ngapType"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi.producer"})
}

// interface to NAS handler
type NasHandler interface {
	HandleNAS(*context.RanUe, int64, []byte)
}

type NasGmm interface {
	BuildConfigurationUpdateCommand(*context.AmfUe, models.AccessType, *nasType.NetworkSlicingIndication) ([]byte, error)
	SendConfigurationUpdateCommand(*context.AmfUe, models.AccessType, *nasType.NetworkSlicingIndication)
}

type NasInf interface {
	NasHandler
	NasGmm
}

type NgapSender interface {
	SendPaging(*context.AmfUe, []byte)
	BuildPaging(*context.AmfUe, *ngapType.PagingPriority, bool) ([]byte, error)
}

type Backend interface {
	Context() *context.AmfContext
}

type Producer struct {
	backend Backend
	ngap    NgapSender //sending NGAP messages to gnB
	nas     NasInf     //sending and handling NAS message
}

func New(b Backend, ngap NgapSender, nas NasInf) *Producer {
	return &Producer{
		backend: b,
		ngap:    ngap,
		nas:     nas,
	}
}

// build services to register to the underlying server (http server in service
// agent)
func (prod *Producer) Services() []common.Service {
	log.Info("build producer service");
	services := make([]common.Service, 4, 4)
	services[0] = httpdp.HttpService{
		Group:  "comm",
		Routes: sbi_httpdp.MakeHttpRoutes(amfcomm.Routes, prod),
	}
	services[1] = httpdp.HttpService{
		Group:  "ee",
		Routes: sbi_httpdp.MakeHttpRoutes(amfee.Routes, prod),
	}

	services[2] = httpdp.HttpService{
		Group:  "loc",
		Routes: sbi_httpdp.MakeHttpRoutes(amfloc.Routes, prod),
	}

	services[3] = httpdp.HttpService{
		Group:  "mt",
		Routes: sbi_httpdp.MakeHttpRoutes(amfmt.Routes, prod),
	}
	return services
}

// access to the internal data structures of the AMF
func (prod *Producer) amf() *context.AmfContext {
	return prod.backend.Context()
}

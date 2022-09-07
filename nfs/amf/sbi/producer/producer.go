package producer

import (
	//	"fmt"
	"etri5gc/nfs/amf/context"

	"etri5gc/fabric/common"
	"etri5gc/openapi/models"

	"etri5gc/openapi"
	openapi_http "etri5gc/openapi/httpdp"
	amfprod "etri5gc/openapi/producers/amf"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/ngap/ngapType"
	"github.com/gin-gonic/gin"
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
	services := make([]common.Service, 6, 6)
	services[0] = callbackService(prod)
	services[1] = communicationService(prod)
	services[2] = eventexposureService(prod)
	services[3] = locationService(prod)
	services[4] = mtService(prod)
	services[5] = oamService(prod)
	return services
}

// access to the internal data structures of the AMF
func (prod *Producer) amf() *context.AmfContext {
	return prod.backend.Context()
}

func ginHandler(fn openapi.OpenApiProducerHandler, p amfprod.AmfProducer) gin.HandlerFunc {
	return openapi_http.CreateGinHandler(fn, p)
}

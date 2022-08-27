package producer

import (
	//	"fmt"
	"etri5gc/nfs/amf/context"
	"etri5gc/openapi"
	amfprod "etri5gc/openapi/producers/amf"
	"strings"

	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/fabric/httpdp/encoding"
	"etri5gc/nfs/amf/sbi/producer/communication"
	"etri5gc/nfs/amf/sbi/producer/eventexposure"
	"etri5gc/nfs/amf/sbi/producer/httpcallback"
	"etri5gc/nfs/amf/sbi/producer/location"
	"etri5gc/nfs/amf/sbi/producer/mt"
	"etri5gc/nfs/amf/sbi/producer/oam"
	"etri5gc/openapi/models"
	openapi_http "etri5gc/openapi/httpdp"

	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/ngap/ngapType"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi.producer"})
}

//interface to NAS handler
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
	Context() *context.AMFContext
}

type Producer struct {
	backend Backend
	ngap    NgapSender //sending NGAP messages to gnB
	nas     NasInf     //sending and handling NAS message
}

func NewProducer(b Backend, ngap NgapSender, nas NasInf) *Producer {
	return &Producer{
		backend: b,
		ngap:    ngap,
		nas:     nas,
	}
}

//build services to register to the underlying server (http server in service
//agent)
func (prod *Producer) Services() []common.Service {
	services := make([]common.Service, 6, 6)
	services[0] = httpcallback.MakeService(prod)
	services[1] = communication.MakeService(prod)
	services[2] = eventexposure.MakeService(prod)
	services[3] = location.MakeService(prod)
	services[4] = mt.MakeService(prod)
	services[5] = oam.MakeService(prod)
	return services
}

// access to the internal data structures of the AMF
func (prod *Producer) amf() *context.AMFContext {
	return prod.backend.Context()
}

func MakeTestService() (service httpdp.HttpService) {
	service.Routes = httpdp.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			httpdp.HttpIndexHandler,
		},

		{
			"AMFStatusChangeSubscribeModify",
			strings.ToUpper("Put"),
			"/subscriptions/:subscriptionId",
			openapi_http.BuildProducerRequestHandler(amfprod.HTTPAMFStatusChangeSubscribeModify, myapphandler, encoding.New()),
		},
	}
	service.Group = "test"
	return
}

func myapphandler(ctx openapi.RequestContext) *openapi.Response {
	return nil
}

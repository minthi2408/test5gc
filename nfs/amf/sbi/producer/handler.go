package producer

import (
	//	"fmt"
	"etri5gc/nfs/amf/context"
	"etri5gc/openapi"
	amfprod "etri5gc/openapi/producers/amf"
	"strings"

	"etri5gc/fabric/common"
	"etri5gc/fabric/httpdp"
	"etri5gc/nfs/amf/sbi/producer/communication"
	"etri5gc/nfs/amf/sbi/producer/eventexposure"
	"etri5gc/nfs/amf/sbi/producer/httpcallback"
	"etri5gc/nfs/amf/sbi/producer/location"
	"etri5gc/nfs/amf/sbi/producer/mt"
	"etri5gc/nfs/amf/sbi/producer/oam"
	openapi_http "etri5gc/openapi/httpdp"
	"etri5gc/openapi/models"

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
func (prod *Producer) amf() *context.AmfContext {
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
			openapi_http.BuildProducerRequestHandler(amfprod.HTTPAMFStatusChangeSubscribeModify, myapphandler),
		},
	}
	service.Group = "test"
	return
}

func myapphandler(ctx openapi.RequestContext) *openapi.Response {
	return nil
}

/*
func (prod *Producer) makeCommServices(p Backend) (service httpdp.HttpService) {
	fn := openapi_http.BuildProducerRequestHandler
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
			fn(amfprod.HTTPAMFStatusChangeSubscribeModify, prod.HandleAMFStatusChangeSubscribeModify),
		},

		{
			"AMFStatusChangeUnSubscribe",
			strings.ToUpper("Delete"),
			"/subscriptions/:subscriptionId",
			fn(amfprod.HTTPAMFStatusChangeUnSubscribe, prod.HTTPAMFStatusChangeUnSubscribe),
		},

		{
			"CreateUEContext",
			strings.ToUpper("Put"),
			"/ue-contexts/:ueContextId",
			fn(amfprod.HTTPCreateUEContext, prod.HTTPCreateUEContext),
		},

		{
			"EBIAssignment",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/assign-ebi",
			fn(amfprod.HTTPEBIAssignment, prod.HTTPEBIAssignment),
		},

		{
			"RegistrationStatusUpdate",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/transfer-update",
			fn(amfprod.HTTPRegistrationStatusUpdate, prod.HTTPRegistrtionStatusUpdat),
		},

		{
			"ReleaseUEContext",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/release",
			fn(amfprod.HTTPReleaseUEContext, prod.HTTPReleaseUEContext),
		},

		{
			"UEContextTransfer",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/transfer",
			fn(amfprod.HTTPUEContextTransfer, prod.HTTPUEContextTransfer),
		},

		{
			"N1N2MessageUnSubscribe",
			strings.ToUpper("Delete"),
			"/ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId",
			h.HTTPN1N2MessageUnSubscribe,
		},

		{
			"N1N2MessageTransfer",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/n1-n2-messages",
			h.HTTPN1N2MessageTransfer,
		},

		{
			"N1N2MessageTransferStatus",
			strings.ToUpper("Get"),
			"/ue-contexts/:ueContextId/n1-n2-messages/:n1N2MessageId",
			h.HTTPN1N2MessageTransferStatus,
		},

		{
			"N1N2MessageSubscribe",
			strings.ToUpper("Post"),
			"/ue-contexts/:ueContextId/n1-n2-messages/subscriptions",
			fn(amfprod.HTTPN1N2MessageSubscribe, prod.HTTPN1N2MessageSubscribe),
		},

		{
			"NonUeN2InfoUnSubscribe",
			strings.ToUpper("Delete"),
			"/non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId",
			h.HTTPNonUeN2InfoUnSubscribe,
		},

		{
			"NonUeN2MessageTransfer",
			strings.ToUpper("Post"),
			"/non-ue-n2-messages/transfer",
			h.HTTPNonUeN2MessageTransfer,
		},

		{
			"NonUeN2InfoSubscribe",
			strings.ToUpper("Post"),
			"/non-ue-n2-messages/subscriptions",
			h.HTTPNonUeN2InfoSubscribe,
		},

		{
			"AMFStatusChangeSubscribe",
			strings.ToUpper("Post"),
			"/subscriptions",
			h.HTTPAMFStatusChangeSubscribe,
		},
	}
	service.Group = SERVICE_NAME
	return
}
*/

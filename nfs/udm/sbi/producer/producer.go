package producer

import (
	//	"fmt"
	"etri5gc/nfs/udm/context"

	"etri5gc/fabric/common"

	"etri5gc/openapi"
	openapi_http "etri5gc/openapi/httpdp"
	udmprod "etri5gc/openapi/producers/udm"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "udm.producer"})
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
	services := make([]common.Service, 6, 6)
	services[0] = eventexposureService(prod)
	//	services[1] = communicationService(prod)
	//	services[2] = eventexposureService(prod)
	//	services[3] = locationService(prod)
	//	services[4] = mtService(prod)
	//	services[5] = oamService(prod)
	return services
}

// access to the internal data structures of the udm
func (prod *Producer) udm() *context.UdmContext {
	return prod.backend.Context()
}

func ginHandler(fn openapi.OpenApiProducerHandler, p udmprod.UdmProducer) gin.HandlerFunc {
	return openapi_http.CreateGinHandler(fn, p)
}

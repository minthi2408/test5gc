package producer

import (
//	"fmt"
	"etri5gc/nfs/amf/context"
	//"etri5gc/nfs/amf/ngap"
	"etri5gc/fabric"
	"etri5gc/nfs/amf/sbi/producer/httpcallback"
	"etri5gc/nfs/amf/sbi/producer/eventexposure"
	"etri5gc/nfs/amf/sbi/producer/location"
	"etri5gc/nfs/amf/sbi/producer/communication"
	"etri5gc/nfs/amf/sbi/producer/oam"
	"etri5gc/nfs/amf/sbi/producer/mt"
	"github.com/free5gc/ngap/ngapType"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/openapi/models"
	"github.com/sirupsen/logrus"
)

var log	*logrus.Entry
func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "sbi.producer"})
}

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
	backend		Backend
	ngap		NgapSender
	nas			NasInf
}

func NewProducer(b Backend, ngap NgapSender, nas NasInf) *Producer {
	return &Producer{
		backend: b,
		ngap: ngap,
		nas: nas,
	}
}

func (prod *Producer) Services() []fabric.Service {
	services := make([]fabric.Service,6,6)
	services[0] = httpcallback.MakeService(prod)
	services[1] = communication.MakeService(prod)
	services[2] = eventexposure.MakeService(prod)
	services[3] = location.MakeService(prod)
	services[4] = mt.MakeService(prod)
	services[5] = oam.MakeService(prod)
	return services
}

func (prod *Producer) amf() *context.AMFContext {
	return prod.backend.Context()
}

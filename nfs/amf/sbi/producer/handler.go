package producer

import (
//	"fmt"
	"etri5gc/nfs/amf/context"
	//"etri5gc/nfs/amf/ngap"
	"etri5gc/nfs/amf/sbi/consumer"
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
	Consumer() *consumer.Consumer
	Context() *context.AMFContext
}

type Handler struct {
	backend		Backend
	ngap		NgapSender
	nas			NasInf
}

func NewHandler(b Backend, ngap NgapSender, nas NasInf) *Handler {
	return &Handler{
		backend: b,
		ngap: ngap,
		nas: nas,
	}
}

func (h *Handler) amf() *context.AMFContext {
	return h.backend.Context()
}

package nfselect

import (
	"etri5gc/nfs/amf/context"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi/Nnrf_NFDiscovery"
)

type NfSelector interface {
	AmfSelector
	UdmSelector
	PcfSelector
	SmfSelector
	AusfSelector
	NssiSelector
}

type UdmSelector interface {
	SearchUdmSdmInstance(*context.AmfUe, models.NfType, models.NfType,*Nnrf_NFDiscovery.SearchNFInstancesParamOpts) error
SelectUdm()


}


type PcfSelector interface {
	SelectPcf()
}

type SmfSelector interface {
	SelectSmf(*context.AmfUe, models.AccessType, int32, models.Snssai, string) (*context.SmContext, uint8, error)
}

type AusfSelector interface {
	SelectAusf()
}

type NssiSelector interface {
	SearchNssfNSSelectionInstance(*context.AmfUe, models.NfType,  models.NfType, *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) error 
}

type AmfSelector interface {
	SearchAmfCommunicationInstance(*context.AmfUe, models.NfType, models.NfType, *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) error
}

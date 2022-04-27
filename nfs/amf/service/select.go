package service

import (
	"fmt"
	"etri5gc/nfs/amf/sbi/nfselect"
	"etri5gc/nfs/amf/context"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi/Nnrf_NFDiscovery"
)

type nfSelector struct {
	nf	*AMF
}


func NewNfSelector(nf *AMF) nfselect.NfSelector {
	return &nfSelector{
		nf:	nf,
	}
}


func (s *nfSelector) SelectUdm() {
}

func (s *nfSelector) SelectPcf() {
}

func (s *nfSelector) SelectSmf() {
}

func (s *nfSelector) SelectAusf() {
}

func (s *nfSelector) SelectNssi() {
}

func (s *nfSelector) SearchUdmSdmInstance(ue *context.AmfUe, targetNfType, requestNfType models.NfType,
	param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) error {
	nrf := s.nf.consumer.Nrf()
	resp, localErr := nrf.SendSearchNFInstances(targetNfType, requestNfType, param)
	if localErr != nil {
		return localErr
	}

	// select the first UDM_SDM, TODO: select base on other info
	var sdmUri string
	for _, nfProfile := range resp.NfInstances {
		ue.UdmId = nfProfile.NfInstanceId
		sdmUri = searchNFServiceUri(nfProfile, models.ServiceName_NUDM_SDM, models.NfServiceStatus_REGISTERED)
		if sdmUri != "" {
			break
		}
	}
	ue.NudmSDMUri = sdmUri
	if ue.NudmSDMUri == "" {
		err := fmt.Errorf("AMF can not select an UDM by NRF")
	//	logger.ConsumerLog.Errorf(err.Error())
		return err
	}
	return nil
}
func (s *nfSelector) SearchNssfNSSelectionInstance(ue *context.AmfUe, targetNfType, requestNfType models.NfType,
	param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) error {

	nrf := s.nf.consumer.Nrf()

	resp, localErr := nrf.SendSearchNFInstances(targetNfType, requestNfType, param)
	if localErr != nil {
		return localErr
	}

	// select the first NSSF, TODO: select base on other info
	var nssfUri string
	for _, nfProfile := range resp.NfInstances {
		ue.NssfId = nfProfile.NfInstanceId
		nssfUri = searchNFServiceUri(nfProfile, models.ServiceName_NNSSF_NSSELECTION, models.NfServiceStatus_REGISTERED)
		if nssfUri != "" {
			break
		}
	}
	ue.NssfUri = nssfUri
	if ue.NssfUri == "" {
		return fmt.Errorf("AMF can not select an NSSF by NRF")
	}
	return nil
}

func (s *nfSelector) SearchAmfCommunicationInstance(ue *context.AmfUe, targetNfType,
	requestNfType models.NfType, param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (err error) {
		nrf := s.nf.consumer.Nrf()
	resp, localErr := nrf.SendSearchNFInstances(targetNfType, requestNfType, param)
	if localErr != nil {
		err = localErr
		return
	}

	// select the first AMF, TODO: select base on other info
	var amfUri string
	for _, nfProfile := range resp.NfInstances {
		ue.TargetAmfProfile = &nfProfile
		amfUri = searchNFServiceUri(nfProfile, models.ServiceName_NAMF_COMM, models.NfServiceStatus_REGISTERED)
		if amfUri != "" {
			break
		}
	}
	ue.TargetAmfUri = amfUri
	if ue.TargetAmfUri == "" {
		err = fmt.Errorf("AMF can not select an target AMF by NRF")
	}
	return
}

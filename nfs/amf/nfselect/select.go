package nfselect

import (
	"fmt"
	"time"
	"net/url"
	"etri5gc/nfs/amf/context"
	"etri5gc/nfs/amf/sbi/consumer"
	"etri5gc/nfs/amf"

	"github.com/antihax/optional"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Nnrf_NFDiscovery"
	"github.com/free5gc/nas/nasMessage"
	"github.com/sirupsen/logrus"
)

var log	*logrus.Entry
func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "selector"})
}

type NfSelector struct {
	nf			amf.Backend
	consumer	*consumer.Consumer
}


func NewNfSelector(b amf.Backend) *NfSelector {
	return &NfSelector{
		nf:	b,
		consumer: b.Consumer(),
	}
}


func (s *NfSelector) SelectUdm() {
}

func (s *NfSelector) SelectPcf() {
}

func (s *NfSelector) SelectAusf() {
}

func (s *NfSelector) SelectNssi() {
}

func (s *NfSelector) SearchUdmSdmInstance(ue *context.AmfUe, targetNfType, requestNfType models.NfType,
	param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) error {
	nrf := s.consumer.Nrf()
	resp, localErr := nrf.SendSearchNFInstances("", targetNfType, requestNfType, param)
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
func (s *NfSelector) SearchNssfNSSelectionInstance(ue *context.AmfUe, targetNfType, requestNfType models.NfType,
	param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) error {

	nrf := s.consumer.Nrf()

	resp, localErr := nrf.SendSearchNFInstances("", targetNfType, requestNfType, param)
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

func (s *NfSelector) SearchAmfCommunicationInstance(ue *context.AmfUe, targetNfType,
	requestNfType models.NfType, param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (err error) {
		nrf := s.consumer.Nrf()
	resp, localErr := nrf.SendSearchNFInstances("", targetNfType, requestNfType, param)
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

func (s *NfSelector) SelectSmf(	ue *context.AmfUe,	anType models.AccessType, pduSessionID int32, snssai models.Snssai, dnn string) (*context.SmContext, uint8, error) {
	var (
		smfID  string
		smfUri string
	)

	log.Infof("Select SMF [snssai: %+v, dnn: %+v]", snssai, dnn)

	nsiInformation := ue.GetNsiInformationFromSnssai(anType, snssai)
	if nsiInformation == nil {
		if ue.NssfUri == "" {
			// TODO: Set a timeout of NSSF Selection or will starvation here
			for {
				if err := s.SearchNssfNSSelectionInstance(ue, models.NfType_NSSF,
					models.NfType_AMF, nil); err != nil {
					log.Errorf("AMF can not select an NSSF Instance by NRF[Error: %+v]", err)
					time.Sleep(2 * time.Second)
				} else {
					break
				}
			}
		}

		response, problemDetails, err := s.consumer.Nssf().NSSelectionGetForPduSession(ue, snssai)
		if err != nil {
			err = fmt.Errorf("NSSelection Get Error[%+v]", err)
			return nil, nasMessage.Cause5GMMPayloadWasNotForwarded, err
		} else if problemDetails != nil {
			err = fmt.Errorf("NSSelection Get Failed Problem[%+v]", problemDetails)
			return nil, nasMessage.Cause5GMMPayloadWasNotForwarded, err
		}
		nsiInformation = response.NsiInformation
	}

	smContext := context.NewSmContext(pduSessionID)
	smContext.SetSnssai(snssai)
	smContext.SetDnn(dnn)
	smContext.SetAccessType(anType)
	nrfUri := s.nf.Context().NrfUri()
	if nsiInformation == nil {
		log.Warnf("nsiInformation is still nil, use default NRF[%s]", nrfUri)
	} else {
		smContext.SetNsInstance(nsiInformation.NsiId)
		nrfApiUri, err := url.Parse(nsiInformation.NrfId)
		if err != nil {
			log.Errorf("Parse NRF URI error, use default NRF[%s]", nrfUri)
		} else {
			nrfUri = fmt.Sprintf("%s://%s", nrfApiUri.Scheme, nrfApiUri.Host)
		}
	}

	param := Nnrf_NFDiscovery.SearchNFInstancesParamOpts{
		ServiceNames: optional.NewInterface([]models.ServiceName{models.ServiceName_NSMF_PDUSESSION}),
		Dnn:          optional.NewString(dnn),
		Snssais:      optional.NewInterface(openapi.MarshToJsonString([]models.Snssai{snssai})),
	}
	if ue.PlmnId.Mcc != "" {
		param.TargetPlmnList = optional.NewInterface(openapi.MarshToJsonString(ue.PlmnId))
	}
	locality := s.nf.Context().Locality()
	if locality != "" {
		param.PreferredLocality = optional.NewString(locality)
	}

	log.Debugf("Search SMF from NRF[%s]", nrfUri)

	result, err := s.consumer.Nrf().SendSearchNFInstances(nrfUri, models.NfType_SMF, models.NfType_AMF, &param)
	if err != nil {
		return nil, nasMessage.Cause5GMMPayloadWasNotForwarded, err
	}

	if len(result.NfInstances) == 0 {
		err = fmt.Errorf("DNN[%s] is not supported or not subscribed in the slice[Snssai: %+v]", dnn, snssai)
		return nil, nasMessage.Cause5GMMDNNNotSupportedOrNotSubscribedInTheSlice, err
	}

	// select the first SMF, TODO: select base on other info
	for _, nfProfile := range result.NfInstances {
		smfUri = searchNFServiceUri(nfProfile, models.ServiceName_NSMF_PDUSESSION, models.NfServiceStatus_REGISTERED)
		if smfUri != "" {
			break
		}
	}
	smContext.SetSmfID(smfID)
	smContext.SetSmfUri(smfUri)
	return smContext, 0, nil
}



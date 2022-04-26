package sbi
import (
	"fmt"
	"time"
	"strings"
	"net/http"
	"context"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/openapi/Nnrf_NFManagement"
	"github.com/free5gc/openapi/Nnrf_NFDiscovery"
)

type NrfConsumer interface {
	SendRegisterNFInstance() (string, string, error) // Register NF to the NRF
	SendDeregisterNFInstance() (*models.ProblemDetails, error)
	SendSearchNFInstances(models.NfType, models.NfType, *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (models.SearchResult, error)
}

type NFInterface interface {
	NrfUri() string
	NfId() string
	BuildProfile() (*models.NfProfile, error)
}

type nrfClient struct {
	nf		NFInterface
}

func NewNrfConsumer(nf NFInterface) NrfConsumer {
	return &nrfClient{
		nf:	nf,
	}
}

func (c *nrfClient) man_client() *Nnrf_NFManagement.APIClient {
	conf := Nnrf_NFManagement.NewConfiguration()
	conf.SetBasePath(c.nf.NrfUri())
	return Nnrf_NFManagement.NewAPIClient(conf)

}

func (c *nrfClient) disc_client() *Nnrf_NFDiscovery.APIClient {
	conf := Nnrf_NFDiscovery.NewConfiguration()
	conf.SetBasePath(c.nf.NrfUri())
	return Nnrf_NFDiscovery.NewAPIClient(conf)
}


func (c *nrfClient) SendRegisterNFInstance() (
	resouceNrfUri string, retrieveNfInstanceId string, err error) {
		//TODO: tungtq - should move registration retrying to outer layer

	// Set client and set url
	client := c.man_client()	
	var res *http.Response
	var profile *models.NfProfile
	if profile, err = c.nf.BuildProfile(); err != nil {
		return 
	}

	for {
		_, res, err = client.NFInstanceIDDocumentApi.RegisterNFInstance(context.TODO(), c.nf.NfId(), *profile)
		if err != nil || res == nil {
			//fmt.Println(fmt.Errorf("AMF register to NRF Error[%s]", err.Error()))
			time.Sleep(2 * time.Second)
			continue
		}
		defer func() {
			if bodyCloseErr := res.Body.Close(); bodyCloseErr != nil {
				err = fmt.Errorf("SearchNFInstances' response body cannot close: %+w", bodyCloseErr)
			}
		}()
		status := res.StatusCode
		if status == http.StatusOK {
			// NFUpdate
			break
		} else if status == http.StatusCreated {
			// NFRegister
			resourceUri := res.Header.Get("Location")
			resouceNrfUri = resourceUri[:strings.Index(resourceUri, "/nnrf-nfm/")]
			retrieveNfInstanceId = resourceUri[strings.LastIndex(resourceUri, "/")+1:]
			break
		} else {
			//fmt.Println(fmt.Errorf("handler returned wrong status code %d", status))
			//fmt.Println(fmt.Errorf("NRF return wrong status code %d", status))
		}
	}
	return resouceNrfUri, retrieveNfInstanceId, err
}

func (c *nrfClient) SendDeregisterNFInstance() (problemDetails *models.ProblemDetails, err error) {
	//logger.ConsumerLog.Infof("[AMF] Send Deregister NFInstance")

	// Set client and set url
	client := c.man_client()

	var res *http.Response

	res, err = client.NFInstanceIDDocumentApi.DeregisterNFInstance(context.Background(), c.nf.NfId())
	if err == nil {
		return
	} else if res != nil {
		defer func() {
			if bodyCloseErr := res.Body.Close(); bodyCloseErr != nil {
				err = fmt.Errorf("SearchNFInstances' response body cannot close: %+w", bodyCloseErr)
			}
		}()
		if res.Status != err.Error() {
			return
		}
		problem := err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("server no response")
	}
	return
}

func (c *nrfClient) SendSearchNFInstances(targetNfType, requestNfType models.NfType,
	param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (models.SearchResult, error) {
	// Set client and set url
	client := c.disc_client()

	result, res, err := client.NFInstancesStoreApi.SearchNFInstances(context.TODO(), targetNfType, requestNfType, param)
	if res != nil && res.StatusCode == http.StatusTemporaryRedirect {
		err = fmt.Errorf("Temporary Redirect For Non NRF Consumer")
	}
	defer func() {
		if bodyCloseErr := res.Body.Close(); bodyCloseErr != nil {
			err = fmt.Errorf("SearchNFInstances' response body cannot close: %+w", bodyCloseErr)
		}
	}()
	return result, err
}
/*
func (c *nrfClient) SearchUdmSdmInstance(ue *amf_context.AmfUe, nrfUri string, targetNfType, requestNfType models.NfType,
	param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) error {
	resp, localErr := SendSearchNFInstances(nrfUri, targetNfType, requestNfType, param)
	if localErr != nil {
		return localErr
	}

	// select the first UDM_SDM, TODO: select base on other info
	var sdmUri string
	for _, nfProfile := range resp.NfInstances {
		ue.UdmId = nfProfile.NfInstanceId
		sdmUri = util.SearchNFServiceUri(nfProfile, models.ServiceName_NUDM_SDM, models.NfServiceStatus_REGISTERED)
		if sdmUri != "" {
			break
		}
	}
	ue.NudmSDMUri = sdmUri
	if ue.NudmSDMUri == "" {
		err := fmt.Errorf("AMF can not select an UDM by NRF")
		logger.ConsumerLog.Errorf(err.Error())
		return err
	}
	return nil
}

func (c *nrfClient) SearchNssfNSSelectionInstance(ue *amf_context.AmfUe, nrfUri string, targetNfType, requestNfType models.NfType,
	param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) error {
	resp, localErr := SendSearchNFInstances(nrfUri, targetNfType, requestNfType, param)
	if localErr != nil {
		return localErr
	}

	// select the first NSSF, TODO: select base on other info
	var nssfUri string
	for _, nfProfile := range resp.NfInstances {
		ue.NssfId = nfProfile.NfInstanceId
		nssfUri = util.SearchNFServiceUri(nfProfile, models.ServiceName_NNSSF_NSSELECTION, models.NfServiceStatus_REGISTERED)
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

func (c *nrfClient) SearchAmfCommunicationInstance(ue *amf_context.AmfUe, nrfUri string, targetNfType,
	requestNfType models.NfType, param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (err error) {
	resp, localErr := SendSearchNFInstances(nrfUri, targetNfType, requestNfType, param)
	if localErr != nil {
		err = localErr
		return
	}

	// select the first AMF, TODO: select base on other info
	var amfUri string
	for _, nfProfile := range resp.NfInstances {
		ue.TargetAmfProfile = &nfProfile
		amfUri = util.SearchNFServiceUri(nfProfile, models.ServiceName_NAMF_COMM, models.NfServiceStatus_REGISTERED)
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
*/

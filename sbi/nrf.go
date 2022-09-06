package sbi

import (
	"fmt"
	//	"time"
	"context"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/Nnrf_NFDiscovery"
	"github.com/free5gc/openapi/Nnrf_NFManagement"
	"github.com/free5gc/openapi/models"
	"net/http"
	"strings"
)

type NrfConsumer interface {
	SendRegisterNFInstance() (string, string, error) // Register NF to the NRF
	SendDeregisterNFInstance() (*models.ProblemDetails, error)
	SendSearchNFInstances(string, models.NfType, models.NfType, *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (models.SearchResult, error)
}

type NFInterface interface {
	NrfUri() string
	NfId() string
	BuildProfile() (*models.NfProfile, error)
}

type nrfClient struct {
	nf NFInterface
}

func NewNrfConsumer(nf NFInterface) NrfConsumer {
	return &nrfClient{
		nf: nf,
	}
}

func (c *nrfClient) man_client() *Nnrf_NFManagement.APIClient {
	conf := Nnrf_NFManagement.NewConfiguration()
	conf.SetBasePath(c.nf.NrfUri())
	return Nnrf_NFManagement.NewAPIClient(conf)

}

func (c *nrfClient) disc_client(uri string) *Nnrf_NFDiscovery.APIClient {
	conf := Nnrf_NFDiscovery.NewConfiguration()
	if len(uri) > 0 {
		conf.SetBasePath(uri)
	} else {
		conf.SetBasePath(c.nf.NrfUri())
	}
	return Nnrf_NFDiscovery.NewAPIClient(conf)
}

func (c *nrfClient) SendRegisterNFInstance() (
	resouceNrfUri string, retrieveNfInstanceId string, err error) {

	// Set client and set url
	client := c.man_client()
	var res *http.Response
	var profile *models.NfProfile
	if profile, err = c.nf.BuildProfile(); err != nil {
		return
	}

	log.Info("Sending a registration request")
	_, res, err = client.NFInstanceIDDocumentApi.RegisterNFInstance(context.TODO(), c.nf.NfId(), *profile)

	log.Info("get a respone")
	if err != nil {
		return
	}

	if res == nil {
		err = fmt.Errorf("NRF response is empty")
		return
	}

	defer func() {
		if bodyCloseErr := res.Body.Close(); bodyCloseErr != nil {
			err = fmt.Errorf("SearchNFInstances' response body cannot close: %+w", bodyCloseErr)
		}
	}()
	status := res.StatusCode
	if status == http.StatusOK {
		log.Info("Profile is updated at NRF")
		// NFUpdate
	} else if status == http.StatusCreated {
		// NFRegister
		resourceUri := res.Header.Get("Location")
		resouceNrfUri = resourceUri[:strings.Index(resourceUri, "/nnrf-nfm/")]
		retrieveNfInstanceId = resourceUri[strings.LastIndex(resourceUri, "/")+1:]
	} else {
		log.Warnf("NRF return wrong status code %d", status)
	}
	return
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

func (c *nrfClient) SendSearchNFInstances(nrfuri string, targetNfType, requestNfType models.NfType,
	param *Nnrf_NFDiscovery.SearchNFInstancesParamOpts) (models.SearchResult, error) {
	// Set client and set url
	client := c.disc_client(nrfuri)

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

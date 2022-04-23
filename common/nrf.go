package common
import (
	"github.com/free5gc/openapi/models"
)

type interface NrfConsumer {
	Register() (string, string, error) // Register NF to the NRF
	Deregister() error //Deregister from the NRF
	Discover() error // query the NRF for an NF
}

//build an NFProfile, return its id, the profile, and an error
type ProfileBuilderFn func() (string, models.NFProfile, error)

type nrfClient struct {
	nrf		string
	fn		ProfileBuilderFn
}

func NewNrfConsumer(nrf_uri string, builder ProfileBuilderFn) NrfConsumer {
	return &nrfClient{
		nrf:	nrf_uri,
		fn:		builder,
	}
}

func (c *nrfClient) Register() (string, string, error) {
	//TODO:
	// 1. build profile
	// 2. create a client
	// 3. send registration
	// 4. extract response
	return "","", nil
}

func (c *nrfClient) Deregister() error {
	//TODO:
	//1. create a client
	//2. send deregistration
	return nil
}

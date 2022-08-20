/*
 * Namf_Location
 *
 * AMF Location Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location

import (
	"strings"
	"etri5gc/fabric"
	"github.com/free5gc/util/httpwrapper"
)

const (
	SERVICE_NAME="/namf-loc/v1"
)

type Backend interface {
	HandleProvideLocationInfoRequest(*httpwrapper.Request) *httpwrapper.Response
}

type Handler struct {
	prod	Backend	
}

func MakeService(p Backend) (service fabric.HttpService) {
	h := &Handler{
		prod:	p,
	}
	service.Routes = fabric.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			fabric.HttpIndexHandler,
		},

		{
			"ProvideLocationInfo",
			strings.ToUpper("Post"),
			"/:ueContextId/provide-loc-info",
			h.HTTPProvideLocationInfo,
		},

		{
			"ProvidePositioningInfo",
			strings.ToUpper("Post"),
			"/:ueContextId/provide-pos-info",
			h.HTTPProvidePositioningInfo,
		},
	}
	service.Group =  SERVICE_NAME
	return
}	
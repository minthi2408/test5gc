/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eventexposure

import (
	//"net/http"
	"etri5gc/fabric/httpdp"
	"github.com/free5gc/util/httpwrapper"
	"strings"
)

const (
	SERVICE_NAME = "/namf-evts/v1"
)

type Backend interface {
	HandleCreateAMFEventSubscription(*httpwrapper.Request) *httpwrapper.Response
	HandleDeleteAMFEventSubscription(*httpwrapper.Request) *httpwrapper.Response
	HandleModifyAMFEventSubscription(*httpwrapper.Request) *httpwrapper.Response
}

type Handler struct {
	prod Backend
}

func MakeService(p Backend) (service httpdp.HttpService) {
	h := &Handler{
		prod: p,
	}
	service.Routes = httpdp.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			httpdp.HttpIndexHandler,
		},

		{
			"HTTPDeleteSubscription",
			strings.ToUpper("Delete"),
			"/subscriptions/:subscriptionId",
			h.HTTPDeleteSubscription,
		},

		{
			"HTTPModifySubscription",
			strings.ToUpper("Patch"),
			"/subscriptions/:subscriptionId",
			h.HTTPModifySubscription,
		},

		{
			"HTTPCreateSubscription",
			strings.ToUpper("Post"),
			"/subscriptions",
			h.HTTPCreateSubscription,
		},
	}
	service.Group = SERVICE_NAME
	return
}

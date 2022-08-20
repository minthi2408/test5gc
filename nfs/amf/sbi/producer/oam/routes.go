package oam

import (
	"etri5gc/fabric"
	"github.com/free5gc/util/httpwrapper"
)

const (
	SERVICE_NAME="/namf-oam/v1"
)

type Backend interface {
	HandleOAMRegisteredUEContext(*httpwrapper.Request) *httpwrapper.Response
}

type Handler struct {
	prod	Backend	
}

func MakeService(p Backend) (service fabric.HttpService) {
	h := Handler{
		prod: p,
	}

	service.Routes  = fabric.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			fabric.HttpIndexHandler,
		},

		{
			"Registered UE Context",
			"GET",
			"/registered-ue-context",
			h.HTTPRegisteredUEContext,
		},

		{
			"Individual Registered UE Context",
			"GET",
			"/registered-ue-context/:supi",
			h.HTTPRegisteredUEContext,
		},
	}
	service.Group = SERVICE_NAME
	return
}

package oam

import (
	
	"etri5gc/nfs/amf/sbi/producer"
	"etri5gc/sbi"


)

const (
	SERVICE_NAME="/namf-oam/v1"
)

type Handler struct {
	prod	*producer.Handler
}

func MakeRoutes(p *producer.Handler) (string, sbi.Routes) {
	h := Handler{
		prod: p,
	}

	var routes = sbi.Routes{
		{
			"Index",
			"GET",
			"/",
			sbi.IndexHandler,
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
	return SERVICE_NAME, routes
}

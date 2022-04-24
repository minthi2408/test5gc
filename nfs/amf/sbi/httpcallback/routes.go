package httpcallback

import (
	"strings"

	"etri5gc/nfs/amf/sbi/producer"
	"etri5gc/sbi"
)

const (
	SERVICE_NAME="/namf-callback/v1"
)

type Handler struct {
	prod	*producer.Handler
}

func MakeRoutes(p *producer.Handler) (string, sbi.Routes) {
	h := &Handler{
		prod: 	p,
	}
	var routes = sbi.Routes{
		{
			"Index",
			"GET",
			"/",
			sbi.IndexHandler,
		},

		{
			"SmContextStatusNotify",
			strings.ToUpper("Post"),
			"/smContextStatus/:guti/:pduSessionId",
			h.HTTPSmContextStatusNotify,
		},

		{
			"AmPolicyControlUpdateNotifyUpdate",
			strings.ToUpper("Post"),
			"/am-policy/:polAssoId/update",
			h.HTTPAmPolicyControlUpdateNotifyUpdate,
		},

		{
			"AmPolicyControlUpdateNotifyTerminate",
			strings.ToUpper("Post"),
			"/am-policy/:polAssoId/terminate",
			h.HTTPAmPolicyControlUpdateNotifyTerminate,
		},

		{
			"N1MessageNotify",
			strings.ToUpper("Post"),
			"/n1-message-notify",
			h.HTTPN1MessageNotify,
		},
	}
	return SERVICE_NAME, routes
}

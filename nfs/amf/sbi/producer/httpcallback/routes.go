package httpcallback

import (
	"strings"
	"etri5gc/fabric"
	"github.com/free5gc/util/httpwrapper"
)

const (
	SERVICE_NAME="/namf-callback/v1"
)

type Backend interface {
	HandleSmContextStatusNotify(*httpwrapper.Request) *httpwrapper.Response
	HandleAmPolicyControlUpdateNotifyUpdate(*httpwrapper.Request) *httpwrapper.Response
	HandleAmPolicyControlUpdateNotifyTerminate(*httpwrapper.Request) *httpwrapper.Response
	HandleN1MessageNotify(*httpwrapper.Request) *httpwrapper.Response 

}


type Handler struct {
	prod	Backend	
}

func MakeService(p Backend) (service fabric.HttpService) {
	h := &Handler{
		prod: 	p,
	}
	service.Routes = fabric.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			fabric.HttpIndexHandler,
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
	service.Group = SERVICE_NAME
	return
}

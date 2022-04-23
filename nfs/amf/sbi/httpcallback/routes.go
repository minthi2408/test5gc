package httpcallback

import (
	"net/http"
	"strings"

	"etri5gc/nfs/amf/sbi/producer"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/free5gc/amf/internal/logger"
	logger_util "github.com/free5gc/util/logger"
)

const (
	SERVICE_NAME="/namf-callback/v1"
)


func MakeRoutes(h *producer.Handler) (string, common.Routes) {
	h := &Handler{app: app}

	var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		common.IndexHandler,
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

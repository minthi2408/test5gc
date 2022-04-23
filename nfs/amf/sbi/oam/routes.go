package oam

import (
	"net/http"
	
	"etri5gc/nfs/amf/sbi/producer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/free5gc/amf/internal/logger"
	logger_util "github.com/free5gc/util/logger"
)

const (
	SERVICE_NAME="/namf-oam/v1"
)


func MakeRoutes(h *producer.Handler) (string, common.Routes) {
	h := &Handler{app: app}

	var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
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

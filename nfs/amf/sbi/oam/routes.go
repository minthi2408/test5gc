package oam

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/free5gc/amf/internal/logger"
	logger_util "github.com/free5gc/util/logger"
)

const (
	SERVICE_NAME="/namf-oam/v1"
)

type Handler struct {
	app		amf.App	
}

func MakeRoutes(app *amf.App) (string, common.Routes) {
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
		HTTPRegisteredUEContext,
	},

	{
		"Individual Registered UE Context",
		"GET",
		"/registered-ue-context/:supi",
		HTTPRegisteredUEContext,
	},
	}
	return SERVICE_NAME, routes
}

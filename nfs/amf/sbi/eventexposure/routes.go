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
	"net/http"
	"strings"

	"etri5gc/nfs/amf/sbi/producer"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/amf/internal/logger"
	logger_util "github.com/free5gc/util/logger"
)

const (
	SERVICE_NAME="/namf-evts/v1"
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
	return SERVICE_NAME, routes
}

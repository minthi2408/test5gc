/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package communication

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

// AMFStatusChangeSubscribeModify - Namf_Communication AMF Status Change Subscribe Modify service Operation
func (h *Handler) HTTPAMFStatusChangeSubscribeModify(c *gin.Context) {
	var subscriptionData models.SubscriptionData

	requestBody, err := c.GetRawData()
	if err != nil {
		//		logger.CommLog.Errorf("Get Request Body error: %+v", err)
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&subscriptionData, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		//logger.CommLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := httpwrapper.NewRequest(c.Request, subscriptionData)
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	rsp := h.prod.HandleAMFStatusChangeSubscribeModify(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		//logger.CommLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// AMFStatusChangeUnSubscribe - Namf_Communication AMF Status Change UnSubscribe service Operation
func (h *Handler) HTTPAMFStatusChangeUnSubscribe(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	rsp := h.prod.HandleAMFStatusChangeUnSubscribeRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		//		logger.CommLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

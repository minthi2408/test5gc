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

)

// NonUeN2InfoSubscribe - Namf_Communication Non UE N2 Info Subscribe service Operation
func (h *Handler) HTTPNonUeN2InfoSubscribe(c *gin.Context) {
//	logger.CommLog.Warnf("Handle Non Ue N2 Info Subscribe is not implemented.")
	c.JSON(http.StatusOK, gin.H{})
}
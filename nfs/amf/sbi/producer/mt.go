package producer

import (
	"net/http"

//	"etri5gc/nfs/amf/context"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

func (h *Producer) HandleProvideDomainSelectionInfoRequest(request *httpwrapper.Request) *httpwrapper.Response {
	log.Info("Handle Provide Domain Selection Info Request")

	ueContextID := request.Params["ueContextId"]
	infoClassQuery := request.Query.Get("info-class")
	supportedFeaturesQuery := request.Query.Get("supported-features")

	if ueContextInfo, prob := h.doProvideDomainSelectionInfo(ueContextID,
		infoClassQuery, supportedFeaturesQuery); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, ueContextInfo)
	}
}

func (h *Producer) doProvideDomainSelectionInfo(ueContextID string, infoClassQuery string, supportedFeaturesQuery string) (
	*models.UeContextInfo, *models.ProblemDetails) {
	amf := h.amf()

	if ue, ok := amf.AmfUeFindByUeContextID(ueContextID); !ok {
		return nil, &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
	} else {
		return  ue.GetContextInfo(infoClassQuery, supportedFeaturesQuery), nil
	}
}

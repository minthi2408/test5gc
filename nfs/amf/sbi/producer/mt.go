package producer

import (
	"net/http"

//	"etri5gc/nfs/amf/context"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

func (h *Handler) HandleProvideDomainSelectionInfoRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.MtLog.Info("Handle Provide Domain Selection Info Request")

	ueContextID := request.Params["ueContextId"]
	infoClassQuery := request.Query.Get("info-class")
	supportedFeaturesQuery := request.Query.Get("supported-features")

	if ueContextInfo, prob := h.ProvideDomainSelectionInfoProcedure(ueContextID,
		infoClassQuery, supportedFeaturesQuery); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, ueContextInfo)
	}
}

func (h *Handler) ProvideDomainSelectionInfoProcedure(ueContextID string, infoClassQuery string, supportedFeaturesQuery string) (
	*models.UeContextInfo, *models.ProblemDetails) {
		/*
	amf := h.backend.Context()

	if ue, ok := amf.AmfUeFindByUeContextID(ueContextID); !ok {
		return nil, &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
	} else {

		ueContextInfo := new(models.UeContextInfo)

		// TODO: Error Status 307, 403 in TS29.518 Table 6.3.3.3.3.1-3
		anType := ue.GetAnType()
		if anType != "" && infoClassQuery != "" {
			ranUe := ue.RanUe[anType]
			ueContextInfo.AccessType = anType
			ueContextInfo.LastActTime = ranUe.LastActTime
			ueContextInfo.RatType = ue.RatType
			ueContextInfo.SupportedFeatures = ranUe.SupportedFeatures
			ueContextInfo.SupportVoPS = ranUe.SupportVoPS
			ueContextInfo.SupportVoPSn3gpp = ranUe.SupportVoPSn3gpp
		}

		return ueContextInfo, nil
	}
	*/
	return nil, nil
}

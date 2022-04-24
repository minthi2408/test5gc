package producer

import (
	"net/http"

//	"etri5gc/nfs/amf/context"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

func (h *Handler) HandleProvideLocationInfoRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.ProducerLog.Info("Handle Provide Location Info Request")

	requestLocInfo := request.Body.(models.RequestLocInfo)
	ueContextID := request.Params["ueContextId"]

	if provideLocInfo, prob := h.ProvideLocationInfoProcedure(requestLocInfo, ueContextID); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, provideLocInfo)
	}
}

func (h *Handler) ProvideLocationInfoProcedure(requestLocInfo models.RequestLocInfo, ueContextID string) (
	*models.ProvideLocInfo, *models.ProblemDetails) {
		/*
	amf := h.backend.Context()

	if ue, ok := amf.AmfUeFindByUeContextID(ueContextID); !ok {
		return nil, &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
	} else {

		anType := ue.GetAnType()
		if anType == "" {
			return nil, &models.ProblemDetails{
				Status: http.StatusNotFound,
				Cause:  "CONTEXT_NOT_FOUND",
			}
		}

		provideLocInfo := new(models.ProvideLocInfo)

		ranUe := ue.RanUe[anType]
		if requestLocInfo.Req5gsLoc || requestLocInfo.ReqCurrentLoc {
			provideLocInfo.CurrentLoc = true
			provideLocInfo.Location = &ue.Location
		}

		if requestLocInfo.ReqRatType {
			provideLocInfo.RatType = ue.RatType
		}

		if requestLocInfo.ReqTimeZone {
			provideLocInfo.Timezone = ue.TimeZone
		}

		if requestLocInfo.SupportedFeatures != "" {
			provideLocInfo.SupportedFeatures = ranUe.SupportedFeatures
		}
		return provideLocInfo, nil
	}
	*/
	return nil, nil
}

package producer

import (
	"net/http"

//	"etri5gc/nfs/amf/context"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

func (h *Handler) HandleProvideLocationInfoRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.ProducerLog.Info("Handle Provide Location Info Request")

	req := request.Body.(models.RequestLocInfo)
	ueId := request.Params["ueContextId"]

	if locinfo, prob := h.doProvideLocationInfo(req, ueId); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, locinfo)
	}
}

func (h *Handler) doProvideLocationInfo(req models.RequestLocInfo, ueId string) (
	*models.ProvideLocInfo, *models.ProblemDetails) {
	amf := h.backend.Context()

	if ue, ok := amf.AmfUeFindByUeContextID(ueId); ok {
		if locinfo := ue.GetLocInfo(&req); locinfo != nil {
			return locinfo, nil
		} 
	}

	return nil, &models.ProblemDetails{
		Status: http.StatusNotFound,
		Cause:  "CONTEXT_NOT_FOUND",
	}
}

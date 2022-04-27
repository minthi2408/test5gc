package producer

import (
	"net/http"
//	"strconv"
//	"time"

//	"etri5gc/nfs/amf/context"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

func (h *Handler) HandleCreateAMFEventSubscription(request *httpwrapper.Request) *httpwrapper.Response {
	reqsub := request.Body.(models.AmfCreateEventSubscription)

	//ressub, prob := CreateAMFEventSubscriptionProcedure(reqsub)
	ressub, prob := h.amf().CreateAMFEventSub(reqsub)

	if ressub != nil {
		return httpwrapper.NewResponse(http.StatusCreated, nil, ressub)
	} else if prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		prob = &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "UNSPECIFIED_NF_FAILURE",
		}
		return httpwrapper.NewResponse(http.StatusInternalServerError, nil, prob)
	}
	return nil
}

func (h *Handler) HandleDeleteAMFEventSubscription(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.EeLog.Infoln("Handle Delete AMF Event Subscription")

	subId := request.Params["subscriptionId"]

	if prob := h.amf().DeleteAMFEventSub(subId); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusOK, nil, nil)
	}
}

func (h *Handler) HandleModifyAMFEventSubscription(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.EeLog.Infoln("Handle Modify AMF Event Subscription")

	subId := request.Params["subscriptionId"]
	modreq := request.Body.(models.ModifySubscriptionRequest)

	upsub, prob := h.amf().ModifyAMFEventSub(subId, modreq)
	if upsub != nil {
		return httpwrapper.NewResponse(http.StatusOK, nil, upsub)
	} else if prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		prob = &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "UNSPECIFIED_NF_FAILURE",
		}
		return httpwrapper.NewResponse(http.StatusInternalServerError, nil, prob)
	}
}


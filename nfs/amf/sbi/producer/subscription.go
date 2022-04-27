package producer

import (
	"net/http"
//	"reflect"

//	"etri5gc/nfs/amf/context"

	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

// TS 29.518 5.2.2.5.1
func (h *Handler) HandleAMFStatusChangeSubscribeRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.CommLog.Info("Handle AMF Status Change Subscribe Request")

	datreq := request.Body.(models.SubscriptionData)

	if datres, locheader, prob := h.amf().AMFStatusChangeSubscribe(datreq); prob!=nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {

		headers := http.Header{
			"Location": {locheader},
		}
		return httpwrapper.NewResponse(http.StatusCreated, headers, datres)
	}
}

// TS 29.518 5.2.2.5.2
func (h *Handler) HandleAMFStatusChangeUnSubscribeRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.CommLog.Info("Handle AMF Status Change UnSubscribe Request")

	subId := request.Params["subscriptionId"]

	if prob := h.amf().AMFStatusChangeUnSubscribe(subId); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
}


// TS 29.518 5.2.2.5.1.3
func (h *Handler) HandleAMFStatusChangeSubscribeModify(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.CommLog.Info("Handle AMF Status Change Subscribe Modify Request")

	dat := request.Body.(models.SubscriptionData)
	subId := request.Params["subscriptionId"]

	if updat, prob := h.amf().AMFStatusChangeSubscribeModify(subId, dat); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusAccepted, nil, updat)
	}
}


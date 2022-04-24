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

	subscriptionDataReq := request.Body.(models.SubscriptionData)

	if subscriptionDataRsp, locationHeader, prob := h.AMFStatusChangeSubscribeProcedure(subscriptionDataReq); prob!=nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {

		headers := http.Header{
			"Location": {locationHeader},
		}
		return httpwrapper.NewResponse(http.StatusCreated, headers, subscriptionDataRsp)
	}
}

func (h *Handler) AMFStatusChangeSubscribeProcedure(subscriptionDataReq models.SubscriptionData) (
	subscriptionDataRsp models.SubscriptionData, locationHeader string, problemDetails *models.ProblemDetails) {
		/*
	amf := h.backend.Context() 

	for _, guami := range subscriptionDataReq.GuamiList {
		for _, servedGumi := range amfSelf.ServedGuamiList {
			if reflect.DeepEqual(guami, servedGumi) {
				// AMF status is available
				subscriptionDataRsp.GuamiList = append(subscriptionDataRsp.GuamiList, guami)
			}
		}
	}

	if subscriptionDataRsp.GuamiList != nil {
		newSubscriptionID := amf.NewAMFStatusSubscription(subscriptionDataReq)
		locationHeader = subscriptionDataReq.AmfStatusUri + "/" + newSubscriptionID
		//logger.CommLog.Infof("new AMF Status Subscription[%s]", newSubscriptionID)
		return
	} else {
		problemDetails = &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "UNSPECIFIED",
		}
		return
	}
	*/

	return
}

// TS 29.518 5.2.2.5.2
func (h *Handler) HandleAMFStatusChangeUnSubscribeRequest(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.CommLog.Info("Handle AMF Status Change UnSubscribe Request")

	subscriptionID := request.Params["subscriptionId"]

	if prob := h.AMFStatusChangeUnSubscribeProcedure(subscriptionID); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
}

func (h *Handler) AMFStatusChangeUnSubscribeProcedure(subscriptionID string) *models.ProblemDetails {
	/*
	amf := h.backend.Context()

	if _, ok := amf.FindAMFStatusSubscription(subscriptionID); !ok {
		return &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "SUBSCRIPTION_NOT_FOUND",
		}
	} else {
		//logger.CommLog.Debugf("Delete AMF status subscription[%s]", subscriptionID)
		amf.DeleteAMFStatusSubscription(subscriptionID)
	}
	return nil
	*/
	return nil
}

// TS 29.518 5.2.2.5.1.3
func (h *Handler) HandleAMFStatusChangeSubscribeModify(request *httpwrapper.Request) *httpwrapper.Response {
	//logger.CommLog.Info("Handle AMF Status Change Subscribe Modify Request")

	updateSubscriptionData := request.Body.(models.SubscriptionData)
	subscriptionID := request.Params["subscriptionId"]

	if updatedSubscriptionData, prob := h.AMFStatusChangeSubscribeModifyProcedure(subscriptionID,
		updateSubscriptionData); prob != nil {
		return httpwrapper.NewResponse(int(prob.Status), nil, prob)
	} else {
		return httpwrapper.NewResponse(http.StatusAccepted, nil, updatedSubscriptionData)
	}
}

func (h *Handler) AMFStatusChangeSubscribeModifyProcedure(subscriptionID string, subscriptionData models.SubscriptionData) (
	*models.SubscriptionData, *models.ProblemDetails) {
		/*
	amf := h.backend.Context()

	if currentSubscriptionData, ok := amfSelf.FindAMFStatusSubscription(subscriptionID); !ok {
		return nil, &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "Forbidden",
		}
	} else {
		//logger.CommLog.Debugf("Modify AMF status subscription[%s]", subscriptionID)

		currentSubscriptionData.GuamiList = currentSubscriptionData.GuamiList[:0]

		currentSubscriptionData.GuamiList = append(currentSubscriptionData.GuamiList, subscriptionData.GuamiList...)
		currentSubscriptionData.AmfStatusUri = subscriptionData.AmfStatusUri

		amf.AMFStatusSubscriptions.Store(subscriptionID, currentSubscriptionData)
		return currentSubscriptionData, nil
	}
	*/
	return nil, nil
}

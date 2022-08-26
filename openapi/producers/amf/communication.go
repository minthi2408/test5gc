package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

// AMFStatusChangeSubscribeModify - Namf_Communication AMF Status Change Subscribe Modify service Operation
func (producer *amfProducerImpl) HTTPAMFStatusChangeSubscribeModify(ctx openapi.RequestContext) {
	var subData models.SubscriptionData
	ctx.Request().Body = &subData

	//decode the request (and body)
	ctx.DecodeRequest()
}

/*
		prob := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: fmt.Sprintf("[Request Body] %s", err.Error()),
		}
		c.JSON(http.StatusBadRequest, prob)
		return
	}

	// call the application handler
	resp := ctx.AppHandler()(ctx, req)

	//encode the response
	if err := producer.encoding.Encode(resp); err != nil {
		prob := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, prob)
	} else {
		//ctx.Data(resp.Status, "application/json", responseBody)
	}
}
*/

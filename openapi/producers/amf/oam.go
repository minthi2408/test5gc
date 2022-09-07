package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func RegisteredUEContext(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {
	amfproducer := handler.(AmfProducer)

	var prob *models.ProblemDetails
    var result []byte
	supi := ctx.Param("supi")

	//call the application handler
	result, prob = amfproducer.HandleRegisteredUEContext(supi)

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(202, result)
	}

    return
}


package udm

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func Update(ctx openapi.RequestContext, handler interface{}) (resp openapi.Response) {

	udmproducer := handler.(UdmProducer)

	var input models.PpData
	var prob *models.ProblemDetails

	gpsi := ctx.Param("gpsi")

	//decode the request (and body)
	if prob = ctx.DecodeRequest(&input); prob == nil {
		//call the application handler
		prob = udmproducer.HandleUpdate(gpsi, &input)
	}

	if prob != nil {
		resp.SetProblem(prob)
	} else {
		resp.SetBody(204, nil)
	}
	return
}



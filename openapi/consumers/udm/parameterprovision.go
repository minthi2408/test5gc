package udm

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)



func Update(client openapi.ConsumerClient, gpsi string, ppdata models.PpData) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PATCH"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_PARAMETERPROVISION, gpsi, openapi.UDM_PP_DATA)
	req.Body = &ppdata
	req.HeaderParams["Content-Type"] = "application/merge-patch+json"

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch resp.StatusCode {
	case 204:
		//good, do no thing
	case 400, 403, 404, 500, 503:
		var prob models.ProblemDetails
		resp.Body = &prob
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
		}

	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

package udm

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

func UecmanGet(client openapi.ConsumerClient, ueId string, supportedFeatures string) (result models.Amf3GppAccessRegistration, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_UECMAN, ueId, openapi.UDM_UECMAN_AMF_3GPP_ACCESS)

	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 200:
		resp.Body = &result
		err = client.DecodeResponse(resp)

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

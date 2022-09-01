package ausf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

func SupiUeUpuPost(client openapi.ConsumerClient, supi string, upuInfo models.UpuInfo) (
	result models.UpuSecurityInfo, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s", supi, openapi.AUSF_UPUPROTECTION, openapi.AUSF_UE_UPUPROTECTION)
	req.Body = &upuInfo

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
	case 503:
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

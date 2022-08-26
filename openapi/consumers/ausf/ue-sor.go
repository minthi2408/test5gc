package ausf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

func (consumer *ausfConsumerImpl) SupiUeSorPost(supi string, sorInfo models.SorInfo) (result models.SorSecurityInfo, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s", supi, openapi.AUSF_SORPROTECTION, openapi.AUSF_UE_SORPROTECTION)
	req.Body = &sorInfo

	//send the request
	var resp *openapi.Response
	if resp, err = consumer.client.Send(req); err != nil {
		return
	}

	//handle the response

	var prob models.ProblemDetails

	switch resp.StatusCode {
	case 200:
        resp.Body = &result
		err = consumer.client.DecodeResponse(resp)
	case 503:
        resp.Body = &prob
		if err = consumer.client.DecodeResponse(resp); err == nil {
			err = openapi.NewError(err.Error(), &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

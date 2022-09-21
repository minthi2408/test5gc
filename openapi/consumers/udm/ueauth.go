package udm

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

func ConfirmAuth(client openapi.ConsumerClient, supi string, authEvent models.AuthEvent) (result models.AuthEvent, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_UEAUTHENTICATION, supi, openapi.UDM_UEAUTH_EVENTS)
	req.Body = &authEvent

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 201:
		resp.Body = &result
		err = client.DecodeResponse(resp)

	case 400, 404, 500, 503:
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

func GenerateAuthData(client openapi.ConsumerClient, supiOrSuci string, authReq models.AuthenticationInfoRequest) (result models.AuthenticationInfoResult, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_UEAUTHENTICATION, supiOrSuci, openapi.UDM_UEAUTH_GEN_AUTH_DATA)
	req.Body = &authReq

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

	case 400, 403, 404, 500, 501, 503:
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

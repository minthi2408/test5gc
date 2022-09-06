package ausf

import (
	"fmt"

	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func EapAuthMethod(client openapi.ConsumerClient, authCtxId string, eapin *models.EapSession) (
	eapout models.EapSession, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s/eap-session", openapi.AUSF_AUTHENTICATIONS, openapi.AUSF_UE_AUTHENTICATIONS, authCtxId)
	req.Body = eapin

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch resp.StatusCode {
	case 200:
		resp.Body = &eapout
		err = client.DecodeResponse(resp)
	case 400, 500:
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

func UeAuthPost(client openapi.ConsumerClient, info models.AuthenticationInfo) (
	authCtx models.UeAuthenticationCtx, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s", openapi.AUSF_AUTHENTICATIONS, openapi.AUSF_UE_AUTHENTICATIONS)
	req.Body = &info

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch resp.StatusCode {
	case 201:
		resp.Body = &authCtx
		err = client.DecodeResponse(resp)
	case 400, 403, 500:
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

func UeAuthAuthCtxId5gAkaConfirmationPut(client openapi.ConsumerClient, authCtxId string, resStar string) (
	confirm models.ConfirmationDataResponse, err error) {

	// create a request
	req := openapi.DefaultRequest()
	req.Method = "PUT"
	req.Path = fmt.Sprintf("%s/%s/%s/5g-aka-confirmation", openapi.AUSF_AUTHENTICATIONS, openapi.AUSF_UE_AUTHENTICATIONS, authCtxId)
	req.Body = &models.ConfirmationData{
		ResStar: resStar,
	}

	// send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	// handle the request

	switch resp.StatusCode {
	case 200:
		resp.Body = &confirm
		err = client.DecodeResponse(resp)
	case 400, 500:
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

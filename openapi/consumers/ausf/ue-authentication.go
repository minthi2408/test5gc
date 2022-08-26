package ausf

import (
	"fmt"

	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func (consumer *ausfConsumerImpl) EapAuthMethod(authCtxId string, eapin *models.EapSession) (
	eapout models.EapSession, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s/eap-session", openapi.AUSF_AUTHENTICATIONS, openapi.AUSF_UE_AUTHENTICATIONS, authCtxId)
	req.Body = eapin

	//send the request
	var resp *openapi.Response
	if resp, err = consumer.client.Send(req); err != nil {
		return
	}

	//handle the response
	//httpresp = resp.Response.(*http.Response)
	//encoding := httpresp.Header.Get("Content-Type")

	var prob models.ProblemDetails
	switch resp.StatusCode {
	case 200:
		//err = consumer.client.DecodeBody(&eapout, resp.BodyBytes, encoding)
        resp.Body = &eapout
        err = consumer.client.DecodeResponse(resp)
	case 400:
		fallthrough
	case 500:
        resp.Body = &prob
		if err = consumer.client.DecodeResponse(resp); err == nil {
			err = openapi.NewError(err.Error(), &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
        return
	}
	return
}

func (consumer *ausfConsumerImpl) UeAuthPost(info models.AuthenticationInfo) (
	authCtx models.UeAuthenticationCtx, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s", openapi.AUSF_AUTHENTICATIONS, openapi.AUSF_UE_AUTHENTICATIONS)
	req.Body = &info

	//send the request
	var resp *openapi.Response
	if resp, err = consumer.client.Send(req); err != nil {
		return
	}

	//handle the response

	var prob models.ProblemDetails
	switch resp.StatusCode {
	case 201:
        resp.Body = &authCtx
		err = consumer.client.DecodeResponse(resp)
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 500:
        resp.Body = &prob
		if err = consumer.client.DecodeResponse(resp); err == nil {
			err = openapi.NewError(err.Error(), &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

func (consumer *ausfConsumerImpl) UeAuthAuthCtxId5gAkaConfirmationPut(authCtxId string, resStar string) (
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
	if resp, err = consumer.client.Send(req); err != nil {
		return
	}

	// handle the request

	var prob models.ProblemDetails
	switch resp.StatusCode {
	case 200:
        resp.Body = &confirm
		err = consumer.client.DecodeResponse(resp)
	case 400:
		fallthrough
	case 500:
        resp.Body = &prob
		if err = consumer.client.DecodeResponse(resp); err == nil {
			err = openapi.NewError(resp.Status, &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

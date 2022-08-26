package ausf

import (
	"fmt"
	"net/http"

	"etri5gc/openapi"
	"etri5gc/openapi/models"
)

func (consumer *ausfConsumerImpl) EapAuthMethod(authCtxId string, eapin *models.EapSession) (
	eapout models.EapSession, httpresp *http.Response, err error) {

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
	httpresp = resp.Response.(*http.Response)
	encoding := httpresp.Header.Get("Content-Type")

	var prob models.ProblemDetails
	switch httpresp.StatusCode {
	case 200:
		err = consumer.client.DecodeBody(&eapout, resp.BodyBytes, encoding)
	case 400:
		fallthrough
	case 500:
		if err = consumer.client.DecodeBody(&prob, resp.BodyBytes, encoding); err == nil {
			err = openapi.NewError(err.Error(), &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", httpresp.StatusCode)
	}
	return
}

func (consumer *ausfConsumerImpl) UeAuthPost(info models.AuthenticationInfo) (
	authCtx models.UeAuthenticationCtx, httpresp *http.Response, err error) {

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
	httpresp = resp.Response.(*http.Response)
	encoding := httpresp.Header.Get("Content-Type")

	var prob models.ProblemDetails
	switch httpresp.StatusCode {
	case 201:
		err = consumer.client.DecodeBody(&authCtx, resp.BodyBytes, encoding)
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 500:
		if err = consumer.client.DecodeBody(&prob, resp.BodyBytes, encoding); err == nil {
			err = openapi.NewError(err.Error(), &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", httpresp.StatusCode)
	}
	return
}

func (consumer *ausfConsumerImpl) UeAuthAuthCtxId5gAkaConfirmationPut(authCtxId string, resStar string) (
	confirm models.ConfirmationDataResponse, httpresp *http.Response, err error) {

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
	httpresp = resp.Response.(*http.Response)
	encoding := httpresp.Header.Get("Content-Type")

	var prob models.ProblemDetails
	switch httpresp.StatusCode {
	case 200:
		err = consumer.client.DecodeBody(&confirm, resp.BodyBytes, encoding)
	case 400:
		fallthrough
	case 500:
		if err = consumer.client.DecodeBody(&prob, resp.BodyBytes, encoding); err == nil {
			err = openapi.NewError(httpresp.Status, &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", httpresp.StatusCode)
	}
	return
}

package ausf

import (
	"fmt"
	"net/http"

	"etri5gc/openapi"
	"etri5gc/openapi/models"

//	"github.com/antihax/optional"
)
const (
    AUSF_UE_AUTHENTICATIONS = "/ue-authentications"
)

func (consumer *ausfImpl) EapAuthMethod(authCtxId string, eapin *models.EapSession,) (eapout models.EapSession, httpresp *http.Response, err error) {
	req := openapi.DefaultRequest()
    req.Method = "POST"
	// create path and map variables
	req.Path = fmt.Sprintf("%s/%s.eap-session", AUSF_UE_AUTHENTICATIONS, authCtxId)
	// body params
    req.Body = eapin

    var resp *openapi.Response
	if resp, err = consumer.client.Send(req); err != nil {
        return
    }
    httpresp = resp.Raw.(*http.Response)
    encoding := httpresp.Header.Get("Content-Type")

	var prob models.ProblemDetails
	switch resp.StatusCode {
	case 200:
        err = consumer.client.DecodeBody(&eapout, resp.Body, encoding )
	case 400:
        fallthrough
	case 500:
        if err = consumer.client.DecodeBody(&prob, resp.Body, encoding); err == nil {
            err = openapi.NewError(err.Error(), &prob)
        }
	default:
        err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
    return
}

func (consumer *ausfImpl) UeAuthPost(info models.AuthenticationInfo) (authCtx models.UeAuthenticationCtx, httpresp *http.Response, err error) {
	req := openapi.DefaultRequest()
    req.Method = "POST"
    req.Path = AUSF_UE_AUTHENTICATIONS
    req.Body = &info

    var resp *openapi.Response
    if resp, err = consumer.client.Send(req); err != nil {
        return
    }
    httpresp = resp.Raw.(*http.Response)
    encoding := httpresp.Header.Get("Content-Type")

	var prob models.ProblemDetails
	switch resp.StatusCode {
	case 201:
        err = consumer.client.DecodeBody(&authCtx, resp.Body, encoding )
	case 400:
        fallthrough
	case 403:
        fallthrough
	case 500:
        if err = consumer.client.DecodeBody(&prob, resp.Body, encoding); err == nil {
            err = openapi.NewError(err.Error(), &prob)
        }
	default:
        err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
    return
}


func (consumer *ausfImpl) UeAuthAuthCtxId5gAkaConfirmationPut(authCtxId string, resStar string) (confirm models.ConfirmationDataResponse, httpresp *http.Response, err error) {

	req := openapi.DefaultRequest()
    req.Method = "PUT"
    req.Path = fmt.Sprintf("%s/%s/5g-aka-confirmation",AUSF_UE_AUTHENTICATIONS, authCtxId)
    req.Body = &models.ConfirmationData{
        ResStar: resStar,
    }

    var resp *openapi.Response
    if resp, err = consumer.client.Send(req); err != nil {
        return
    }

    httpresp = resp.Raw.(*http.Response)
    encoding := httpresp.Header.Get("Content-Type")

	var prob models.ProblemDetails
	switch httpresp.StatusCode {
	case 200:
        err = consumer.client.DecodeBody(&confirm, resp.Body, encoding)
	case 400:
        fallthrough
	case 500:
		if err = consumer.client.DecodeBody(&prob, resp.Body, encoding); err == nil {
            err = openapi.NewError(httpresp.Status, &prob)
        }
	default:
        err = fmt.Errorf("invalid status code: %d", httpresp.StatusCode)
	}
    return
}

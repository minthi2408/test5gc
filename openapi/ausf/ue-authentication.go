package ausf

import (
	"context"
	"etri5gc/openapi"
	"fmt"
	"net/http"
	"strings"

	"etri5gc/openapi/models"

	"github.com/antihax/optional"
)

type Sender interface {
	Send(*openapi.Request) (*openapi.Response, error)
}
type UeAuth struct {
	client Sender
}

type EapAuthMethodParamOpts struct {
	EapSession optional.Interface
}

func (consumer *UeAuth) EapAuthMethod(ctx context.Context, authCtxId string, eap *EapAuthMethodParamOpts) (models.EapSession, *http.Response, error) {
	req := openapi.DefaultRequest()
	// create path and map variables
	req.Path = "/ue-authentications/{authCtxId}/eap-session"
	req.Path = strings.Replace(req.Path, "{"+"authCtxId"+"}", fmt.Sprintf("%v", authCtxId), -1)

	// body params
	if eap != nil && eap.EapSession.IsSet() {
		eapsession := localVarOptionals.EapSession.Value().(models.EapSession)
		req.Body = &eapsession
	}

	resp, err := consumer.client.Send(req)

	//localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	//localVarHTTPResponse.Body.Close()

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch resp.StatusCode {
	case 200:
		//err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		//if err != nil {
	//		apiError.ErrorStatus = err.Error()
	//	}
	//	return localVarReturnValue, localVarHTTPResponse, nil
	case 400:
	//	var v models.ProblemDetails
	//	err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	//	if err != nil {
	//		apiError.ErrorStatus = err.Error()
	//		return localVarReturnValue, localVarHTTPResponse, apiError
	//	}
	//	apiError.ErrorModel = v
	//	return localVarReturnValue, localVarHTTPResponse, apiError
	case 500:
	//	var v models.ProblemDetails
	//	err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	//	if err != nil {
	//		apiError.ErrorStatus = err.Error()
	//		return localVarReturnValue, localVarHTTPResponse, apiError
	//	}
	//	apiError.ErrorModel = v
	//	return localVarReturnValue, localVarHTTPResponse, apiError
	default:
		//	return localVarReturnValue, localVarHTTPResponse, openapi.ReportError("%d is not a valid status code in EapAuthMethod", localVarHTTPResponse.StatusCode)
	}
}

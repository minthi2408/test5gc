package amf

import (
	"context"
	"etri5gc/openapi/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
IndividualSubscriptionDocumentApiService Namf_Communication AMF Status Change Subscribe Modify service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param subscriptionId AMF Status Change Subscription Identifier
 * @param subscriptionData
@return SubscriptionData
*/

func AMFStatusChangeSubscribeModfy(client openapi.ConsumerClient, subscriptionId string, body models.SubscriptionData) (result models.SubscriptionData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PUT"
	req.Path = fmt.Sprintf("%s/%s/subscriptions/%s", openapi.AMF_COMMUNICATION, openapi.AMF_COMM_, subscriptionId)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}
	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", fmt.Sprintf("%v", subscriptionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHttpContentTypes := []string{"application/json"}
	localVarHeaderParams["Content-Type"] = localVarHttpContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := openapi.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	// body params
	localVarPostBody = &subscriptionData

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHttpResponse.Status,
	}
	switch localVarHttpResponse.StatusCode {
	case 202:
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHttpResponse, nil
	case 400:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 403:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 411:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 413:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 415:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 429:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	default:
		return localVarReturnValue, localVarHttpResponse, openapi.ReportError("%d is not a valid status code in AMFStatusChangeSubscribeModfy", localVarHttpResponse.StatusCode)
	}
}

/*
IndividualSubscriptionDocumentApiService Namf_Communication AMF Status Change UnSubscribe service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param subscriptionId AMF Status Change Subscription Identifier
*/

func (a *IndividualSubscriptionDocumentApiService) AMFStatusChangeUnSubscribe(ctx context.Context, subscriptionId string) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", fmt.Sprintf("%v", subscriptionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHttpContentTypes := []string{"application / json"}
	localVarHeaderParams["Content-Type"] = localVarHttpContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := openapi.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHttpResponse.Status,
	}
	switch localVarHttpResponse.StatusCode {
	case 204:
		return localVarHttpResponse, nil
	case 400:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 404:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 429:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	default:
		return localVarHttpResponse, openapi.ReportError("%d is not a valid status code in AMFStatusChangeUnSubscribe", localVarHttpResponse.StatusCode)
	}
}

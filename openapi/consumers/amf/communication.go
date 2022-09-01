package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
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
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_SUBSCRIPTIONS, subscriptionId)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	var prob models.ProblemDetails
	switch resp.StatusCode {
	case 202:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
	case 503:
		resp.Body = &prob
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

/*
IndividualSubscriptionDocumentApiService Namf_Communication AMF Status Change UnSubscribe service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param subscriptionId AMF Status Change Subscription Identifier
*/

func AMFStatusChangeUnSubscribe(client openapi.ConsumerClient, subscriptionId string) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "DELETE"
	req.HeaderParams["Accept"] = "application/problem+json"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_SUBSCRIPTIONS, subscriptionId)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 204:
		return
	case 400:
		fallthrough
	case 404:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
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

/*
IndividualUeContextDocumentApiService Namf_Communication CreateUEContext service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
@return UeContextCreatedData
*/

func CreateUEContext(client openapi.ConsumerClient, ueContextId string, body models.CreateUeContextRequest) (result models.CreateUeContextResponse, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PUT"
	req.HeaderParams["Accept"] = "application/problem+json"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_UE_CONTEXTS, ueContextId)
	req.Body = &body

	if body.BinaryDataN2Information != nil ||
		body.BinaryDataN2InformationExt1 != nil ||
		body.BinaryDataN2InformationExt2 != nil ||
		body.BinaryDataN2InformationExt3 != nil ||
		body.BinaryDataN2InformationExt4 != nil ||
		body.BinaryDataN2InformationExt5 != nil ||
		body.BinaryDataN2InformationExt6 != nil ||
		body.BinaryDataN2InformationExt7 != nil ||
		body.BinaryDataN2InformationExt8 != nil ||
		body.BinaryDataN2InformationExt9 != nil ||
		body.BinaryDataN2InformationExt10 != nil ||
		body.BinaryDataN2InformationExt11 != nil ||
		body.BinaryDataN2InformationExt12 != nil ||
		body.BinaryDataN2InformationExt13 != nil ||
		body.BinaryDataN2InformationExt14 != nil ||
		body.BinaryDataN2InformationExt15 != nil ||
		body.BinaryDataN2InformationExt16 != nil {
		req.HeaderParams["Content-Type"] = "multipart/related"
		req.Body = &body

	} else {
		req.HeaderParams["Content-Type"] = "application/json"
		req.Body = body.JsonData
	}

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 201:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 500:
		var cErr models.UeContextCreateError
		resp.Body = &cErr
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &cErr)
		}
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
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

/*
IndividualUeContextDocumentApiService Namf_Communication EBI Assignment service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
 * @param assignEbiData
@return AssignedEbiData
*/

func EBIAssignment(client openapi.ConsumerClient, ueContextId string, body models.AssignEbiData) (result models.AssignedEbiData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.HeaderParams["Accept"] = "application/problem+json"
	req.Path = fmt.Sprintf("%s/%s/%s/assign-ebi", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_UE_CONTEXTS, ueContextId)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 200:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 500:
		var apiErr models.AssignEbiError
		resp.Body = &apiErr
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &apiErr)
		}
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
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

/*
IndividualUeContextDocumentApiService Namf_Communication RegistrationStatusUpdate service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
 * @param ueRegStatusUpdateReqData
@return UeRegStatusUpdateRspData
*/

func RegistrationStatusUpdate(client openapi.ConsumerClient, ueContextId string, body models.UeRegStatusUpdateReqData) (result models.UeRegStatusUpdateRspData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.HeaderParams["Accept"] = "application/problem+json"
	req.Path = fmt.Sprintf("%s/%s/%s/transfer-update", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_UE_CONTEXTS, ueContextId)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 200:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 404:
		fallthrough
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
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

/*
IndividualUeContextDocumentApiService Namf_Communication ReleaseUEContext service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
 * @param ueContextRelease
*/

func ReleaseUEContext(client openapi.ConsumerClient, ueContextId string, body models.UeContextRelease) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.HeaderParams["Accept"] = "application/problem+json"
	req.Path = fmt.Sprintf("%s/%s/%s/release", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_UE_CONTEXTS, ueContextId)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 204:
		return
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 404:
		fallthrough
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
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

/*
IndividualUeContextDocumentApiService Namf_Communication UEContextTransfer service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
 * @param ueContextTransferRequest
@return UeContextTransferRspData
*/

func UEContextTransfer(client openapi.ConsumerClient, ueContextId string, body models.UeContextTransferRequest) (result models.UeContextTransferResponse, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.HeaderParams["Accept"] = "application/problem+json"
	req.Path = fmt.Sprintf("%s/%s/%s/transfer", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_UE_CONTEXTS, ueContextId)
	req.Body = &body

	// to determine is multipart request
	if body.BinaryDataN1Message != nil {
		req.HeaderParams["Content-Type"] = "multipart/related"
	} else {
		req.HeaderParams["Content-Type"] = "application/json"
		req.Body = body.JsonData
	}

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 200:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
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

//NOTE: call back consumer

func N2InfoNotify(client openapi.ConsumerClient, n2InfoNotifyUrl string, body models.N2InformationNotification) (result models.N2InfoNotifyResponse, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.HeaderParams["Accept"] = "application/problem+json"
	//TODO: check this URL, should it be part of the consumer client
	//localVarPath := n2InfoNotifyUrl
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 200:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 204:
	case 400:
		fallthrough
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
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

/*
N1N2IndividualSubscriptionDocumentApiService Namf_Communication N1N2 Message UnSubscribe (UE Specific) service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
 * @param subscriptionId Subscription Identifier
*/

func N1N2MessageUnSubscribe(client openapi.ConsumerClient, ueContextId string, subscriptionId string) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "DELETE"
	req.HeaderParams["Accept"] = "application/problem+json"
	req.Path = fmt.Sprintf("%s/%s/%s/n1-n2-messages/subscriptions/%s", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_UE_CONTEXTS, ueContextId, subscriptionId)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 204:
	case 400:
	case 411:
	case 413:
	case 415:
	case 429:
	case 500:
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

/*
N1N2MessageCollectionDocumentApiService Namf_Communication N1N2 Message Transfer (UE Specific) service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
 * @param n1MessageContainer
 * @param n2InfoContainer
 * @param skipInd
 * @param lastMsgIndication
 * @param pduSessionId
 * @param lcsCorrelationId
 * @param ppi
 * @param arp
 * @param var5qi
 * @param n1n2FailureTxfNotifURI
 * @param smfReallocationInd
 * @param areaOfValidity
 * @param supportedFeatures
@return N1N2MessageTransferRspData
*/

func N1N2MessageTransfer(client openapi.ConsumerClient, ueContextId string, body models.N1N2MessageTransferRequest) (result models.N1N2MessageTransferRspData, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.HeaderParams["Accept"] = "application/problem+json"
	req.Path = fmt.Sprintf("%s/%s/%s/n1-n2-messages", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_SUBSCRIPTIONS, ueContextId)
	req.Body = &body

	// to determine is multipart request
	if body.BinaryDataN1Message != nil || body.BinaryDataN2Information != nil {
		req.HeaderParams["Content-Type"] = "multipart/related"
	} else {
		req.HeaderParams["Content-Type"] = "application/json"
		req.Body = body.JsonData
	}

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 202:
		fallthrough
	case 200:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 307:
		fallthrough
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 404:
		fallthrough
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
	case 503:
		var prob models.ProblemDetails
		resp.Body = &prob
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
		}
	case 409:
		fallthrough
	case 504:
		var apiErr models.N1N2MessageTransferError
		resp.Body = &apiErr
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &apiErr)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

func N1MessageNotify(client openapi.ConsumerClient, n1MessageNotificationUrl string, body models.N1MessageNotify) (err error) {
	/*
		var (
			localVarHttpMethod   = strings.ToUpper("Post")
			localVarPostBody     interface{}
			localVarFormFileName string
			localVarFileName     string
			localVarFileBytes    []byte
		)

		// create path and map variables
		localVarPath := n1MessageNotificationUrl
		localVarHeaderParams := make(map[string]string)
		localVarQueryParams := url.Values{}
		localVarFormParams := url.Values{}

		// to determine is multipart request
		if request.BinaryDataN1Message != nil {
			localVarHeaderParams["Content-Type"] = "multipart/related"
			localVarPostBody = &request
		} else {
			localVarHeaderParams["Content-Type"] = "application/json"
			localVarPostBody = request.JsonData
		}

		// to determine the Accept header
		localVarHttpHeaderAccepts := []string{"application/json", "multipart/related", "application/problem+json"}

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
			return localVarHttpResponse, err
		case 400:
			var v models.ProblemDetails
			err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				apiError.ErrorStatus = err.Error()
				return localVarHttpResponse, apiError
			}
			apiError.ErrorModel = v
			return localVarHttpResponse, apiError
		case 411:
			var v models.ProblemDetails
			err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				apiError.ErrorStatus = err.Error()
				return localVarHttpResponse, apiError
			}
			apiError.ErrorModel = v
			return localVarHttpResponse, apiError
		case 413:
			var v models.ProblemDetails
			err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				apiError.ErrorStatus = err.Error()
				return localVarHttpResponse, apiError
			}
			apiError.ErrorModel = v
			return localVarHttpResponse, apiError
		case 415:
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
			return localVarHttpResponse, openapi.ReportError("%d is not a valid status code in N1MessageNotify", localVarHttpResponse.StatusCode)
		}
	*/
	return
}

/*
N1N2MessageTransferStatusNotificationCallbackDocumentApiService Namf_Communication N1N2Transfer Failure Notification service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
 * @param subscriptionId Subscription Identifier
*/
func N1N2TransferFailureNotification(client openapi.ConsumerClient, n1N2MessageTransferNotificationUrl string, body models.N1N2MsgTxfrFailureNotification) (err error) {
	/*
		var (
			localVarHttpMethod   = strings.ToUpper("Post")
			localVarPostBody     interface{}
			localVarFormFileName string
			localVarFileName     string
			localVarFileBytes    []byte
		)

		// create path and map variables
		localVarPath := n1N2MessageTransferNotificationUrl

		localVarHeaderParams := make(map[string]string)
		localVarQueryParams := url.Values{}
		localVarFormParams := url.Values{}

		localVarHttpContentTypes := []string{"application/json"}
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
			return localVarHttpResponse, err
		case 400:
			var v models.ProblemDetails
			err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				apiError.ErrorStatus = err.Error()
				return localVarHttpResponse, apiError
			}
			apiError.ErrorModel = v
			return localVarHttpResponse, apiError
		case 411:
			var v models.ProblemDetails
			err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				apiError.ErrorStatus = err.Error()
				return localVarHttpResponse, apiError
			}
			apiError.ErrorModel = v
			return localVarHttpResponse, apiError
		case 413:
			var v models.ProblemDetails
			err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				apiError.ErrorStatus = err.Error()
				return localVarHttpResponse, apiError
			}
			apiError.ErrorModel = v
			return localVarHttpResponse, apiError
		case 415:
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
			return localVarHttpResponse, openapi.ReportError("%d is not a valid status code in N1N2TransferFailureNotification", localVarHttpResponse.StatusCode)
		}
	*/
	return

}

/*
N1N2SubscriptionsCollectionForIndividualUEContextsDocumentApiService Namf_Communication N1N2 Message Subscribe (UE Specific) service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueContextId UE Context Identifier
 * @param ueN1N2InfoSubscriptionCreateData
@return UeN1N2InfoSubscriptionCreatedData
*/

func N1N2MessageSubscribe(client openapi.ConsumerClient, ueContextId string, body models.UeN1N2InfoSubscriptionCreateData) (result models.UeN1N2InfoSubscriptionCreatedData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s/n1-n2-messages/subscriptions", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_UE_CONTEXTS, ueContextId)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 201:
	case 400:
		fallthrough
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
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

/*
NonUEN2MessageNotificationIndividualSubscriptionDocumentApiService Namf_Communication Non UE N2 Info UnSubscribe service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param n2NotifySubscriptionId N2 info Subscription Identifier
*/

func NonUeN2InfoUnSubscribe(client openapi.ConsumerClient, subId string) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/subscriptions/%s", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_NONUE_N2_MESSAGES, subId)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//TODO: look strange, no success status code?
	switch resp.StatusCode {
	case 400:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
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

/*
NonUEN2MessagesCollectionDocumentApiService Namf_Communication Non UE N2 Message Transfer service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return N2InformationTransferRspData
*/

func NonUeN2MessageTransfer(client openapi.ConsumerClient, body models.NonUeN2MessageTransferRequest) (result models.N2InformationTransferRspData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/transfer", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_NONUE_N2_MESSAGES)

	// set content-type
	if body.BinaryDataN2Information != nil {
		req.HeaderParams["Content-Type"] = "multipart/related"
		req.Body = &body
	} else {
		req.HeaderParams["Content-Type"] = "application/json"
		req.Body = body.JsonData
	}

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 200:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 403:
		fallthrough
	case 404:
		fallthrough
	case 500:
		fallthrough
	case 503:
		var apiErr models.N2InformationTransferError
		resp.Body = &apiErr
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &apiErr)
		}
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
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

/*
NonUEN2MessagesSubscriptionsCollectionDocumentApiService Namf_Communication Non UE N2 Info Subscribe service Operation
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param nonUeN2InfoSubscriptionCreateData
@return NonUeN2InfoSubscriptionCreatedData
*/

func NonUeN2InfoSubscribe(client openapi.ConsumerClient, body models.NonUeN2InfoSubscriptionCreateData) (result models.NonUeN2InfoSubscriptionCreatedData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/subscriptions", openapi.AMF_COMMUNICATION, openapi.AMF_COMMUNICATION_NONUE_N2_MESSAGES)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	switch resp.StatusCode {
	case 201:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 400:
		fallthrough
	case 403:
		fallthrough
	case 411:
		fallthrough
	case 413:
		fallthrough
	case 415:
		fallthrough
	case 429:
		fallthrough
	case 500:
		fallthrough
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

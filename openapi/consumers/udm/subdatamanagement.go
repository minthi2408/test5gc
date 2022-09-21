package udm

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"etri5gc/openapi/utils"
	"fmt"
)

type GetAmDataParams struct {
	SupportedFeatures string
	PlmnId            string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetAmData(client openapi.ConsumerClient, supi string, params *GetAmDataParams) (result models.AccessAndMobilitySubscriptionData, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_AMDATA)

	//set parameters for the request
	if len(params.IfNoneMatch) > 0 {
		req.HeaderParams["If-Non-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

	if len(params.PlmnId) > 0 {
		req.QueryParams.Add("plmn-id", params.PlmnId)
	}
	if len(params.SupportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", params.SupportedFeatures)
	}

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

type GetIdTranslationResultParams struct {
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetIdTranslationResult(client openapi.ConsumerClient, gpsi string, params *GetIdTranslationResultParams) (result models.IdTranslationResult, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, gpsi, openapi.UDM_SDM_ID_TRANS_RESULT)

	//set parameters for the request
	if len(params.IfNoneMatch) > 0 {
		req.HeaderParams["If-Non-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

	if len(params.SupportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", params.SupportedFeatures)
	}

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

func Info(client openapi.ConsumerClient, supi string, info *models.AcknowledgeInfo) (err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PUT"
	req.Path = fmt.Sprintf("%s/%s/%s/sor-ack", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_AMDATA)
	req.Body = info

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch resp.StatusCode {
	case 204:
		//good, do nothing
	case 400, 500, 503:
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

func PutUpuAck(client openapi.ConsumerClient, supi string, info *models.AcknowledgeInfo) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PUT"
	req.Path = fmt.Sprintf("%s/%s/%s/upu-ack", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_AMDATA)
	req.Body = info

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 204:
		//good, do nothing
	case 400, 500, 503:
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

type GetParams struct {
	PlmnId            string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func SdmGet(client openapi.ConsumerClient, supi string, datnames []models.DataSetName, params *GetParams) (result models.SubscriptionDataSets, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi)

	if len(datnames) < 2 {
		err = fmt.Errorf("datnames must have at least 2 elements")
		return
	}

	req.QueryParams.Add("dataset-names", utils.Param2Str(datnames, "csv"))

	if len(params.PlmnId) > 0 {
		req.QueryParams.Add("plmn-id", params.PlmnId)
	}

	if len(params.SupportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}
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

type GetSharedDataParams struct {
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetSharedData(client openapi.ConsumerClient, datIds []string, params *GetSharedDataParams) (result []models.SharedData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s", openapi.UDM_SUBDATAMANAGEMENT, openapi.UDM_SDM_SHARED_DATA)

	req.QueryParams.Add("shared-data-ids", utils.Param2Str(datIds, "csv"))

	if len(params.SupportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

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
	case 400, 403, 500, 503:
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

type GetSmDataParams struct {
	SupportedFeatures string
	SingleNssai       string
	Dnn               string
	PlmnId            string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetSmData(client openapi.ConsumerClient, supi string, params *GetSmDataParams) (result []models.SessionManagementSubscriptionData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_SM_DATA)

	if len(params.PlmnId) > 0 {
		req.QueryParams.Add("plmn-id", params.PlmnId)
	}

	if len(params.Dnn) > 0 {
		req.QueryParams.Add("dnn", params.Dnn)
	}

	if len(params.SingleNssai) > 0 {
		req.QueryParams.Add("single-nssai", params.SingleNssai)
	}

	if len(params.SupportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

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

type GetNssaiParams struct {
	SupportedFeatures string
	PlmnId            string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetNssai(client openapi.ConsumerClient, supi string, params *GetNssaiParams) (result models.Nssai, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_NSSAI)

	if len(params.PlmnId) > 0 {
		req.QueryParams.Add("plmn-id", params.PlmnId)
	}

	if len(params.SupportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

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

type GetSmsMngDataParams struct {
	SupportedFeatures string
	PlmnId            string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetSmsMngData(client openapi.ConsumerClient, supi string, params *GetSmsMngDataParams) (result models.SmsManagementSubscriptionData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_SMS_MNG_DATA)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	if len(params.PlmnId) > 0 {
		req.QueryParams.Add("plmn-id", params.PlmnId)
	}

	if len(params.SupportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

	//handle the response
	switch resp.StatusCode {
	case 200:
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

type GetSmsDataParams struct {
	SupportedFeatures string
	PlmnId            string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetSmsData(client openapi.ConsumerClient, supi string, params *GetSmsDataParams) (result models.SmsSubscriptionData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_SMS_DATA)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	if len(params.PlmnId) > 0 {
		req.QueryParams.Add("plmn-id", params.PlmnId)
	}

	if len(params.SupportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

	//handle the response
	switch resp.StatusCode {
	case 200:
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

func OnDataChangeNotification(client openapi.ConsumerClient, target_url string, request models.ModificationNotification) (err error) {
	/*
		TODO: build the client with the tartget_url
		var (
			localVarHttpMethod   = strings.ToUpper("Post")
			localVarPostBody     interface{}
			localVarFormFileName string
			localVarFileName     string
			localVarFileBytes    []byte
		)

		// create path and map variables
		localVarPath := onDataChangeNotificationUrl
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

		// body params
		localVarPostBody = &request

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
		case 404:
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
			return localVarHttpResponse, openapi.ReportError("%d is not a valid status code in OnDataChangeNotify", localVarHttpResponse.StatusCode)
		}
	*/
	return
}

func Subscribe(client openapi.ConsumerClient, supi string, sdmsub models.SdmSubscription) (result models.SdmSubscription, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_SUBSCRIPTIONS)
	req.Body = &sdmsub

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

func SubscribeToSharedData(client openapi.ConsumerClient, sdmsub models.SdmSubscription) (result models.SdmSubscription, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s", openapi.UDM_SUBDATAMANAGEMENT, openapi.UDM_SDM_SHARED_DATA_SUB)
	req.Body = &sdmsub

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
	case 400, 404:
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

func Unsubscribe(client openapi.ConsumerClient, supi string, subId string) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "DELETE"
	req.Path = fmt.Sprintf("%s/%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_SUBSCRIPTIONS, subId)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 204:
		//good, do nothing
	case 400, 500, 503:
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

func UnsubscribeForSharedData(client openapi.ConsumerClient, subId string) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "DELETE"
	req.Path = fmt.Sprintf("%s/%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, openapi.UDM_SDM_SHARED_DATA_SUB, subId)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 204:
		//good, do nothing
	case 400, 500, 503:
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

func Modify(client openapi.ConsumerClient, supi string, subId string, submod models.SdmSubsModification) (result models.SdmSubscription, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PATCH"
	req.Path = fmt.Sprintf("%s/%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_SUBSCRIPTIONS, subId)
	req.HeaderParams["Content-Type"] = "application/merge-patch+json"
	req.Body = &submod

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

func ModifyForSharedData(client openapi.ConsumerClient, subId string, submod models.SdmSubsModification) (result models.SdmSubscription, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PATCH"
	req.HeaderParams["Content-Type"] = "application/merge-patch+json"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, openapi.UDM_SDM_SHARED_DATA_SUB, subId)
	req.Body = &submod

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

type GetTraceDataParams struct {
	SupportedFeatures string
	PlmnId            string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetTraceData(client openapi.ConsumerClient, supi string, params *GetTraceDataParams) (result models.TraceDataResponse, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_TRACE_DATA)

	if len(params.PlmnId) > 0 {
		req.QueryParams.Add("plmn-id", params.PlmnId)
	}

	if len(params.SupportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", params.SupportedFeatures)
	}

	if len(params.IfNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = params.IfNoneMatch
	}

	if len(params.IfModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

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

func GetUeContextInSmfData(client openapi.ConsumerClient, supi string, supportedFeatures string) (result models.UeContextInSmfData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_UECONTEXT_SMF)

	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}

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

func GetUeContextInSmsfData(client openapi.ConsumerClient, supi string, supportedFeatures string) (result models.UeContextInSmsfData, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_UECONTEXT_SMSF)

	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}

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

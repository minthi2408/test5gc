package udm

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

type GetAmDataParams struct {
	SupportedFeatures	string 
	PlmnId             	string 
	IfNoneMatch         string
	IfModifiedSince  	string 
}

func GetAmData(client openapi.ConsumerClient, supi string, params * GetAmDataParams ) (result models.AccessAndMobilitySubscriptionData, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, supi, openapi.UDM_SDM_AMDATA)

	//set parameters for the request
	if len(params.IfNoneMatch) >0 {
		req.HeaderParams["If-Non-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) >0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

	if len(params.PlmnId) >0 {
		req.QueryParams.Add("plmn-id", params.PlmnId)
	}
	if len(params.SupportedFeatures) >0 {
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
	SupportedFeatures	string 
	IfNoneMatch			string 
	IfModifiedSince		string
}

func GetIdTranslationResult(client openapi.ConsumerClient, gpsi string, params *GetIdTranslationResultParams) (result models.IdTranslationResult, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_SUBDATAMANAGEMENT, gpsi, openapi.UDM_SDM_ID_TRANS_RESULT)

	//set parameters for the request
	if len(params.IfNoneMatch) >0 {
		req.HeaderParams["If-Non-Match"] = params.IfNoneMatch
	}
	if len(params.IfModifiedSince) >0 {
		req.HeaderParams["If-Modified-Since"] = params.IfModifiedSince
	}

	if len(params.SupportedFeatures) >0 {
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

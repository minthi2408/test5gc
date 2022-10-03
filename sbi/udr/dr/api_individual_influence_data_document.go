/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package dr

import (
	"fmt"
	"net/http"
	"net/url"
	"etri5gc/sbi"
	"etri5gc/sbi/models"
	"strings"
)


/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param influenceId The Identifier of an Individual Influence Data to be created or updated. It shall apply the format of Data type string.
@return *models.TrafficInfluData, 
*/
func CreateOrReplaceIndividualInfluenceData(client sbi.ConsumerClient, influenceId string, body models.TrafficInfluData) (result models.TrafficInfluData, err error) {
	
	if len(influenceId) == 0 {
		err = fmt.Errorf("influenceId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/application-data/influenceData/{influenceId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"influenceId"+"}", url.PathEscape(influenceId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 400 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 401 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 411 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 413 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 414 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 415 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 429 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 500 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 503 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to CreateOrReplaceIndividualInfluenceData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for CreateOrReplaceIndividualInfluenceData
func OnCreateOrReplaceIndividualInfluenceData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualInfluenceDataDocumentApiHandler)
	
	influenceId := ctx.Param("influenceId")
	if len(influenceId) == 0 {
		//influenceId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "influenceId is required",
		}))
		return
	}

	var input models.TrafficInfluData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.TrafficInfluData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleCreateOrReplaceIndividualInfluenceData(influenceId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}
	

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param influenceId The Identifier of an Individual Influence Data to be updated. It shall apply the format of Data type string.
@return 
*/
func DeleteIndividualInfluenceData(client sbi.ConsumerClient, influenceId string) (err error) {
	
	if len(influenceId) == 0 {
		err = fmt.Errorf("influenceId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/application-data/influenceData/{influenceId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"influenceId"+"}", url.PathEscape(influenceId), -1)	
	req.HeaderParams["Accept"] = "application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 400 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 401 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 429 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 500 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 503 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to DeleteIndividualInfluenceData", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for DeleteIndividualInfluenceData
func OnDeleteIndividualInfluenceData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualInfluenceDataDocumentApiHandler)
	
	influenceId := ctx.Param("influenceId")
	if len(influenceId) == 0 {
		//influenceId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "influenceId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.DR_HandleDeleteIndividualInfluenceData(influenceId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param influenceId The Identifier of an Individual Influence Data to be updated. It shall apply the format of Data type string.
@return *models.TrafficInfluData, 
*/
func UpdateIndividualInfluenceData(client sbi.ConsumerClient, influenceId string, body models.TrafficInfluDataPatch) (result models.TrafficInfluData, err error) {
	
	if len(influenceId) == 0 {
		err = fmt.Errorf("influenceId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPatch

	req.Path = fmt.Sprintf("%s/application-data/influenceData/{influenceId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"influenceId"+"}", url.PathEscape(influenceId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/merge-patch+json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 400 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 401 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 411 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 413 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 415 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 429 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 500 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 503 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to UpdateIndividualInfluenceData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for UpdateIndividualInfluenceData
func OnUpdateIndividualInfluenceData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualInfluenceDataDocumentApiHandler)
	
	influenceId := ctx.Param("influenceId")
	if len(influenceId) == 0 {
		//influenceId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "influenceId is required",
		}))
		return
	}

	var input models.TrafficInfluDataPatch

	var apierr *sbi.ApiError
	var successCode int32
	var result models.TrafficInfluData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleUpdateIndividualInfluenceData(influenceId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}
	

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type IndividualInfluenceDataDocumentApiHandler interface {
	DR_HandleCreateOrReplaceIndividualInfluenceData(influenceId string, body models.TrafficInfluData) (successCode int32, result models.TrafficInfluData, err *sbi.ApiError)
	DR_HandleDeleteIndividualInfluenceData(influenceId string) (successCode int32, err *sbi.ApiError)
	DR_HandleUpdateIndividualInfluenceData(influenceId string, body models.TrafficInfluDataPatch) (successCode int32, result models.TrafficInfluData, err *sbi.ApiError)
}

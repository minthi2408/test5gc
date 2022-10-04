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
	"etri5gc/sbi"
	"etri5gc/sbi/models"
	"etri5gc/sbi/utils"
)


/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param sharedDataIds List of shared data ids
@param supportedFeatures Supported Features
@return []models.SharedData, 
*/
func GetSharedData(client sbi.ConsumerClient, sharedDataIds []string, supportedFeatures string) (result []models.SharedData, err error) {
	
	sharedDataIdsStr := utils.Param2String(sharedDataIds)
	if len(sharedDataIdsStr) == 0 {
		err = fmt.Errorf("sharedDataIds is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/subscription-data/shared-data", ServicePath)
	req.QueryParams.Add("shared-data-ids", sharedDataIdsStr)

	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}	
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
		if resp.StatusCode == 404 {
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
			err = fmt.Errorf("%d is unknown to GetSharedData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for GetSharedData
func OnGetSharedData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(RetrievalOfSharedDataApiHandler)
	
	sharedDataIdsStr := ctx.Param("shared-data-ids")
	if len(sharedDataIdsStr) == 0 {
		//sharedDataIds is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "sharedDataIds is required",
		}))
		return
	}
	var sharedDataIds []string
	var sharedDataIdsErr error
	if sharedDataIds, sharedDataIdsErr = utils.String2ArrayOfstring(sharedDataIdsStr); sharedDataIdsErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: sharedDataIdsErr.Error(), 
		}))
		return
	}
	
	supportedFeatures := ctx.Param("supported-features")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result []models.SharedData


	successCode, result, apierr = prod.DR_HandleGetSharedData(sharedDataIds, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type RetrievalOfSharedDataApiHandler interface {
	DR_HandleGetSharedData(sharedDataIds []string, supportedFeatures string) (successCode int32, result []models.SharedData, err *sbi.ApiError)
}
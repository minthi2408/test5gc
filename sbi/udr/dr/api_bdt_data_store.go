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
)


/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param bdtRefIds List of the BDT reference identifiers.
@param suppFeat Supported Features
@return []BdtData, 
*/
func ReadBdtData(client sbi.ConsumerClient, bdtRefIds []string, suppFeat string) (result []models.BdtData, err error) {
		
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/policy-data/bdt-data", ServicePath)
	bdtRefIdsStr := utils.Param2String(bdtRefIds)
	if len(bdtRefIdsStr) > 0 {
		req.QueryParams.Add("bdt-ref-ids", bdtRefIdsStr)
	}
	if len(suppFeat) > 0 {
		req.QueryParams.Add("supp-feat", suppFeat)
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
			err = fmt.Errorf("%d is unknown to ReadBdtData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for ReadBdtData
func OnReadBdtData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(BdtDataStoreApiHandler)
	
	bdtRefIdsStr := ctx.Param("bdt-ref-ids")
	var bdtRefIds []string
	var bdtRefIdsErr error
	if bdtRefIds, bdtRefIdsErr = utils.String2ArrayOfstring(bdtRefIdsStr); bdtRefIdsErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: bdtRefIdsErr.Error(), 
		}))
		return
	}
	
	suppFeat := ctx.Param("supp-feat")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result []BdtData


	successCode, result, apierr = prod.DR_HandleReadBdtData(bdtRefIds, suppFeat)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type BdtDataStoreApiHandler interface {
	DR_HandleReadBdtData(bdtRefIds []string, suppFeat string) (successCode int32, result []models.BdtData, err *sbi.ApiError)
}

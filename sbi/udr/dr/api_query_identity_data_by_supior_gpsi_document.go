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
@param ueId UE ID
@param appPortId Application port identifier
@param ifNoneMatch Validator for conditional requests, as described in RFC 7232, 3.2
@param ifModifiedSince Validator for conditional requests, as described in RFC 7232, 3.3
@return *models.IdentityData, 
*/
func GetIdentityData(client sbi.ConsumerClient, ueId string, appPortId *models.AppPortId, ifNoneMatch string, ifModifiedSince string) (result models.IdentityData, err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}			
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/identity-data", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	appPortIdStr := utils.Param2String(appPortId)
	if len(appPortIdStr) > 0 {
		req.QueryParams.Add("app-port-id", appPortIdStr)
	}
	if len(ifNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = ifNoneMatch
	}
	if len(ifModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = ifModifiedSince
	}	
	req.HeaderParams["Accept"] = "application/json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to GetIdentityData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for GetIdentityData
func OnGetIdentityData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(QueryIdentityDataBySUPIOrGPSIDocumentApiHandler)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	appPortIdStr := ctx.Param("app-port-id")
	var appPortId *models.AppPortId
	var appPortIdErr error
	if appPortId, appPortIdErr = utils.String2AppPortId(appPortIdStr); appPortIdErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: appPortIdErr.Error(), 
		}))
		return
	}
	
	ifNoneMatch := ctx.Param("If-None-Match")
	ifModifiedSince := ctx.Param("If-Modified-Since")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.IdentityData


	successCode, result, apierr = prod.DR_HandleGetIdentityData(ueId, appPortId, ifNoneMatch, ifModifiedSince)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type QueryIdentityDataBySUPIOrGPSIDocumentApiHandler interface {
	DR_HandleGetIdentityData(ueId string, appPortId *models.AppPortId, ifNoneMatch string, ifModifiedSince string) (successCode int32, result models.IdentityData, err *sbi.ApiError)
}

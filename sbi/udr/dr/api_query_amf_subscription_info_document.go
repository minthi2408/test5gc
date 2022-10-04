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
@param ueId
@param subsId
@return []models.AmfSubscriptionInfo, 
*/
func GetAmfSubscriptionInfo(client sbi.ConsumerClient, ueId string, subsId string) (result []models.AmfSubscriptionInfo, err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}	
	if len(subsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/amf-subscriptions", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"subsId"+"}", url.PathEscape(subsId), -1)	
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
			err = fmt.Errorf("%d is unknown to GetAmfSubscriptionInfo", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for GetAmfSubscriptionInfo
func OnGetAmfSubscriptionInfo(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(QueryAMFSubscriptionInfoDocumentApiHandler)
	
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
	subsId := ctx.Param("subsId")
	if len(subsId) == 0 {
		//subsId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subsId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32
	var result []models.AmfSubscriptionInfo


	successCode, result, apierr = prod.DR_HandleGetAmfSubscriptionInfo(ueId, subsId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type QueryAMFSubscriptionInfoDocumentApiHandler interface {
	DR_HandleGetAmfSubscriptionInfo(ueId string, subsId string) (successCode int32, result []models.AmfSubscriptionInfo, err *sbi.ApiError)
}

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
@param sponsorId
@return *models.SponsorConnectivityData, 
*/
func ReadSponsorConnectivityData(client sbi.ConsumerClient, sponsorId string) (result models.SponsorConnectivityData, err error) {
	
	if len(sponsorId) == 0 {
		err = fmt.Errorf("sponsorId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/policy-data/sponsor-connectivity-data/{sponsorId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"sponsorId"+"}", url.PathEscape(sponsorId), -1)	
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
			err = fmt.Errorf("%d is unknown to ReadSponsorConnectivityData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for ReadSponsorConnectivityData
func OnReadSponsorConnectivityData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(SponsorConnectivityDataDocumentApiHandler)
	
	sponsorId := ctx.Param("sponsorId")
	if len(sponsorId) == 0 {
		//sponsorId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "sponsorId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SponsorConnectivityData


	successCode, result, apierr = prod.DR_HandleReadSponsorConnectivityData(sponsorId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type SponsorConnectivityDataDocumentApiHandler interface {
	DR_HandleReadSponsorConnectivityData(sponsorId string) (successCode int32, result models.SponsorConnectivityData, err *sbi.ApiError)
}

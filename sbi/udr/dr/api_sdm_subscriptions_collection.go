/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package dr

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId UE ID
@return *models.SdmSubscription,
*/
func CreateSdmSubscriptions(client sbi.ConsumerClient, ueId string, body models.SdmSubscription) (result models.SdmSubscription, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/sdm-subscriptions", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
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
			err = fmt.Errorf("%d is unknown to CreateSdmSubscriptions", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

//sbi producer handler for CreateSdmSubscriptions
func OnCreateSdmSubscriptions(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(SDMSubscriptionsCollectionApiHandler)

	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	var input models.SdmSubscription

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SdmSubscription

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleCreateSdmSubscriptions(ueId, input)
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
@param ueId UE id
@param supportedFeatures Supported Features
@return []models.SdmSubscription,
*/
func Querysdmsubscriptions(client sbi.ConsumerClient, ueId string, supportedFeatures string) (result []models.SdmSubscription, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/sdm-subscriptions", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
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
			err = fmt.Errorf("%d is unknown to Querysdmsubscriptions", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

//sbi producer handler for Querysdmsubscriptions
func OnQuerysdmsubscriptions(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(SDMSubscriptionsCollectionApiHandler)

	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	supportedFeatures := ctx.Param("supported-features")

	var apierr *sbi.ApiError
	var successCode int32
	var result []models.SdmSubscription

	successCode, result, apierr = prod.DR_HandleQuerysdmsubscriptions(ueId, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

type SDMSubscriptionsCollectionApiHandler interface {
	DR_HandleCreateSdmSubscriptions(ueId string, body models.SdmSubscription) (successCode int32, result models.SdmSubscription, err *sbi.ApiError)
	DR_HandleQuerysdmsubscriptions(ueId string, supportedFeatures string) (successCode int32, result []models.SdmSubscription, err *sbi.ApiError)
}

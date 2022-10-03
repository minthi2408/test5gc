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
@param ueId UE id
@param subsId
@param supportedFeatures Features required to be supported by the target NF
@return *models.PatchResult, 
*/
func ModifyEesubscription(client sbi.ConsumerClient, ueId string, subsId string, supportedFeatures string, body []PatchItem) (result models.PatchResult, err error) {
	
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
	req.Method = http.MethodPatch

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"subsId"+"}", url.PathEscape(subsId), -1)
	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json-patch+json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to ModifyEesubscription", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for ModifyEesubscription
func OnModifyEesubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(EventExposureSubscriptionDocumentApiHandler)
	
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
	supportedFeatures := ctx.Param("supported-features")

	var input []PatchItem

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PatchResult

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleModifyEesubscription(ueId, subsId, supportedFeatures, input)
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
@param ueId
@param subsId Unique ID of the subscription to remove
@return map[string]interface{}, 
*/
func QueryeeSubscription(client sbi.ConsumerClient, ueId string, subsId string) (result map[string]interface{}, err error) {
	
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

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}", ServicePath)
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
			err = fmt.Errorf("%d is unknown to QueryeeSubscription", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for QueryeeSubscription
func OnQueryeeSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(EventExposureSubscriptionDocumentApiHandler)
	
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
	var result map[string]interface{}


	successCode, result, apierr = prod.DR_HandleQueryeeSubscription(ueId, subsId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId
@param subsId Unique ID of the subscription to remove
@return 
*/
func RemoveeeSubscriptions(client sbi.ConsumerClient, ueId string, subsId string) (err error) {
	
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
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"subsId"+"}", url.PathEscape(subsId), -1)	
	
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
			err = fmt.Errorf("%d is unknown to RemoveeeSubscriptions", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for RemoveeeSubscriptions
func OnRemoveeeSubscriptions(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(EventExposureSubscriptionDocumentApiHandler)
	
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

	successCode, apierr = prod.DR_HandleRemoveeeSubscriptions(ueId, subsId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId
@param subsId
@return 
*/
func UpdateEesubscriptions(client sbi.ConsumerClient, ueId string, subsId string, body models.EeSubscription) (err error) {
	
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
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"subsId"+"}", url.PathEscape(subsId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to UpdateEesubscriptions", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for UpdateEesubscriptions
func OnUpdateEesubscriptions(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(EventExposureSubscriptionDocumentApiHandler)
	
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

	var input models.EeSubscription

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.DR_HandleUpdateEesubscriptions(ueId, subsId, input)
	} else {
		apierr = sbi.ApiErrFromProb(prob)
	}
	

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}




type EventExposureSubscriptionDocumentApiHandler interface {
	DR_HandleModifyEesubscription(ueId string, subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError)
	DR_HandleQueryeeSubscription(ueId string, subsId string) (successCode int32, result map[string]interface{}, err *sbi.ApiError)
	DR_HandleRemoveeeSubscriptions(ueId string, subsId string) (successCode int32, err *sbi.ApiError)
	DR_HandleUpdateEesubscriptions(ueId string, subsId string, body models.EeSubscription) (successCode int32, err *sbi.ApiError)
}

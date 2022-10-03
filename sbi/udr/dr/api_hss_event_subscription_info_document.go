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
@return 
*/
func CreateHSSSubscriptions(client sbi.ConsumerClient, ueId string, subsId string, body models.HssSubscriptionInfo) (err error) {
	
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

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"subsId"+"}", url.PathEscape(subsId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	
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
			err = fmt.Errorf("%d is unknown to CreateHSSSubscriptions", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for CreateHSSSubscriptions
func OnCreateHSSSubscriptions(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(HSSEventSubscriptionInfoDocumentApiHandler)
	
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

	var input models.HssSubscriptionInfo

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.DR_HandleCreateHSSSubscriptions(ueId, subsId, input)
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




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId
@param subsId
@return *models.SmfSubscriptionInfo, 
*/
func GetHssSubscriptionInfo(client sbi.ConsumerClient, ueId string, subsId string) (result models.SmfSubscriptionInfo, err error) {
	
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

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions", ServicePath)
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
			err = fmt.Errorf("%d is unknown to GetHssSubscriptionInfo", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for GetHssSubscriptionInfo
func OnGetHssSubscriptionInfo(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(HSSEventSubscriptionInfoDocumentApiHandler)
	
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
	var result models.SmfSubscriptionInfo


	successCode, result, apierr = prod.DR_HandleGetHssSubscriptionInfo(ueId, subsId)

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
@param subsId
@param supportedFeatures Features required to be supported by the target NF
@return *models.PatchResult, 
*/
func ModifyHssSubscriptionInfo(client sbi.ConsumerClient, ueId string, subsId string, supportedFeatures string, body []models.PatchItem) (result models.PatchResult, err error) {
	
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

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions", ServicePath)
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
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to ModifyHssSubscriptionInfo", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for ModifyHssSubscriptionInfo
func OnModifyHssSubscriptionInfo(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(HSSEventSubscriptionInfoDocumentApiHandler)
	
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

	var input []models.PatchItem

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PatchResult

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleModifyHssSubscriptionInfo(ueId, subsId, supportedFeatures, input)
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
@param subsId
@return 
*/
func RemoveHssSubscriptionsInfo(client sbi.ConsumerClient, ueId string, subsId string) (err error) {
	
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

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions", ServicePath)
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
			err = fmt.Errorf("%d is unknown to RemoveHssSubscriptionsInfo", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for RemoveHssSubscriptionsInfo
func OnRemoveHssSubscriptionsInfo(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(HSSEventSubscriptionInfoDocumentApiHandler)
	
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

	successCode, apierr = prod.DR_HandleRemoveHssSubscriptionsInfo(ueId, subsId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}




type HSSEventSubscriptionInfoDocumentApiHandler interface {
	DR_HandleCreateHSSSubscriptions(ueId string, subsId string, body models.HssSubscriptionInfo) (successCode int32, err *sbi.ApiError)
	DR_HandleGetHssSubscriptionInfo(ueId string, subsId string) (successCode int32, result models.SmfSubscriptionInfo, err *sbi.ApiError)
	DR_HandleModifyHssSubscriptionInfo(ueId string, subsId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError)
	DR_HandleRemoveHssSubscriptionsInfo(ueId string, subsId string) (successCode int32, err *sbi.ApiError)
}

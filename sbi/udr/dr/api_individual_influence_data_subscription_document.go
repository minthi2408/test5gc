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
@param subscriptionId String identifying a subscription to the Individual Influence Data Subscription
@return 
*/
func DeleteIndividualInfluenceDataSubscription(client sbi.ConsumerClient, subscriptionId string) (err error) {
	
	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/application-data/influenceData/subs-to-notify/{subscriptionId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"subscriptionId"+"}", url.PathEscape(subscriptionId), -1)	
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
			err = fmt.Errorf("%d is unknown to DeleteIndividualInfluenceDataSubscription", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for DeleteIndividualInfluenceDataSubscription
func OnDeleteIndividualInfluenceDataSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualInfluenceDataSubscriptionDocumentApiHandler)
	
	subscriptionId := ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		//subscriptionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subscriptionId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.DR_HandleDeleteIndividualInfluenceDataSubscription(subscriptionId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param subscriptionId String identifying a subscription to the Individual Influence Data Subscription
@return *models.TrafficInfluSub, 
*/
func ReadIndividualInfluenceDataSubscription(client sbi.ConsumerClient, subscriptionId string) (result models.TrafficInfluSub, err error) {
	
	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/application-data/influenceData/subs-to-notify/{subscriptionId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"subscriptionId"+"}", url.PathEscape(subscriptionId), -1)	
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
		if resp.StatusCode == 414 {
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
			err = fmt.Errorf("%d is unknown to ReadIndividualInfluenceDataSubscription", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for ReadIndividualInfluenceDataSubscription
func OnReadIndividualInfluenceDataSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualInfluenceDataSubscriptionDocumentApiHandler)
	
	subscriptionId := ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		//subscriptionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subscriptionId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.TrafficInfluSub


	successCode, result, apierr = prod.DR_HandleReadIndividualInfluenceDataSubscription(subscriptionId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param subscriptionId String identifying a subscription to the Individual Influence Data Subscription
@return *models.TrafficInfluSub, 
*/
func ReplaceIndividualInfluenceDataSubscription(client sbi.ConsumerClient, subscriptionId string, body models.TrafficInfluSub) (result models.TrafficInfluSub, err error) {
	
	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/application-data/influenceData/subs-to-notify/{subscriptionId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"subscriptionId"+"}", url.PathEscape(subscriptionId), -1)
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
			err = fmt.Errorf("%d is unknown to ReplaceIndividualInfluenceDataSubscription", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for ReplaceIndividualInfluenceDataSubscription
func OnReplaceIndividualInfluenceDataSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualInfluenceDataSubscriptionDocumentApiHandler)
	
	subscriptionId := ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		//subscriptionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subscriptionId is required",
		}))
		return
	}

	var input models.TrafficInfluSub

	var apierr *sbi.ApiError
	var successCode int32
	var result models.TrafficInfluSub

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleReplaceIndividualInfluenceDataSubscription(subscriptionId, input)
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




type IndividualInfluenceDataSubscriptionDocumentApiHandler interface {
	DR_HandleDeleteIndividualInfluenceDataSubscription(subscriptionId string) (successCode int32, err *sbi.ApiError)
	DR_HandleReadIndividualInfluenceDataSubscription(subscriptionId string) (successCode int32, result models.TrafficInfluSub, err *sbi.ApiError)
	DR_HandleReplaceIndividualInfluenceDataSubscription(subscriptionId string, body models.TrafficInfluSub) (successCode int32, result models.TrafficInfluSub, err *sbi.ApiError)
}
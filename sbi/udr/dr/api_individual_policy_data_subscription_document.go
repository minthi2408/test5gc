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
@param subsId
@return
*/
func DeleteIndividualPolicyDataSubscription(client sbi.ConsumerClient, subsId string) (err error) {

	if len(subsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/policy-data/subs-to-notify/{subsId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"subsId"+"}", url.PathEscape(subsId), -1)
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
			err = fmt.Errorf("%d is unknown to DeleteIndividualPolicyDataSubscription", resp.StatusCode)
			return
		}
	}

	return
}

//sbi producer handler for DeleteIndividualPolicyDataSubscription
func OnDeleteIndividualPolicyDataSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualPolicyDataSubscriptionDocumentApiHandler)

	subsId := ctx.Param("subsId")
	if len(subsId) == 0 {
		//subsId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subsId is required",
		}))
		return
	}

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.DR_HandleDeleteIndividualPolicyDataSubscription(subsId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param subsId
@return *models.PolicyDataSubscription,
*/
func ReplaceIndividualPolicyDataSubscription(client sbi.ConsumerClient, subsId string, body models.PolicyDataSubscription) (result models.PolicyDataSubscription, err error) {

	if len(subsId) == 0 {
		err = fmt.Errorf("subsId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/policy-data/subs-to-notify/{subsId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"subsId"+"}", url.PathEscape(subsId), -1)
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
			err = fmt.Errorf("%d is unknown to ReplaceIndividualPolicyDataSubscription", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

//sbi producer handler for ReplaceIndividualPolicyDataSubscription
func OnReplaceIndividualPolicyDataSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualPolicyDataSubscriptionDocumentApiHandler)

	subsId := ctx.Param("subsId")
	if len(subsId) == 0 {
		//subsId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "subsId is required",
		}))
		return
	}

	var input models.PolicyDataSubscription

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PolicyDataSubscription

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleReplaceIndividualPolicyDataSubscription(subsId, input)
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

type IndividualPolicyDataSubscriptionDocumentApiHandler interface {
	DR_HandleDeleteIndividualPolicyDataSubscription(subsId string) (successCode int32, err *sbi.ApiError)
	DR_HandleReplaceIndividualPolicyDataSubscription(subsId string, body models.PolicyDataSubscription) (successCode int32, result models.PolicyDataSubscription, err *sbi.ApiError)
}

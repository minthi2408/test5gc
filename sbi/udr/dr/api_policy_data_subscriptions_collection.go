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
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@return *models.PolicyDataSubscription,
*/
func CreateIndividualPolicyDataSubscription(client sbi.ConsumerClient, body models.PolicyDataSubscription) (result models.PolicyDataSubscription, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/policy-data/subs-to-notify", ServicePath)
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
			err = fmt.Errorf("%d is unknown to CreateIndividualPolicyDataSubscription", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

//sbi producer handler for CreateIndividualPolicyDataSubscription
func OnCreateIndividualPolicyDataSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(PolicyDataSubscriptionsCollectionApiHandler)

	var input models.PolicyDataSubscription

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PolicyDataSubscription

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleCreateIndividualPolicyDataSubscription(input)
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

type PolicyDataSubscriptionsCollectionApiHandler interface {
	DR_HandleCreateIndividualPolicyDataSubscription(body models.PolicyDataSubscription) (successCode int32, result models.PolicyDataSubscription, err *sbi.ApiError)
}

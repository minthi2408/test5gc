/*
Npcf_BDTPolicyControl Service API

PCF BDT Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package btdpc

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICE_PATH = "{apiRoot}/npcf-bdtpolicycontrol/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@return *models.BdtPolicy,
*/
func CreateBDTPolicy(client sbi.ConsumerClient, body models.BdtReqData) (result models.BdtPolicy, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/bdtpolicies", SERVICE_PATH)
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
			err = fmt.Errorf("%d is unknown to CreateBDTPolicy", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param bdtPolicyId String identifying the individual BDT policy resource in the PCF.
@return *models.BdtPolicy,
*/
func GetBDTPolicy(client sbi.ConsumerClient, bdtPolicyId string) (result models.BdtPolicy, err error) {

	if len(bdtPolicyId) == 0 {
		err = fmt.Errorf("bdtPolicyId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/bdtpolicies/{bdtPolicyId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"bdtPolicyId"+"}", url.PathEscape(bdtPolicyId), -1)
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 307 {
			resp.Body = &models.RedirectResponse{}
		}
		if resp.StatusCode == 308 {
			resp.Body = &models.RedirectResponse{}
		}
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
			err = fmt.Errorf("%d is unknown to GetBDTPolicy", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param bdtPolicyId String identifying the individual BDT policy resource in the PCF.
@return *models.BdtPolicy,
*/
func UpdateBDTPolicy(client sbi.ConsumerClient, bdtPolicyId string, body models.PatchBdtPolicy) (result models.BdtPolicy, err error) {

	if len(bdtPolicyId) == 0 {
		err = fmt.Errorf("bdtPolicyId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPatch

	req.Path = fmt.Sprintf("%s/bdtpolicies/{bdtPolicyId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"bdtPolicyId"+"}", url.PathEscape(bdtPolicyId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/merge-patch+json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 307 {
			resp.Body = &models.RedirectResponse{}
		}
		if resp.StatusCode == 308 {
			resp.Body = &models.RedirectResponse{}
		}
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
			err = fmt.Errorf("%d is unknown to UpdateBDTPolicy", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package ampc

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICE_PATH = "{apiRoot}/npcf-am-policy-control/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@return *models.PolicyAssociation,
*/
func CreateIndividualAMPolicyAssociation(client sbi.ConsumerClient, body models.PolicyAssociationRequest) (result models.PolicyAssociation, err error) {

	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/policies", SERVICE_PATH)
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
			err = fmt.Errorf("%d is unknown to CreateIndividualAMPolicyAssociation", resp.StatusCode)
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
@param polAssoId Identifier of a policy association
@return
*/
func DeleteIndividualAMPolicyAssociation(client sbi.ConsumerClient, polAssoId string) (err error) {

	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/policies/{polAssoId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"polAssoId"+"}", url.PathEscape(polAssoId), -1)
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
			err = fmt.Errorf("%d is unknown to DeleteIndividualAMPolicyAssociation", resp.StatusCode)
			return
		}
	}

	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param polAssoId Identifier of a policy association
@return *models.PolicyAssociation,
*/
func ReadIndividualAMPolicyAssociation(client sbi.ConsumerClient, polAssoId string) (result models.PolicyAssociation, err error) {

	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/policies/{polAssoId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"polAssoId"+"}", url.PathEscape(polAssoId), -1)
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
			err = fmt.Errorf("%d is unknown to ReadIndividualAMPolicyAssociation", resp.StatusCode)
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
@param polAssoId Identifier of a policy association
@return *models.PolicyUpdate,
*/
func ReportObservedEventTriggersForIndividualAMPolicyAssociation(client sbi.ConsumerClient, polAssoId string, body models.PolicyAssociationUpdateRequest) (result models.PolicyUpdate, err error) {

	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/policies/{polAssoId}/update", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"polAssoId"+"}", url.PathEscape(polAssoId), -1)
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
			err = fmt.Errorf("%d is unknown to ReportObservedEventTriggersForIndividualAMPolicyAssociation", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

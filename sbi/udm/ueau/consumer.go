/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package ueau

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"etrib5gc/sbi/utils"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	SERVICE_PATH = "{apiRoot}/nudm-ueau/v1"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param supi SUPI of the user
@return *models.AuthEvent,
*/
func ConfirmAuth(client sbi.ConsumerClient, supi string, body models.AuthEvent) (result models.AuthEvent, err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/{supi}/auth-events", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
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
		if resp.StatusCode == 404 {
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
			err = fmt.Errorf("%d is unknown to ConfirmAuth", resp.StatusCode)
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
@param supi SUPI of the user
@param authEventId authEvent Id
@return
*/
func DeleteAuth(client sbi.ConsumerClient, supi string, authEventId string, body models.AuthEvent) (err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	if len(authEventId) == 0 {
		err = fmt.Errorf("authEventId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/{supi}/auth-events/{authEventId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
	req.Path = strings.Replace(req.Path, "{"+"authEventId"+"}", url.PathEscape(authEventId), -1)
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
		if resp.StatusCode == 400 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
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
			err = fmt.Errorf("%d is unknown to DeleteAuth", resp.StatusCode)
			return
		}
	}

	return
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param supiOrSuci SUPI or SUCI of the user
@return *models.AuthenticationInfoResult,
*/
func GenerateAuthData(client sbi.ConsumerClient, supiOrSuci string, body models.AuthenticationInfoRequest) (result models.AuthenticationInfoResult, err error) {

	if len(supiOrSuci) == 0 {
		err = fmt.Errorf("supiOrSuci is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/{supiOrSuci}/security-information/generate-auth-data", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"supiOrSuci"+"}", url.PathEscape(supiOrSuci), -1)
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 500 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 501 {
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
			err = fmt.Errorf("%d is unknown to GenerateAuthData", resp.StatusCode)
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
@param supi SUPI of the user
@param hssAuthType Type of AV requested by HSS
@return *models.HssAuthenticationInfoResult,
*/
func GenerateAv(client sbi.ConsumerClient, supi string, hssAuthType models.HssAuthTypeInUri, body models.HssAuthenticationInfoRequest) (result models.HssAuthenticationInfoResult, err error) {

	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	hssAuthTypeStr := utils.Param2String(hssAuthType)
	if len(hssAuthTypeStr) == 0 {
		err = fmt.Errorf("hssAuthType is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/{supi}/hss-security-information/{hssAuthType}/generate-av", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"supi"+"}", url.PathEscape(supi), -1)
	req.Path = strings.Replace(req.Path, "{"+"hssAuthType"+"}", url.PathEscape(hssAuthTypeStr), -1)
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 500 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 501 {
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
			err = fmt.Errorf("%d is unknown to GenerateAv", resp.StatusCode)
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
@param supiOrSuci SUPI or SUCI of the user
@param authenticatedInd Authenticated indication
@param supportedFeatures Supported Features
@param plmnId serving PLMN ID
@param ifNoneMatch Validator for conditional requests, as described in RFC 7232, 3.2
@param ifModifiedSince Validator for conditional requests, as described in RFC 7232, 3.3
@return *models.RgAuthCtx,
*/
func GetRgAuthData(client sbi.ConsumerClient, supiOrSuci string, authenticatedInd bool, supportedFeatures string, plmnId *models.PlmnId, ifNoneMatch string, ifModifiedSince string) (result models.RgAuthCtx, err error) {

	if len(supiOrSuci) == 0 {
		err = fmt.Errorf("supiOrSuci is required")
		return
	}
	authenticatedIndStr := utils.Param2String(authenticatedInd)
	if len(authenticatedIndStr) == 0 {
		err = fmt.Errorf("authenticatedInd is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/{supiOrSuci}/security-information-rg", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"supiOrSuci"+"}", url.PathEscape(supiOrSuci), -1)
	req.QueryParams.Add("authenticated-ind", authenticatedIndStr)

	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}
	plmnIdStr := utils.Param2String(plmnId)
	if len(plmnIdStr) > 0 {
		req.QueryParams.Add("plmn-id", plmnIdStr)
	}
	if len(ifNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = ifNoneMatch
	}
	if len(ifModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = ifModifiedSince
	}
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
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
			err = fmt.Errorf("%d is unknown to GetRgAuthData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

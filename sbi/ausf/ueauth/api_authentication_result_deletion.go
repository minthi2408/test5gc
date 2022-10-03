/*
AUSF API

AUSF UE Authentication Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package ueauth

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
@param authCtxId
@return 
*/
func Delete5gAkaAuthenticationResult(client sbi.ConsumerClient, authCtxId string) (err error) {
	
	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/ue-authentications/{authCtxId}/5g-aka-confirmation", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"authCtxId"+"}", url.PathEscape(authCtxId), -1)	
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
			err = fmt.Errorf("%d is unknown to Delete5gAkaAuthenticationResult", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for Delete5gAkaAuthenticationResult
func OnDelete5gAkaAuthenticationResult(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(AuthenticationResultDeletionApiHandler)
	
	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "authCtxId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.UEAUTH_HandleDelete5gAkaAuthenticationResult(authCtxId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param authCtxId
@return 
*/
func DeleteEapAuthenticationResult(client sbi.ConsumerClient, authCtxId string) (err error) {
	
	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/ue-authentications/{authCtxId}/eap-session", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"authCtxId"+"}", url.PathEscape(authCtxId), -1)	
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
			err = fmt.Errorf("%d is unknown to DeleteEapAuthenticationResult", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for DeleteEapAuthenticationResult
func OnDeleteEapAuthenticationResult(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(AuthenticationResultDeletionApiHandler)
	
	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "authCtxId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.UEAUTH_HandleDeleteEapAuthenticationResult(authCtxId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}




type AuthenticationResultDeletionApiHandler interface {
	UEAUTH_HandleDelete5gAkaAuthenticationResult(authCtxId string) (successCode int32, err *sbi.ApiError)
	UEAUTH_HandleDeleteEapAuthenticationResult(authCtxId string) (successCode int32, err *sbi.ApiError)
}

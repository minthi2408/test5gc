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
@return *models.EapSession, 
*/
func EapAuthMethod(client sbi.ConsumerClient, authCtxId string, body *models.EapSession) (result models.EapSession, err error) {
	
	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/ue-authentications/{authCtxId}/eap-session", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"authCtxId"+"}", url.PathEscape(authCtxId), -1)
	req.Body = body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/json, application/3gppHal+json, application/problem+json"
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
		if resp.StatusCode == 500 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to EapAuthMethod", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for EapAuthMethod
func OnEapAuthMethod(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(DefaultApiHandler)
	
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

	var input models.EapSession

	var apierr *sbi.ApiError
	var successCode int32
	var result models.EapSession

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UEAUTH_HandleEapAuthMethod(authCtxId, &input)
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
@return *models.RgAuthCtx, 
*/
func RgAuthenticationsPost(client sbi.ConsumerClient, body models.RgAuthenticationInfo) (result models.RgAuthCtx, err error) {
	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/rg-authentications", ServicePath)
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 400 {
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
			err = fmt.Errorf("%d is unknown to RgAuthenticationsPost", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for RgAuthenticationsPost
func OnRgAuthenticationsPost(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(DefaultApiHandler)
	

	var input models.RgAuthenticationInfo

	var apierr *sbi.ApiError
	var successCode int32
	var result models.RgAuthCtx

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UEAUTH_HandleRgAuthenticationsPost(input)
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
@param authCtxId
@return *models.ConfirmationDataResponse, 
*/
func UeAuthenticationsAuthCtxId5gAkaConfirmationPut(client sbi.ConsumerClient, authCtxId string, body *models.ConfirmationData) (result models.ConfirmationDataResponse, err error) {
	
	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/ue-authentications/{authCtxId}/5g-aka-confirmation", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"authCtxId"+"}", url.PathEscape(authCtxId), -1)
	req.Body = body
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
		if resp.StatusCode == 500 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to UeAuthenticationsAuthCtxId5gAkaConfirmationPut", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for UeAuthenticationsAuthCtxId5gAkaConfirmationPut
func OnUeAuthenticationsAuthCtxId5gAkaConfirmationPut(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(DefaultApiHandler)
	
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

	var input models.ConfirmationData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.ConfirmationDataResponse

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UEAUTH_HandleUeAuthenticationsAuthCtxId5gAkaConfirmationPut(authCtxId, &input)
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
@return 
*/
func UeAuthenticationsDeregisterPost(client sbi.ConsumerClient, body models.DeregistrationInfo) (err error) {
	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/ue-authentications/deregister", ServicePath)
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
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to UeAuthenticationsDeregisterPost", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for UeAuthenticationsDeregisterPost
func OnUeAuthenticationsDeregisterPost(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(DefaultApiHandler)
	

	var input models.DeregistrationInfo

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.UEAUTH_HandleUeAuthenticationsDeregisterPost(input)
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
@return *models.UEAuthenticationCtx, 
*/
func UeAuthenticationsPost(client sbi.ConsumerClient, body models.AuthenticationInfo) (result models.UEAuthenticationCtx, err error) {
	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/ue-authentications", ServicePath)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
	req.HeaderParams["Accept"] = "application/3gppHal+json, application/json, application/problem+json"
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
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to UeAuthenticationsPost", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for UeAuthenticationsPost
func OnUeAuthenticationsPost(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(DefaultApiHandler)
	

	var input models.AuthenticationInfo

	var apierr *sbi.ApiError
	var successCode int32
	var result models.UEAuthenticationCtx

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UEAUTH_HandleUeAuthenticationsPost(input)
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




type DefaultApiHandler interface {
	UEAUTH_HandleEapAuthMethod(authCtxId string, body *models.EapSession) (successCode int32, result models.EapSession, err *sbi.ApiError)
	UEAUTH_HandleRgAuthenticationsPost(body models.RgAuthenticationInfo) (successCode int32, result models.RgAuthCtx, err *sbi.ApiError)
	UEAUTH_HandleUeAuthenticationsAuthCtxId5gAkaConfirmationPut(authCtxId string, body *models.ConfirmationData) (successCode int32, result models.ConfirmationDataResponse, err *sbi.ApiError)
	UEAUTH_HandleUeAuthenticationsDeregisterPost(body models.DeregistrationInfo) (successCode int32, err *sbi.ApiError)
	UEAUTH_HandleUeAuthenticationsPost(body models.AuthenticationInfo) (successCode int32, result models.UEAuthenticationCtx, err *sbi.ApiError)
}

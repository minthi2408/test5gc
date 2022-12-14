/*
AUSF API

AUSF UE Authentication Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package uea

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"net/http"
)

//sbi producer handler for Delete5gAkaAuthenticationResult
func OnDelete5gAkaAuthenticationResult(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "authCtxId is required",
		}))
		return
	}

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.UEA_HandleDelete5gAkaAuthenticationResult(authCtxId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}

//sbi producer handler for DeleteEapAuthenticationResult
func OnDeleteEapAuthenticationResult(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "authCtxId is required",
		}))
		return
	}

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.UEA_HandleDeleteEapAuthenticationResult(authCtxId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}

//sbi producer handler for EapAuthMethod
func OnEapAuthMethod(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
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
		successCode, result, apierr = prod.UEA_HandleEapAuthMethod(authCtxId, &input)
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

//sbi producer handler for RgAuthenticationsPost
func OnRgAuthenticationsPost(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	var input models.RgAuthenticationInfo

	var apierr *sbi.ApiError
	var successCode int32
	var result models.RgAuthCtx

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UEA_HandleRgAuthenticationsPost(input)
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

//sbi producer handler for UeAuthenticationsAuthCtxId5gAkaConfirmationPut
func OnUeAuthenticationsAuthCtxId5gAkaConfirmationPut(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	authCtxId := ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		//authCtxId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
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
		successCode, result, apierr = prod.UEA_HandleUeAuthenticationsAuthCtxId5gAkaConfirmationPut(authCtxId, &input)
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

//sbi producer handler for UeAuthenticationsDeregisterPost
func OnUeAuthenticationsDeregisterPost(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	var input models.DeregistrationInfo

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.UEA_HandleUeAuthenticationsDeregisterPost(input)
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

//sbi producer handler for UeAuthenticationsPost
func OnUeAuthenticationsPost(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	var input models.AuthenticationInfo

	var apierr *sbi.ApiError
	var successCode int32
	var result models.UEAuthenticationCtx

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UEA_HandleUeAuthenticationsPost(input)
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

type Producer interface {
	UEA_HandleDelete5gAkaAuthenticationResult(authCtxId string) (successCode int32, err *sbi.ApiError)
	UEA_HandleDeleteEapAuthenticationResult(authCtxId string) (successCode int32, err *sbi.ApiError)
	UEA_HandleEapAuthMethod(authCtxId string, body *models.EapSession) (successCode int32, result models.EapSession, err *sbi.ApiError)
	UEA_HandleRgAuthenticationsPost(body models.RgAuthenticationInfo) (successCode int32, result models.RgAuthCtx, err *sbi.ApiError)
	UEA_HandleUeAuthenticationsAuthCtxId5gAkaConfirmationPut(authCtxId string, body *models.ConfirmationData) (successCode int32, result models.ConfirmationDataResponse, err *sbi.ApiError)
	UEA_HandleUeAuthenticationsDeregisterPost(body models.DeregistrationInfo) (successCode int32, err *sbi.ApiError)
	UEA_HandleUeAuthenticationsPost(body models.AuthenticationInfo) (successCode int32, result models.UEAuthenticationCtx, err *sbi.ApiError)
}

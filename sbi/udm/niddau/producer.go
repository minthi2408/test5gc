/*
Nudm_NIDDAU

Nudm NIDD Authorization Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package niddau

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
	"net/http"
)

//sbi producer handler for AuthorizeNiddData
func OnAuthorizeNiddData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	ueIdentity := ctx.Param("ueIdentity")
	if len(ueIdentity) == 0 {
		//ueIdentity is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueIdentity is required",
		}))
		return
	}

	var input models.AuthorizationInfo

	var apierr *sbi.ApiError
	var successCode int32
	var result models.AuthorizationData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.NIDDAU_HandleAuthorizeNiddData(ueIdentity, input)
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
	NIDDAU_HandleAuthorizeNiddData(ueIdentity string, body models.AuthorizationInfo) (successCode int32, result models.AuthorizationData, err *sbi.ApiError)
}
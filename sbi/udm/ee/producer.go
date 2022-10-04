/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package ee

import (
	"net/http"
	"etri5gc/sbi"
	"etri5gc/sbi/models"
)

//sbi producer handler for CreateEeSubscription
func OnCreateEeSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueIdentity := ctx.Param("ueIdentity")
	if len(ueIdentity) == 0 {
		//ueIdentity is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueIdentity is required",
		}))
		return
	}

	var input models.EeSubscription

	var apierr *sbi.ApiError
	var successCode int32
	var result models.CreatedEeSubscription

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.EE_HandleCreateEeSubscription(ueIdentity, input)
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


//sbi producer handler for DeleteEeSubscription
func OnDeleteEeSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueIdentity := ctx.Param("ueIdentity")
	if len(ueIdentity) == 0 {
		//ueIdentity is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueIdentity is required",
		}))
		return
	}
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

	successCode, apierr = prod.EE_HandleDeleteEeSubscription(ueIdentity, subscriptionId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}


//sbi producer handler for UpdateEeSubscription
func OnUpdateEeSubscription(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueIdentity := ctx.Param("ueIdentity")
	if len(ueIdentity) == 0 {
		//ueIdentity is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueIdentity is required",
		}))
		return
	}
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
	supportedFeatures := ctx.Param("supported-features")

	var input []models.PatchItem

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PatchResult

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.EE_HandleUpdateEeSubscription(ueIdentity, subscriptionId, supportedFeatures, input)
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
	EE_HandleCreateEeSubscription(ueIdentity string, body models.EeSubscription) (successCode int32, result models.CreatedEeSubscription, err *sbi.ApiError)
	EE_HandleDeleteEeSubscription(ueIdentity string, subscriptionId string) (successCode int32, err *sbi.ApiError)
	EE_HandleUpdateEeSubscription(ueIdentity string, subscriptionId string, supportedFeatures string, body []models.PatchItem) (successCode int32, result models.PatchResult, err *sbi.ApiError)
}



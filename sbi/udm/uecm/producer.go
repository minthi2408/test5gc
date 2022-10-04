/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package uecm

import (
	"net/http"
	"etri5gc/sbi"
	"etri5gc/sbi/models"
	"etri5gc/sbi/utils"
)

//sbi producer handler for Get3GppRegistration
func OnGet3GppRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	supportedFeatures := ctx.Param("supported-features")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.Amf3GppAccessRegistration


	successCode, result, apierr = prod.UECM_HandleGet3GppRegistration(ueId, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


//sbi producer handler for GetNon3GppRegistration
func OnGetNon3GppRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	supportedFeatures := ctx.Param("supported-features")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.AmfNon3GppAccessRegistration


	successCode, result, apierr = prod.UECM_HandleGetNon3GppRegistration(ueId, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


//sbi producer handler for Call3GppRegistration
func OnCall3GppRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	var input models.Amf3GppAccessRegistration

	var apierr *sbi.ApiError
	var successCode int32
	var result models.Amf3GppAccessRegistration

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UECM_HandleCall3GppRegistration(ueId, input)
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


//sbi producer handler for Non3GppRegistration
func OnNon3GppRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	var input models.AmfNon3GppAccessRegistration

	var apierr *sbi.ApiError
	var successCode int32
	var result models.AmfNon3GppAccessRegistration

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UECM_HandleNon3GppRegistration(ueId, input)
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


//sbi producer handler for IpSmGwDeregistration
func OnIpSmGwDeregistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.UECM_HandleIpSmGwDeregistration(ueId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}


//sbi producer handler for IpSmGwRegistration
func OnIpSmGwRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	var input models.IpSmGwRegistration

	var apierr *sbi.ApiError
	var successCode int32
	var result models.IpSmGwRegistration

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UECM_HandleIpSmGwRegistration(ueId, input)
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


//sbi producer handler for GetIpSmGwRegistration
func OnGetIpSmGwRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.IpSmGwRegistration


	successCode, result, apierr = prod.UECM_HandleGetIpSmGwRegistration(ueId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


//sbi producer handler for PeiUpdate
func OnPeiUpdate(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	var input models.PeiUpdateInfo

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.UECM_HandlePeiUpdate(ueId, input)
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


//sbi producer handler for Update3GppRegistration
func OnUpdate3GppRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	supportedFeatures := ctx.Param("supported-features")

	var input models.Amf3GppAccessRegistrationModification

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PatchResult

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UECM_HandleUpdate3GppRegistration(ueId, supportedFeatures, input)
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


//sbi producer handler for UpdateNon3GppRegistration
func OnUpdateNon3GppRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	supportedFeatures := ctx.Param("supported-features")

	var input models.AmfNon3GppAccessRegistrationModification

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PatchResult

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UECM_HandleUpdateNon3GppRegistration(ueId, supportedFeatures, input)
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


//sbi producer handler for RetrieveSmfRegistration
func OnRetrieveSmfRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		//pduSessionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "pduSessionId is required",
		}))
		return
	}
	var pduSessionId *int32
	var pduSessionIdErr error
	if pduSessionId, pduSessionIdErr = utils.String2Int32(pduSessionIdStr); pduSessionIdErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: pduSessionIdErr.Error(), 
		}))
		return
	}
	

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SmfRegistration


	successCode, result, apierr = prod.UECM_HandleRetrieveSmfRegistration(ueId, *pduSessionId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


//sbi producer handler for SmfDeregistration
func OnSmfDeregistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		//pduSessionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "pduSessionId is required",
		}))
		return
	}
	var pduSessionId *int32
	var pduSessionIdErr error
	if pduSessionId, pduSessionIdErr = utils.String2Int32(pduSessionIdStr); pduSessionIdErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: pduSessionIdErr.Error(), 
		}))
		return
	}
	
	smfSetId := ctx.Param("smf-set-id")
	smfInstanceId := ctx.Param("smf-instance-id")

	

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.UECM_HandleSmfDeregistration(ueId, *pduSessionId, smfSetId, smfInstanceId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}


//sbi producer handler for GetSmfRegistration
func OnGetSmfRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	singleNssaiStr := ctx.Param("single-nssai")
	var singleNssai *models.Snssai
	var singleNssaiErr error
	if singleNssai, singleNssaiErr = utils.String2Snssai(singleNssaiStr); singleNssaiErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: singleNssaiErr.Error(), 
		}))
		return
	}
	
	dnn := ctx.Param("dnn")
	supportedFeatures := ctx.Param("supported-features")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SmfRegistrationInfo


	successCode, result, apierr = prod.UECM_HandleGetSmfRegistration(ueId, singleNssai, dnn, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


//sbi producer handler for Registration
func OnRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		//pduSessionId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "pduSessionId is required",
		}))
		return
	}
	var pduSessionId *int32
	var pduSessionIdErr error
	if pduSessionId, pduSessionIdErr = utils.String2Int32(pduSessionIdStr); pduSessionIdErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: pduSessionIdErr.Error(), 
		}))
		return
	}
	

	var input models.SmfRegistration

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SmfRegistration

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UECM_HandleRegistration(ueId, *pduSessionId, input)
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


//sbi producer handler for Get3GppSmsfRegistration
func OnGet3GppSmsfRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	supportedFeatures := ctx.Param("supported-features")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SmsfRegistration


	successCode, result, apierr = prod.UECM_HandleGet3GppSmsfRegistration(ueId, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


//sbi producer handler for Call3GppSmsfDeregistration
func OnCall3GppSmsfDeregistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	smsfSetId := ctx.Param("smsf-set-id")

	

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.UECM_HandleCall3GppSmsfDeregistration(ueId, smsfSetId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}


//sbi producer handler for Non3GppSmsfDeregistration
func OnNon3GppSmsfDeregistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	smsfSetId := ctx.Param("smsf-set-id")

	

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.UECM_HandleNon3GppSmsfDeregistration(ueId, smsfSetId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}


//sbi producer handler for GetNon3GppSmsfRegistration
func OnGetNon3GppSmsfRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	supportedFeatures := ctx.Param("supported-features")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SmsfRegistration


	successCode, result, apierr = prod.UECM_HandleGetNon3GppSmsfRegistration(ueId, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


//sbi producer handler for Call3GppSmsfRegistration
func OnCall3GppSmsfRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	var input models.SmsfRegistration

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SmsfRegistration

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UECM_HandleCall3GppSmsfRegistration(ueId, input)
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


//sbi producer handler for Non3GppSmsfRegistration
func OnNon3GppSmsfRegistration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	var input models.SmsfRegistration

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SmsfRegistration

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.UECM_HandleNon3GppSmsfRegistration(ueId, input)
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


//sbi producer handler for DeregAMF
func OnDeregAMF(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}

	var input models.AmfDeregInfo

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.UECM_HandleDeregAMF(ueId, input)
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


//sbi producer handler for TriggerPCSCFRestoration
func OnTriggerPCSCFRestoration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	

	var input models.TriggerRequest

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.UECM_HandleTriggerPCSCFRestoration(input)
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


//sbi producer handler for GetRegistrations
func OnGetRegistrations(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	registrationDatasetNamesStr := ctx.Param("registration-dataset-names")
	if len(registrationDatasetNamesStr) == 0 {
		//registrationDatasetNames is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "registrationDatasetNames is required",
		}))
		return
	}
	var registrationDatasetNames []models.RegistrationDataSetName
	var registrationDatasetNamesErr error
	if registrationDatasetNames, registrationDatasetNamesErr = utils.String2ArrayOfRegistrationDataSetName(registrationDatasetNamesStr); registrationDatasetNamesErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: registrationDatasetNamesErr.Error(), 
		}))
		return
	}
	
	supportedFeatures := ctx.Param("supported-features")
	singleNssaiStr := ctx.Param("single-nssai")
	var singleNssai *models.Snssai
	var singleNssaiErr error
	if singleNssai, singleNssaiErr = utils.String2Snssai(singleNssaiStr); singleNssaiErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: singleNssaiErr.Error(), 
		}))
		return
	}
	
	dnn := ctx.Param("dnn")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.RegistrationDataSets


	successCode, result, apierr = prod.UECM_HandleGetRegistrations(ueId, registrationDatasetNames, supportedFeatures, singleNssai, dnn)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


//sbi producer handler for GetLocationInfo
func OnGetLocationInfo(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)
	
	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	supportedFeatures := ctx.Param("supported-features")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.LocationInfo


	successCode, result, apierr = prod.UECM_HandleGetLocationInfo(ueId, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}


type Producer interface {
	UECM_HandleGet3GppRegistration(ueId string, supportedFeatures string) (successCode int32, result models.Amf3GppAccessRegistration, err *sbi.ApiError)
	UECM_HandleGetNon3GppRegistration(ueId string, supportedFeatures string) (successCode int32, result models.AmfNon3GppAccessRegistration, err *sbi.ApiError)
	UECM_HandleCall3GppRegistration(ueId string, body models.Amf3GppAccessRegistration) (successCode int32, result models.Amf3GppAccessRegistration, err *sbi.ApiError)
	UECM_HandleNon3GppRegistration(ueId string, body models.AmfNon3GppAccessRegistration) (successCode int32, result models.AmfNon3GppAccessRegistration, err *sbi.ApiError)
	UECM_HandleIpSmGwDeregistration(ueId string) (successCode int32, err *sbi.ApiError)
	UECM_HandleIpSmGwRegistration(ueId string, body models.IpSmGwRegistration) (successCode int32, result models.IpSmGwRegistration, err *sbi.ApiError)
	UECM_HandleGetIpSmGwRegistration(ueId string) (successCode int32, result models.IpSmGwRegistration, err *sbi.ApiError)
	UECM_HandlePeiUpdate(ueId string, body models.PeiUpdateInfo) (successCode int32, err *sbi.ApiError)
	UECM_HandleUpdate3GppRegistration(ueId string, supportedFeatures string, body models.Amf3GppAccessRegistrationModification) (successCode int32, result models.PatchResult, err *sbi.ApiError)
	UECM_HandleUpdateNon3GppRegistration(ueId string, supportedFeatures string, body models.AmfNon3GppAccessRegistrationModification) (successCode int32, result models.PatchResult, err *sbi.ApiError)
	UECM_HandleRetrieveSmfRegistration(ueId string, pduSessionId int32) (successCode int32, result models.SmfRegistration, err *sbi.ApiError)
	UECM_HandleSmfDeregistration(ueId string, pduSessionId int32, smfSetId string, smfInstanceId string) (successCode int32, err *sbi.ApiError)
	UECM_HandleGetSmfRegistration(ueId string, singleNssai *models.Snssai, dnn string, supportedFeatures string) (successCode int32, result models.SmfRegistrationInfo, err *sbi.ApiError)
	UECM_HandleRegistration(ueId string, pduSessionId int32, body models.SmfRegistration) (successCode int32, result models.SmfRegistration, err *sbi.ApiError)
	UECM_HandleGet3GppSmsfRegistration(ueId string, supportedFeatures string) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError)
	UECM_HandleCall3GppSmsfDeregistration(ueId string, smsfSetId string) (successCode int32, err *sbi.ApiError)
	UECM_HandleNon3GppSmsfDeregistration(ueId string, smsfSetId string) (successCode int32, err *sbi.ApiError)
	UECM_HandleGetNon3GppSmsfRegistration(ueId string, supportedFeatures string) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError)
	UECM_HandleCall3GppSmsfRegistration(ueId string, body models.SmsfRegistration) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError)
	UECM_HandleNon3GppSmsfRegistration(ueId string, body models.SmsfRegistration) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError)
	UECM_HandleDeregAMF(ueId string, body models.AmfDeregInfo) (successCode int32, err *sbi.ApiError)
	UECM_HandleTriggerPCSCFRestoration(body models.TriggerRequest) (successCode int32, err *sbi.ApiError)
	UECM_HandleGetRegistrations(ueId string, registrationDatasetNames []models.RegistrationDataSetName, supportedFeatures string, singleNssai *models.Snssai, dnn string) (successCode int32, result models.RegistrationDataSets, err *sbi.ApiError)
	UECM_HandleGetLocationInfo(ueId string, supportedFeatures string) (successCode int32, result models.LocationInfo, err *sbi.ApiError)
}



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
	"net/http"
)

//sbi producer handler for CreateIndividualAMPolicyAssociation
func OnCreateIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	var input models.PolicyAssociationRequest

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PolicyAssociation

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.AMPC_HandleCreateIndividualAMPolicyAssociation(input)
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

//sbi producer handler for DeleteIndividualAMPolicyAssociation
func OnDeleteIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	polAssoId := ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		//polAssoId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "polAssoId is required",
		}))
		return
	}

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.AMPC_HandleDeleteIndividualAMPolicyAssociation(polAssoId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}

//sbi producer handler for ReadIndividualAMPolicyAssociation
func OnReadIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	polAssoId := ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		//polAssoId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "polAssoId is required",
		}))
		return
	}

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PolicyAssociation

	successCode, result, apierr = prod.AMPC_HandleReadIndividualAMPolicyAssociation(polAssoId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

//sbi producer handler for ReportObservedEventTriggersForIndividualAMPolicyAssociation
func OnReportObservedEventTriggersForIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	polAssoId := ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		//polAssoId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "polAssoId is required",
		}))
		return
	}

	var input models.PolicyAssociationUpdateRequest

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PolicyUpdate

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.AMPC_HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(polAssoId, input)
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
	AMPC_HandleCreateIndividualAMPolicyAssociation(body models.PolicyAssociationRequest) (successCode int32, result models.PolicyAssociation, err *sbi.ApiError)
	AMPC_HandleDeleteIndividualAMPolicyAssociation(polAssoId string) (successCode int32, err *sbi.ApiError)
	AMPC_HandleReadIndividualAMPolicyAssociation(polAssoId string) (successCode int32, result models.PolicyAssociation, err *sbi.ApiError)
	AMPC_HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(polAssoId string, body models.PolicyAssociationUpdateRequest) (successCode int32, result models.PolicyUpdate, err *sbi.ApiError)
}

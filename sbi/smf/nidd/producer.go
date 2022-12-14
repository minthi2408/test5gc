/*
Nsmf_NIDD

SMF NIDD Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package nidd

import (
	"etrib5gc/sbi"
	"etrib5gc/sbi/models"
	"net/http"
)

//sbi producer handler for Deliver
func OnDeliver(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Producer)

	pduSessionRef := ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		//pduSessionRef is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "pduSessionRef is required",
		}))
		return
	}

	var input models.DeliverRequest

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.NIDD_HandleDeliver(pduSessionRef, input)
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

type Producer interface {
	NIDD_HandleDeliver(pduSessionRef string, body models.DeliverRequest) (successCode int32, err *sbi.ApiError)
}

/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package dr

import (
	"etri5gc/sbi"
	"etri5gc/sbi/models"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId UE id
@param supportedFeatures Supported Features
@return
*/
func CreateOrUpdateNssaiAck(client sbi.ConsumerClient, ueId string, supportedFeatures string, body models.NssaiAckData) (err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/ue-update-confirmation-data/subscribed-snssais", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"

	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to CreateOrUpdateNssaiAck", resp.StatusCode)
			return
		}
	}

	return
}

//sbi producer handler for CreateOrUpdateNssaiAck
func OnCreateOrUpdateNssaiAck(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(NSSAIUpdateAckDocumentApiHandler)

	ueId := ctx.Param("ueId")
	if len(ueId) == 0 {
		//ueId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueId is required",
		}))
		return
	}
	supportedFeatures := ctx.Param("supported-features")

	var input models.NssaiAckData

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.DR_HandleCreateOrUpdateNssaiAck(ueId, supportedFeatures, input)
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

type NSSAIUpdateAckDocumentApiHandler interface {
	DR_HandleCreateOrUpdateNssaiAck(ueId string, supportedFeatures string, body models.NssaiAckData) (successCode int32, err *sbi.ApiError)
}

/*
Namf_Communication

AMF Communication Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package comm

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
@param ueContextId UE Context Identifier
@return *models.N1N2MessageTransferRspData, 
*/
func N1N2MessageTransfer(client sbi.ConsumerClient, ueContextId string, body models.N1N2MessageTransferReqData) (result models.N1N2MessageTransferRspData, err error) {
	
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPost

	req.Path = fmt.Sprintf("%s/ue-contexts/{ueContextId}/n1-n2-messages", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueContextId"+"}", url.PathEscape(ueContextId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "multipart/related"
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 409 {
			resp.Body = &models.N1N2MessageTransferError{}
		}
		if resp.StatusCode == 411 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 413 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 415 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 429 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 500 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 503 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 504 {
			resp.Body = &models.N1N2MessageTransferError{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to N1N2MessageTransfer", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for N1N2MessageTransfer
func OnN1N2MessageTransfer(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(N1N2MessageCollectionDocumentApiHandler)
	
	ueContextId := ctx.Param("ueContextId")
	if len(ueContextId) == 0 {
		//ueContextId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "ueContextId is required",
		}))
		return
	}

	var input models.N1N2MessageTransferReqData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.N1N2MessageTransferRspData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.COMM_HandleN1N2MessageTransfer(ueContextId, input)
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




type N1N2MessageCollectionDocumentApiHandler interface {
	COMM_HandleN1N2MessageTransfer(ueContextId string, body models.N1N2MessageTransferReqData) (successCode int32, result models.N1N2MessageTransferRspData, err *sbi.ApiError)
}

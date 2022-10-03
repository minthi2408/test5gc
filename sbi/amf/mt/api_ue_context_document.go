/*
Namf_MT

AMF Mobile Terminated Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package mt

import (
	"fmt"
	"net/http"
	"net/url"
	"etri5gc/sbi"
	"etri5gc/sbi/utils"
	"etri5gc/sbi/models"
	"strings"
)


/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueContextId UE Context Identifier
@param infoClass UE Context Information Class
@param supportedFeatures Supported Features
@param oldGuami Old GUAMI
@return *models.UeContextInfo, 
*/
func ProvideDomainSelectionInfo(client sbi.ConsumerClient, ueContextId string, infoClass *models.UeContextInfoClass, supportedFeatures string, oldGuami *models.Guami) (result models.UeContextInfo, err error) {
	
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}			
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/ue-contexts/{ueContextId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueContextId"+"}", url.PathEscape(ueContextId), -1)
	infoClassStr := utils.Param2String(infoClass)
	if len(infoClassStr) > 0 {
		req.QueryParams.Add("info-class", infoClassStr)
	}
	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}
	oldGuamiStr := utils.Param2String(oldGuami)
	if len(oldGuamiStr) > 0 {
		req.QueryParams.Add("old-guami", oldGuamiStr)
	}	
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
		if resp.StatusCode == 414 {
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
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to ProvideDomainSelectionInfo", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for ProvideDomainSelectionInfo
func OnProvideDomainSelectionInfo(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(UeContextDocumentApiHandler)
	
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
	infoClassStr := ctx.Param("info-class")
	var infoClass *models.UeContextInfoClass
	var infoClassErr error
	if infoClass, infoClassErr = utils.String2UeContextInfoClass(infoClassStr); infoClassErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: infoClassErr.Error(), 
		}))
		return
	}
	
	supportedFeatures := ctx.Param("supported-features")
	oldGuamiStr := ctx.Param("old-guami")
	var oldGuami *models.Guami
	var oldGuamiErr error
	if oldGuami, oldGuamiErr = utils.String2Guami(oldGuamiStr); oldGuamiErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: oldGuamiErr.Error(), 
		}))
		return
	}
	

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.UeContextInfo


	successCode, result, apierr = prod.MT_HandleProvideDomainSelectionInfo(ueContextId, infoClass, supportedFeatures, oldGuami)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type UeContextDocumentApiHandler interface {
	MT_HandleProvideDomainSelectionInfo(ueContextId string, infoClass *models.UeContextInfoClass, supportedFeatures string, oldGuami *models.Guami) (successCode int32, result models.UeContextInfo, err *sbi.ApiError)
}

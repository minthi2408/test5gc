/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package dr

import (
	"fmt"
	"net/http"
	"net/url"
	"etri5gc/sbi"
	"etri5gc/sbi/models"
	"etri5gc/sbi/utils"
	"strings"
)


/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId
@param snssai
@param dnn
@param fields attributes to be retrieved
@param suppFeat Supported Features
@return *models.SmPolicyData, 
*/
func ReadSessionManagementPolicyData(client sbi.ConsumerClient, ueId string, snssai *models.Snssai, dnn string, fields []string, suppFeat string) (result models.SmPolicyData, err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}				
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/policy-data/ues/{ueId}/sm-data", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	snssaiStr := utils.Param2String(snssai)
	if len(snssaiStr) > 0 {
		req.QueryParams.Add("snssai", snssaiStr)
	}
	if len(dnn) > 0 {
		req.QueryParams.Add("dnn", dnn)
	}
	fieldsStr := utils.Param2String(fields)
	if len(fieldsStr) > 0 {
		req.QueryParams.Add("fields", fieldsStr)
	}
	if len(suppFeat) > 0 {
		req.QueryParams.Add("supp-feat", suppFeat)
	}	
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 400 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 401 {
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
			err = fmt.Errorf("%d is unknown to ReadSessionManagementPolicyData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for ReadSessionManagementPolicyData
func OnReadSessionManagementPolicyData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(SessionManagementPolicyDataDocumentApiHandler)
	
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
	snssaiStr := ctx.Param("snssai")
	var snssai *models.Snssai
	var snssaiErr error
	if snssai, snssaiErr = utils.String2Snssai(snssaiStr); snssaiErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: snssaiErr.Error(), 
		}))
		return
	}
	
	dnn := ctx.Param("dnn")
	fieldsStr := ctx.Param("fields")
	var fields []string
	var fieldsErr error
	if fields, fieldsErr = utils.String2ArrayOfstring(fieldsStr); fieldsErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: fieldsErr.Error(), 
		}))
		return
	}
	
	suppFeat := ctx.Param("supp-feat")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SmPolicyData


	successCode, result, apierr = prod.DR_HandleReadSessionManagementPolicyData(ueId, snssai, dnn, fields, suppFeat)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId
@return *models.SmPolicyData, 
*/
func UpdateSessionManagementPolicyData(client sbi.ConsumerClient, ueId string, body models.SmPolicyDataPatch) (result models.SmPolicyData, err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPatch

	req.Path = fmt.Sprintf("%s/policy-data/ues/{ueId}/sm-data", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/merge-patch+json"
	req.HeaderParams["Accept"] = "application/json, application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 400 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 401 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
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
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to UpdateSessionManagementPolicyData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for UpdateSessionManagementPolicyData
func OnUpdateSessionManagementPolicyData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(SessionManagementPolicyDataDocumentApiHandler)
	
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

	var input models.SmPolicyDataPatch

	var apierr *sbi.ApiError
	var successCode int32
	var result models.SmPolicyData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleUpdateSessionManagementPolicyData(ueId, input)
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




type SessionManagementPolicyDataDocumentApiHandler interface {
	DR_HandleReadSessionManagementPolicyData(ueId string, snssai *models.Snssai, dnn string, fields []string, suppFeat string) (successCode int32, result models.SmPolicyData, err *sbi.ApiError)
	DR_HandleUpdateSessionManagementPolicyData(ueId string, body models.SmPolicyDataPatch) (successCode int32, result models.SmPolicyData, err *sbi.ApiError)
}

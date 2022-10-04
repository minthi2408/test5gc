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
@param ueId UE id
@param pduSessionId PDU session id
@return *models.AccessAndMobilityData, 
*/
func CreateOrReplaceSessionManagementData(client sbi.ConsumerClient, ueId string, pduSessionId int32, body models.PduSessionManagementData) (result models.AccessAndMobilityData, err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}	
	pduSessionIdStr := utils.Param2String(pduSessionId)
	if len(pduSessionIdStr) == 0 {
		err = fmt.Errorf("pduSessionId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/exposure-data/{ueId}/session-management-data/{pduSessionId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"pduSessionId"+"}", url.PathEscape(pduSessionIdStr), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
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
			err = fmt.Errorf("%d is unknown to CreateOrReplaceSessionManagementData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for CreateOrReplaceSessionManagementData
func OnCreateOrReplaceSessionManagementData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(PduSessionManagementDataApiHandler)
	
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
	

	var input models.PduSessionManagementData

	var apierr *sbi.ApiError
	var successCode int32
	var result models.AccessAndMobilityData

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleCreateOrReplaceSessionManagementData(ueId, *pduSessionId, input)
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




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId UE id
@param pduSessionId PDU session id
@return 
*/
func DeleteSessionManagementData(client sbi.ConsumerClient, ueId string, pduSessionId int32) (err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}	
	pduSessionIdStr := utils.Param2String(pduSessionId)
	if len(pduSessionIdStr) == 0 {
		err = fmt.Errorf("pduSessionId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/exposure-data/{ueId}/session-management-data/{pduSessionId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"pduSessionId"+"}", url.PathEscape(pduSessionIdStr), -1)	
	req.HeaderParams["Accept"] = "application/problem+json"
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
			err = fmt.Errorf("%d is unknown to DeleteSessionManagementData", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for DeleteSessionManagementData
func OnDeleteSessionManagementData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(PduSessionManagementDataApiHandler)
	
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

	successCode, apierr = prod.DR_HandleDeleteSessionManagementData(ueId, *pduSessionId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId UE id
@param pduSessionId PDU session id
@param ipv4Addr IPv4 Address of the UE
@param ipv6Prefix IPv6 Address Prefix of the UE
@param dnn DNN of the UE
@param fields attributes to be retrieved
@param suppFeat Supported Features
@return *models.PduSessionManagementData, 
*/
func QuerySessionManagementData(client sbi.ConsumerClient, ueId string, pduSessionId int32, ipv4Addr string, ipv6Prefix string, dnn string, fields []string, suppFeat string) (result models.PduSessionManagementData, err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}	
	pduSessionIdStr := utils.Param2String(pduSessionId)
	if len(pduSessionIdStr) == 0 {
		err = fmt.Errorf("pduSessionId is required")
		return
	}					
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/exposure-data/{ueId}/session-management-data/{pduSessionId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"pduSessionId"+"}", url.PathEscape(pduSessionIdStr), -1)
	if len(ipv4Addr) > 0 {
		req.QueryParams.Add("ipv4-addr", ipv4Addr)
	}
	if len(ipv6Prefix) > 0 {
		req.QueryParams.Add("ipv6-prefix", ipv6Prefix)
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
			err = fmt.Errorf("%d is unknown to QuerySessionManagementData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for QuerySessionManagementData
func OnQuerySessionManagementData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(PduSessionManagementDataApiHandler)
	
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
	
	ipv4Addr := ctx.Param("ipv4-addr")
	ipv6Prefix := ctx.Param("ipv6-prefix")
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
	var result models.PduSessionManagementData


	successCode, result, apierr = prod.DR_HandleQuerySessionManagementData(ueId, *pduSessionId, ipv4Addr, ipv6Prefix, dnn, fields, suppFeat)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type PduSessionManagementDataApiHandler interface {
	DR_HandleCreateOrReplaceSessionManagementData(ueId string, pduSessionId int32, body models.PduSessionManagementData) (successCode int32, result models.AccessAndMobilityData, err *sbi.ApiError)
	DR_HandleDeleteSessionManagementData(ueId string, pduSessionId int32) (successCode int32, err *sbi.ApiError)
	DR_HandleQuerySessionManagementData(ueId string, pduSessionId int32, ipv4Addr string, ipv6Prefix string, dnn string, fields []string, suppFeat string) (successCode int32, result models.PduSessionManagementData, err *sbi.ApiError)
}

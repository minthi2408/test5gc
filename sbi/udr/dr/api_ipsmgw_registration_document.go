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
	"strings"
)


/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId UE id
@return 
*/
func CreateIpSmGwContext(client sbi.ConsumerClient, ueId string, body models.IpSmGwRegistration) (err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ip-sm-gw", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
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
			err = fmt.Errorf("%d is unknown to CreateIpSmGwContext", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for CreateIpSmGwContext
func OnCreateIpSmGwContext(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IPSMGWRegistrationDocumentApiHandler)
	
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
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.DR_HandleCreateIpSmGwContext(ueId, input)
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




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId UE id
@return 
*/
func DeleteIpSmGwContext(client sbi.ConsumerClient, ueId string) (err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ip-sm-gw", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)	
	
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
			err = fmt.Errorf("%d is unknown to DeleteIpSmGwContext", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for DeleteIpSmGwContext
func OnDeleteIpSmGwContext(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IPSMGWRegistrationDocumentApiHandler)
	
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

	successCode, apierr = prod.DR_HandleDeleteIpSmGwContext(ueId)

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
@return 
*/
func ModifyIpSmGwContext(client sbi.ConsumerClient, ueId string, body []models.PatchItem) (err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPatch

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ip-sm-gw", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json-patch+json"
	req.HeaderParams["Accept"] = "application/problem+json"
	//send the request
	var resp *sbi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	if resp.StatusCode >= 300 {
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.Body != nil {
			if err = client.DecodeResponse(resp); err == nil {
				err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
			}
			return
		} else {
			err = fmt.Errorf("%d is unknown to ModifyIpSmGwContext", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for ModifyIpSmGwContext
func OnModifyIpSmGwContext(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IPSMGWRegistrationDocumentApiHandler)
	
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

	var input []models.PatchItem

	var apierr *sbi.ApiError
	var successCode int32
	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, apierr = prod.DR_HandleModifyIpSmGwContext(ueId, input)
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




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId UE id
@param fields attributes to be retrieved
@param supportedFeatures Supported Features
@return *models.IpSmGwRegistration, 
*/
func QueryIpSmGwContext(client sbi.ConsumerClient, ueId string, fields []string, supportedFeatures string) (result models.IpSmGwRegistration, err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}		
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/ip-sm-gw", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	fieldsStr := utils.Param2String(fields)
	if len(fieldsStr) > 0 {
		req.QueryParams.Add("fields", fieldsStr)
	}
	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}	
	req.HeaderParams["Accept"] = "application/json"
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
			err = fmt.Errorf("%d is unknown to QueryIpSmGwContext", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for QueryIpSmGwContext
func OnQueryIpSmGwContext(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IPSMGWRegistrationDocumentApiHandler)
	
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
	
	supportedFeatures := ctx.Param("supported-features")

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.IpSmGwRegistration


	successCode, result, apierr = prod.DR_HandleQueryIpSmGwContext(ueId, fields, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type IPSMGWRegistrationDocumentApiHandler interface {
	DR_HandleCreateIpSmGwContext(ueId string, body models.IpSmGwRegistration) (successCode int32, err *sbi.ApiError)
	DR_HandleDeleteIpSmGwContext(ueId string) (successCode int32, err *sbi.ApiError)
	DR_HandleModifyIpSmGwContext(ueId string, body []models.PatchItem) (successCode int32, err *sbi.ApiError)
	DR_HandleQueryIpSmGwContext(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.IpSmGwRegistration, err *sbi.ApiError)
}

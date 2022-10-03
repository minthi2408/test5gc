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
@return *models.SmsfRegistration, 
*/
func CreateSmsfContextNon3gpp(client sbi.ConsumerClient, ueId string, body models.SmsfRegistration) (result models.SmsfRegistration, err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/smsf-non-3gpp-access", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
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
			err = fmt.Errorf("%d is unknown to CreateSmsfContextNon3gpp", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for CreateSmsfContextNon3gpp
func OnCreateSmsfContextNon3gpp(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(SMSFNon3GPPRegistrationDocumentApiHandler)
	
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
		successCode, result, apierr = prod.DR_HandleCreateSmsfContextNon3gpp(ueId, input)
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
@return 
*/
func DeleteSmsfContextNon3gpp(client sbi.ConsumerClient, ueId string) (err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/smsf-non-3gpp-access", ServicePath)
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
			err = fmt.Errorf("%d is unknown to DeleteSmsfContextNon3gpp", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for DeleteSmsfContextNon3gpp
func OnDeleteSmsfContextNon3gpp(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(SMSFNon3GPPRegistrationDocumentApiHandler)
	
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

	successCode, apierr = prod.DR_HandleDeleteSmsfContextNon3gpp(ueId)

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
@return *models.SmsfRegistration, 
*/
func QuerySmsfContextNon3gpp(client sbi.ConsumerClient, ueId string, fields []string, supportedFeatures string) (result models.SmsfRegistration, err error) {
	
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}		
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/context-data/smsf-non-3gpp-access", ServicePath)
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
			err = fmt.Errorf("%d is unknown to QuerySmsfContextNon3gpp", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for QuerySmsfContextNon3gpp
func OnQuerySmsfContextNon3gpp(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(SMSFNon3GPPRegistrationDocumentApiHandler)
	
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
	var result models.SmsfRegistration


	successCode, result, apierr = prod.DR_HandleQuerySmsfContextNon3gpp(ueId, fields, supportedFeatures)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type SMSFNon3GPPRegistrationDocumentApiHandler interface {
	DR_HandleCreateSmsfContextNon3gpp(ueId string, body models.SmsfRegistration) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError)
	DR_HandleDeleteSmsfContextNon3gpp(ueId string) (successCode int32, err *sbi.ApiError)
	DR_HandleQuerySmsfContextNon3gpp(ueId string, fields []string, supportedFeatures string) (successCode int32, result models.SmsfRegistration, err *sbi.ApiError)
}

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
@param appId Indicate the application identifier for the request pfd(s). It shall apply the format of Data type ApplicationId.
@return *models.PfdDataForAppExt, 
*/
func CreateOrReplaceIndividualPFDData(client sbi.ConsumerClient, appId string, body models.PfdDataForAppExt) (result models.PfdDataForAppExt, err error) {
	
	if len(appId) == 0 {
		err = fmt.Errorf("appId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/application-data/pfds/{appId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"appId"+"}", url.PathEscape(appId), -1)
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
		if resp.StatusCode == 414 {
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
			err = fmt.Errorf("%d is unknown to CreateOrReplaceIndividualPFDData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for CreateOrReplaceIndividualPFDData
func OnCreateOrReplaceIndividualPFDData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualPFDDataDocumentApiHandler)
	
	appId := ctx.Param("appId")
	if len(appId) == 0 {
		//appId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "appId is required",
		}))
		return
	}

	var input models.PfdDataForAppExt

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PfdDataForAppExt

	if prob := ctx.DecodeRequest(&input); prob == nil {
		successCode, result, apierr = prod.DR_HandleCreateOrReplaceIndividualPFDData(appId, input)
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
@param appId Indicate the application identifier for the request pfd(s). It shall apply the format of Data type ApplicationId.
@return 
*/
func DeleteIndividualPFDData(client sbi.ConsumerClient, appId string) (err error) {
	
	if len(appId) == 0 {
		err = fmt.Errorf("appId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/application-data/pfds/{appId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"appId"+"}", url.PathEscape(appId), -1)	
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
			err = fmt.Errorf("%d is unknown to DeleteIndividualPFDData", resp.StatusCode)
			return
		}
	}

	return 
}


//sbi producer handler for DeleteIndividualPFDData
func OnDeleteIndividualPFDData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualPFDDataDocumentApiHandler)
	
	appId := ctx.Param("appId")
	if len(appId) == 0 {
		//appId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "appId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32

	successCode, apierr = prod.DR_HandleDeleteIndividualPFDData(appId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), nil)
	}
	return
}




/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param appId Indicate the application identifier for the request pfd(s). It shall apply the format of Data type ApplicationId.
@return *models.PfdDataForAppExt, 
*/
func ReadIndividualPFDData(client sbi.ConsumerClient, appId string) (result models.PfdDataForAppExt, err error) {
	
	if len(appId) == 0 {
		err = fmt.Errorf("appId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/application-data/pfds/{appId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"appId"+"}", url.PathEscape(appId), -1)	
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
			err = fmt.Errorf("%d is unknown to ReadIndividualPFDData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for ReadIndividualPFDData
func OnReadIndividualPFDData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(IndividualPFDDataDocumentApiHandler)
	
	appId := ctx.Param("appId")
	if len(appId) == 0 {
		//appId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "appId is required",
		}))
		return
	}

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.PfdDataForAppExt


	successCode, result, apierr = prod.DR_HandleReadIndividualPFDData(appId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type IndividualPFDDataDocumentApiHandler interface {
	DR_HandleCreateOrReplaceIndividualPFDData(appId string, body models.PfdDataForAppExt) (successCode int32, result models.PfdDataForAppExt, err *sbi.ApiError)
	DR_HandleDeleteIndividualPFDData(appId string) (successCode int32, err *sbi.ApiError)
	DR_HandleReadIndividualPFDData(appId string) (successCode int32, result models.PfdDataForAppExt, err *sbi.ApiError)
}
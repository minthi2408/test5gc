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
	"etri5gc/sbi/utils"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId UE id
@param servingPlmnId PLMN ID
@param singleNssai single NSSAI
@param dnn DNN
@param fields attributes to be retrieved
@param supportedFeatures Supported Features
@param ifNoneMatch Validator for conditional requests, as described in RFC 7232, 3.2
@param ifModifiedSince Validator for conditional requests, as described in RFC 7232, 3.3
@return []models.SessionManagementSubscriptionData,
*/
func QuerySmData(client sbi.ConsumerClient, ueId string, servingPlmnId string, singleNssai *models.Snssai, dnn string, fields []string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (result []models.SessionManagementSubscriptionData, err error) {

	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if len(servingPlmnId) == 0 {
		err = fmt.Errorf("servingPlmnId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/{servingPlmnId}/provisioned-data/sm-data", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"servingPlmnId"+"}", url.PathEscape(servingPlmnId), -1)
	singleNssaiStr := utils.Param2String(singleNssai)
	if len(singleNssaiStr) > 0 {
		req.QueryParams.Add("single-nssai", singleNssaiStr)
	}
	if len(dnn) > 0 {
		req.QueryParams.Add("dnn", dnn)
	}
	fieldsStr := utils.Param2String(fields)
	if len(fieldsStr) > 0 {
		req.QueryParams.Add("fields", fieldsStr)
	}
	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}
	if len(ifNoneMatch) > 0 {
		req.HeaderParams["If-None-Match"] = ifNoneMatch
	}
	if len(ifModifiedSince) > 0 {
		req.HeaderParams["If-Modified-Since"] = ifModifiedSince
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
			err = fmt.Errorf("%d is unknown to QuerySmData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

//sbi producer handler for QuerySmData
func OnQuerySmData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(SessionManagementSubscriptionDataApiHandler)

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
	servingPlmnId := ctx.Param("servingPlmnId")
	if len(servingPlmnId) == 0 {
		//servingPlmnId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "servingPlmnId is required",
		}))
		return
	}
	singleNssaiStr := ctx.Param("single-nssai")
	var singleNssai *models.Snssai
	var singleNssaiErr error
	if singleNssai, singleNssaiErr = utils.String2Snssai(singleNssaiStr); singleNssaiErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: singleNssaiErr.Error(),
		}))
		return
	}

	dnn := ctx.Param("dnn")
	fieldsStr := ctx.Param("fields")
	var fields []string
	var fieldsErr error
	if fields, fieldsErr = utils.String2ArrayOfstring(fieldsStr); fieldsErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: fieldsErr.Error(),
		}))
		return
	}

	supportedFeatures := ctx.Param("supported-features")
	ifNoneMatch := ctx.Param("If-None-Match")
	ifModifiedSince := ctx.Param("If-Modified-Since")

	var apierr *sbi.ApiError
	var successCode int32
	var result []models.SessionManagementSubscriptionData

	successCode, result, apierr = prod.DR_HandleQuerySmData(ueId, servingPlmnId, singleNssai, dnn, fields, supportedFeatures, ifNoneMatch, ifModifiedSince)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

type SessionManagementSubscriptionDataApiHandler interface {
	DR_HandleQuerySmData(ueId string, servingPlmnId string, singleNssai *models.Snssai, dnn string, fields []string, supportedFeatures string, ifNoneMatch string, ifModifiedSince string) (successCode int32, result []models.SessionManagementSubscriptionData, err *sbi.ApiError)
}

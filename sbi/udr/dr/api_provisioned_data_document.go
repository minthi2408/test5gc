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
@param servingPlmnId PLMN ID
@param datasetNames List of dataset names
@return *models.ProvisionedDataSets, 
*/
func QueryProvisionedData(client sbi.ConsumerClient, ueId string, servingPlmnId string, datasetNames []models.DataSetName) (result models.ProvisionedDataSets, err error) {
	
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

	req.Path = fmt.Sprintf("%s/subscription-data/{ueId}/{servingPlmnId}/provisioned-data", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueId), -1)
	req.Path = strings.Replace(req.Path, "{"+"servingPlmnId"+"}", url.PathEscape(servingPlmnId), -1)
	datasetNamesStr := utils.Param2String(datasetNames)
	if len(datasetNamesStr) > 0 {
		req.QueryParams.Add("dataset-names", datasetNamesStr)
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
			err = fmt.Errorf("%d is unknown to QueryProvisionedData", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}


//sbi producer handler for QueryProvisionedData
func OnQueryProvisionedData(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(ProvisionedDataDocumentApiHandler)
	
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
	servingPlmnId := ctx.Param("servingPlmnId")
	if len(servingPlmnId) == 0 {
		//servingPlmnId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: "servingPlmnId is required",
		}))
		return
	}
	datasetNamesStr := ctx.Param("dataset-names")
	var datasetNames []models.DataSetName
	var datasetNamesErr error
	if datasetNames, datasetNamesErr = utils.String2ArrayOfDataSetName(datasetNamesStr); datasetNamesErr != nil {
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title: "Bad request",
			Status: http.StatusBadRequest,
			Detail: datasetNamesErr.Error(), 
		}))
		return
	}
	

	

	var apierr *sbi.ApiError
	var successCode int32
	var result models.ProvisionedDataSets


	successCode, result, apierr = prod.DR_HandleQueryProvisionedData(ueId, servingPlmnId, datasetNames)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}




type ProvisionedDataDocumentApiHandler interface {
	DR_HandleQueryProvisionedData(ueId string, servingPlmnId string, datasetNames []models.DataSetName) (successCode int32, result models.ProvisionedDataSets, err *sbi.ApiError)
}

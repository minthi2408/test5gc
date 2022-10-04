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
@param externalGroupId
@return *models.Model5GVnGroupConfiguration,
*/
func Get5GVnGroupConfiguration(client sbi.ConsumerClient, externalGroupId string) (result models.Model5GVnGroupConfiguration, err error) {

	if len(externalGroupId) == 0 {
		err = fmt.Errorf("externalGroupId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/subscription-data/group-data/5g-vn-groups/{externalGroupId}", ServicePath)
	req.Path = strings.Replace(req.Path, "{"+"externalGroupId"+"}", url.PathEscape(externalGroupId), -1)
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
			err = fmt.Errorf("%d is unknown to Get5GVnGroupConfiguration", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)
	}
	return
}

//sbi producer handler for Get5GVnGroupConfiguration
func OnGet5GVnGroupConfiguration(ctx sbi.RequestContext, handler interface{}) (resp sbi.Response) {
	prod := handler.(Query5GVnGroupConfigurationDocumentApiHandler)

	externalGroupId := ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		//externalGroupId is required
		resp.SetApiError(sbi.ApiErrFromProb(&models.ProblemDetails{
			Title:  "Bad request",
			Status: http.StatusBadRequest,
			Detail: "externalGroupId is required",
		}))
		return
	}

	var apierr *sbi.ApiError
	var successCode int32
	var result models.Model5GVnGroupConfiguration

	successCode, result, apierr = prod.DR_HandleGet5GVnGroupConfiguration(externalGroupId)

	if apierr != nil {
		resp.SetApiError(apierr)
	} else {
		resp.SetBody(int(successCode), &result)
	}
	return
}

type Query5GVnGroupConfigurationDocumentApiHandler interface {
	DR_HandleGet5GVnGroupConfiguration(externalGroupId string) (successCode int32, result models.Model5GVnGroupConfiguration, err *sbi.ApiError)
}

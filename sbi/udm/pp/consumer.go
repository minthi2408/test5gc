/*
Nudm_PP

Nudm Parameter Provision Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// Templates and customized generator are developed by Quang Tung Thai (tqtung@etri.re.kr)

package pp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"etri5gc/sbi"
	"etri5gc/sbi/models"
	"etri5gc/sbi/utils"
)


const (
 SERVICE_PATH = "{apiRoot}/nudm-pp/v1"
)



/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param extGroupId External Identifier of the Group
@return 
*/
func Create5GVNGroup(client sbi.ConsumerClient, extGroupId string, body models.Model5GVnGroupConfiguration) (err error) {
	
	if len(extGroupId) == 0 {
		err = fmt.Errorf("extGroupId is required")
		return
	}	
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPut

	req.Path = fmt.Sprintf("%s/5g-vn-groups/{extGroupId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"extGroupId"+"}", url.PathEscape(extGroupId), -1)
	req.Body = &body
	req.HeaderParams["Content-Type"] = "application/json"
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
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
			err = fmt.Errorf("%d is unknown to Create5GVNGroup", resp.StatusCode)
			return
		}
	}

	return 
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param extGroupId External Identifier of the Group
@param mtcProviderInfo MTC Provider Information that originated the service operation
@param afId AF ID that originated the service operation
@return 
*/
func Delete5GVNGroup(client sbi.ConsumerClient, extGroupId string, mtcProviderInfo string, afId string) (err error) {
	
	if len(extGroupId) == 0 {
		err = fmt.Errorf("extGroupId is required")
		return
	}		
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodDelete

	req.Path = fmt.Sprintf("%s/5g-vn-groups/{extGroupId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"extGroupId"+"}", url.PathEscape(extGroupId), -1)
	if len(mtcProviderInfo) > 0 {
		req.QueryParams.Add("mtc-provider-info", mtcProviderInfo)
	}
	if len(afId) > 0 {
		req.QueryParams.Add("af-id", afId)
	}	
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
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
			err = fmt.Errorf("%d is unknown to Delete5GVNGroup", resp.StatusCode)
			return
		}
	}

	return 
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param extGroupId External Identifier of the group
@return *models.Model5GVnGroupConfiguration, 
*/
func Get5GVNGroup(client sbi.ConsumerClient, extGroupId string) (result models.Model5GVnGroupConfiguration, err error) {
	
	if len(extGroupId) == 0 {
		err = fmt.Errorf("extGroupId is required")
		return
	}
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodGet

	req.Path = fmt.Sprintf("%s/5g-vn-groups/{extGroupId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"extGroupId"+"}", url.PathEscape(extGroupId), -1)	
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
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
			err = fmt.Errorf("%d is unknown to Get5GVNGroup", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param extGroupId External Identifier of the group
@param supportedFeatures Features required to be supported by the target NF
@return *models.PatchResult, 
*/
func Modify5GVNGroup(client sbi.ConsumerClient, extGroupId string, supportedFeatures string, body models.Model5GVnGroupConfiguration) (result models.PatchResult, err error) {
	
	if len(extGroupId) == 0 {
		err = fmt.Errorf("extGroupId is required")
		return
	}		
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPatch

	req.Path = fmt.Sprintf("%s/5g-vn-groups/{extGroupId}", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"extGroupId"+"}", url.PathEscape(extGroupId), -1)
	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
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
			err = fmt.Errorf("%d is unknown to Modify5GVNGroup", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}

/*
@param client sbi.ConsumerClient - for encoding request/encoding response and sending request to remote agent.
@param ueId Identifier of the UE
@param supportedFeatures Features required to be supported by the target NF
@return *models.PatchResult, 
*/
func Update(client sbi.ConsumerClient, ueId models.UpdateUeIdParameter, supportedFeatures string, body models.PpData) (result models.PatchResult, err error) {
	
	ueIdStr := utils.Param2String(ueId)
	if len(ueIdStr) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}		
	//create a request
	req := sbi.DefaultRequest()
	req.Method = http.MethodPatch

	req.Path = fmt.Sprintf("%s/{ueId}/pp-data", SERVICE_PATH)
	req.Path = strings.Replace(req.Path, "{"+"ueId"+"}", url.PathEscape(ueIdStr), -1)
	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}
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
		if resp.StatusCode == 403 {
			resp.Body = &models.ProblemDetails{}
		}
		if resp.StatusCode == 404 {
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
			err = fmt.Errorf("%d is unknown to Update", resp.StatusCode)
			return
		}
	}

	resp.Body = &result
	if err = client.DecodeResponse(resp); err == nil {
		err = sbi.NewApiError(resp.StatusCode, resp.Status, resp.Body)	
	}
	return 
}



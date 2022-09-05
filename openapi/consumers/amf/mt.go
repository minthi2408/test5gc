package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

/*
UeContextDocumentApiService Namf_MT Provide Domain Selection Info service Operation
 * @param client openapi.ConsumerClient an abstraction of a connection in the
 * dataplane of the signaling network
 * @param ueContextId UE Context Identifier
 * @param optional nil or *ProvideDomainSelectionInfoParamOpts - Optional Parameters:
 * @param "InfoClass" (optional.Interface of models.UeContextInfoClass) -  UE Context Information Class
 * @param "SupportedFeatures" (optional.String) -  Supported Features
@return models.UeContextInfo
*/


//TODO: review this client API
type ProvideDomainSelectionInfoParamOpts struct {
	//InfoClass         optional.Interface
	//SupportedFeatures optional.String
}

func  ProvideDomainSelectionInfo(client openapi.ConsumerClient, ueContextId string, params *ProvideDomainSelectionInfoParamOpts) (result models.UeContextInfo, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.AMF_MT, openapi.AMF_MT_UE_CONTEXTS, ueContextId)
/*
	if params != nil && params.InfoClass.IsSet() {
		req.QueryParams["info-class"] = openapi.ParameterToString(params.InfoClass.Value(), ""))
	}
	if params != nil && params.SupportedFeatures.IsSet() && params.SupportedFeatures.Value() != "" {
		req.QueryParams["supported-features"], openapi.ParameterToString(params.SupportedFeatures.Value(), ""))
	}
*/
	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

    switch resp.StatusCode {
	case 200:
        resp.Body = &result
        err = client.DecodeResponse(resp)
	case 307:
        fallthrough
	case 400:
        fallthrough
	case 403:
        fallthrough
	case 404:
        fallthrough
	case 414:
        fallthrough
	case 429:
        fallthrough
	case 500:
        fallthrough
	case 503:
		var prob models.ProblemDetails
        resp.Body = &prob
        if err = client.DecodeResponse(resp); err == nil {
            err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
        }
	default:
        //NOTE: strange
	}
    return
}
/*
UeReachIndDocumentApiService Namf_MT EnableUEReachability service Operation
 * @param client openapi.ConsumerClient an abstraction of a connection in the
 * dataplane of the signaling network
 * @param ueContextId UE Context Identifier
 * @param enableUeReachabilityReqData
@return models.EnableUeReachabilityRspData
*/

func EnableUeReachability(client openapi.ConsumerClient, ueContextId string, body models.EnableUeReachabilityReqData) (result models.EnableUeReachabilityRspData, err error) {
    //create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s/ue-reachind", openapi.AMF_MT, openapi.AMF_MT_UE_CONTEXTS, ueContextId)

    //send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

    switch resp.StatusCode {
	case 200:
        resp.Body = &body
        err = client.DecodeResponse(resp)
	case 307:
        fallthrough
	case 400:
        fallthrough
	case 403:
        fallthrough
	case 404:
        fallthrough
	case 411:
        fallthrough
	case 413:
        fallthrough
	case 415:
        fallthrough
	case 429:
        fallthrough
	case 500:
        fallthrough
	case 503:
        fallthrough
	case 504:
		var prob models.ProblemDetails
        resp.Body = prob
        if err = client.DecodeResponse(resp); err == nil {
            err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
        }
	default:
        //NOTE: strange
	}
    return
}

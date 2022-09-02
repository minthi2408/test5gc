package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

/*
IndividualUEContextDocumentApiService Namf_Location ProvideLocationInfo service Operation
 * @param client openapi.ConsumerClient an abstraction of a connection in the
 * dataplane of the signaling network
 * @param ueContextId UE Context Identifier
 * @param requestLocInfo
@return models.ProvideLocInfo
*/

func ProvideLocationInfo(client openapi.ConsumerClient, ueContextId string, body models.RequestLocInfo) (result models.ProvideLocInfo, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/provide-loc-info", openapi.AMF_LOCATION, ueContextId)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}


	switch resp.StatusCode {
	case 200:
        resp.Body = &result
        err =  client.DecodeResponse(resp)
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
IndividualUEContextDocumentApiService Namf_Location ProvidePositioningInfo service Operation
 * @param client openapi.ConsumerClient an abstraction of a connection in the
 * dataplane of the signaling network
 * @param ueContextId UE Context Identifier
 * @param requestPosInfo
@return models.ProvidePosInfo
*/

func ProvidePositioningInfo(client openapi.ConsumerClient, ueContextId string, body models.RequestPosInfo) (result models.ProvidePosInfo, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/provide-pos-info", openapi.AMF_LOCATION, ueContextId)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}


	switch resp.StatusCode {
	case 200:
        resp.Body = &result
        err = client.DecodeResponse(resp)
	case 400:
        fallthrough
	case 403:
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
        resp.Body = &prob
        if err = client.DecodeResponse(resp); err == nil {
            err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
        }
	default:
        //NOTE: strange
	}
    return
}

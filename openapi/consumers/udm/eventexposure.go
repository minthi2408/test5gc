package udm

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

func CreateEeSubscription(client openapi.ConsumerClient, ueId string, eesub models.EeSubscription) (result models.CreatedEeSubscription, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/%s/%s", openapi.UDM_EVENTEXPOSURE, ueId, openapi.UDM_EE_SUBSCRIPTIONS)
	req.Body = &eesub

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch resp.StatusCode {
	case 201:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 400, 403, 404, 500, 501, 503:
		var prob models.ProblemDetails
		resp.Body = &prob
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

func DeleteEeSubscription(client openapi.ConsumerClient, ueId string, subId string) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "DELETE"
	req.Path = fmt.Sprintf("%s/%s/%s/%s", openapi.UDM_EVENTEXPOSURE, ueId, openapi.UDM_EE_SUBSCRIPTIONS, subId)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch resp.StatusCode {
	case 204:
		//good, do nothing
	case 400, 404, 500, 503:
		var prob models.ProblemDetails
		resp.Body = &prob
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

func UpdateEeSubscription(client openapi.ConsumerClient, ueId string, subId string, items []models.PatchItem) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PATCH"
	req.Path = fmt.Sprintf("%s/%s/%s/%s", openapi.UDM_EVENTEXPOSURE, ueId, openapi.UDM_EE_SUBSCRIPTIONS, subId)
	req.HeaderParams["Content-Type"] = "application/json-patch+json"
	req.Body = &items

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response

	switch resp.StatusCode {
	case 204:
		//good, do nothing
	case 403, 404:
		var prob models.ProblemDetails
		resp.Body = &prob
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
		}
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

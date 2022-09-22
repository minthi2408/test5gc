package udm

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

func UecmGet(client openapi.ConsumerClient, ueId string, supportedFeatures string) (result models.Amf3GppAccessRegistration, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s/%s", openapi.UDM_UECMAN, ueId, openapi.UDM_UECMAN_REGISTRATIONS, openapi.UDM_UECMAN_AMF_3GPP_ACCESS)

	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 200:
		resp.Body = &result
		err = client.DecodeResponse(resp)

	case 400, 403, 404, 500, 503:
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

func UecmGetAmfNon3gppAccess(client openapi.ConsumerClient, ueId string, supportedFeatures string) (result models.AmfNon3GppAccessRegistration, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "GET"
	req.Path = fmt.Sprintf("%s/%s/%s/%s", openapi.UDM_UECMAN, ueId, openapi.UDM_UECMAN_REGISTRATIONS, openapi.UDM_UECMAN_AMF_NON3GPP_ACCESS)

	if len(supportedFeatures) > 0 {
		req.QueryParams.Add("supported-features", supportedFeatures)
	}

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 200:
		resp.Body = &result
		err = client.DecodeResponse(resp)

	case 400, 403, 404, 500, 503:
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

func UecmRegistration(client openapi.ConsumerClient, ueId string, input models.Amf3GppAccessRegistration) (result models.Amf3GppAccessRegistration, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PUT"
	req.Path = fmt.Sprintf("%s/%s/%s/%s", openapi.UDM_UECMAN, ueId, openapi.UDM_UECMAN_REGISTRATIONS, openapi.UDM_UECMAN_AMF_3GPP_ACCESS)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 200, 201:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 204:
		//Good, do nothing
	case 400, 403, 404, 500, 503:
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

func UecmRegister(client openapi.ConsumerClient, ueId string, input models.AmfNon3GppAccessRegistration) (result models.AmfNon3GppAccessRegistration, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PUT"
	req.Path = fmt.Sprintf("%s/%s/%s/%s", openapi.UDM_UECMAN, ueId, openapi.UDM_UECMAN_REGISTRATIONS, openapi.UDM_UECMAN_AMF_NON3GPP_ACCESS)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 200, 201:
		resp.Body = &result
		err = client.DecodeResponse(resp)
	case 204:
		//Good, do nothing
	case 400, 403, 404, 500, 503:
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

func UecmDeregistrationNotify(client openapi.ConsumerClient, cliUrl string, request models.DeregistrationData) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	//TODO: set the client path as cliUrl

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}

	//handle the response
	switch resp.StatusCode {
	case 204:
		//Good, do nothing
	case 400, 404:
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

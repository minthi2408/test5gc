package amf

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"fmt"
)

/*
SubscriptionsCollectionDocumentApiService Namf_EventExposure Subscribe service Operation
 * @param client openapi.ConsumerClient an abstraction of a connection in the
 * dataplane of the signaling network
 * @param amfCreateEventSubscription
@return AmfCreatedEventSubscription
*/

func CreateSubscription(client openapi.ConsumerClient, body models.AmfCreateEventSubscription) (result models.AmfCreatedEventSubscription, err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/subscriptions", openapi.AMF_EVENTEXPOSURE)
	req.Body = &body

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}
	switch resp.StatusCode {
	case 201:
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
		var prob models.ProblemDetails
		resp.Body = &prob
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &body)
		}
	case 0:
		//do nothing
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

/*
IndividualSubscriptionDocumentApiService Namf_EventExposure Unsubscribe service Operation
 * @param client openapi.ConsumerClient an abstraction of a connection in the
 * dataplane of the signaling network
 * @param subscriptionId Unique ID of the subscription to be deleted
*/

func DeleteSubscription(client openapi.ConsumerClient, subscriptionId string) (err error) {
	//create a request
	req := openapi.DefaultRequest()
	req.Method = "POST"
	req.Path = fmt.Sprintf("%s/subscriptions/%s", openapi.AMF_EVENTEXPOSURE, subscriptionId)

	//send the request
	var resp *openapi.Response
	if resp, err = client.Send(req); err != nil {
		return
	}
	switch resp.StatusCode {
	case 200:
		//do nothing
	case 400:
	case 404:
	case 411:
	case 413:
	case 415:
	case 429:
	case 500:
	case 503:
		var prob models.ProblemDetails
		resp.Body = &prob
		if err = client.DecodeResponse(resp); err == nil {
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &prob)
		}
	case 0:
		// do nothing
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

/*
IndividualSubscriptionDocumentApiService Namf_EventExposure Subscribe Modify service Operation
 * @param client openapi.ConsumerClient an abstraction of a connection in the
 * dataplane of the signaling network
 * @param subscriptionId Unique ID of the subscription to be modified
 * @param modifySubscriptionRequest
@return AmfUpdatedEventSubscription
*/

func ModifySubscription(client openapi.ConsumerClient, subscriptionId string, body interface{}) (result models.AmfUpdatedEventSubscription, err error) {

	//create a request
	req := openapi.DefaultRequest()
	req.Method = "PATCH"
	req.HeaderParams["Accept"] = "application/problem+json"
	req.Path = fmt.Sprintf("%s/subscriptions/%s", openapi.AMF_EVENTEXPOSURE, subscriptionId)
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
			err = openapi.NewApiError(resp.StatusCode, resp.Status, &body)
		}
	case 0:
		//do nothing
	default:
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	return
}

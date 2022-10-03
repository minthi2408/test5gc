# \IndividualInfluenceDataSubscriptionDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualInfluenceDataSubscription**](IndividualInfluenceDataSubscriptionDocumentApi.md#DeleteIndividualInfluenceDataSubscription) | **Delete** /application-data/influenceData/subs-to-notify/{subscriptionId} | Delete an individual Influence Data Subscription resource
[**ReadIndividualInfluenceDataSubscription**](IndividualInfluenceDataSubscriptionDocumentApi.md#ReadIndividualInfluenceDataSubscription) | **Get** /application-data/influenceData/subs-to-notify/{subscriptionId} | Get an existing individual Influence Data Subscription resource
[**ReplaceIndividualInfluenceDataSubscription**](IndividualInfluenceDataSubscriptionDocumentApi.md#ReplaceIndividualInfluenceDataSubscription) | **Put** /application-data/influenceData/subs-to-notify/{subscriptionId} | Modify an existing individual Influence Data Subscription resource



## DeleteIndividualInfluenceDataSubscription

> DeleteIndividualInfluenceDataSubscription(ctx, subscriptionId).Execute()

Delete an individual Influence Data Subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Individual Influence Data Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualInfluenceDataSubscriptionDocumentApi.DeleteIndividualInfluenceDataSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualInfluenceDataSubscriptionDocumentApi.DeleteIndividualInfluenceDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Individual Influence Data Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualInfluenceDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIndividualInfluenceDataSubscription

> TrafficInfluSub ReadIndividualInfluenceDataSubscription(ctx, subscriptionId).Execute()

Get an existing individual Influence Data Subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Individual Influence Data Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualInfluenceDataSubscriptionDocumentApi.ReadIndividualInfluenceDataSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualInfluenceDataSubscriptionDocumentApi.ReadIndividualInfluenceDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualInfluenceDataSubscription`: TrafficInfluSub
    fmt.Fprintf(os.Stdout, "Response from `IndividualInfluenceDataSubscriptionDocumentApi.ReadIndividualInfluenceDataSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Individual Influence Data Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualInfluenceDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrafficInfluSub**](TrafficInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceIndividualInfluenceDataSubscription

> TrafficInfluSub ReplaceIndividualInfluenceDataSubscription(ctx, subscriptionId).TrafficInfluSub(trafficInfluSub).Execute()

Modify an existing individual Influence Data Subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying a subscription to the Individual Influence Data Subscription
    trafficInfluSub := TODO // TrafficInfluSub | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualInfluenceDataSubscriptionDocumentApi.ReplaceIndividualInfluenceDataSubscription(context.Background(), subscriptionId).TrafficInfluSub(trafficInfluSub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualInfluenceDataSubscriptionDocumentApi.ReplaceIndividualInfluenceDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIndividualInfluenceDataSubscription`: TrafficInfluSub
    fmt.Fprintf(os.Stdout, "Response from `IndividualInfluenceDataSubscriptionDocumentApi.ReplaceIndividualInfluenceDataSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a subscription to the Individual Influence Data Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIndividualInfluenceDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trafficInfluSub** | [**TrafficInfluSub**](TrafficInfluSub.md) |  | 

### Return type

[**TrafficInfluSub**](TrafficInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


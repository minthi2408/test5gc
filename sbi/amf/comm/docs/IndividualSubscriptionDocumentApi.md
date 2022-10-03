# \IndividualSubscriptionDocumentApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AMFStatusChangeSubscribeModfy**](IndividualSubscriptionDocumentApi.md#AMFStatusChangeSubscribeModfy) | **Put** /subscriptions/{subscriptionId} | Namf_Communication AMF Status Change Subscribe Modify service Operation
[**AMFStatusChangeUnSubscribe**](IndividualSubscriptionDocumentApi.md#AMFStatusChangeUnSubscribe) | **Delete** /subscriptions/{subscriptionId} | Namf_Communication AMF Status Change UnSubscribe service Operation



## AMFStatusChangeSubscribeModfy

> SubscriptionData AMFStatusChangeSubscribeModfy(ctx, subscriptionId).SubscriptionData(subscriptionData).Execute()

Namf_Communication AMF Status Change Subscribe Modify service Operation

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
    subscriptionId := "subscriptionId_example" // string | AMF Status Change Subscription Identifier
    subscriptionData := TODO // SubscriptionData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.AMFStatusChangeSubscribeModfy(context.Background(), subscriptionId).SubscriptionData(subscriptionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.AMFStatusChangeSubscribeModfy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AMFStatusChangeSubscribeModfy`: SubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionDocumentApi.AMFStatusChangeSubscribeModfy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | AMF Status Change Subscription Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAMFStatusChangeSubscribeModfyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionData** | [**SubscriptionData**](SubscriptionData.md) |  | 

### Return type

[**SubscriptionData**](SubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AMFStatusChangeUnSubscribe

> AMFStatusChangeUnSubscribe(ctx, subscriptionId).Execute()

Namf_Communication AMF Status Change UnSubscribe service Operation

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
    subscriptionId := "subscriptionId_example" // string | AMF Status Change Subscription Identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.AMFStatusChangeUnSubscribe(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.AMFStatusChangeUnSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | AMF Status Change Subscription Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAMFStatusChangeUnSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


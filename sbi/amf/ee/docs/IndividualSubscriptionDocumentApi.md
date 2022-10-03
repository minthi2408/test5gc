# \IndividualSubscriptionDocumentApi

All URIs are relative to *https://example.com/namf-evts/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubscription**](IndividualSubscriptionDocumentApi.md#DeleteSubscription) | **Delete** /subscriptions/{subscriptionId} | Namf_EventExposure Unsubscribe service Operation
[**ModifySubscription**](IndividualSubscriptionDocumentApi.md#ModifySubscription) | **Patch** /subscriptions/{subscriptionId} | Namf_EventExposure Subscribe Modify service Operation



## DeleteSubscription

> DeleteSubscription(ctx, subscriptionId).Execute()

Namf_EventExposure Unsubscribe service Operation

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the subscription to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.DeleteSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the subscription to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


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


## ModifySubscription

> AmfUpdatedEventSubscription ModifySubscription(ctx, subscriptionId).ModifySubscriptionRequest(modifySubscriptionRequest).Execute()

Namf_EventExposure Subscribe Modify service Operation

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the subscription to be modified
    modifySubscriptionRequest := TODO // ModifySubscriptionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.ModifySubscription(context.Background(), subscriptionId).ModifySubscriptionRequest(modifySubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.ModifySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifySubscription`: AmfUpdatedEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionDocumentApi.ModifySubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the subscription to be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifySubscriptionRequest** | [**ModifySubscriptionRequest**](ModifySubscriptionRequest.md) |  | 

### Return type

[**AmfUpdatedEventSubscription**](AmfUpdatedEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


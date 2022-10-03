# \IndividualExposureDataSubscriptionDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualExposureDataSubscription**](IndividualExposureDataSubscriptionDocumentApi.md#DeleteIndividualExposureDataSubscription) | **Delete** /exposure-data/subs-to-notify/{subId} | Deletes the individual Exposure Data subscription
[**ReplaceIndividualExposureDataSubscription**](IndividualExposureDataSubscriptionDocumentApi.md#ReplaceIndividualExposureDataSubscription) | **Put** /exposure-data/subs-to-notify/{subId} | updates a subscription to receive notifications of exposure data changes



## DeleteIndividualExposureDataSubscription

> DeleteIndividualExposureDataSubscription(ctx, subId).Execute()

Deletes the individual Exposure Data subscription

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
    subId := "subId_example" // string | Subscription id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualExposureDataSubscriptionDocumentApi.DeleteIndividualExposureDataSubscription(context.Background(), subId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualExposureDataSubscriptionDocumentApi.DeleteIndividualExposureDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | Subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualExposureDataSubscriptionRequest struct via the builder pattern


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


## ReplaceIndividualExposureDataSubscription

> ExposureDataSubscription ReplaceIndividualExposureDataSubscription(ctx, subId).ExposureDataSubscription(exposureDataSubscription).Execute()

updates a subscription to receive notifications of exposure data changes

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
    subId := "subId_example" // string | Subscription id
    exposureDataSubscription := TODO // ExposureDataSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualExposureDataSubscriptionDocumentApi.ReplaceIndividualExposureDataSubscription(context.Background(), subId).ExposureDataSubscription(exposureDataSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualExposureDataSubscriptionDocumentApi.ReplaceIndividualExposureDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIndividualExposureDataSubscription`: ExposureDataSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualExposureDataSubscriptionDocumentApi.ReplaceIndividualExposureDataSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | Subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIndividualExposureDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exposureDataSubscription** | [**ExposureDataSubscription**](ExposureDataSubscription.md) |  | 

### Return type

[**ExposureDataSubscription**](ExposureDataSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


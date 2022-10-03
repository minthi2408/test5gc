# \EventExposureSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEeSubscriptions**](EventExposureSubscriptionsCollectionApi.md#CreateEeSubscriptions) | **Post** /subscription-data/{ueId}/context-data/ee-subscriptions | Create individual EE subscription
[**Queryeesubscriptions**](EventExposureSubscriptionsCollectionApi.md#Queryeesubscriptions) | **Get** /subscription-data/{ueId}/context-data/ee-subscriptions | Retrieves the ee subscriptions of a UE



## CreateEeSubscriptions

> EeSubscription CreateEeSubscriptions(ctx, ueId).EeSubscription(eeSubscription).Execute()

Create individual EE subscription

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
    ueId := "ueId_example" // string | UE ID
    eeSubscription := TODO // EeSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExposureSubscriptionsCollectionApi.CreateEeSubscriptions(context.Background(), ueId).EeSubscription(eeSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExposureSubscriptionsCollectionApi.CreateEeSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEeSubscriptions`: EeSubscription
    fmt.Fprintf(os.Stdout, "Response from `EventExposureSubscriptionsCollectionApi.CreateEeSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEeSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eeSubscription** | [**EeSubscription**](EeSubscription.md) |  | 

### Return type

[**EeSubscription**](EeSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Queryeesubscriptions

> []EeSubscription Queryeesubscriptions(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

Retrieves the ee subscriptions of a UE

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
    ueId := "ueId_example" // string | UE id
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExposureSubscriptionsCollectionApi.Queryeesubscriptions(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExposureSubscriptionsCollectionApi.Queryeesubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Queryeesubscriptions`: []EeSubscription
    fmt.Fprintf(os.Stdout, "Response from `EventExposureSubscriptionsCollectionApi.Queryeesubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryeesubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**[]EeSubscription**](EeSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


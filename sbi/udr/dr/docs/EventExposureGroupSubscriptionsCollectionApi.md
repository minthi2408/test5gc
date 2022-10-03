# \EventExposureGroupSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEeGroupSubscriptions**](EventExposureGroupSubscriptionsCollectionApi.md#CreateEeGroupSubscriptions) | **Post** /subscription-data/group-data/{ueGroupId}/ee-subscriptions | Create individual EE subscription for a group of UEs or any UE
[**QueryEeGroupSubscriptions**](EventExposureGroupSubscriptionsCollectionApi.md#QueryEeGroupSubscriptions) | **Get** /subscription-data/group-data/{ueGroupId}/ee-subscriptions | Retrieves the ee subscriptions of a group of UEs or any UE



## CreateEeGroupSubscriptions

> EeSubscription CreateEeGroupSubscriptions(ctx, ueGroupId).EeSubscription(eeSubscription).Execute()

Create individual EE subscription for a group of UEs or any UE

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
    ueGroupId := "ueGroupId_example" // string | Group of UEs or any UE
    eeSubscription := TODO // EeSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExposureGroupSubscriptionsCollectionApi.CreateEeGroupSubscriptions(context.Background(), ueGroupId).EeSubscription(eeSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExposureGroupSubscriptionsCollectionApi.CreateEeGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEeGroupSubscriptions`: EeSubscription
    fmt.Fprintf(os.Stdout, "Response from `EventExposureGroupSubscriptionsCollectionApi.CreateEeGroupSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** | Group of UEs or any UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEeGroupSubscriptionsRequest struct via the builder pattern


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


## QueryEeGroupSubscriptions

> []EeSubscription QueryEeGroupSubscriptions(ctx, ueGroupId).SupportedFeatures(supportedFeatures).Execute()

Retrieves the ee subscriptions of a group of UEs or any UE

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
    ueGroupId := "ueGroupId_example" // string | Group of UEs or any UE
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExposureGroupSubscriptionsCollectionApi.QueryEeGroupSubscriptions(context.Background(), ueGroupId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExposureGroupSubscriptionsCollectionApi.QueryEeGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryEeGroupSubscriptions`: []EeSubscription
    fmt.Fprintf(os.Stdout, "Response from `EventExposureGroupSubscriptionsCollectionApi.QueryEeGroupSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** | Group of UEs or any UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryEeGroupSubscriptionsRequest struct via the builder pattern


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


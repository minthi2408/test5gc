# \SDMSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSdmSubscriptions**](SDMSubscriptionsCollectionApi.md#CreateSdmSubscriptions) | **Post** /subscription-data/{ueId}/context-data/sdm-subscriptions | Create individual sdm subscription
[**Querysdmsubscriptions**](SDMSubscriptionsCollectionApi.md#Querysdmsubscriptions) | **Get** /subscription-data/{ueId}/context-data/sdm-subscriptions | Retrieves the sdm subscriptions of a UE



## CreateSdmSubscriptions

> SdmSubscription CreateSdmSubscriptions(ctx, ueId).SdmSubscription(sdmSubscription).Execute()

Create individual sdm subscription

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
    sdmSubscription := TODO // SdmSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDMSubscriptionsCollectionApi.CreateSdmSubscriptions(context.Background(), ueId).SdmSubscription(sdmSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDMSubscriptionsCollectionApi.CreateSdmSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSdmSubscriptions`: SdmSubscription
    fmt.Fprintf(os.Stdout, "Response from `SDMSubscriptionsCollectionApi.CreateSdmSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdmSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sdmSubscription** | [**SdmSubscription**](SdmSubscription.md) |  | 

### Return type

[**SdmSubscription**](SdmSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Querysdmsubscriptions

> []SdmSubscription Querysdmsubscriptions(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

Retrieves the sdm subscriptions of a UE

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
    resp, r, err := apiClient.SDMSubscriptionsCollectionApi.Querysdmsubscriptions(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDMSubscriptionsCollectionApi.Querysdmsubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Querysdmsubscriptions`: []SdmSubscription
    fmt.Fprintf(os.Stdout, "Response from `SDMSubscriptionsCollectionApi.Querysdmsubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerysdmsubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**[]SdmSubscription**](SdmSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


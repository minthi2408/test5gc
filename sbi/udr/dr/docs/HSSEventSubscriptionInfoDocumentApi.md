# \HSSEventSubscriptionInfoDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHSSSubscriptions**](HSSEventSubscriptionInfoDocumentApi.md#CreateHSSSubscriptions) | **Put** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions | Create HSS Subscription Info
[**GetHssSubscriptionInfo**](HSSEventSubscriptionInfoDocumentApi.md#GetHssSubscriptionInfo) | **Get** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions | Retrieve HSS Subscription Info
[**ModifyHssSubscriptionInfo**](HSSEventSubscriptionInfoDocumentApi.md#ModifyHssSubscriptionInfo) | **Patch** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions | Modify HSS Subscription Info
[**RemoveHssSubscriptionsInfo**](HSSEventSubscriptionInfoDocumentApi.md#RemoveHssSubscriptionsInfo) | **Delete** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions | Delete HSS Subscription Info



## CreateHSSSubscriptions

> CreateHSSSubscriptions(ctx, ueId, subsId).HssSubscriptionInfo(hssSubscriptionInfo).Execute()

Create HSS Subscription Info

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
    ueId := "ueId_example" // string | 
    subsId := "subsId_example" // string | 
    hssSubscriptionInfo := TODO // HssSubscriptionInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HSSEventSubscriptionInfoDocumentApi.CreateHSSSubscriptions(context.Background(), ueId, subsId).HssSubscriptionInfo(hssSubscriptionInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HSSEventSubscriptionInfoDocumentApi.CreateHSSSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHSSSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hssSubscriptionInfo** | [**HssSubscriptionInfo**](HssSubscriptionInfo.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHssSubscriptionInfo

> SmfSubscriptionInfo GetHssSubscriptionInfo(ctx, ueId, subsId).Execute()

Retrieve HSS Subscription Info

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
    ueId := "ueId_example" // string | 
    subsId := "subsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HSSEventSubscriptionInfoDocumentApi.GetHssSubscriptionInfo(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HSSEventSubscriptionInfoDocumentApi.GetHssSubscriptionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHssSubscriptionInfo`: SmfSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `HSSEventSubscriptionInfoDocumentApi.GetHssSubscriptionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHssSubscriptionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SmfSubscriptionInfo**](SmfSubscriptionInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyHssSubscriptionInfo

> PatchResult ModifyHssSubscriptionInfo(ctx, ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify HSS Subscription Info

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
    ueId := "ueId_example" // string | 
    subsId := "subsId_example" // string | 
    patchItem := []PatchItem{"TODO"} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HSSEventSubscriptionInfoDocumentApi.ModifyHssSubscriptionInfo(context.Background(), ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HSSEventSubscriptionInfoDocumentApi.ModifyHssSubscriptionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyHssSubscriptionInfo`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `HSSEventSubscriptionInfoDocumentApi.ModifyHssSubscriptionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyHssSubscriptionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveHssSubscriptionsInfo

> RemoveHssSubscriptionsInfo(ctx, ueId, subsId).Execute()

Delete HSS Subscription Info

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
    ueId := "ueId_example" // string | 
    subsId := "subsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HSSEventSubscriptionInfoDocumentApi.RemoveHssSubscriptionsInfo(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HSSEventSubscriptionInfoDocumentApi.RemoveHssSubscriptionsInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveHssSubscriptionsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


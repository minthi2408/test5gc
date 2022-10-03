# \HSSSDMSubscriptionInfoDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHSSSDMSubscriptions**](HSSSDMSubscriptionInfoDocumentApi.md#CreateHSSSDMSubscriptions) | **Put** /subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}/hss-sdm-subscriptions | Create HSS SDM Subscription Info
[**GetHssSDMSubscriptionInfo**](HSSSDMSubscriptionInfoDocumentApi.md#GetHssSDMSubscriptionInfo) | **Get** /subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}/hss-sdm-subscriptions | Retrieve HSS SDM Subscription Info
[**ModifyHssSDMSubscriptionInfo**](HSSSDMSubscriptionInfoDocumentApi.md#ModifyHssSDMSubscriptionInfo) | **Patch** /subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}/hss-sdm-subscriptions | Modify HSS SDM Subscription Info
[**RemoveHssSDMSubscriptionsInfo**](HSSSDMSubscriptionInfoDocumentApi.md#RemoveHssSDMSubscriptionsInfo) | **Delete** /subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}/hss-sdm-subscriptions | Delete HSS SDM Subscription Info



## CreateHSSSDMSubscriptions

> CreateHSSSDMSubscriptions(ctx, ueId, subsId).HssSubscriptionInfo(hssSubscriptionInfo).Execute()

Create HSS SDM Subscription Info

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
    resp, r, err := apiClient.HSSSDMSubscriptionInfoDocumentApi.CreateHSSSDMSubscriptions(context.Background(), ueId, subsId).HssSubscriptionInfo(hssSubscriptionInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HSSSDMSubscriptionInfoDocumentApi.CreateHSSSDMSubscriptions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateHSSSDMSubscriptionsRequest struct via the builder pattern


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


## GetHssSDMSubscriptionInfo

> SmfSubscriptionInfo GetHssSDMSubscriptionInfo(ctx, ueId, subsId).Execute()

Retrieve HSS SDM Subscription Info

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
    resp, r, err := apiClient.HSSSDMSubscriptionInfoDocumentApi.GetHssSDMSubscriptionInfo(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HSSSDMSubscriptionInfoDocumentApi.GetHssSDMSubscriptionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHssSDMSubscriptionInfo`: SmfSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `HSSSDMSubscriptionInfoDocumentApi.GetHssSDMSubscriptionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHssSDMSubscriptionInfoRequest struct via the builder pattern


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


## ModifyHssSDMSubscriptionInfo

> PatchResult ModifyHssSDMSubscriptionInfo(ctx, ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify HSS SDM Subscription Info

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
    resp, r, err := apiClient.HSSSDMSubscriptionInfoDocumentApi.ModifyHssSDMSubscriptionInfo(context.Background(), ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HSSSDMSubscriptionInfoDocumentApi.ModifyHssSDMSubscriptionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyHssSDMSubscriptionInfo`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `HSSSDMSubscriptionInfoDocumentApi.ModifyHssSDMSubscriptionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyHssSDMSubscriptionInfoRequest struct via the builder pattern


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


## RemoveHssSDMSubscriptionsInfo

> RemoveHssSDMSubscriptionsInfo(ctx, ueId, subsId).Execute()

Delete HSS SDM Subscription Info

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
    resp, r, err := apiClient.HSSSDMSubscriptionInfoDocumentApi.RemoveHssSDMSubscriptionsInfo(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HSSSDMSubscriptionInfoDocumentApi.RemoveHssSDMSubscriptionsInfo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveHssSDMSubscriptionsInfoRequest struct via the builder pattern


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


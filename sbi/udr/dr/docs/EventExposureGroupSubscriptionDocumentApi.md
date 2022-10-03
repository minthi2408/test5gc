# \EventExposureGroupSubscriptionDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModifyEeGroupSubscription**](EventExposureGroupSubscriptionDocumentApi.md#ModifyEeGroupSubscription) | **Patch** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId} | Modify an individual ee subscription for a group of a UEs
[**QueryEeGroupSubscription**](EventExposureGroupSubscriptionDocumentApi.md#QueryEeGroupSubscription) | **Get** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId} | Retrieve a individual eeSubscription for a group of UEs or any UE
[**RemoveEeGroupSubscriptions**](EventExposureGroupSubscriptionDocumentApi.md#RemoveEeGroupSubscriptions) | **Delete** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId} | Deletes a eeSubscription for a group of UEs or any UE
[**UpdateEeGroupSubscriptions**](EventExposureGroupSubscriptionDocumentApi.md#UpdateEeGroupSubscriptions) | **Put** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId} | Update an individual ee subscription of a group of UEs or any UE



## ModifyEeGroupSubscription

> PatchResult ModifyEeGroupSubscription(ctx, ueGroupId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify an individual ee subscription for a group of a UEs

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
    ueGroupId := "ueGroupId_example" // string | 
    subsId := "subsId_example" // string | 
    patchItem := []PatchItem{"TODO"} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExposureGroupSubscriptionDocumentApi.ModifyEeGroupSubscription(context.Background(), ueGroupId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExposureGroupSubscriptionDocumentApi.ModifyEeGroupSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyEeGroupSubscription`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `EventExposureGroupSubscriptionDocumentApi.ModifyEeGroupSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyEeGroupSubscriptionRequest struct via the builder pattern


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


## QueryEeGroupSubscription

> map[string]interface{} QueryEeGroupSubscription(ctx, ueGroupId, subsId).Execute()

Retrieve a individual eeSubscription for a group of UEs or any UE

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
    ueGroupId := "ueGroupId_example" // string | 
    subsId := "subsId_example" // string | Unique ID of the subscription to remove

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExposureGroupSubscriptionDocumentApi.QueryEeGroupSubscription(context.Background(), ueGroupId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExposureGroupSubscriptionDocumentApi.QueryEeGroupSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryEeGroupSubscription`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EventExposureGroupSubscriptionDocumentApi.QueryEeGroupSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** | Unique ID of the subscription to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryEeGroupSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveEeGroupSubscriptions

> RemoveEeGroupSubscriptions(ctx, ueGroupId, subsId).Execute()

Deletes a eeSubscription for a group of UEs or any UE

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
    ueGroupId := "ueGroupId_example" // string | 
    subsId := "subsId_example" // string | Unique ID of the subscription to remove

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExposureGroupSubscriptionDocumentApi.RemoveEeGroupSubscriptions(context.Background(), ueGroupId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExposureGroupSubscriptionDocumentApi.RemoveEeGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** | Unique ID of the subscription to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEeGroupSubscriptionsRequest struct via the builder pattern


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


## UpdateEeGroupSubscriptions

> UpdateEeGroupSubscriptions(ctx, ueGroupId, subsId).EeSubscription(eeSubscription).Execute()

Update an individual ee subscription of a group of UEs or any UE

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
    ueGroupId := "ueGroupId_example" // string | 
    subsId := "subsId_example" // string | 
    eeSubscription := TODO // EeSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExposureGroupSubscriptionDocumentApi.UpdateEeGroupSubscriptions(context.Background(), ueGroupId, subsId).EeSubscription(eeSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExposureGroupSubscriptionDocumentApi.UpdateEeGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEeGroupSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **eeSubscription** | [**EeSubscription**](EeSubscription.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


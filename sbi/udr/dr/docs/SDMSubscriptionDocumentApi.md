# \SDMSubscriptionDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModifysdmSubscription**](SDMSubscriptionDocumentApi.md#ModifysdmSubscription) | **Patch** /subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId} | Modify an individual sdm subscription
[**QuerysdmSubscription**](SDMSubscriptionDocumentApi.md#QuerysdmSubscription) | **Get** /subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId} | Retrieves a individual sdmSubscription identified by subsId
[**RemovesdmSubscriptions**](SDMSubscriptionDocumentApi.md#RemovesdmSubscriptions) | **Delete** /subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId} | Deletes a sdmsubscriptions
[**Updatesdmsubscriptions**](SDMSubscriptionDocumentApi.md#Updatesdmsubscriptions) | **Put** /subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId} | Update an individual sdm subscriptions of a UE



## ModifysdmSubscription

> PatchResult ModifysdmSubscription(ctx, ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Modify an individual sdm subscription

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
    subsId := "subsId_example" // string | 
    patchItem := []PatchItem{"TODO"} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDMSubscriptionDocumentApi.ModifysdmSubscription(context.Background(), ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDMSubscriptionDocumentApi.ModifysdmSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifysdmSubscription`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `SDMSubscriptionDocumentApi.ModifysdmSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifysdmSubscriptionRequest struct via the builder pattern


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


## QuerysdmSubscription

> map[string]interface{} QuerysdmSubscription(ctx, ueId, subsId).Execute()

Retrieves a individual sdmSubscription identified by subsId

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
    subsId := "subsId_example" // string | Unique ID of the subscription to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDMSubscriptionDocumentApi.QuerysdmSubscription(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDMSubscriptionDocumentApi.QuerysdmSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerysdmSubscription`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SDMSubscriptionDocumentApi.QuerysdmSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** | Unique ID of the subscription to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerysdmSubscriptionRequest struct via the builder pattern


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


## RemovesdmSubscriptions

> RemovesdmSubscriptions(ctx, ueId, subsId).Execute()

Deletes a sdmsubscriptions

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
    subsId := "subsId_example" // string | Unique ID of the subscription to remove

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDMSubscriptionDocumentApi.RemovesdmSubscriptions(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDMSubscriptionDocumentApi.RemovesdmSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** | Unique ID of the subscription to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovesdmSubscriptionsRequest struct via the builder pattern


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


## Updatesdmsubscriptions

> Updatesdmsubscriptions(ctx, ueId, subsId).SdmSubscription(sdmSubscription).Execute()

Update an individual sdm subscriptions of a UE

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
    sdmSubscription := TODO // SdmSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SDMSubscriptionDocumentApi.Updatesdmsubscriptions(context.Background(), ueId, subsId).SdmSubscription(sdmSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SDMSubscriptionDocumentApi.Updatesdmsubscriptions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdatesdmsubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sdmSubscription** | [**SdmSubscription**](SdmSubscription.md) |  | 

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


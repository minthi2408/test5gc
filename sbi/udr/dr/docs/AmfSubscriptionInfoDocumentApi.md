# \AmfSubscriptionInfoDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModifyAmfSubscriptionInfo**](AmfSubscriptionInfoDocumentApi.md#ModifyAmfSubscriptionInfo) | **Patch** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/amf-subscriptions | modify the AMF Subscription Info



## ModifyAmfSubscriptionInfo

> PatchResult ModifyAmfSubscriptionInfo(ctx, ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

modify the AMF Subscription Info

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
    resp, r, err := apiClient.AmfSubscriptionInfoDocumentApi.ModifyAmfSubscriptionInfo(context.Background(), ueId, subsId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AmfSubscriptionInfoDocumentApi.ModifyAmfSubscriptionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyAmfSubscriptionInfo`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `AmfSubscriptionInfoDocumentApi.ModifyAmfSubscriptionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAmfSubscriptionInfoRequest struct via the builder pattern


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


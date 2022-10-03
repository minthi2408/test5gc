# \AMFSubscriptionInfoDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAMFSubscriptions**](AMFSubscriptionInfoDocumentApi.md#CreateAMFSubscriptions) | **Put** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/amf-subscriptions | Create AmfSubscriptions for an individual ee subscriptions of a UE



## CreateAMFSubscriptions

> CreateAMFSubscriptions(ctx, ueId, subsId).AmfSubscriptionInfo(amfSubscriptionInfo).Execute()

Create AmfSubscriptions for an individual ee subscriptions of a UE

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
    amfSubscriptionInfo := []AmfSubscriptionInfo{"TODO"} // []AmfSubscriptionInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMFSubscriptionInfoDocumentApi.CreateAMFSubscriptions(context.Background(), ueId, subsId).AmfSubscriptionInfo(amfSubscriptionInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMFSubscriptionInfoDocumentApi.CreateAMFSubscriptions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateAMFSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amfSubscriptionInfo** | [**[]AmfSubscriptionInfo**](AmfSubscriptionInfo.md) |  | 

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


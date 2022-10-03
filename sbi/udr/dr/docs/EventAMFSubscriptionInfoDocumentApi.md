# \EventAMFSubscriptionInfoDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveAmfSubscriptionsInfo**](EventAMFSubscriptionInfoDocumentApi.md#RemoveAmfSubscriptionsInfo) | **Delete** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/amf-subscriptions | Deletes AMF Subscription Info for an eeSubscription



## RemoveAmfSubscriptionsInfo

> RemoveAmfSubscriptionsInfo(ctx, ueId, subsId).Execute()

Deletes AMF Subscription Info for an eeSubscription

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
    resp, r, err := apiClient.EventAMFSubscriptionInfoDocumentApi.RemoveAmfSubscriptionsInfo(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventAMFSubscriptionInfoDocumentApi.RemoveAmfSubscriptionsInfo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveAmfSubscriptionsInfoRequest struct via the builder pattern


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


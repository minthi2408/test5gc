# \SubscriptionsCollectionDocumentApi

All URIs are relative to *https://example.com/namf-evts/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsCollectionDocumentApi.md#CreateSubscription) | **Post** /subscriptions | Namf_EventExposure Subscribe service Operation



## CreateSubscription

> AmfCreatedEventSubscription CreateSubscription(ctx).AmfCreateEventSubscription(amfCreateEventSubscription).Execute()

Namf_EventExposure Subscribe service Operation

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
    amfCreateEventSubscription := TODO // AmfCreateEventSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionDocumentApi.CreateSubscription(context.Background()).AmfCreateEventSubscription(amfCreateEventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionDocumentApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: AmfCreatedEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionDocumentApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amfCreateEventSubscription** | [**AmfCreateEventSubscription**](AmfCreateEventSubscription.md) |  | 

### Return type

[**AmfCreatedEventSubscription**](AmfCreatedEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


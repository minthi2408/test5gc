# \SubsToNotifyCollectionApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuerySubsToNotify**](SubsToNotifyCollectionApi.md#QuerySubsToNotify) | **Get** /subscription-data/subs-to-notify | Retrieves the list of subscriptions
[**RemoveMultipleSubscriptionDataSubscriptions**](SubsToNotifyCollectionApi.md#RemoveMultipleSubscriptionDataSubscriptions) | **Delete** /subscription-data/subs-to-notify | Deletes subscriptions identified by a given ue-id parameter
[**SubscriptionDataSubscriptions**](SubsToNotifyCollectionApi.md#SubscriptionDataSubscriptions) | **Post** /subscription-data/subs-to-notify | Subscription data subscriptions



## QuerySubsToNotify

> []SubscriptionDataSubscriptions QuerySubsToNotify(ctx).UeId(ueId).SupportedFeatures(supportedFeatures).Execute()

Retrieves the list of subscriptions

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
    resp, r, err := apiClient.SubsToNotifyCollectionApi.QuerySubsToNotify(context.Background()).UeId(ueId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubsToNotifyCollectionApi.QuerySubsToNotify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySubsToNotify`: []SubscriptionDataSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `SubsToNotifyCollectionApi.QuerySubsToNotify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuerySubsToNotifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ueId** | **string** | UE id | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**[]SubscriptionDataSubscriptions**](SubscriptionDataSubscriptions.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMultipleSubscriptionDataSubscriptions

> RemoveMultipleSubscriptionDataSubscriptions(ctx).UeId(ueId).NfInstanceId(nfInstanceId).DeleteAllNfs(deleteAllNfs).ImplicitUnsubscribeIndication(implicitUnsubscribeIndication).Execute()

Deletes subscriptions identified by a given ue-id parameter

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
    nfInstanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | NF Instance ID (optional)
    deleteAllNfs := true // bool | Flag to delete subscriptions from all NFs (optional)
    implicitUnsubscribeIndication := true // bool | Implicit Unsubscribe Indication (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubsToNotifyCollectionApi.RemoveMultipleSubscriptionDataSubscriptions(context.Background()).UeId(ueId).NfInstanceId(nfInstanceId).DeleteAllNfs(deleteAllNfs).ImplicitUnsubscribeIndication(implicitUnsubscribeIndication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubsToNotifyCollectionApi.RemoveMultipleSubscriptionDataSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMultipleSubscriptionDataSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ueId** | **string** | UE ID | 
 **nfInstanceId** | **string** | NF Instance ID | 
 **deleteAllNfs** | **bool** | Flag to delete subscriptions from all NFs | 
 **implicitUnsubscribeIndication** | **bool** | Implicit Unsubscribe Indication | 

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


## SubscriptionDataSubscriptions

> SubscriptionDataSubscriptions SubscriptionDataSubscriptions(ctx).SubscriptionDataSubscriptions(subscriptionDataSubscriptions).Execute()

Subscription data subscriptions

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
    subscriptionDataSubscriptions := TODO // SubscriptionDataSubscriptions | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubsToNotifyCollectionApi.SubscriptionDataSubscriptions(context.Background()).SubscriptionDataSubscriptions(subscriptionDataSubscriptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubsToNotifyCollectionApi.SubscriptionDataSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionDataSubscriptions`: SubscriptionDataSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `SubsToNotifyCollectionApi.SubscriptionDataSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionDataSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionDataSubscriptions** | [**SubscriptionDataSubscriptions**](SubscriptionDataSubscriptions.md) |  | 

### Return type

[**SubscriptionDataSubscriptions**](SubscriptionDataSubscriptions.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


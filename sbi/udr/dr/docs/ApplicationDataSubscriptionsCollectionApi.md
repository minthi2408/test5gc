# \ApplicationDataSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualApplicationDataSubscription**](ApplicationDataSubscriptionsCollectionApi.md#CreateIndividualApplicationDataSubscription) | **Post** /application-data/subs-to-notify | Create a subscription to receive notification of application data changes
[**ReadApplicationDataChangeSubscriptions**](ApplicationDataSubscriptionsCollectionApi.md#ReadApplicationDataChangeSubscriptions) | **Get** /application-data/subs-to-notify | Read Application Data change Subscriptions



## CreateIndividualApplicationDataSubscription

> ApplicationDataSubs CreateIndividualApplicationDataSubscription(ctx).ApplicationDataSubs(applicationDataSubs).Execute()

Create a subscription to receive notification of application data changes

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
    applicationDataSubs := TODO // ApplicationDataSubs | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationDataSubscriptionsCollectionApi.CreateIndividualApplicationDataSubscription(context.Background()).ApplicationDataSubs(applicationDataSubs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDataSubscriptionsCollectionApi.CreateIndividualApplicationDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualApplicationDataSubscription`: ApplicationDataSubs
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDataSubscriptionsCollectionApi.CreateIndividualApplicationDataSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualApplicationDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDataSubs** | [**ApplicationDataSubs**](ApplicationDataSubs.md) |  | 

### Return type

[**ApplicationDataSubs**](ApplicationDataSubs.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadApplicationDataChangeSubscriptions

> []ApplicationDataSubs ReadApplicationDataChangeSubscriptions(ctx).DataFilter(dataFilter).Execute()

Read Application Data change Subscriptions

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
    dataFilter := DataFilter{ ... } // DataFilter | The data filter for the query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationDataSubscriptionsCollectionApi.ReadApplicationDataChangeSubscriptions(context.Background()).DataFilter(dataFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationDataSubscriptionsCollectionApi.ReadApplicationDataChangeSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadApplicationDataChangeSubscriptions`: []ApplicationDataSubs
    fmt.Fprintf(os.Stdout, "Response from `ApplicationDataSubscriptionsCollectionApi.ReadApplicationDataChangeSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadApplicationDataChangeSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataFilter** | [**DataFilter**](DataFilter.md) | The data filter for the query. | 

### Return type

[**[]ApplicationDataSubs**](ApplicationDataSubs.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


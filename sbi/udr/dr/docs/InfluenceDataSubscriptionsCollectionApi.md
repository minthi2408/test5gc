# \InfluenceDataSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualInfluenceDataSubscription**](InfluenceDataSubscriptionsCollectionApi.md#CreateIndividualInfluenceDataSubscription) | **Post** /application-data/influenceData/subs-to-notify | Create a new Individual Influence Data Subscription resource
[**ReadInfluenceDataSubscriptions**](InfluenceDataSubscriptionsCollectionApi.md#ReadInfluenceDataSubscriptions) | **Get** /application-data/influenceData/subs-to-notify | Read Influence Data Subscriptions



## CreateIndividualInfluenceDataSubscription

> TrafficInfluSub CreateIndividualInfluenceDataSubscription(ctx).TrafficInfluSub(trafficInfluSub).Execute()

Create a new Individual Influence Data Subscription resource

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
    trafficInfluSub := TODO // TrafficInfluSub | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfluenceDataSubscriptionsCollectionApi.CreateIndividualInfluenceDataSubscription(context.Background()).TrafficInfluSub(trafficInfluSub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfluenceDataSubscriptionsCollectionApi.CreateIndividualInfluenceDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualInfluenceDataSubscription`: TrafficInfluSub
    fmt.Fprintf(os.Stdout, "Response from `InfluenceDataSubscriptionsCollectionApi.CreateIndividualInfluenceDataSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualInfluenceDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trafficInfluSub** | [**TrafficInfluSub**](TrafficInfluSub.md) |  | 

### Return type

[**TrafficInfluSub**](TrafficInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadInfluenceDataSubscriptions

> []TrafficInfluSub ReadInfluenceDataSubscriptions(ctx).Dnn(dnn).Snssai(snssai).InternalGroupId(internalGroupId).Supi(supi).Execute()

Read Influence Data Subscriptions

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
    dnn := "dnn_example" // string | Identifies a DNN. (optional)
    snssai := Snssai{ ... } // Snssai | Identifies a slice. (optional)
    internalGroupId := "internalGroupId_example" // string | Identifies a group of users. (optional)
    supi := "supi_example" // string | Identifies a user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfluenceDataSubscriptionsCollectionApi.ReadInfluenceDataSubscriptions(context.Background()).Dnn(dnn).Snssai(snssai).InternalGroupId(internalGroupId).Supi(supi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfluenceDataSubscriptionsCollectionApi.ReadInfluenceDataSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadInfluenceDataSubscriptions`: []TrafficInfluSub
    fmt.Fprintf(os.Stdout, "Response from `InfluenceDataSubscriptionsCollectionApi.ReadInfluenceDataSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadInfluenceDataSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnn** | **string** | Identifies a DNN. | 
 **snssai** | [**Snssai**](Snssai.md) | Identifies a slice. | 
 **internalGroupId** | **string** | Identifies a group of users. | 
 **supi** | **string** | Identifies a user. | 

### Return type

[**[]TrafficInfluSub**](TrafficInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


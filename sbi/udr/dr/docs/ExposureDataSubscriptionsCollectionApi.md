# \ExposureDataSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualExposureDataSubscription**](ExposureDataSubscriptionsCollectionApi.md#CreateIndividualExposureDataSubscription) | **Post** /exposure-data/subs-to-notify | Create a subscription to receive notification of exposure data changes



## CreateIndividualExposureDataSubscription

> ExposureDataSubscription CreateIndividualExposureDataSubscription(ctx).ExposureDataSubscription(exposureDataSubscription).Execute()

Create a subscription to receive notification of exposure data changes

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
    exposureDataSubscription := TODO // ExposureDataSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExposureDataSubscriptionsCollectionApi.CreateIndividualExposureDataSubscription(context.Background()).ExposureDataSubscription(exposureDataSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExposureDataSubscriptionsCollectionApi.CreateIndividualExposureDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualExposureDataSubscription`: ExposureDataSubscription
    fmt.Fprintf(os.Stdout, "Response from `ExposureDataSubscriptionsCollectionApi.CreateIndividualExposureDataSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualExposureDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exposureDataSubscription** | [**ExposureDataSubscription**](ExposureDataSubscription.md) |  | 

### Return type

[**ExposureDataSubscription**](ExposureDataSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


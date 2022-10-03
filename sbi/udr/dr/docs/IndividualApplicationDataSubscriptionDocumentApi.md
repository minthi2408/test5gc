# \IndividualApplicationDataSubscriptionDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualApplicationDataSubscription**](IndividualApplicationDataSubscriptionDocumentApi.md#DeleteIndividualApplicationDataSubscription) | **Delete** /application-data/subs-to-notify/{subsId} | Delete the individual Application Data subscription
[**ReadIndividualApplicationDataSubscription**](IndividualApplicationDataSubscriptionDocumentApi.md#ReadIndividualApplicationDataSubscription) | **Get** /application-data/subs-to-notify/{subsId} | Get an existing individual Application Data Subscription resource
[**ReplaceIndividualApplicationDataSubscription**](IndividualApplicationDataSubscriptionDocumentApi.md#ReplaceIndividualApplicationDataSubscription) | **Put** /application-data/subs-to-notify/{subsId} | Modify a subscription to receive notification of application data changes



## DeleteIndividualApplicationDataSubscription

> DeleteIndividualApplicationDataSubscription(ctx, subsId).Execute()

Delete the individual Application Data subscription

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
    subsId := "subsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationDataSubscriptionDocumentApi.DeleteIndividualApplicationDataSubscription(context.Background(), subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationDataSubscriptionDocumentApi.DeleteIndividualApplicationDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualApplicationDataSubscriptionRequest struct via the builder pattern


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


## ReadIndividualApplicationDataSubscription

> ApplicationDataSubs ReadIndividualApplicationDataSubscription(ctx, subsId).Execute()

Get an existing individual Application Data Subscription resource

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
    subsId := "subsId_example" // string | String identifying a subscription to the Individual Application Data Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationDataSubscriptionDocumentApi.ReadIndividualApplicationDataSubscription(context.Background(), subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationDataSubscriptionDocumentApi.ReadIndividualApplicationDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualApplicationDataSubscription`: ApplicationDataSubs
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationDataSubscriptionDocumentApi.ReadIndividualApplicationDataSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subsId** | **string** | String identifying a subscription to the Individual Application Data Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualApplicationDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationDataSubs**](ApplicationDataSubs.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceIndividualApplicationDataSubscription

> ApplicationDataSubs ReplaceIndividualApplicationDataSubscription(ctx, subsId).ApplicationDataSubs(applicationDataSubs).Execute()

Modify a subscription to receive notification of application data changes

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
    subsId := "subsId_example" // string | 
    applicationDataSubs := TODO // ApplicationDataSubs | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationDataSubscriptionDocumentApi.ReplaceIndividualApplicationDataSubscription(context.Background(), subsId).ApplicationDataSubs(applicationDataSubs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationDataSubscriptionDocumentApi.ReplaceIndividualApplicationDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIndividualApplicationDataSubscription`: ApplicationDataSubs
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationDataSubscriptionDocumentApi.ReplaceIndividualApplicationDataSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIndividualApplicationDataSubscriptionRequest struct via the builder pattern


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


# \IndividualPolicyDataSubscriptionDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualPolicyDataSubscription**](IndividualPolicyDataSubscriptionDocumentApi.md#DeleteIndividualPolicyDataSubscription) | **Delete** /policy-data/subs-to-notify/{subsId} | Delete the individual Policy Data subscription
[**ReplaceIndividualPolicyDataSubscription**](IndividualPolicyDataSubscriptionDocumentApi.md#ReplaceIndividualPolicyDataSubscription) | **Put** /policy-data/subs-to-notify/{subsId} | Modify a subscription to receive notification of policy data changes



## DeleteIndividualPolicyDataSubscription

> DeleteIndividualPolicyDataSubscription(ctx, subsId).Execute()

Delete the individual Policy Data subscription

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
    resp, r, err := apiClient.IndividualPolicyDataSubscriptionDocumentApi.DeleteIndividualPolicyDataSubscription(context.Background(), subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPolicyDataSubscriptionDocumentApi.DeleteIndividualPolicyDataSubscription``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIndividualPolicyDataSubscriptionRequest struct via the builder pattern


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


## ReplaceIndividualPolicyDataSubscription

> PolicyDataSubscription ReplaceIndividualPolicyDataSubscription(ctx, subsId).PolicyDataSubscription(policyDataSubscription).Execute()

Modify a subscription to receive notification of policy data changes

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
    policyDataSubscription := TODO // PolicyDataSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPolicyDataSubscriptionDocumentApi.ReplaceIndividualPolicyDataSubscription(context.Background(), subsId).PolicyDataSubscription(policyDataSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPolicyDataSubscriptionDocumentApi.ReplaceIndividualPolicyDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIndividualPolicyDataSubscription`: PolicyDataSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualPolicyDataSubscriptionDocumentApi.ReplaceIndividualPolicyDataSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIndividualPolicyDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyDataSubscription** | [**PolicyDataSubscription**](PolicyDataSubscription.md) |  | 

### Return type

[**PolicyDataSubscription**](PolicyDataSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \N1N2IndividualSubscriptionDocumentApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**N1N2MessageUnSubscribe**](N1N2IndividualSubscriptionDocumentApi.md#N1N2MessageUnSubscribe) | **Delete** /ue-contexts/{ueContextId}/n1-n2-messages/subscriptions/{subscriptionId} | Namf_Communication N1N2 Message UnSubscribe (UE Specific) service Operation



## N1N2MessageUnSubscribe

> N1N2MessageUnSubscribe(ctx, ueContextId, subscriptionId).Execute()

Namf_Communication N1N2 Message UnSubscribe (UE Specific) service Operation

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
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    subscriptionId := "subscriptionId_example" // string | Subscription Identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N1N2IndividualSubscriptionDocumentApi.N1N2MessageUnSubscribe(context.Background(), ueContextId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N1N2IndividualSubscriptionDocumentApi.N1N2MessageUnSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 
**subscriptionId** | **string** | Subscription Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiN1N2MessageUnSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


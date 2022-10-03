# \N1N2SubscriptionsCollectionForIndividualUEContextsDocumentApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**N1N2MessageSubscribe**](N1N2SubscriptionsCollectionForIndividualUEContextsDocumentApi.md#N1N2MessageSubscribe) | **Post** /ue-contexts/{ueContextId}/n1-n2-messages/subscriptions | Namf_Communication N1N2 Message Subscribe (UE Specific) service Operation



## N1N2MessageSubscribe

> UeN1N2InfoSubscriptionCreatedData N1N2MessageSubscribe(ctx, ueContextId).UeN1N2InfoSubscriptionCreateData(ueN1N2InfoSubscriptionCreateData).Execute()

Namf_Communication N1N2 Message Subscribe (UE Specific) service Operation

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
    ueN1N2InfoSubscriptionCreateData := TODO // UeN1N2InfoSubscriptionCreateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N1N2SubscriptionsCollectionForIndividualUEContextsDocumentApi.N1N2MessageSubscribe(context.Background(), ueContextId).UeN1N2InfoSubscriptionCreateData(ueN1N2InfoSubscriptionCreateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N1N2SubscriptionsCollectionForIndividualUEContextsDocumentApi.N1N2MessageSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `N1N2MessageSubscribe`: UeN1N2InfoSubscriptionCreatedData
    fmt.Fprintf(os.Stdout, "Response from `N1N2SubscriptionsCollectionForIndividualUEContextsDocumentApi.N1N2MessageSubscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiN1N2MessageSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ueN1N2InfoSubscriptionCreateData** | [**UeN1N2InfoSubscriptionCreateData**](UeN1N2InfoSubscriptionCreateData.md) |  | 

### Return type

[**UeN1N2InfoSubscriptionCreatedData**](UeN1N2InfoSubscriptionCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


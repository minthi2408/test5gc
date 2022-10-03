# \NonUEN2MessageNotificationIndividualSubscriptionDocumentApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NonUeN2InfoUnSubscribe**](NonUEN2MessageNotificationIndividualSubscriptionDocumentApi.md#NonUeN2InfoUnSubscribe) | **Delete** /non-ue-n2-messages/subscriptions/{n2NotifySubscriptionId} | Namf_Communication Non UE N2 Info UnSubscribe service Operation



## NonUeN2InfoUnSubscribe

> NonUeN2InfoUnSubscribe(ctx, n2NotifySubscriptionId).Execute()

Namf_Communication Non UE N2 Info UnSubscribe service Operation

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
    n2NotifySubscriptionId := "n2NotifySubscriptionId_example" // string | N2 info Subscription Identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonUEN2MessageNotificationIndividualSubscriptionDocumentApi.NonUeN2InfoUnSubscribe(context.Background(), n2NotifySubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonUEN2MessageNotificationIndividualSubscriptionDocumentApi.NonUeN2InfoUnSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**n2NotifySubscriptionId** | **string** | N2 info Subscription Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonUeN2InfoUnSubscribeRequest struct via the builder pattern


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


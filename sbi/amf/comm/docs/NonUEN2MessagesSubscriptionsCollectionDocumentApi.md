# \NonUEN2MessagesSubscriptionsCollectionDocumentApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NonUeN2InfoSubscribe**](NonUEN2MessagesSubscriptionsCollectionDocumentApi.md#NonUeN2InfoSubscribe) | **Post** /non-ue-n2-messages/subscriptions | Namf_Communication Non UE N2 Info Subscribe service Operation



## NonUeN2InfoSubscribe

> NonUeN2InfoSubscriptionCreatedData NonUeN2InfoSubscribe(ctx).NonUeN2InfoSubscriptionCreateData(nonUeN2InfoSubscriptionCreateData).Execute()

Namf_Communication Non UE N2 Info Subscribe service Operation

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
    nonUeN2InfoSubscriptionCreateData := TODO // NonUeN2InfoSubscriptionCreateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonUEN2MessagesSubscriptionsCollectionDocumentApi.NonUeN2InfoSubscribe(context.Background()).NonUeN2InfoSubscriptionCreateData(nonUeN2InfoSubscriptionCreateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonUEN2MessagesSubscriptionsCollectionDocumentApi.NonUeN2InfoSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonUeN2InfoSubscribe`: NonUeN2InfoSubscriptionCreatedData
    fmt.Fprintf(os.Stdout, "Response from `NonUEN2MessagesSubscriptionsCollectionDocumentApi.NonUeN2InfoSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonUeN2InfoSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonUeN2InfoSubscriptionCreateData** | [**NonUeN2InfoSubscriptionCreateData**](NonUeN2InfoSubscriptionCreateData.md) |  | 

### Return type

[**NonUeN2InfoSubscriptionCreatedData**](NonUeN2InfoSubscriptionCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


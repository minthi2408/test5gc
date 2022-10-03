# \PolicyDataSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualPolicyDataSubscription**](PolicyDataSubscriptionsCollectionApi.md#CreateIndividualPolicyDataSubscription) | **Post** /policy-data/subs-to-notify | Create a subscription to receive notification of policy data changes



## CreateIndividualPolicyDataSubscription

> PolicyDataSubscription CreateIndividualPolicyDataSubscription(ctx).PolicyDataSubscription(policyDataSubscription).Execute()

Create a subscription to receive notification of policy data changes

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
    policyDataSubscription := TODO // PolicyDataSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyDataSubscriptionsCollectionApi.CreateIndividualPolicyDataSubscription(context.Background()).PolicyDataSubscription(policyDataSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyDataSubscriptionsCollectionApi.CreateIndividualPolicyDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualPolicyDataSubscription`: PolicyDataSubscription
    fmt.Fprintf(os.Stdout, "Response from `PolicyDataSubscriptionsCollectionApi.CreateIndividualPolicyDataSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualPolicyDataSubscriptionRequest struct via the builder pattern


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


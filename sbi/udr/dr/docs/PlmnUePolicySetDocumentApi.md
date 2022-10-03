# \PlmnUePolicySetDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadPlmnUePolicySet**](PlmnUePolicySetDocumentApi.md#ReadPlmnUePolicySet) | **Get** /policy-data/plmns/{plmnId}/ue-policy-set | Retrieve the UE policy set data for an H-PLMN



## ReadPlmnUePolicySet

> UePolicySet ReadPlmnUePolicySet(ctx, plmnId).Execute()

Retrieve the UE policy set data for an H-PLMN

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
    plmnId := "plmnId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlmnUePolicySetDocumentApi.ReadPlmnUePolicySet(context.Background(), plmnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlmnUePolicySetDocumentApi.ReadPlmnUePolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPlmnUePolicySet`: UePolicySet
    fmt.Fprintf(os.Stdout, "Response from `PlmnUePolicySetDocumentApi.ReadPlmnUePolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plmnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadPlmnUePolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UePolicySet**](UePolicySet.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


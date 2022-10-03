# \AccessAndMobilityPolicyDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAccessAndMobilityPolicyData**](AccessAndMobilityPolicyDataDocumentApi.md#ReadAccessAndMobilityPolicyData) | **Get** /policy-data/ues/{ueId}/am-data | Retrieves the access and mobility policy data for a subscriber



## ReadAccessAndMobilityPolicyData

> AmPolicyData ReadAccessAndMobilityPolicyData(ctx, ueId).Execute()

Retrieves the access and mobility policy data for a subscriber

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
    ueId := "ueId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessAndMobilityPolicyDataDocumentApi.ReadAccessAndMobilityPolicyData(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessAndMobilityPolicyDataDocumentApi.ReadAccessAndMobilityPolicyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAccessAndMobilityPolicyData`: AmPolicyData
    fmt.Fprintf(os.Stdout, "Response from `AccessAndMobilityPolicyDataDocumentApi.ReadAccessAndMobilityPolicyData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAccessAndMobilityPolicyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AmPolicyData**](AmPolicyData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


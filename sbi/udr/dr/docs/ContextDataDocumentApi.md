# \ContextDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryContextData**](ContextDataDocumentApi.md#QueryContextData) | **Get** /subscription-data/{ueId}/context-data | Retrieve multiple context data sets of a UE



## QueryContextData

> ContextDataSets QueryContextData(ctx, ueId).ContextDatasetNames(contextDatasetNames).Execute()

Retrieve multiple context data sets of a UE

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
    ueId := "ueId_example" // string | UE id
    contextDatasetNames := []ContextDataSetName{"TODO"} // []ContextDataSetName | List of context dataset names

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextDataDocumentApi.QueryContextData(context.Background(), ueId).ContextDatasetNames(contextDatasetNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextDataDocumentApi.QueryContextData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryContextData`: ContextDataSets
    fmt.Fprintf(os.Stdout, "Response from `ContextDataDocumentApi.QueryContextData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryContextDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contextDatasetNames** | [**[]ContextDataSetName**](ContextDataSetName.md) | List of context dataset names | 

### Return type

[**ContextDataSets**](ContextDataSets.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


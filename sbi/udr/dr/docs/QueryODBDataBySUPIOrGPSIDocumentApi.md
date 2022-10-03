# \QueryODBDataBySUPIOrGPSIDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOdbData**](QueryODBDataBySUPIOrGPSIDocumentApi.md#GetOdbData) | **Get** /subscription-data/{ueId}/operator-determined-barring-data | Retrieve ODB Data data by SUPI or GPSI



## GetOdbData

> OdbData GetOdbData(ctx, ueId).Execute()

Retrieve ODB Data data by SUPI or GPSI

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
    ueId := "ueId_example" // string | UE ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryODBDataBySUPIOrGPSIDocumentApi.GetOdbData(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryODBDataBySUPIOrGPSIDocumentApi.GetOdbData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOdbData`: OdbData
    fmt.Fprintf(os.Stdout, "Response from `QueryODBDataBySUPIOrGPSIDocumentApi.GetOdbData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOdbDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OdbData**](OdbData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


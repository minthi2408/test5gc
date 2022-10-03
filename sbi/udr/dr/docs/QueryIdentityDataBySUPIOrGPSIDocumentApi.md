# \QueryIdentityDataBySUPIOrGPSIDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdentityData**](QueryIdentityDataBySUPIOrGPSIDocumentApi.md#GetIdentityData) | **Get** /subscription-data/{ueId}/identity-data | Retrieve identity data by SUPI or GPSI



## GetIdentityData

> IdentityData GetIdentityData(ctx, ueId).AppPortId(appPortId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Retrieve identity data by SUPI or GPSI

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
    appPortId := AppPortId{ ... } // AppPortId | Application port identifier (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryIdentityDataBySUPIOrGPSIDocumentApi.GetIdentityData(context.Background(), ueId).AppPortId(appPortId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryIdentityDataBySUPIOrGPSIDocumentApi.GetIdentityData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityData`: IdentityData
    fmt.Fprintf(os.Stdout, "Response from `QueryIdentityDataBySUPIOrGPSIDocumentApi.GetIdentityData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appPortId** | [**AppPortId**](AppPortId.md) | Application port identifier | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**IdentityData**](IdentityData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ProvisionedParameterDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModifyPpData**](ProvisionedParameterDataDocumentApi.md#ModifyPpData) | **Patch** /subscription-data/{ueId}/pp-data | modify the provisioned parameter data



## ModifyPpData

> PatchResult ModifyPpData(ctx, ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

modify the provisioned parameter data

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
    ueId := "ueId_example" // string | pp data for a UE
    patchItem := []PatchItem{"TODO"} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisionedParameterDataDocumentApi.ModifyPpData(context.Background(), ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedParameterDataDocumentApi.ModifyPpData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyPpData`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `ProvisionedParameterDataDocumentApi.ModifyPpData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | pp data for a UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyPpDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


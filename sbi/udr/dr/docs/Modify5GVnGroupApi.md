# \Modify5GVnGroupApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Modify5GVnGroup**](Modify5GVnGroupApi.md#Modify5GVnGroup) | **Patch** /subscription-data/group-data/5g-vn-groups/{externalGroupId} | modify the 5GVnGroup



## Modify5GVnGroup

> PatchResult Modify5GVnGroup(ctx, externalGroupId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

modify the 5GVnGroup

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
    externalGroupId := "externalGroupId_example" // string | 
    patchItem := []PatchItem{"TODO"} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Modify5GVnGroupApi.Modify5GVnGroup(context.Background(), externalGroupId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Modify5GVnGroupApi.Modify5GVnGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Modify5GVnGroup`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `Modify5GVnGroupApi.Modify5GVnGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModify5GVnGroupRequest struct via the builder pattern


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


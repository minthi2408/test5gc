# \GroupIdentifiersApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupIdentifiers**](GroupIdentifiersApi.md#GetGroupIdentifiers) | **Get** /subscription-data/group-data/group-identifiers | Mapping of Group Identifiers



## GetGroupIdentifiers

> GroupIdentifiers GetGroupIdentifiers(ctx).ExtGroupId(extGroupId).IntGroupId(intGroupId).UeIdInd(ueIdInd).SupportedFeatures(supportedFeatures).Execute()

Mapping of Group Identifiers

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
    extGroupId := "extGroupId_example" // string | External Group Identifier (optional)
    intGroupId := "intGroupId_example" // string | Internal Group Identifier (optional)
    ueIdInd := true // bool | Indication whether UE identifiers are required or not (optional) (default to false)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupIdentifiersApi.GetGroupIdentifiers(context.Background()).ExtGroupId(extGroupId).IntGroupId(intGroupId).UeIdInd(ueIdInd).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupIdentifiersApi.GetGroupIdentifiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupIdentifiers`: GroupIdentifiers
    fmt.Fprintf(os.Stdout, "Response from `GroupIdentifiersApi.GetGroupIdentifiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extGroupId** | **string** | External Group Identifier | 
 **intGroupId** | **string** | Internal Group Identifier | 
 **ueIdInd** | **bool** | Indication whether UE identifiers are required or not | [default to false]
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**GroupIdentifiers**](GroupIdentifiers.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


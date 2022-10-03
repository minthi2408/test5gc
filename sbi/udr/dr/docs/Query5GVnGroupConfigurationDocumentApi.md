# \Query5GVnGroupConfigurationDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get5GVnGroupConfiguration**](Query5GVnGroupConfigurationDocumentApi.md#Get5GVnGroupConfiguration) | **Get** /subscription-data/group-data/5g-vn-groups/{externalGroupId} | Retrieve a 5GVnGroup configuration



## Get5GVnGroupConfiguration

> Model5GVnGroupConfiguration Get5GVnGroupConfiguration(ctx, externalGroupId).Execute()

Retrieve a 5GVnGroup configuration

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Query5GVnGroupConfigurationDocumentApi.Get5GVnGroupConfiguration(context.Background(), externalGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Query5GVnGroupConfigurationDocumentApi.Get5GVnGroupConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get5GVnGroupConfiguration`: Model5GVnGroupConfiguration
    fmt.Fprintf(os.Stdout, "Response from `Query5GVnGroupConfigurationDocumentApi.Get5GVnGroupConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet5GVnGroupConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Model5GVnGroupConfiguration**](Model5GVnGroupConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


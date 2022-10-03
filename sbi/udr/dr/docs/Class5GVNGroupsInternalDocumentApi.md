# \Class5GVNGroupsInternalDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Query5GVnGroupInternal**](Class5GVNGroupsInternalDocumentApi.md#Query5GVnGroupInternal) | **Get** /subscription-data/group-data/5g-vn-groups/internal | Retrieves the data of 5G VN Group



## Query5GVnGroupInternal

> map[string]Model5GVnGroupConfiguration Query5GVnGroupInternal(ctx).InternalGroupIds(internalGroupIds).Execute()

Retrieves the data of 5G VN Group

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
    internalGroupIds := []string{"Inner_example"} // []string | List of Internal Group IDs

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Class5GVNGroupsInternalDocumentApi.Query5GVnGroupInternal(context.Background()).InternalGroupIds(internalGroupIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Class5GVNGroupsInternalDocumentApi.Query5GVnGroupInternal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Query5GVnGroupInternal`: map[string]Model5GVnGroupConfiguration
    fmt.Fprintf(os.Stdout, "Response from `Class5GVNGroupsInternalDocumentApi.Query5GVnGroupInternal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuery5GVnGroupInternalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internalGroupIds** | **[]string** | List of Internal Group IDs | 

### Return type

[**map[string]Model5GVnGroupConfiguration**](Model5GVnGroupConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


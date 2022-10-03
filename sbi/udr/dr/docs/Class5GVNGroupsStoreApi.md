# \Class5GVNGroupsStoreApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Query5GVnGroup**](Class5GVNGroupsStoreApi.md#Query5GVnGroup) | **Get** /subscription-data/group-data/5g-vn-groups | Retrieves the data of a 5G VN Group



## Query5GVnGroup

> map[string]Model5GVnGroupConfiguration Query5GVnGroup(ctx).Gpsis(gpsis).Execute()

Retrieves the data of a 5G VN Group

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
    gpsis := []string{"Inner_example"} // []string | List of GPSIs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Class5GVNGroupsStoreApi.Query5GVnGroup(context.Background()).Gpsis(gpsis).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Class5GVNGroupsStoreApi.Query5GVnGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Query5GVnGroup`: map[string]Model5GVnGroupConfiguration
    fmt.Fprintf(os.Stdout, "Response from `Class5GVNGroupsStoreApi.Query5GVnGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuery5GVnGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpsis** | **[]string** | List of GPSIs | 

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


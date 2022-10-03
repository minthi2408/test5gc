# \Delete5GVnGroupApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete5GVnGroup**](Delete5GVnGroupApi.md#Delete5GVnGroup) | **Delete** /subscription-data/group-data/5g-vn-groups/{externalGroupId} | Deletes the 5GVnGroup



## Delete5GVnGroup

> Delete5GVnGroup(ctx, externalGroupId).Execute()

Deletes the 5GVnGroup

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
    resp, r, err := apiClient.Delete5GVnGroupApi.Delete5GVnGroup(context.Background(), externalGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Delete5GVnGroupApi.Delete5GVnGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete5GVnGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


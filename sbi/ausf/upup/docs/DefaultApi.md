# \DefaultApi

All URIs are relative to *https://example.com/nausf-upuprotection/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SupiUeUpuPost**](DefaultApi.md#SupiUeUpuPost) | **Post** /{supi}/ue-upu | 



## SupiUeUpuPost

> UpuSecurityInfo SupiUeUpuPost(ctx, supi).UpuInfo(upuInfo).Execute()



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
    supi := "supi_example" // string | Identifier of the UE
    upuInfo := TODO // UpuInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SupiUeUpuPost(context.Background(), supi).UpuInfo(upuInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SupiUeUpuPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupiUeUpuPost`: UpuSecurityInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SupiUeUpuPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupiUeUpuPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upuInfo** | [**UpuInfo**](UpuInfo.md) |  | 

### Return type

[**UpuSecurityInfo**](UpuSecurityInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


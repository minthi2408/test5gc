# \DefaultApi

All URIs are relative to *https://example.com/nausf-sorprotection/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SupiUeSorPost**](DefaultApi.md#SupiUeSorPost) | **Post** /{supi}/ue-sor | 



## SupiUeSorPost

> SorSecurityInfo SupiUeSorPost(ctx, supi).SorInfo(sorInfo).Execute()



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
    sorInfo := TODO // SorInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SupiUeSorPost(context.Background(), supi).SorInfo(sorInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SupiUeSorPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupiUeSorPost`: SorSecurityInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SupiUeSorPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupiUeSorPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sorInfo** | [**SorInfo**](SorInfo.md) |  | 

### Return type

[**SorSecurityInfo**](SorSecurityInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


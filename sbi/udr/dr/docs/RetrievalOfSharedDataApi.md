# \RetrievalOfSharedDataApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSharedData**](RetrievalOfSharedDataApi.md#GetSharedData) | **Get** /subscription-data/shared-data | retrieve shared data



## GetSharedData

> []SharedData GetSharedData(ctx).SharedDataIds(sharedDataIds).SupportedFeatures(supportedFeatures).Execute()

retrieve shared data

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
    sharedDataIds := []string{"Inner_example"} // []string | List of shared data ids
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrievalOfSharedDataApi.GetSharedData(context.Background()).SharedDataIds(sharedDataIds).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrievalOfSharedDataApi.GetSharedData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSharedData`: []SharedData
    fmt.Fprintf(os.Stdout, "Response from `RetrievalOfSharedDataApi.GetSharedData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sharedDataIds** | **[]string** | List of shared data ids | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**[]SharedData**](SharedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


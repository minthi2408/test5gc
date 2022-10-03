# \EventExposureDataForAGroupDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryGroupEEData**](EventExposureDataForAGroupDocumentApi.md#QueryGroupEEData) | **Get** /subscription-data/group-data/{ueGroupId}/ee-profile-data | Retrieves the ee profile data profile data of a group or anyUE



## QueryGroupEEData

> EeGroupProfileData QueryGroupEEData(ctx, ueGroupId).SupportedFeatures(supportedFeatures).Execute()

Retrieves the ee profile data profile data of a group or anyUE

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
    ueGroupId := "ueGroupId_example" // string | Group of UEs or any UE
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventExposureDataForAGroupDocumentApi.QueryGroupEEData(context.Background(), ueGroupId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventExposureDataForAGroupDocumentApi.QueryGroupEEData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryGroupEEData`: EeGroupProfileData
    fmt.Fprintf(os.Stdout, "Response from `EventExposureDataForAGroupDocumentApi.QueryGroupEEData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** | Group of UEs or any UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryGroupEEDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**EeGroupProfileData**](EeGroupProfileData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


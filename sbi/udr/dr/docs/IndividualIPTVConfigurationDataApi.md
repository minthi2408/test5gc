# \IndividualIPTVConfigurationDataApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartialReplaceIndividualIPTVConfigurationData**](IndividualIPTVConfigurationDataApi.md#PartialReplaceIndividualIPTVConfigurationData) | **Patch** /application-data/iptvConfigData/{configurationId} | Partial update an individual IPTV configuration resource



## PartialReplaceIndividualIPTVConfigurationData

> IptvConfigData PartialReplaceIndividualIPTVConfigurationData(ctx, configurationId).IptvConfigDataPatch(iptvConfigDataPatch).Execute()

Partial update an individual IPTV configuration resource

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
    configurationId := "configurationId_example" // string | The Identifier of an Individual IPTV Configuration Data to be updated. It shall apply the format of Data type string.
    iptvConfigDataPatch := TODO // IptvConfigDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualIPTVConfigurationDataApi.PartialReplaceIndividualIPTVConfigurationData(context.Background(), configurationId).IptvConfigDataPatch(iptvConfigDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualIPTVConfigurationDataApi.PartialReplaceIndividualIPTVConfigurationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialReplaceIndividualIPTVConfigurationData`: IptvConfigData
    fmt.Fprintf(os.Stdout, "Response from `IndividualIPTVConfigurationDataApi.PartialReplaceIndividualIPTVConfigurationData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | The Identifier of an Individual IPTV Configuration Data to be updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialReplaceIndividualIPTVConfigurationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iptvConfigDataPatch** | [**IptvConfigDataPatch**](IptvConfigDataPatch.md) |  | 

### Return type

[**IptvConfigData**](IptvConfigData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


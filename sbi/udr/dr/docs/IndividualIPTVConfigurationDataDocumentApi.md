# \IndividualIPTVConfigurationDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceIndividualIPTVConfigurationData**](IndividualIPTVConfigurationDataDocumentApi.md#CreateOrReplaceIndividualIPTVConfigurationData) | **Put** /application-data/iptvConfigData/{configurationId} | Create or update an individual IPTV configuration resource
[**DeleteIndividualIPTVConfigurationData**](IndividualIPTVConfigurationDataDocumentApi.md#DeleteIndividualIPTVConfigurationData) | **Delete** /application-data/iptvConfigData/{configurationId} | Delete an individual IPTV configuration resource



## CreateOrReplaceIndividualIPTVConfigurationData

> IptvConfigData CreateOrReplaceIndividualIPTVConfigurationData(ctx, configurationId).IptvConfigData(iptvConfigData).Execute()

Create or update an individual IPTV configuration resource

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
    configurationId := "configurationId_example" // string | The Identifier of an Individual IPTV Configuration Data to be created or updated. It shall apply the format of Data type string.
    iptvConfigData := TODO // IptvConfigData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualIPTVConfigurationDataDocumentApi.CreateOrReplaceIndividualIPTVConfigurationData(context.Background(), configurationId).IptvConfigData(iptvConfigData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualIPTVConfigurationDataDocumentApi.CreateOrReplaceIndividualIPTVConfigurationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceIndividualIPTVConfigurationData`: IptvConfigData
    fmt.Fprintf(os.Stdout, "Response from `IndividualIPTVConfigurationDataDocumentApi.CreateOrReplaceIndividualIPTVConfigurationData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | The Identifier of an Individual IPTV Configuration Data to be created or updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceIndividualIPTVConfigurationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iptvConfigData** | [**IptvConfigData**](IptvConfigData.md) |  | 

### Return type

[**IptvConfigData**](IptvConfigData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndividualIPTVConfigurationData

> DeleteIndividualIPTVConfigurationData(ctx, configurationId).Execute()

Delete an individual IPTV configuration resource

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
    configurationId := "configurationId_example" // string | The Identifier of an Individual IPTV Configuration to be updated. It shall apply the format of Data type string.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualIPTVConfigurationDataDocumentApi.DeleteIndividualIPTVConfigurationData(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualIPTVConfigurationDataDocumentApi.DeleteIndividualIPTVConfigurationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | The Identifier of an Individual IPTV Configuration to be updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualIPTVConfigurationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


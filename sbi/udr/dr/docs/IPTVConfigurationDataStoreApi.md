# \IPTVConfigurationDataStoreApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadIPTVCongifurationData**](IPTVConfigurationDataStoreApi.md#ReadIPTVCongifurationData) | **Get** /application-data/iptvConfigData | Retrieve IPTV configuration Data



## ReadIPTVCongifurationData

> []IptvConfigData ReadIPTVCongifurationData(ctx).ConfigIds(configIds).Dnns(dnns).Snssais(snssais).Supis(supis).InterGroupIds(interGroupIds).Execute()

Retrieve IPTV configuration Data

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
    configIds := []string{"Inner_example"} // []string | Each element identifies a configuration. (optional)
    dnns := []string{"Inner_example"} // []string | Each element identifies a DNN. (optional)
    snssais := []Snssai{"TODO"} // []Snssai | Each element identifies a slice. (optional)
    supis := []string{"Inner_example"} // []string | Each element identifies the user. (optional)
    interGroupIds := []string{"Inner_example"} // []string | Each element identifies a group of users. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPTVConfigurationDataStoreApi.ReadIPTVCongifurationData(context.Background()).ConfigIds(configIds).Dnns(dnns).Snssais(snssais).Supis(supis).InterGroupIds(interGroupIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPTVConfigurationDataStoreApi.ReadIPTVCongifurationData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIPTVCongifurationData`: []IptvConfigData
    fmt.Fprintf(os.Stdout, "Response from `IPTVConfigurationDataStoreApi.ReadIPTVCongifurationData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadIPTVCongifurationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configIds** | **[]string** | Each element identifies a configuration. | 
 **dnns** | **[]string** | Each element identifies a DNN. | 
 **snssais** | [**[]Snssai**](Snssai.md) | Each element identifies a slice. | 
 **supis** | **[]string** | Each element identifies the user. | 
 **interGroupIds** | **[]string** | Each element identifies a group of users. | 

### Return type

[**[]IptvConfigData**](IptvConfigData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


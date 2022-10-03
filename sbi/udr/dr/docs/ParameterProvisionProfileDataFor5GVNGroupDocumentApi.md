# \ParameterProvisionProfileDataFor5GVNGroupDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Query5GVNGroupPPData**](ParameterProvisionProfileDataFor5GVNGroupDocumentApi.md#Query5GVNGroupPPData) | **Get** /subscription-data/group-data/5g-vn-groups/pp-profile-data | Retrieves the parameter provision profile data for 5G VN Group



## Query5GVNGroupPPData

> Pp5gVnGroupProfileData Query5GVNGroupPPData(ctx).ExtGroupIds(extGroupIds).SupportedFeatures(supportedFeatures).Execute()

Retrieves the parameter provision profile data for 5G VN Group

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
    extGroupIds := []string{"Inner_example"} // []string | List of external VN group identifiers (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParameterProvisionProfileDataFor5GVNGroupDocumentApi.Query5GVNGroupPPData(context.Background()).ExtGroupIds(extGroupIds).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterProvisionProfileDataFor5GVNGroupDocumentApi.Query5GVNGroupPPData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Query5GVNGroupPPData`: Pp5gVnGroupProfileData
    fmt.Fprintf(os.Stdout, "Response from `ParameterProvisionProfileDataFor5GVNGroupDocumentApi.Query5GVNGroupPPData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuery5GVNGroupPPDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extGroupIds** | **[]string** | List of external VN group identifiers | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**Pp5gVnGroupProfileData**](Pp5gVnGroupProfileData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


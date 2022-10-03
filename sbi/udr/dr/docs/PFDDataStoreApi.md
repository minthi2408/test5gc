# \PFDDataStoreApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadPFDData**](PFDDataStoreApi.md#ReadPFDData) | **Get** /application-data/pfds | Retrieve PFDs for application identifier(s)



## ReadPFDData

> []PfdDataForAppExt ReadPFDData(ctx).AppId(appId).Execute()

Retrieve PFDs for application identifier(s)

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
    appId := []string{"Inner_example"} // []string | Contains the information of the application identifier(s) for the querying PFD Data resource. If none appId is included in the URI, it applies to all application identifier(s) for the querying PFD Data resource. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PFDDataStoreApi.ReadPFDData(context.Background()).AppId(appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PFDDataStoreApi.ReadPFDData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPFDData`: []PfdDataForAppExt
    fmt.Fprintf(os.Stdout, "Response from `PFDDataStoreApi.ReadPFDData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadPFDDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **[]string** | Contains the information of the application identifier(s) for the querying PFD Data resource. If none appId is included in the URI, it applies to all application identifier(s) for the querying PFD Data resource. | 

### Return type

[**[]PfdDataForAppExt**](PfdDataForAppExt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \BdtDataStoreApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadBdtData**](BdtDataStoreApi.md#ReadBdtData) | **Get** /policy-data/bdt-data | Retrieves the BDT data collection



## ReadBdtData

> []BdtData ReadBdtData(ctx).BdtRefIds(bdtRefIds).SuppFeat(suppFeat).Execute()

Retrieves the BDT data collection

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
    bdtRefIds := []string{"Inner_example"} // []string | List of the BDT reference identifiers. (optional)
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BdtDataStoreApi.ReadBdtData(context.Background()).BdtRefIds(bdtRefIds).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BdtDataStoreApi.ReadBdtData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadBdtData`: []BdtData
    fmt.Fprintf(os.Stdout, "Response from `BdtDataStoreApi.ReadBdtData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadBdtDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bdtRefIds** | **[]string** | List of the BDT reference identifiers. | 
 **suppFeat** | **string** | Supported Features | 

### Return type

[**[]BdtData**](BdtData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


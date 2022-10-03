# \BdtPolicyDataStoreApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadBdtPolicyData**](BdtPolicyDataStoreApi.md#ReadBdtPolicyData) | **Get** /application-data/bdtPolicyData | Retrieve applied BDT Policy Data



## ReadBdtPolicyData

> []BdtPolicyData ReadBdtPolicyData(ctx).BdtPolicyIds(bdtPolicyIds).InternalGroupIds(internalGroupIds).Supis(supis).Execute()

Retrieve applied BDT Policy Data

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
    bdtPolicyIds := []string{"Inner_example"} // []string | Each element identifies a service. (optional)
    internalGroupIds := []string{"Inner_example"} // []string | Each element identifies a group of users. (optional)
    supis := []string{"Inner_example"} // []string | Each element identifies the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BdtPolicyDataStoreApi.ReadBdtPolicyData(context.Background()).BdtPolicyIds(bdtPolicyIds).InternalGroupIds(internalGroupIds).Supis(supis).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BdtPolicyDataStoreApi.ReadBdtPolicyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadBdtPolicyData`: []BdtPolicyData
    fmt.Fprintf(os.Stdout, "Response from `BdtPolicyDataStoreApi.ReadBdtPolicyData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadBdtPolicyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bdtPolicyIds** | **[]string** | Each element identifies a service. | 
 **internalGroupIds** | **[]string** | Each element identifies a group of users. | 
 **supis** | **[]string** | Each element identifies the user. | 

### Return type

[**[]BdtPolicyData**](BdtPolicyData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


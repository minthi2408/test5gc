# \SessionManagementPolicyDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadSessionManagementPolicyData**](SessionManagementPolicyDataDocumentApi.md#ReadSessionManagementPolicyData) | **Get** /policy-data/ues/{ueId}/sm-data | Retrieves the session management policy data for a subscriber
[**UpdateSessionManagementPolicyData**](SessionManagementPolicyDataDocumentApi.md#UpdateSessionManagementPolicyData) | **Patch** /policy-data/ues/{ueId}/sm-data | Modify the session management policy data for a subscriber



## ReadSessionManagementPolicyData

> SmPolicyData ReadSessionManagementPolicyData(ctx, ueId).Snssai(snssai).Dnn(dnn).Fields(fields).SuppFeat(suppFeat).Execute()

Retrieves the session management policy data for a subscriber

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
    ueId := "ueId_example" // string | 
    snssai := Snssai{ ... } // Snssai |  (optional)
    dnn := "dnn_example" // string |  (optional)
    fields := []string{"Inner_example"} // []string | attributes to be retrieved (optional)
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionManagementPolicyDataDocumentApi.ReadSessionManagementPolicyData(context.Background(), ueId).Snssai(snssai).Dnn(dnn).Fields(fields).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementPolicyDataDocumentApi.ReadSessionManagementPolicyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSessionManagementPolicyData`: SmPolicyData
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementPolicyDataDocumentApi.ReadSessionManagementPolicyData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSessionManagementPolicyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snssai** | [**Snssai**](Snssai.md) |  | 
 **dnn** | **string** |  | 
 **fields** | **[]string** | attributes to be retrieved | 
 **suppFeat** | **string** | Supported Features | 

### Return type

[**SmPolicyData**](SmPolicyData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSessionManagementPolicyData

> SmPolicyData UpdateSessionManagementPolicyData(ctx, ueId).SmPolicyDataPatch(smPolicyDataPatch).Execute()

Modify the session management policy data for a subscriber

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
    ueId := "ueId_example" // string | 
    smPolicyDataPatch := TODO // SmPolicyDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionManagementPolicyDataDocumentApi.UpdateSessionManagementPolicyData(context.Background(), ueId).SmPolicyDataPatch(smPolicyDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementPolicyDataDocumentApi.UpdateSessionManagementPolicyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionManagementPolicyData`: SmPolicyData
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementPolicyDataDocumentApi.UpdateSessionManagementPolicyData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionManagementPolicyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smPolicyDataPatch** | [**SmPolicyDataPatch**](SmPolicyDataPatch.md) |  | 

### Return type

[**SmPolicyData**](SmPolicyData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


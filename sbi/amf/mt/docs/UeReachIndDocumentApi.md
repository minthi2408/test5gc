# \UeReachIndDocumentApi

All URIs are relative to *https://example.com/namf-mt/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableUeReachability**](UeReachIndDocumentApi.md#EnableUeReachability) | **Put** /ue-contexts/{ueContextId}/ue-reachind | Namf_MT EnableUEReachability service Operation



## EnableUeReachability

> EnableUeReachabilityRspData EnableUeReachability(ctx, ueContextId).EnableUeReachabilityReqData(enableUeReachabilityReqData).Execute()

Namf_MT EnableUEReachability service Operation

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
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    enableUeReachabilityReqData := TODO // EnableUeReachabilityReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UeReachIndDocumentApi.EnableUeReachability(context.Background(), ueContextId).EnableUeReachabilityReqData(enableUeReachabilityReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UeReachIndDocumentApi.EnableUeReachability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableUeReachability`: EnableUeReachabilityRspData
    fmt.Fprintf(os.Stdout, "Response from `UeReachIndDocumentApi.EnableUeReachability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableUeReachabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableUeReachabilityReqData** | [**EnableUeReachabilityReqData**](EnableUeReachabilityReqData.md) |  | 

### Return type

[**EnableUeReachabilityRspData**](EnableUeReachabilityRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


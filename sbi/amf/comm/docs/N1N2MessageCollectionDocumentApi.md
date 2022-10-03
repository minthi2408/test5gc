# \N1N2MessageCollectionDocumentApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**N1N2MessageTransfer**](N1N2MessageCollectionDocumentApi.md#N1N2MessageTransfer) | **Post** /ue-contexts/{ueContextId}/n1-n2-messages | Namf_Communication N1N2 Message Transfer (UE Specific) service Operation



## N1N2MessageTransfer

> N1N2MessageTransferRspData N1N2MessageTransfer(ctx, ueContextId).N1N2MessageTransferReqData(n1N2MessageTransferReqData).Execute()

Namf_Communication N1N2 Message Transfer (UE Specific) service Operation

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
    n1N2MessageTransferReqData := TODO // N1N2MessageTransferReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N1N2MessageCollectionDocumentApi.N1N2MessageTransfer(context.Background(), ueContextId).N1N2MessageTransferReqData(n1N2MessageTransferReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N1N2MessageCollectionDocumentApi.N1N2MessageTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `N1N2MessageTransfer`: N1N2MessageTransferRspData
    fmt.Fprintf(os.Stdout, "Response from `N1N2MessageCollectionDocumentApi.N1N2MessageTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiN1N2MessageTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n1N2MessageTransferReqData** | [**N1N2MessageTransferReqData**](N1N2MessageTransferReqData.md) |  | 

### Return type

[**N1N2MessageTransferRspData**](N1N2MessageTransferRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


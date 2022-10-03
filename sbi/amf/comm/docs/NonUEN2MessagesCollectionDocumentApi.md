# \NonUEN2MessagesCollectionDocumentApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NonUeN2MessageTransfer**](NonUEN2MessagesCollectionDocumentApi.md#NonUeN2MessageTransfer) | **Post** /non-ue-n2-messages/transfer | Namf_Communication Non UE N2 Message Transfer service Operation



## NonUeN2MessageTransfer

> N2InformationTransferRspData NonUeN2MessageTransfer(ctx).N2InformationTransferReqData(n2InformationTransferReqData).Execute()

Namf_Communication Non UE N2 Message Transfer service Operation

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
    n2InformationTransferReqData := TODO // N2InformationTransferReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonUEN2MessagesCollectionDocumentApi.NonUeN2MessageTransfer(context.Background()).N2InformationTransferReqData(n2InformationTransferReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonUEN2MessagesCollectionDocumentApi.NonUeN2MessageTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonUeN2MessageTransfer`: N2InformationTransferRspData
    fmt.Fprintf(os.Stdout, "Response from `NonUEN2MessagesCollectionDocumentApi.NonUeN2MessageTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonUeN2MessageTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n2InformationTransferReqData** | [**N2InformationTransferReqData**](N2InformationTransferReqData.md) |  | 

### Return type

[**N2InformationTransferRspData**](N2InformationTransferRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


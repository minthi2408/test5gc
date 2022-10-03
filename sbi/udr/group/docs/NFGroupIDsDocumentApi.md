# \NFGroupIDsDocumentApi

All URIs are relative to *https://example.com/nudr-group-id-map/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNfGroupIDs**](NFGroupIDsDocumentApi.md#GetNfGroupIDs) | **Get** /nf-group-ids | Retrieves NF-Group IDs for provided Subscriber and NF types



## GetNfGroupIDs

> map[string]string GetNfGroupIDs(ctx).NfType(nfType).SubscriberId(subscriberId).Execute()

Retrieves NF-Group IDs for provided Subscriber and NF types

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
    nfType := []NFType{"TODO"} // []NFType | Type of NF
    subscriberId := "subscriberId_example" // string | Identifier of the subscriber

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFGroupIDsDocumentApi.GetNfGroupIDs(context.Background()).NfType(nfType).SubscriberId(subscriberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFGroupIDsDocumentApi.GetNfGroupIDs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNfGroupIDs`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `NFGroupIDsDocumentApi.GetNfGroupIDs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNfGroupIDsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfType** | [**[]NFType**](NFType.md) | Type of NF | 
 **subscriberId** | **string** | Identifier of the subscriber | 

### Return type

**map[string]string**

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


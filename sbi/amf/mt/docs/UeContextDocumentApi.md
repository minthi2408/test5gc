# \UeContextDocumentApi

All URIs are relative to *https://example.com/namf-mt/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvideDomainSelectionInfo**](UeContextDocumentApi.md#ProvideDomainSelectionInfo) | **Get** /ue-contexts/{ueContextId} | Namf_MT Provide Domain Selection Info service Operation



## ProvideDomainSelectionInfo

> UeContextInfo ProvideDomainSelectionInfo(ctx, ueContextId).InfoClass(infoClass).SupportedFeatures(supportedFeatures).OldGuami(oldGuami).Execute()

Namf_MT Provide Domain Selection Info service Operation

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
    infoClass := TODO // UeContextInfoClass | UE Context Information Class (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    oldGuami := Guami{ ... } // Guami | Old GUAMI (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UeContextDocumentApi.ProvideDomainSelectionInfo(context.Background(), ueContextId).InfoClass(infoClass).SupportedFeatures(supportedFeatures).OldGuami(oldGuami).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UeContextDocumentApi.ProvideDomainSelectionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvideDomainSelectionInfo`: UeContextInfo
    fmt.Fprintf(os.Stdout, "Response from `UeContextDocumentApi.ProvideDomainSelectionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvideDomainSelectionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **infoClass** | [**UeContextInfoClass**](UeContextInfoClass.md) | UE Context Information Class | 
 **supportedFeatures** | **string** | Supported Features | 
 **oldGuami** | [**Guami**](Guami.md) | Old GUAMI | 

### Return type

[**UeContextInfo**](UeContextInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


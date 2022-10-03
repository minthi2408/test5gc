# \AMF3GPPAccessRegistrationDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AmfContext3gpp**](AMF3GPPAccessRegistrationDocumentApi.md#AmfContext3gpp) | **Patch** /subscription-data/{ueId}/context-data/amf-3gpp-access | To modify the AMF context data of a UE using 3gpp access in the UDR
[**CreateAmfContext3gpp**](AMF3GPPAccessRegistrationDocumentApi.md#CreateAmfContext3gpp) | **Put** /subscription-data/{ueId}/context-data/amf-3gpp-access | To store the AMF context data of a UE using 3gpp access in the UDR
[**QueryAmfContext3gpp**](AMF3GPPAccessRegistrationDocumentApi.md#QueryAmfContext3gpp) | **Get** /subscription-data/{ueId}/context-data/amf-3gpp-access | Retrieves the AMF context data of a UE using 3gpp access



## AmfContext3gpp

> PatchResult AmfContext3gpp(ctx, ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

To modify the AMF context data of a UE using 3gpp access in the UDR

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
    ueId := "ueId_example" // string | UE id
    patchItem := []PatchItem{"TODO"} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMF3GPPAccessRegistrationDocumentApi.AmfContext3gpp(context.Background(), ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMF3GPPAccessRegistrationDocumentApi.AmfContext3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AmfContext3gpp`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `AMF3GPPAccessRegistrationDocumentApi.AmfContext3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAmfContext3gppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAmfContext3gpp

> Amf3GppAccessRegistration CreateAmfContext3gpp(ctx, ueId).Amf3GppAccessRegistration(amf3GppAccessRegistration).Execute()

To store the AMF context data of a UE using 3gpp access in the UDR

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
    ueId := "ueId_example" // string | UE id
    amf3GppAccessRegistration := TODO // Amf3GppAccessRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMF3GPPAccessRegistrationDocumentApi.CreateAmfContext3gpp(context.Background(), ueId).Amf3GppAccessRegistration(amf3GppAccessRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMF3GPPAccessRegistrationDocumentApi.CreateAmfContext3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAmfContext3gpp`: Amf3GppAccessRegistration
    fmt.Fprintf(os.Stdout, "Response from `AMF3GPPAccessRegistrationDocumentApi.CreateAmfContext3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAmfContext3gppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amf3GppAccessRegistration** | [**Amf3GppAccessRegistration**](Amf3GppAccessRegistration.md) |  | 

### Return type

[**Amf3GppAccessRegistration**](Amf3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryAmfContext3gpp

> Amf3GppAccessRegistration QueryAmfContext3gpp(ctx, ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()

Retrieves the AMF context data of a UE using 3gpp access

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
    ueId := "ueId_example" // string | UE id
    fields := []string{"Inner_example"} // []string | attributes to be retrieved (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMF3GPPAccessRegistrationDocumentApi.QueryAmfContext3gpp(context.Background(), ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMF3GPPAccessRegistrationDocumentApi.QueryAmfContext3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAmfContext3gpp`: Amf3GppAccessRegistration
    fmt.Fprintf(os.Stdout, "Response from `AMF3GPPAccessRegistrationDocumentApi.QueryAmfContext3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryAmfContext3gppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | attributes to be retrieved | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**Amf3GppAccessRegistration**](Amf3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


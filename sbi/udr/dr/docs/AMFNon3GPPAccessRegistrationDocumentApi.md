# \AMFNon3GPPAccessRegistrationDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AmfContextNon3gpp**](AMFNon3GPPAccessRegistrationDocumentApi.md#AmfContextNon3gpp) | **Patch** /subscription-data/{ueId}/context-data/amf-non-3gpp-access | To modify the AMF context data of a UE using non 3gpp access in the UDR
[**CreateAmfContextNon3gpp**](AMFNon3GPPAccessRegistrationDocumentApi.md#CreateAmfContextNon3gpp) | **Put** /subscription-data/{ueId}/context-data/amf-non-3gpp-access | To store the AMF context data of a UE using non-3gpp access in the UDR
[**QueryAmfContextNon3gpp**](AMFNon3GPPAccessRegistrationDocumentApi.md#QueryAmfContextNon3gpp) | **Get** /subscription-data/{ueId}/context-data/amf-non-3gpp-access | Retrieves the AMF context data of a UE using non-3gpp access



## AmfContextNon3gpp

> PatchResult AmfContextNon3gpp(ctx, ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

To modify the AMF context data of a UE using non 3gpp access in the UDR

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
    resp, r, err := apiClient.AMFNon3GPPAccessRegistrationDocumentApi.AmfContextNon3gpp(context.Background(), ueId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMFNon3GPPAccessRegistrationDocumentApi.AmfContextNon3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AmfContextNon3gpp`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `AMFNon3GPPAccessRegistrationDocumentApi.AmfContextNon3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAmfContextNon3gppRequest struct via the builder pattern


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


## CreateAmfContextNon3gpp

> Amf3GppAccessRegistration CreateAmfContextNon3gpp(ctx, ueId).AmfNon3GppAccessRegistration(amfNon3GppAccessRegistration).Execute()

To store the AMF context data of a UE using non-3gpp access in the UDR

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
    amfNon3GppAccessRegistration := TODO // AmfNon3GppAccessRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMFNon3GPPAccessRegistrationDocumentApi.CreateAmfContextNon3gpp(context.Background(), ueId).AmfNon3GppAccessRegistration(amfNon3GppAccessRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMFNon3GPPAccessRegistrationDocumentApi.CreateAmfContextNon3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAmfContextNon3gpp`: Amf3GppAccessRegistration
    fmt.Fprintf(os.Stdout, "Response from `AMFNon3GPPAccessRegistrationDocumentApi.CreateAmfContextNon3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAmfContextNon3gppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amfNon3GppAccessRegistration** | [**AmfNon3GppAccessRegistration**](AmfNon3GppAccessRegistration.md) |  | 

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


## QueryAmfContextNon3gpp

> AmfNon3GppAccessRegistration QueryAmfContextNon3gpp(ctx, ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()

Retrieves the AMF context data of a UE using non-3gpp access

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
    resp, r, err := apiClient.AMFNon3GPPAccessRegistrationDocumentApi.QueryAmfContextNon3gpp(context.Background(), ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMFNon3GPPAccessRegistrationDocumentApi.QueryAmfContextNon3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAmfContextNon3gpp`: AmfNon3GppAccessRegistration
    fmt.Fprintf(os.Stdout, "Response from `AMFNon3GPPAccessRegistrationDocumentApi.QueryAmfContextNon3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryAmfContextNon3gppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | attributes to be retrieved | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**AmfNon3GppAccessRegistration**](AmfNon3GppAccessRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


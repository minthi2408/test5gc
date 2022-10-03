# \SMSFNon3GPPRegistrationDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmsfContextNon3gpp**](SMSFNon3GPPRegistrationDocumentApi.md#CreateSmsfContextNon3gpp) | **Put** /subscription-data/{ueId}/context-data/smsf-non-3gpp-access | Create the SMSF context data of a UE via non-3GPP access
[**DeleteSmsfContextNon3gpp**](SMSFNon3GPPRegistrationDocumentApi.md#DeleteSmsfContextNon3gpp) | **Delete** /subscription-data/{ueId}/context-data/smsf-non-3gpp-access | To remove the SMSF context data of a UE via non-3GPP access
[**QuerySmsfContextNon3gpp**](SMSFNon3GPPRegistrationDocumentApi.md#QuerySmsfContextNon3gpp) | **Get** /subscription-data/{ueId}/context-data/smsf-non-3gpp-access | Retrieves the SMSF context data of a UE using non-3gpp access



## CreateSmsfContextNon3gpp

> SmsfRegistration CreateSmsfContextNon3gpp(ctx, ueId).SmsfRegistration(smsfRegistration).Execute()

Create the SMSF context data of a UE via non-3GPP access

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
    smsfRegistration := TODO // SmsfRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSFNon3GPPRegistrationDocumentApi.CreateSmsfContextNon3gpp(context.Background(), ueId).SmsfRegistration(smsfRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSFNon3GPPRegistrationDocumentApi.CreateSmsfContextNon3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSmsfContextNon3gpp`: SmsfRegistration
    fmt.Fprintf(os.Stdout, "Response from `SMSFNon3GPPRegistrationDocumentApi.CreateSmsfContextNon3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSmsfContextNon3gppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsfRegistration** | [**SmsfRegistration**](SmsfRegistration.md) |  | 

### Return type

[**SmsfRegistration**](SmsfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSmsfContextNon3gpp

> DeleteSmsfContextNon3gpp(ctx, ueId).Execute()

To remove the SMSF context data of a UE via non-3GPP access

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMSFNon3GPPRegistrationDocumentApi.DeleteSmsfContextNon3gpp(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSFNon3GPPRegistrationDocumentApi.DeleteSmsfContextNon3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmsfContextNon3gppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuerySmsfContextNon3gpp

> SmsfRegistration QuerySmsfContextNon3gpp(ctx, ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()

Retrieves the SMSF context data of a UE using non-3gpp access

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
    resp, r, err := apiClient.SMSFNon3GPPRegistrationDocumentApi.QuerySmsfContextNon3gpp(context.Background(), ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSFNon3GPPRegistrationDocumentApi.QuerySmsfContextNon3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySmsfContextNon3gpp`: SmsfRegistration
    fmt.Fprintf(os.Stdout, "Response from `SMSFNon3GPPRegistrationDocumentApi.QuerySmsfContextNon3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySmsfContextNon3gppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | attributes to be retrieved | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**SmsfRegistration**](SmsfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


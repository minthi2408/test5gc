# \SMSF3GPPRegistrationDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmsfContext3gpp**](SMSF3GPPRegistrationDocumentApi.md#CreateSmsfContext3gpp) | **Put** /subscription-data/{ueId}/context-data/smsf-3gpp-access | Create the SMSF context data of a UE via 3GPP access
[**DeleteSmsfContext3gpp**](SMSF3GPPRegistrationDocumentApi.md#DeleteSmsfContext3gpp) | **Delete** /subscription-data/{ueId}/context-data/smsf-3gpp-access | To remove the SMSF context data of a UE via 3GPP access
[**QuerySmsfContext3gpp**](SMSF3GPPRegistrationDocumentApi.md#QuerySmsfContext3gpp) | **Get** /subscription-data/{ueId}/context-data/smsf-3gpp-access | Retrieves the SMSF context data of a UE using 3gpp access



## CreateSmsfContext3gpp

> SmsfRegistration CreateSmsfContext3gpp(ctx, ueId).SmsfRegistration(smsfRegistration).Execute()

Create the SMSF context data of a UE via 3GPP access

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
    resp, r, err := apiClient.SMSF3GPPRegistrationDocumentApi.CreateSmsfContext3gpp(context.Background(), ueId).SmsfRegistration(smsfRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSF3GPPRegistrationDocumentApi.CreateSmsfContext3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSmsfContext3gpp`: SmsfRegistration
    fmt.Fprintf(os.Stdout, "Response from `SMSF3GPPRegistrationDocumentApi.CreateSmsfContext3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSmsfContext3gppRequest struct via the builder pattern


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


## DeleteSmsfContext3gpp

> DeleteSmsfContext3gpp(ctx, ueId).Execute()

To remove the SMSF context data of a UE via 3GPP access

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
    resp, r, err := apiClient.SMSF3GPPRegistrationDocumentApi.DeleteSmsfContext3gpp(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSF3GPPRegistrationDocumentApi.DeleteSmsfContext3gpp``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSmsfContext3gppRequest struct via the builder pattern


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


## QuerySmsfContext3gpp

> SmsfRegistration QuerySmsfContext3gpp(ctx, ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()

Retrieves the SMSF context data of a UE using 3gpp access

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
    resp, r, err := apiClient.SMSF3GPPRegistrationDocumentApi.QuerySmsfContext3gpp(context.Background(), ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMSF3GPPRegistrationDocumentApi.QuerySmsfContext3gpp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySmsfContext3gpp`: SmsfRegistration
    fmt.Fprintf(os.Stdout, "Response from `SMSF3GPPRegistrationDocumentApi.QuerySmsfContext3gpp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySmsfContext3gppRequest struct via the builder pattern


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


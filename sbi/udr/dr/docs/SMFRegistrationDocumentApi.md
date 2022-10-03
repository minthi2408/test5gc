# \SMFRegistrationDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateSmfRegistration**](SMFRegistrationDocumentApi.md#CreateOrUpdateSmfRegistration) | **Put** /subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId} | To create an individual SMF context data of a UE in the UDR
[**DeleteSmfRegistration**](SMFRegistrationDocumentApi.md#DeleteSmfRegistration) | **Delete** /subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId} | To remove an individual SMF context data of a UE the UDR
[**QuerySmfRegistration**](SMFRegistrationDocumentApi.md#QuerySmfRegistration) | **Get** /subscription-data/{ueId}/context-data/smf-registrations/{pduSessionId} | Retrieves the individual SMF registration of a UE



## CreateOrUpdateSmfRegistration

> SmfRegistration CreateOrUpdateSmfRegistration(ctx, ueId, pduSessionId).SmfRegistration(smfRegistration).Execute()

To create an individual SMF context data of a UE in the UDR

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
    pduSessionId := int32(56) // int32 | PDU session id
    smfRegistration := TODO // SmfRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFRegistrationDocumentApi.CreateOrUpdateSmfRegistration(context.Background(), ueId, pduSessionId).SmfRegistration(smfRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFRegistrationDocumentApi.CreateOrUpdateSmfRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateSmfRegistration`: SmfRegistration
    fmt.Fprintf(os.Stdout, "Response from `SMFRegistrationDocumentApi.CreateOrUpdateSmfRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**pduSessionId** | **int32** | PDU session id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateSmfRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smfRegistration** | [**SmfRegistration**](SmfRegistration.md) |  | 

### Return type

[**SmfRegistration**](SmfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSmfRegistration

> DeleteSmfRegistration(ctx, ueId, pduSessionId).Execute()

To remove an individual SMF context data of a UE the UDR

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
    pduSessionId := int32(56) // int32 | PDU session id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFRegistrationDocumentApi.DeleteSmfRegistration(context.Background(), ueId, pduSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFRegistrationDocumentApi.DeleteSmfRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**pduSessionId** | **int32** | PDU session id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmfRegistrationRequest struct via the builder pattern


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


## QuerySmfRegistration

> SmfRegistration QuerySmfRegistration(ctx, ueId, pduSessionId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()

Retrieves the individual SMF registration of a UE

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
    pduSessionId := int32(56) // int32 | PDU session id
    fields := []string{"Inner_example"} // []string | attributes to be retrieved (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFRegistrationDocumentApi.QuerySmfRegistration(context.Background(), ueId, pduSessionId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFRegistrationDocumentApi.QuerySmfRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySmfRegistration`: SmfRegistration
    fmt.Fprintf(os.Stdout, "Response from `SMFRegistrationDocumentApi.QuerySmfRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**pduSessionId** | **int32** | PDU session id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySmfRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | attributes to be retrieved | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**SmfRegistration**](SmfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AuthEventDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthenticationStatus**](AuthEventDocumentApi.md#DeleteAuthenticationStatus) | **Delete** /subscription-data/{ueId}/authentication-data/authentication-status | To remove the Authentication Status of a UE
[**QueryAuthenticationStatus**](AuthEventDocumentApi.md#QueryAuthenticationStatus) | **Get** /subscription-data/{ueId}/authentication-data/authentication-status | Retrieves the Authentication Status of a UE



## DeleteAuthenticationStatus

> DeleteAuthenticationStatus(ctx, ueId).Execute()

To remove the Authentication Status of a UE

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
    resp, r, err := apiClient.AuthEventDocumentApi.DeleteAuthenticationStatus(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthEventDocumentApi.DeleteAuthenticationStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthenticationStatusRequest struct via the builder pattern


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


## QueryAuthenticationStatus

> AuthEvent QueryAuthenticationStatus(ctx, ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()

Retrieves the Authentication Status of a UE

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
    resp, r, err := apiClient.AuthEventDocumentApi.QueryAuthenticationStatus(context.Background(), ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthEventDocumentApi.QueryAuthenticationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAuthenticationStatus`: AuthEvent
    fmt.Fprintf(os.Stdout, "Response from `AuthEventDocumentApi.QueryAuthenticationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryAuthenticationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | attributes to be retrieved | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**AuthEvent**](AuthEvent.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


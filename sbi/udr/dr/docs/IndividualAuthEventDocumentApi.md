# \IndividualAuthEventDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualAuthenticationStatus**](IndividualAuthEventDocumentApi.md#DeleteIndividualAuthenticationStatus) | **Delete** /subscription-data/{ueId}/authentication-data/authentication-status/{servingNetworkName} | To remove the Individual Authentication Status of a UE
[**QueryIndividualAuthenticationStatus**](IndividualAuthEventDocumentApi.md#QueryIndividualAuthenticationStatus) | **Get** /subscription-data/{ueId}/authentication-data/authentication-status/{servingNetworkName} | Retrieves the Individual Authentication Status of a UE



## DeleteIndividualAuthenticationStatus

> DeleteIndividualAuthenticationStatus(ctx, ueId, servingNetworkName).Execute()

To remove the Individual Authentication Status of a UE

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
    servingNetworkName := "servingNetworkName_example" // string | Serving Network Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAuthEventDocumentApi.DeleteIndividualAuthenticationStatus(context.Background(), ueId, servingNetworkName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAuthEventDocumentApi.DeleteIndividualAuthenticationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**servingNetworkName** | **string** | Serving Network Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualAuthenticationStatusRequest struct via the builder pattern


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


## QueryIndividualAuthenticationStatus

> AuthEvent QueryIndividualAuthenticationStatus(ctx, ueId, servingNetworkName).Fields(fields).SupportedFeatures(supportedFeatures).Execute()

Retrieves the Individual Authentication Status of a UE

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
    servingNetworkName := "servingNetworkName_example" // string | Serving Network Name
    fields := []string{"Inner_example"} // []string | attributes to be retrieved (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAuthEventDocumentApi.QueryIndividualAuthenticationStatus(context.Background(), ueId, servingNetworkName).Fields(fields).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAuthEventDocumentApi.QueryIndividualAuthenticationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryIndividualAuthenticationStatus`: AuthEvent
    fmt.Fprintf(os.Stdout, "Response from `IndividualAuthEventDocumentApi.QueryIndividualAuthenticationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**servingNetworkName** | **string** | Serving Network Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryIndividualAuthenticationStatusRequest struct via the builder pattern


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


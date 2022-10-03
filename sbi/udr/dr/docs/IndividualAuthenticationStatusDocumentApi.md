# \IndividualAuthenticationStatusDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualAuthenticationStatus**](IndividualAuthenticationStatusDocumentApi.md#CreateIndividualAuthenticationStatus) | **Put** /subscription-data/{ueId}/authentication-data/authentication-status/{servingNetworkName} | To store the individual Authentication Status data of a UE



## CreateIndividualAuthenticationStatus

> CreateIndividualAuthenticationStatus(ctx, ueId, servingNetworkName).AuthEvent(authEvent).Execute()

To store the individual Authentication Status data of a UE

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
    authEvent := TODO // AuthEvent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAuthenticationStatusDocumentApi.CreateIndividualAuthenticationStatus(context.Background(), ueId, servingNetworkName).AuthEvent(authEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAuthenticationStatusDocumentApi.CreateIndividualAuthenticationStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateIndividualAuthenticationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authEvent** | [**AuthEvent**](AuthEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


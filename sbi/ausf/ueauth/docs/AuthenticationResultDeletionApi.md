# \AuthenticationResultDeletionApi

All URIs are relative to *https://example.com/nausf-auth/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete5gAkaAuthenticationResult**](AuthenticationResultDeletionApi.md#Delete5gAkaAuthenticationResult) | **Delete** /ue-authentications/{authCtxId}/5g-aka-confirmation | Deletes the authentication result in the UDM
[**DeleteEapAuthenticationResult**](AuthenticationResultDeletionApi.md#DeleteEapAuthenticationResult) | **Delete** /ue-authentications/{authCtxId}/eap-session | Deletes the authentication result in the UDM



## Delete5gAkaAuthenticationResult

> Delete5gAkaAuthenticationResult(ctx, authCtxId).Execute()

Deletes the authentication result in the UDM

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
    authCtxId := "authCtxId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationResultDeletionApi.Delete5gAkaAuthenticationResult(context.Background(), authCtxId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationResultDeletionApi.Delete5gAkaAuthenticationResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCtxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete5gAkaAuthenticationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEapAuthenticationResult

> DeleteEapAuthenticationResult(ctx, authCtxId).Execute()

Deletes the authentication result in the UDM

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
    authCtxId := "authCtxId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationResultDeletionApi.DeleteEapAuthenticationResult(context.Background(), authCtxId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationResultDeletionApi.DeleteEapAuthenticationResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCtxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEapAuthenticationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


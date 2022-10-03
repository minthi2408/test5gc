# \UEPolicySetDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceUEPolicySet**](UEPolicySetDocumentApi.md#CreateOrReplaceUEPolicySet) | **Put** /policy-data/ues/{ueId}/ue-policy-set | Create or modify the UE policy set data for a subscriber
[**ReadUEPolicySet**](UEPolicySetDocumentApi.md#ReadUEPolicySet) | **Get** /policy-data/ues/{ueId}/ue-policy-set | Retrieves the UE policy set data for a subscriber
[**UpdateUEPolicySet**](UEPolicySetDocumentApi.md#UpdateUEPolicySet) | **Patch** /policy-data/ues/{ueId}/ue-policy-set | Modify the UE policy set data for a subscriber



## CreateOrReplaceUEPolicySet

> UePolicySet CreateOrReplaceUEPolicySet(ctx, ueId).UePolicySet(uePolicySet).Execute()

Create or modify the UE policy set data for a subscriber

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
    ueId := "ueId_example" // string | 
    uePolicySet := TODO // UePolicySet | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UEPolicySetDocumentApi.CreateOrReplaceUEPolicySet(context.Background(), ueId).UePolicySet(uePolicySet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UEPolicySetDocumentApi.CreateOrReplaceUEPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceUEPolicySet`: UePolicySet
    fmt.Fprintf(os.Stdout, "Response from `UEPolicySetDocumentApi.CreateOrReplaceUEPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceUEPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uePolicySet** | [**UePolicySet**](UePolicySet.md) |  | 

### Return type

[**UePolicySet**](UePolicySet.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUEPolicySet

> UePolicySet ReadUEPolicySet(ctx, ueId).SuppFeat(suppFeat).Execute()

Retrieves the UE policy set data for a subscriber

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
    ueId := "ueId_example" // string | 
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UEPolicySetDocumentApi.ReadUEPolicySet(context.Background(), ueId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UEPolicySetDocumentApi.ReadUEPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUEPolicySet`: UePolicySet
    fmt.Fprintf(os.Stdout, "Response from `UEPolicySetDocumentApi.ReadUEPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUEPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppFeat** | **string** | Supported Features | 

### Return type

[**UePolicySet**](UePolicySet.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUEPolicySet

> UpdateUEPolicySet(ctx, ueId).UePolicySetPatch(uePolicySetPatch).Execute()

Modify the UE policy set data for a subscriber

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
    ueId := "ueId_example" // string | 
    uePolicySetPatch := TODO // UePolicySetPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UEPolicySetDocumentApi.UpdateUEPolicySet(context.Background(), ueId).UePolicySetPatch(uePolicySetPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UEPolicySetDocumentApi.UpdateUEPolicySet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUEPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uePolicySetPatch** | [**UePolicySetPatch**](UePolicySetPatch.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


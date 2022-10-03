# \IndividualUEContextDocumentApi

All URIs are relative to *https://example.com/namf-loc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelLocation**](IndividualUEContextDocumentApi.md#CancelLocation) | **Post** /{ueContextId}/cancel-pos-info | Namf_Location CancelLocation service operation
[**ProvideLocationInfo**](IndividualUEContextDocumentApi.md#ProvideLocationInfo) | **Post** /{ueContextId}/provide-loc-info | Namf_Location ProvideLocationInfo service Operation
[**ProvidePositioningInfo**](IndividualUEContextDocumentApi.md#ProvidePositioningInfo) | **Post** /{ueContextId}/provide-pos-info | Namf_Location ProvidePositioningInfo service Operation



## CancelLocation

> CancelLocation(ctx, ueContextId).CancelPosInfo(cancelPosInfo).Execute()

Namf_Location CancelLocation service operation

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
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    cancelPosInfo := TODO // CancelPosInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUEContextDocumentApi.CancelLocation(context.Background(), ueContextId).CancelPosInfo(cancelPosInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUEContextDocumentApi.CancelLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelPosInfo** | [**CancelPosInfo**](CancelPosInfo.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvideLocationInfo

> ProvideLocInfo ProvideLocationInfo(ctx, ueContextId).RequestLocInfo(requestLocInfo).Execute()

Namf_Location ProvideLocationInfo service Operation

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
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    requestLocInfo := TODO // RequestLocInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUEContextDocumentApi.ProvideLocationInfo(context.Background(), ueContextId).RequestLocInfo(requestLocInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUEContextDocumentApi.ProvideLocationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvideLocationInfo`: ProvideLocInfo
    fmt.Fprintf(os.Stdout, "Response from `IndividualUEContextDocumentApi.ProvideLocationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvideLocationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestLocInfo** | [**RequestLocInfo**](RequestLocInfo.md) |  | 

### Return type

[**ProvideLocInfo**](ProvideLocInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidePositioningInfo

> ProvidePosInfo ProvidePositioningInfo(ctx, ueContextId).RequestPosInfo(requestPosInfo).Execute()

Namf_Location ProvidePositioningInfo service Operation

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
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    requestPosInfo := TODO // RequestPosInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUEContextDocumentApi.ProvidePositioningInfo(context.Background(), ueContextId).RequestPosInfo(requestPosInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUEContextDocumentApi.ProvidePositioningInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidePositioningInfo`: ProvidePosInfo
    fmt.Fprintf(os.Stdout, "Response from `IndividualUEContextDocumentApi.ProvidePositioningInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidePositioningInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestPosInfo** | [**RequestPosInfo**](RequestPosInfo.md) |  | 

### Return type

[**ProvidePosInfo**](ProvidePosInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


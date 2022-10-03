# \IndividualUeContextDocumentApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRelocateUEContext**](IndividualUeContextDocumentApi.md#CancelRelocateUEContext) | **Post** /ue-contexts/{ueContextId}/cancel-relocate | Namf_Communication CancelRelocateUEContext service Operation
[**CreateUEContext**](IndividualUeContextDocumentApi.md#CreateUEContext) | **Put** /ue-contexts/{ueContextId} | Namf_Communication CreateUEContext service Operation
[**EBIAssignment**](IndividualUeContextDocumentApi.md#EBIAssignment) | **Post** /ue-contexts/{ueContextId}/assign-ebi | Namf_Communication EBI Assignment service Operation
[**RegistrationStatusUpdate**](IndividualUeContextDocumentApi.md#RegistrationStatusUpdate) | **Post** /ue-contexts/{ueContextId}/transfer-update | Namf_Communication RegistrationStatusUpdate service Operation
[**ReleaseUEContext**](IndividualUeContextDocumentApi.md#ReleaseUEContext) | **Post** /ue-contexts/{ueContextId}/release | Namf_Communication ReleaseUEContext service Operation
[**RelocateUEContext**](IndividualUeContextDocumentApi.md#RelocateUEContext) | **Post** /ue-contexts/{ueContextId}/relocate | Namf_Communication RelocateUEContext service Operation
[**UEContextTransfer**](IndividualUeContextDocumentApi.md#UEContextTransfer) | **Post** /ue-contexts/{ueContextId}/transfer | Namf_Communication UEContextTransfer service Operation



## CancelRelocateUEContext

> CancelRelocateUEContext(ctx, ueContextId).CancelRelocateUEContextRequest(cancelRelocateUEContextRequest).Execute()

Namf_Communication CancelRelocateUEContext service Operation

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
    cancelRelocateUEContextRequest := TODO // CancelRelocateUEContextRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.CancelRelocateUEContext(context.Background(), ueContextId).CancelRelocateUEContextRequest(cancelRelocateUEContextRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.CancelRelocateUEContext``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCancelRelocateUEContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelRelocateUEContextRequest** | [**CancelRelocateUEContextRequest**](CancelRelocateUEContextRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUEContext

> UeContextCreatedData CreateUEContext(ctx, ueContextId).CreateUEContextRequest(createUEContextRequest).Execute()

Namf_Communication CreateUEContext service Operation

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
    createUEContextRequest := TODO // CreateUEContextRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.CreateUEContext(context.Background(), ueContextId).CreateUEContextRequest(createUEContextRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.CreateUEContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUEContext`: UeContextCreatedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.CreateUEContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUEContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUEContextRequest** | [**CreateUEContextRequest**](CreateUEContextRequest.md) |  | 

### Return type

[**UeContextCreatedData**](UeContextCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EBIAssignment

> AssignedEbiData EBIAssignment(ctx, ueContextId).AssignEbiData(assignEbiData).Execute()

Namf_Communication EBI Assignment service Operation

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
    assignEbiData := TODO // AssignEbiData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.EBIAssignment(context.Background(), ueContextId).AssignEbiData(assignEbiData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.EBIAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EBIAssignment`: AssignedEbiData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.EBIAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEBIAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignEbiData** | [**AssignEbiData**](AssignEbiData.md) |  | 

### Return type

[**AssignedEbiData**](AssignedEbiData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationStatusUpdate

> UeRegStatusUpdateRspData RegistrationStatusUpdate(ctx, ueContextId).UeRegStatusUpdateReqData(ueRegStatusUpdateReqData).Execute()

Namf_Communication RegistrationStatusUpdate service Operation

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
    ueRegStatusUpdateReqData := TODO // UeRegStatusUpdateReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.RegistrationStatusUpdate(context.Background(), ueContextId).UeRegStatusUpdateReqData(ueRegStatusUpdateReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.RegistrationStatusUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationStatusUpdate`: UeRegStatusUpdateRspData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.RegistrationStatusUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationStatusUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ueRegStatusUpdateReqData** | [**UeRegStatusUpdateReqData**](UeRegStatusUpdateReqData.md) |  | 

### Return type

[**UeRegStatusUpdateRspData**](UeRegStatusUpdateRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseUEContext

> ReleaseUEContext(ctx, ueContextId).UEContextRelease(uEContextRelease).Execute()

Namf_Communication ReleaseUEContext service Operation

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
    uEContextRelease := TODO // UEContextRelease | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.ReleaseUEContext(context.Background(), ueContextId).UEContextRelease(uEContextRelease).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.ReleaseUEContext``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReleaseUEContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uEContextRelease** | [**UEContextRelease**](UEContextRelease.md) |  | 

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


## RelocateUEContext

> UeContextRelocatedData RelocateUEContext(ctx, ueContextId).RelocateUEContextRequest(relocateUEContextRequest).Execute()

Namf_Communication RelocateUEContext service Operation

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
    relocateUEContextRequest := TODO // RelocateUEContextRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.RelocateUEContext(context.Background(), ueContextId).RelocateUEContextRequest(relocateUEContextRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.RelocateUEContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RelocateUEContext`: UeContextRelocatedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.RelocateUEContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelocateUEContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relocateUEContextRequest** | [**RelocateUEContextRequest**](RelocateUEContextRequest.md) |  | 

### Return type

[**UeContextRelocatedData**](UeContextRelocatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UEContextTransfer

> UeContextTransferRspData UEContextTransfer(ctx, ueContextId).UeContextTransferReqData(ueContextTransferReqData).Execute()

Namf_Communication UEContextTransfer service Operation

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
    ueContextTransferReqData := TODO // UeContextTransferReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.UEContextTransfer(context.Background(), ueContextId).UeContextTransferReqData(ueContextTransferReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.UEContextTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UEContextTransfer`: UeContextTransferRspData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.UEContextTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUEContextTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ueContextTransferReqData** | [**UeContextTransferReqData**](UeContextTransferReqData.md) |  | 

### Return type

[**UeContextTransferRspData**](UeContextTransferRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


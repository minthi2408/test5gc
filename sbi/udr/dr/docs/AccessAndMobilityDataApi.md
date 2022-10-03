# \AccessAndMobilityDataApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceAccessAndMobilityData**](AccessAndMobilityDataApi.md#CreateOrReplaceAccessAndMobilityData) | **Put** /exposure-data/{ueId}/access-and-mobility-data | Creates and updates the access and mobility exposure data for a UE
[**DeleteAccessAndMobilityData**](AccessAndMobilityDataApi.md#DeleteAccessAndMobilityData) | **Delete** /exposure-data/{ueId}/access-and-mobility-data | Deletes the access and mobility exposure data for a UE
[**QueryAccessAndMobilityData**](AccessAndMobilityDataApi.md#QueryAccessAndMobilityData) | **Get** /exposure-data/{ueId}/access-and-mobility-data | Retrieves the access and mobility exposure data for a UE
[**UpdateAccessAndMobilityData**](AccessAndMobilityDataApi.md#UpdateAccessAndMobilityData) | **Patch** /exposure-data/{ueId}/access-and-mobility-data | Updates the access and mobility exposure data for a UE



## CreateOrReplaceAccessAndMobilityData

> AccessAndMobilityData CreateOrReplaceAccessAndMobilityData(ctx, ueId).AccessAndMobilityData(accessAndMobilityData).Execute()

Creates and updates the access and mobility exposure data for a UE

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
    accessAndMobilityData := TODO // AccessAndMobilityData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessAndMobilityDataApi.CreateOrReplaceAccessAndMobilityData(context.Background(), ueId).AccessAndMobilityData(accessAndMobilityData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessAndMobilityDataApi.CreateOrReplaceAccessAndMobilityData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceAccessAndMobilityData`: AccessAndMobilityData
    fmt.Fprintf(os.Stdout, "Response from `AccessAndMobilityDataApi.CreateOrReplaceAccessAndMobilityData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceAccessAndMobilityDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessAndMobilityData** | [**AccessAndMobilityData**](AccessAndMobilityData.md) |  | 

### Return type

[**AccessAndMobilityData**](AccessAndMobilityData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessAndMobilityData

> DeleteAccessAndMobilityData(ctx, ueId).Execute()

Deletes the access and mobility exposure data for a UE

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
    resp, r, err := apiClient.AccessAndMobilityDataApi.DeleteAccessAndMobilityData(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessAndMobilityDataApi.DeleteAccessAndMobilityData``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAccessAndMobilityDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryAccessAndMobilityData

> AccessAndMobilityData QueryAccessAndMobilityData(ctx, ueId).SuppFeat(suppFeat).Execute()

Retrieves the access and mobility exposure data for a UE

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
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessAndMobilityDataApi.QueryAccessAndMobilityData(context.Background(), ueId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessAndMobilityDataApi.QueryAccessAndMobilityData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAccessAndMobilityData`: AccessAndMobilityData
    fmt.Fprintf(os.Stdout, "Response from `AccessAndMobilityDataApi.QueryAccessAndMobilityData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryAccessAndMobilityDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppFeat** | **string** | Supported Features | 

### Return type

[**AccessAndMobilityData**](AccessAndMobilityData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessAndMobilityData

> UpdateAccessAndMobilityData(ctx, ueId).AccessAndMobilityData(accessAndMobilityData).Execute()

Updates the access and mobility exposure data for a UE

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
    accessAndMobilityData := TODO // AccessAndMobilityData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessAndMobilityDataApi.UpdateAccessAndMobilityData(context.Background(), ueId).AccessAndMobilityData(accessAndMobilityData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessAndMobilityDataApi.UpdateAccessAndMobilityData``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateAccessAndMobilityDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessAndMobilityData** | [**AccessAndMobilityData**](AccessAndMobilityData.md) |  | 

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


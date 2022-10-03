# \IndividualPFDDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceIndividualPFDData**](IndividualPFDDataDocumentApi.md#CreateOrReplaceIndividualPFDData) | **Put** /application-data/pfds/{appId} | Create or update the corresponding PFDs for the specified application identifier
[**DeleteIndividualPFDData**](IndividualPFDDataDocumentApi.md#DeleteIndividualPFDData) | **Delete** /application-data/pfds/{appId} | Delete the corresponding PFDs of the specified application identifier
[**ReadIndividualPFDData**](IndividualPFDDataDocumentApi.md#ReadIndividualPFDData) | **Get** /application-data/pfds/{appId} | Retrieve the corresponding PFDs of the specified application identifier



## CreateOrReplaceIndividualPFDData

> PfdDataForAppExt CreateOrReplaceIndividualPFDData(ctx, appId).PfdDataForAppExt(pfdDataForAppExt).Execute()

Create or update the corresponding PFDs for the specified application identifier

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
    appId := "appId_example" // string | Indicate the application identifier for the request pfd(s). It shall apply the format of Data type ApplicationId.
    pfdDataForAppExt := TODO // PfdDataForAppExt | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPFDDataDocumentApi.CreateOrReplaceIndividualPFDData(context.Background(), appId).PfdDataForAppExt(pfdDataForAppExt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPFDDataDocumentApi.CreateOrReplaceIndividualPFDData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceIndividualPFDData`: PfdDataForAppExt
    fmt.Fprintf(os.Stdout, "Response from `IndividualPFDDataDocumentApi.CreateOrReplaceIndividualPFDData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Indicate the application identifier for the request pfd(s). It shall apply the format of Data type ApplicationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceIndividualPFDDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pfdDataForAppExt** | [**PfdDataForAppExt**](PfdDataForAppExt.md) |  | 

### Return type

[**PfdDataForAppExt**](PfdDataForAppExt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndividualPFDData

> DeleteIndividualPFDData(ctx, appId).Execute()

Delete the corresponding PFDs of the specified application identifier

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
    appId := "appId_example" // string | Indicate the application identifier for the request pfd(s). It shall apply the format of Data type ApplicationId.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPFDDataDocumentApi.DeleteIndividualPFDData(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPFDDataDocumentApi.DeleteIndividualPFDData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Indicate the application identifier for the request pfd(s). It shall apply the format of Data type ApplicationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualPFDDataRequest struct via the builder pattern


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


## ReadIndividualPFDData

> PfdDataForAppExt ReadIndividualPFDData(ctx, appId).Execute()

Retrieve the corresponding PFDs of the specified application identifier

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
    appId := "appId_example" // string | Indicate the application identifier for the request pfd(s). It shall apply the format of Data type ApplicationId.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPFDDataDocumentApi.ReadIndividualPFDData(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPFDDataDocumentApi.ReadIndividualPFDData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualPFDData`: PfdDataForAppExt
    fmt.Fprintf(os.Stdout, "Response from `IndividualPFDDataDocumentApi.ReadIndividualPFDData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Indicate the application identifier for the request pfd(s). It shall apply the format of Data type ApplicationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualPFDDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PfdDataForAppExt**](PfdDataForAppExt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \IndividualServiceParameterDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceServiceParameterData**](IndividualServiceParameterDataDocumentApi.md#CreateOrReplaceServiceParameterData) | **Put** /application-data/serviceParamData/{serviceParamId} | Create or update an individual Service Parameter Data resource
[**DeleteIndividualServiceParameterData**](IndividualServiceParameterDataDocumentApi.md#DeleteIndividualServiceParameterData) | **Delete** /application-data/serviceParamData/{serviceParamId} | Delete an individual Service Parameter Data resource
[**UpdateIndividualServiceParameterData**](IndividualServiceParameterDataDocumentApi.md#UpdateIndividualServiceParameterData) | **Patch** /application-data/serviceParamData/{serviceParamId} | Modify part of the properties of an individual Service Parameter Data resource



## CreateOrReplaceServiceParameterData

> ServiceParameterData CreateOrReplaceServiceParameterData(ctx, serviceParamId).ServiceParameterData(serviceParameterData).Execute()

Create or update an individual Service Parameter Data resource

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
    serviceParamId := "serviceParamId_example" // string | The Identifier of an Individual Service Parameter Data to be created or updated. It shall apply the format of Data type string.
    serviceParameterData := TODO // ServiceParameterData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceParameterDataDocumentApi.CreateOrReplaceServiceParameterData(context.Background(), serviceParamId).ServiceParameterData(serviceParameterData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceParameterDataDocumentApi.CreateOrReplaceServiceParameterData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceServiceParameterData`: ServiceParameterData
    fmt.Fprintf(os.Stdout, "Response from `IndividualServiceParameterDataDocumentApi.CreateOrReplaceServiceParameterData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceParamId** | **string** | The Identifier of an Individual Service Parameter Data to be created or updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceServiceParameterDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceParameterData** | [**ServiceParameterData**](ServiceParameterData.md) |  | 

### Return type

[**ServiceParameterData**](ServiceParameterData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndividualServiceParameterData

> DeleteIndividualServiceParameterData(ctx, serviceParamId).Execute()

Delete an individual Service Parameter Data resource

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
    serviceParamId := "serviceParamId_example" // string | The Identifier of an Individual Service Parameter Data to be updated. It shall apply the format of Data type string.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceParameterDataDocumentApi.DeleteIndividualServiceParameterData(context.Background(), serviceParamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceParameterDataDocumentApi.DeleteIndividualServiceParameterData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceParamId** | **string** | The Identifier of an Individual Service Parameter Data to be updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualServiceParameterDataRequest struct via the builder pattern


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


## UpdateIndividualServiceParameterData

> ServiceParameterData UpdateIndividualServiceParameterData(ctx, serviceParamId).ServiceParameterDataPatch(serviceParameterDataPatch).Execute()

Modify part of the properties of an individual Service Parameter Data resource

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
    serviceParamId := "serviceParamId_example" // string | The Identifier of an Individual Service Parameter Data to be updated. It shall apply the format of Data type string.
    serviceParameterDataPatch := TODO // ServiceParameterDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceParameterDataDocumentApi.UpdateIndividualServiceParameterData(context.Background(), serviceParamId).ServiceParameterDataPatch(serviceParameterDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceParameterDataDocumentApi.UpdateIndividualServiceParameterData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndividualServiceParameterData`: ServiceParameterData
    fmt.Fprintf(os.Stdout, "Response from `IndividualServiceParameterDataDocumentApi.UpdateIndividualServiceParameterData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceParamId** | **string** | The Identifier of an Individual Service Parameter Data to be updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndividualServiceParameterDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceParameterDataPatch** | [**ServiceParameterDataPatch**](ServiceParameterDataPatch.md) |  | 

### Return type

[**ServiceParameterData**](ServiceParameterData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


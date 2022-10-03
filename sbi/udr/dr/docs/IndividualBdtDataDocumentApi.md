# \IndividualBdtDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualBdtData**](IndividualBdtDataDocumentApi.md#CreateIndividualBdtData) | **Put** /policy-data/bdt-data/{bdtReferenceId} | Creates an BDT data resource associated with an BDT reference Id
[**DeleteIndividualBdtData**](IndividualBdtDataDocumentApi.md#DeleteIndividualBdtData) | **Delete** /policy-data/bdt-data/{bdtReferenceId} | Deletes an BDT data resource associated with an BDT reference Id
[**ReadIndividualBdtData**](IndividualBdtDataDocumentApi.md#ReadIndividualBdtData) | **Get** /policy-data/bdt-data/{bdtReferenceId} | Retrieves the BDT data information associated with a BDT reference Id
[**UpdateIndividualBdtData**](IndividualBdtDataDocumentApi.md#UpdateIndividualBdtData) | **Patch** /policy-data/bdt-data/{bdtReferenceId} | Modifies an BDT data resource associated with an BDT reference Id



## CreateIndividualBdtData

> BdtData CreateIndividualBdtData(ctx, bdtReferenceId).BdtData(bdtData).Execute()

Creates an BDT data resource associated with an BDT reference Id

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
    bdtReferenceId := "bdtReferenceId_example" // string | 
    bdtData := TODO // BdtData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBdtDataDocumentApi.CreateIndividualBdtData(context.Background(), bdtReferenceId).BdtData(bdtData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBdtDataDocumentApi.CreateIndividualBdtData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualBdtData`: BdtData
    fmt.Fprintf(os.Stdout, "Response from `IndividualBdtDataDocumentApi.CreateIndividualBdtData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bdtReferenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualBdtDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bdtData** | [**BdtData**](BdtData.md) |  | 

### Return type

[**BdtData**](BdtData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndividualBdtData

> DeleteIndividualBdtData(ctx, bdtReferenceId).Execute()

Deletes an BDT data resource associated with an BDT reference Id

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
    bdtReferenceId := "bdtReferenceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBdtDataDocumentApi.DeleteIndividualBdtData(context.Background(), bdtReferenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBdtDataDocumentApi.DeleteIndividualBdtData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bdtReferenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualBdtDataRequest struct via the builder pattern


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


## ReadIndividualBdtData

> BdtData ReadIndividualBdtData(ctx, bdtReferenceId).SuppFeat(suppFeat).Execute()

Retrieves the BDT data information associated with a BDT reference Id

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
    bdtReferenceId := "bdtReferenceId_example" // string | 
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBdtDataDocumentApi.ReadIndividualBdtData(context.Background(), bdtReferenceId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBdtDataDocumentApi.ReadIndividualBdtData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualBdtData`: BdtData
    fmt.Fprintf(os.Stdout, "Response from `IndividualBdtDataDocumentApi.ReadIndividualBdtData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bdtReferenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualBdtDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppFeat** | **string** | Supported Features | 

### Return type

[**BdtData**](BdtData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndividualBdtData

> BdtData UpdateIndividualBdtData(ctx, bdtReferenceId).BdtDataPatch(bdtDataPatch).Execute()

Modifies an BDT data resource associated with an BDT reference Id

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
    bdtReferenceId := "bdtReferenceId_example" // string | 
    bdtDataPatch := TODO // BdtDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBdtDataDocumentApi.UpdateIndividualBdtData(context.Background(), bdtReferenceId).BdtDataPatch(bdtDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBdtDataDocumentApi.UpdateIndividualBdtData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndividualBdtData`: BdtData
    fmt.Fprintf(os.Stdout, "Response from `IndividualBdtDataDocumentApi.UpdateIndividualBdtData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bdtReferenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndividualBdtDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bdtDataPatch** | [**BdtDataPatch**](BdtDataPatch.md) |  | 

### Return type

[**BdtData**](BdtData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


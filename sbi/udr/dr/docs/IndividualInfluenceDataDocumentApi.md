# \IndividualInfluenceDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceIndividualInfluenceData**](IndividualInfluenceDataDocumentApi.md#CreateOrReplaceIndividualInfluenceData) | **Put** /application-data/influenceData/{influenceId} | Create or update an individual Influence Data resource
[**DeleteIndividualInfluenceData**](IndividualInfluenceDataDocumentApi.md#DeleteIndividualInfluenceData) | **Delete** /application-data/influenceData/{influenceId} | Delete an individual Influence Data resource
[**UpdateIndividualInfluenceData**](IndividualInfluenceDataDocumentApi.md#UpdateIndividualInfluenceData) | **Patch** /application-data/influenceData/{influenceId} | Modify part of the properties of an individual Influence Data resource



## CreateOrReplaceIndividualInfluenceData

> TrafficInfluData CreateOrReplaceIndividualInfluenceData(ctx, influenceId).TrafficInfluData(trafficInfluData).Execute()

Create or update an individual Influence Data resource

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
    influenceId := "influenceId_example" // string | The Identifier of an Individual Influence Data to be created or updated. It shall apply the format of Data type string.
    trafficInfluData := TODO // TrafficInfluData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualInfluenceDataDocumentApi.CreateOrReplaceIndividualInfluenceData(context.Background(), influenceId).TrafficInfluData(trafficInfluData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualInfluenceDataDocumentApi.CreateOrReplaceIndividualInfluenceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceIndividualInfluenceData`: TrafficInfluData
    fmt.Fprintf(os.Stdout, "Response from `IndividualInfluenceDataDocumentApi.CreateOrReplaceIndividualInfluenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**influenceId** | **string** | The Identifier of an Individual Influence Data to be created or updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceIndividualInfluenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trafficInfluData** | [**TrafficInfluData**](TrafficInfluData.md) |  | 

### Return type

[**TrafficInfluData**](TrafficInfluData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndividualInfluenceData

> DeleteIndividualInfluenceData(ctx, influenceId).Execute()

Delete an individual Influence Data resource

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
    influenceId := "influenceId_example" // string | The Identifier of an Individual Influence Data to be updated. It shall apply the format of Data type string.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualInfluenceDataDocumentApi.DeleteIndividualInfluenceData(context.Background(), influenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualInfluenceDataDocumentApi.DeleteIndividualInfluenceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**influenceId** | **string** | The Identifier of an Individual Influence Data to be updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualInfluenceDataRequest struct via the builder pattern


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


## UpdateIndividualInfluenceData

> TrafficInfluData UpdateIndividualInfluenceData(ctx, influenceId).TrafficInfluDataPatch(trafficInfluDataPatch).Execute()

Modify part of the properties of an individual Influence Data resource

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
    influenceId := "influenceId_example" // string | The Identifier of an Individual Influence Data to be updated. It shall apply the format of Data type string.
    trafficInfluDataPatch := TODO // TrafficInfluDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualInfluenceDataDocumentApi.UpdateIndividualInfluenceData(context.Background(), influenceId).TrafficInfluDataPatch(trafficInfluDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualInfluenceDataDocumentApi.UpdateIndividualInfluenceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndividualInfluenceData`: TrafficInfluData
    fmt.Fprintf(os.Stdout, "Response from `IndividualInfluenceDataDocumentApi.UpdateIndividualInfluenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**influenceId** | **string** | The Identifier of an Individual Influence Data to be updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndividualInfluenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trafficInfluDataPatch** | [**TrafficInfluDataPatch**](TrafficInfluDataPatch.md) |  | 

### Return type

[**TrafficInfluData**](TrafficInfluData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


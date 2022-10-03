# \IndividualAppliedBDTPolicyDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualAppliedBdtPolicyData**](IndividualAppliedBDTPolicyDataDocumentApi.md#CreateIndividualAppliedBdtPolicyData) | **Put** /application-data/bdtPolicyData/{bdtPolicyId} | Create an individual applied BDT Policy Data resource
[**DeleteIndividualAppliedBdtPolicyData**](IndividualAppliedBDTPolicyDataDocumentApi.md#DeleteIndividualAppliedBdtPolicyData) | **Delete** /application-data/bdtPolicyData/{bdtPolicyId} | Delete an individual Applied BDT Policy Data resource
[**UpdateIndividualAppliedBdtPolicyData**](IndividualAppliedBDTPolicyDataDocumentApi.md#UpdateIndividualAppliedBdtPolicyData) | **Patch** /application-data/bdtPolicyData/{bdtPolicyId} | Modify part of the properties of an individual Applied BDT Policy Data resource



## CreateIndividualAppliedBdtPolicyData

> BdtPolicyData CreateIndividualAppliedBdtPolicyData(ctx, bdtPolicyId).BdtPolicyData(bdtPolicyData).Execute()

Create an individual applied BDT Policy Data resource

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
    bdtPolicyId := "bdtPolicyId_example" // string | The Identifier of an Individual Applied BDT Policy Data to be created or updated. It shall apply the format of Data type string.
    bdtPolicyData := TODO // BdtPolicyData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAppliedBDTPolicyDataDocumentApi.CreateIndividualAppliedBdtPolicyData(context.Background(), bdtPolicyId).BdtPolicyData(bdtPolicyData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAppliedBDTPolicyDataDocumentApi.CreateIndividualAppliedBdtPolicyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualAppliedBdtPolicyData`: BdtPolicyData
    fmt.Fprintf(os.Stdout, "Response from `IndividualAppliedBDTPolicyDataDocumentApi.CreateIndividualAppliedBdtPolicyData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bdtPolicyId** | **string** | The Identifier of an Individual Applied BDT Policy Data to be created or updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualAppliedBdtPolicyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bdtPolicyData** | [**BdtPolicyData**](BdtPolicyData.md) |  | 

### Return type

[**BdtPolicyData**](BdtPolicyData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndividualAppliedBdtPolicyData

> DeleteIndividualAppliedBdtPolicyData(ctx, bdtPolicyId).Execute()

Delete an individual Applied BDT Policy Data resource

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
    bdtPolicyId := "bdtPolicyId_example" // string | The Identifier of an Individual Applied BDT Policy Data to be updated. It shall apply the format of Data type string.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAppliedBDTPolicyDataDocumentApi.DeleteIndividualAppliedBdtPolicyData(context.Background(), bdtPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAppliedBDTPolicyDataDocumentApi.DeleteIndividualAppliedBdtPolicyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bdtPolicyId** | **string** | The Identifier of an Individual Applied BDT Policy Data to be updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualAppliedBdtPolicyDataRequest struct via the builder pattern


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


## UpdateIndividualAppliedBdtPolicyData

> BdtPolicyData UpdateIndividualAppliedBdtPolicyData(ctx, bdtPolicyId).BdtPolicyDataPatch(bdtPolicyDataPatch).Execute()

Modify part of the properties of an individual Applied BDT Policy Data resource

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
    bdtPolicyId := "bdtPolicyId_example" // string | The Identifier of an Individual Applied BDT Policy Data to be updated. It shall apply the format of Data type string.
    bdtPolicyDataPatch := TODO // BdtPolicyDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAppliedBDTPolicyDataDocumentApi.UpdateIndividualAppliedBdtPolicyData(context.Background(), bdtPolicyId).BdtPolicyDataPatch(bdtPolicyDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAppliedBDTPolicyDataDocumentApi.UpdateIndividualAppliedBdtPolicyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndividualAppliedBdtPolicyData`: BdtPolicyData
    fmt.Fprintf(os.Stdout, "Response from `IndividualAppliedBDTPolicyDataDocumentApi.UpdateIndividualAppliedBdtPolicyData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bdtPolicyId** | **string** | The Identifier of an Individual Applied BDT Policy Data to be updated. It shall apply the format of Data type string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndividualAppliedBdtPolicyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bdtPolicyDataPatch** | [**BdtPolicyDataPatch**](BdtPolicyDataPatch.md) |  | 

### Return type

[**BdtPolicyData**](BdtPolicyData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


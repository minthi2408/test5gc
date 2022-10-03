# \ProvisionedDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryProvisionedData**](ProvisionedDataDocumentApi.md#QueryProvisionedData) | **Get** /subscription-data/{ueId}/{servingPlmnId}/provisioned-data | Retrieve multiple provisioned data sets of a UE



## QueryProvisionedData

> ProvisionedDataSets QueryProvisionedData(ctx, ueId, servingPlmnId).DatasetNames(datasetNames).Execute()

Retrieve multiple provisioned data sets of a UE

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
    servingPlmnId := "servingPlmnId_example" // string | PLMN ID
    datasetNames := []DataSetName{"TODO"} // []DataSetName | List of dataset names (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisionedDataDocumentApi.QueryProvisionedData(context.Background(), ueId, servingPlmnId).DatasetNames(datasetNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedDataDocumentApi.QueryProvisionedData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryProvisionedData`: ProvisionedDataSets
    fmt.Fprintf(os.Stdout, "Response from `ProvisionedDataDocumentApi.QueryProvisionedData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**servingPlmnId** | **string** | PLMN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryProvisionedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasetNames** | [**[]DataSetName**](DataSetName.md) | List of dataset names | 

### Return type

[**ProvisionedDataSets**](ProvisionedDataSets.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


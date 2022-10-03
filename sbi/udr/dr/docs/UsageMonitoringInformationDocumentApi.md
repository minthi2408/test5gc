# \UsageMonitoringInformationDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsageMonitoringResource**](UsageMonitoringInformationDocumentApi.md#CreateUsageMonitoringResource) | **Put** /policy-data/ues/{ueId}/sm-data/{usageMonId} | Create a usage monitoring resource
[**DeleteUsageMonitoringInformation**](UsageMonitoringInformationDocumentApi.md#DeleteUsageMonitoringInformation) | **Delete** /policy-data/ues/{ueId}/sm-data/{usageMonId} | Delete a usage monitoring resource
[**ReadUsageMonitoringInformation**](UsageMonitoringInformationDocumentApi.md#ReadUsageMonitoringInformation) | **Get** /policy-data/ues/{ueId}/sm-data/{usageMonId} | Retrieve a usage monitoring resource



## CreateUsageMonitoringResource

> UsageMonData CreateUsageMonitoringResource(ctx, ueId, usageMonId).UsageMonData(usageMonData).Execute()

Create a usage monitoring resource

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
    usageMonId := "usageMonId_example" // string | 
    usageMonData := TODO // UsageMonData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMonitoringInformationDocumentApi.CreateUsageMonitoringResource(context.Background(), ueId, usageMonId).UsageMonData(usageMonData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMonitoringInformationDocumentApi.CreateUsageMonitoringResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsageMonitoringResource`: UsageMonData
    fmt.Fprintf(os.Stdout, "Response from `UsageMonitoringInformationDocumentApi.CreateUsageMonitoringResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**usageMonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsageMonitoringResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **usageMonData** | [**UsageMonData**](UsageMonData.md) |  | 

### Return type

[**UsageMonData**](UsageMonData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsageMonitoringInformation

> DeleteUsageMonitoringInformation(ctx, ueId, usageMonId).Execute()

Delete a usage monitoring resource

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
    usageMonId := "usageMonId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMonitoringInformationDocumentApi.DeleteUsageMonitoringInformation(context.Background(), ueId, usageMonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMonitoringInformationDocumentApi.DeleteUsageMonitoringInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**usageMonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsageMonitoringInformationRequest struct via the builder pattern


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


## ReadUsageMonitoringInformation

> UsageMonData ReadUsageMonitoringInformation(ctx, ueId, usageMonId).SuppFeat(suppFeat).Execute()

Retrieve a usage monitoring resource

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
    usageMonId := "usageMonId_example" // string | 
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageMonitoringInformationDocumentApi.ReadUsageMonitoringInformation(context.Background(), ueId, usageMonId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageMonitoringInformationDocumentApi.ReadUsageMonitoringInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUsageMonitoringInformation`: UsageMonData
    fmt.Fprintf(os.Stdout, "Response from `UsageMonitoringInformationDocumentApi.ReadUsageMonitoringInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**usageMonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUsageMonitoringInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suppFeat** | **string** | Supported Features | 

### Return type

[**UsageMonData**](UsageMonData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


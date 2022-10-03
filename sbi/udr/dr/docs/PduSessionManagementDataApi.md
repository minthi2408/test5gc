# \PduSessionManagementDataApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceSessionManagementData**](PduSessionManagementDataApi.md#CreateOrReplaceSessionManagementData) | **Put** /exposure-data/{ueId}/session-management-data/{pduSessionId} | Creates and updates the session management data for a UE and for an individual PDU session
[**DeleteSessionManagementData**](PduSessionManagementDataApi.md#DeleteSessionManagementData) | **Delete** /exposure-data/{ueId}/session-management-data/{pduSessionId} | Deletes the session management data for a UE and for an individual PDU session
[**QuerySessionManagementData**](PduSessionManagementDataApi.md#QuerySessionManagementData) | **Get** /exposure-data/{ueId}/session-management-data/{pduSessionId} | Retrieves the session management data for a UE and for an individual PDU session



## CreateOrReplaceSessionManagementData

> AccessAndMobilityData CreateOrReplaceSessionManagementData(ctx, ueId, pduSessionId).PduSessionManagementData(pduSessionManagementData).Execute()

Creates and updates the session management data for a UE and for an individual PDU session

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
    pduSessionId := int32(56) // int32 | PDU session id
    pduSessionManagementData := TODO // PduSessionManagementData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PduSessionManagementDataApi.CreateOrReplaceSessionManagementData(context.Background(), ueId, pduSessionId).PduSessionManagementData(pduSessionManagementData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PduSessionManagementDataApi.CreateOrReplaceSessionManagementData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceSessionManagementData`: AccessAndMobilityData
    fmt.Fprintf(os.Stdout, "Response from `PduSessionManagementDataApi.CreateOrReplaceSessionManagementData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**pduSessionId** | **int32** | PDU session id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceSessionManagementDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pduSessionManagementData** | [**PduSessionManagementData**](PduSessionManagementData.md) |  | 

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


## DeleteSessionManagementData

> DeleteSessionManagementData(ctx, ueId, pduSessionId).Execute()

Deletes the session management data for a UE and for an individual PDU session

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
    pduSessionId := int32(56) // int32 | PDU session id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PduSessionManagementDataApi.DeleteSessionManagementData(context.Background(), ueId, pduSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PduSessionManagementDataApi.DeleteSessionManagementData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**pduSessionId** | **int32** | PDU session id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionManagementDataRequest struct via the builder pattern


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


## QuerySessionManagementData

> PduSessionManagementData QuerySessionManagementData(ctx, ueId, pduSessionId).Ipv4Addr(ipv4Addr).Ipv6Prefix(ipv6Prefix).Dnn(dnn).Fields(fields).SuppFeat(suppFeat).Execute()

Retrieves the session management data for a UE and for an individual PDU session

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
    pduSessionId := int32(56) // int32 | PDU session id
    ipv4Addr := "ipv4Addr_example" // string | IPv4 Address of the UE (optional)
    ipv6Prefix := "ipv6Prefix_example" // Ipv6Prefix | IPv6 Address Prefix of the UE (optional)
    dnn := "dnn_example" // string | DNN of the UE (optional)
    fields := []string{"Inner_example"} // []string | attributes to be retrieved (optional)
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PduSessionManagementDataApi.QuerySessionManagementData(context.Background(), ueId, pduSessionId).Ipv4Addr(ipv4Addr).Ipv6Prefix(ipv6Prefix).Dnn(dnn).Fields(fields).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PduSessionManagementDataApi.QuerySessionManagementData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySessionManagementData`: PduSessionManagementData
    fmt.Fprintf(os.Stdout, "Response from `PduSessionManagementDataApi.QuerySessionManagementData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**pduSessionId** | **int32** | PDU session id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySessionManagementDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ipv4Addr** | **string** | IPv4 Address of the UE | 
 **ipv6Prefix** | **Ipv6Prefix** | IPv6 Address Prefix of the UE | 
 **dnn** | **string** | DNN of the UE | 
 **fields** | **[]string** | attributes to be retrieved | 
 **suppFeat** | **string** | Supported Features | 

### Return type

[**PduSessionManagementData**](PduSessionManagementData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


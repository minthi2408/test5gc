# \NSSAIUpdateAckDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateNssaiAck**](NSSAIUpdateAckDocumentApi.md#CreateOrUpdateNssaiAck) | **Put** /subscription-data/{ueId}/ue-update-confirmation-data/subscribed-snssais | To store the NSSAI update acknowledgement information of a UE



## CreateOrUpdateNssaiAck

> CreateOrUpdateNssaiAck(ctx, ueId).NssaiAckData(nssaiAckData).SupportedFeatures(supportedFeatures).Execute()

To store the NSSAI update acknowledgement information of a UE

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
    nssaiAckData := TODO // NssaiAckData | 
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NSSAIUpdateAckDocumentApi.CreateOrUpdateNssaiAck(context.Background(), ueId).NssaiAckData(nssaiAckData).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NSSAIUpdateAckDocumentApi.CreateOrUpdateNssaiAck``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateOrUpdateNssaiAckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nssaiAckData** | [**NssaiAckData**](NssaiAckData.md) |  | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


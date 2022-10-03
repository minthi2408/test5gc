# \SponsorConnectivityDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadSponsorConnectivityData**](SponsorConnectivityDataDocumentApi.md#ReadSponsorConnectivityData) | **Get** /policy-data/sponsor-connectivity-data/{sponsorId} | Retrieves the sponsored connectivity information for a given sponsorId



## ReadSponsorConnectivityData

> SponsorConnectivityData ReadSponsorConnectivityData(ctx, sponsorId).Execute()

Retrieves the sponsored connectivity information for a given sponsorId

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
    sponsorId := "sponsorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SponsorConnectivityDataDocumentApi.ReadSponsorConnectivityData(context.Background(), sponsorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SponsorConnectivityDataDocumentApi.ReadSponsorConnectivityData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSponsorConnectivityData`: SponsorConnectivityData
    fmt.Fprintf(os.Stdout, "Response from `SponsorConnectivityDataDocumentApi.ReadSponsorConnectivityData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sponsorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSponsorConnectivityDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SponsorConnectivityData**](SponsorConnectivityData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


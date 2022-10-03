# \QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNiddAuData**](QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApi.md#GetNiddAuData) | **Get** /subscription-data/{ueId}/nidd-authorization-data | Retrieve NIDD Authorization Data GPSI or External Group identifier



## GetNiddAuData

> AuthorizationData GetNiddAuData(ctx, ueId).SingleNssai(singleNssai).Dnn(dnn).MtcProviderInformation(mtcProviderInformation).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Retrieve NIDD Authorization Data GPSI or External Group identifier

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
    ueId := "ueId_example" // string | UE ID
    singleNssai := Snssai{ ... } // Snssai | single NSSAI
    dnn := "dnn_example" // string | DNN
    mtcProviderInformation := "mtcProviderInformation_example" // string | MTC Provider Information
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApi.GetNiddAuData(context.Background(), ueId).SingleNssai(singleNssai).Dnn(dnn).MtcProviderInformation(mtcProviderInformation).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApi.GetNiddAuData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNiddAuData`: AuthorizationData
    fmt.Fprintf(os.Stdout, "Response from `QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApi.GetNiddAuData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNiddAuDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **singleNssai** | [**Snssai**](Snssai.md) | single NSSAI | 
 **dnn** | **string** | DNN | 
 **mtcProviderInformation** | **string** | MTC Provider Information | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**AuthorizationData**](AuthorizationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


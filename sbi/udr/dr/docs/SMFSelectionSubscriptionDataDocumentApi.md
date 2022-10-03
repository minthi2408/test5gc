# \SMFSelectionSubscriptionDataDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuerySmfSelectData**](SMFSelectionSubscriptionDataDocumentApi.md#QuerySmfSelectData) | **Get** /subscription-data/{ueId}/{servingPlmnId}/provisioned-data/smf-selection-subscription-data | Retrieves the SMF selection subscription data of a UE



## QuerySmfSelectData

> SmfSelectionSubscriptionData QuerySmfSelectData(ctx, ueId, servingPlmnId).Fields(fields).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Retrieves the SMF selection subscription data of a UE

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
    fields := []string{"Inner_example"} // []string | attributes to be retrieved (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFSelectionSubscriptionDataDocumentApi.QuerySmfSelectData(context.Background(), ueId, servingPlmnId).Fields(fields).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFSelectionSubscriptionDataDocumentApi.QuerySmfSelectData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuerySmfSelectData`: SmfSelectionSubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `SMFSelectionSubscriptionDataDocumentApi.QuerySmfSelectData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 
**servingPlmnId** | **string** | PLMN ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuerySmfSelectDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | attributes to be retrieved | 
 **supportedFeatures** | **string** | Supported Features | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**SmfSelectionSubscriptionData**](SmfSelectionSubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


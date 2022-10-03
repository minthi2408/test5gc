# \IPSMGWRegistrationDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpSmGwContext**](IPSMGWRegistrationDocumentApi.md#CreateIpSmGwContext) | **Put** /subscription-data/{ueId}/context-data/ip-sm-gw | Create the IP-SM-GW context data of a UE
[**DeleteIpSmGwContext**](IPSMGWRegistrationDocumentApi.md#DeleteIpSmGwContext) | **Delete** /subscription-data/{ueId}/context-data/ip-sm-gw | To remove the IP-SM-GW context data of a UE
[**ModifyIpSmGwContext**](IPSMGWRegistrationDocumentApi.md#ModifyIpSmGwContext) | **Patch** /subscription-data/{ueId}/context-data/ip-sm-gw | Modify the IP-SM-GW context data of a UE
[**QueryIpSmGwContext**](IPSMGWRegistrationDocumentApi.md#QueryIpSmGwContext) | **Get** /subscription-data/{ueId}/context-data/ip-sm-gw | Retrieves the IP-SM-GW context data of a UE



## CreateIpSmGwContext

> CreateIpSmGwContext(ctx, ueId).IpSmGwRegistration(ipSmGwRegistration).Execute()

Create the IP-SM-GW context data of a UE

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
    ipSmGwRegistration := TODO // IpSmGwRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPSMGWRegistrationDocumentApi.CreateIpSmGwContext(context.Background(), ueId).IpSmGwRegistration(ipSmGwRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSMGWRegistrationDocumentApi.CreateIpSmGwContext``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateIpSmGwContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipSmGwRegistration** | [**IpSmGwRegistration**](IpSmGwRegistration.md) |  | 

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


## DeleteIpSmGwContext

> DeleteIpSmGwContext(ctx, ueId).Execute()

To remove the IP-SM-GW context data of a UE

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPSMGWRegistrationDocumentApi.DeleteIpSmGwContext(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSMGWRegistrationDocumentApi.DeleteIpSmGwContext``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIpSmGwContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIpSmGwContext

> ModifyIpSmGwContext(ctx, ueId).PatchItem(patchItem).Execute()

Modify the IP-SM-GW context data of a UE

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
    patchItem := []PatchItem{"TODO"} // []PatchItem | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPSMGWRegistrationDocumentApi.ModifyIpSmGwContext(context.Background(), ueId).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSMGWRegistrationDocumentApi.ModifyIpSmGwContext``: %v\n", err)
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

Other parameters are passed through a pointer to a apiModifyIpSmGwContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryIpSmGwContext

> IpSmGwRegistration QueryIpSmGwContext(ctx, ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()

Retrieves the IP-SM-GW context data of a UE

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
    fields := []string{"Inner_example"} // []string | attributes to be retrieved (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPSMGWRegistrationDocumentApi.QueryIpSmGwContext(context.Background(), ueId).Fields(fields).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSMGWRegistrationDocumentApi.QueryIpSmGwContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryIpSmGwContext`: IpSmGwRegistration
    fmt.Fprintf(os.Stdout, "Response from `IPSMGWRegistrationDocumentApi.QueryIpSmGwContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryIpSmGwContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | attributes to be retrieved | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**IpSmGwRegistration**](IpSmGwRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


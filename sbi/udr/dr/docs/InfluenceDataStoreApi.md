# \InfluenceDataStoreApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadInfluenceData**](InfluenceDataStoreApi.md#ReadInfluenceData) | **Get** /application-data/influenceData | Retrieve Traffic Influence Data



## ReadInfluenceData

> []TrafficInfluData ReadInfluenceData(ctx).InfluenceIds(influenceIds).Dnns(dnns).Snssais(snssais).InternalGroupIds(internalGroupIds).Supis(supis).SuppFeat(suppFeat).Execute()

Retrieve Traffic Influence Data

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
    influenceIds := []string{"Inner_example"} // []string | Each element identifies a service. (optional)
    dnns := []string{"Inner_example"} // []string | Each element identifies a DNN. (optional)
    snssais := []Snssai{"TODO"} // []Snssai | Each element identifies a slice. (optional)
    internalGroupIds := []string{"Inner_example"} // []string | Each element identifies a group of users. (optional)
    supis := []string{"Inner_example"} // []string | Each element identifies the user. (optional)
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InfluenceDataStoreApi.ReadInfluenceData(context.Background()).InfluenceIds(influenceIds).Dnns(dnns).Snssais(snssais).InternalGroupIds(internalGroupIds).Supis(supis).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfluenceDataStoreApi.ReadInfluenceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadInfluenceData`: []TrafficInfluData
    fmt.Fprintf(os.Stdout, "Response from `InfluenceDataStoreApi.ReadInfluenceData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadInfluenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **influenceIds** | **[]string** | Each element identifies a service. | 
 **dnns** | **[]string** | Each element identifies a DNN. | 
 **snssais** | [**[]Snssai**](Snssai.md) | Each element identifies a slice. | 
 **internalGroupIds** | **[]string** | Each element identifies a group of users. | 
 **supis** | **[]string** | Each element identifies the user. | 
 **suppFeat** | **string** | Supported Features | 

### Return type

[**[]TrafficInfluData**](TrafficInfluData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


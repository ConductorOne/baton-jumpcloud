# \ImageApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsDeleteLogo**](ImageApi.md#ApplicationsDeleteLogo) | **Delete** /applications/{application_id}/logo | Delete application image



## ApplicationsDeleteLogo

> ApplicationsDeleteLogo(ctx, applicationId).XOrgId(xOrgId).Execute()

Delete application image



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    applicationId := "applicationId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImageApi.ApplicationsDeleteLogo(context.Background(), applicationId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.ApplicationsDeleteLogo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsDeleteLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


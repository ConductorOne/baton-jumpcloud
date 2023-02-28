# \CustomEmailsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomEmailsCreate**](CustomEmailsApi.md#CustomEmailsCreate) | **Post** /customemails | Create custom email configuration
[**CustomEmailsDestroy**](CustomEmailsApi.md#CustomEmailsDestroy) | **Delete** /customemails/{custom_email_type} | Delete custom email configuration
[**CustomEmailsGetTemplates**](CustomEmailsApi.md#CustomEmailsGetTemplates) | **Get** /customemail/templates | List custom email templates
[**CustomEmailsRead**](CustomEmailsApi.md#CustomEmailsRead) | **Get** /customemails/{custom_email_type} | Get custom email configuration
[**CustomEmailsUpdate**](CustomEmailsApi.md#CustomEmailsUpdate) | **Put** /customemails/{custom_email_type} | Update custom email configuration



## CustomEmailsCreate

> CustomEmail CustomEmailsCreate(ctx).XOrgId(xOrgId).Body(body).Execute()

Create custom email configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewCustomEmail("Subject_example", openapiclient.CustomEmailType("activate_gapps_user")) // CustomEmail |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEmailsApi.CustomEmailsCreate(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEmailsApi.CustomEmailsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomEmailsCreate`: CustomEmail
    fmt.Fprintf(os.Stdout, "Response from `CustomEmailsApi.CustomEmailsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomEmailsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**CustomEmail**](CustomEmail.md) |  | 

### Return type

[**CustomEmail**](CustomEmail.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomEmailsDestroy

> CustomEmailsDestroy(ctx, customEmailType).XOrgId(xOrgId).Execute()

Delete custom email configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    customEmailType := "customEmailType_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomEmailsApi.CustomEmailsDestroy(context.Background(), customEmailType).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEmailsApi.CustomEmailsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customEmailType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomEmailsDestroyRequest struct via the builder pattern


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


## CustomEmailsGetTemplates

> []CustomEmailTemplate CustomEmailsGetTemplates(ctx).Execute()

List custom email templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEmailsApi.CustomEmailsGetTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEmailsApi.CustomEmailsGetTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomEmailsGetTemplates`: []CustomEmailTemplate
    fmt.Fprintf(os.Stdout, "Response from `CustomEmailsApi.CustomEmailsGetTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCustomEmailsGetTemplatesRequest struct via the builder pattern


### Return type

[**[]CustomEmailTemplate**](CustomEmailTemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomEmailsRead

> CustomEmail CustomEmailsRead(ctx, customEmailType).XOrgId(xOrgId).Execute()

Get custom email configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    customEmailType := "customEmailType_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEmailsApi.CustomEmailsRead(context.Background(), customEmailType).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEmailsApi.CustomEmailsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomEmailsRead`: CustomEmail
    fmt.Fprintf(os.Stdout, "Response from `CustomEmailsApi.CustomEmailsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customEmailType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomEmailsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**CustomEmail**](CustomEmail.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomEmailsUpdate

> CustomEmail CustomEmailsUpdate(ctx, customEmailType).XOrgId(xOrgId).Body(body).Execute()

Update custom email configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    customEmailType := "customEmailType_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewCustomEmail("Subject_example", openapiclient.CustomEmailType("activate_gapps_user")) // CustomEmail |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEmailsApi.CustomEmailsUpdate(context.Background(), customEmailType).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEmailsApi.CustomEmailsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomEmailsUpdate`: CustomEmail
    fmt.Fprintf(os.Stdout, "Response from `CustomEmailsApi.CustomEmailsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customEmailType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomEmailsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**CustomEmail**](CustomEmail.md) |  | 

### Return type

[**CustomEmail**](CustomEmail.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


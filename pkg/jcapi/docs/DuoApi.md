# \DuoApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DuoAccountDelete**](DuoApi.md#DuoAccountDelete) | **Delete** /duo/accounts/{id} | Delete a Duo Account
[**DuoAccountGet**](DuoApi.md#DuoAccountGet) | **Get** /duo/accounts/{id} | Get a Duo Acount
[**DuoAccountList**](DuoApi.md#DuoAccountList) | **Get** /duo/accounts | List Duo Accounts
[**DuoAccountPost**](DuoApi.md#DuoAccountPost) | **Post** /duo/accounts | Create Duo Account
[**DuoApplicationDelete**](DuoApi.md#DuoApplicationDelete) | **Delete** /duo/accounts/{account_id}/applications/{application_id} | Delete a Duo Application
[**DuoApplicationGet**](DuoApi.md#DuoApplicationGet) | **Get** /duo/accounts/{account_id}/applications/{application_id} | Get a Duo application
[**DuoApplicationList**](DuoApi.md#DuoApplicationList) | **Get** /duo/accounts/{account_id}/applications | List Duo Applications
[**DuoApplicationPost**](DuoApi.md#DuoApplicationPost) | **Post** /duo/accounts/{account_id}/applications | Create Duo Application
[**DuoApplicationUpdate**](DuoApi.md#DuoApplicationUpdate) | **Put** /duo/accounts/{account_id}/applications/{application_id} | Update Duo Application



## DuoAccountDelete

> DuoAccount DuoAccountDelete(ctx, id).XOrgId(xOrgId).Execute()

Delete a Duo Account



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
    id := "id_example" // string | ObjectID of the Duo Account
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DuoApi.DuoAccountDelete(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DuoApi.DuoAccountDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuoAccountDelete`: DuoAccount
    fmt.Fprintf(os.Stdout, "Response from `DuoApi.DuoAccountDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of the Duo Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuoAccountDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoAccount**](DuoAccount.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuoAccountGet

> DuoAccount DuoAccountGet(ctx, id).XOrgId(xOrgId).Execute()

Get a Duo Acount



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
    id := "id_example" // string | ObjectID of the Duo Account
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DuoApi.DuoAccountGet(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DuoApi.DuoAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuoAccountGet`: DuoAccount
    fmt.Fprintf(os.Stdout, "Response from `DuoApi.DuoAccountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of the Duo Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuoAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoAccount**](DuoAccount.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuoAccountList

> []DuoAccount DuoAccountList(ctx).XOrgId(xOrgId).Execute()

List Duo Accounts



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DuoApi.DuoAccountList(context.Background()).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DuoApi.DuoAccountList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuoAccountList`: []DuoAccount
    fmt.Fprintf(os.Stdout, "Response from `DuoApi.DuoAccountList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDuoAccountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]DuoAccount**](DuoAccount.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuoAccountPost

> DuoAccount DuoAccountPost(ctx).XOrgId(xOrgId).Execute()

Create Duo Account



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DuoApi.DuoAccountPost(context.Background()).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DuoApi.DuoAccountPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuoAccountPost`: DuoAccount
    fmt.Fprintf(os.Stdout, "Response from `DuoApi.DuoAccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDuoAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoAccount**](DuoAccount.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuoApplicationDelete

> DuoApplication DuoApplicationDelete(ctx, accountId, applicationId).XOrgId(xOrgId).Execute()

Delete a Duo Application



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
    accountId := "accountId_example" // string | 
    applicationId := "applicationId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DuoApi.DuoApplicationDelete(context.Background(), accountId, applicationId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DuoApi.DuoApplicationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuoApplicationDelete`: DuoApplication
    fmt.Fprintf(os.Stdout, "Response from `DuoApi.DuoApplicationDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuoApplicationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuoApplicationGet

> DuoApplication DuoApplicationGet(ctx, accountId, applicationId).XOrgId(xOrgId).Execute()

Get a Duo application



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
    accountId := "accountId_example" // string | 
    applicationId := "applicationId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DuoApi.DuoApplicationGet(context.Background(), accountId, applicationId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DuoApi.DuoApplicationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuoApplicationGet`: DuoApplication
    fmt.Fprintf(os.Stdout, "Response from `DuoApi.DuoApplicationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuoApplicationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuoApplicationList

> []DuoApplication DuoApplicationList(ctx, accountId).XOrgId(xOrgId).Execute()

List Duo Applications



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
    accountId := "accountId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DuoApi.DuoApplicationList(context.Background(), accountId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DuoApi.DuoApplicationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuoApplicationList`: []DuoApplication
    fmt.Fprintf(os.Stdout, "Response from `DuoApi.DuoApplicationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuoApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuoApplicationPost

> DuoApplication DuoApplicationPost(ctx, accountId).XOrgId(xOrgId).Body(body).Execute()

Create Duo Application



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
    accountId := "accountId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewDuoApplicationReq("ApiHost_example", "IntegrationKey_example", "Name_example", "SecretKey_example") // DuoApplicationReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DuoApi.DuoApplicationPost(context.Background(), accountId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DuoApi.DuoApplicationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuoApplicationPost`: DuoApplication
    fmt.Fprintf(os.Stdout, "Response from `DuoApi.DuoApplicationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuoApplicationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**DuoApplicationReq**](DuoApplicationReq.md) |  | 

### Return type

[**DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuoApplicationUpdate

> DuoApplication DuoApplicationUpdate(ctx, accountId, applicationId).XOrgId(xOrgId).Body(body).Execute()

Update Duo Application



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
    accountId := "accountId_example" // string | 
    applicationId := "applicationId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewDuoApplicationUpdateReq("ApiHost_example", "IntegrationKey_example", "Name_example") // DuoApplicationUpdateReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DuoApi.DuoApplicationUpdate(context.Background(), accountId, applicationId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DuoApi.DuoApplicationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuoApplicationUpdate`: DuoApplication
    fmt.Fprintf(os.Stdout, "Response from `DuoApi.DuoApplicationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuoApplicationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**DuoApplicationUpdateReq**](DuoApplicationUpdateReq.md) |  | 

### Return type

[**DuoApplication**](DuoApplication.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


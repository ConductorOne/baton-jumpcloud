# \SoftwareAppsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphSoftwareappsAssociationsList**](SoftwareAppsApi.md#GraphSoftwareappsAssociationsList) | **Get** /softwareapps/{software_app_id}/associations | List the associations of a Software Application
[**GraphSoftwareappsAssociationsPost**](SoftwareAppsApi.md#GraphSoftwareappsAssociationsPost) | **Post** /softwareapps/{software_app_id}/associations | Manage the associations of a software application.
[**GraphSoftwareappsTraverseSystem**](SoftwareAppsApi.md#GraphSoftwareappsTraverseSystem) | **Get** /softwareapps/{software_app_id}/systems | List the Systems bound to a Software App.
[**GraphSoftwareappsTraverseSystemGroup**](SoftwareAppsApi.md#GraphSoftwareappsTraverseSystemGroup) | **Get** /softwareapps/{software_app_id}/systemgroups | List the System Groups bound to a Software App.
[**SoftwareAppStatusesList**](SoftwareAppsApi.md#SoftwareAppStatusesList) | **Get** /softwareapps/{software_app_id}/statuses | Get the status of the provided Software Application
[**SoftwareAppsDelete**](SoftwareAppsApi.md#SoftwareAppsDelete) | **Delete** /softwareapps/{id} | Delete a configured Software Application
[**SoftwareAppsGet**](SoftwareAppsApi.md#SoftwareAppsGet) | **Get** /softwareapps/{id} | Retrieve a configured Software Application.
[**SoftwareAppsList**](SoftwareAppsApi.md#SoftwareAppsList) | **Get** /softwareapps | Get all configured Software Applications.
[**SoftwareAppsPost**](SoftwareAppsApi.md#SoftwareAppsPost) | **Post** /softwareapps | Create a Software Application that will be managed by JumpCloud.
[**SoftwareAppsReclaimLicenses**](SoftwareAppsApi.md#SoftwareAppsReclaimLicenses) | **Post** /softwareapps/{software_app_id}/reclaim-licenses | Reclaim Licenses for a Software Application.
[**SoftwareAppsRetryInstallation**](SoftwareAppsApi.md#SoftwareAppsRetryInstallation) | **Post** /softwareapps/{software_app_id}/retry-installation | Retry Installation for a Software Application
[**SoftwareAppsUpdate**](SoftwareAppsApi.md#SoftwareAppsUpdate) | **Put** /softwareapps/{id} | Update a Software Application Configuration.



## GraphSoftwareappsAssociationsList

> []GraphConnection GraphSoftwareappsAssociationsList(ctx, softwareAppId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a Software Application



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
    softwareAppId := "softwareAppId_example" // string | ObjectID of the Software App.
    targets := []string{"Targets_example"} // []string | Targets which a \"software_app\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareAppsApi.GraphSoftwareappsAssociationsList(context.Background(), softwareAppId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.GraphSoftwareappsAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSoftwareappsAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `SoftwareAppsApi.GraphSoftwareappsAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softwareAppId** | **string** | ObjectID of the Software App. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSoftwareappsAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;software_app\&quot; can be associated to. | 
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]GraphConnection**](GraphConnection.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphSoftwareappsAssociationsPost

> GraphSoftwareappsAssociationsPost(ctx, softwareAppId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a software application.



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
    softwareAppId := "softwareAppId_example" // string | ObjectID of the Software App.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationSoftwareApp("Id_example", "Op_example", "Type_example") // GraphOperationSoftwareApp |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SoftwareAppsApi.GraphSoftwareappsAssociationsPost(context.Background(), softwareAppId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.GraphSoftwareappsAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softwareAppId** | **string** | ObjectID of the Software App. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSoftwareappsAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationSoftwareApp**](GraphOperationSoftwareApp.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphSoftwareappsTraverseSystem

> []GraphObjectWithPaths GraphSoftwareappsTraverseSystem(ctx, softwareAppId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Systems bound to a Software App.



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
    softwareAppId := "softwareAppId_example" // string | ObjectID of the Software App.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareAppsApi.GraphSoftwareappsTraverseSystem(context.Background(), softwareAppId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.GraphSoftwareappsTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSoftwareappsTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `SoftwareAppsApi.GraphSoftwareappsTraverseSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softwareAppId** | **string** | ObjectID of the Software App. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSoftwareappsTraverseSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphSoftwareappsTraverseSystemGroup

> []GraphObjectWithPaths GraphSoftwareappsTraverseSystemGroup(ctx, softwareAppId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the System Groups bound to a Software App.



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
    softwareAppId := "softwareAppId_example" // string | ObjectID of the Software App.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareAppsApi.GraphSoftwareappsTraverseSystemGroup(context.Background(), softwareAppId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.GraphSoftwareappsTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSoftwareappsTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `SoftwareAppsApi.GraphSoftwareappsTraverseSystemGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softwareAppId** | **string** | ObjectID of the Software App. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSoftwareappsTraverseSystemGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareAppStatusesList

> []SoftwareAppStatus SoftwareAppStatusesList(ctx, softwareAppId).XOrgId(xOrgId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Get the status of the provided Software Application



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
    softwareAppId := "softwareAppId_example" // string | ObjectID of the Software App.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareAppsApi.SoftwareAppStatusesList(context.Background(), softwareAppId).XOrgId(xOrgId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.SoftwareAppStatusesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareAppStatusesList`: []SoftwareAppStatus
    fmt.Fprintf(os.Stdout, "Response from `SoftwareAppsApi.SoftwareAppStatusesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softwareAppId** | **string** | ObjectID of the Software App. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareAppStatusesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**[]SoftwareAppStatus**](SoftwareAppStatus.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareAppsDelete

> SoftwareAppsDelete(ctx, id).XOrgId(xOrgId).Execute()

Delete a configured Software Application



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
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SoftwareAppsApi.SoftwareAppsDelete(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.SoftwareAppsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareAppsDeleteRequest struct via the builder pattern


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


## SoftwareAppsGet

> SoftwareApp SoftwareAppsGet(ctx, id).XOrgId(xOrgId).Execute()

Retrieve a configured Software Application.



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
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareAppsApi.SoftwareAppsGet(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.SoftwareAppsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareAppsGet`: SoftwareApp
    fmt.Fprintf(os.Stdout, "Response from `SoftwareAppsApi.SoftwareAppsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareAppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**SoftwareApp**](SoftwareApp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareAppsList

> []SoftwareApp SoftwareAppsList(ctx).XOrgId(xOrgId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Get all configured Software Applications.



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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareAppsApi.SoftwareAppsList(context.Background()).XOrgId(xOrgId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.SoftwareAppsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareAppsList`: []SoftwareApp
    fmt.Fprintf(os.Stdout, "Response from `SoftwareAppsApi.SoftwareAppsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareAppsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**[]SoftwareApp**](SoftwareApp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareAppsPost

> SoftwareApp SoftwareAppsPost(ctx).XOrgId(xOrgId).Body(body).Execute()

Create a Software Application that will be managed by JumpCloud.



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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewSoftwareApp() // SoftwareApp |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareAppsApi.SoftwareAppsPost(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.SoftwareAppsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareAppsPost`: SoftwareApp
    fmt.Fprintf(os.Stdout, "Response from `SoftwareAppsApi.SoftwareAppsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareAppsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**SoftwareApp**](SoftwareApp.md) |  | 

### Return type

[**SoftwareApp**](SoftwareApp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareAppsReclaimLicenses

> SoftwareAppReclaimLicenses SoftwareAppsReclaimLicenses(ctx, softwareAppId).Execute()

Reclaim Licenses for a Software Application.



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
    softwareAppId := "softwareAppId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareAppsApi.SoftwareAppsReclaimLicenses(context.Background(), softwareAppId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.SoftwareAppsReclaimLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareAppsReclaimLicenses`: SoftwareAppReclaimLicenses
    fmt.Fprintf(os.Stdout, "Response from `SoftwareAppsApi.SoftwareAppsReclaimLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softwareAppId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareAppsReclaimLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoftwareAppReclaimLicenses**](SoftwareAppReclaimLicenses.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareAppsRetryInstallation

> SoftwareAppsRetryInstallation(ctx, softwareAppId).Body(body).Execute()

Retry Installation for a Software Application



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
    softwareAppId := "softwareAppId_example" // string | 
    body := *openapiclient.NewSoftwareAppsRetryInstallationRequest() // SoftwareAppsRetryInstallationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SoftwareAppsApi.SoftwareAppsRetryInstallation(context.Background(), softwareAppId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.SoftwareAppsRetryInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softwareAppId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareAppsRetryInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SoftwareAppsRetryInstallationRequest**](SoftwareAppsRetryInstallationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoftwareAppsUpdate

> SoftwareApp SoftwareAppsUpdate(ctx, id).XOrgId(xOrgId).Body(body).Execute()

Update a Software Application Configuration.



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
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewSoftwareApp() // SoftwareApp |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwareAppsApi.SoftwareAppsUpdate(context.Background(), id).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwareAppsApi.SoftwareAppsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareAppsUpdate`: SoftwareApp
    fmt.Fprintf(os.Stdout, "Response from `SoftwareAppsApi.SoftwareAppsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareAppsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**SoftwareApp**](SoftwareApp.md) |  | 

### Return type

[**SoftwareApp**](SoftwareApp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


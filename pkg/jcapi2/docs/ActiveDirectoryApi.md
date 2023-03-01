# \ActiveDirectoryApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivedirectoriesAgentsDelete**](ActiveDirectoryApi.md#ActivedirectoriesAgentsDelete) | **Delete** /activedirectories/{activedirectory_id}/agents/{agent_id} | Delete Active Directory Agent
[**ActivedirectoriesAgentsGet**](ActiveDirectoryApi.md#ActivedirectoriesAgentsGet) | **Get** /activedirectories/{activedirectory_id}/agents/{agent_id} | Get Active Directory Agent
[**ActivedirectoriesAgentsList**](ActiveDirectoryApi.md#ActivedirectoriesAgentsList) | **Get** /activedirectories/{activedirectory_id}/agents | List Active Directory Agents
[**ActivedirectoriesAgentsPost**](ActiveDirectoryApi.md#ActivedirectoriesAgentsPost) | **Post** /activedirectories/{activedirectory_id}/agents | Create a new Active Directory Agent
[**ActivedirectoriesDelete**](ActiveDirectoryApi.md#ActivedirectoriesDelete) | **Delete** /activedirectories/{id} | Delete an Active Directory
[**ActivedirectoriesGet**](ActiveDirectoryApi.md#ActivedirectoriesGet) | **Get** /activedirectories/{id} | Get an Active Directory
[**ActivedirectoriesList**](ActiveDirectoryApi.md#ActivedirectoriesList) | **Get** /activedirectories | List Active Directories
[**ActivedirectoriesPost**](ActiveDirectoryApi.md#ActivedirectoriesPost) | **Post** /activedirectories | Create a new Active Directory
[**GraphActiveDirectoryAssociationsList**](ActiveDirectoryApi.md#GraphActiveDirectoryAssociationsList) | **Get** /activedirectories/{activedirectory_id}/associations | List the associations of an Active Directory instance
[**GraphActiveDirectoryAssociationsPost**](ActiveDirectoryApi.md#GraphActiveDirectoryAssociationsPost) | **Post** /activedirectories/{activedirectory_id}/associations | Manage the associations of an Active Directory instance
[**GraphActiveDirectoryTraverseUser**](ActiveDirectoryApi.md#GraphActiveDirectoryTraverseUser) | **Get** /activedirectories/{activedirectory_id}/users | List the Users bound to an Active Directory instance
[**GraphActiveDirectoryTraverseUserGroup**](ActiveDirectoryApi.md#GraphActiveDirectoryTraverseUserGroup) | **Get** /activedirectories/{activedirectory_id}/usergroups | List the User Groups bound to an Active Directory instance



## ActivedirectoriesAgentsDelete

> ActivedirectoriesAgentsDelete(ctx, activedirectoryId, agentId).XOrgId(xOrgId).Execute()

Delete Active Directory Agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | 
    agentId := "agentId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveDirectoryApi.ActivedirectoriesAgentsDelete(context.Background(), activedirectoryId, agentId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.ActivedirectoriesAgentsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activedirectoryId** | **string** |  | 
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoriesAgentsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivedirectoriesAgentsGet

> ActiveDirectoryAgentList ActivedirectoriesAgentsGet(ctx, activedirectoryId, agentId).XOrgId(xOrgId).Execute()

Get Active Directory Agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | 
    agentId := "agentId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.ActivedirectoriesAgentsGet(context.Background(), activedirectoryId, agentId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.ActivedirectoriesAgentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivedirectoriesAgentsGet`: ActiveDirectoryAgentList
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.ActivedirectoriesAgentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activedirectoryId** | **string** |  | 
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoriesAgentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**ActiveDirectoryAgentList**](ActiveDirectoryAgentList.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivedirectoriesAgentsList

> []ActiveDirectoryAgentList ActivedirectoriesAgentsList(ctx, activedirectoryId).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()

List Active Directory Agents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | 
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.ActivedirectoriesAgentsList(context.Background(), activedirectoryId).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.ActivedirectoriesAgentsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivedirectoriesAgentsList`: []ActiveDirectoryAgentList
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.ActivedirectoriesAgentsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activedirectoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoriesAgentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]ActiveDirectoryAgentList**](ActiveDirectoryAgentList.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivedirectoriesAgentsPost

> ActiveDirectoryAgentGet ActivedirectoriesAgentsPost(ctx, activedirectoryId).XOrgId(xOrgId).Body(body).Execute()

Create a new Active Directory Agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.ActivedirectoriesAgentsPost(context.Background(), activedirectoryId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.ActivedirectoriesAgentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivedirectoriesAgentsPost`: ActiveDirectoryAgentGet
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.ActivedirectoriesAgentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activedirectoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoriesAgentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**ActiveDirectoryAgentGet**](ActiveDirectoryAgentGet.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivedirectoriesDelete

> ActiveDirectory ActivedirectoriesDelete(ctx, id).XOrgId(xOrgId).Execute()

Delete an Active Directory



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    id := "id_example" // string | ObjectID of this Active Directory instance.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.ActivedirectoriesDelete(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.ActivedirectoriesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivedirectoriesDelete`: ActiveDirectory
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.ActivedirectoriesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of this Active Directory instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoriesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**ActiveDirectory**](ActiveDirectory.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivedirectoriesGet

> ActiveDirectory ActivedirectoriesGet(ctx, id).XOrgId(xOrgId).Execute()

Get an Active Directory



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    id := "id_example" // string | ObjectID of this Active Directory instance.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.ActivedirectoriesGet(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.ActivedirectoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivedirectoriesGet`: ActiveDirectory
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.ActivedirectoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of this Active Directory instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**ActiveDirectory**](ActiveDirectory.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivedirectoriesList

> []ActiveDirectory ActivedirectoriesList(ctx).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()

List Active Directories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.ActivedirectoriesList(context.Background()).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.ActivedirectoriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivedirectoriesList`: []ActiveDirectory
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.ActivedirectoriesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]ActiveDirectory**](ActiveDirectory.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivedirectoriesPost

> ActiveDirectory ActivedirectoriesPost(ctx).XOrgId(xOrgId).Body(body).Execute()

Create a new Active Directory



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewActiveDirectory() // ActiveDirectory |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.ActivedirectoriesPost(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.ActivedirectoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivedirectoriesPost`: ActiveDirectory
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.ActivedirectoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivedirectoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**ActiveDirectory**](ActiveDirectory.md) |  | 

### Return type

[**ActiveDirectory**](ActiveDirectory.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GraphActiveDirectoryAssociationsList

> []GraphConnection GraphActiveDirectoryAssociationsList(ctx, activedirectoryId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of an Active Directory instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | 
    targets := []string{"Targets_example"} // []string | Targets which a \"active_directory\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.GraphActiveDirectoryAssociationsList(context.Background(), activedirectoryId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.GraphActiveDirectoryAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphActiveDirectoryAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.GraphActiveDirectoryAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activedirectoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphActiveDirectoryAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;active_directory\&quot; can be associated to. | 
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


## GraphActiveDirectoryAssociationsPost

> GraphActiveDirectoryAssociationsPost(ctx, activedirectoryId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of an Active Directory instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationActiveDirectory("Id_example", "Op_example", "Type_example") // GraphOperationActiveDirectory |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveDirectoryApi.GraphActiveDirectoryAssociationsPost(context.Background(), activedirectoryId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.GraphActiveDirectoryAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activedirectoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphActiveDirectoryAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationActiveDirectory**](GraphOperationActiveDirectory.md) |  | 

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


## GraphActiveDirectoryTraverseUser

> []GraphObjectWithPaths GraphActiveDirectoryTraverseUser(ctx, activedirectoryId).Filter(filter).Limit(limit).XOrgId(xOrgId).Skip(skip).Execute()

List the Users bound to an Active Directory instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | ObjectID of the Active Directory instance.
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.GraphActiveDirectoryTraverseUser(context.Background(), activedirectoryId).Filter(filter).Limit(limit).XOrgId(xOrgId).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.GraphActiveDirectoryTraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphActiveDirectoryTraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.GraphActiveDirectoryTraverseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activedirectoryId** | **string** | ObjectID of the Active Directory instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphActiveDirectoryTraverseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **skip** | **int32** | The offset into the records to return. | [default to 0]

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


## GraphActiveDirectoryTraverseUserGroup

> []GraphObjectWithPaths GraphActiveDirectoryTraverseUserGroup(ctx, activedirectoryId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the User Groups bound to an Active Directory instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | ObjectID of the Active Directory instance.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveDirectoryApi.GraphActiveDirectoryTraverseUserGroup(context.Background(), activedirectoryId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.GraphActiveDirectoryTraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphActiveDirectoryTraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.GraphActiveDirectoryTraverseUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activedirectoryId** | **string** | ObjectID of the Active Directory instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphActiveDirectoryTraverseUserGroupRequest struct via the builder pattern


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


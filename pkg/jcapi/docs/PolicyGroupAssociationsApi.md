# \PolicyGroupAssociationsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphPolicyGroupAssociationsList**](PolicyGroupAssociationsApi.md#GraphPolicyGroupAssociationsList) | **Get** /policygroups/{group_id}/associations | List the associations of a Policy Group.
[**GraphPolicyGroupAssociationsPost**](PolicyGroupAssociationsApi.md#GraphPolicyGroupAssociationsPost) | **Post** /policygroups/{group_id}/associations | Manage the associations of a Policy Group
[**GraphPolicyGroupTraverseSystem**](PolicyGroupAssociationsApi.md#GraphPolicyGroupTraverseSystem) | **Get** /policygroups/{group_id}/systems | List the Systems bound to a Policy Group
[**GraphPolicyGroupTraverseSystemGroup**](PolicyGroupAssociationsApi.md#GraphPolicyGroupTraverseSystemGroup) | **Get** /policygroups/{group_id}/systemgroups | List the System Groups bound to Policy Groups



## GraphPolicyGroupAssociationsList

> []GraphConnection GraphPolicyGroupAssociationsList(ctx, groupId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a Policy Group.



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
    groupId := "groupId_example" // string | ObjectID of the Policy Group.
    targets := []string{"Targets_example"} // []string | Targets which a \"policy_group\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupAssociationsApi.GraphPolicyGroupAssociationsList(context.Background(), groupId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupAssociationsApi.GraphPolicyGroupAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupAssociationsApi.GraphPolicyGroupAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the Policy Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyGroupAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;policy_group\&quot; can be associated to. | 
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


## GraphPolicyGroupAssociationsPost

> GraphPolicyGroupAssociationsPost(ctx, groupId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a Policy Group



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
    groupId := "groupId_example" // string | ObjectID of the Policy Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationPolicyGroup("Id_example", "Op_example", "Type_example") // GraphOperationPolicyGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyGroupAssociationsApi.GraphPolicyGroupAssociationsPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupAssociationsApi.GraphPolicyGroupAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the Policy Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyGroupAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationPolicyGroup**](GraphOperationPolicyGroup.md) |  | 

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


## GraphPolicyGroupTraverseSystem

> []GraphObjectWithPaths GraphPolicyGroupTraverseSystem(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Systems bound to a Policy Group



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
    groupId := "groupId_example" // string | ObjectID of the Policy Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupAssociationsApi.GraphPolicyGroupTraverseSystem(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupAssociationsApi.GraphPolicyGroupTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupAssociationsApi.GraphPolicyGroupTraverseSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the Policy Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyGroupTraverseSystemRequest struct via the builder pattern


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


## GraphPolicyGroupTraverseSystemGroup

> []GraphObjectWithPaths GraphPolicyGroupTraverseSystemGroup(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the System Groups bound to Policy Groups



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
    groupId := "groupId_example" // string | ObjectID of the Policy Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupAssociationsApi.GraphPolicyGroupTraverseSystemGroup(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupAssociationsApi.GraphPolicyGroupTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupAssociationsApi.GraphPolicyGroupTraverseSystemGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the Policy Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyGroupTraverseSystemGroupRequest struct via the builder pattern


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


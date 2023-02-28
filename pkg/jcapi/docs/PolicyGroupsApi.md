# \PolicyGroupsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphPolicyGroupAssociationsList**](PolicyGroupsApi.md#GraphPolicyGroupAssociationsList) | **Get** /policygroups/{group_id}/associations | List the associations of a Policy Group.
[**GraphPolicyGroupAssociationsPost**](PolicyGroupsApi.md#GraphPolicyGroupAssociationsPost) | **Post** /policygroups/{group_id}/associations | Manage the associations of a Policy Group
[**GraphPolicyGroupMembersList**](PolicyGroupsApi.md#GraphPolicyGroupMembersList) | **Get** /policygroups/{group_id}/members | List the members of a Policy Group
[**GraphPolicyGroupMembersPost**](PolicyGroupsApi.md#GraphPolicyGroupMembersPost) | **Post** /policygroups/{group_id}/members | Manage the members of a Policy Group
[**GraphPolicyGroupMembership**](PolicyGroupsApi.md#GraphPolicyGroupMembership) | **Get** /policygroups/{group_id}/membership | List the Policy Group&#39;s membership
[**GraphPolicyGroupTraverseSystem**](PolicyGroupsApi.md#GraphPolicyGroupTraverseSystem) | **Get** /policygroups/{group_id}/systems | List the Systems bound to a Policy Group
[**GraphPolicyGroupTraverseSystemGroup**](PolicyGroupsApi.md#GraphPolicyGroupTraverseSystemGroup) | **Get** /policygroups/{group_id}/systemgroups | List the System Groups bound to Policy Groups
[**GroupsPolicyDelete**](PolicyGroupsApi.md#GroupsPolicyDelete) | **Delete** /policygroups/{id} | Delete a Policy Group
[**GroupsPolicyGet**](PolicyGroupsApi.md#GroupsPolicyGet) | **Get** /policygroups/{id} | View an individual Policy Group details
[**GroupsPolicyList**](PolicyGroupsApi.md#GroupsPolicyList) | **Get** /policygroups | List all Policy Groups
[**GroupsPolicyPost**](PolicyGroupsApi.md#GroupsPolicyPost) | **Post** /policygroups | Create a new Policy Group
[**GroupsPolicyPut**](PolicyGroupsApi.md#GroupsPolicyPut) | **Put** /policygroups/{id} | Update a Policy Group



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
    resp, r, err := apiClient.PolicyGroupsApi.GraphPolicyGroupAssociationsList(context.Background(), groupId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GraphPolicyGroupAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GraphPolicyGroupAssociationsList`: %v\n", resp)
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
    r, err := apiClient.PolicyGroupsApi.GraphPolicyGroupAssociationsPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GraphPolicyGroupAssociationsPost``: %v\n", err)
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


## GraphPolicyGroupMembersList

> []GraphConnection GraphPolicyGroupMembersList(ctx, groupId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the members of a Policy Group



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
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupsApi.GraphPolicyGroupMembersList(context.Background(), groupId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GraphPolicyGroupMembersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupMembersList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GraphPolicyGroupMembersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the Policy Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyGroupMembersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GraphPolicyGroupMembersPost

> GraphPolicyGroupMembersPost(ctx, groupId).XOrgId(xOrgId).Body(body).Execute()

Manage the members of a Policy Group



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
    body := *openapiclient.NewGraphOperationPolicyGroupMember("Id_example", "Op_example", "Type_example") // GraphOperationPolicyGroupMember |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyGroupsApi.GraphPolicyGroupMembersPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GraphPolicyGroupMembersPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGraphPolicyGroupMembersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationPolicyGroupMember**](GraphOperationPolicyGroupMember.md) |  | 

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


## GraphPolicyGroupMembership

> []GraphObjectWithPaths GraphPolicyGroupMembership(ctx, groupId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()

List the Policy Group's membership



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
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupsApi.GraphPolicyGroupMembership(context.Background(), groupId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GraphPolicyGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupMembership`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GraphPolicyGroupMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the Policy Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyGroupMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

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
    resp, r, err := apiClient.PolicyGroupsApi.GraphPolicyGroupTraverseSystem(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GraphPolicyGroupTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GraphPolicyGroupTraverseSystem`: %v\n", resp)
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
    resp, r, err := apiClient.PolicyGroupsApi.GraphPolicyGroupTraverseSystemGroup(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GraphPolicyGroupTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GraphPolicyGroupTraverseSystemGroup`: %v\n", resp)
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


## GroupsPolicyDelete

> PolicyGroup GroupsPolicyDelete(ctx, id).XOrgId(xOrgId).Execute()

Delete a Policy Group



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
    id := "id_example" // string | ObjectID of the Policy Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupsApi.GroupsPolicyDelete(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GroupsPolicyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsPolicyDelete`: PolicyGroup
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GroupsPolicyDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of the Policy Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPolicyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsPolicyGet

> PolicyGroup GroupsPolicyGet(ctx, id).XOrgId(xOrgId).Execute()

View an individual Policy Group details



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
    id := "id_example" // string | ObjectID of the Policy Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupsApi.GroupsPolicyGet(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GroupsPolicyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsPolicyGet`: PolicyGroup
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GroupsPolicyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of the Policy Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPolicyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsPolicyList

> []PolicyGroup GroupsPolicyList(ctx).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()

List all Policy Groups



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
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupsApi.GroupsPolicyList(context.Background()).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GroupsPolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsPolicyList`: []PolicyGroup
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GroupsPolicyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsPolicyPost

> PolicyGroup GroupsPolicyPost(ctx).XOrgId(xOrgId).Body(body).Execute()

Create a new Policy Group



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
    body := *openapiclient.NewPolicyGroupData("Name_example") // PolicyGroupData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupsApi.GroupsPolicyPost(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GroupsPolicyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsPolicyPost`: PolicyGroup
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GroupsPolicyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPolicyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**PolicyGroupData**](PolicyGroupData.md) |  | 

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsPolicyPut

> PolicyGroup GroupsPolicyPut(ctx, id).XOrgId(xOrgId).Body(body).Execute()

Update a Policy Group



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
    id := "id_example" // string | ObjectID of the Policy Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewPolicyGroupData("Name_example") // PolicyGroupData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupsApi.GroupsPolicyPut(context.Background(), id).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupsApi.GroupsPolicyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsPolicyPut`: PolicyGroup
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupsApi.GroupsPolicyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of the Policy Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPolicyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**PolicyGroupData**](PolicyGroupData.md) |  | 

### Return type

[**PolicyGroup**](PolicyGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


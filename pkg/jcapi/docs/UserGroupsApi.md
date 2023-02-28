# \UserGroupsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphUserGroupAssociationsList**](UserGroupsApi.md#GraphUserGroupAssociationsList) | **Get** /usergroups/{group_id}/associations | List the associations of a User Group.
[**GraphUserGroupAssociationsPost**](UserGroupsApi.md#GraphUserGroupAssociationsPost) | **Post** /usergroups/{group_id}/associations | Manage the associations of a User Group
[**GraphUserGroupMembersList**](UserGroupsApi.md#GraphUserGroupMembersList) | **Get** /usergroups/{group_id}/members | List the members of a User Group
[**GraphUserGroupMembersPost**](UserGroupsApi.md#GraphUserGroupMembersPost) | **Post** /usergroups/{group_id}/members | Manage the members of a User Group
[**GraphUserGroupMembership**](UserGroupsApi.md#GraphUserGroupMembership) | **Get** /usergroups/{group_id}/membership | List the User Group&#39;s membership
[**GraphUserGroupTraverseActiveDirectory**](UserGroupsApi.md#GraphUserGroupTraverseActiveDirectory) | **Get** /usergroups/{group_id}/activedirectories | List the Active Directories bound to a User Group
[**GraphUserGroupTraverseApplication**](UserGroupsApi.md#GraphUserGroupTraverseApplication) | **Get** /usergroups/{group_id}/applications | List the Applications bound to a User Group
[**GraphUserGroupTraverseDirectory**](UserGroupsApi.md#GraphUserGroupTraverseDirectory) | **Get** /usergroups/{group_id}/directories | List the Directories bound to a User Group
[**GraphUserGroupTraverseGSuite**](UserGroupsApi.md#GraphUserGroupTraverseGSuite) | **Get** /usergroups/{group_id}/gsuites | List the G Suite instances bound to a User Group
[**GraphUserGroupTraverseLdapServer**](UserGroupsApi.md#GraphUserGroupTraverseLdapServer) | **Get** /usergroups/{group_id}/ldapservers | List the LDAP Servers bound to a User Group
[**GraphUserGroupTraverseOffice365**](UserGroupsApi.md#GraphUserGroupTraverseOffice365) | **Get** /usergroups/{group_id}/office365s | List the Office 365 instances bound to a User Group
[**GraphUserGroupTraverseRadiusServer**](UserGroupsApi.md#GraphUserGroupTraverseRadiusServer) | **Get** /usergroups/{group_id}/radiusservers | List the RADIUS Servers bound to a User Group
[**GraphUserGroupTraverseSystem**](UserGroupsApi.md#GraphUserGroupTraverseSystem) | **Get** /usergroups/{group_id}/systems | List the Systems bound to a User Group
[**GraphUserGroupTraverseSystemGroup**](UserGroupsApi.md#GraphUserGroupTraverseSystemGroup) | **Get** /usergroups/{group_id}/systemgroups | List the System Groups bound to User Groups
[**GroupsSuggestionsGet**](UserGroupsApi.md#GroupsSuggestionsGet) | **Get** /usergroups/{group_id}/suggestions | List Suggestions for a User Group
[**GroupsSuggestionsPost**](UserGroupsApi.md#GroupsSuggestionsPost) | **Post** /usergroups/{group_id}/suggestions | List Suggestions for a User Group
[**GroupsUserDelete**](UserGroupsApi.md#GroupsUserDelete) | **Delete** /usergroups/{id} | Delete a User Group
[**GroupsUserGet**](UserGroupsApi.md#GroupsUserGet) | **Get** /usergroups/{id} | View an individual User Group details
[**GroupsUserList**](UserGroupsApi.md#GroupsUserList) | **Get** /usergroups | List all User Groups
[**GroupsUserPost**](UserGroupsApi.md#GroupsUserPost) | **Post** /usergroups | Create a new User Group
[**GroupsUserPut**](UserGroupsApi.md#GroupsUserPut) | **Put** /usergroups/{id} | Update a User Group



## GraphUserGroupAssociationsList

> []GraphConnection GraphUserGroupAssociationsList(ctx, groupId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a User Group.



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    targets := []string{"Targets_example"} // []string | Targets which a \"user_group\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupAssociationsList(context.Background(), groupId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;user_group\&quot; can be associated to. | 
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


## GraphUserGroupAssociationsPost

> GraphUserGroupAssociationsPost(ctx, groupId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationUserGroup("Id_example", "Op_example", "Type_example") // GraphOperationUserGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserGroupsApi.GraphUserGroupAssociationsPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationUserGroup**](GraphOperationUserGroup.md) |  | 

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


## GraphUserGroupMembersList

> []GraphConnection GraphUserGroupMembersList(ctx, groupId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the members of a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupMembersList(context.Background(), groupId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupMembersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupMembersList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupMembersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupMembersListRequest struct via the builder pattern


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


## GraphUserGroupMembersPost

> GraphUserGroupMembersPost(ctx, groupId).XOrgId(xOrgId).Body(body).Execute()

Manage the members of a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationUserGroupMember("Id_example", "Op_example", "Type_example") // GraphOperationUserGroupMember |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserGroupsApi.GraphUserGroupMembersPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupMembersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupMembersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationUserGroupMember**](GraphOperationUserGroupMember.md) |  | 

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


## GraphUserGroupMembership

> []GraphObjectWithPaths GraphUserGroupMembership(ctx, groupId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()

List the User Group's membership



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupMembership(context.Background(), groupId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupMembership`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupMembershipRequest struct via the builder pattern


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


## GraphUserGroupTraverseActiveDirectory

> []GraphObjectWithPaths GraphUserGroupTraverseActiveDirectory(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Active Directories bound to a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupTraverseActiveDirectory(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupTraverseActiveDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseActiveDirectory`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupTraverseActiveDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupTraverseActiveDirectoryRequest struct via the builder pattern


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


## GraphUserGroupTraverseApplication

> []GraphObjectWithPaths GraphUserGroupTraverseApplication(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Applications bound to a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupTraverseApplication(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupTraverseApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseApplication`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupTraverseApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupTraverseApplicationRequest struct via the builder pattern


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


## GraphUserGroupTraverseDirectory

> []GraphObjectWithPaths GraphUserGroupTraverseDirectory(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Directories bound to a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupTraverseDirectory(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupTraverseDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseDirectory`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupTraverseDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupTraverseDirectoryRequest struct via the builder pattern


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


## GraphUserGroupTraverseGSuite

> []GraphObjectWithPaths GraphUserGroupTraverseGSuite(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the G Suite instances bound to a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupTraverseGSuite(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupTraverseGSuite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseGSuite`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupTraverseGSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupTraverseGSuiteRequest struct via the builder pattern


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


## GraphUserGroupTraverseLdapServer

> []GraphObjectWithPaths GraphUserGroupTraverseLdapServer(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the LDAP Servers bound to a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupTraverseLdapServer(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupTraverseLdapServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseLdapServer`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupTraverseLdapServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupTraverseLdapServerRequest struct via the builder pattern


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


## GraphUserGroupTraverseOffice365

> []GraphObjectWithPaths GraphUserGroupTraverseOffice365(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Office 365 instances bound to a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupTraverseOffice365(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupTraverseOffice365``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseOffice365`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupTraverseOffice365`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupTraverseOffice365Request struct via the builder pattern


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


## GraphUserGroupTraverseRadiusServer

> []GraphObjectWithPaths GraphUserGroupTraverseRadiusServer(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the RADIUS Servers bound to a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupTraverseRadiusServer(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupTraverseRadiusServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseRadiusServer`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupTraverseRadiusServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupTraverseRadiusServerRequest struct via the builder pattern


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


## GraphUserGroupTraverseSystem

> []GraphObjectWithPaths GraphUserGroupTraverseSystem(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Systems bound to a User Group



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupTraverseSystem(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupTraverseSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupTraverseSystemRequest struct via the builder pattern


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


## GraphUserGroupTraverseSystemGroup

> []GraphObjectWithPaths GraphUserGroupTraverseSystemGroup(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the System Groups bound to User Groups



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
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GraphUserGroupTraverseSystemGroup(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GraphUserGroupTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GraphUserGroupTraverseSystemGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserGroupTraverseSystemGroupRequest struct via the builder pattern


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


## GroupsSuggestionsGet

> []MemberSuggestion GroupsSuggestionsGet(ctx, groupId).XOrgId(xOrgId).Limit(limit).Skip(skip).Execute()

List Suggestions for a User Group



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
    groupId := "groupId_example" // string | ID of the group
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GroupsSuggestionsGet(context.Background(), groupId).XOrgId(xOrgId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GroupsSuggestionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsSuggestionsGet`: []MemberSuggestion
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GroupsSuggestionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsSuggestionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]

### Return type

[**[]MemberSuggestion**](MemberSuggestion.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsSuggestionsPost

> []MemberSuggestionsPostResult GroupsSuggestionsPost(ctx, groupId).Body(body).XOrgId(xOrgId).Execute()

List Suggestions for a User Group



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
    groupId := "groupId_example" // string | ID of the group
    body := *openapiclient.NewGroupsSuggestionsPostRequest() // GroupsSuggestionsPostRequest | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GroupsSuggestionsPost(context.Background(), groupId).Body(body).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GroupsSuggestionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsSuggestionsPost`: []MemberSuggestionsPostResult
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GroupsSuggestionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsSuggestionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GroupsSuggestionsPostRequest**](GroupsSuggestionsPostRequest.md) |  | 
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]MemberSuggestionsPostResult**](MemberSuggestionsPostResult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUserDelete

> UserGroup GroupsUserDelete(ctx, id).XOrgId(xOrgId).Execute()

Delete a User Group



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
    id := "id_example" // string | ObjectID of the User Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GroupsUserDelete(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GroupsUserDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUserDelete`: UserGroup
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GroupsUserDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUserDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUserGet

> UserGroup GroupsUserGet(ctx, id).XOrgId(xOrgId).Execute()

View an individual User Group details



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
    id := "id_example" // string | ObjectID of the User Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GroupsUserGet(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GroupsUserGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUserGet`: UserGroup
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GroupsUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUserList

> []UserGroup GroupsUserList(ctx).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()

List all User Groups



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
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GroupsUserList(context.Background()).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GroupsUserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUserList`: []UserGroup
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GroupsUserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUserPost

> UserGroup GroupsUserPost(ctx).XOrgId(xOrgId).Body(body).Execute()

Create a new User Group



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
    body := *openapiclient.NewUserGroupPost("Name_example") // UserGroupPost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GroupsUserPost(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GroupsUserPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUserPost`: UserGroup
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GroupsUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**UserGroupPost**](UserGroupPost.md) |  | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUserPut

> UserGroup GroupsUserPut(ctx, id).XOrgId(xOrgId).Body(body).Execute()

Update a User Group



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
    id := "id_example" // string | ObjectID of the User Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewUserGroupPut("Name_example") // UserGroupPut |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserGroupsApi.GroupsUserPut(context.Background(), id).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.GroupsUserPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsUserPut`: UserGroup
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.GroupsUserPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ObjectID of the User Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUserPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**UserGroupPut**](UserGroupPut.md) |  | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


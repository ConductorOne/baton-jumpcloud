# \GraphApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphActiveDirectoryAssociationsList**](GraphApi.md#GraphActiveDirectoryAssociationsList) | **Get** /activedirectories/{activedirectory_id}/associations | List the associations of an Active Directory instance
[**GraphActiveDirectoryAssociationsPost**](GraphApi.md#GraphActiveDirectoryAssociationsPost) | **Post** /activedirectories/{activedirectory_id}/associations | Manage the associations of an Active Directory instance
[**GraphActiveDirectoryTraverseUser**](GraphApi.md#GraphActiveDirectoryTraverseUser) | **Get** /activedirectories/{activedirectory_id}/users | List the Users bound to an Active Directory instance
[**GraphActiveDirectoryTraverseUserGroup**](GraphApi.md#GraphActiveDirectoryTraverseUserGroup) | **Get** /activedirectories/{activedirectory_id}/usergroups | List the User Groups bound to an Active Directory instance
[**GraphApplicationAssociationsList**](GraphApi.md#GraphApplicationAssociationsList) | **Get** /applications/{application_id}/associations | List the associations of an Application
[**GraphApplicationAssociationsPost**](GraphApi.md#GraphApplicationAssociationsPost) | **Post** /applications/{application_id}/associations | Manage the associations of an Application
[**GraphApplicationTraverseUser**](GraphApi.md#GraphApplicationTraverseUser) | **Get** /applications/{application_id}/users | List the Users bound to an Application
[**GraphApplicationTraverseUserGroup**](GraphApi.md#GraphApplicationTraverseUserGroup) | **Get** /applications/{application_id}/usergroups | List the User Groups bound to an Application
[**GraphCommandAssociationsList**](GraphApi.md#GraphCommandAssociationsList) | **Get** /commands/{command_id}/associations | List the associations of a Command
[**GraphCommandAssociationsPost**](GraphApi.md#GraphCommandAssociationsPost) | **Post** /commands/{command_id}/associations | Manage the associations of a Command
[**GraphCommandTraverseSystem**](GraphApi.md#GraphCommandTraverseSystem) | **Get** /commands/{command_id}/systems | List the Systems bound to a Command
[**GraphCommandTraverseSystemGroup**](GraphApi.md#GraphCommandTraverseSystemGroup) | **Get** /commands/{command_id}/systemgroups | List the System Groups bound to a Command
[**GraphGSuiteAssociationsList**](GraphApi.md#GraphGSuiteAssociationsList) | **Get** /gsuites/{gsuite_id}/associations | List the associations of a G Suite instance
[**GraphGSuiteAssociationsPost**](GraphApi.md#GraphGSuiteAssociationsPost) | **Post** /gsuites/{gsuite_id}/associations | Manage the associations of a G Suite instance
[**GraphGSuiteTraverseUser**](GraphApi.md#GraphGSuiteTraverseUser) | **Get** /gsuites/{gsuite_id}/users | List the Users bound to a G Suite instance
[**GraphGSuiteTraverseUserGroup**](GraphApi.md#GraphGSuiteTraverseUserGroup) | **Get** /gsuites/{gsuite_id}/usergroups | List the User Groups bound to a G Suite instance
[**GraphLdapServerAssociationsList**](GraphApi.md#GraphLdapServerAssociationsList) | **Get** /ldapservers/{ldapserver_id}/associations | List the associations of a LDAP Server
[**GraphLdapServerAssociationsPost**](GraphApi.md#GraphLdapServerAssociationsPost) | **Post** /ldapservers/{ldapserver_id}/associations | Manage the associations of a LDAP Server
[**GraphLdapServerTraverseUser**](GraphApi.md#GraphLdapServerTraverseUser) | **Get** /ldapservers/{ldapserver_id}/users | List the Users bound to a LDAP Server
[**GraphLdapServerTraverseUserGroup**](GraphApi.md#GraphLdapServerTraverseUserGroup) | **Get** /ldapservers/{ldapserver_id}/usergroups | List the User Groups bound to a LDAP Server
[**GraphOffice365AssociationsList**](GraphApi.md#GraphOffice365AssociationsList) | **Get** /office365s/{office365_id}/associations | List the associations of an Office 365 instance
[**GraphOffice365AssociationsPost**](GraphApi.md#GraphOffice365AssociationsPost) | **Post** /office365s/{office365_id}/associations | Manage the associations of an Office 365 instance
[**GraphOffice365TraverseUser**](GraphApi.md#GraphOffice365TraverseUser) | **Get** /office365s/{office365_id}/users | List the Users bound to an Office 365 instance
[**GraphOffice365TraverseUserGroup**](GraphApi.md#GraphOffice365TraverseUserGroup) | **Get** /office365s/{office365_id}/usergroups | List the User Groups bound to an Office 365 instance
[**GraphPolicyAssociationsList**](GraphApi.md#GraphPolicyAssociationsList) | **Get** /policies/{policy_id}/associations | List the associations of a Policy
[**GraphPolicyAssociationsPost**](GraphApi.md#GraphPolicyAssociationsPost) | **Post** /policies/{policy_id}/associations | Manage the associations of a Policy
[**GraphPolicyGroupAssociationsList**](GraphApi.md#GraphPolicyGroupAssociationsList) | **Get** /policygroups/{group_id}/associations | List the associations of a Policy Group.
[**GraphPolicyGroupAssociationsPost**](GraphApi.md#GraphPolicyGroupAssociationsPost) | **Post** /policygroups/{group_id}/associations | Manage the associations of a Policy Group
[**GraphPolicyGroupMembersList**](GraphApi.md#GraphPolicyGroupMembersList) | **Get** /policygroups/{group_id}/members | List the members of a Policy Group
[**GraphPolicyGroupMembersPost**](GraphApi.md#GraphPolicyGroupMembersPost) | **Post** /policygroups/{group_id}/members | Manage the members of a Policy Group
[**GraphPolicyGroupMembership**](GraphApi.md#GraphPolicyGroupMembership) | **Get** /policygroups/{group_id}/membership | List the Policy Group&#39;s membership
[**GraphPolicyGroupTraverseSystem**](GraphApi.md#GraphPolicyGroupTraverseSystem) | **Get** /policygroups/{group_id}/systems | List the Systems bound to a Policy Group
[**GraphPolicyGroupTraverseSystemGroup**](GraphApi.md#GraphPolicyGroupTraverseSystemGroup) | **Get** /policygroups/{group_id}/systemgroups | List the System Groups bound to Policy Groups
[**GraphPolicyMemberOf**](GraphApi.md#GraphPolicyMemberOf) | **Get** /policies/{policy_id}/memberof | List the parent Groups of a Policy
[**GraphPolicyTraverseSystem**](GraphApi.md#GraphPolicyTraverseSystem) | **Get** /policies/{policy_id}/systems | List the Systems bound to a Policy
[**GraphPolicyTraverseSystemGroup**](GraphApi.md#GraphPolicyTraverseSystemGroup) | **Get** /policies/{policy_id}/systemgroups | List the System Groups bound to a Policy
[**GraphRadiusServerAssociationsList**](GraphApi.md#GraphRadiusServerAssociationsList) | **Get** /radiusservers/{radiusserver_id}/associations | List the associations of a RADIUS  Server
[**GraphRadiusServerAssociationsPost**](GraphApi.md#GraphRadiusServerAssociationsPost) | **Post** /radiusservers/{radiusserver_id}/associations | Manage the associations of a RADIUS Server
[**GraphRadiusServerTraverseUser**](GraphApi.md#GraphRadiusServerTraverseUser) | **Get** /radiusservers/{radiusserver_id}/users | List the Users bound to a RADIUS  Server
[**GraphRadiusServerTraverseUserGroup**](GraphApi.md#GraphRadiusServerTraverseUserGroup) | **Get** /radiusservers/{radiusserver_id}/usergroups | List the User Groups bound to a RADIUS  Server
[**GraphSoftwareappsAssociationsList**](GraphApi.md#GraphSoftwareappsAssociationsList) | **Get** /softwareapps/{software_app_id}/associations | List the associations of a Software Application
[**GraphSoftwareappsAssociationsPost**](GraphApi.md#GraphSoftwareappsAssociationsPost) | **Post** /softwareapps/{software_app_id}/associations | Manage the associations of a software application.
[**GraphSoftwareappsTraverseSystem**](GraphApi.md#GraphSoftwareappsTraverseSystem) | **Get** /softwareapps/{software_app_id}/systems | List the Systems bound to a Software App.
[**GraphSoftwareappsTraverseSystemGroup**](GraphApi.md#GraphSoftwareappsTraverseSystemGroup) | **Get** /softwareapps/{software_app_id}/systemgroups | List the System Groups bound to a Software App.
[**GraphSystemAssociationsList**](GraphApi.md#GraphSystemAssociationsList) | **Get** /systems/{system_id}/associations | List the associations of a System
[**GraphSystemAssociationsPost**](GraphApi.md#GraphSystemAssociationsPost) | **Post** /systems/{system_id}/associations | Manage associations of a System
[**GraphSystemGroupAssociationsList**](GraphApi.md#GraphSystemGroupAssociationsList) | **Get** /systemgroups/{group_id}/associations | List the associations of a System Group
[**GraphSystemGroupAssociationsPost**](GraphApi.md#GraphSystemGroupAssociationsPost) | **Post** /systemgroups/{group_id}/associations | Manage the associations of a System Group
[**GraphSystemGroupMembersList**](GraphApi.md#GraphSystemGroupMembersList) | **Get** /systemgroups/{group_id}/members | List the members of a System Group
[**GraphSystemGroupMembersPost**](GraphApi.md#GraphSystemGroupMembersPost) | **Post** /systemgroups/{group_id}/members | Manage the members of a System Group
[**GraphSystemGroupMembership**](GraphApi.md#GraphSystemGroupMembership) | **Get** /systemgroups/{group_id}/membership | List the System Group&#39;s membership
[**GraphSystemGroupTraverseCommand**](GraphApi.md#GraphSystemGroupTraverseCommand) | **Get** /systemgroups/{group_id}/commands | List the Commands bound to a System Group
[**GraphSystemGroupTraversePolicy**](GraphApi.md#GraphSystemGroupTraversePolicy) | **Get** /systemgroups/{group_id}/policies | List the Policies bound to a System Group
[**GraphSystemGroupTraversePolicyGroup**](GraphApi.md#GraphSystemGroupTraversePolicyGroup) | **Get** /systemgroups/{group_id}/policygroups | List the Policy Groups bound to a System Group
[**GraphSystemGroupTraverseUser**](GraphApi.md#GraphSystemGroupTraverseUser) | **Get** /systemgroups/{group_id}/users | List the Users bound to a System Group
[**GraphSystemGroupTraverseUserGroup**](GraphApi.md#GraphSystemGroupTraverseUserGroup) | **Get** /systemgroups/{group_id}/usergroups | List the User Groups bound to a System Group
[**GraphSystemMemberOf**](GraphApi.md#GraphSystemMemberOf) | **Get** /systems/{system_id}/memberof | List the parent Groups of a System
[**GraphSystemTraverseCommand**](GraphApi.md#GraphSystemTraverseCommand) | **Get** /systems/{system_id}/commands | List the Commands bound to a System
[**GraphSystemTraversePolicy**](GraphApi.md#GraphSystemTraversePolicy) | **Get** /systems/{system_id}/policies | List the Policies bound to a System
[**GraphSystemTraversePolicyGroup**](GraphApi.md#GraphSystemTraversePolicyGroup) | **Get** /systems/{system_id}/policygroups | List the Policy Groups bound to a System
[**GraphSystemTraverseUser**](GraphApi.md#GraphSystemTraverseUser) | **Get** /systems/{system_id}/users | List the Users bound to a System
[**GraphSystemTraverseUserGroup**](GraphApi.md#GraphSystemTraverseUserGroup) | **Get** /systems/{system_id}/usergroups | List the User Groups bound to a System
[**GraphUserAssociationsList**](GraphApi.md#GraphUserAssociationsList) | **Get** /users/{user_id}/associations | List the associations of a User
[**GraphUserAssociationsPost**](GraphApi.md#GraphUserAssociationsPost) | **Post** /users/{user_id}/associations | Manage the associations of a User
[**GraphUserGroupAssociationsList**](GraphApi.md#GraphUserGroupAssociationsList) | **Get** /usergroups/{group_id}/associations | List the associations of a User Group.
[**GraphUserGroupAssociationsPost**](GraphApi.md#GraphUserGroupAssociationsPost) | **Post** /usergroups/{group_id}/associations | Manage the associations of a User Group
[**GraphUserGroupMembersList**](GraphApi.md#GraphUserGroupMembersList) | **Get** /usergroups/{group_id}/members | List the members of a User Group
[**GraphUserGroupMembersPost**](GraphApi.md#GraphUserGroupMembersPost) | **Post** /usergroups/{group_id}/members | Manage the members of a User Group
[**GraphUserGroupMembership**](GraphApi.md#GraphUserGroupMembership) | **Get** /usergroups/{group_id}/membership | List the User Group&#39;s membership
[**GraphUserGroupTraverseActiveDirectory**](GraphApi.md#GraphUserGroupTraverseActiveDirectory) | **Get** /usergroups/{group_id}/activedirectories | List the Active Directories bound to a User Group
[**GraphUserGroupTraverseApplication**](GraphApi.md#GraphUserGroupTraverseApplication) | **Get** /usergroups/{group_id}/applications | List the Applications bound to a User Group
[**GraphUserGroupTraverseDirectory**](GraphApi.md#GraphUserGroupTraverseDirectory) | **Get** /usergroups/{group_id}/directories | List the Directories bound to a User Group
[**GraphUserGroupTraverseGSuite**](GraphApi.md#GraphUserGroupTraverseGSuite) | **Get** /usergroups/{group_id}/gsuites | List the G Suite instances bound to a User Group
[**GraphUserGroupTraverseLdapServer**](GraphApi.md#GraphUserGroupTraverseLdapServer) | **Get** /usergroups/{group_id}/ldapservers | List the LDAP Servers bound to a User Group
[**GraphUserGroupTraverseOffice365**](GraphApi.md#GraphUserGroupTraverseOffice365) | **Get** /usergroups/{group_id}/office365s | List the Office 365 instances bound to a User Group
[**GraphUserGroupTraverseRadiusServer**](GraphApi.md#GraphUserGroupTraverseRadiusServer) | **Get** /usergroups/{group_id}/radiusservers | List the RADIUS Servers bound to a User Group
[**GraphUserGroupTraverseSystem**](GraphApi.md#GraphUserGroupTraverseSystem) | **Get** /usergroups/{group_id}/systems | List the Systems bound to a User Group
[**GraphUserGroupTraverseSystemGroup**](GraphApi.md#GraphUserGroupTraverseSystemGroup) | **Get** /usergroups/{group_id}/systemgroups | List the System Groups bound to User Groups
[**GraphUserMemberOf**](GraphApi.md#GraphUserMemberOf) | **Get** /users/{user_id}/memberof | List the parent Groups of a User
[**GraphUserTraverseActiveDirectory**](GraphApi.md#GraphUserTraverseActiveDirectory) | **Get** /users/{user_id}/activedirectories | List the Active Directory instances bound to a User
[**GraphUserTraverseApplication**](GraphApi.md#GraphUserTraverseApplication) | **Get** /users/{user_id}/applications | List the Applications bound to a User
[**GraphUserTraverseDirectory**](GraphApi.md#GraphUserTraverseDirectory) | **Get** /users/{user_id}/directories | List the Directories bound to a User
[**GraphUserTraverseGSuite**](GraphApi.md#GraphUserTraverseGSuite) | **Get** /users/{user_id}/gsuites | List the G Suite instances bound to a User
[**GraphUserTraverseLdapServer**](GraphApi.md#GraphUserTraverseLdapServer) | **Get** /users/{user_id}/ldapservers | List the LDAP servers bound to a User
[**GraphUserTraverseOffice365**](GraphApi.md#GraphUserTraverseOffice365) | **Get** /users/{user_id}/office365s | List the Office 365 instances bound to a User
[**GraphUserTraverseRadiusServer**](GraphApi.md#GraphUserTraverseRadiusServer) | **Get** /users/{user_id}/radiusservers | List the RADIUS Servers bound to a User
[**GraphUserTraverseSystem**](GraphApi.md#GraphUserTraverseSystem) | **Get** /users/{user_id}/systems | List the Systems bound to a User
[**GraphUserTraverseSystemGroup**](GraphApi.md#GraphUserTraverseSystemGroup) | **Get** /users/{user_id}/systemgroups | List the System Groups bound to a User
[**PolicystatusesSystemsList**](GraphApi.md#PolicystatusesSystemsList) | **Get** /systems/{system_id}/policystatuses | List the policy statuses for a system



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | 
    targets := []string{"Targets_example"} // []string | Targets which a \"active_directory\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphActiveDirectoryAssociationsList(context.Background(), activedirectoryId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphActiveDirectoryAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphActiveDirectoryAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphActiveDirectoryAssociationsList`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationActiveDirectory("Id_example", "Op_example", "Type_example") // GraphOperationActiveDirectory |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphActiveDirectoryAssociationsPost(context.Background(), activedirectoryId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphActiveDirectoryAssociationsPost``: %v\n", err)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | ObjectID of the Active Directory instance.
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphActiveDirectoryTraverseUser(context.Background(), activedirectoryId).Filter(filter).Limit(limit).XOrgId(xOrgId).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphActiveDirectoryTraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphActiveDirectoryTraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphActiveDirectoryTraverseUser`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    activedirectoryId := "activedirectoryId_example" // string | ObjectID of the Active Directory instance.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphActiveDirectoryTraverseUserGroup(context.Background(), activedirectoryId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphActiveDirectoryTraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphActiveDirectoryTraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphActiveDirectoryTraverseUserGroup`: %v\n", resp)
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


## GraphApplicationAssociationsList

> []GraphConnection GraphApplicationAssociationsList(ctx, applicationId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of an Application



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
    applicationId := "applicationId_example" // string | ObjectID of the Application.
    targets := []string{"Targets_example"} // []string | Targets which a \"application\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphApplicationAssociationsList(context.Background(), applicationId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphApplicationAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphApplicationAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphApplicationAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | ObjectID of the Application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphApplicationAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;application\&quot; can be associated to. | 
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


## GraphApplicationAssociationsPost

> GraphApplicationAssociationsPost(ctx, applicationId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of an Application



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
    applicationId := "applicationId_example" // string | ObjectID of the Application.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationApplication("Id_example", "Op_example", "Type_example") // GraphOperationApplication |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphApplicationAssociationsPost(context.Background(), applicationId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphApplicationAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | ObjectID of the Application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphApplicationAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationApplication**](GraphOperationApplication.md) |  | 

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


## GraphApplicationTraverseUser

> []GraphObjectWithPaths GraphApplicationTraverseUser(ctx, applicationId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Users bound to an Application



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
    applicationId := "applicationId_example" // string | ObjectID of the Application.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphApplicationTraverseUser(context.Background(), applicationId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphApplicationTraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphApplicationTraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphApplicationTraverseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | ObjectID of the Application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphApplicationTraverseUserRequest struct via the builder pattern


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


## GraphApplicationTraverseUserGroup

> []GraphObjectWithPaths GraphApplicationTraverseUserGroup(ctx, applicationId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the User Groups bound to an Application



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
    applicationId := "applicationId_example" // string | ObjectID of the Application.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphApplicationTraverseUserGroup(context.Background(), applicationId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphApplicationTraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphApplicationTraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphApplicationTraverseUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | ObjectID of the Application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphApplicationTraverseUserGroupRequest struct via the builder pattern


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


## GraphCommandAssociationsList

> []GraphConnection GraphCommandAssociationsList(ctx, commandId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a Command



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
    commandId := "commandId_example" // string | ObjectID of the Command.
    targets := []string{"Targets_example"} // []string | Targets which a \"command\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphCommandAssociationsList(context.Background(), commandId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphCommandAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphCommandAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphCommandAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | ObjectID of the Command. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphCommandAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;command\&quot; can be associated to. | 
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


## GraphCommandAssociationsPost

> GraphCommandAssociationsPost(ctx, commandId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a Command



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
    commandId := "commandId_example" // string | ObjectID of the Command.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationCommand("Id_example", "Op_example", "Type_example") // GraphOperationCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphCommandAssociationsPost(context.Background(), commandId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphCommandAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | ObjectID of the Command. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphCommandAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationCommand**](GraphOperationCommand.md) |  | 

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


## GraphCommandTraverseSystem

> []GraphObjectWithPaths GraphCommandTraverseSystem(ctx, commandId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Systems bound to a Command



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
    commandId := "commandId_example" // string | ObjectID of the Command.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphCommandTraverseSystem(context.Background(), commandId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphCommandTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphCommandTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphCommandTraverseSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | ObjectID of the Command. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphCommandTraverseSystemRequest struct via the builder pattern


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


## GraphCommandTraverseSystemGroup

> []GraphObjectWithPaths GraphCommandTraverseSystemGroup(ctx, commandId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the System Groups bound to a Command



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
    commandId := "commandId_example" // string | ObjectID of the Command.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphCommandTraverseSystemGroup(context.Background(), commandId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphCommandTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphCommandTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphCommandTraverseSystemGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | ObjectID of the Command. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphCommandTraverseSystemGroupRequest struct via the builder pattern


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


## GraphGSuiteAssociationsList

> []GraphConnection GraphGSuiteAssociationsList(ctx, gsuiteId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a G Suite instance



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
    gsuiteId := "gsuiteId_example" // string | ObjectID of the G Suite instance.
    targets := []string{"Targets_example"} // []string | Targets which a \"g_suite\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphGSuiteAssociationsList(context.Background(), gsuiteId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphGSuiteAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphGSuiteAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphGSuiteAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gsuiteId** | **string** | ObjectID of the G Suite instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphGSuiteAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;g_suite\&quot; can be associated to. | 
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


## GraphGSuiteAssociationsPost

> GraphGSuiteAssociationsPost(ctx, gsuiteId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a G Suite instance



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
    gsuiteId := "gsuiteId_example" // string | ObjectID of the G Suite instance.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationGSuite("Id_example", "Op_example", "Type_example") // GraphOperationGSuite |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphGSuiteAssociationsPost(context.Background(), gsuiteId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphGSuiteAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gsuiteId** | **string** | ObjectID of the G Suite instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphGSuiteAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationGSuite**](GraphOperationGSuite.md) |  | 

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


## GraphGSuiteTraverseUser

> []GraphObjectWithPaths GraphGSuiteTraverseUser(ctx, gsuiteId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Users bound to a G Suite instance



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
    gsuiteId := "gsuiteId_example" // string | ObjectID of the G Suite instance.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphGSuiteTraverseUser(context.Background(), gsuiteId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphGSuiteTraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphGSuiteTraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphGSuiteTraverseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gsuiteId** | **string** | ObjectID of the G Suite instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphGSuiteTraverseUserRequest struct via the builder pattern


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


## GraphGSuiteTraverseUserGroup

> []GraphObjectWithPaths GraphGSuiteTraverseUserGroup(ctx, gsuiteId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the User Groups bound to a G Suite instance



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
    gsuiteId := "gsuiteId_example" // string | ObjectID of the G Suite instance.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphGSuiteTraverseUserGroup(context.Background(), gsuiteId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphGSuiteTraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphGSuiteTraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphGSuiteTraverseUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gsuiteId** | **string** | ObjectID of the G Suite instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphGSuiteTraverseUserGroupRequest struct via the builder pattern


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


## GraphLdapServerAssociationsList

> []GraphConnection GraphLdapServerAssociationsList(ctx, ldapserverId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a LDAP Server



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
    ldapserverId := "ldapserverId_example" // string | ObjectID of the LDAP Server.
    targets := []string{"Targets_example"} // []string | Targets which a \"ldap_server\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphLdapServerAssociationsList(context.Background(), ldapserverId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphLdapServerAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphLdapServerAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphLdapServerAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapserverId** | **string** | ObjectID of the LDAP Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphLdapServerAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;ldap_server\&quot; can be associated to. | 
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


## GraphLdapServerAssociationsPost

> GraphLdapServerAssociationsPost(ctx, ldapserverId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a LDAP Server



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
    ldapserverId := "ldapserverId_example" // string | ObjectID of the LDAP Server.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationLdapServer("Id_example", "Op_example", "Type_example") // GraphOperationLdapServer |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphLdapServerAssociationsPost(context.Background(), ldapserverId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphLdapServerAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapserverId** | **string** | ObjectID of the LDAP Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphLdapServerAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationLdapServer**](GraphOperationLdapServer.md) |  | 

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


## GraphLdapServerTraverseUser

> []GraphObjectWithPaths GraphLdapServerTraverseUser(ctx, ldapserverId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Users bound to a LDAP Server



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
    ldapserverId := "ldapserverId_example" // string | ObjectID of the LDAP Server.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphLdapServerTraverseUser(context.Background(), ldapserverId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphLdapServerTraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphLdapServerTraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphLdapServerTraverseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapserverId** | **string** | ObjectID of the LDAP Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphLdapServerTraverseUserRequest struct via the builder pattern


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


## GraphLdapServerTraverseUserGroup

> []GraphObjectWithPaths GraphLdapServerTraverseUserGroup(ctx, ldapserverId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the User Groups bound to a LDAP Server



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
    ldapserverId := "ldapserverId_example" // string | ObjectID of the LDAP Server.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphLdapServerTraverseUserGroup(context.Background(), ldapserverId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphLdapServerTraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphLdapServerTraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphLdapServerTraverseUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapserverId** | **string** | ObjectID of the LDAP Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphLdapServerTraverseUserGroupRequest struct via the builder pattern


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


## GraphOffice365AssociationsList

> []GraphConnection GraphOffice365AssociationsList(ctx, office365Id).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of an Office 365 instance



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
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 instance.
    targets := []string{"Targets_example"} // []string | Targets which a \"office_365\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphOffice365AssociationsList(context.Background(), office365Id).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphOffice365AssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphOffice365AssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphOffice365AssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** | ObjectID of the Office 365 instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphOffice365AssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;office_365\&quot; can be associated to. | 
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


## GraphOffice365AssociationsPost

> GraphOffice365AssociationsPost(ctx, office365Id).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of an Office 365 instance



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
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 instance.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationOffice365("Id_example", "Op_example", "Type_example") // GraphOperationOffice365 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphOffice365AssociationsPost(context.Background(), office365Id).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphOffice365AssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** | ObjectID of the Office 365 instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphOffice365AssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationOffice365**](GraphOperationOffice365.md) |  | 

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


## GraphOffice365TraverseUser

> []GraphObjectWithPaths GraphOffice365TraverseUser(ctx, office365Id).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Users bound to an Office 365 instance



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
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 suite.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphOffice365TraverseUser(context.Background(), office365Id).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphOffice365TraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphOffice365TraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphOffice365TraverseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** | ObjectID of the Office 365 suite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphOffice365TraverseUserRequest struct via the builder pattern


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


## GraphOffice365TraverseUserGroup

> []GraphObjectWithPaths GraphOffice365TraverseUserGroup(ctx, office365Id).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the User Groups bound to an Office 365 instance



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
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 suite.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphOffice365TraverseUserGroup(context.Background(), office365Id).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphOffice365TraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphOffice365TraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphOffice365TraverseUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** | ObjectID of the Office 365 suite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphOffice365TraverseUserGroupRequest struct via the builder pattern


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


## GraphPolicyAssociationsList

> []GraphConnection GraphPolicyAssociationsList(ctx, policyId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a Policy



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
    policyId := "policyId_example" // string | ObjectID of the Policy.
    targets := []string{"Targets_example"} // []string | Targets which a \"policy\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphPolicyAssociationsList(context.Background(), policyId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphPolicyAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | ObjectID of the Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;policy\&quot; can be associated to. | 
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


## GraphPolicyAssociationsPost

> GraphPolicyAssociationsPost(ctx, policyId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a Policy



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
    policyId := "policyId_example" // string | ObjectID of the Policy.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationPolicy("Id_example", "Op_example", "Type_example") // GraphOperationPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphPolicyAssociationsPost(context.Background(), policyId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | ObjectID of the Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationPolicy**](GraphOperationPolicy.md) |  | 

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
    resp, r, err := apiClient.GraphApi.GraphPolicyGroupAssociationsList(context.Background(), groupId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyGroupAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphPolicyGroupAssociationsList`: %v\n", resp)
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
    r, err := apiClient.GraphApi.GraphPolicyGroupAssociationsPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyGroupAssociationsPost``: %v\n", err)
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
    resp, r, err := apiClient.GraphApi.GraphPolicyGroupMembersList(context.Background(), groupId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyGroupMembersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupMembersList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphPolicyGroupMembersList`: %v\n", resp)
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
    r, err := apiClient.GraphApi.GraphPolicyGroupMembersPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyGroupMembersPost``: %v\n", err)
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
    resp, r, err := apiClient.GraphApi.GraphPolicyGroupMembership(context.Background(), groupId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupMembership`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphPolicyGroupMembership`: %v\n", resp)
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
    resp, r, err := apiClient.GraphApi.GraphPolicyGroupTraverseSystem(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyGroupTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphPolicyGroupTraverseSystem`: %v\n", resp)
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
    resp, r, err := apiClient.GraphApi.GraphPolicyGroupTraverseSystemGroup(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyGroupTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyGroupTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphPolicyGroupTraverseSystemGroup`: %v\n", resp)
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


## GraphPolicyMemberOf

> []GraphObjectWithPaths GraphPolicyMemberOf(ctx, policyId).Filter(filter).Limit(limit).Skip(skip).Date(date).Authorization(authorization).Sort(sort).XOrgId(xOrgId).Execute()

List the parent Groups of a Policy



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
    policyId := "policyId_example" // string | ObjectID of the Policy.
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    date := "date_example" // string | Current date header for the System Context API (optional)
    authorization := "authorization_example" // string | Authorization header for the System Context API (optional)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphPolicyMemberOf(context.Background(), policyId).Filter(filter).Limit(limit).Skip(skip).Date(date).Authorization(authorization).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyMemberOf`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphPolicyMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | ObjectID of the Policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **date** | **string** | Current date header for the System Context API | 
 **authorization** | **string** | Authorization header for the System Context API | 
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


## GraphPolicyTraverseSystem

> []GraphObjectWithPaths GraphPolicyTraverseSystem(ctx, policyId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Systems bound to a Policy



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
    policyId := "policyId_example" // string | ObjectID of the Command.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphPolicyTraverseSystem(context.Background(), policyId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphPolicyTraverseSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | ObjectID of the Command. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyTraverseSystemRequest struct via the builder pattern


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


## GraphPolicyTraverseSystemGroup

> []GraphObjectWithPaths GraphPolicyTraverseSystemGroup(ctx, policyId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the System Groups bound to a Policy



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
    policyId := "policyId_example" // string | ObjectID of the Command.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphPolicyTraverseSystemGroup(context.Background(), policyId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphPolicyTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphPolicyTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphPolicyTraverseSystemGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | ObjectID of the Command. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphPolicyTraverseSystemGroupRequest struct via the builder pattern


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


## GraphRadiusServerAssociationsList

> []GraphConnection GraphRadiusServerAssociationsList(ctx, radiusserverId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a RADIUS  Server



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
    radiusserverId := "radiusserverId_example" // string | ObjectID of the Radius Server.
    targets := []string{"Targets_example"} // []string | Targets which a \"radius_server\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphRadiusServerAssociationsList(context.Background(), radiusserverId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphRadiusServerAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphRadiusServerAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphRadiusServerAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**radiusserverId** | **string** | ObjectID of the Radius Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphRadiusServerAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;radius_server\&quot; can be associated to. | 
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


## GraphRadiusServerAssociationsPost

> GraphRadiusServerAssociationsPost(ctx, radiusserverId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a RADIUS Server



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
    radiusserverId := "radiusserverId_example" // string | ObjectID of the Radius Server.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationRadiusServer("Id_example", "Op_example", "Type_example") // GraphOperationRadiusServer |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphRadiusServerAssociationsPost(context.Background(), radiusserverId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphRadiusServerAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**radiusserverId** | **string** | ObjectID of the Radius Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphRadiusServerAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationRadiusServer**](GraphOperationRadiusServer.md) |  | 

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


## GraphRadiusServerTraverseUser

> []GraphObjectWithPaths GraphRadiusServerTraverseUser(ctx, radiusserverId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Users bound to a RADIUS  Server



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
    radiusserverId := "radiusserverId_example" // string | ObjectID of the Radius Server.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphRadiusServerTraverseUser(context.Background(), radiusserverId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphRadiusServerTraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphRadiusServerTraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphRadiusServerTraverseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**radiusserverId** | **string** | ObjectID of the Radius Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphRadiusServerTraverseUserRequest struct via the builder pattern


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


## GraphRadiusServerTraverseUserGroup

> []GraphObjectWithPaths GraphRadiusServerTraverseUserGroup(ctx, radiusserverId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the User Groups bound to a RADIUS  Server



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
    radiusserverId := "radiusserverId_example" // string | ObjectID of the Radius Server.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphRadiusServerTraverseUserGroup(context.Background(), radiusserverId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphRadiusServerTraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphRadiusServerTraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphRadiusServerTraverseUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**radiusserverId** | **string** | ObjectID of the Radius Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphRadiusServerTraverseUserGroupRequest struct via the builder pattern


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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    softwareAppId := "softwareAppId_example" // string | ObjectID of the Software App.
    targets := []string{"Targets_example"} // []string | Targets which a \"software_app\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSoftwareappsAssociationsList(context.Background(), softwareAppId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSoftwareappsAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSoftwareappsAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSoftwareappsAssociationsList`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    softwareAppId := "softwareAppId_example" // string | ObjectID of the Software App.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationSoftwareApp("Id_example", "Op_example", "Type_example") // GraphOperationSoftwareApp |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphSoftwareappsAssociationsPost(context.Background(), softwareAppId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSoftwareappsAssociationsPost``: %v\n", err)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    softwareAppId := "softwareAppId_example" // string | ObjectID of the Software App.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSoftwareappsTraverseSystem(context.Background(), softwareAppId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSoftwareappsTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSoftwareappsTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSoftwareappsTraverseSystem`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    softwareAppId := "softwareAppId_example" // string | ObjectID of the Software App.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSoftwareappsTraverseSystemGroup(context.Background(), softwareAppId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSoftwareappsTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSoftwareappsTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSoftwareappsTraverseSystemGroup`: %v\n", resp)
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


## GraphSystemAssociationsList

> []GraphConnection GraphSystemAssociationsList(ctx, systemId).Targets(targets).Limit(limit).Skip(skip).Date(date).Authorization(authorization).XOrgId(xOrgId).Execute()

List the associations of a System



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
    systemId := "systemId_example" // string | ObjectID of the System.
    targets := []string{"Targets_example"} // []string | Targets which a \"system\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    date := "date_example" // string | Current date header for the System Context API (optional)
    authorization := "authorization_example" // string | Authorization header for the System Context API (optional)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemAssociationsList(context.Background(), systemId).Targets(targets).Limit(limit).Skip(skip).Date(date).Authorization(authorization).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ObjectID of the System. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;system\&quot; can be associated to. | 
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **date** | **string** | Current date header for the System Context API | 
 **authorization** | **string** | Authorization header for the System Context API | 
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


## GraphSystemAssociationsPost

> GraphSystemAssociationsPost(ctx, systemId).Date(date).Authorization(authorization).XOrgId(xOrgId).Body(body).Execute()

Manage associations of a System



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
    systemId := "systemId_example" // string | ObjectID of the System.
    date := "date_example" // string | Current date header for the System Context API (optional)
    authorization := "authorization_example" // string | Authorization header for the System Context API (optional)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationSystem("Id_example", "Op_example", "Type_example") // GraphOperationSystem |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphSystemAssociationsPost(context.Background(), systemId).Date(date).Authorization(authorization).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ObjectID of the System. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | Current date header for the System Context API | 
 **authorization** | **string** | Authorization header for the System Context API | 
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationSystem**](GraphOperationSystem.md) |  | 

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


## GraphSystemGroupAssociationsList

> []GraphConnection GraphSystemGroupAssociationsList(ctx, groupId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a System Group



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    targets := []string{"Targets_example"} // []string | Targets which a \"system_group\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemGroupAssociationsList(context.Background(), groupId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemGroupAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemGroupAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;system_group\&quot; can be associated to. | 
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


## GraphSystemGroupAssociationsPost

> GraphSystemGroupAssociationsPost(ctx, groupId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a System Group



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationSystemGroup("Id_example", "Op_example", "Type_example") // GraphOperationSystemGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphSystemGroupAssociationsPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationSystemGroup**](GraphOperationSystemGroup.md) |  | 

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


## GraphSystemGroupMembersList

> []GraphConnection GraphSystemGroupMembersList(ctx, groupId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the members of a System Group



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemGroupMembersList(context.Background(), groupId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupMembersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemGroupMembersList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemGroupMembersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupMembersListRequest struct via the builder pattern


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


## GraphSystemGroupMembersPost

> GraphSystemGroupMembersPost(ctx, groupId).Date(date).Authorization(authorization).XOrgId(xOrgId).Body(body).Execute()

Manage the members of a System Group



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    date := "date_example" // string | Current date header for the System Context API (optional)
    authorization := "authorization_example" // string | Authorization header for the System Context API (optional)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationSystemGroupMember("Id_example", "Op_example", "Type_example") // GraphOperationSystemGroupMember |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphSystemGroupMembersPost(context.Background(), groupId).Date(date).Authorization(authorization).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupMembersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupMembersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | Current date header for the System Context API | 
 **authorization** | **string** | Authorization header for the System Context API | 
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationSystemGroupMember**](GraphOperationSystemGroupMember.md) |  | 

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


## GraphSystemGroupMembership

> []GraphObjectWithPaths GraphSystemGroupMembership(ctx, groupId).Limit(limit).Skip(skip).Sort(sort).Filter(filter).XOrgId(xOrgId).Execute()

List the System Group's membership



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemGroupMembership(context.Background(), groupId).Limit(limit).Skip(skip).Sort(sort).Filter(filter).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemGroupMembership`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemGroupMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
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


## GraphSystemGroupTraverseCommand

> []GraphObjectWithPaths GraphSystemGroupTraverseCommand(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Commands bound to a System Group



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemGroupTraverseCommand(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupTraverseCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemGroupTraverseCommand`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemGroupTraverseCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupTraverseCommandRequest struct via the builder pattern


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


## GraphSystemGroupTraversePolicy

> []GraphObjectWithPaths GraphSystemGroupTraversePolicy(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Policies bound to a System Group



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemGroupTraversePolicy(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupTraversePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemGroupTraversePolicy`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemGroupTraversePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupTraversePolicyRequest struct via the builder pattern


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


## GraphSystemGroupTraversePolicyGroup

> []GraphObjectWithPaths GraphSystemGroupTraversePolicyGroup(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Policy Groups bound to a System Group



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemGroupTraversePolicyGroup(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupTraversePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemGroupTraversePolicyGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemGroupTraversePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupTraversePolicyGroupRequest struct via the builder pattern


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


## GraphSystemGroupTraverseUser

> []GraphObjectWithPaths GraphSystemGroupTraverseUser(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Users bound to a System Group



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemGroupTraverseUser(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupTraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemGroupTraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemGroupTraverseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupTraverseUserRequest struct via the builder pattern


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


## GraphSystemGroupTraverseUserGroup

> []GraphObjectWithPaths GraphSystemGroupTraverseUserGroup(ctx, groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the User Groups bound to a System Group



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
    groupId := "groupId_example" // string | ObjectID of the System Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemGroupTraverseUserGroup(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemGroupTraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemGroupTraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemGroupTraverseUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ObjectID of the System Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemGroupTraverseUserGroupRequest struct via the builder pattern


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


## GraphSystemMemberOf

> []GraphObjectWithPaths GraphSystemMemberOf(ctx, systemId).Filter(filter).Limit(limit).Skip(skip).Date(date).Authorization(authorization).Sort(sort).XOrgId(xOrgId).Execute()

List the parent Groups of a System



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
    systemId := "systemId_example" // string | ObjectID of the System.
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    date := "date_example" // string | Current date header for the System Context API (optional)
    authorization := "authorization_example" // string | Authorization header for the System Context API (optional)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemMemberOf(context.Background(), systemId).Filter(filter).Limit(limit).Skip(skip).Date(date).Authorization(authorization).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemMemberOf`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ObjectID of the System. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **date** | **string** | Current date header for the System Context API | 
 **authorization** | **string** | Authorization header for the System Context API | 
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


## GraphSystemTraverseCommand

> []GraphObjectWithPaths GraphSystemTraverseCommand(ctx, systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Commands bound to a System



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
    systemId := "systemId_example" // string | ObjectID of the System.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemTraverseCommand(context.Background(), systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemTraverseCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemTraverseCommand`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemTraverseCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ObjectID of the System. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemTraverseCommandRequest struct via the builder pattern


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


## GraphSystemTraversePolicy

> []GraphObjectWithPaths GraphSystemTraversePolicy(ctx, systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Policies bound to a System



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
    systemId := "systemId_example" // string | ObjectID of the System.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemTraversePolicy(context.Background(), systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemTraversePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemTraversePolicy`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemTraversePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ObjectID of the System. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemTraversePolicyRequest struct via the builder pattern


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


## GraphSystemTraversePolicyGroup

> []GraphObjectWithPaths GraphSystemTraversePolicyGroup(ctx, systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Date(date).Authorization(authorization).Filter(filter).Execute()

List the Policy Groups bound to a System



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
    systemId := "systemId_example" // string | ObjectID of the System.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    date := "date_example" // string | Current date header for the System Context API (optional)
    authorization := "authorization_example" // string | Authorization header for the System Context API (optional)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemTraversePolicyGroup(context.Background(), systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Date(date).Authorization(authorization).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemTraversePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemTraversePolicyGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemTraversePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ObjectID of the System. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemTraversePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **date** | **string** | Current date header for the System Context API | 
 **authorization** | **string** | Authorization header for the System Context API | 
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


## GraphSystemTraverseUser

> []GraphObjectWithPaths GraphSystemTraverseUser(ctx, systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Date(date).Authorization(authorization).Filter(filter).Execute()

List the Users bound to a System



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
    systemId := "systemId_example" // string | ObjectID of the System.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    date := "date_example" // string | Current date header for the System Context API (optional)
    authorization := "authorization_example" // string | Authorization header for the System Context API (optional)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemTraverseUser(context.Background(), systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Date(date).Authorization(authorization).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemTraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemTraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemTraverseUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ObjectID of the System. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemTraverseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **date** | **string** | Current date header for the System Context API | 
 **authorization** | **string** | Authorization header for the System Context API | 
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


## GraphSystemTraverseUserGroup

> []GraphObjectWithPaths GraphSystemTraverseUserGroup(ctx, systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Date(date).Authorization(authorization).Filter(filter).Execute()

List the User Groups bound to a System



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
    systemId := "systemId_example" // string | ObjectID of the System.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    date := "date_example" // string | Current date header for the System Context API (optional)
    authorization := "authorization_example" // string | Authorization header for the System Context API (optional)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphSystemTraverseUserGroup(context.Background(), systemId).Limit(limit).XOrgId(xOrgId).Skip(skip).Date(date).Authorization(authorization).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphSystemTraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphSystemTraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphSystemTraverseUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ObjectID of the System. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphSystemTraverseUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **date** | **string** | Current date header for the System Context API | 
 **authorization** | **string** | Authorization header for the System Context API | 
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


## GraphUserAssociationsList

> []GraphConnection GraphUserAssociationsList(ctx, userId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List the associations of a User



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
    userId := "userId_example" // string | ObjectID of the User.
    targets := []string{"Targets_example"} // []string | Targets which a \"user\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserAssociationsList(context.Background(), userId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserAssociationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targets** | **[]string** | Targets which a \&quot;user\&quot; can be associated to. | 
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


## GraphUserAssociationsPost

> GraphUserAssociationsPost(ctx, userId).XOrgId(xOrgId).Body(body).Execute()

Manage the associations of a User



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
    userId := "userId_example" // string | ObjectID of the User.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationUser("Id_example", "Op_example", "Type_example") // GraphOperationUser |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphUserAssociationsPost(context.Background(), userId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserAssociationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**GraphOperationUser**](GraphOperationUser.md) |  | 

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    targets := []string{"Targets_example"} // []string | Targets which a \"user_group\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupAssociationsList(context.Background(), groupId).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupAssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupAssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupAssociationsList`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationUserGroup("Id_example", "Op_example", "Type_example") // GraphOperationUserGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphUserGroupAssociationsPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupAssociationsPost``: %v\n", err)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupMembersList(context.Background(), groupId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupMembersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupMembersList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupMembersList`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationUserGroupMember("Id_example", "Op_example", "Type_example") // GraphOperationUserGroupMember |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GraphApi.GraphUserGroupMembersPost(context.Background(), groupId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupMembersPost``: %v\n", err)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
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
    resp, r, err := apiClient.GraphApi.GraphUserGroupMembership(context.Background(), groupId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupMembership`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupMembership`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupTraverseActiveDirectory(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupTraverseActiveDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseActiveDirectory`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupTraverseActiveDirectory`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupTraverseApplication(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupTraverseApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseApplication`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupTraverseApplication`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupTraverseDirectory(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupTraverseDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseDirectory`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupTraverseDirectory`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupTraverseGSuite(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupTraverseGSuite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseGSuite`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupTraverseGSuite`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupTraverseLdapServer(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupTraverseLdapServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseLdapServer`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupTraverseLdapServer`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupTraverseOffice365(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupTraverseOffice365``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseOffice365`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupTraverseOffice365`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupTraverseRadiusServer(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupTraverseRadiusServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseRadiusServer`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupTraverseRadiusServer`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupTraverseSystem(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupTraverseSystem`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    groupId := "groupId_example" // string | ObjectID of the User Group.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserGroupTraverseSystemGroup(context.Background(), groupId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserGroupTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserGroupTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserGroupTraverseSystemGroup`: %v\n", resp)
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


## GraphUserMemberOf

> []GraphObjectWithPaths GraphUserMemberOf(ctx, userId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()

List the parent Groups of a User



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
    userId := "userId_example" // string | ObjectID of the User.
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserMemberOf(context.Background(), userId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserMemberOf`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserMemberOfRequest struct via the builder pattern


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


## GraphUserTraverseActiveDirectory

> []GraphObjectWithPaths GraphUserTraverseActiveDirectory(ctx, userId).Filter(filter).Limit(limit).XOrgId(xOrgId).Skip(skip).Execute()

List the Active Directory instances bound to a User



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
    userId := "userId_example" // string | ObjectID of the User.
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserTraverseActiveDirectory(context.Background(), userId).Filter(filter).Limit(limit).XOrgId(xOrgId).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserTraverseActiveDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserTraverseActiveDirectory`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserTraverseActiveDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserTraverseActiveDirectoryRequest struct via the builder pattern


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


## GraphUserTraverseApplication

> []GraphObjectWithPaths GraphUserTraverseApplication(ctx, userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Applications bound to a User



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
    userId := "userId_example" // string | ObjectID of the User.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserTraverseApplication(context.Background(), userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserTraverseApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserTraverseApplication`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserTraverseApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserTraverseApplicationRequest struct via the builder pattern


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


## GraphUserTraverseDirectory

> []GraphObjectWithPaths GraphUserTraverseDirectory(ctx, userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Directories bound to a User



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
    userId := "userId_example" // string | ObjectID of the User.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserTraverseDirectory(context.Background(), userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserTraverseDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserTraverseDirectory`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserTraverseDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserTraverseDirectoryRequest struct via the builder pattern


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


## GraphUserTraverseGSuite

> []GraphObjectWithPaths GraphUserTraverseGSuite(ctx, userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the G Suite instances bound to a User



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
    userId := "userId_example" // string | ObjectID of the User.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserTraverseGSuite(context.Background(), userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserTraverseGSuite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserTraverseGSuite`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserTraverseGSuite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserTraverseGSuiteRequest struct via the builder pattern


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


## GraphUserTraverseLdapServer

> []GraphObjectWithPaths GraphUserTraverseLdapServer(ctx, userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the LDAP servers bound to a User



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
    userId := "userId_example" // string | ObjectID of the User.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserTraverseLdapServer(context.Background(), userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserTraverseLdapServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserTraverseLdapServer`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserTraverseLdapServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserTraverseLdapServerRequest struct via the builder pattern


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


## GraphUserTraverseOffice365

> []GraphObjectWithPaths GraphUserTraverseOffice365(ctx, userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Office 365 instances bound to a User



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
    userId := "userId_example" // string | ObjectID of the User.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserTraverseOffice365(context.Background(), userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserTraverseOffice365``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserTraverseOffice365`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserTraverseOffice365`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserTraverseOffice365Request struct via the builder pattern


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


## GraphUserTraverseRadiusServer

> []GraphObjectWithPaths GraphUserTraverseRadiusServer(ctx, userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the RADIUS Servers bound to a User



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
    userId := "userId_example" // string | ObjectID of the User.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserTraverseRadiusServer(context.Background(), userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserTraverseRadiusServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserTraverseRadiusServer`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserTraverseRadiusServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserTraverseRadiusServerRequest struct via the builder pattern


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


## GraphUserTraverseSystem

> []GraphObjectWithPaths GraphUserTraverseSystem(ctx, userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the Systems bound to a User



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
    userId := "userId_example" // string | ObjectID of the User.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserTraverseSystem(context.Background(), userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserTraverseSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserTraverseSystem`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserTraverseSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserTraverseSystemRequest struct via the builder pattern


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


## GraphUserTraverseSystemGroup

> []GraphObjectWithPaths GraphUserTraverseSystemGroup(ctx, userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()

List the System Groups bound to a User



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
    userId := "userId_example" // string | ObjectID of the User.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.GraphUserTraverseSystemGroup(context.Background(), userId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.GraphUserTraverseSystemGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphUserTraverseSystemGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.GraphUserTraverseSystemGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ObjectID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGraphUserTraverseSystemGroupRequest struct via the builder pattern


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


## PolicystatusesSystemsList

> []PolicyResult PolicystatusesSystemsList(ctx, systemId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()

List the policy statuses for a system



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
    systemId := "systemId_example" // string | ObjectID of the System.
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GraphApi.PolicystatusesSystemsList(context.Background(), systemId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphApi.PolicystatusesSystemsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicystatusesSystemsList`: []PolicyResult
    fmt.Fprintf(os.Stdout, "Response from `GraphApi.PolicystatusesSystemsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** | ObjectID of the System. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicystatusesSystemsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]PolicyResult**](PolicyResult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


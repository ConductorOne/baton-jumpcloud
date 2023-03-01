# \Office365Api

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphOffice365AssociationsList**](Office365Api.md#GraphOffice365AssociationsList) | **Get** /office365s/{office365_id}/associations | List the associations of an Office 365 instance
[**GraphOffice365AssociationsPost**](Office365Api.md#GraphOffice365AssociationsPost) | **Post** /office365s/{office365_id}/associations | Manage the associations of an Office 365 instance
[**GraphOffice365TraverseUser**](Office365Api.md#GraphOffice365TraverseUser) | **Get** /office365s/{office365_id}/users | List the Users bound to an Office 365 instance
[**GraphOffice365TraverseUserGroup**](Office365Api.md#GraphOffice365TraverseUserGroup) | **Get** /office365s/{office365_id}/usergroups | List the User Groups bound to an Office 365 instance
[**Office365sGet**](Office365Api.md#Office365sGet) | **Get** /office365s/{office365_id} | Get Office 365 instance
[**Office365sListImportUsers**](Office365Api.md#Office365sListImportUsers) | **Get** /office365s/{office365_id}/import/users | Get a list of users to import from an Office 365 instance
[**Office365sPatch**](Office365Api.md#Office365sPatch) | **Patch** /office365s/{office365_id} | Update existing Office 365 instance.
[**TranslationRulesOffice365Delete**](Office365Api.md#TranslationRulesOffice365Delete) | **Delete** /office365s/{office365_id}/translationrules/{id} | Deletes a Office 365 translation rule
[**TranslationRulesOffice365Get**](Office365Api.md#TranslationRulesOffice365Get) | **Get** /office365s/{office365_id}/translationrules/{id} | Gets a specific Office 365 translation rule
[**TranslationRulesOffice365List**](Office365Api.md#TranslationRulesOffice365List) | **Get** /office365s/{office365_id}/translationrules | List all the Office 365 Translation Rules
[**TranslationRulesOffice365Post**](Office365Api.md#TranslationRulesOffice365Post) | **Post** /office365s/{office365_id}/translationrules | Create a new Office 365 Translation Rule



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
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 instance.
    targets := []string{"Targets_example"} // []string | Targets which a \"office_365\" can be associated to.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Office365Api.GraphOffice365AssociationsList(context.Background(), office365Id).Targets(targets).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.GraphOffice365AssociationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphOffice365AssociationsList`: []GraphConnection
    fmt.Fprintf(os.Stdout, "Response from `Office365Api.GraphOffice365AssociationsList`: %v\n", resp)
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
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 instance.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewGraphOperationOffice365("Id_example", "Op_example", "Type_example") // GraphOperationOffice365 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.Office365Api.GraphOffice365AssociationsPost(context.Background(), office365Id).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.GraphOffice365AssociationsPost``: %v\n", err)
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
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 suite.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Office365Api.GraphOffice365TraverseUser(context.Background(), office365Id).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.GraphOffice365TraverseUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphOffice365TraverseUser`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `Office365Api.GraphOffice365TraverseUser`: %v\n", resp)
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
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 suite.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Office365Api.GraphOffice365TraverseUserGroup(context.Background(), office365Id).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.GraphOffice365TraverseUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GraphOffice365TraverseUserGroup`: []GraphObjectWithPaths
    fmt.Fprintf(os.Stdout, "Response from `Office365Api.GraphOffice365TraverseUserGroup`: %v\n", resp)
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


## Office365sGet

> Office365 Office365sGet(ctx, office365Id).XOrgId(xOrgId).Execute()

Get Office 365 instance



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
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 instance.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Office365Api.Office365sGet(context.Background(), office365Id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.Office365sGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Office365sGet`: Office365
    fmt.Fprintf(os.Stdout, "Response from `Office365Api.Office365sGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** | ObjectID of the Office 365 instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOffice365sGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**Office365**](Office365.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Office365sListImportUsers

> Office365sListImportUsers200Response Office365sListImportUsers(ctx, office365Id).ConsistencyLevel(consistencyLevel).Top(top).SkipToken(skipToken).Filter(filter).Search(search).Orderby(orderby).Count(count).Execute()

Get a list of users to import from an Office 365 instance



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
    office365Id := "office365Id_example" // string | 
    consistencyLevel := "consistencyLevel_example" // string | Defines the consistency header for O365 requests. See https://docs.microsoft.com/en-us/graph/api/user-list?view=graph-rest-1.0&tabs=http#request-headers (optional)
    top := int32(56) // int32 | Office 365 API maximum number of results per page. See https://docs.microsoft.com/en-us/graph/paging. (optional)
    skipToken := "skipToken_example" // string | Office 365 API token used to access the next page of results. See https://docs.microsoft.com/en-us/graph/paging. (optional)
    filter := "filter_example" // string | Office 365 API filter parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view=graph-rest-1.0&tabs=http#optional-query-parameters. (optional)
    search := "search_example" // string | Office 365 API search parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view=graph-rest-1.0&tabs=http#optional-query-parameters. (optional)
    orderby := "orderby_example" // string | Office 365 API orderby parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view=graph-rest-1.0&tabs=http#optional-query-parameters. (optional)
    count := true // bool | Office 365 API count parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view=graph-rest-1.0&tabs=http#optional-query-parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Office365Api.Office365sListImportUsers(context.Background(), office365Id).ConsistencyLevel(consistencyLevel).Top(top).SkipToken(skipToken).Filter(filter).Search(search).Orderby(orderby).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.Office365sListImportUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Office365sListImportUsers`: Office365sListImportUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `Office365Api.Office365sListImportUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOffice365sListImportUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consistencyLevel** | **string** | Defines the consistency header for O365 requests. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#request-headers | 
 **top** | **int32** | Office 365 API maximum number of results per page. See https://docs.microsoft.com/en-us/graph/paging. | 
 **skipToken** | **string** | Office 365 API token used to access the next page of results. See https://docs.microsoft.com/en-us/graph/paging. | 
 **filter** | **string** | Office 365 API filter parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters. | 
 **search** | **string** | Office 365 API search parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters. | 
 **orderby** | **string** | Office 365 API orderby parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters. | 
 **count** | **bool** | Office 365 API count parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters. | 

### Return type

[**Office365sListImportUsers200Response**](Office365sListImportUsers200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Office365sPatch

> Office365 Office365sPatch(ctx, office365Id).XOrgId(xOrgId).Body(body).Execute()

Update existing Office 365 instance.



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
    office365Id := "office365Id_example" // string | ObjectID of the Office 365 instance.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewOffice365() // Office365 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Office365Api.Office365sPatch(context.Background(), office365Id).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.Office365sPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Office365sPatch`: Office365
    fmt.Fprintf(os.Stdout, "Response from `Office365Api.Office365sPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** | ObjectID of the Office 365 instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOffice365sPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**Office365**](Office365.md) |  | 

### Return type

[**Office365**](Office365.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationRulesOffice365Delete

> TranslationRulesOffice365Delete(ctx, office365Id, id).Execute()

Deletes a Office 365 translation rule



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
    office365Id := "office365Id_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.Office365Api.TranslationRulesOffice365Delete(context.Background(), office365Id, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.TranslationRulesOffice365Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranslationRulesOffice365DeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## TranslationRulesOffice365Get

> Office365TranslationRule TranslationRulesOffice365Get(ctx, office365Id, id).Execute()

Gets a specific Office 365 translation rule



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
    office365Id := "office365Id_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Office365Api.TranslationRulesOffice365Get(context.Background(), office365Id, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.TranslationRulesOffice365Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TranslationRulesOffice365Get`: Office365TranslationRule
    fmt.Fprintf(os.Stdout, "Response from `Office365Api.TranslationRulesOffice365Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranslationRulesOffice365GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Office365TranslationRule**](Office365TranslationRule.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationRulesOffice365List

> []Office365TranslationRule TranslationRulesOffice365List(ctx, office365Id).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

List all the Office 365 Translation Rules



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
    office365Id := "office365Id_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Office365Api.TranslationRulesOffice365List(context.Background(), office365Id).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.TranslationRulesOffice365List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TranslationRulesOffice365List`: []Office365TranslationRule
    fmt.Fprintf(os.Stdout, "Response from `Office365Api.TranslationRulesOffice365List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranslationRulesOffice365ListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**[]Office365TranslationRule**](Office365TranslationRule.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslationRulesOffice365Post

> Office365TranslationRule TranslationRulesOffice365Post(ctx, office365Id).Body(body).Execute()

Create a new Office 365 Translation Rule



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
    office365Id := "office365Id_example" // string | 
    body := *openapiclient.NewOffice365TranslationRuleRequest() // Office365TranslationRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Office365Api.TranslationRulesOffice365Post(context.Background(), office365Id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365Api.TranslationRulesOffice365Post``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TranslationRulesOffice365Post`: Office365TranslationRule
    fmt.Fprintf(os.Stdout, "Response from `Office365Api.TranslationRulesOffice365Post`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**office365Id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTranslationRulesOffice365PostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Office365TranslationRuleRequest**](Office365TranslationRuleRequest.md) |  | 

### Return type

[**Office365TranslationRule**](Office365TranslationRule.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


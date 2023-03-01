# \GSuiteImportApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GsuitesListImportJumpcloudUsers**](GSuiteImportApi.md#GsuitesListImportJumpcloudUsers) | **Get** /gsuites/{gsuite_id}/import/jumpcloudusers | Get a list of users in Jumpcloud format to import from a Google Workspace account.
[**GsuitesListImportUsers**](GSuiteImportApi.md#GsuitesListImportUsers) | **Get** /gsuites/{gsuite_id}/import/users | Get a list of users to import from a G Suite instance



## GsuitesListImportJumpcloudUsers

> GsuitesListImportJumpcloudUsers200Response GsuitesListImportJumpcloudUsers(ctx, gsuiteId).MaxResults(maxResults).OrderBy(orderBy).PageToken(pageToken).Query(query).SortOrder(sortOrder).Execute()

Get a list of users in Jumpcloud format to import from a Google Workspace account.



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
    gsuiteId := "gsuiteId_example" // string | 
    maxResults := int32(56) // int32 | Google Directory API maximum number of results per page. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. (optional)
    orderBy := "orderBy_example" // string | Google Directory API sort field parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. (optional)
    pageToken := "pageToken_example" // string | Google Directory API token used to access the next page of results. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. (optional)
    query := "query_example" // string | Google Directory API search parameter. See https://developers.google.com/admin-sdk/directory/v1/guides/search-users. (optional)
    sortOrder := "sortOrder_example" // string | Google Directory API sort direction parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GSuiteImportApi.GsuitesListImportJumpcloudUsers(context.Background(), gsuiteId).MaxResults(maxResults).OrderBy(orderBy).PageToken(pageToken).Query(query).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GSuiteImportApi.GsuitesListImportJumpcloudUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GsuitesListImportJumpcloudUsers`: GsuitesListImportJumpcloudUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `GSuiteImportApi.GsuitesListImportJumpcloudUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gsuiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGsuitesListImportJumpcloudUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxResults** | **int32** | Google Directory API maximum number of results per page. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **orderBy** | **string** | Google Directory API sort field parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **pageToken** | **string** | Google Directory API token used to access the next page of results. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **query** | **string** | Google Directory API search parameter. See https://developers.google.com/admin-sdk/directory/v1/guides/search-users. | 
 **sortOrder** | **string** | Google Directory API sort direction parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 

### Return type

[**GsuitesListImportJumpcloudUsers200Response**](GsuitesListImportJumpcloudUsers200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GsuitesListImportUsers

> GsuitesListImportUsers200Response GsuitesListImportUsers(ctx, gsuiteId).Limit(limit).MaxResults(maxResults).OrderBy(orderBy).PageToken(pageToken).Query(query).SortOrder(sortOrder).Execute()

Get a list of users to import from a G Suite instance



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
    gsuiteId := "gsuiteId_example" // string | 
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    maxResults := int32(56) // int32 | Google Directory API maximum number of results per page. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. (optional)
    orderBy := "orderBy_example" // string | Google Directory API sort field parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. (optional)
    pageToken := "pageToken_example" // string | Google Directory API token used to access the next page of results. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. (optional)
    query := "query_example" // string | Google Directory API search parameter. See https://developers.google.com/admin-sdk/directory/v1/guides/search-users. (optional)
    sortOrder := "sortOrder_example" // string | Google Directory API sort direction parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GSuiteImportApi.GsuitesListImportUsers(context.Background(), gsuiteId).Limit(limit).MaxResults(maxResults).OrderBy(orderBy).PageToken(pageToken).Query(query).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GSuiteImportApi.GsuitesListImportUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GsuitesListImportUsers`: GsuitesListImportUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `GSuiteImportApi.GsuitesListImportUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gsuiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGsuitesListImportUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **maxResults** | **int32** | Google Directory API maximum number of results per page. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **orderBy** | **string** | Google Directory API sort field parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **pageToken** | **string** | Google Directory API token used to access the next page of results. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 
 **query** | **string** | Google Directory API search parameter. See https://developers.google.com/admin-sdk/directory/v1/guides/search-users. | 
 **sortOrder** | **string** | Google Directory API sort direction parameter. See https://developers.google.com/admin-sdk/directory/reference/rest/v1/users/list. | 

### Return type

[**GsuitesListImportUsers200Response**](GsuitesListImportUsers200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


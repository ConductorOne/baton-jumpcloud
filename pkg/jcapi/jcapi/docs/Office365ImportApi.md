# \Office365ImportApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Office365sListImportUsers**](Office365ImportApi.md#Office365sListImportUsers) | **Get** /office365s/{office365_id}/import/users | Get a list of users to import from an Office 365 instance



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
    resp, r, err := apiClient.Office365ImportApi.Office365sListImportUsers(context.Background(), office365Id).ConsistencyLevel(consistencyLevel).Top(top).SkipToken(skipToken).Filter(filter).Search(search).Orderby(orderby).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Office365ImportApi.Office365sListImportUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Office365sListImportUsers`: Office365sListImportUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `Office365ImportApi.Office365sListImportUsers`: %v\n", resp)
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


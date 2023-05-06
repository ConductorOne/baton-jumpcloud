# \SCIMImportApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportUsers**](SCIMImportApi.md#ImportUsers) | **Get** /applications/{application_id}/import/users | Get a list of users to import from an Application IdM service provider



## ImportUsers

> ImportUsersResponse ImportUsers(ctx, applicationId).Filter(filter).Query(query).Sort(sort).SortOrder(sortOrder).XOrgId(xOrgId).Limit(limit).Skip(skip).Execute()

Get a list of users to import from an Application IdM service provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    applicationId := "applicationId_example" // string | ObjectID of the Application.
    filter := "filter_example" // string | Filter users by a search term (optional)
    query := "query_example" // string | URL query to merge with the service provider request (optional)
    sort := "sort_example" // string | Sort users by supported fields (optional)
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SCIMImportApi.ImportUsers(context.Background(), applicationId).Filter(filter).Query(query).Sort(sort).SortOrder(sortOrder).XOrgId(xOrgId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMImportApi.ImportUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportUsers`: ImportUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMImportApi.ImportUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | ObjectID of the Application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter users by a search term | 
 **query** | **string** | URL query to merge with the service provider request | 
 **sort** | **string** | Sort users by supported fields | 
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]

### Return type

[**ImportUsersResponse**](ImportUsersResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


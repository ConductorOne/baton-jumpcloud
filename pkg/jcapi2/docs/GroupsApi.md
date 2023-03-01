# \GroupsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsList**](GroupsApi.md#GroupsList) | **Get** /groups | List All Groups



## GroupsList

> []Group GroupsList(ctx).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).XUnfilteredTotalCount(xUnfilteredTotalCount).Execute()

List All Groups



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
    xUnfilteredTotalCount := int32(56) // int32 | If provided in the request with any non-empty value, this header will be returned on the response populated with the total count of objects without filters taken into account (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GroupsList(context.Background()).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).XUnfilteredTotalCount(xUnfilteredTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsList`: []Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **xUnfilteredTotalCount** | **int32** | If provided in the request with any non-empty value, this header will be returned on the response populated with the total count of objects without filters taken into account | 

### Return type

[**[]Group**](Group.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


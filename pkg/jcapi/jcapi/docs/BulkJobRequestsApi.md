# \BulkJobRequestsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUserStatesCreate**](BulkJobRequestsApi.md#BulkUserStatesCreate) | **Post** /bulk/userstates | Create Scheduled Userstate Job
[**BulkUserStatesDelete**](BulkJobRequestsApi.md#BulkUserStatesDelete) | **Delete** /bulk/userstates/{id} | Delete Scheduled Userstate Job
[**BulkUserStatesGetNextScheduled**](BulkJobRequestsApi.md#BulkUserStatesGetNextScheduled) | **Get** /bulk/userstates/eventlist/next | Get the next scheduled state change for a list of users
[**BulkUserStatesList**](BulkJobRequestsApi.md#BulkUserStatesList) | **Get** /bulk/userstates | List Scheduled Userstate Change Jobs
[**BulkUsersCreate**](BulkJobRequestsApi.md#BulkUsersCreate) | **Post** /bulk/users | Bulk Users Create
[**BulkUsersCreateResults**](BulkJobRequestsApi.md#BulkUsersCreateResults) | **Get** /bulk/users/{job_id}/results | List Bulk Users Results
[**BulkUsersUpdate**](BulkJobRequestsApi.md#BulkUsersUpdate) | **Patch** /bulk/users | Bulk Users Update



## BulkUserStatesCreate

> []ScheduledUserstateResult BulkUserStatesCreate(ctx).XOrgId(xOrgId).Body(body).Execute()

Create Scheduled Userstate Job



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewBulkScheduledStatechangeCreate(time.Now(), "State_example", []string{"UserIds_example"}) // BulkScheduledStatechangeCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkJobRequestsApi.BulkUserStatesCreate(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkJobRequestsApi.BulkUserStatesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUserStatesCreate`: []ScheduledUserstateResult
    fmt.Fprintf(os.Stdout, "Response from `BulkJobRequestsApi.BulkUserStatesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUserStatesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**BulkScheduledStatechangeCreate**](BulkScheduledStatechangeCreate.md) |  | 

### Return type

[**[]ScheduledUserstateResult**](ScheduledUserstateResult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUserStatesDelete

> BulkUserStatesDelete(ctx, id).XOrgId(xOrgId).Execute()

Delete Scheduled Userstate Job



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
    id := "id_example" // string | Unique identifier of the scheduled statechange job.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BulkJobRequestsApi.BulkUserStatesDelete(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkJobRequestsApi.BulkUserStatesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the scheduled statechange job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUserStatesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUserStatesGetNextScheduled

> BulkUserStatesGetNextScheduled200Response BulkUserStatesGetNextScheduled(ctx).Users(users).Limit(limit).Skip(skip).Execute()

Get the next scheduled state change for a list of users



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
    users := []string{"Inner_example"} // []string | A list of system user IDs, limited to 100 items.
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkJobRequestsApi.BulkUserStatesGetNextScheduled(context.Background()).Users(users).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkJobRequestsApi.BulkUserStatesGetNextScheduled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUserStatesGetNextScheduled`: BulkUserStatesGetNextScheduled200Response
    fmt.Fprintf(os.Stdout, "Response from `BulkJobRequestsApi.BulkUserStatesGetNextScheduled`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUserStatesGetNextScheduledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **users** | **[]string** | A list of system user IDs, limited to 100 items. | 
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]

### Return type

[**BulkUserStatesGetNextScheduled200Response**](BulkUserStatesGetNextScheduled200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUserStatesList

> []ScheduledUserstateResult BulkUserStatesList(ctx).Limit(limit).Filter(filter).Skip(skip).XOrgId(xOrgId).Userid(userid).Execute()

List Scheduled Userstate Change Jobs



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
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    userid := "userid_example" // string | The systemuser id to filter by. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkJobRequestsApi.BulkUserStatesList(context.Background()).Limit(limit).Filter(filter).Skip(skip).XOrgId(xOrgId).Userid(userid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkJobRequestsApi.BulkUserStatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUserStatesList`: []ScheduledUserstateResult
    fmt.Fprintf(os.Stdout, "Response from `BulkJobRequestsApi.BulkUserStatesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUserStatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **userid** | **string** | The systemuser id to filter by. | 

### Return type

[**[]ScheduledUserstateResult**](ScheduledUserstateResult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUsersCreate

> JobId BulkUsersCreate(ctx).XOrgId(xOrgId).CreationSource(creationSource).Body(body).Execute()

Bulk Users Create



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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    creationSource := "creationSource_example" // string | Defines the creation-source header for gapps, o365 and workdays requests. If the header isn't sent, the default value is `jumpcloud:bulk`, if you send the header with a malformed value you receive a 400 error.  (optional) (default to "jumpcloud:bulk")
    body := []openapiclient.BulkUserCreate{*openapiclient.NewBulkUserCreate()} // []BulkUserCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkJobRequestsApi.BulkUsersCreate(context.Background()).XOrgId(xOrgId).CreationSource(creationSource).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkJobRequestsApi.BulkUsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUsersCreate`: JobId
    fmt.Fprintf(os.Stdout, "Response from `BulkJobRequestsApi.BulkUsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **creationSource** | **string** | Defines the creation-source header for gapps, o365 and workdays requests. If the header isn&#39;t sent, the default value is &#x60;jumpcloud:bulk&#x60;, if you send the header with a malformed value you receive a 400 error.  | [default to &quot;jumpcloud:bulk&quot;]
 **body** | [**[]BulkUserCreate**](BulkUserCreate.md) |  | 

### Return type

[**JobId**](JobId.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUsersCreateResults

> []JobWorkresult BulkUsersCreateResults(ctx, jobId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()

List Bulk Users Results



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
    jobId := "jobId_example" // string | 
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkJobRequestsApi.BulkUsersCreateResults(context.Background(), jobId).Limit(limit).Skip(skip).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkJobRequestsApi.BulkUsersCreateResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUsersCreateResults`: []JobWorkresult
    fmt.Fprintf(os.Stdout, "Response from `BulkJobRequestsApi.BulkUsersCreateResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUsersCreateResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]JobWorkresult**](JobWorkresult.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUsersUpdate

> JobId BulkUsersUpdate(ctx).XOrgId(xOrgId).Body(body).Execute()

Bulk Users Update



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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := []openapiclient.BulkUserUpdate{*openapiclient.NewBulkUserUpdate()} // []BulkUserUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkJobRequestsApi.BulkUsersUpdate(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkJobRequestsApi.BulkUsersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUsersUpdate`: JobId
    fmt.Fprintf(os.Stdout, "Response from `BulkJobRequestsApi.BulkUsersUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**[]BulkUserUpdate**](BulkUserUpdate.md) |  | 

### Return type

[**JobId**](JobId.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


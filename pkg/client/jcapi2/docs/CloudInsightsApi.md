# \CloudInsightsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudInsightsCreateAccount**](CloudInsightsApi.md#CloudInsightsCreateAccount) | **Post** /cloudinsights/accounts | Create AWS account record
[**CloudInsightsCreateSavedView**](CloudInsightsApi.md#CloudInsightsCreateSavedView) | **Post** /cloudinsights/views | 
[**CloudInsightsDeleteSavedView**](CloudInsightsApi.md#CloudInsightsDeleteSavedView) | **Delete** /cloudinsights/views/{view_id} | 
[**CloudInsightsFetchDistinctFieldValues**](CloudInsightsApi.md#CloudInsightsFetchDistinctFieldValues) | **Post** /cloudinsights/events/distinct | 
[**CloudInsightsFetchSavedView**](CloudInsightsApi.md#CloudInsightsFetchSavedView) | **Get** /cloudinsights/views/{view_id} | 
[**CloudInsightsFetchSavedViewsList**](CloudInsightsApi.md#CloudInsightsFetchSavedViewsList) | **Get** /cloudinsights/views | 
[**CloudInsightsGetAccountByID**](CloudInsightsApi.md#CloudInsightsGetAccountByID) | **Get** /cloudinsights/accounts/{cloud_insights_id} | 
[**CloudInsightsListAccounts**](CloudInsightsApi.md#CloudInsightsListAccounts) | **Get** /cloudinsights/accounts | 
[**CloudInsightsListEvents**](CloudInsightsApi.md#CloudInsightsListEvents) | **Post** /cloudinsights/events | 
[**CloudInsightsUpdateAccountByID**](CloudInsightsApi.md#CloudInsightsUpdateAccountByID) | **Put** /cloudinsights/accounts/{cloud_insights_id} | Update Cloud Insights Account record by ID
[**CloudInsightsUpdateSavedView**](CloudInsightsApi.md#CloudInsightsUpdateSavedView) | **Put** /cloudinsights/views/{view_id} | 



## CloudInsightsCreateAccount

> CloudInsightsAccountPostResponse CloudInsightsCreateAccount(ctx).XOrgId(xOrgId).Body(body).Execute()

Create AWS account record



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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewCloudInsightsAccountPost("AwsAccountId_example", "AwsRegion_example") // CloudInsightsAccountPost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsCreateAccount(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsCreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsCreateAccount`: CloudInsightsAccountPostResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsCreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**CloudInsightsAccountPost**](CloudInsightsAccountPost.md) |  | 

### Return type

[**CloudInsightsAccountPostResponse**](CloudInsightsAccountPostResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, text/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsCreateSavedView

> CloudInsightsSavedViewResponse CloudInsightsCreateSavedView(ctx).XOrgId(xOrgId).Body(body).Execute()





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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewCloudInsightsSavedViewPost("Name_example") // CloudInsightsSavedViewPost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsCreateSavedView(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsCreateSavedView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsCreateSavedView`: CloudInsightsSavedViewResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsCreateSavedView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsCreateSavedViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**CloudInsightsSavedViewPost**](CloudInsightsSavedViewPost.md) |  | 

### Return type

[**CloudInsightsSavedViewResponse**](CloudInsightsSavedViewResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsDeleteSavedView

> CloudInsightsSavedViewId CloudInsightsDeleteSavedView(ctx, viewId).XOrgId(xOrgId).Execute()





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
    viewId := "viewId_example" // string | objectId of the Cloud Insights Saved View record.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsDeleteSavedView(context.Background(), viewId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsDeleteSavedView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsDeleteSavedView`: CloudInsightsSavedViewId
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsDeleteSavedView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | objectId of the Cloud Insights Saved View record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsDeleteSavedViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsSavedViewId**](CloudInsightsSavedViewId.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsFetchDistinctFieldValues

> CloudInsightsEventsDistinctResponse CloudInsightsFetchDistinctFieldValues(ctx).XOrgId(xOrgId).Body(body).Execute()





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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewCloudInsightsEventsDistinctPost([]string{"AwsAccounts_example"}, *openapiclient.NewCloudInsightsEventsDistinctPostEndTime()) // CloudInsightsEventsDistinctPost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsFetchDistinctFieldValues(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsFetchDistinctFieldValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsFetchDistinctFieldValues`: CloudInsightsEventsDistinctResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsFetchDistinctFieldValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsFetchDistinctFieldValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**CloudInsightsEventsDistinctPost**](CloudInsightsEventsDistinctPost.md) |  | 

### Return type

[**CloudInsightsEventsDistinctResponse**](CloudInsightsEventsDistinctResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsFetchSavedView

> CloudInsightsSavedViewResponse CloudInsightsFetchSavedView(ctx, viewId).XOrgId(xOrgId).Execute()





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
    viewId := "viewId_example" // string | objectId of the Cloud Insights Saved View record.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsFetchSavedView(context.Background(), viewId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsFetchSavedView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsFetchSavedView`: CloudInsightsSavedViewResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsFetchSavedView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | objectId of the Cloud Insights Saved View record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsFetchSavedViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsSavedViewResponse**](CloudInsightsSavedViewResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsFetchSavedViewsList

> []CloudInsightsSavedViewResponse CloudInsightsFetchSavedViewsList(ctx).XOrgId(xOrgId).Execute()





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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsFetchSavedViewsList(context.Background()).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsFetchSavedViewsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsFetchSavedViewsList`: []CloudInsightsSavedViewResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsFetchSavedViewsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsFetchSavedViewsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]CloudInsightsSavedViewResponse**](CloudInsightsSavedViewResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsGetAccountByID

> CloudInsightsAccountGet CloudInsightsGetAccountByID(ctx, cloudInsightsId).XOrgId(xOrgId).Execute()





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
    cloudInsightsId := "cloudInsightsId_example" // string | objectId of the Cloud Insights instance.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsGetAccountByID(context.Background(), cloudInsightsId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsGetAccountByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsGetAccountByID`: CloudInsightsAccountGet
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsGetAccountByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudInsightsId** | **string** | objectId of the Cloud Insights instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsGetAccountByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsAccountGet**](CloudInsightsAccountGet.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsListAccounts

> CloudInsightsAccountListResponse CloudInsightsListAccounts(ctx).XOrgId(xOrgId).Execute()





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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsListAccounts(context.Background()).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsListAccounts`: CloudInsightsAccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**CloudInsightsAccountListResponse**](CloudInsightsAccountListResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsListEvents

> CloudInsightsEventsListResponse CloudInsightsListEvents(ctx).XOrgId(xOrgId).Body(body).Execute()





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
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewCloudInsightsEventsListPost() // CloudInsightsEventsListPost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsListEvents(context.Background()).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsListEvents`: CloudInsightsEventsListResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**CloudInsightsEventsListPost**](CloudInsightsEventsListPost.md) |  | 

### Return type

[**CloudInsightsEventsListResponse**](CloudInsightsEventsListResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsUpdateAccountByID

> CloudInsightsAccountGet CloudInsightsUpdateAccountByID(ctx, cloudInsightsId).XOrgId(xOrgId).Body(body).Execute()

Update Cloud Insights Account record by ID



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
    cloudInsightsId := "cloudInsightsId_example" // string | objectId of the Cloud Insights instance.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewCloudInsightsAccountPut() // CloudInsightsAccountPut |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsUpdateAccountByID(context.Background(), cloudInsightsId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsUpdateAccountByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsUpdateAccountByID`: CloudInsightsAccountGet
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsUpdateAccountByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudInsightsId** | **string** | objectId of the Cloud Insights instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsUpdateAccountByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**CloudInsightsAccountPut**](CloudInsightsAccountPut.md) |  | 

### Return type

[**CloudInsightsAccountGet**](CloudInsightsAccountGet.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudInsightsUpdateSavedView

> CloudInsightsSavedViewResponse CloudInsightsUpdateSavedView(ctx, viewId).XOrgId(xOrgId).Body(body).Execute()





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
    viewId := "viewId_example" // string | objectId of the Cloud Insights Saved View record.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewCloudInsightsSavedViewPut() // CloudInsightsSavedViewPut |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudInsightsApi.CloudInsightsUpdateSavedView(context.Background(), viewId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudInsightsApi.CloudInsightsUpdateSavedView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudInsightsUpdateSavedView`: CloudInsightsSavedViewResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudInsightsApi.CloudInsightsUpdateSavedView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | objectId of the Cloud Insights Saved View record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudInsightsUpdateSavedViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**CloudInsightsSavedViewPut**](CloudInsightsSavedViewPut.md) |  | 

### Return type

[**CloudInsightsSavedViewResponse**](CloudInsightsSavedViewResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


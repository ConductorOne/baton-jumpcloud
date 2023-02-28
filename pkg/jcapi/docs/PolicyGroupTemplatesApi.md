# \PolicyGroupTemplatesApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PolicyGroupTemplatesGet**](PolicyGroupTemplatesApi.md#PolicyGroupTemplatesGet) | **Get** /providers/{provider_id}/policygrouptemplates/{id} | Gets a provider&#39;s policy group template.
[**PolicyGroupTemplatesGetConfiguredPolicyTemplate**](PolicyGroupTemplatesApi.md#PolicyGroupTemplatesGetConfiguredPolicyTemplate) | **Get** /providers/{provider_id}/configuredpolicytemplates/{id} | Retrieve a configured policy template by id.
[**PolicyGroupTemplatesList**](PolicyGroupTemplatesApi.md#PolicyGroupTemplatesList) | **Get** /providers/{provider_id}/policygrouptemplates | List a provider&#39;s policy group templates.
[**PolicyGroupTemplatesListConfiguredPolicyTemplates**](PolicyGroupTemplatesApi.md#PolicyGroupTemplatesListConfiguredPolicyTemplates) | **Get** /providers/{provider_id}/configuredpolicytemplates | List a provider&#39;s configured policy templates.
[**PolicyGroupTemplatesListMembers**](PolicyGroupTemplatesApi.md#PolicyGroupTemplatesListMembers) | **Get** /providers/{provider_id}/policygrouptemplates/{id}/members | Gets the list of members from a policy group template.



## PolicyGroupTemplatesGet

> PolicyGroupTemplate PolicyGroupTemplatesGet(ctx, providerId, id).Execute()

Gets a provider's policy group template.



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
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesGet(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupTemplatesApi.PolicyGroupTemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesGet`: PolicyGroupTemplate
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupTemplatesApi.PolicyGroupTemplatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PolicyGroupTemplate**](PolicyGroupTemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesGetConfiguredPolicyTemplate

> ConfiguredPolicyTemplate PolicyGroupTemplatesGetConfiguredPolicyTemplate(ctx, providerId, id).Execute()

Retrieve a configured policy template by id.



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
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupTemplatesApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesGetConfiguredPolicyTemplate`: ConfiguredPolicyTemplate
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupTemplatesApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesGetConfiguredPolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConfiguredPolicyTemplate**](ConfiguredPolicyTemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesList

> PolicyGroupTemplates PolicyGroupTemplatesList(ctx, providerId).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()

List a provider's policy group templates.



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
    providerId := "providerId_example" // string | 
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesList(context.Background(), providerId).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupTemplatesApi.PolicyGroupTemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesList`: PolicyGroupTemplates
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupTemplatesApi.PolicyGroupTemplatesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]

### Return type

[**PolicyGroupTemplates**](PolicyGroupTemplates.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesListConfiguredPolicyTemplates

> PolicyGroupTemplatesListConfiguredPolicyTemplates200Response PolicyGroupTemplatesListConfiguredPolicyTemplates(ctx, providerId).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()

List a provider's configured policy templates.



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
    providerId := "providerId_example" // string | 
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesListConfiguredPolicyTemplates(context.Background(), providerId).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupTemplatesApi.PolicyGroupTemplatesListConfiguredPolicyTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesListConfiguredPolicyTemplates`: PolicyGroupTemplatesListConfiguredPolicyTemplates200Response
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupTemplatesApi.PolicyGroupTemplatesListConfiguredPolicyTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesListConfiguredPolicyTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]

### Return type

[**PolicyGroupTemplatesListConfiguredPolicyTemplates200Response**](PolicyGroupTemplatesListConfiguredPolicyTemplates200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesListMembers

> PolicyGroupTemplateMembers PolicyGroupTemplatesListMembers(ctx, providerId, id).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()

Gets the list of members from a policy group template.



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
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesListMembers(context.Background(), providerId, id).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyGroupTemplatesApi.PolicyGroupTemplatesListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesListMembers`: PolicyGroupTemplateMembers
    fmt.Fprintf(os.Stdout, "Response from `PolicyGroupTemplatesApi.PolicyGroupTemplatesListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]

### Return type

[**PolicyGroupTemplateMembers**](PolicyGroupTemplateMembers.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ManagedServiceProviderApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdministratorOrganizationsCreateByAdministrator**](ManagedServiceProviderApi.md#AdministratorOrganizationsCreateByAdministrator) | **Post** /administrators/{id}/organizationlinks | Allow Adminstrator access to an Organization.
[**AdministratorOrganizationsListByAdministrator**](ManagedServiceProviderApi.md#AdministratorOrganizationsListByAdministrator) | **Get** /administrators/{id}/organizationlinks | List the association links between an Administrator and Organizations.
[**AdministratorOrganizationsListByOrganization**](ManagedServiceProviderApi.md#AdministratorOrganizationsListByOrganization) | **Get** /organizations/{id}/administratorlinks | List the association links between an Organization and Administrators.
[**AdministratorOrganizationsRemoveByAdministrator**](ManagedServiceProviderApi.md#AdministratorOrganizationsRemoveByAdministrator) | **Delete** /administrators/{administrator_id}/organizationlinks/{id} | Remove association between an Administrator and an Organization.
[**PolicyGroupTemplatesGet**](ManagedServiceProviderApi.md#PolicyGroupTemplatesGet) | **Get** /providers/{provider_id}/policygrouptemplates/{id} | Gets a provider&#39;s policy group template.
[**PolicyGroupTemplatesGetConfiguredPolicyTemplate**](ManagedServiceProviderApi.md#PolicyGroupTemplatesGetConfiguredPolicyTemplate) | **Get** /providers/{provider_id}/configuredpolicytemplates/{id} | Retrieve a configured policy template by id.
[**PolicyGroupTemplatesList**](ManagedServiceProviderApi.md#PolicyGroupTemplatesList) | **Get** /providers/{provider_id}/policygrouptemplates | List a provider&#39;s policy group templates.
[**PolicyGroupTemplatesListConfiguredPolicyTemplates**](ManagedServiceProviderApi.md#PolicyGroupTemplatesListConfiguredPolicyTemplates) | **Get** /providers/{provider_id}/configuredpolicytemplates | List a provider&#39;s configured policy templates.
[**PolicyGroupTemplatesListMembers**](ManagedServiceProviderApi.md#PolicyGroupTemplatesListMembers) | **Get** /providers/{provider_id}/policygrouptemplates/{id}/members | Gets the list of members from a policy group template.
[**ProviderOrganizationsCreateOrg**](ManagedServiceProviderApi.md#ProviderOrganizationsCreateOrg) | **Post** /providers/{provider_id}/organizations | Create Provider Organization
[**ProviderOrganizationsUpdateOrg**](ManagedServiceProviderApi.md#ProviderOrganizationsUpdateOrg) | **Put** /providers/{provider_id}/organizations/{id} | Update Provider Organization
[**ProvidersGetProvider**](ManagedServiceProviderApi.md#ProvidersGetProvider) | **Get** /providers/{provider_id} | Retrieve Provider
[**ProvidersListAdministrators**](ManagedServiceProviderApi.md#ProvidersListAdministrators) | **Get** /providers/{provider_id}/administrators | List Provider Administrators
[**ProvidersListOrganizations**](ManagedServiceProviderApi.md#ProvidersListOrganizations) | **Get** /providers/{provider_id}/organizations | List Provider Organizations
[**ProvidersPostAdmins**](ManagedServiceProviderApi.md#ProvidersPostAdmins) | **Post** /providers/{provider_id}/administrators | Create a new Provider Administrator
[**ProvidersRetrieveInvoice**](ManagedServiceProviderApi.md#ProvidersRetrieveInvoice) | **Get** /providers/{provider_id}/invoices/{ID} | Download a provider&#39;s invoice.
[**ProvidersRetrieveInvoices**](ManagedServiceProviderApi.md#ProvidersRetrieveInvoices) | **Get** /providers/{provider_id}/invoices | List a provider&#39;s invoices.



## AdministratorOrganizationsCreateByAdministrator

> AdministratorOrganizationLink AdministratorOrganizationsCreateByAdministrator(ctx, id).Body(body).Execute()

Allow Adminstrator access to an Organization.



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
    id := "id_example" // string | 
    body := *openapiclient.NewAdministratorOrganizationLinkReq() // AdministratorOrganizationLinkReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.AdministratorOrganizationsCreateByAdministrator(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.AdministratorOrganizationsCreateByAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdministratorOrganizationsCreateByAdministrator`: AdministratorOrganizationLink
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.AdministratorOrganizationsCreateByAdministrator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdministratorOrganizationsCreateByAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AdministratorOrganizationLinkReq**](AdministratorOrganizationLinkReq.md) |  | 

### Return type

[**AdministratorOrganizationLink**](AdministratorOrganizationLink.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdministratorOrganizationsListByAdministrator

> []AdministratorOrganizationLink AdministratorOrganizationsListByAdministrator(ctx, id).Limit(limit).Skip(skip).Execute()

List the association links between an Administrator and Organizations.



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
    id := "id_example" // string | 
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.AdministratorOrganizationsListByAdministrator(context.Background(), id).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.AdministratorOrganizationsListByAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdministratorOrganizationsListByAdministrator`: []AdministratorOrganizationLink
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.AdministratorOrganizationsListByAdministrator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdministratorOrganizationsListByAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]

### Return type

[**[]AdministratorOrganizationLink**](AdministratorOrganizationLink.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdministratorOrganizationsListByOrganization

> []AdministratorOrganizationLink AdministratorOrganizationsListByOrganization(ctx, id).Limit(limit).Skip(skip).Execute()

List the association links between an Organization and Administrators.



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
    id := "id_example" // string | 
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.AdministratorOrganizationsListByOrganization(context.Background(), id).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.AdministratorOrganizationsListByOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdministratorOrganizationsListByOrganization`: []AdministratorOrganizationLink
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.AdministratorOrganizationsListByOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdministratorOrganizationsListByOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]

### Return type

[**[]AdministratorOrganizationLink**](AdministratorOrganizationLink.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdministratorOrganizationsRemoveByAdministrator

> AdministratorOrganizationsRemoveByAdministrator(ctx, administratorId, id).Execute()

Remove association between an Administrator and an Organization.



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
    administratorId := "administratorId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedServiceProviderApi.AdministratorOrganizationsRemoveByAdministrator(context.Background(), administratorId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.AdministratorOrganizationsRemoveByAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**administratorId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdministratorOrganizationsRemoveByAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.PolicyGroupTemplatesGet(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.PolicyGroupTemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesGet`: PolicyGroupTemplate
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.PolicyGroupTemplatesGet`: %v\n", resp)
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
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesGetConfiguredPolicyTemplate`: ConfiguredPolicyTemplate
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate`: %v\n", resp)
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
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.PolicyGroupTemplatesList(context.Background(), providerId).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.PolicyGroupTemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesList`: PolicyGroupTemplates
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.PolicyGroupTemplatesList`: %v\n", resp)
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
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.PolicyGroupTemplatesListConfiguredPolicyTemplates(context.Background(), providerId).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.PolicyGroupTemplatesListConfiguredPolicyTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesListConfiguredPolicyTemplates`: PolicyGroupTemplatesListConfiguredPolicyTemplates200Response
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.PolicyGroupTemplatesListConfiguredPolicyTemplates`: %v\n", resp)
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
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
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
    resp, r, err := apiClient.ManagedServiceProviderApi.PolicyGroupTemplatesListMembers(context.Background(), providerId, id).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.PolicyGroupTemplatesListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesListMembers`: PolicyGroupTemplateMembers
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.PolicyGroupTemplatesListMembers`: %v\n", resp)
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


## ProviderOrganizationsCreateOrg

> Organization ProviderOrganizationsCreateOrg(ctx, providerId).Body(body).Execute()

Create Provider Organization



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
    providerId := "providerId_example" // string | 
    body := *openapiclient.NewCreateOrganization() // CreateOrganization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.ProviderOrganizationsCreateOrg(context.Background(), providerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.ProviderOrganizationsCreateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderOrganizationsCreateOrg`: Organization
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.ProviderOrganizationsCreateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderOrganizationsCreateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrganization**](CreateOrganization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderOrganizationsUpdateOrg

> Organization ProviderOrganizationsUpdateOrg(ctx, providerId, id).Body(body).Execute()

Update Provider Organization



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
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 
    body := *openapiclient.NewOrganization() // Organization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.ProviderOrganizationsUpdateOrg(context.Background(), providerId, id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.ProviderOrganizationsUpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderOrganizationsUpdateOrg`: Organization
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.ProviderOrganizationsUpdateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderOrganizationsUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Organization**](Organization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersGetProvider

> Provider ProvidersGetProvider(ctx, providerId).Fields(fields).Execute()

Retrieve Provider



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
    providerId := "providerId_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.ProvidersGetProvider(context.Background(), providerId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.ProvidersGetProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersGetProvider`: Provider
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.ProvidersGetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]

### Return type

[**Provider**](Provider.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersListAdministrators

> ProvidersListAdministrators200Response ProvidersListAdministrators(ctx, providerId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).SortIgnoreCase(sortIgnoreCase).Execute()

List Provider Administrators



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
    providerId := "providerId_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    sortIgnoreCase := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.ProvidersListAdministrators(context.Background(), providerId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).SortIgnoreCase(sortIgnoreCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.ProvidersListAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersListAdministrators`: ProvidersListAdministrators200Response
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.ProvidersListAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersListAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **sortIgnoreCase** | **[]string** | The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**ProvidersListAdministrators200Response**](ProvidersListAdministrators200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersListOrganizations

> ProvidersListOrganizations200Response ProvidersListOrganizations(ctx, providerId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).SortIgnoreCase(sortIgnoreCase).Execute()

List Provider Organizations



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
    providerId := "providerId_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    sortIgnoreCase := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.ProvidersListOrganizations(context.Background(), providerId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).SortIgnoreCase(sortIgnoreCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.ProvidersListOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersListOrganizations`: ProvidersListOrganizations200Response
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.ProvidersListOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersListOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **sortIgnoreCase** | **[]string** | The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**ProvidersListOrganizations200Response**](ProvidersListOrganizations200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersPostAdmins

> Administrator ProvidersPostAdmins(ctx, providerId).Body(body).Execute()

Create a new Provider Administrator



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
    providerId := "providerId_example" // string | 
    body := *openapiclient.NewProviderAdminReq("Email_example") // ProviderAdminReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.ProvidersPostAdmins(context.Background(), providerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.ProvidersPostAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersPostAdmins`: Administrator
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.ProvidersPostAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersPostAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProviderAdminReq**](ProviderAdminReq.md) |  | 

### Return type

[**Administrator**](Administrator.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersRetrieveInvoice

> *os.File ProvidersRetrieveInvoice(ctx, providerId, iD).Execute()

Download a provider's invoice.



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
    providerId := "providerId_example" // string | 
    iD := "iD_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.ProvidersRetrieveInvoice(context.Background(), providerId, iD).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.ProvidersRetrieveInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersRetrieveInvoice`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.ProvidersRetrieveInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**iD** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersRetrieveInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersRetrieveInvoices

> ProviderInvoiceResponse ProvidersRetrieveInvoices(ctx, providerId).Skip(skip).Sort(sort).Limit(limit).Execute()

List a provider's invoices.



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
    providerId := "providerId_example" // string | 
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManagedServiceProviderApi.ProvidersRetrieveInvoices(context.Background(), providerId).Skip(skip).Sort(sort).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedServiceProviderApi.ProvidersRetrieveInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersRetrieveInvoices`: ProviderInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagedServiceProviderApi.ProvidersRetrieveInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersRetrieveInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]

### Return type

[**ProviderInvoiceResponse**](ProviderInvoiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


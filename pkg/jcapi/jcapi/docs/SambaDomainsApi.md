# \SambaDomainsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LdapserversSambaDomainsDelete**](SambaDomainsApi.md#LdapserversSambaDomainsDelete) | **Delete** /ldapservers/{ldapserver_id}/sambadomains/{id} | Delete Samba Domain
[**LdapserversSambaDomainsGet**](SambaDomainsApi.md#LdapserversSambaDomainsGet) | **Get** /ldapservers/{ldapserver_id}/sambadomains/{id} | Get Samba Domain
[**LdapserversSambaDomainsList**](SambaDomainsApi.md#LdapserversSambaDomainsList) | **Get** /ldapservers/{ldapserver_id}/sambadomains | List Samba Domains
[**LdapserversSambaDomainsPost**](SambaDomainsApi.md#LdapserversSambaDomainsPost) | **Post** /ldapservers/{ldapserver_id}/sambadomains | Create Samba Domain
[**LdapserversSambaDomainsPut**](SambaDomainsApi.md#LdapserversSambaDomainsPut) | **Put** /ldapservers/{ldapserver_id}/sambadomains/{id} | Update Samba Domain



## LdapserversSambaDomainsDelete

> string LdapserversSambaDomainsDelete(ctx, ldapserverId, id).XOrgId(xOrgId).Execute()

Delete Samba Domain



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
    ldapserverId := "ldapserverId_example" // string | Unique identifier of the LDAP server.
    id := "id_example" // string | Unique identifier of the samba domain.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SambaDomainsApi.LdapserversSambaDomainsDelete(context.Background(), ldapserverId, id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SambaDomainsApi.LdapserversSambaDomainsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LdapserversSambaDomainsDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `SambaDomainsApi.LdapserversSambaDomainsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapserverId** | **string** | Unique identifier of the LDAP server. | 
**id** | **string** | Unique identifier of the samba domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLdapserversSambaDomainsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapserversSambaDomainsGet

> SambaDomain LdapserversSambaDomainsGet(ctx, ldapserverId, id).XOrgId(xOrgId).Execute()

Get Samba Domain



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
    ldapserverId := "ldapserverId_example" // string | Unique identifier of the LDAP server.
    id := "id_example" // string | Unique identifier of the samba domain.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SambaDomainsApi.LdapserversSambaDomainsGet(context.Background(), ldapserverId, id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SambaDomainsApi.LdapserversSambaDomainsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LdapserversSambaDomainsGet`: SambaDomain
    fmt.Fprintf(os.Stdout, "Response from `SambaDomainsApi.LdapserversSambaDomainsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapserverId** | **string** | Unique identifier of the LDAP server. | 
**id** | **string** | Unique identifier of the samba domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLdapserversSambaDomainsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**SambaDomain**](SambaDomain.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapserversSambaDomainsList

> []SambaDomain LdapserversSambaDomainsList(ctx, ldapserverId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()

List Samba Domains



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
    ldapserverId := "ldapserverId_example" // string | Unique identifier of the LDAP server.
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SambaDomainsApi.LdapserversSambaDomainsList(context.Background(), ldapserverId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SambaDomainsApi.LdapserversSambaDomainsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LdapserversSambaDomainsList`: []SambaDomain
    fmt.Fprintf(os.Stdout, "Response from `SambaDomainsApi.LdapserversSambaDomainsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapserverId** | **string** | Unique identifier of the LDAP server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLdapserversSambaDomainsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]SambaDomain**](SambaDomain.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapserversSambaDomainsPost

> SambaDomain LdapserversSambaDomainsPost(ctx, ldapserverId).XOrgId(xOrgId).Body(body).Execute()

Create Samba Domain



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
    ldapserverId := "ldapserverId_example" // string | Unique identifier of the LDAP server.
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewSambaDomain("Name_example", "Sid_example") // SambaDomain |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SambaDomainsApi.LdapserversSambaDomainsPost(context.Background(), ldapserverId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SambaDomainsApi.LdapserversSambaDomainsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LdapserversSambaDomainsPost`: SambaDomain
    fmt.Fprintf(os.Stdout, "Response from `SambaDomainsApi.LdapserversSambaDomainsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapserverId** | **string** | Unique identifier of the LDAP server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLdapserversSambaDomainsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**SambaDomain**](SambaDomain.md) |  | 

### Return type

[**SambaDomain**](SambaDomain.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LdapserversSambaDomainsPut

> SambaDomain LdapserversSambaDomainsPut(ctx, ldapserverId, id).Body(body).Execute()

Update Samba Domain



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
    ldapserverId := "ldapserverId_example" // string | Unique identifier of the LDAP server.
    id := "id_example" // string | Unique identifier of the samba domain.
    body := *openapiclient.NewSambaDomain("Name_example", "Sid_example") // SambaDomain |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SambaDomainsApi.LdapserversSambaDomainsPut(context.Background(), ldapserverId, id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SambaDomainsApi.LdapserversSambaDomainsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LdapserversSambaDomainsPut`: SambaDomain
    fmt.Fprintf(os.Stdout, "Response from `SambaDomainsApi.LdapserversSambaDomainsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapserverId** | **string** | Unique identifier of the LDAP server. | 
**id** | **string** | Unique identifier of the samba domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLdapserversSambaDomainsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SambaDomain**](SambaDomain.md) |  | 

### Return type

[**SambaDomain**](SambaDomain.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AdministratorsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdministratorOrganizationsCreateByAdministrator**](AdministratorsApi.md#AdministratorOrganizationsCreateByAdministrator) | **Post** /administrators/{id}/organizationlinks | Allow Adminstrator access to an Organization.
[**AdministratorOrganizationsListByAdministrator**](AdministratorsApi.md#AdministratorOrganizationsListByAdministrator) | **Get** /administrators/{id}/organizationlinks | List the association links between an Administrator and Organizations.
[**AdministratorOrganizationsListByOrganization**](AdministratorsApi.md#AdministratorOrganizationsListByOrganization) | **Get** /organizations/{id}/administratorlinks | List the association links between an Organization and Administrators.
[**AdministratorOrganizationsRemoveByAdministrator**](AdministratorsApi.md#AdministratorOrganizationsRemoveByAdministrator) | **Delete** /administrators/{administrator_id}/organizationlinks/{id} | Remove association between an Administrator and an Organization.



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    id := "id_example" // string | 
    body := *openapiclient.NewAdministratorOrganizationLinkReq() // AdministratorOrganizationLinkReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministratorsApi.AdministratorOrganizationsCreateByAdministrator(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsApi.AdministratorOrganizationsCreateByAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdministratorOrganizationsCreateByAdministrator`: AdministratorOrganizationLink
    fmt.Fprintf(os.Stdout, "Response from `AdministratorsApi.AdministratorOrganizationsCreateByAdministrator`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    id := "id_example" // string | 
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministratorsApi.AdministratorOrganizationsListByAdministrator(context.Background(), id).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsApi.AdministratorOrganizationsListByAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdministratorOrganizationsListByAdministrator`: []AdministratorOrganizationLink
    fmt.Fprintf(os.Stdout, "Response from `AdministratorsApi.AdministratorOrganizationsListByAdministrator`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    id := "id_example" // string | 
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministratorsApi.AdministratorOrganizationsListByOrganization(context.Background(), id).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsApi.AdministratorOrganizationsListByOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdministratorOrganizationsListByOrganization`: []AdministratorOrganizationLink
    fmt.Fprintf(os.Stdout, "Response from `AdministratorsApi.AdministratorOrganizationsListByOrganization`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    administratorId := "administratorId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdministratorsApi.AdministratorOrganizationsRemoveByAdministrator(context.Background(), administratorId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorsApi.AdministratorOrganizationsRemoveByAdministrator``: %v\n", err)
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


# \SystemusersApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SshkeyDelete**](SystemusersApi.md#SshkeyDelete) | **Delete** /systemusers/{systemuser_id}/sshkeys/{id} | Delete a system user&#39;s Public SSH Keys
[**SshkeyList**](SystemusersApi.md#SshkeyList) | **Get** /systemusers/{id}/sshkeys | List a system user&#39;s public SSH keys
[**SshkeyPost**](SystemusersApi.md#SshkeyPost) | **Post** /systemusers/{id}/sshkeys | Create a system user&#39;s Public SSH Key
[**SystemusersDelete**](SystemusersApi.md#SystemusersDelete) | **Delete** /systemusers/{id} | Delete a system user
[**SystemusersExpire**](SystemusersApi.md#SystemusersExpire) | **Post** /systemusers/{id}/expire | Expire a system user&#39;s password
[**SystemusersGet**](SystemusersApi.md#SystemusersGet) | **Get** /systemusers/{id} | List a system user
[**SystemusersList**](SystemusersApi.md#SystemusersList) | **Get** /systemusers | List all system users
[**SystemusersMfasync**](SystemusersApi.md#SystemusersMfasync) | **Post** /systemusers/{id}/mfasync | Sync a systemuser&#39;s mfa enrollment status
[**SystemusersPost**](SystemusersApi.md#SystemusersPost) | **Post** /systemusers | Create a system user
[**SystemusersPut**](SystemusersApi.md#SystemusersPut) | **Put** /systemusers/{id} | Update a system user
[**SystemusersResetmfa**](SystemusersApi.md#SystemusersResetmfa) | **Post** /systemusers/{id}/resetmfa | Reset a system user&#39;s MFA token
[**SystemusersStateActivate**](SystemusersApi.md#SystemusersStateActivate) | **Post** /systemusers/{id}/state/activate | Activate System User
[**SystemusersUnlock**](SystemusersApi.md#SystemusersUnlock) | **Post** /systemusers/{id}/unlock | Unlock a system user



## SshkeyDelete

> string SshkeyDelete(ctx, systemuserId, id).XOrgId(xOrgId).Execute()

Delete a system user's Public SSH Keys



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    systemuserId := "systemuserId_example" // string | 
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SshkeyDelete(context.Background(), systemuserId, id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SshkeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshkeyDelete`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SshkeyDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemuserId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshkeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** |  | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshkeyList

> []Sshkeylist SshkeyList(ctx, id).XOrgId(xOrgId).Execute()

List a system user's public SSH keys



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SshkeyList(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SshkeyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshkeyList`: []Sshkeylist
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SshkeyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshkeyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** |  | 

### Return type

[**[]Sshkeylist**](Sshkeylist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshkeyPost

> Sshkeylist SshkeyPost(ctx, id).XOrgId(xOrgId).Body(body).Execute()

Create a system user's Public SSH Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string |  (optional)
    body := *openapiclient.NewSshkeypost("Name_example", "PublicKey_example") // Sshkeypost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SshkeyPost(context.Background(), id).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SshkeyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshkeyPost`: Sshkeylist
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SshkeyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshkeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** |  | 
 **body** | [**Sshkeypost**](Sshkeypost.md) |  | 

### Return type

[**Sshkeylist**](Sshkeylist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemusersDelete

> Systemuserreturn SystemusersDelete(ctx, id).XOrgId(xOrgId).CascadeManager(cascadeManager).Execute()

Delete a system user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string |  (optional)
    cascadeManager := "cascadeManager_example" // string | This is an optional flag that can be enabled on the DELETE call, DELETE /systemusers/{id}?cascade_manager=null. This parameter will clear the Manager attribute on all direct reports and then delete the account. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SystemusersDelete(context.Background(), id).XOrgId(xOrgId).CascadeManager(cascadeManager).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemusersDelete`: Systemuserreturn
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SystemusersDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** |  | 
 **cascadeManager** | **string** | This is an optional flag that can be enabled on the DELETE call, DELETE /systemusers/{id}?cascade_manager&#x3D;null. This parameter will clear the Manager attribute on all direct reports and then delete the account. | 

### Return type

[**Systemuserreturn**](Systemuserreturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemusersExpire

> string SystemusersExpire(ctx, id).XOrgId(xOrgId).Execute()

Expire a system user's password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SystemusersExpire(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersExpire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemusersExpire`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SystemusersExpire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersExpireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** |  | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemusersGet

> Systemuserreturn SystemusersGet(ctx, id).Fields(fields).Filter(filter).XOrgId(xOrgId).Execute()

List a system user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 
    fields := "fields_example" // string | Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  (optional)
    filter := "filter_example" // string | A filter to apply to the query. See the supported operators below. For more complex searches, see the related `/search/<domain>` endpoints, e.g. `/search/systems`.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** = Supported operators are: - `$eq` (equals) - `$ne` (does not equal) - `$gt` (is greater than) - `$gte` (is greater than or equal to) - `$lt` (is less than) - `$lte` (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the `$` will result in undefined behavior._  **value** = Populate with the value you want to search for. Is case sensitive.  **Examples** - `GET /users?filter=username:$eq:testuser` - `GET /systemusers?filter=password_expiration_date:$lte:2021-10-24` - `GET /systemusers?filter=department:$ne:Accounting` - `GET /systems?filter[0]=firstname:$eq:foo&filter[1]=lastname:$eq:bar` - this will    AND the filters together. - `GET /systems?filter[or][0]=lastname:$eq:foo&filter[or][1]=lastname:$eq:bar` - this will    OR the filters together. (optional)
    xOrgId := "xOrgId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SystemusersGet(context.Background(), id).Fields(fields).Filter(filter).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemusersGet`: Systemuserreturn
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SystemusersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Use a space seperated string of field parameters to include the data in the response. If omitted, the default list of fields will be returned.  | 
 **filter** | **string** | A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 
 **xOrgId** | **string** |  | 

### Return type

[**Systemuserreturn**](Systemuserreturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemusersList

> Systemuserslist SystemusersList(ctx).Limit(limit).Skip(skip).Sort(sort).Fields(fields).Filter(filter).XOrgId(xOrgId).Search(search).Execute()

List all system users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    limit := int32(56) // int32 | The number of records to return at once. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := "sort_example" // string | The space separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional)
    fields := "fields_example" // string | The space separated fields included in the returned records. If omitted the default list of fields will be returned.  (optional)
    filter := "filter_example" // string | A filter to apply to the query. See the supported operators below. For more complex searches, see the related `/search/<domain>` endpoints, e.g. `/search/systems`.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** = Supported operators are: - `$eq` (equals) - `$ne` (does not equal) - `$gt` (is greater than) - `$gte` (is greater than or equal to) - `$lt` (is less than) - `$lte` (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the `$` will result in undefined behavior._  **value** = Populate with the value you want to search for. Is case sensitive.  **Examples** - `GET /users?filter=username:$eq:testuser` - `GET /systemusers?filter=password_expiration_date:$lte:2021-10-24` - `GET /systemusers?filter=department:$ne:Accounting` - `GET /systems?filter[0]=firstname:$eq:foo&filter[1]=lastname:$eq:bar` - this will    AND the filters together. - `GET /systems?filter[or][0]=lastname:$eq:foo&filter[or][1]=lastname:$eq:bar` - this will    OR the filters together. (optional)
    xOrgId := "xOrgId_example" // string |  (optional)
    search := "search_example" // string | A nested object containing a `searchTerm` string or array of strings and a list of `fields` to search on. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SystemusersList(context.Background()).Limit(limit).Skip(skip).Sort(sort).Fields(fields).Filter(filter).XOrgId(xOrgId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemusersList`: Systemuserslist
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SystemusersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The number of records to return at once. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **string** | The space separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | 
 **fields** | **string** | The space separated fields included in the returned records. If omitted the default list of fields will be returned.  | 
 **filter** | **string** | A filter to apply to the query. See the supported operators below. For more complex searches, see the related &#x60;/search/&lt;domain&gt;&#x60; endpoints, e.g. &#x60;/search/systems&#x60;.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D; Supported operators are: - &#x60;$eq&#x60; (equals) - &#x60;$ne&#x60; (does not equal) - &#x60;$gt&#x60; (is greater than) - &#x60;$gte&#x60; (is greater than or equal to) - &#x60;$lt&#x60; (is less than) - &#x60;$lte&#x60; (is less than or equal to)  _Note: v1 operators differ from v2 operators._  _Note: For v1 operators, excluding the &#x60;$&#x60; will result in undefined behavior._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive.  **Examples** - &#x60;GET /users?filter&#x3D;username:$eq:testuser&#x60; - &#x60;GET /systemusers?filter&#x3D;password_expiration_date:$lte:2021-10-24&#x60; - &#x60;GET /systemusers?filter&#x3D;department:$ne:Accounting&#x60; - &#x60;GET /systems?filter[0]&#x3D;firstname:$eq:foo&amp;filter[1]&#x3D;lastname:$eq:bar&#x60; - this will    AND the filters together. - &#x60;GET /systems?filter[or][0]&#x3D;lastname:$eq:foo&amp;filter[or][1]&#x3D;lastname:$eq:bar&#x60; - this will    OR the filters together. | 
 **xOrgId** | **string** |  | 
 **search** | **string** | A nested object containing a &#x60;searchTerm&#x60; string or array of strings and a list of &#x60;fields&#x60; to search on. | 

### Return type

[**Systemuserslist**](Systemuserslist.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemusersMfasync

> SystemusersMfasync(ctx, id).Execute()

Sync a systemuser's mfa enrollment status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SystemusersApi.SystemusersMfasync(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersMfasync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersMfasyncRequest struct via the builder pattern


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


## SystemusersPost

> Systemuserreturn SystemusersPost(ctx).XOrgId(xOrgId).FullValidationDetails(fullValidationDetails).Body(body).Execute()

Create a system user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    xOrgId := "xOrgId_example" // string |  (optional)
    fullValidationDetails := "fullValidationDetails_example" // string | Pass this query parameter when a client wants all validation errors to be returned with a detailed error response for the form field specified. The current form fields are allowed:  * `password`  #### Password validation flag Use the `password` validation flag to receive details on a possible bad request response ``` ?fullValidationDetails=password ``` Without the flag, default behavior will be a normal 400 with only a single validation string error #### Expected Behavior Clients can expect a list of validation error mappings for the validation query field in the details provided on the response: ``` {   \"code\": 400,   \"message\": \"Password validation fail\",   \"status\": \"INVALID_ARGUMENT\",   \"details\": [       {         \"fieldViolationsList\": [           {\"field\": \"password\", \"description\": \"specialCharacter\"}         ],         '@type': 'type.googleapis.com/google.rpc.BadRequest',       },   ], }, ``` (optional)
    body := *openapiclient.NewSystemuserputpost("Email_example", "Username_example") // Systemuserputpost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SystemusersPost(context.Background()).XOrgId(xOrgId).FullValidationDetails(fullValidationDetails).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemusersPost`: Systemuserreturn
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SystemusersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** |  | 
 **fullValidationDetails** | **string** | Pass this query parameter when a client wants all validation errors to be returned with a detailed error response for the form field specified. The current form fields are allowed:  * &#x60;password&#x60;  #### Password validation flag Use the &#x60;password&#x60; validation flag to receive details on a possible bad request response &#x60;&#x60;&#x60; ?fullValidationDetails&#x3D;password &#x60;&#x60;&#x60; Without the flag, default behavior will be a normal 400 with only a single validation string error #### Expected Behavior Clients can expect a list of validation error mappings for the validation query field in the details provided on the response: &#x60;&#x60;&#x60; {   \&quot;code\&quot;: 400,   \&quot;message\&quot;: \&quot;Password validation fail\&quot;,   \&quot;status\&quot;: \&quot;INVALID_ARGUMENT\&quot;,   \&quot;details\&quot;: [       {         \&quot;fieldViolationsList\&quot;: [           {\&quot;field\&quot;: \&quot;password\&quot;, \&quot;description\&quot;: \&quot;specialCharacter\&quot;}         ],         &#39;@type&#39;: &#39;type.googleapis.com/google.rpc.BadRequest&#39;,       },   ], }, &#x60;&#x60;&#x60; | 
 **body** | [**Systemuserputpost**](Systemuserputpost.md) |  | 

### Return type

[**Systemuserreturn**](Systemuserreturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemusersPut

> Systemuserreturn SystemusersPut(ctx, id).XOrgId(xOrgId).FullValidationDetails(fullValidationDetails).Body(body).Execute()

Update a system user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string |  (optional)
    fullValidationDetails := "fullValidationDetails_example" // string | This endpoint can take in a query when a client wants all validation errors to be returned with error response for the form field specified, i.e. 'password' #### Password validation flag Use the \"password\" validation flag to receive details on a possible bad request response Without the `password` flag, default behavior will be a normal 400 with only a validation string message ``` ?fullValidationDetails=password ``` #### Expected Behavior Clients can expect a list of validation error mappings for the validation query field in the details provided on the response: ``` {   \"code\": 400,   \"message\": \"Password validation fail\",   \"status\": \"INVALID_ARGUMENT\",   \"details\": [       {         \"fieldViolationsList\": [{ \"field\": \"password\", \"description\": \"passwordHistory\" }],         '@type': 'type.googleapis.com/google.rpc.BadRequest',       },   ], }, ``` (optional)
    body := *openapiclient.NewSystemuserput() // Systemuserput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SystemusersPut(context.Background(), id).XOrgId(xOrgId).FullValidationDetails(fullValidationDetails).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemusersPut`: Systemuserreturn
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SystemusersPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** |  | 
 **fullValidationDetails** | **string** | This endpoint can take in a query when a client wants all validation errors to be returned with error response for the form field specified, i.e. &#39;password&#39; #### Password validation flag Use the \&quot;password\&quot; validation flag to receive details on a possible bad request response Without the &#x60;password&#x60; flag, default behavior will be a normal 400 with only a validation string message &#x60;&#x60;&#x60; ?fullValidationDetails&#x3D;password &#x60;&#x60;&#x60; #### Expected Behavior Clients can expect a list of validation error mappings for the validation query field in the details provided on the response: &#x60;&#x60;&#x60; {   \&quot;code\&quot;: 400,   \&quot;message\&quot;: \&quot;Password validation fail\&quot;,   \&quot;status\&quot;: \&quot;INVALID_ARGUMENT\&quot;,   \&quot;details\&quot;: [       {         \&quot;fieldViolationsList\&quot;: [{ \&quot;field\&quot;: \&quot;password\&quot;, \&quot;description\&quot;: \&quot;passwordHistory\&quot; }],         &#39;@type&#39;: &#39;type.googleapis.com/google.rpc.BadRequest&#39;,       },   ], }, &#x60;&#x60;&#x60; | 
 **body** | [**Systemuserput**](Systemuserput.md) |  | 

### Return type

[**Systemuserreturn**](Systemuserreturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemusersResetmfa

> string SystemusersResetmfa(ctx, id).XOrgId(xOrgId).Body(body).Execute()

Reset a system user's MFA token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string |  (optional)
    body := *openapiclient.NewSystemusersResetmfaRequest() // SystemusersResetmfaRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SystemusersResetmfa(context.Background(), id).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersResetmfa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemusersResetmfa`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SystemusersResetmfa`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersResetmfaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** |  | 
 **body** | [**SystemusersResetmfaRequest**](SystemusersResetmfaRequest.md) |  | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemusersStateActivate

> string SystemusersStateActivate(ctx, id).Body(body).Execute()

Activate System User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 
    body := *openapiclient.NewSystemusersStateActivateRequest() // SystemusersStateActivateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SystemusersStateActivate(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersStateActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemusersStateActivate`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SystemusersStateActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersStateActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SystemusersStateActivateRequest**](SystemusersStateActivateRequest.md) |  | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemusersUnlock

> string SystemusersUnlock(ctx, id).XOrgId(xOrgId).Execute()

Unlock a system user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemusersApi.SystemusersUnlock(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemusersApi.SystemusersUnlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemusersUnlock`: string
    fmt.Fprintf(os.Stdout, "Response from `SystemusersApi.SystemusersUnlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemusersUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** |  | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


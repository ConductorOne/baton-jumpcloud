# \GoogleEMMApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesGetDevice**](GoogleEMMApi.md#DevicesGetDevice) | **Get** /google-emm/devices/{deviceId} | Get device
[**DevicesGetDeviceAndroidPolicy**](GoogleEMMApi.md#DevicesGetDeviceAndroidPolicy) | **Get** /google-emm/devices/{deviceId}/policy_results | Get the policy JSON of a device
[**DevicesLockDevice**](GoogleEMMApi.md#DevicesLockDevice) | **Post** /google-emm/devices/{deviceId}/lock | Lock device
[**DevicesRelinquishOwnership**](GoogleEMMApi.md#DevicesRelinquishOwnership) | **Post** /google-emm/devices/{deviceId}/relinquishownership | Relinquish Ownership of COPE device
[**DevicesResetPassword**](GoogleEMMApi.md#DevicesResetPassword) | **Post** /google-emm/devices/{deviceId}/resetpassword | Reset Password of a device
[**EnterprisesCreateEnterprise**](GoogleEMMApi.md#EnterprisesCreateEnterprise) | **Post** /google-emm/enterprises | Create a Google Enterprise
[**EnterprisesDeleteEnterprise**](GoogleEMMApi.md#EnterprisesDeleteEnterprise) | **Delete** /google-emm/enterprises/{enterpriseId} | Delete a Google Enterprise
[**EnterprisesGetConnectionStatus**](GoogleEMMApi.md#EnterprisesGetConnectionStatus) | **Get** /google-emm/enterprises/{enterpriseId}/connection-status | Test connection with Google
[**EnterprisesListEnterprises**](GoogleEMMApi.md#EnterprisesListEnterprises) | **Get** /google-emm/enterprises | List Google Enterprises
[**EnterprisesPatchEnterprise**](GoogleEMMApi.md#EnterprisesPatchEnterprise) | **Patch** /google-emm/enterprises/{enterpriseId} | Update a Google Enterprise
[**SignupURLsCreate**](GoogleEMMApi.md#SignupURLsCreate) | **Post** /google-emm/signup-urls | Get a Signup URL to enroll Google enterprise
[**WebTokensCreateWebToken**](GoogleEMMApi.md#WebTokensCreateWebToken) | **Post** /google-emm/web-tokens | Get a web token to render Google Play iFrame



## DevicesGetDevice

> JumpcloudGoogleEmmDevice DevicesGetDevice(ctx, deviceId).Execute()

Get device



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
    deviceId := string(BYTE_ARRAY_DATA_HERE) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.DevicesGetDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.DevicesGetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesGetDevice`: JumpcloudGoogleEmmDevice
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.DevicesGetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JumpcloudGoogleEmmDevice**](JumpcloudGoogleEmmDevice.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesGetDeviceAndroidPolicy

> JumpcloudGoogleEmmDeviceAndroidPolicy DevicesGetDeviceAndroidPolicy(ctx, deviceId).Execute()

Get the policy JSON of a device



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
    deviceId := string(BYTE_ARRAY_DATA_HERE) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.DevicesGetDeviceAndroidPolicy(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.DevicesGetDeviceAndroidPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesGetDeviceAndroidPolicy`: JumpcloudGoogleEmmDeviceAndroidPolicy
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.DevicesGetDeviceAndroidPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetDeviceAndroidPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JumpcloudGoogleEmmDeviceAndroidPolicy**](JumpcloudGoogleEmmDeviceAndroidPolicy.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesLockDevice

> map[string]interface{} DevicesLockDevice(ctx, deviceId).Body(body).Execute()

Lock device



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
    deviceId := string(BYTE_ARRAY_DATA_HERE) // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.DevicesLockDevice(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.DevicesLockDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesLockDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.DevicesLockDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesLockDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesRelinquishOwnership

> map[string]interface{} DevicesRelinquishOwnership(ctx, deviceId).Body(body).Execute()

Relinquish Ownership of COPE device



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
    deviceId := string(BYTE_ARRAY_DATA_HERE) // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.DevicesRelinquishOwnership(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.DevicesRelinquishOwnership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesRelinquishOwnership`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.DevicesRelinquishOwnership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesRelinquishOwnershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesResetPassword

> map[string]interface{} DevicesResetPassword(ctx, deviceId).Body(body).Execute()

Reset Password of a device



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
    deviceId := string(BYTE_ARRAY_DATA_HERE) // string | 
    body := *openapiclient.NewDevicesResetPasswordRequest() // DevicesResetPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.DevicesResetPassword(context.Background(), deviceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.DevicesResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesResetPassword`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.DevicesResetPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DevicesResetPasswordRequest**](DevicesResetPasswordRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterprisesCreateEnterprise

> JumpcloudGoogleEmmEnterprise EnterprisesCreateEnterprise(ctx).Body(body).Execute()

Create a Google Enterprise



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
    body := *openapiclient.NewJumpcloudGoogleEmmCreateEnterpriseRequest() // JumpcloudGoogleEmmCreateEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.EnterprisesCreateEnterprise(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.EnterprisesCreateEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterprisesCreateEnterprise`: JumpcloudGoogleEmmEnterprise
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.EnterprisesCreateEnterprise`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnterprisesCreateEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JumpcloudGoogleEmmCreateEnterpriseRequest**](JumpcloudGoogleEmmCreateEnterpriseRequest.md) |  | 

### Return type

[**JumpcloudGoogleEmmEnterprise**](JumpcloudGoogleEmmEnterprise.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterprisesDeleteEnterprise

> map[string]interface{} EnterprisesDeleteEnterprise(ctx, enterpriseId).Execute()

Delete a Google Enterprise



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
    enterpriseId := string(BYTE_ARRAY_DATA_HERE) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.EnterprisesDeleteEnterprise(context.Background(), enterpriseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.EnterprisesDeleteEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterprisesDeleteEnterprise`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.EnterprisesDeleteEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterpriseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterprisesDeleteEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterprisesGetConnectionStatus

> JumpcloudGoogleEmmConnectionStatus EnterprisesGetConnectionStatus(ctx, enterpriseId).Execute()

Test connection with Google



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
    enterpriseId := string(BYTE_ARRAY_DATA_HERE) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.EnterprisesGetConnectionStatus(context.Background(), enterpriseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.EnterprisesGetConnectionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterprisesGetConnectionStatus`: JumpcloudGoogleEmmConnectionStatus
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.EnterprisesGetConnectionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterpriseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterprisesGetConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JumpcloudGoogleEmmConnectionStatus**](JumpcloudGoogleEmmConnectionStatus.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterprisesListEnterprises

> JumpcloudGoogleEmmListEnterprisesResponse EnterprisesListEnterprises(ctx).QueryParametersLimit(queryParametersLimit).QueryParametersSkip(queryParametersSkip).QueryParametersFields(queryParametersFields).Execute()

List Google Enterprises



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
    queryParametersLimit := "queryParametersLimit_example" // string |  (optional)
    queryParametersSkip := "queryParametersSkip_example" // string |  (optional)
    queryParametersFields := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.EnterprisesListEnterprises(context.Background()).QueryParametersLimit(queryParametersLimit).QueryParametersSkip(queryParametersSkip).QueryParametersFields(queryParametersFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.EnterprisesListEnterprises``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterprisesListEnterprises`: JumpcloudGoogleEmmListEnterprisesResponse
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.EnterprisesListEnterprises`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnterprisesListEnterprisesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryParametersLimit** | **string** |  | 
 **queryParametersSkip** | **string** |  | 
 **queryParametersFields** | **[]string** |  | 

### Return type

[**JumpcloudGoogleEmmListEnterprisesResponse**](JumpcloudGoogleEmmListEnterprisesResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterprisesPatchEnterprise

> JumpcloudGoogleEmmEnterprise EnterprisesPatchEnterprise(ctx, enterpriseId).Body(body).Execute()

Update a Google Enterprise



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
    enterpriseId := string(BYTE_ARRAY_DATA_HERE) // string | 
    body := *openapiclient.NewEnterprisesPatchEnterpriseRequest() // EnterprisesPatchEnterpriseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.EnterprisesPatchEnterprise(context.Background(), enterpriseId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.EnterprisesPatchEnterprise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterprisesPatchEnterprise`: JumpcloudGoogleEmmEnterprise
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.EnterprisesPatchEnterprise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enterpriseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnterprisesPatchEnterpriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**EnterprisesPatchEnterpriseRequest**](EnterprisesPatchEnterpriseRequest.md) |  | 

### Return type

[**JumpcloudGoogleEmmEnterprise**](JumpcloudGoogleEmmEnterprise.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignupURLsCreate

> JumpcloudGoogleEmmSignupURL SignupURLsCreate(ctx).Execute()

Get a Signup URL to enroll Google enterprise



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.SignupURLsCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.SignupURLsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignupURLsCreate`: JumpcloudGoogleEmmSignupURL
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.SignupURLsCreate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSignupURLsCreateRequest struct via the builder pattern


### Return type

[**JumpcloudGoogleEmmSignupURL**](JumpcloudGoogleEmmSignupURL.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebTokensCreateWebToken

> JumpcloudGoogleEmmWebToken WebTokensCreateWebToken(ctx).Body(body).Execute()

Get a web token to render Google Play iFrame



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
    body := *openapiclient.NewJumpcloudGoogleEmmCreateWebTokenRequest() // JumpcloudGoogleEmmCreateWebTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleEMMApi.WebTokensCreateWebToken(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleEMMApi.WebTokensCreateWebToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebTokensCreateWebToken`: JumpcloudGoogleEmmWebToken
    fmt.Fprintf(os.Stdout, "Response from `GoogleEMMApi.WebTokensCreateWebToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebTokensCreateWebTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**JumpcloudGoogleEmmCreateWebTokenRequest**](JumpcloudGoogleEmmCreateWebTokenRequest.md) |  | 

### Return type

[**JumpcloudGoogleEmmWebToken**](JumpcloudGoogleEmmWebToken.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AppleMDMApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplemdmsCsrget**](AppleMDMApi.md#ApplemdmsCsrget) | **Get** /applemdms/{apple_mdm_id}/csr | Get Apple MDM CSR Plist
[**ApplemdmsDelete**](AppleMDMApi.md#ApplemdmsDelete) | **Delete** /applemdms/{id} | Delete an Apple MDM
[**ApplemdmsDeletedevice**](AppleMDMApi.md#ApplemdmsDeletedevice) | **Delete** /applemdms/{apple_mdm_id}/devices/{device_id} | Remove an Apple MDM Device&#39;s Enrollment
[**ApplemdmsDepkeyget**](AppleMDMApi.md#ApplemdmsDepkeyget) | **Get** /applemdms/{apple_mdm_id}/depkey | Get Apple MDM DEP Public Key
[**ApplemdmsDevicesClearActivationLock**](AppleMDMApi.md#ApplemdmsDevicesClearActivationLock) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/clearActivationLock | Clears the Activation Lock for a Device
[**ApplemdmsDevicesOSUpdateStatus**](AppleMDMApi.md#ApplemdmsDevicesOSUpdateStatus) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/osUpdateStatus | Request the status of an OS update for a device
[**ApplemdmsDevicesRefreshActivationLockInformation**](AppleMDMApi.md#ApplemdmsDevicesRefreshActivationLockInformation) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/refreshActivationLockInformation | Refresh activation lock information for a device
[**ApplemdmsDevicesScheduleOSUpdate**](AppleMDMApi.md#ApplemdmsDevicesScheduleOSUpdate) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/scheduleOSUpdate | Schedule an OS update for a device
[**ApplemdmsDeviceserase**](AppleMDMApi.md#ApplemdmsDeviceserase) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/erase | Erase Device
[**ApplemdmsDeviceslist**](AppleMDMApi.md#ApplemdmsDeviceslist) | **Get** /applemdms/{apple_mdm_id}/devices | List AppleMDM Devices
[**ApplemdmsDeviceslock**](AppleMDMApi.md#ApplemdmsDeviceslock) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/lock | Lock Device
[**ApplemdmsDevicesrestart**](AppleMDMApi.md#ApplemdmsDevicesrestart) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/restart | Restart Device
[**ApplemdmsDevicesshutdown**](AppleMDMApi.md#ApplemdmsDevicesshutdown) | **Post** /applemdms/{apple_mdm_id}/devices/{device_id}/shutdown | Shut Down Device
[**ApplemdmsEnrollmentprofilesget**](AppleMDMApi.md#ApplemdmsEnrollmentprofilesget) | **Get** /applemdms/{apple_mdm_id}/enrollmentprofiles/{id} | Get an Apple MDM Enrollment Profile
[**ApplemdmsEnrollmentprofileslist**](AppleMDMApi.md#ApplemdmsEnrollmentprofileslist) | **Get** /applemdms/{apple_mdm_id}/enrollmentprofiles | List Apple MDM Enrollment Profiles
[**ApplemdmsGetdevice**](AppleMDMApi.md#ApplemdmsGetdevice) | **Get** /applemdms/{apple_mdm_id}/devices/{device_id} | Details of an AppleMDM Device
[**ApplemdmsList**](AppleMDMApi.md#ApplemdmsList) | **Get** /applemdms | List Apple MDMs
[**ApplemdmsPut**](AppleMDMApi.md#ApplemdmsPut) | **Put** /applemdms/{id} | Update an Apple MDM
[**ApplemdmsRefreshdepdevices**](AppleMDMApi.md#ApplemdmsRefreshdepdevices) | **Post** /applemdms/{apple_mdm_id}/refreshdepdevices | Refresh DEP Devices



## ApplemdmsCsrget

> string ApplemdmsCsrget(ctx, appleMdmId).XOrgId(xOrgId).Execute()

Get Apple MDM CSR Plist



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
    appleMdmId := "appleMdmId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsCsrget(context.Background(), appleMdmId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsCsrget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsCsrget`: string
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsCsrget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsCsrgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsDelete

> AppleMDM ApplemdmsDelete(ctx, id).XOrgId(xOrgId).Execute()

Delete an Apple MDM



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
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsDelete(context.Background(), id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsDelete`: AppleMDM
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**AppleMDM**](AppleMDM.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsDeletedevice

> AppleMdmDevice ApplemdmsDeletedevice(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Execute()

Remove an Apple MDM Device's Enrollment



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsDeletedevice(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDeletedevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsDeletedevice`: AppleMdmDevice
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsDeletedevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDeletedeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**AppleMdmDevice**](AppleMdmDevice.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsDepkeyget

> string ApplemdmsDepkeyget(ctx, appleMdmId).XOrgId(xOrgId).Execute()

Get Apple MDM DEP Public Key



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
    appleMdmId := "appleMdmId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsDepkeyget(context.Background(), appleMdmId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDepkeyget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsDepkeyget`: string
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsDepkeyget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDepkeygetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-pem-file

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsDevicesClearActivationLock

> ApplemdmsDevicesClearActivationLock(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Execute()

Clears the Activation Lock for a Device



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppleMDMApi.ApplemdmsDevicesClearActivationLock(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDevicesClearActivationLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDevicesClearActivationLockRequest struct via the builder pattern


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


## ApplemdmsDevicesOSUpdateStatus

> ApplemdmsDevicesOSUpdateStatus(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Execute()

Request the status of an OS update for a device



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppleMDMApi.ApplemdmsDevicesOSUpdateStatus(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDevicesOSUpdateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDevicesOSUpdateStatusRequest struct via the builder pattern


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


## ApplemdmsDevicesRefreshActivationLockInformation

> ApplemdmsDevicesRefreshActivationLockInformation(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Execute()

Refresh activation lock information for a device



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppleMDMApi.ApplemdmsDevicesRefreshActivationLockInformation(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDevicesRefreshActivationLockInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDevicesRefreshActivationLockInformationRequest struct via the builder pattern


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


## ApplemdmsDevicesScheduleOSUpdate

> ApplemdmsDevicesScheduleOSUpdate(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Body(body).Execute()

Schedule an OS update for a device



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewScheduleOSUpdate(openapiclient.InstallActionType("DOWNLOAD_ONLY"), "ProductKey_example") // ScheduleOSUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppleMDMApi.ApplemdmsDevicesScheduleOSUpdate(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDevicesScheduleOSUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDevicesScheduleOSUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**ScheduleOSUpdate**](ScheduleOSUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsDeviceserase

> ApplemdmsDeviceserase(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Body(body).Execute()

Erase Device



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewApplemdmsDeviceseraseRequest() // ApplemdmsDeviceseraseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppleMDMApi.ApplemdmsDeviceserase(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDeviceserase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDeviceseraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**ApplemdmsDeviceseraseRequest**](ApplemdmsDeviceseraseRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsDeviceslist

> []AppleMdmDevice ApplemdmsDeviceslist(ctx, appleMdmId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Sort(sort).XTotalCount(xTotalCount).Execute()

List AppleMDM Devices



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
    appleMdmId := "appleMdmId_example" // string | 
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    xTotalCount := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsDeviceslist(context.Background(), appleMdmId).Limit(limit).XOrgId(xOrgId).Skip(skip).Filter(filter).Sort(sort).XTotalCount(xTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDeviceslist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsDeviceslist`: []AppleMdmDevice
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsDeviceslist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDeviceslistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **xTotalCount** | **int32** |  | 

### Return type

[**[]AppleMdmDevice**](AppleMdmDevice.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsDeviceslock

> ApplemdmsDeviceslock(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Body(body).Execute()

Lock Device



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewApplemdmsDeviceslockRequest() // ApplemdmsDeviceslockRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppleMDMApi.ApplemdmsDeviceslock(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDeviceslock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDeviceslockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**ApplemdmsDeviceslockRequest**](ApplemdmsDeviceslockRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsDevicesrestart

> ApplemdmsDevicesrestart(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Body(body).Execute()

Restart Device



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewApplemdmsDevicesrestartRequest() // ApplemdmsDevicesrestartRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppleMDMApi.ApplemdmsDevicesrestart(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDevicesrestart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDevicesrestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**ApplemdmsDevicesrestartRequest**](ApplemdmsDevicesrestartRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsDevicesshutdown

> ApplemdmsDevicesshutdown(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Execute()

Shut Down Device



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppleMDMApi.ApplemdmsDevicesshutdown(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsDevicesshutdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsDevicesshutdownRequest struct via the builder pattern


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


## ApplemdmsEnrollmentprofilesget

> string ApplemdmsEnrollmentprofilesget(ctx, appleMdmId, id).XOrgId(xOrgId).Execute()

Get an Apple MDM Enrollment Profile



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
    appleMdmId := "appleMdmId_example" // string | 
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsEnrollmentprofilesget(context.Background(), appleMdmId, id).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsEnrollmentprofilesget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsEnrollmentprofilesget`: string
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsEnrollmentprofilesget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsEnrollmentprofilesgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

**string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-apple-aspen-config

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsEnrollmentprofileslist

> []AppleMDM ApplemdmsEnrollmentprofileslist(ctx, appleMdmId).XOrgId(xOrgId).Execute()

List Apple MDM Enrollment Profiles



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
    appleMdmId := "appleMdmId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsEnrollmentprofileslist(context.Background(), appleMdmId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsEnrollmentprofileslist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsEnrollmentprofileslist`: []AppleMDM
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsEnrollmentprofileslist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsEnrollmentprofileslistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**[]AppleMDM**](AppleMDM.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsGetdevice

> AppleMdmDevice ApplemdmsGetdevice(ctx, appleMdmId, deviceId).XOrgId(xOrgId).Execute()

Details of an AppleMDM Device



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
    appleMdmId := "appleMdmId_example" // string | 
    deviceId := "deviceId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsGetdevice(context.Background(), appleMdmId, deviceId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsGetdevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsGetdevice`: AppleMdmDevice
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsGetdevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsGetdeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 

### Return type

[**AppleMdmDevice**](AppleMdmDevice.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsList

> []AppleMDM ApplemdmsList(ctx).XOrgId(xOrgId).Limit(limit).Skip(skip).Execute()

List Apple MDMs



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
    limit := int32(56) // int32 |  (optional) (default to 1)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsList(context.Background()).XOrgId(xOrgId).Limit(limit).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsList`: []AppleMDM
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **limit** | **int32** |  | [default to 1]
 **skip** | **int32** | The offset into the records to return. | [default to 0]

### Return type

[**[]AppleMDM**](AppleMDM.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsPut

> AppleMDM ApplemdmsPut(ctx, id).XOrgId(xOrgId).Body(body).Execute()

Update an Apple MDM



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
    id := "id_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)
    body := *openapiclient.NewAppleMdmPatch() // AppleMdmPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleMDMApi.ApplemdmsPut(context.Background(), id).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplemdmsPut`: AppleMDM
    fmt.Fprintf(os.Stdout, "Response from `AppleMDMApi.ApplemdmsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** | Organization identifier that can be obtained from console settings. | 
 **body** | [**AppleMdmPatch**](AppleMdmPatch.md) |  | 

### Return type

[**AppleMDM**](AppleMDM.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplemdmsRefreshdepdevices

> ApplemdmsRefreshdepdevices(ctx, appleMdmId).XOrgId(xOrgId).Execute()

Refresh DEP Devices



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
    appleMdmId := "appleMdmId_example" // string | 
    xOrgId := "xOrgId_example" // string | Organization identifier that can be obtained from console settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppleMDMApi.ApplemdmsRefreshdepdevices(context.Background(), appleMdmId).XOrgId(xOrgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleMDMApi.ApplemdmsRefreshdepdevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appleMdmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplemdmsRefreshdepdevicesRequest struct via the builder pattern


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


# SystemInsightsServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ModulePath** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **int32** |  | [optional] 
**ServiceExitCode** | Pointer to **int32** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**StartType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SystemId** | Pointer to **string** |  | [optional] 
**UserAccount** | Pointer to **string** |  | [optional] 
**Win32ExitCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewSystemInsightsServices

`func NewSystemInsightsServices() *SystemInsightsServices`

NewSystemInsightsServices instantiates a new SystemInsightsServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInsightsServicesWithDefaults

`func NewSystemInsightsServicesWithDefaults() *SystemInsightsServices`

NewSystemInsightsServicesWithDefaults instantiates a new SystemInsightsServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SystemInsightsServices) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemInsightsServices) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemInsightsServices) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemInsightsServices) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *SystemInsightsServices) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SystemInsightsServices) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SystemInsightsServices) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SystemInsightsServices) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetModulePath

`func (o *SystemInsightsServices) GetModulePath() string`

GetModulePath returns the ModulePath field if non-nil, zero value otherwise.

### GetModulePathOk

`func (o *SystemInsightsServices) GetModulePathOk() (*string, bool)`

GetModulePathOk returns a tuple with the ModulePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulePath

`func (o *SystemInsightsServices) SetModulePath(v string)`

SetModulePath sets ModulePath field to given value.

### HasModulePath

`func (o *SystemInsightsServices) HasModulePath() bool`

HasModulePath returns a boolean if a field has been set.

### GetName

`func (o *SystemInsightsServices) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemInsightsServices) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemInsightsServices) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemInsightsServices) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *SystemInsightsServices) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SystemInsightsServices) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SystemInsightsServices) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SystemInsightsServices) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPid

`func (o *SystemInsightsServices) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *SystemInsightsServices) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *SystemInsightsServices) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *SystemInsightsServices) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetServiceExitCode

`func (o *SystemInsightsServices) GetServiceExitCode() int32`

GetServiceExitCode returns the ServiceExitCode field if non-nil, zero value otherwise.

### GetServiceExitCodeOk

`func (o *SystemInsightsServices) GetServiceExitCodeOk() (*int32, bool)`

GetServiceExitCodeOk returns a tuple with the ServiceExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceExitCode

`func (o *SystemInsightsServices) SetServiceExitCode(v int32)`

SetServiceExitCode sets ServiceExitCode field to given value.

### HasServiceExitCode

`func (o *SystemInsightsServices) HasServiceExitCode() bool`

HasServiceExitCode returns a boolean if a field has been set.

### GetServiceType

`func (o *SystemInsightsServices) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *SystemInsightsServices) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *SystemInsightsServices) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *SystemInsightsServices) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStartType

`func (o *SystemInsightsServices) GetStartType() string`

GetStartType returns the StartType field if non-nil, zero value otherwise.

### GetStartTypeOk

`func (o *SystemInsightsServices) GetStartTypeOk() (*string, bool)`

GetStartTypeOk returns a tuple with the StartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartType

`func (o *SystemInsightsServices) SetStartType(v string)`

SetStartType sets StartType field to given value.

### HasStartType

`func (o *SystemInsightsServices) HasStartType() bool`

HasStartType returns a boolean if a field has been set.

### GetStatus

`func (o *SystemInsightsServices) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemInsightsServices) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemInsightsServices) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SystemInsightsServices) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemId

`func (o *SystemInsightsServices) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *SystemInsightsServices) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *SystemInsightsServices) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *SystemInsightsServices) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetUserAccount

`func (o *SystemInsightsServices) GetUserAccount() string`

GetUserAccount returns the UserAccount field if non-nil, zero value otherwise.

### GetUserAccountOk

`func (o *SystemInsightsServices) GetUserAccountOk() (*string, bool)`

GetUserAccountOk returns a tuple with the UserAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccount

`func (o *SystemInsightsServices) SetUserAccount(v string)`

SetUserAccount sets UserAccount field to given value.

### HasUserAccount

`func (o *SystemInsightsServices) HasUserAccount() bool`

HasUserAccount returns a boolean if a field has been set.

### GetWin32ExitCode

`func (o *SystemInsightsServices) GetWin32ExitCode() int32`

GetWin32ExitCode returns the Win32ExitCode field if non-nil, zero value otherwise.

### GetWin32ExitCodeOk

`func (o *SystemInsightsServices) GetWin32ExitCodeOk() (*int32, bool)`

GetWin32ExitCodeOk returns a tuple with the Win32ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin32ExitCode

`func (o *SystemInsightsServices) SetWin32ExitCode(v int32)`

SetWin32ExitCode sets Win32ExitCode field to given value.

### HasWin32ExitCode

`func (o *SystemInsightsServices) HasWin32ExitCode() bool`

HasWin32ExitCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



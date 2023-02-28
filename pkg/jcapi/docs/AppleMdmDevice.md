# AppleMdmDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DepRegistered** | Pointer to **bool** |  | [optional] 
**DeviceInformation** | Pointer to [**AppleMdmDeviceInfo**](AppleMdmDeviceInfo.md) |  | [optional] 
**Enrolled** | Pointer to **bool** |  | [optional] 
**HasActivationLockBypassCodes** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**SecurityInfo** | Pointer to [**AppleMdmDeviceSecurityInfo**](AppleMdmDeviceSecurityInfo.md) |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Udid** | Pointer to **string** |  | [optional] 

## Methods

### NewAppleMdmDevice

`func NewAppleMdmDevice() *AppleMdmDevice`

NewAppleMdmDevice instantiates a new AppleMdmDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleMdmDeviceWithDefaults

`func NewAppleMdmDeviceWithDefaults() *AppleMdmDevice`

NewAppleMdmDeviceWithDefaults instantiates a new AppleMdmDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AppleMdmDevice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppleMdmDevice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppleMdmDevice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AppleMdmDevice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepRegistered

`func (o *AppleMdmDevice) GetDepRegistered() bool`

GetDepRegistered returns the DepRegistered field if non-nil, zero value otherwise.

### GetDepRegisteredOk

`func (o *AppleMdmDevice) GetDepRegisteredOk() (*bool, bool)`

GetDepRegisteredOk returns a tuple with the DepRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepRegistered

`func (o *AppleMdmDevice) SetDepRegistered(v bool)`

SetDepRegistered sets DepRegistered field to given value.

### HasDepRegistered

`func (o *AppleMdmDevice) HasDepRegistered() bool`

HasDepRegistered returns a boolean if a field has been set.

### GetDeviceInformation

`func (o *AppleMdmDevice) GetDeviceInformation() AppleMdmDeviceInfo`

GetDeviceInformation returns the DeviceInformation field if non-nil, zero value otherwise.

### GetDeviceInformationOk

`func (o *AppleMdmDevice) GetDeviceInformationOk() (*AppleMdmDeviceInfo, bool)`

GetDeviceInformationOk returns a tuple with the DeviceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInformation

`func (o *AppleMdmDevice) SetDeviceInformation(v AppleMdmDeviceInfo)`

SetDeviceInformation sets DeviceInformation field to given value.

### HasDeviceInformation

`func (o *AppleMdmDevice) HasDeviceInformation() bool`

HasDeviceInformation returns a boolean if a field has been set.

### GetEnrolled

`func (o *AppleMdmDevice) GetEnrolled() bool`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *AppleMdmDevice) GetEnrolledOk() (*bool, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *AppleMdmDevice) SetEnrolled(v bool)`

SetEnrolled sets Enrolled field to given value.

### HasEnrolled

`func (o *AppleMdmDevice) HasEnrolled() bool`

HasEnrolled returns a boolean if a field has been set.

### GetHasActivationLockBypassCodes

`func (o *AppleMdmDevice) GetHasActivationLockBypassCodes() bool`

GetHasActivationLockBypassCodes returns the HasActivationLockBypassCodes field if non-nil, zero value otherwise.

### GetHasActivationLockBypassCodesOk

`func (o *AppleMdmDevice) GetHasActivationLockBypassCodesOk() (*bool, bool)`

GetHasActivationLockBypassCodesOk returns a tuple with the HasActivationLockBypassCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActivationLockBypassCodes

`func (o *AppleMdmDevice) SetHasActivationLockBypassCodes(v bool)`

SetHasActivationLockBypassCodes sets HasActivationLockBypassCodes field to given value.

### HasHasActivationLockBypassCodes

`func (o *AppleMdmDevice) HasHasActivationLockBypassCodes() bool`

HasHasActivationLockBypassCodes returns a boolean if a field has been set.

### GetId

`func (o *AppleMdmDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppleMdmDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppleMdmDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppleMdmDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOsVersion

`func (o *AppleMdmDevice) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *AppleMdmDevice) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *AppleMdmDevice) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *AppleMdmDevice) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetSecurityInfo

`func (o *AppleMdmDevice) GetSecurityInfo() AppleMdmDeviceSecurityInfo`

GetSecurityInfo returns the SecurityInfo field if non-nil, zero value otherwise.

### GetSecurityInfoOk

`func (o *AppleMdmDevice) GetSecurityInfoOk() (*AppleMdmDeviceSecurityInfo, bool)`

GetSecurityInfoOk returns a tuple with the SecurityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityInfo

`func (o *AppleMdmDevice) SetSecurityInfo(v AppleMdmDeviceSecurityInfo)`

SetSecurityInfo sets SecurityInfo field to given value.

### HasSecurityInfo

`func (o *AppleMdmDevice) HasSecurityInfo() bool`

HasSecurityInfo returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AppleMdmDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AppleMdmDevice) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AppleMdmDevice) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AppleMdmDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetUdid

`func (o *AppleMdmDevice) GetUdid() string`

GetUdid returns the Udid field if non-nil, zero value otherwise.

### GetUdidOk

`func (o *AppleMdmDevice) GetUdidOk() (*string, bool)`

GetUdidOk returns a tuple with the Udid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdid

`func (o *AppleMdmDevice) SetUdid(v string)`

SetUdid sets Udid field to given value.

### HasUdid

`func (o *AppleMdmDevice) HasUdid() bool`

HasUdid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DevicesResetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to **[]string** |  | [optional] 
**NewPassword** | Pointer to **string** | Not logging as it contains sensitive information. | [optional] 

## Methods

### NewDevicesResetPasswordRequest

`func NewDevicesResetPasswordRequest() *DevicesResetPasswordRequest`

NewDevicesResetPasswordRequest instantiates a new DevicesResetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesResetPasswordRequestWithDefaults

`func NewDevicesResetPasswordRequestWithDefaults() *DevicesResetPasswordRequest`

NewDevicesResetPasswordRequestWithDefaults instantiates a new DevicesResetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *DevicesResetPasswordRequest) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *DevicesResetPasswordRequest) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *DevicesResetPasswordRequest) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *DevicesResetPasswordRequest) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetNewPassword

`func (o *DevicesResetPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *DevicesResetPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *DevicesResetPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *DevicesResetPasswordRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



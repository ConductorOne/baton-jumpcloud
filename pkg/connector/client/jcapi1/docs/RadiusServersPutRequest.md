# RadiusServersPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCert** | Pointer to **string** |  | [optional] 
**DeviceCertEnabled** | Pointer to **bool** |  | [optional] 
**Mfa** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**NetworkSourceIp** | **string** |  | 
**SharedSecret** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**UserCertEnabled** | Pointer to **bool** |  | [optional] 
**UserLockoutAction** | Pointer to **string** |  | [optional] 
**UserPasswordEnabled** | Pointer to **bool** |  | [optional] 
**UserPasswordExpirationAction** | Pointer to **string** |  | [optional] 

## Methods

### NewRadiusServersPutRequest

`func NewRadiusServersPutRequest(name string, networkSourceIp string, sharedSecret string, ) *RadiusServersPutRequest`

NewRadiusServersPutRequest instantiates a new RadiusServersPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusServersPutRequestWithDefaults

`func NewRadiusServersPutRequestWithDefaults() *RadiusServersPutRequest`

NewRadiusServersPutRequestWithDefaults instantiates a new RadiusServersPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCert

`func (o *RadiusServersPutRequest) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *RadiusServersPutRequest) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *RadiusServersPutRequest) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *RadiusServersPutRequest) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetDeviceCertEnabled

`func (o *RadiusServersPutRequest) GetDeviceCertEnabled() bool`

GetDeviceCertEnabled returns the DeviceCertEnabled field if non-nil, zero value otherwise.

### GetDeviceCertEnabledOk

`func (o *RadiusServersPutRequest) GetDeviceCertEnabledOk() (*bool, bool)`

GetDeviceCertEnabledOk returns a tuple with the DeviceCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCertEnabled

`func (o *RadiusServersPutRequest) SetDeviceCertEnabled(v bool)`

SetDeviceCertEnabled sets DeviceCertEnabled field to given value.

### HasDeviceCertEnabled

`func (o *RadiusServersPutRequest) HasDeviceCertEnabled() bool`

HasDeviceCertEnabled returns a boolean if a field has been set.

### GetMfa

`func (o *RadiusServersPutRequest) GetMfa() string`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *RadiusServersPutRequest) GetMfaOk() (*string, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *RadiusServersPutRequest) SetMfa(v string)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *RadiusServersPutRequest) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetName

`func (o *RadiusServersPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadiusServersPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadiusServersPutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkSourceIp

`func (o *RadiusServersPutRequest) GetNetworkSourceIp() string`

GetNetworkSourceIp returns the NetworkSourceIp field if non-nil, zero value otherwise.

### GetNetworkSourceIpOk

`func (o *RadiusServersPutRequest) GetNetworkSourceIpOk() (*string, bool)`

GetNetworkSourceIpOk returns a tuple with the NetworkSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSourceIp

`func (o *RadiusServersPutRequest) SetNetworkSourceIp(v string)`

SetNetworkSourceIp sets NetworkSourceIp field to given value.


### GetSharedSecret

`func (o *RadiusServersPutRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *RadiusServersPutRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *RadiusServersPutRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.


### GetTags

`func (o *RadiusServersPutRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RadiusServersPutRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RadiusServersPutRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RadiusServersPutRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserCertEnabled

`func (o *RadiusServersPutRequest) GetUserCertEnabled() bool`

GetUserCertEnabled returns the UserCertEnabled field if non-nil, zero value otherwise.

### GetUserCertEnabledOk

`func (o *RadiusServersPutRequest) GetUserCertEnabledOk() (*bool, bool)`

GetUserCertEnabledOk returns a tuple with the UserCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCertEnabled

`func (o *RadiusServersPutRequest) SetUserCertEnabled(v bool)`

SetUserCertEnabled sets UserCertEnabled field to given value.

### HasUserCertEnabled

`func (o *RadiusServersPutRequest) HasUserCertEnabled() bool`

HasUserCertEnabled returns a boolean if a field has been set.

### GetUserLockoutAction

`func (o *RadiusServersPutRequest) GetUserLockoutAction() string`

GetUserLockoutAction returns the UserLockoutAction field if non-nil, zero value otherwise.

### GetUserLockoutActionOk

`func (o *RadiusServersPutRequest) GetUserLockoutActionOk() (*string, bool)`

GetUserLockoutActionOk returns a tuple with the UserLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutAction

`func (o *RadiusServersPutRequest) SetUserLockoutAction(v string)`

SetUserLockoutAction sets UserLockoutAction field to given value.

### HasUserLockoutAction

`func (o *RadiusServersPutRequest) HasUserLockoutAction() bool`

HasUserLockoutAction returns a boolean if a field has been set.

### GetUserPasswordEnabled

`func (o *RadiusServersPutRequest) GetUserPasswordEnabled() bool`

GetUserPasswordEnabled returns the UserPasswordEnabled field if non-nil, zero value otherwise.

### GetUserPasswordEnabledOk

`func (o *RadiusServersPutRequest) GetUserPasswordEnabledOk() (*bool, bool)`

GetUserPasswordEnabledOk returns a tuple with the UserPasswordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordEnabled

`func (o *RadiusServersPutRequest) SetUserPasswordEnabled(v bool)`

SetUserPasswordEnabled sets UserPasswordEnabled field to given value.

### HasUserPasswordEnabled

`func (o *RadiusServersPutRequest) HasUserPasswordEnabled() bool`

HasUserPasswordEnabled returns a boolean if a field has been set.

### GetUserPasswordExpirationAction

`func (o *RadiusServersPutRequest) GetUserPasswordExpirationAction() string`

GetUserPasswordExpirationAction returns the UserPasswordExpirationAction field if non-nil, zero value otherwise.

### GetUserPasswordExpirationActionOk

`func (o *RadiusServersPutRequest) GetUserPasswordExpirationActionOk() (*string, bool)`

GetUserPasswordExpirationActionOk returns a tuple with the UserPasswordExpirationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordExpirationAction

`func (o *RadiusServersPutRequest) SetUserPasswordExpirationAction(v string)`

SetUserPasswordExpirationAction sets UserPasswordExpirationAction field to given value.

### HasUserPasswordExpirationAction

`func (o *RadiusServersPutRequest) HasUserPasswordExpirationAction() bool`

HasUserPasswordExpirationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



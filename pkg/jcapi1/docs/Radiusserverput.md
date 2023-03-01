# Radiusserverput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AuthIdp** | Pointer to **string** |  | [optional] 
**CaCert** | Pointer to **string** |  | [optional] 
**DeviceCertEnabled** | Pointer to **bool** |  | [optional] 
**Mfa** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetworkSourceIp** | Pointer to **string** |  | [optional] 
**TagNames** | Pointer to **[]string** |  | [optional] 
**UserCertEnabled** | Pointer to **bool** |  | [optional] 
**UserLockoutAction** | Pointer to **string** |  | [optional] 
**UserPasswordEnabled** | Pointer to **bool** |  | [optional] 
**UserPasswordExpirationAction** | Pointer to **string** |  | [optional] 

## Methods

### NewRadiusserverput

`func NewRadiusserverput() *Radiusserverput`

NewRadiusserverput instantiates a new Radiusserverput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusserverputWithDefaults

`func NewRadiusserverputWithDefaults() *Radiusserverput`

NewRadiusserverputWithDefaults instantiates a new Radiusserverput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Radiusserverput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Radiusserverput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Radiusserverput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Radiusserverput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthIdp

`func (o *Radiusserverput) GetAuthIdp() string`

GetAuthIdp returns the AuthIdp field if non-nil, zero value otherwise.

### GetAuthIdpOk

`func (o *Radiusserverput) GetAuthIdpOk() (*string, bool)`

GetAuthIdpOk returns a tuple with the AuthIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIdp

`func (o *Radiusserverput) SetAuthIdp(v string)`

SetAuthIdp sets AuthIdp field to given value.

### HasAuthIdp

`func (o *Radiusserverput) HasAuthIdp() bool`

HasAuthIdp returns a boolean if a field has been set.

### GetCaCert

`func (o *Radiusserverput) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *Radiusserverput) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *Radiusserverput) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *Radiusserverput) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetDeviceCertEnabled

`func (o *Radiusserverput) GetDeviceCertEnabled() bool`

GetDeviceCertEnabled returns the DeviceCertEnabled field if non-nil, zero value otherwise.

### GetDeviceCertEnabledOk

`func (o *Radiusserverput) GetDeviceCertEnabledOk() (*bool, bool)`

GetDeviceCertEnabledOk returns a tuple with the DeviceCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCertEnabled

`func (o *Radiusserverput) SetDeviceCertEnabled(v bool)`

SetDeviceCertEnabled sets DeviceCertEnabled field to given value.

### HasDeviceCertEnabled

`func (o *Radiusserverput) HasDeviceCertEnabled() bool`

HasDeviceCertEnabled returns a boolean if a field has been set.

### GetMfa

`func (o *Radiusserverput) GetMfa() string`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *Radiusserverput) GetMfaOk() (*string, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *Radiusserverput) SetMfa(v string)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *Radiusserverput) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetName

`func (o *Radiusserverput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Radiusserverput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Radiusserverput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Radiusserverput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkSourceIp

`func (o *Radiusserverput) GetNetworkSourceIp() string`

GetNetworkSourceIp returns the NetworkSourceIp field if non-nil, zero value otherwise.

### GetNetworkSourceIpOk

`func (o *Radiusserverput) GetNetworkSourceIpOk() (*string, bool)`

GetNetworkSourceIpOk returns a tuple with the NetworkSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSourceIp

`func (o *Radiusserverput) SetNetworkSourceIp(v string)`

SetNetworkSourceIp sets NetworkSourceIp field to given value.

### HasNetworkSourceIp

`func (o *Radiusserverput) HasNetworkSourceIp() bool`

HasNetworkSourceIp returns a boolean if a field has been set.

### GetTagNames

`func (o *Radiusserverput) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *Radiusserverput) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *Radiusserverput) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *Radiusserverput) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetUserCertEnabled

`func (o *Radiusserverput) GetUserCertEnabled() bool`

GetUserCertEnabled returns the UserCertEnabled field if non-nil, zero value otherwise.

### GetUserCertEnabledOk

`func (o *Radiusserverput) GetUserCertEnabledOk() (*bool, bool)`

GetUserCertEnabledOk returns a tuple with the UserCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCertEnabled

`func (o *Radiusserverput) SetUserCertEnabled(v bool)`

SetUserCertEnabled sets UserCertEnabled field to given value.

### HasUserCertEnabled

`func (o *Radiusserverput) HasUserCertEnabled() bool`

HasUserCertEnabled returns a boolean if a field has been set.

### GetUserLockoutAction

`func (o *Radiusserverput) GetUserLockoutAction() string`

GetUserLockoutAction returns the UserLockoutAction field if non-nil, zero value otherwise.

### GetUserLockoutActionOk

`func (o *Radiusserverput) GetUserLockoutActionOk() (*string, bool)`

GetUserLockoutActionOk returns a tuple with the UserLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutAction

`func (o *Radiusserverput) SetUserLockoutAction(v string)`

SetUserLockoutAction sets UserLockoutAction field to given value.

### HasUserLockoutAction

`func (o *Radiusserverput) HasUserLockoutAction() bool`

HasUserLockoutAction returns a boolean if a field has been set.

### GetUserPasswordEnabled

`func (o *Radiusserverput) GetUserPasswordEnabled() bool`

GetUserPasswordEnabled returns the UserPasswordEnabled field if non-nil, zero value otherwise.

### GetUserPasswordEnabledOk

`func (o *Radiusserverput) GetUserPasswordEnabledOk() (*bool, bool)`

GetUserPasswordEnabledOk returns a tuple with the UserPasswordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordEnabled

`func (o *Radiusserverput) SetUserPasswordEnabled(v bool)`

SetUserPasswordEnabled sets UserPasswordEnabled field to given value.

### HasUserPasswordEnabled

`func (o *Radiusserverput) HasUserPasswordEnabled() bool`

HasUserPasswordEnabled returns a boolean if a field has been set.

### GetUserPasswordExpirationAction

`func (o *Radiusserverput) GetUserPasswordExpirationAction() string`

GetUserPasswordExpirationAction returns the UserPasswordExpirationAction field if non-nil, zero value otherwise.

### GetUserPasswordExpirationActionOk

`func (o *Radiusserverput) GetUserPasswordExpirationActionOk() (*string, bool)`

GetUserPasswordExpirationActionOk returns a tuple with the UserPasswordExpirationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordExpirationAction

`func (o *Radiusserverput) SetUserPasswordExpirationAction(v string)`

SetUserPasswordExpirationAction sets UserPasswordExpirationAction field to given value.

### HasUserPasswordExpirationAction

`func (o *Radiusserverput) HasUserPasswordExpirationAction() bool`

HasUserPasswordExpirationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



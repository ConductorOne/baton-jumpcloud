# Radiusserverpost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthIdp** | Pointer to **string** |  | [optional] 
**CaCert** | Pointer to **string** |  | [optional] 
**DeviceCertEnabled** | Pointer to **bool** |  | [optional] 
**Mfa** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**NetworkSourceIp** | **string** |  | 
**SharedSecret** | **string** | RADIUS shared secret between the server and client. | 
**TagNames** | Pointer to **[]string** |  | [optional] 
**UserCertEnabled** | Pointer to **bool** |  | [optional] 
**UserLockoutAction** | Pointer to **string** |  | [optional] 
**UserPasswordEnabled** | Pointer to **bool** |  | [optional] 
**UserPasswordExpirationAction** | Pointer to **string** |  | [optional] 

## Methods

### NewRadiusserverpost

`func NewRadiusserverpost(name string, networkSourceIp string, sharedSecret string, ) *Radiusserverpost`

NewRadiusserverpost instantiates a new Radiusserverpost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusserverpostWithDefaults

`func NewRadiusserverpostWithDefaults() *Radiusserverpost`

NewRadiusserverpostWithDefaults instantiates a new Radiusserverpost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthIdp

`func (o *Radiusserverpost) GetAuthIdp() string`

GetAuthIdp returns the AuthIdp field if non-nil, zero value otherwise.

### GetAuthIdpOk

`func (o *Radiusserverpost) GetAuthIdpOk() (*string, bool)`

GetAuthIdpOk returns a tuple with the AuthIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIdp

`func (o *Radiusserverpost) SetAuthIdp(v string)`

SetAuthIdp sets AuthIdp field to given value.

### HasAuthIdp

`func (o *Radiusserverpost) HasAuthIdp() bool`

HasAuthIdp returns a boolean if a field has been set.

### GetCaCert

`func (o *Radiusserverpost) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *Radiusserverpost) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *Radiusserverpost) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *Radiusserverpost) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetDeviceCertEnabled

`func (o *Radiusserverpost) GetDeviceCertEnabled() bool`

GetDeviceCertEnabled returns the DeviceCertEnabled field if non-nil, zero value otherwise.

### GetDeviceCertEnabledOk

`func (o *Radiusserverpost) GetDeviceCertEnabledOk() (*bool, bool)`

GetDeviceCertEnabledOk returns a tuple with the DeviceCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCertEnabled

`func (o *Radiusserverpost) SetDeviceCertEnabled(v bool)`

SetDeviceCertEnabled sets DeviceCertEnabled field to given value.

### HasDeviceCertEnabled

`func (o *Radiusserverpost) HasDeviceCertEnabled() bool`

HasDeviceCertEnabled returns a boolean if a field has been set.

### GetMfa

`func (o *Radiusserverpost) GetMfa() string`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *Radiusserverpost) GetMfaOk() (*string, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *Radiusserverpost) SetMfa(v string)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *Radiusserverpost) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetName

`func (o *Radiusserverpost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Radiusserverpost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Radiusserverpost) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkSourceIp

`func (o *Radiusserverpost) GetNetworkSourceIp() string`

GetNetworkSourceIp returns the NetworkSourceIp field if non-nil, zero value otherwise.

### GetNetworkSourceIpOk

`func (o *Radiusserverpost) GetNetworkSourceIpOk() (*string, bool)`

GetNetworkSourceIpOk returns a tuple with the NetworkSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSourceIp

`func (o *Radiusserverpost) SetNetworkSourceIp(v string)`

SetNetworkSourceIp sets NetworkSourceIp field to given value.


### GetSharedSecret

`func (o *Radiusserverpost) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *Radiusserverpost) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *Radiusserverpost) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.


### GetTagNames

`func (o *Radiusserverpost) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *Radiusserverpost) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *Radiusserverpost) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *Radiusserverpost) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetUserCertEnabled

`func (o *Radiusserverpost) GetUserCertEnabled() bool`

GetUserCertEnabled returns the UserCertEnabled field if non-nil, zero value otherwise.

### GetUserCertEnabledOk

`func (o *Radiusserverpost) GetUserCertEnabledOk() (*bool, bool)`

GetUserCertEnabledOk returns a tuple with the UserCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCertEnabled

`func (o *Radiusserverpost) SetUserCertEnabled(v bool)`

SetUserCertEnabled sets UserCertEnabled field to given value.

### HasUserCertEnabled

`func (o *Radiusserverpost) HasUserCertEnabled() bool`

HasUserCertEnabled returns a boolean if a field has been set.

### GetUserLockoutAction

`func (o *Radiusserverpost) GetUserLockoutAction() string`

GetUserLockoutAction returns the UserLockoutAction field if non-nil, zero value otherwise.

### GetUserLockoutActionOk

`func (o *Radiusserverpost) GetUserLockoutActionOk() (*string, bool)`

GetUserLockoutActionOk returns a tuple with the UserLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutAction

`func (o *Radiusserverpost) SetUserLockoutAction(v string)`

SetUserLockoutAction sets UserLockoutAction field to given value.

### HasUserLockoutAction

`func (o *Radiusserverpost) HasUserLockoutAction() bool`

HasUserLockoutAction returns a boolean if a field has been set.

### GetUserPasswordEnabled

`func (o *Radiusserverpost) GetUserPasswordEnabled() bool`

GetUserPasswordEnabled returns the UserPasswordEnabled field if non-nil, zero value otherwise.

### GetUserPasswordEnabledOk

`func (o *Radiusserverpost) GetUserPasswordEnabledOk() (*bool, bool)`

GetUserPasswordEnabledOk returns a tuple with the UserPasswordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordEnabled

`func (o *Radiusserverpost) SetUserPasswordEnabled(v bool)`

SetUserPasswordEnabled sets UserPasswordEnabled field to given value.

### HasUserPasswordEnabled

`func (o *Radiusserverpost) HasUserPasswordEnabled() bool`

HasUserPasswordEnabled returns a boolean if a field has been set.

### GetUserPasswordExpirationAction

`func (o *Radiusserverpost) GetUserPasswordExpirationAction() string`

GetUserPasswordExpirationAction returns the UserPasswordExpirationAction field if non-nil, zero value otherwise.

### GetUserPasswordExpirationActionOk

`func (o *Radiusserverpost) GetUserPasswordExpirationActionOk() (*string, bool)`

GetUserPasswordExpirationActionOk returns a tuple with the UserPasswordExpirationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordExpirationAction

`func (o *Radiusserverpost) SetUserPasswordExpirationAction(v string)`

SetUserPasswordExpirationAction sets UserPasswordExpirationAction field to given value.

### HasUserPasswordExpirationAction

`func (o *Radiusserverpost) HasUserPasswordExpirationAction() bool`

HasUserPasswordExpirationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



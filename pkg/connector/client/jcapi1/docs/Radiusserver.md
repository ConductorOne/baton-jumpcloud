# Radiusserver

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
**Organization** | Pointer to **string** |  | [optional] 
**SharedSecret** | Pointer to **string** |  | [optional] 
**TagNames** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UserCertEnabled** | Pointer to **bool** |  | [optional] 
**UserLockoutAction** | Pointer to **string** |  | [optional] 
**UserPasswordEnabled** | Pointer to **bool** |  | [optional] 
**UserPasswordExpirationAction** | Pointer to **string** |  | [optional] 

## Methods

### NewRadiusserver

`func NewRadiusserver() *Radiusserver`

NewRadiusserver instantiates a new Radiusserver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusserverWithDefaults

`func NewRadiusserverWithDefaults() *Radiusserver`

NewRadiusserverWithDefaults instantiates a new Radiusserver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Radiusserver) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Radiusserver) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Radiusserver) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Radiusserver) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthIdp

`func (o *Radiusserver) GetAuthIdp() string`

GetAuthIdp returns the AuthIdp field if non-nil, zero value otherwise.

### GetAuthIdpOk

`func (o *Radiusserver) GetAuthIdpOk() (*string, bool)`

GetAuthIdpOk returns a tuple with the AuthIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIdp

`func (o *Radiusserver) SetAuthIdp(v string)`

SetAuthIdp sets AuthIdp field to given value.

### HasAuthIdp

`func (o *Radiusserver) HasAuthIdp() bool`

HasAuthIdp returns a boolean if a field has been set.

### GetCaCert

`func (o *Radiusserver) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *Radiusserver) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *Radiusserver) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *Radiusserver) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetDeviceCertEnabled

`func (o *Radiusserver) GetDeviceCertEnabled() bool`

GetDeviceCertEnabled returns the DeviceCertEnabled field if non-nil, zero value otherwise.

### GetDeviceCertEnabledOk

`func (o *Radiusserver) GetDeviceCertEnabledOk() (*bool, bool)`

GetDeviceCertEnabledOk returns a tuple with the DeviceCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCertEnabled

`func (o *Radiusserver) SetDeviceCertEnabled(v bool)`

SetDeviceCertEnabled sets DeviceCertEnabled field to given value.

### HasDeviceCertEnabled

`func (o *Radiusserver) HasDeviceCertEnabled() bool`

HasDeviceCertEnabled returns a boolean if a field has been set.

### GetMfa

`func (o *Radiusserver) GetMfa() string`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *Radiusserver) GetMfaOk() (*string, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *Radiusserver) SetMfa(v string)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *Radiusserver) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetName

`func (o *Radiusserver) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Radiusserver) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Radiusserver) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Radiusserver) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkSourceIp

`func (o *Radiusserver) GetNetworkSourceIp() string`

GetNetworkSourceIp returns the NetworkSourceIp field if non-nil, zero value otherwise.

### GetNetworkSourceIpOk

`func (o *Radiusserver) GetNetworkSourceIpOk() (*string, bool)`

GetNetworkSourceIpOk returns a tuple with the NetworkSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSourceIp

`func (o *Radiusserver) SetNetworkSourceIp(v string)`

SetNetworkSourceIp sets NetworkSourceIp field to given value.

### HasNetworkSourceIp

`func (o *Radiusserver) HasNetworkSourceIp() bool`

HasNetworkSourceIp returns a boolean if a field has been set.

### GetOrganization

`func (o *Radiusserver) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Radiusserver) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Radiusserver) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Radiusserver) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSharedSecret

`func (o *Radiusserver) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *Radiusserver) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *Radiusserver) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *Radiusserver) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetTagNames

`func (o *Radiusserver) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *Radiusserver) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *Radiusserver) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *Radiusserver) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetTags

`func (o *Radiusserver) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Radiusserver) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Radiusserver) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Radiusserver) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserCertEnabled

`func (o *Radiusserver) GetUserCertEnabled() bool`

GetUserCertEnabled returns the UserCertEnabled field if non-nil, zero value otherwise.

### GetUserCertEnabledOk

`func (o *Radiusserver) GetUserCertEnabledOk() (*bool, bool)`

GetUserCertEnabledOk returns a tuple with the UserCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCertEnabled

`func (o *Radiusserver) SetUserCertEnabled(v bool)`

SetUserCertEnabled sets UserCertEnabled field to given value.

### HasUserCertEnabled

`func (o *Radiusserver) HasUserCertEnabled() bool`

HasUserCertEnabled returns a boolean if a field has been set.

### GetUserLockoutAction

`func (o *Radiusserver) GetUserLockoutAction() string`

GetUserLockoutAction returns the UserLockoutAction field if non-nil, zero value otherwise.

### GetUserLockoutActionOk

`func (o *Radiusserver) GetUserLockoutActionOk() (*string, bool)`

GetUserLockoutActionOk returns a tuple with the UserLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutAction

`func (o *Radiusserver) SetUserLockoutAction(v string)`

SetUserLockoutAction sets UserLockoutAction field to given value.

### HasUserLockoutAction

`func (o *Radiusserver) HasUserLockoutAction() bool`

HasUserLockoutAction returns a boolean if a field has been set.

### GetUserPasswordEnabled

`func (o *Radiusserver) GetUserPasswordEnabled() bool`

GetUserPasswordEnabled returns the UserPasswordEnabled field if non-nil, zero value otherwise.

### GetUserPasswordEnabledOk

`func (o *Radiusserver) GetUserPasswordEnabledOk() (*bool, bool)`

GetUserPasswordEnabledOk returns a tuple with the UserPasswordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordEnabled

`func (o *Radiusserver) SetUserPasswordEnabled(v bool)`

SetUserPasswordEnabled sets UserPasswordEnabled field to given value.

### HasUserPasswordEnabled

`func (o *Radiusserver) HasUserPasswordEnabled() bool`

HasUserPasswordEnabled returns a boolean if a field has been set.

### GetUserPasswordExpirationAction

`func (o *Radiusserver) GetUserPasswordExpirationAction() string`

GetUserPasswordExpirationAction returns the UserPasswordExpirationAction field if non-nil, zero value otherwise.

### GetUserPasswordExpirationActionOk

`func (o *Radiusserver) GetUserPasswordExpirationActionOk() (*string, bool)`

GetUserPasswordExpirationActionOk returns a tuple with the UserPasswordExpirationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordExpirationAction

`func (o *Radiusserver) SetUserPasswordExpirationAction(v string)`

SetUserPasswordExpirationAction sets UserPasswordExpirationAction field to given value.

### HasUserPasswordExpirationAction

`func (o *Radiusserver) HasUserPasswordExpirationAction() bool`

HasUserPasswordExpirationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



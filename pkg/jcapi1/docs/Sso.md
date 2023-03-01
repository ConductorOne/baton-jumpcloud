# Sso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beta** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**IdpCertExpirationAt** | Pointer to **time.Time** |  | [optional] 
**Jit** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSso

`func NewSso() *Sso`

NewSso instantiates a new Sso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoWithDefaults

`func NewSsoWithDefaults() *Sso`

NewSsoWithDefaults instantiates a new Sso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeta

`func (o *Sso) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *Sso) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *Sso) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *Sso) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetHidden

`func (o *Sso) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Sso) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Sso) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Sso) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetIdpCertExpirationAt

`func (o *Sso) GetIdpCertExpirationAt() time.Time`

GetIdpCertExpirationAt returns the IdpCertExpirationAt field if non-nil, zero value otherwise.

### GetIdpCertExpirationAtOk

`func (o *Sso) GetIdpCertExpirationAtOk() (*time.Time, bool)`

GetIdpCertExpirationAtOk returns a tuple with the IdpCertExpirationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertExpirationAt

`func (o *Sso) SetIdpCertExpirationAt(v time.Time)`

SetIdpCertExpirationAt sets IdpCertExpirationAt field to given value.

### HasIdpCertExpirationAt

`func (o *Sso) HasIdpCertExpirationAt() bool`

HasIdpCertExpirationAt returns a boolean if a field has been set.

### GetJit

`func (o *Sso) GetJit() bool`

GetJit returns the Jit field if non-nil, zero value otherwise.

### GetJitOk

`func (o *Sso) GetJitOk() (*bool, bool)`

GetJitOk returns a tuple with the Jit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJit

`func (o *Sso) SetJit(v bool)`

SetJit sets Jit field to given value.

### HasJit

`func (o *Sso) HasJit() bool`

HasJit returns a boolean if a field has been set.

### GetType

`func (o *Sso) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Sso) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Sso) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Sso) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



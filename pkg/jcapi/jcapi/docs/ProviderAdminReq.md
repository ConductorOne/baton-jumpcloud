# ProviderAdminReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindNoOrgs** | Pointer to **bool** |  | [optional] [default to false]
**Email** | **string** |  | 
**EnableMultiFactor** | Pointer to **bool** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 

## Methods

### NewProviderAdminReq

`func NewProviderAdminReq(email string, ) *ProviderAdminReq`

NewProviderAdminReq instantiates a new ProviderAdminReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderAdminReqWithDefaults

`func NewProviderAdminReqWithDefaults() *ProviderAdminReq`

NewProviderAdminReqWithDefaults instantiates a new ProviderAdminReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindNoOrgs

`func (o *ProviderAdminReq) GetBindNoOrgs() bool`

GetBindNoOrgs returns the BindNoOrgs field if non-nil, zero value otherwise.

### GetBindNoOrgsOk

`func (o *ProviderAdminReq) GetBindNoOrgsOk() (*bool, bool)`

GetBindNoOrgsOk returns a tuple with the BindNoOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindNoOrgs

`func (o *ProviderAdminReq) SetBindNoOrgs(v bool)`

SetBindNoOrgs sets BindNoOrgs field to given value.

### HasBindNoOrgs

`func (o *ProviderAdminReq) HasBindNoOrgs() bool`

HasBindNoOrgs returns a boolean if a field has been set.

### GetEmail

`func (o *ProviderAdminReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProviderAdminReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProviderAdminReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnableMultiFactor

`func (o *ProviderAdminReq) GetEnableMultiFactor() bool`

GetEnableMultiFactor returns the EnableMultiFactor field if non-nil, zero value otherwise.

### GetEnableMultiFactorOk

`func (o *ProviderAdminReq) GetEnableMultiFactorOk() (*bool, bool)`

GetEnableMultiFactorOk returns a tuple with the EnableMultiFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiFactor

`func (o *ProviderAdminReq) SetEnableMultiFactor(v bool)`

SetEnableMultiFactor sets EnableMultiFactor field to given value.

### HasEnableMultiFactor

`func (o *ProviderAdminReq) HasEnableMultiFactor() bool`

HasEnableMultiFactor returns a boolean if a field has been set.

### GetFirstname

`func (o *ProviderAdminReq) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *ProviderAdminReq) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *ProviderAdminReq) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *ProviderAdminReq) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *ProviderAdminReq) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *ProviderAdminReq) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *ProviderAdminReq) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *ProviderAdminReq) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetRole

`func (o *ProviderAdminReq) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProviderAdminReq) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProviderAdminReq) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ProviderAdminReq) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoleName

`func (o *ProviderAdminReq) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *ProviderAdminReq) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *ProviderAdminReq) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *ProviderAdminReq) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



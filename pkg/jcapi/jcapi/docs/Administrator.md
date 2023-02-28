# Administrator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**EnableMultiFactor** | Pointer to **bool** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**OrganizationAccessTotal** | Pointer to **float32** |  | [optional] 
**Registered** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 

## Methods

### NewAdministrator

`func NewAdministrator() *Administrator`

NewAdministrator instantiates a new Administrator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorWithDefaults

`func NewAdministratorWithDefaults() *Administrator`

NewAdministratorWithDefaults instantiates a new Administrator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Administrator) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Administrator) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Administrator) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Administrator) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnableMultiFactor

`func (o *Administrator) GetEnableMultiFactor() bool`

GetEnableMultiFactor returns the EnableMultiFactor field if non-nil, zero value otherwise.

### GetEnableMultiFactorOk

`func (o *Administrator) GetEnableMultiFactorOk() (*bool, bool)`

GetEnableMultiFactorOk returns a tuple with the EnableMultiFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiFactor

`func (o *Administrator) SetEnableMultiFactor(v bool)`

SetEnableMultiFactor sets EnableMultiFactor field to given value.

### HasEnableMultiFactor

`func (o *Administrator) HasEnableMultiFactor() bool`

HasEnableMultiFactor returns a boolean if a field has been set.

### GetFirstname

`func (o *Administrator) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *Administrator) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *Administrator) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *Administrator) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetId

`func (o *Administrator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Administrator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Administrator) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Administrator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastname

`func (o *Administrator) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *Administrator) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *Administrator) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *Administrator) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetOrganizationAccessTotal

`func (o *Administrator) GetOrganizationAccessTotal() float32`

GetOrganizationAccessTotal returns the OrganizationAccessTotal field if non-nil, zero value otherwise.

### GetOrganizationAccessTotalOk

`func (o *Administrator) GetOrganizationAccessTotalOk() (*float32, bool)`

GetOrganizationAccessTotalOk returns a tuple with the OrganizationAccessTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationAccessTotal

`func (o *Administrator) SetOrganizationAccessTotal(v float32)`

SetOrganizationAccessTotal sets OrganizationAccessTotal field to given value.

### HasOrganizationAccessTotal

`func (o *Administrator) HasOrganizationAccessTotal() bool`

HasOrganizationAccessTotal returns a boolean if a field has been set.

### GetRegistered

`func (o *Administrator) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *Administrator) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *Administrator) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *Administrator) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.

### GetRole

`func (o *Administrator) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Administrator) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Administrator) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Administrator) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoleName

`func (o *Administrator) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *Administrator) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *Administrator) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *Administrator) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetSuspended

`func (o *Administrator) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Administrator) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Administrator) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Administrator) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Userreturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**DisableIntroduction** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EnableMultiFactor** | Pointer to **bool** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**GrowthData** | Pointer to [**UserreturnGrowthData**](UserreturnGrowthData.md) |  | [optional] 
**LastWhatsNewChecked** | Pointer to **time.Time** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**SessionCount** | Pointer to **int32** |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**TotpEnrolled** | Pointer to **bool** |  | [optional] 
**UsersTimeZone** | Pointer to **string** |  | [optional] 

## Methods

### NewUserreturn

`func NewUserreturn() *Userreturn`

NewUserreturn instantiates a new Userreturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserreturnWithDefaults

`func NewUserreturnWithDefaults() *Userreturn`

NewUserreturnWithDefaults instantiates a new Userreturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Userreturn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Userreturn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Userreturn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Userreturn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *Userreturn) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Userreturn) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Userreturn) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Userreturn) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDisableIntroduction

`func (o *Userreturn) GetDisableIntroduction() bool`

GetDisableIntroduction returns the DisableIntroduction field if non-nil, zero value otherwise.

### GetDisableIntroductionOk

`func (o *Userreturn) GetDisableIntroductionOk() (*bool, bool)`

GetDisableIntroductionOk returns a tuple with the DisableIntroduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableIntroduction

`func (o *Userreturn) SetDisableIntroduction(v bool)`

SetDisableIntroduction sets DisableIntroduction field to given value.

### HasDisableIntroduction

`func (o *Userreturn) HasDisableIntroduction() bool`

HasDisableIntroduction returns a boolean if a field has been set.

### GetEmail

`func (o *Userreturn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Userreturn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Userreturn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Userreturn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnableMultiFactor

`func (o *Userreturn) GetEnableMultiFactor() bool`

GetEnableMultiFactor returns the EnableMultiFactor field if non-nil, zero value otherwise.

### GetEnableMultiFactorOk

`func (o *Userreturn) GetEnableMultiFactorOk() (*bool, bool)`

GetEnableMultiFactorOk returns a tuple with the EnableMultiFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiFactor

`func (o *Userreturn) SetEnableMultiFactor(v bool)`

SetEnableMultiFactor sets EnableMultiFactor field to given value.

### HasEnableMultiFactor

`func (o *Userreturn) HasEnableMultiFactor() bool`

HasEnableMultiFactor returns a boolean if a field has been set.

### GetFirstname

`func (o *Userreturn) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *Userreturn) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *Userreturn) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *Userreturn) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetGrowthData

`func (o *Userreturn) GetGrowthData() UserreturnGrowthData`

GetGrowthData returns the GrowthData field if non-nil, zero value otherwise.

### GetGrowthDataOk

`func (o *Userreturn) GetGrowthDataOk() (*UserreturnGrowthData, bool)`

GetGrowthDataOk returns a tuple with the GrowthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthData

`func (o *Userreturn) SetGrowthData(v UserreturnGrowthData)`

SetGrowthData sets GrowthData field to given value.

### HasGrowthData

`func (o *Userreturn) HasGrowthData() bool`

HasGrowthData returns a boolean if a field has been set.

### GetLastWhatsNewChecked

`func (o *Userreturn) GetLastWhatsNewChecked() time.Time`

GetLastWhatsNewChecked returns the LastWhatsNewChecked field if non-nil, zero value otherwise.

### GetLastWhatsNewCheckedOk

`func (o *Userreturn) GetLastWhatsNewCheckedOk() (*time.Time, bool)`

GetLastWhatsNewCheckedOk returns a tuple with the LastWhatsNewChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWhatsNewChecked

`func (o *Userreturn) SetLastWhatsNewChecked(v time.Time)`

SetLastWhatsNewChecked sets LastWhatsNewChecked field to given value.

### HasLastWhatsNewChecked

`func (o *Userreturn) HasLastWhatsNewChecked() bool`

HasLastWhatsNewChecked returns a boolean if a field has been set.

### GetLastname

`func (o *Userreturn) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *Userreturn) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *Userreturn) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *Userreturn) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetOrganization

`func (o *Userreturn) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Userreturn) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Userreturn) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Userreturn) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProvider

`func (o *Userreturn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Userreturn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Userreturn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Userreturn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRole

`func (o *Userreturn) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Userreturn) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Userreturn) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Userreturn) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoleName

`func (o *Userreturn) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *Userreturn) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *Userreturn) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *Userreturn) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetSessionCount

`func (o *Userreturn) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *Userreturn) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *Userreturn) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *Userreturn) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetSuspended

`func (o *Userreturn) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Userreturn) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Userreturn) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Userreturn) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetTotpEnrolled

`func (o *Userreturn) GetTotpEnrolled() bool`

GetTotpEnrolled returns the TotpEnrolled field if non-nil, zero value otherwise.

### GetTotpEnrolledOk

`func (o *Userreturn) GetTotpEnrolledOk() (*bool, bool)`

GetTotpEnrolledOk returns a tuple with the TotpEnrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpEnrolled

`func (o *Userreturn) SetTotpEnrolled(v bool)`

SetTotpEnrolled sets TotpEnrolled field to given value.

### HasTotpEnrolled

`func (o *Userreturn) HasTotpEnrolled() bool`

HasTotpEnrolled returns a boolean if a field has been set.

### GetUsersTimeZone

`func (o *Userreturn) GetUsersTimeZone() string`

GetUsersTimeZone returns the UsersTimeZone field if non-nil, zero value otherwise.

### GetUsersTimeZoneOk

`func (o *Userreturn) GetUsersTimeZoneOk() (*string, bool)`

GetUsersTimeZoneOk returns a tuple with the UsersTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersTimeZone

`func (o *Userreturn) SetUsersTimeZone(v string)`

SetUsersTimeZone sets UsersTimeZone field to given value.

### HasUsersTimeZone

`func (o *Userreturn) HasUsersTimeZone() bool`

HasUsersTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



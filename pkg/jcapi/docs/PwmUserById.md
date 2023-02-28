# PwmUserById

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]AppsInner**](AppsInner.md) |  | [optional] 
**Email** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]GroupsInner**](GroupsInner.md) |  | [optional] 
**Id** | **string** |  | 
**ItemsCount** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**PasswordsCount** | Pointer to **int32** |  | [optional] 
**PasswordsScore** | Pointer to **float32** |  | [optional] 
**ScoreUpdatedAt** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 
**CompromisedPasswords** | Pointer to **int32** |  | [optional] 
**OldPasswords** | Pointer to **int32** |  | [optional] 
**ReusedPasswords** | Pointer to **int32** |  | [optional] 
**WeakPasswords** | Pointer to **int32** |  | [optional] 

## Methods

### NewPwmUserById

`func NewPwmUserById(email string, id string, name string, status string, ) *PwmUserById`

NewPwmUserById instantiates a new PwmUserById object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPwmUserByIdWithDefaults

`func NewPwmUserByIdWithDefaults() *PwmUserById`

NewPwmUserByIdWithDefaults instantiates a new PwmUserById object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *PwmUserById) GetApps() []AppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PwmUserById) GetAppsOk() (*[]AppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PwmUserById) SetApps(v []AppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PwmUserById) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetEmail

`func (o *PwmUserById) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PwmUserById) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PwmUserById) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalId

`func (o *PwmUserById) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PwmUserById) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PwmUserById) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PwmUserById) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGroups

`func (o *PwmUserById) GetGroups() []GroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PwmUserById) GetGroupsOk() (*[]GroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PwmUserById) SetGroups(v []GroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PwmUserById) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *PwmUserById) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PwmUserById) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PwmUserById) SetId(v string)`

SetId sets Id field to given value.


### GetItemsCount

`func (o *PwmUserById) GetItemsCount() int32`

GetItemsCount returns the ItemsCount field if non-nil, zero value otherwise.

### GetItemsCountOk

`func (o *PwmUserById) GetItemsCountOk() (*int32, bool)`

GetItemsCountOk returns a tuple with the ItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsCount

`func (o *PwmUserById) SetItemsCount(v int32)`

SetItemsCount sets ItemsCount field to given value.

### HasItemsCount

`func (o *PwmUserById) HasItemsCount() bool`

HasItemsCount returns a boolean if a field has been set.

### GetName

`func (o *PwmUserById) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PwmUserById) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PwmUserById) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordsCount

`func (o *PwmUserById) GetPasswordsCount() int32`

GetPasswordsCount returns the PasswordsCount field if non-nil, zero value otherwise.

### GetPasswordsCountOk

`func (o *PwmUserById) GetPasswordsCountOk() (*int32, bool)`

GetPasswordsCountOk returns a tuple with the PasswordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsCount

`func (o *PwmUserById) SetPasswordsCount(v int32)`

SetPasswordsCount sets PasswordsCount field to given value.

### HasPasswordsCount

`func (o *PwmUserById) HasPasswordsCount() bool`

HasPasswordsCount returns a boolean if a field has been set.

### GetPasswordsScore

`func (o *PwmUserById) GetPasswordsScore() float32`

GetPasswordsScore returns the PasswordsScore field if non-nil, zero value otherwise.

### GetPasswordsScoreOk

`func (o *PwmUserById) GetPasswordsScoreOk() (*float32, bool)`

GetPasswordsScoreOk returns a tuple with the PasswordsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsScore

`func (o *PwmUserById) SetPasswordsScore(v float32)`

SetPasswordsScore sets PasswordsScore field to given value.

### HasPasswordsScore

`func (o *PwmUserById) HasPasswordsScore() bool`

HasPasswordsScore returns a boolean if a field has been set.

### GetScoreUpdatedAt

`func (o *PwmUserById) GetScoreUpdatedAt() string`

GetScoreUpdatedAt returns the ScoreUpdatedAt field if non-nil, zero value otherwise.

### GetScoreUpdatedAtOk

`func (o *PwmUserById) GetScoreUpdatedAtOk() (*string, bool)`

GetScoreUpdatedAtOk returns a tuple with the ScoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreUpdatedAt

`func (o *PwmUserById) SetScoreUpdatedAt(v string)`

SetScoreUpdatedAt sets ScoreUpdatedAt field to given value.

### HasScoreUpdatedAt

`func (o *PwmUserById) HasScoreUpdatedAt() bool`

HasScoreUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PwmUserById) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PwmUserById) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PwmUserById) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUsername

`func (o *PwmUserById) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PwmUserById) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PwmUserById) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PwmUserById) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCompromisedPasswords

`func (o *PwmUserById) GetCompromisedPasswords() int32`

GetCompromisedPasswords returns the CompromisedPasswords field if non-nil, zero value otherwise.

### GetCompromisedPasswordsOk

`func (o *PwmUserById) GetCompromisedPasswordsOk() (*int32, bool)`

GetCompromisedPasswordsOk returns a tuple with the CompromisedPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedPasswords

`func (o *PwmUserById) SetCompromisedPasswords(v int32)`

SetCompromisedPasswords sets CompromisedPasswords field to given value.

### HasCompromisedPasswords

`func (o *PwmUserById) HasCompromisedPasswords() bool`

HasCompromisedPasswords returns a boolean if a field has been set.

### GetOldPasswords

`func (o *PwmUserById) GetOldPasswords() int32`

GetOldPasswords returns the OldPasswords field if non-nil, zero value otherwise.

### GetOldPasswordsOk

`func (o *PwmUserById) GetOldPasswordsOk() (*int32, bool)`

GetOldPasswordsOk returns a tuple with the OldPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPasswords

`func (o *PwmUserById) SetOldPasswords(v int32)`

SetOldPasswords sets OldPasswords field to given value.

### HasOldPasswords

`func (o *PwmUserById) HasOldPasswords() bool`

HasOldPasswords returns a boolean if a field has been set.

### GetReusedPasswords

`func (o *PwmUserById) GetReusedPasswords() int32`

GetReusedPasswords returns the ReusedPasswords field if non-nil, zero value otherwise.

### GetReusedPasswordsOk

`func (o *PwmUserById) GetReusedPasswordsOk() (*int32, bool)`

GetReusedPasswordsOk returns a tuple with the ReusedPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusedPasswords

`func (o *PwmUserById) SetReusedPasswords(v int32)`

SetReusedPasswords sets ReusedPasswords field to given value.

### HasReusedPasswords

`func (o *PwmUserById) HasReusedPasswords() bool`

HasReusedPasswords returns a boolean if a field has been set.

### GetWeakPasswords

`func (o *PwmUserById) GetWeakPasswords() int32`

GetWeakPasswords returns the WeakPasswords field if non-nil, zero value otherwise.

### GetWeakPasswordsOk

`func (o *PwmUserById) GetWeakPasswordsOk() (*int32, bool)`

GetWeakPasswordsOk returns a tuple with the WeakPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakPasswords

`func (o *PwmUserById) SetWeakPasswords(v int32)`

SetWeakPasswords sets WeakPasswords field to given value.

### HasWeakPasswords

`func (o *PwmUserById) HasWeakPasswords() bool`

HasWeakPasswords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PwmAllUsersResultsInner

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

### NewPwmAllUsersResultsInner

`func NewPwmAllUsersResultsInner(email string, id string, name string, status string, ) *PwmAllUsersResultsInner`

NewPwmAllUsersResultsInner instantiates a new PwmAllUsersResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPwmAllUsersResultsInnerWithDefaults

`func NewPwmAllUsersResultsInnerWithDefaults() *PwmAllUsersResultsInner`

NewPwmAllUsersResultsInnerWithDefaults instantiates a new PwmAllUsersResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *PwmAllUsersResultsInner) GetApps() []AppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PwmAllUsersResultsInner) GetAppsOk() (*[]AppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PwmAllUsersResultsInner) SetApps(v []AppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PwmAllUsersResultsInner) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetEmail

`func (o *PwmAllUsersResultsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PwmAllUsersResultsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PwmAllUsersResultsInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalId

`func (o *PwmAllUsersResultsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PwmAllUsersResultsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PwmAllUsersResultsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PwmAllUsersResultsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGroups

`func (o *PwmAllUsersResultsInner) GetGroups() []GroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PwmAllUsersResultsInner) GetGroupsOk() (*[]GroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PwmAllUsersResultsInner) SetGroups(v []GroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PwmAllUsersResultsInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *PwmAllUsersResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PwmAllUsersResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PwmAllUsersResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetItemsCount

`func (o *PwmAllUsersResultsInner) GetItemsCount() int32`

GetItemsCount returns the ItemsCount field if non-nil, zero value otherwise.

### GetItemsCountOk

`func (o *PwmAllUsersResultsInner) GetItemsCountOk() (*int32, bool)`

GetItemsCountOk returns a tuple with the ItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsCount

`func (o *PwmAllUsersResultsInner) SetItemsCount(v int32)`

SetItemsCount sets ItemsCount field to given value.

### HasItemsCount

`func (o *PwmAllUsersResultsInner) HasItemsCount() bool`

HasItemsCount returns a boolean if a field has been set.

### GetName

`func (o *PwmAllUsersResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PwmAllUsersResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PwmAllUsersResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordsCount

`func (o *PwmAllUsersResultsInner) GetPasswordsCount() int32`

GetPasswordsCount returns the PasswordsCount field if non-nil, zero value otherwise.

### GetPasswordsCountOk

`func (o *PwmAllUsersResultsInner) GetPasswordsCountOk() (*int32, bool)`

GetPasswordsCountOk returns a tuple with the PasswordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsCount

`func (o *PwmAllUsersResultsInner) SetPasswordsCount(v int32)`

SetPasswordsCount sets PasswordsCount field to given value.

### HasPasswordsCount

`func (o *PwmAllUsersResultsInner) HasPasswordsCount() bool`

HasPasswordsCount returns a boolean if a field has been set.

### GetPasswordsScore

`func (o *PwmAllUsersResultsInner) GetPasswordsScore() float32`

GetPasswordsScore returns the PasswordsScore field if non-nil, zero value otherwise.

### GetPasswordsScoreOk

`func (o *PwmAllUsersResultsInner) GetPasswordsScoreOk() (*float32, bool)`

GetPasswordsScoreOk returns a tuple with the PasswordsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsScore

`func (o *PwmAllUsersResultsInner) SetPasswordsScore(v float32)`

SetPasswordsScore sets PasswordsScore field to given value.

### HasPasswordsScore

`func (o *PwmAllUsersResultsInner) HasPasswordsScore() bool`

HasPasswordsScore returns a boolean if a field has been set.

### GetScoreUpdatedAt

`func (o *PwmAllUsersResultsInner) GetScoreUpdatedAt() string`

GetScoreUpdatedAt returns the ScoreUpdatedAt field if non-nil, zero value otherwise.

### GetScoreUpdatedAtOk

`func (o *PwmAllUsersResultsInner) GetScoreUpdatedAtOk() (*string, bool)`

GetScoreUpdatedAtOk returns a tuple with the ScoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreUpdatedAt

`func (o *PwmAllUsersResultsInner) SetScoreUpdatedAt(v string)`

SetScoreUpdatedAt sets ScoreUpdatedAt field to given value.

### HasScoreUpdatedAt

`func (o *PwmAllUsersResultsInner) HasScoreUpdatedAt() bool`

HasScoreUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PwmAllUsersResultsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PwmAllUsersResultsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PwmAllUsersResultsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUsername

`func (o *PwmAllUsersResultsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PwmAllUsersResultsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PwmAllUsersResultsInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PwmAllUsersResultsInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCompromisedPasswords

`func (o *PwmAllUsersResultsInner) GetCompromisedPasswords() int32`

GetCompromisedPasswords returns the CompromisedPasswords field if non-nil, zero value otherwise.

### GetCompromisedPasswordsOk

`func (o *PwmAllUsersResultsInner) GetCompromisedPasswordsOk() (*int32, bool)`

GetCompromisedPasswordsOk returns a tuple with the CompromisedPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedPasswords

`func (o *PwmAllUsersResultsInner) SetCompromisedPasswords(v int32)`

SetCompromisedPasswords sets CompromisedPasswords field to given value.

### HasCompromisedPasswords

`func (o *PwmAllUsersResultsInner) HasCompromisedPasswords() bool`

HasCompromisedPasswords returns a boolean if a field has been set.

### GetOldPasswords

`func (o *PwmAllUsersResultsInner) GetOldPasswords() int32`

GetOldPasswords returns the OldPasswords field if non-nil, zero value otherwise.

### GetOldPasswordsOk

`func (o *PwmAllUsersResultsInner) GetOldPasswordsOk() (*int32, bool)`

GetOldPasswordsOk returns a tuple with the OldPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPasswords

`func (o *PwmAllUsersResultsInner) SetOldPasswords(v int32)`

SetOldPasswords sets OldPasswords field to given value.

### HasOldPasswords

`func (o *PwmAllUsersResultsInner) HasOldPasswords() bool`

HasOldPasswords returns a boolean if a field has been set.

### GetReusedPasswords

`func (o *PwmAllUsersResultsInner) GetReusedPasswords() int32`

GetReusedPasswords returns the ReusedPasswords field if non-nil, zero value otherwise.

### GetReusedPasswordsOk

`func (o *PwmAllUsersResultsInner) GetReusedPasswordsOk() (*int32, bool)`

GetReusedPasswordsOk returns a tuple with the ReusedPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusedPasswords

`func (o *PwmAllUsersResultsInner) SetReusedPasswords(v int32)`

SetReusedPasswords sets ReusedPasswords field to given value.

### HasReusedPasswords

`func (o *PwmAllUsersResultsInner) HasReusedPasswords() bool`

HasReusedPasswords returns a boolean if a field has been set.

### GetWeakPasswords

`func (o *PwmAllUsersResultsInner) GetWeakPasswords() int32`

GetWeakPasswords returns the WeakPasswords field if non-nil, zero value otherwise.

### GetWeakPasswordsOk

`func (o *PwmAllUsersResultsInner) GetWeakPasswordsOk() (*int32, bool)`

GetWeakPasswordsOk returns a tuple with the WeakPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakPasswords

`func (o *PwmAllUsersResultsInner) SetWeakPasswords(v int32)`

SetWeakPasswords sets WeakPasswords field to given value.

### HasWeakPasswords

`func (o *PwmAllUsersResultsInner) HasWeakPasswords() bool`

HasWeakPasswords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



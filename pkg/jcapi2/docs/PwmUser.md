# PwmUser

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

## Methods

### NewPwmUser

`func NewPwmUser(email string, id string, name string, status string, ) *PwmUser`

NewPwmUser instantiates a new PwmUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPwmUserWithDefaults

`func NewPwmUserWithDefaults() *PwmUser`

NewPwmUserWithDefaults instantiates a new PwmUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *PwmUser) GetApps() []AppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PwmUser) GetAppsOk() (*[]AppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PwmUser) SetApps(v []AppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PwmUser) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetEmail

`func (o *PwmUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PwmUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PwmUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalId

`func (o *PwmUser) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PwmUser) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PwmUser) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PwmUser) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGroups

`func (o *PwmUser) GetGroups() []GroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PwmUser) GetGroupsOk() (*[]GroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PwmUser) SetGroups(v []GroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PwmUser) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *PwmUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PwmUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PwmUser) SetId(v string)`

SetId sets Id field to given value.


### GetItemsCount

`func (o *PwmUser) GetItemsCount() int32`

GetItemsCount returns the ItemsCount field if non-nil, zero value otherwise.

### GetItemsCountOk

`func (o *PwmUser) GetItemsCountOk() (*int32, bool)`

GetItemsCountOk returns a tuple with the ItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsCount

`func (o *PwmUser) SetItemsCount(v int32)`

SetItemsCount sets ItemsCount field to given value.

### HasItemsCount

`func (o *PwmUser) HasItemsCount() bool`

HasItemsCount returns a boolean if a field has been set.

### GetName

`func (o *PwmUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PwmUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PwmUser) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordsCount

`func (o *PwmUser) GetPasswordsCount() int32`

GetPasswordsCount returns the PasswordsCount field if non-nil, zero value otherwise.

### GetPasswordsCountOk

`func (o *PwmUser) GetPasswordsCountOk() (*int32, bool)`

GetPasswordsCountOk returns a tuple with the PasswordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsCount

`func (o *PwmUser) SetPasswordsCount(v int32)`

SetPasswordsCount sets PasswordsCount field to given value.

### HasPasswordsCount

`func (o *PwmUser) HasPasswordsCount() bool`

HasPasswordsCount returns a boolean if a field has been set.

### GetPasswordsScore

`func (o *PwmUser) GetPasswordsScore() float32`

GetPasswordsScore returns the PasswordsScore field if non-nil, zero value otherwise.

### GetPasswordsScoreOk

`func (o *PwmUser) GetPasswordsScoreOk() (*float32, bool)`

GetPasswordsScoreOk returns a tuple with the PasswordsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsScore

`func (o *PwmUser) SetPasswordsScore(v float32)`

SetPasswordsScore sets PasswordsScore field to given value.

### HasPasswordsScore

`func (o *PwmUser) HasPasswordsScore() bool`

HasPasswordsScore returns a boolean if a field has been set.

### GetScoreUpdatedAt

`func (o *PwmUser) GetScoreUpdatedAt() string`

GetScoreUpdatedAt returns the ScoreUpdatedAt field if non-nil, zero value otherwise.

### GetScoreUpdatedAtOk

`func (o *PwmUser) GetScoreUpdatedAtOk() (*string, bool)`

GetScoreUpdatedAtOk returns a tuple with the ScoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreUpdatedAt

`func (o *PwmUser) SetScoreUpdatedAt(v string)`

SetScoreUpdatedAt sets ScoreUpdatedAt field to given value.

### HasScoreUpdatedAt

`func (o *PwmUser) HasScoreUpdatedAt() bool`

HasScoreUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PwmUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PwmUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PwmUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUsername

`func (o *PwmUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PwmUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PwmUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PwmUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



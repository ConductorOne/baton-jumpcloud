# PwmUserSharedFoldersResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevelId** | Pointer to **string** |  | [optional] 
**AccessLevelName** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** |  | 
**Id** | **string** |  | 
**ItemsInFolder** | **int32** |  | 
**Name** | **string** |  | 
**PasswordsScore** | Pointer to **float32** |  | [optional] 
**ScoreUpdatedAt** | Pointer to **string** |  | [optional] 
**UsersWithAccess** | **int32** |  | 

## Methods

### NewPwmUserSharedFoldersResultsInner

`func NewPwmUserSharedFoldersResultsInner(createdAt string, id string, itemsInFolder int32, name string, usersWithAccess int32, ) *PwmUserSharedFoldersResultsInner`

NewPwmUserSharedFoldersResultsInner instantiates a new PwmUserSharedFoldersResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPwmUserSharedFoldersResultsInnerWithDefaults

`func NewPwmUserSharedFoldersResultsInnerWithDefaults() *PwmUserSharedFoldersResultsInner`

NewPwmUserSharedFoldersResultsInnerWithDefaults instantiates a new PwmUserSharedFoldersResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevelId

`func (o *PwmUserSharedFoldersResultsInner) GetAccessLevelId() string`

GetAccessLevelId returns the AccessLevelId field if non-nil, zero value otherwise.

### GetAccessLevelIdOk

`func (o *PwmUserSharedFoldersResultsInner) GetAccessLevelIdOk() (*string, bool)`

GetAccessLevelIdOk returns a tuple with the AccessLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelId

`func (o *PwmUserSharedFoldersResultsInner) SetAccessLevelId(v string)`

SetAccessLevelId sets AccessLevelId field to given value.

### HasAccessLevelId

`func (o *PwmUserSharedFoldersResultsInner) HasAccessLevelId() bool`

HasAccessLevelId returns a boolean if a field has been set.

### GetAccessLevelName

`func (o *PwmUserSharedFoldersResultsInner) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *PwmUserSharedFoldersResultsInner) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *PwmUserSharedFoldersResultsInner) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.

### HasAccessLevelName

`func (o *PwmUserSharedFoldersResultsInner) HasAccessLevelName() bool`

HasAccessLevelName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PwmUserSharedFoldersResultsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PwmUserSharedFoldersResultsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PwmUserSharedFoldersResultsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *PwmUserSharedFoldersResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PwmUserSharedFoldersResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PwmUserSharedFoldersResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetItemsInFolder

`func (o *PwmUserSharedFoldersResultsInner) GetItemsInFolder() int32`

GetItemsInFolder returns the ItemsInFolder field if non-nil, zero value otherwise.

### GetItemsInFolderOk

`func (o *PwmUserSharedFoldersResultsInner) GetItemsInFolderOk() (*int32, bool)`

GetItemsInFolderOk returns a tuple with the ItemsInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsInFolder

`func (o *PwmUserSharedFoldersResultsInner) SetItemsInFolder(v int32)`

SetItemsInFolder sets ItemsInFolder field to given value.


### GetName

`func (o *PwmUserSharedFoldersResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PwmUserSharedFoldersResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PwmUserSharedFoldersResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordsScore

`func (o *PwmUserSharedFoldersResultsInner) GetPasswordsScore() float32`

GetPasswordsScore returns the PasswordsScore field if non-nil, zero value otherwise.

### GetPasswordsScoreOk

`func (o *PwmUserSharedFoldersResultsInner) GetPasswordsScoreOk() (*float32, bool)`

GetPasswordsScoreOk returns a tuple with the PasswordsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsScore

`func (o *PwmUserSharedFoldersResultsInner) SetPasswordsScore(v float32)`

SetPasswordsScore sets PasswordsScore field to given value.

### HasPasswordsScore

`func (o *PwmUserSharedFoldersResultsInner) HasPasswordsScore() bool`

HasPasswordsScore returns a boolean if a field has been set.

### GetScoreUpdatedAt

`func (o *PwmUserSharedFoldersResultsInner) GetScoreUpdatedAt() string`

GetScoreUpdatedAt returns the ScoreUpdatedAt field if non-nil, zero value otherwise.

### GetScoreUpdatedAtOk

`func (o *PwmUserSharedFoldersResultsInner) GetScoreUpdatedAtOk() (*string, bool)`

GetScoreUpdatedAtOk returns a tuple with the ScoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreUpdatedAt

`func (o *PwmUserSharedFoldersResultsInner) SetScoreUpdatedAt(v string)`

SetScoreUpdatedAt sets ScoreUpdatedAt field to given value.

### HasScoreUpdatedAt

`func (o *PwmUserSharedFoldersResultsInner) HasScoreUpdatedAt() bool`

HasScoreUpdatedAt returns a boolean if a field has been set.

### GetUsersWithAccess

`func (o *PwmUserSharedFoldersResultsInner) GetUsersWithAccess() int32`

GetUsersWithAccess returns the UsersWithAccess field if non-nil, zero value otherwise.

### GetUsersWithAccessOk

`func (o *PwmUserSharedFoldersResultsInner) GetUsersWithAccessOk() (*int32, bool)`

GetUsersWithAccessOk returns a tuple with the UsersWithAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersWithAccess

`func (o *PwmUserSharedFoldersResultsInner) SetUsersWithAccess(v int32)`

SetUsersWithAccess sets UsersWithAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



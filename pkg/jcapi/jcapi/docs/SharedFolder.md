# SharedFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**ItemsInFolder** | **int32** |  | 
**Name** | **string** |  | 
**PasswordsCount** | Pointer to **int32** |  | [optional] 
**PasswordsScore** | Pointer to **float32** |  | [optional] 
**ScoreUpdatedAt** | Pointer to **string** |  | [optional] 
**UsersWithAccess** | **int32** |  | 
**Uuid** | **string** |  | 

## Methods

### NewSharedFolder

`func NewSharedFolder(createdAt string, itemsInFolder int32, name string, usersWithAccess int32, uuid string, ) *SharedFolder`

NewSharedFolder instantiates a new SharedFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedFolderWithDefaults

`func NewSharedFolderWithDefaults() *SharedFolder`

NewSharedFolderWithDefaults instantiates a new SharedFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SharedFolder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SharedFolder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SharedFolder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetItemsInFolder

`func (o *SharedFolder) GetItemsInFolder() int32`

GetItemsInFolder returns the ItemsInFolder field if non-nil, zero value otherwise.

### GetItemsInFolderOk

`func (o *SharedFolder) GetItemsInFolderOk() (*int32, bool)`

GetItemsInFolderOk returns a tuple with the ItemsInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsInFolder

`func (o *SharedFolder) SetItemsInFolder(v int32)`

SetItemsInFolder sets ItemsInFolder field to given value.


### GetName

`func (o *SharedFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedFolder) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordsCount

`func (o *SharedFolder) GetPasswordsCount() int32`

GetPasswordsCount returns the PasswordsCount field if non-nil, zero value otherwise.

### GetPasswordsCountOk

`func (o *SharedFolder) GetPasswordsCountOk() (*int32, bool)`

GetPasswordsCountOk returns a tuple with the PasswordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsCount

`func (o *SharedFolder) SetPasswordsCount(v int32)`

SetPasswordsCount sets PasswordsCount field to given value.

### HasPasswordsCount

`func (o *SharedFolder) HasPasswordsCount() bool`

HasPasswordsCount returns a boolean if a field has been set.

### GetPasswordsScore

`func (o *SharedFolder) GetPasswordsScore() float32`

GetPasswordsScore returns the PasswordsScore field if non-nil, zero value otherwise.

### GetPasswordsScoreOk

`func (o *SharedFolder) GetPasswordsScoreOk() (*float32, bool)`

GetPasswordsScoreOk returns a tuple with the PasswordsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsScore

`func (o *SharedFolder) SetPasswordsScore(v float32)`

SetPasswordsScore sets PasswordsScore field to given value.

### HasPasswordsScore

`func (o *SharedFolder) HasPasswordsScore() bool`

HasPasswordsScore returns a boolean if a field has been set.

### GetScoreUpdatedAt

`func (o *SharedFolder) GetScoreUpdatedAt() string`

GetScoreUpdatedAt returns the ScoreUpdatedAt field if non-nil, zero value otherwise.

### GetScoreUpdatedAtOk

`func (o *SharedFolder) GetScoreUpdatedAtOk() (*string, bool)`

GetScoreUpdatedAtOk returns a tuple with the ScoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreUpdatedAt

`func (o *SharedFolder) SetScoreUpdatedAt(v string)`

SetScoreUpdatedAt sets ScoreUpdatedAt field to given value.

### HasScoreUpdatedAt

`func (o *SharedFolder) HasScoreUpdatedAt() bool`

HasScoreUpdatedAt returns a boolean if a field has been set.

### GetUsersWithAccess

`func (o *SharedFolder) GetUsersWithAccess() int32`

GetUsersWithAccess returns the UsersWithAccess field if non-nil, zero value otherwise.

### GetUsersWithAccessOk

`func (o *SharedFolder) GetUsersWithAccessOk() (*int32, bool)`

GetUsersWithAccessOk returns a tuple with the UsersWithAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersWithAccess

`func (o *SharedFolder) SetUsersWithAccess(v int32)`

SetUsersWithAccess sets UsersWithAccess field to given value.


### GetUuid

`func (o *SharedFolder) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SharedFolder) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SharedFolder) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



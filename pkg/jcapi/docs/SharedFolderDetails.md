# SharedFolderDetails

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
**CompromisedPasswords** | Pointer to **int32** |  | [optional] 
**OldPasswords** | Pointer to **int32** |  | [optional] 
**ReusedPasswords** | Pointer to **int32** |  | [optional] 
**WeakPasswords** | Pointer to **int32** |  | [optional] 

## Methods

### NewSharedFolderDetails

`func NewSharedFolderDetails(createdAt string, itemsInFolder int32, name string, usersWithAccess int32, uuid string, ) *SharedFolderDetails`

NewSharedFolderDetails instantiates a new SharedFolderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedFolderDetailsWithDefaults

`func NewSharedFolderDetailsWithDefaults() *SharedFolderDetails`

NewSharedFolderDetailsWithDefaults instantiates a new SharedFolderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SharedFolderDetails) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SharedFolderDetails) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SharedFolderDetails) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetItemsInFolder

`func (o *SharedFolderDetails) GetItemsInFolder() int32`

GetItemsInFolder returns the ItemsInFolder field if non-nil, zero value otherwise.

### GetItemsInFolderOk

`func (o *SharedFolderDetails) GetItemsInFolderOk() (*int32, bool)`

GetItemsInFolderOk returns a tuple with the ItemsInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsInFolder

`func (o *SharedFolderDetails) SetItemsInFolder(v int32)`

SetItemsInFolder sets ItemsInFolder field to given value.


### GetName

`func (o *SharedFolderDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedFolderDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedFolderDetails) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordsCount

`func (o *SharedFolderDetails) GetPasswordsCount() int32`

GetPasswordsCount returns the PasswordsCount field if non-nil, zero value otherwise.

### GetPasswordsCountOk

`func (o *SharedFolderDetails) GetPasswordsCountOk() (*int32, bool)`

GetPasswordsCountOk returns a tuple with the PasswordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsCount

`func (o *SharedFolderDetails) SetPasswordsCount(v int32)`

SetPasswordsCount sets PasswordsCount field to given value.

### HasPasswordsCount

`func (o *SharedFolderDetails) HasPasswordsCount() bool`

HasPasswordsCount returns a boolean if a field has been set.

### GetPasswordsScore

`func (o *SharedFolderDetails) GetPasswordsScore() float32`

GetPasswordsScore returns the PasswordsScore field if non-nil, zero value otherwise.

### GetPasswordsScoreOk

`func (o *SharedFolderDetails) GetPasswordsScoreOk() (*float32, bool)`

GetPasswordsScoreOk returns a tuple with the PasswordsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsScore

`func (o *SharedFolderDetails) SetPasswordsScore(v float32)`

SetPasswordsScore sets PasswordsScore field to given value.

### HasPasswordsScore

`func (o *SharedFolderDetails) HasPasswordsScore() bool`

HasPasswordsScore returns a boolean if a field has been set.

### GetScoreUpdatedAt

`func (o *SharedFolderDetails) GetScoreUpdatedAt() string`

GetScoreUpdatedAt returns the ScoreUpdatedAt field if non-nil, zero value otherwise.

### GetScoreUpdatedAtOk

`func (o *SharedFolderDetails) GetScoreUpdatedAtOk() (*string, bool)`

GetScoreUpdatedAtOk returns a tuple with the ScoreUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreUpdatedAt

`func (o *SharedFolderDetails) SetScoreUpdatedAt(v string)`

SetScoreUpdatedAt sets ScoreUpdatedAt field to given value.

### HasScoreUpdatedAt

`func (o *SharedFolderDetails) HasScoreUpdatedAt() bool`

HasScoreUpdatedAt returns a boolean if a field has been set.

### GetUsersWithAccess

`func (o *SharedFolderDetails) GetUsersWithAccess() int32`

GetUsersWithAccess returns the UsersWithAccess field if non-nil, zero value otherwise.

### GetUsersWithAccessOk

`func (o *SharedFolderDetails) GetUsersWithAccessOk() (*int32, bool)`

GetUsersWithAccessOk returns a tuple with the UsersWithAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersWithAccess

`func (o *SharedFolderDetails) SetUsersWithAccess(v int32)`

SetUsersWithAccess sets UsersWithAccess field to given value.


### GetUuid

`func (o *SharedFolderDetails) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SharedFolderDetails) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SharedFolderDetails) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCompromisedPasswords

`func (o *SharedFolderDetails) GetCompromisedPasswords() int32`

GetCompromisedPasswords returns the CompromisedPasswords field if non-nil, zero value otherwise.

### GetCompromisedPasswordsOk

`func (o *SharedFolderDetails) GetCompromisedPasswordsOk() (*int32, bool)`

GetCompromisedPasswordsOk returns a tuple with the CompromisedPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedPasswords

`func (o *SharedFolderDetails) SetCompromisedPasswords(v int32)`

SetCompromisedPasswords sets CompromisedPasswords field to given value.

### HasCompromisedPasswords

`func (o *SharedFolderDetails) HasCompromisedPasswords() bool`

HasCompromisedPasswords returns a boolean if a field has been set.

### GetOldPasswords

`func (o *SharedFolderDetails) GetOldPasswords() int32`

GetOldPasswords returns the OldPasswords field if non-nil, zero value otherwise.

### GetOldPasswordsOk

`func (o *SharedFolderDetails) GetOldPasswordsOk() (*int32, bool)`

GetOldPasswordsOk returns a tuple with the OldPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPasswords

`func (o *SharedFolderDetails) SetOldPasswords(v int32)`

SetOldPasswords sets OldPasswords field to given value.

### HasOldPasswords

`func (o *SharedFolderDetails) HasOldPasswords() bool`

HasOldPasswords returns a boolean if a field has been set.

### GetReusedPasswords

`func (o *SharedFolderDetails) GetReusedPasswords() int32`

GetReusedPasswords returns the ReusedPasswords field if non-nil, zero value otherwise.

### GetReusedPasswordsOk

`func (o *SharedFolderDetails) GetReusedPasswordsOk() (*int32, bool)`

GetReusedPasswordsOk returns a tuple with the ReusedPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusedPasswords

`func (o *SharedFolderDetails) SetReusedPasswords(v int32)`

SetReusedPasswords sets ReusedPasswords field to given value.

### HasReusedPasswords

`func (o *SharedFolderDetails) HasReusedPasswords() bool`

HasReusedPasswords returns a boolean if a field has been set.

### GetWeakPasswords

`func (o *SharedFolderDetails) GetWeakPasswords() int32`

GetWeakPasswords returns the WeakPasswords field if non-nil, zero value otherwise.

### GetWeakPasswordsOk

`func (o *SharedFolderDetails) GetWeakPasswordsOk() (*int32, bool)`

GetWeakPasswordsOk returns a tuple with the WeakPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakPasswords

`func (o *SharedFolderDetails) SetWeakPasswords(v int32)`

SetWeakPasswords sets WeakPasswords field to given value.

### HasWeakPasswords

`func (o *SharedFolderDetails) HasWeakPasswords() bool`

HasWeakPasswords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



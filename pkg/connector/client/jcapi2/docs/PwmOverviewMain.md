# PwmOverviewMain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompromisedPasswords** | Pointer to **int32** |  | [optional] 
**Devices** | [**[]PwmOverviewMainDevicesInner**](PwmOverviewMainDevicesInner.md) |  | 
**EnrolledGroups** | Pointer to **int32** |  | [optional] 
**OldPasswords** | Pointer to **int32** |  | [optional] 
**PasswordsCount** | Pointer to **int32** |  | [optional] 
**PasswordsScore** | Pointer to **float32** |  | [optional] 
**PendingInvites** | **int32** |  | 
**SharedFolders** | **int32** |  | 
**TotalUsers** | **int32** |  | 
**WeakPasswords** | Pointer to **int32** |  | [optional] 

## Methods

### NewPwmOverviewMain

`func NewPwmOverviewMain(devices []PwmOverviewMainDevicesInner, pendingInvites int32, sharedFolders int32, totalUsers int32, ) *PwmOverviewMain`

NewPwmOverviewMain instantiates a new PwmOverviewMain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPwmOverviewMainWithDefaults

`func NewPwmOverviewMainWithDefaults() *PwmOverviewMain`

NewPwmOverviewMainWithDefaults instantiates a new PwmOverviewMain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompromisedPasswords

`func (o *PwmOverviewMain) GetCompromisedPasswords() int32`

GetCompromisedPasswords returns the CompromisedPasswords field if non-nil, zero value otherwise.

### GetCompromisedPasswordsOk

`func (o *PwmOverviewMain) GetCompromisedPasswordsOk() (*int32, bool)`

GetCompromisedPasswordsOk returns a tuple with the CompromisedPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedPasswords

`func (o *PwmOverviewMain) SetCompromisedPasswords(v int32)`

SetCompromisedPasswords sets CompromisedPasswords field to given value.

### HasCompromisedPasswords

`func (o *PwmOverviewMain) HasCompromisedPasswords() bool`

HasCompromisedPasswords returns a boolean if a field has been set.

### GetDevices

`func (o *PwmOverviewMain) GetDevices() []PwmOverviewMainDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *PwmOverviewMain) GetDevicesOk() (*[]PwmOverviewMainDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *PwmOverviewMain) SetDevices(v []PwmOverviewMainDevicesInner)`

SetDevices sets Devices field to given value.


### GetEnrolledGroups

`func (o *PwmOverviewMain) GetEnrolledGroups() int32`

GetEnrolledGroups returns the EnrolledGroups field if non-nil, zero value otherwise.

### GetEnrolledGroupsOk

`func (o *PwmOverviewMain) GetEnrolledGroupsOk() (*int32, bool)`

GetEnrolledGroupsOk returns a tuple with the EnrolledGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledGroups

`func (o *PwmOverviewMain) SetEnrolledGroups(v int32)`

SetEnrolledGroups sets EnrolledGroups field to given value.

### HasEnrolledGroups

`func (o *PwmOverviewMain) HasEnrolledGroups() bool`

HasEnrolledGroups returns a boolean if a field has been set.

### GetOldPasswords

`func (o *PwmOverviewMain) GetOldPasswords() int32`

GetOldPasswords returns the OldPasswords field if non-nil, zero value otherwise.

### GetOldPasswordsOk

`func (o *PwmOverviewMain) GetOldPasswordsOk() (*int32, bool)`

GetOldPasswordsOk returns a tuple with the OldPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPasswords

`func (o *PwmOverviewMain) SetOldPasswords(v int32)`

SetOldPasswords sets OldPasswords field to given value.

### HasOldPasswords

`func (o *PwmOverviewMain) HasOldPasswords() bool`

HasOldPasswords returns a boolean if a field has been set.

### GetPasswordsCount

`func (o *PwmOverviewMain) GetPasswordsCount() int32`

GetPasswordsCount returns the PasswordsCount field if non-nil, zero value otherwise.

### GetPasswordsCountOk

`func (o *PwmOverviewMain) GetPasswordsCountOk() (*int32, bool)`

GetPasswordsCountOk returns a tuple with the PasswordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsCount

`func (o *PwmOverviewMain) SetPasswordsCount(v int32)`

SetPasswordsCount sets PasswordsCount field to given value.

### HasPasswordsCount

`func (o *PwmOverviewMain) HasPasswordsCount() bool`

HasPasswordsCount returns a boolean if a field has been set.

### GetPasswordsScore

`func (o *PwmOverviewMain) GetPasswordsScore() float32`

GetPasswordsScore returns the PasswordsScore field if non-nil, zero value otherwise.

### GetPasswordsScoreOk

`func (o *PwmOverviewMain) GetPasswordsScoreOk() (*float32, bool)`

GetPasswordsScoreOk returns a tuple with the PasswordsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordsScore

`func (o *PwmOverviewMain) SetPasswordsScore(v float32)`

SetPasswordsScore sets PasswordsScore field to given value.

### HasPasswordsScore

`func (o *PwmOverviewMain) HasPasswordsScore() bool`

HasPasswordsScore returns a boolean if a field has been set.

### GetPendingInvites

`func (o *PwmOverviewMain) GetPendingInvites() int32`

GetPendingInvites returns the PendingInvites field if non-nil, zero value otherwise.

### GetPendingInvitesOk

`func (o *PwmOverviewMain) GetPendingInvitesOk() (*int32, bool)`

GetPendingInvitesOk returns a tuple with the PendingInvites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingInvites

`func (o *PwmOverviewMain) SetPendingInvites(v int32)`

SetPendingInvites sets PendingInvites field to given value.


### GetSharedFolders

`func (o *PwmOverviewMain) GetSharedFolders() int32`

GetSharedFolders returns the SharedFolders field if non-nil, zero value otherwise.

### GetSharedFoldersOk

`func (o *PwmOverviewMain) GetSharedFoldersOk() (*int32, bool)`

GetSharedFoldersOk returns a tuple with the SharedFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedFolders

`func (o *PwmOverviewMain) SetSharedFolders(v int32)`

SetSharedFolders sets SharedFolders field to given value.


### GetTotalUsers

`func (o *PwmOverviewMain) GetTotalUsers() int32`

GetTotalUsers returns the TotalUsers field if non-nil, zero value otherwise.

### GetTotalUsersOk

`func (o *PwmOverviewMain) GetTotalUsersOk() (*int32, bool)`

GetTotalUsersOk returns a tuple with the TotalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers

`func (o *PwmOverviewMain) SetTotalUsers(v int32)`

SetTotalUsers sets TotalUsers field to given value.


### GetWeakPasswords

`func (o *PwmOverviewMain) GetWeakPasswords() int32`

GetWeakPasswords returns the WeakPasswords field if non-nil, zero value otherwise.

### GetWeakPasswordsOk

`func (o *PwmOverviewMain) GetWeakPasswordsOk() (*int32, bool)`

GetWeakPasswordsOk returns a tuple with the WeakPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakPasswords

`func (o *PwmOverviewMain) SetWeakPasswords(v int32)`

SetWeakPasswords sets WeakPasswords field to given value.

### HasWeakPasswords

`func (o *PwmOverviewMain) HasWeakPasswords() bool`

HasWeakPasswords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



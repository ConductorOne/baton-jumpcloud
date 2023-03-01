# SystemInsightsUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdManaged** | Pointer to **bool** | Indicates this account belongs to a AD-managed user | [optional] 
**Admin** | Pointer to **bool** | Indicates this account has local administrator privileges | [optional] 
**CollectionTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Directory** | Pointer to **string** |  | [optional] 
**Gid** | Pointer to **string** |  | [optional] 
**GidSigned** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **string** | A Unix timestamp showing the last time this user logged in | [optional] 
**Managed** | Pointer to **bool** | Indicates this account belongs to a JumpCloud-managed user | [optional] 
**RealUser** | Pointer to **bool** | Indicates this account represents an interactive user account vs. a system or daemon account | [optional] 
**Shell** | Pointer to **string** |  | [optional] 
**Suspended** | Pointer to **bool** | Indicates this account is suspended or locked out | [optional] 
**SystemId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**UidSigned** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemInsightsUsers

`func NewSystemInsightsUsers() *SystemInsightsUsers`

NewSystemInsightsUsers instantiates a new SystemInsightsUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInsightsUsersWithDefaults

`func NewSystemInsightsUsersWithDefaults() *SystemInsightsUsers`

NewSystemInsightsUsersWithDefaults instantiates a new SystemInsightsUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdManaged

`func (o *SystemInsightsUsers) GetAdManaged() bool`

GetAdManaged returns the AdManaged field if non-nil, zero value otherwise.

### GetAdManagedOk

`func (o *SystemInsightsUsers) GetAdManagedOk() (*bool, bool)`

GetAdManagedOk returns a tuple with the AdManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdManaged

`func (o *SystemInsightsUsers) SetAdManaged(v bool)`

SetAdManaged sets AdManaged field to given value.

### HasAdManaged

`func (o *SystemInsightsUsers) HasAdManaged() bool`

HasAdManaged returns a boolean if a field has been set.

### GetAdmin

`func (o *SystemInsightsUsers) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *SystemInsightsUsers) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *SystemInsightsUsers) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *SystemInsightsUsers) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCollectionTime

`func (o *SystemInsightsUsers) GetCollectionTime() string`

GetCollectionTime returns the CollectionTime field if non-nil, zero value otherwise.

### GetCollectionTimeOk

`func (o *SystemInsightsUsers) GetCollectionTimeOk() (*string, bool)`

GetCollectionTimeOk returns a tuple with the CollectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTime

`func (o *SystemInsightsUsers) SetCollectionTime(v string)`

SetCollectionTime sets CollectionTime field to given value.

### HasCollectionTime

`func (o *SystemInsightsUsers) HasCollectionTime() bool`

HasCollectionTime returns a boolean if a field has been set.

### GetDescription

`func (o *SystemInsightsUsers) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemInsightsUsers) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemInsightsUsers) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemInsightsUsers) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectory

`func (o *SystemInsightsUsers) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *SystemInsightsUsers) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *SystemInsightsUsers) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *SystemInsightsUsers) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetGid

`func (o *SystemInsightsUsers) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *SystemInsightsUsers) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *SystemInsightsUsers) SetGid(v string)`

SetGid sets Gid field to given value.

### HasGid

`func (o *SystemInsightsUsers) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetGidSigned

`func (o *SystemInsightsUsers) GetGidSigned() string`

GetGidSigned returns the GidSigned field if non-nil, zero value otherwise.

### GetGidSignedOk

`func (o *SystemInsightsUsers) GetGidSignedOk() (*string, bool)`

GetGidSignedOk returns a tuple with the GidSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidSigned

`func (o *SystemInsightsUsers) SetGidSigned(v string)`

SetGidSigned sets GidSigned field to given value.

### HasGidSigned

`func (o *SystemInsightsUsers) HasGidSigned() bool`

HasGidSigned returns a boolean if a field has been set.

### GetLastLogin

`func (o *SystemInsightsUsers) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *SystemInsightsUsers) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *SystemInsightsUsers) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *SystemInsightsUsers) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetManaged

`func (o *SystemInsightsUsers) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *SystemInsightsUsers) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *SystemInsightsUsers) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *SystemInsightsUsers) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetRealUser

`func (o *SystemInsightsUsers) GetRealUser() bool`

GetRealUser returns the RealUser field if non-nil, zero value otherwise.

### GetRealUserOk

`func (o *SystemInsightsUsers) GetRealUserOk() (*bool, bool)`

GetRealUserOk returns a tuple with the RealUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealUser

`func (o *SystemInsightsUsers) SetRealUser(v bool)`

SetRealUser sets RealUser field to given value.

### HasRealUser

`func (o *SystemInsightsUsers) HasRealUser() bool`

HasRealUser returns a boolean if a field has been set.

### GetShell

`func (o *SystemInsightsUsers) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *SystemInsightsUsers) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *SystemInsightsUsers) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *SystemInsightsUsers) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetSuspended

`func (o *SystemInsightsUsers) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *SystemInsightsUsers) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *SystemInsightsUsers) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *SystemInsightsUsers) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetSystemId

`func (o *SystemInsightsUsers) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *SystemInsightsUsers) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *SystemInsightsUsers) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *SystemInsightsUsers) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetType

`func (o *SystemInsightsUsers) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemInsightsUsers) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemInsightsUsers) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemInsightsUsers) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *SystemInsightsUsers) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SystemInsightsUsers) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SystemInsightsUsers) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SystemInsightsUsers) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUidSigned

`func (o *SystemInsightsUsers) GetUidSigned() string`

GetUidSigned returns the UidSigned field if non-nil, zero value otherwise.

### GetUidSignedOk

`func (o *SystemInsightsUsers) GetUidSignedOk() (*string, bool)`

GetUidSignedOk returns a tuple with the UidSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidSigned

`func (o *SystemInsightsUsers) SetUidSigned(v string)`

SetUidSigned sets UidSigned field to given value.

### HasUidSigned

`func (o *SystemInsightsUsers) HasUidSigned() bool`

HasUidSigned returns a boolean if a field has been set.

### GetUsername

`func (o *SystemInsightsUsers) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SystemInsightsUsers) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SystemInsightsUsers) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SystemInsightsUsers) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *SystemInsightsUsers) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SystemInsightsUsers) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SystemInsightsUsers) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SystemInsightsUsers) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GroupAttributesUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sudo** | Pointer to [**GraphAttributeSudoSudo**](GraphAttributeSudoSudo.md) |  | [optional] 
**LdapGroups** | Pointer to [**[]LdapGroup**](LdapGroup.md) |  | [optional] 
**PosixGroups** | Pointer to [**[]GraphAttributePosixGroupsPosixGroupsInner**](GraphAttributePosixGroupsPosixGroupsInner.md) |  | [optional] 
**Radius** | Pointer to [**GraphAttributeRadiusRadius**](GraphAttributeRadiusRadius.md) |  | [optional] 
**SambaEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewGroupAttributesUserGroup

`func NewGroupAttributesUserGroup() *GroupAttributesUserGroup`

NewGroupAttributesUserGroup instantiates a new GroupAttributesUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupAttributesUserGroupWithDefaults

`func NewGroupAttributesUserGroupWithDefaults() *GroupAttributesUserGroup`

NewGroupAttributesUserGroupWithDefaults instantiates a new GroupAttributesUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSudo

`func (o *GroupAttributesUserGroup) GetSudo() GraphAttributeSudoSudo`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *GroupAttributesUserGroup) GetSudoOk() (*GraphAttributeSudoSudo, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *GroupAttributesUserGroup) SetSudo(v GraphAttributeSudoSudo)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *GroupAttributesUserGroup) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetLdapGroups

`func (o *GroupAttributesUserGroup) GetLdapGroups() []LdapGroup`

GetLdapGroups returns the LdapGroups field if non-nil, zero value otherwise.

### GetLdapGroupsOk

`func (o *GroupAttributesUserGroup) GetLdapGroupsOk() (*[]LdapGroup, bool)`

GetLdapGroupsOk returns a tuple with the LdapGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroups

`func (o *GroupAttributesUserGroup) SetLdapGroups(v []LdapGroup)`

SetLdapGroups sets LdapGroups field to given value.

### HasLdapGroups

`func (o *GroupAttributesUserGroup) HasLdapGroups() bool`

HasLdapGroups returns a boolean if a field has been set.

### GetPosixGroups

`func (o *GroupAttributesUserGroup) GetPosixGroups() []GraphAttributePosixGroupsPosixGroupsInner`

GetPosixGroups returns the PosixGroups field if non-nil, zero value otherwise.

### GetPosixGroupsOk

`func (o *GroupAttributesUserGroup) GetPosixGroupsOk() (*[]GraphAttributePosixGroupsPosixGroupsInner, bool)`

GetPosixGroupsOk returns a tuple with the PosixGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosixGroups

`func (o *GroupAttributesUserGroup) SetPosixGroups(v []GraphAttributePosixGroupsPosixGroupsInner)`

SetPosixGroups sets PosixGroups field to given value.

### HasPosixGroups

`func (o *GroupAttributesUserGroup) HasPosixGroups() bool`

HasPosixGroups returns a boolean if a field has been set.

### GetRadius

`func (o *GroupAttributesUserGroup) GetRadius() GraphAttributeRadiusRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *GroupAttributesUserGroup) GetRadiusOk() (*GraphAttributeRadiusRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *GroupAttributesUserGroup) SetRadius(v GraphAttributeRadiusRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *GroupAttributesUserGroup) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetSambaEnabled

`func (o *GroupAttributesUserGroup) GetSambaEnabled() bool`

GetSambaEnabled returns the SambaEnabled field if non-nil, zero value otherwise.

### GetSambaEnabledOk

`func (o *GroupAttributesUserGroup) GetSambaEnabledOk() (*bool, bool)`

GetSambaEnabledOk returns a tuple with the SambaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSambaEnabled

`func (o *GroupAttributesUserGroup) SetSambaEnabled(v bool)`

SetSambaEnabled sets SambaEnabled field to given value.

### HasSambaEnabled

`func (o *GroupAttributesUserGroup) HasSambaEnabled() bool`

HasSambaEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



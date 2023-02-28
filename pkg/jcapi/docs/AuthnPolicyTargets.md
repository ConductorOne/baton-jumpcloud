# AuthnPolicyTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]AuthnPolicyResourceTarget**](AuthnPolicyResourceTarget.md) |  | [optional] 
**UserAttributes** | Pointer to [**AuthnPolicyUserAttributeTarget**](AuthnPolicyUserAttributeTarget.md) |  | [optional] 
**UserGroups** | Pointer to [**AuthnPolicyUserGroupTarget**](AuthnPolicyUserGroupTarget.md) |  | [optional] 
**Users** | Pointer to [**AuthnPolicyUserTarget**](AuthnPolicyUserTarget.md) |  | [optional] 

## Methods

### NewAuthnPolicyTargets

`func NewAuthnPolicyTargets() *AuthnPolicyTargets`

NewAuthnPolicyTargets instantiates a new AuthnPolicyTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnPolicyTargetsWithDefaults

`func NewAuthnPolicyTargetsWithDefaults() *AuthnPolicyTargets`

NewAuthnPolicyTargetsWithDefaults instantiates a new AuthnPolicyTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *AuthnPolicyTargets) GetResources() []AuthnPolicyResourceTarget`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AuthnPolicyTargets) GetResourcesOk() (*[]AuthnPolicyResourceTarget, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AuthnPolicyTargets) SetResources(v []AuthnPolicyResourceTarget)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AuthnPolicyTargets) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetUserAttributes

`func (o *AuthnPolicyTargets) GetUserAttributes() AuthnPolicyUserAttributeTarget`

GetUserAttributes returns the UserAttributes field if non-nil, zero value otherwise.

### GetUserAttributesOk

`func (o *AuthnPolicyTargets) GetUserAttributesOk() (*AuthnPolicyUserAttributeTarget, bool)`

GetUserAttributesOk returns a tuple with the UserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributes

`func (o *AuthnPolicyTargets) SetUserAttributes(v AuthnPolicyUserAttributeTarget)`

SetUserAttributes sets UserAttributes field to given value.

### HasUserAttributes

`func (o *AuthnPolicyTargets) HasUserAttributes() bool`

HasUserAttributes returns a boolean if a field has been set.

### GetUserGroups

`func (o *AuthnPolicyTargets) GetUserGroups() AuthnPolicyUserGroupTarget`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *AuthnPolicyTargets) GetUserGroupsOk() (*AuthnPolicyUserGroupTarget, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *AuthnPolicyTargets) SetUserGroups(v AuthnPolicyUserGroupTarget)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *AuthnPolicyTargets) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetUsers

`func (o *AuthnPolicyTargets) GetUsers() AuthnPolicyUserTarget`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AuthnPolicyTargets) GetUsersOk() (*AuthnPolicyUserTarget, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AuthnPolicyTargets) SetUsers(v AuthnPolicyUserTarget)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AuthnPolicyTargets) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



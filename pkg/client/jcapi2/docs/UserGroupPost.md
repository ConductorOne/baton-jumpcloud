# UserGroupPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**GroupAttributesUserGroup**](GroupAttributesUserGroup.md) |  | [optional] 
**Description** | Pointer to **string** | Description of a User Group | [optional] 
**Email** | Pointer to **string** | Email address of a User Group | [optional] 
**MemberQuery** | Pointer to [**FilterQuery**](FilterQuery.md) |  | [optional] 
**MemberQueryExemptions** | Pointer to [**[]GraphObject**](GraphObject.md) | Array of GraphObjects exempted from the query | [optional] 
**MemberSuggestionsNotify** | Pointer to **bool** | True if notification emails are to be sent for membership suggestions. | [optional] 
**MembershipAutomated** | Pointer to **bool** | True if membership of this group is automatically updated based on the Member Query and Member Query Exemptions, if configured | [optional] 
**MembershipMethod** | Pointer to [**GroupMembershipMethodType**](GroupMembershipMethodType.md) |  | [optional] 
**Name** | **string** | Display name of a User Group. | 

## Methods

### NewUserGroupPost

`func NewUserGroupPost(name string, ) *UserGroupPost`

NewUserGroupPost instantiates a new UserGroupPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupPostWithDefaults

`func NewUserGroupPostWithDefaults() *UserGroupPost`

NewUserGroupPostWithDefaults instantiates a new UserGroupPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UserGroupPost) GetAttributes() GroupAttributesUserGroup`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserGroupPost) GetAttributesOk() (*GroupAttributesUserGroup, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserGroupPost) SetAttributes(v GroupAttributesUserGroup)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserGroupPost) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *UserGroupPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroupPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroupPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroupPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *UserGroupPost) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserGroupPost) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserGroupPost) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserGroupPost) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMemberQuery

`func (o *UserGroupPost) GetMemberQuery() FilterQuery`

GetMemberQuery returns the MemberQuery field if non-nil, zero value otherwise.

### GetMemberQueryOk

`func (o *UserGroupPost) GetMemberQueryOk() (*FilterQuery, bool)`

GetMemberQueryOk returns a tuple with the MemberQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberQuery

`func (o *UserGroupPost) SetMemberQuery(v FilterQuery)`

SetMemberQuery sets MemberQuery field to given value.

### HasMemberQuery

`func (o *UserGroupPost) HasMemberQuery() bool`

HasMemberQuery returns a boolean if a field has been set.

### GetMemberQueryExemptions

`func (o *UserGroupPost) GetMemberQueryExemptions() []GraphObject`

GetMemberQueryExemptions returns the MemberQueryExemptions field if non-nil, zero value otherwise.

### GetMemberQueryExemptionsOk

`func (o *UserGroupPost) GetMemberQueryExemptionsOk() (*[]GraphObject, bool)`

GetMemberQueryExemptionsOk returns a tuple with the MemberQueryExemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberQueryExemptions

`func (o *UserGroupPost) SetMemberQueryExemptions(v []GraphObject)`

SetMemberQueryExemptions sets MemberQueryExemptions field to given value.

### HasMemberQueryExemptions

`func (o *UserGroupPost) HasMemberQueryExemptions() bool`

HasMemberQueryExemptions returns a boolean if a field has been set.

### GetMemberSuggestionsNotify

`func (o *UserGroupPost) GetMemberSuggestionsNotify() bool`

GetMemberSuggestionsNotify returns the MemberSuggestionsNotify field if non-nil, zero value otherwise.

### GetMemberSuggestionsNotifyOk

`func (o *UserGroupPost) GetMemberSuggestionsNotifyOk() (*bool, bool)`

GetMemberSuggestionsNotifyOk returns a tuple with the MemberSuggestionsNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSuggestionsNotify

`func (o *UserGroupPost) SetMemberSuggestionsNotify(v bool)`

SetMemberSuggestionsNotify sets MemberSuggestionsNotify field to given value.

### HasMemberSuggestionsNotify

`func (o *UserGroupPost) HasMemberSuggestionsNotify() bool`

HasMemberSuggestionsNotify returns a boolean if a field has been set.

### GetMembershipAutomated

`func (o *UserGroupPost) GetMembershipAutomated() bool`

GetMembershipAutomated returns the MembershipAutomated field if non-nil, zero value otherwise.

### GetMembershipAutomatedOk

`func (o *UserGroupPost) GetMembershipAutomatedOk() (*bool, bool)`

GetMembershipAutomatedOk returns a tuple with the MembershipAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipAutomated

`func (o *UserGroupPost) SetMembershipAutomated(v bool)`

SetMembershipAutomated sets MembershipAutomated field to given value.

### HasMembershipAutomated

`func (o *UserGroupPost) HasMembershipAutomated() bool`

HasMembershipAutomated returns a boolean if a field has been set.

### GetMembershipMethod

`func (o *UserGroupPost) GetMembershipMethod() GroupMembershipMethodType`

GetMembershipMethod returns the MembershipMethod field if non-nil, zero value otherwise.

### GetMembershipMethodOk

`func (o *UserGroupPost) GetMembershipMethodOk() (*GroupMembershipMethodType, bool)`

GetMembershipMethodOk returns a tuple with the MembershipMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipMethod

`func (o *UserGroupPost) SetMembershipMethod(v GroupMembershipMethodType)`

SetMembershipMethod sets MembershipMethod field to given value.

### HasMembershipMethod

`func (o *UserGroupPost) HasMembershipMethod() bool`

HasMembershipMethod returns a boolean if a field has been set.

### GetName

`func (o *UserGroupPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupPost) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



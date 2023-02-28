# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**GroupAttributesUserGroup**](GroupAttributesUserGroup.md) |  | [optional] 
**Description** | Pointer to **string** | Description of a User Group | [optional] 
**Email** | Pointer to **string** | Email address of a User Group | [optional] 
**Id** | Pointer to **string** | ObjectId uniquely identifying a User Group. | [optional] 
**MemberQuery** | Pointer to [**FilterQuery**](FilterQuery.md) |  | [optional] 
**MemberQueryExemptions** | Pointer to [**[]GraphObject**](GraphObject.md) | Array of GraphObjects exempted from the query | [optional] 
**MemberSuggestionsNotify** | Pointer to **bool** | True if notification emails are to be sent for membership suggestions. | [optional] 
**MembershipAutomated** | Pointer to **bool** | True if membership of this group is automatically updated based on the Member Query and Member Query Exemptions, if configured | [optional] 
**Name** | Pointer to **string** | Display name of a User Group. | [optional] 
**SuggestionCounts** | Pointer to [**SuggestionCounts**](SuggestionCounts.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the group. | [optional] 

## Methods

### NewUserGroup

`func NewUserGroup() *UserGroup`

NewUserGroup instantiates a new UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupWithDefaults

`func NewUserGroupWithDefaults() *UserGroup`

NewUserGroupWithDefaults instantiates a new UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UserGroup) GetAttributes() GroupAttributesUserGroup`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserGroup) GetAttributesOk() (*GroupAttributesUserGroup, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserGroup) SetAttributes(v GroupAttributesUserGroup)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserGroup) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *UserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *UserGroup) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserGroup) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserGroup) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserGroup) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *UserGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemberQuery

`func (o *UserGroup) GetMemberQuery() FilterQuery`

GetMemberQuery returns the MemberQuery field if non-nil, zero value otherwise.

### GetMemberQueryOk

`func (o *UserGroup) GetMemberQueryOk() (*FilterQuery, bool)`

GetMemberQueryOk returns a tuple with the MemberQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberQuery

`func (o *UserGroup) SetMemberQuery(v FilterQuery)`

SetMemberQuery sets MemberQuery field to given value.

### HasMemberQuery

`func (o *UserGroup) HasMemberQuery() bool`

HasMemberQuery returns a boolean if a field has been set.

### GetMemberQueryExemptions

`func (o *UserGroup) GetMemberQueryExemptions() []GraphObject`

GetMemberQueryExemptions returns the MemberQueryExemptions field if non-nil, zero value otherwise.

### GetMemberQueryExemptionsOk

`func (o *UserGroup) GetMemberQueryExemptionsOk() (*[]GraphObject, bool)`

GetMemberQueryExemptionsOk returns a tuple with the MemberQueryExemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberQueryExemptions

`func (o *UserGroup) SetMemberQueryExemptions(v []GraphObject)`

SetMemberQueryExemptions sets MemberQueryExemptions field to given value.

### HasMemberQueryExemptions

`func (o *UserGroup) HasMemberQueryExemptions() bool`

HasMemberQueryExemptions returns a boolean if a field has been set.

### GetMemberSuggestionsNotify

`func (o *UserGroup) GetMemberSuggestionsNotify() bool`

GetMemberSuggestionsNotify returns the MemberSuggestionsNotify field if non-nil, zero value otherwise.

### GetMemberSuggestionsNotifyOk

`func (o *UserGroup) GetMemberSuggestionsNotifyOk() (*bool, bool)`

GetMemberSuggestionsNotifyOk returns a tuple with the MemberSuggestionsNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSuggestionsNotify

`func (o *UserGroup) SetMemberSuggestionsNotify(v bool)`

SetMemberSuggestionsNotify sets MemberSuggestionsNotify field to given value.

### HasMemberSuggestionsNotify

`func (o *UserGroup) HasMemberSuggestionsNotify() bool`

HasMemberSuggestionsNotify returns a boolean if a field has been set.

### GetMembershipAutomated

`func (o *UserGroup) GetMembershipAutomated() bool`

GetMembershipAutomated returns the MembershipAutomated field if non-nil, zero value otherwise.

### GetMembershipAutomatedOk

`func (o *UserGroup) GetMembershipAutomatedOk() (*bool, bool)`

GetMembershipAutomatedOk returns a tuple with the MembershipAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipAutomated

`func (o *UserGroup) SetMembershipAutomated(v bool)`

SetMembershipAutomated sets MembershipAutomated field to given value.

### HasMembershipAutomated

`func (o *UserGroup) HasMembershipAutomated() bool`

HasMembershipAutomated returns a boolean if a field has been set.

### GetName

`func (o *UserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSuggestionCounts

`func (o *UserGroup) GetSuggestionCounts() SuggestionCounts`

GetSuggestionCounts returns the SuggestionCounts field if non-nil, zero value otherwise.

### GetSuggestionCountsOk

`func (o *UserGroup) GetSuggestionCountsOk() (*SuggestionCounts, bool)`

GetSuggestionCountsOk returns a tuple with the SuggestionCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionCounts

`func (o *UserGroup) SetSuggestionCounts(v SuggestionCounts)`

SetSuggestionCounts sets SuggestionCounts field to given value.

### HasSuggestionCounts

`func (o *UserGroup) HasSuggestionCounts() bool`

HasSuggestionCounts returns a boolean if a field has been set.

### GetType

`func (o *UserGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserGroup) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



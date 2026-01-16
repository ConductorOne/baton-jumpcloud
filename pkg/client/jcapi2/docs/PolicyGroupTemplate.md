# PolicyGroupTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]PolicyGroupTemplateMember**](PolicyGroupTemplateMember.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyGroupTemplate

`func NewPolicyGroupTemplate() *PolicyGroupTemplate`

NewPolicyGroupTemplate instantiates a new PolicyGroupTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyGroupTemplateWithDefaults

`func NewPolicyGroupTemplateWithDefaults() *PolicyGroupTemplate`

NewPolicyGroupTemplateWithDefaults instantiates a new PolicyGroupTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PolicyGroupTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyGroupTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyGroupTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyGroupTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PolicyGroupTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyGroupTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyGroupTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyGroupTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMembers

`func (o *PolicyGroupTemplate) GetMembers() []PolicyGroupTemplateMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *PolicyGroupTemplate) GetMembersOk() (*[]PolicyGroupTemplateMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *PolicyGroupTemplate) SetMembers(v []PolicyGroupTemplateMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *PolicyGroupTemplate) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *PolicyGroupTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyGroupTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyGroupTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyGroupTemplate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



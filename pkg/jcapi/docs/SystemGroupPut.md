# SystemGroupPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** | The graph attributes. | [optional] 
**Description** | Pointer to **string** | Description of a System Group | [optional] 
**Email** | Pointer to **string** | Email address of a System Group | [optional] 
**MemberQuery** | Pointer to [**FilterQuery**](FilterQuery.md) |  | [optional] 
**MemberQueryExemptions** | Pointer to [**[]GraphObject**](GraphObject.md) | Array of GraphObjects exempted from the query | [optional] 
**Name** | **string** | Display name of a System Group. | 

## Methods

### NewSystemGroupPut

`func NewSystemGroupPut(name string, ) *SystemGroupPut`

NewSystemGroupPut instantiates a new SystemGroupPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemGroupPutWithDefaults

`func NewSystemGroupPutWithDefaults() *SystemGroupPut`

NewSystemGroupPutWithDefaults instantiates a new SystemGroupPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SystemGroupPut) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SystemGroupPut) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SystemGroupPut) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SystemGroupPut) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *SystemGroupPut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemGroupPut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemGroupPut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemGroupPut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *SystemGroupPut) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SystemGroupPut) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SystemGroupPut) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SystemGroupPut) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMemberQuery

`func (o *SystemGroupPut) GetMemberQuery() FilterQuery`

GetMemberQuery returns the MemberQuery field if non-nil, zero value otherwise.

### GetMemberQueryOk

`func (o *SystemGroupPut) GetMemberQueryOk() (*FilterQuery, bool)`

GetMemberQueryOk returns a tuple with the MemberQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberQuery

`func (o *SystemGroupPut) SetMemberQuery(v FilterQuery)`

SetMemberQuery sets MemberQuery field to given value.

### HasMemberQuery

`func (o *SystemGroupPut) HasMemberQuery() bool`

HasMemberQuery returns a boolean if a field has been set.

### GetMemberQueryExemptions

`func (o *SystemGroupPut) GetMemberQueryExemptions() []GraphObject`

GetMemberQueryExemptions returns the MemberQueryExemptions field if non-nil, zero value otherwise.

### GetMemberQueryExemptionsOk

`func (o *SystemGroupPut) GetMemberQueryExemptionsOk() (*[]GraphObject, bool)`

GetMemberQueryExemptionsOk returns a tuple with the MemberQueryExemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberQueryExemptions

`func (o *SystemGroupPut) SetMemberQueryExemptions(v []GraphObject)`

SetMemberQueryExemptions sets MemberQueryExemptions field to given value.

### HasMemberQueryExemptions

`func (o *SystemGroupPut) HasMemberQueryExemptions() bool`

HasMemberQueryExemptions returns a boolean if a field has been set.

### GetName

`func (o *SystemGroupPut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemGroupPut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemGroupPut) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


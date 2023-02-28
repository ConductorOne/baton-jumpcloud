# PolicyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** | The graph attributes. | [optional] 
**Description** | Pointer to **string** | Description of a Policy Group | [optional] 
**Email** | Pointer to **string** | E-mail address associated with a Policy Group | [optional] 
**Id** | Pointer to **string** | ObjectId uniquely identifying a Policy Group. | [optional] 
**Name** | Pointer to **string** | Display name of a Policy Group. | [optional] 
**Type** | Pointer to **string** | The type of the group; always &#39;policy&#39; for a Policy Group. | [optional] 

## Methods

### NewPolicyGroup

`func NewPolicyGroup() *PolicyGroup`

NewPolicyGroup instantiates a new PolicyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyGroupWithDefaults

`func NewPolicyGroupWithDefaults() *PolicyGroup`

NewPolicyGroupWithDefaults instantiates a new PolicyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *PolicyGroup) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PolicyGroup) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PolicyGroup) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PolicyGroup) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *PolicyGroup) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PolicyGroup) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PolicyGroup) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PolicyGroup) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *PolicyGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PolicyGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PolicyGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyGroup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyGroup) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



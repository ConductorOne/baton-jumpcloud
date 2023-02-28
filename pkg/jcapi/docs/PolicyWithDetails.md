# PolicyWithDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigFields** | Pointer to [**[]PolicyTemplateConfigField**](PolicyTemplateConfigField.md) |  | [optional] 
**Id** | Pointer to **string** | ObjectId uniquely identifying a Policy. | [optional] 
**Name** | Pointer to **string** | The description for this specific Policy. | [optional] 
**Template** | Pointer to [**PolicyTemplate**](PolicyTemplate.md) |  | [optional] 
**Values** | Pointer to [**[]PolicyValue**](PolicyValue.md) |  | [optional] 

## Methods

### NewPolicyWithDetails

`func NewPolicyWithDetails() *PolicyWithDetails`

NewPolicyWithDetails instantiates a new PolicyWithDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDetailsWithDefaults

`func NewPolicyWithDetailsWithDefaults() *PolicyWithDetails`

NewPolicyWithDetailsWithDefaults instantiates a new PolicyWithDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigFields

`func (o *PolicyWithDetails) GetConfigFields() []PolicyTemplateConfigField`

GetConfigFields returns the ConfigFields field if non-nil, zero value otherwise.

### GetConfigFieldsOk

`func (o *PolicyWithDetails) GetConfigFieldsOk() (*[]PolicyTemplateConfigField, bool)`

GetConfigFieldsOk returns a tuple with the ConfigFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFields

`func (o *PolicyWithDetails) SetConfigFields(v []PolicyTemplateConfigField)`

SetConfigFields sets ConfigFields field to given value.

### HasConfigFields

`func (o *PolicyWithDetails) HasConfigFields() bool`

HasConfigFields returns a boolean if a field has been set.

### GetId

`func (o *PolicyWithDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyWithDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyWithDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyWithDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PolicyWithDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyWithDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyWithDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyWithDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplate

`func (o *PolicyWithDetails) GetTemplate() PolicyTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PolicyWithDetails) GetTemplateOk() (*PolicyTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PolicyWithDetails) SetTemplate(v PolicyTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PolicyWithDetails) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetValues

`func (o *PolicyWithDetails) GetValues() []PolicyValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PolicyWithDetails) GetValuesOk() (*[]PolicyValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PolicyWithDetails) SetValues(v []PolicyValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *PolicyWithDetails) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



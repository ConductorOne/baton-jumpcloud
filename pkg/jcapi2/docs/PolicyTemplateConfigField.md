# PolicyTemplateConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** | The default value for this field. | [optional] 
**DisplayOptions** | Pointer to **map[string]interface{}** | The options that correspond to the display_type. | [optional] 
**DisplayType** | Pointer to **string** | The default rendering for this field. | [optional] 
**Id** | **string** | ObjectId uniquely identifying a Policy Template Configuration Field | 
**Label** | Pointer to **string** | The default label for this field. | [optional] 
**Name** | **string** | A unique name identifying this config field. | 
**Position** | Pointer to **float32** | The default position to render this field. | [optional] 
**ReadOnly** | Pointer to **bool** | If an admin is allowed to modify this field. | [optional] 
**Required** | Pointer to **bool** | If this field is required for this field. | [optional] 
**Sensitive** | Pointer to **bool** | Defines if the policy template config field is sensitive or not. | [optional] 
**Tooltip** | Pointer to [**PolicyTemplateConfigFieldTooltip**](PolicyTemplateConfigFieldTooltip.md) |  | [optional] 

## Methods

### NewPolicyTemplateConfigField

`func NewPolicyTemplateConfigField(id string, name string, ) *PolicyTemplateConfigField`

NewPolicyTemplateConfigField instantiates a new PolicyTemplateConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateConfigFieldWithDefaults

`func NewPolicyTemplateConfigFieldWithDefaults() *PolicyTemplateConfigField`

NewPolicyTemplateConfigFieldWithDefaults instantiates a new PolicyTemplateConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *PolicyTemplateConfigField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PolicyTemplateConfigField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PolicyTemplateConfigField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PolicyTemplateConfigField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDisplayOptions

`func (o *PolicyTemplateConfigField) GetDisplayOptions() map[string]interface{}`

GetDisplayOptions returns the DisplayOptions field if non-nil, zero value otherwise.

### GetDisplayOptionsOk

`func (o *PolicyTemplateConfigField) GetDisplayOptionsOk() (*map[string]interface{}, bool)`

GetDisplayOptionsOk returns a tuple with the DisplayOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOptions

`func (o *PolicyTemplateConfigField) SetDisplayOptions(v map[string]interface{})`

SetDisplayOptions sets DisplayOptions field to given value.

### HasDisplayOptions

`func (o *PolicyTemplateConfigField) HasDisplayOptions() bool`

HasDisplayOptions returns a boolean if a field has been set.

### GetDisplayType

`func (o *PolicyTemplateConfigField) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *PolicyTemplateConfigField) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *PolicyTemplateConfigField) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *PolicyTemplateConfigField) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetId

`func (o *PolicyTemplateConfigField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyTemplateConfigField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyTemplateConfigField) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *PolicyTemplateConfigField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PolicyTemplateConfigField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PolicyTemplateConfigField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PolicyTemplateConfigField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *PolicyTemplateConfigField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyTemplateConfigField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyTemplateConfigField) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *PolicyTemplateConfigField) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PolicyTemplateConfigField) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PolicyTemplateConfigField) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PolicyTemplateConfigField) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetReadOnly

`func (o *PolicyTemplateConfigField) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PolicyTemplateConfigField) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PolicyTemplateConfigField) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PolicyTemplateConfigField) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRequired

`func (o *PolicyTemplateConfigField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PolicyTemplateConfigField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PolicyTemplateConfigField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PolicyTemplateConfigField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSensitive

`func (o *PolicyTemplateConfigField) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *PolicyTemplateConfigField) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *PolicyTemplateConfigField) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *PolicyTemplateConfigField) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetTooltip

`func (o *PolicyTemplateConfigField) GetTooltip() PolicyTemplateConfigFieldTooltip`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *PolicyTemplateConfigField) GetTooltipOk() (*PolicyTemplateConfigFieldTooltip, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *PolicyTemplateConfigField) SetTooltip(v PolicyTemplateConfigFieldTooltip)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *PolicyTemplateConfigField) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



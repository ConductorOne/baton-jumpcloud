# ConfiguredPolicyTemplateValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigFieldId** | Pointer to **string** | The ObjectId of the corresponding Policy Template configuration field. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** | The value of this Configured Policy Template Value | [optional] 

## Methods

### NewConfiguredPolicyTemplateValue

`func NewConfiguredPolicyTemplateValue() *ConfiguredPolicyTemplateValue`

NewConfiguredPolicyTemplateValue instantiates a new ConfiguredPolicyTemplateValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfiguredPolicyTemplateValueWithDefaults

`func NewConfiguredPolicyTemplateValueWithDefaults() *ConfiguredPolicyTemplateValue`

NewConfiguredPolicyTemplateValueWithDefaults instantiates a new ConfiguredPolicyTemplateValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigFieldId

`func (o *ConfiguredPolicyTemplateValue) GetConfigFieldId() string`

GetConfigFieldId returns the ConfigFieldId field if non-nil, zero value otherwise.

### GetConfigFieldIdOk

`func (o *ConfiguredPolicyTemplateValue) GetConfigFieldIdOk() (*string, bool)`

GetConfigFieldIdOk returns a tuple with the ConfigFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFieldId

`func (o *ConfiguredPolicyTemplateValue) SetConfigFieldId(v string)`

SetConfigFieldId sets ConfigFieldId field to given value.

### HasConfigFieldId

`func (o *ConfiguredPolicyTemplateValue) HasConfigFieldId() bool`

HasConfigFieldId returns a boolean if a field has been set.

### GetName

`func (o *ConfiguredPolicyTemplateValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfiguredPolicyTemplateValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfiguredPolicyTemplateValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfiguredPolicyTemplateValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ConfiguredPolicyTemplateValue) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfiguredPolicyTemplateValue) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfiguredPolicyTemplateValue) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfiguredPolicyTemplateValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



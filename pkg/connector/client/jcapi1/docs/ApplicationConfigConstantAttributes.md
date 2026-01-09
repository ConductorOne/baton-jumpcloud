# ApplicationConfigConstantAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Mutable** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Tooltip** | Pointer to [**ApplicationConfigAcsUrlTooltip**](ApplicationConfigAcsUrlTooltip.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**[]ApplicationConfigConstantAttributesValueInner**](ApplicationConfigConstantAttributesValueInner.md) |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplicationConfigConstantAttributes

`func NewApplicationConfigConstantAttributes() *ApplicationConfigConstantAttributes`

NewApplicationConfigConstantAttributes instantiates a new ApplicationConfigConstantAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationConfigConstantAttributesWithDefaults

`func NewApplicationConfigConstantAttributesWithDefaults() *ApplicationConfigConstantAttributes`

NewApplicationConfigConstantAttributesWithDefaults instantiates a new ApplicationConfigConstantAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ApplicationConfigConstantAttributes) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApplicationConfigConstantAttributes) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApplicationConfigConstantAttributes) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApplicationConfigConstantAttributes) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMutable

`func (o *ApplicationConfigConstantAttributes) GetMutable() bool`

GetMutable returns the Mutable field if non-nil, zero value otherwise.

### GetMutableOk

`func (o *ApplicationConfigConstantAttributes) GetMutableOk() (*bool, bool)`

GetMutableOk returns a tuple with the Mutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutable

`func (o *ApplicationConfigConstantAttributes) SetMutable(v bool)`

SetMutable sets Mutable field to given value.

### HasMutable

`func (o *ApplicationConfigConstantAttributes) HasMutable() bool`

HasMutable returns a boolean if a field has been set.

### GetPosition

`func (o *ApplicationConfigConstantAttributes) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ApplicationConfigConstantAttributes) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ApplicationConfigConstantAttributes) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ApplicationConfigConstantAttributes) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetReadOnly

`func (o *ApplicationConfigConstantAttributes) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApplicationConfigConstantAttributes) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApplicationConfigConstantAttributes) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApplicationConfigConstantAttributes) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRequired

`func (o *ApplicationConfigConstantAttributes) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationConfigConstantAttributes) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationConfigConstantAttributes) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationConfigConstantAttributes) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetTooltip

`func (o *ApplicationConfigConstantAttributes) GetTooltip() ApplicationConfigAcsUrlTooltip`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *ApplicationConfigConstantAttributes) GetTooltipOk() (*ApplicationConfigAcsUrlTooltip, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *ApplicationConfigConstantAttributes) SetTooltip(v ApplicationConfigAcsUrlTooltip)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *ApplicationConfigConstantAttributes) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### GetType

`func (o *ApplicationConfigConstantAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationConfigConstantAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationConfigConstantAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationConfigConstantAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ApplicationConfigConstantAttributes) GetValue() []ApplicationConfigConstantAttributesValueInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationConfigConstantAttributes) GetValueOk() (*[]ApplicationConfigConstantAttributesValueInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationConfigConstantAttributes) SetValue(v []ApplicationConfigConstantAttributesValueInner)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApplicationConfigConstantAttributes) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVisible

`func (o *ApplicationConfigConstantAttributes) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ApplicationConfigConstantAttributes) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ApplicationConfigConstantAttributes) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ApplicationConfigConstantAttributes) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



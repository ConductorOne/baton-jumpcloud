# PolicyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigFieldID** | Pointer to **string** | The ObjectId of the corresponding Policy Template configuration field. | [optional] 
**Sensitive** | Pointer to **bool** | Defines if the value is sensitive or not. | [optional] 
**Value** | Pointer to **string** | The value for the configuration field for this Policy instance. | [optional] 

## Methods

### NewPolicyValue

`func NewPolicyValue() *PolicyValue`

NewPolicyValue instantiates a new PolicyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyValueWithDefaults

`func NewPolicyValueWithDefaults() *PolicyValue`

NewPolicyValueWithDefaults instantiates a new PolicyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigFieldID

`func (o *PolicyValue) GetConfigFieldID() string`

GetConfigFieldID returns the ConfigFieldID field if non-nil, zero value otherwise.

### GetConfigFieldIDOk

`func (o *PolicyValue) GetConfigFieldIDOk() (*string, bool)`

GetConfigFieldIDOk returns a tuple with the ConfigFieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFieldID

`func (o *PolicyValue) SetConfigFieldID(v string)`

SetConfigFieldID sets ConfigFieldID field to given value.

### HasConfigFieldID

`func (o *PolicyValue) HasConfigFieldID() bool`

HasConfigFieldID returns a boolean if a field has been set.

### GetSensitive

`func (o *PolicyValue) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *PolicyValue) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *PolicyValue) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.

### HasSensitive

`func (o *PolicyValue) HasSensitive() bool`

HasSensitive returns a boolean if a field has been set.

### GetValue

`func (o *PolicyValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PolicyValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PolicyValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PolicyValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



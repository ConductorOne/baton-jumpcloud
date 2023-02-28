# AuthnPolicyUserAttributeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | The only field that is currently supported is ldap_binding_user | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** | Can be any value - string, number, boolean, array or object. | [optional] 

## Methods

### NewAuthnPolicyUserAttributeFilter

`func NewAuthnPolicyUserAttributeFilter() *AuthnPolicyUserAttributeFilter`

NewAuthnPolicyUserAttributeFilter instantiates a new AuthnPolicyUserAttributeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnPolicyUserAttributeFilterWithDefaults

`func NewAuthnPolicyUserAttributeFilterWithDefaults() *AuthnPolicyUserAttributeFilter`

NewAuthnPolicyUserAttributeFilterWithDefaults instantiates a new AuthnPolicyUserAttributeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AuthnPolicyUserAttributeFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AuthnPolicyUserAttributeFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AuthnPolicyUserAttributeFilter) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *AuthnPolicyUserAttributeFilter) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOperator

`func (o *AuthnPolicyUserAttributeFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AuthnPolicyUserAttributeFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AuthnPolicyUserAttributeFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AuthnPolicyUserAttributeFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *AuthnPolicyUserAttributeFilter) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthnPolicyUserAttributeFilter) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthnPolicyUserAttributeFilter) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *AuthnPolicyUserAttributeFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



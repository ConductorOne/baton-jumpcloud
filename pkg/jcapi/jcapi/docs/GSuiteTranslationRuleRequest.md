# GSuiteTranslationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltIn** | Pointer to [**GSuiteBuiltinTranslation**](GSuiteBuiltinTranslation.md) |  | [optional] 
**Direction** | Pointer to [**GSuiteDirectionTranslation**](GSuiteDirectionTranslation.md) |  | [optional] [default to EXPORT]

## Methods

### NewGSuiteTranslationRuleRequest

`func NewGSuiteTranslationRuleRequest() *GSuiteTranslationRuleRequest`

NewGSuiteTranslationRuleRequest instantiates a new GSuiteTranslationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGSuiteTranslationRuleRequestWithDefaults

`func NewGSuiteTranslationRuleRequestWithDefaults() *GSuiteTranslationRuleRequest`

NewGSuiteTranslationRuleRequestWithDefaults instantiates a new GSuiteTranslationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltIn

`func (o *GSuiteTranslationRuleRequest) GetBuiltIn() GSuiteBuiltinTranslation`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *GSuiteTranslationRuleRequest) GetBuiltInOk() (*GSuiteBuiltinTranslation, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *GSuiteTranslationRuleRequest) SetBuiltIn(v GSuiteBuiltinTranslation)`

SetBuiltIn sets BuiltIn field to given value.

### HasBuiltIn

`func (o *GSuiteTranslationRuleRequest) HasBuiltIn() bool`

HasBuiltIn returns a boolean if a field has been set.

### GetDirection

`func (o *GSuiteTranslationRuleRequest) GetDirection() GSuiteDirectionTranslation`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GSuiteTranslationRuleRequest) GetDirectionOk() (*GSuiteDirectionTranslation, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GSuiteTranslationRuleRequest) SetDirection(v GSuiteDirectionTranslation)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GSuiteTranslationRuleRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



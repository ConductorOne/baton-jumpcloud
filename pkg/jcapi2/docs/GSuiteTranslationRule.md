# GSuiteTranslationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltIn** | Pointer to [**GSuiteBuiltinTranslation**](GSuiteBuiltinTranslation.md) |  | [optional] 
**Direction** | Pointer to [**GSuiteDirectionTranslation**](GSuiteDirectionTranslation.md) |  | [optional] [default to GSUITEDIRECTIONTRANSLATION_EXPORT]
**Id** | Pointer to **string** | ObjectId uniquely identifying a Translation Rule. | [optional] 

## Methods

### NewGSuiteTranslationRule

`func NewGSuiteTranslationRule() *GSuiteTranslationRule`

NewGSuiteTranslationRule instantiates a new GSuiteTranslationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGSuiteTranslationRuleWithDefaults

`func NewGSuiteTranslationRuleWithDefaults() *GSuiteTranslationRule`

NewGSuiteTranslationRuleWithDefaults instantiates a new GSuiteTranslationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltIn

`func (o *GSuiteTranslationRule) GetBuiltIn() GSuiteBuiltinTranslation`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *GSuiteTranslationRule) GetBuiltInOk() (*GSuiteBuiltinTranslation, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *GSuiteTranslationRule) SetBuiltIn(v GSuiteBuiltinTranslation)`

SetBuiltIn sets BuiltIn field to given value.

### HasBuiltIn

`func (o *GSuiteTranslationRule) HasBuiltIn() bool`

HasBuiltIn returns a boolean if a field has been set.

### GetDirection

`func (o *GSuiteTranslationRule) GetDirection() GSuiteDirectionTranslation`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GSuiteTranslationRule) GetDirectionOk() (*GSuiteDirectionTranslation, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GSuiteTranslationRule) SetDirection(v GSuiteDirectionTranslation)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GSuiteTranslationRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *GSuiteTranslationRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GSuiteTranslationRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GSuiteTranslationRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GSuiteTranslationRule) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



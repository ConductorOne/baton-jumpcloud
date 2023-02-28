# Office365TranslationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltIn** | Pointer to [**Office365BuiltinTranslation**](Office365BuiltinTranslation.md) |  | [optional] 
**Direction** | Pointer to [**Office365DirectionTranslation**](Office365DirectionTranslation.md) |  | [optional] [default to EXPORT]

## Methods

### NewOffice365TranslationRuleRequest

`func NewOffice365TranslationRuleRequest() *Office365TranslationRuleRequest`

NewOffice365TranslationRuleRequest instantiates a new Office365TranslationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365TranslationRuleRequestWithDefaults

`func NewOffice365TranslationRuleRequestWithDefaults() *Office365TranslationRuleRequest`

NewOffice365TranslationRuleRequestWithDefaults instantiates a new Office365TranslationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltIn

`func (o *Office365TranslationRuleRequest) GetBuiltIn() Office365BuiltinTranslation`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *Office365TranslationRuleRequest) GetBuiltInOk() (*Office365BuiltinTranslation, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *Office365TranslationRuleRequest) SetBuiltIn(v Office365BuiltinTranslation)`

SetBuiltIn sets BuiltIn field to given value.

### HasBuiltIn

`func (o *Office365TranslationRuleRequest) HasBuiltIn() bool`

HasBuiltIn returns a boolean if a field has been set.

### GetDirection

`func (o *Office365TranslationRuleRequest) GetDirection() Office365DirectionTranslation`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Office365TranslationRuleRequest) GetDirectionOk() (*Office365DirectionTranslation, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Office365TranslationRuleRequest) SetDirection(v Office365DirectionTranslation)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Office365TranslationRuleRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Office365TranslationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltIn** | Pointer to [**Office365BuiltinTranslation**](Office365BuiltinTranslation.md) |  | [optional] 
**Direction** | Pointer to [**Office365DirectionTranslation**](Office365DirectionTranslation.md) |  | [optional] [default to OFFICE365DIRECTIONTRANSLATION_EXPORT]
**Id** | Pointer to **string** | ObjectId uniquely identifying a Translation Rule. | [optional] 

## Methods

### NewOffice365TranslationRule

`func NewOffice365TranslationRule() *Office365TranslationRule`

NewOffice365TranslationRule instantiates a new Office365TranslationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365TranslationRuleWithDefaults

`func NewOffice365TranslationRuleWithDefaults() *Office365TranslationRule`

NewOffice365TranslationRuleWithDefaults instantiates a new Office365TranslationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltIn

`func (o *Office365TranslationRule) GetBuiltIn() Office365BuiltinTranslation`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *Office365TranslationRule) GetBuiltInOk() (*Office365BuiltinTranslation, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *Office365TranslationRule) SetBuiltIn(v Office365BuiltinTranslation)`

SetBuiltIn sets BuiltIn field to given value.

### HasBuiltIn

`func (o *Office365TranslationRule) HasBuiltIn() bool`

HasBuiltIn returns a boolean if a field has been set.

### GetDirection

`func (o *Office365TranslationRule) GetDirection() Office365DirectionTranslation`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Office365TranslationRule) GetDirectionOk() (*Office365DirectionTranslation, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Office365TranslationRule) SetDirection(v Office365DirectionTranslation)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Office365TranslationRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *Office365TranslationRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Office365TranslationRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Office365TranslationRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Office365TranslationRule) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CustomEmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]CustomEmailTemplateField**](CustomEmailTemplateField.md) |  | [optional] 
**Type** | Pointer to [**CustomEmailType**](CustomEmailType.md) |  | [optional] 

## Methods

### NewCustomEmailTemplate

`func NewCustomEmailTemplate() *CustomEmailTemplate`

NewCustomEmailTemplate instantiates a new CustomEmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEmailTemplateWithDefaults

`func NewCustomEmailTemplateWithDefaults() *CustomEmailTemplate`

NewCustomEmailTemplateWithDefaults instantiates a new CustomEmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomEmailTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEmailTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEmailTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEmailTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CustomEmailTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CustomEmailTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CustomEmailTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CustomEmailTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFields

`func (o *CustomEmailTemplate) GetFields() []CustomEmailTemplateField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CustomEmailTemplate) GetFieldsOk() (*[]CustomEmailTemplateField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CustomEmailTemplate) SetFields(v []CustomEmailTemplateField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CustomEmailTemplate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetType

`func (o *CustomEmailTemplate) GetType() CustomEmailType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomEmailTemplate) GetTypeOk() (*CustomEmailType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomEmailTemplate) SetType(v CustomEmailType)`

SetType sets Type field to given value.

### HasType

`func (o *CustomEmailTemplate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



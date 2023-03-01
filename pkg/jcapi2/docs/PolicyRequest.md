# PolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The description for this specific Policy. | 
**Template** | Pointer to [**PolicyRequestTemplate**](PolicyRequestTemplate.md) |  | [optional] 
**Values** | Pointer to [**[]PolicyValue**](PolicyValue.md) |  | [optional] 

## Methods

### NewPolicyRequest

`func NewPolicyRequest(name string, ) *PolicyRequest`

NewPolicyRequest instantiates a new PolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRequestWithDefaults

`func NewPolicyRequestWithDefaults() *PolicyRequest`

NewPolicyRequestWithDefaults instantiates a new PolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTemplate

`func (o *PolicyRequest) GetTemplate() PolicyRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PolicyRequest) GetTemplateOk() (*PolicyRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PolicyRequest) SetTemplate(v PolicyRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PolicyRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetValues

`func (o *PolicyRequest) GetValues() []PolicyValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PolicyRequest) GetValuesOk() (*[]PolicyValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PolicyRequest) SetValues(v []PolicyValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *PolicyRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



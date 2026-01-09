# CloudInsightsTermConjunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** |  | [optional] 
**TermItems** | Pointer to [**[]CloudInsightsTermItem**](CloudInsightsTermItem.md) |  | [optional] 

## Methods

### NewCloudInsightsTermConjunction

`func NewCloudInsightsTermConjunction() *CloudInsightsTermConjunction`

NewCloudInsightsTermConjunction instantiates a new CloudInsightsTermConjunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsTermConjunctionWithDefaults

`func NewCloudInsightsTermConjunctionWithDefaults() *CloudInsightsTermConjunction`

NewCloudInsightsTermConjunctionWithDefaults instantiates a new CloudInsightsTermConjunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *CloudInsightsTermConjunction) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CloudInsightsTermConjunction) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CloudInsightsTermConjunction) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CloudInsightsTermConjunction) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetTermItems

`func (o *CloudInsightsTermConjunction) GetTermItems() []CloudInsightsTermItem`

GetTermItems returns the TermItems field if non-nil, zero value otherwise.

### GetTermItemsOk

`func (o *CloudInsightsTermConjunction) GetTermItemsOk() (*[]CloudInsightsTermItem, bool)`

GetTermItemsOk returns a tuple with the TermItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermItems

`func (o *CloudInsightsTermConjunction) SetTermItems(v []CloudInsightsTermItem)`

SetTermItems sets TermItems field to given value.

### HasTermItems

`func (o *CloudInsightsTermConjunction) HasTermItems() bool`

HasTermItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



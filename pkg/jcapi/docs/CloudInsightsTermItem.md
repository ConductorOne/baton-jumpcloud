# CloudInsightsTermItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermArray** | Pointer to [**CloudInsightsTermItemTermArray**](CloudInsightsTermItemTermArray.md) |  | [optional] 
**TermConjunction** | Pointer to **map[string]interface{}** |  | [optional] 
**TermValue** | Pointer to [**CloudInsightsTermItemTermValue**](CloudInsightsTermItemTermValue.md) |  | [optional] 

## Methods

### NewCloudInsightsTermItem

`func NewCloudInsightsTermItem() *CloudInsightsTermItem`

NewCloudInsightsTermItem instantiates a new CloudInsightsTermItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsTermItemWithDefaults

`func NewCloudInsightsTermItemWithDefaults() *CloudInsightsTermItem`

NewCloudInsightsTermItemWithDefaults instantiates a new CloudInsightsTermItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermArray

`func (o *CloudInsightsTermItem) GetTermArray() CloudInsightsTermItemTermArray`

GetTermArray returns the TermArray field if non-nil, zero value otherwise.

### GetTermArrayOk

`func (o *CloudInsightsTermItem) GetTermArrayOk() (*CloudInsightsTermItemTermArray, bool)`

GetTermArrayOk returns a tuple with the TermArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermArray

`func (o *CloudInsightsTermItem) SetTermArray(v CloudInsightsTermItemTermArray)`

SetTermArray sets TermArray field to given value.

### HasTermArray

`func (o *CloudInsightsTermItem) HasTermArray() bool`

HasTermArray returns a boolean if a field has been set.

### GetTermConjunction

`func (o *CloudInsightsTermItem) GetTermConjunction() map[string]interface{}`

GetTermConjunction returns the TermConjunction field if non-nil, zero value otherwise.

### GetTermConjunctionOk

`func (o *CloudInsightsTermItem) GetTermConjunctionOk() (*map[string]interface{}, bool)`

GetTermConjunctionOk returns a tuple with the TermConjunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermConjunction

`func (o *CloudInsightsTermItem) SetTermConjunction(v map[string]interface{})`

SetTermConjunction sets TermConjunction field to given value.

### HasTermConjunction

`func (o *CloudInsightsTermItem) HasTermConjunction() bool`

HasTermConjunction returns a boolean if a field has been set.

### GetTermValue

`func (o *CloudInsightsTermItem) GetTermValue() CloudInsightsTermItemTermValue`

GetTermValue returns the TermValue field if non-nil, zero value otherwise.

### GetTermValueOk

`func (o *CloudInsightsTermItem) GetTermValueOk() (*CloudInsightsTermItemTermValue, bool)`

GetTermValueOk returns a tuple with the TermValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermValue

`func (o *CloudInsightsTermItem) SetTermValue(v CloudInsightsTermItemTermValue)`

SetTermValue sets TermValue field to given value.

### HasTermValue

`func (o *CloudInsightsTermItem) HasTermValue() bool`

HasTermValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



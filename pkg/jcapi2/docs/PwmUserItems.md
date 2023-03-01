# PwmUserItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]PwmItemsCountByType**](PwmItemsCountByType.md) |  | [optional] 
**Results** | [**[]PwmUserItem**](PwmUserItem.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewPwmUserItems

`func NewPwmUserItems(results []PwmUserItem, totalCount int32, ) *PwmUserItems`

NewPwmUserItems instantiates a new PwmUserItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPwmUserItemsWithDefaults

`func NewPwmUserItemsWithDefaults() *PwmUserItems`

NewPwmUserItemsWithDefaults instantiates a new PwmUserItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PwmUserItems) GetItems() []PwmItemsCountByType`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PwmUserItems) GetItemsOk() (*[]PwmItemsCountByType, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PwmUserItems) SetItems(v []PwmItemsCountByType)`

SetItems sets Items field to given value.

### HasItems

`func (o *PwmUserItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResults

`func (o *PwmUserItems) GetResults() []PwmUserItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PwmUserItems) GetResultsOk() (*[]PwmUserItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PwmUserItems) SetResults(v []PwmUserItem)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PwmUserItems) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PwmUserItems) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PwmUserItems) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



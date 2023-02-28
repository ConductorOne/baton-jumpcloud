# PwmItemsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]PwmItemsCountByType**](PwmItemsCountByType.md) |  | [optional] 
**Results** | [**[]PwmItemsMetadataResultsInner**](PwmItemsMetadataResultsInner.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewPwmItemsMetadata

`func NewPwmItemsMetadata(results []PwmItemsMetadataResultsInner, totalCount int32, ) *PwmItemsMetadata`

NewPwmItemsMetadata instantiates a new PwmItemsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPwmItemsMetadataWithDefaults

`func NewPwmItemsMetadataWithDefaults() *PwmItemsMetadata`

NewPwmItemsMetadataWithDefaults instantiates a new PwmItemsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PwmItemsMetadata) GetItems() []PwmItemsCountByType`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PwmItemsMetadata) GetItemsOk() (*[]PwmItemsCountByType, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PwmItemsMetadata) SetItems(v []PwmItemsCountByType)`

SetItems sets Items field to given value.

### HasItems

`func (o *PwmItemsMetadata) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResults

`func (o *PwmItemsMetadata) GetResults() []PwmItemsMetadataResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PwmItemsMetadata) GetResultsOk() (*[]PwmItemsMetadataResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PwmItemsMetadata) SetResults(v []PwmItemsMetadataResultsInner)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PwmItemsMetadata) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PwmItemsMetadata) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PwmItemsMetadata) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



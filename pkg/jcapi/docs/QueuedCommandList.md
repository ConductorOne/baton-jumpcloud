# QueuedCommandList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]QueuedCommandListResultsInner**](QueuedCommandListResultsInner.md) |  | [optional] 
**TotalCount** | Pointer to **int32** | The total number of queued command results. | [optional] 

## Methods

### NewQueuedCommandList

`func NewQueuedCommandList() *QueuedCommandList`

NewQueuedCommandList instantiates a new QueuedCommandList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedCommandListWithDefaults

`func NewQueuedCommandListWithDefaults() *QueuedCommandList`

NewQueuedCommandListWithDefaults instantiates a new QueuedCommandList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *QueuedCommandList) GetResults() []QueuedCommandListResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueuedCommandList) GetResultsOk() (*[]QueuedCommandListResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueuedCommandList) SetResults(v []QueuedCommandListResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *QueuedCommandList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *QueuedCommandList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *QueuedCommandList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *QueuedCommandList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *QueuedCommandList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



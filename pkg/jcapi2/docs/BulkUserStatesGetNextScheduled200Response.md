# BulkUserStatesGetNextScheduled200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventsCount** | Pointer to **int32** | The total number of ACTIVATED and SUSPENDED events to a max depth of 1 for all of the users in the query. A value larger than the limit specified on the query indicates that additional calls are needed, using a skip greater than 0, to retrieve the full set of results. | [optional] 
**Results** | Pointer to [**[]ScheduledUserstateResult**](ScheduledUserstateResult.md) |  | [optional] 

## Methods

### NewBulkUserStatesGetNextScheduled200Response

`func NewBulkUserStatesGetNextScheduled200Response() *BulkUserStatesGetNextScheduled200Response`

NewBulkUserStatesGetNextScheduled200Response instantiates a new BulkUserStatesGetNextScheduled200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUserStatesGetNextScheduled200ResponseWithDefaults

`func NewBulkUserStatesGetNextScheduled200ResponseWithDefaults() *BulkUserStatesGetNextScheduled200Response`

NewBulkUserStatesGetNextScheduled200ResponseWithDefaults instantiates a new BulkUserStatesGetNextScheduled200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventsCount

`func (o *BulkUserStatesGetNextScheduled200Response) GetEventsCount() int32`

GetEventsCount returns the EventsCount field if non-nil, zero value otherwise.

### GetEventsCountOk

`func (o *BulkUserStatesGetNextScheduled200Response) GetEventsCountOk() (*int32, bool)`

GetEventsCountOk returns a tuple with the EventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsCount

`func (o *BulkUserStatesGetNextScheduled200Response) SetEventsCount(v int32)`

SetEventsCount sets EventsCount field to given value.

### HasEventsCount

`func (o *BulkUserStatesGetNextScheduled200Response) HasEventsCount() bool`

HasEventsCount returns a boolean if a field has been set.

### GetResults

`func (o *BulkUserStatesGetNextScheduled200Response) GetResults() []ScheduledUserstateResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkUserStatesGetNextScheduled200Response) GetResultsOk() (*[]ScheduledUserstateResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkUserStatesGetNextScheduled200Response) SetResults(v []ScheduledUserstateResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *BulkUserStatesGetNextScheduled200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



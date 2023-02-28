# CloudInsightsEventsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]CloudInsightsEventsListEvents**](CloudInsightsEventsListEvents.md) |  | [optional] 
**SearchAfter** | Pointer to **map[string]interface{}** |  | [optional] 
**TotalCount** | Pointer to **float32** |  | [optional] 

## Methods

### NewCloudInsightsEventsListResponse

`func NewCloudInsightsEventsListResponse() *CloudInsightsEventsListResponse`

NewCloudInsightsEventsListResponse instantiates a new CloudInsightsEventsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsEventsListResponseWithDefaults

`func NewCloudInsightsEventsListResponseWithDefaults() *CloudInsightsEventsListResponse`

NewCloudInsightsEventsListResponseWithDefaults instantiates a new CloudInsightsEventsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *CloudInsightsEventsListResponse) GetEvents() []CloudInsightsEventsListEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CloudInsightsEventsListResponse) GetEventsOk() (*[]CloudInsightsEventsListEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CloudInsightsEventsListResponse) SetEvents(v []CloudInsightsEventsListEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CloudInsightsEventsListResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetSearchAfter

`func (o *CloudInsightsEventsListResponse) GetSearchAfter() map[string]interface{}`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *CloudInsightsEventsListResponse) GetSearchAfterOk() (*map[string]interface{}, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *CloudInsightsEventsListResponse) SetSearchAfter(v map[string]interface{})`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *CloudInsightsEventsListResponse) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetTotalCount

`func (o *CloudInsightsEventsListResponse) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CloudInsightsEventsListResponse) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CloudInsightsEventsListResponse) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CloudInsightsEventsListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



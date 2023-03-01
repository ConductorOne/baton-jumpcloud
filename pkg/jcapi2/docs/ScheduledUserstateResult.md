# ScheduledUserstateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledDate** | Pointer to **string** | The UTC date and time when the scheduled job will execute. | [optional] 
**ScheduledJobId** | Pointer to **string** | The id of the scheduled job that scheduled the state change. | [optional] 
**State** | Pointer to **string** | The state that the user will be in once the scheduled job executes. | [optional] 
**SystemUserId** | Pointer to **string** | The id of the user that the scheduled job will update. | [optional] 

## Methods

### NewScheduledUserstateResult

`func NewScheduledUserstateResult() *ScheduledUserstateResult`

NewScheduledUserstateResult instantiates a new ScheduledUserstateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledUserstateResultWithDefaults

`func NewScheduledUserstateResultWithDefaults() *ScheduledUserstateResult`

NewScheduledUserstateResultWithDefaults instantiates a new ScheduledUserstateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledDate

`func (o *ScheduledUserstateResult) GetScheduledDate() string`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *ScheduledUserstateResult) GetScheduledDateOk() (*string, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *ScheduledUserstateResult) SetScheduledDate(v string)`

SetScheduledDate sets ScheduledDate field to given value.

### HasScheduledDate

`func (o *ScheduledUserstateResult) HasScheduledDate() bool`

HasScheduledDate returns a boolean if a field has been set.

### GetScheduledJobId

`func (o *ScheduledUserstateResult) GetScheduledJobId() string`

GetScheduledJobId returns the ScheduledJobId field if non-nil, zero value otherwise.

### GetScheduledJobIdOk

`func (o *ScheduledUserstateResult) GetScheduledJobIdOk() (*string, bool)`

GetScheduledJobIdOk returns a tuple with the ScheduledJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledJobId

`func (o *ScheduledUserstateResult) SetScheduledJobId(v string)`

SetScheduledJobId sets ScheduledJobId field to given value.

### HasScheduledJobId

`func (o *ScheduledUserstateResult) HasScheduledJobId() bool`

HasScheduledJobId returns a boolean if a field has been set.

### GetState

`func (o *ScheduledUserstateResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ScheduledUserstateResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ScheduledUserstateResult) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ScheduledUserstateResult) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSystemUserId

`func (o *ScheduledUserstateResult) GetSystemUserId() string`

GetSystemUserId returns the SystemUserId field if non-nil, zero value otherwise.

### GetSystemUserIdOk

`func (o *ScheduledUserstateResult) GetSystemUserIdOk() (*string, bool)`

GetSystemUserIdOk returns a tuple with the SystemUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUserId

`func (o *ScheduledUserstateResult) SetSystemUserId(v string)`

SetSystemUserId sets SystemUserId field to given value.

### HasSystemUserId

`func (o *ScheduledUserstateResult) HasSystemUserId() bool`

HasSystemUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



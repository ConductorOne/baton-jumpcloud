# SyncroTicketingAlertConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueDays** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**ProblemType** | **string** |  | 
**ShouldCreateTickets** | **bool** |  | 
**Status** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **float32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewSyncroTicketingAlertConfigurationRequest

`func NewSyncroTicketingAlertConfigurationRequest(problemType string, shouldCreateTickets bool, ) *SyncroTicketingAlertConfigurationRequest`

NewSyncroTicketingAlertConfigurationRequest instantiates a new SyncroTicketingAlertConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncroTicketingAlertConfigurationRequestWithDefaults

`func NewSyncroTicketingAlertConfigurationRequestWithDefaults() *SyncroTicketingAlertConfigurationRequest`

NewSyncroTicketingAlertConfigurationRequestWithDefaults instantiates a new SyncroTicketingAlertConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueDays

`func (o *SyncroTicketingAlertConfigurationRequest) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *SyncroTicketingAlertConfigurationRequest) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *SyncroTicketingAlertConfigurationRequest) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *SyncroTicketingAlertConfigurationRequest) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### GetPriority

`func (o *SyncroTicketingAlertConfigurationRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SyncroTicketingAlertConfigurationRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SyncroTicketingAlertConfigurationRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SyncroTicketingAlertConfigurationRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProblemType

`func (o *SyncroTicketingAlertConfigurationRequest) GetProblemType() string`

GetProblemType returns the ProblemType field if non-nil, zero value otherwise.

### GetProblemTypeOk

`func (o *SyncroTicketingAlertConfigurationRequest) GetProblemTypeOk() (*string, bool)`

GetProblemTypeOk returns a tuple with the ProblemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemType

`func (o *SyncroTicketingAlertConfigurationRequest) SetProblemType(v string)`

SetProblemType sets ProblemType field to given value.


### GetShouldCreateTickets

`func (o *SyncroTicketingAlertConfigurationRequest) GetShouldCreateTickets() bool`

GetShouldCreateTickets returns the ShouldCreateTickets field if non-nil, zero value otherwise.

### GetShouldCreateTicketsOk

`func (o *SyncroTicketingAlertConfigurationRequest) GetShouldCreateTicketsOk() (*bool, bool)`

GetShouldCreateTicketsOk returns a tuple with the ShouldCreateTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateTickets

`func (o *SyncroTicketingAlertConfigurationRequest) SetShouldCreateTickets(v bool)`

SetShouldCreateTickets sets ShouldCreateTickets field to given value.


### GetStatus

`func (o *SyncroTicketingAlertConfigurationRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncroTicketingAlertConfigurationRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncroTicketingAlertConfigurationRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyncroTicketingAlertConfigurationRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *SyncroTicketingAlertConfigurationRequest) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SyncroTicketingAlertConfigurationRequest) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SyncroTicketingAlertConfigurationRequest) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SyncroTicketingAlertConfigurationRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *SyncroTicketingAlertConfigurationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SyncroTicketingAlertConfigurationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SyncroTicketingAlertConfigurationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SyncroTicketingAlertConfigurationRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



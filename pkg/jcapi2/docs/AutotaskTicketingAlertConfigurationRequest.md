# AutotaskTicketingAlertConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | **string** |  | 
**DueDays** | **int32** |  | 
**Priority** | [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | 
**Queue** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 
**Resource** | Pointer to [**AutotaskTicketingAlertConfigurationResource**](AutotaskTicketingAlertConfigurationResource.md) |  | [optional] 
**ShouldCreateTickets** | **bool** |  | 
**Source** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 
**Status** | [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | 

## Methods

### NewAutotaskTicketingAlertConfigurationRequest

`func NewAutotaskTicketingAlertConfigurationRequest(destination string, dueDays int32, priority AutotaskTicketingAlertConfigurationPriority, shouldCreateTickets bool, status AutotaskTicketingAlertConfigurationPriority, ) *AutotaskTicketingAlertConfigurationRequest`

NewAutotaskTicketingAlertConfigurationRequest instantiates a new AutotaskTicketingAlertConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotaskTicketingAlertConfigurationRequestWithDefaults

`func NewAutotaskTicketingAlertConfigurationRequestWithDefaults() *AutotaskTicketingAlertConfigurationRequest`

NewAutotaskTicketingAlertConfigurationRequestWithDefaults instantiates a new AutotaskTicketingAlertConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *AutotaskTicketingAlertConfigurationRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AutotaskTicketingAlertConfigurationRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AutotaskTicketingAlertConfigurationRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetDueDays

`func (o *AutotaskTicketingAlertConfigurationRequest) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *AutotaskTicketingAlertConfigurationRequest) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *AutotaskTicketingAlertConfigurationRequest) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.


### GetPriority

`func (o *AutotaskTicketingAlertConfigurationRequest) GetPriority() AutotaskTicketingAlertConfigurationPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AutotaskTicketingAlertConfigurationRequest) GetPriorityOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AutotaskTicketingAlertConfigurationRequest) SetPriority(v AutotaskTicketingAlertConfigurationPriority)`

SetPriority sets Priority field to given value.


### GetQueue

`func (o *AutotaskTicketingAlertConfigurationRequest) GetQueue() AutotaskTicketingAlertConfigurationPriority`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *AutotaskTicketingAlertConfigurationRequest) GetQueueOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *AutotaskTicketingAlertConfigurationRequest) SetQueue(v AutotaskTicketingAlertConfigurationPriority)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *AutotaskTicketingAlertConfigurationRequest) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetResource

`func (o *AutotaskTicketingAlertConfigurationRequest) GetResource() AutotaskTicketingAlertConfigurationResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AutotaskTicketingAlertConfigurationRequest) GetResourceOk() (*AutotaskTicketingAlertConfigurationResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AutotaskTicketingAlertConfigurationRequest) SetResource(v AutotaskTicketingAlertConfigurationResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AutotaskTicketingAlertConfigurationRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetShouldCreateTickets

`func (o *AutotaskTicketingAlertConfigurationRequest) GetShouldCreateTickets() bool`

GetShouldCreateTickets returns the ShouldCreateTickets field if non-nil, zero value otherwise.

### GetShouldCreateTicketsOk

`func (o *AutotaskTicketingAlertConfigurationRequest) GetShouldCreateTicketsOk() (*bool, bool)`

GetShouldCreateTicketsOk returns a tuple with the ShouldCreateTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateTickets

`func (o *AutotaskTicketingAlertConfigurationRequest) SetShouldCreateTickets(v bool)`

SetShouldCreateTickets sets ShouldCreateTickets field to given value.


### GetSource

`func (o *AutotaskTicketingAlertConfigurationRequest) GetSource() AutotaskTicketingAlertConfigurationPriority`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AutotaskTicketingAlertConfigurationRequest) GetSourceOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AutotaskTicketingAlertConfigurationRequest) SetSource(v AutotaskTicketingAlertConfigurationPriority)`

SetSource sets Source field to given value.

### HasSource

`func (o *AutotaskTicketingAlertConfigurationRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *AutotaskTicketingAlertConfigurationRequest) GetStatus() AutotaskTicketingAlertConfigurationPriority`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutotaskTicketingAlertConfigurationRequest) GetStatusOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutotaskTicketingAlertConfigurationRequest) SetStatus(v AutotaskTicketingAlertConfigurationPriority)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



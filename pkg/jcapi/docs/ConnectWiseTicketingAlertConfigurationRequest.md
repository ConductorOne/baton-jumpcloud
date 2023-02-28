# ConnectWiseTicketingAlertConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueDays** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 
**ShouldCreateTickets** | **bool** |  | 
**Source** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 

## Methods

### NewConnectWiseTicketingAlertConfigurationRequest

`func NewConnectWiseTicketingAlertConfigurationRequest(shouldCreateTickets bool, ) *ConnectWiseTicketingAlertConfigurationRequest`

NewConnectWiseTicketingAlertConfigurationRequest instantiates a new ConnectWiseTicketingAlertConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectWiseTicketingAlertConfigurationRequestWithDefaults

`func NewConnectWiseTicketingAlertConfigurationRequestWithDefaults() *ConnectWiseTicketingAlertConfigurationRequest`

NewConnectWiseTicketingAlertConfigurationRequestWithDefaults instantiates a new ConnectWiseTicketingAlertConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueDays

`func (o *ConnectWiseTicketingAlertConfigurationRequest) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *ConnectWiseTicketingAlertConfigurationRequest) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *ConnectWiseTicketingAlertConfigurationRequest) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *ConnectWiseTicketingAlertConfigurationRequest) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### GetPriority

`func (o *ConnectWiseTicketingAlertConfigurationRequest) GetPriority() AutotaskTicketingAlertConfigurationPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConnectWiseTicketingAlertConfigurationRequest) GetPriorityOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConnectWiseTicketingAlertConfigurationRequest) SetPriority(v AutotaskTicketingAlertConfigurationPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConnectWiseTicketingAlertConfigurationRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetShouldCreateTickets

`func (o *ConnectWiseTicketingAlertConfigurationRequest) GetShouldCreateTickets() bool`

GetShouldCreateTickets returns the ShouldCreateTickets field if non-nil, zero value otherwise.

### GetShouldCreateTicketsOk

`func (o *ConnectWiseTicketingAlertConfigurationRequest) GetShouldCreateTicketsOk() (*bool, bool)`

GetShouldCreateTicketsOk returns a tuple with the ShouldCreateTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateTickets

`func (o *ConnectWiseTicketingAlertConfigurationRequest) SetShouldCreateTickets(v bool)`

SetShouldCreateTickets sets ShouldCreateTickets field to given value.


### GetSource

`func (o *ConnectWiseTicketingAlertConfigurationRequest) GetSource() AutotaskTicketingAlertConfigurationPriority`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConnectWiseTicketingAlertConfigurationRequest) GetSourceOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConnectWiseTicketingAlertConfigurationRequest) SetSource(v AutotaskTicketingAlertConfigurationPriority)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConnectWiseTicketingAlertConfigurationRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AutotaskTicketingAlertConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DueDays** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 
**Queue** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 
**Resource** | Pointer to [**AutotaskTicketingAlertConfigurationResource**](AutotaskTicketingAlertConfigurationResource.md) |  | [optional] 
**ShouldCreateTickets** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 
**Status** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 

## Methods

### NewAutotaskTicketingAlertConfiguration

`func NewAutotaskTicketingAlertConfiguration() *AutotaskTicketingAlertConfiguration`

NewAutotaskTicketingAlertConfiguration instantiates a new AutotaskTicketingAlertConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotaskTicketingAlertConfigurationWithDefaults

`func NewAutotaskTicketingAlertConfigurationWithDefaults() *AutotaskTicketingAlertConfiguration`

NewAutotaskTicketingAlertConfigurationWithDefaults instantiates a new AutotaskTicketingAlertConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AutotaskTicketingAlertConfiguration) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AutotaskTicketingAlertConfiguration) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AutotaskTicketingAlertConfiguration) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AutotaskTicketingAlertConfiguration) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *AutotaskTicketingAlertConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutotaskTicketingAlertConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutotaskTicketingAlertConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutotaskTicketingAlertConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestination

`func (o *AutotaskTicketingAlertConfiguration) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AutotaskTicketingAlertConfiguration) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AutotaskTicketingAlertConfiguration) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *AutotaskTicketingAlertConfiguration) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutotaskTicketingAlertConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutotaskTicketingAlertConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutotaskTicketingAlertConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutotaskTicketingAlertConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDueDays

`func (o *AutotaskTicketingAlertConfiguration) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *AutotaskTicketingAlertConfiguration) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *AutotaskTicketingAlertConfiguration) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *AutotaskTicketingAlertConfiguration) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### GetId

`func (o *AutotaskTicketingAlertConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutotaskTicketingAlertConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutotaskTicketingAlertConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AutotaskTicketingAlertConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *AutotaskTicketingAlertConfiguration) GetPriority() AutotaskTicketingAlertConfigurationPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AutotaskTicketingAlertConfiguration) GetPriorityOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AutotaskTicketingAlertConfiguration) SetPriority(v AutotaskTicketingAlertConfigurationPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AutotaskTicketingAlertConfiguration) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQueue

`func (o *AutotaskTicketingAlertConfiguration) GetQueue() AutotaskTicketingAlertConfigurationPriority`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *AutotaskTicketingAlertConfiguration) GetQueueOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *AutotaskTicketingAlertConfiguration) SetQueue(v AutotaskTicketingAlertConfigurationPriority)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *AutotaskTicketingAlertConfiguration) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetResource

`func (o *AutotaskTicketingAlertConfiguration) GetResource() AutotaskTicketingAlertConfigurationResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AutotaskTicketingAlertConfiguration) GetResourceOk() (*AutotaskTicketingAlertConfigurationResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AutotaskTicketingAlertConfiguration) SetResource(v AutotaskTicketingAlertConfigurationResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AutotaskTicketingAlertConfiguration) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetShouldCreateTickets

`func (o *AutotaskTicketingAlertConfiguration) GetShouldCreateTickets() bool`

GetShouldCreateTickets returns the ShouldCreateTickets field if non-nil, zero value otherwise.

### GetShouldCreateTicketsOk

`func (o *AutotaskTicketingAlertConfiguration) GetShouldCreateTicketsOk() (*bool, bool)`

GetShouldCreateTicketsOk returns a tuple with the ShouldCreateTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateTickets

`func (o *AutotaskTicketingAlertConfiguration) SetShouldCreateTickets(v bool)`

SetShouldCreateTickets sets ShouldCreateTickets field to given value.

### HasShouldCreateTickets

`func (o *AutotaskTicketingAlertConfiguration) HasShouldCreateTickets() bool`

HasShouldCreateTickets returns a boolean if a field has been set.

### GetSource

`func (o *AutotaskTicketingAlertConfiguration) GetSource() AutotaskTicketingAlertConfigurationPriority`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AutotaskTicketingAlertConfiguration) GetSourceOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AutotaskTicketingAlertConfiguration) SetSource(v AutotaskTicketingAlertConfigurationPriority)`

SetSource sets Source field to given value.

### HasSource

`func (o *AutotaskTicketingAlertConfiguration) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *AutotaskTicketingAlertConfiguration) GetStatus() AutotaskTicketingAlertConfigurationPriority`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutotaskTicketingAlertConfiguration) GetStatusOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutotaskTicketingAlertConfiguration) SetStatus(v AutotaskTicketingAlertConfigurationPriority)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutotaskTicketingAlertConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



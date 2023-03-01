# ConnectWiseTicketingAlertConfigurationListRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DueDays** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 
**ShouldCreateTickets** | **bool** |  | 
**Source** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 

## Methods

### NewConnectWiseTicketingAlertConfigurationListRecordsInner

`func NewConnectWiseTicketingAlertConfigurationListRecordsInner(shouldCreateTickets bool, ) *ConnectWiseTicketingAlertConfigurationListRecordsInner`

NewConnectWiseTicketingAlertConfigurationListRecordsInner instantiates a new ConnectWiseTicketingAlertConfigurationListRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectWiseTicketingAlertConfigurationListRecordsInnerWithDefaults

`func NewConnectWiseTicketingAlertConfigurationListRecordsInnerWithDefaults() *ConnectWiseTicketingAlertConfigurationListRecordsInner`

NewConnectWiseTicketingAlertConfigurationListRecordsInnerWithDefaults instantiates a new ConnectWiseTicketingAlertConfigurationListRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetCategory

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDueDays

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### GetId

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetPriority() AutotaskTicketingAlertConfigurationPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetPriorityOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) SetPriority(v AutotaskTicketingAlertConfigurationPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetShouldCreateTickets

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetShouldCreateTickets() bool`

GetShouldCreateTickets returns the ShouldCreateTickets field if non-nil, zero value otherwise.

### GetShouldCreateTicketsOk

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetShouldCreateTicketsOk() (*bool, bool)`

GetShouldCreateTicketsOk returns a tuple with the ShouldCreateTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateTickets

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) SetShouldCreateTickets(v bool)`

SetShouldCreateTickets sets ShouldCreateTickets field to given value.


### GetSource

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetSource() AutotaskTicketingAlertConfigurationPriority`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) GetSourceOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) SetSource(v AutotaskTicketingAlertConfigurationPriority)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConnectWiseTicketingAlertConfigurationListRecordsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



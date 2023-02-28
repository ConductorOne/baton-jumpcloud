# ConnectWiseTicketingAlertConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DueDays** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 
**ShouldCreateTickets** | **bool** |  | 
**Source** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 

## Methods

### NewConnectWiseTicketingAlertConfiguration

`func NewConnectWiseTicketingAlertConfiguration(shouldCreateTickets bool, ) *ConnectWiseTicketingAlertConfiguration`

NewConnectWiseTicketingAlertConfiguration instantiates a new ConnectWiseTicketingAlertConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectWiseTicketingAlertConfigurationWithDefaults

`func NewConnectWiseTicketingAlertConfigurationWithDefaults() *ConnectWiseTicketingAlertConfiguration`

NewConnectWiseTicketingAlertConfigurationWithDefaults instantiates a new ConnectWiseTicketingAlertConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ConnectWiseTicketingAlertConfiguration) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConnectWiseTicketingAlertConfiguration) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConnectWiseTicketingAlertConfiguration) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ConnectWiseTicketingAlertConfiguration) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectWiseTicketingAlertConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectWiseTicketingAlertConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectWiseTicketingAlertConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectWiseTicketingAlertConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ConnectWiseTicketingAlertConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectWiseTicketingAlertConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectWiseTicketingAlertConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectWiseTicketingAlertConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDueDays

`func (o *ConnectWiseTicketingAlertConfiguration) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *ConnectWiseTicketingAlertConfiguration) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *ConnectWiseTicketingAlertConfiguration) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.

### HasDueDays

`func (o *ConnectWiseTicketingAlertConfiguration) HasDueDays() bool`

HasDueDays returns a boolean if a field has been set.

### GetId

`func (o *ConnectWiseTicketingAlertConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectWiseTicketingAlertConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectWiseTicketingAlertConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectWiseTicketingAlertConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *ConnectWiseTicketingAlertConfiguration) GetPriority() AutotaskTicketingAlertConfigurationPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConnectWiseTicketingAlertConfiguration) GetPriorityOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConnectWiseTicketingAlertConfiguration) SetPriority(v AutotaskTicketingAlertConfigurationPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConnectWiseTicketingAlertConfiguration) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetShouldCreateTickets

`func (o *ConnectWiseTicketingAlertConfiguration) GetShouldCreateTickets() bool`

GetShouldCreateTickets returns the ShouldCreateTickets field if non-nil, zero value otherwise.

### GetShouldCreateTicketsOk

`func (o *ConnectWiseTicketingAlertConfiguration) GetShouldCreateTicketsOk() (*bool, bool)`

GetShouldCreateTicketsOk returns a tuple with the ShouldCreateTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateTickets

`func (o *ConnectWiseTicketingAlertConfiguration) SetShouldCreateTickets(v bool)`

SetShouldCreateTickets sets ShouldCreateTickets field to given value.


### GetSource

`func (o *ConnectWiseTicketingAlertConfiguration) GetSource() AutotaskTicketingAlertConfigurationPriority`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConnectWiseTicketingAlertConfiguration) GetSourceOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConnectWiseTicketingAlertConfiguration) SetSource(v AutotaskTicketingAlertConfigurationPriority)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConnectWiseTicketingAlertConfiguration) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



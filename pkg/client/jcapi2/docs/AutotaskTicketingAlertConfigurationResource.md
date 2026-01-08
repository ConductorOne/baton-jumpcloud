# AutotaskTicketingAlertConfigurationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**AutotaskTicketingAlertConfigurationPriority**](AutotaskTicketingAlertConfigurationPriority.md) |  | [optional] 

## Methods

### NewAutotaskTicketingAlertConfigurationResource

`func NewAutotaskTicketingAlertConfigurationResource() *AutotaskTicketingAlertConfigurationResource`

NewAutotaskTicketingAlertConfigurationResource instantiates a new AutotaskTicketingAlertConfigurationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotaskTicketingAlertConfigurationResourceWithDefaults

`func NewAutotaskTicketingAlertConfigurationResourceWithDefaults() *AutotaskTicketingAlertConfigurationResource`

NewAutotaskTicketingAlertConfigurationResourceWithDefaults instantiates a new AutotaskTicketingAlertConfigurationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutotaskTicketingAlertConfigurationResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutotaskTicketingAlertConfigurationResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutotaskTicketingAlertConfigurationResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AutotaskTicketingAlertConfigurationResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AutotaskTicketingAlertConfigurationResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutotaskTicketingAlertConfigurationResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutotaskTicketingAlertConfigurationResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutotaskTicketingAlertConfigurationResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *AutotaskTicketingAlertConfigurationResource) GetRole() AutotaskTicketingAlertConfigurationPriority`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AutotaskTicketingAlertConfigurationResource) GetRoleOk() (*AutotaskTicketingAlertConfigurationPriority, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AutotaskTicketingAlertConfigurationResource) SetRole(v AutotaskTicketingAlertConfigurationPriority)`

SetRole sets Role field to given value.

### HasRole

`func (o *AutotaskTicketingAlertConfigurationResource) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



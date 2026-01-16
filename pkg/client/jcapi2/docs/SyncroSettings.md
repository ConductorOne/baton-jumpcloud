# SyncroSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticTicketing** | Pointer to **bool** | Determine whether Syncro uses automatic ticketing | [optional] 

## Methods

### NewSyncroSettings

`func NewSyncroSettings() *SyncroSettings`

NewSyncroSettings instantiates a new SyncroSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncroSettingsWithDefaults

`func NewSyncroSettingsWithDefaults() *SyncroSettings`

NewSyncroSettingsWithDefaults instantiates a new SyncroSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticTicketing

`func (o *SyncroSettings) GetAutomaticTicketing() bool`

GetAutomaticTicketing returns the AutomaticTicketing field if non-nil, zero value otherwise.

### GetAutomaticTicketingOk

`func (o *SyncroSettings) GetAutomaticTicketingOk() (*bool, bool)`

GetAutomaticTicketingOk returns a tuple with the AutomaticTicketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTicketing

`func (o *SyncroSettings) SetAutomaticTicketing(v bool)`

SetAutomaticTicketing sets AutomaticTicketing field to given value.

### HasAutomaticTicketing

`func (o *SyncroSettings) HasAutomaticTicketing() bool`

HasAutomaticTicketing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



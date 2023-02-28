# OSList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**[]OSListEntriesInner**](OSListEntriesInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewOSList

`func NewOSList() *OSList`

NewOSList instantiates a new OSList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSListWithDefaults

`func NewOSListWithDefaults() *OSList`

NewOSListWithDefaults instantiates a new OSList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OSList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OSList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OSList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OSList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *OSList) GetEntries() []OSListEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *OSList) GetEntriesOk() (*[]OSListEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *OSList) SetEntries(v []OSListEntriesInner)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *OSList) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetId

`func (o *OSList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OSList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OSList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSList) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



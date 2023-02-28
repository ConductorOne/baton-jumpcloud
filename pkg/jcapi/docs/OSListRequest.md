# OSListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**[]OSListEntriesInner**](OSListEntriesInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewOSListRequest

`func NewOSListRequest() *OSListRequest`

NewOSListRequest instantiates a new OSListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSListRequestWithDefaults

`func NewOSListRequestWithDefaults() *OSListRequest`

NewOSListRequestWithDefaults instantiates a new OSListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OSListRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OSListRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OSListRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OSListRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *OSListRequest) GetEntries() []OSListEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *OSListRequest) GetEntriesOk() (*[]OSListEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *OSListRequest) SetEntries(v []OSListEntriesInner)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *OSListRequest) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetName

`func (o *OSListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSListRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OSListRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Systemslist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]System**](System.md) | The list of systems. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of systems. | [optional] 

## Methods

### NewSystemslist

`func NewSystemslist() *Systemslist`

NewSystemslist instantiates a new Systemslist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemslistWithDefaults

`func NewSystemslistWithDefaults() *Systemslist`

NewSystemslistWithDefaults instantiates a new Systemslist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *Systemslist) GetResults() []System`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Systemslist) GetResultsOk() (*[]System, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Systemslist) SetResults(v []System)`

SetResults sets Results field to given value.

### HasResults

`func (o *Systemslist) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *Systemslist) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Systemslist) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Systemslist) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Systemslist) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



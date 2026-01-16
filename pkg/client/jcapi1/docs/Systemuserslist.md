# Systemuserslist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Systemuserreturn**](Systemuserreturn.md) | The list of system users. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of system users. | [optional] 

## Methods

### NewSystemuserslist

`func NewSystemuserslist() *Systemuserslist`

NewSystemuserslist instantiates a new Systemuserslist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemuserslistWithDefaults

`func NewSystemuserslistWithDefaults() *Systemuserslist`

NewSystemuserslistWithDefaults instantiates a new Systemuserslist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *Systemuserslist) GetResults() []Systemuserreturn`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Systemuserslist) GetResultsOk() (*[]Systemuserreturn, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Systemuserslist) SetResults(v []Systemuserreturn)`

SetResults sets Results field to given value.

### HasResults

`func (o *Systemuserslist) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *Systemuserslist) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Systemuserslist) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Systemuserslist) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Systemuserslist) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



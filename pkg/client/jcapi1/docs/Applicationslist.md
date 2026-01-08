# Applicationslist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]Application**](Application.md) | The list of applications. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of applications. | [optional] 

## Methods

### NewApplicationslist

`func NewApplicationslist() *Applicationslist`

NewApplicationslist instantiates a new Applicationslist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationslistWithDefaults

`func NewApplicationslistWithDefaults() *Applicationslist`

NewApplicationslistWithDefaults instantiates a new Applicationslist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Applicationslist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Applicationslist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Applicationslist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Applicationslist) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResults

`func (o *Applicationslist) GetResults() []Application`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Applicationslist) GetResultsOk() (*[]Application, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Applicationslist) SetResults(v []Application)`

SetResults sets Results field to given value.

### HasResults

`func (o *Applicationslist) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *Applicationslist) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Applicationslist) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Applicationslist) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Applicationslist) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



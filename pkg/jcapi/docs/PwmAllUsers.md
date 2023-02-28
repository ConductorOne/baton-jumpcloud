# PwmAllUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]PwmAllUsersResultsInner**](PwmAllUsersResultsInner.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewPwmAllUsers

`func NewPwmAllUsers(results []PwmAllUsersResultsInner, totalCount int32, ) *PwmAllUsers`

NewPwmAllUsers instantiates a new PwmAllUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPwmAllUsersWithDefaults

`func NewPwmAllUsersWithDefaults() *PwmAllUsers`

NewPwmAllUsersWithDefaults instantiates a new PwmAllUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PwmAllUsers) GetResults() []PwmAllUsersResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PwmAllUsers) GetResultsOk() (*[]PwmAllUsersResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PwmAllUsers) SetResults(v []PwmAllUsersResultsInner)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *PwmAllUsers) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PwmAllUsers) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PwmAllUsers) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



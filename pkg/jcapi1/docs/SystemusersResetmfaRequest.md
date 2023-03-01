# SystemusersResetmfaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclusion** | Pointer to **bool** |  | [optional] 
**ExclusionDays** | Pointer to **float32** |  | [optional] 
**ExclusionUntil** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSystemusersResetmfaRequest

`func NewSystemusersResetmfaRequest() *SystemusersResetmfaRequest`

NewSystemusersResetmfaRequest instantiates a new SystemusersResetmfaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemusersResetmfaRequestWithDefaults

`func NewSystemusersResetmfaRequestWithDefaults() *SystemusersResetmfaRequest`

NewSystemusersResetmfaRequestWithDefaults instantiates a new SystemusersResetmfaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclusion

`func (o *SystemusersResetmfaRequest) GetExclusion() bool`

GetExclusion returns the Exclusion field if non-nil, zero value otherwise.

### GetExclusionOk

`func (o *SystemusersResetmfaRequest) GetExclusionOk() (*bool, bool)`

GetExclusionOk returns a tuple with the Exclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusion

`func (o *SystemusersResetmfaRequest) SetExclusion(v bool)`

SetExclusion sets Exclusion field to given value.

### HasExclusion

`func (o *SystemusersResetmfaRequest) HasExclusion() bool`

HasExclusion returns a boolean if a field has been set.

### GetExclusionDays

`func (o *SystemusersResetmfaRequest) GetExclusionDays() float32`

GetExclusionDays returns the ExclusionDays field if non-nil, zero value otherwise.

### GetExclusionDaysOk

`func (o *SystemusersResetmfaRequest) GetExclusionDaysOk() (*float32, bool)`

GetExclusionDaysOk returns a tuple with the ExclusionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionDays

`func (o *SystemusersResetmfaRequest) SetExclusionDays(v float32)`

SetExclusionDays sets ExclusionDays field to given value.

### HasExclusionDays

`func (o *SystemusersResetmfaRequest) HasExclusionDays() bool`

HasExclusionDays returns a boolean if a field has been set.

### GetExclusionUntil

`func (o *SystemusersResetmfaRequest) GetExclusionUntil() time.Time`

GetExclusionUntil returns the ExclusionUntil field if non-nil, zero value otherwise.

### GetExclusionUntilOk

`func (o *SystemusersResetmfaRequest) GetExclusionUntilOk() (*time.Time, bool)`

GetExclusionUntilOk returns a tuple with the ExclusionUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionUntil

`func (o *SystemusersResetmfaRequest) SetExclusionUntil(v time.Time)`

SetExclusionUntil sets ExclusionUntil field to given value.

### HasExclusionUntil

`func (o *SystemusersResetmfaRequest) HasExclusionUntil() bool`

HasExclusionUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



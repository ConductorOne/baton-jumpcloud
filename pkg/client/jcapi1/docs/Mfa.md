# Mfa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** |  | [optional] 
**Exclusion** | Pointer to **bool** |  | [optional] 
**ExclusionDays** | Pointer to **int32** |  | [optional] 
**ExclusionUntil** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMfa

`func NewMfa() *Mfa`

NewMfa instantiates a new Mfa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaWithDefaults

`func NewMfaWithDefaults() *Mfa`

NewMfaWithDefaults instantiates a new Mfa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *Mfa) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *Mfa) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *Mfa) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *Mfa) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetExclusion

`func (o *Mfa) GetExclusion() bool`

GetExclusion returns the Exclusion field if non-nil, zero value otherwise.

### GetExclusionOk

`func (o *Mfa) GetExclusionOk() (*bool, bool)`

GetExclusionOk returns a tuple with the Exclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusion

`func (o *Mfa) SetExclusion(v bool)`

SetExclusion sets Exclusion field to given value.

### HasExclusion

`func (o *Mfa) HasExclusion() bool`

HasExclusion returns a boolean if a field has been set.

### GetExclusionDays

`func (o *Mfa) GetExclusionDays() int32`

GetExclusionDays returns the ExclusionDays field if non-nil, zero value otherwise.

### GetExclusionDaysOk

`func (o *Mfa) GetExclusionDaysOk() (*int32, bool)`

GetExclusionDaysOk returns a tuple with the ExclusionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionDays

`func (o *Mfa) SetExclusionDays(v int32)`

SetExclusionDays sets ExclusionDays field to given value.

### HasExclusionDays

`func (o *Mfa) HasExclusionDays() bool`

HasExclusionDays returns a boolean if a field has been set.

### GetExclusionUntil

`func (o *Mfa) GetExclusionUntil() time.Time`

GetExclusionUntil returns the ExclusionUntil field if non-nil, zero value otherwise.

### GetExclusionUntilOk

`func (o *Mfa) GetExclusionUntilOk() (*time.Time, bool)`

GetExclusionUntilOk returns a tuple with the ExclusionUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionUntil

`func (o *Mfa) SetExclusionUntil(v time.Time)`

SetExclusionUntil sets ExclusionUntil field to given value.

### HasExclusionUntil

`func (o *Mfa) HasExclusionUntil() bool`

HasExclusionUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PolicyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Details pertaining to the policy result. | [optional] 
**EndedAt** | Pointer to **time.Time** | The end of the policy application. | [optional] 
**ExitStatus** | Pointer to **int32** | The 32-bit unsigned exit status from the applying the policy. | [optional] 
**Id** | Pointer to **string** | ObjectId uniquely identifying a Policy Result. | [optional] 
**PolicyID** | Pointer to **string** | ObjectId uniquely identifying the parent Policy. | [optional] 
**StartedAt** | Pointer to **time.Time** | The start of the policy application. | [optional] 
**State** | Pointer to **string** | Enumeration describing the state of the policy. Success, failed, or pending. | [optional] 
**StdErr** | Pointer to **string** | The STDERR output from applying the policy. | [optional] 
**StdOut** | Pointer to **string** | The STDOUT output from applying the policy. | [optional] 
**Success** | Pointer to **bool** | True if the policy was successfully applied; false otherwise. | [optional] 
**SystemID** | Pointer to **string** | ObjectId uniquely identifying the parent System. | [optional] 

## Methods

### NewPolicyResult

`func NewPolicyResult() *PolicyResult`

NewPolicyResult instantiates a new PolicyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyResultWithDefaults

`func NewPolicyResultWithDefaults() *PolicyResult`

NewPolicyResultWithDefaults instantiates a new PolicyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *PolicyResult) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *PolicyResult) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *PolicyResult) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *PolicyResult) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetEndedAt

`func (o *PolicyResult) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *PolicyResult) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *PolicyResult) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *PolicyResult) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetExitStatus

`func (o *PolicyResult) GetExitStatus() int32`

GetExitStatus returns the ExitStatus field if non-nil, zero value otherwise.

### GetExitStatusOk

`func (o *PolicyResult) GetExitStatusOk() (*int32, bool)`

GetExitStatusOk returns a tuple with the ExitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitStatus

`func (o *PolicyResult) SetExitStatus(v int32)`

SetExitStatus sets ExitStatus field to given value.

### HasExitStatus

`func (o *PolicyResult) HasExitStatus() bool`

HasExitStatus returns a boolean if a field has been set.

### GetId

`func (o *PolicyResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicyID

`func (o *PolicyResult) GetPolicyID() string`

GetPolicyID returns the PolicyID field if non-nil, zero value otherwise.

### GetPolicyIDOk

`func (o *PolicyResult) GetPolicyIDOk() (*string, bool)`

GetPolicyIDOk returns a tuple with the PolicyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyID

`func (o *PolicyResult) SetPolicyID(v string)`

SetPolicyID sets PolicyID field to given value.

### HasPolicyID

`func (o *PolicyResult) HasPolicyID() bool`

HasPolicyID returns a boolean if a field has been set.

### GetStartedAt

`func (o *PolicyResult) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PolicyResult) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PolicyResult) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *PolicyResult) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetState

`func (o *PolicyResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PolicyResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PolicyResult) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PolicyResult) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStdErr

`func (o *PolicyResult) GetStdErr() string`

GetStdErr returns the StdErr field if non-nil, zero value otherwise.

### GetStdErrOk

`func (o *PolicyResult) GetStdErrOk() (*string, bool)`

GetStdErrOk returns a tuple with the StdErr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdErr

`func (o *PolicyResult) SetStdErr(v string)`

SetStdErr sets StdErr field to given value.

### HasStdErr

`func (o *PolicyResult) HasStdErr() bool`

HasStdErr returns a boolean if a field has been set.

### GetStdOut

`func (o *PolicyResult) GetStdOut() string`

GetStdOut returns the StdOut field if non-nil, zero value otherwise.

### GetStdOutOk

`func (o *PolicyResult) GetStdOutOk() (*string, bool)`

GetStdOutOk returns a tuple with the StdOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdOut

`func (o *PolicyResult) SetStdOut(v string)`

SetStdOut sets StdOut field to given value.

### HasStdOut

`func (o *PolicyResult) HasStdOut() bool`

HasStdOut returns a boolean if a field has been set.

### GetSuccess

`func (o *PolicyResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PolicyResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PolicyResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PolicyResult) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSystemID

`func (o *PolicyResult) GetSystemID() string`

GetSystemID returns the SystemID field if non-nil, zero value otherwise.

### GetSystemIDOk

`func (o *PolicyResult) GetSystemIDOk() (*string, bool)`

GetSystemIDOk returns a tuple with the SystemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemID

`func (o *PolicyResult) SetSystemID(v string)`

SetSystemID sets SystemID field to given value.

### HasSystemID

`func (o *PolicyResult) HasSystemID() bool`

HasSystemID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



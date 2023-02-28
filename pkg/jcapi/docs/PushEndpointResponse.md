# PushEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**PushEndpointResponseDevice**](PushEndpointResponseDevice.md) |  | [optional] 
**EnrollmentDate** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUsedDate** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewPushEndpointResponse

`func NewPushEndpointResponse() *PushEndpointResponse`

NewPushEndpointResponse instantiates a new PushEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushEndpointResponseWithDefaults

`func NewPushEndpointResponseWithDefaults() *PushEndpointResponse`

NewPushEndpointResponseWithDefaults instantiates a new PushEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PushEndpointResponse) GetDevice() PushEndpointResponseDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PushEndpointResponse) GetDeviceOk() (*PushEndpointResponseDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PushEndpointResponse) SetDevice(v PushEndpointResponseDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PushEndpointResponse) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetEnrollmentDate

`func (o *PushEndpointResponse) GetEnrollmentDate() time.Time`

GetEnrollmentDate returns the EnrollmentDate field if non-nil, zero value otherwise.

### GetEnrollmentDateOk

`func (o *PushEndpointResponse) GetEnrollmentDateOk() (*time.Time, bool)`

GetEnrollmentDateOk returns a tuple with the EnrollmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentDate

`func (o *PushEndpointResponse) SetEnrollmentDate(v time.Time)`

SetEnrollmentDate sets EnrollmentDate field to given value.

### HasEnrollmentDate

`func (o *PushEndpointResponse) HasEnrollmentDate() bool`

HasEnrollmentDate returns a boolean if a field has been set.

### GetId

`func (o *PushEndpointResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PushEndpointResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PushEndpointResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PushEndpointResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedDate

`func (o *PushEndpointResponse) GetLastUsedDate() time.Time`

GetLastUsedDate returns the LastUsedDate field if non-nil, zero value otherwise.

### GetLastUsedDateOk

`func (o *PushEndpointResponse) GetLastUsedDateOk() (*time.Time, bool)`

GetLastUsedDateOk returns a tuple with the LastUsedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedDate

`func (o *PushEndpointResponse) SetLastUsedDate(v time.Time)`

SetLastUsedDate sets LastUsedDate field to given value.

### HasLastUsedDate

`func (o *PushEndpointResponse) HasLastUsedDate() bool`

HasLastUsedDate returns a boolean if a field has been set.

### GetName

`func (o *PushEndpointResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PushEndpointResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PushEndpointResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PushEndpointResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *PushEndpointResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PushEndpointResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PushEndpointResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PushEndpointResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IntegrationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]Integration**](Integration.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewIntegrationsResponse

`func NewIntegrationsResponse() *IntegrationsResponse`

NewIntegrationsResponse instantiates a new IntegrationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsResponseWithDefaults

`func NewIntegrationsResponseWithDefaults() *IntegrationsResponse`

NewIntegrationsResponseWithDefaults instantiates a new IntegrationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *IntegrationsResponse) GetRecords() []Integration`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *IntegrationsResponse) GetRecordsOk() (*[]Integration, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *IntegrationsResponse) SetRecords(v []Integration)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *IntegrationsResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotalCount

`func (o *IntegrationsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *IntegrationsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *IntegrationsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *IntegrationsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



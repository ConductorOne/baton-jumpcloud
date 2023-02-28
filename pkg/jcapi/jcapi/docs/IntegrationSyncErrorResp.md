# IntegrationSyncErrorResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | [**[]IntegrationSyncError**](IntegrationSyncError.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewIntegrationSyncErrorResp

`func NewIntegrationSyncErrorResp(records []IntegrationSyncError, totalCount int32, ) *IntegrationSyncErrorResp`

NewIntegrationSyncErrorResp instantiates a new IntegrationSyncErrorResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSyncErrorRespWithDefaults

`func NewIntegrationSyncErrorRespWithDefaults() *IntegrationSyncErrorResp`

NewIntegrationSyncErrorRespWithDefaults instantiates a new IntegrationSyncErrorResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *IntegrationSyncErrorResp) GetRecords() []IntegrationSyncError`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *IntegrationSyncErrorResp) GetRecordsOk() (*[]IntegrationSyncError, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *IntegrationSyncErrorResp) SetRecords(v []IntegrationSyncError)`

SetRecords sets Records field to given value.


### GetTotalCount

`func (o *IntegrationSyncErrorResp) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *IntegrationSyncErrorResp) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *IntegrationSyncErrorResp) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ProvidersListAdministrators200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Administrator**](Administrator.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewProvidersListAdministrators200Response

`func NewProvidersListAdministrators200Response() *ProvidersListAdministrators200Response`

NewProvidersListAdministrators200Response instantiates a new ProvidersListAdministrators200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvidersListAdministrators200ResponseWithDefaults

`func NewProvidersListAdministrators200ResponseWithDefaults() *ProvidersListAdministrators200Response`

NewProvidersListAdministrators200ResponseWithDefaults instantiates a new ProvidersListAdministrators200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ProvidersListAdministrators200Response) GetResults() []Administrator`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ProvidersListAdministrators200Response) GetResultsOk() (*[]Administrator, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ProvidersListAdministrators200Response) SetResults(v []Administrator)`

SetResults sets Results field to given value.

### HasResults

`func (o *ProvidersListAdministrators200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *ProvidersListAdministrators200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProvidersListAdministrators200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProvidersListAdministrators200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ProvidersListAdministrators200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



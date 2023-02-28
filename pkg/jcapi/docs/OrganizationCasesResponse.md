# OrganizationCasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]OrganizationCase**](OrganizationCase.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrganizationCasesResponse

`func NewOrganizationCasesResponse() *OrganizationCasesResponse`

NewOrganizationCasesResponse instantiates a new OrganizationCasesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCasesResponseWithDefaults

`func NewOrganizationCasesResponseWithDefaults() *OrganizationCasesResponse`

NewOrganizationCasesResponseWithDefaults instantiates a new OrganizationCasesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *OrganizationCasesResponse) GetResults() []OrganizationCase`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OrganizationCasesResponse) GetResultsOk() (*[]OrganizationCase, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OrganizationCasesResponse) SetResults(v []OrganizationCase)`

SetResults sets Results field to given value.

### HasResults

`func (o *OrganizationCasesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *OrganizationCasesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OrganizationCasesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OrganizationCasesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OrganizationCasesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



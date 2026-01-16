# AutotaskSettingsPatchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticTicketing** | Pointer to **bool** | Determine whether Autotask uses automatic ticketing | [optional] 
**CompanyTypeIds** | Pointer to **[]int32** | The array of Autotask companyType IDs applicable to the Provider. | [optional] 

## Methods

### NewAutotaskSettingsPatchReq

`func NewAutotaskSettingsPatchReq() *AutotaskSettingsPatchReq`

NewAutotaskSettingsPatchReq instantiates a new AutotaskSettingsPatchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotaskSettingsPatchReqWithDefaults

`func NewAutotaskSettingsPatchReqWithDefaults() *AutotaskSettingsPatchReq`

NewAutotaskSettingsPatchReqWithDefaults instantiates a new AutotaskSettingsPatchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticTicketing

`func (o *AutotaskSettingsPatchReq) GetAutomaticTicketing() bool`

GetAutomaticTicketing returns the AutomaticTicketing field if non-nil, zero value otherwise.

### GetAutomaticTicketingOk

`func (o *AutotaskSettingsPatchReq) GetAutomaticTicketingOk() (*bool, bool)`

GetAutomaticTicketingOk returns a tuple with the AutomaticTicketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTicketing

`func (o *AutotaskSettingsPatchReq) SetAutomaticTicketing(v bool)`

SetAutomaticTicketing sets AutomaticTicketing field to given value.

### HasAutomaticTicketing

`func (o *AutotaskSettingsPatchReq) HasAutomaticTicketing() bool`

HasAutomaticTicketing returns a boolean if a field has been set.

### GetCompanyTypeIds

`func (o *AutotaskSettingsPatchReq) GetCompanyTypeIds() []int32`

GetCompanyTypeIds returns the CompanyTypeIds field if non-nil, zero value otherwise.

### GetCompanyTypeIdsOk

`func (o *AutotaskSettingsPatchReq) GetCompanyTypeIdsOk() (*[]int32, bool)`

GetCompanyTypeIdsOk returns a tuple with the CompanyTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyTypeIds

`func (o *AutotaskSettingsPatchReq) SetCompanyTypeIds(v []int32)`

SetCompanyTypeIds sets CompanyTypeIds field to given value.

### HasCompanyTypeIds

`func (o *AutotaskSettingsPatchReq) HasCompanyTypeIds() bool`

HasCompanyTypeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



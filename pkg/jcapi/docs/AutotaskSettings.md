# AutotaskSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticTicketing** | Pointer to **bool** | Determine whether Autotask uses automatic ticketing | [optional] 
**CompanyTypeIds** | Pointer to **[]int32** | The array of Autotask companyType IDs applicable to the Provider. | [optional] 

## Methods

### NewAutotaskSettings

`func NewAutotaskSettings() *AutotaskSettings`

NewAutotaskSettings instantiates a new AutotaskSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotaskSettingsWithDefaults

`func NewAutotaskSettingsWithDefaults() *AutotaskSettings`

NewAutotaskSettingsWithDefaults instantiates a new AutotaskSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticTicketing

`func (o *AutotaskSettings) GetAutomaticTicketing() bool`

GetAutomaticTicketing returns the AutomaticTicketing field if non-nil, zero value otherwise.

### GetAutomaticTicketingOk

`func (o *AutotaskSettings) GetAutomaticTicketingOk() (*bool, bool)`

GetAutomaticTicketingOk returns a tuple with the AutomaticTicketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTicketing

`func (o *AutotaskSettings) SetAutomaticTicketing(v bool)`

SetAutomaticTicketing sets AutomaticTicketing field to given value.

### HasAutomaticTicketing

`func (o *AutotaskSettings) HasAutomaticTicketing() bool`

HasAutomaticTicketing returns a boolean if a field has been set.

### GetCompanyTypeIds

`func (o *AutotaskSettings) GetCompanyTypeIds() []int32`

GetCompanyTypeIds returns the CompanyTypeIds field if non-nil, zero value otherwise.

### GetCompanyTypeIdsOk

`func (o *AutotaskSettings) GetCompanyTypeIdsOk() (*[]int32, bool)`

GetCompanyTypeIdsOk returns a tuple with the CompanyTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyTypeIds

`func (o *AutotaskSettings) SetCompanyTypeIds(v []int32)`

SetCompanyTypeIds sets CompanyTypeIds field to given value.

### HasCompanyTypeIds

`func (o *AutotaskSettings) HasCompanyTypeIds() bool`

HasCompanyTypeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



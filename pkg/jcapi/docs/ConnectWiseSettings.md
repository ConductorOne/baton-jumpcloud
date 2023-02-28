# ConnectWiseSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticTicketing** | Pointer to **bool** | Determine whether ConnectWise uses automatic ticketing | [optional] 
**CompanyTypeIds** | Pointer to **[]int32** | The array of ConnectWise companyType IDs applicable to the Provider. | [optional] 

## Methods

### NewConnectWiseSettings

`func NewConnectWiseSettings() *ConnectWiseSettings`

NewConnectWiseSettings instantiates a new ConnectWiseSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectWiseSettingsWithDefaults

`func NewConnectWiseSettingsWithDefaults() *ConnectWiseSettings`

NewConnectWiseSettingsWithDefaults instantiates a new ConnectWiseSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticTicketing

`func (o *ConnectWiseSettings) GetAutomaticTicketing() bool`

GetAutomaticTicketing returns the AutomaticTicketing field if non-nil, zero value otherwise.

### GetAutomaticTicketingOk

`func (o *ConnectWiseSettings) GetAutomaticTicketingOk() (*bool, bool)`

GetAutomaticTicketingOk returns a tuple with the AutomaticTicketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTicketing

`func (o *ConnectWiseSettings) SetAutomaticTicketing(v bool)`

SetAutomaticTicketing sets AutomaticTicketing field to given value.

### HasAutomaticTicketing

`func (o *ConnectWiseSettings) HasAutomaticTicketing() bool`

HasAutomaticTicketing returns a boolean if a field has been set.

### GetCompanyTypeIds

`func (o *ConnectWiseSettings) GetCompanyTypeIds() []int32`

GetCompanyTypeIds returns the CompanyTypeIds field if non-nil, zero value otherwise.

### GetCompanyTypeIdsOk

`func (o *ConnectWiseSettings) GetCompanyTypeIdsOk() (*[]int32, bool)`

GetCompanyTypeIdsOk returns a tuple with the CompanyTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyTypeIds

`func (o *ConnectWiseSettings) SetCompanyTypeIds(v []int32)`

SetCompanyTypeIds sets CompanyTypeIds field to given value.

### HasCompanyTypeIds

`func (o *ConnectWiseSettings) HasCompanyTypeIds() bool`

HasCompanyTypeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



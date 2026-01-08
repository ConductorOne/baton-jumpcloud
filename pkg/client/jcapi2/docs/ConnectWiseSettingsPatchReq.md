# ConnectWiseSettingsPatchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticTicketing** | Pointer to **bool** | Determine whether ConnectWise uses automatic ticketing | [optional] 
**CompanyTypeIds** | Pointer to **[]int32** | The array of ConnectWise companyType IDs applicable to the Provider. | [optional] 

## Methods

### NewConnectWiseSettingsPatchReq

`func NewConnectWiseSettingsPatchReq() *ConnectWiseSettingsPatchReq`

NewConnectWiseSettingsPatchReq instantiates a new ConnectWiseSettingsPatchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectWiseSettingsPatchReqWithDefaults

`func NewConnectWiseSettingsPatchReqWithDefaults() *ConnectWiseSettingsPatchReq`

NewConnectWiseSettingsPatchReqWithDefaults instantiates a new ConnectWiseSettingsPatchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticTicketing

`func (o *ConnectWiseSettingsPatchReq) GetAutomaticTicketing() bool`

GetAutomaticTicketing returns the AutomaticTicketing field if non-nil, zero value otherwise.

### GetAutomaticTicketingOk

`func (o *ConnectWiseSettingsPatchReq) GetAutomaticTicketingOk() (*bool, bool)`

GetAutomaticTicketingOk returns a tuple with the AutomaticTicketing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticTicketing

`func (o *ConnectWiseSettingsPatchReq) SetAutomaticTicketing(v bool)`

SetAutomaticTicketing sets AutomaticTicketing field to given value.

### HasAutomaticTicketing

`func (o *ConnectWiseSettingsPatchReq) HasAutomaticTicketing() bool`

HasAutomaticTicketing returns a boolean if a field has been set.

### GetCompanyTypeIds

`func (o *ConnectWiseSettingsPatchReq) GetCompanyTypeIds() []int32`

GetCompanyTypeIds returns the CompanyTypeIds field if non-nil, zero value otherwise.

### GetCompanyTypeIdsOk

`func (o *ConnectWiseSettingsPatchReq) GetCompanyTypeIdsOk() (*[]int32, bool)`

GetCompanyTypeIdsOk returns a tuple with the CompanyTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyTypeIds

`func (o *ConnectWiseSettingsPatchReq) SetCompanyTypeIds(v []int32)`

SetCompanyTypeIds sets CompanyTypeIds field to given value.

### HasCompanyTypeIds

`func (o *ConnectWiseSettingsPatchReq) HasCompanyTypeIds() bool`

HasCompanyTypeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



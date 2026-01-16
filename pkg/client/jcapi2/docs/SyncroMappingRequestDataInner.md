# SyncroMappingRequestDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingConfigurations** | Pointer to [**SyncroMappingRequestDataInnerBillingConfigurations**](SyncroMappingRequestDataInnerBillingConfigurations.md) |  | [optional] 
**Company** | [**ConnectWiseMappingRequestDataInnerCompany**](ConnectWiseMappingRequestDataInnerCompany.md) |  | 
**Delete** | Pointer to **bool** |  | [optional] 
**Organization** | [**ConnectWiseMappingRequestDataInnerOrganization**](ConnectWiseMappingRequestDataInnerOrganization.md) |  | 

## Methods

### NewSyncroMappingRequestDataInner

`func NewSyncroMappingRequestDataInner(company ConnectWiseMappingRequestDataInnerCompany, organization ConnectWiseMappingRequestDataInnerOrganization, ) *SyncroMappingRequestDataInner`

NewSyncroMappingRequestDataInner instantiates a new SyncroMappingRequestDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncroMappingRequestDataInnerWithDefaults

`func NewSyncroMappingRequestDataInnerWithDefaults() *SyncroMappingRequestDataInner`

NewSyncroMappingRequestDataInnerWithDefaults instantiates a new SyncroMappingRequestDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingConfigurations

`func (o *SyncroMappingRequestDataInner) GetBillingConfigurations() SyncroMappingRequestDataInnerBillingConfigurations`

GetBillingConfigurations returns the BillingConfigurations field if non-nil, zero value otherwise.

### GetBillingConfigurationsOk

`func (o *SyncroMappingRequestDataInner) GetBillingConfigurationsOk() (*SyncroMappingRequestDataInnerBillingConfigurations, bool)`

GetBillingConfigurationsOk returns a tuple with the BillingConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingConfigurations

`func (o *SyncroMappingRequestDataInner) SetBillingConfigurations(v SyncroMappingRequestDataInnerBillingConfigurations)`

SetBillingConfigurations sets BillingConfigurations field to given value.

### HasBillingConfigurations

`func (o *SyncroMappingRequestDataInner) HasBillingConfigurations() bool`

HasBillingConfigurations returns a boolean if a field has been set.

### GetCompany

`func (o *SyncroMappingRequestDataInner) GetCompany() ConnectWiseMappingRequestDataInnerCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SyncroMappingRequestDataInner) GetCompanyOk() (*ConnectWiseMappingRequestDataInnerCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SyncroMappingRequestDataInner) SetCompany(v ConnectWiseMappingRequestDataInnerCompany)`

SetCompany sets Company field to given value.


### GetDelete

`func (o *SyncroMappingRequestDataInner) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *SyncroMappingRequestDataInner) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *SyncroMappingRequestDataInner) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *SyncroMappingRequestDataInner) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetOrganization

`func (o *SyncroMappingRequestDataInner) GetOrganization() ConnectWiseMappingRequestDataInnerOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SyncroMappingRequestDataInner) GetOrganizationOk() (*ConnectWiseMappingRequestDataInnerOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SyncroMappingRequestDataInner) SetOrganization(v ConnectWiseMappingRequestDataInnerOrganization)`

SetOrganization sets Organization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



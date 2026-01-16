# AutotaskMappingRequestDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | [**AutotaskMappingRequestCompany**](AutotaskMappingRequestCompany.md) |  | 
**Contract** | Pointer to **map[string]interface{}** |  | [optional] 
**Delete** | Pointer to **bool** |  | [optional] 
**Organization** | [**AutotaskMappingRequestOrganization**](AutotaskMappingRequestOrganization.md) |  | 
**Service** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAutotaskMappingRequestDataInner

`func NewAutotaskMappingRequestDataInner(company AutotaskMappingRequestCompany, organization AutotaskMappingRequestOrganization, ) *AutotaskMappingRequestDataInner`

NewAutotaskMappingRequestDataInner instantiates a new AutotaskMappingRequestDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotaskMappingRequestDataInnerWithDefaults

`func NewAutotaskMappingRequestDataInnerWithDefaults() *AutotaskMappingRequestDataInner`

NewAutotaskMappingRequestDataInnerWithDefaults instantiates a new AutotaskMappingRequestDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *AutotaskMappingRequestDataInner) GetCompany() AutotaskMappingRequestCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AutotaskMappingRequestDataInner) GetCompanyOk() (*AutotaskMappingRequestCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AutotaskMappingRequestDataInner) SetCompany(v AutotaskMappingRequestCompany)`

SetCompany sets Company field to given value.


### GetContract

`func (o *AutotaskMappingRequestDataInner) GetContract() map[string]interface{}`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *AutotaskMappingRequestDataInner) GetContractOk() (*map[string]interface{}, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *AutotaskMappingRequestDataInner) SetContract(v map[string]interface{})`

SetContract sets Contract field to given value.

### HasContract

`func (o *AutotaskMappingRequestDataInner) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetDelete

`func (o *AutotaskMappingRequestDataInner) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *AutotaskMappingRequestDataInner) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *AutotaskMappingRequestDataInner) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *AutotaskMappingRequestDataInner) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetOrganization

`func (o *AutotaskMappingRequestDataInner) GetOrganization() AutotaskMappingRequestOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AutotaskMappingRequestDataInner) GetOrganizationOk() (*AutotaskMappingRequestOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AutotaskMappingRequestDataInner) SetOrganization(v AutotaskMappingRequestOrganization)`

SetOrganization sets Organization field to given value.


### GetService

`func (o *AutotaskMappingRequestDataInner) GetService() map[string]interface{}`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AutotaskMappingRequestDataInner) GetServiceOk() (*map[string]interface{}, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AutotaskMappingRequestDataInner) SetService(v map[string]interface{})`

SetService sets Service field to given value.

### HasService

`func (o *AutotaskMappingRequestDataInner) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



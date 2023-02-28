# ConnectWiseMappingRequestDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addition** | Pointer to **map[string]interface{}** |  | [optional] 
**Agreement** | Pointer to **map[string]interface{}** |  | [optional] 
**Company** | [**ConnectWiseMappingRequestDataInnerCompany**](ConnectWiseMappingRequestDataInnerCompany.md) |  | 
**Delete** | Pointer to **bool** |  | [optional] 
**Organization** | [**ConnectWiseMappingRequestDataInnerOrganization**](ConnectWiseMappingRequestDataInnerOrganization.md) |  | 

## Methods

### NewConnectWiseMappingRequestDataInner

`func NewConnectWiseMappingRequestDataInner(company ConnectWiseMappingRequestDataInnerCompany, organization ConnectWiseMappingRequestDataInnerOrganization, ) *ConnectWiseMappingRequestDataInner`

NewConnectWiseMappingRequestDataInner instantiates a new ConnectWiseMappingRequestDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectWiseMappingRequestDataInnerWithDefaults

`func NewConnectWiseMappingRequestDataInnerWithDefaults() *ConnectWiseMappingRequestDataInner`

NewConnectWiseMappingRequestDataInnerWithDefaults instantiates a new ConnectWiseMappingRequestDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddition

`func (o *ConnectWiseMappingRequestDataInner) GetAddition() map[string]interface{}`

GetAddition returns the Addition field if non-nil, zero value otherwise.

### GetAdditionOk

`func (o *ConnectWiseMappingRequestDataInner) GetAdditionOk() (*map[string]interface{}, bool)`

GetAdditionOk returns a tuple with the Addition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddition

`func (o *ConnectWiseMappingRequestDataInner) SetAddition(v map[string]interface{})`

SetAddition sets Addition field to given value.

### HasAddition

`func (o *ConnectWiseMappingRequestDataInner) HasAddition() bool`

HasAddition returns a boolean if a field has been set.

### GetAgreement

`func (o *ConnectWiseMappingRequestDataInner) GetAgreement() map[string]interface{}`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ConnectWiseMappingRequestDataInner) GetAgreementOk() (*map[string]interface{}, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ConnectWiseMappingRequestDataInner) SetAgreement(v map[string]interface{})`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ConnectWiseMappingRequestDataInner) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetCompany

`func (o *ConnectWiseMappingRequestDataInner) GetCompany() ConnectWiseMappingRequestDataInnerCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ConnectWiseMappingRequestDataInner) GetCompanyOk() (*ConnectWiseMappingRequestDataInnerCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ConnectWiseMappingRequestDataInner) SetCompany(v ConnectWiseMappingRequestDataInnerCompany)`

SetCompany sets Company field to given value.


### GetDelete

`func (o *ConnectWiseMappingRequestDataInner) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ConnectWiseMappingRequestDataInner) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ConnectWiseMappingRequestDataInner) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ConnectWiseMappingRequestDataInner) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetOrganization

`func (o *ConnectWiseMappingRequestDataInner) GetOrganization() ConnectWiseMappingRequestDataInnerOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ConnectWiseMappingRequestDataInner) GetOrganizationOk() (*ConnectWiseMappingRequestDataInnerOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ConnectWiseMappingRequestDataInner) SetOrganization(v ConnectWiseMappingRequestDataInnerOrganization)`

SetOrganization sets Organization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



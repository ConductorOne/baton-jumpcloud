# Organizationentitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingModel** | Pointer to **string** |  | [optional] 
**EntitlementProducts** | Pointer to [**[]OrganizationentitlementEntitlementProductsInner**](OrganizationentitlementEntitlementProductsInner.md) |  | [optional] 
**IsManuallyBilled** | Pointer to **bool** |  | [optional] 
**PricePerUserSum** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrganizationentitlement

`func NewOrganizationentitlement() *Organizationentitlement`

NewOrganizationentitlement instantiates a new Organizationentitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationentitlementWithDefaults

`func NewOrganizationentitlementWithDefaults() *Organizationentitlement`

NewOrganizationentitlementWithDefaults instantiates a new Organizationentitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingModel

`func (o *Organizationentitlement) GetBillingModel() string`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *Organizationentitlement) GetBillingModelOk() (*string, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *Organizationentitlement) SetBillingModel(v string)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *Organizationentitlement) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetEntitlementProducts

`func (o *Organizationentitlement) GetEntitlementProducts() []OrganizationentitlementEntitlementProductsInner`

GetEntitlementProducts returns the EntitlementProducts field if non-nil, zero value otherwise.

### GetEntitlementProductsOk

`func (o *Organizationentitlement) GetEntitlementProductsOk() (*[]OrganizationentitlementEntitlementProductsInner, bool)`

GetEntitlementProductsOk returns a tuple with the EntitlementProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementProducts

`func (o *Organizationentitlement) SetEntitlementProducts(v []OrganizationentitlementEntitlementProductsInner)`

SetEntitlementProducts sets EntitlementProducts field to given value.

### HasEntitlementProducts

`func (o *Organizationentitlement) HasEntitlementProducts() bool`

HasEntitlementProducts returns a boolean if a field has been set.

### GetIsManuallyBilled

`func (o *Organizationentitlement) GetIsManuallyBilled() bool`

GetIsManuallyBilled returns the IsManuallyBilled field if non-nil, zero value otherwise.

### GetIsManuallyBilledOk

`func (o *Organizationentitlement) GetIsManuallyBilledOk() (*bool, bool)`

GetIsManuallyBilledOk returns a tuple with the IsManuallyBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyBilled

`func (o *Organizationentitlement) SetIsManuallyBilled(v bool)`

SetIsManuallyBilled sets IsManuallyBilled field to given value.

### HasIsManuallyBilled

`func (o *Organizationentitlement) HasIsManuallyBilled() bool`

HasIsManuallyBilled returns a boolean if a field has been set.

### GetPricePerUserSum

`func (o *Organizationentitlement) GetPricePerUserSum() int32`

GetPricePerUserSum returns the PricePerUserSum field if non-nil, zero value otherwise.

### GetPricePerUserSumOk

`func (o *Organizationentitlement) GetPricePerUserSumOk() (*int32, bool)`

GetPricePerUserSumOk returns a tuple with the PricePerUserSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUserSum

`func (o *Organizationentitlement) SetPricePerUserSum(v int32)`

SetPricePerUserSum sets PricePerUserSum field to given value.

### HasPricePerUserSum

`func (o *Organizationentitlement) HasPricePerUserSum() bool`

HasPricePerUserSum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



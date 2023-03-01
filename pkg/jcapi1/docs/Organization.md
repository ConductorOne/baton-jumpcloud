# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AccountsReceivable** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Entitlement** | Pointer to [**Organizationentitlement**](Organizationentitlement.md) |  | [optional] 
**HasCreditCard** | Pointer to **bool** |  | [optional] 
**HasStripeCustomerId** | Pointer to **bool** |  | [optional] 
**LastEstimateCalculationTimeStamp** | Pointer to **string** |  | [optional] 
**LastSfdcSyncStatus** | Pointer to **map[string]interface{}** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**Organizationsettings**](Organizationsettings.md) |  | [optional] 
**TotalBillingEstimate** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountsReceivable

`func (o *Organization) GetAccountsReceivable() string`

GetAccountsReceivable returns the AccountsReceivable field if non-nil, zero value otherwise.

### GetAccountsReceivableOk

`func (o *Organization) GetAccountsReceivableOk() (*string, bool)`

GetAccountsReceivableOk returns a tuple with the AccountsReceivable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsReceivable

`func (o *Organization) SetAccountsReceivable(v string)`

SetAccountsReceivable sets AccountsReceivable field to given value.

### HasAccountsReceivable

`func (o *Organization) HasAccountsReceivable() bool`

HasAccountsReceivable returns a boolean if a field has been set.

### GetCreated

`func (o *Organization) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Organization) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Organization) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Organization) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDisplayName

`func (o *Organization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Organization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Organization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Organization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEntitlement

`func (o *Organization) GetEntitlement() Organizationentitlement`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *Organization) GetEntitlementOk() (*Organizationentitlement, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *Organization) SetEntitlement(v Organizationentitlement)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *Organization) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.

### GetHasCreditCard

`func (o *Organization) GetHasCreditCard() bool`

GetHasCreditCard returns the HasCreditCard field if non-nil, zero value otherwise.

### GetHasCreditCardOk

`func (o *Organization) GetHasCreditCardOk() (*bool, bool)`

GetHasCreditCardOk returns a tuple with the HasCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCreditCard

`func (o *Organization) SetHasCreditCard(v bool)`

SetHasCreditCard sets HasCreditCard field to given value.

### HasHasCreditCard

`func (o *Organization) HasHasCreditCard() bool`

HasHasCreditCard returns a boolean if a field has been set.

### GetHasStripeCustomerId

`func (o *Organization) GetHasStripeCustomerId() bool`

GetHasStripeCustomerId returns the HasStripeCustomerId field if non-nil, zero value otherwise.

### GetHasStripeCustomerIdOk

`func (o *Organization) GetHasStripeCustomerIdOk() (*bool, bool)`

GetHasStripeCustomerIdOk returns a tuple with the HasStripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStripeCustomerId

`func (o *Organization) SetHasStripeCustomerId(v bool)`

SetHasStripeCustomerId sets HasStripeCustomerId field to given value.

### HasHasStripeCustomerId

`func (o *Organization) HasHasStripeCustomerId() bool`

HasHasStripeCustomerId returns a boolean if a field has been set.

### GetLastEstimateCalculationTimeStamp

`func (o *Organization) GetLastEstimateCalculationTimeStamp() string`

GetLastEstimateCalculationTimeStamp returns the LastEstimateCalculationTimeStamp field if non-nil, zero value otherwise.

### GetLastEstimateCalculationTimeStampOk

`func (o *Organization) GetLastEstimateCalculationTimeStampOk() (*string, bool)`

GetLastEstimateCalculationTimeStampOk returns a tuple with the LastEstimateCalculationTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEstimateCalculationTimeStamp

`func (o *Organization) SetLastEstimateCalculationTimeStamp(v string)`

SetLastEstimateCalculationTimeStamp sets LastEstimateCalculationTimeStamp field to given value.

### HasLastEstimateCalculationTimeStamp

`func (o *Organization) HasLastEstimateCalculationTimeStamp() bool`

HasLastEstimateCalculationTimeStamp returns a boolean if a field has been set.

### GetLastSfdcSyncStatus

`func (o *Organization) GetLastSfdcSyncStatus() map[string]interface{}`

GetLastSfdcSyncStatus returns the LastSfdcSyncStatus field if non-nil, zero value otherwise.

### GetLastSfdcSyncStatusOk

`func (o *Organization) GetLastSfdcSyncStatusOk() (*map[string]interface{}, bool)`

GetLastSfdcSyncStatusOk returns a tuple with the LastSfdcSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSfdcSyncStatus

`func (o *Organization) SetLastSfdcSyncStatus(v map[string]interface{})`

SetLastSfdcSyncStatus sets LastSfdcSyncStatus field to given value.

### HasLastSfdcSyncStatus

`func (o *Organization) HasLastSfdcSyncStatus() bool`

HasLastSfdcSyncStatus returns a boolean if a field has been set.

### GetLogoUrl

`func (o *Organization) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Organization) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Organization) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Organization) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetProvider

`func (o *Organization) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Organization) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Organization) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Organization) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSettings

`func (o *Organization) GetSettings() Organizationsettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Organization) GetSettingsOk() (*Organizationsettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Organization) SetSettings(v Organizationsettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Organization) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTotalBillingEstimate

`func (o *Organization) GetTotalBillingEstimate() int32`

GetTotalBillingEstimate returns the TotalBillingEstimate field if non-nil, zero value otherwise.

### GetTotalBillingEstimateOk

`func (o *Organization) GetTotalBillingEstimateOk() (*int32, bool)`

GetTotalBillingEstimateOk returns a tuple with the TotalBillingEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBillingEstimate

`func (o *Organization) SetTotalBillingEstimate(v int32)`

SetTotalBillingEstimate sets TotalBillingEstimate field to given value.

### HasTotalBillingEstimate

`func (o *Organization) HasTotalBillingEstimate() bool`

HasTotalBillingEstimate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



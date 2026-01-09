# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnualPrice** | **float32** | The annual (discounted) price of this subscription. | 
**DisplayName** | **string** | The display name of this subscription. | 
**Features** | [**[]Feature**](Feature.md) | Array of the features included in the subscription. | 
**ListPrice** | **float32** | The list price of this subscription. | 
**ProductCode** | **string** | Unique identifier corresponding to this subscription. | 

## Methods

### NewSubscription

`func NewSubscription(annualPrice float32, displayName string, features []Feature, listPrice float32, productCode string, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnualPrice

`func (o *Subscription) GetAnnualPrice() float32`

GetAnnualPrice returns the AnnualPrice field if non-nil, zero value otherwise.

### GetAnnualPriceOk

`func (o *Subscription) GetAnnualPriceOk() (*float32, bool)`

GetAnnualPriceOk returns a tuple with the AnnualPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualPrice

`func (o *Subscription) SetAnnualPrice(v float32)`

SetAnnualPrice sets AnnualPrice field to given value.


### GetDisplayName

`func (o *Subscription) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Subscription) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Subscription) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFeatures

`func (o *Subscription) GetFeatures() []Feature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Subscription) GetFeaturesOk() (*[]Feature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Subscription) SetFeatures(v []Feature)`

SetFeatures sets Features field to given value.


### GetListPrice

`func (o *Subscription) GetListPrice() float32`

GetListPrice returns the ListPrice field if non-nil, zero value otherwise.

### GetListPriceOk

`func (o *Subscription) GetListPriceOk() (*float32, bool)`

GetListPriceOk returns a tuple with the ListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPrice

`func (o *Subscription) SetListPrice(v float32)`

SetListPrice sets ListPrice field to given value.


### GetProductCode

`func (o *Subscription) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *Subscription) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *Subscription) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



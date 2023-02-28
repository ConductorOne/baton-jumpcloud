# AuthnPolicyUserAttributeTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclusions** | Pointer to [**[]AuthnPolicyUserAttributeFilter**](AuthnPolicyUserAttributeFilter.md) |  | [optional] 
**Inclusions** | Pointer to [**[]AuthnPolicyUserAttributeFilter**](AuthnPolicyUserAttributeFilter.md) |  | [optional] 

## Methods

### NewAuthnPolicyUserAttributeTarget

`func NewAuthnPolicyUserAttributeTarget() *AuthnPolicyUserAttributeTarget`

NewAuthnPolicyUserAttributeTarget instantiates a new AuthnPolicyUserAttributeTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnPolicyUserAttributeTargetWithDefaults

`func NewAuthnPolicyUserAttributeTargetWithDefaults() *AuthnPolicyUserAttributeTarget`

NewAuthnPolicyUserAttributeTargetWithDefaults instantiates a new AuthnPolicyUserAttributeTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclusions

`func (o *AuthnPolicyUserAttributeTarget) GetExclusions() []AuthnPolicyUserAttributeFilter`

GetExclusions returns the Exclusions field if non-nil, zero value otherwise.

### GetExclusionsOk

`func (o *AuthnPolicyUserAttributeTarget) GetExclusionsOk() (*[]AuthnPolicyUserAttributeFilter, bool)`

GetExclusionsOk returns a tuple with the Exclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusions

`func (o *AuthnPolicyUserAttributeTarget) SetExclusions(v []AuthnPolicyUserAttributeFilter)`

SetExclusions sets Exclusions field to given value.

### HasExclusions

`func (o *AuthnPolicyUserAttributeTarget) HasExclusions() bool`

HasExclusions returns a boolean if a field has been set.

### GetInclusions

`func (o *AuthnPolicyUserAttributeTarget) GetInclusions() []AuthnPolicyUserAttributeFilter`

GetInclusions returns the Inclusions field if non-nil, zero value otherwise.

### GetInclusionsOk

`func (o *AuthnPolicyUserAttributeTarget) GetInclusionsOk() (*[]AuthnPolicyUserAttributeFilter, bool)`

GetInclusionsOk returns a tuple with the Inclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusions

`func (o *AuthnPolicyUserAttributeTarget) SetInclusions(v []AuthnPolicyUserAttributeFilter)`

SetInclusions sets Inclusions field to given value.

### HasInclusions

`func (o *AuthnPolicyUserAttributeTarget) HasInclusions() bool`

HasInclusions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



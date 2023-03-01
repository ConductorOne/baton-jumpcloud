# AuthnPolicyObligations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mfa** | Pointer to [**AuthnPolicyObligationsMfa**](AuthnPolicyObligationsMfa.md) |  | [optional] 
**UserVerification** | Pointer to [**AuthnPolicyObligationsUserVerification**](AuthnPolicyObligationsUserVerification.md) |  | [optional] 

## Methods

### NewAuthnPolicyObligations

`func NewAuthnPolicyObligations() *AuthnPolicyObligations`

NewAuthnPolicyObligations instantiates a new AuthnPolicyObligations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnPolicyObligationsWithDefaults

`func NewAuthnPolicyObligationsWithDefaults() *AuthnPolicyObligations`

NewAuthnPolicyObligationsWithDefaults instantiates a new AuthnPolicyObligations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfa

`func (o *AuthnPolicyObligations) GetMfa() AuthnPolicyObligationsMfa`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *AuthnPolicyObligations) GetMfaOk() (*AuthnPolicyObligationsMfa, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *AuthnPolicyObligations) SetMfa(v AuthnPolicyObligationsMfa)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *AuthnPolicyObligations) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetUserVerification

`func (o *AuthnPolicyObligations) GetUserVerification() AuthnPolicyObligationsUserVerification`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthnPolicyObligations) GetUserVerificationOk() (*AuthnPolicyObligationsUserVerification, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthnPolicyObligations) SetUserVerification(v AuthnPolicyObligationsUserVerification)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthnPolicyObligations) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



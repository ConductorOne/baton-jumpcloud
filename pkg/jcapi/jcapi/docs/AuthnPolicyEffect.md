# AuthnPolicyEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Obligations** | Pointer to [**AuthnPolicyObligations**](AuthnPolicyObligations.md) |  | [optional] 

## Methods

### NewAuthnPolicyEffect

`func NewAuthnPolicyEffect(action string, ) *AuthnPolicyEffect`

NewAuthnPolicyEffect instantiates a new AuthnPolicyEffect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnPolicyEffectWithDefaults

`func NewAuthnPolicyEffectWithDefaults() *AuthnPolicyEffect`

NewAuthnPolicyEffectWithDefaults instantiates a new AuthnPolicyEffect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuthnPolicyEffect) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuthnPolicyEffect) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuthnPolicyEffect) SetAction(v string)`

SetAction sets Action field to given value.


### GetObligations

`func (o *AuthnPolicyEffect) GetObligations() AuthnPolicyObligations`

GetObligations returns the Obligations field if non-nil, zero value otherwise.

### GetObligationsOk

`func (o *AuthnPolicyEffect) GetObligationsOk() (*AuthnPolicyObligations, bool)`

GetObligationsOk returns a tuple with the Obligations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObligations

`func (o *AuthnPolicyEffect) SetObligations(v AuthnPolicyObligations)`

SetObligations sets Obligations field to given value.

### HasObligations

`func (o *AuthnPolicyEffect) HasObligations() bool`

HasObligations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



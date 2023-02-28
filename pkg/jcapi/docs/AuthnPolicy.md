# AuthnPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to **map[string]interface{}** | Conditions may be added to an authentication policy using the following conditional language:  &#x60;&#x60;&#x60; &lt;conditions&gt; ::&#x3D; &lt;expression&gt; &lt;expression&gt; ::&#x3D; &lt;deviceEncrypted&gt; | &lt;deviceManaged&gt; | &lt;ipAddressIn&gt; |                  &lt;locationIn&gt; | &lt;notExpression&gt; | &lt;allExpression&gt; |                  &lt;anyExpression&gt; &lt;deviceEncrypted&gt; ::&#x3D; { \&quot;deviceEncrypted\&quot;: &lt;boolean&gt; } &lt;deviceManaged&gt; ::&#x3D; { \&quot;deviceManaged\&quot;: &lt;boolean&gt; } &lt;ipAddressIn&gt; ::&#x3D; { \&quot;ipAddressIn\&quot;: [ &lt;objectId&gt;, ... ] } &lt;locationIn&gt; ::&#x3D; { \&quot;locationIn\&quot;: {                      \&quot;countries\&quot;: [                        &lt;iso_3166_country_code&gt;, ...                      ]                    }                  } &lt;notExpression&gt; ::&#x3D; { \&quot;not\&quot;: &lt;expression&gt; } &lt;allExpression&gt; ::&#x3D; { \&quot;all\&quot;: [ &lt;expression&gt;, ... ] } &lt;anyExpression&gt; ::&#x3D; { \&quot;any\&quot;: [ &lt;expression&gt;, ... ] } &#x60;&#x60;&#x60;  For example, to add a condition that applies to IP addresses in a given list, the following condition can be added:  &#x60;&#x60;&#x60; {\&quot;ipAddressIn\&quot;: [ &lt;ip_list_object_id&gt; ]} &#x60;&#x60;&#x60;  If you would rather exclude IP addresses in the given lists, the following condition could be added:  &#x60;&#x60;&#x60; {   \&quot;not\&quot;: {     \&quot;ipAddressIn\&quot;: [ &lt;ip_list_object_id_1&gt;, &lt;ip_list_object_id_2&gt; ]   } } &#x60;&#x60;&#x60;  You may also include more than one condition and choose whether \&quot;all\&quot; or \&quot;any\&quot; of them must be met for the policy to apply:  &#x60;&#x60;&#x60; {   \&quot;all\&quot;: [     {       \&quot;ipAddressIn\&quot;: [ &lt;ip_list_object_id&gt;, ... ]     },     {       \&quot;deviceManaged\&quot;: true     },     {       \&quot;locationIn\&quot;: {         countries: [ &lt;iso_3166_country_code&gt;, ... ]       }     }   ] } &#x60;&#x60;&#x60; | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Effect** | Pointer to [**AuthnPolicyEffect**](AuthnPolicyEffect.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Targets** | Pointer to [**AuthnPolicyTargets**](AuthnPolicyTargets.md) |  | [optional] 
**Type** | Pointer to [**AuthnPolicyType**](AuthnPolicyType.md) |  | [optional] [default to AUTHNPOLICYTYPE_USER_PORTAL]

## Methods

### NewAuthnPolicy

`func NewAuthnPolicy() *AuthnPolicy`

NewAuthnPolicy instantiates a new AuthnPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnPolicyWithDefaults

`func NewAuthnPolicyWithDefaults() *AuthnPolicy`

NewAuthnPolicyWithDefaults instantiates a new AuthnPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *AuthnPolicy) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthnPolicy) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthnPolicy) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AuthnPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDescription

`func (o *AuthnPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthnPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthnPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthnPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *AuthnPolicy) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AuthnPolicy) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AuthnPolicy) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AuthnPolicy) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEffect

`func (o *AuthnPolicy) GetEffect() AuthnPolicyEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *AuthnPolicy) GetEffectOk() (*AuthnPolicyEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *AuthnPolicy) SetEffect(v AuthnPolicyEffect)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *AuthnPolicy) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetId

`func (o *AuthnPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthnPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthnPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthnPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthnPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthnPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthnPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthnPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTargets

`func (o *AuthnPolicy) GetTargets() AuthnPolicyTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AuthnPolicy) GetTargetsOk() (*AuthnPolicyTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AuthnPolicy) SetTargets(v AuthnPolicyTargets)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *AuthnPolicy) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetType

`func (o *AuthnPolicy) GetType() AuthnPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthnPolicy) GetTypeOk() (*AuthnPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthnPolicy) SetType(v AuthnPolicyType)`

SetType sets Type field to given value.

### HasType

`func (o *AuthnPolicy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



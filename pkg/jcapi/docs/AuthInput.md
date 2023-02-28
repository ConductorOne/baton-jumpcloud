# AuthInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basic** | Pointer to [**AuthInputBasic**](AuthInputBasic.md) |  | [optional] 
**Oauth** | Pointer to [**AuthInputOauth**](AuthInputOauth.md) |  | [optional] 

## Methods

### NewAuthInput

`func NewAuthInput() *AuthInput`

NewAuthInput instantiates a new AuthInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthInputWithDefaults

`func NewAuthInputWithDefaults() *AuthInput`

NewAuthInputWithDefaults instantiates a new AuthInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasic

`func (o *AuthInput) GetBasic() AuthInputBasic`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *AuthInput) GetBasicOk() (*AuthInputBasic, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *AuthInput) SetBasic(v AuthInputBasic)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *AuthInput) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetOauth

`func (o *AuthInput) GetOauth() AuthInputOauth`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *AuthInput) GetOauthOk() (*AuthInputOauth, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *AuthInput) SetOauth(v AuthInputOauth)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *AuthInput) HasOauth() bool`

HasOauth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



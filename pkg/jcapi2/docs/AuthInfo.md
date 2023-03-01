# AuthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **string** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthInfo

`func NewAuthInfo() *AuthInfo`

NewAuthInfo instantiates a new AuthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthInfoWithDefaults

`func NewAuthInfoWithDefaults() *AuthInfo`

NewAuthInfoWithDefaults instantiates a new AuthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *AuthInfo) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AuthInfo) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AuthInfo) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AuthInfo) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetIsValid

`func (o *AuthInfo) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *AuthInfo) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *AuthInfo) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *AuthInfo) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetMessage

`func (o *AuthInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



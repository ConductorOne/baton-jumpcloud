# LdapserversPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserLockoutAction** | Pointer to [**LdapServerAction**](LdapServerAction.md) |  | [optional] 
**UserPasswordExpirationAction** | Pointer to [**LdapServerAction**](LdapServerAction.md) |  | [optional] 

## Methods

### NewLdapserversPatchRequest

`func NewLdapserversPatchRequest() *LdapserversPatchRequest`

NewLdapserversPatchRequest instantiates a new LdapserversPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapserversPatchRequestWithDefaults

`func NewLdapserversPatchRequestWithDefaults() *LdapserversPatchRequest`

NewLdapserversPatchRequestWithDefaults instantiates a new LdapserversPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapserversPatchRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapserversPatchRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapserversPatchRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LdapserversPatchRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserLockoutAction

`func (o *LdapserversPatchRequest) GetUserLockoutAction() LdapServerAction`

GetUserLockoutAction returns the UserLockoutAction field if non-nil, zero value otherwise.

### GetUserLockoutActionOk

`func (o *LdapserversPatchRequest) GetUserLockoutActionOk() (*LdapServerAction, bool)`

GetUserLockoutActionOk returns a tuple with the UserLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutAction

`func (o *LdapserversPatchRequest) SetUserLockoutAction(v LdapServerAction)`

SetUserLockoutAction sets UserLockoutAction field to given value.

### HasUserLockoutAction

`func (o *LdapserversPatchRequest) HasUserLockoutAction() bool`

HasUserLockoutAction returns a boolean if a field has been set.

### GetUserPasswordExpirationAction

`func (o *LdapserversPatchRequest) GetUserPasswordExpirationAction() LdapServerAction`

GetUserPasswordExpirationAction returns the UserPasswordExpirationAction field if non-nil, zero value otherwise.

### GetUserPasswordExpirationActionOk

`func (o *LdapserversPatchRequest) GetUserPasswordExpirationActionOk() (*LdapServerAction, bool)`

GetUserPasswordExpirationActionOk returns a tuple with the UserPasswordExpirationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordExpirationAction

`func (o *LdapserversPatchRequest) SetUserPasswordExpirationAction(v LdapServerAction)`

SetUserPasswordExpirationAction sets UserPasswordExpirationAction field to given value.

### HasUserPasswordExpirationAction

`func (o *LdapserversPatchRequest) HasUserPasswordExpirationAction() bool`

HasUserPasswordExpirationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



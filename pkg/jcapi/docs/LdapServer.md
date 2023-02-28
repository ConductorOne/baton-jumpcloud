# LdapServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of this LDAP server | [optional] [readonly] 
**Name** | Pointer to **string** | The name of this LDAP server | [optional] 
**UserLockoutAction** | Pointer to **string** | action to take; one of &#39;remove&#39; or &#39;disable&#39; | [optional] 
**UserPasswordExpirationAction** | Pointer to **string** | action to take; one of &#39;remove&#39; or &#39;disable&#39; | [optional] 

## Methods

### NewLdapServer

`func NewLdapServer() *LdapServer`

NewLdapServer instantiates a new LdapServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapServerWithDefaults

`func NewLdapServerWithDefaults() *LdapServer`

NewLdapServerWithDefaults instantiates a new LdapServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LdapServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LdapServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserLockoutAction

`func (o *LdapServer) GetUserLockoutAction() string`

GetUserLockoutAction returns the UserLockoutAction field if non-nil, zero value otherwise.

### GetUserLockoutActionOk

`func (o *LdapServer) GetUserLockoutActionOk() (*string, bool)`

GetUserLockoutActionOk returns a tuple with the UserLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutAction

`func (o *LdapServer) SetUserLockoutAction(v string)`

SetUserLockoutAction sets UserLockoutAction field to given value.

### HasUserLockoutAction

`func (o *LdapServer) HasUserLockoutAction() bool`

HasUserLockoutAction returns a boolean if a field has been set.

### GetUserPasswordExpirationAction

`func (o *LdapServer) GetUserPasswordExpirationAction() string`

GetUserPasswordExpirationAction returns the UserPasswordExpirationAction field if non-nil, zero value otherwise.

### GetUserPasswordExpirationActionOk

`func (o *LdapServer) GetUserPasswordExpirationActionOk() (*string, bool)`

GetUserPasswordExpirationActionOk returns a tuple with the UserPasswordExpirationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordExpirationAction

`func (o *LdapServer) SetUserPasswordExpirationAction(v string)`

SetUserPasswordExpirationAction sets UserPasswordExpirationAction field to given value.

### HasUserPasswordExpirationAction

`func (o *LdapServer) HasUserPasswordExpirationAction() bool`

HasUserPasswordExpirationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



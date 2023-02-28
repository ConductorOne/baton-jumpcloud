# LdapserversPatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UserLockoutAction** | Pointer to [**LdapServerAction**](LdapServerAction.md) |  | [optional] 
**UserPasswordExpirationAction** | Pointer to [**LdapServerAction**](LdapServerAction.md) |  | [optional] 

## Methods

### NewLdapserversPatch200Response

`func NewLdapserversPatch200Response() *LdapserversPatch200Response`

NewLdapserversPatch200Response instantiates a new LdapserversPatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapserversPatch200ResponseWithDefaults

`func NewLdapserversPatch200ResponseWithDefaults() *LdapserversPatch200Response`

NewLdapserversPatch200ResponseWithDefaults instantiates a new LdapserversPatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapserversPatch200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapserversPatch200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapserversPatch200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LdapserversPatch200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LdapserversPatch200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapserversPatch200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapserversPatch200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapserversPatch200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserLockoutAction

`func (o *LdapserversPatch200Response) GetUserLockoutAction() LdapServerAction`

GetUserLockoutAction returns the UserLockoutAction field if non-nil, zero value otherwise.

### GetUserLockoutActionOk

`func (o *LdapserversPatch200Response) GetUserLockoutActionOk() (*LdapServerAction, bool)`

GetUserLockoutActionOk returns a tuple with the UserLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutAction

`func (o *LdapserversPatch200Response) SetUserLockoutAction(v LdapServerAction)`

SetUserLockoutAction sets UserLockoutAction field to given value.

### HasUserLockoutAction

`func (o *LdapserversPatch200Response) HasUserLockoutAction() bool`

HasUserLockoutAction returns a boolean if a field has been set.

### GetUserPasswordExpirationAction

`func (o *LdapserversPatch200Response) GetUserPasswordExpirationAction() LdapServerAction`

GetUserPasswordExpirationAction returns the UserPasswordExpirationAction field if non-nil, zero value otherwise.

### GetUserPasswordExpirationActionOk

`func (o *LdapserversPatch200Response) GetUserPasswordExpirationActionOk() (*LdapServerAction, bool)`

GetUserPasswordExpirationActionOk returns a tuple with the UserPasswordExpirationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordExpirationAction

`func (o *LdapserversPatch200Response) SetUserPasswordExpirationAction(v LdapServerAction)`

SetUserPasswordExpirationAction sets UserPasswordExpirationAction field to given value.

### HasUserPasswordExpirationAction

`func (o *LdapserversPatch200Response) HasUserPasswordExpirationAction() bool`

HasUserPasswordExpirationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



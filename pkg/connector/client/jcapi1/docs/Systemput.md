# Systemput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentBoundMessages** | Pointer to [**[]SystemputAgentBoundMessagesInner**](SystemputAgentBoundMessagesInner.md) |  | [optional] 
**AllowMultiFactorAuthentication** | Pointer to **bool** |  | [optional] 
**AllowPublicKeyAuthentication** | Pointer to **bool** |  | [optional] 
**AllowSshPasswordAuthentication** | Pointer to **bool** |  | [optional] 
**AllowSshRootLogin** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSystemput

`func NewSystemput() *Systemput`

NewSystemput instantiates a new Systemput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemputWithDefaults

`func NewSystemputWithDefaults() *Systemput`

NewSystemputWithDefaults instantiates a new Systemput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBoundMessages

`func (o *Systemput) GetAgentBoundMessages() []SystemputAgentBoundMessagesInner`

GetAgentBoundMessages returns the AgentBoundMessages field if non-nil, zero value otherwise.

### GetAgentBoundMessagesOk

`func (o *Systemput) GetAgentBoundMessagesOk() (*[]SystemputAgentBoundMessagesInner, bool)`

GetAgentBoundMessagesOk returns a tuple with the AgentBoundMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBoundMessages

`func (o *Systemput) SetAgentBoundMessages(v []SystemputAgentBoundMessagesInner)`

SetAgentBoundMessages sets AgentBoundMessages field to given value.

### HasAgentBoundMessages

`func (o *Systemput) HasAgentBoundMessages() bool`

HasAgentBoundMessages returns a boolean if a field has been set.

### GetAllowMultiFactorAuthentication

`func (o *Systemput) GetAllowMultiFactorAuthentication() bool`

GetAllowMultiFactorAuthentication returns the AllowMultiFactorAuthentication field if non-nil, zero value otherwise.

### GetAllowMultiFactorAuthenticationOk

`func (o *Systemput) GetAllowMultiFactorAuthenticationOk() (*bool, bool)`

GetAllowMultiFactorAuthenticationOk returns a tuple with the AllowMultiFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiFactorAuthentication

`func (o *Systemput) SetAllowMultiFactorAuthentication(v bool)`

SetAllowMultiFactorAuthentication sets AllowMultiFactorAuthentication field to given value.

### HasAllowMultiFactorAuthentication

`func (o *Systemput) HasAllowMultiFactorAuthentication() bool`

HasAllowMultiFactorAuthentication returns a boolean if a field has been set.

### GetAllowPublicKeyAuthentication

`func (o *Systemput) GetAllowPublicKeyAuthentication() bool`

GetAllowPublicKeyAuthentication returns the AllowPublicKeyAuthentication field if non-nil, zero value otherwise.

### GetAllowPublicKeyAuthenticationOk

`func (o *Systemput) GetAllowPublicKeyAuthenticationOk() (*bool, bool)`

GetAllowPublicKeyAuthenticationOk returns a tuple with the AllowPublicKeyAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPublicKeyAuthentication

`func (o *Systemput) SetAllowPublicKeyAuthentication(v bool)`

SetAllowPublicKeyAuthentication sets AllowPublicKeyAuthentication field to given value.

### HasAllowPublicKeyAuthentication

`func (o *Systemput) HasAllowPublicKeyAuthentication() bool`

HasAllowPublicKeyAuthentication returns a boolean if a field has been set.

### GetAllowSshPasswordAuthentication

`func (o *Systemput) GetAllowSshPasswordAuthentication() bool`

GetAllowSshPasswordAuthentication returns the AllowSshPasswordAuthentication field if non-nil, zero value otherwise.

### GetAllowSshPasswordAuthenticationOk

`func (o *Systemput) GetAllowSshPasswordAuthenticationOk() (*bool, bool)`

GetAllowSshPasswordAuthenticationOk returns a tuple with the AllowSshPasswordAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSshPasswordAuthentication

`func (o *Systemput) SetAllowSshPasswordAuthentication(v bool)`

SetAllowSshPasswordAuthentication sets AllowSshPasswordAuthentication field to given value.

### HasAllowSshPasswordAuthentication

`func (o *Systemput) HasAllowSshPasswordAuthentication() bool`

HasAllowSshPasswordAuthentication returns a boolean if a field has been set.

### GetAllowSshRootLogin

`func (o *Systemput) GetAllowSshRootLogin() bool`

GetAllowSshRootLogin returns the AllowSshRootLogin field if non-nil, zero value otherwise.

### GetAllowSshRootLoginOk

`func (o *Systemput) GetAllowSshRootLoginOk() (*bool, bool)`

GetAllowSshRootLoginOk returns a tuple with the AllowSshRootLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSshRootLogin

`func (o *Systemput) SetAllowSshRootLogin(v bool)`

SetAllowSshRootLogin sets AllowSshRootLogin field to given value.

### HasAllowSshRootLogin

`func (o *Systemput) HasAllowSshRootLogin() bool`

HasAllowSshRootLogin returns a boolean if a field has been set.

### GetDisplayName

`func (o *Systemput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Systemput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Systemput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Systemput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTags

`func (o *Systemput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Systemput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Systemput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Systemput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



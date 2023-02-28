# ActiveDirectoryAgentGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectKey** | Pointer to **string** | The connect key to use when installing the Agent on a Domain Controller. | [optional] 
**ContactAt** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Id** | **string** | ObjectID of this Active Directory Agent. | 
**SourceIp** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewActiveDirectoryAgentGet

`func NewActiveDirectoryAgentGet(id string, ) *ActiveDirectoryAgentGet`

NewActiveDirectoryAgentGet instantiates a new ActiveDirectoryAgentGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryAgentGetWithDefaults

`func NewActiveDirectoryAgentGetWithDefaults() *ActiveDirectoryAgentGet`

NewActiveDirectoryAgentGetWithDefaults instantiates a new ActiveDirectoryAgentGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectKey

`func (o *ActiveDirectoryAgentGet) GetConnectKey() string`

GetConnectKey returns the ConnectKey field if non-nil, zero value otherwise.

### GetConnectKeyOk

`func (o *ActiveDirectoryAgentGet) GetConnectKeyOk() (*string, bool)`

GetConnectKeyOk returns a tuple with the ConnectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectKey

`func (o *ActiveDirectoryAgentGet) SetConnectKey(v string)`

SetConnectKey sets ConnectKey field to given value.

### HasConnectKey

`func (o *ActiveDirectoryAgentGet) HasConnectKey() bool`

HasConnectKey returns a boolean if a field has been set.

### GetContactAt

`func (o *ActiveDirectoryAgentGet) GetContactAt() string`

GetContactAt returns the ContactAt field if non-nil, zero value otherwise.

### GetContactAtOk

`func (o *ActiveDirectoryAgentGet) GetContactAtOk() (*string, bool)`

GetContactAtOk returns a tuple with the ContactAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAt

`func (o *ActiveDirectoryAgentGet) SetContactAt(v string)`

SetContactAt sets ContactAt field to given value.

### HasContactAt

`func (o *ActiveDirectoryAgentGet) HasContactAt() bool`

HasContactAt returns a boolean if a field has been set.

### GetHostname

`func (o *ActiveDirectoryAgentGet) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ActiveDirectoryAgentGet) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ActiveDirectoryAgentGet) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ActiveDirectoryAgentGet) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *ActiveDirectoryAgentGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveDirectoryAgentGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveDirectoryAgentGet) SetId(v string)`

SetId sets Id field to given value.


### GetSourceIp

`func (o *ActiveDirectoryAgentGet) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *ActiveDirectoryAgentGet) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *ActiveDirectoryAgentGet) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *ActiveDirectoryAgentGet) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetState

`func (o *ActiveDirectoryAgentGet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActiveDirectoryAgentGet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActiveDirectoryAgentGet) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ActiveDirectoryAgentGet) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *ActiveDirectoryAgentGet) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ActiveDirectoryAgentGet) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ActiveDirectoryAgentGet) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ActiveDirectoryAgentGet) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



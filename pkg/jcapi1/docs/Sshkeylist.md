# Sshkeylist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the SSH key. | [optional] 
**CreateDate** | Pointer to **string** | The date the SSH key was created. | [optional] 
**Name** | Pointer to **string** | The name of the SSH key. | [optional] 
**PublicKey** | Pointer to **string** | The Public SSH key. | [optional] 

## Methods

### NewSshkeylist

`func NewSshkeylist() *Sshkeylist`

NewSshkeylist instantiates a new Sshkeylist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshkeylistWithDefaults

`func NewSshkeylistWithDefaults() *Sshkeylist`

NewSshkeylistWithDefaults instantiates a new Sshkeylist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Sshkeylist) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sshkeylist) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sshkeylist) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Sshkeylist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateDate

`func (o *Sshkeylist) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *Sshkeylist) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *Sshkeylist) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *Sshkeylist) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetName

`func (o *Sshkeylist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sshkeylist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sshkeylist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sshkeylist) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *Sshkeylist) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Sshkeylist) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Sshkeylist) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *Sshkeylist) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Directory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ObjectID of the directory. | 
**Name** | **string** | The name of the directory. | 
**OAuthStatus** | Pointer to **map[string]interface{}** | the expiry and error status of the bearer token | [optional] 
**Type** | **string** | The type of directory. | 

## Methods

### NewDirectory

`func NewDirectory(id string, name string, type_ string, ) *Directory`

NewDirectory instantiates a new Directory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryWithDefaults

`func NewDirectoryWithDefaults() *Directory`

NewDirectoryWithDefaults instantiates a new Directory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Directory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Directory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Directory) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Directory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Directory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Directory) SetName(v string)`

SetName sets Name field to given value.


### GetOAuthStatus

`func (o *Directory) GetOAuthStatus() map[string]interface{}`

GetOAuthStatus returns the OAuthStatus field if non-nil, zero value otherwise.

### GetOAuthStatusOk

`func (o *Directory) GetOAuthStatusOk() (*map[string]interface{}, bool)`

GetOAuthStatusOk returns a tuple with the OAuthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthStatus

`func (o *Directory) SetOAuthStatus(v map[string]interface{})`

SetOAuthStatus sets OAuthStatus field to given value.

### HasOAuthStatus

`func (o *Directory) HasOAuthStatus() bool`

HasOAuthStatus returns a boolean if a field has been set.

### GetType

`func (o *Directory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Directory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Directory) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



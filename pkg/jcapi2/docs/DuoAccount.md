# DuoAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | object ID | 
**Name** | Pointer to **string** | Duo application name. | [optional] 

## Methods

### NewDuoAccount

`func NewDuoAccount(id string, ) *DuoAccount`

NewDuoAccount instantiates a new DuoAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuoAccountWithDefaults

`func NewDuoAccountWithDefaults() *DuoAccount`

NewDuoAccountWithDefaults instantiates a new DuoAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DuoAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DuoAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DuoAccount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DuoAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DuoAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DuoAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DuoAccount) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SambaDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of this domain | [optional] [readonly] 
**Name** | **string** | Name of this domain&#39;s WorkGroup | 
**Sid** | **string** | Security identifier of this domain | 

## Methods

### NewSambaDomain

`func NewSambaDomain(name string, sid string, ) *SambaDomain`

NewSambaDomain instantiates a new SambaDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSambaDomainWithDefaults

`func NewSambaDomainWithDefaults() *SambaDomain`

NewSambaDomainWithDefaults instantiates a new SambaDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SambaDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SambaDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SambaDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SambaDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SambaDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SambaDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SambaDomain) SetName(v string)`

SetName sets Name field to given value.


### GetSid

`func (o *SambaDomain) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SambaDomain) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SambaDomain) SetSid(v string)`

SetSid sets Sid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



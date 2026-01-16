# ActiveDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain name for this Active Directory instance. | [optional] 
**Id** | Pointer to **string** | ObjectID of this Active Directory instance. | [optional] [readonly] 

## Methods

### NewActiveDirectory

`func NewActiveDirectory() *ActiveDirectory`

NewActiveDirectory instantiates a new ActiveDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryWithDefaults

`func NewActiveDirectoryWithDefaults() *ActiveDirectory`

NewActiveDirectoryWithDefaults instantiates a new ActiveDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ActiveDirectory) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ActiveDirectory) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ActiveDirectory) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ActiveDirectory) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetId

`func (o *ActiveDirectory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveDirectory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveDirectory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActiveDirectory) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



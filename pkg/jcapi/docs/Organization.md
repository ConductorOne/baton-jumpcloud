# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MaxSystemUsers** | Pointer to **int32** | The maximum number of users allowed in this organization. Requires organizations.billing scope to modify. | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxSystemUsers

`func (o *Organization) GetMaxSystemUsers() int32`

GetMaxSystemUsers returns the MaxSystemUsers field if non-nil, zero value otherwise.

### GetMaxSystemUsersOk

`func (o *Organization) GetMaxSystemUsersOk() (*int32, bool)`

GetMaxSystemUsersOk returns a tuple with the MaxSystemUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSystemUsers

`func (o *Organization) SetMaxSystemUsers(v int32)`

SetMaxSystemUsers sets MaxSystemUsers field to given value.

### HasMaxSystemUsers

`func (o *Organization) HasMaxSystemUsers() bool`

HasMaxSystemUsers returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSystemUsers** | Pointer to **int32** | The maximum number of users allowed in this organization. Requires organizations.billing scope to modify. | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOrganization

`func NewCreateOrganization() *CreateOrganization`

NewCreateOrganization instantiates a new CreateOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationWithDefaults

`func NewCreateOrganizationWithDefaults() *CreateOrganization`

NewCreateOrganizationWithDefaults instantiates a new CreateOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSystemUsers

`func (o *CreateOrganization) GetMaxSystemUsers() int32`

GetMaxSystemUsers returns the MaxSystemUsers field if non-nil, zero value otherwise.

### GetMaxSystemUsersOk

`func (o *CreateOrganization) GetMaxSystemUsersOk() (*int32, bool)`

GetMaxSystemUsersOk returns a tuple with the MaxSystemUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSystemUsers

`func (o *CreateOrganization) SetMaxSystemUsers(v int32)`

SetMaxSystemUsers sets MaxSystemUsers field to given value.

### HasMaxSystemUsers

`func (o *CreateOrganization) HasMaxSystemUsers() bool`

HasMaxSystemUsers returns a boolean if a field has been set.

### GetName

`func (o *CreateOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrganization) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



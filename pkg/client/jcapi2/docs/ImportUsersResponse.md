# ImportUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **float32** |  | [optional] 
**Users** | Pointer to [**[]ImportUser**](ImportUser.md) |  | [optional] 

## Methods

### NewImportUsersResponse

`func NewImportUsersResponse() *ImportUsersResponse`

NewImportUsersResponse instantiates a new ImportUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportUsersResponseWithDefaults

`func NewImportUsersResponseWithDefaults() *ImportUsersResponse`

NewImportUsersResponseWithDefaults instantiates a new ImportUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ImportUsersResponse) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ImportUsersResponse) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ImportUsersResponse) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ImportUsersResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetUsers

`func (o *ImportUsersResponse) GetUsers() []ImportUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ImportUsersResponse) GetUsersOk() (*[]ImportUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ImportUsersResponse) SetUsers(v []ImportUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ImportUsersResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



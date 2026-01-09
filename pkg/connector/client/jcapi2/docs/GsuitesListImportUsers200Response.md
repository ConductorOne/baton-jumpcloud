# GsuitesListImportUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]GsuitesListImportUsers200ResponseUsersInner**](GsuitesListImportUsers200ResponseUsersInner.md) |  | [optional] 

## Methods

### NewGsuitesListImportUsers200Response

`func NewGsuitesListImportUsers200Response() *GsuitesListImportUsers200Response`

NewGsuitesListImportUsers200Response instantiates a new GsuitesListImportUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGsuitesListImportUsers200ResponseWithDefaults

`func NewGsuitesListImportUsers200ResponseWithDefaults() *GsuitesListImportUsers200Response`

NewGsuitesListImportUsers200ResponseWithDefaults instantiates a new GsuitesListImportUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *GsuitesListImportUsers200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *GsuitesListImportUsers200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *GsuitesListImportUsers200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *GsuitesListImportUsers200Response) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetUsers

`func (o *GsuitesListImportUsers200Response) GetUsers() []GsuitesListImportUsers200ResponseUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GsuitesListImportUsers200Response) GetUsersOk() (*[]GsuitesListImportUsers200ResponseUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GsuitesListImportUsers200Response) SetUsers(v []GsuitesListImportUsers200ResponseUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GsuitesListImportUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



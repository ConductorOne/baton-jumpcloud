# Office365sListImportUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipToken** | Pointer to **string** |  | [optional] 
**Top** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to [**[]Office365sListImportUsers200ResponseUsersInner**](Office365sListImportUsers200ResponseUsersInner.md) |  | [optional] 

## Methods

### NewOffice365sListImportUsers200Response

`func NewOffice365sListImportUsers200Response() *Office365sListImportUsers200Response`

NewOffice365sListImportUsers200Response instantiates a new Office365sListImportUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365sListImportUsers200ResponseWithDefaults

`func NewOffice365sListImportUsers200ResponseWithDefaults() *Office365sListImportUsers200Response`

NewOffice365sListImportUsers200ResponseWithDefaults instantiates a new Office365sListImportUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipToken

`func (o *Office365sListImportUsers200Response) GetSkipToken() string`

GetSkipToken returns the SkipToken field if non-nil, zero value otherwise.

### GetSkipTokenOk

`func (o *Office365sListImportUsers200Response) GetSkipTokenOk() (*string, bool)`

GetSkipTokenOk returns a tuple with the SkipToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipToken

`func (o *Office365sListImportUsers200Response) SetSkipToken(v string)`

SetSkipToken sets SkipToken field to given value.

### HasSkipToken

`func (o *Office365sListImportUsers200Response) HasSkipToken() bool`

HasSkipToken returns a boolean if a field has been set.

### GetTop

`func (o *Office365sListImportUsers200Response) GetTop() int32`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *Office365sListImportUsers200Response) GetTopOk() (*int32, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *Office365sListImportUsers200Response) SetTop(v int32)`

SetTop sets Top field to given value.

### HasTop

`func (o *Office365sListImportUsers200Response) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetUsers

`func (o *Office365sListImportUsers200Response) GetUsers() []Office365sListImportUsers200ResponseUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Office365sListImportUsers200Response) GetUsersOk() (*[]Office365sListImportUsers200ResponseUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Office365sListImportUsers200Response) SetUsers(v []Office365sListImportUsers200ResponseUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Office365sListImportUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



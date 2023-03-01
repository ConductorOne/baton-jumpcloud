# ImportUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUserReactivation** | Pointer to **bool** | A boolean value to allow the reactivation of suspended users | [optional] [default to true]
**Operations** | Pointer to [**[]ImportOperation**](ImportOperation.md) | Operations to be performed on the user list returned from the application | [optional] 
**QueryString** | Pointer to **string** | Query string to filter and sort the user list returned from the application.  The supported filtering and sorting varies by application.  If no value is sent, all users are returned. **Example:** \&quot;location&#x3D;Chicago&amp;department&#x3D;IT\&quot;Query string used to retrieve users from service | [optional] [default to ""]

## Methods

### NewImportUsersRequest

`func NewImportUsersRequest() *ImportUsersRequest`

NewImportUsersRequest instantiates a new ImportUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportUsersRequestWithDefaults

`func NewImportUsersRequestWithDefaults() *ImportUsersRequest`

NewImportUsersRequestWithDefaults instantiates a new ImportUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUserReactivation

`func (o *ImportUsersRequest) GetAllowUserReactivation() bool`

GetAllowUserReactivation returns the AllowUserReactivation field if non-nil, zero value otherwise.

### GetAllowUserReactivationOk

`func (o *ImportUsersRequest) GetAllowUserReactivationOk() (*bool, bool)`

GetAllowUserReactivationOk returns a tuple with the AllowUserReactivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserReactivation

`func (o *ImportUsersRequest) SetAllowUserReactivation(v bool)`

SetAllowUserReactivation sets AllowUserReactivation field to given value.

### HasAllowUserReactivation

`func (o *ImportUsersRequest) HasAllowUserReactivation() bool`

HasAllowUserReactivation returns a boolean if a field has been set.

### GetOperations

`func (o *ImportUsersRequest) GetOperations() []ImportOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ImportUsersRequest) GetOperationsOk() (*[]ImportOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ImportUsersRequest) SetOperations(v []ImportOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ImportUsersRequest) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetQueryString

`func (o *ImportUsersRequest) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *ImportUsersRequest) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *ImportUsersRequest) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *ImportUsersRequest) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Office365

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupsEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**UserLockoutAction** | Pointer to **string** |  | [optional] [readonly] 
**UserPasswordExpirationAction** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewOffice365

`func NewOffice365() *Office365`

NewOffice365 instantiates a new Office365 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365WithDefaults

`func NewOffice365WithDefaults() *Office365`

NewOffice365WithDefaults instantiates a new Office365 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupsEnabled

`func (o *Office365) GetGroupsEnabled() bool`

GetGroupsEnabled returns the GroupsEnabled field if non-nil, zero value otherwise.

### GetGroupsEnabledOk

`func (o *Office365) GetGroupsEnabledOk() (*bool, bool)`

GetGroupsEnabledOk returns a tuple with the GroupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsEnabled

`func (o *Office365) SetGroupsEnabled(v bool)`

SetGroupsEnabled sets GroupsEnabled field to given value.

### HasGroupsEnabled

`func (o *Office365) HasGroupsEnabled() bool`

HasGroupsEnabled returns a boolean if a field has been set.

### GetId

`func (o *Office365) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Office365) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Office365) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Office365) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Office365) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Office365) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Office365) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Office365) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserLockoutAction

`func (o *Office365) GetUserLockoutAction() string`

GetUserLockoutAction returns the UserLockoutAction field if non-nil, zero value otherwise.

### GetUserLockoutActionOk

`func (o *Office365) GetUserLockoutActionOk() (*string, bool)`

GetUserLockoutActionOk returns a tuple with the UserLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutAction

`func (o *Office365) SetUserLockoutAction(v string)`

SetUserLockoutAction sets UserLockoutAction field to given value.

### HasUserLockoutAction

`func (o *Office365) HasUserLockoutAction() bool`

HasUserLockoutAction returns a boolean if a field has been set.

### GetUserPasswordExpirationAction

`func (o *Office365) GetUserPasswordExpirationAction() string`

GetUserPasswordExpirationAction returns the UserPasswordExpirationAction field if non-nil, zero value otherwise.

### GetUserPasswordExpirationActionOk

`func (o *Office365) GetUserPasswordExpirationActionOk() (*string, bool)`

GetUserPasswordExpirationActionOk returns a tuple with the UserPasswordExpirationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordExpirationAction

`func (o *Office365) SetUserPasswordExpirationAction(v string)`

SetUserPasswordExpirationAction sets UserPasswordExpirationAction field to given value.

### HasUserPasswordExpirationAction

`func (o *Office365) HasUserPasswordExpirationAction() bool`

HasUserPasswordExpirationAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



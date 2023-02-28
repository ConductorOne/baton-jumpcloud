# SharedFolderUsersResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevelId** | **string** |  | 
**AccessLevelName** | **string** |  | 
**Apps** | Pointer to [**[]AppsInner**](AppsInner.md) |  | [optional] 
**Email** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Status** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewSharedFolderUsersResultsInner

`func NewSharedFolderUsersResultsInner(accessLevelId string, accessLevelName string, email string, id string, name string, status string, ) *SharedFolderUsersResultsInner`

NewSharedFolderUsersResultsInner instantiates a new SharedFolderUsersResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedFolderUsersResultsInnerWithDefaults

`func NewSharedFolderUsersResultsInnerWithDefaults() *SharedFolderUsersResultsInner`

NewSharedFolderUsersResultsInnerWithDefaults instantiates a new SharedFolderUsersResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevelId

`func (o *SharedFolderUsersResultsInner) GetAccessLevelId() string`

GetAccessLevelId returns the AccessLevelId field if non-nil, zero value otherwise.

### GetAccessLevelIdOk

`func (o *SharedFolderUsersResultsInner) GetAccessLevelIdOk() (*string, bool)`

GetAccessLevelIdOk returns a tuple with the AccessLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelId

`func (o *SharedFolderUsersResultsInner) SetAccessLevelId(v string)`

SetAccessLevelId sets AccessLevelId field to given value.


### GetAccessLevelName

`func (o *SharedFolderUsersResultsInner) GetAccessLevelName() string`

GetAccessLevelName returns the AccessLevelName field if non-nil, zero value otherwise.

### GetAccessLevelNameOk

`func (o *SharedFolderUsersResultsInner) GetAccessLevelNameOk() (*string, bool)`

GetAccessLevelNameOk returns a tuple with the AccessLevelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevelName

`func (o *SharedFolderUsersResultsInner) SetAccessLevelName(v string)`

SetAccessLevelName sets AccessLevelName field to given value.


### GetApps

`func (o *SharedFolderUsersResultsInner) GetApps() []AppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SharedFolderUsersResultsInner) GetAppsOk() (*[]AppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SharedFolderUsersResultsInner) SetApps(v []AppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SharedFolderUsersResultsInner) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetEmail

`func (o *SharedFolderUsersResultsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SharedFolderUsersResultsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SharedFolderUsersResultsInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalId

`func (o *SharedFolderUsersResultsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SharedFolderUsersResultsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SharedFolderUsersResultsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SharedFolderUsersResultsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *SharedFolderUsersResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedFolderUsersResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedFolderUsersResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SharedFolderUsersResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedFolderUsersResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedFolderUsersResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *SharedFolderUsersResultsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SharedFolderUsersResultsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SharedFolderUsersResultsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUsername

`func (o *SharedFolderUsersResultsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SharedFolderUsersResultsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SharedFolderUsersResultsInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SharedFolderUsersResultsInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



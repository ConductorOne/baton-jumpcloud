# SharedFolderUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SharedFolderUsersResultsInner**](SharedFolderUsersResultsInner.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewSharedFolderUsers

`func NewSharedFolderUsers(results []SharedFolderUsersResultsInner, totalCount int32, ) *SharedFolderUsers`

NewSharedFolderUsers instantiates a new SharedFolderUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedFolderUsersWithDefaults

`func NewSharedFolderUsersWithDefaults() *SharedFolderUsers`

NewSharedFolderUsersWithDefaults instantiates a new SharedFolderUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SharedFolderUsers) GetResults() []SharedFolderUsersResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SharedFolderUsers) GetResultsOk() (*[]SharedFolderUsersResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SharedFolderUsers) SetResults(v []SharedFolderUsersResultsInner)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *SharedFolderUsers) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SharedFolderUsers) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SharedFolderUsers) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



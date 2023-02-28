# SharedFoldersList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SharedFolder**](SharedFolder.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewSharedFoldersList

`func NewSharedFoldersList(results []SharedFolder, totalCount int32, ) *SharedFoldersList`

NewSharedFoldersList instantiates a new SharedFoldersList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedFoldersListWithDefaults

`func NewSharedFoldersListWithDefaults() *SharedFoldersList`

NewSharedFoldersListWithDefaults instantiates a new SharedFoldersList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SharedFoldersList) GetResults() []SharedFolder`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SharedFoldersList) GetResultsOk() (*[]SharedFolder, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SharedFoldersList) SetResults(v []SharedFolder)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *SharedFoldersList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SharedFoldersList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SharedFoldersList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



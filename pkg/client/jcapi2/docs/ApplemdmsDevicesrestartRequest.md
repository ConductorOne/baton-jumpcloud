# ApplemdmsDevicesrestartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KextPaths** | Pointer to **[]string** | The string to pass when doing a restart and performing a RebuildKernelCache. | [optional] 

## Methods

### NewApplemdmsDevicesrestartRequest

`func NewApplemdmsDevicesrestartRequest() *ApplemdmsDevicesrestartRequest`

NewApplemdmsDevicesrestartRequest instantiates a new ApplemdmsDevicesrestartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplemdmsDevicesrestartRequestWithDefaults

`func NewApplemdmsDevicesrestartRequestWithDefaults() *ApplemdmsDevicesrestartRequest`

NewApplemdmsDevicesrestartRequestWithDefaults instantiates a new ApplemdmsDevicesrestartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKextPaths

`func (o *ApplemdmsDevicesrestartRequest) GetKextPaths() []string`

GetKextPaths returns the KextPaths field if non-nil, zero value otherwise.

### GetKextPathsOk

`func (o *ApplemdmsDevicesrestartRequest) GetKextPathsOk() (*[]string, bool)`

GetKextPathsOk returns a tuple with the KextPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKextPaths

`func (o *ApplemdmsDevicesrestartRequest) SetKextPaths(v []string)`

SetKextPaths sets KextPaths field to given value.

### HasKextPaths

`func (o *ApplemdmsDevicesrestartRequest) HasKextPaths() bool`

HasKextPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



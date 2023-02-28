# WorkdayInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**AuthInput**](AuthInput.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReportUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkdayInput

`func NewWorkdayInput() *WorkdayInput`

NewWorkdayInput instantiates a new WorkdayInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkdayInputWithDefaults

`func NewWorkdayInputWithDefaults() *WorkdayInput`

NewWorkdayInputWithDefaults instantiates a new WorkdayInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *WorkdayInput) GetAuth() AuthInput`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *WorkdayInput) GetAuthOk() (*AuthInput, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *WorkdayInput) SetAuth(v AuthInput)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *WorkdayInput) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetName

`func (o *WorkdayInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkdayInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkdayInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkdayInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReportUrl

`func (o *WorkdayInput) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *WorkdayInput) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *WorkdayInput) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *WorkdayInput) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



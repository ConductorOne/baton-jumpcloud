# CloudInsightsEventsDistinctPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccounts** | **[]string** |  | 
**EndTime** | Pointer to [**CloudInsightsEventsDistinctPostEndTime**](CloudInsightsEventsDistinctPostEndTime.md) |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**NpeFilter** | Pointer to **bool** |  | [optional] 
**SearchTerm** | Pointer to [**CloudInsightsEventsSearchTerm**](CloudInsightsEventsSearchTerm.md) |  | [optional] 
**StartTime** | [**CloudInsightsEventsDistinctPostEndTime**](CloudInsightsEventsDistinctPostEndTime.md) |  | 

## Methods

### NewCloudInsightsEventsDistinctPost

`func NewCloudInsightsEventsDistinctPost(awsAccounts []string, startTime CloudInsightsEventsDistinctPostEndTime, ) *CloudInsightsEventsDistinctPost`

NewCloudInsightsEventsDistinctPost instantiates a new CloudInsightsEventsDistinctPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsEventsDistinctPostWithDefaults

`func NewCloudInsightsEventsDistinctPostWithDefaults() *CloudInsightsEventsDistinctPost`

NewCloudInsightsEventsDistinctPostWithDefaults instantiates a new CloudInsightsEventsDistinctPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccounts

`func (o *CloudInsightsEventsDistinctPost) GetAwsAccounts() []string`

GetAwsAccounts returns the AwsAccounts field if non-nil, zero value otherwise.

### GetAwsAccountsOk

`func (o *CloudInsightsEventsDistinctPost) GetAwsAccountsOk() (*[]string, bool)`

GetAwsAccountsOk returns a tuple with the AwsAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccounts

`func (o *CloudInsightsEventsDistinctPost) SetAwsAccounts(v []string)`

SetAwsAccounts sets AwsAccounts field to given value.


### GetEndTime

`func (o *CloudInsightsEventsDistinctPost) GetEndTime() CloudInsightsEventsDistinctPostEndTime`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CloudInsightsEventsDistinctPost) GetEndTimeOk() (*CloudInsightsEventsDistinctPostEndTime, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CloudInsightsEventsDistinctPost) SetEndTime(v CloudInsightsEventsDistinctPostEndTime)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CloudInsightsEventsDistinctPost) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetField

`func (o *CloudInsightsEventsDistinctPost) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *CloudInsightsEventsDistinctPost) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *CloudInsightsEventsDistinctPost) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *CloudInsightsEventsDistinctPost) HasField() bool`

HasField returns a boolean if a field has been set.

### GetNpeFilter

`func (o *CloudInsightsEventsDistinctPost) GetNpeFilter() bool`

GetNpeFilter returns the NpeFilter field if non-nil, zero value otherwise.

### GetNpeFilterOk

`func (o *CloudInsightsEventsDistinctPost) GetNpeFilterOk() (*bool, bool)`

GetNpeFilterOk returns a tuple with the NpeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpeFilter

`func (o *CloudInsightsEventsDistinctPost) SetNpeFilter(v bool)`

SetNpeFilter sets NpeFilter field to given value.

### HasNpeFilter

`func (o *CloudInsightsEventsDistinctPost) HasNpeFilter() bool`

HasNpeFilter returns a boolean if a field has been set.

### GetSearchTerm

`func (o *CloudInsightsEventsDistinctPost) GetSearchTerm() CloudInsightsEventsSearchTerm`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *CloudInsightsEventsDistinctPost) GetSearchTermOk() (*CloudInsightsEventsSearchTerm, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *CloudInsightsEventsDistinctPost) SetSearchTerm(v CloudInsightsEventsSearchTerm)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *CloudInsightsEventsDistinctPost) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.

### GetStartTime

`func (o *CloudInsightsEventsDistinctPost) GetStartTime() CloudInsightsEventsDistinctPostEndTime`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CloudInsightsEventsDistinctPost) GetStartTimeOk() (*CloudInsightsEventsDistinctPostEndTime, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CloudInsightsEventsDistinctPost) SetStartTime(v CloudInsightsEventsDistinctPostEndTime)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudInsightsEventsListPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccounts** | Pointer to **[]string** |  | [optional] 
**EndTime** | Pointer to [**CloudInsightsEventsDistinctPostEndTime**](CloudInsightsEventsDistinctPostEndTime.md) |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **float32** |  | [optional] 
**NpeFilter** | Pointer to **bool** |  | [optional] 
**SearchAfter** | Pointer to **map[string]interface{}** |  | [optional] 
**SearchTerm** | Pointer to [**CloudInsightsEventsSearchTerm**](CloudInsightsEventsSearchTerm.md) |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Skip** | Pointer to **float32** |  | [optional] 
**Sort** | Pointer to **map[string]interface{}** |  | [optional] 
**StartTime** | Pointer to [**CloudInsightsEventsDistinctPostEndTime**](CloudInsightsEventsDistinctPostEndTime.md) |  | [optional] 

## Methods

### NewCloudInsightsEventsListPost

`func NewCloudInsightsEventsListPost() *CloudInsightsEventsListPost`

NewCloudInsightsEventsListPost instantiates a new CloudInsightsEventsListPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsEventsListPostWithDefaults

`func NewCloudInsightsEventsListPostWithDefaults() *CloudInsightsEventsListPost`

NewCloudInsightsEventsListPostWithDefaults instantiates a new CloudInsightsEventsListPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccounts

`func (o *CloudInsightsEventsListPost) GetAwsAccounts() []string`

GetAwsAccounts returns the AwsAccounts field if non-nil, zero value otherwise.

### GetAwsAccountsOk

`func (o *CloudInsightsEventsListPost) GetAwsAccountsOk() (*[]string, bool)`

GetAwsAccountsOk returns a tuple with the AwsAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccounts

`func (o *CloudInsightsEventsListPost) SetAwsAccounts(v []string)`

SetAwsAccounts sets AwsAccounts field to given value.

### HasAwsAccounts

`func (o *CloudInsightsEventsListPost) HasAwsAccounts() bool`

HasAwsAccounts returns a boolean if a field has been set.

### GetEndTime

`func (o *CloudInsightsEventsListPost) GetEndTime() CloudInsightsEventsDistinctPostEndTime`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CloudInsightsEventsListPost) GetEndTimeOk() (*CloudInsightsEventsDistinctPostEndTime, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CloudInsightsEventsListPost) SetEndTime(v CloudInsightsEventsDistinctPostEndTime)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CloudInsightsEventsListPost) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFields

`func (o *CloudInsightsEventsListPost) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CloudInsightsEventsListPost) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CloudInsightsEventsListPost) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CloudInsightsEventsListPost) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLimit

`func (o *CloudInsightsEventsListPost) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CloudInsightsEventsListPost) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CloudInsightsEventsListPost) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CloudInsightsEventsListPost) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNpeFilter

`func (o *CloudInsightsEventsListPost) GetNpeFilter() bool`

GetNpeFilter returns the NpeFilter field if non-nil, zero value otherwise.

### GetNpeFilterOk

`func (o *CloudInsightsEventsListPost) GetNpeFilterOk() (*bool, bool)`

GetNpeFilterOk returns a tuple with the NpeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpeFilter

`func (o *CloudInsightsEventsListPost) SetNpeFilter(v bool)`

SetNpeFilter sets NpeFilter field to given value.

### HasNpeFilter

`func (o *CloudInsightsEventsListPost) HasNpeFilter() bool`

HasNpeFilter returns a boolean if a field has been set.

### GetSearchAfter

`func (o *CloudInsightsEventsListPost) GetSearchAfter() map[string]interface{}`

GetSearchAfter returns the SearchAfter field if non-nil, zero value otherwise.

### GetSearchAfterOk

`func (o *CloudInsightsEventsListPost) GetSearchAfterOk() (*map[string]interface{}, bool)`

GetSearchAfterOk returns a tuple with the SearchAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAfter

`func (o *CloudInsightsEventsListPost) SetSearchAfter(v map[string]interface{})`

SetSearchAfter sets SearchAfter field to given value.

### HasSearchAfter

`func (o *CloudInsightsEventsListPost) HasSearchAfter() bool`

HasSearchAfter returns a boolean if a field has been set.

### GetSearchTerm

`func (o *CloudInsightsEventsListPost) GetSearchTerm() CloudInsightsEventsSearchTerm`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *CloudInsightsEventsListPost) GetSearchTermOk() (*CloudInsightsEventsSearchTerm, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *CloudInsightsEventsListPost) SetSearchTerm(v CloudInsightsEventsSearchTerm)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *CloudInsightsEventsListPost) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.

### GetService

`func (o *CloudInsightsEventsListPost) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudInsightsEventsListPost) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudInsightsEventsListPost) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudInsightsEventsListPost) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSkip

`func (o *CloudInsightsEventsListPost) GetSkip() float32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *CloudInsightsEventsListPost) GetSkipOk() (*float32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *CloudInsightsEventsListPost) SetSkip(v float32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *CloudInsightsEventsListPost) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSort

`func (o *CloudInsightsEventsListPost) GetSort() map[string]interface{}`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CloudInsightsEventsListPost) GetSortOk() (*map[string]interface{}, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CloudInsightsEventsListPost) SetSort(v map[string]interface{})`

SetSort sets Sort field to given value.

### HasSort

`func (o *CloudInsightsEventsListPost) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *CloudInsightsEventsListPost) GetStartTime() CloudInsightsEventsDistinctPostEndTime`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CloudInsightsEventsListPost) GetStartTimeOk() (*CloudInsightsEventsDistinctPostEndTime, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CloudInsightsEventsListPost) SetStartTime(v CloudInsightsEventsDistinctPostEndTime)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CloudInsightsEventsListPost) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



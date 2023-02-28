# CloudInsightsEventsSearchTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AndConjunction** | Pointer to [**CloudInsightsTermConjunction**](CloudInsightsTermConjunction.md) |  | [optional] 
**OrConjunction** | Pointer to [**CloudInsightsTermConjunction**](CloudInsightsTermConjunction.md) |  | [optional] 
**SingleTerm** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudInsightsEventsSearchTerm

`func NewCloudInsightsEventsSearchTerm() *CloudInsightsEventsSearchTerm`

NewCloudInsightsEventsSearchTerm instantiates a new CloudInsightsEventsSearchTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsEventsSearchTermWithDefaults

`func NewCloudInsightsEventsSearchTermWithDefaults() *CloudInsightsEventsSearchTerm`

NewCloudInsightsEventsSearchTermWithDefaults instantiates a new CloudInsightsEventsSearchTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAndConjunction

`func (o *CloudInsightsEventsSearchTerm) GetAndConjunction() CloudInsightsTermConjunction`

GetAndConjunction returns the AndConjunction field if non-nil, zero value otherwise.

### GetAndConjunctionOk

`func (o *CloudInsightsEventsSearchTerm) GetAndConjunctionOk() (*CloudInsightsTermConjunction, bool)`

GetAndConjunctionOk returns a tuple with the AndConjunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndConjunction

`func (o *CloudInsightsEventsSearchTerm) SetAndConjunction(v CloudInsightsTermConjunction)`

SetAndConjunction sets AndConjunction field to given value.

### HasAndConjunction

`func (o *CloudInsightsEventsSearchTerm) HasAndConjunction() bool`

HasAndConjunction returns a boolean if a field has been set.

### GetOrConjunction

`func (o *CloudInsightsEventsSearchTerm) GetOrConjunction() CloudInsightsTermConjunction`

GetOrConjunction returns the OrConjunction field if non-nil, zero value otherwise.

### GetOrConjunctionOk

`func (o *CloudInsightsEventsSearchTerm) GetOrConjunctionOk() (*CloudInsightsTermConjunction, bool)`

GetOrConjunctionOk returns a tuple with the OrConjunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrConjunction

`func (o *CloudInsightsEventsSearchTerm) SetOrConjunction(v CloudInsightsTermConjunction)`

SetOrConjunction sets OrConjunction field to given value.

### HasOrConjunction

`func (o *CloudInsightsEventsSearchTerm) HasOrConjunction() bool`

HasOrConjunction returns a boolean if a field has been set.

### GetSingleTerm

`func (o *CloudInsightsEventsSearchTerm) GetSingleTerm() string`

GetSingleTerm returns the SingleTerm field if non-nil, zero value otherwise.

### GetSingleTermOk

`func (o *CloudInsightsEventsSearchTerm) GetSingleTermOk() (*string, bool)`

GetSingleTermOk returns a tuple with the SingleTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleTerm

`func (o *CloudInsightsEventsSearchTerm) SetSingleTerm(v string)`

SetSingleTerm sets SingleTerm field to given value.

### HasSingleTerm

`func (o *CloudInsightsEventsSearchTerm) HasSingleTerm() bool`

HasSingleTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



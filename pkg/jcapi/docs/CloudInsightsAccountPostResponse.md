# CloudInsightsAccountPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**CloudInsightsAccountGet**](CloudInsightsAccountGet.md) |  | [optional] 
**CftUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudInsightsAccountPostResponse

`func NewCloudInsightsAccountPostResponse() *CloudInsightsAccountPostResponse`

NewCloudInsightsAccountPostResponse instantiates a new CloudInsightsAccountPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsAccountPostResponseWithDefaults

`func NewCloudInsightsAccountPostResponseWithDefaults() *CloudInsightsAccountPostResponse`

NewCloudInsightsAccountPostResponseWithDefaults instantiates a new CloudInsightsAccountPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudInsightsAccountPostResponse) GetAccount() CloudInsightsAccountGet`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudInsightsAccountPostResponse) GetAccountOk() (*CloudInsightsAccountGet, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudInsightsAccountPostResponse) SetAccount(v CloudInsightsAccountGet)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudInsightsAccountPostResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCftUrl

`func (o *CloudInsightsAccountPostResponse) GetCftUrl() string`

GetCftUrl returns the CftUrl field if non-nil, zero value otherwise.

### GetCftUrlOk

`func (o *CloudInsightsAccountPostResponse) GetCftUrlOk() (*string, bool)`

GetCftUrlOk returns a tuple with the CftUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCftUrl

`func (o *CloudInsightsAccountPostResponse) SetCftUrl(v string)`

SetCftUrl sets CftUrl field to given value.

### HasCftUrl

`func (o *CloudInsightsAccountPostResponse) HasCftUrl() bool`

HasCftUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



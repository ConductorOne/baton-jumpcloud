# CloudInsightsAccountListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]CloudInsightsAccountGet**](CloudInsightsAccountGet.md) |  | [optional] 
**TotalCount** | Pointer to **float32** |  | [optional] 

## Methods

### NewCloudInsightsAccountListResponse

`func NewCloudInsightsAccountListResponse() *CloudInsightsAccountListResponse`

NewCloudInsightsAccountListResponse instantiates a new CloudInsightsAccountListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsAccountListResponseWithDefaults

`func NewCloudInsightsAccountListResponseWithDefaults() *CloudInsightsAccountListResponse`

NewCloudInsightsAccountListResponseWithDefaults instantiates a new CloudInsightsAccountListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *CloudInsightsAccountListResponse) GetAccounts() []CloudInsightsAccountGet`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *CloudInsightsAccountListResponse) GetAccountsOk() (*[]CloudInsightsAccountGet, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *CloudInsightsAccountListResponse) SetAccounts(v []CloudInsightsAccountGet)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *CloudInsightsAccountListResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetTotalCount

`func (o *CloudInsightsAccountListResponse) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CloudInsightsAccountListResponse) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CloudInsightsAccountListResponse) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CloudInsightsAccountListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ProviderInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]ProviderInvoice**](ProviderInvoice.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewProviderInvoiceResponse

`func NewProviderInvoiceResponse() *ProviderInvoiceResponse`

NewProviderInvoiceResponse instantiates a new ProviderInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderInvoiceResponseWithDefaults

`func NewProviderInvoiceResponseWithDefaults() *ProviderInvoiceResponse`

NewProviderInvoiceResponseWithDefaults instantiates a new ProviderInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *ProviderInvoiceResponse) GetRecords() []ProviderInvoice`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ProviderInvoiceResponse) GetRecordsOk() (*[]ProviderInvoice, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ProviderInvoiceResponse) SetRecords(v []ProviderInvoice)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ProviderInvoiceResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotalCount

`func (o *ProviderInvoiceResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProviderInvoiceResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProviderInvoiceResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ProviderInvoiceResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



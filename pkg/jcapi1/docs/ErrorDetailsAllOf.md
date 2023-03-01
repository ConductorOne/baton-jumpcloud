# ErrorDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **[]map[string]interface{}** | Describes a list of objects with more detailed information of the given error. Each detail schema is according to one of the messages defined in Google&#39;s API: https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto\&quot;  | [optional] 

## Methods

### NewErrorDetailsAllOf

`func NewErrorDetailsAllOf() *ErrorDetailsAllOf`

NewErrorDetailsAllOf instantiates a new ErrorDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailsAllOfWithDefaults

`func NewErrorDetailsAllOfWithDefaults() *ErrorDetailsAllOf`

NewErrorDetailsAllOfWithDefaults instantiates a new ErrorDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ErrorDetailsAllOf) GetDetails() []map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorDetailsAllOf) GetDetailsOk() (*[]map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorDetailsAllOf) SetDetails(v []map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorDetailsAllOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



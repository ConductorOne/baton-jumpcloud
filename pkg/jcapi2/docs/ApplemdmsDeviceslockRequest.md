# ApplemdmsDeviceslockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | Pointer to **string** | 6-digit PIN, required for MacOS, to lock the device | [optional] 

## Methods

### NewApplemdmsDeviceslockRequest

`func NewApplemdmsDeviceslockRequest() *ApplemdmsDeviceslockRequest`

NewApplemdmsDeviceslockRequest instantiates a new ApplemdmsDeviceslockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplemdmsDeviceslockRequestWithDefaults

`func NewApplemdmsDeviceslockRequestWithDefaults() *ApplemdmsDeviceslockRequest`

NewApplemdmsDeviceslockRequestWithDefaults instantiates a new ApplemdmsDeviceslockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *ApplemdmsDeviceslockRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *ApplemdmsDeviceslockRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *ApplemdmsDeviceslockRequest) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *ApplemdmsDeviceslockRequest) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



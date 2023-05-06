# JumpcloudGoogleEmmNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imei** | Pointer to **string** | Not logging as it contains sensitive information. | [optional] 
**Meid** | Pointer to **string** | Not logging as it contains sensitive information. | [optional] 
**TelephonyInfo** | Pointer to [**[]JumpcloudGoogleEmmTelephonyInfo**](JumpcloudGoogleEmmTelephonyInfo.md) |  | [optional] 
**WifiMacAddress** | Pointer to **string** | Not logging as it contains sensitive information. | [optional] 

## Methods

### NewJumpcloudGoogleEmmNetworkInfo

`func NewJumpcloudGoogleEmmNetworkInfo() *JumpcloudGoogleEmmNetworkInfo`

NewJumpcloudGoogleEmmNetworkInfo instantiates a new JumpcloudGoogleEmmNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJumpcloudGoogleEmmNetworkInfoWithDefaults

`func NewJumpcloudGoogleEmmNetworkInfoWithDefaults() *JumpcloudGoogleEmmNetworkInfo`

NewJumpcloudGoogleEmmNetworkInfoWithDefaults instantiates a new JumpcloudGoogleEmmNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImei

`func (o *JumpcloudGoogleEmmNetworkInfo) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *JumpcloudGoogleEmmNetworkInfo) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *JumpcloudGoogleEmmNetworkInfo) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *JumpcloudGoogleEmmNetworkInfo) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetMeid

`func (o *JumpcloudGoogleEmmNetworkInfo) GetMeid() string`

GetMeid returns the Meid field if non-nil, zero value otherwise.

### GetMeidOk

`func (o *JumpcloudGoogleEmmNetworkInfo) GetMeidOk() (*string, bool)`

GetMeidOk returns a tuple with the Meid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeid

`func (o *JumpcloudGoogleEmmNetworkInfo) SetMeid(v string)`

SetMeid sets Meid field to given value.

### HasMeid

`func (o *JumpcloudGoogleEmmNetworkInfo) HasMeid() bool`

HasMeid returns a boolean if a field has been set.

### GetTelephonyInfo

`func (o *JumpcloudGoogleEmmNetworkInfo) GetTelephonyInfo() []JumpcloudGoogleEmmTelephonyInfo`

GetTelephonyInfo returns the TelephonyInfo field if non-nil, zero value otherwise.

### GetTelephonyInfoOk

`func (o *JumpcloudGoogleEmmNetworkInfo) GetTelephonyInfoOk() (*[]JumpcloudGoogleEmmTelephonyInfo, bool)`

GetTelephonyInfoOk returns a tuple with the TelephonyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephonyInfo

`func (o *JumpcloudGoogleEmmNetworkInfo) SetTelephonyInfo(v []JumpcloudGoogleEmmTelephonyInfo)`

SetTelephonyInfo sets TelephonyInfo field to given value.

### HasTelephonyInfo

`func (o *JumpcloudGoogleEmmNetworkInfo) HasTelephonyInfo() bool`

HasTelephonyInfo returns a boolean if a field has been set.

### GetWifiMacAddress

`func (o *JumpcloudGoogleEmmNetworkInfo) GetWifiMacAddress() string`

GetWifiMacAddress returns the WifiMacAddress field if non-nil, zero value otherwise.

### GetWifiMacAddressOk

`func (o *JumpcloudGoogleEmmNetworkInfo) GetWifiMacAddressOk() (*string, bool)`

GetWifiMacAddressOk returns a tuple with the WifiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacAddress

`func (o *JumpcloudGoogleEmmNetworkInfo) SetWifiMacAddress(v string)`

SetWifiMacAddress sets WifiMacAddress field to given value.

### HasWifiMacAddress

`func (o *JumpcloudGoogleEmmNetworkInfo) HasWifiMacAddress() bool`

HasWifiMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AppleMdmDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationLockAllowedWhileSupervised** | Pointer to **bool** |  | [optional] 
**AvailableDeviceCapacity** | Pointer to **float32** |  | [optional] 
**DeviceCapacity** | Pointer to **float32** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**Iccid** | Pointer to **string** |  | [optional] 
**Imei** | Pointer to **string** |  | [optional] 
**IsSupervised** | Pointer to **bool** |  | [optional] 
**ModelName** | Pointer to **string** |  | [optional] 
**SecondIccid** | Pointer to **string** |  | [optional] 
**SecondImei** | Pointer to **string** |  | [optional] 
**SecondSubscriberCarrierNetwork** | Pointer to **string** |  | [optional] 
**SubscriberCarrierNetwork** | Pointer to **string** |  | [optional] 
**WifiMac** | Pointer to **string** |  | [optional] 

## Methods

### NewAppleMdmDeviceInfo

`func NewAppleMdmDeviceInfo() *AppleMdmDeviceInfo`

NewAppleMdmDeviceInfo instantiates a new AppleMdmDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleMdmDeviceInfoWithDefaults

`func NewAppleMdmDeviceInfoWithDefaults() *AppleMdmDeviceInfo`

NewAppleMdmDeviceInfoWithDefaults instantiates a new AppleMdmDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationLockAllowedWhileSupervised

`func (o *AppleMdmDeviceInfo) GetActivationLockAllowedWhileSupervised() bool`

GetActivationLockAllowedWhileSupervised returns the ActivationLockAllowedWhileSupervised field if non-nil, zero value otherwise.

### GetActivationLockAllowedWhileSupervisedOk

`func (o *AppleMdmDeviceInfo) GetActivationLockAllowedWhileSupervisedOk() (*bool, bool)`

GetActivationLockAllowedWhileSupervisedOk returns a tuple with the ActivationLockAllowedWhileSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLockAllowedWhileSupervised

`func (o *AppleMdmDeviceInfo) SetActivationLockAllowedWhileSupervised(v bool)`

SetActivationLockAllowedWhileSupervised sets ActivationLockAllowedWhileSupervised field to given value.

### HasActivationLockAllowedWhileSupervised

`func (o *AppleMdmDeviceInfo) HasActivationLockAllowedWhileSupervised() bool`

HasActivationLockAllowedWhileSupervised returns a boolean if a field has been set.

### GetAvailableDeviceCapacity

`func (o *AppleMdmDeviceInfo) GetAvailableDeviceCapacity() float32`

GetAvailableDeviceCapacity returns the AvailableDeviceCapacity field if non-nil, zero value otherwise.

### GetAvailableDeviceCapacityOk

`func (o *AppleMdmDeviceInfo) GetAvailableDeviceCapacityOk() (*float32, bool)`

GetAvailableDeviceCapacityOk returns a tuple with the AvailableDeviceCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDeviceCapacity

`func (o *AppleMdmDeviceInfo) SetAvailableDeviceCapacity(v float32)`

SetAvailableDeviceCapacity sets AvailableDeviceCapacity field to given value.

### HasAvailableDeviceCapacity

`func (o *AppleMdmDeviceInfo) HasAvailableDeviceCapacity() bool`

HasAvailableDeviceCapacity returns a boolean if a field has been set.

### GetDeviceCapacity

`func (o *AppleMdmDeviceInfo) GetDeviceCapacity() float32`

GetDeviceCapacity returns the DeviceCapacity field if non-nil, zero value otherwise.

### GetDeviceCapacityOk

`func (o *AppleMdmDeviceInfo) GetDeviceCapacityOk() (*float32, bool)`

GetDeviceCapacityOk returns a tuple with the DeviceCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCapacity

`func (o *AppleMdmDeviceInfo) SetDeviceCapacity(v float32)`

SetDeviceCapacity sets DeviceCapacity field to given value.

### HasDeviceCapacity

`func (o *AppleMdmDeviceInfo) HasDeviceCapacity() bool`

HasDeviceCapacity returns a boolean if a field has been set.

### GetDeviceName

`func (o *AppleMdmDeviceInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *AppleMdmDeviceInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *AppleMdmDeviceInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *AppleMdmDeviceInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetIccid

`func (o *AppleMdmDeviceInfo) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *AppleMdmDeviceInfo) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *AppleMdmDeviceInfo) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *AppleMdmDeviceInfo) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetImei

`func (o *AppleMdmDeviceInfo) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *AppleMdmDeviceInfo) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *AppleMdmDeviceInfo) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *AppleMdmDeviceInfo) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetIsSupervised

`func (o *AppleMdmDeviceInfo) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *AppleMdmDeviceInfo) GetIsSupervisedOk() (*bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupervised

`func (o *AppleMdmDeviceInfo) SetIsSupervised(v bool)`

SetIsSupervised sets IsSupervised field to given value.

### HasIsSupervised

`func (o *AppleMdmDeviceInfo) HasIsSupervised() bool`

HasIsSupervised returns a boolean if a field has been set.

### GetModelName

`func (o *AppleMdmDeviceInfo) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *AppleMdmDeviceInfo) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *AppleMdmDeviceInfo) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *AppleMdmDeviceInfo) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetSecondIccid

`func (o *AppleMdmDeviceInfo) GetSecondIccid() string`

GetSecondIccid returns the SecondIccid field if non-nil, zero value otherwise.

### GetSecondIccidOk

`func (o *AppleMdmDeviceInfo) GetSecondIccidOk() (*string, bool)`

GetSecondIccidOk returns a tuple with the SecondIccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondIccid

`func (o *AppleMdmDeviceInfo) SetSecondIccid(v string)`

SetSecondIccid sets SecondIccid field to given value.

### HasSecondIccid

`func (o *AppleMdmDeviceInfo) HasSecondIccid() bool`

HasSecondIccid returns a boolean if a field has been set.

### GetSecondImei

`func (o *AppleMdmDeviceInfo) GetSecondImei() string`

GetSecondImei returns the SecondImei field if non-nil, zero value otherwise.

### GetSecondImeiOk

`func (o *AppleMdmDeviceInfo) GetSecondImeiOk() (*string, bool)`

GetSecondImeiOk returns a tuple with the SecondImei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondImei

`func (o *AppleMdmDeviceInfo) SetSecondImei(v string)`

SetSecondImei sets SecondImei field to given value.

### HasSecondImei

`func (o *AppleMdmDeviceInfo) HasSecondImei() bool`

HasSecondImei returns a boolean if a field has been set.

### GetSecondSubscriberCarrierNetwork

`func (o *AppleMdmDeviceInfo) GetSecondSubscriberCarrierNetwork() string`

GetSecondSubscriberCarrierNetwork returns the SecondSubscriberCarrierNetwork field if non-nil, zero value otherwise.

### GetSecondSubscriberCarrierNetworkOk

`func (o *AppleMdmDeviceInfo) GetSecondSubscriberCarrierNetworkOk() (*string, bool)`

GetSecondSubscriberCarrierNetworkOk returns a tuple with the SecondSubscriberCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondSubscriberCarrierNetwork

`func (o *AppleMdmDeviceInfo) SetSecondSubscriberCarrierNetwork(v string)`

SetSecondSubscriberCarrierNetwork sets SecondSubscriberCarrierNetwork field to given value.

### HasSecondSubscriberCarrierNetwork

`func (o *AppleMdmDeviceInfo) HasSecondSubscriberCarrierNetwork() bool`

HasSecondSubscriberCarrierNetwork returns a boolean if a field has been set.

### GetSubscriberCarrierNetwork

`func (o *AppleMdmDeviceInfo) GetSubscriberCarrierNetwork() string`

GetSubscriberCarrierNetwork returns the SubscriberCarrierNetwork field if non-nil, zero value otherwise.

### GetSubscriberCarrierNetworkOk

`func (o *AppleMdmDeviceInfo) GetSubscriberCarrierNetworkOk() (*string, bool)`

GetSubscriberCarrierNetworkOk returns a tuple with the SubscriberCarrierNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCarrierNetwork

`func (o *AppleMdmDeviceInfo) SetSubscriberCarrierNetwork(v string)`

SetSubscriberCarrierNetwork sets SubscriberCarrierNetwork field to given value.

### HasSubscriberCarrierNetwork

`func (o *AppleMdmDeviceInfo) HasSubscriberCarrierNetwork() bool`

HasSubscriberCarrierNetwork returns a boolean if a field has been set.

### GetWifiMac

`func (o *AppleMdmDeviceInfo) GetWifiMac() string`

GetWifiMac returns the WifiMac field if non-nil, zero value otherwise.

### GetWifiMacOk

`func (o *AppleMdmDeviceInfo) GetWifiMacOk() (*string, bool)`

GetWifiMacOk returns a tuple with the WifiMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMac

`func (o *AppleMdmDeviceInfo) SetWifiMac(v string)`

SetWifiMac sets WifiMac field to given value.

### HasWifiMac

`func (o *AppleMdmDeviceInfo) HasWifiMac() bool`

HasWifiMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



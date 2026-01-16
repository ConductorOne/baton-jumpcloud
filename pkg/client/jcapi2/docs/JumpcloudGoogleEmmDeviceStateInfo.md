# JumpcloudGoogleEmmDeviceStateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedDeviceState** | Pointer to **string** |  | [optional] 
**CommonCriteriaModeInfo** | Pointer to [**JumpcloudGoogleEmmCommonCriteriaModeInfo**](JumpcloudGoogleEmmCommonCriteriaModeInfo.md) |  | [optional] 
**DeviceSettings** | Pointer to [**JumpcloudGoogleEmmDeviceSettings**](JumpcloudGoogleEmmDeviceSettings.md) |  | [optional] 
**DeviceState** | Pointer to **string** |  | [optional] 
**DisabledReason** | Pointer to **string** |  | [optional] 
**LastPolicySyncTime** | Pointer to **string** |  | [optional] 
**LastStatusReportTime** | Pointer to **string** |  | [optional] 
**PolicyCompliant** | Pointer to **bool** |  | [optional] 
**SecurityPosture** | Pointer to [**JumpcloudGoogleEmmSecurityPosture**](JumpcloudGoogleEmmSecurityPosture.md) |  | [optional] 

## Methods

### NewJumpcloudGoogleEmmDeviceStateInfo

`func NewJumpcloudGoogleEmmDeviceStateInfo() *JumpcloudGoogleEmmDeviceStateInfo`

NewJumpcloudGoogleEmmDeviceStateInfo instantiates a new JumpcloudGoogleEmmDeviceStateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJumpcloudGoogleEmmDeviceStateInfoWithDefaults

`func NewJumpcloudGoogleEmmDeviceStateInfoWithDefaults() *JumpcloudGoogleEmmDeviceStateInfo`

NewJumpcloudGoogleEmmDeviceStateInfoWithDefaults instantiates a new JumpcloudGoogleEmmDeviceStateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedDeviceState

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetAppliedDeviceState() string`

GetAppliedDeviceState returns the AppliedDeviceState field if non-nil, zero value otherwise.

### GetAppliedDeviceStateOk

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetAppliedDeviceStateOk() (*string, bool)`

GetAppliedDeviceStateOk returns a tuple with the AppliedDeviceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedDeviceState

`func (o *JumpcloudGoogleEmmDeviceStateInfo) SetAppliedDeviceState(v string)`

SetAppliedDeviceState sets AppliedDeviceState field to given value.

### HasAppliedDeviceState

`func (o *JumpcloudGoogleEmmDeviceStateInfo) HasAppliedDeviceState() bool`

HasAppliedDeviceState returns a boolean if a field has been set.

### GetCommonCriteriaModeInfo

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetCommonCriteriaModeInfo() JumpcloudGoogleEmmCommonCriteriaModeInfo`

GetCommonCriteriaModeInfo returns the CommonCriteriaModeInfo field if non-nil, zero value otherwise.

### GetCommonCriteriaModeInfoOk

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetCommonCriteriaModeInfoOk() (*JumpcloudGoogleEmmCommonCriteriaModeInfo, bool)`

GetCommonCriteriaModeInfoOk returns a tuple with the CommonCriteriaModeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonCriteriaModeInfo

`func (o *JumpcloudGoogleEmmDeviceStateInfo) SetCommonCriteriaModeInfo(v JumpcloudGoogleEmmCommonCriteriaModeInfo)`

SetCommonCriteriaModeInfo sets CommonCriteriaModeInfo field to given value.

### HasCommonCriteriaModeInfo

`func (o *JumpcloudGoogleEmmDeviceStateInfo) HasCommonCriteriaModeInfo() bool`

HasCommonCriteriaModeInfo returns a boolean if a field has been set.

### GetDeviceSettings

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetDeviceSettings() JumpcloudGoogleEmmDeviceSettings`

GetDeviceSettings returns the DeviceSettings field if non-nil, zero value otherwise.

### GetDeviceSettingsOk

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetDeviceSettingsOk() (*JumpcloudGoogleEmmDeviceSettings, bool)`

GetDeviceSettingsOk returns a tuple with the DeviceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSettings

`func (o *JumpcloudGoogleEmmDeviceStateInfo) SetDeviceSettings(v JumpcloudGoogleEmmDeviceSettings)`

SetDeviceSettings sets DeviceSettings field to given value.

### HasDeviceSettings

`func (o *JumpcloudGoogleEmmDeviceStateInfo) HasDeviceSettings() bool`

HasDeviceSettings returns a boolean if a field has been set.

### GetDeviceState

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetDeviceState() string`

GetDeviceState returns the DeviceState field if non-nil, zero value otherwise.

### GetDeviceStateOk

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetDeviceStateOk() (*string, bool)`

GetDeviceStateOk returns a tuple with the DeviceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceState

`func (o *JumpcloudGoogleEmmDeviceStateInfo) SetDeviceState(v string)`

SetDeviceState sets DeviceState field to given value.

### HasDeviceState

`func (o *JumpcloudGoogleEmmDeviceStateInfo) HasDeviceState() bool`

HasDeviceState returns a boolean if a field has been set.

### GetDisabledReason

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetDisabledReason() string`

GetDisabledReason returns the DisabledReason field if non-nil, zero value otherwise.

### GetDisabledReasonOk

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetDisabledReasonOk() (*string, bool)`

GetDisabledReasonOk returns a tuple with the DisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReason

`func (o *JumpcloudGoogleEmmDeviceStateInfo) SetDisabledReason(v string)`

SetDisabledReason sets DisabledReason field to given value.

### HasDisabledReason

`func (o *JumpcloudGoogleEmmDeviceStateInfo) HasDisabledReason() bool`

HasDisabledReason returns a boolean if a field has been set.

### GetLastPolicySyncTime

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetLastPolicySyncTime() string`

GetLastPolicySyncTime returns the LastPolicySyncTime field if non-nil, zero value otherwise.

### GetLastPolicySyncTimeOk

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetLastPolicySyncTimeOk() (*string, bool)`

GetLastPolicySyncTimeOk returns a tuple with the LastPolicySyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPolicySyncTime

`func (o *JumpcloudGoogleEmmDeviceStateInfo) SetLastPolicySyncTime(v string)`

SetLastPolicySyncTime sets LastPolicySyncTime field to given value.

### HasLastPolicySyncTime

`func (o *JumpcloudGoogleEmmDeviceStateInfo) HasLastPolicySyncTime() bool`

HasLastPolicySyncTime returns a boolean if a field has been set.

### GetLastStatusReportTime

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetLastStatusReportTime() string`

GetLastStatusReportTime returns the LastStatusReportTime field if non-nil, zero value otherwise.

### GetLastStatusReportTimeOk

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetLastStatusReportTimeOk() (*string, bool)`

GetLastStatusReportTimeOk returns a tuple with the LastStatusReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusReportTime

`func (o *JumpcloudGoogleEmmDeviceStateInfo) SetLastStatusReportTime(v string)`

SetLastStatusReportTime sets LastStatusReportTime field to given value.

### HasLastStatusReportTime

`func (o *JumpcloudGoogleEmmDeviceStateInfo) HasLastStatusReportTime() bool`

HasLastStatusReportTime returns a boolean if a field has been set.

### GetPolicyCompliant

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetPolicyCompliant() bool`

GetPolicyCompliant returns the PolicyCompliant field if non-nil, zero value otherwise.

### GetPolicyCompliantOk

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetPolicyCompliantOk() (*bool, bool)`

GetPolicyCompliantOk returns a tuple with the PolicyCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCompliant

`func (o *JumpcloudGoogleEmmDeviceStateInfo) SetPolicyCompliant(v bool)`

SetPolicyCompliant sets PolicyCompliant field to given value.

### HasPolicyCompliant

`func (o *JumpcloudGoogleEmmDeviceStateInfo) HasPolicyCompliant() bool`

HasPolicyCompliant returns a boolean if a field has been set.

### GetSecurityPosture

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetSecurityPosture() JumpcloudGoogleEmmSecurityPosture`

GetSecurityPosture returns the SecurityPosture field if non-nil, zero value otherwise.

### GetSecurityPostureOk

`func (o *JumpcloudGoogleEmmDeviceStateInfo) GetSecurityPostureOk() (*JumpcloudGoogleEmmSecurityPosture, bool)`

GetSecurityPostureOk returns a tuple with the SecurityPosture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPosture

`func (o *JumpcloudGoogleEmmDeviceStateInfo) SetSecurityPosture(v JumpcloudGoogleEmmSecurityPosture)`

SetSecurityPosture sets SecurityPosture field to given value.

### HasSecurityPosture

`func (o *JumpcloudGoogleEmmDeviceStateInfo) HasSecurityPosture() bool`

HasSecurityPosture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



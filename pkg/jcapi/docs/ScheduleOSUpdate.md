# ScheduleOSUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallAction** | [**InstallActionType**](InstallActionType.md) |  | 
**MaxUserDeferrals** | Pointer to **int32** |  | [optional] 
**ProductKey** | **string** |  | 

## Methods

### NewScheduleOSUpdate

`func NewScheduleOSUpdate(installAction InstallActionType, productKey string, ) *ScheduleOSUpdate`

NewScheduleOSUpdate instantiates a new ScheduleOSUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleOSUpdateWithDefaults

`func NewScheduleOSUpdateWithDefaults() *ScheduleOSUpdate`

NewScheduleOSUpdateWithDefaults instantiates a new ScheduleOSUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallAction

`func (o *ScheduleOSUpdate) GetInstallAction() InstallActionType`

GetInstallAction returns the InstallAction field if non-nil, zero value otherwise.

### GetInstallActionOk

`func (o *ScheduleOSUpdate) GetInstallActionOk() (*InstallActionType, bool)`

GetInstallActionOk returns a tuple with the InstallAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAction

`func (o *ScheduleOSUpdate) SetInstallAction(v InstallActionType)`

SetInstallAction sets InstallAction field to given value.


### GetMaxUserDeferrals

`func (o *ScheduleOSUpdate) GetMaxUserDeferrals() int32`

GetMaxUserDeferrals returns the MaxUserDeferrals field if non-nil, zero value otherwise.

### GetMaxUserDeferralsOk

`func (o *ScheduleOSUpdate) GetMaxUserDeferralsOk() (*int32, bool)`

GetMaxUserDeferralsOk returns a tuple with the MaxUserDeferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUserDeferrals

`func (o *ScheduleOSUpdate) SetMaxUserDeferrals(v int32)`

SetMaxUserDeferrals sets MaxUserDeferrals field to given value.

### HasMaxUserDeferrals

`func (o *ScheduleOSUpdate) HasMaxUserDeferrals() bool`

HasMaxUserDeferrals returns a boolean if a field has been set.

### GetProductKey

`func (o *ScheduleOSUpdate) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *ScheduleOSUpdate) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *ScheduleOSUpdate) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ADE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDeviceGroupObjectIds** | Pointer to **[]string** | An array of ObjectIDs identifying the default device groups for this specific type (based on the OS family) of automated device enrollment. Currently, only a single DeviceGroupID is supported. | [optional] 
**EnableZeroTouchEnrollment** | Pointer to **bool** | A toggle to determine if ADE registered devices should go through JumpCloud Zero Touch Enrollment. | [optional] 
**SetupAssistantOptions** | Pointer to [**[]DEPSetupAssistantOption**](DEPSetupAssistantOption.md) |  | [optional] 
**WelcomeScreen** | Pointer to [**DEPWelcomeScreen**](DEPWelcomeScreen.md) |  | [optional] 

## Methods

### NewADE

`func NewADE() *ADE`

NewADE instantiates a new ADE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADEWithDefaults

`func NewADEWithDefaults() *ADE`

NewADEWithDefaults instantiates a new ADE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDeviceGroupObjectIds

`func (o *ADE) GetDefaultDeviceGroupObjectIds() []string`

GetDefaultDeviceGroupObjectIds returns the DefaultDeviceGroupObjectIds field if non-nil, zero value otherwise.

### GetDefaultDeviceGroupObjectIdsOk

`func (o *ADE) GetDefaultDeviceGroupObjectIdsOk() (*[]string, bool)`

GetDefaultDeviceGroupObjectIdsOk returns a tuple with the DefaultDeviceGroupObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeviceGroupObjectIds

`func (o *ADE) SetDefaultDeviceGroupObjectIds(v []string)`

SetDefaultDeviceGroupObjectIds sets DefaultDeviceGroupObjectIds field to given value.

### HasDefaultDeviceGroupObjectIds

`func (o *ADE) HasDefaultDeviceGroupObjectIds() bool`

HasDefaultDeviceGroupObjectIds returns a boolean if a field has been set.

### SetDefaultDeviceGroupObjectIdsNil

`func (o *ADE) SetDefaultDeviceGroupObjectIdsNil(b bool)`

 SetDefaultDeviceGroupObjectIdsNil sets the value for DefaultDeviceGroupObjectIds to be an explicit nil

### UnsetDefaultDeviceGroupObjectIds
`func (o *ADE) UnsetDefaultDeviceGroupObjectIds()`

UnsetDefaultDeviceGroupObjectIds ensures that no value is present for DefaultDeviceGroupObjectIds, not even an explicit nil
### GetEnableZeroTouchEnrollment

`func (o *ADE) GetEnableZeroTouchEnrollment() bool`

GetEnableZeroTouchEnrollment returns the EnableZeroTouchEnrollment field if non-nil, zero value otherwise.

### GetEnableZeroTouchEnrollmentOk

`func (o *ADE) GetEnableZeroTouchEnrollmentOk() (*bool, bool)`

GetEnableZeroTouchEnrollmentOk returns a tuple with the EnableZeroTouchEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableZeroTouchEnrollment

`func (o *ADE) SetEnableZeroTouchEnrollment(v bool)`

SetEnableZeroTouchEnrollment sets EnableZeroTouchEnrollment field to given value.

### HasEnableZeroTouchEnrollment

`func (o *ADE) HasEnableZeroTouchEnrollment() bool`

HasEnableZeroTouchEnrollment returns a boolean if a field has been set.

### GetSetupAssistantOptions

`func (o *ADE) GetSetupAssistantOptions() []DEPSetupAssistantOption`

GetSetupAssistantOptions returns the SetupAssistantOptions field if non-nil, zero value otherwise.

### GetSetupAssistantOptionsOk

`func (o *ADE) GetSetupAssistantOptionsOk() (*[]DEPSetupAssistantOption, bool)`

GetSetupAssistantOptionsOk returns a tuple with the SetupAssistantOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupAssistantOptions

`func (o *ADE) SetSetupAssistantOptions(v []DEPSetupAssistantOption)`

SetSetupAssistantOptions sets SetupAssistantOptions field to given value.

### HasSetupAssistantOptions

`func (o *ADE) HasSetupAssistantOptions() bool`

HasSetupAssistantOptions returns a boolean if a field has been set.

### GetWelcomeScreen

`func (o *ADE) GetWelcomeScreen() DEPWelcomeScreen`

GetWelcomeScreen returns the WelcomeScreen field if non-nil, zero value otherwise.

### GetWelcomeScreenOk

`func (o *ADE) GetWelcomeScreenOk() (*DEPWelcomeScreen, bool)`

GetWelcomeScreenOk returns a tuple with the WelcomeScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeScreen

`func (o *ADE) SetWelcomeScreen(v DEPWelcomeScreen)`

SetWelcomeScreen sets WelcomeScreen field to given value.

### HasWelcomeScreen

`func (o *ADE) HasWelcomeScreen() bool`

HasWelcomeScreen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



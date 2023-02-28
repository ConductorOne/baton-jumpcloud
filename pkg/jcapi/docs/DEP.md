# DEP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableZeroTouchEnrollment** | Pointer to **bool** | A toggle to determine if DEP registered devices should go through JumpCloud Zero Touch Enrollment. | [optional] 
**SetupAssistantOptions** | Pointer to [**[]DEPSetupAssistantOption**](DEPSetupAssistantOption.md) |  | [optional] 
**WelcomeScreen** | Pointer to [**DEPWelcomeScreen**](DEPWelcomeScreen.md) |  | [optional] 

## Methods

### NewDEP

`func NewDEP() *DEP`

NewDEP instantiates a new DEP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDEPWithDefaults

`func NewDEPWithDefaults() *DEP`

NewDEPWithDefaults instantiates a new DEP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableZeroTouchEnrollment

`func (o *DEP) GetEnableZeroTouchEnrollment() bool`

GetEnableZeroTouchEnrollment returns the EnableZeroTouchEnrollment field if non-nil, zero value otherwise.

### GetEnableZeroTouchEnrollmentOk

`func (o *DEP) GetEnableZeroTouchEnrollmentOk() (*bool, bool)`

GetEnableZeroTouchEnrollmentOk returns a tuple with the EnableZeroTouchEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableZeroTouchEnrollment

`func (o *DEP) SetEnableZeroTouchEnrollment(v bool)`

SetEnableZeroTouchEnrollment sets EnableZeroTouchEnrollment field to given value.

### HasEnableZeroTouchEnrollment

`func (o *DEP) HasEnableZeroTouchEnrollment() bool`

HasEnableZeroTouchEnrollment returns a boolean if a field has been set.

### GetSetupAssistantOptions

`func (o *DEP) GetSetupAssistantOptions() []DEPSetupAssistantOption`

GetSetupAssistantOptions returns the SetupAssistantOptions field if non-nil, zero value otherwise.

### GetSetupAssistantOptionsOk

`func (o *DEP) GetSetupAssistantOptionsOk() (*[]DEPSetupAssistantOption, bool)`

GetSetupAssistantOptionsOk returns a tuple with the SetupAssistantOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupAssistantOptions

`func (o *DEP) SetSetupAssistantOptions(v []DEPSetupAssistantOption)`

SetSetupAssistantOptions sets SetupAssistantOptions field to given value.

### HasSetupAssistantOptions

`func (o *DEP) HasSetupAssistantOptions() bool`

HasSetupAssistantOptions returns a boolean if a field has been set.

### GetWelcomeScreen

`func (o *DEP) GetWelcomeScreen() DEPWelcomeScreen`

GetWelcomeScreen returns the WelcomeScreen field if non-nil, zero value otherwise.

### GetWelcomeScreenOk

`func (o *DEP) GetWelcomeScreenOk() (*DEPWelcomeScreen, bool)`

GetWelcomeScreenOk returns a tuple with the WelcomeScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeScreen

`func (o *DEP) SetWelcomeScreen(v DEPWelcomeScreen)`

SetWelcomeScreen sets WelcomeScreen field to given value.

### HasWelcomeScreen

`func (o *DEP) HasWelcomeScreen() bool`

HasWelcomeScreen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



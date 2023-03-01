# OrganizationsettingsputPasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsernameSubstring** | Pointer to **bool** |  | [optional] 
**DaysAfterExpirationToSelfRecover** | Pointer to **int32** | Deprecated field used for the legacy grace period feature. | [optional] 
**DaysBeforeExpirationToForceReset** | Pointer to **int32** |  | [optional] 
**EffectiveDate** | Pointer to **string** |  | [optional] 
**EnableDaysAfterExpirationToSelfRecover** | Pointer to **bool** |  | [optional] 
**EnableDaysBeforeExpirationToForceReset** | Pointer to **bool** |  | [optional] 
**EnableLockoutTimeInSeconds** | Pointer to **bool** |  | [optional] 
**EnableMaxHistory** | Pointer to **bool** |  | [optional] 
**EnableMaxLoginAttempts** | Pointer to **bool** |  | [optional] 
**EnableMinChangePeriodInDays** | Pointer to **bool** |  | [optional] 
**EnableMinLength** | Pointer to **bool** |  | [optional] 
**EnablePasswordExpirationInDays** | Pointer to **bool** |  | [optional] 
**GracePeriodDate** | Pointer to **string** |  | [optional] 
**LockoutTimeInSeconds** | Pointer to **int32** |  | [optional] 
**MaxHistory** | Pointer to **int32** |  | [optional] 
**MaxLoginAttempts** | Pointer to **int32** |  | [optional] 
**MinChangePeriodInDays** | Pointer to **int32** |  | [optional] 
**MinLength** | Pointer to **int32** |  | [optional] 
**NeedsLowercase** | Pointer to **bool** |  | [optional] 
**NeedsNumeric** | Pointer to **bool** |  | [optional] 
**NeedsSymbolic** | Pointer to **bool** |  | [optional] 
**NeedsUppercase** | Pointer to **bool** |  | [optional] 
**PasswordExpirationInDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrganizationsettingsputPasswordPolicy

`func NewOrganizationsettingsputPasswordPolicy() *OrganizationsettingsputPasswordPolicy`

NewOrganizationsettingsputPasswordPolicy instantiates a new OrganizationsettingsputPasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsettingsputPasswordPolicyWithDefaults

`func NewOrganizationsettingsputPasswordPolicyWithDefaults() *OrganizationsettingsputPasswordPolicy`

NewOrganizationsettingsputPasswordPolicyWithDefaults instantiates a new OrganizationsettingsputPasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsernameSubstring

`func (o *OrganizationsettingsputPasswordPolicy) GetAllowUsernameSubstring() bool`

GetAllowUsernameSubstring returns the AllowUsernameSubstring field if non-nil, zero value otherwise.

### GetAllowUsernameSubstringOk

`func (o *OrganizationsettingsputPasswordPolicy) GetAllowUsernameSubstringOk() (*bool, bool)`

GetAllowUsernameSubstringOk returns a tuple with the AllowUsernameSubstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsernameSubstring

`func (o *OrganizationsettingsputPasswordPolicy) SetAllowUsernameSubstring(v bool)`

SetAllowUsernameSubstring sets AllowUsernameSubstring field to given value.

### HasAllowUsernameSubstring

`func (o *OrganizationsettingsputPasswordPolicy) HasAllowUsernameSubstring() bool`

HasAllowUsernameSubstring returns a boolean if a field has been set.

### GetDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsputPasswordPolicy) GetDaysAfterExpirationToSelfRecover() int32`

GetDaysAfterExpirationToSelfRecover returns the DaysAfterExpirationToSelfRecover field if non-nil, zero value otherwise.

### GetDaysAfterExpirationToSelfRecoverOk

`func (o *OrganizationsettingsputPasswordPolicy) GetDaysAfterExpirationToSelfRecoverOk() (*int32, bool)`

GetDaysAfterExpirationToSelfRecoverOk returns a tuple with the DaysAfterExpirationToSelfRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsputPasswordPolicy) SetDaysAfterExpirationToSelfRecover(v int32)`

SetDaysAfterExpirationToSelfRecover sets DaysAfterExpirationToSelfRecover field to given value.

### HasDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsputPasswordPolicy) HasDaysAfterExpirationToSelfRecover() bool`

HasDaysAfterExpirationToSelfRecover returns a boolean if a field has been set.

### GetDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsputPasswordPolicy) GetDaysBeforeExpirationToForceReset() int32`

GetDaysBeforeExpirationToForceReset returns the DaysBeforeExpirationToForceReset field if non-nil, zero value otherwise.

### GetDaysBeforeExpirationToForceResetOk

`func (o *OrganizationsettingsputPasswordPolicy) GetDaysBeforeExpirationToForceResetOk() (*int32, bool)`

GetDaysBeforeExpirationToForceResetOk returns a tuple with the DaysBeforeExpirationToForceReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsputPasswordPolicy) SetDaysBeforeExpirationToForceReset(v int32)`

SetDaysBeforeExpirationToForceReset sets DaysBeforeExpirationToForceReset field to given value.

### HasDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsputPasswordPolicy) HasDaysBeforeExpirationToForceReset() bool`

HasDaysBeforeExpirationToForceReset returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *OrganizationsettingsputPasswordPolicy) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *OrganizationsettingsputPasswordPolicy) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *OrganizationsettingsputPasswordPolicy) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *OrganizationsettingsputPasswordPolicy) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetEnableDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableDaysAfterExpirationToSelfRecover() bool`

GetEnableDaysAfterExpirationToSelfRecover returns the EnableDaysAfterExpirationToSelfRecover field if non-nil, zero value otherwise.

### GetEnableDaysAfterExpirationToSelfRecoverOk

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableDaysAfterExpirationToSelfRecoverOk() (*bool, bool)`

GetEnableDaysAfterExpirationToSelfRecoverOk returns a tuple with the EnableDaysAfterExpirationToSelfRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsputPasswordPolicy) SetEnableDaysAfterExpirationToSelfRecover(v bool)`

SetEnableDaysAfterExpirationToSelfRecover sets EnableDaysAfterExpirationToSelfRecover field to given value.

### HasEnableDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsputPasswordPolicy) HasEnableDaysAfterExpirationToSelfRecover() bool`

HasEnableDaysAfterExpirationToSelfRecover returns a boolean if a field has been set.

### GetEnableDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableDaysBeforeExpirationToForceReset() bool`

GetEnableDaysBeforeExpirationToForceReset returns the EnableDaysBeforeExpirationToForceReset field if non-nil, zero value otherwise.

### GetEnableDaysBeforeExpirationToForceResetOk

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableDaysBeforeExpirationToForceResetOk() (*bool, bool)`

GetEnableDaysBeforeExpirationToForceResetOk returns a tuple with the EnableDaysBeforeExpirationToForceReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsputPasswordPolicy) SetEnableDaysBeforeExpirationToForceReset(v bool)`

SetEnableDaysBeforeExpirationToForceReset sets EnableDaysBeforeExpirationToForceReset field to given value.

### HasEnableDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsputPasswordPolicy) HasEnableDaysBeforeExpirationToForceReset() bool`

HasEnableDaysBeforeExpirationToForceReset returns a boolean if a field has been set.

### GetEnableLockoutTimeInSeconds

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableLockoutTimeInSeconds() bool`

GetEnableLockoutTimeInSeconds returns the EnableLockoutTimeInSeconds field if non-nil, zero value otherwise.

### GetEnableLockoutTimeInSecondsOk

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableLockoutTimeInSecondsOk() (*bool, bool)`

GetEnableLockoutTimeInSecondsOk returns a tuple with the EnableLockoutTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLockoutTimeInSeconds

`func (o *OrganizationsettingsputPasswordPolicy) SetEnableLockoutTimeInSeconds(v bool)`

SetEnableLockoutTimeInSeconds sets EnableLockoutTimeInSeconds field to given value.

### HasEnableLockoutTimeInSeconds

`func (o *OrganizationsettingsputPasswordPolicy) HasEnableLockoutTimeInSeconds() bool`

HasEnableLockoutTimeInSeconds returns a boolean if a field has been set.

### GetEnableMaxHistory

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableMaxHistory() bool`

GetEnableMaxHistory returns the EnableMaxHistory field if non-nil, zero value otherwise.

### GetEnableMaxHistoryOk

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableMaxHistoryOk() (*bool, bool)`

GetEnableMaxHistoryOk returns a tuple with the EnableMaxHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMaxHistory

`func (o *OrganizationsettingsputPasswordPolicy) SetEnableMaxHistory(v bool)`

SetEnableMaxHistory sets EnableMaxHistory field to given value.

### HasEnableMaxHistory

`func (o *OrganizationsettingsputPasswordPolicy) HasEnableMaxHistory() bool`

HasEnableMaxHistory returns a boolean if a field has been set.

### GetEnableMaxLoginAttempts

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableMaxLoginAttempts() bool`

GetEnableMaxLoginAttempts returns the EnableMaxLoginAttempts field if non-nil, zero value otherwise.

### GetEnableMaxLoginAttemptsOk

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableMaxLoginAttemptsOk() (*bool, bool)`

GetEnableMaxLoginAttemptsOk returns a tuple with the EnableMaxLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMaxLoginAttempts

`func (o *OrganizationsettingsputPasswordPolicy) SetEnableMaxLoginAttempts(v bool)`

SetEnableMaxLoginAttempts sets EnableMaxLoginAttempts field to given value.

### HasEnableMaxLoginAttempts

`func (o *OrganizationsettingsputPasswordPolicy) HasEnableMaxLoginAttempts() bool`

HasEnableMaxLoginAttempts returns a boolean if a field has been set.

### GetEnableMinChangePeriodInDays

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableMinChangePeriodInDays() bool`

GetEnableMinChangePeriodInDays returns the EnableMinChangePeriodInDays field if non-nil, zero value otherwise.

### GetEnableMinChangePeriodInDaysOk

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableMinChangePeriodInDaysOk() (*bool, bool)`

GetEnableMinChangePeriodInDaysOk returns a tuple with the EnableMinChangePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMinChangePeriodInDays

`func (o *OrganizationsettingsputPasswordPolicy) SetEnableMinChangePeriodInDays(v bool)`

SetEnableMinChangePeriodInDays sets EnableMinChangePeriodInDays field to given value.

### HasEnableMinChangePeriodInDays

`func (o *OrganizationsettingsputPasswordPolicy) HasEnableMinChangePeriodInDays() bool`

HasEnableMinChangePeriodInDays returns a boolean if a field has been set.

### GetEnableMinLength

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableMinLength() bool`

GetEnableMinLength returns the EnableMinLength field if non-nil, zero value otherwise.

### GetEnableMinLengthOk

`func (o *OrganizationsettingsputPasswordPolicy) GetEnableMinLengthOk() (*bool, bool)`

GetEnableMinLengthOk returns a tuple with the EnableMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMinLength

`func (o *OrganizationsettingsputPasswordPolicy) SetEnableMinLength(v bool)`

SetEnableMinLength sets EnableMinLength field to given value.

### HasEnableMinLength

`func (o *OrganizationsettingsputPasswordPolicy) HasEnableMinLength() bool`

HasEnableMinLength returns a boolean if a field has been set.

### GetEnablePasswordExpirationInDays

`func (o *OrganizationsettingsputPasswordPolicy) GetEnablePasswordExpirationInDays() bool`

GetEnablePasswordExpirationInDays returns the EnablePasswordExpirationInDays field if non-nil, zero value otherwise.

### GetEnablePasswordExpirationInDaysOk

`func (o *OrganizationsettingsputPasswordPolicy) GetEnablePasswordExpirationInDaysOk() (*bool, bool)`

GetEnablePasswordExpirationInDaysOk returns a tuple with the EnablePasswordExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordExpirationInDays

`func (o *OrganizationsettingsputPasswordPolicy) SetEnablePasswordExpirationInDays(v bool)`

SetEnablePasswordExpirationInDays sets EnablePasswordExpirationInDays field to given value.

### HasEnablePasswordExpirationInDays

`func (o *OrganizationsettingsputPasswordPolicy) HasEnablePasswordExpirationInDays() bool`

HasEnablePasswordExpirationInDays returns a boolean if a field has been set.

### GetGracePeriodDate

`func (o *OrganizationsettingsputPasswordPolicy) GetGracePeriodDate() string`

GetGracePeriodDate returns the GracePeriodDate field if non-nil, zero value otherwise.

### GetGracePeriodDateOk

`func (o *OrganizationsettingsputPasswordPolicy) GetGracePeriodDateOk() (*string, bool)`

GetGracePeriodDateOk returns a tuple with the GracePeriodDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodDate

`func (o *OrganizationsettingsputPasswordPolicy) SetGracePeriodDate(v string)`

SetGracePeriodDate sets GracePeriodDate field to given value.

### HasGracePeriodDate

`func (o *OrganizationsettingsputPasswordPolicy) HasGracePeriodDate() bool`

HasGracePeriodDate returns a boolean if a field has been set.

### GetLockoutTimeInSeconds

`func (o *OrganizationsettingsputPasswordPolicy) GetLockoutTimeInSeconds() int32`

GetLockoutTimeInSeconds returns the LockoutTimeInSeconds field if non-nil, zero value otherwise.

### GetLockoutTimeInSecondsOk

`func (o *OrganizationsettingsputPasswordPolicy) GetLockoutTimeInSecondsOk() (*int32, bool)`

GetLockoutTimeInSecondsOk returns a tuple with the LockoutTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutTimeInSeconds

`func (o *OrganizationsettingsputPasswordPolicy) SetLockoutTimeInSeconds(v int32)`

SetLockoutTimeInSeconds sets LockoutTimeInSeconds field to given value.

### HasLockoutTimeInSeconds

`func (o *OrganizationsettingsputPasswordPolicy) HasLockoutTimeInSeconds() bool`

HasLockoutTimeInSeconds returns a boolean if a field has been set.

### GetMaxHistory

`func (o *OrganizationsettingsputPasswordPolicy) GetMaxHistory() int32`

GetMaxHistory returns the MaxHistory field if non-nil, zero value otherwise.

### GetMaxHistoryOk

`func (o *OrganizationsettingsputPasswordPolicy) GetMaxHistoryOk() (*int32, bool)`

GetMaxHistoryOk returns a tuple with the MaxHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHistory

`func (o *OrganizationsettingsputPasswordPolicy) SetMaxHistory(v int32)`

SetMaxHistory sets MaxHistory field to given value.

### HasMaxHistory

`func (o *OrganizationsettingsputPasswordPolicy) HasMaxHistory() bool`

HasMaxHistory returns a boolean if a field has been set.

### GetMaxLoginAttempts

`func (o *OrganizationsettingsputPasswordPolicy) GetMaxLoginAttempts() int32`

GetMaxLoginAttempts returns the MaxLoginAttempts field if non-nil, zero value otherwise.

### GetMaxLoginAttemptsOk

`func (o *OrganizationsettingsputPasswordPolicy) GetMaxLoginAttemptsOk() (*int32, bool)`

GetMaxLoginAttemptsOk returns a tuple with the MaxLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoginAttempts

`func (o *OrganizationsettingsputPasswordPolicy) SetMaxLoginAttempts(v int32)`

SetMaxLoginAttempts sets MaxLoginAttempts field to given value.

### HasMaxLoginAttempts

`func (o *OrganizationsettingsputPasswordPolicy) HasMaxLoginAttempts() bool`

HasMaxLoginAttempts returns a boolean if a field has been set.

### GetMinChangePeriodInDays

`func (o *OrganizationsettingsputPasswordPolicy) GetMinChangePeriodInDays() int32`

GetMinChangePeriodInDays returns the MinChangePeriodInDays field if non-nil, zero value otherwise.

### GetMinChangePeriodInDaysOk

`func (o *OrganizationsettingsputPasswordPolicy) GetMinChangePeriodInDaysOk() (*int32, bool)`

GetMinChangePeriodInDaysOk returns a tuple with the MinChangePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinChangePeriodInDays

`func (o *OrganizationsettingsputPasswordPolicy) SetMinChangePeriodInDays(v int32)`

SetMinChangePeriodInDays sets MinChangePeriodInDays field to given value.

### HasMinChangePeriodInDays

`func (o *OrganizationsettingsputPasswordPolicy) HasMinChangePeriodInDays() bool`

HasMinChangePeriodInDays returns a boolean if a field has been set.

### GetMinLength

`func (o *OrganizationsettingsputPasswordPolicy) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *OrganizationsettingsputPasswordPolicy) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *OrganizationsettingsputPasswordPolicy) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *OrganizationsettingsputPasswordPolicy) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetNeedsLowercase

`func (o *OrganizationsettingsputPasswordPolicy) GetNeedsLowercase() bool`

GetNeedsLowercase returns the NeedsLowercase field if non-nil, zero value otherwise.

### GetNeedsLowercaseOk

`func (o *OrganizationsettingsputPasswordPolicy) GetNeedsLowercaseOk() (*bool, bool)`

GetNeedsLowercaseOk returns a tuple with the NeedsLowercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsLowercase

`func (o *OrganizationsettingsputPasswordPolicy) SetNeedsLowercase(v bool)`

SetNeedsLowercase sets NeedsLowercase field to given value.

### HasNeedsLowercase

`func (o *OrganizationsettingsputPasswordPolicy) HasNeedsLowercase() bool`

HasNeedsLowercase returns a boolean if a field has been set.

### GetNeedsNumeric

`func (o *OrganizationsettingsputPasswordPolicy) GetNeedsNumeric() bool`

GetNeedsNumeric returns the NeedsNumeric field if non-nil, zero value otherwise.

### GetNeedsNumericOk

`func (o *OrganizationsettingsputPasswordPolicy) GetNeedsNumericOk() (*bool, bool)`

GetNeedsNumericOk returns a tuple with the NeedsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsNumeric

`func (o *OrganizationsettingsputPasswordPolicy) SetNeedsNumeric(v bool)`

SetNeedsNumeric sets NeedsNumeric field to given value.

### HasNeedsNumeric

`func (o *OrganizationsettingsputPasswordPolicy) HasNeedsNumeric() bool`

HasNeedsNumeric returns a boolean if a field has been set.

### GetNeedsSymbolic

`func (o *OrganizationsettingsputPasswordPolicy) GetNeedsSymbolic() bool`

GetNeedsSymbolic returns the NeedsSymbolic field if non-nil, zero value otherwise.

### GetNeedsSymbolicOk

`func (o *OrganizationsettingsputPasswordPolicy) GetNeedsSymbolicOk() (*bool, bool)`

GetNeedsSymbolicOk returns a tuple with the NeedsSymbolic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsSymbolic

`func (o *OrganizationsettingsputPasswordPolicy) SetNeedsSymbolic(v bool)`

SetNeedsSymbolic sets NeedsSymbolic field to given value.

### HasNeedsSymbolic

`func (o *OrganizationsettingsputPasswordPolicy) HasNeedsSymbolic() bool`

HasNeedsSymbolic returns a boolean if a field has been set.

### GetNeedsUppercase

`func (o *OrganizationsettingsputPasswordPolicy) GetNeedsUppercase() bool`

GetNeedsUppercase returns the NeedsUppercase field if non-nil, zero value otherwise.

### GetNeedsUppercaseOk

`func (o *OrganizationsettingsputPasswordPolicy) GetNeedsUppercaseOk() (*bool, bool)`

GetNeedsUppercaseOk returns a tuple with the NeedsUppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsUppercase

`func (o *OrganizationsettingsputPasswordPolicy) SetNeedsUppercase(v bool)`

SetNeedsUppercase sets NeedsUppercase field to given value.

### HasNeedsUppercase

`func (o *OrganizationsettingsputPasswordPolicy) HasNeedsUppercase() bool`

HasNeedsUppercase returns a boolean if a field has been set.

### GetPasswordExpirationInDays

`func (o *OrganizationsettingsputPasswordPolicy) GetPasswordExpirationInDays() int32`

GetPasswordExpirationInDays returns the PasswordExpirationInDays field if non-nil, zero value otherwise.

### GetPasswordExpirationInDaysOk

`func (o *OrganizationsettingsputPasswordPolicy) GetPasswordExpirationInDaysOk() (*int32, bool)`

GetPasswordExpirationInDaysOk returns a tuple with the PasswordExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationInDays

`func (o *OrganizationsettingsputPasswordPolicy) SetPasswordExpirationInDays(v int32)`

SetPasswordExpirationInDays sets PasswordExpirationInDays field to given value.

### HasPasswordExpirationInDays

`func (o *OrganizationsettingsputPasswordPolicy) HasPasswordExpirationInDays() bool`

HasPasswordExpirationInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



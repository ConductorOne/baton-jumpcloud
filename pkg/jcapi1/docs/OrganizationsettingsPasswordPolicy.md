# OrganizationsettingsPasswordPolicy

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
**EnableRecoveryEmail** | Pointer to **bool** |  | [optional] 
**EnableResetLockoutCounter** | Pointer to **bool** |  | [optional] 
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
**ResetLockoutCounterMinutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrganizationsettingsPasswordPolicy

`func NewOrganizationsettingsPasswordPolicy() *OrganizationsettingsPasswordPolicy`

NewOrganizationsettingsPasswordPolicy instantiates a new OrganizationsettingsPasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsettingsPasswordPolicyWithDefaults

`func NewOrganizationsettingsPasswordPolicyWithDefaults() *OrganizationsettingsPasswordPolicy`

NewOrganizationsettingsPasswordPolicyWithDefaults instantiates a new OrganizationsettingsPasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsernameSubstring

`func (o *OrganizationsettingsPasswordPolicy) GetAllowUsernameSubstring() bool`

GetAllowUsernameSubstring returns the AllowUsernameSubstring field if non-nil, zero value otherwise.

### GetAllowUsernameSubstringOk

`func (o *OrganizationsettingsPasswordPolicy) GetAllowUsernameSubstringOk() (*bool, bool)`

GetAllowUsernameSubstringOk returns a tuple with the AllowUsernameSubstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsernameSubstring

`func (o *OrganizationsettingsPasswordPolicy) SetAllowUsernameSubstring(v bool)`

SetAllowUsernameSubstring sets AllowUsernameSubstring field to given value.

### HasAllowUsernameSubstring

`func (o *OrganizationsettingsPasswordPolicy) HasAllowUsernameSubstring() bool`

HasAllowUsernameSubstring returns a boolean if a field has been set.

### GetDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsPasswordPolicy) GetDaysAfterExpirationToSelfRecover() int32`

GetDaysAfterExpirationToSelfRecover returns the DaysAfterExpirationToSelfRecover field if non-nil, zero value otherwise.

### GetDaysAfterExpirationToSelfRecoverOk

`func (o *OrganizationsettingsPasswordPolicy) GetDaysAfterExpirationToSelfRecoverOk() (*int32, bool)`

GetDaysAfterExpirationToSelfRecoverOk returns a tuple with the DaysAfterExpirationToSelfRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsPasswordPolicy) SetDaysAfterExpirationToSelfRecover(v int32)`

SetDaysAfterExpirationToSelfRecover sets DaysAfterExpirationToSelfRecover field to given value.

### HasDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsPasswordPolicy) HasDaysAfterExpirationToSelfRecover() bool`

HasDaysAfterExpirationToSelfRecover returns a boolean if a field has been set.

### GetDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsPasswordPolicy) GetDaysBeforeExpirationToForceReset() int32`

GetDaysBeforeExpirationToForceReset returns the DaysBeforeExpirationToForceReset field if non-nil, zero value otherwise.

### GetDaysBeforeExpirationToForceResetOk

`func (o *OrganizationsettingsPasswordPolicy) GetDaysBeforeExpirationToForceResetOk() (*int32, bool)`

GetDaysBeforeExpirationToForceResetOk returns a tuple with the DaysBeforeExpirationToForceReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsPasswordPolicy) SetDaysBeforeExpirationToForceReset(v int32)`

SetDaysBeforeExpirationToForceReset sets DaysBeforeExpirationToForceReset field to given value.

### HasDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsPasswordPolicy) HasDaysBeforeExpirationToForceReset() bool`

HasDaysBeforeExpirationToForceReset returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *OrganizationsettingsPasswordPolicy) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *OrganizationsettingsPasswordPolicy) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *OrganizationsettingsPasswordPolicy) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *OrganizationsettingsPasswordPolicy) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetEnableDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsPasswordPolicy) GetEnableDaysAfterExpirationToSelfRecover() bool`

GetEnableDaysAfterExpirationToSelfRecover returns the EnableDaysAfterExpirationToSelfRecover field if non-nil, zero value otherwise.

### GetEnableDaysAfterExpirationToSelfRecoverOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnableDaysAfterExpirationToSelfRecoverOk() (*bool, bool)`

GetEnableDaysAfterExpirationToSelfRecoverOk returns a tuple with the EnableDaysAfterExpirationToSelfRecover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsPasswordPolicy) SetEnableDaysAfterExpirationToSelfRecover(v bool)`

SetEnableDaysAfterExpirationToSelfRecover sets EnableDaysAfterExpirationToSelfRecover field to given value.

### HasEnableDaysAfterExpirationToSelfRecover

`func (o *OrganizationsettingsPasswordPolicy) HasEnableDaysAfterExpirationToSelfRecover() bool`

HasEnableDaysAfterExpirationToSelfRecover returns a boolean if a field has been set.

### GetEnableDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsPasswordPolicy) GetEnableDaysBeforeExpirationToForceReset() bool`

GetEnableDaysBeforeExpirationToForceReset returns the EnableDaysBeforeExpirationToForceReset field if non-nil, zero value otherwise.

### GetEnableDaysBeforeExpirationToForceResetOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnableDaysBeforeExpirationToForceResetOk() (*bool, bool)`

GetEnableDaysBeforeExpirationToForceResetOk returns a tuple with the EnableDaysBeforeExpirationToForceReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsPasswordPolicy) SetEnableDaysBeforeExpirationToForceReset(v bool)`

SetEnableDaysBeforeExpirationToForceReset sets EnableDaysBeforeExpirationToForceReset field to given value.

### HasEnableDaysBeforeExpirationToForceReset

`func (o *OrganizationsettingsPasswordPolicy) HasEnableDaysBeforeExpirationToForceReset() bool`

HasEnableDaysBeforeExpirationToForceReset returns a boolean if a field has been set.

### GetEnableLockoutTimeInSeconds

`func (o *OrganizationsettingsPasswordPolicy) GetEnableLockoutTimeInSeconds() bool`

GetEnableLockoutTimeInSeconds returns the EnableLockoutTimeInSeconds field if non-nil, zero value otherwise.

### GetEnableLockoutTimeInSecondsOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnableLockoutTimeInSecondsOk() (*bool, bool)`

GetEnableLockoutTimeInSecondsOk returns a tuple with the EnableLockoutTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLockoutTimeInSeconds

`func (o *OrganizationsettingsPasswordPolicy) SetEnableLockoutTimeInSeconds(v bool)`

SetEnableLockoutTimeInSeconds sets EnableLockoutTimeInSeconds field to given value.

### HasEnableLockoutTimeInSeconds

`func (o *OrganizationsettingsPasswordPolicy) HasEnableLockoutTimeInSeconds() bool`

HasEnableLockoutTimeInSeconds returns a boolean if a field has been set.

### GetEnableMaxHistory

`func (o *OrganizationsettingsPasswordPolicy) GetEnableMaxHistory() bool`

GetEnableMaxHistory returns the EnableMaxHistory field if non-nil, zero value otherwise.

### GetEnableMaxHistoryOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnableMaxHistoryOk() (*bool, bool)`

GetEnableMaxHistoryOk returns a tuple with the EnableMaxHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMaxHistory

`func (o *OrganizationsettingsPasswordPolicy) SetEnableMaxHistory(v bool)`

SetEnableMaxHistory sets EnableMaxHistory field to given value.

### HasEnableMaxHistory

`func (o *OrganizationsettingsPasswordPolicy) HasEnableMaxHistory() bool`

HasEnableMaxHistory returns a boolean if a field has been set.

### GetEnableMaxLoginAttempts

`func (o *OrganizationsettingsPasswordPolicy) GetEnableMaxLoginAttempts() bool`

GetEnableMaxLoginAttempts returns the EnableMaxLoginAttempts field if non-nil, zero value otherwise.

### GetEnableMaxLoginAttemptsOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnableMaxLoginAttemptsOk() (*bool, bool)`

GetEnableMaxLoginAttemptsOk returns a tuple with the EnableMaxLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMaxLoginAttempts

`func (o *OrganizationsettingsPasswordPolicy) SetEnableMaxLoginAttempts(v bool)`

SetEnableMaxLoginAttempts sets EnableMaxLoginAttempts field to given value.

### HasEnableMaxLoginAttempts

`func (o *OrganizationsettingsPasswordPolicy) HasEnableMaxLoginAttempts() bool`

HasEnableMaxLoginAttempts returns a boolean if a field has been set.

### GetEnableMinChangePeriodInDays

`func (o *OrganizationsettingsPasswordPolicy) GetEnableMinChangePeriodInDays() bool`

GetEnableMinChangePeriodInDays returns the EnableMinChangePeriodInDays field if non-nil, zero value otherwise.

### GetEnableMinChangePeriodInDaysOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnableMinChangePeriodInDaysOk() (*bool, bool)`

GetEnableMinChangePeriodInDaysOk returns a tuple with the EnableMinChangePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMinChangePeriodInDays

`func (o *OrganizationsettingsPasswordPolicy) SetEnableMinChangePeriodInDays(v bool)`

SetEnableMinChangePeriodInDays sets EnableMinChangePeriodInDays field to given value.

### HasEnableMinChangePeriodInDays

`func (o *OrganizationsettingsPasswordPolicy) HasEnableMinChangePeriodInDays() bool`

HasEnableMinChangePeriodInDays returns a boolean if a field has been set.

### GetEnableMinLength

`func (o *OrganizationsettingsPasswordPolicy) GetEnableMinLength() bool`

GetEnableMinLength returns the EnableMinLength field if non-nil, zero value otherwise.

### GetEnableMinLengthOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnableMinLengthOk() (*bool, bool)`

GetEnableMinLengthOk returns a tuple with the EnableMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMinLength

`func (o *OrganizationsettingsPasswordPolicy) SetEnableMinLength(v bool)`

SetEnableMinLength sets EnableMinLength field to given value.

### HasEnableMinLength

`func (o *OrganizationsettingsPasswordPolicy) HasEnableMinLength() bool`

HasEnableMinLength returns a boolean if a field has been set.

### GetEnablePasswordExpirationInDays

`func (o *OrganizationsettingsPasswordPolicy) GetEnablePasswordExpirationInDays() bool`

GetEnablePasswordExpirationInDays returns the EnablePasswordExpirationInDays field if non-nil, zero value otherwise.

### GetEnablePasswordExpirationInDaysOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnablePasswordExpirationInDaysOk() (*bool, bool)`

GetEnablePasswordExpirationInDaysOk returns a tuple with the EnablePasswordExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePasswordExpirationInDays

`func (o *OrganizationsettingsPasswordPolicy) SetEnablePasswordExpirationInDays(v bool)`

SetEnablePasswordExpirationInDays sets EnablePasswordExpirationInDays field to given value.

### HasEnablePasswordExpirationInDays

`func (o *OrganizationsettingsPasswordPolicy) HasEnablePasswordExpirationInDays() bool`

HasEnablePasswordExpirationInDays returns a boolean if a field has been set.

### GetEnableRecoveryEmail

`func (o *OrganizationsettingsPasswordPolicy) GetEnableRecoveryEmail() bool`

GetEnableRecoveryEmail returns the EnableRecoveryEmail field if non-nil, zero value otherwise.

### GetEnableRecoveryEmailOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnableRecoveryEmailOk() (*bool, bool)`

GetEnableRecoveryEmailOk returns a tuple with the EnableRecoveryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecoveryEmail

`func (o *OrganizationsettingsPasswordPolicy) SetEnableRecoveryEmail(v bool)`

SetEnableRecoveryEmail sets EnableRecoveryEmail field to given value.

### HasEnableRecoveryEmail

`func (o *OrganizationsettingsPasswordPolicy) HasEnableRecoveryEmail() bool`

HasEnableRecoveryEmail returns a boolean if a field has been set.

### GetEnableResetLockoutCounter

`func (o *OrganizationsettingsPasswordPolicy) GetEnableResetLockoutCounter() bool`

GetEnableResetLockoutCounter returns the EnableResetLockoutCounter field if non-nil, zero value otherwise.

### GetEnableResetLockoutCounterOk

`func (o *OrganizationsettingsPasswordPolicy) GetEnableResetLockoutCounterOk() (*bool, bool)`

GetEnableResetLockoutCounterOk returns a tuple with the EnableResetLockoutCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableResetLockoutCounter

`func (o *OrganizationsettingsPasswordPolicy) SetEnableResetLockoutCounter(v bool)`

SetEnableResetLockoutCounter sets EnableResetLockoutCounter field to given value.

### HasEnableResetLockoutCounter

`func (o *OrganizationsettingsPasswordPolicy) HasEnableResetLockoutCounter() bool`

HasEnableResetLockoutCounter returns a boolean if a field has been set.

### GetGracePeriodDate

`func (o *OrganizationsettingsPasswordPolicy) GetGracePeriodDate() string`

GetGracePeriodDate returns the GracePeriodDate field if non-nil, zero value otherwise.

### GetGracePeriodDateOk

`func (o *OrganizationsettingsPasswordPolicy) GetGracePeriodDateOk() (*string, bool)`

GetGracePeriodDateOk returns a tuple with the GracePeriodDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodDate

`func (o *OrganizationsettingsPasswordPolicy) SetGracePeriodDate(v string)`

SetGracePeriodDate sets GracePeriodDate field to given value.

### HasGracePeriodDate

`func (o *OrganizationsettingsPasswordPolicy) HasGracePeriodDate() bool`

HasGracePeriodDate returns a boolean if a field has been set.

### GetLockoutTimeInSeconds

`func (o *OrganizationsettingsPasswordPolicy) GetLockoutTimeInSeconds() int32`

GetLockoutTimeInSeconds returns the LockoutTimeInSeconds field if non-nil, zero value otherwise.

### GetLockoutTimeInSecondsOk

`func (o *OrganizationsettingsPasswordPolicy) GetLockoutTimeInSecondsOk() (*int32, bool)`

GetLockoutTimeInSecondsOk returns a tuple with the LockoutTimeInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutTimeInSeconds

`func (o *OrganizationsettingsPasswordPolicy) SetLockoutTimeInSeconds(v int32)`

SetLockoutTimeInSeconds sets LockoutTimeInSeconds field to given value.

### HasLockoutTimeInSeconds

`func (o *OrganizationsettingsPasswordPolicy) HasLockoutTimeInSeconds() bool`

HasLockoutTimeInSeconds returns a boolean if a field has been set.

### GetMaxHistory

`func (o *OrganizationsettingsPasswordPolicy) GetMaxHistory() int32`

GetMaxHistory returns the MaxHistory field if non-nil, zero value otherwise.

### GetMaxHistoryOk

`func (o *OrganizationsettingsPasswordPolicy) GetMaxHistoryOk() (*int32, bool)`

GetMaxHistoryOk returns a tuple with the MaxHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHistory

`func (o *OrganizationsettingsPasswordPolicy) SetMaxHistory(v int32)`

SetMaxHistory sets MaxHistory field to given value.

### HasMaxHistory

`func (o *OrganizationsettingsPasswordPolicy) HasMaxHistory() bool`

HasMaxHistory returns a boolean if a field has been set.

### GetMaxLoginAttempts

`func (o *OrganizationsettingsPasswordPolicy) GetMaxLoginAttempts() int32`

GetMaxLoginAttempts returns the MaxLoginAttempts field if non-nil, zero value otherwise.

### GetMaxLoginAttemptsOk

`func (o *OrganizationsettingsPasswordPolicy) GetMaxLoginAttemptsOk() (*int32, bool)`

GetMaxLoginAttemptsOk returns a tuple with the MaxLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoginAttempts

`func (o *OrganizationsettingsPasswordPolicy) SetMaxLoginAttempts(v int32)`

SetMaxLoginAttempts sets MaxLoginAttempts field to given value.

### HasMaxLoginAttempts

`func (o *OrganizationsettingsPasswordPolicy) HasMaxLoginAttempts() bool`

HasMaxLoginAttempts returns a boolean if a field has been set.

### GetMinChangePeriodInDays

`func (o *OrganizationsettingsPasswordPolicy) GetMinChangePeriodInDays() int32`

GetMinChangePeriodInDays returns the MinChangePeriodInDays field if non-nil, zero value otherwise.

### GetMinChangePeriodInDaysOk

`func (o *OrganizationsettingsPasswordPolicy) GetMinChangePeriodInDaysOk() (*int32, bool)`

GetMinChangePeriodInDaysOk returns a tuple with the MinChangePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinChangePeriodInDays

`func (o *OrganizationsettingsPasswordPolicy) SetMinChangePeriodInDays(v int32)`

SetMinChangePeriodInDays sets MinChangePeriodInDays field to given value.

### HasMinChangePeriodInDays

`func (o *OrganizationsettingsPasswordPolicy) HasMinChangePeriodInDays() bool`

HasMinChangePeriodInDays returns a boolean if a field has been set.

### GetMinLength

`func (o *OrganizationsettingsPasswordPolicy) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *OrganizationsettingsPasswordPolicy) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *OrganizationsettingsPasswordPolicy) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *OrganizationsettingsPasswordPolicy) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetNeedsLowercase

`func (o *OrganizationsettingsPasswordPolicy) GetNeedsLowercase() bool`

GetNeedsLowercase returns the NeedsLowercase field if non-nil, zero value otherwise.

### GetNeedsLowercaseOk

`func (o *OrganizationsettingsPasswordPolicy) GetNeedsLowercaseOk() (*bool, bool)`

GetNeedsLowercaseOk returns a tuple with the NeedsLowercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsLowercase

`func (o *OrganizationsettingsPasswordPolicy) SetNeedsLowercase(v bool)`

SetNeedsLowercase sets NeedsLowercase field to given value.

### HasNeedsLowercase

`func (o *OrganizationsettingsPasswordPolicy) HasNeedsLowercase() bool`

HasNeedsLowercase returns a boolean if a field has been set.

### GetNeedsNumeric

`func (o *OrganizationsettingsPasswordPolicy) GetNeedsNumeric() bool`

GetNeedsNumeric returns the NeedsNumeric field if non-nil, zero value otherwise.

### GetNeedsNumericOk

`func (o *OrganizationsettingsPasswordPolicy) GetNeedsNumericOk() (*bool, bool)`

GetNeedsNumericOk returns a tuple with the NeedsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsNumeric

`func (o *OrganizationsettingsPasswordPolicy) SetNeedsNumeric(v bool)`

SetNeedsNumeric sets NeedsNumeric field to given value.

### HasNeedsNumeric

`func (o *OrganizationsettingsPasswordPolicy) HasNeedsNumeric() bool`

HasNeedsNumeric returns a boolean if a field has been set.

### GetNeedsSymbolic

`func (o *OrganizationsettingsPasswordPolicy) GetNeedsSymbolic() bool`

GetNeedsSymbolic returns the NeedsSymbolic field if non-nil, zero value otherwise.

### GetNeedsSymbolicOk

`func (o *OrganizationsettingsPasswordPolicy) GetNeedsSymbolicOk() (*bool, bool)`

GetNeedsSymbolicOk returns a tuple with the NeedsSymbolic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsSymbolic

`func (o *OrganizationsettingsPasswordPolicy) SetNeedsSymbolic(v bool)`

SetNeedsSymbolic sets NeedsSymbolic field to given value.

### HasNeedsSymbolic

`func (o *OrganizationsettingsPasswordPolicy) HasNeedsSymbolic() bool`

HasNeedsSymbolic returns a boolean if a field has been set.

### GetNeedsUppercase

`func (o *OrganizationsettingsPasswordPolicy) GetNeedsUppercase() bool`

GetNeedsUppercase returns the NeedsUppercase field if non-nil, zero value otherwise.

### GetNeedsUppercaseOk

`func (o *OrganizationsettingsPasswordPolicy) GetNeedsUppercaseOk() (*bool, bool)`

GetNeedsUppercaseOk returns a tuple with the NeedsUppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsUppercase

`func (o *OrganizationsettingsPasswordPolicy) SetNeedsUppercase(v bool)`

SetNeedsUppercase sets NeedsUppercase field to given value.

### HasNeedsUppercase

`func (o *OrganizationsettingsPasswordPolicy) HasNeedsUppercase() bool`

HasNeedsUppercase returns a boolean if a field has been set.

### GetPasswordExpirationInDays

`func (o *OrganizationsettingsPasswordPolicy) GetPasswordExpirationInDays() int32`

GetPasswordExpirationInDays returns the PasswordExpirationInDays field if non-nil, zero value otherwise.

### GetPasswordExpirationInDaysOk

`func (o *OrganizationsettingsPasswordPolicy) GetPasswordExpirationInDaysOk() (*int32, bool)`

GetPasswordExpirationInDaysOk returns a tuple with the PasswordExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationInDays

`func (o *OrganizationsettingsPasswordPolicy) SetPasswordExpirationInDays(v int32)`

SetPasswordExpirationInDays sets PasswordExpirationInDays field to given value.

### HasPasswordExpirationInDays

`func (o *OrganizationsettingsPasswordPolicy) HasPasswordExpirationInDays() bool`

HasPasswordExpirationInDays returns a boolean if a field has been set.

### GetResetLockoutCounterMinutes

`func (o *OrganizationsettingsPasswordPolicy) GetResetLockoutCounterMinutes() int32`

GetResetLockoutCounterMinutes returns the ResetLockoutCounterMinutes field if non-nil, zero value otherwise.

### GetResetLockoutCounterMinutesOk

`func (o *OrganizationsettingsPasswordPolicy) GetResetLockoutCounterMinutesOk() (*int32, bool)`

GetResetLockoutCounterMinutesOk returns a tuple with the ResetLockoutCounterMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetLockoutCounterMinutes

`func (o *OrganizationsettingsPasswordPolicy) SetResetLockoutCounterMinutes(v int32)`

SetResetLockoutCounterMinutes sets ResetLockoutCounterMinutes field to given value.

### HasResetLockoutCounterMinutes

`func (o *OrganizationsettingsPasswordPolicy) HasResetLockoutCounterMinutes() bool`

HasResetLockoutCounterMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



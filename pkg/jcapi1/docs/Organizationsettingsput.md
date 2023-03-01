# Organizationsettingsput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to **string** |  | [optional] 
**ContactName** | Pointer to **string** |  | [optional] 
**DeviceIdentificationEnabled** | Pointer to **bool** |  | [optional] 
**DisableGoogleLogin** | Pointer to **bool** |  | [optional] 
**DisableLdap** | Pointer to **bool** |  | [optional] 
**DisableUM** | Pointer to **bool** |  | [optional] 
**DuplicateLDAPGroups** | Pointer to **bool** |  | [optional] 
**EmailDisclaimer** | Pointer to **string** |  | [optional] 
**EnableManagedUID** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**OrganizationsettingsFeatures**](OrganizationsettingsFeatures.md) |  | [optional] 
**GrowthData** | Pointer to **map[string]interface{}** | Object containing Optimizely experimentIds and states corresponding to them | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NewSystemUserStateDefaults** | Pointer to [**OrganizationsettingsputNewSystemUserStateDefaults**](OrganizationsettingsputNewSystemUserStateDefaults.md) |  | [optional] 
**PasswordCompliance** | Pointer to **string** |  | [optional] 
**PasswordPolicy** | Pointer to [**OrganizationsettingsputPasswordPolicy**](OrganizationsettingsputPasswordPolicy.md) |  | [optional] 
**ShowIntro** | Pointer to **bool** |  | [optional] 
**SystemUserPasswordExpirationInDays** | Pointer to **int32** |  | [optional] 
**SystemUsersCanEdit** | Pointer to **bool** |  | [optional] 
**SystemUsersCap** | Pointer to **int32** |  | [optional] 
**TrustedAppConfig** | Pointer to [**TrustedappConfigPut**](TrustedappConfigPut.md) |  | [optional] 
**UserPortal** | Pointer to [**OrganizationsettingsUserPortal**](OrganizationsettingsUserPortal.md) |  | [optional] 

## Methods

### NewOrganizationsettingsput

`func NewOrganizationsettingsput() *Organizationsettingsput`

NewOrganizationsettingsput instantiates a new Organizationsettingsput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsettingsputWithDefaults

`func NewOrganizationsettingsputWithDefaults() *Organizationsettingsput`

NewOrganizationsettingsputWithDefaults instantiates a new Organizationsettingsput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *Organizationsettingsput) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Organizationsettingsput) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Organizationsettingsput) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Organizationsettingsput) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetContactName

`func (o *Organizationsettingsput) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *Organizationsettingsput) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *Organizationsettingsput) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *Organizationsettingsput) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetDeviceIdentificationEnabled

`func (o *Organizationsettingsput) GetDeviceIdentificationEnabled() bool`

GetDeviceIdentificationEnabled returns the DeviceIdentificationEnabled field if non-nil, zero value otherwise.

### GetDeviceIdentificationEnabledOk

`func (o *Organizationsettingsput) GetDeviceIdentificationEnabledOk() (*bool, bool)`

GetDeviceIdentificationEnabledOk returns a tuple with the DeviceIdentificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentificationEnabled

`func (o *Organizationsettingsput) SetDeviceIdentificationEnabled(v bool)`

SetDeviceIdentificationEnabled sets DeviceIdentificationEnabled field to given value.

### HasDeviceIdentificationEnabled

`func (o *Organizationsettingsput) HasDeviceIdentificationEnabled() bool`

HasDeviceIdentificationEnabled returns a boolean if a field has been set.

### GetDisableGoogleLogin

`func (o *Organizationsettingsput) GetDisableGoogleLogin() bool`

GetDisableGoogleLogin returns the DisableGoogleLogin field if non-nil, zero value otherwise.

### GetDisableGoogleLoginOk

`func (o *Organizationsettingsput) GetDisableGoogleLoginOk() (*bool, bool)`

GetDisableGoogleLoginOk returns a tuple with the DisableGoogleLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableGoogleLogin

`func (o *Organizationsettingsput) SetDisableGoogleLogin(v bool)`

SetDisableGoogleLogin sets DisableGoogleLogin field to given value.

### HasDisableGoogleLogin

`func (o *Organizationsettingsput) HasDisableGoogleLogin() bool`

HasDisableGoogleLogin returns a boolean if a field has been set.

### GetDisableLdap

`func (o *Organizationsettingsput) GetDisableLdap() bool`

GetDisableLdap returns the DisableLdap field if non-nil, zero value otherwise.

### GetDisableLdapOk

`func (o *Organizationsettingsput) GetDisableLdapOk() (*bool, bool)`

GetDisableLdapOk returns a tuple with the DisableLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLdap

`func (o *Organizationsettingsput) SetDisableLdap(v bool)`

SetDisableLdap sets DisableLdap field to given value.

### HasDisableLdap

`func (o *Organizationsettingsput) HasDisableLdap() bool`

HasDisableLdap returns a boolean if a field has been set.

### GetDisableUM

`func (o *Organizationsettingsput) GetDisableUM() bool`

GetDisableUM returns the DisableUM field if non-nil, zero value otherwise.

### GetDisableUMOk

`func (o *Organizationsettingsput) GetDisableUMOk() (*bool, bool)`

GetDisableUMOk returns a tuple with the DisableUM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUM

`func (o *Organizationsettingsput) SetDisableUM(v bool)`

SetDisableUM sets DisableUM field to given value.

### HasDisableUM

`func (o *Organizationsettingsput) HasDisableUM() bool`

HasDisableUM returns a boolean if a field has been set.

### GetDuplicateLDAPGroups

`func (o *Organizationsettingsput) GetDuplicateLDAPGroups() bool`

GetDuplicateLDAPGroups returns the DuplicateLDAPGroups field if non-nil, zero value otherwise.

### GetDuplicateLDAPGroupsOk

`func (o *Organizationsettingsput) GetDuplicateLDAPGroupsOk() (*bool, bool)`

GetDuplicateLDAPGroupsOk returns a tuple with the DuplicateLDAPGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateLDAPGroups

`func (o *Organizationsettingsput) SetDuplicateLDAPGroups(v bool)`

SetDuplicateLDAPGroups sets DuplicateLDAPGroups field to given value.

### HasDuplicateLDAPGroups

`func (o *Organizationsettingsput) HasDuplicateLDAPGroups() bool`

HasDuplicateLDAPGroups returns a boolean if a field has been set.

### GetEmailDisclaimer

`func (o *Organizationsettingsput) GetEmailDisclaimer() string`

GetEmailDisclaimer returns the EmailDisclaimer field if non-nil, zero value otherwise.

### GetEmailDisclaimerOk

`func (o *Organizationsettingsput) GetEmailDisclaimerOk() (*string, bool)`

GetEmailDisclaimerOk returns a tuple with the EmailDisclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDisclaimer

`func (o *Organizationsettingsput) SetEmailDisclaimer(v string)`

SetEmailDisclaimer sets EmailDisclaimer field to given value.

### HasEmailDisclaimer

`func (o *Organizationsettingsput) HasEmailDisclaimer() bool`

HasEmailDisclaimer returns a boolean if a field has been set.

### GetEnableManagedUID

`func (o *Organizationsettingsput) GetEnableManagedUID() bool`

GetEnableManagedUID returns the EnableManagedUID field if non-nil, zero value otherwise.

### GetEnableManagedUIDOk

`func (o *Organizationsettingsput) GetEnableManagedUIDOk() (*bool, bool)`

GetEnableManagedUIDOk returns a tuple with the EnableManagedUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableManagedUID

`func (o *Organizationsettingsput) SetEnableManagedUID(v bool)`

SetEnableManagedUID sets EnableManagedUID field to given value.

### HasEnableManagedUID

`func (o *Organizationsettingsput) HasEnableManagedUID() bool`

HasEnableManagedUID returns a boolean if a field has been set.

### GetFeatures

`func (o *Organizationsettingsput) GetFeatures() OrganizationsettingsFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Organizationsettingsput) GetFeaturesOk() (*OrganizationsettingsFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Organizationsettingsput) SetFeatures(v OrganizationsettingsFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Organizationsettingsput) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetGrowthData

`func (o *Organizationsettingsput) GetGrowthData() map[string]interface{}`

GetGrowthData returns the GrowthData field if non-nil, zero value otherwise.

### GetGrowthDataOk

`func (o *Organizationsettingsput) GetGrowthDataOk() (*map[string]interface{}, bool)`

GetGrowthDataOk returns a tuple with the GrowthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthData

`func (o *Organizationsettingsput) SetGrowthData(v map[string]interface{})`

SetGrowthData sets GrowthData field to given value.

### HasGrowthData

`func (o *Organizationsettingsput) HasGrowthData() bool`

HasGrowthData returns a boolean if a field has been set.

### GetLogo

`func (o *Organizationsettingsput) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Organizationsettingsput) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Organizationsettingsput) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Organizationsettingsput) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *Organizationsettingsput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organizationsettingsput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organizationsettingsput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organizationsettingsput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewSystemUserStateDefaults

`func (o *Organizationsettingsput) GetNewSystemUserStateDefaults() OrganizationsettingsputNewSystemUserStateDefaults`

GetNewSystemUserStateDefaults returns the NewSystemUserStateDefaults field if non-nil, zero value otherwise.

### GetNewSystemUserStateDefaultsOk

`func (o *Organizationsettingsput) GetNewSystemUserStateDefaultsOk() (*OrganizationsettingsputNewSystemUserStateDefaults, bool)`

GetNewSystemUserStateDefaultsOk returns a tuple with the NewSystemUserStateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSystemUserStateDefaults

`func (o *Organizationsettingsput) SetNewSystemUserStateDefaults(v OrganizationsettingsputNewSystemUserStateDefaults)`

SetNewSystemUserStateDefaults sets NewSystemUserStateDefaults field to given value.

### HasNewSystemUserStateDefaults

`func (o *Organizationsettingsput) HasNewSystemUserStateDefaults() bool`

HasNewSystemUserStateDefaults returns a boolean if a field has been set.

### GetPasswordCompliance

`func (o *Organizationsettingsput) GetPasswordCompliance() string`

GetPasswordCompliance returns the PasswordCompliance field if non-nil, zero value otherwise.

### GetPasswordComplianceOk

`func (o *Organizationsettingsput) GetPasswordComplianceOk() (*string, bool)`

GetPasswordComplianceOk returns a tuple with the PasswordCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCompliance

`func (o *Organizationsettingsput) SetPasswordCompliance(v string)`

SetPasswordCompliance sets PasswordCompliance field to given value.

### HasPasswordCompliance

`func (o *Organizationsettingsput) HasPasswordCompliance() bool`

HasPasswordCompliance returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *Organizationsettingsput) GetPasswordPolicy() OrganizationsettingsputPasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *Organizationsettingsput) GetPasswordPolicyOk() (*OrganizationsettingsputPasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *Organizationsettingsput) SetPasswordPolicy(v OrganizationsettingsputPasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *Organizationsettingsput) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetShowIntro

`func (o *Organizationsettingsput) GetShowIntro() bool`

GetShowIntro returns the ShowIntro field if non-nil, zero value otherwise.

### GetShowIntroOk

`func (o *Organizationsettingsput) GetShowIntroOk() (*bool, bool)`

GetShowIntroOk returns a tuple with the ShowIntro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIntro

`func (o *Organizationsettingsput) SetShowIntro(v bool)`

SetShowIntro sets ShowIntro field to given value.

### HasShowIntro

`func (o *Organizationsettingsput) HasShowIntro() bool`

HasShowIntro returns a boolean if a field has been set.

### GetSystemUserPasswordExpirationInDays

`func (o *Organizationsettingsput) GetSystemUserPasswordExpirationInDays() int32`

GetSystemUserPasswordExpirationInDays returns the SystemUserPasswordExpirationInDays field if non-nil, zero value otherwise.

### GetSystemUserPasswordExpirationInDaysOk

`func (o *Organizationsettingsput) GetSystemUserPasswordExpirationInDaysOk() (*int32, bool)`

GetSystemUserPasswordExpirationInDaysOk returns a tuple with the SystemUserPasswordExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUserPasswordExpirationInDays

`func (o *Organizationsettingsput) SetSystemUserPasswordExpirationInDays(v int32)`

SetSystemUserPasswordExpirationInDays sets SystemUserPasswordExpirationInDays field to given value.

### HasSystemUserPasswordExpirationInDays

`func (o *Organizationsettingsput) HasSystemUserPasswordExpirationInDays() bool`

HasSystemUserPasswordExpirationInDays returns a boolean if a field has been set.

### GetSystemUsersCanEdit

`func (o *Organizationsettingsput) GetSystemUsersCanEdit() bool`

GetSystemUsersCanEdit returns the SystemUsersCanEdit field if non-nil, zero value otherwise.

### GetSystemUsersCanEditOk

`func (o *Organizationsettingsput) GetSystemUsersCanEditOk() (*bool, bool)`

GetSystemUsersCanEditOk returns a tuple with the SystemUsersCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUsersCanEdit

`func (o *Organizationsettingsput) SetSystemUsersCanEdit(v bool)`

SetSystemUsersCanEdit sets SystemUsersCanEdit field to given value.

### HasSystemUsersCanEdit

`func (o *Organizationsettingsput) HasSystemUsersCanEdit() bool`

HasSystemUsersCanEdit returns a boolean if a field has been set.

### GetSystemUsersCap

`func (o *Organizationsettingsput) GetSystemUsersCap() int32`

GetSystemUsersCap returns the SystemUsersCap field if non-nil, zero value otherwise.

### GetSystemUsersCapOk

`func (o *Organizationsettingsput) GetSystemUsersCapOk() (*int32, bool)`

GetSystemUsersCapOk returns a tuple with the SystemUsersCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUsersCap

`func (o *Organizationsettingsput) SetSystemUsersCap(v int32)`

SetSystemUsersCap sets SystemUsersCap field to given value.

### HasSystemUsersCap

`func (o *Organizationsettingsput) HasSystemUsersCap() bool`

HasSystemUsersCap returns a boolean if a field has been set.

### GetTrustedAppConfig

`func (o *Organizationsettingsput) GetTrustedAppConfig() TrustedappConfigPut`

GetTrustedAppConfig returns the TrustedAppConfig field if non-nil, zero value otherwise.

### GetTrustedAppConfigOk

`func (o *Organizationsettingsput) GetTrustedAppConfigOk() (*TrustedappConfigPut, bool)`

GetTrustedAppConfigOk returns a tuple with the TrustedAppConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedAppConfig

`func (o *Organizationsettingsput) SetTrustedAppConfig(v TrustedappConfigPut)`

SetTrustedAppConfig sets TrustedAppConfig field to given value.

### HasTrustedAppConfig

`func (o *Organizationsettingsput) HasTrustedAppConfig() bool`

HasTrustedAppConfig returns a boolean if a field has been set.

### GetUserPortal

`func (o *Organizationsettingsput) GetUserPortal() OrganizationsettingsUserPortal`

GetUserPortal returns the UserPortal field if non-nil, zero value otherwise.

### GetUserPortalOk

`func (o *Organizationsettingsput) GetUserPortalOk() (*OrganizationsettingsUserPortal, bool)`

GetUserPortalOk returns a tuple with the UserPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPortal

`func (o *Organizationsettingsput) SetUserPortal(v OrganizationsettingsUserPortal)`

SetUserPortal sets UserPortal field to given value.

### HasUserPortal

`func (o *Organizationsettingsput) HasUserPortal() bool`

HasUserPortal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Organizationsettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentVersion** | Pointer to **string** |  | [optional] 
**BetaFeatures** | Pointer to **map[string]interface{}** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**ContactName** | Pointer to **string** |  | [optional] 
**DeviceIdentificationEnabled** | Pointer to **bool** |  | [optional] 
**DisableCommandRunner** | Pointer to **bool** |  | [optional] 
**DisableGoogleLogin** | Pointer to **bool** |  | [optional] 
**DisableLdap** | Pointer to **bool** |  | [optional] 
**DisableUM** | Pointer to **bool** |  | [optional] 
**DuplicateLDAPGroups** | Pointer to **bool** |  | [optional] 
**EmailDisclaimer** | Pointer to **string** |  | [optional] 
**EnableGoogleApps** | Pointer to **bool** |  | [optional] 
**EnableManagedUID** | Pointer to **bool** |  | [optional] 
**EnableO365** | Pointer to **bool** |  | [optional] 
**EnableUserPortalAgentInstall** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**OrganizationsettingsFeatures**](OrganizationsettingsFeatures.md) |  | [optional] 
**GrowthData** | Pointer to **map[string]interface{}** | Object containing Optimizely experimentIds and states corresponding to them | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NewSystemUserStateDefaults** | Pointer to [**OrganizationsettingsNewSystemUserStateDefaults**](OrganizationsettingsNewSystemUserStateDefaults.md) |  | [optional] 
**PasswordCompliance** | Pointer to **string** |  | [optional] 
**PasswordPolicy** | Pointer to [**OrganizationsettingsPasswordPolicy**](OrganizationsettingsPasswordPolicy.md) |  | [optional] 
**PendingDelete** | Pointer to **bool** |  | [optional] 
**ShowIntro** | Pointer to **bool** |  | [optional] 
**SystemUserPasswordExpirationInDays** | Pointer to **int32** |  | [optional] 
**SystemUsersCanEdit** | Pointer to **bool** |  | [optional] 
**SystemUsersCap** | Pointer to **int32** |  | [optional] 
**TrustedAppConfig** | Pointer to [**TrustedappConfigGet**](TrustedappConfigGet.md) |  | [optional] 
**UserPortal** | Pointer to [**OrganizationsettingsUserPortal**](OrganizationsettingsUserPortal.md) |  | [optional] 
**WindowsMDM** | Pointer to [**OrganizationsettingsWindowsMDM**](OrganizationsettingsWindowsMDM.md) |  | [optional] 

## Methods

### NewOrganizationsettings

`func NewOrganizationsettings() *Organizationsettings`

NewOrganizationsettings instantiates a new Organizationsettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsettingsWithDefaults

`func NewOrganizationsettingsWithDefaults() *Organizationsettings`

NewOrganizationsettingsWithDefaults instantiates a new Organizationsettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentVersion

`func (o *Organizationsettings) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *Organizationsettings) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *Organizationsettings) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *Organizationsettings) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetBetaFeatures

`func (o *Organizationsettings) GetBetaFeatures() map[string]interface{}`

GetBetaFeatures returns the BetaFeatures field if non-nil, zero value otherwise.

### GetBetaFeaturesOk

`func (o *Organizationsettings) GetBetaFeaturesOk() (*map[string]interface{}, bool)`

GetBetaFeaturesOk returns a tuple with the BetaFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaFeatures

`func (o *Organizationsettings) SetBetaFeatures(v map[string]interface{})`

SetBetaFeatures sets BetaFeatures field to given value.

### HasBetaFeatures

`func (o *Organizationsettings) HasBetaFeatures() bool`

HasBetaFeatures returns a boolean if a field has been set.

### GetContactEmail

`func (o *Organizationsettings) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Organizationsettings) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Organizationsettings) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Organizationsettings) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetContactName

`func (o *Organizationsettings) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *Organizationsettings) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *Organizationsettings) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *Organizationsettings) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetDeviceIdentificationEnabled

`func (o *Organizationsettings) GetDeviceIdentificationEnabled() bool`

GetDeviceIdentificationEnabled returns the DeviceIdentificationEnabled field if non-nil, zero value otherwise.

### GetDeviceIdentificationEnabledOk

`func (o *Organizationsettings) GetDeviceIdentificationEnabledOk() (*bool, bool)`

GetDeviceIdentificationEnabledOk returns a tuple with the DeviceIdentificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentificationEnabled

`func (o *Organizationsettings) SetDeviceIdentificationEnabled(v bool)`

SetDeviceIdentificationEnabled sets DeviceIdentificationEnabled field to given value.

### HasDeviceIdentificationEnabled

`func (o *Organizationsettings) HasDeviceIdentificationEnabled() bool`

HasDeviceIdentificationEnabled returns a boolean if a field has been set.

### GetDisableCommandRunner

`func (o *Organizationsettings) GetDisableCommandRunner() bool`

GetDisableCommandRunner returns the DisableCommandRunner field if non-nil, zero value otherwise.

### GetDisableCommandRunnerOk

`func (o *Organizationsettings) GetDisableCommandRunnerOk() (*bool, bool)`

GetDisableCommandRunnerOk returns a tuple with the DisableCommandRunner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCommandRunner

`func (o *Organizationsettings) SetDisableCommandRunner(v bool)`

SetDisableCommandRunner sets DisableCommandRunner field to given value.

### HasDisableCommandRunner

`func (o *Organizationsettings) HasDisableCommandRunner() bool`

HasDisableCommandRunner returns a boolean if a field has been set.

### GetDisableGoogleLogin

`func (o *Organizationsettings) GetDisableGoogleLogin() bool`

GetDisableGoogleLogin returns the DisableGoogleLogin field if non-nil, zero value otherwise.

### GetDisableGoogleLoginOk

`func (o *Organizationsettings) GetDisableGoogleLoginOk() (*bool, bool)`

GetDisableGoogleLoginOk returns a tuple with the DisableGoogleLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableGoogleLogin

`func (o *Organizationsettings) SetDisableGoogleLogin(v bool)`

SetDisableGoogleLogin sets DisableGoogleLogin field to given value.

### HasDisableGoogleLogin

`func (o *Organizationsettings) HasDisableGoogleLogin() bool`

HasDisableGoogleLogin returns a boolean if a field has been set.

### GetDisableLdap

`func (o *Organizationsettings) GetDisableLdap() bool`

GetDisableLdap returns the DisableLdap field if non-nil, zero value otherwise.

### GetDisableLdapOk

`func (o *Organizationsettings) GetDisableLdapOk() (*bool, bool)`

GetDisableLdapOk returns a tuple with the DisableLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLdap

`func (o *Organizationsettings) SetDisableLdap(v bool)`

SetDisableLdap sets DisableLdap field to given value.

### HasDisableLdap

`func (o *Organizationsettings) HasDisableLdap() bool`

HasDisableLdap returns a boolean if a field has been set.

### GetDisableUM

`func (o *Organizationsettings) GetDisableUM() bool`

GetDisableUM returns the DisableUM field if non-nil, zero value otherwise.

### GetDisableUMOk

`func (o *Organizationsettings) GetDisableUMOk() (*bool, bool)`

GetDisableUMOk returns a tuple with the DisableUM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUM

`func (o *Organizationsettings) SetDisableUM(v bool)`

SetDisableUM sets DisableUM field to given value.

### HasDisableUM

`func (o *Organizationsettings) HasDisableUM() bool`

HasDisableUM returns a boolean if a field has been set.

### GetDuplicateLDAPGroups

`func (o *Organizationsettings) GetDuplicateLDAPGroups() bool`

GetDuplicateLDAPGroups returns the DuplicateLDAPGroups field if non-nil, zero value otherwise.

### GetDuplicateLDAPGroupsOk

`func (o *Organizationsettings) GetDuplicateLDAPGroupsOk() (*bool, bool)`

GetDuplicateLDAPGroupsOk returns a tuple with the DuplicateLDAPGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateLDAPGroups

`func (o *Organizationsettings) SetDuplicateLDAPGroups(v bool)`

SetDuplicateLDAPGroups sets DuplicateLDAPGroups field to given value.

### HasDuplicateLDAPGroups

`func (o *Organizationsettings) HasDuplicateLDAPGroups() bool`

HasDuplicateLDAPGroups returns a boolean if a field has been set.

### GetEmailDisclaimer

`func (o *Organizationsettings) GetEmailDisclaimer() string`

GetEmailDisclaimer returns the EmailDisclaimer field if non-nil, zero value otherwise.

### GetEmailDisclaimerOk

`func (o *Organizationsettings) GetEmailDisclaimerOk() (*string, bool)`

GetEmailDisclaimerOk returns a tuple with the EmailDisclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDisclaimer

`func (o *Organizationsettings) SetEmailDisclaimer(v string)`

SetEmailDisclaimer sets EmailDisclaimer field to given value.

### HasEmailDisclaimer

`func (o *Organizationsettings) HasEmailDisclaimer() bool`

HasEmailDisclaimer returns a boolean if a field has been set.

### GetEnableGoogleApps

`func (o *Organizationsettings) GetEnableGoogleApps() bool`

GetEnableGoogleApps returns the EnableGoogleApps field if non-nil, zero value otherwise.

### GetEnableGoogleAppsOk

`func (o *Organizationsettings) GetEnableGoogleAppsOk() (*bool, bool)`

GetEnableGoogleAppsOk returns a tuple with the EnableGoogleApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGoogleApps

`func (o *Organizationsettings) SetEnableGoogleApps(v bool)`

SetEnableGoogleApps sets EnableGoogleApps field to given value.

### HasEnableGoogleApps

`func (o *Organizationsettings) HasEnableGoogleApps() bool`

HasEnableGoogleApps returns a boolean if a field has been set.

### GetEnableManagedUID

`func (o *Organizationsettings) GetEnableManagedUID() bool`

GetEnableManagedUID returns the EnableManagedUID field if non-nil, zero value otherwise.

### GetEnableManagedUIDOk

`func (o *Organizationsettings) GetEnableManagedUIDOk() (*bool, bool)`

GetEnableManagedUIDOk returns a tuple with the EnableManagedUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableManagedUID

`func (o *Organizationsettings) SetEnableManagedUID(v bool)`

SetEnableManagedUID sets EnableManagedUID field to given value.

### HasEnableManagedUID

`func (o *Organizationsettings) HasEnableManagedUID() bool`

HasEnableManagedUID returns a boolean if a field has been set.

### GetEnableO365

`func (o *Organizationsettings) GetEnableO365() bool`

GetEnableO365 returns the EnableO365 field if non-nil, zero value otherwise.

### GetEnableO365Ok

`func (o *Organizationsettings) GetEnableO365Ok() (*bool, bool)`

GetEnableO365Ok returns a tuple with the EnableO365 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableO365

`func (o *Organizationsettings) SetEnableO365(v bool)`

SetEnableO365 sets EnableO365 field to given value.

### HasEnableO365

`func (o *Organizationsettings) HasEnableO365() bool`

HasEnableO365 returns a boolean if a field has been set.

### GetEnableUserPortalAgentInstall

`func (o *Organizationsettings) GetEnableUserPortalAgentInstall() bool`

GetEnableUserPortalAgentInstall returns the EnableUserPortalAgentInstall field if non-nil, zero value otherwise.

### GetEnableUserPortalAgentInstallOk

`func (o *Organizationsettings) GetEnableUserPortalAgentInstallOk() (*bool, bool)`

GetEnableUserPortalAgentInstallOk returns a tuple with the EnableUserPortalAgentInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserPortalAgentInstall

`func (o *Organizationsettings) SetEnableUserPortalAgentInstall(v bool)`

SetEnableUserPortalAgentInstall sets EnableUserPortalAgentInstall field to given value.

### HasEnableUserPortalAgentInstall

`func (o *Organizationsettings) HasEnableUserPortalAgentInstall() bool`

HasEnableUserPortalAgentInstall returns a boolean if a field has been set.

### GetFeatures

`func (o *Organizationsettings) GetFeatures() OrganizationsettingsFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Organizationsettings) GetFeaturesOk() (*OrganizationsettingsFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Organizationsettings) SetFeatures(v OrganizationsettingsFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Organizationsettings) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetGrowthData

`func (o *Organizationsettings) GetGrowthData() map[string]interface{}`

GetGrowthData returns the GrowthData field if non-nil, zero value otherwise.

### GetGrowthDataOk

`func (o *Organizationsettings) GetGrowthDataOk() (*map[string]interface{}, bool)`

GetGrowthDataOk returns a tuple with the GrowthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthData

`func (o *Organizationsettings) SetGrowthData(v map[string]interface{})`

SetGrowthData sets GrowthData field to given value.

### HasGrowthData

`func (o *Organizationsettings) HasGrowthData() bool`

HasGrowthData returns a boolean if a field has been set.

### GetLogo

`func (o *Organizationsettings) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Organizationsettings) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Organizationsettings) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Organizationsettings) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *Organizationsettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organizationsettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organizationsettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organizationsettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewSystemUserStateDefaults

`func (o *Organizationsettings) GetNewSystemUserStateDefaults() OrganizationsettingsNewSystemUserStateDefaults`

GetNewSystemUserStateDefaults returns the NewSystemUserStateDefaults field if non-nil, zero value otherwise.

### GetNewSystemUserStateDefaultsOk

`func (o *Organizationsettings) GetNewSystemUserStateDefaultsOk() (*OrganizationsettingsNewSystemUserStateDefaults, bool)`

GetNewSystemUserStateDefaultsOk returns a tuple with the NewSystemUserStateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSystemUserStateDefaults

`func (o *Organizationsettings) SetNewSystemUserStateDefaults(v OrganizationsettingsNewSystemUserStateDefaults)`

SetNewSystemUserStateDefaults sets NewSystemUserStateDefaults field to given value.

### HasNewSystemUserStateDefaults

`func (o *Organizationsettings) HasNewSystemUserStateDefaults() bool`

HasNewSystemUserStateDefaults returns a boolean if a field has been set.

### GetPasswordCompliance

`func (o *Organizationsettings) GetPasswordCompliance() string`

GetPasswordCompliance returns the PasswordCompliance field if non-nil, zero value otherwise.

### GetPasswordComplianceOk

`func (o *Organizationsettings) GetPasswordComplianceOk() (*string, bool)`

GetPasswordComplianceOk returns a tuple with the PasswordCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCompliance

`func (o *Organizationsettings) SetPasswordCompliance(v string)`

SetPasswordCompliance sets PasswordCompliance field to given value.

### HasPasswordCompliance

`func (o *Organizationsettings) HasPasswordCompliance() bool`

HasPasswordCompliance returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *Organizationsettings) GetPasswordPolicy() OrganizationsettingsPasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *Organizationsettings) GetPasswordPolicyOk() (*OrganizationsettingsPasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *Organizationsettings) SetPasswordPolicy(v OrganizationsettingsPasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *Organizationsettings) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetPendingDelete

`func (o *Organizationsettings) GetPendingDelete() bool`

GetPendingDelete returns the PendingDelete field if non-nil, zero value otherwise.

### GetPendingDeleteOk

`func (o *Organizationsettings) GetPendingDeleteOk() (*bool, bool)`

GetPendingDeleteOk returns a tuple with the PendingDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDelete

`func (o *Organizationsettings) SetPendingDelete(v bool)`

SetPendingDelete sets PendingDelete field to given value.

### HasPendingDelete

`func (o *Organizationsettings) HasPendingDelete() bool`

HasPendingDelete returns a boolean if a field has been set.

### GetShowIntro

`func (o *Organizationsettings) GetShowIntro() bool`

GetShowIntro returns the ShowIntro field if non-nil, zero value otherwise.

### GetShowIntroOk

`func (o *Organizationsettings) GetShowIntroOk() (*bool, bool)`

GetShowIntroOk returns a tuple with the ShowIntro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIntro

`func (o *Organizationsettings) SetShowIntro(v bool)`

SetShowIntro sets ShowIntro field to given value.

### HasShowIntro

`func (o *Organizationsettings) HasShowIntro() bool`

HasShowIntro returns a boolean if a field has been set.

### GetSystemUserPasswordExpirationInDays

`func (o *Organizationsettings) GetSystemUserPasswordExpirationInDays() int32`

GetSystemUserPasswordExpirationInDays returns the SystemUserPasswordExpirationInDays field if non-nil, zero value otherwise.

### GetSystemUserPasswordExpirationInDaysOk

`func (o *Organizationsettings) GetSystemUserPasswordExpirationInDaysOk() (*int32, bool)`

GetSystemUserPasswordExpirationInDaysOk returns a tuple with the SystemUserPasswordExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUserPasswordExpirationInDays

`func (o *Organizationsettings) SetSystemUserPasswordExpirationInDays(v int32)`

SetSystemUserPasswordExpirationInDays sets SystemUserPasswordExpirationInDays field to given value.

### HasSystemUserPasswordExpirationInDays

`func (o *Organizationsettings) HasSystemUserPasswordExpirationInDays() bool`

HasSystemUserPasswordExpirationInDays returns a boolean if a field has been set.

### GetSystemUsersCanEdit

`func (o *Organizationsettings) GetSystemUsersCanEdit() bool`

GetSystemUsersCanEdit returns the SystemUsersCanEdit field if non-nil, zero value otherwise.

### GetSystemUsersCanEditOk

`func (o *Organizationsettings) GetSystemUsersCanEditOk() (*bool, bool)`

GetSystemUsersCanEditOk returns a tuple with the SystemUsersCanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUsersCanEdit

`func (o *Organizationsettings) SetSystemUsersCanEdit(v bool)`

SetSystemUsersCanEdit sets SystemUsersCanEdit field to given value.

### HasSystemUsersCanEdit

`func (o *Organizationsettings) HasSystemUsersCanEdit() bool`

HasSystemUsersCanEdit returns a boolean if a field has been set.

### GetSystemUsersCap

`func (o *Organizationsettings) GetSystemUsersCap() int32`

GetSystemUsersCap returns the SystemUsersCap field if non-nil, zero value otherwise.

### GetSystemUsersCapOk

`func (o *Organizationsettings) GetSystemUsersCapOk() (*int32, bool)`

GetSystemUsersCapOk returns a tuple with the SystemUsersCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUsersCap

`func (o *Organizationsettings) SetSystemUsersCap(v int32)`

SetSystemUsersCap sets SystemUsersCap field to given value.

### HasSystemUsersCap

`func (o *Organizationsettings) HasSystemUsersCap() bool`

HasSystemUsersCap returns a boolean if a field has been set.

### GetTrustedAppConfig

`func (o *Organizationsettings) GetTrustedAppConfig() TrustedappConfigGet`

GetTrustedAppConfig returns the TrustedAppConfig field if non-nil, zero value otherwise.

### GetTrustedAppConfigOk

`func (o *Organizationsettings) GetTrustedAppConfigOk() (*TrustedappConfigGet, bool)`

GetTrustedAppConfigOk returns a tuple with the TrustedAppConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedAppConfig

`func (o *Organizationsettings) SetTrustedAppConfig(v TrustedappConfigGet)`

SetTrustedAppConfig sets TrustedAppConfig field to given value.

### HasTrustedAppConfig

`func (o *Organizationsettings) HasTrustedAppConfig() bool`

HasTrustedAppConfig returns a boolean if a field has been set.

### GetUserPortal

`func (o *Organizationsettings) GetUserPortal() OrganizationsettingsUserPortal`

GetUserPortal returns the UserPortal field if non-nil, zero value otherwise.

### GetUserPortalOk

`func (o *Organizationsettings) GetUserPortalOk() (*OrganizationsettingsUserPortal, bool)`

GetUserPortalOk returns a tuple with the UserPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPortal

`func (o *Organizationsettings) SetUserPortal(v OrganizationsettingsUserPortal)`

SetUserPortal sets UserPortal field to given value.

### HasUserPortal

`func (o *Organizationsettings) HasUserPortal() bool`

HasUserPortal returns a boolean if a field has been set.

### GetWindowsMDM

`func (o *Organizationsettings) GetWindowsMDM() OrganizationsettingsWindowsMDM`

GetWindowsMDM returns the WindowsMDM field if non-nil, zero value otherwise.

### GetWindowsMDMOk

`func (o *Organizationsettings) GetWindowsMDMOk() (*OrganizationsettingsWindowsMDM, bool)`

GetWindowsMDMOk returns a tuple with the WindowsMDM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsMDM

`func (o *Organizationsettings) SetWindowsMDM(v OrganizationsettingsWindowsMDM)`

SetWindowsMDM sets WindowsMDM field to given value.

### HasWindowsMDM

`func (o *Organizationsettings) HasWindowsMDM() bool`

HasWindowsMDM returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



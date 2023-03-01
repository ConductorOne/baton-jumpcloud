# Systemuserputpost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountLocked** | Pointer to **bool** |  | [optional] 
**Activated** | Pointer to **bool** |  | [optional] 
**Addresses** | Pointer to [**[]SystemuserputpostAddressesInner**](SystemuserputpostAddressesInner.md) |  | [optional] 
**AllowPublicKey** | Pointer to **bool** |  | [optional] 
**AlternateEmail** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]SystemuserputAttributesInner**](SystemuserputAttributesInner.md) |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**CostCenter** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisableDeviceMaxLoginAttempts** | Pointer to **bool** |  | [optional] 
**Displayname** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**EmployeeIdentifier** | Pointer to **string** | Must be unique per user.  | [optional] 
**EmployeeType** | Pointer to **string** |  | [optional] 
**EnableManagedUid** | Pointer to **bool** |  | [optional] 
**EnableUserPortalMultifactor** | Pointer to **bool** |  | [optional] 
**ExternalDn** | Pointer to **string** |  | [optional] 
**ExternalPasswordExpirationDate** | Pointer to **time.Time** |  | [optional] 
**ExternalSourceType** | Pointer to **string** |  | [optional] 
**ExternallyManaged** | Pointer to **bool** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**LdapBindingUser** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**ManagedAppleId** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to **string** | Relation with another systemuser to identify the last as a manager. | [optional] 
**Mfa** | Pointer to [**Mfa**](Mfa.md) |  | [optional] 
**Middlename** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordNeverExpires** | Pointer to **bool** |  | [optional] 
**PasswordlessSudo** | Pointer to **bool** |  | [optional] 
**PhoneNumbers** | Pointer to [**[]SystemuserputpostPhoneNumbersInner**](SystemuserputpostPhoneNumbersInner.md) |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**RecoveryEmail** | Pointer to [**SystemuserputpostRecoveryEmail**](SystemuserputpostRecoveryEmail.md) |  | [optional] 
**Relationships** | Pointer to [**[]SystemuserputRelationshipsInner**](SystemuserputRelationshipsInner.md) |  | [optional] 
**SambaServiceUser** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Sudo** | Pointer to **bool** |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UnixGuid** | Pointer to **int32** |  | [optional] 
**UnixUid** | Pointer to **int32** |  | [optional] 
**Username** | **string** |  | 

## Methods

### NewSystemuserputpost

`func NewSystemuserputpost(email string, username string, ) *Systemuserputpost`

NewSystemuserputpost instantiates a new Systemuserputpost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemuserputpostWithDefaults

`func NewSystemuserputpostWithDefaults() *Systemuserputpost`

NewSystemuserputpostWithDefaults instantiates a new Systemuserputpost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLocked

`func (o *Systemuserputpost) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *Systemuserputpost) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *Systemuserputpost) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *Systemuserputpost) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetActivated

`func (o *Systemuserputpost) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *Systemuserputpost) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *Systemuserputpost) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *Systemuserputpost) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetAddresses

`func (o *Systemuserputpost) GetAddresses() []SystemuserputpostAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Systemuserputpost) GetAddressesOk() (*[]SystemuserputpostAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Systemuserputpost) SetAddresses(v []SystemuserputpostAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Systemuserputpost) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAllowPublicKey

`func (o *Systemuserputpost) GetAllowPublicKey() bool`

GetAllowPublicKey returns the AllowPublicKey field if non-nil, zero value otherwise.

### GetAllowPublicKeyOk

`func (o *Systemuserputpost) GetAllowPublicKeyOk() (*bool, bool)`

GetAllowPublicKeyOk returns a tuple with the AllowPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPublicKey

`func (o *Systemuserputpost) SetAllowPublicKey(v bool)`

SetAllowPublicKey sets AllowPublicKey field to given value.

### HasAllowPublicKey

`func (o *Systemuserputpost) HasAllowPublicKey() bool`

HasAllowPublicKey returns a boolean if a field has been set.

### GetAlternateEmail

`func (o *Systemuserputpost) GetAlternateEmail() string`

GetAlternateEmail returns the AlternateEmail field if non-nil, zero value otherwise.

### GetAlternateEmailOk

`func (o *Systemuserputpost) GetAlternateEmailOk() (*string, bool)`

GetAlternateEmailOk returns a tuple with the AlternateEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateEmail

`func (o *Systemuserputpost) SetAlternateEmail(v string)`

SetAlternateEmail sets AlternateEmail field to given value.

### HasAlternateEmail

`func (o *Systemuserputpost) HasAlternateEmail() bool`

HasAlternateEmail returns a boolean if a field has been set.

### GetAttributes

`func (o *Systemuserputpost) GetAttributes() []SystemuserputAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Systemuserputpost) GetAttributesOk() (*[]SystemuserputAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Systemuserputpost) SetAttributes(v []SystemuserputAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Systemuserputpost) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCompany

`func (o *Systemuserputpost) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Systemuserputpost) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Systemuserputpost) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Systemuserputpost) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCostCenter

`func (o *Systemuserputpost) GetCostCenter() string`

GetCostCenter returns the CostCenter field if non-nil, zero value otherwise.

### GetCostCenterOk

`func (o *Systemuserputpost) GetCostCenterOk() (*string, bool)`

GetCostCenterOk returns a tuple with the CostCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCenter

`func (o *Systemuserputpost) SetCostCenter(v string)`

SetCostCenter sets CostCenter field to given value.

### HasCostCenter

`func (o *Systemuserputpost) HasCostCenter() bool`

HasCostCenter returns a boolean if a field has been set.

### GetDepartment

`func (o *Systemuserputpost) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Systemuserputpost) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Systemuserputpost) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Systemuserputpost) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDescription

`func (o *Systemuserputpost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Systemuserputpost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Systemuserputpost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Systemuserputpost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableDeviceMaxLoginAttempts

`func (o *Systemuserputpost) GetDisableDeviceMaxLoginAttempts() bool`

GetDisableDeviceMaxLoginAttempts returns the DisableDeviceMaxLoginAttempts field if non-nil, zero value otherwise.

### GetDisableDeviceMaxLoginAttemptsOk

`func (o *Systemuserputpost) GetDisableDeviceMaxLoginAttemptsOk() (*bool, bool)`

GetDisableDeviceMaxLoginAttemptsOk returns a tuple with the DisableDeviceMaxLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDeviceMaxLoginAttempts

`func (o *Systemuserputpost) SetDisableDeviceMaxLoginAttempts(v bool)`

SetDisableDeviceMaxLoginAttempts sets DisableDeviceMaxLoginAttempts field to given value.

### HasDisableDeviceMaxLoginAttempts

`func (o *Systemuserputpost) HasDisableDeviceMaxLoginAttempts() bool`

HasDisableDeviceMaxLoginAttempts returns a boolean if a field has been set.

### GetDisplayname

`func (o *Systemuserputpost) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *Systemuserputpost) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *Systemuserputpost) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *Systemuserputpost) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### GetEmail

`func (o *Systemuserputpost) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Systemuserputpost) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Systemuserputpost) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmployeeIdentifier

`func (o *Systemuserputpost) GetEmployeeIdentifier() string`

GetEmployeeIdentifier returns the EmployeeIdentifier field if non-nil, zero value otherwise.

### GetEmployeeIdentifierOk

`func (o *Systemuserputpost) GetEmployeeIdentifierOk() (*string, bool)`

GetEmployeeIdentifierOk returns a tuple with the EmployeeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIdentifier

`func (o *Systemuserputpost) SetEmployeeIdentifier(v string)`

SetEmployeeIdentifier sets EmployeeIdentifier field to given value.

### HasEmployeeIdentifier

`func (o *Systemuserputpost) HasEmployeeIdentifier() bool`

HasEmployeeIdentifier returns a boolean if a field has been set.

### GetEmployeeType

`func (o *Systemuserputpost) GetEmployeeType() string`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *Systemuserputpost) GetEmployeeTypeOk() (*string, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *Systemuserputpost) SetEmployeeType(v string)`

SetEmployeeType sets EmployeeType field to given value.

### HasEmployeeType

`func (o *Systemuserputpost) HasEmployeeType() bool`

HasEmployeeType returns a boolean if a field has been set.

### GetEnableManagedUid

`func (o *Systemuserputpost) GetEnableManagedUid() bool`

GetEnableManagedUid returns the EnableManagedUid field if non-nil, zero value otherwise.

### GetEnableManagedUidOk

`func (o *Systemuserputpost) GetEnableManagedUidOk() (*bool, bool)`

GetEnableManagedUidOk returns a tuple with the EnableManagedUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableManagedUid

`func (o *Systemuserputpost) SetEnableManagedUid(v bool)`

SetEnableManagedUid sets EnableManagedUid field to given value.

### HasEnableManagedUid

`func (o *Systemuserputpost) HasEnableManagedUid() bool`

HasEnableManagedUid returns a boolean if a field has been set.

### GetEnableUserPortalMultifactor

`func (o *Systemuserputpost) GetEnableUserPortalMultifactor() bool`

GetEnableUserPortalMultifactor returns the EnableUserPortalMultifactor field if non-nil, zero value otherwise.

### GetEnableUserPortalMultifactorOk

`func (o *Systemuserputpost) GetEnableUserPortalMultifactorOk() (*bool, bool)`

GetEnableUserPortalMultifactorOk returns a tuple with the EnableUserPortalMultifactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserPortalMultifactor

`func (o *Systemuserputpost) SetEnableUserPortalMultifactor(v bool)`

SetEnableUserPortalMultifactor sets EnableUserPortalMultifactor field to given value.

### HasEnableUserPortalMultifactor

`func (o *Systemuserputpost) HasEnableUserPortalMultifactor() bool`

HasEnableUserPortalMultifactor returns a boolean if a field has been set.

### GetExternalDn

`func (o *Systemuserputpost) GetExternalDn() string`

GetExternalDn returns the ExternalDn field if non-nil, zero value otherwise.

### GetExternalDnOk

`func (o *Systemuserputpost) GetExternalDnOk() (*string, bool)`

GetExternalDnOk returns a tuple with the ExternalDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDn

`func (o *Systemuserputpost) SetExternalDn(v string)`

SetExternalDn sets ExternalDn field to given value.

### HasExternalDn

`func (o *Systemuserputpost) HasExternalDn() bool`

HasExternalDn returns a boolean if a field has been set.

### GetExternalPasswordExpirationDate

`func (o *Systemuserputpost) GetExternalPasswordExpirationDate() time.Time`

GetExternalPasswordExpirationDate returns the ExternalPasswordExpirationDate field if non-nil, zero value otherwise.

### GetExternalPasswordExpirationDateOk

`func (o *Systemuserputpost) GetExternalPasswordExpirationDateOk() (*time.Time, bool)`

GetExternalPasswordExpirationDateOk returns a tuple with the ExternalPasswordExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPasswordExpirationDate

`func (o *Systemuserputpost) SetExternalPasswordExpirationDate(v time.Time)`

SetExternalPasswordExpirationDate sets ExternalPasswordExpirationDate field to given value.

### HasExternalPasswordExpirationDate

`func (o *Systemuserputpost) HasExternalPasswordExpirationDate() bool`

HasExternalPasswordExpirationDate returns a boolean if a field has been set.

### GetExternalSourceType

`func (o *Systemuserputpost) GetExternalSourceType() string`

GetExternalSourceType returns the ExternalSourceType field if non-nil, zero value otherwise.

### GetExternalSourceTypeOk

`func (o *Systemuserputpost) GetExternalSourceTypeOk() (*string, bool)`

GetExternalSourceTypeOk returns a tuple with the ExternalSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSourceType

`func (o *Systemuserputpost) SetExternalSourceType(v string)`

SetExternalSourceType sets ExternalSourceType field to given value.

### HasExternalSourceType

`func (o *Systemuserputpost) HasExternalSourceType() bool`

HasExternalSourceType returns a boolean if a field has been set.

### GetExternallyManaged

`func (o *Systemuserputpost) GetExternallyManaged() bool`

GetExternallyManaged returns the ExternallyManaged field if non-nil, zero value otherwise.

### GetExternallyManagedOk

`func (o *Systemuserputpost) GetExternallyManagedOk() (*bool, bool)`

GetExternallyManagedOk returns a tuple with the ExternallyManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyManaged

`func (o *Systemuserputpost) SetExternallyManaged(v bool)`

SetExternallyManaged sets ExternallyManaged field to given value.

### HasExternallyManaged

`func (o *Systemuserputpost) HasExternallyManaged() bool`

HasExternallyManaged returns a boolean if a field has been set.

### GetFirstname

`func (o *Systemuserputpost) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *Systemuserputpost) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *Systemuserputpost) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *Systemuserputpost) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetJobTitle

`func (o *Systemuserputpost) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Systemuserputpost) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Systemuserputpost) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Systemuserputpost) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastname

`func (o *Systemuserputpost) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *Systemuserputpost) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *Systemuserputpost) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *Systemuserputpost) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetLdapBindingUser

`func (o *Systemuserputpost) GetLdapBindingUser() bool`

GetLdapBindingUser returns the LdapBindingUser field if non-nil, zero value otherwise.

### GetLdapBindingUserOk

`func (o *Systemuserputpost) GetLdapBindingUserOk() (*bool, bool)`

GetLdapBindingUserOk returns a tuple with the LdapBindingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindingUser

`func (o *Systemuserputpost) SetLdapBindingUser(v bool)`

SetLdapBindingUser sets LdapBindingUser field to given value.

### HasLdapBindingUser

`func (o *Systemuserputpost) HasLdapBindingUser() bool`

HasLdapBindingUser returns a boolean if a field has been set.

### GetLocation

`func (o *Systemuserputpost) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Systemuserputpost) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Systemuserputpost) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Systemuserputpost) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetManagedAppleId

`func (o *Systemuserputpost) GetManagedAppleId() string`

GetManagedAppleId returns the ManagedAppleId field if non-nil, zero value otherwise.

### GetManagedAppleIdOk

`func (o *Systemuserputpost) GetManagedAppleIdOk() (*string, bool)`

GetManagedAppleIdOk returns a tuple with the ManagedAppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppleId

`func (o *Systemuserputpost) SetManagedAppleId(v string)`

SetManagedAppleId sets ManagedAppleId field to given value.

### HasManagedAppleId

`func (o *Systemuserputpost) HasManagedAppleId() bool`

HasManagedAppleId returns a boolean if a field has been set.

### GetManager

`func (o *Systemuserputpost) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Systemuserputpost) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Systemuserputpost) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *Systemuserputpost) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetMfa

`func (o *Systemuserputpost) GetMfa() Mfa`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *Systemuserputpost) GetMfaOk() (*Mfa, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *Systemuserputpost) SetMfa(v Mfa)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *Systemuserputpost) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetMiddlename

`func (o *Systemuserputpost) GetMiddlename() string`

GetMiddlename returns the Middlename field if non-nil, zero value otherwise.

### GetMiddlenameOk

`func (o *Systemuserputpost) GetMiddlenameOk() (*string, bool)`

GetMiddlenameOk returns a tuple with the Middlename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddlename

`func (o *Systemuserputpost) SetMiddlename(v string)`

SetMiddlename sets Middlename field to given value.

### HasMiddlename

`func (o *Systemuserputpost) HasMiddlename() bool`

HasMiddlename returns a boolean if a field has been set.

### GetPassword

`func (o *Systemuserputpost) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Systemuserputpost) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Systemuserputpost) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Systemuserputpost) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordNeverExpires

`func (o *Systemuserputpost) GetPasswordNeverExpires() bool`

GetPasswordNeverExpires returns the PasswordNeverExpires field if non-nil, zero value otherwise.

### GetPasswordNeverExpiresOk

`func (o *Systemuserputpost) GetPasswordNeverExpiresOk() (*bool, bool)`

GetPasswordNeverExpiresOk returns a tuple with the PasswordNeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordNeverExpires

`func (o *Systemuserputpost) SetPasswordNeverExpires(v bool)`

SetPasswordNeverExpires sets PasswordNeverExpires field to given value.

### HasPasswordNeverExpires

`func (o *Systemuserputpost) HasPasswordNeverExpires() bool`

HasPasswordNeverExpires returns a boolean if a field has been set.

### GetPasswordlessSudo

`func (o *Systemuserputpost) GetPasswordlessSudo() bool`

GetPasswordlessSudo returns the PasswordlessSudo field if non-nil, zero value otherwise.

### GetPasswordlessSudoOk

`func (o *Systemuserputpost) GetPasswordlessSudoOk() (*bool, bool)`

GetPasswordlessSudoOk returns a tuple with the PasswordlessSudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessSudo

`func (o *Systemuserputpost) SetPasswordlessSudo(v bool)`

SetPasswordlessSudo sets PasswordlessSudo field to given value.

### HasPasswordlessSudo

`func (o *Systemuserputpost) HasPasswordlessSudo() bool`

HasPasswordlessSudo returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *Systemuserputpost) GetPhoneNumbers() []SystemuserputpostPhoneNumbersInner`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *Systemuserputpost) GetPhoneNumbersOk() (*[]SystemuserputpostPhoneNumbersInner, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *Systemuserputpost) SetPhoneNumbers(v []SystemuserputpostPhoneNumbersInner)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *Systemuserputpost) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetPublicKey

`func (o *Systemuserputpost) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Systemuserputpost) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Systemuserputpost) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *Systemuserputpost) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRecoveryEmail

`func (o *Systemuserputpost) GetRecoveryEmail() SystemuserputpostRecoveryEmail`

GetRecoveryEmail returns the RecoveryEmail field if non-nil, zero value otherwise.

### GetRecoveryEmailOk

`func (o *Systemuserputpost) GetRecoveryEmailOk() (*SystemuserputpostRecoveryEmail, bool)`

GetRecoveryEmailOk returns a tuple with the RecoveryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryEmail

`func (o *Systemuserputpost) SetRecoveryEmail(v SystemuserputpostRecoveryEmail)`

SetRecoveryEmail sets RecoveryEmail field to given value.

### HasRecoveryEmail

`func (o *Systemuserputpost) HasRecoveryEmail() bool`

HasRecoveryEmail returns a boolean if a field has been set.

### GetRelationships

`func (o *Systemuserputpost) GetRelationships() []SystemuserputRelationshipsInner`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Systemuserputpost) GetRelationshipsOk() (*[]SystemuserputRelationshipsInner, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Systemuserputpost) SetRelationships(v []SystemuserputRelationshipsInner)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Systemuserputpost) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetSambaServiceUser

`func (o *Systemuserputpost) GetSambaServiceUser() bool`

GetSambaServiceUser returns the SambaServiceUser field if non-nil, zero value otherwise.

### GetSambaServiceUserOk

`func (o *Systemuserputpost) GetSambaServiceUserOk() (*bool, bool)`

GetSambaServiceUserOk returns a tuple with the SambaServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSambaServiceUser

`func (o *Systemuserputpost) SetSambaServiceUser(v bool)`

SetSambaServiceUser sets SambaServiceUser field to given value.

### HasSambaServiceUser

`func (o *Systemuserputpost) HasSambaServiceUser() bool`

HasSambaServiceUser returns a boolean if a field has been set.

### GetState

`func (o *Systemuserputpost) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Systemuserputpost) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Systemuserputpost) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Systemuserputpost) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSudo

`func (o *Systemuserputpost) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *Systemuserputpost) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *Systemuserputpost) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *Systemuserputpost) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSuspended

`func (o *Systemuserputpost) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Systemuserputpost) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Systemuserputpost) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Systemuserputpost) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetTags

`func (o *Systemuserputpost) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Systemuserputpost) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Systemuserputpost) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Systemuserputpost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnixGuid

`func (o *Systemuserputpost) GetUnixGuid() int32`

GetUnixGuid returns the UnixGuid field if non-nil, zero value otherwise.

### GetUnixGuidOk

`func (o *Systemuserputpost) GetUnixGuidOk() (*int32, bool)`

GetUnixGuidOk returns a tuple with the UnixGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixGuid

`func (o *Systemuserputpost) SetUnixGuid(v int32)`

SetUnixGuid sets UnixGuid field to given value.

### HasUnixGuid

`func (o *Systemuserputpost) HasUnixGuid() bool`

HasUnixGuid returns a boolean if a field has been set.

### GetUnixUid

`func (o *Systemuserputpost) GetUnixUid() int32`

GetUnixUid returns the UnixUid field if non-nil, zero value otherwise.

### GetUnixUidOk

`func (o *Systemuserputpost) GetUnixUidOk() (*int32, bool)`

GetUnixUidOk returns a tuple with the UnixUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUid

`func (o *Systemuserputpost) SetUnixUid(v int32)`

SetUnixUid sets UnixUid field to given value.

### HasUnixUid

`func (o *Systemuserputpost) HasUnixUid() bool`

HasUnixUid returns a boolean if a field has been set.

### GetUsername

`func (o *Systemuserputpost) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Systemuserputpost) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Systemuserputpost) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



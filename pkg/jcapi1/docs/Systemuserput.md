# Systemuserput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountLocked** | Pointer to **bool** |  | [optional] 
**Addresses** | Pointer to [**[]SystemuserputAddressesInner**](SystemuserputAddressesInner.md) | type, poBox, extendedAddress, streetAddress, locality, region, postalCode, country | [optional] 
**AllowPublicKey** | Pointer to **bool** |  | [optional] 
**AlternateEmail** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]SystemuserputAttributesInner**](SystemuserputAttributesInner.md) |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**CostCenter** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisableDeviceMaxLoginAttempts** | Pointer to **bool** |  | [optional] 
**Displayname** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmployeeIdentifier** | Pointer to **string** | Must be unique per user.  | [optional] 
**EmployeeType** | Pointer to **string** |  | [optional] 
**EnableManagedUid** | Pointer to **bool** |  | [optional] 
**EnableUserPortalMultifactor** | Pointer to **bool** |  | [optional] 
**ExternalDn** | Pointer to **string** |  | [optional] 
**ExternalPasswordExpirationDate** | Pointer to **string** |  | [optional] 
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
**PhoneNumbers** | Pointer to [**[]SystemuserputPhoneNumbersInner**](SystemuserputPhoneNumbersInner.md) |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**[]SystemuserputRelationshipsInner**](SystemuserputRelationshipsInner.md) |  | [optional] 
**SambaServiceUser** | Pointer to **bool** |  | [optional] 
**SshKeys** | Pointer to [**[]Sshkeypost**](Sshkeypost.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Sudo** | Pointer to **bool** |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UnixGuid** | Pointer to **int32** |  | [optional] 
**UnixUid** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemuserput

`func NewSystemuserput() *Systemuserput`

NewSystemuserput instantiates a new Systemuserput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemuserputWithDefaults

`func NewSystemuserputWithDefaults() *Systemuserput`

NewSystemuserputWithDefaults instantiates a new Systemuserput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountLocked

`func (o *Systemuserput) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *Systemuserput) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *Systemuserput) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *Systemuserput) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetAddresses

`func (o *Systemuserput) GetAddresses() []SystemuserputAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Systemuserput) GetAddressesOk() (*[]SystemuserputAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Systemuserput) SetAddresses(v []SystemuserputAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Systemuserput) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAllowPublicKey

`func (o *Systemuserput) GetAllowPublicKey() bool`

GetAllowPublicKey returns the AllowPublicKey field if non-nil, zero value otherwise.

### GetAllowPublicKeyOk

`func (o *Systemuserput) GetAllowPublicKeyOk() (*bool, bool)`

GetAllowPublicKeyOk returns a tuple with the AllowPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPublicKey

`func (o *Systemuserput) SetAllowPublicKey(v bool)`

SetAllowPublicKey sets AllowPublicKey field to given value.

### HasAllowPublicKey

`func (o *Systemuserput) HasAllowPublicKey() bool`

HasAllowPublicKey returns a boolean if a field has been set.

### GetAlternateEmail

`func (o *Systemuserput) GetAlternateEmail() string`

GetAlternateEmail returns the AlternateEmail field if non-nil, zero value otherwise.

### GetAlternateEmailOk

`func (o *Systemuserput) GetAlternateEmailOk() (*string, bool)`

GetAlternateEmailOk returns a tuple with the AlternateEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateEmail

`func (o *Systemuserput) SetAlternateEmail(v string)`

SetAlternateEmail sets AlternateEmail field to given value.

### HasAlternateEmail

`func (o *Systemuserput) HasAlternateEmail() bool`

HasAlternateEmail returns a boolean if a field has been set.

### GetAttributes

`func (o *Systemuserput) GetAttributes() []SystemuserputAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Systemuserput) GetAttributesOk() (*[]SystemuserputAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Systemuserput) SetAttributes(v []SystemuserputAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Systemuserput) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCompany

`func (o *Systemuserput) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Systemuserput) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Systemuserput) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Systemuserput) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCostCenter

`func (o *Systemuserput) GetCostCenter() string`

GetCostCenter returns the CostCenter field if non-nil, zero value otherwise.

### GetCostCenterOk

`func (o *Systemuserput) GetCostCenterOk() (*string, bool)`

GetCostCenterOk returns a tuple with the CostCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCenter

`func (o *Systemuserput) SetCostCenter(v string)`

SetCostCenter sets CostCenter field to given value.

### HasCostCenter

`func (o *Systemuserput) HasCostCenter() bool`

HasCostCenter returns a boolean if a field has been set.

### GetDepartment

`func (o *Systemuserput) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Systemuserput) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Systemuserput) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Systemuserput) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDescription

`func (o *Systemuserput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Systemuserput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Systemuserput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Systemuserput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableDeviceMaxLoginAttempts

`func (o *Systemuserput) GetDisableDeviceMaxLoginAttempts() bool`

GetDisableDeviceMaxLoginAttempts returns the DisableDeviceMaxLoginAttempts field if non-nil, zero value otherwise.

### GetDisableDeviceMaxLoginAttemptsOk

`func (o *Systemuserput) GetDisableDeviceMaxLoginAttemptsOk() (*bool, bool)`

GetDisableDeviceMaxLoginAttemptsOk returns a tuple with the DisableDeviceMaxLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDeviceMaxLoginAttempts

`func (o *Systemuserput) SetDisableDeviceMaxLoginAttempts(v bool)`

SetDisableDeviceMaxLoginAttempts sets DisableDeviceMaxLoginAttempts field to given value.

### HasDisableDeviceMaxLoginAttempts

`func (o *Systemuserput) HasDisableDeviceMaxLoginAttempts() bool`

HasDisableDeviceMaxLoginAttempts returns a boolean if a field has been set.

### GetDisplayname

`func (o *Systemuserput) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *Systemuserput) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *Systemuserput) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *Systemuserput) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### GetEmail

`func (o *Systemuserput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Systemuserput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Systemuserput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Systemuserput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmployeeIdentifier

`func (o *Systemuserput) GetEmployeeIdentifier() string`

GetEmployeeIdentifier returns the EmployeeIdentifier field if non-nil, zero value otherwise.

### GetEmployeeIdentifierOk

`func (o *Systemuserput) GetEmployeeIdentifierOk() (*string, bool)`

GetEmployeeIdentifierOk returns a tuple with the EmployeeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIdentifier

`func (o *Systemuserput) SetEmployeeIdentifier(v string)`

SetEmployeeIdentifier sets EmployeeIdentifier field to given value.

### HasEmployeeIdentifier

`func (o *Systemuserput) HasEmployeeIdentifier() bool`

HasEmployeeIdentifier returns a boolean if a field has been set.

### GetEmployeeType

`func (o *Systemuserput) GetEmployeeType() string`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *Systemuserput) GetEmployeeTypeOk() (*string, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *Systemuserput) SetEmployeeType(v string)`

SetEmployeeType sets EmployeeType field to given value.

### HasEmployeeType

`func (o *Systemuserput) HasEmployeeType() bool`

HasEmployeeType returns a boolean if a field has been set.

### GetEnableManagedUid

`func (o *Systemuserput) GetEnableManagedUid() bool`

GetEnableManagedUid returns the EnableManagedUid field if non-nil, zero value otherwise.

### GetEnableManagedUidOk

`func (o *Systemuserput) GetEnableManagedUidOk() (*bool, bool)`

GetEnableManagedUidOk returns a tuple with the EnableManagedUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableManagedUid

`func (o *Systemuserput) SetEnableManagedUid(v bool)`

SetEnableManagedUid sets EnableManagedUid field to given value.

### HasEnableManagedUid

`func (o *Systemuserput) HasEnableManagedUid() bool`

HasEnableManagedUid returns a boolean if a field has been set.

### GetEnableUserPortalMultifactor

`func (o *Systemuserput) GetEnableUserPortalMultifactor() bool`

GetEnableUserPortalMultifactor returns the EnableUserPortalMultifactor field if non-nil, zero value otherwise.

### GetEnableUserPortalMultifactorOk

`func (o *Systemuserput) GetEnableUserPortalMultifactorOk() (*bool, bool)`

GetEnableUserPortalMultifactorOk returns a tuple with the EnableUserPortalMultifactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserPortalMultifactor

`func (o *Systemuserput) SetEnableUserPortalMultifactor(v bool)`

SetEnableUserPortalMultifactor sets EnableUserPortalMultifactor field to given value.

### HasEnableUserPortalMultifactor

`func (o *Systemuserput) HasEnableUserPortalMultifactor() bool`

HasEnableUserPortalMultifactor returns a boolean if a field has been set.

### GetExternalDn

`func (o *Systemuserput) GetExternalDn() string`

GetExternalDn returns the ExternalDn field if non-nil, zero value otherwise.

### GetExternalDnOk

`func (o *Systemuserput) GetExternalDnOk() (*string, bool)`

GetExternalDnOk returns a tuple with the ExternalDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDn

`func (o *Systemuserput) SetExternalDn(v string)`

SetExternalDn sets ExternalDn field to given value.

### HasExternalDn

`func (o *Systemuserput) HasExternalDn() bool`

HasExternalDn returns a boolean if a field has been set.

### GetExternalPasswordExpirationDate

`func (o *Systemuserput) GetExternalPasswordExpirationDate() string`

GetExternalPasswordExpirationDate returns the ExternalPasswordExpirationDate field if non-nil, zero value otherwise.

### GetExternalPasswordExpirationDateOk

`func (o *Systemuserput) GetExternalPasswordExpirationDateOk() (*string, bool)`

GetExternalPasswordExpirationDateOk returns a tuple with the ExternalPasswordExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPasswordExpirationDate

`func (o *Systemuserput) SetExternalPasswordExpirationDate(v string)`

SetExternalPasswordExpirationDate sets ExternalPasswordExpirationDate field to given value.

### HasExternalPasswordExpirationDate

`func (o *Systemuserput) HasExternalPasswordExpirationDate() bool`

HasExternalPasswordExpirationDate returns a boolean if a field has been set.

### GetExternalSourceType

`func (o *Systemuserput) GetExternalSourceType() string`

GetExternalSourceType returns the ExternalSourceType field if non-nil, zero value otherwise.

### GetExternalSourceTypeOk

`func (o *Systemuserput) GetExternalSourceTypeOk() (*string, bool)`

GetExternalSourceTypeOk returns a tuple with the ExternalSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSourceType

`func (o *Systemuserput) SetExternalSourceType(v string)`

SetExternalSourceType sets ExternalSourceType field to given value.

### HasExternalSourceType

`func (o *Systemuserput) HasExternalSourceType() bool`

HasExternalSourceType returns a boolean if a field has been set.

### GetExternallyManaged

`func (o *Systemuserput) GetExternallyManaged() bool`

GetExternallyManaged returns the ExternallyManaged field if non-nil, zero value otherwise.

### GetExternallyManagedOk

`func (o *Systemuserput) GetExternallyManagedOk() (*bool, bool)`

GetExternallyManagedOk returns a tuple with the ExternallyManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyManaged

`func (o *Systemuserput) SetExternallyManaged(v bool)`

SetExternallyManaged sets ExternallyManaged field to given value.

### HasExternallyManaged

`func (o *Systemuserput) HasExternallyManaged() bool`

HasExternallyManaged returns a boolean if a field has been set.

### GetFirstname

`func (o *Systemuserput) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *Systemuserput) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *Systemuserput) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *Systemuserput) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetJobTitle

`func (o *Systemuserput) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Systemuserput) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Systemuserput) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Systemuserput) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastname

`func (o *Systemuserput) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *Systemuserput) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *Systemuserput) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *Systemuserput) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetLdapBindingUser

`func (o *Systemuserput) GetLdapBindingUser() bool`

GetLdapBindingUser returns the LdapBindingUser field if non-nil, zero value otherwise.

### GetLdapBindingUserOk

`func (o *Systemuserput) GetLdapBindingUserOk() (*bool, bool)`

GetLdapBindingUserOk returns a tuple with the LdapBindingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindingUser

`func (o *Systemuserput) SetLdapBindingUser(v bool)`

SetLdapBindingUser sets LdapBindingUser field to given value.

### HasLdapBindingUser

`func (o *Systemuserput) HasLdapBindingUser() bool`

HasLdapBindingUser returns a boolean if a field has been set.

### GetLocation

`func (o *Systemuserput) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Systemuserput) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Systemuserput) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Systemuserput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetManagedAppleId

`func (o *Systemuserput) GetManagedAppleId() string`

GetManagedAppleId returns the ManagedAppleId field if non-nil, zero value otherwise.

### GetManagedAppleIdOk

`func (o *Systemuserput) GetManagedAppleIdOk() (*string, bool)`

GetManagedAppleIdOk returns a tuple with the ManagedAppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppleId

`func (o *Systemuserput) SetManagedAppleId(v string)`

SetManagedAppleId sets ManagedAppleId field to given value.

### HasManagedAppleId

`func (o *Systemuserput) HasManagedAppleId() bool`

HasManagedAppleId returns a boolean if a field has been set.

### GetManager

`func (o *Systemuserput) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Systemuserput) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Systemuserput) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *Systemuserput) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetMfa

`func (o *Systemuserput) GetMfa() Mfa`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *Systemuserput) GetMfaOk() (*Mfa, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *Systemuserput) SetMfa(v Mfa)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *Systemuserput) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetMiddlename

`func (o *Systemuserput) GetMiddlename() string`

GetMiddlename returns the Middlename field if non-nil, zero value otherwise.

### GetMiddlenameOk

`func (o *Systemuserput) GetMiddlenameOk() (*string, bool)`

GetMiddlenameOk returns a tuple with the Middlename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddlename

`func (o *Systemuserput) SetMiddlename(v string)`

SetMiddlename sets Middlename field to given value.

### HasMiddlename

`func (o *Systemuserput) HasMiddlename() bool`

HasMiddlename returns a boolean if a field has been set.

### GetPassword

`func (o *Systemuserput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Systemuserput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Systemuserput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Systemuserput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordNeverExpires

`func (o *Systemuserput) GetPasswordNeverExpires() bool`

GetPasswordNeverExpires returns the PasswordNeverExpires field if non-nil, zero value otherwise.

### GetPasswordNeverExpiresOk

`func (o *Systemuserput) GetPasswordNeverExpiresOk() (*bool, bool)`

GetPasswordNeverExpiresOk returns a tuple with the PasswordNeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordNeverExpires

`func (o *Systemuserput) SetPasswordNeverExpires(v bool)`

SetPasswordNeverExpires sets PasswordNeverExpires field to given value.

### HasPasswordNeverExpires

`func (o *Systemuserput) HasPasswordNeverExpires() bool`

HasPasswordNeverExpires returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *Systemuserput) GetPhoneNumbers() []SystemuserputPhoneNumbersInner`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *Systemuserput) GetPhoneNumbersOk() (*[]SystemuserputPhoneNumbersInner, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *Systemuserput) SetPhoneNumbers(v []SystemuserputPhoneNumbersInner)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *Systemuserput) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetPublicKey

`func (o *Systemuserput) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Systemuserput) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Systemuserput) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *Systemuserput) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRelationships

`func (o *Systemuserput) GetRelationships() []SystemuserputRelationshipsInner`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Systemuserput) GetRelationshipsOk() (*[]SystemuserputRelationshipsInner, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Systemuserput) SetRelationships(v []SystemuserputRelationshipsInner)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Systemuserput) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetSambaServiceUser

`func (o *Systemuserput) GetSambaServiceUser() bool`

GetSambaServiceUser returns the SambaServiceUser field if non-nil, zero value otherwise.

### GetSambaServiceUserOk

`func (o *Systemuserput) GetSambaServiceUserOk() (*bool, bool)`

GetSambaServiceUserOk returns a tuple with the SambaServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSambaServiceUser

`func (o *Systemuserput) SetSambaServiceUser(v bool)`

SetSambaServiceUser sets SambaServiceUser field to given value.

### HasSambaServiceUser

`func (o *Systemuserput) HasSambaServiceUser() bool`

HasSambaServiceUser returns a boolean if a field has been set.

### GetSshKeys

`func (o *Systemuserput) GetSshKeys() []Sshkeypost`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *Systemuserput) GetSshKeysOk() (*[]Sshkeypost, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *Systemuserput) SetSshKeys(v []Sshkeypost)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *Systemuserput) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetState

`func (o *Systemuserput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Systemuserput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Systemuserput) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Systemuserput) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSudo

`func (o *Systemuserput) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *Systemuserput) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *Systemuserput) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *Systemuserput) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSuspended

`func (o *Systemuserput) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Systemuserput) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Systemuserput) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Systemuserput) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetTags

`func (o *Systemuserput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Systemuserput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Systemuserput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Systemuserput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnixGuid

`func (o *Systemuserput) GetUnixGuid() int32`

GetUnixGuid returns the UnixGuid field if non-nil, zero value otherwise.

### GetUnixGuidOk

`func (o *Systemuserput) GetUnixGuidOk() (*int32, bool)`

GetUnixGuidOk returns a tuple with the UnixGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixGuid

`func (o *Systemuserput) SetUnixGuid(v int32)`

SetUnixGuid sets UnixGuid field to given value.

### HasUnixGuid

`func (o *Systemuserput) HasUnixGuid() bool`

HasUnixGuid returns a boolean if a field has been set.

### GetUnixUid

`func (o *Systemuserput) GetUnixUid() int32`

GetUnixUid returns the UnixUid field if non-nil, zero value otherwise.

### GetUnixUidOk

`func (o *Systemuserput) GetUnixUidOk() (*int32, bool)`

GetUnixUidOk returns a tuple with the UnixUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUid

`func (o *Systemuserput) SetUnixUid(v int32)`

SetUnixUid sets UnixUid field to given value.

### HasUnixUid

`func (o *Systemuserput) HasUnixUid() bool`

HasUnixUid returns a boolean if a field has been set.

### GetUsername

`func (o *Systemuserput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Systemuserput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Systemuserput) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Systemuserput) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



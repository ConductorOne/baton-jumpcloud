# Systemuserreturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AccountLocked** | Pointer to **bool** |  | [optional] 
**AccountLockedDate** | Pointer to **NullableString** |  | [optional] 
**Activated** | Pointer to **bool** |  | [optional] 
**Addresses** | Pointer to [**[]SystemuserreturnAddressesInner**](SystemuserreturnAddressesInner.md) |  | [optional] 
**AllowPublicKey** | Pointer to **bool** |  | [optional] 
**AlternateEmail** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]SystemuserputAttributesInner**](SystemuserputAttributesInner.md) |  | [optional] 
**BadLoginAttempts** | Pointer to **int32** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**CostCenter** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**CreationSource** | Pointer to **string** |  | [optional] 
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
**MfaEnrollment** | Pointer to [**MfaEnrollment**](MfaEnrollment.md) |  | [optional] 
**Middlename** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**PasswordDate** | Pointer to **NullableString** |  | [optional] 
**PasswordExpirationDate** | Pointer to **NullableString** |  | [optional] 
**PasswordExpired** | Pointer to **bool** |  | [optional] 
**PasswordNeverExpires** | Pointer to **bool** |  | [optional] 
**PasswordlessSudo** | Pointer to **bool** |  | [optional] 
**PhoneNumbers** | Pointer to [**[]SystemuserreturnPhoneNumbersInner**](SystemuserreturnPhoneNumbersInner.md) |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**RecoveryEmail** | Pointer to [**SystemuserreturnRecoveryEmail**](SystemuserreturnRecoveryEmail.md) |  | [optional] 
**Relationships** | Pointer to [**[]SystemuserputRelationshipsInner**](SystemuserputRelationshipsInner.md) |  | [optional] 
**SambaServiceUser** | Pointer to **bool** |  | [optional] 
**SshKeys** | Pointer to [**[]Sshkeylist**](Sshkeylist.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Sudo** | Pointer to **bool** |  | [optional] 
**Suspended** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TotpEnabled** | Pointer to **bool** |  | [optional] 
**UnixGuid** | Pointer to **int32** |  | [optional] 
**UnixUid** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemuserreturn

`func NewSystemuserreturn() *Systemuserreturn`

NewSystemuserreturn instantiates a new Systemuserreturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemuserreturnWithDefaults

`func NewSystemuserreturnWithDefaults() *Systemuserreturn`

NewSystemuserreturnWithDefaults instantiates a new Systemuserreturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Systemuserreturn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Systemuserreturn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Systemuserreturn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Systemuserreturn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountLocked

`func (o *Systemuserreturn) GetAccountLocked() bool`

GetAccountLocked returns the AccountLocked field if non-nil, zero value otherwise.

### GetAccountLockedOk

`func (o *Systemuserreturn) GetAccountLockedOk() (*bool, bool)`

GetAccountLockedOk returns a tuple with the AccountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLocked

`func (o *Systemuserreturn) SetAccountLocked(v bool)`

SetAccountLocked sets AccountLocked field to given value.

### HasAccountLocked

`func (o *Systemuserreturn) HasAccountLocked() bool`

HasAccountLocked returns a boolean if a field has been set.

### GetAccountLockedDate

`func (o *Systemuserreturn) GetAccountLockedDate() string`

GetAccountLockedDate returns the AccountLockedDate field if non-nil, zero value otherwise.

### GetAccountLockedDateOk

`func (o *Systemuserreturn) GetAccountLockedDateOk() (*string, bool)`

GetAccountLockedDateOk returns a tuple with the AccountLockedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLockedDate

`func (o *Systemuserreturn) SetAccountLockedDate(v string)`

SetAccountLockedDate sets AccountLockedDate field to given value.

### HasAccountLockedDate

`func (o *Systemuserreturn) HasAccountLockedDate() bool`

HasAccountLockedDate returns a boolean if a field has been set.

### SetAccountLockedDateNil

`func (o *Systemuserreturn) SetAccountLockedDateNil(b bool)`

 SetAccountLockedDateNil sets the value for AccountLockedDate to be an explicit nil

### UnsetAccountLockedDate
`func (o *Systemuserreturn) UnsetAccountLockedDate()`

UnsetAccountLockedDate ensures that no value is present for AccountLockedDate, not even an explicit nil
### GetActivated

`func (o *Systemuserreturn) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *Systemuserreturn) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *Systemuserreturn) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *Systemuserreturn) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetAddresses

`func (o *Systemuserreturn) GetAddresses() []SystemuserreturnAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Systemuserreturn) GetAddressesOk() (*[]SystemuserreturnAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Systemuserreturn) SetAddresses(v []SystemuserreturnAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Systemuserreturn) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAllowPublicKey

`func (o *Systemuserreturn) GetAllowPublicKey() bool`

GetAllowPublicKey returns the AllowPublicKey field if non-nil, zero value otherwise.

### GetAllowPublicKeyOk

`func (o *Systemuserreturn) GetAllowPublicKeyOk() (*bool, bool)`

GetAllowPublicKeyOk returns a tuple with the AllowPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPublicKey

`func (o *Systemuserreturn) SetAllowPublicKey(v bool)`

SetAllowPublicKey sets AllowPublicKey field to given value.

### HasAllowPublicKey

`func (o *Systemuserreturn) HasAllowPublicKey() bool`

HasAllowPublicKey returns a boolean if a field has been set.

### GetAlternateEmail

`func (o *Systemuserreturn) GetAlternateEmail() string`

GetAlternateEmail returns the AlternateEmail field if non-nil, zero value otherwise.

### GetAlternateEmailOk

`func (o *Systemuserreturn) GetAlternateEmailOk() (*string, bool)`

GetAlternateEmailOk returns a tuple with the AlternateEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateEmail

`func (o *Systemuserreturn) SetAlternateEmail(v string)`

SetAlternateEmail sets AlternateEmail field to given value.

### HasAlternateEmail

`func (o *Systemuserreturn) HasAlternateEmail() bool`

HasAlternateEmail returns a boolean if a field has been set.

### GetAttributes

`func (o *Systemuserreturn) GetAttributes() []SystemuserputAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Systemuserreturn) GetAttributesOk() (*[]SystemuserputAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Systemuserreturn) SetAttributes(v []SystemuserputAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Systemuserreturn) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetBadLoginAttempts

`func (o *Systemuserreturn) GetBadLoginAttempts() int32`

GetBadLoginAttempts returns the BadLoginAttempts field if non-nil, zero value otherwise.

### GetBadLoginAttemptsOk

`func (o *Systemuserreturn) GetBadLoginAttemptsOk() (*int32, bool)`

GetBadLoginAttemptsOk returns a tuple with the BadLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadLoginAttempts

`func (o *Systemuserreturn) SetBadLoginAttempts(v int32)`

SetBadLoginAttempts sets BadLoginAttempts field to given value.

### HasBadLoginAttempts

`func (o *Systemuserreturn) HasBadLoginAttempts() bool`

HasBadLoginAttempts returns a boolean if a field has been set.

### GetCompany

`func (o *Systemuserreturn) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Systemuserreturn) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Systemuserreturn) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Systemuserreturn) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCostCenter

`func (o *Systemuserreturn) GetCostCenter() string`

GetCostCenter returns the CostCenter field if non-nil, zero value otherwise.

### GetCostCenterOk

`func (o *Systemuserreturn) GetCostCenterOk() (*string, bool)`

GetCostCenterOk returns a tuple with the CostCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCenter

`func (o *Systemuserreturn) SetCostCenter(v string)`

SetCostCenter sets CostCenter field to given value.

### HasCostCenter

`func (o *Systemuserreturn) HasCostCenter() bool`

HasCostCenter returns a boolean if a field has been set.

### GetCreated

`func (o *Systemuserreturn) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Systemuserreturn) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Systemuserreturn) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Systemuserreturn) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreationSource

`func (o *Systemuserreturn) GetCreationSource() string`

GetCreationSource returns the CreationSource field if non-nil, zero value otherwise.

### GetCreationSourceOk

`func (o *Systemuserreturn) GetCreationSourceOk() (*string, bool)`

GetCreationSourceOk returns a tuple with the CreationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationSource

`func (o *Systemuserreturn) SetCreationSource(v string)`

SetCreationSource sets CreationSource field to given value.

### HasCreationSource

`func (o *Systemuserreturn) HasCreationSource() bool`

HasCreationSource returns a boolean if a field has been set.

### GetDepartment

`func (o *Systemuserreturn) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Systemuserreturn) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Systemuserreturn) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Systemuserreturn) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDescription

`func (o *Systemuserreturn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Systemuserreturn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Systemuserreturn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Systemuserreturn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableDeviceMaxLoginAttempts

`func (o *Systemuserreturn) GetDisableDeviceMaxLoginAttempts() bool`

GetDisableDeviceMaxLoginAttempts returns the DisableDeviceMaxLoginAttempts field if non-nil, zero value otherwise.

### GetDisableDeviceMaxLoginAttemptsOk

`func (o *Systemuserreturn) GetDisableDeviceMaxLoginAttemptsOk() (*bool, bool)`

GetDisableDeviceMaxLoginAttemptsOk returns a tuple with the DisableDeviceMaxLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDeviceMaxLoginAttempts

`func (o *Systemuserreturn) SetDisableDeviceMaxLoginAttempts(v bool)`

SetDisableDeviceMaxLoginAttempts sets DisableDeviceMaxLoginAttempts field to given value.

### HasDisableDeviceMaxLoginAttempts

`func (o *Systemuserreturn) HasDisableDeviceMaxLoginAttempts() bool`

HasDisableDeviceMaxLoginAttempts returns a boolean if a field has been set.

### GetDisplayname

`func (o *Systemuserreturn) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *Systemuserreturn) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *Systemuserreturn) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *Systemuserreturn) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### GetEmail

`func (o *Systemuserreturn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Systemuserreturn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Systemuserreturn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Systemuserreturn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmployeeIdentifier

`func (o *Systemuserreturn) GetEmployeeIdentifier() string`

GetEmployeeIdentifier returns the EmployeeIdentifier field if non-nil, zero value otherwise.

### GetEmployeeIdentifierOk

`func (o *Systemuserreturn) GetEmployeeIdentifierOk() (*string, bool)`

GetEmployeeIdentifierOk returns a tuple with the EmployeeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIdentifier

`func (o *Systemuserreturn) SetEmployeeIdentifier(v string)`

SetEmployeeIdentifier sets EmployeeIdentifier field to given value.

### HasEmployeeIdentifier

`func (o *Systemuserreturn) HasEmployeeIdentifier() bool`

HasEmployeeIdentifier returns a boolean if a field has been set.

### GetEmployeeType

`func (o *Systemuserreturn) GetEmployeeType() string`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *Systemuserreturn) GetEmployeeTypeOk() (*string, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *Systemuserreturn) SetEmployeeType(v string)`

SetEmployeeType sets EmployeeType field to given value.

### HasEmployeeType

`func (o *Systemuserreturn) HasEmployeeType() bool`

HasEmployeeType returns a boolean if a field has been set.

### GetEnableManagedUid

`func (o *Systemuserreturn) GetEnableManagedUid() bool`

GetEnableManagedUid returns the EnableManagedUid field if non-nil, zero value otherwise.

### GetEnableManagedUidOk

`func (o *Systemuserreturn) GetEnableManagedUidOk() (*bool, bool)`

GetEnableManagedUidOk returns a tuple with the EnableManagedUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableManagedUid

`func (o *Systemuserreturn) SetEnableManagedUid(v bool)`

SetEnableManagedUid sets EnableManagedUid field to given value.

### HasEnableManagedUid

`func (o *Systemuserreturn) HasEnableManagedUid() bool`

HasEnableManagedUid returns a boolean if a field has been set.

### GetEnableUserPortalMultifactor

`func (o *Systemuserreturn) GetEnableUserPortalMultifactor() bool`

GetEnableUserPortalMultifactor returns the EnableUserPortalMultifactor field if non-nil, zero value otherwise.

### GetEnableUserPortalMultifactorOk

`func (o *Systemuserreturn) GetEnableUserPortalMultifactorOk() (*bool, bool)`

GetEnableUserPortalMultifactorOk returns a tuple with the EnableUserPortalMultifactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserPortalMultifactor

`func (o *Systemuserreturn) SetEnableUserPortalMultifactor(v bool)`

SetEnableUserPortalMultifactor sets EnableUserPortalMultifactor field to given value.

### HasEnableUserPortalMultifactor

`func (o *Systemuserreturn) HasEnableUserPortalMultifactor() bool`

HasEnableUserPortalMultifactor returns a boolean if a field has been set.

### GetExternalDn

`func (o *Systemuserreturn) GetExternalDn() string`

GetExternalDn returns the ExternalDn field if non-nil, zero value otherwise.

### GetExternalDnOk

`func (o *Systemuserreturn) GetExternalDnOk() (*string, bool)`

GetExternalDnOk returns a tuple with the ExternalDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDn

`func (o *Systemuserreturn) SetExternalDn(v string)`

SetExternalDn sets ExternalDn field to given value.

### HasExternalDn

`func (o *Systemuserreturn) HasExternalDn() bool`

HasExternalDn returns a boolean if a field has been set.

### GetExternalPasswordExpirationDate

`func (o *Systemuserreturn) GetExternalPasswordExpirationDate() string`

GetExternalPasswordExpirationDate returns the ExternalPasswordExpirationDate field if non-nil, zero value otherwise.

### GetExternalPasswordExpirationDateOk

`func (o *Systemuserreturn) GetExternalPasswordExpirationDateOk() (*string, bool)`

GetExternalPasswordExpirationDateOk returns a tuple with the ExternalPasswordExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPasswordExpirationDate

`func (o *Systemuserreturn) SetExternalPasswordExpirationDate(v string)`

SetExternalPasswordExpirationDate sets ExternalPasswordExpirationDate field to given value.

### HasExternalPasswordExpirationDate

`func (o *Systemuserreturn) HasExternalPasswordExpirationDate() bool`

HasExternalPasswordExpirationDate returns a boolean if a field has been set.

### GetExternalSourceType

`func (o *Systemuserreturn) GetExternalSourceType() string`

GetExternalSourceType returns the ExternalSourceType field if non-nil, zero value otherwise.

### GetExternalSourceTypeOk

`func (o *Systemuserreturn) GetExternalSourceTypeOk() (*string, bool)`

GetExternalSourceTypeOk returns a tuple with the ExternalSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSourceType

`func (o *Systemuserreturn) SetExternalSourceType(v string)`

SetExternalSourceType sets ExternalSourceType field to given value.

### HasExternalSourceType

`func (o *Systemuserreturn) HasExternalSourceType() bool`

HasExternalSourceType returns a boolean if a field has been set.

### GetExternallyManaged

`func (o *Systemuserreturn) GetExternallyManaged() bool`

GetExternallyManaged returns the ExternallyManaged field if non-nil, zero value otherwise.

### GetExternallyManagedOk

`func (o *Systemuserreturn) GetExternallyManagedOk() (*bool, bool)`

GetExternallyManagedOk returns a tuple with the ExternallyManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyManaged

`func (o *Systemuserreturn) SetExternallyManaged(v bool)`

SetExternallyManaged sets ExternallyManaged field to given value.

### HasExternallyManaged

`func (o *Systemuserreturn) HasExternallyManaged() bool`

HasExternallyManaged returns a boolean if a field has been set.

### GetFirstname

`func (o *Systemuserreturn) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *Systemuserreturn) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *Systemuserreturn) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *Systemuserreturn) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetJobTitle

`func (o *Systemuserreturn) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Systemuserreturn) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Systemuserreturn) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Systemuserreturn) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastname

`func (o *Systemuserreturn) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *Systemuserreturn) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *Systemuserreturn) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *Systemuserreturn) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetLdapBindingUser

`func (o *Systemuserreturn) GetLdapBindingUser() bool`

GetLdapBindingUser returns the LdapBindingUser field if non-nil, zero value otherwise.

### GetLdapBindingUserOk

`func (o *Systemuserreturn) GetLdapBindingUserOk() (*bool, bool)`

GetLdapBindingUserOk returns a tuple with the LdapBindingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindingUser

`func (o *Systemuserreturn) SetLdapBindingUser(v bool)`

SetLdapBindingUser sets LdapBindingUser field to given value.

### HasLdapBindingUser

`func (o *Systemuserreturn) HasLdapBindingUser() bool`

HasLdapBindingUser returns a boolean if a field has been set.

### GetLocation

`func (o *Systemuserreturn) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Systemuserreturn) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Systemuserreturn) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Systemuserreturn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetManagedAppleId

`func (o *Systemuserreturn) GetManagedAppleId() string`

GetManagedAppleId returns the ManagedAppleId field if non-nil, zero value otherwise.

### GetManagedAppleIdOk

`func (o *Systemuserreturn) GetManagedAppleIdOk() (*string, bool)`

GetManagedAppleIdOk returns a tuple with the ManagedAppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAppleId

`func (o *Systemuserreturn) SetManagedAppleId(v string)`

SetManagedAppleId sets ManagedAppleId field to given value.

### HasManagedAppleId

`func (o *Systemuserreturn) HasManagedAppleId() bool`

HasManagedAppleId returns a boolean if a field has been set.

### GetManager

`func (o *Systemuserreturn) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Systemuserreturn) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Systemuserreturn) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *Systemuserreturn) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetMfa

`func (o *Systemuserreturn) GetMfa() Mfa`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *Systemuserreturn) GetMfaOk() (*Mfa, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *Systemuserreturn) SetMfa(v Mfa)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *Systemuserreturn) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetMfaEnrollment

`func (o *Systemuserreturn) GetMfaEnrollment() MfaEnrollment`

GetMfaEnrollment returns the MfaEnrollment field if non-nil, zero value otherwise.

### GetMfaEnrollmentOk

`func (o *Systemuserreturn) GetMfaEnrollmentOk() (*MfaEnrollment, bool)`

GetMfaEnrollmentOk returns a tuple with the MfaEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnrollment

`func (o *Systemuserreturn) SetMfaEnrollment(v MfaEnrollment)`

SetMfaEnrollment sets MfaEnrollment field to given value.

### HasMfaEnrollment

`func (o *Systemuserreturn) HasMfaEnrollment() bool`

HasMfaEnrollment returns a boolean if a field has been set.

### GetMiddlename

`func (o *Systemuserreturn) GetMiddlename() string`

GetMiddlename returns the Middlename field if non-nil, zero value otherwise.

### GetMiddlenameOk

`func (o *Systemuserreturn) GetMiddlenameOk() (*string, bool)`

GetMiddlenameOk returns a tuple with the Middlename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddlename

`func (o *Systemuserreturn) SetMiddlename(v string)`

SetMiddlename sets Middlename field to given value.

### HasMiddlename

`func (o *Systemuserreturn) HasMiddlename() bool`

HasMiddlename returns a boolean if a field has been set.

### GetOrganization

`func (o *Systemuserreturn) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Systemuserreturn) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Systemuserreturn) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Systemuserreturn) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPasswordDate

`func (o *Systemuserreturn) GetPasswordDate() string`

GetPasswordDate returns the PasswordDate field if non-nil, zero value otherwise.

### GetPasswordDateOk

`func (o *Systemuserreturn) GetPasswordDateOk() (*string, bool)`

GetPasswordDateOk returns a tuple with the PasswordDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordDate

`func (o *Systemuserreturn) SetPasswordDate(v string)`

SetPasswordDate sets PasswordDate field to given value.

### HasPasswordDate

`func (o *Systemuserreturn) HasPasswordDate() bool`

HasPasswordDate returns a boolean if a field has been set.

### SetPasswordDateNil

`func (o *Systemuserreturn) SetPasswordDateNil(b bool)`

 SetPasswordDateNil sets the value for PasswordDate to be an explicit nil

### UnsetPasswordDate
`func (o *Systemuserreturn) UnsetPasswordDate()`

UnsetPasswordDate ensures that no value is present for PasswordDate, not even an explicit nil
### GetPasswordExpirationDate

`func (o *Systemuserreturn) GetPasswordExpirationDate() string`

GetPasswordExpirationDate returns the PasswordExpirationDate field if non-nil, zero value otherwise.

### GetPasswordExpirationDateOk

`func (o *Systemuserreturn) GetPasswordExpirationDateOk() (*string, bool)`

GetPasswordExpirationDateOk returns a tuple with the PasswordExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationDate

`func (o *Systemuserreturn) SetPasswordExpirationDate(v string)`

SetPasswordExpirationDate sets PasswordExpirationDate field to given value.

### HasPasswordExpirationDate

`func (o *Systemuserreturn) HasPasswordExpirationDate() bool`

HasPasswordExpirationDate returns a boolean if a field has been set.

### SetPasswordExpirationDateNil

`func (o *Systemuserreturn) SetPasswordExpirationDateNil(b bool)`

 SetPasswordExpirationDateNil sets the value for PasswordExpirationDate to be an explicit nil

### UnsetPasswordExpirationDate
`func (o *Systemuserreturn) UnsetPasswordExpirationDate()`

UnsetPasswordExpirationDate ensures that no value is present for PasswordExpirationDate, not even an explicit nil
### GetPasswordExpired

`func (o *Systemuserreturn) GetPasswordExpired() bool`

GetPasswordExpired returns the PasswordExpired field if non-nil, zero value otherwise.

### GetPasswordExpiredOk

`func (o *Systemuserreturn) GetPasswordExpiredOk() (*bool, bool)`

GetPasswordExpiredOk returns a tuple with the PasswordExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpired

`func (o *Systemuserreturn) SetPasswordExpired(v bool)`

SetPasswordExpired sets PasswordExpired field to given value.

### HasPasswordExpired

`func (o *Systemuserreturn) HasPasswordExpired() bool`

HasPasswordExpired returns a boolean if a field has been set.

### GetPasswordNeverExpires

`func (o *Systemuserreturn) GetPasswordNeverExpires() bool`

GetPasswordNeverExpires returns the PasswordNeverExpires field if non-nil, zero value otherwise.

### GetPasswordNeverExpiresOk

`func (o *Systemuserreturn) GetPasswordNeverExpiresOk() (*bool, bool)`

GetPasswordNeverExpiresOk returns a tuple with the PasswordNeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordNeverExpires

`func (o *Systemuserreturn) SetPasswordNeverExpires(v bool)`

SetPasswordNeverExpires sets PasswordNeverExpires field to given value.

### HasPasswordNeverExpires

`func (o *Systemuserreturn) HasPasswordNeverExpires() bool`

HasPasswordNeverExpires returns a boolean if a field has been set.

### GetPasswordlessSudo

`func (o *Systemuserreturn) GetPasswordlessSudo() bool`

GetPasswordlessSudo returns the PasswordlessSudo field if non-nil, zero value otherwise.

### GetPasswordlessSudoOk

`func (o *Systemuserreturn) GetPasswordlessSudoOk() (*bool, bool)`

GetPasswordlessSudoOk returns a tuple with the PasswordlessSudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessSudo

`func (o *Systemuserreturn) SetPasswordlessSudo(v bool)`

SetPasswordlessSudo sets PasswordlessSudo field to given value.

### HasPasswordlessSudo

`func (o *Systemuserreturn) HasPasswordlessSudo() bool`

HasPasswordlessSudo returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *Systemuserreturn) GetPhoneNumbers() []SystemuserreturnPhoneNumbersInner`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *Systemuserreturn) GetPhoneNumbersOk() (*[]SystemuserreturnPhoneNumbersInner, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *Systemuserreturn) SetPhoneNumbers(v []SystemuserreturnPhoneNumbersInner)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *Systemuserreturn) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetPublicKey

`func (o *Systemuserreturn) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Systemuserreturn) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Systemuserreturn) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *Systemuserreturn) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRecoveryEmail

`func (o *Systemuserreturn) GetRecoveryEmail() SystemuserreturnRecoveryEmail`

GetRecoveryEmail returns the RecoveryEmail field if non-nil, zero value otherwise.

### GetRecoveryEmailOk

`func (o *Systemuserreturn) GetRecoveryEmailOk() (*SystemuserreturnRecoveryEmail, bool)`

GetRecoveryEmailOk returns a tuple with the RecoveryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryEmail

`func (o *Systemuserreturn) SetRecoveryEmail(v SystemuserreturnRecoveryEmail)`

SetRecoveryEmail sets RecoveryEmail field to given value.

### HasRecoveryEmail

`func (o *Systemuserreturn) HasRecoveryEmail() bool`

HasRecoveryEmail returns a boolean if a field has been set.

### GetRelationships

`func (o *Systemuserreturn) GetRelationships() []SystemuserputRelationshipsInner`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Systemuserreturn) GetRelationshipsOk() (*[]SystemuserputRelationshipsInner, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Systemuserreturn) SetRelationships(v []SystemuserputRelationshipsInner)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Systemuserreturn) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetSambaServiceUser

`func (o *Systemuserreturn) GetSambaServiceUser() bool`

GetSambaServiceUser returns the SambaServiceUser field if non-nil, zero value otherwise.

### GetSambaServiceUserOk

`func (o *Systemuserreturn) GetSambaServiceUserOk() (*bool, bool)`

GetSambaServiceUserOk returns a tuple with the SambaServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSambaServiceUser

`func (o *Systemuserreturn) SetSambaServiceUser(v bool)`

SetSambaServiceUser sets SambaServiceUser field to given value.

### HasSambaServiceUser

`func (o *Systemuserreturn) HasSambaServiceUser() bool`

HasSambaServiceUser returns a boolean if a field has been set.

### GetSshKeys

`func (o *Systemuserreturn) GetSshKeys() []Sshkeylist`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *Systemuserreturn) GetSshKeysOk() (*[]Sshkeylist, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *Systemuserreturn) SetSshKeys(v []Sshkeylist)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *Systemuserreturn) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetState

`func (o *Systemuserreturn) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Systemuserreturn) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Systemuserreturn) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Systemuserreturn) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSudo

`func (o *Systemuserreturn) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *Systemuserreturn) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *Systemuserreturn) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *Systemuserreturn) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSuspended

`func (o *Systemuserreturn) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Systemuserreturn) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Systemuserreturn) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Systemuserreturn) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetTags

`func (o *Systemuserreturn) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Systemuserreturn) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Systemuserreturn) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Systemuserreturn) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTotpEnabled

`func (o *Systemuserreturn) GetTotpEnabled() bool`

GetTotpEnabled returns the TotpEnabled field if non-nil, zero value otherwise.

### GetTotpEnabledOk

`func (o *Systemuserreturn) GetTotpEnabledOk() (*bool, bool)`

GetTotpEnabledOk returns a tuple with the TotpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpEnabled

`func (o *Systemuserreturn) SetTotpEnabled(v bool)`

SetTotpEnabled sets TotpEnabled field to given value.

### HasTotpEnabled

`func (o *Systemuserreturn) HasTotpEnabled() bool`

HasTotpEnabled returns a boolean if a field has been set.

### GetUnixGuid

`func (o *Systemuserreturn) GetUnixGuid() int32`

GetUnixGuid returns the UnixGuid field if non-nil, zero value otherwise.

### GetUnixGuidOk

`func (o *Systemuserreturn) GetUnixGuidOk() (*int32, bool)`

GetUnixGuidOk returns a tuple with the UnixGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixGuid

`func (o *Systemuserreturn) SetUnixGuid(v int32)`

SetUnixGuid sets UnixGuid field to given value.

### HasUnixGuid

`func (o *Systemuserreturn) HasUnixGuid() bool`

HasUnixGuid returns a boolean if a field has been set.

### GetUnixUid

`func (o *Systemuserreturn) GetUnixUid() int32`

GetUnixUid returns the UnixUid field if non-nil, zero value otherwise.

### GetUnixUidOk

`func (o *Systemuserreturn) GetUnixUidOk() (*int32, bool)`

GetUnixUidOk returns a tuple with the UnixUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixUid

`func (o *Systemuserreturn) SetUnixUid(v int32)`

SetUnixUid sets UnixUid field to given value.

### HasUnixUid

`func (o *Systemuserreturn) HasUnixUid() bool`

HasUnixUid returns a boolean if a field has been set.

### GetUsername

`func (o *Systemuserreturn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Systemuserreturn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Systemuserreturn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Systemuserreturn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



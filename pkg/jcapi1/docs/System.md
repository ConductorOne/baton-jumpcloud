# System

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**AllowMultiFactorAuthentication** | Pointer to **bool** |  | [optional] 
**AllowPublicKeyAuthentication** | Pointer to **bool** |  | [optional] 
**AllowSshPasswordAuthentication** | Pointer to **bool** |  | [optional] 
**AllowSshRootLogin** | Pointer to **bool** |  | [optional] 
**AmazonInstanceID** | Pointer to **string** |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**AzureAdJoined** | Pointer to **bool** |  | [optional] 
**BuiltInCommands** | Pointer to [**[]SystemBuiltInCommandsInner**](SystemBuiltInCommandsInner.md) |  | [optional] [readonly] 
**ConnectionHistory** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DesktopCapable** | Pointer to **bool** |  | [optional] 
**DisplayManager** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DomainInfo** | Pointer to [**SystemDomainInfo**](SystemDomainInfo.md) |  | [optional] 
**Fde** | Pointer to [**Fde**](Fde.md) |  | [optional] 
**FileSystem** | Pointer to **NullableString** |  | [optional] 
**HasServiceAccount** | Pointer to **bool** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**LastContact** | Pointer to **NullableTime** |  | [optional] 
**Mdm** | Pointer to [**SystemMdm**](SystemMdm.md) |  | [optional] 
**ModifySSHDConfig** | Pointer to **bool** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]SystemNetworkInterfacesInner**](SystemNetworkInterfacesInner.md) |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**OsFamily** | Pointer to **string** |  | [optional] 
**OsVersionDetail** | Pointer to [**SystemOsVersionDetail**](SystemOsVersionDetail.md) |  | [optional] 
**ProvisionMetadata** | Pointer to [**SystemProvisionMetadata**](SystemProvisionMetadata.md) |  | [optional] 
**RemoteIP** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**ServiceAccountState** | Pointer to [**SystemServiceAccountState**](SystemServiceAccountState.md) |  | [optional] 
**SshRootEnabled** | Pointer to **bool** |  | [optional] 
**SshdParams** | Pointer to [**[]SystemSshdParamsInner**](SystemSshdParamsInner.md) |  | [optional] 
**SystemInsights** | Pointer to [**SystemSystemInsights**](SystemSystemInsights.md) |  | [optional] 
**SystemTimezone** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TemplateName** | Pointer to **string** |  | [optional] 
**UserMetrics** | Pointer to [**[]SystemUserMetricsInner**](SystemUserMetricsInner.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewSystem

`func NewSystem() *System`

NewSystem instantiates a new System object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemWithDefaults

`func NewSystemWithDefaults() *System`

NewSystemWithDefaults instantiates a new System object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *System) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *System) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *System) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *System) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *System) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *System) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *System) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *System) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAgentVersion

`func (o *System) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *System) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *System) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *System) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetAllowMultiFactorAuthentication

`func (o *System) GetAllowMultiFactorAuthentication() bool`

GetAllowMultiFactorAuthentication returns the AllowMultiFactorAuthentication field if non-nil, zero value otherwise.

### GetAllowMultiFactorAuthenticationOk

`func (o *System) GetAllowMultiFactorAuthenticationOk() (*bool, bool)`

GetAllowMultiFactorAuthenticationOk returns a tuple with the AllowMultiFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiFactorAuthentication

`func (o *System) SetAllowMultiFactorAuthentication(v bool)`

SetAllowMultiFactorAuthentication sets AllowMultiFactorAuthentication field to given value.

### HasAllowMultiFactorAuthentication

`func (o *System) HasAllowMultiFactorAuthentication() bool`

HasAllowMultiFactorAuthentication returns a boolean if a field has been set.

### GetAllowPublicKeyAuthentication

`func (o *System) GetAllowPublicKeyAuthentication() bool`

GetAllowPublicKeyAuthentication returns the AllowPublicKeyAuthentication field if non-nil, zero value otherwise.

### GetAllowPublicKeyAuthenticationOk

`func (o *System) GetAllowPublicKeyAuthenticationOk() (*bool, bool)`

GetAllowPublicKeyAuthenticationOk returns a tuple with the AllowPublicKeyAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPublicKeyAuthentication

`func (o *System) SetAllowPublicKeyAuthentication(v bool)`

SetAllowPublicKeyAuthentication sets AllowPublicKeyAuthentication field to given value.

### HasAllowPublicKeyAuthentication

`func (o *System) HasAllowPublicKeyAuthentication() bool`

HasAllowPublicKeyAuthentication returns a boolean if a field has been set.

### GetAllowSshPasswordAuthentication

`func (o *System) GetAllowSshPasswordAuthentication() bool`

GetAllowSshPasswordAuthentication returns the AllowSshPasswordAuthentication field if non-nil, zero value otherwise.

### GetAllowSshPasswordAuthenticationOk

`func (o *System) GetAllowSshPasswordAuthenticationOk() (*bool, bool)`

GetAllowSshPasswordAuthenticationOk returns a tuple with the AllowSshPasswordAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSshPasswordAuthentication

`func (o *System) SetAllowSshPasswordAuthentication(v bool)`

SetAllowSshPasswordAuthentication sets AllowSshPasswordAuthentication field to given value.

### HasAllowSshPasswordAuthentication

`func (o *System) HasAllowSshPasswordAuthentication() bool`

HasAllowSshPasswordAuthentication returns a boolean if a field has been set.

### GetAllowSshRootLogin

`func (o *System) GetAllowSshRootLogin() bool`

GetAllowSshRootLogin returns the AllowSshRootLogin field if non-nil, zero value otherwise.

### GetAllowSshRootLoginOk

`func (o *System) GetAllowSshRootLoginOk() (*bool, bool)`

GetAllowSshRootLoginOk returns a tuple with the AllowSshRootLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSshRootLogin

`func (o *System) SetAllowSshRootLogin(v bool)`

SetAllowSshRootLogin sets AllowSshRootLogin field to given value.

### HasAllowSshRootLogin

`func (o *System) HasAllowSshRootLogin() bool`

HasAllowSshRootLogin returns a boolean if a field has been set.

### GetAmazonInstanceID

`func (o *System) GetAmazonInstanceID() string`

GetAmazonInstanceID returns the AmazonInstanceID field if non-nil, zero value otherwise.

### GetAmazonInstanceIDOk

`func (o *System) GetAmazonInstanceIDOk() (*string, bool)`

GetAmazonInstanceIDOk returns a tuple with the AmazonInstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonInstanceID

`func (o *System) SetAmazonInstanceID(v string)`

SetAmazonInstanceID sets AmazonInstanceID field to given value.

### HasAmazonInstanceID

`func (o *System) HasAmazonInstanceID() bool`

HasAmazonInstanceID returns a boolean if a field has been set.

### GetArch

`func (o *System) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *System) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *System) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *System) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetAzureAdJoined

`func (o *System) GetAzureAdJoined() bool`

GetAzureAdJoined returns the AzureAdJoined field if non-nil, zero value otherwise.

### GetAzureAdJoinedOk

`func (o *System) GetAzureAdJoinedOk() (*bool, bool)`

GetAzureAdJoinedOk returns a tuple with the AzureAdJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdJoined

`func (o *System) SetAzureAdJoined(v bool)`

SetAzureAdJoined sets AzureAdJoined field to given value.

### HasAzureAdJoined

`func (o *System) HasAzureAdJoined() bool`

HasAzureAdJoined returns a boolean if a field has been set.

### GetBuiltInCommands

`func (o *System) GetBuiltInCommands() []SystemBuiltInCommandsInner`

GetBuiltInCommands returns the BuiltInCommands field if non-nil, zero value otherwise.

### GetBuiltInCommandsOk

`func (o *System) GetBuiltInCommandsOk() (*[]SystemBuiltInCommandsInner, bool)`

GetBuiltInCommandsOk returns a tuple with the BuiltInCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInCommands

`func (o *System) SetBuiltInCommands(v []SystemBuiltInCommandsInner)`

SetBuiltInCommands sets BuiltInCommands field to given value.

### HasBuiltInCommands

`func (o *System) HasBuiltInCommands() bool`

HasBuiltInCommands returns a boolean if a field has been set.

### GetConnectionHistory

`func (o *System) GetConnectionHistory() []map[string]interface{}`

GetConnectionHistory returns the ConnectionHistory field if non-nil, zero value otherwise.

### GetConnectionHistoryOk

`func (o *System) GetConnectionHistoryOk() (*[]map[string]interface{}, bool)`

GetConnectionHistoryOk returns a tuple with the ConnectionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionHistory

`func (o *System) SetConnectionHistory(v []map[string]interface{})`

SetConnectionHistory sets ConnectionHistory field to given value.

### HasConnectionHistory

`func (o *System) HasConnectionHistory() bool`

HasConnectionHistory returns a boolean if a field has been set.

### GetCreated

`func (o *System) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *System) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *System) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *System) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *System) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *System) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *System) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *System) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesktopCapable

`func (o *System) GetDesktopCapable() bool`

GetDesktopCapable returns the DesktopCapable field if non-nil, zero value otherwise.

### GetDesktopCapableOk

`func (o *System) GetDesktopCapableOk() (*bool, bool)`

GetDesktopCapableOk returns a tuple with the DesktopCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopCapable

`func (o *System) SetDesktopCapable(v bool)`

SetDesktopCapable sets DesktopCapable field to given value.

### HasDesktopCapable

`func (o *System) HasDesktopCapable() bool`

HasDesktopCapable returns a boolean if a field has been set.

### GetDisplayManager

`func (o *System) GetDisplayManager() string`

GetDisplayManager returns the DisplayManager field if non-nil, zero value otherwise.

### GetDisplayManagerOk

`func (o *System) GetDisplayManagerOk() (*string, bool)`

GetDisplayManagerOk returns a tuple with the DisplayManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayManager

`func (o *System) SetDisplayManager(v string)`

SetDisplayManager sets DisplayManager field to given value.

### HasDisplayManager

`func (o *System) HasDisplayManager() bool`

HasDisplayManager returns a boolean if a field has been set.

### GetDisplayName

`func (o *System) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *System) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *System) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *System) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomainInfo

`func (o *System) GetDomainInfo() SystemDomainInfo`

GetDomainInfo returns the DomainInfo field if non-nil, zero value otherwise.

### GetDomainInfoOk

`func (o *System) GetDomainInfoOk() (*SystemDomainInfo, bool)`

GetDomainInfoOk returns a tuple with the DomainInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainInfo

`func (o *System) SetDomainInfo(v SystemDomainInfo)`

SetDomainInfo sets DomainInfo field to given value.

### HasDomainInfo

`func (o *System) HasDomainInfo() bool`

HasDomainInfo returns a boolean if a field has been set.

### GetFde

`func (o *System) GetFde() Fde`

GetFde returns the Fde field if non-nil, zero value otherwise.

### GetFdeOk

`func (o *System) GetFdeOk() (*Fde, bool)`

GetFdeOk returns a tuple with the Fde field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFde

`func (o *System) SetFde(v Fde)`

SetFde sets Fde field to given value.

### HasFde

`func (o *System) HasFde() bool`

HasFde returns a boolean if a field has been set.

### GetFileSystem

`func (o *System) GetFileSystem() string`

GetFileSystem returns the FileSystem field if non-nil, zero value otherwise.

### GetFileSystemOk

`func (o *System) GetFileSystemOk() (*string, bool)`

GetFileSystemOk returns a tuple with the FileSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystem

`func (o *System) SetFileSystem(v string)`

SetFileSystem sets FileSystem field to given value.

### HasFileSystem

`func (o *System) HasFileSystem() bool`

HasFileSystem returns a boolean if a field has been set.

### SetFileSystemNil

`func (o *System) SetFileSystemNil(b bool)`

 SetFileSystemNil sets the value for FileSystem to be an explicit nil

### UnsetFileSystem
`func (o *System) UnsetFileSystem()`

UnsetFileSystem ensures that no value is present for FileSystem, not even an explicit nil
### GetHasServiceAccount

`func (o *System) GetHasServiceAccount() bool`

GetHasServiceAccount returns the HasServiceAccount field if non-nil, zero value otherwise.

### GetHasServiceAccountOk

`func (o *System) GetHasServiceAccountOk() (*bool, bool)`

GetHasServiceAccountOk returns a tuple with the HasServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasServiceAccount

`func (o *System) SetHasServiceAccount(v bool)`

SetHasServiceAccount sets HasServiceAccount field to given value.

### HasHasServiceAccount

`func (o *System) HasHasServiceAccount() bool`

HasHasServiceAccount returns a boolean if a field has been set.

### GetHostname

`func (o *System) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *System) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *System) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *System) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLastContact

`func (o *System) GetLastContact() time.Time`

GetLastContact returns the LastContact field if non-nil, zero value otherwise.

### GetLastContactOk

`func (o *System) GetLastContactOk() (*time.Time, bool)`

GetLastContactOk returns a tuple with the LastContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContact

`func (o *System) SetLastContact(v time.Time)`

SetLastContact sets LastContact field to given value.

### HasLastContact

`func (o *System) HasLastContact() bool`

HasLastContact returns a boolean if a field has been set.

### SetLastContactNil

`func (o *System) SetLastContactNil(b bool)`

 SetLastContactNil sets the value for LastContact to be an explicit nil

### UnsetLastContact
`func (o *System) UnsetLastContact()`

UnsetLastContact ensures that no value is present for LastContact, not even an explicit nil
### GetMdm

`func (o *System) GetMdm() SystemMdm`

GetMdm returns the Mdm field if non-nil, zero value otherwise.

### GetMdmOk

`func (o *System) GetMdmOk() (*SystemMdm, bool)`

GetMdmOk returns a tuple with the Mdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdm

`func (o *System) SetMdm(v SystemMdm)`

SetMdm sets Mdm field to given value.

### HasMdm

`func (o *System) HasMdm() bool`

HasMdm returns a boolean if a field has been set.

### GetModifySSHDConfig

`func (o *System) GetModifySSHDConfig() bool`

GetModifySSHDConfig returns the ModifySSHDConfig field if non-nil, zero value otherwise.

### GetModifySSHDConfigOk

`func (o *System) GetModifySSHDConfigOk() (*bool, bool)`

GetModifySSHDConfigOk returns a tuple with the ModifySSHDConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifySSHDConfig

`func (o *System) SetModifySSHDConfig(v bool)`

SetModifySSHDConfig sets ModifySSHDConfig field to given value.

### HasModifySSHDConfig

`func (o *System) HasModifySSHDConfig() bool`

HasModifySSHDConfig returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *System) GetNetworkInterfaces() []SystemNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *System) GetNetworkInterfacesOk() (*[]SystemNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *System) SetNetworkInterfaces(v []SystemNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *System) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetOrganization

`func (o *System) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *System) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *System) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *System) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOs

`func (o *System) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *System) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *System) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *System) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOsFamily

`func (o *System) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *System) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *System) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *System) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetOsVersionDetail

`func (o *System) GetOsVersionDetail() SystemOsVersionDetail`

GetOsVersionDetail returns the OsVersionDetail field if non-nil, zero value otherwise.

### GetOsVersionDetailOk

`func (o *System) GetOsVersionDetailOk() (*SystemOsVersionDetail, bool)`

GetOsVersionDetailOk returns a tuple with the OsVersionDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersionDetail

`func (o *System) SetOsVersionDetail(v SystemOsVersionDetail)`

SetOsVersionDetail sets OsVersionDetail field to given value.

### HasOsVersionDetail

`func (o *System) HasOsVersionDetail() bool`

HasOsVersionDetail returns a boolean if a field has been set.

### GetProvisionMetadata

`func (o *System) GetProvisionMetadata() SystemProvisionMetadata`

GetProvisionMetadata returns the ProvisionMetadata field if non-nil, zero value otherwise.

### GetProvisionMetadataOk

`func (o *System) GetProvisionMetadataOk() (*SystemProvisionMetadata, bool)`

GetProvisionMetadataOk returns a tuple with the ProvisionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionMetadata

`func (o *System) SetProvisionMetadata(v SystemProvisionMetadata)`

SetProvisionMetadata sets ProvisionMetadata field to given value.

### HasProvisionMetadata

`func (o *System) HasProvisionMetadata() bool`

HasProvisionMetadata returns a boolean if a field has been set.

### GetRemoteIP

`func (o *System) GetRemoteIP() string`

GetRemoteIP returns the RemoteIP field if non-nil, zero value otherwise.

### GetRemoteIPOk

`func (o *System) GetRemoteIPOk() (*string, bool)`

GetRemoteIPOk returns a tuple with the RemoteIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIP

`func (o *System) SetRemoteIP(v string)`

SetRemoteIP sets RemoteIP field to given value.

### HasRemoteIP

`func (o *System) HasRemoteIP() bool`

HasRemoteIP returns a boolean if a field has been set.

### GetSerialNumber

`func (o *System) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *System) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *System) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *System) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetServiceAccountState

`func (o *System) GetServiceAccountState() SystemServiceAccountState`

GetServiceAccountState returns the ServiceAccountState field if non-nil, zero value otherwise.

### GetServiceAccountStateOk

`func (o *System) GetServiceAccountStateOk() (*SystemServiceAccountState, bool)`

GetServiceAccountStateOk returns a tuple with the ServiceAccountState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountState

`func (o *System) SetServiceAccountState(v SystemServiceAccountState)`

SetServiceAccountState sets ServiceAccountState field to given value.

### HasServiceAccountState

`func (o *System) HasServiceAccountState() bool`

HasServiceAccountState returns a boolean if a field has been set.

### GetSshRootEnabled

`func (o *System) GetSshRootEnabled() bool`

GetSshRootEnabled returns the SshRootEnabled field if non-nil, zero value otherwise.

### GetSshRootEnabledOk

`func (o *System) GetSshRootEnabledOk() (*bool, bool)`

GetSshRootEnabledOk returns a tuple with the SshRootEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshRootEnabled

`func (o *System) SetSshRootEnabled(v bool)`

SetSshRootEnabled sets SshRootEnabled field to given value.

### HasSshRootEnabled

`func (o *System) HasSshRootEnabled() bool`

HasSshRootEnabled returns a boolean if a field has been set.

### GetSshdParams

`func (o *System) GetSshdParams() []SystemSshdParamsInner`

GetSshdParams returns the SshdParams field if non-nil, zero value otherwise.

### GetSshdParamsOk

`func (o *System) GetSshdParamsOk() (*[]SystemSshdParamsInner, bool)`

GetSshdParamsOk returns a tuple with the SshdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshdParams

`func (o *System) SetSshdParams(v []SystemSshdParamsInner)`

SetSshdParams sets SshdParams field to given value.

### HasSshdParams

`func (o *System) HasSshdParams() bool`

HasSshdParams returns a boolean if a field has been set.

### GetSystemInsights

`func (o *System) GetSystemInsights() SystemSystemInsights`

GetSystemInsights returns the SystemInsights field if non-nil, zero value otherwise.

### GetSystemInsightsOk

`func (o *System) GetSystemInsightsOk() (*SystemSystemInsights, bool)`

GetSystemInsightsOk returns a tuple with the SystemInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInsights

`func (o *System) SetSystemInsights(v SystemSystemInsights)`

SetSystemInsights sets SystemInsights field to given value.

### HasSystemInsights

`func (o *System) HasSystemInsights() bool`

HasSystemInsights returns a boolean if a field has been set.

### GetSystemTimezone

`func (o *System) GetSystemTimezone() int32`

GetSystemTimezone returns the SystemTimezone field if non-nil, zero value otherwise.

### GetSystemTimezoneOk

`func (o *System) GetSystemTimezoneOk() (*int32, bool)`

GetSystemTimezoneOk returns a tuple with the SystemTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTimezone

`func (o *System) SetSystemTimezone(v int32)`

SetSystemTimezone sets SystemTimezone field to given value.

### HasSystemTimezone

`func (o *System) HasSystemTimezone() bool`

HasSystemTimezone returns a boolean if a field has been set.

### GetTags

`func (o *System) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *System) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *System) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *System) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplateName

`func (o *System) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *System) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *System) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *System) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetUserMetrics

`func (o *System) GetUserMetrics() []SystemUserMetricsInner`

GetUserMetrics returns the UserMetrics field if non-nil, zero value otherwise.

### GetUserMetricsOk

`func (o *System) GetUserMetricsOk() (*[]SystemUserMetricsInner, bool)`

GetUserMetricsOk returns a tuple with the UserMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetrics

`func (o *System) SetUserMetrics(v []SystemUserMetricsInner)`

SetUserMetrics sets UserMetrics field to given value.

### HasUserMetrics

`func (o *System) HasUserMetrics() bool`

HasUserMetrics returns a boolean if a field has been set.

### GetVersion

`func (o *System) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *System) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *System) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *System) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



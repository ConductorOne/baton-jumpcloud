# SystemInsightsInterfaceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collisions** | Pointer to **string** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**ConnectionStatus** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DhcpEnabled** | Pointer to **int32** |  | [optional] 
**DhcpLeaseExpires** | Pointer to **string** |  | [optional] 
**DhcpLeaseObtained** | Pointer to **string** |  | [optional] 
**DhcpServer** | Pointer to **string** |  | [optional] 
**DnsDomain** | Pointer to **string** |  | [optional] 
**DnsDomainSuffixSearchOrder** | Pointer to **string** |  | [optional] 
**DnsHostName** | Pointer to **string** |  | [optional] 
**DnsServerSearchOrder** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **int32** |  | [optional] 
**Flags** | Pointer to **int32** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Ibytes** | Pointer to **string** |  | [optional] 
**Idrops** | Pointer to **string** |  | [optional] 
**Ierrors** | Pointer to **string** |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**Ipackets** | Pointer to **string** |  | [optional] 
**LastChange** | Pointer to **string** |  | [optional] 
**LinkSpeed** | Pointer to **string** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **int32** |  | [optional] 
**Mtu** | Pointer to **int32** |  | [optional] 
**Obytes** | Pointer to **string** |  | [optional] 
**Odrops** | Pointer to **string** |  | [optional] 
**Oerrors** | Pointer to **string** |  | [optional] 
**Opackets** | Pointer to **string** |  | [optional] 
**PciSlot** | Pointer to **string** |  | [optional] 
**PhysicalAdapter** | Pointer to **int32** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**SystemId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewSystemInsightsInterfaceDetails

`func NewSystemInsightsInterfaceDetails() *SystemInsightsInterfaceDetails`

NewSystemInsightsInterfaceDetails instantiates a new SystemInsightsInterfaceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInsightsInterfaceDetailsWithDefaults

`func NewSystemInsightsInterfaceDetailsWithDefaults() *SystemInsightsInterfaceDetails`

NewSystemInsightsInterfaceDetailsWithDefaults instantiates a new SystemInsightsInterfaceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollisions

`func (o *SystemInsightsInterfaceDetails) GetCollisions() string`

GetCollisions returns the Collisions field if non-nil, zero value otherwise.

### GetCollisionsOk

`func (o *SystemInsightsInterfaceDetails) GetCollisionsOk() (*string, bool)`

GetCollisionsOk returns a tuple with the Collisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollisions

`func (o *SystemInsightsInterfaceDetails) SetCollisions(v string)`

SetCollisions sets Collisions field to given value.

### HasCollisions

`func (o *SystemInsightsInterfaceDetails) HasCollisions() bool`

HasCollisions returns a boolean if a field has been set.

### GetConnectionId

`func (o *SystemInsightsInterfaceDetails) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SystemInsightsInterfaceDetails) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SystemInsightsInterfaceDetails) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SystemInsightsInterfaceDetails) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *SystemInsightsInterfaceDetails) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *SystemInsightsInterfaceDetails) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *SystemInsightsInterfaceDetails) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *SystemInsightsInterfaceDetails) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDescription

`func (o *SystemInsightsInterfaceDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemInsightsInterfaceDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemInsightsInterfaceDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemInsightsInterfaceDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDhcpEnabled

`func (o *SystemInsightsInterfaceDetails) GetDhcpEnabled() int32`

GetDhcpEnabled returns the DhcpEnabled field if non-nil, zero value otherwise.

### GetDhcpEnabledOk

`func (o *SystemInsightsInterfaceDetails) GetDhcpEnabledOk() (*int32, bool)`

GetDhcpEnabledOk returns a tuple with the DhcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnabled

`func (o *SystemInsightsInterfaceDetails) SetDhcpEnabled(v int32)`

SetDhcpEnabled sets DhcpEnabled field to given value.

### HasDhcpEnabled

`func (o *SystemInsightsInterfaceDetails) HasDhcpEnabled() bool`

HasDhcpEnabled returns a boolean if a field has been set.

### GetDhcpLeaseExpires

`func (o *SystemInsightsInterfaceDetails) GetDhcpLeaseExpires() string`

GetDhcpLeaseExpires returns the DhcpLeaseExpires field if non-nil, zero value otherwise.

### GetDhcpLeaseExpiresOk

`func (o *SystemInsightsInterfaceDetails) GetDhcpLeaseExpiresOk() (*string, bool)`

GetDhcpLeaseExpiresOk returns a tuple with the DhcpLeaseExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseExpires

`func (o *SystemInsightsInterfaceDetails) SetDhcpLeaseExpires(v string)`

SetDhcpLeaseExpires sets DhcpLeaseExpires field to given value.

### HasDhcpLeaseExpires

`func (o *SystemInsightsInterfaceDetails) HasDhcpLeaseExpires() bool`

HasDhcpLeaseExpires returns a boolean if a field has been set.

### GetDhcpLeaseObtained

`func (o *SystemInsightsInterfaceDetails) GetDhcpLeaseObtained() string`

GetDhcpLeaseObtained returns the DhcpLeaseObtained field if non-nil, zero value otherwise.

### GetDhcpLeaseObtainedOk

`func (o *SystemInsightsInterfaceDetails) GetDhcpLeaseObtainedOk() (*string, bool)`

GetDhcpLeaseObtainedOk returns a tuple with the DhcpLeaseObtained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseObtained

`func (o *SystemInsightsInterfaceDetails) SetDhcpLeaseObtained(v string)`

SetDhcpLeaseObtained sets DhcpLeaseObtained field to given value.

### HasDhcpLeaseObtained

`func (o *SystemInsightsInterfaceDetails) HasDhcpLeaseObtained() bool`

HasDhcpLeaseObtained returns a boolean if a field has been set.

### GetDhcpServer

`func (o *SystemInsightsInterfaceDetails) GetDhcpServer() string`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *SystemInsightsInterfaceDetails) GetDhcpServerOk() (*string, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *SystemInsightsInterfaceDetails) SetDhcpServer(v string)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *SystemInsightsInterfaceDetails) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDnsDomain

`func (o *SystemInsightsInterfaceDetails) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *SystemInsightsInterfaceDetails) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *SystemInsightsInterfaceDetails) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *SystemInsightsInterfaceDetails) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.

### GetDnsDomainSuffixSearchOrder

`func (o *SystemInsightsInterfaceDetails) GetDnsDomainSuffixSearchOrder() string`

GetDnsDomainSuffixSearchOrder returns the DnsDomainSuffixSearchOrder field if non-nil, zero value otherwise.

### GetDnsDomainSuffixSearchOrderOk

`func (o *SystemInsightsInterfaceDetails) GetDnsDomainSuffixSearchOrderOk() (*string, bool)`

GetDnsDomainSuffixSearchOrderOk returns a tuple with the DnsDomainSuffixSearchOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomainSuffixSearchOrder

`func (o *SystemInsightsInterfaceDetails) SetDnsDomainSuffixSearchOrder(v string)`

SetDnsDomainSuffixSearchOrder sets DnsDomainSuffixSearchOrder field to given value.

### HasDnsDomainSuffixSearchOrder

`func (o *SystemInsightsInterfaceDetails) HasDnsDomainSuffixSearchOrder() bool`

HasDnsDomainSuffixSearchOrder returns a boolean if a field has been set.

### GetDnsHostName

`func (o *SystemInsightsInterfaceDetails) GetDnsHostName() string`

GetDnsHostName returns the DnsHostName field if non-nil, zero value otherwise.

### GetDnsHostNameOk

`func (o *SystemInsightsInterfaceDetails) GetDnsHostNameOk() (*string, bool)`

GetDnsHostNameOk returns a tuple with the DnsHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHostName

`func (o *SystemInsightsInterfaceDetails) SetDnsHostName(v string)`

SetDnsHostName sets DnsHostName field to given value.

### HasDnsHostName

`func (o *SystemInsightsInterfaceDetails) HasDnsHostName() bool`

HasDnsHostName returns a boolean if a field has been set.

### GetDnsServerSearchOrder

`func (o *SystemInsightsInterfaceDetails) GetDnsServerSearchOrder() string`

GetDnsServerSearchOrder returns the DnsServerSearchOrder field if non-nil, zero value otherwise.

### GetDnsServerSearchOrderOk

`func (o *SystemInsightsInterfaceDetails) GetDnsServerSearchOrderOk() (*string, bool)`

GetDnsServerSearchOrderOk returns a tuple with the DnsServerSearchOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerSearchOrder

`func (o *SystemInsightsInterfaceDetails) SetDnsServerSearchOrder(v string)`

SetDnsServerSearchOrder sets DnsServerSearchOrder field to given value.

### HasDnsServerSearchOrder

`func (o *SystemInsightsInterfaceDetails) HasDnsServerSearchOrder() bool`

HasDnsServerSearchOrder returns a boolean if a field has been set.

### GetEnabled

`func (o *SystemInsightsInterfaceDetails) GetEnabled() int32`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SystemInsightsInterfaceDetails) GetEnabledOk() (*int32, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SystemInsightsInterfaceDetails) SetEnabled(v int32)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SystemInsightsInterfaceDetails) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFlags

`func (o *SystemInsightsInterfaceDetails) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SystemInsightsInterfaceDetails) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SystemInsightsInterfaceDetails) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *SystemInsightsInterfaceDetails) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetFriendlyName

`func (o *SystemInsightsInterfaceDetails) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *SystemInsightsInterfaceDetails) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *SystemInsightsInterfaceDetails) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *SystemInsightsInterfaceDetails) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetIbytes

`func (o *SystemInsightsInterfaceDetails) GetIbytes() string`

GetIbytes returns the Ibytes field if non-nil, zero value otherwise.

### GetIbytesOk

`func (o *SystemInsightsInterfaceDetails) GetIbytesOk() (*string, bool)`

GetIbytesOk returns a tuple with the Ibytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbytes

`func (o *SystemInsightsInterfaceDetails) SetIbytes(v string)`

SetIbytes sets Ibytes field to given value.

### HasIbytes

`func (o *SystemInsightsInterfaceDetails) HasIbytes() bool`

HasIbytes returns a boolean if a field has been set.

### GetIdrops

`func (o *SystemInsightsInterfaceDetails) GetIdrops() string`

GetIdrops returns the Idrops field if non-nil, zero value otherwise.

### GetIdropsOk

`func (o *SystemInsightsInterfaceDetails) GetIdropsOk() (*string, bool)`

GetIdropsOk returns a tuple with the Idrops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdrops

`func (o *SystemInsightsInterfaceDetails) SetIdrops(v string)`

SetIdrops sets Idrops field to given value.

### HasIdrops

`func (o *SystemInsightsInterfaceDetails) HasIdrops() bool`

HasIdrops returns a boolean if a field has been set.

### GetIerrors

`func (o *SystemInsightsInterfaceDetails) GetIerrors() string`

GetIerrors returns the Ierrors field if non-nil, zero value otherwise.

### GetIerrorsOk

`func (o *SystemInsightsInterfaceDetails) GetIerrorsOk() (*string, bool)`

GetIerrorsOk returns a tuple with the Ierrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIerrors

`func (o *SystemInsightsInterfaceDetails) SetIerrors(v string)`

SetIerrors sets Ierrors field to given value.

### HasIerrors

`func (o *SystemInsightsInterfaceDetails) HasIerrors() bool`

HasIerrors returns a boolean if a field has been set.

### GetInterface

`func (o *SystemInsightsInterfaceDetails) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *SystemInsightsInterfaceDetails) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *SystemInsightsInterfaceDetails) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *SystemInsightsInterfaceDetails) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIpackets

`func (o *SystemInsightsInterfaceDetails) GetIpackets() string`

GetIpackets returns the Ipackets field if non-nil, zero value otherwise.

### GetIpacketsOk

`func (o *SystemInsightsInterfaceDetails) GetIpacketsOk() (*string, bool)`

GetIpacketsOk returns a tuple with the Ipackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpackets

`func (o *SystemInsightsInterfaceDetails) SetIpackets(v string)`

SetIpackets sets Ipackets field to given value.

### HasIpackets

`func (o *SystemInsightsInterfaceDetails) HasIpackets() bool`

HasIpackets returns a boolean if a field has been set.

### GetLastChange

`func (o *SystemInsightsInterfaceDetails) GetLastChange() string`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *SystemInsightsInterfaceDetails) GetLastChangeOk() (*string, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *SystemInsightsInterfaceDetails) SetLastChange(v string)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *SystemInsightsInterfaceDetails) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *SystemInsightsInterfaceDetails) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *SystemInsightsInterfaceDetails) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *SystemInsightsInterfaceDetails) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *SystemInsightsInterfaceDetails) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMac

`func (o *SystemInsightsInterfaceDetails) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SystemInsightsInterfaceDetails) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SystemInsightsInterfaceDetails) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SystemInsightsInterfaceDetails) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManufacturer

`func (o *SystemInsightsInterfaceDetails) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *SystemInsightsInterfaceDetails) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *SystemInsightsInterfaceDetails) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *SystemInsightsInterfaceDetails) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetMetric

`func (o *SystemInsightsInterfaceDetails) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SystemInsightsInterfaceDetails) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SystemInsightsInterfaceDetails) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *SystemInsightsInterfaceDetails) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetMtu

`func (o *SystemInsightsInterfaceDetails) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *SystemInsightsInterfaceDetails) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *SystemInsightsInterfaceDetails) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *SystemInsightsInterfaceDetails) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetObytes

`func (o *SystemInsightsInterfaceDetails) GetObytes() string`

GetObytes returns the Obytes field if non-nil, zero value otherwise.

### GetObytesOk

`func (o *SystemInsightsInterfaceDetails) GetObytesOk() (*string, bool)`

GetObytesOk returns a tuple with the Obytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObytes

`func (o *SystemInsightsInterfaceDetails) SetObytes(v string)`

SetObytes sets Obytes field to given value.

### HasObytes

`func (o *SystemInsightsInterfaceDetails) HasObytes() bool`

HasObytes returns a boolean if a field has been set.

### GetOdrops

`func (o *SystemInsightsInterfaceDetails) GetOdrops() string`

GetOdrops returns the Odrops field if non-nil, zero value otherwise.

### GetOdropsOk

`func (o *SystemInsightsInterfaceDetails) GetOdropsOk() (*string, bool)`

GetOdropsOk returns a tuple with the Odrops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdrops

`func (o *SystemInsightsInterfaceDetails) SetOdrops(v string)`

SetOdrops sets Odrops field to given value.

### HasOdrops

`func (o *SystemInsightsInterfaceDetails) HasOdrops() bool`

HasOdrops returns a boolean if a field has been set.

### GetOerrors

`func (o *SystemInsightsInterfaceDetails) GetOerrors() string`

GetOerrors returns the Oerrors field if non-nil, zero value otherwise.

### GetOerrorsOk

`func (o *SystemInsightsInterfaceDetails) GetOerrorsOk() (*string, bool)`

GetOerrorsOk returns a tuple with the Oerrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOerrors

`func (o *SystemInsightsInterfaceDetails) SetOerrors(v string)`

SetOerrors sets Oerrors field to given value.

### HasOerrors

`func (o *SystemInsightsInterfaceDetails) HasOerrors() bool`

HasOerrors returns a boolean if a field has been set.

### GetOpackets

`func (o *SystemInsightsInterfaceDetails) GetOpackets() string`

GetOpackets returns the Opackets field if non-nil, zero value otherwise.

### GetOpacketsOk

`func (o *SystemInsightsInterfaceDetails) GetOpacketsOk() (*string, bool)`

GetOpacketsOk returns a tuple with the Opackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpackets

`func (o *SystemInsightsInterfaceDetails) SetOpackets(v string)`

SetOpackets sets Opackets field to given value.

### HasOpackets

`func (o *SystemInsightsInterfaceDetails) HasOpackets() bool`

HasOpackets returns a boolean if a field has been set.

### GetPciSlot

`func (o *SystemInsightsInterfaceDetails) GetPciSlot() string`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *SystemInsightsInterfaceDetails) GetPciSlotOk() (*string, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *SystemInsightsInterfaceDetails) SetPciSlot(v string)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *SystemInsightsInterfaceDetails) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetPhysicalAdapter

`func (o *SystemInsightsInterfaceDetails) GetPhysicalAdapter() int32`

GetPhysicalAdapter returns the PhysicalAdapter field if non-nil, zero value otherwise.

### GetPhysicalAdapterOk

`func (o *SystemInsightsInterfaceDetails) GetPhysicalAdapterOk() (*int32, bool)`

GetPhysicalAdapterOk returns a tuple with the PhysicalAdapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAdapter

`func (o *SystemInsightsInterfaceDetails) SetPhysicalAdapter(v int32)`

SetPhysicalAdapter sets PhysicalAdapter field to given value.

### HasPhysicalAdapter

`func (o *SystemInsightsInterfaceDetails) HasPhysicalAdapter() bool`

HasPhysicalAdapter returns a boolean if a field has been set.

### GetService

`func (o *SystemInsightsInterfaceDetails) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SystemInsightsInterfaceDetails) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SystemInsightsInterfaceDetails) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SystemInsightsInterfaceDetails) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSpeed

`func (o *SystemInsightsInterfaceDetails) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *SystemInsightsInterfaceDetails) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *SystemInsightsInterfaceDetails) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *SystemInsightsInterfaceDetails) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSystemId

`func (o *SystemInsightsInterfaceDetails) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *SystemInsightsInterfaceDetails) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *SystemInsightsInterfaceDetails) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *SystemInsightsInterfaceDetails) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetType

`func (o *SystemInsightsInterfaceDetails) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemInsightsInterfaceDetails) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemInsightsInterfaceDetails) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SystemInsightsInterfaceDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



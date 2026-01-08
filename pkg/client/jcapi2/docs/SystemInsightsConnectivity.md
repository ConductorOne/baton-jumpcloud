# SystemInsightsConnectivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTime** | Pointer to **string** |  | [optional] 
**Disconnected** | Pointer to **int32** |  | [optional] 
**Ipv4Internet** | Pointer to **int32** |  | [optional] 
**Ipv4LocalNetwork** | Pointer to **int32** |  | [optional] 
**Ipv4NoTraffic** | Pointer to **int32** |  | [optional] 
**Ipv4Subnet** | Pointer to **int32** |  | [optional] 
**Ipv6Internet** | Pointer to **int32** |  | [optional] 
**Ipv6LocalNetwork** | Pointer to **int32** |  | [optional] 
**Ipv6NoTraffic** | Pointer to **int32** |  | [optional] 
**Ipv6Subnet** | Pointer to **int32** |  | [optional] 
**SystemId** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemInsightsConnectivity

`func NewSystemInsightsConnectivity() *SystemInsightsConnectivity`

NewSystemInsightsConnectivity instantiates a new SystemInsightsConnectivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInsightsConnectivityWithDefaults

`func NewSystemInsightsConnectivityWithDefaults() *SystemInsightsConnectivity`

NewSystemInsightsConnectivityWithDefaults instantiates a new SystemInsightsConnectivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTime

`func (o *SystemInsightsConnectivity) GetCollectionTime() string`

GetCollectionTime returns the CollectionTime field if non-nil, zero value otherwise.

### GetCollectionTimeOk

`func (o *SystemInsightsConnectivity) GetCollectionTimeOk() (*string, bool)`

GetCollectionTimeOk returns a tuple with the CollectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTime

`func (o *SystemInsightsConnectivity) SetCollectionTime(v string)`

SetCollectionTime sets CollectionTime field to given value.

### HasCollectionTime

`func (o *SystemInsightsConnectivity) HasCollectionTime() bool`

HasCollectionTime returns a boolean if a field has been set.

### GetDisconnected

`func (o *SystemInsightsConnectivity) GetDisconnected() int32`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *SystemInsightsConnectivity) GetDisconnectedOk() (*int32, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *SystemInsightsConnectivity) SetDisconnected(v int32)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *SystemInsightsConnectivity) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetIpv4Internet

`func (o *SystemInsightsConnectivity) GetIpv4Internet() int32`

GetIpv4Internet returns the Ipv4Internet field if non-nil, zero value otherwise.

### GetIpv4InternetOk

`func (o *SystemInsightsConnectivity) GetIpv4InternetOk() (*int32, bool)`

GetIpv4InternetOk returns a tuple with the Ipv4Internet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Internet

`func (o *SystemInsightsConnectivity) SetIpv4Internet(v int32)`

SetIpv4Internet sets Ipv4Internet field to given value.

### HasIpv4Internet

`func (o *SystemInsightsConnectivity) HasIpv4Internet() bool`

HasIpv4Internet returns a boolean if a field has been set.

### GetIpv4LocalNetwork

`func (o *SystemInsightsConnectivity) GetIpv4LocalNetwork() int32`

GetIpv4LocalNetwork returns the Ipv4LocalNetwork field if non-nil, zero value otherwise.

### GetIpv4LocalNetworkOk

`func (o *SystemInsightsConnectivity) GetIpv4LocalNetworkOk() (*int32, bool)`

GetIpv4LocalNetworkOk returns a tuple with the Ipv4LocalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4LocalNetwork

`func (o *SystemInsightsConnectivity) SetIpv4LocalNetwork(v int32)`

SetIpv4LocalNetwork sets Ipv4LocalNetwork field to given value.

### HasIpv4LocalNetwork

`func (o *SystemInsightsConnectivity) HasIpv4LocalNetwork() bool`

HasIpv4LocalNetwork returns a boolean if a field has been set.

### GetIpv4NoTraffic

`func (o *SystemInsightsConnectivity) GetIpv4NoTraffic() int32`

GetIpv4NoTraffic returns the Ipv4NoTraffic field if non-nil, zero value otherwise.

### GetIpv4NoTrafficOk

`func (o *SystemInsightsConnectivity) GetIpv4NoTrafficOk() (*int32, bool)`

GetIpv4NoTrafficOk returns a tuple with the Ipv4NoTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4NoTraffic

`func (o *SystemInsightsConnectivity) SetIpv4NoTraffic(v int32)`

SetIpv4NoTraffic sets Ipv4NoTraffic field to given value.

### HasIpv4NoTraffic

`func (o *SystemInsightsConnectivity) HasIpv4NoTraffic() bool`

HasIpv4NoTraffic returns a boolean if a field has been set.

### GetIpv4Subnet

`func (o *SystemInsightsConnectivity) GetIpv4Subnet() int32`

GetIpv4Subnet returns the Ipv4Subnet field if non-nil, zero value otherwise.

### GetIpv4SubnetOk

`func (o *SystemInsightsConnectivity) GetIpv4SubnetOk() (*int32, bool)`

GetIpv4SubnetOk returns a tuple with the Ipv4Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Subnet

`func (o *SystemInsightsConnectivity) SetIpv4Subnet(v int32)`

SetIpv4Subnet sets Ipv4Subnet field to given value.

### HasIpv4Subnet

`func (o *SystemInsightsConnectivity) HasIpv4Subnet() bool`

HasIpv4Subnet returns a boolean if a field has been set.

### GetIpv6Internet

`func (o *SystemInsightsConnectivity) GetIpv6Internet() int32`

GetIpv6Internet returns the Ipv6Internet field if non-nil, zero value otherwise.

### GetIpv6InternetOk

`func (o *SystemInsightsConnectivity) GetIpv6InternetOk() (*int32, bool)`

GetIpv6InternetOk returns a tuple with the Ipv6Internet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Internet

`func (o *SystemInsightsConnectivity) SetIpv6Internet(v int32)`

SetIpv6Internet sets Ipv6Internet field to given value.

### HasIpv6Internet

`func (o *SystemInsightsConnectivity) HasIpv6Internet() bool`

HasIpv6Internet returns a boolean if a field has been set.

### GetIpv6LocalNetwork

`func (o *SystemInsightsConnectivity) GetIpv6LocalNetwork() int32`

GetIpv6LocalNetwork returns the Ipv6LocalNetwork field if non-nil, zero value otherwise.

### GetIpv6LocalNetworkOk

`func (o *SystemInsightsConnectivity) GetIpv6LocalNetworkOk() (*int32, bool)`

GetIpv6LocalNetworkOk returns a tuple with the Ipv6LocalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6LocalNetwork

`func (o *SystemInsightsConnectivity) SetIpv6LocalNetwork(v int32)`

SetIpv6LocalNetwork sets Ipv6LocalNetwork field to given value.

### HasIpv6LocalNetwork

`func (o *SystemInsightsConnectivity) HasIpv6LocalNetwork() bool`

HasIpv6LocalNetwork returns a boolean if a field has been set.

### GetIpv6NoTraffic

`func (o *SystemInsightsConnectivity) GetIpv6NoTraffic() int32`

GetIpv6NoTraffic returns the Ipv6NoTraffic field if non-nil, zero value otherwise.

### GetIpv6NoTrafficOk

`func (o *SystemInsightsConnectivity) GetIpv6NoTrafficOk() (*int32, bool)`

GetIpv6NoTrafficOk returns a tuple with the Ipv6NoTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6NoTraffic

`func (o *SystemInsightsConnectivity) SetIpv6NoTraffic(v int32)`

SetIpv6NoTraffic sets Ipv6NoTraffic field to given value.

### HasIpv6NoTraffic

`func (o *SystemInsightsConnectivity) HasIpv6NoTraffic() bool`

HasIpv6NoTraffic returns a boolean if a field has been set.

### GetIpv6Subnet

`func (o *SystemInsightsConnectivity) GetIpv6Subnet() int32`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *SystemInsightsConnectivity) GetIpv6SubnetOk() (*int32, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *SystemInsightsConnectivity) SetIpv6Subnet(v int32)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *SystemInsightsConnectivity) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### GetSystemId

`func (o *SystemInsightsConnectivity) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *SystemInsightsConnectivity) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *SystemInsightsConnectivity) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *SystemInsightsConnectivity) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



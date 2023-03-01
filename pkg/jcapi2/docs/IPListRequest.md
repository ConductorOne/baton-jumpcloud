# IPListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Ips** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewIPListRequest

`func NewIPListRequest() *IPListRequest`

NewIPListRequest instantiates a new IPListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPListRequestWithDefaults

`func NewIPListRequestWithDefaults() *IPListRequest`

NewIPListRequestWithDefaults instantiates a new IPListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IPListRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPListRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPListRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPListRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIps

`func (o *IPListRequest) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IPListRequest) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IPListRequest) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *IPListRequest) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetName

`func (o *IPListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPListRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPListRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



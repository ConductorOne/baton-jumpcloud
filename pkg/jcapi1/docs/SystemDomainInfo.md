# SystemDomainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** |  | [optional] 
**PartOfDomain** | Pointer to **bool** |  | [optional] 

## Methods

### NewSystemDomainInfo

`func NewSystemDomainInfo() *SystemDomainInfo`

NewSystemDomainInfo instantiates a new SystemDomainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemDomainInfoWithDefaults

`func NewSystemDomainInfoWithDefaults() *SystemDomainInfo`

NewSystemDomainInfoWithDefaults instantiates a new SystemDomainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *SystemDomainInfo) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *SystemDomainInfo) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *SystemDomainInfo) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *SystemDomainInfo) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetPartOfDomain

`func (o *SystemDomainInfo) GetPartOfDomain() bool`

GetPartOfDomain returns the PartOfDomain field if non-nil, zero value otherwise.

### GetPartOfDomainOk

`func (o *SystemDomainInfo) GetPartOfDomainOk() (*bool, bool)`

GetPartOfDomainOk returns a tuple with the PartOfDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOfDomain

`func (o *SystemDomainInfo) SetPartOfDomain(v bool)`

SetPartOfDomain sets PartOfDomain field to given value.

### HasPartOfDomain

`func (o *SystemDomainInfo) HasPartOfDomain() bool`

HasPartOfDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



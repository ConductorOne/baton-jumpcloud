# TrustedappConfigGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | **string** | Checksum to validate the trustedApp configuration for the organization | 
**TrustedApps** | [**[]TrustedappConfigGetTrustedAppsInner**](TrustedappConfigGetTrustedAppsInner.md) | List of authorized apps for the organization  | 

## Methods

### NewTrustedappConfigGet

`func NewTrustedappConfigGet(checksum string, trustedApps []TrustedappConfigGetTrustedAppsInner, ) *TrustedappConfigGet`

NewTrustedappConfigGet instantiates a new TrustedappConfigGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedappConfigGetWithDefaults

`func NewTrustedappConfigGetWithDefaults() *TrustedappConfigGet`

NewTrustedappConfigGetWithDefaults instantiates a new TrustedappConfigGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *TrustedappConfigGet) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *TrustedappConfigGet) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *TrustedappConfigGet) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.


### GetTrustedApps

`func (o *TrustedappConfigGet) GetTrustedApps() []TrustedappConfigGetTrustedAppsInner`

GetTrustedApps returns the TrustedApps field if non-nil, zero value otherwise.

### GetTrustedAppsOk

`func (o *TrustedappConfigGet) GetTrustedAppsOk() (*[]TrustedappConfigGetTrustedAppsInner, bool)`

GetTrustedAppsOk returns a tuple with the TrustedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedApps

`func (o *TrustedappConfigGet) SetTrustedApps(v []TrustedappConfigGetTrustedAppsInner)`

SetTrustedApps sets TrustedApps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



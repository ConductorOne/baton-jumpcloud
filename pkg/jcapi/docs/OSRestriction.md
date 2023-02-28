# OSRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppleRestrictions** | Pointer to [**OSRestrictionAppleRestrictions**](OSRestrictionAppleRestrictions.md) |  | [optional] 
**DeprecatedVersion** | Pointer to **string** | The version of the OS in which the policy was deprecated | [optional] 
**EarliestVersion** | Pointer to **string** | The earliest version of the OS in which the policy can be applied | [optional] 
**OsName** | Pointer to **string** | The name of the OS in which this restriction applies | [optional] 
**SupportedEnrollmentTypes** | Pointer to **[]string** | This field is deprecated and will be ignored. Use appleRestrictions.supportedEnrollmentTypes instead | [optional] 

## Methods

### NewOSRestriction

`func NewOSRestriction() *OSRestriction`

NewOSRestriction instantiates a new OSRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSRestrictionWithDefaults

`func NewOSRestrictionWithDefaults() *OSRestriction`

NewOSRestrictionWithDefaults instantiates a new OSRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppleRestrictions

`func (o *OSRestriction) GetAppleRestrictions() OSRestrictionAppleRestrictions`

GetAppleRestrictions returns the AppleRestrictions field if non-nil, zero value otherwise.

### GetAppleRestrictionsOk

`func (o *OSRestriction) GetAppleRestrictionsOk() (*OSRestrictionAppleRestrictions, bool)`

GetAppleRestrictionsOk returns a tuple with the AppleRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleRestrictions

`func (o *OSRestriction) SetAppleRestrictions(v OSRestrictionAppleRestrictions)`

SetAppleRestrictions sets AppleRestrictions field to given value.

### HasAppleRestrictions

`func (o *OSRestriction) HasAppleRestrictions() bool`

HasAppleRestrictions returns a boolean if a field has been set.

### GetDeprecatedVersion

`func (o *OSRestriction) GetDeprecatedVersion() string`

GetDeprecatedVersion returns the DeprecatedVersion field if non-nil, zero value otherwise.

### GetDeprecatedVersionOk

`func (o *OSRestriction) GetDeprecatedVersionOk() (*string, bool)`

GetDeprecatedVersionOk returns a tuple with the DeprecatedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedVersion

`func (o *OSRestriction) SetDeprecatedVersion(v string)`

SetDeprecatedVersion sets DeprecatedVersion field to given value.

### HasDeprecatedVersion

`func (o *OSRestriction) HasDeprecatedVersion() bool`

HasDeprecatedVersion returns a boolean if a field has been set.

### GetEarliestVersion

`func (o *OSRestriction) GetEarliestVersion() string`

GetEarliestVersion returns the EarliestVersion field if non-nil, zero value otherwise.

### GetEarliestVersionOk

`func (o *OSRestriction) GetEarliestVersionOk() (*string, bool)`

GetEarliestVersionOk returns a tuple with the EarliestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestVersion

`func (o *OSRestriction) SetEarliestVersion(v string)`

SetEarliestVersion sets EarliestVersion field to given value.

### HasEarliestVersion

`func (o *OSRestriction) HasEarliestVersion() bool`

HasEarliestVersion returns a boolean if a field has been set.

### GetOsName

`func (o *OSRestriction) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *OSRestriction) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *OSRestriction) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *OSRestriction) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetSupportedEnrollmentTypes

`func (o *OSRestriction) GetSupportedEnrollmentTypes() []string`

GetSupportedEnrollmentTypes returns the SupportedEnrollmentTypes field if non-nil, zero value otherwise.

### GetSupportedEnrollmentTypesOk

`func (o *OSRestriction) GetSupportedEnrollmentTypesOk() (*[]string, bool)`

GetSupportedEnrollmentTypesOk returns a tuple with the SupportedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEnrollmentTypes

`func (o *OSRestriction) SetSupportedEnrollmentTypes(v []string)`

SetSupportedEnrollmentTypes sets SupportedEnrollmentTypes field to given value.

### HasSupportedEnrollmentTypes

`func (o *OSRestriction) HasSupportedEnrollmentTypes() bool`

HasSupportedEnrollmentTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



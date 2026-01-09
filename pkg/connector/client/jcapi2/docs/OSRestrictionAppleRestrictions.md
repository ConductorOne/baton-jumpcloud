# OSRestrictionAppleRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiresSupervision** | Pointer to **bool** | Boolean representing if the policy requires the Apple devices to be MDM supervised | [optional] 
**SupportedEnrollmentTypes** | Pointer to **[]string** | The supported Apple enrollment types for this policy | [optional] 

## Methods

### NewOSRestrictionAppleRestrictions

`func NewOSRestrictionAppleRestrictions() *OSRestrictionAppleRestrictions`

NewOSRestrictionAppleRestrictions instantiates a new OSRestrictionAppleRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSRestrictionAppleRestrictionsWithDefaults

`func NewOSRestrictionAppleRestrictionsWithDefaults() *OSRestrictionAppleRestrictions`

NewOSRestrictionAppleRestrictionsWithDefaults instantiates a new OSRestrictionAppleRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiresSupervision

`func (o *OSRestrictionAppleRestrictions) GetRequiresSupervision() bool`

GetRequiresSupervision returns the RequiresSupervision field if non-nil, zero value otherwise.

### GetRequiresSupervisionOk

`func (o *OSRestrictionAppleRestrictions) GetRequiresSupervisionOk() (*bool, bool)`

GetRequiresSupervisionOk returns a tuple with the RequiresSupervision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSupervision

`func (o *OSRestrictionAppleRestrictions) SetRequiresSupervision(v bool)`

SetRequiresSupervision sets RequiresSupervision field to given value.

### HasRequiresSupervision

`func (o *OSRestrictionAppleRestrictions) HasRequiresSupervision() bool`

HasRequiresSupervision returns a boolean if a field has been set.

### GetSupportedEnrollmentTypes

`func (o *OSRestrictionAppleRestrictions) GetSupportedEnrollmentTypes() []string`

GetSupportedEnrollmentTypes returns the SupportedEnrollmentTypes field if non-nil, zero value otherwise.

### GetSupportedEnrollmentTypesOk

`func (o *OSRestrictionAppleRestrictions) GetSupportedEnrollmentTypesOk() (*[]string, bool)`

GetSupportedEnrollmentTypesOk returns a tuple with the SupportedEnrollmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEnrollmentTypes

`func (o *OSRestrictionAppleRestrictions) SetSupportedEnrollmentTypes(v []string)`

SetSupportedEnrollmentTypes sets SupportedEnrollmentTypes field to given value.

### HasSupportedEnrollmentTypes

`func (o *OSRestrictionAppleRestrictions) HasSupportedEnrollmentTypes() bool`

HasSupportedEnrollmentTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



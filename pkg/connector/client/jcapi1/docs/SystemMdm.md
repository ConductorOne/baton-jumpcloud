# SystemMdm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dep** | Pointer to **bool** |  | [optional] 
**EnrollmentType** | Pointer to **string** |  | [optional] 
**Internal** | Pointer to [**SystemMdmInternal**](SystemMdmInternal.md) |  | [optional] 
**ProfileIdentifier** | Pointer to **string** |  | [optional] 
**UserApproved** | Pointer to **bool** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemMdm

`func NewSystemMdm() *SystemMdm`

NewSystemMdm instantiates a new SystemMdm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemMdmWithDefaults

`func NewSystemMdmWithDefaults() *SystemMdm`

NewSystemMdmWithDefaults instantiates a new SystemMdm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDep

`func (o *SystemMdm) GetDep() bool`

GetDep returns the Dep field if non-nil, zero value otherwise.

### GetDepOk

`func (o *SystemMdm) GetDepOk() (*bool, bool)`

GetDepOk returns a tuple with the Dep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDep

`func (o *SystemMdm) SetDep(v bool)`

SetDep sets Dep field to given value.

### HasDep

`func (o *SystemMdm) HasDep() bool`

HasDep returns a boolean if a field has been set.

### GetEnrollmentType

`func (o *SystemMdm) GetEnrollmentType() string`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *SystemMdm) GetEnrollmentTypeOk() (*string, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentType

`func (o *SystemMdm) SetEnrollmentType(v string)`

SetEnrollmentType sets EnrollmentType field to given value.

### HasEnrollmentType

`func (o *SystemMdm) HasEnrollmentType() bool`

HasEnrollmentType returns a boolean if a field has been set.

### GetInternal

`func (o *SystemMdm) GetInternal() SystemMdmInternal`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *SystemMdm) GetInternalOk() (*SystemMdmInternal, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *SystemMdm) SetInternal(v SystemMdmInternal)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *SystemMdm) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetProfileIdentifier

`func (o *SystemMdm) GetProfileIdentifier() string`

GetProfileIdentifier returns the ProfileIdentifier field if non-nil, zero value otherwise.

### GetProfileIdentifierOk

`func (o *SystemMdm) GetProfileIdentifierOk() (*string, bool)`

GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIdentifier

`func (o *SystemMdm) SetProfileIdentifier(v string)`

SetProfileIdentifier sets ProfileIdentifier field to given value.

### HasProfileIdentifier

`func (o *SystemMdm) HasProfileIdentifier() bool`

HasProfileIdentifier returns a boolean if a field has been set.

### GetUserApproved

`func (o *SystemMdm) GetUserApproved() bool`

GetUserApproved returns the UserApproved field if non-nil, zero value otherwise.

### GetUserApprovedOk

`func (o *SystemMdm) GetUserApprovedOk() (*bool, bool)`

GetUserApprovedOk returns a tuple with the UserApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserApproved

`func (o *SystemMdm) SetUserApproved(v bool)`

SetUserApproved sets UserApproved field to given value.

### HasUserApproved

`func (o *SystemMdm) HasUserApproved() bool`

HasUserApproved returns a boolean if a field has been set.

### GetVendor

`func (o *SystemMdm) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SystemMdm) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SystemMdm) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *SystemMdm) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



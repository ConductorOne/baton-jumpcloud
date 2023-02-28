# AppleMdmPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ades** | Pointer to [**ADES**](ADES.md) |  | [optional] 
**AllowMobileUserEnrollment** | Pointer to **bool** | A toggle to allow mobile device enrollment for an organization. | [optional] 
**AppleCertCreatorAppleID** | Pointer to **string** | The Apple ID of the admin who created the Apple signed certificate. | [optional] 
**AppleSignedCert** | Pointer to **string** | A signed certificate obtained from Apple after providing Apple with the plist file provided on POST. | [optional] 
**DefaultIosUserEnrollmentDeviceGroupID** | Pointer to **string** | ObjectId uniquely identifying the MDM default iOS user enrollment device group. | [optional] 
**DefaultSystemGroupID** | Pointer to **string** | ObjectId uniquely identifying the MDM default System Group. | [optional] 
**Dep** | Pointer to [**DEP**](DEP.md) |  | [optional] 
**EncryptedDepServerToken** | Pointer to **string** | The S/MIME encoded DEP Server Token returned by Apple Business Manager when creating an MDM instance. | [optional] 
**Name** | Pointer to **string** | A new name for the Apple MDM configuration. | [optional] 

## Methods

### NewAppleMdmPatch

`func NewAppleMdmPatch() *AppleMdmPatch`

NewAppleMdmPatch instantiates a new AppleMdmPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleMdmPatchWithDefaults

`func NewAppleMdmPatchWithDefaults() *AppleMdmPatch`

NewAppleMdmPatchWithDefaults instantiates a new AppleMdmPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdes

`func (o *AppleMdmPatch) GetAdes() ADES`

GetAdes returns the Ades field if non-nil, zero value otherwise.

### GetAdesOk

`func (o *AppleMdmPatch) GetAdesOk() (*ADES, bool)`

GetAdesOk returns a tuple with the Ades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdes

`func (o *AppleMdmPatch) SetAdes(v ADES)`

SetAdes sets Ades field to given value.

### HasAdes

`func (o *AppleMdmPatch) HasAdes() bool`

HasAdes returns a boolean if a field has been set.

### GetAllowMobileUserEnrollment

`func (o *AppleMdmPatch) GetAllowMobileUserEnrollment() bool`

GetAllowMobileUserEnrollment returns the AllowMobileUserEnrollment field if non-nil, zero value otherwise.

### GetAllowMobileUserEnrollmentOk

`func (o *AppleMdmPatch) GetAllowMobileUserEnrollmentOk() (*bool, bool)`

GetAllowMobileUserEnrollmentOk returns a tuple with the AllowMobileUserEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMobileUserEnrollment

`func (o *AppleMdmPatch) SetAllowMobileUserEnrollment(v bool)`

SetAllowMobileUserEnrollment sets AllowMobileUserEnrollment field to given value.

### HasAllowMobileUserEnrollment

`func (o *AppleMdmPatch) HasAllowMobileUserEnrollment() bool`

HasAllowMobileUserEnrollment returns a boolean if a field has been set.

### GetAppleCertCreatorAppleID

`func (o *AppleMdmPatch) GetAppleCertCreatorAppleID() string`

GetAppleCertCreatorAppleID returns the AppleCertCreatorAppleID field if non-nil, zero value otherwise.

### GetAppleCertCreatorAppleIDOk

`func (o *AppleMdmPatch) GetAppleCertCreatorAppleIDOk() (*string, bool)`

GetAppleCertCreatorAppleIDOk returns a tuple with the AppleCertCreatorAppleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCertCreatorAppleID

`func (o *AppleMdmPatch) SetAppleCertCreatorAppleID(v string)`

SetAppleCertCreatorAppleID sets AppleCertCreatorAppleID field to given value.

### HasAppleCertCreatorAppleID

`func (o *AppleMdmPatch) HasAppleCertCreatorAppleID() bool`

HasAppleCertCreatorAppleID returns a boolean if a field has been set.

### GetAppleSignedCert

`func (o *AppleMdmPatch) GetAppleSignedCert() string`

GetAppleSignedCert returns the AppleSignedCert field if non-nil, zero value otherwise.

### GetAppleSignedCertOk

`func (o *AppleMdmPatch) GetAppleSignedCertOk() (*string, bool)`

GetAppleSignedCertOk returns a tuple with the AppleSignedCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleSignedCert

`func (o *AppleMdmPatch) SetAppleSignedCert(v string)`

SetAppleSignedCert sets AppleSignedCert field to given value.

### HasAppleSignedCert

`func (o *AppleMdmPatch) HasAppleSignedCert() bool`

HasAppleSignedCert returns a boolean if a field has been set.

### GetDefaultIosUserEnrollmentDeviceGroupID

`func (o *AppleMdmPatch) GetDefaultIosUserEnrollmentDeviceGroupID() string`

GetDefaultIosUserEnrollmentDeviceGroupID returns the DefaultIosUserEnrollmentDeviceGroupID field if non-nil, zero value otherwise.

### GetDefaultIosUserEnrollmentDeviceGroupIDOk

`func (o *AppleMdmPatch) GetDefaultIosUserEnrollmentDeviceGroupIDOk() (*string, bool)`

GetDefaultIosUserEnrollmentDeviceGroupIDOk returns a tuple with the DefaultIosUserEnrollmentDeviceGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIosUserEnrollmentDeviceGroupID

`func (o *AppleMdmPatch) SetDefaultIosUserEnrollmentDeviceGroupID(v string)`

SetDefaultIosUserEnrollmentDeviceGroupID sets DefaultIosUserEnrollmentDeviceGroupID field to given value.

### HasDefaultIosUserEnrollmentDeviceGroupID

`func (o *AppleMdmPatch) HasDefaultIosUserEnrollmentDeviceGroupID() bool`

HasDefaultIosUserEnrollmentDeviceGroupID returns a boolean if a field has been set.

### GetDefaultSystemGroupID

`func (o *AppleMdmPatch) GetDefaultSystemGroupID() string`

GetDefaultSystemGroupID returns the DefaultSystemGroupID field if non-nil, zero value otherwise.

### GetDefaultSystemGroupIDOk

`func (o *AppleMdmPatch) GetDefaultSystemGroupIDOk() (*string, bool)`

GetDefaultSystemGroupIDOk returns a tuple with the DefaultSystemGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSystemGroupID

`func (o *AppleMdmPatch) SetDefaultSystemGroupID(v string)`

SetDefaultSystemGroupID sets DefaultSystemGroupID field to given value.

### HasDefaultSystemGroupID

`func (o *AppleMdmPatch) HasDefaultSystemGroupID() bool`

HasDefaultSystemGroupID returns a boolean if a field has been set.

### GetDep

`func (o *AppleMdmPatch) GetDep() DEP`

GetDep returns the Dep field if non-nil, zero value otherwise.

### GetDepOk

`func (o *AppleMdmPatch) GetDepOk() (*DEP, bool)`

GetDepOk returns a tuple with the Dep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDep

`func (o *AppleMdmPatch) SetDep(v DEP)`

SetDep sets Dep field to given value.

### HasDep

`func (o *AppleMdmPatch) HasDep() bool`

HasDep returns a boolean if a field has been set.

### GetEncryptedDepServerToken

`func (o *AppleMdmPatch) GetEncryptedDepServerToken() string`

GetEncryptedDepServerToken returns the EncryptedDepServerToken field if non-nil, zero value otherwise.

### GetEncryptedDepServerTokenOk

`func (o *AppleMdmPatch) GetEncryptedDepServerTokenOk() (*string, bool)`

GetEncryptedDepServerTokenOk returns a tuple with the EncryptedDepServerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedDepServerToken

`func (o *AppleMdmPatch) SetEncryptedDepServerToken(v string)`

SetEncryptedDepServerToken sets EncryptedDepServerToken field to given value.

### HasEncryptedDepServerToken

`func (o *AppleMdmPatch) HasEncryptedDepServerToken() bool`

HasEncryptedDepServerToken returns a boolean if a field has been set.

### GetName

`func (o *AppleMdmPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppleMdmPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppleMdmPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppleMdmPatch) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



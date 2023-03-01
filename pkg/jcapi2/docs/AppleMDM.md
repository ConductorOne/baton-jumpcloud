# AppleMDM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ades** | Pointer to [**ADES**](ADES.md) |  | [optional] 
**AllowMobileUserEnrollment** | Pointer to **bool** | A toggle to allow mobile device enrollment for an organization. | [optional] 
**ApnsCertExpiry** | Pointer to **string** | The expiration date and time for the APNS Certificate. | [optional] 
**ApnsPushTopic** | Pointer to **string** | The push topic assigned to this enrollment by Apple after uploading the Signed CSR plist. | [optional] 
**AppleCertCreatorAppleID** | Pointer to **string** | The Apple ID of the admin who created the Apple signed certificate associated to the Device Manager. | [optional] 
**AppleCertSerialNumber** | Pointer to **string** | The serial number of the Apple signed certificate associated to the Device Manager. | [optional] 
**DefaultIosUserEnrollmentDeviceGroupID** | Pointer to **string** | ObjectId uniquely identifying the MDM default iOS user enrollment device group. | [optional] 
**DefaultSystemGroupID** | Pointer to **string** | ObjectId uniquely identifying the MDM default System Group. | [optional] 
**Dep** | Pointer to [**DEP**](DEP.md) |  | [optional] 
**DepAccessTokenExpiry** | Pointer to **string** | The expiration date and time for the DEP Access Token. This aligns with the DEP Server Token State. | [optional] 
**DepServerTokenState** | Pointer to **string** | The state of the dep server token, presence and expiry. | [optional] 
**Id** | **string** | ObjectId uniquely identifying an MDM Enrollment, | 
**Name** | Pointer to **string** | A friendly name to identify this enrollment.  Not required to be unique. | [optional] 
**Organization** | Pointer to **string** | The identifier for an organization | [optional] 

## Methods

### NewAppleMDM

`func NewAppleMDM(id string, ) *AppleMDM`

NewAppleMDM instantiates a new AppleMDM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleMDMWithDefaults

`func NewAppleMDMWithDefaults() *AppleMDM`

NewAppleMDMWithDefaults instantiates a new AppleMDM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdes

`func (o *AppleMDM) GetAdes() ADES`

GetAdes returns the Ades field if non-nil, zero value otherwise.

### GetAdesOk

`func (o *AppleMDM) GetAdesOk() (*ADES, bool)`

GetAdesOk returns a tuple with the Ades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdes

`func (o *AppleMDM) SetAdes(v ADES)`

SetAdes sets Ades field to given value.

### HasAdes

`func (o *AppleMDM) HasAdes() bool`

HasAdes returns a boolean if a field has been set.

### GetAllowMobileUserEnrollment

`func (o *AppleMDM) GetAllowMobileUserEnrollment() bool`

GetAllowMobileUserEnrollment returns the AllowMobileUserEnrollment field if non-nil, zero value otherwise.

### GetAllowMobileUserEnrollmentOk

`func (o *AppleMDM) GetAllowMobileUserEnrollmentOk() (*bool, bool)`

GetAllowMobileUserEnrollmentOk returns a tuple with the AllowMobileUserEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMobileUserEnrollment

`func (o *AppleMDM) SetAllowMobileUserEnrollment(v bool)`

SetAllowMobileUserEnrollment sets AllowMobileUserEnrollment field to given value.

### HasAllowMobileUserEnrollment

`func (o *AppleMDM) HasAllowMobileUserEnrollment() bool`

HasAllowMobileUserEnrollment returns a boolean if a field has been set.

### GetApnsCertExpiry

`func (o *AppleMDM) GetApnsCertExpiry() string`

GetApnsCertExpiry returns the ApnsCertExpiry field if non-nil, zero value otherwise.

### GetApnsCertExpiryOk

`func (o *AppleMDM) GetApnsCertExpiryOk() (*string, bool)`

GetApnsCertExpiryOk returns a tuple with the ApnsCertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsCertExpiry

`func (o *AppleMDM) SetApnsCertExpiry(v string)`

SetApnsCertExpiry sets ApnsCertExpiry field to given value.

### HasApnsCertExpiry

`func (o *AppleMDM) HasApnsCertExpiry() bool`

HasApnsCertExpiry returns a boolean if a field has been set.

### GetApnsPushTopic

`func (o *AppleMDM) GetApnsPushTopic() string`

GetApnsPushTopic returns the ApnsPushTopic field if non-nil, zero value otherwise.

### GetApnsPushTopicOk

`func (o *AppleMDM) GetApnsPushTopicOk() (*string, bool)`

GetApnsPushTopicOk returns a tuple with the ApnsPushTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsPushTopic

`func (o *AppleMDM) SetApnsPushTopic(v string)`

SetApnsPushTopic sets ApnsPushTopic field to given value.

### HasApnsPushTopic

`func (o *AppleMDM) HasApnsPushTopic() bool`

HasApnsPushTopic returns a boolean if a field has been set.

### GetAppleCertCreatorAppleID

`func (o *AppleMDM) GetAppleCertCreatorAppleID() string`

GetAppleCertCreatorAppleID returns the AppleCertCreatorAppleID field if non-nil, zero value otherwise.

### GetAppleCertCreatorAppleIDOk

`func (o *AppleMDM) GetAppleCertCreatorAppleIDOk() (*string, bool)`

GetAppleCertCreatorAppleIDOk returns a tuple with the AppleCertCreatorAppleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCertCreatorAppleID

`func (o *AppleMDM) SetAppleCertCreatorAppleID(v string)`

SetAppleCertCreatorAppleID sets AppleCertCreatorAppleID field to given value.

### HasAppleCertCreatorAppleID

`func (o *AppleMDM) HasAppleCertCreatorAppleID() bool`

HasAppleCertCreatorAppleID returns a boolean if a field has been set.

### GetAppleCertSerialNumber

`func (o *AppleMDM) GetAppleCertSerialNumber() string`

GetAppleCertSerialNumber returns the AppleCertSerialNumber field if non-nil, zero value otherwise.

### GetAppleCertSerialNumberOk

`func (o *AppleMDM) GetAppleCertSerialNumberOk() (*string, bool)`

GetAppleCertSerialNumberOk returns a tuple with the AppleCertSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleCertSerialNumber

`func (o *AppleMDM) SetAppleCertSerialNumber(v string)`

SetAppleCertSerialNumber sets AppleCertSerialNumber field to given value.

### HasAppleCertSerialNumber

`func (o *AppleMDM) HasAppleCertSerialNumber() bool`

HasAppleCertSerialNumber returns a boolean if a field has been set.

### GetDefaultIosUserEnrollmentDeviceGroupID

`func (o *AppleMDM) GetDefaultIosUserEnrollmentDeviceGroupID() string`

GetDefaultIosUserEnrollmentDeviceGroupID returns the DefaultIosUserEnrollmentDeviceGroupID field if non-nil, zero value otherwise.

### GetDefaultIosUserEnrollmentDeviceGroupIDOk

`func (o *AppleMDM) GetDefaultIosUserEnrollmentDeviceGroupIDOk() (*string, bool)`

GetDefaultIosUserEnrollmentDeviceGroupIDOk returns a tuple with the DefaultIosUserEnrollmentDeviceGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIosUserEnrollmentDeviceGroupID

`func (o *AppleMDM) SetDefaultIosUserEnrollmentDeviceGroupID(v string)`

SetDefaultIosUserEnrollmentDeviceGroupID sets DefaultIosUserEnrollmentDeviceGroupID field to given value.

### HasDefaultIosUserEnrollmentDeviceGroupID

`func (o *AppleMDM) HasDefaultIosUserEnrollmentDeviceGroupID() bool`

HasDefaultIosUserEnrollmentDeviceGroupID returns a boolean if a field has been set.

### GetDefaultSystemGroupID

`func (o *AppleMDM) GetDefaultSystemGroupID() string`

GetDefaultSystemGroupID returns the DefaultSystemGroupID field if non-nil, zero value otherwise.

### GetDefaultSystemGroupIDOk

`func (o *AppleMDM) GetDefaultSystemGroupIDOk() (*string, bool)`

GetDefaultSystemGroupIDOk returns a tuple with the DefaultSystemGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSystemGroupID

`func (o *AppleMDM) SetDefaultSystemGroupID(v string)`

SetDefaultSystemGroupID sets DefaultSystemGroupID field to given value.

### HasDefaultSystemGroupID

`func (o *AppleMDM) HasDefaultSystemGroupID() bool`

HasDefaultSystemGroupID returns a boolean if a field has been set.

### GetDep

`func (o *AppleMDM) GetDep() DEP`

GetDep returns the Dep field if non-nil, zero value otherwise.

### GetDepOk

`func (o *AppleMDM) GetDepOk() (*DEP, bool)`

GetDepOk returns a tuple with the Dep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDep

`func (o *AppleMDM) SetDep(v DEP)`

SetDep sets Dep field to given value.

### HasDep

`func (o *AppleMDM) HasDep() bool`

HasDep returns a boolean if a field has been set.

### GetDepAccessTokenExpiry

`func (o *AppleMDM) GetDepAccessTokenExpiry() string`

GetDepAccessTokenExpiry returns the DepAccessTokenExpiry field if non-nil, zero value otherwise.

### GetDepAccessTokenExpiryOk

`func (o *AppleMDM) GetDepAccessTokenExpiryOk() (*string, bool)`

GetDepAccessTokenExpiryOk returns a tuple with the DepAccessTokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepAccessTokenExpiry

`func (o *AppleMDM) SetDepAccessTokenExpiry(v string)`

SetDepAccessTokenExpiry sets DepAccessTokenExpiry field to given value.

### HasDepAccessTokenExpiry

`func (o *AppleMDM) HasDepAccessTokenExpiry() bool`

HasDepAccessTokenExpiry returns a boolean if a field has been set.

### GetDepServerTokenState

`func (o *AppleMDM) GetDepServerTokenState() string`

GetDepServerTokenState returns the DepServerTokenState field if non-nil, zero value otherwise.

### GetDepServerTokenStateOk

`func (o *AppleMDM) GetDepServerTokenStateOk() (*string, bool)`

GetDepServerTokenStateOk returns a tuple with the DepServerTokenState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepServerTokenState

`func (o *AppleMDM) SetDepServerTokenState(v string)`

SetDepServerTokenState sets DepServerTokenState field to given value.

### HasDepServerTokenState

`func (o *AppleMDM) HasDepServerTokenState() bool`

HasDepServerTokenState returns a boolean if a field has been set.

### GetId

`func (o *AppleMDM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppleMDM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppleMDM) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AppleMDM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppleMDM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppleMDM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppleMDM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *AppleMDM) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AppleMDM) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AppleMDM) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AppleMDM) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



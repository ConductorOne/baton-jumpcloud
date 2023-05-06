# JumpcloudGoogleEmmEnterprise

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDeviceEnrollment** | Pointer to **bool** |  | [optional] 
**ContactEmail** | Pointer to **string** | not logging because it contains PII non-sensitive information. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeviceGroupId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**OrganizationObjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewJumpcloudGoogleEmmEnterprise

`func NewJumpcloudGoogleEmmEnterprise() *JumpcloudGoogleEmmEnterprise`

NewJumpcloudGoogleEmmEnterprise instantiates a new JumpcloudGoogleEmmEnterprise object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJumpcloudGoogleEmmEnterpriseWithDefaults

`func NewJumpcloudGoogleEmmEnterpriseWithDefaults() *JumpcloudGoogleEmmEnterprise`

NewJumpcloudGoogleEmmEnterpriseWithDefaults instantiates a new JumpcloudGoogleEmmEnterprise object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDeviceEnrollment

`func (o *JumpcloudGoogleEmmEnterprise) GetAllowDeviceEnrollment() bool`

GetAllowDeviceEnrollment returns the AllowDeviceEnrollment field if non-nil, zero value otherwise.

### GetAllowDeviceEnrollmentOk

`func (o *JumpcloudGoogleEmmEnterprise) GetAllowDeviceEnrollmentOk() (*bool, bool)`

GetAllowDeviceEnrollmentOk returns a tuple with the AllowDeviceEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDeviceEnrollment

`func (o *JumpcloudGoogleEmmEnterprise) SetAllowDeviceEnrollment(v bool)`

SetAllowDeviceEnrollment sets AllowDeviceEnrollment field to given value.

### HasAllowDeviceEnrollment

`func (o *JumpcloudGoogleEmmEnterprise) HasAllowDeviceEnrollment() bool`

HasAllowDeviceEnrollment returns a boolean if a field has been set.

### GetContactEmail

`func (o *JumpcloudGoogleEmmEnterprise) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *JumpcloudGoogleEmmEnterprise) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *JumpcloudGoogleEmmEnterprise) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *JumpcloudGoogleEmmEnterprise) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *JumpcloudGoogleEmmEnterprise) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JumpcloudGoogleEmmEnterprise) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JumpcloudGoogleEmmEnterprise) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *JumpcloudGoogleEmmEnterprise) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeviceGroupId

`func (o *JumpcloudGoogleEmmEnterprise) GetDeviceGroupId() string`

GetDeviceGroupId returns the DeviceGroupId field if non-nil, zero value otherwise.

### GetDeviceGroupIdOk

`func (o *JumpcloudGoogleEmmEnterprise) GetDeviceGroupIdOk() (*string, bool)`

GetDeviceGroupIdOk returns a tuple with the DeviceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroupId

`func (o *JumpcloudGoogleEmmEnterprise) SetDeviceGroupId(v string)`

SetDeviceGroupId sets DeviceGroupId field to given value.

### HasDeviceGroupId

`func (o *JumpcloudGoogleEmmEnterprise) HasDeviceGroupId() bool`

HasDeviceGroupId returns a boolean if a field has been set.

### GetDisplayName

`func (o *JumpcloudGoogleEmmEnterprise) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *JumpcloudGoogleEmmEnterprise) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *JumpcloudGoogleEmmEnterprise) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *JumpcloudGoogleEmmEnterprise) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *JumpcloudGoogleEmmEnterprise) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JumpcloudGoogleEmmEnterprise) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JumpcloudGoogleEmmEnterprise) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JumpcloudGoogleEmmEnterprise) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectId

`func (o *JumpcloudGoogleEmmEnterprise) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *JumpcloudGoogleEmmEnterprise) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *JumpcloudGoogleEmmEnterprise) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *JumpcloudGoogleEmmEnterprise) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetOrganizationObjectId

`func (o *JumpcloudGoogleEmmEnterprise) GetOrganizationObjectId() string`

GetOrganizationObjectId returns the OrganizationObjectId field if non-nil, zero value otherwise.

### GetOrganizationObjectIdOk

`func (o *JumpcloudGoogleEmmEnterprise) GetOrganizationObjectIdOk() (*string, bool)`

GetOrganizationObjectIdOk returns a tuple with the OrganizationObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationObjectId

`func (o *JumpcloudGoogleEmmEnterprise) SetOrganizationObjectId(v string)`

SetOrganizationObjectId sets OrganizationObjectId field to given value.

### HasOrganizationObjectId

`func (o *JumpcloudGoogleEmmEnterprise) HasOrganizationObjectId() bool`

HasOrganizationObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



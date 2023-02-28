# SoftwareAppSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUpdateDelay** | Pointer to **bool** |  | [optional] [default to false]
**AppleVpp** | Pointer to [**SoftwareAppAppleVpp**](SoftwareAppAppleVpp.md) |  | [optional] 
**AssetKind** | Pointer to **string** | The manifest asset kind (ex: software). | [optional] 
**AssetSha256Size** | Pointer to **int32** | The incremental size to use for summing the package as it is downloaded. | [optional] 
**AssetSha256Strings** | Pointer to **[]string** | The array of checksums, one each for the hash size up to the total size of the package. | [optional] 
**AutoUpdate** | Pointer to **bool** |  | [optional] [default to false]
**Description** | Pointer to **string** | The software app description. | [optional] 
**DesiredState** | Pointer to **string** | State of Install or Uninstall | [optional] 
**EnterpriseObjectId** | Pointer to **string** | ID of the Enterprise with which this app is associated | [optional] 
**GoogleAndroid** | Pointer to [**SoftwareAppGoogleAndroid**](SoftwareAppGoogleAndroid.md) |  | [optional] 
**Location** | Pointer to **string** | Repository where the app is located within the package manager | [optional] 
**LocationObjectId** | Pointer to **string** | ID of the repository where the app is located within the package manager | [optional] 
**PackageId** | Pointer to **string** |  | [optional] 
**PackageKind** | Pointer to **string** | The package manifest kind (ex: software-package). | [optional] 
**PackageManager** | Pointer to **string** | App store serving the app: APPLE_VPP, CHOCOLATEY, etc. | [optional] 
**PackageSubtitle** | Pointer to **string** | The package manifest subtitle. | [optional] 
**PackageVersion** | Pointer to **string** | The package manifest version. | [optional] 

## Methods

### NewSoftwareAppSettings

`func NewSoftwareAppSettings() *SoftwareAppSettings`

NewSoftwareAppSettings instantiates a new SoftwareAppSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareAppSettingsWithDefaults

`func NewSoftwareAppSettingsWithDefaults() *SoftwareAppSettings`

NewSoftwareAppSettingsWithDefaults instantiates a new SoftwareAppSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUpdateDelay

`func (o *SoftwareAppSettings) GetAllowUpdateDelay() bool`

GetAllowUpdateDelay returns the AllowUpdateDelay field if non-nil, zero value otherwise.

### GetAllowUpdateDelayOk

`func (o *SoftwareAppSettings) GetAllowUpdateDelayOk() (*bool, bool)`

GetAllowUpdateDelayOk returns a tuple with the AllowUpdateDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdateDelay

`func (o *SoftwareAppSettings) SetAllowUpdateDelay(v bool)`

SetAllowUpdateDelay sets AllowUpdateDelay field to given value.

### HasAllowUpdateDelay

`func (o *SoftwareAppSettings) HasAllowUpdateDelay() bool`

HasAllowUpdateDelay returns a boolean if a field has been set.

### GetAppleVpp

`func (o *SoftwareAppSettings) GetAppleVpp() SoftwareAppAppleVpp`

GetAppleVpp returns the AppleVpp field if non-nil, zero value otherwise.

### GetAppleVppOk

`func (o *SoftwareAppSettings) GetAppleVppOk() (*SoftwareAppAppleVpp, bool)`

GetAppleVppOk returns a tuple with the AppleVpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleVpp

`func (o *SoftwareAppSettings) SetAppleVpp(v SoftwareAppAppleVpp)`

SetAppleVpp sets AppleVpp field to given value.

### HasAppleVpp

`func (o *SoftwareAppSettings) HasAppleVpp() bool`

HasAppleVpp returns a boolean if a field has been set.

### GetAssetKind

`func (o *SoftwareAppSettings) GetAssetKind() string`

GetAssetKind returns the AssetKind field if non-nil, zero value otherwise.

### GetAssetKindOk

`func (o *SoftwareAppSettings) GetAssetKindOk() (*string, bool)`

GetAssetKindOk returns a tuple with the AssetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetKind

`func (o *SoftwareAppSettings) SetAssetKind(v string)`

SetAssetKind sets AssetKind field to given value.

### HasAssetKind

`func (o *SoftwareAppSettings) HasAssetKind() bool`

HasAssetKind returns a boolean if a field has been set.

### GetAssetSha256Size

`func (o *SoftwareAppSettings) GetAssetSha256Size() int32`

GetAssetSha256Size returns the AssetSha256Size field if non-nil, zero value otherwise.

### GetAssetSha256SizeOk

`func (o *SoftwareAppSettings) GetAssetSha256SizeOk() (*int32, bool)`

GetAssetSha256SizeOk returns a tuple with the AssetSha256Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSha256Size

`func (o *SoftwareAppSettings) SetAssetSha256Size(v int32)`

SetAssetSha256Size sets AssetSha256Size field to given value.

### HasAssetSha256Size

`func (o *SoftwareAppSettings) HasAssetSha256Size() bool`

HasAssetSha256Size returns a boolean if a field has been set.

### GetAssetSha256Strings

`func (o *SoftwareAppSettings) GetAssetSha256Strings() []string`

GetAssetSha256Strings returns the AssetSha256Strings field if non-nil, zero value otherwise.

### GetAssetSha256StringsOk

`func (o *SoftwareAppSettings) GetAssetSha256StringsOk() (*[]string, bool)`

GetAssetSha256StringsOk returns a tuple with the AssetSha256Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSha256Strings

`func (o *SoftwareAppSettings) SetAssetSha256Strings(v []string)`

SetAssetSha256Strings sets AssetSha256Strings field to given value.

### HasAssetSha256Strings

`func (o *SoftwareAppSettings) HasAssetSha256Strings() bool`

HasAssetSha256Strings returns a boolean if a field has been set.

### GetAutoUpdate

`func (o *SoftwareAppSettings) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *SoftwareAppSettings) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *SoftwareAppSettings) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *SoftwareAppSettings) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetDescription

`func (o *SoftwareAppSettings) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SoftwareAppSettings) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SoftwareAppSettings) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SoftwareAppSettings) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *SoftwareAppSettings) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *SoftwareAppSettings) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *SoftwareAppSettings) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *SoftwareAppSettings) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetEnterpriseObjectId

`func (o *SoftwareAppSettings) GetEnterpriseObjectId() string`

GetEnterpriseObjectId returns the EnterpriseObjectId field if non-nil, zero value otherwise.

### GetEnterpriseObjectIdOk

`func (o *SoftwareAppSettings) GetEnterpriseObjectIdOk() (*string, bool)`

GetEnterpriseObjectIdOk returns a tuple with the EnterpriseObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseObjectId

`func (o *SoftwareAppSettings) SetEnterpriseObjectId(v string)`

SetEnterpriseObjectId sets EnterpriseObjectId field to given value.

### HasEnterpriseObjectId

`func (o *SoftwareAppSettings) HasEnterpriseObjectId() bool`

HasEnterpriseObjectId returns a boolean if a field has been set.

### GetGoogleAndroid

`func (o *SoftwareAppSettings) GetGoogleAndroid() SoftwareAppGoogleAndroid`

GetGoogleAndroid returns the GoogleAndroid field if non-nil, zero value otherwise.

### GetGoogleAndroidOk

`func (o *SoftwareAppSettings) GetGoogleAndroidOk() (*SoftwareAppGoogleAndroid, bool)`

GetGoogleAndroidOk returns a tuple with the GoogleAndroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAndroid

`func (o *SoftwareAppSettings) SetGoogleAndroid(v SoftwareAppGoogleAndroid)`

SetGoogleAndroid sets GoogleAndroid field to given value.

### HasGoogleAndroid

`func (o *SoftwareAppSettings) HasGoogleAndroid() bool`

HasGoogleAndroid returns a boolean if a field has been set.

### GetLocation

`func (o *SoftwareAppSettings) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SoftwareAppSettings) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SoftwareAppSettings) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SoftwareAppSettings) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocationObjectId

`func (o *SoftwareAppSettings) GetLocationObjectId() string`

GetLocationObjectId returns the LocationObjectId field if non-nil, zero value otherwise.

### GetLocationObjectIdOk

`func (o *SoftwareAppSettings) GetLocationObjectIdOk() (*string, bool)`

GetLocationObjectIdOk returns a tuple with the LocationObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationObjectId

`func (o *SoftwareAppSettings) SetLocationObjectId(v string)`

SetLocationObjectId sets LocationObjectId field to given value.

### HasLocationObjectId

`func (o *SoftwareAppSettings) HasLocationObjectId() bool`

HasLocationObjectId returns a boolean if a field has been set.

### GetPackageId

`func (o *SoftwareAppSettings) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *SoftwareAppSettings) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *SoftwareAppSettings) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *SoftwareAppSettings) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPackageKind

`func (o *SoftwareAppSettings) GetPackageKind() string`

GetPackageKind returns the PackageKind field if non-nil, zero value otherwise.

### GetPackageKindOk

`func (o *SoftwareAppSettings) GetPackageKindOk() (*string, bool)`

GetPackageKindOk returns a tuple with the PackageKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageKind

`func (o *SoftwareAppSettings) SetPackageKind(v string)`

SetPackageKind sets PackageKind field to given value.

### HasPackageKind

`func (o *SoftwareAppSettings) HasPackageKind() bool`

HasPackageKind returns a boolean if a field has been set.

### GetPackageManager

`func (o *SoftwareAppSettings) GetPackageManager() string`

GetPackageManager returns the PackageManager field if non-nil, zero value otherwise.

### GetPackageManagerOk

`func (o *SoftwareAppSettings) GetPackageManagerOk() (*string, bool)`

GetPackageManagerOk returns a tuple with the PackageManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManager

`func (o *SoftwareAppSettings) SetPackageManager(v string)`

SetPackageManager sets PackageManager field to given value.

### HasPackageManager

`func (o *SoftwareAppSettings) HasPackageManager() bool`

HasPackageManager returns a boolean if a field has been set.

### GetPackageSubtitle

`func (o *SoftwareAppSettings) GetPackageSubtitle() string`

GetPackageSubtitle returns the PackageSubtitle field if non-nil, zero value otherwise.

### GetPackageSubtitleOk

`func (o *SoftwareAppSettings) GetPackageSubtitleOk() (*string, bool)`

GetPackageSubtitleOk returns a tuple with the PackageSubtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSubtitle

`func (o *SoftwareAppSettings) SetPackageSubtitle(v string)`

SetPackageSubtitle sets PackageSubtitle field to given value.

### HasPackageSubtitle

`func (o *SoftwareAppSettings) HasPackageSubtitle() bool`

HasPackageSubtitle returns a boolean if a field has been set.

### GetPackageVersion

`func (o *SoftwareAppSettings) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *SoftwareAppSettings) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *SoftwareAppSettings) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *SoftwareAppSettings) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



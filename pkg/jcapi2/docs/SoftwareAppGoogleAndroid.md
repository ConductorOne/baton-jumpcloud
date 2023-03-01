# SoftwareAppGoogleAndroid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppPricing** | Pointer to **string** | Whether this app is free, free with in-app purchases, or paid. | [optional] 
**AppVersion** | Pointer to **string** | Latest version currently available for this app. | [optional] 
**Author** | Pointer to **string** | The name of the author of this app. | [optional] 
**AutoUpdateMode** | Pointer to **string** | Controls the auto-update mode for the app. | [optional] 
**Category** | Pointer to **string** | The app category (e.g. COMMUNICATION, SOCIAL, etc.). | [optional] 
**ContentRating** | Pointer to **string** | The content rating for this app. | [optional] 
**DisplayMode** | Pointer to **string** | The display mode of the web app. | [optional] 
**DistributionChannel** | Pointer to **string** | How and to whom the package is made available. | [optional] 
**FullDescription** | Pointer to **string** | Full app description, if available. | [optional] 
**IconUrl** | Pointer to **string** | A link to an image that can be used as an icon for the app. | [optional] 
**InstallType** | Pointer to **string** | The type of installation to perform for an app. | [optional] 
**ManagedConfigurationTemplateId** | Pointer to **string** | The managed configurations template for the app. | [optional] 
**ManagedProperties** | Pointer to **bool** | Indicates whether this app has managed properties or not. | [optional] 
**MinSdkVersion** | Pointer to **int32** | The minimum Android SDK necessary to run the app. | [optional] 
**Name** | Pointer to **string** | The name of the app in the form enterprises/{enterprise}/applications/{packageName}. | [optional] 
**PermissionGrants** | Pointer to [**[]SoftwareAppPermissionGrants**](SoftwareAppPermissionGrants.md) |  | [optional] 
**RuntimePermission** | Pointer to **string** | The policy for granting permission requests to apps. | [optional] 
**StartUrl** | Pointer to **string** | The start URL, i.e. the URL that should load when the user opens the application. Applicable only for webapps. | [optional] 
**Type** | Pointer to **string** | Type of this android application. | [optional] 
**UpdateTime** | Pointer to **string** | The approximate time (within 7 days) the app was last published. | [optional] 
**VersionCode** | Pointer to **int32** | The current version of the web app. | [optional] 

## Methods

### NewSoftwareAppGoogleAndroid

`func NewSoftwareAppGoogleAndroid() *SoftwareAppGoogleAndroid`

NewSoftwareAppGoogleAndroid instantiates a new SoftwareAppGoogleAndroid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareAppGoogleAndroidWithDefaults

`func NewSoftwareAppGoogleAndroidWithDefaults() *SoftwareAppGoogleAndroid`

NewSoftwareAppGoogleAndroidWithDefaults instantiates a new SoftwareAppGoogleAndroid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPricing

`func (o *SoftwareAppGoogleAndroid) GetAppPricing() string`

GetAppPricing returns the AppPricing field if non-nil, zero value otherwise.

### GetAppPricingOk

`func (o *SoftwareAppGoogleAndroid) GetAppPricingOk() (*string, bool)`

GetAppPricingOk returns a tuple with the AppPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPricing

`func (o *SoftwareAppGoogleAndroid) SetAppPricing(v string)`

SetAppPricing sets AppPricing field to given value.

### HasAppPricing

`func (o *SoftwareAppGoogleAndroid) HasAppPricing() bool`

HasAppPricing returns a boolean if a field has been set.

### GetAppVersion

`func (o *SoftwareAppGoogleAndroid) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *SoftwareAppGoogleAndroid) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *SoftwareAppGoogleAndroid) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *SoftwareAppGoogleAndroid) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetAuthor

`func (o *SoftwareAppGoogleAndroid) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SoftwareAppGoogleAndroid) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SoftwareAppGoogleAndroid) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SoftwareAppGoogleAndroid) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAutoUpdateMode

`func (o *SoftwareAppGoogleAndroid) GetAutoUpdateMode() string`

GetAutoUpdateMode returns the AutoUpdateMode field if non-nil, zero value otherwise.

### GetAutoUpdateModeOk

`func (o *SoftwareAppGoogleAndroid) GetAutoUpdateModeOk() (*string, bool)`

GetAutoUpdateModeOk returns a tuple with the AutoUpdateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateMode

`func (o *SoftwareAppGoogleAndroid) SetAutoUpdateMode(v string)`

SetAutoUpdateMode sets AutoUpdateMode field to given value.

### HasAutoUpdateMode

`func (o *SoftwareAppGoogleAndroid) HasAutoUpdateMode() bool`

HasAutoUpdateMode returns a boolean if a field has been set.

### GetCategory

`func (o *SoftwareAppGoogleAndroid) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SoftwareAppGoogleAndroid) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SoftwareAppGoogleAndroid) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SoftwareAppGoogleAndroid) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetContentRating

`func (o *SoftwareAppGoogleAndroid) GetContentRating() string`

GetContentRating returns the ContentRating field if non-nil, zero value otherwise.

### GetContentRatingOk

`func (o *SoftwareAppGoogleAndroid) GetContentRatingOk() (*string, bool)`

GetContentRatingOk returns a tuple with the ContentRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRating

`func (o *SoftwareAppGoogleAndroid) SetContentRating(v string)`

SetContentRating sets ContentRating field to given value.

### HasContentRating

`func (o *SoftwareAppGoogleAndroid) HasContentRating() bool`

HasContentRating returns a boolean if a field has been set.

### GetDisplayMode

`func (o *SoftwareAppGoogleAndroid) GetDisplayMode() string`

GetDisplayMode returns the DisplayMode field if non-nil, zero value otherwise.

### GetDisplayModeOk

`func (o *SoftwareAppGoogleAndroid) GetDisplayModeOk() (*string, bool)`

GetDisplayModeOk returns a tuple with the DisplayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMode

`func (o *SoftwareAppGoogleAndroid) SetDisplayMode(v string)`

SetDisplayMode sets DisplayMode field to given value.

### HasDisplayMode

`func (o *SoftwareAppGoogleAndroid) HasDisplayMode() bool`

HasDisplayMode returns a boolean if a field has been set.

### GetDistributionChannel

`func (o *SoftwareAppGoogleAndroid) GetDistributionChannel() string`

GetDistributionChannel returns the DistributionChannel field if non-nil, zero value otherwise.

### GetDistributionChannelOk

`func (o *SoftwareAppGoogleAndroid) GetDistributionChannelOk() (*string, bool)`

GetDistributionChannelOk returns a tuple with the DistributionChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionChannel

`func (o *SoftwareAppGoogleAndroid) SetDistributionChannel(v string)`

SetDistributionChannel sets DistributionChannel field to given value.

### HasDistributionChannel

`func (o *SoftwareAppGoogleAndroid) HasDistributionChannel() bool`

HasDistributionChannel returns a boolean if a field has been set.

### GetFullDescription

`func (o *SoftwareAppGoogleAndroid) GetFullDescription() string`

GetFullDescription returns the FullDescription field if non-nil, zero value otherwise.

### GetFullDescriptionOk

`func (o *SoftwareAppGoogleAndroid) GetFullDescriptionOk() (*string, bool)`

GetFullDescriptionOk returns a tuple with the FullDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDescription

`func (o *SoftwareAppGoogleAndroid) SetFullDescription(v string)`

SetFullDescription sets FullDescription field to given value.

### HasFullDescription

`func (o *SoftwareAppGoogleAndroid) HasFullDescription() bool`

HasFullDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *SoftwareAppGoogleAndroid) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *SoftwareAppGoogleAndroid) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *SoftwareAppGoogleAndroid) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *SoftwareAppGoogleAndroid) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetInstallType

`func (o *SoftwareAppGoogleAndroid) GetInstallType() string`

GetInstallType returns the InstallType field if non-nil, zero value otherwise.

### GetInstallTypeOk

`func (o *SoftwareAppGoogleAndroid) GetInstallTypeOk() (*string, bool)`

GetInstallTypeOk returns a tuple with the InstallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallType

`func (o *SoftwareAppGoogleAndroid) SetInstallType(v string)`

SetInstallType sets InstallType field to given value.

### HasInstallType

`func (o *SoftwareAppGoogleAndroid) HasInstallType() bool`

HasInstallType returns a boolean if a field has been set.

### GetManagedConfigurationTemplateId

`func (o *SoftwareAppGoogleAndroid) GetManagedConfigurationTemplateId() string`

GetManagedConfigurationTemplateId returns the ManagedConfigurationTemplateId field if non-nil, zero value otherwise.

### GetManagedConfigurationTemplateIdOk

`func (o *SoftwareAppGoogleAndroid) GetManagedConfigurationTemplateIdOk() (*string, bool)`

GetManagedConfigurationTemplateIdOk returns a tuple with the ManagedConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedConfigurationTemplateId

`func (o *SoftwareAppGoogleAndroid) SetManagedConfigurationTemplateId(v string)`

SetManagedConfigurationTemplateId sets ManagedConfigurationTemplateId field to given value.

### HasManagedConfigurationTemplateId

`func (o *SoftwareAppGoogleAndroid) HasManagedConfigurationTemplateId() bool`

HasManagedConfigurationTemplateId returns a boolean if a field has been set.

### GetManagedProperties

`func (o *SoftwareAppGoogleAndroid) GetManagedProperties() bool`

GetManagedProperties returns the ManagedProperties field if non-nil, zero value otherwise.

### GetManagedPropertiesOk

`func (o *SoftwareAppGoogleAndroid) GetManagedPropertiesOk() (*bool, bool)`

GetManagedPropertiesOk returns a tuple with the ManagedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedProperties

`func (o *SoftwareAppGoogleAndroid) SetManagedProperties(v bool)`

SetManagedProperties sets ManagedProperties field to given value.

### HasManagedProperties

`func (o *SoftwareAppGoogleAndroid) HasManagedProperties() bool`

HasManagedProperties returns a boolean if a field has been set.

### GetMinSdkVersion

`func (o *SoftwareAppGoogleAndroid) GetMinSdkVersion() int32`

GetMinSdkVersion returns the MinSdkVersion field if non-nil, zero value otherwise.

### GetMinSdkVersionOk

`func (o *SoftwareAppGoogleAndroid) GetMinSdkVersionOk() (*int32, bool)`

GetMinSdkVersionOk returns a tuple with the MinSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSdkVersion

`func (o *SoftwareAppGoogleAndroid) SetMinSdkVersion(v int32)`

SetMinSdkVersion sets MinSdkVersion field to given value.

### HasMinSdkVersion

`func (o *SoftwareAppGoogleAndroid) HasMinSdkVersion() bool`

HasMinSdkVersion returns a boolean if a field has been set.

### GetName

`func (o *SoftwareAppGoogleAndroid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SoftwareAppGoogleAndroid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SoftwareAppGoogleAndroid) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SoftwareAppGoogleAndroid) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissionGrants

`func (o *SoftwareAppGoogleAndroid) GetPermissionGrants() []SoftwareAppPermissionGrants`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *SoftwareAppGoogleAndroid) GetPermissionGrantsOk() (*[]SoftwareAppPermissionGrants, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *SoftwareAppGoogleAndroid) SetPermissionGrants(v []SoftwareAppPermissionGrants)`

SetPermissionGrants sets PermissionGrants field to given value.

### HasPermissionGrants

`func (o *SoftwareAppGoogleAndroid) HasPermissionGrants() bool`

HasPermissionGrants returns a boolean if a field has been set.

### GetRuntimePermission

`func (o *SoftwareAppGoogleAndroid) GetRuntimePermission() string`

GetRuntimePermission returns the RuntimePermission field if non-nil, zero value otherwise.

### GetRuntimePermissionOk

`func (o *SoftwareAppGoogleAndroid) GetRuntimePermissionOk() (*string, bool)`

GetRuntimePermissionOk returns a tuple with the RuntimePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimePermission

`func (o *SoftwareAppGoogleAndroid) SetRuntimePermission(v string)`

SetRuntimePermission sets RuntimePermission field to given value.

### HasRuntimePermission

`func (o *SoftwareAppGoogleAndroid) HasRuntimePermission() bool`

HasRuntimePermission returns a boolean if a field has been set.

### GetStartUrl

`func (o *SoftwareAppGoogleAndroid) GetStartUrl() string`

GetStartUrl returns the StartUrl field if non-nil, zero value otherwise.

### GetStartUrlOk

`func (o *SoftwareAppGoogleAndroid) GetStartUrlOk() (*string, bool)`

GetStartUrlOk returns a tuple with the StartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUrl

`func (o *SoftwareAppGoogleAndroid) SetStartUrl(v string)`

SetStartUrl sets StartUrl field to given value.

### HasStartUrl

`func (o *SoftwareAppGoogleAndroid) HasStartUrl() bool`

HasStartUrl returns a boolean if a field has been set.

### GetType

`func (o *SoftwareAppGoogleAndroid) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SoftwareAppGoogleAndroid) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SoftwareAppGoogleAndroid) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SoftwareAppGoogleAndroid) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *SoftwareAppGoogleAndroid) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *SoftwareAppGoogleAndroid) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *SoftwareAppGoogleAndroid) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *SoftwareAppGoogleAndroid) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetVersionCode

`func (o *SoftwareAppGoogleAndroid) GetVersionCode() int32`

GetVersionCode returns the VersionCode field if non-nil, zero value otherwise.

### GetVersionCodeOk

`func (o *SoftwareAppGoogleAndroid) GetVersionCodeOk() (*int32, bool)`

GetVersionCodeOk returns a tuple with the VersionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCode

`func (o *SoftwareAppGoogleAndroid) SetVersionCode(v int32)`

SetVersionCode sets VersionCode field to given value.

### HasVersionCode

`func (o *SoftwareAppGoogleAndroid) HasVersionCode() bool`

HasVersionCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



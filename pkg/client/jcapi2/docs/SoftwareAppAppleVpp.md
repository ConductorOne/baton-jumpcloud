# SoftwareAppAppleVpp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConfiguration** | Pointer to **string** | Text sent to configure the application, the text should be a valid plist.  Returned only by &#39;GET /softwareapps/{id}&#39;. | [optional] 
**AssignedLicenses** | Pointer to **int32** |  | [optional] 
**AvailableLicenses** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** | App details returned by iTunes API. See example. The properties in this field are out of our control and we cannot guarantee consistency, so it should be checked by the client and manage the details accordingly. | [optional] 
**IsConfigEnabled** | Pointer to **bool** | Denotes if configuration has been enabled for the application.  Returned only by &#39;&#39;GET /softwareapps/{id}&#39;&#39;. | [optional] 
**SupportedDeviceFamilies** | Pointer to **[]string** | The supported device families for this VPP Application. | [optional] 
**TotalLicenses** | Pointer to **int32** |  | [optional] 

## Methods

### NewSoftwareAppAppleVpp

`func NewSoftwareAppAppleVpp() *SoftwareAppAppleVpp`

NewSoftwareAppAppleVpp instantiates a new SoftwareAppAppleVpp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareAppAppleVppWithDefaults

`func NewSoftwareAppAppleVppWithDefaults() *SoftwareAppAppleVpp`

NewSoftwareAppAppleVppWithDefaults instantiates a new SoftwareAppAppleVpp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConfiguration

`func (o *SoftwareAppAppleVpp) GetAppConfiguration() string`

GetAppConfiguration returns the AppConfiguration field if non-nil, zero value otherwise.

### GetAppConfigurationOk

`func (o *SoftwareAppAppleVpp) GetAppConfigurationOk() (*string, bool)`

GetAppConfigurationOk returns a tuple with the AppConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConfiguration

`func (o *SoftwareAppAppleVpp) SetAppConfiguration(v string)`

SetAppConfiguration sets AppConfiguration field to given value.

### HasAppConfiguration

`func (o *SoftwareAppAppleVpp) HasAppConfiguration() bool`

HasAppConfiguration returns a boolean if a field has been set.

### GetAssignedLicenses

`func (o *SoftwareAppAppleVpp) GetAssignedLicenses() int32`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *SoftwareAppAppleVpp) GetAssignedLicensesOk() (*int32, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedLicenses

`func (o *SoftwareAppAppleVpp) SetAssignedLicenses(v int32)`

SetAssignedLicenses sets AssignedLicenses field to given value.

### HasAssignedLicenses

`func (o *SoftwareAppAppleVpp) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### GetAvailableLicenses

`func (o *SoftwareAppAppleVpp) GetAvailableLicenses() int32`

GetAvailableLicenses returns the AvailableLicenses field if non-nil, zero value otherwise.

### GetAvailableLicensesOk

`func (o *SoftwareAppAppleVpp) GetAvailableLicensesOk() (*int32, bool)`

GetAvailableLicensesOk returns a tuple with the AvailableLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableLicenses

`func (o *SoftwareAppAppleVpp) SetAvailableLicenses(v int32)`

SetAvailableLicenses sets AvailableLicenses field to given value.

### HasAvailableLicenses

`func (o *SoftwareAppAppleVpp) HasAvailableLicenses() bool`

HasAvailableLicenses returns a boolean if a field has been set.

### GetDetails

`func (o *SoftwareAppAppleVpp) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SoftwareAppAppleVpp) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SoftwareAppAppleVpp) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SoftwareAppAppleVpp) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetIsConfigEnabled

`func (o *SoftwareAppAppleVpp) GetIsConfigEnabled() bool`

GetIsConfigEnabled returns the IsConfigEnabled field if non-nil, zero value otherwise.

### GetIsConfigEnabledOk

`func (o *SoftwareAppAppleVpp) GetIsConfigEnabledOk() (*bool, bool)`

GetIsConfigEnabledOk returns a tuple with the IsConfigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigEnabled

`func (o *SoftwareAppAppleVpp) SetIsConfigEnabled(v bool)`

SetIsConfigEnabled sets IsConfigEnabled field to given value.

### HasIsConfigEnabled

`func (o *SoftwareAppAppleVpp) HasIsConfigEnabled() bool`

HasIsConfigEnabled returns a boolean if a field has been set.

### GetSupportedDeviceFamilies

`func (o *SoftwareAppAppleVpp) GetSupportedDeviceFamilies() []string`

GetSupportedDeviceFamilies returns the SupportedDeviceFamilies field if non-nil, zero value otherwise.

### GetSupportedDeviceFamiliesOk

`func (o *SoftwareAppAppleVpp) GetSupportedDeviceFamiliesOk() (*[]string, bool)`

GetSupportedDeviceFamiliesOk returns a tuple with the SupportedDeviceFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDeviceFamilies

`func (o *SoftwareAppAppleVpp) SetSupportedDeviceFamilies(v []string)`

SetSupportedDeviceFamilies sets SupportedDeviceFamilies field to given value.

### HasSupportedDeviceFamilies

`func (o *SoftwareAppAppleVpp) HasSupportedDeviceFamilies() bool`

HasSupportedDeviceFamilies returns a boolean if a field has been set.

### GetTotalLicenses

`func (o *SoftwareAppAppleVpp) GetTotalLicenses() int32`

GetTotalLicenses returns the TotalLicenses field if non-nil, zero value otherwise.

### GetTotalLicensesOk

`func (o *SoftwareAppAppleVpp) GetTotalLicensesOk() (*int32, bool)`

GetTotalLicensesOk returns a tuple with the TotalLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLicenses

`func (o *SoftwareAppAppleVpp) SetTotalLicenses(v int32)`

SetTotalLicenses sets TotalLicenses field to given value.

### HasTotalLicenses

`func (o *SoftwareAppAppleVpp) HasTotalLicenses() bool`

HasTotalLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



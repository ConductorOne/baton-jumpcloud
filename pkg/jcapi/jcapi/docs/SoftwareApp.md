# SoftwareApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**[]SoftwareAppSettings**](SoftwareAppSettings.md) |  | [optional] 

## Methods

### NewSoftwareApp

`func NewSoftwareApp() *SoftwareApp`

NewSoftwareApp instantiates a new SoftwareApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareAppWithDefaults

`func NewSoftwareAppWithDefaults() *SoftwareApp`

NewSoftwareAppWithDefaults instantiates a new SoftwareApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *SoftwareApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SoftwareApp) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SoftwareApp) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SoftwareApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *SoftwareApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SoftwareApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SoftwareApp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SoftwareApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSettings

`func (o *SoftwareApp) GetSettings() []SoftwareAppSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SoftwareApp) GetSettingsOk() (*[]SoftwareAppSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SoftwareApp) SetSettings(v []SoftwareAppSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SoftwareApp) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TrustedappConfigGetTrustedAppsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the trusted application | 
**Path** | Pointer to **string** | Absolute path for the app&#39;s location in user&#39;s device | [optional] 
**Teamid** | Pointer to **string** | App&#39;s Team ID | [optional] 

## Methods

### NewTrustedappConfigGetTrustedAppsInner

`func NewTrustedappConfigGetTrustedAppsInner(name string, ) *TrustedappConfigGetTrustedAppsInner`

NewTrustedappConfigGetTrustedAppsInner instantiates a new TrustedappConfigGetTrustedAppsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedappConfigGetTrustedAppsInnerWithDefaults

`func NewTrustedappConfigGetTrustedAppsInnerWithDefaults() *TrustedappConfigGetTrustedAppsInner`

NewTrustedappConfigGetTrustedAppsInnerWithDefaults instantiates a new TrustedappConfigGetTrustedAppsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrustedappConfigGetTrustedAppsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustedappConfigGetTrustedAppsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustedappConfigGetTrustedAppsInner) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *TrustedappConfigGetTrustedAppsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TrustedappConfigGetTrustedAppsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TrustedappConfigGetTrustedAppsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TrustedappConfigGetTrustedAppsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTeamid

`func (o *TrustedappConfigGetTrustedAppsInner) GetTeamid() string`

GetTeamid returns the Teamid field if non-nil, zero value otherwise.

### GetTeamidOk

`func (o *TrustedappConfigGetTrustedAppsInner) GetTeamidOk() (*string, bool)`

GetTeamidOk returns a tuple with the Teamid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamid

`func (o *TrustedappConfigGetTrustedAppsInner) SetTeamid(v string)`

SetTeamid sets Teamid field to given value.

### HasTeamid

`func (o *TrustedappConfigGetTrustedAppsInner) HasTeamid() bool`

HasTeamid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



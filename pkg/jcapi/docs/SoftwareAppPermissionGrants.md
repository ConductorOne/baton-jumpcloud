# SoftwareAppPermissionGrants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An opaque string uniquely identifying the Android permission, e.g. android.permission.READ_CALENDAR. | [optional] 
**Policy** | Pointer to **string** | The policy for granting the permission. | [optional] 

## Methods

### NewSoftwareAppPermissionGrants

`func NewSoftwareAppPermissionGrants() *SoftwareAppPermissionGrants`

NewSoftwareAppPermissionGrants instantiates a new SoftwareAppPermissionGrants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareAppPermissionGrantsWithDefaults

`func NewSoftwareAppPermissionGrantsWithDefaults() *SoftwareAppPermissionGrants`

NewSoftwareAppPermissionGrantsWithDefaults instantiates a new SoftwareAppPermissionGrants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SoftwareAppPermissionGrants) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SoftwareAppPermissionGrants) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SoftwareAppPermissionGrants) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SoftwareAppPermissionGrants) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicy

`func (o *SoftwareAppPermissionGrants) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SoftwareAppPermissionGrants) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SoftwareAppPermissionGrants) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *SoftwareAppPermissionGrants) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



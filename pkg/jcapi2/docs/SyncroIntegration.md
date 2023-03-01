# SyncroIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier for this Syncro integration. | 
**IsMspAuthConfigured** | Pointer to **bool** | Has the msp-api been configured with auth data yet | [optional] 
**Subdomain** | **string** | The subdomain for the URL to connect to Syncro. | 

## Methods

### NewSyncroIntegration

`func NewSyncroIntegration(id string, subdomain string, ) *SyncroIntegration`

NewSyncroIntegration instantiates a new SyncroIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncroIntegrationWithDefaults

`func NewSyncroIntegrationWithDefaults() *SyncroIntegration`

NewSyncroIntegrationWithDefaults instantiates a new SyncroIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyncroIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncroIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncroIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetIsMspAuthConfigured

`func (o *SyncroIntegration) GetIsMspAuthConfigured() bool`

GetIsMspAuthConfigured returns the IsMspAuthConfigured field if non-nil, zero value otherwise.

### GetIsMspAuthConfiguredOk

`func (o *SyncroIntegration) GetIsMspAuthConfiguredOk() (*bool, bool)`

GetIsMspAuthConfiguredOk returns a tuple with the IsMspAuthConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMspAuthConfigured

`func (o *SyncroIntegration) SetIsMspAuthConfigured(v bool)`

SetIsMspAuthConfigured sets IsMspAuthConfigured field to given value.

### HasIsMspAuthConfigured

`func (o *SyncroIntegration) HasIsMspAuthConfigured() bool`

HasIsMspAuthConfigured returns a boolean if a field has been set.

### GetSubdomain

`func (o *SyncroIntegration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *SyncroIntegration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *SyncroIntegration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



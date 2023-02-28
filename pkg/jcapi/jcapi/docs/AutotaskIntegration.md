# AutotaskIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier for this Autotask integration. | 
**IsMspAuthConfigured** | Pointer to **bool** | Has the msp-api been configured with auth data yet | [optional] 
**Username** | **string** | The username for connecting to Autotask. | 

## Methods

### NewAutotaskIntegration

`func NewAutotaskIntegration(id string, username string, ) *AutotaskIntegration`

NewAutotaskIntegration instantiates a new AutotaskIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutotaskIntegrationWithDefaults

`func NewAutotaskIntegrationWithDefaults() *AutotaskIntegration`

NewAutotaskIntegrationWithDefaults instantiates a new AutotaskIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutotaskIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutotaskIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutotaskIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetIsMspAuthConfigured

`func (o *AutotaskIntegration) GetIsMspAuthConfigured() bool`

GetIsMspAuthConfigured returns the IsMspAuthConfigured field if non-nil, zero value otherwise.

### GetIsMspAuthConfiguredOk

`func (o *AutotaskIntegration) GetIsMspAuthConfiguredOk() (*bool, bool)`

GetIsMspAuthConfiguredOk returns a tuple with the IsMspAuthConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMspAuthConfigured

`func (o *AutotaskIntegration) SetIsMspAuthConfigured(v bool)`

SetIsMspAuthConfigured sets IsMspAuthConfigured field to given value.

### HasIsMspAuthConfigured

`func (o *AutotaskIntegration) HasIsMspAuthConfigured() bool`

HasIsMspAuthConfigured returns a boolean if a field has been set.

### GetUsername

`func (o *AutotaskIntegration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AutotaskIntegration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AutotaskIntegration) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



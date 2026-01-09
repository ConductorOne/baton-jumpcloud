# ConnectwiseIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **string** | The ConnectWise company identifier. | 
**Id** | **string** | The identifier for this ConnectWise integration. | 
**IsMspAuthConfigured** | Pointer to **bool** | Has the msp-api been configured with auth data yet | [optional] 
**Url** | **string** | The base url for connecting to ConnectWise. | 

## Methods

### NewConnectwiseIntegration

`func NewConnectwiseIntegration(companyId string, id string, url string, ) *ConnectwiseIntegration`

NewConnectwiseIntegration instantiates a new ConnectwiseIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectwiseIntegrationWithDefaults

`func NewConnectwiseIntegrationWithDefaults() *ConnectwiseIntegration`

NewConnectwiseIntegrationWithDefaults instantiates a new ConnectwiseIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ConnectwiseIntegration) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ConnectwiseIntegration) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ConnectwiseIntegration) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetId

`func (o *ConnectwiseIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectwiseIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectwiseIntegration) SetId(v string)`

SetId sets Id field to given value.


### GetIsMspAuthConfigured

`func (o *ConnectwiseIntegration) GetIsMspAuthConfigured() bool`

GetIsMspAuthConfigured returns the IsMspAuthConfigured field if non-nil, zero value otherwise.

### GetIsMspAuthConfiguredOk

`func (o *ConnectwiseIntegration) GetIsMspAuthConfiguredOk() (*bool, bool)`

GetIsMspAuthConfiguredOk returns a tuple with the IsMspAuthConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMspAuthConfigured

`func (o *ConnectwiseIntegration) SetIsMspAuthConfigured(v bool)`

SetIsMspAuthConfigured sets IsMspAuthConfigured field to given value.

### HasIsMspAuthConfigured

`func (o *ConnectwiseIntegration) HasIsMspAuthConfigured() bool`

HasIsMspAuthConfigured returns a boolean if a field has been set.

### GetUrl

`func (o *ConnectwiseIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConnectwiseIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConnectwiseIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ApplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsUrl** | Pointer to [**ApplicationConfigAcsUrl**](ApplicationConfigAcsUrl.md) |  | [optional] 
**ConstantAttributes** | Pointer to [**ApplicationConfigConstantAttributes**](ApplicationConfigConstantAttributes.md) |  | [optional] 
**DatabaseAttributes** | Pointer to [**ApplicationConfigDatabaseAttributes**](ApplicationConfigDatabaseAttributes.md) |  | [optional] 
**IdpCertificate** | Pointer to [**ApplicationConfigAcsUrl**](ApplicationConfigAcsUrl.md) |  | [optional] 
**IdpEntityId** | Pointer to [**ApplicationConfigAcsUrl**](ApplicationConfigAcsUrl.md) |  | [optional] 
**IdpPrivateKey** | Pointer to [**ApplicationConfigAcsUrl**](ApplicationConfigAcsUrl.md) |  | [optional] 
**SpEntityId** | Pointer to [**ApplicationConfigAcsUrl**](ApplicationConfigAcsUrl.md) |  | [optional] 

## Methods

### NewApplicationConfig

`func NewApplicationConfig() *ApplicationConfig`

NewApplicationConfig instantiates a new ApplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationConfigWithDefaults

`func NewApplicationConfigWithDefaults() *ApplicationConfig`

NewApplicationConfigWithDefaults instantiates a new ApplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsUrl

`func (o *ApplicationConfig) GetAcsUrl() ApplicationConfigAcsUrl`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *ApplicationConfig) GetAcsUrlOk() (*ApplicationConfigAcsUrl, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *ApplicationConfig) SetAcsUrl(v ApplicationConfigAcsUrl)`

SetAcsUrl sets AcsUrl field to given value.

### HasAcsUrl

`func (o *ApplicationConfig) HasAcsUrl() bool`

HasAcsUrl returns a boolean if a field has been set.

### GetConstantAttributes

`func (o *ApplicationConfig) GetConstantAttributes() ApplicationConfigConstantAttributes`

GetConstantAttributes returns the ConstantAttributes field if non-nil, zero value otherwise.

### GetConstantAttributesOk

`func (o *ApplicationConfig) GetConstantAttributesOk() (*ApplicationConfigConstantAttributes, bool)`

GetConstantAttributesOk returns a tuple with the ConstantAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstantAttributes

`func (o *ApplicationConfig) SetConstantAttributes(v ApplicationConfigConstantAttributes)`

SetConstantAttributes sets ConstantAttributes field to given value.

### HasConstantAttributes

`func (o *ApplicationConfig) HasConstantAttributes() bool`

HasConstantAttributes returns a boolean if a field has been set.

### GetDatabaseAttributes

`func (o *ApplicationConfig) GetDatabaseAttributes() ApplicationConfigDatabaseAttributes`

GetDatabaseAttributes returns the DatabaseAttributes field if non-nil, zero value otherwise.

### GetDatabaseAttributesOk

`func (o *ApplicationConfig) GetDatabaseAttributesOk() (*ApplicationConfigDatabaseAttributes, bool)`

GetDatabaseAttributesOk returns a tuple with the DatabaseAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseAttributes

`func (o *ApplicationConfig) SetDatabaseAttributes(v ApplicationConfigDatabaseAttributes)`

SetDatabaseAttributes sets DatabaseAttributes field to given value.

### HasDatabaseAttributes

`func (o *ApplicationConfig) HasDatabaseAttributes() bool`

HasDatabaseAttributes returns a boolean if a field has been set.

### GetIdpCertificate

`func (o *ApplicationConfig) GetIdpCertificate() ApplicationConfigAcsUrl`

GetIdpCertificate returns the IdpCertificate field if non-nil, zero value otherwise.

### GetIdpCertificateOk

`func (o *ApplicationConfig) GetIdpCertificateOk() (*ApplicationConfigAcsUrl, bool)`

GetIdpCertificateOk returns a tuple with the IdpCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertificate

`func (o *ApplicationConfig) SetIdpCertificate(v ApplicationConfigAcsUrl)`

SetIdpCertificate sets IdpCertificate field to given value.

### HasIdpCertificate

`func (o *ApplicationConfig) HasIdpCertificate() bool`

HasIdpCertificate returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *ApplicationConfig) GetIdpEntityId() ApplicationConfigAcsUrl`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *ApplicationConfig) GetIdpEntityIdOk() (*ApplicationConfigAcsUrl, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *ApplicationConfig) SetIdpEntityId(v ApplicationConfigAcsUrl)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *ApplicationConfig) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetIdpPrivateKey

`func (o *ApplicationConfig) GetIdpPrivateKey() ApplicationConfigAcsUrl`

GetIdpPrivateKey returns the IdpPrivateKey field if non-nil, zero value otherwise.

### GetIdpPrivateKeyOk

`func (o *ApplicationConfig) GetIdpPrivateKeyOk() (*ApplicationConfigAcsUrl, bool)`

GetIdpPrivateKeyOk returns a tuple with the IdpPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpPrivateKey

`func (o *ApplicationConfig) SetIdpPrivateKey(v ApplicationConfigAcsUrl)`

SetIdpPrivateKey sets IdpPrivateKey field to given value.

### HasIdpPrivateKey

`func (o *ApplicationConfig) HasIdpPrivateKey() bool`

HasIdpPrivateKey returns a boolean if a field has been set.

### GetSpEntityId

`func (o *ApplicationConfig) GetSpEntityId() ApplicationConfigAcsUrl`

GetSpEntityId returns the SpEntityId field if non-nil, zero value otherwise.

### GetSpEntityIdOk

`func (o *ApplicationConfig) GetSpEntityIdOk() (*ApplicationConfigAcsUrl, bool)`

GetSpEntityIdOk returns a tuple with the SpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEntityId

`func (o *ApplicationConfig) SetSpEntityId(v ApplicationConfigAcsUrl)`

SetSpEntityId sets SpEntityId field to given value.

### HasSpEntityId

`func (o *ApplicationConfig) HasSpEntityId() bool`

HasSpEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



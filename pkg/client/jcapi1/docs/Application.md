# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Beta** | Pointer to **bool** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Config** | [**ApplicationConfig**](ApplicationConfig.md) |  | 
**Created** | Pointer to **string** |  | [optional] 
**DatabaseAttributes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**LearnMore** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to [**ApplicationLogo**](ApplicationLogo.md) |  | [optional] 
**Name** | **string** |  | 
**Organization** | Pointer to **string** |  | [optional] 
**Sso** | Pointer to [**Sso**](Sso.md) |  | [optional] 
**SsoUrl** | **string** |  | 

## Methods

### NewApplication

`func NewApplication(config ApplicationConfig, name string, ssoUrl string, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Application) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Application) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *Application) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Application) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Application) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Application) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBeta

`func (o *Application) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *Application) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *Application) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *Application) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetColor

`func (o *Application) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Application) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Application) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Application) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetConfig

`func (o *Application) GetConfig() ApplicationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Application) GetConfigOk() (*ApplicationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Application) SetConfig(v ApplicationConfig)`

SetConfig sets Config field to given value.


### GetCreated

`func (o *Application) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Application) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Application) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Application) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDatabaseAttributes

`func (o *Application) GetDatabaseAttributes() []map[string]interface{}`

GetDatabaseAttributes returns the DatabaseAttributes field if non-nil, zero value otherwise.

### GetDatabaseAttributesOk

`func (o *Application) GetDatabaseAttributesOk() (*[]map[string]interface{}, bool)`

GetDatabaseAttributesOk returns a tuple with the DatabaseAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseAttributes

`func (o *Application) SetDatabaseAttributes(v []map[string]interface{})`

SetDatabaseAttributes sets DatabaseAttributes field to given value.

### HasDatabaseAttributes

`func (o *Application) HasDatabaseAttributes() bool`

HasDatabaseAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *Application) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *Application) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *Application) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *Application) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetDisplayName

`func (o *Application) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Application) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Application) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Application) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLearnMore

`func (o *Application) GetLearnMore() string`

GetLearnMore returns the LearnMore field if non-nil, zero value otherwise.

### GetLearnMoreOk

`func (o *Application) GetLearnMoreOk() (*string, bool)`

GetLearnMoreOk returns a tuple with the LearnMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnMore

`func (o *Application) SetLearnMore(v string)`

SetLearnMore sets LearnMore field to given value.

### HasLearnMore

`func (o *Application) HasLearnMore() bool`

HasLearnMore returns a boolean if a field has been set.

### GetLogo

`func (o *Application) GetLogo() ApplicationLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Application) GetLogoOk() (*ApplicationLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Application) SetLogo(v ApplicationLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Application) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *Application) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Application) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Application) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Application) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSso

`func (o *Application) GetSso() Sso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *Application) GetSsoOk() (*Sso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *Application) SetSso(v Sso)`

SetSso sets Sso field to given value.

### HasSso

`func (o *Application) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetSsoUrl

`func (o *Application) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *Application) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *Application) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



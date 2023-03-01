# Applicationtemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Beta** | Pointer to **bool** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ApplicationConfig**](ApplicationConfig.md) |  | [optional] 
**DisplayLabel** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IsConfigured** | Pointer to **bool** |  | [optional] 
**Jit** | Pointer to [**ApplicationtemplateJit**](ApplicationtemplateJit.md) |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**LearnMore** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to [**ApplicationtemplateLogo**](ApplicationtemplateLogo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Oidc** | Pointer to [**ApplicationtemplateOidc**](ApplicationtemplateOidc.md) |  | [optional] 
**Provision** | Pointer to [**ApplicationtemplateProvision**](ApplicationtemplateProvision.md) |  | [optional] 
**Sso** | Pointer to [**Sso**](Sso.md) |  | [optional] 
**SsoUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Test** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationtemplate

`func NewApplicationtemplate() *Applicationtemplate`

NewApplicationtemplate instantiates a new Applicationtemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationtemplateWithDefaults

`func NewApplicationtemplateWithDefaults() *Applicationtemplate`

NewApplicationtemplateWithDefaults instantiates a new Applicationtemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Applicationtemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Applicationtemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Applicationtemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Applicationtemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *Applicationtemplate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Applicationtemplate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Applicationtemplate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Applicationtemplate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBeta

`func (o *Applicationtemplate) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *Applicationtemplate) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *Applicationtemplate) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *Applicationtemplate) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetColor

`func (o *Applicationtemplate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Applicationtemplate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Applicationtemplate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Applicationtemplate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetConfig

`func (o *Applicationtemplate) GetConfig() ApplicationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Applicationtemplate) GetConfigOk() (*ApplicationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Applicationtemplate) SetConfig(v ApplicationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Applicationtemplate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *Applicationtemplate) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *Applicationtemplate) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *Applicationtemplate) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *Applicationtemplate) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetDisplayName

`func (o *Applicationtemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Applicationtemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Applicationtemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Applicationtemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsConfigured

`func (o *Applicationtemplate) GetIsConfigured() bool`

GetIsConfigured returns the IsConfigured field if non-nil, zero value otherwise.

### GetIsConfiguredOk

`func (o *Applicationtemplate) GetIsConfiguredOk() (*bool, bool)`

GetIsConfiguredOk returns a tuple with the IsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigured

`func (o *Applicationtemplate) SetIsConfigured(v bool)`

SetIsConfigured sets IsConfigured field to given value.

### HasIsConfigured

`func (o *Applicationtemplate) HasIsConfigured() bool`

HasIsConfigured returns a boolean if a field has been set.

### GetJit

`func (o *Applicationtemplate) GetJit() ApplicationtemplateJit`

GetJit returns the Jit field if non-nil, zero value otherwise.

### GetJitOk

`func (o *Applicationtemplate) GetJitOk() (*ApplicationtemplateJit, bool)`

GetJitOk returns a tuple with the Jit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJit

`func (o *Applicationtemplate) SetJit(v ApplicationtemplateJit)`

SetJit sets Jit field to given value.

### HasJit

`func (o *Applicationtemplate) HasJit() bool`

HasJit returns a boolean if a field has been set.

### GetKeywords

`func (o *Applicationtemplate) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *Applicationtemplate) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *Applicationtemplate) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *Applicationtemplate) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLearnMore

`func (o *Applicationtemplate) GetLearnMore() string`

GetLearnMore returns the LearnMore field if non-nil, zero value otherwise.

### GetLearnMoreOk

`func (o *Applicationtemplate) GetLearnMoreOk() (*string, bool)`

GetLearnMoreOk returns a tuple with the LearnMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnMore

`func (o *Applicationtemplate) SetLearnMore(v string)`

SetLearnMore sets LearnMore field to given value.

### HasLearnMore

`func (o *Applicationtemplate) HasLearnMore() bool`

HasLearnMore returns a boolean if a field has been set.

### GetLogo

`func (o *Applicationtemplate) GetLogo() ApplicationtemplateLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Applicationtemplate) GetLogoOk() (*ApplicationtemplateLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Applicationtemplate) SetLogo(v ApplicationtemplateLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Applicationtemplate) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *Applicationtemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Applicationtemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Applicationtemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Applicationtemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOidc

`func (o *Applicationtemplate) GetOidc() ApplicationtemplateOidc`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *Applicationtemplate) GetOidcOk() (*ApplicationtemplateOidc, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *Applicationtemplate) SetOidc(v ApplicationtemplateOidc)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *Applicationtemplate) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### GetProvision

`func (o *Applicationtemplate) GetProvision() ApplicationtemplateProvision`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *Applicationtemplate) GetProvisionOk() (*ApplicationtemplateProvision, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *Applicationtemplate) SetProvision(v ApplicationtemplateProvision)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *Applicationtemplate) HasProvision() bool`

HasProvision returns a boolean if a field has been set.

### GetSso

`func (o *Applicationtemplate) GetSso() Sso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *Applicationtemplate) GetSsoOk() (*Sso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *Applicationtemplate) SetSso(v Sso)`

SetSso sets Sso field to given value.

### HasSso

`func (o *Applicationtemplate) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetSsoUrl

`func (o *Applicationtemplate) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *Applicationtemplate) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *Applicationtemplate) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *Applicationtemplate) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetStatus

`func (o *Applicationtemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Applicationtemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Applicationtemplate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Applicationtemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTest

`func (o *Applicationtemplate) GetTest() string`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *Applicationtemplate) GetTestOk() (*string, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *Applicationtemplate) SetTest(v string)`

SetTest sets Test field to given value.

### HasTest

`func (o *Applicationtemplate) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



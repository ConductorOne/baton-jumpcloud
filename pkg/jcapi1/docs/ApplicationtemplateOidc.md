# ApplicationtemplateOidc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantTypes** | Pointer to **[]string** | The grant types allowed. Currently only authorization_code is allowed. | [optional] 
**RedirectUris** | Pointer to **[]string** | List of allowed redirectUris | [optional] 
**SsoUrl** | Pointer to **string** | The relying party url to trigger an oidc login. | [optional] 
**TokenEndpointAuthMethod** | Pointer to **string** | Method that the client uses to authenticate when requesting a token. If &#39;none&#39;, then the client must use PKCE. If &#39;client_secret_post&#39;, then the secret is passed in the post body when requesting the token. | [optional] 

## Methods

### NewApplicationtemplateOidc

`func NewApplicationtemplateOidc() *ApplicationtemplateOidc`

NewApplicationtemplateOidc instantiates a new ApplicationtemplateOidc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationtemplateOidcWithDefaults

`func NewApplicationtemplateOidcWithDefaults() *ApplicationtemplateOidc`

NewApplicationtemplateOidcWithDefaults instantiates a new ApplicationtemplateOidc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantTypes

`func (o *ApplicationtemplateOidc) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *ApplicationtemplateOidc) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *ApplicationtemplateOidc) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *ApplicationtemplateOidc) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ApplicationtemplateOidc) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ApplicationtemplateOidc) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ApplicationtemplateOidc) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ApplicationtemplateOidc) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetSsoUrl

`func (o *ApplicationtemplateOidc) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *ApplicationtemplateOidc) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *ApplicationtemplateOidc) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *ApplicationtemplateOidc) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *ApplicationtemplateOidc) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *ApplicationtemplateOidc) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *ApplicationtemplateOidc) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *ApplicationtemplateOidc) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



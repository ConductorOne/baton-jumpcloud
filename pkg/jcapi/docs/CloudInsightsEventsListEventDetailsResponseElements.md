# CloudInsightsEventsListEventDetailsResponseElements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsoleLogin** | Pointer to **string** |  | [optional] 
**AssumedRoleUser** | Pointer to [**CloudInsightsEventsListEventDetailsResponseElementsAssumedRoleUser**](CloudInsightsEventsListEventDetailsResponseElementsAssumedRoleUser.md) |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**CloudInsightsEventsListEventDetailsResponseElementsCredentials**](CloudInsightsEventsListEventDetailsResponseElementsCredentials.md) |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**NameQualifier** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**SubjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudInsightsEventsListEventDetailsResponseElements

`func NewCloudInsightsEventsListEventDetailsResponseElements() *CloudInsightsEventsListEventDetailsResponseElements`

NewCloudInsightsEventsListEventDetailsResponseElements instantiates a new CloudInsightsEventsListEventDetailsResponseElements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsEventsListEventDetailsResponseElementsWithDefaults

`func NewCloudInsightsEventsListEventDetailsResponseElementsWithDefaults() *CloudInsightsEventsListEventDetailsResponseElements`

NewCloudInsightsEventsListEventDetailsResponseElementsWithDefaults instantiates a new CloudInsightsEventsListEventDetailsResponseElements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsoleLogin

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetConsoleLogin() string`

GetConsoleLogin returns the ConsoleLogin field if non-nil, zero value otherwise.

### GetConsoleLoginOk

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetConsoleLoginOk() (*string, bool)`

GetConsoleLoginOk returns a tuple with the ConsoleLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLogin

`func (o *CloudInsightsEventsListEventDetailsResponseElements) SetConsoleLogin(v string)`

SetConsoleLogin sets ConsoleLogin field to given value.

### HasConsoleLogin

`func (o *CloudInsightsEventsListEventDetailsResponseElements) HasConsoleLogin() bool`

HasConsoleLogin returns a boolean if a field has been set.

### GetAssumedRoleUser

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetAssumedRoleUser() CloudInsightsEventsListEventDetailsResponseElementsAssumedRoleUser`

GetAssumedRoleUser returns the AssumedRoleUser field if non-nil, zero value otherwise.

### GetAssumedRoleUserOk

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetAssumedRoleUserOk() (*CloudInsightsEventsListEventDetailsResponseElementsAssumedRoleUser, bool)`

GetAssumedRoleUserOk returns a tuple with the AssumedRoleUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumedRoleUser

`func (o *CloudInsightsEventsListEventDetailsResponseElements) SetAssumedRoleUser(v CloudInsightsEventsListEventDetailsResponseElementsAssumedRoleUser)`

SetAssumedRoleUser sets AssumedRoleUser field to given value.

### HasAssumedRoleUser

`func (o *CloudInsightsEventsListEventDetailsResponseElements) HasAssumedRoleUser() bool`

HasAssumedRoleUser returns a boolean if a field has been set.

### GetAudience

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CloudInsightsEventsListEventDetailsResponseElements) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CloudInsightsEventsListEventDetailsResponseElements) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetCredentials

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetCredentials() CloudInsightsEventsListEventDetailsResponseElementsCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetCredentialsOk() (*CloudInsightsEventsListEventDetailsResponseElementsCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CloudInsightsEventsListEventDetailsResponseElements) SetCredentials(v CloudInsightsEventsListEventDetailsResponseElementsCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CloudInsightsEventsListEventDetailsResponseElements) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetIssuer

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CloudInsightsEventsListEventDetailsResponseElements) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CloudInsightsEventsListEventDetailsResponseElements) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetNameQualifier

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetNameQualifier() string`

GetNameQualifier returns the NameQualifier field if non-nil, zero value otherwise.

### GetNameQualifierOk

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetNameQualifierOk() (*string, bool)`

GetNameQualifierOk returns a tuple with the NameQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameQualifier

`func (o *CloudInsightsEventsListEventDetailsResponseElements) SetNameQualifier(v string)`

SetNameQualifier sets NameQualifier field to given value.

### HasNameQualifier

`func (o *CloudInsightsEventsListEventDetailsResponseElements) HasNameQualifier() bool`

HasNameQualifier returns a boolean if a field has been set.

### GetSubject

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudInsightsEventsListEventDetailsResponseElements) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudInsightsEventsListEventDetailsResponseElements) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectType

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *CloudInsightsEventsListEventDetailsResponseElements) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *CloudInsightsEventsListEventDetailsResponseElements) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *CloudInsightsEventsListEventDetailsResponseElements) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CloudInsightsEventsListEventDetailsRequestParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**BucketName** | Pointer to **string** |  | [optional] 
**CertificateArn** | Pointer to **string** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**ComplianceTypes** | Pointer to **[]string** |  | [optional] 
**ConfigRuleName** | Pointer to **string** |  | [optional] 
**DurationSeconds** | Pointer to **float32** |  | [optional] 
**EncryptionAlgorithm** | Pointer to **string** |  | [optional] 
**EncryptionContext** | Pointer to [**CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext**](CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext.md) |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Include** | Pointer to **[]string** |  | [optional] 
**IncludeDeletedResponses** | Pointer to **bool** |  | [optional] 
**IncludePublic** | Pointer to **bool** |  | [optional] 
**IncludeShared** | Pointer to **bool** |  | [optional] 
**Includes** | Pointer to [**CloudInsightsEventsListEventDetailsRequestParametersIncludes**](CloudInsightsEventsListEventDetailsRequestParametersIncludes.md) |  | [optional] 
**KeyId** | Pointer to **string** |  | [optional] 
**KeySpec** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **float32** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**LogGroupName** | Pointer to **string** |  | [optional] 
**LogStreamName** | Pointer to **string** |  | [optional] 
**PrincipalArn** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ResultToken** | Pointer to **string** |  | [optional] 
**RoleArn** | Pointer to **string** |  | [optional] 
**RoleSessionName** | Pointer to **string** |  | [optional] 
**SAMLAssertionID** | Pointer to **string** |  | [optional] 
**Services** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TestMode** | Pointer to **bool** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudInsightsEventsListEventDetailsRequestParameters

`func NewCloudInsightsEventsListEventDetailsRequestParameters() *CloudInsightsEventsListEventDetailsRequestParameters`

NewCloudInsightsEventsListEventDetailsRequestParameters instantiates a new CloudInsightsEventsListEventDetailsRequestParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInsightsEventsListEventDetailsRequestParametersWithDefaults

`func NewCloudInsightsEventsListEventDetailsRequestParametersWithDefaults() *CloudInsightsEventsListEventDetailsRequestParameters`

NewCloudInsightsEventsListEventDetailsRequestParametersWithDefaults instantiates a new CloudInsightsEventsListEventDetailsRequestParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetBucketName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetCertificateArn

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetCertificateArn() string`

GetCertificateArn returns the CertificateArn field if non-nil, zero value otherwise.

### GetCertificateArnOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetCertificateArnOk() (*string, bool)`

GetCertificateArnOk returns a tuple with the CertificateArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateArn

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetCertificateArn(v string)`

SetCertificateArn sets CertificateArn field to given value.

### HasCertificateArn

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasCertificateArn() bool`

HasCertificateArn returns a boolean if a field has been set.

### GetCluster

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetComplianceTypes

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetComplianceTypes() []string`

GetComplianceTypes returns the ComplianceTypes field if non-nil, zero value otherwise.

### GetComplianceTypesOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetComplianceTypesOk() (*[]string, bool)`

GetComplianceTypesOk returns a tuple with the ComplianceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceTypes

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetComplianceTypes(v []string)`

SetComplianceTypes sets ComplianceTypes field to given value.

### HasComplianceTypes

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasComplianceTypes() bool`

HasComplianceTypes returns a boolean if a field has been set.

### GetConfigRuleName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetConfigRuleName() string`

GetConfigRuleName returns the ConfigRuleName field if non-nil, zero value otherwise.

### GetConfigRuleNameOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetConfigRuleNameOk() (*string, bool)`

GetConfigRuleNameOk returns a tuple with the ConfigRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigRuleName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetConfigRuleName(v string)`

SetConfigRuleName sets ConfigRuleName field to given value.

### HasConfigRuleName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasConfigRuleName() bool`

HasConfigRuleName returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetEncryptionContext

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetEncryptionContext() CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext`

GetEncryptionContext returns the EncryptionContext field if non-nil, zero value otherwise.

### GetEncryptionContextOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetEncryptionContextOk() (*CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext, bool)`

GetEncryptionContextOk returns a tuple with the EncryptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionContext

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetEncryptionContext(v CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext)`

SetEncryptionContext sets EncryptionContext field to given value.

### HasEncryptionContext

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasEncryptionContext() bool`

HasEncryptionContext returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInclude

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetIncludeDeletedResponses

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeDeletedResponses() bool`

GetIncludeDeletedResponses returns the IncludeDeletedResponses field if non-nil, zero value otherwise.

### GetIncludeDeletedResponsesOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeDeletedResponsesOk() (*bool, bool)`

GetIncludeDeletedResponsesOk returns a tuple with the IncludeDeletedResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeletedResponses

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetIncludeDeletedResponses(v bool)`

SetIncludeDeletedResponses sets IncludeDeletedResponses field to given value.

### HasIncludeDeletedResponses

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasIncludeDeletedResponses() bool`

HasIncludeDeletedResponses returns a boolean if a field has been set.

### GetIncludePublic

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludePublic() bool`

GetIncludePublic returns the IncludePublic field if non-nil, zero value otherwise.

### GetIncludePublicOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludePublicOk() (*bool, bool)`

GetIncludePublicOk returns a tuple with the IncludePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePublic

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetIncludePublic(v bool)`

SetIncludePublic sets IncludePublic field to given value.

### HasIncludePublic

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasIncludePublic() bool`

HasIncludePublic returns a boolean if a field has been set.

### GetIncludeShared

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeShared() bool`

GetIncludeShared returns the IncludeShared field if non-nil, zero value otherwise.

### GetIncludeSharedOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeSharedOk() (*bool, bool)`

GetIncludeSharedOk returns a tuple with the IncludeShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeShared

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetIncludeShared(v bool)`

SetIncludeShared sets IncludeShared field to given value.

### HasIncludeShared

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasIncludeShared() bool`

HasIncludeShared returns a boolean if a field has been set.

### GetIncludes

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludes() CloudInsightsEventsListEventDetailsRequestParametersIncludes`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludesOk() (*CloudInsightsEventsListEventDetailsRequestParametersIncludes, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetIncludes(v CloudInsightsEventsListEventDetailsRequestParametersIncludes)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetKeyId

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetKeySpec

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetKeySpec() string`

GetKeySpec returns the KeySpec field if non-nil, zero value otherwise.

### GetKeySpecOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetKeySpecOk() (*string, bool)`

GetKeySpecOk returns a tuple with the KeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySpec

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetKeySpec(v string)`

SetKeySpec sets KeySpec field to given value.

### HasKeySpec

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasKeySpec() bool`

HasKeySpec returns a boolean if a field has been set.

### GetLimit

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLocation

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLogGroupName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLogGroupName() string`

GetLogGroupName returns the LogGroupName field if non-nil, zero value otherwise.

### GetLogGroupNameOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLogGroupNameOk() (*string, bool)`

GetLogGroupNameOk returns a tuple with the LogGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogGroupName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetLogGroupName(v string)`

SetLogGroupName sets LogGroupName field to given value.

### HasLogGroupName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasLogGroupName() bool`

HasLogGroupName returns a boolean if a field has been set.

### GetLogStreamName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLogStreamName() string`

GetLogStreamName returns the LogStreamName field if non-nil, zero value otherwise.

### GetLogStreamNameOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLogStreamNameOk() (*string, bool)`

GetLogStreamNameOk returns a tuple with the LogStreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStreamName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetLogStreamName(v string)`

SetLogStreamName sets LogStreamName field to given value.

### HasLogStreamName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasLogStreamName() bool`

HasLogStreamName returns a boolean if a field has been set.

### GetPrincipalArn

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetPrincipalArn() string`

GetPrincipalArn returns the PrincipalArn field if non-nil, zero value otherwise.

### GetPrincipalArnOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetPrincipalArnOk() (*string, bool)`

GetPrincipalArnOk returns a tuple with the PrincipalArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalArn

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetPrincipalArn(v string)`

SetPrincipalArn sets PrincipalArn field to given value.

### HasPrincipalArn

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasPrincipalArn() bool`

HasPrincipalArn returns a boolean if a field has been set.

### GetResourceType

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResultToken

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetResultToken() string`

GetResultToken returns the ResultToken field if non-nil, zero value otherwise.

### GetResultTokenOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetResultTokenOk() (*string, bool)`

GetResultTokenOk returns a tuple with the ResultToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultToken

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetResultToken(v string)`

SetResultToken sets ResultToken field to given value.

### HasResultToken

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasResultToken() bool`

HasResultToken returns a boolean if a field has been set.

### GetRoleArn

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetRoleSessionName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetRoleSessionName() string`

GetRoleSessionName returns the RoleSessionName field if non-nil, zero value otherwise.

### GetRoleSessionNameOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetRoleSessionNameOk() (*string, bool)`

GetRoleSessionNameOk returns a tuple with the RoleSessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSessionName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetRoleSessionName(v string)`

SetRoleSessionName sets RoleSessionName field to given value.

### HasRoleSessionName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasRoleSessionName() bool`

HasRoleSessionName returns a boolean if a field has been set.

### GetSAMLAssertionID

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetSAMLAssertionID() string`

GetSAMLAssertionID returns the SAMLAssertionID field if non-nil, zero value otherwise.

### GetSAMLAssertionIDOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetSAMLAssertionIDOk() (*string, bool)`

GetSAMLAssertionIDOk returns a tuple with the SAMLAssertionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLAssertionID

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetSAMLAssertionID(v string)`

SetSAMLAssertionID sets SAMLAssertionID field to given value.

### HasSAMLAssertionID

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasSAMLAssertionID() bool`

HasSAMLAssertionID returns a boolean if a field has been set.

### GetServices

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStatus

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTestMode

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetTestMode() bool`

GetTestMode returns the TestMode field if non-nil, zero value otherwise.

### GetTestModeOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetTestModeOk() (*bool, bool)`

GetTestModeOk returns a tuple with the TestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMode

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetTestMode(v bool)`

SetTestMode sets TestMode field to given value.

### HasTestMode

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasTestMode() bool`

HasTestMode returns a boolean if a field has been set.

### GetUserName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



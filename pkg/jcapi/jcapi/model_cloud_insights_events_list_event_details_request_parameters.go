/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CloudInsightsEventsListEventDetailsRequestParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudInsightsEventsListEventDetailsRequestParameters{}

// CloudInsightsEventsListEventDetailsRequestParameters struct for CloudInsightsEventsListEventDetailsRequestParameters
type CloudInsightsEventsListEventDetailsRequestParameters struct {
	Host *string `json:"Host,omitempty"`
	BucketName *string `json:"bucketName,omitempty"`
	CertificateArn *string `json:"certificateArn,omitempty"`
	Cluster *string `json:"cluster,omitempty"`
	ComplianceTypes []string `json:"complianceTypes,omitempty"`
	ConfigRuleName *string `json:"configRuleName,omitempty"`
	DurationSeconds *float32 `json:"durationSeconds,omitempty"`
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty"`
	EncryptionContext *CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext `json:"encryptionContext,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	Include []string `json:"include,omitempty"`
	IncludeDeletedResponses *bool `json:"includeDeletedResponses,omitempty"`
	IncludePublic *bool `json:"includePublic,omitempty"`
	IncludeShared *bool `json:"includeShared,omitempty"`
	Includes *CloudInsightsEventsListEventDetailsRequestParametersIncludes `json:"includes,omitempty"`
	KeyId *string `json:"keyId,omitempty"`
	KeySpec *string `json:"keySpec,omitempty"`
	Limit *float32 `json:"limit,omitempty"`
	Location *string `json:"location,omitempty"`
	LogGroupName *string `json:"logGroupName,omitempty"`
	LogStreamName *string `json:"logStreamName,omitempty"`
	PrincipalArn *string `json:"principalArn,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ResultToken *string `json:"resultToken,omitempty"`
	RoleArn *string `json:"roleArn,omitempty"`
	RoleSessionName *string `json:"roleSessionName,omitempty"`
	SAMLAssertionID *string `json:"sAMLAssertionID,omitempty"`
	Services []string `json:"services,omitempty"`
	Status *string `json:"status,omitempty"`
	TestMode *bool `json:"testMode,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// NewCloudInsightsEventsListEventDetailsRequestParameters instantiates a new CloudInsightsEventsListEventDetailsRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudInsightsEventsListEventDetailsRequestParameters() *CloudInsightsEventsListEventDetailsRequestParameters {
	this := CloudInsightsEventsListEventDetailsRequestParameters{}
	return &this
}

// NewCloudInsightsEventsListEventDetailsRequestParametersWithDefaults instantiates a new CloudInsightsEventsListEventDetailsRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudInsightsEventsListEventDetailsRequestParametersWithDefaults() *CloudInsightsEventsListEventDetailsRequestParameters {
	this := CloudInsightsEventsListEventDetailsRequestParameters{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetHost(v string) {
	o.Host = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetBucketName(v string) {
	o.BucketName = &v
}

// GetCertificateArn returns the CertificateArn field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetCertificateArn() string {
	if o == nil || IsNil(o.CertificateArn) {
		var ret string
		return ret
	}
	return *o.CertificateArn
}

// GetCertificateArnOk returns a tuple with the CertificateArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetCertificateArnOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateArn) {
		return nil, false
	}
	return o.CertificateArn, true
}

// HasCertificateArn returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasCertificateArn() bool {
	if o != nil && !IsNil(o.CertificateArn) {
		return true
	}

	return false
}

// SetCertificateArn gets a reference to the given string and assigns it to the CertificateArn field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetCertificateArn(v string) {
	o.CertificateArn = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetCluster(v string) {
	o.Cluster = &v
}

// GetComplianceTypes returns the ComplianceTypes field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetComplianceTypes() []string {
	if o == nil || IsNil(o.ComplianceTypes) {
		var ret []string
		return ret
	}
	return o.ComplianceTypes
}

// GetComplianceTypesOk returns a tuple with the ComplianceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetComplianceTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ComplianceTypes) {
		return nil, false
	}
	return o.ComplianceTypes, true
}

// HasComplianceTypes returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasComplianceTypes() bool {
	if o != nil && !IsNil(o.ComplianceTypes) {
		return true
	}

	return false
}

// SetComplianceTypes gets a reference to the given []string and assigns it to the ComplianceTypes field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetComplianceTypes(v []string) {
	o.ComplianceTypes = v
}

// GetConfigRuleName returns the ConfigRuleName field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetConfigRuleName() string {
	if o == nil || IsNil(o.ConfigRuleName) {
		var ret string
		return ret
	}
	return *o.ConfigRuleName
}

// GetConfigRuleNameOk returns a tuple with the ConfigRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetConfigRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigRuleName) {
		return nil, false
	}
	return o.ConfigRuleName, true
}

// HasConfigRuleName returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasConfigRuleName() bool {
	if o != nil && !IsNil(o.ConfigRuleName) {
		return true
	}

	return false
}

// SetConfigRuleName gets a reference to the given string and assigns it to the ConfigRuleName field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetConfigRuleName(v string) {
	o.ConfigRuleName = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetDurationSeconds() float32 {
	if o == nil || IsNil(o.DurationSeconds) {
		var ret float32
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetDurationSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.DurationSeconds) {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasDurationSeconds() bool {
	if o != nil && !IsNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given float32 and assigns it to the DurationSeconds field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetDurationSeconds(v float32) {
	o.DurationSeconds = &v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetEncryptionAlgorithm() string {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		var ret string
		return ret
	}
	return *o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetEncryptionAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		return nil, false
	}
	return o.EncryptionAlgorithm, true
}

// HasEncryptionAlgorithm returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.EncryptionAlgorithm) {
		return true
	}

	return false
}

// SetEncryptionAlgorithm gets a reference to the given string and assigns it to the EncryptionAlgorithm field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetEncryptionAlgorithm(v string) {
	o.EncryptionAlgorithm = &v
}

// GetEncryptionContext returns the EncryptionContext field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetEncryptionContext() CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext {
	if o == nil || IsNil(o.EncryptionContext) {
		var ret CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext
		return ret
	}
	return *o.EncryptionContext
}

// GetEncryptionContextOk returns a tuple with the EncryptionContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetEncryptionContextOk() (*CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext, bool) {
	if o == nil || IsNil(o.EncryptionContext) {
		return nil, false
	}
	return o.EncryptionContext, true
}

// HasEncryptionContext returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasEncryptionContext() bool {
	if o != nil && !IsNil(o.EncryptionContext) {
		return true
	}

	return false
}

// SetEncryptionContext gets a reference to the given CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext and assigns it to the EncryptionContext field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetEncryptionContext(v CloudInsightsEventsListEventDetailsRequestParametersEncryptionContext) {
	o.EncryptionContext = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetInclude(v []string) {
	o.Include = v
}

// GetIncludeDeletedResponses returns the IncludeDeletedResponses field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeDeletedResponses() bool {
	if o == nil || IsNil(o.IncludeDeletedResponses) {
		var ret bool
		return ret
	}
	return *o.IncludeDeletedResponses
}

// GetIncludeDeletedResponsesOk returns a tuple with the IncludeDeletedResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeDeletedResponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDeletedResponses) {
		return nil, false
	}
	return o.IncludeDeletedResponses, true
}

// HasIncludeDeletedResponses returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasIncludeDeletedResponses() bool {
	if o != nil && !IsNil(o.IncludeDeletedResponses) {
		return true
	}

	return false
}

// SetIncludeDeletedResponses gets a reference to the given bool and assigns it to the IncludeDeletedResponses field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetIncludeDeletedResponses(v bool) {
	o.IncludeDeletedResponses = &v
}

// GetIncludePublic returns the IncludePublic field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludePublic() bool {
	if o == nil || IsNil(o.IncludePublic) {
		var ret bool
		return ret
	}
	return *o.IncludePublic
}

// GetIncludePublicOk returns a tuple with the IncludePublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludePublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludePublic) {
		return nil, false
	}
	return o.IncludePublic, true
}

// HasIncludePublic returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasIncludePublic() bool {
	if o != nil && !IsNil(o.IncludePublic) {
		return true
	}

	return false
}

// SetIncludePublic gets a reference to the given bool and assigns it to the IncludePublic field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetIncludePublic(v bool) {
	o.IncludePublic = &v
}

// GetIncludeShared returns the IncludeShared field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeShared() bool {
	if o == nil || IsNil(o.IncludeShared) {
		var ret bool
		return ret
	}
	return *o.IncludeShared
}

// GetIncludeSharedOk returns a tuple with the IncludeShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludeSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeShared) {
		return nil, false
	}
	return o.IncludeShared, true
}

// HasIncludeShared returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasIncludeShared() bool {
	if o != nil && !IsNil(o.IncludeShared) {
		return true
	}

	return false
}

// SetIncludeShared gets a reference to the given bool and assigns it to the IncludeShared field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetIncludeShared(v bool) {
	o.IncludeShared = &v
}

// GetIncludes returns the Includes field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludes() CloudInsightsEventsListEventDetailsRequestParametersIncludes {
	if o == nil || IsNil(o.Includes) {
		var ret CloudInsightsEventsListEventDetailsRequestParametersIncludes
		return ret
	}
	return *o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetIncludesOk() (*CloudInsightsEventsListEventDetailsRequestParametersIncludes, bool) {
	if o == nil || IsNil(o.Includes) {
		return nil, false
	}
	return o.Includes, true
}

// HasIncludes returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasIncludes() bool {
	if o != nil && !IsNil(o.Includes) {
		return true
	}

	return false
}

// SetIncludes gets a reference to the given CloudInsightsEventsListEventDetailsRequestParametersIncludes and assigns it to the Includes field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetIncludes(v CloudInsightsEventsListEventDetailsRequestParametersIncludes) {
	o.Includes = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetKeyId(v string) {
	o.KeyId = &v
}

// GetKeySpec returns the KeySpec field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetKeySpec() string {
	if o == nil || IsNil(o.KeySpec) {
		var ret string
		return ret
	}
	return *o.KeySpec
}

// GetKeySpecOk returns a tuple with the KeySpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetKeySpecOk() (*string, bool) {
	if o == nil || IsNil(o.KeySpec) {
		return nil, false
	}
	return o.KeySpec, true
}

// HasKeySpec returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasKeySpec() bool {
	if o != nil && !IsNil(o.KeySpec) {
		return true
	}

	return false
}

// SetKeySpec gets a reference to the given string and assigns it to the KeySpec field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetKeySpec(v string) {
	o.KeySpec = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLimit() float32 {
	if o == nil || IsNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetLimit(v float32) {
	o.Limit = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetLocation(v string) {
	o.Location = &v
}

// GetLogGroupName returns the LogGroupName field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLogGroupName() string {
	if o == nil || IsNil(o.LogGroupName) {
		var ret string
		return ret
	}
	return *o.LogGroupName
}

// GetLogGroupNameOk returns a tuple with the LogGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLogGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogGroupName) {
		return nil, false
	}
	return o.LogGroupName, true
}

// HasLogGroupName returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasLogGroupName() bool {
	if o != nil && !IsNil(o.LogGroupName) {
		return true
	}

	return false
}

// SetLogGroupName gets a reference to the given string and assigns it to the LogGroupName field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetLogGroupName(v string) {
	o.LogGroupName = &v
}

// GetLogStreamName returns the LogStreamName field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLogStreamName() string {
	if o == nil || IsNil(o.LogStreamName) {
		var ret string
		return ret
	}
	return *o.LogStreamName
}

// GetLogStreamNameOk returns a tuple with the LogStreamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetLogStreamNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogStreamName) {
		return nil, false
	}
	return o.LogStreamName, true
}

// HasLogStreamName returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasLogStreamName() bool {
	if o != nil && !IsNil(o.LogStreamName) {
		return true
	}

	return false
}

// SetLogStreamName gets a reference to the given string and assigns it to the LogStreamName field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetLogStreamName(v string) {
	o.LogStreamName = &v
}

// GetPrincipalArn returns the PrincipalArn field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetPrincipalArn() string {
	if o == nil || IsNil(o.PrincipalArn) {
		var ret string
		return ret
	}
	return *o.PrincipalArn
}

// GetPrincipalArnOk returns a tuple with the PrincipalArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetPrincipalArnOk() (*string, bool) {
	if o == nil || IsNil(o.PrincipalArn) {
		return nil, false
	}
	return o.PrincipalArn, true
}

// HasPrincipalArn returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasPrincipalArn() bool {
	if o != nil && !IsNil(o.PrincipalArn) {
		return true
	}

	return false
}

// SetPrincipalArn gets a reference to the given string and assigns it to the PrincipalArn field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetPrincipalArn(v string) {
	o.PrincipalArn = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetResultToken returns the ResultToken field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetResultToken() string {
	if o == nil || IsNil(o.ResultToken) {
		var ret string
		return ret
	}
	return *o.ResultToken
}

// GetResultTokenOk returns a tuple with the ResultToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetResultTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ResultToken) {
		return nil, false
	}
	return o.ResultToken, true
}

// HasResultToken returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasResultToken() bool {
	if o != nil && !IsNil(o.ResultToken) {
		return true
	}

	return false
}

// SetResultToken gets a reference to the given string and assigns it to the ResultToken field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetResultToken(v string) {
	o.ResultToken = &v
}

// GetRoleArn returns the RoleArn field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetRoleArn() string {
	if o == nil || IsNil(o.RoleArn) {
		var ret string
		return ret
	}
	return *o.RoleArn
}

// GetRoleArnOk returns a tuple with the RoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetRoleArnOk() (*string, bool) {
	if o == nil || IsNil(o.RoleArn) {
		return nil, false
	}
	return o.RoleArn, true
}

// HasRoleArn returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasRoleArn() bool {
	if o != nil && !IsNil(o.RoleArn) {
		return true
	}

	return false
}

// SetRoleArn gets a reference to the given string and assigns it to the RoleArn field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetRoleArn(v string) {
	o.RoleArn = &v
}

// GetRoleSessionName returns the RoleSessionName field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetRoleSessionName() string {
	if o == nil || IsNil(o.RoleSessionName) {
		var ret string
		return ret
	}
	return *o.RoleSessionName
}

// GetRoleSessionNameOk returns a tuple with the RoleSessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetRoleSessionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleSessionName) {
		return nil, false
	}
	return o.RoleSessionName, true
}

// HasRoleSessionName returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasRoleSessionName() bool {
	if o != nil && !IsNil(o.RoleSessionName) {
		return true
	}

	return false
}

// SetRoleSessionName gets a reference to the given string and assigns it to the RoleSessionName field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetRoleSessionName(v string) {
	o.RoleSessionName = &v
}

// GetSAMLAssertionID returns the SAMLAssertionID field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetSAMLAssertionID() string {
	if o == nil || IsNil(o.SAMLAssertionID) {
		var ret string
		return ret
	}
	return *o.SAMLAssertionID
}

// GetSAMLAssertionIDOk returns a tuple with the SAMLAssertionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetSAMLAssertionIDOk() (*string, bool) {
	if o == nil || IsNil(o.SAMLAssertionID) {
		return nil, false
	}
	return o.SAMLAssertionID, true
}

// HasSAMLAssertionID returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasSAMLAssertionID() bool {
	if o != nil && !IsNil(o.SAMLAssertionID) {
		return true
	}

	return false
}

// SetSAMLAssertionID gets a reference to the given string and assigns it to the SAMLAssertionID field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetSAMLAssertionID(v string) {
	o.SAMLAssertionID = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetServices() []string {
	if o == nil || IsNil(o.Services) {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetServices(v []string) {
	o.Services = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetStatus(v string) {
	o.Status = &v
}

// GetTestMode returns the TestMode field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetTestMode() bool {
	if o == nil || IsNil(o.TestMode) {
		var ret bool
		return ret
	}
	return *o.TestMode
}

// GetTestModeOk returns a tuple with the TestMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetTestModeOk() (*bool, bool) {
	if o == nil || IsNil(o.TestMode) {
		return nil, false
	}
	return o.TestMode, true
}

// HasTestMode returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasTestMode() bool {
	if o != nil && !IsNil(o.TestMode) {
		return true
	}

	return false
}

// SetTestMode gets a reference to the given bool and assigns it to the TestMode field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetTestMode(v bool) {
	o.TestMode = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *CloudInsightsEventsListEventDetailsRequestParameters) SetUserName(v string) {
	o.UserName = &v
}

func (o CloudInsightsEventsListEventDetailsRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudInsightsEventsListEventDetailsRequestParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["Host"] = o.Host
	}
	if !IsNil(o.BucketName) {
		toSerialize["bucketName"] = o.BucketName
	}
	if !IsNil(o.CertificateArn) {
		toSerialize["certificateArn"] = o.CertificateArn
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.ComplianceTypes) {
		toSerialize["complianceTypes"] = o.ComplianceTypes
	}
	if !IsNil(o.ConfigRuleName) {
		toSerialize["configRuleName"] = o.ConfigRuleName
	}
	if !IsNil(o.DurationSeconds) {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	if !IsNil(o.EncryptionAlgorithm) {
		toSerialize["encryptionAlgorithm"] = o.EncryptionAlgorithm
	}
	if !IsNil(o.EncryptionContext) {
		toSerialize["encryptionContext"] = o.EncryptionContext
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	if !IsNil(o.IncludeDeletedResponses) {
		toSerialize["includeDeletedResponses"] = o.IncludeDeletedResponses
	}
	if !IsNil(o.IncludePublic) {
		toSerialize["includePublic"] = o.IncludePublic
	}
	if !IsNil(o.IncludeShared) {
		toSerialize["includeShared"] = o.IncludeShared
	}
	if !IsNil(o.Includes) {
		toSerialize["includes"] = o.Includes
	}
	if !IsNil(o.KeyId) {
		toSerialize["keyId"] = o.KeyId
	}
	if !IsNil(o.KeySpec) {
		toSerialize["keySpec"] = o.KeySpec
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.LogGroupName) {
		toSerialize["logGroupName"] = o.LogGroupName
	}
	if !IsNil(o.LogStreamName) {
		toSerialize["logStreamName"] = o.LogStreamName
	}
	if !IsNil(o.PrincipalArn) {
		toSerialize["principalArn"] = o.PrincipalArn
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.ResultToken) {
		toSerialize["resultToken"] = o.ResultToken
	}
	if !IsNil(o.RoleArn) {
		toSerialize["roleArn"] = o.RoleArn
	}
	if !IsNil(o.RoleSessionName) {
		toSerialize["roleSessionName"] = o.RoleSessionName
	}
	if !IsNil(o.SAMLAssertionID) {
		toSerialize["sAMLAssertionID"] = o.SAMLAssertionID
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TestMode) {
		toSerialize["testMode"] = o.TestMode
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	return toSerialize, nil
}

type NullableCloudInsightsEventsListEventDetailsRequestParameters struct {
	value *CloudInsightsEventsListEventDetailsRequestParameters
	isSet bool
}

func (v NullableCloudInsightsEventsListEventDetailsRequestParameters) Get() *CloudInsightsEventsListEventDetailsRequestParameters {
	return v.value
}

func (v *NullableCloudInsightsEventsListEventDetailsRequestParameters) Set(val *CloudInsightsEventsListEventDetailsRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudInsightsEventsListEventDetailsRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudInsightsEventsListEventDetailsRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudInsightsEventsListEventDetailsRequestParameters(val *CloudInsightsEventsListEventDetailsRequestParameters) *NullableCloudInsightsEventsListEventDetailsRequestParameters {
	return &NullableCloudInsightsEventsListEventDetailsRequestParameters{value: val, isSet: true}
}

func (v NullableCloudInsightsEventsListEventDetailsRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudInsightsEventsListEventDetailsRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



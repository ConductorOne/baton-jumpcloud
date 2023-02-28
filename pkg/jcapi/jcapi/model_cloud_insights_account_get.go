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

// checks if the CloudInsightsAccountGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudInsightsAccountGet{}

// CloudInsightsAccountGet struct for CloudInsightsAccountGet
type CloudInsightsAccountGet struct {
	AwsAccountAlias *string `json:"awsAccountAlias,omitempty"`
	AwsAccountId *string `json:"awsAccountId,omitempty"`
	AwsAccountIsActivated *string `json:"awsAccountIsActivated,omitempty"`
	AwsAccountIsValidated *string `json:"awsAccountIsValidated,omitempty"`
	AwsAccountRegion *string `json:"awsAccountRegion,omitempty"`
	AwsAccountRoleArn *string `json:"awsAccountRoleArn,omitempty"`
	AwsAccountRoleExternalId *string `json:"awsAccountRoleExternalId,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	LastEventReceivedAt *string `json:"lastEventReceivedAt,omitempty"`
	ObjectId *string `json:"objectId,omitempty"`
	OrganizationObjectId *string `json:"organizationObjectId,omitempty"`
	SqsQueueName *string `json:"sqsQueueName,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewCloudInsightsAccountGet instantiates a new CloudInsightsAccountGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudInsightsAccountGet() *CloudInsightsAccountGet {
	this := CloudInsightsAccountGet{}
	return &this
}

// NewCloudInsightsAccountGetWithDefaults instantiates a new CloudInsightsAccountGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudInsightsAccountGetWithDefaults() *CloudInsightsAccountGet {
	this := CloudInsightsAccountGet{}
	return &this
}

// GetAwsAccountAlias returns the AwsAccountAlias field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetAwsAccountAlias() string {
	if o == nil || IsNil(o.AwsAccountAlias) {
		var ret string
		return ret
	}
	return *o.AwsAccountAlias
}

// GetAwsAccountAliasOk returns a tuple with the AwsAccountAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetAwsAccountAliasOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountAlias) {
		return nil, false
	}
	return o.AwsAccountAlias, true
}

// HasAwsAccountAlias returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasAwsAccountAlias() bool {
	if o != nil && !IsNil(o.AwsAccountAlias) {
		return true
	}

	return false
}

// SetAwsAccountAlias gets a reference to the given string and assigns it to the AwsAccountAlias field.
func (o *CloudInsightsAccountGet) SetAwsAccountAlias(v string) {
	o.AwsAccountAlias = &v
}

// GetAwsAccountId returns the AwsAccountId field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetAwsAccountId() string {
	if o == nil || IsNil(o.AwsAccountId) {
		var ret string
		return ret
	}
	return *o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetAwsAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountId) {
		return nil, false
	}
	return o.AwsAccountId, true
}

// HasAwsAccountId returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasAwsAccountId() bool {
	if o != nil && !IsNil(o.AwsAccountId) {
		return true
	}

	return false
}

// SetAwsAccountId gets a reference to the given string and assigns it to the AwsAccountId field.
func (o *CloudInsightsAccountGet) SetAwsAccountId(v string) {
	o.AwsAccountId = &v
}

// GetAwsAccountIsActivated returns the AwsAccountIsActivated field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetAwsAccountIsActivated() string {
	if o == nil || IsNil(o.AwsAccountIsActivated) {
		var ret string
		return ret
	}
	return *o.AwsAccountIsActivated
}

// GetAwsAccountIsActivatedOk returns a tuple with the AwsAccountIsActivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetAwsAccountIsActivatedOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountIsActivated) {
		return nil, false
	}
	return o.AwsAccountIsActivated, true
}

// HasAwsAccountIsActivated returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasAwsAccountIsActivated() bool {
	if o != nil && !IsNil(o.AwsAccountIsActivated) {
		return true
	}

	return false
}

// SetAwsAccountIsActivated gets a reference to the given string and assigns it to the AwsAccountIsActivated field.
func (o *CloudInsightsAccountGet) SetAwsAccountIsActivated(v string) {
	o.AwsAccountIsActivated = &v
}

// GetAwsAccountIsValidated returns the AwsAccountIsValidated field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetAwsAccountIsValidated() string {
	if o == nil || IsNil(o.AwsAccountIsValidated) {
		var ret string
		return ret
	}
	return *o.AwsAccountIsValidated
}

// GetAwsAccountIsValidatedOk returns a tuple with the AwsAccountIsValidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetAwsAccountIsValidatedOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountIsValidated) {
		return nil, false
	}
	return o.AwsAccountIsValidated, true
}

// HasAwsAccountIsValidated returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasAwsAccountIsValidated() bool {
	if o != nil && !IsNil(o.AwsAccountIsValidated) {
		return true
	}

	return false
}

// SetAwsAccountIsValidated gets a reference to the given string and assigns it to the AwsAccountIsValidated field.
func (o *CloudInsightsAccountGet) SetAwsAccountIsValidated(v string) {
	o.AwsAccountIsValidated = &v
}

// GetAwsAccountRegion returns the AwsAccountRegion field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetAwsAccountRegion() string {
	if o == nil || IsNil(o.AwsAccountRegion) {
		var ret string
		return ret
	}
	return *o.AwsAccountRegion
}

// GetAwsAccountRegionOk returns a tuple with the AwsAccountRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetAwsAccountRegionOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountRegion) {
		return nil, false
	}
	return o.AwsAccountRegion, true
}

// HasAwsAccountRegion returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasAwsAccountRegion() bool {
	if o != nil && !IsNil(o.AwsAccountRegion) {
		return true
	}

	return false
}

// SetAwsAccountRegion gets a reference to the given string and assigns it to the AwsAccountRegion field.
func (o *CloudInsightsAccountGet) SetAwsAccountRegion(v string) {
	o.AwsAccountRegion = &v
}

// GetAwsAccountRoleArn returns the AwsAccountRoleArn field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetAwsAccountRoleArn() string {
	if o == nil || IsNil(o.AwsAccountRoleArn) {
		var ret string
		return ret
	}
	return *o.AwsAccountRoleArn
}

// GetAwsAccountRoleArnOk returns a tuple with the AwsAccountRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetAwsAccountRoleArnOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountRoleArn) {
		return nil, false
	}
	return o.AwsAccountRoleArn, true
}

// HasAwsAccountRoleArn returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasAwsAccountRoleArn() bool {
	if o != nil && !IsNil(o.AwsAccountRoleArn) {
		return true
	}

	return false
}

// SetAwsAccountRoleArn gets a reference to the given string and assigns it to the AwsAccountRoleArn field.
func (o *CloudInsightsAccountGet) SetAwsAccountRoleArn(v string) {
	o.AwsAccountRoleArn = &v
}

// GetAwsAccountRoleExternalId returns the AwsAccountRoleExternalId field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetAwsAccountRoleExternalId() string {
	if o == nil || IsNil(o.AwsAccountRoleExternalId) {
		var ret string
		return ret
	}
	return *o.AwsAccountRoleExternalId
}

// GetAwsAccountRoleExternalIdOk returns a tuple with the AwsAccountRoleExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetAwsAccountRoleExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountRoleExternalId) {
		return nil, false
	}
	return o.AwsAccountRoleExternalId, true
}

// HasAwsAccountRoleExternalId returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasAwsAccountRoleExternalId() bool {
	if o != nil && !IsNil(o.AwsAccountRoleExternalId) {
		return true
	}

	return false
}

// SetAwsAccountRoleExternalId gets a reference to the given string and assigns it to the AwsAccountRoleExternalId field.
func (o *CloudInsightsAccountGet) SetAwsAccountRoleExternalId(v string) {
	o.AwsAccountRoleExternalId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CloudInsightsAccountGet) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLastEventReceivedAt returns the LastEventReceivedAt field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetLastEventReceivedAt() string {
	if o == nil || IsNil(o.LastEventReceivedAt) {
		var ret string
		return ret
	}
	return *o.LastEventReceivedAt
}

// GetLastEventReceivedAtOk returns a tuple with the LastEventReceivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetLastEventReceivedAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastEventReceivedAt) {
		return nil, false
	}
	return o.LastEventReceivedAt, true
}

// HasLastEventReceivedAt returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasLastEventReceivedAt() bool {
	if o != nil && !IsNil(o.LastEventReceivedAt) {
		return true
	}

	return false
}

// SetLastEventReceivedAt gets a reference to the given string and assigns it to the LastEventReceivedAt field.
func (o *CloudInsightsAccountGet) SetLastEventReceivedAt(v string) {
	o.LastEventReceivedAt = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetObjectId() string {
	if o == nil || IsNil(o.ObjectId) {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasObjectId() bool {
	if o != nil && !IsNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *CloudInsightsAccountGet) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetOrganizationObjectId returns the OrganizationObjectId field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetOrganizationObjectId() string {
	if o == nil || IsNil(o.OrganizationObjectId) {
		var ret string
		return ret
	}
	return *o.OrganizationObjectId
}

// GetOrganizationObjectIdOk returns a tuple with the OrganizationObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetOrganizationObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationObjectId) {
		return nil, false
	}
	return o.OrganizationObjectId, true
}

// HasOrganizationObjectId returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasOrganizationObjectId() bool {
	if o != nil && !IsNil(o.OrganizationObjectId) {
		return true
	}

	return false
}

// SetOrganizationObjectId gets a reference to the given string and assigns it to the OrganizationObjectId field.
func (o *CloudInsightsAccountGet) SetOrganizationObjectId(v string) {
	o.OrganizationObjectId = &v
}

// GetSqsQueueName returns the SqsQueueName field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetSqsQueueName() string {
	if o == nil || IsNil(o.SqsQueueName) {
		var ret string
		return ret
	}
	return *o.SqsQueueName
}

// GetSqsQueueNameOk returns a tuple with the SqsQueueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetSqsQueueNameOk() (*string, bool) {
	if o == nil || IsNil(o.SqsQueueName) {
		return nil, false
	}
	return o.SqsQueueName, true
}

// HasSqsQueueName returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasSqsQueueName() bool {
	if o != nil && !IsNil(o.SqsQueueName) {
		return true
	}

	return false
}

// SetSqsQueueName gets a reference to the given string and assigns it to the SqsQueueName field.
func (o *CloudInsightsAccountGet) SetSqsQueueName(v string) {
	o.SqsQueueName = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CloudInsightsAccountGet) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsAccountGet) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CloudInsightsAccountGet) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CloudInsightsAccountGet) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o CloudInsightsAccountGet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudInsightsAccountGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsAccountAlias) {
		toSerialize["awsAccountAlias"] = o.AwsAccountAlias
	}
	if !IsNil(o.AwsAccountId) {
		toSerialize["awsAccountId"] = o.AwsAccountId
	}
	if !IsNil(o.AwsAccountIsActivated) {
		toSerialize["awsAccountIsActivated"] = o.AwsAccountIsActivated
	}
	if !IsNil(o.AwsAccountIsValidated) {
		toSerialize["awsAccountIsValidated"] = o.AwsAccountIsValidated
	}
	if !IsNil(o.AwsAccountRegion) {
		toSerialize["awsAccountRegion"] = o.AwsAccountRegion
	}
	if !IsNil(o.AwsAccountRoleArn) {
		toSerialize["awsAccountRoleArn"] = o.AwsAccountRoleArn
	}
	if !IsNil(o.AwsAccountRoleExternalId) {
		toSerialize["awsAccountRoleExternalId"] = o.AwsAccountRoleExternalId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.LastEventReceivedAt) {
		toSerialize["lastEventReceivedAt"] = o.LastEventReceivedAt
	}
	if !IsNil(o.ObjectId) {
		toSerialize["objectId"] = o.ObjectId
	}
	if !IsNil(o.OrganizationObjectId) {
		toSerialize["organizationObjectId"] = o.OrganizationObjectId
	}
	if !IsNil(o.SqsQueueName) {
		toSerialize["sqsQueueName"] = o.SqsQueueName
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableCloudInsightsAccountGet struct {
	value *CloudInsightsAccountGet
	isSet bool
}

func (v NullableCloudInsightsAccountGet) Get() *CloudInsightsAccountGet {
	return v.value
}

func (v *NullableCloudInsightsAccountGet) Set(val *CloudInsightsAccountGet) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudInsightsAccountGet) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudInsightsAccountGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudInsightsAccountGet(val *CloudInsightsAccountGet) *NullableCloudInsightsAccountGet {
	return &NullableCloudInsightsAccountGet{value: val, isSet: true}
}

func (v NullableCloudInsightsAccountGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudInsightsAccountGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



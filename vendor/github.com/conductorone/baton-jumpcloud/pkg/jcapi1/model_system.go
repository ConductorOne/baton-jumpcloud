/*
JumpCloud API

# Overview  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, and system users.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/systemusers\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 1.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jcapi1

import (
	"encoding/json"
	"time"
)

// checks if the System type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &System{}

// System struct for System
type System struct {
	Id *string `json:"_id,omitempty"`
	Active *bool `json:"active,omitempty"`
	AgentVersion *string `json:"agentVersion,omitempty"`
	AllowMultiFactorAuthentication *bool `json:"allowMultiFactorAuthentication,omitempty"`
	AllowPublicKeyAuthentication *bool `json:"allowPublicKeyAuthentication,omitempty"`
	AllowSshPasswordAuthentication *bool `json:"allowSshPasswordAuthentication,omitempty"`
	AllowSshRootLogin *bool `json:"allowSshRootLogin,omitempty"`
	AmazonInstanceID *string `json:"amazonInstanceID,omitempty"`
	Arch *string `json:"arch,omitempty"`
	AzureAdJoined *bool `json:"azureAdJoined,omitempty"`
	BuiltInCommands []SystemBuiltInCommandsInner `json:"builtInCommands,omitempty"`
	ConnectionHistory []map[string]interface{} `json:"connectionHistory,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Description *string `json:"description,omitempty"`
	DesktopCapable *bool `json:"desktopCapable,omitempty"`
	DisplayManager *string `json:"displayManager,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	DomainInfo *SystemDomainInfo `json:"domainInfo,omitempty"`
	Fde *Fde `json:"fde,omitempty"`
	FileSystem NullableString `json:"fileSystem,omitempty"`
	HasServiceAccount *bool `json:"hasServiceAccount,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	LastContact NullableTime `json:"lastContact,omitempty"`
	Mdm *SystemMdm `json:"mdm,omitempty"`
	ModifySSHDConfig *bool `json:"modifySSHDConfig,omitempty"`
	NetworkInterfaces []SystemNetworkInterfacesInner `json:"networkInterfaces,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Os *string `json:"os,omitempty"`
	OsFamily *string `json:"osFamily,omitempty"`
	OsVersionDetail *SystemOsVersionDetail `json:"osVersionDetail,omitempty"`
	ProvisionMetadata *SystemProvisionMetadata `json:"provisionMetadata,omitempty"`
	RemoteIP *string `json:"remoteIP,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	ServiceAccountState *SystemServiceAccountState `json:"serviceAccountState,omitempty"`
	SshRootEnabled *bool `json:"sshRootEnabled,omitempty"`
	SshdParams []SystemSshdParamsInner `json:"sshdParams,omitempty"`
	SystemInsights *SystemSystemInsights `json:"systemInsights,omitempty"`
	SystemTimezone *int32 `json:"systemTimezone,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
	UserMetrics []SystemUserMetricsInner `json:"userMetrics,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _System System

// NewSystem instantiates a new System object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystem() *System {
	this := System{}
	return &this
}

// NewSystemWithDefaults instantiates a new System object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemWithDefaults() *System {
	this := System{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *System) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *System) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *System) SetId(v string) {
	o.Id = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *System) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *System) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *System) SetActive(v bool) {
	o.Active = &v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *System) GetAgentVersion() string {
	if o == nil || IsNil(o.AgentVersion) {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetAgentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AgentVersion) {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *System) HasAgentVersion() bool {
	if o != nil && !IsNil(o.AgentVersion) {
		return true
	}

	return false
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *System) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetAllowMultiFactorAuthentication returns the AllowMultiFactorAuthentication field value if set, zero value otherwise.
func (o *System) GetAllowMultiFactorAuthentication() bool {
	if o == nil || IsNil(o.AllowMultiFactorAuthentication) {
		var ret bool
		return ret
	}
	return *o.AllowMultiFactorAuthentication
}

// GetAllowMultiFactorAuthenticationOk returns a tuple with the AllowMultiFactorAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetAllowMultiFactorAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMultiFactorAuthentication) {
		return nil, false
	}
	return o.AllowMultiFactorAuthentication, true
}

// HasAllowMultiFactorAuthentication returns a boolean if a field has been set.
func (o *System) HasAllowMultiFactorAuthentication() bool {
	if o != nil && !IsNil(o.AllowMultiFactorAuthentication) {
		return true
	}

	return false
}

// SetAllowMultiFactorAuthentication gets a reference to the given bool and assigns it to the AllowMultiFactorAuthentication field.
func (o *System) SetAllowMultiFactorAuthentication(v bool) {
	o.AllowMultiFactorAuthentication = &v
}

// GetAllowPublicKeyAuthentication returns the AllowPublicKeyAuthentication field value if set, zero value otherwise.
func (o *System) GetAllowPublicKeyAuthentication() bool {
	if o == nil || IsNil(o.AllowPublicKeyAuthentication) {
		var ret bool
		return ret
	}
	return *o.AllowPublicKeyAuthentication
}

// GetAllowPublicKeyAuthenticationOk returns a tuple with the AllowPublicKeyAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetAllowPublicKeyAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPublicKeyAuthentication) {
		return nil, false
	}
	return o.AllowPublicKeyAuthentication, true
}

// HasAllowPublicKeyAuthentication returns a boolean if a field has been set.
func (o *System) HasAllowPublicKeyAuthentication() bool {
	if o != nil && !IsNil(o.AllowPublicKeyAuthentication) {
		return true
	}

	return false
}

// SetAllowPublicKeyAuthentication gets a reference to the given bool and assigns it to the AllowPublicKeyAuthentication field.
func (o *System) SetAllowPublicKeyAuthentication(v bool) {
	o.AllowPublicKeyAuthentication = &v
}

// GetAllowSshPasswordAuthentication returns the AllowSshPasswordAuthentication field value if set, zero value otherwise.
func (o *System) GetAllowSshPasswordAuthentication() bool {
	if o == nil || IsNil(o.AllowSshPasswordAuthentication) {
		var ret bool
		return ret
	}
	return *o.AllowSshPasswordAuthentication
}

// GetAllowSshPasswordAuthenticationOk returns a tuple with the AllowSshPasswordAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetAllowSshPasswordAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSshPasswordAuthentication) {
		return nil, false
	}
	return o.AllowSshPasswordAuthentication, true
}

// HasAllowSshPasswordAuthentication returns a boolean if a field has been set.
func (o *System) HasAllowSshPasswordAuthentication() bool {
	if o != nil && !IsNil(o.AllowSshPasswordAuthentication) {
		return true
	}

	return false
}

// SetAllowSshPasswordAuthentication gets a reference to the given bool and assigns it to the AllowSshPasswordAuthentication field.
func (o *System) SetAllowSshPasswordAuthentication(v bool) {
	o.AllowSshPasswordAuthentication = &v
}

// GetAllowSshRootLogin returns the AllowSshRootLogin field value if set, zero value otherwise.
func (o *System) GetAllowSshRootLogin() bool {
	if o == nil || IsNil(o.AllowSshRootLogin) {
		var ret bool
		return ret
	}
	return *o.AllowSshRootLogin
}

// GetAllowSshRootLoginOk returns a tuple with the AllowSshRootLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetAllowSshRootLoginOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSshRootLogin) {
		return nil, false
	}
	return o.AllowSshRootLogin, true
}

// HasAllowSshRootLogin returns a boolean if a field has been set.
func (o *System) HasAllowSshRootLogin() bool {
	if o != nil && !IsNil(o.AllowSshRootLogin) {
		return true
	}

	return false
}

// SetAllowSshRootLogin gets a reference to the given bool and assigns it to the AllowSshRootLogin field.
func (o *System) SetAllowSshRootLogin(v bool) {
	o.AllowSshRootLogin = &v
}

// GetAmazonInstanceID returns the AmazonInstanceID field value if set, zero value otherwise.
func (o *System) GetAmazonInstanceID() string {
	if o == nil || IsNil(o.AmazonInstanceID) {
		var ret string
		return ret
	}
	return *o.AmazonInstanceID
}

// GetAmazonInstanceIDOk returns a tuple with the AmazonInstanceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetAmazonInstanceIDOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonInstanceID) {
		return nil, false
	}
	return o.AmazonInstanceID, true
}

// HasAmazonInstanceID returns a boolean if a field has been set.
func (o *System) HasAmazonInstanceID() bool {
	if o != nil && !IsNil(o.AmazonInstanceID) {
		return true
	}

	return false
}

// SetAmazonInstanceID gets a reference to the given string and assigns it to the AmazonInstanceID field.
func (o *System) SetAmazonInstanceID(v string) {
	o.AmazonInstanceID = &v
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *System) GetArch() string {
	if o == nil || IsNil(o.Arch) {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetArchOk() (*string, bool) {
	if o == nil || IsNil(o.Arch) {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *System) HasArch() bool {
	if o != nil && !IsNil(o.Arch) {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *System) SetArch(v string) {
	o.Arch = &v
}

// GetAzureAdJoined returns the AzureAdJoined field value if set, zero value otherwise.
func (o *System) GetAzureAdJoined() bool {
	if o == nil || IsNil(o.AzureAdJoined) {
		var ret bool
		return ret
	}
	return *o.AzureAdJoined
}

// GetAzureAdJoinedOk returns a tuple with the AzureAdJoined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetAzureAdJoinedOk() (*bool, bool) {
	if o == nil || IsNil(o.AzureAdJoined) {
		return nil, false
	}
	return o.AzureAdJoined, true
}

// HasAzureAdJoined returns a boolean if a field has been set.
func (o *System) HasAzureAdJoined() bool {
	if o != nil && !IsNil(o.AzureAdJoined) {
		return true
	}

	return false
}

// SetAzureAdJoined gets a reference to the given bool and assigns it to the AzureAdJoined field.
func (o *System) SetAzureAdJoined(v bool) {
	o.AzureAdJoined = &v
}

// GetBuiltInCommands returns the BuiltInCommands field value if set, zero value otherwise.
func (o *System) GetBuiltInCommands() []SystemBuiltInCommandsInner {
	if o == nil || IsNil(o.BuiltInCommands) {
		var ret []SystemBuiltInCommandsInner
		return ret
	}
	return o.BuiltInCommands
}

// GetBuiltInCommandsOk returns a tuple with the BuiltInCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetBuiltInCommandsOk() ([]SystemBuiltInCommandsInner, bool) {
	if o == nil || IsNil(o.BuiltInCommands) {
		return nil, false
	}
	return o.BuiltInCommands, true
}

// HasBuiltInCommands returns a boolean if a field has been set.
func (o *System) HasBuiltInCommands() bool {
	if o != nil && !IsNil(o.BuiltInCommands) {
		return true
	}

	return false
}

// SetBuiltInCommands gets a reference to the given []SystemBuiltInCommandsInner and assigns it to the BuiltInCommands field.
func (o *System) SetBuiltInCommands(v []SystemBuiltInCommandsInner) {
	o.BuiltInCommands = v
}

// GetConnectionHistory returns the ConnectionHistory field value if set, zero value otherwise.
func (o *System) GetConnectionHistory() []map[string]interface{} {
	if o == nil || IsNil(o.ConnectionHistory) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ConnectionHistory
}

// GetConnectionHistoryOk returns a tuple with the ConnectionHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetConnectionHistoryOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ConnectionHistory) {
		return nil, false
	}
	return o.ConnectionHistory, true
}

// HasConnectionHistory returns a boolean if a field has been set.
func (o *System) HasConnectionHistory() bool {
	if o != nil && !IsNil(o.ConnectionHistory) {
		return true
	}

	return false
}

// SetConnectionHistory gets a reference to the given []map[string]interface{} and assigns it to the ConnectionHistory field.
func (o *System) SetConnectionHistory(v []map[string]interface{}) {
	o.ConnectionHistory = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *System) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *System) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *System) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *System) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *System) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *System) SetDescription(v string) {
	o.Description = &v
}

// GetDesktopCapable returns the DesktopCapable field value if set, zero value otherwise.
func (o *System) GetDesktopCapable() bool {
	if o == nil || IsNil(o.DesktopCapable) {
		var ret bool
		return ret
	}
	return *o.DesktopCapable
}

// GetDesktopCapableOk returns a tuple with the DesktopCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetDesktopCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.DesktopCapable) {
		return nil, false
	}
	return o.DesktopCapable, true
}

// HasDesktopCapable returns a boolean if a field has been set.
func (o *System) HasDesktopCapable() bool {
	if o != nil && !IsNil(o.DesktopCapable) {
		return true
	}

	return false
}

// SetDesktopCapable gets a reference to the given bool and assigns it to the DesktopCapable field.
func (o *System) SetDesktopCapable(v bool) {
	o.DesktopCapable = &v
}

// GetDisplayManager returns the DisplayManager field value if set, zero value otherwise.
func (o *System) GetDisplayManager() string {
	if o == nil || IsNil(o.DisplayManager) {
		var ret string
		return ret
	}
	return *o.DisplayManager
}

// GetDisplayManagerOk returns a tuple with the DisplayManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetDisplayManagerOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayManager) {
		return nil, false
	}
	return o.DisplayManager, true
}

// HasDisplayManager returns a boolean if a field has been set.
func (o *System) HasDisplayManager() bool {
	if o != nil && !IsNil(o.DisplayManager) {
		return true
	}

	return false
}

// SetDisplayManager gets a reference to the given string and assigns it to the DisplayManager field.
func (o *System) SetDisplayManager(v string) {
	o.DisplayManager = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *System) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *System) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *System) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDomainInfo returns the DomainInfo field value if set, zero value otherwise.
func (o *System) GetDomainInfo() SystemDomainInfo {
	if o == nil || IsNil(o.DomainInfo) {
		var ret SystemDomainInfo
		return ret
	}
	return *o.DomainInfo
}

// GetDomainInfoOk returns a tuple with the DomainInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetDomainInfoOk() (*SystemDomainInfo, bool) {
	if o == nil || IsNil(o.DomainInfo) {
		return nil, false
	}
	return o.DomainInfo, true
}

// HasDomainInfo returns a boolean if a field has been set.
func (o *System) HasDomainInfo() bool {
	if o != nil && !IsNil(o.DomainInfo) {
		return true
	}

	return false
}

// SetDomainInfo gets a reference to the given SystemDomainInfo and assigns it to the DomainInfo field.
func (o *System) SetDomainInfo(v SystemDomainInfo) {
	o.DomainInfo = &v
}

// GetFde returns the Fde field value if set, zero value otherwise.
func (o *System) GetFde() Fde {
	if o == nil || IsNil(o.Fde) {
		var ret Fde
		return ret
	}
	return *o.Fde
}

// GetFdeOk returns a tuple with the Fde field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetFdeOk() (*Fde, bool) {
	if o == nil || IsNil(o.Fde) {
		return nil, false
	}
	return o.Fde, true
}

// HasFde returns a boolean if a field has been set.
func (o *System) HasFde() bool {
	if o != nil && !IsNil(o.Fde) {
		return true
	}

	return false
}

// SetFde gets a reference to the given Fde and assigns it to the Fde field.
func (o *System) SetFde(v Fde) {
	o.Fde = &v
}

// GetFileSystem returns the FileSystem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetFileSystem() string {
	if o == nil || IsNil(o.FileSystem.Get()) {
		var ret string
		return ret
	}
	return *o.FileSystem.Get()
}

// GetFileSystemOk returns a tuple with the FileSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetFileSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSystem.Get(), o.FileSystem.IsSet()
}

// HasFileSystem returns a boolean if a field has been set.
func (o *System) HasFileSystem() bool {
	if o != nil && o.FileSystem.IsSet() {
		return true
	}

	return false
}

// SetFileSystem gets a reference to the given NullableString and assigns it to the FileSystem field.
func (o *System) SetFileSystem(v string) {
	o.FileSystem.Set(&v)
}
// SetFileSystemNil sets the value for FileSystem to be an explicit nil
func (o *System) SetFileSystemNil() {
	o.FileSystem.Set(nil)
}

// UnsetFileSystem ensures that no value is present for FileSystem, not even an explicit nil
func (o *System) UnsetFileSystem() {
	o.FileSystem.Unset()
}

// GetHasServiceAccount returns the HasServiceAccount field value if set, zero value otherwise.
func (o *System) GetHasServiceAccount() bool {
	if o == nil || IsNil(o.HasServiceAccount) {
		var ret bool
		return ret
	}
	return *o.HasServiceAccount
}

// GetHasServiceAccountOk returns a tuple with the HasServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetHasServiceAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.HasServiceAccount) {
		return nil, false
	}
	return o.HasServiceAccount, true
}

// HasHasServiceAccount returns a boolean if a field has been set.
func (o *System) HasHasServiceAccount() bool {
	if o != nil && !IsNil(o.HasServiceAccount) {
		return true
	}

	return false
}

// SetHasServiceAccount gets a reference to the given bool and assigns it to the HasServiceAccount field.
func (o *System) SetHasServiceAccount(v bool) {
	o.HasServiceAccount = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *System) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *System) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *System) SetHostname(v string) {
	o.Hostname = &v
}

// GetLastContact returns the LastContact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *System) GetLastContact() time.Time {
	if o == nil || IsNil(o.LastContact.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastContact.Get()
}

// GetLastContactOk returns a tuple with the LastContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *System) GetLastContactOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastContact.Get(), o.LastContact.IsSet()
}

// HasLastContact returns a boolean if a field has been set.
func (o *System) HasLastContact() bool {
	if o != nil && o.LastContact.IsSet() {
		return true
	}

	return false
}

// SetLastContact gets a reference to the given NullableTime and assigns it to the LastContact field.
func (o *System) SetLastContact(v time.Time) {
	o.LastContact.Set(&v)
}
// SetLastContactNil sets the value for LastContact to be an explicit nil
func (o *System) SetLastContactNil() {
	o.LastContact.Set(nil)
}

// UnsetLastContact ensures that no value is present for LastContact, not even an explicit nil
func (o *System) UnsetLastContact() {
	o.LastContact.Unset()
}

// GetMdm returns the Mdm field value if set, zero value otherwise.
func (o *System) GetMdm() SystemMdm {
	if o == nil || IsNil(o.Mdm) {
		var ret SystemMdm
		return ret
	}
	return *o.Mdm
}

// GetMdmOk returns a tuple with the Mdm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetMdmOk() (*SystemMdm, bool) {
	if o == nil || IsNil(o.Mdm) {
		return nil, false
	}
	return o.Mdm, true
}

// HasMdm returns a boolean if a field has been set.
func (o *System) HasMdm() bool {
	if o != nil && !IsNil(o.Mdm) {
		return true
	}

	return false
}

// SetMdm gets a reference to the given SystemMdm and assigns it to the Mdm field.
func (o *System) SetMdm(v SystemMdm) {
	o.Mdm = &v
}

// GetModifySSHDConfig returns the ModifySSHDConfig field value if set, zero value otherwise.
func (o *System) GetModifySSHDConfig() bool {
	if o == nil || IsNil(o.ModifySSHDConfig) {
		var ret bool
		return ret
	}
	return *o.ModifySSHDConfig
}

// GetModifySSHDConfigOk returns a tuple with the ModifySSHDConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetModifySSHDConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.ModifySSHDConfig) {
		return nil, false
	}
	return o.ModifySSHDConfig, true
}

// HasModifySSHDConfig returns a boolean if a field has been set.
func (o *System) HasModifySSHDConfig() bool {
	if o != nil && !IsNil(o.ModifySSHDConfig) {
		return true
	}

	return false
}

// SetModifySSHDConfig gets a reference to the given bool and assigns it to the ModifySSHDConfig field.
func (o *System) SetModifySSHDConfig(v bool) {
	o.ModifySSHDConfig = &v
}

// GetNetworkInterfaces returns the NetworkInterfaces field value if set, zero value otherwise.
func (o *System) GetNetworkInterfaces() []SystemNetworkInterfacesInner {
	if o == nil || IsNil(o.NetworkInterfaces) {
		var ret []SystemNetworkInterfacesInner
		return ret
	}
	return o.NetworkInterfaces
}

// GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetNetworkInterfacesOk() ([]SystemNetworkInterfacesInner, bool) {
	if o == nil || IsNil(o.NetworkInterfaces) {
		return nil, false
	}
	return o.NetworkInterfaces, true
}

// HasNetworkInterfaces returns a boolean if a field has been set.
func (o *System) HasNetworkInterfaces() bool {
	if o != nil && !IsNil(o.NetworkInterfaces) {
		return true
	}

	return false
}

// SetNetworkInterfaces gets a reference to the given []SystemNetworkInterfacesInner and assigns it to the NetworkInterfaces field.
func (o *System) SetNetworkInterfaces(v []SystemNetworkInterfacesInner) {
	o.NetworkInterfaces = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *System) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *System) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *System) SetOrganization(v string) {
	o.Organization = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *System) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *System) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *System) SetOs(v string) {
	o.Os = &v
}

// GetOsFamily returns the OsFamily field value if set, zero value otherwise.
func (o *System) GetOsFamily() string {
	if o == nil || IsNil(o.OsFamily) {
		var ret string
		return ret
	}
	return *o.OsFamily
}

// GetOsFamilyOk returns a tuple with the OsFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetOsFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.OsFamily) {
		return nil, false
	}
	return o.OsFamily, true
}

// HasOsFamily returns a boolean if a field has been set.
func (o *System) HasOsFamily() bool {
	if o != nil && !IsNil(o.OsFamily) {
		return true
	}

	return false
}

// SetOsFamily gets a reference to the given string and assigns it to the OsFamily field.
func (o *System) SetOsFamily(v string) {
	o.OsFamily = &v
}

// GetOsVersionDetail returns the OsVersionDetail field value if set, zero value otherwise.
func (o *System) GetOsVersionDetail() SystemOsVersionDetail {
	if o == nil || IsNil(o.OsVersionDetail) {
		var ret SystemOsVersionDetail
		return ret
	}
	return *o.OsVersionDetail
}

// GetOsVersionDetailOk returns a tuple with the OsVersionDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetOsVersionDetailOk() (*SystemOsVersionDetail, bool) {
	if o == nil || IsNil(o.OsVersionDetail) {
		return nil, false
	}
	return o.OsVersionDetail, true
}

// HasOsVersionDetail returns a boolean if a field has been set.
func (o *System) HasOsVersionDetail() bool {
	if o != nil && !IsNil(o.OsVersionDetail) {
		return true
	}

	return false
}

// SetOsVersionDetail gets a reference to the given SystemOsVersionDetail and assigns it to the OsVersionDetail field.
func (o *System) SetOsVersionDetail(v SystemOsVersionDetail) {
	o.OsVersionDetail = &v
}

// GetProvisionMetadata returns the ProvisionMetadata field value if set, zero value otherwise.
func (o *System) GetProvisionMetadata() SystemProvisionMetadata {
	if o == nil || IsNil(o.ProvisionMetadata) {
		var ret SystemProvisionMetadata
		return ret
	}
	return *o.ProvisionMetadata
}

// GetProvisionMetadataOk returns a tuple with the ProvisionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetProvisionMetadataOk() (*SystemProvisionMetadata, bool) {
	if o == nil || IsNil(o.ProvisionMetadata) {
		return nil, false
	}
	return o.ProvisionMetadata, true
}

// HasProvisionMetadata returns a boolean if a field has been set.
func (o *System) HasProvisionMetadata() bool {
	if o != nil && !IsNil(o.ProvisionMetadata) {
		return true
	}

	return false
}

// SetProvisionMetadata gets a reference to the given SystemProvisionMetadata and assigns it to the ProvisionMetadata field.
func (o *System) SetProvisionMetadata(v SystemProvisionMetadata) {
	o.ProvisionMetadata = &v
}

// GetRemoteIP returns the RemoteIP field value if set, zero value otherwise.
func (o *System) GetRemoteIP() string {
	if o == nil || IsNil(o.RemoteIP) {
		var ret string
		return ret
	}
	return *o.RemoteIP
}

// GetRemoteIPOk returns a tuple with the RemoteIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetRemoteIPOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteIP) {
		return nil, false
	}
	return o.RemoteIP, true
}

// HasRemoteIP returns a boolean if a field has been set.
func (o *System) HasRemoteIP() bool {
	if o != nil && !IsNil(o.RemoteIP) {
		return true
	}

	return false
}

// SetRemoteIP gets a reference to the given string and assigns it to the RemoteIP field.
func (o *System) SetRemoteIP(v string) {
	o.RemoteIP = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *System) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *System) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *System) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetServiceAccountState returns the ServiceAccountState field value if set, zero value otherwise.
func (o *System) GetServiceAccountState() SystemServiceAccountState {
	if o == nil || IsNil(o.ServiceAccountState) {
		var ret SystemServiceAccountState
		return ret
	}
	return *o.ServiceAccountState
}

// GetServiceAccountStateOk returns a tuple with the ServiceAccountState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetServiceAccountStateOk() (*SystemServiceAccountState, bool) {
	if o == nil || IsNil(o.ServiceAccountState) {
		return nil, false
	}
	return o.ServiceAccountState, true
}

// HasServiceAccountState returns a boolean if a field has been set.
func (o *System) HasServiceAccountState() bool {
	if o != nil && !IsNil(o.ServiceAccountState) {
		return true
	}

	return false
}

// SetServiceAccountState gets a reference to the given SystemServiceAccountState and assigns it to the ServiceAccountState field.
func (o *System) SetServiceAccountState(v SystemServiceAccountState) {
	o.ServiceAccountState = &v
}

// GetSshRootEnabled returns the SshRootEnabled field value if set, zero value otherwise.
func (o *System) GetSshRootEnabled() bool {
	if o == nil || IsNil(o.SshRootEnabled) {
		var ret bool
		return ret
	}
	return *o.SshRootEnabled
}

// GetSshRootEnabledOk returns a tuple with the SshRootEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSshRootEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SshRootEnabled) {
		return nil, false
	}
	return o.SshRootEnabled, true
}

// HasSshRootEnabled returns a boolean if a field has been set.
func (o *System) HasSshRootEnabled() bool {
	if o != nil && !IsNil(o.SshRootEnabled) {
		return true
	}

	return false
}

// SetSshRootEnabled gets a reference to the given bool and assigns it to the SshRootEnabled field.
func (o *System) SetSshRootEnabled(v bool) {
	o.SshRootEnabled = &v
}

// GetSshdParams returns the SshdParams field value if set, zero value otherwise.
func (o *System) GetSshdParams() []SystemSshdParamsInner {
	if o == nil || IsNil(o.SshdParams) {
		var ret []SystemSshdParamsInner
		return ret
	}
	return o.SshdParams
}

// GetSshdParamsOk returns a tuple with the SshdParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSshdParamsOk() ([]SystemSshdParamsInner, bool) {
	if o == nil || IsNil(o.SshdParams) {
		return nil, false
	}
	return o.SshdParams, true
}

// HasSshdParams returns a boolean if a field has been set.
func (o *System) HasSshdParams() bool {
	if o != nil && !IsNil(o.SshdParams) {
		return true
	}

	return false
}

// SetSshdParams gets a reference to the given []SystemSshdParamsInner and assigns it to the SshdParams field.
func (o *System) SetSshdParams(v []SystemSshdParamsInner) {
	o.SshdParams = v
}

// GetSystemInsights returns the SystemInsights field value if set, zero value otherwise.
func (o *System) GetSystemInsights() SystemSystemInsights {
	if o == nil || IsNil(o.SystemInsights) {
		var ret SystemSystemInsights
		return ret
	}
	return *o.SystemInsights
}

// GetSystemInsightsOk returns a tuple with the SystemInsights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSystemInsightsOk() (*SystemSystemInsights, bool) {
	if o == nil || IsNil(o.SystemInsights) {
		return nil, false
	}
	return o.SystemInsights, true
}

// HasSystemInsights returns a boolean if a field has been set.
func (o *System) HasSystemInsights() bool {
	if o != nil && !IsNil(o.SystemInsights) {
		return true
	}

	return false
}

// SetSystemInsights gets a reference to the given SystemSystemInsights and assigns it to the SystemInsights field.
func (o *System) SetSystemInsights(v SystemSystemInsights) {
	o.SystemInsights = &v
}

// GetSystemTimezone returns the SystemTimezone field value if set, zero value otherwise.
func (o *System) GetSystemTimezone() int32 {
	if o == nil || IsNil(o.SystemTimezone) {
		var ret int32
		return ret
	}
	return *o.SystemTimezone
}

// GetSystemTimezoneOk returns a tuple with the SystemTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetSystemTimezoneOk() (*int32, bool) {
	if o == nil || IsNil(o.SystemTimezone) {
		return nil, false
	}
	return o.SystemTimezone, true
}

// HasSystemTimezone returns a boolean if a field has been set.
func (o *System) HasSystemTimezone() bool {
	if o != nil && !IsNil(o.SystemTimezone) {
		return true
	}

	return false
}

// SetSystemTimezone gets a reference to the given int32 and assigns it to the SystemTimezone field.
func (o *System) SetSystemTimezone(v int32) {
	o.SystemTimezone = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *System) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *System) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *System) SetTags(v []string) {
	o.Tags = v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *System) GetTemplateName() string {
	if o == nil || IsNil(o.TemplateName) {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateName) {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *System) HasTemplateName() bool {
	if o != nil && !IsNil(o.TemplateName) {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *System) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetUserMetrics returns the UserMetrics field value if set, zero value otherwise.
func (o *System) GetUserMetrics() []SystemUserMetricsInner {
	if o == nil || IsNil(o.UserMetrics) {
		var ret []SystemUserMetricsInner
		return ret
	}
	return o.UserMetrics
}

// GetUserMetricsOk returns a tuple with the UserMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetUserMetricsOk() ([]SystemUserMetricsInner, bool) {
	if o == nil || IsNil(o.UserMetrics) {
		return nil, false
	}
	return o.UserMetrics, true
}

// HasUserMetrics returns a boolean if a field has been set.
func (o *System) HasUserMetrics() bool {
	if o != nil && !IsNil(o.UserMetrics) {
		return true
	}

	return false
}

// SetUserMetrics gets a reference to the given []SystemUserMetricsInner and assigns it to the UserMetrics field.
func (o *System) SetUserMetrics(v []SystemUserMetricsInner) {
	o.UserMetrics = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *System) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *System) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *System) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *System) SetVersion(v string) {
	o.Version = &v
}

func (o System) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o System) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.AgentVersion) {
		toSerialize["agentVersion"] = o.AgentVersion
	}
	if !IsNil(o.AllowMultiFactorAuthentication) {
		toSerialize["allowMultiFactorAuthentication"] = o.AllowMultiFactorAuthentication
	}
	if !IsNil(o.AllowPublicKeyAuthentication) {
		toSerialize["allowPublicKeyAuthentication"] = o.AllowPublicKeyAuthentication
	}
	if !IsNil(o.AllowSshPasswordAuthentication) {
		toSerialize["allowSshPasswordAuthentication"] = o.AllowSshPasswordAuthentication
	}
	if !IsNil(o.AllowSshRootLogin) {
		toSerialize["allowSshRootLogin"] = o.AllowSshRootLogin
	}
	if !IsNil(o.AmazonInstanceID) {
		toSerialize["amazonInstanceID"] = o.AmazonInstanceID
	}
	if !IsNil(o.Arch) {
		toSerialize["arch"] = o.Arch
	}
	if !IsNil(o.AzureAdJoined) {
		toSerialize["azureAdJoined"] = o.AzureAdJoined
	}
	// skip: builtInCommands is readOnly
	if !IsNil(o.ConnectionHistory) {
		toSerialize["connectionHistory"] = o.ConnectionHistory
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DesktopCapable) {
		toSerialize["desktopCapable"] = o.DesktopCapable
	}
	if !IsNil(o.DisplayManager) {
		toSerialize["displayManager"] = o.DisplayManager
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DomainInfo) {
		toSerialize["domainInfo"] = o.DomainInfo
	}
	if !IsNil(o.Fde) {
		toSerialize["fde"] = o.Fde
	}
	if o.FileSystem.IsSet() {
		toSerialize["fileSystem"] = o.FileSystem.Get()
	}
	if !IsNil(o.HasServiceAccount) {
		toSerialize["hasServiceAccount"] = o.HasServiceAccount
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if o.LastContact.IsSet() {
		toSerialize["lastContact"] = o.LastContact.Get()
	}
	if !IsNil(o.Mdm) {
		toSerialize["mdm"] = o.Mdm
	}
	if !IsNil(o.ModifySSHDConfig) {
		toSerialize["modifySSHDConfig"] = o.ModifySSHDConfig
	}
	if !IsNil(o.NetworkInterfaces) {
		toSerialize["networkInterfaces"] = o.NetworkInterfaces
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.OsFamily) {
		toSerialize["osFamily"] = o.OsFamily
	}
	if !IsNil(o.OsVersionDetail) {
		toSerialize["osVersionDetail"] = o.OsVersionDetail
	}
	if !IsNil(o.ProvisionMetadata) {
		toSerialize["provisionMetadata"] = o.ProvisionMetadata
	}
	if !IsNil(o.RemoteIP) {
		toSerialize["remoteIP"] = o.RemoteIP
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.ServiceAccountState) {
		toSerialize["serviceAccountState"] = o.ServiceAccountState
	}
	if !IsNil(o.SshRootEnabled) {
		toSerialize["sshRootEnabled"] = o.SshRootEnabled
	}
	if !IsNil(o.SshdParams) {
		toSerialize["sshdParams"] = o.SshdParams
	}
	if !IsNil(o.SystemInsights) {
		toSerialize["systemInsights"] = o.SystemInsights
	}
	if !IsNil(o.SystemTimezone) {
		toSerialize["systemTimezone"] = o.SystemTimezone
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TemplateName) {
		toSerialize["templateName"] = o.TemplateName
	}
	if !IsNil(o.UserMetrics) {
		toSerialize["userMetrics"] = o.UserMetrics
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *System) UnmarshalJSON(bytes []byte) (err error) {
	varSystem := _System{}

	if err = json.Unmarshal(bytes, &varSystem); err == nil {
		*o = System(varSystem)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "active")
		delete(additionalProperties, "agentVersion")
		delete(additionalProperties, "allowMultiFactorAuthentication")
		delete(additionalProperties, "allowPublicKeyAuthentication")
		delete(additionalProperties, "allowSshPasswordAuthentication")
		delete(additionalProperties, "allowSshRootLogin")
		delete(additionalProperties, "amazonInstanceID")
		delete(additionalProperties, "arch")
		delete(additionalProperties, "azureAdJoined")
		delete(additionalProperties, "builtInCommands")
		delete(additionalProperties, "connectionHistory")
		delete(additionalProperties, "created")
		delete(additionalProperties, "description")
		delete(additionalProperties, "desktopCapable")
		delete(additionalProperties, "displayManager")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "domainInfo")
		delete(additionalProperties, "fde")
		delete(additionalProperties, "fileSystem")
		delete(additionalProperties, "hasServiceAccount")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "lastContact")
		delete(additionalProperties, "mdm")
		delete(additionalProperties, "modifySSHDConfig")
		delete(additionalProperties, "networkInterfaces")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "os")
		delete(additionalProperties, "osFamily")
		delete(additionalProperties, "osVersionDetail")
		delete(additionalProperties, "provisionMetadata")
		delete(additionalProperties, "remoteIP")
		delete(additionalProperties, "serialNumber")
		delete(additionalProperties, "serviceAccountState")
		delete(additionalProperties, "sshRootEnabled")
		delete(additionalProperties, "sshdParams")
		delete(additionalProperties, "systemInsights")
		delete(additionalProperties, "systemTimezone")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "templateName")
		delete(additionalProperties, "userMetrics")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystem struct {
	value *System
	isSet bool
}

func (v NullableSystem) Get() *System {
	return v.value
}

func (v *NullableSystem) Set(val *System) {
	v.value = val
	v.isSet = true
}

func (v NullableSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystem(val *System) *NullableSystem {
	return &NullableSystem{value: val, isSet: true}
}

func (v NullableSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



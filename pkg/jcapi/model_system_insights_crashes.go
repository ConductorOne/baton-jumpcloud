/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jcapi

import (
	"encoding/json"
)

// checks if the SystemInsightsCrashes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInsightsCrashes{}

// SystemInsightsCrashes struct for SystemInsightsCrashes
type SystemInsightsCrashes struct {
	CollectionTime *string `json:"collection_time,omitempty"`
	CrashPath *string `json:"crash_path,omitempty"`
	CrashedThread *string `json:"crashed_thread,omitempty"`
	Datetime *string `json:"datetime,omitempty"`
	ExceptionCodes *string `json:"exception_codes,omitempty"`
	ExceptionNotes *string `json:"exception_notes,omitempty"`
	ExceptionType *string `json:"exception_type,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	Parent *string `json:"parent,omitempty"`
	Path *string `json:"path,omitempty"`
	Pid *string `json:"pid,omitempty"`
	Registers *string `json:"registers,omitempty"`
	Responsible *string `json:"responsible,omitempty"`
	StackTrace *string `json:"stack_trace,omitempty"`
	SystemId *string `json:"system_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Uid *int32 `json:"uid,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemInsightsCrashes SystemInsightsCrashes

// NewSystemInsightsCrashes instantiates a new SystemInsightsCrashes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInsightsCrashes() *SystemInsightsCrashes {
	this := SystemInsightsCrashes{}
	return &this
}

// NewSystemInsightsCrashesWithDefaults instantiates a new SystemInsightsCrashes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInsightsCrashesWithDefaults() *SystemInsightsCrashes {
	this := SystemInsightsCrashes{}
	return &this
}

// GetCollectionTime returns the CollectionTime field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetCollectionTime() string {
	if o == nil || IsNil(o.CollectionTime) {
		var ret string
		return ret
	}
	return *o.CollectionTime
}

// GetCollectionTimeOk returns a tuple with the CollectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetCollectionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionTime) {
		return nil, false
	}
	return o.CollectionTime, true
}

// HasCollectionTime returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasCollectionTime() bool {
	if o != nil && !IsNil(o.CollectionTime) {
		return true
	}

	return false
}

// SetCollectionTime gets a reference to the given string and assigns it to the CollectionTime field.
func (o *SystemInsightsCrashes) SetCollectionTime(v string) {
	o.CollectionTime = &v
}

// GetCrashPath returns the CrashPath field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetCrashPath() string {
	if o == nil || IsNil(o.CrashPath) {
		var ret string
		return ret
	}
	return *o.CrashPath
}

// GetCrashPathOk returns a tuple with the CrashPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetCrashPathOk() (*string, bool) {
	if o == nil || IsNil(o.CrashPath) {
		return nil, false
	}
	return o.CrashPath, true
}

// HasCrashPath returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasCrashPath() bool {
	if o != nil && !IsNil(o.CrashPath) {
		return true
	}

	return false
}

// SetCrashPath gets a reference to the given string and assigns it to the CrashPath field.
func (o *SystemInsightsCrashes) SetCrashPath(v string) {
	o.CrashPath = &v
}

// GetCrashedThread returns the CrashedThread field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetCrashedThread() string {
	if o == nil || IsNil(o.CrashedThread) {
		var ret string
		return ret
	}
	return *o.CrashedThread
}

// GetCrashedThreadOk returns a tuple with the CrashedThread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetCrashedThreadOk() (*string, bool) {
	if o == nil || IsNil(o.CrashedThread) {
		return nil, false
	}
	return o.CrashedThread, true
}

// HasCrashedThread returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasCrashedThread() bool {
	if o != nil && !IsNil(o.CrashedThread) {
		return true
	}

	return false
}

// SetCrashedThread gets a reference to the given string and assigns it to the CrashedThread field.
func (o *SystemInsightsCrashes) SetCrashedThread(v string) {
	o.CrashedThread = &v
}

// GetDatetime returns the Datetime field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetDatetime() string {
	if o == nil || IsNil(o.Datetime) {
		var ret string
		return ret
	}
	return *o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetDatetimeOk() (*string, bool) {
	if o == nil || IsNil(o.Datetime) {
		return nil, false
	}
	return o.Datetime, true
}

// HasDatetime returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasDatetime() bool {
	if o != nil && !IsNil(o.Datetime) {
		return true
	}

	return false
}

// SetDatetime gets a reference to the given string and assigns it to the Datetime field.
func (o *SystemInsightsCrashes) SetDatetime(v string) {
	o.Datetime = &v
}

// GetExceptionCodes returns the ExceptionCodes field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetExceptionCodes() string {
	if o == nil || IsNil(o.ExceptionCodes) {
		var ret string
		return ret
	}
	return *o.ExceptionCodes
}

// GetExceptionCodesOk returns a tuple with the ExceptionCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetExceptionCodesOk() (*string, bool) {
	if o == nil || IsNil(o.ExceptionCodes) {
		return nil, false
	}
	return o.ExceptionCodes, true
}

// HasExceptionCodes returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasExceptionCodes() bool {
	if o != nil && !IsNil(o.ExceptionCodes) {
		return true
	}

	return false
}

// SetExceptionCodes gets a reference to the given string and assigns it to the ExceptionCodes field.
func (o *SystemInsightsCrashes) SetExceptionCodes(v string) {
	o.ExceptionCodes = &v
}

// GetExceptionNotes returns the ExceptionNotes field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetExceptionNotes() string {
	if o == nil || IsNil(o.ExceptionNotes) {
		var ret string
		return ret
	}
	return *o.ExceptionNotes
}

// GetExceptionNotesOk returns a tuple with the ExceptionNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetExceptionNotesOk() (*string, bool) {
	if o == nil || IsNil(o.ExceptionNotes) {
		return nil, false
	}
	return o.ExceptionNotes, true
}

// HasExceptionNotes returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasExceptionNotes() bool {
	if o != nil && !IsNil(o.ExceptionNotes) {
		return true
	}

	return false
}

// SetExceptionNotes gets a reference to the given string and assigns it to the ExceptionNotes field.
func (o *SystemInsightsCrashes) SetExceptionNotes(v string) {
	o.ExceptionNotes = &v
}

// GetExceptionType returns the ExceptionType field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetExceptionType() string {
	if o == nil || IsNil(o.ExceptionType) {
		var ret string
		return ret
	}
	return *o.ExceptionType
}

// GetExceptionTypeOk returns a tuple with the ExceptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetExceptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExceptionType) {
		return nil, false
	}
	return o.ExceptionType, true
}

// HasExceptionType returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasExceptionType() bool {
	if o != nil && !IsNil(o.ExceptionType) {
		return true
	}

	return false
}

// SetExceptionType gets a reference to the given string and assigns it to the ExceptionType field.
func (o *SystemInsightsCrashes) SetExceptionType(v string) {
	o.ExceptionType = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *SystemInsightsCrashes) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *SystemInsightsCrashes) SetParent(v string) {
	o.Parent = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SystemInsightsCrashes) SetPath(v string) {
	o.Path = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *SystemInsightsCrashes) SetPid(v string) {
	o.Pid = &v
}

// GetRegisters returns the Registers field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetRegisters() string {
	if o == nil || IsNil(o.Registers) {
		var ret string
		return ret
	}
	return *o.Registers
}

// GetRegistersOk returns a tuple with the Registers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetRegistersOk() (*string, bool) {
	if o == nil || IsNil(o.Registers) {
		return nil, false
	}
	return o.Registers, true
}

// HasRegisters returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasRegisters() bool {
	if o != nil && !IsNil(o.Registers) {
		return true
	}

	return false
}

// SetRegisters gets a reference to the given string and assigns it to the Registers field.
func (o *SystemInsightsCrashes) SetRegisters(v string) {
	o.Registers = &v
}

// GetResponsible returns the Responsible field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetResponsible() string {
	if o == nil || IsNil(o.Responsible) {
		var ret string
		return ret
	}
	return *o.Responsible
}

// GetResponsibleOk returns a tuple with the Responsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetResponsibleOk() (*string, bool) {
	if o == nil || IsNil(o.Responsible) {
		return nil, false
	}
	return o.Responsible, true
}

// HasResponsible returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasResponsible() bool {
	if o != nil && !IsNil(o.Responsible) {
		return true
	}

	return false
}

// SetResponsible gets a reference to the given string and assigns it to the Responsible field.
func (o *SystemInsightsCrashes) SetResponsible(v string) {
	o.Responsible = &v
}

// GetStackTrace returns the StackTrace field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetStackTrace() string {
	if o == nil || IsNil(o.StackTrace) {
		var ret string
		return ret
	}
	return *o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetStackTraceOk() (*string, bool) {
	if o == nil || IsNil(o.StackTrace) {
		return nil, false
	}
	return o.StackTrace, true
}

// HasStackTrace returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasStackTrace() bool {
	if o != nil && !IsNil(o.StackTrace) {
		return true
	}

	return false
}

// SetStackTrace gets a reference to the given string and assigns it to the StackTrace field.
func (o *SystemInsightsCrashes) SetStackTrace(v string) {
	o.StackTrace = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *SystemInsightsCrashes) SetSystemId(v string) {
	o.SystemId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SystemInsightsCrashes) SetType(v string) {
	o.Type = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetUid() int32 {
	if o == nil || IsNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetUidOk() (*int32, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *SystemInsightsCrashes) SetUid(v int32) {
	o.Uid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SystemInsightsCrashes) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCrashes) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SystemInsightsCrashes) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SystemInsightsCrashes) SetVersion(v string) {
	o.Version = &v
}

func (o SystemInsightsCrashes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInsightsCrashes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionTime) {
		toSerialize["collection_time"] = o.CollectionTime
	}
	if !IsNil(o.CrashPath) {
		toSerialize["crash_path"] = o.CrashPath
	}
	if !IsNil(o.CrashedThread) {
		toSerialize["crashed_thread"] = o.CrashedThread
	}
	if !IsNil(o.Datetime) {
		toSerialize["datetime"] = o.Datetime
	}
	if !IsNil(o.ExceptionCodes) {
		toSerialize["exception_codes"] = o.ExceptionCodes
	}
	if !IsNil(o.ExceptionNotes) {
		toSerialize["exception_notes"] = o.ExceptionNotes
	}
	if !IsNil(o.ExceptionType) {
		toSerialize["exception_type"] = o.ExceptionType
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.Registers) {
		toSerialize["registers"] = o.Registers
	}
	if !IsNil(o.Responsible) {
		toSerialize["responsible"] = o.Responsible
	}
	if !IsNil(o.StackTrace) {
		toSerialize["stack_trace"] = o.StackTrace
	}
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemInsightsCrashes) UnmarshalJSON(bytes []byte) (err error) {
	varSystemInsightsCrashes := _SystemInsightsCrashes{}

	if err = json.Unmarshal(bytes, &varSystemInsightsCrashes); err == nil {
		*o = SystemInsightsCrashes(varSystemInsightsCrashes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "collection_time")
		delete(additionalProperties, "crash_path")
		delete(additionalProperties, "crashed_thread")
		delete(additionalProperties, "datetime")
		delete(additionalProperties, "exception_codes")
		delete(additionalProperties, "exception_notes")
		delete(additionalProperties, "exception_type")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "path")
		delete(additionalProperties, "pid")
		delete(additionalProperties, "registers")
		delete(additionalProperties, "responsible")
		delete(additionalProperties, "stack_trace")
		delete(additionalProperties, "system_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uid")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemInsightsCrashes struct {
	value *SystemInsightsCrashes
	isSet bool
}

func (v NullableSystemInsightsCrashes) Get() *SystemInsightsCrashes {
	return v.value
}

func (v *NullableSystemInsightsCrashes) Set(val *SystemInsightsCrashes) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInsightsCrashes) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInsightsCrashes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInsightsCrashes(val *SystemInsightsCrashes) *NullableSystemInsightsCrashes {
	return &NullableSystemInsightsCrashes{value: val, isSet: true}
}

func (v NullableSystemInsightsCrashes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInsightsCrashes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



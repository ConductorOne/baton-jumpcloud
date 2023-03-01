/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jcapi2

import (
	"encoding/json"
)

// checks if the SystemInsightsScheduledTasks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInsightsScheduledTasks{}

// SystemInsightsScheduledTasks struct for SystemInsightsScheduledTasks
type SystemInsightsScheduledTasks struct {
	Action *string `json:"action,omitempty"`
	Enabled *int32 `json:"enabled,omitempty"`
	Hidden *int32 `json:"hidden,omitempty"`
	LastRunCode *string `json:"last_run_code,omitempty"`
	LastRunMessage *string `json:"last_run_message,omitempty"`
	LastRunTime *string `json:"last_run_time,omitempty"`
	Name *string `json:"name,omitempty"`
	NextRunTime *string `json:"next_run_time,omitempty"`
	Path *string `json:"path,omitempty"`
	State *string `json:"state,omitempty"`
	SystemId *string `json:"system_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemInsightsScheduledTasks SystemInsightsScheduledTasks

// NewSystemInsightsScheduledTasks instantiates a new SystemInsightsScheduledTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInsightsScheduledTasks() *SystemInsightsScheduledTasks {
	this := SystemInsightsScheduledTasks{}
	return &this
}

// NewSystemInsightsScheduledTasksWithDefaults instantiates a new SystemInsightsScheduledTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInsightsScheduledTasksWithDefaults() *SystemInsightsScheduledTasks {
	this := SystemInsightsScheduledTasks{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SystemInsightsScheduledTasks) SetAction(v string) {
	o.Action = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetEnabled() int32 {
	if o == nil || IsNil(o.Enabled) {
		var ret int32
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetEnabledOk() (*int32, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given int32 and assigns it to the Enabled field.
func (o *SystemInsightsScheduledTasks) SetEnabled(v int32) {
	o.Enabled = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetHidden() int32 {
	if o == nil || IsNil(o.Hidden) {
		var ret int32
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetHiddenOk() (*int32, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given int32 and assigns it to the Hidden field.
func (o *SystemInsightsScheduledTasks) SetHidden(v int32) {
	o.Hidden = &v
}

// GetLastRunCode returns the LastRunCode field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetLastRunCode() string {
	if o == nil || IsNil(o.LastRunCode) {
		var ret string
		return ret
	}
	return *o.LastRunCode
}

// GetLastRunCodeOk returns a tuple with the LastRunCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetLastRunCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LastRunCode) {
		return nil, false
	}
	return o.LastRunCode, true
}

// HasLastRunCode returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasLastRunCode() bool {
	if o != nil && !IsNil(o.LastRunCode) {
		return true
	}

	return false
}

// SetLastRunCode gets a reference to the given string and assigns it to the LastRunCode field.
func (o *SystemInsightsScheduledTasks) SetLastRunCode(v string) {
	o.LastRunCode = &v
}

// GetLastRunMessage returns the LastRunMessage field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetLastRunMessage() string {
	if o == nil || IsNil(o.LastRunMessage) {
		var ret string
		return ret
	}
	return *o.LastRunMessage
}

// GetLastRunMessageOk returns a tuple with the LastRunMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetLastRunMessageOk() (*string, bool) {
	if o == nil || IsNil(o.LastRunMessage) {
		return nil, false
	}
	return o.LastRunMessage, true
}

// HasLastRunMessage returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasLastRunMessage() bool {
	if o != nil && !IsNil(o.LastRunMessage) {
		return true
	}

	return false
}

// SetLastRunMessage gets a reference to the given string and assigns it to the LastRunMessage field.
func (o *SystemInsightsScheduledTasks) SetLastRunMessage(v string) {
	o.LastRunMessage = &v
}

// GetLastRunTime returns the LastRunTime field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetLastRunTime() string {
	if o == nil || IsNil(o.LastRunTime) {
		var ret string
		return ret
	}
	return *o.LastRunTime
}

// GetLastRunTimeOk returns a tuple with the LastRunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetLastRunTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastRunTime) {
		return nil, false
	}
	return o.LastRunTime, true
}

// HasLastRunTime returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasLastRunTime() bool {
	if o != nil && !IsNil(o.LastRunTime) {
		return true
	}

	return false
}

// SetLastRunTime gets a reference to the given string and assigns it to the LastRunTime field.
func (o *SystemInsightsScheduledTasks) SetLastRunTime(v string) {
	o.LastRunTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SystemInsightsScheduledTasks) SetName(v string) {
	o.Name = &v
}

// GetNextRunTime returns the NextRunTime field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetNextRunTime() string {
	if o == nil || IsNil(o.NextRunTime) {
		var ret string
		return ret
	}
	return *o.NextRunTime
}

// GetNextRunTimeOk returns a tuple with the NextRunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetNextRunTimeOk() (*string, bool) {
	if o == nil || IsNil(o.NextRunTime) {
		return nil, false
	}
	return o.NextRunTime, true
}

// HasNextRunTime returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasNextRunTime() bool {
	if o != nil && !IsNil(o.NextRunTime) {
		return true
	}

	return false
}

// SetNextRunTime gets a reference to the given string and assigns it to the NextRunTime field.
func (o *SystemInsightsScheduledTasks) SetNextRunTime(v string) {
	o.NextRunTime = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SystemInsightsScheduledTasks) SetPath(v string) {
	o.Path = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SystemInsightsScheduledTasks) SetState(v string) {
	o.State = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *SystemInsightsScheduledTasks) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsScheduledTasks) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *SystemInsightsScheduledTasks) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *SystemInsightsScheduledTasks) SetSystemId(v string) {
	o.SystemId = &v
}

func (o SystemInsightsScheduledTasks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInsightsScheduledTasks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.LastRunCode) {
		toSerialize["last_run_code"] = o.LastRunCode
	}
	if !IsNil(o.LastRunMessage) {
		toSerialize["last_run_message"] = o.LastRunMessage
	}
	if !IsNil(o.LastRunTime) {
		toSerialize["last_run_time"] = o.LastRunTime
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NextRunTime) {
		toSerialize["next_run_time"] = o.NextRunTime
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemInsightsScheduledTasks) UnmarshalJSON(bytes []byte) (err error) {
	varSystemInsightsScheduledTasks := _SystemInsightsScheduledTasks{}

	if err = json.Unmarshal(bytes, &varSystemInsightsScheduledTasks); err == nil {
		*o = SystemInsightsScheduledTasks(varSystemInsightsScheduledTasks)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "last_run_code")
		delete(additionalProperties, "last_run_message")
		delete(additionalProperties, "last_run_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "next_run_time")
		delete(additionalProperties, "path")
		delete(additionalProperties, "state")
		delete(additionalProperties, "system_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemInsightsScheduledTasks struct {
	value *SystemInsightsScheduledTasks
	isSet bool
}

func (v NullableSystemInsightsScheduledTasks) Get() *SystemInsightsScheduledTasks {
	return v.value
}

func (v *NullableSystemInsightsScheduledTasks) Set(val *SystemInsightsScheduledTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInsightsScheduledTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInsightsScheduledTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInsightsScheduledTasks(val *SystemInsightsScheduledTasks) *NullableSystemInsightsScheduledTasks {
	return &NullableSystemInsightsScheduledTasks{value: val, isSet: true}
}

func (v NullableSystemInsightsScheduledTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInsightsScheduledTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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
	"time"
)

// checks if the PolicyResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyResult{}

// PolicyResult struct for PolicyResult
type PolicyResult struct {
	// Details pertaining to the policy result.
	Detail *string `json:"detail,omitempty"`
	// The end of the policy application.
	EndedAt *time.Time `json:"endedAt,omitempty"`
	// The 32-bit unsigned exit status from the applying the policy.
	ExitStatus *int32 `json:"exitStatus,omitempty"`
	// ObjectId uniquely identifying a Policy Result.
	Id *string `json:"id,omitempty"`
	// ObjectId uniquely identifying the parent Policy.
	PolicyID *string `json:"policyID,omitempty"`
	// The start of the policy application.
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Enumeration describing the state of the policy. Success, failed, or pending.
	State *string `json:"state,omitempty"`
	// The STDERR output from applying the policy.
	StdErr *string `json:"stdErr,omitempty"`
	// The STDOUT output from applying the policy.
	StdOut *string `json:"stdOut,omitempty"`
	// True if the policy was successfully applied; false otherwise.
	Success *bool `json:"success,omitempty"`
	// ObjectId uniquely identifying the parent System.
	SystemID *string `json:"systemID,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyResult PolicyResult

// NewPolicyResult instantiates a new PolicyResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyResult() *PolicyResult {
	this := PolicyResult{}
	return &this
}

// NewPolicyResultWithDefaults instantiates a new PolicyResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyResultWithDefaults() *PolicyResult {
	this := PolicyResult{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *PolicyResult) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *PolicyResult) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *PolicyResult) SetDetail(v string) {
	o.Detail = &v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *PolicyResult) GetEndedAt() time.Time {
	if o == nil || IsNil(o.EndedAt) {
		var ret time.Time
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetEndedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndedAt) {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *PolicyResult) HasEndedAt() bool {
	if o != nil && !IsNil(o.EndedAt) {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given time.Time and assigns it to the EndedAt field.
func (o *PolicyResult) SetEndedAt(v time.Time) {
	o.EndedAt = &v
}

// GetExitStatus returns the ExitStatus field value if set, zero value otherwise.
func (o *PolicyResult) GetExitStatus() int32 {
	if o == nil || IsNil(o.ExitStatus) {
		var ret int32
		return ret
	}
	return *o.ExitStatus
}

// GetExitStatusOk returns a tuple with the ExitStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetExitStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.ExitStatus) {
		return nil, false
	}
	return o.ExitStatus, true
}

// HasExitStatus returns a boolean if a field has been set.
func (o *PolicyResult) HasExitStatus() bool {
	if o != nil && !IsNil(o.ExitStatus) {
		return true
	}

	return false
}

// SetExitStatus gets a reference to the given int32 and assigns it to the ExitStatus field.
func (o *PolicyResult) SetExitStatus(v int32) {
	o.ExitStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PolicyResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PolicyResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PolicyResult) SetId(v string) {
	o.Id = &v
}

// GetPolicyID returns the PolicyID field value if set, zero value otherwise.
func (o *PolicyResult) GetPolicyID() string {
	if o == nil || IsNil(o.PolicyID) {
		var ret string
		return ret
	}
	return *o.PolicyID
}

// GetPolicyIDOk returns a tuple with the PolicyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetPolicyIDOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyID) {
		return nil, false
	}
	return o.PolicyID, true
}

// HasPolicyID returns a boolean if a field has been set.
func (o *PolicyResult) HasPolicyID() bool {
	if o != nil && !IsNil(o.PolicyID) {
		return true
	}

	return false
}

// SetPolicyID gets a reference to the given string and assigns it to the PolicyID field.
func (o *PolicyResult) SetPolicyID(v string) {
	o.PolicyID = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *PolicyResult) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *PolicyResult) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *PolicyResult) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PolicyResult) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PolicyResult) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PolicyResult) SetState(v string) {
	o.State = &v
}

// GetStdErr returns the StdErr field value if set, zero value otherwise.
func (o *PolicyResult) GetStdErr() string {
	if o == nil || IsNil(o.StdErr) {
		var ret string
		return ret
	}
	return *o.StdErr
}

// GetStdErrOk returns a tuple with the StdErr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetStdErrOk() (*string, bool) {
	if o == nil || IsNil(o.StdErr) {
		return nil, false
	}
	return o.StdErr, true
}

// HasStdErr returns a boolean if a field has been set.
func (o *PolicyResult) HasStdErr() bool {
	if o != nil && !IsNil(o.StdErr) {
		return true
	}

	return false
}

// SetStdErr gets a reference to the given string and assigns it to the StdErr field.
func (o *PolicyResult) SetStdErr(v string) {
	o.StdErr = &v
}

// GetStdOut returns the StdOut field value if set, zero value otherwise.
func (o *PolicyResult) GetStdOut() string {
	if o == nil || IsNil(o.StdOut) {
		var ret string
		return ret
	}
	return *o.StdOut
}

// GetStdOutOk returns a tuple with the StdOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetStdOutOk() (*string, bool) {
	if o == nil || IsNil(o.StdOut) {
		return nil, false
	}
	return o.StdOut, true
}

// HasStdOut returns a boolean if a field has been set.
func (o *PolicyResult) HasStdOut() bool {
	if o != nil && !IsNil(o.StdOut) {
		return true
	}

	return false
}

// SetStdOut gets a reference to the given string and assigns it to the StdOut field.
func (o *PolicyResult) SetStdOut(v string) {
	o.StdOut = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *PolicyResult) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *PolicyResult) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *PolicyResult) SetSuccess(v bool) {
	o.Success = &v
}

// GetSystemID returns the SystemID field value if set, zero value otherwise.
func (o *PolicyResult) GetSystemID() string {
	if o == nil || IsNil(o.SystemID) {
		var ret string
		return ret
	}
	return *o.SystemID
}

// GetSystemIDOk returns a tuple with the SystemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResult) GetSystemIDOk() (*string, bool) {
	if o == nil || IsNil(o.SystemID) {
		return nil, false
	}
	return o.SystemID, true
}

// HasSystemID returns a boolean if a field has been set.
func (o *PolicyResult) HasSystemID() bool {
	if o != nil && !IsNil(o.SystemID) {
		return true
	}

	return false
}

// SetSystemID gets a reference to the given string and assigns it to the SystemID field.
func (o *PolicyResult) SetSystemID(v string) {
	o.SystemID = &v
}

func (o PolicyResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.EndedAt) {
		toSerialize["endedAt"] = o.EndedAt
	}
	if !IsNil(o.ExitStatus) {
		toSerialize["exitStatus"] = o.ExitStatus
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PolicyID) {
		toSerialize["policyID"] = o.PolicyID
	}
	if !IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StdErr) {
		toSerialize["stdErr"] = o.StdErr
	}
	if !IsNil(o.StdOut) {
		toSerialize["stdOut"] = o.StdOut
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.SystemID) {
		toSerialize["systemID"] = o.SystemID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyResult) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyResult := _PolicyResult{}

	if err = json.Unmarshal(bytes, &varPolicyResult); err == nil {
		*o = PolicyResult(varPolicyResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		delete(additionalProperties, "endedAt")
		delete(additionalProperties, "exitStatus")
		delete(additionalProperties, "id")
		delete(additionalProperties, "policyID")
		delete(additionalProperties, "startedAt")
		delete(additionalProperties, "state")
		delete(additionalProperties, "stdErr")
		delete(additionalProperties, "stdOut")
		delete(additionalProperties, "success")
		delete(additionalProperties, "systemID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyResult struct {
	value *PolicyResult
	isSet bool
}

func (v NullablePolicyResult) Get() *PolicyResult {
	return v.value
}

func (v *NullablePolicyResult) Set(val *PolicyResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyResult(val *PolicyResult) *NullablePolicyResult {
	return &NullablePolicyResult{value: val, isSet: true}
}

func (v NullablePolicyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



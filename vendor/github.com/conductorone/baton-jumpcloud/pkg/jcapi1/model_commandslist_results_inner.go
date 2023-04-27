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
)

// checks if the CommandslistResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandslistResultsInner{}

// CommandslistResultsInner struct for CommandslistResultsInner
type CommandslistResultsInner struct {
	// The ID of the command.
	Id *string `json:"_id,omitempty"`
	// The Command to execute.
	Command *string `json:"command,omitempty"`
	// The Command OS.
	CommandType *string `json:"commandType,omitempty"`
	// How the Command is executed.
	LaunchType *string `json:"launchType,omitempty"`
	ListensTo *string `json:"listensTo,omitempty"`
	// The name of the Command.
	Name *string `json:"name,omitempty"`
	// The ID of the Organization.
	Organization *string `json:"organization,omitempty"`
	// A crontab that consists of: [ (seconds) (minutes) (hours) (days of month) (months) (weekdays) ] or [ immediate ]. If you send this as an empty string, it will run immediately.
	Schedule *string `json:"schedule,omitempty"`
	// When the command will repeat.
	ScheduleRepeatType *string `json:"scheduleRepeatType,omitempty"`
	// Trigger to execute command.
	Trigger *string `json:"trigger,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommandslistResultsInner CommandslistResultsInner

// NewCommandslistResultsInner instantiates a new CommandslistResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandslistResultsInner() *CommandslistResultsInner {
	this := CommandslistResultsInner{}
	return &this
}

// NewCommandslistResultsInnerWithDefaults instantiates a new CommandslistResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandslistResultsInnerWithDefaults() *CommandslistResultsInner {
	this := CommandslistResultsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CommandslistResultsInner) SetId(v string) {
	o.Id = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetCommand() string {
	if o == nil || IsNil(o.Command) {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetCommandOk() (*string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *CommandslistResultsInner) SetCommand(v string) {
	o.Command = &v
}

// GetCommandType returns the CommandType field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetCommandType() string {
	if o == nil || IsNil(o.CommandType) {
		var ret string
		return ret
	}
	return *o.CommandType
}

// GetCommandTypeOk returns a tuple with the CommandType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetCommandTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CommandType) {
		return nil, false
	}
	return o.CommandType, true
}

// HasCommandType returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasCommandType() bool {
	if o != nil && !IsNil(o.CommandType) {
		return true
	}

	return false
}

// SetCommandType gets a reference to the given string and assigns it to the CommandType field.
func (o *CommandslistResultsInner) SetCommandType(v string) {
	o.CommandType = &v
}

// GetLaunchType returns the LaunchType field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetLaunchType() string {
	if o == nil || IsNil(o.LaunchType) {
		var ret string
		return ret
	}
	return *o.LaunchType
}

// GetLaunchTypeOk returns a tuple with the LaunchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetLaunchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LaunchType) {
		return nil, false
	}
	return o.LaunchType, true
}

// HasLaunchType returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasLaunchType() bool {
	if o != nil && !IsNil(o.LaunchType) {
		return true
	}

	return false
}

// SetLaunchType gets a reference to the given string and assigns it to the LaunchType field.
func (o *CommandslistResultsInner) SetLaunchType(v string) {
	o.LaunchType = &v
}

// GetListensTo returns the ListensTo field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetListensTo() string {
	if o == nil || IsNil(o.ListensTo) {
		var ret string
		return ret
	}
	return *o.ListensTo
}

// GetListensToOk returns a tuple with the ListensTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetListensToOk() (*string, bool) {
	if o == nil || IsNil(o.ListensTo) {
		return nil, false
	}
	return o.ListensTo, true
}

// HasListensTo returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasListensTo() bool {
	if o != nil && !IsNil(o.ListensTo) {
		return true
	}

	return false
}

// SetListensTo gets a reference to the given string and assigns it to the ListensTo field.
func (o *CommandslistResultsInner) SetListensTo(v string) {
	o.ListensTo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommandslistResultsInner) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *CommandslistResultsInner) SetOrganization(v string) {
	o.Organization = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetSchedule() string {
	if o == nil || IsNil(o.Schedule) {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *CommandslistResultsInner) SetSchedule(v string) {
	o.Schedule = &v
}

// GetScheduleRepeatType returns the ScheduleRepeatType field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetScheduleRepeatType() string {
	if o == nil || IsNil(o.ScheduleRepeatType) {
		var ret string
		return ret
	}
	return *o.ScheduleRepeatType
}

// GetScheduleRepeatTypeOk returns a tuple with the ScheduleRepeatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetScheduleRepeatTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleRepeatType) {
		return nil, false
	}
	return o.ScheduleRepeatType, true
}

// HasScheduleRepeatType returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasScheduleRepeatType() bool {
	if o != nil && !IsNil(o.ScheduleRepeatType) {
		return true
	}

	return false
}

// SetScheduleRepeatType gets a reference to the given string and assigns it to the ScheduleRepeatType field.
func (o *CommandslistResultsInner) SetScheduleRepeatType(v string) {
	o.ScheduleRepeatType = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *CommandslistResultsInner) GetTrigger() string {
	if o == nil || IsNil(o.Trigger) {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandslistResultsInner) GetTriggerOk() (*string, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *CommandslistResultsInner) HasTrigger() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *CommandslistResultsInner) SetTrigger(v string) {
	o.Trigger = &v
}

func (o CommandslistResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandslistResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.CommandType) {
		toSerialize["commandType"] = o.CommandType
	}
	if !IsNil(o.LaunchType) {
		toSerialize["launchType"] = o.LaunchType
	}
	if !IsNil(o.ListensTo) {
		toSerialize["listensTo"] = o.ListensTo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.ScheduleRepeatType) {
		toSerialize["scheduleRepeatType"] = o.ScheduleRepeatType
	}
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommandslistResultsInner) UnmarshalJSON(bytes []byte) (err error) {
	varCommandslistResultsInner := _CommandslistResultsInner{}

	if err = json.Unmarshal(bytes, &varCommandslistResultsInner); err == nil {
		*o = CommandslistResultsInner(varCommandslistResultsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "command")
		delete(additionalProperties, "commandType")
		delete(additionalProperties, "launchType")
		delete(additionalProperties, "listensTo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "scheduleRepeatType")
		delete(additionalProperties, "trigger")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommandslistResultsInner struct {
	value *CommandslistResultsInner
	isSet bool
}

func (v NullableCommandslistResultsInner) Get() *CommandslistResultsInner {
	return v.value
}

func (v *NullableCommandslistResultsInner) Set(val *CommandslistResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandslistResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandslistResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandslistResultsInner(val *CommandslistResultsInner) *NullableCommandslistResultsInner {
	return &NullableCommandslistResultsInner{value: val, isSet: true}
}

func (v NullableCommandslistResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandslistResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the SyncroTicketingAlertConfigurationListRecordsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncroTicketingAlertConfigurationListRecordsInner{}

// SyncroTicketingAlertConfigurationListRecordsInner struct for SyncroTicketingAlertConfigurationListRecordsInner
type SyncroTicketingAlertConfigurationListRecordsInner struct {
	AlertId *string `json:"alertId,omitempty"`
	Category *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	DueDays *int32 `json:"dueDays,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Priority *string `json:"priority,omitempty"`
	ProblemType *string `json:"problemType,omitempty"`
	ShouldCreateTickets *bool `json:"shouldCreateTickets,omitempty"`
	Status *string `json:"status,omitempty"`
	UserId *int32 `json:"userId,omitempty"`
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyncroTicketingAlertConfigurationListRecordsInner SyncroTicketingAlertConfigurationListRecordsInner

// NewSyncroTicketingAlertConfigurationListRecordsInner instantiates a new SyncroTicketingAlertConfigurationListRecordsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncroTicketingAlertConfigurationListRecordsInner() *SyncroTicketingAlertConfigurationListRecordsInner {
	this := SyncroTicketingAlertConfigurationListRecordsInner{}
	return &this
}

// NewSyncroTicketingAlertConfigurationListRecordsInnerWithDefaults instantiates a new SyncroTicketingAlertConfigurationListRecordsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncroTicketingAlertConfigurationListRecordsInnerWithDefaults() *SyncroTicketingAlertConfigurationListRecordsInner {
	this := SyncroTicketingAlertConfigurationListRecordsInner{}
	return &this
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetAlertId() string {
	if o == nil || IsNil(o.AlertId) {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetAlertIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlertId) {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasAlertId() bool {
	if o != nil && !IsNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetAlertId(v string) {
	o.AlertId = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDueDays returns the DueDays field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetDueDays() int32 {
	if o == nil || IsNil(o.DueDays) {
		var ret int32
		return ret
	}
	return *o.DueDays
}

// GetDueDaysOk returns a tuple with the DueDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetDueDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.DueDays) {
		return nil, false
	}
	return o.DueDays, true
}

// HasDueDays returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasDueDays() bool {
	if o != nil && !IsNil(o.DueDays) {
		return true
	}

	return false
}

// SetDueDays gets a reference to the given int32 and assigns it to the DueDays field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetDueDays(v int32) {
	o.DueDays = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetId(v int32) {
	o.Id = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetPriority(v string) {
	o.Priority = &v
}

// GetProblemType returns the ProblemType field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetProblemType() string {
	if o == nil || IsNil(o.ProblemType) {
		var ret string
		return ret
	}
	return *o.ProblemType
}

// GetProblemTypeOk returns a tuple with the ProblemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetProblemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProblemType) {
		return nil, false
	}
	return o.ProblemType, true
}

// HasProblemType returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasProblemType() bool {
	if o != nil && !IsNil(o.ProblemType) {
		return true
	}

	return false
}

// SetProblemType gets a reference to the given string and assigns it to the ProblemType field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetProblemType(v string) {
	o.ProblemType = &v
}

// GetShouldCreateTickets returns the ShouldCreateTickets field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetShouldCreateTickets() bool {
	if o == nil || IsNil(o.ShouldCreateTickets) {
		var ret bool
		return ret
	}
	return *o.ShouldCreateTickets
}

// GetShouldCreateTicketsOk returns a tuple with the ShouldCreateTickets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetShouldCreateTicketsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldCreateTickets) {
		return nil, false
	}
	return o.ShouldCreateTickets, true
}

// HasShouldCreateTickets returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasShouldCreateTickets() bool {
	if o != nil && !IsNil(o.ShouldCreateTickets) {
		return true
	}

	return false
}

// SetShouldCreateTickets gets a reference to the given bool and assigns it to the ShouldCreateTickets field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetShouldCreateTickets(v bool) {
	o.ShouldCreateTickets = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetStatus(v string) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetUserId(v int32) {
	o.UserId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SyncroTicketingAlertConfigurationListRecordsInner) SetUsername(v string) {
	o.Username = &v
}

func (o SyncroTicketingAlertConfigurationListRecordsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncroTicketingAlertConfigurationListRecordsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertId) {
		toSerialize["alertId"] = o.AlertId
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DueDays) {
		toSerialize["dueDays"] = o.DueDays
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.ProblemType) {
		toSerialize["problemType"] = o.ProblemType
	}
	if !IsNil(o.ShouldCreateTickets) {
		toSerialize["shouldCreateTickets"] = o.ShouldCreateTickets
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SyncroTicketingAlertConfigurationListRecordsInner) UnmarshalJSON(bytes []byte) (err error) {
	varSyncroTicketingAlertConfigurationListRecordsInner := _SyncroTicketingAlertConfigurationListRecordsInner{}

	if err = json.Unmarshal(bytes, &varSyncroTicketingAlertConfigurationListRecordsInner); err == nil {
		*o = SyncroTicketingAlertConfigurationListRecordsInner(varSyncroTicketingAlertConfigurationListRecordsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "alertId")
		delete(additionalProperties, "category")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "dueDays")
		delete(additionalProperties, "id")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "problemType")
		delete(additionalProperties, "shouldCreateTickets")
		delete(additionalProperties, "status")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSyncroTicketingAlertConfigurationListRecordsInner struct {
	value *SyncroTicketingAlertConfigurationListRecordsInner
	isSet bool
}

func (v NullableSyncroTicketingAlertConfigurationListRecordsInner) Get() *SyncroTicketingAlertConfigurationListRecordsInner {
	return v.value
}

func (v *NullableSyncroTicketingAlertConfigurationListRecordsInner) Set(val *SyncroTicketingAlertConfigurationListRecordsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncroTicketingAlertConfigurationListRecordsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncroTicketingAlertConfigurationListRecordsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncroTicketingAlertConfigurationListRecordsInner(val *SyncroTicketingAlertConfigurationListRecordsInner) *NullableSyncroTicketingAlertConfigurationListRecordsInner {
	return &NullableSyncroTicketingAlertConfigurationListRecordsInner{value: val, isSet: true}
}

func (v NullableSyncroTicketingAlertConfigurationListRecordsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncroTicketingAlertConfigurationListRecordsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


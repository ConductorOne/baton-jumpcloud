/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/ConductorOne/baton-jumpcloud/pkg/jcapi

import (
	"encoding/json"
)

// checks if the PwmUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PwmUser{}

// PwmUser struct for PwmUser
type PwmUser struct {
	Apps []AppsInner `json:"apps,omitempty"`
	Email string `json:"email"`
	ExternalId *string `json:"externalId,omitempty"`
	Groups []GroupsInner `json:"groups,omitempty"`
	Id string `json:"id"`
	ItemsCount *int32 `json:"itemsCount,omitempty"`
	Name string `json:"name"`
	PasswordsCount *int32 `json:"passwordsCount,omitempty"`
	PasswordsScore *float32 `json:"passwordsScore,omitempty"`
	ScoreUpdatedAt *string `json:"scoreUpdatedAt,omitempty"`
	Status string `json:"status"`
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PwmUser PwmUser

// NewPwmUser instantiates a new PwmUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPwmUser(email string, id string, name string, status string) *PwmUser {
	this := PwmUser{}
	this.Email = email
	this.Id = id
	this.Name = name
	this.Status = status
	return &this
}

// NewPwmUserWithDefaults instantiates a new PwmUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPwmUserWithDefaults() *PwmUser {
	this := PwmUser{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *PwmUser) GetApps() []AppsInner {
	if o == nil || IsNil(o.Apps) {
		var ret []AppsInner
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwmUser) GetAppsOk() ([]AppsInner, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *PwmUser) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []AppsInner and assigns it to the Apps field.
func (o *PwmUser) SetApps(v []AppsInner) {
	o.Apps = v
}

// GetEmail returns the Email field value
func (o *PwmUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PwmUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PwmUser) SetEmail(v string) {
	o.Email = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *PwmUser) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwmUser) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *PwmUser) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *PwmUser) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PwmUser) GetGroups() []GroupsInner {
	if o == nil || IsNil(o.Groups) {
		var ret []GroupsInner
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwmUser) GetGroupsOk() ([]GroupsInner, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PwmUser) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupsInner and assigns it to the Groups field.
func (o *PwmUser) SetGroups(v []GroupsInner) {
	o.Groups = v
}

// GetId returns the Id field value
func (o *PwmUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PwmUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PwmUser) SetId(v string) {
	o.Id = v
}

// GetItemsCount returns the ItemsCount field value if set, zero value otherwise.
func (o *PwmUser) GetItemsCount() int32 {
	if o == nil || IsNil(o.ItemsCount) {
		var ret int32
		return ret
	}
	return *o.ItemsCount
}

// GetItemsCountOk returns a tuple with the ItemsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwmUser) GetItemsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemsCount) {
		return nil, false
	}
	return o.ItemsCount, true
}

// HasItemsCount returns a boolean if a field has been set.
func (o *PwmUser) HasItemsCount() bool {
	if o != nil && !IsNil(o.ItemsCount) {
		return true
	}

	return false
}

// SetItemsCount gets a reference to the given int32 and assigns it to the ItemsCount field.
func (o *PwmUser) SetItemsCount(v int32) {
	o.ItemsCount = &v
}

// GetName returns the Name field value
func (o *PwmUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PwmUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PwmUser) SetName(v string) {
	o.Name = v
}

// GetPasswordsCount returns the PasswordsCount field value if set, zero value otherwise.
func (o *PwmUser) GetPasswordsCount() int32 {
	if o == nil || IsNil(o.PasswordsCount) {
		var ret int32
		return ret
	}
	return *o.PasswordsCount
}

// GetPasswordsCountOk returns a tuple with the PasswordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwmUser) GetPasswordsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PasswordsCount) {
		return nil, false
	}
	return o.PasswordsCount, true
}

// HasPasswordsCount returns a boolean if a field has been set.
func (o *PwmUser) HasPasswordsCount() bool {
	if o != nil && !IsNil(o.PasswordsCount) {
		return true
	}

	return false
}

// SetPasswordsCount gets a reference to the given int32 and assigns it to the PasswordsCount field.
func (o *PwmUser) SetPasswordsCount(v int32) {
	o.PasswordsCount = &v
}

// GetPasswordsScore returns the PasswordsScore field value if set, zero value otherwise.
func (o *PwmUser) GetPasswordsScore() float32 {
	if o == nil || IsNil(o.PasswordsScore) {
		var ret float32
		return ret
	}
	return *o.PasswordsScore
}

// GetPasswordsScoreOk returns a tuple with the PasswordsScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwmUser) GetPasswordsScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.PasswordsScore) {
		return nil, false
	}
	return o.PasswordsScore, true
}

// HasPasswordsScore returns a boolean if a field has been set.
func (o *PwmUser) HasPasswordsScore() bool {
	if o != nil && !IsNil(o.PasswordsScore) {
		return true
	}

	return false
}

// SetPasswordsScore gets a reference to the given float32 and assigns it to the PasswordsScore field.
func (o *PwmUser) SetPasswordsScore(v float32) {
	o.PasswordsScore = &v
}

// GetScoreUpdatedAt returns the ScoreUpdatedAt field value if set, zero value otherwise.
func (o *PwmUser) GetScoreUpdatedAt() string {
	if o == nil || IsNil(o.ScoreUpdatedAt) {
		var ret string
		return ret
	}
	return *o.ScoreUpdatedAt
}

// GetScoreUpdatedAtOk returns a tuple with the ScoreUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwmUser) GetScoreUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ScoreUpdatedAt) {
		return nil, false
	}
	return o.ScoreUpdatedAt, true
}

// HasScoreUpdatedAt returns a boolean if a field has been set.
func (o *PwmUser) HasScoreUpdatedAt() bool {
	if o != nil && !IsNil(o.ScoreUpdatedAt) {
		return true
	}

	return false
}

// SetScoreUpdatedAt gets a reference to the given string and assigns it to the ScoreUpdatedAt field.
func (o *PwmUser) SetScoreUpdatedAt(v string) {
	o.ScoreUpdatedAt = &v
}

// GetStatus returns the Status field value
func (o *PwmUser) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PwmUser) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PwmUser) SetStatus(v string) {
	o.Status = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PwmUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwmUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PwmUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PwmUser) SetUsername(v string) {
	o.Username = &v
}

func (o PwmUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PwmUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.ItemsCount) {
		toSerialize["itemsCount"] = o.ItemsCount
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PasswordsCount) {
		toSerialize["passwordsCount"] = o.PasswordsCount
	}
	if !IsNil(o.PasswordsScore) {
		toSerialize["passwordsScore"] = o.PasswordsScore
	}
	if !IsNil(o.ScoreUpdatedAt) {
		toSerialize["scoreUpdatedAt"] = o.ScoreUpdatedAt
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PwmUser) UnmarshalJSON(bytes []byte) (err error) {
	varPwmUser := _PwmUser{}

	if err = json.Unmarshal(bytes, &varPwmUser); err == nil {
		*o = PwmUser(varPwmUser)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "apps")
		delete(additionalProperties, "email")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "id")
		delete(additionalProperties, "itemsCount")
		delete(additionalProperties, "name")
		delete(additionalProperties, "passwordsCount")
		delete(additionalProperties, "passwordsScore")
		delete(additionalProperties, "scoreUpdatedAt")
		delete(additionalProperties, "status")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePwmUser struct {
	value *PwmUser
	isSet bool
}

func (v NullablePwmUser) Get() *PwmUser {
	return v.value
}

func (v *NullablePwmUser) Set(val *PwmUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePwmUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePwmUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePwmUser(val *PwmUser) *NullablePwmUser {
	return &NullablePwmUser{value: val, isSet: true}
}

func (v NullablePwmUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePwmUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



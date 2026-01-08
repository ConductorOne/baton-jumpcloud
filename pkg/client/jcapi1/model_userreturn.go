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

// checks if the Userreturn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Userreturn{}

// Userreturn struct for Userreturn
type Userreturn struct {
	Id *string `json:"_id,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	DisableIntroduction *bool `json:"disableIntroduction,omitempty"`
	Email *string `json:"email,omitempty"`
	EnableMultiFactor *bool `json:"enableMultiFactor,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	GrowthData *UserreturnGrowthData `json:"growthData,omitempty"`
	LastWhatsNewChecked *time.Time `json:"lastWhatsNewChecked,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Role *string `json:"role,omitempty"`
	RoleName *string `json:"roleName,omitempty"`
	SessionCount *int32 `json:"sessionCount,omitempty"`
	Suspended *bool `json:"suspended,omitempty"`
	TotpEnrolled *bool `json:"totpEnrolled,omitempty"`
	UsersTimeZone *string `json:"usersTimeZone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Userreturn Userreturn

// NewUserreturn instantiates a new Userreturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserreturn() *Userreturn {
	this := Userreturn{}
	return &this
}

// NewUserreturnWithDefaults instantiates a new Userreturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserreturnWithDefaults() *Userreturn {
	this := Userreturn{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Userreturn) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Userreturn) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Userreturn) SetId(v string) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Userreturn) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Userreturn) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Userreturn) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDisableIntroduction returns the DisableIntroduction field value if set, zero value otherwise.
func (o *Userreturn) GetDisableIntroduction() bool {
	if o == nil || IsNil(o.DisableIntroduction) {
		var ret bool
		return ret
	}
	return *o.DisableIntroduction
}

// GetDisableIntroductionOk returns a tuple with the DisableIntroduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetDisableIntroductionOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableIntroduction) {
		return nil, false
	}
	return o.DisableIntroduction, true
}

// HasDisableIntroduction returns a boolean if a field has been set.
func (o *Userreturn) HasDisableIntroduction() bool {
	if o != nil && !IsNil(o.DisableIntroduction) {
		return true
	}

	return false
}

// SetDisableIntroduction gets a reference to the given bool and assigns it to the DisableIntroduction field.
func (o *Userreturn) SetDisableIntroduction(v bool) {
	o.DisableIntroduction = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Userreturn) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Userreturn) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Userreturn) SetEmail(v string) {
	o.Email = &v
}

// GetEnableMultiFactor returns the EnableMultiFactor field value if set, zero value otherwise.
func (o *Userreturn) GetEnableMultiFactor() bool {
	if o == nil || IsNil(o.EnableMultiFactor) {
		var ret bool
		return ret
	}
	return *o.EnableMultiFactor
}

// GetEnableMultiFactorOk returns a tuple with the EnableMultiFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetEnableMultiFactorOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMultiFactor) {
		return nil, false
	}
	return o.EnableMultiFactor, true
}

// HasEnableMultiFactor returns a boolean if a field has been set.
func (o *Userreturn) HasEnableMultiFactor() bool {
	if o != nil && !IsNil(o.EnableMultiFactor) {
		return true
	}

	return false
}

// SetEnableMultiFactor gets a reference to the given bool and assigns it to the EnableMultiFactor field.
func (o *Userreturn) SetEnableMultiFactor(v bool) {
	o.EnableMultiFactor = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *Userreturn) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *Userreturn) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *Userreturn) SetFirstname(v string) {
	o.Firstname = &v
}

// GetGrowthData returns the GrowthData field value if set, zero value otherwise.
func (o *Userreturn) GetGrowthData() UserreturnGrowthData {
	if o == nil || IsNil(o.GrowthData) {
		var ret UserreturnGrowthData
		return ret
	}
	return *o.GrowthData
}

// GetGrowthDataOk returns a tuple with the GrowthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetGrowthDataOk() (*UserreturnGrowthData, bool) {
	if o == nil || IsNil(o.GrowthData) {
		return nil, false
	}
	return o.GrowthData, true
}

// HasGrowthData returns a boolean if a field has been set.
func (o *Userreturn) HasGrowthData() bool {
	if o != nil && !IsNil(o.GrowthData) {
		return true
	}

	return false
}

// SetGrowthData gets a reference to the given UserreturnGrowthData and assigns it to the GrowthData field.
func (o *Userreturn) SetGrowthData(v UserreturnGrowthData) {
	o.GrowthData = &v
}

// GetLastWhatsNewChecked returns the LastWhatsNewChecked field value if set, zero value otherwise.
func (o *Userreturn) GetLastWhatsNewChecked() time.Time {
	if o == nil || IsNil(o.LastWhatsNewChecked) {
		var ret time.Time
		return ret
	}
	return *o.LastWhatsNewChecked
}

// GetLastWhatsNewCheckedOk returns a tuple with the LastWhatsNewChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetLastWhatsNewCheckedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastWhatsNewChecked) {
		return nil, false
	}
	return o.LastWhatsNewChecked, true
}

// HasLastWhatsNewChecked returns a boolean if a field has been set.
func (o *Userreturn) HasLastWhatsNewChecked() bool {
	if o != nil && !IsNil(o.LastWhatsNewChecked) {
		return true
	}

	return false
}

// SetLastWhatsNewChecked gets a reference to the given time.Time and assigns it to the LastWhatsNewChecked field.
func (o *Userreturn) SetLastWhatsNewChecked(v time.Time) {
	o.LastWhatsNewChecked = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *Userreturn) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *Userreturn) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *Userreturn) SetLastname(v string) {
	o.Lastname = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Userreturn) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Userreturn) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *Userreturn) SetOrganization(v string) {
	o.Organization = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *Userreturn) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *Userreturn) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *Userreturn) SetProvider(v string) {
	o.Provider = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Userreturn) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Userreturn) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *Userreturn) SetRole(v string) {
	o.Role = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *Userreturn) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *Userreturn) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *Userreturn) SetRoleName(v string) {
	o.RoleName = &v
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise.
func (o *Userreturn) GetSessionCount() int32 {
	if o == nil || IsNil(o.SessionCount) {
		var ret int32
		return ret
	}
	return *o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetSessionCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SessionCount) {
		return nil, false
	}
	return o.SessionCount, true
}

// HasSessionCount returns a boolean if a field has been set.
func (o *Userreturn) HasSessionCount() bool {
	if o != nil && !IsNil(o.SessionCount) {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given int32 and assigns it to the SessionCount field.
func (o *Userreturn) SetSessionCount(v int32) {
	o.SessionCount = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *Userreturn) GetSuspended() bool {
	if o == nil || IsNil(o.Suspended) {
		var ret bool
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetSuspendedOk() (*bool, bool) {
	if o == nil || IsNil(o.Suspended) {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *Userreturn) HasSuspended() bool {
	if o != nil && !IsNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given bool and assigns it to the Suspended field.
func (o *Userreturn) SetSuspended(v bool) {
	o.Suspended = &v
}

// GetTotpEnrolled returns the TotpEnrolled field value if set, zero value otherwise.
func (o *Userreturn) GetTotpEnrolled() bool {
	if o == nil || IsNil(o.TotpEnrolled) {
		var ret bool
		return ret
	}
	return *o.TotpEnrolled
}

// GetTotpEnrolledOk returns a tuple with the TotpEnrolled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetTotpEnrolledOk() (*bool, bool) {
	if o == nil || IsNil(o.TotpEnrolled) {
		return nil, false
	}
	return o.TotpEnrolled, true
}

// HasTotpEnrolled returns a boolean if a field has been set.
func (o *Userreturn) HasTotpEnrolled() bool {
	if o != nil && !IsNil(o.TotpEnrolled) {
		return true
	}

	return false
}

// SetTotpEnrolled gets a reference to the given bool and assigns it to the TotpEnrolled field.
func (o *Userreturn) SetTotpEnrolled(v bool) {
	o.TotpEnrolled = &v
}

// GetUsersTimeZone returns the UsersTimeZone field value if set, zero value otherwise.
func (o *Userreturn) GetUsersTimeZone() string {
	if o == nil || IsNil(o.UsersTimeZone) {
		var ret string
		return ret
	}
	return *o.UsersTimeZone
}

// GetUsersTimeZoneOk returns a tuple with the UsersTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userreturn) GetUsersTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.UsersTimeZone) {
		return nil, false
	}
	return o.UsersTimeZone, true
}

// HasUsersTimeZone returns a boolean if a field has been set.
func (o *Userreturn) HasUsersTimeZone() bool {
	if o != nil && !IsNil(o.UsersTimeZone) {
		return true
	}

	return false
}

// SetUsersTimeZone gets a reference to the given string and assigns it to the UsersTimeZone field.
func (o *Userreturn) SetUsersTimeZone(v string) {
	o.UsersTimeZone = &v
}

func (o Userreturn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Userreturn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.DisableIntroduction) {
		toSerialize["disableIntroduction"] = o.DisableIntroduction
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EnableMultiFactor) {
		toSerialize["enableMultiFactor"] = o.EnableMultiFactor
	}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.GrowthData) {
		toSerialize["growthData"] = o.GrowthData
	}
	if !IsNil(o.LastWhatsNewChecked) {
		toSerialize["lastWhatsNewChecked"] = o.LastWhatsNewChecked
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.RoleName) {
		toSerialize["roleName"] = o.RoleName
	}
	if !IsNil(o.SessionCount) {
		toSerialize["sessionCount"] = o.SessionCount
	}
	if !IsNil(o.Suspended) {
		toSerialize["suspended"] = o.Suspended
	}
	if !IsNil(o.TotpEnrolled) {
		toSerialize["totpEnrolled"] = o.TotpEnrolled
	}
	if !IsNil(o.UsersTimeZone) {
		toSerialize["usersTimeZone"] = o.UsersTimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Userreturn) UnmarshalJSON(bytes []byte) (err error) {
	varUserreturn := _Userreturn{}

	if err = json.Unmarshal(bytes, &varUserreturn); err == nil {
		*o = Userreturn(varUserreturn)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "disableIntroduction")
		delete(additionalProperties, "email")
		delete(additionalProperties, "enableMultiFactor")
		delete(additionalProperties, "firstname")
		delete(additionalProperties, "growthData")
		delete(additionalProperties, "lastWhatsNewChecked")
		delete(additionalProperties, "lastname")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "role")
		delete(additionalProperties, "roleName")
		delete(additionalProperties, "sessionCount")
		delete(additionalProperties, "suspended")
		delete(additionalProperties, "totpEnrolled")
		delete(additionalProperties, "usersTimeZone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserreturn struct {
	value *Userreturn
	isSet bool
}

func (v NullableUserreturn) Get() *Userreturn {
	return v.value
}

func (v *NullableUserreturn) Set(val *Userreturn) {
	v.value = val
	v.isSet = true
}

func (v NullableUserreturn) IsSet() bool {
	return v.isSet
}

func (v *NullableUserreturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserreturn(val *Userreturn) *NullableUserreturn {
	return &NullableUserreturn{value: val, isSet: true}
}

func (v NullableUserreturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserreturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



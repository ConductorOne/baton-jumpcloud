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

// checks if the ImportUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportUser{}

// ImportUser struct for ImportUser
type ImportUser struct {
	Addresses []ImportUserAddress `json:"addresses,omitempty"`
	AlternateEmail *string `json:"alternateEmail,omitempty"`
	Company *string `json:"company,omitempty"`
	CostCenter *string `json:"costCenter,omitempty"`
	Department *string `json:"department,omitempty"`
	Displayname *string `json:"displayname,omitempty"`
	Email *string `json:"email,omitempty"`
	EmployeeIdentifier *string `json:"employeeIdentifier,omitempty"`
	EmployeeType *string `json:"employeeType,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	Id *string `json:"id,omitempty"`
	JobTitle *string `json:"jobTitle,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Location *string `json:"location,omitempty"`
	Manager *string `json:"manager,omitempty"`
	Middlename *string `json:"middlename,omitempty"`
	PhoneNumbers []ImportUserPhoneNumber `json:"phoneNumbers,omitempty"`
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportUser ImportUser

// NewImportUser instantiates a new ImportUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportUser() *ImportUser {
	this := ImportUser{}
	return &this
}

// NewImportUserWithDefaults instantiates a new ImportUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportUserWithDefaults() *ImportUser {
	this := ImportUser{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ImportUser) GetAddresses() []ImportUserAddress {
	if o == nil || IsNil(o.Addresses) {
		var ret []ImportUserAddress
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetAddressesOk() ([]ImportUserAddress, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ImportUser) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []ImportUserAddress and assigns it to the Addresses field.
func (o *ImportUser) SetAddresses(v []ImportUserAddress) {
	o.Addresses = v
}

// GetAlternateEmail returns the AlternateEmail field value if set, zero value otherwise.
func (o *ImportUser) GetAlternateEmail() string {
	if o == nil || IsNil(o.AlternateEmail) {
		var ret string
		return ret
	}
	return *o.AlternateEmail
}

// GetAlternateEmailOk returns a tuple with the AlternateEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetAlternateEmailOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateEmail) {
		return nil, false
	}
	return o.AlternateEmail, true
}

// HasAlternateEmail returns a boolean if a field has been set.
func (o *ImportUser) HasAlternateEmail() bool {
	if o != nil && !IsNil(o.AlternateEmail) {
		return true
	}

	return false
}

// SetAlternateEmail gets a reference to the given string and assigns it to the AlternateEmail field.
func (o *ImportUser) SetAlternateEmail(v string) {
	o.AlternateEmail = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ImportUser) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ImportUser) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *ImportUser) SetCompany(v string) {
	o.Company = &v
}

// GetCostCenter returns the CostCenter field value if set, zero value otherwise.
func (o *ImportUser) GetCostCenter() string {
	if o == nil || IsNil(o.CostCenter) {
		var ret string
		return ret
	}
	return *o.CostCenter
}

// GetCostCenterOk returns a tuple with the CostCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetCostCenterOk() (*string, bool) {
	if o == nil || IsNil(o.CostCenter) {
		return nil, false
	}
	return o.CostCenter, true
}

// HasCostCenter returns a boolean if a field has been set.
func (o *ImportUser) HasCostCenter() bool {
	if o != nil && !IsNil(o.CostCenter) {
		return true
	}

	return false
}

// SetCostCenter gets a reference to the given string and assigns it to the CostCenter field.
func (o *ImportUser) SetCostCenter(v string) {
	o.CostCenter = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *ImportUser) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *ImportUser) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *ImportUser) SetDepartment(v string) {
	o.Department = &v
}

// GetDisplayname returns the Displayname field value if set, zero value otherwise.
func (o *ImportUser) GetDisplayname() string {
	if o == nil || IsNil(o.Displayname) {
		var ret string
		return ret
	}
	return *o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetDisplaynameOk() (*string, bool) {
	if o == nil || IsNil(o.Displayname) {
		return nil, false
	}
	return o.Displayname, true
}

// HasDisplayname returns a boolean if a field has been set.
func (o *ImportUser) HasDisplayname() bool {
	if o != nil && !IsNil(o.Displayname) {
		return true
	}

	return false
}

// SetDisplayname gets a reference to the given string and assigns it to the Displayname field.
func (o *ImportUser) SetDisplayname(v string) {
	o.Displayname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ImportUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ImportUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ImportUser) SetEmail(v string) {
	o.Email = &v
}

// GetEmployeeIdentifier returns the EmployeeIdentifier field value if set, zero value otherwise.
func (o *ImportUser) GetEmployeeIdentifier() string {
	if o == nil || IsNil(o.EmployeeIdentifier) {
		var ret string
		return ret
	}
	return *o.EmployeeIdentifier
}

// GetEmployeeIdentifierOk returns a tuple with the EmployeeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetEmployeeIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeIdentifier) {
		return nil, false
	}
	return o.EmployeeIdentifier, true
}

// HasEmployeeIdentifier returns a boolean if a field has been set.
func (o *ImportUser) HasEmployeeIdentifier() bool {
	if o != nil && !IsNil(o.EmployeeIdentifier) {
		return true
	}

	return false
}

// SetEmployeeIdentifier gets a reference to the given string and assigns it to the EmployeeIdentifier field.
func (o *ImportUser) SetEmployeeIdentifier(v string) {
	o.EmployeeIdentifier = &v
}

// GetEmployeeType returns the EmployeeType field value if set, zero value otherwise.
func (o *ImportUser) GetEmployeeType() string {
	if o == nil || IsNil(o.EmployeeType) {
		var ret string
		return ret
	}
	return *o.EmployeeType
}

// GetEmployeeTypeOk returns a tuple with the EmployeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetEmployeeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeType) {
		return nil, false
	}
	return o.EmployeeType, true
}

// HasEmployeeType returns a boolean if a field has been set.
func (o *ImportUser) HasEmployeeType() bool {
	if o != nil && !IsNil(o.EmployeeType) {
		return true
	}

	return false
}

// SetEmployeeType gets a reference to the given string and assigns it to the EmployeeType field.
func (o *ImportUser) SetEmployeeType(v string) {
	o.EmployeeType = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *ImportUser) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *ImportUser) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *ImportUser) SetFirstname(v string) {
	o.Firstname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImportUser) SetId(v string) {
	o.Id = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *ImportUser) GetJobTitle() string {
	if o == nil || IsNil(o.JobTitle) {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetJobTitleOk() (*string, bool) {
	if o == nil || IsNil(o.JobTitle) {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *ImportUser) HasJobTitle() bool {
	if o != nil && !IsNil(o.JobTitle) {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *ImportUser) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *ImportUser) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *ImportUser) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *ImportUser) SetLastname(v string) {
	o.Lastname = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ImportUser) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ImportUser) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ImportUser) SetLocation(v string) {
	o.Location = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *ImportUser) GetManager() string {
	if o == nil || IsNil(o.Manager) {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetManagerOk() (*string, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *ImportUser) HasManager() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given string and assigns it to the Manager field.
func (o *ImportUser) SetManager(v string) {
	o.Manager = &v
}

// GetMiddlename returns the Middlename field value if set, zero value otherwise.
func (o *ImportUser) GetMiddlename() string {
	if o == nil || IsNil(o.Middlename) {
		var ret string
		return ret
	}
	return *o.Middlename
}

// GetMiddlenameOk returns a tuple with the Middlename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetMiddlenameOk() (*string, bool) {
	if o == nil || IsNil(o.Middlename) {
		return nil, false
	}
	return o.Middlename, true
}

// HasMiddlename returns a boolean if a field has been set.
func (o *ImportUser) HasMiddlename() bool {
	if o != nil && !IsNil(o.Middlename) {
		return true
	}

	return false
}

// SetMiddlename gets a reference to the given string and assigns it to the Middlename field.
func (o *ImportUser) SetMiddlename(v string) {
	o.Middlename = &v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *ImportUser) GetPhoneNumbers() []ImportUserPhoneNumber {
	if o == nil || IsNil(o.PhoneNumbers) {
		var ret []ImportUserPhoneNumber
		return ret
	}
	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetPhoneNumbersOk() ([]ImportUserPhoneNumber, bool) {
	if o == nil || IsNil(o.PhoneNumbers) {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *ImportUser) HasPhoneNumbers() bool {
	if o != nil && !IsNil(o.PhoneNumbers) {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []ImportUserPhoneNumber and assigns it to the PhoneNumbers field.
func (o *ImportUser) SetPhoneNumbers(v []ImportUserPhoneNumber) {
	o.PhoneNumbers = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ImportUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ImportUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ImportUser) SetUsername(v string) {
	o.Username = &v
}

func (o ImportUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.AlternateEmail) {
		toSerialize["alternateEmail"] = o.AlternateEmail
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.CostCenter) {
		toSerialize["costCenter"] = o.CostCenter
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.Displayname) {
		toSerialize["displayname"] = o.Displayname
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EmployeeIdentifier) {
		toSerialize["employeeIdentifier"] = o.EmployeeIdentifier
	}
	if !IsNil(o.EmployeeType) {
		toSerialize["employeeType"] = o.EmployeeType
	}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.JobTitle) {
		toSerialize["jobTitle"] = o.JobTitle
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if !IsNil(o.Middlename) {
		toSerialize["middlename"] = o.Middlename
	}
	if !IsNil(o.PhoneNumbers) {
		toSerialize["phoneNumbers"] = o.PhoneNumbers
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportUser) UnmarshalJSON(bytes []byte) (err error) {
	varImportUser := _ImportUser{}

	if err = json.Unmarshal(bytes, &varImportUser); err == nil {
		*o = ImportUser(varImportUser)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "addresses")
		delete(additionalProperties, "alternateEmail")
		delete(additionalProperties, "company")
		delete(additionalProperties, "costCenter")
		delete(additionalProperties, "department")
		delete(additionalProperties, "displayname")
		delete(additionalProperties, "email")
		delete(additionalProperties, "employeeIdentifier")
		delete(additionalProperties, "employeeType")
		delete(additionalProperties, "firstname")
		delete(additionalProperties, "id")
		delete(additionalProperties, "jobTitle")
		delete(additionalProperties, "lastname")
		delete(additionalProperties, "location")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "middlename")
		delete(additionalProperties, "phoneNumbers")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportUser struct {
	value *ImportUser
	isSet bool
}

func (v NullableImportUser) Get() *ImportUser {
	return v.value
}

func (v *NullableImportUser) Set(val *ImportUser) {
	v.value = val
	v.isSet = true
}

func (v NullableImportUser) IsSet() bool {
	return v.isSet
}

func (v *NullableImportUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportUser(val *ImportUser) *NullableImportUser {
	return &NullableImportUser{value: val, isSet: true}
}

func (v NullableImportUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



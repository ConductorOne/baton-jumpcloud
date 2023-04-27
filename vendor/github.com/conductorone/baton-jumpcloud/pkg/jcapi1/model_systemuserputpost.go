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

// checks if the Systemuserputpost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Systemuserputpost{}

// Systemuserputpost struct for Systemuserputpost
type Systemuserputpost struct {
	AccountLocked *bool `json:"account_locked,omitempty"`
	Activated *bool `json:"activated,omitempty"`
	Addresses []SystemuserputpostAddressesInner `json:"addresses,omitempty"`
	AllowPublicKey *bool `json:"allow_public_key,omitempty"`
	AlternateEmail *string `json:"alternateEmail,omitempty"`
	Attributes []SystemuserputAttributesInner `json:"attributes,omitempty"`
	Company *string `json:"company,omitempty"`
	CostCenter *string `json:"costCenter,omitempty"`
	Department *string `json:"department,omitempty"`
	Description *string `json:"description,omitempty"`
	DisableDeviceMaxLoginAttempts *bool `json:"disableDeviceMaxLoginAttempts,omitempty"`
	Displayname *string `json:"displayname,omitempty"`
	Email string `json:"email"`
	// Must be unique per user. 
	EmployeeIdentifier *string `json:"employeeIdentifier,omitempty"`
	EmployeeType *string `json:"employeeType,omitempty"`
	EnableManagedUid *bool `json:"enable_managed_uid,omitempty"`
	EnableUserPortalMultifactor *bool `json:"enable_user_portal_multifactor,omitempty"`
	ExternalDn *string `json:"external_dn,omitempty"`
	ExternalPasswordExpirationDate *time.Time `json:"external_password_expiration_date,omitempty"`
	ExternalSourceType *string `json:"external_source_type,omitempty"`
	ExternallyManaged *bool `json:"externally_managed,omitempty"`
	Firstname *string `json:"firstname,omitempty"`
	JobTitle *string `json:"jobTitle,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	LdapBindingUser *bool `json:"ldap_binding_user,omitempty"`
	Location *string `json:"location,omitempty"`
	ManagedAppleId *string `json:"managedAppleId,omitempty"`
	// Relation with another systemuser to identify the last as a manager.
	Manager *string `json:"manager,omitempty"`
	Mfa *Mfa `json:"mfa,omitempty"`
	Middlename *string `json:"middlename,omitempty"`
	Password *string `json:"password,omitempty"`
	PasswordNeverExpires *bool `json:"password_never_expires,omitempty"`
	PasswordlessSudo *bool `json:"passwordless_sudo,omitempty"`
	PhoneNumbers []SystemuserputpostPhoneNumbersInner `json:"phoneNumbers,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
	RecoveryEmail *SystemuserputpostRecoveryEmail `json:"recoveryEmail,omitempty"`
	Relationships []SystemuserputRelationshipsInner `json:"relationships,omitempty"`
	SambaServiceUser *bool `json:"samba_service_user,omitempty"`
	State *string `json:"state,omitempty"`
	Sudo *bool `json:"sudo,omitempty"`
	Suspended *bool `json:"suspended,omitempty"`
	Tags []string `json:"tags,omitempty"`
	UnixGuid *int32 `json:"unix_guid,omitempty"`
	UnixUid *int32 `json:"unix_uid,omitempty"`
	Username string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _Systemuserputpost Systemuserputpost

// NewSystemuserputpost instantiates a new Systemuserputpost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemuserputpost(email string, username string) *Systemuserputpost {
	this := Systemuserputpost{}
	this.Email = email
	this.Username = username
	return &this
}

// NewSystemuserputpostWithDefaults instantiates a new Systemuserputpost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemuserputpostWithDefaults() *Systemuserputpost {
	this := Systemuserputpost{}
	return &this
}

// GetAccountLocked returns the AccountLocked field value if set, zero value otherwise.
func (o *Systemuserputpost) GetAccountLocked() bool {
	if o == nil || IsNil(o.AccountLocked) {
		var ret bool
		return ret
	}
	return *o.AccountLocked
}

// GetAccountLockedOk returns a tuple with the AccountLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetAccountLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccountLocked) {
		return nil, false
	}
	return o.AccountLocked, true
}

// HasAccountLocked returns a boolean if a field has been set.
func (o *Systemuserputpost) HasAccountLocked() bool {
	if o != nil && !IsNil(o.AccountLocked) {
		return true
	}

	return false
}

// SetAccountLocked gets a reference to the given bool and assigns it to the AccountLocked field.
func (o *Systemuserputpost) SetAccountLocked(v bool) {
	o.AccountLocked = &v
}

// GetActivated returns the Activated field value if set, zero value otherwise.
func (o *Systemuserputpost) GetActivated() bool {
	if o == nil || IsNil(o.Activated) {
		var ret bool
		return ret
	}
	return *o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetActivatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Activated) {
		return nil, false
	}
	return o.Activated, true
}

// HasActivated returns a boolean if a field has been set.
func (o *Systemuserputpost) HasActivated() bool {
	if o != nil && !IsNil(o.Activated) {
		return true
	}

	return false
}

// SetActivated gets a reference to the given bool and assigns it to the Activated field.
func (o *Systemuserputpost) SetActivated(v bool) {
	o.Activated = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *Systemuserputpost) GetAddresses() []SystemuserputpostAddressesInner {
	if o == nil || IsNil(o.Addresses) {
		var ret []SystemuserputpostAddressesInner
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetAddressesOk() ([]SystemuserputpostAddressesInner, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *Systemuserputpost) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []SystemuserputpostAddressesInner and assigns it to the Addresses field.
func (o *Systemuserputpost) SetAddresses(v []SystemuserputpostAddressesInner) {
	o.Addresses = v
}

// GetAllowPublicKey returns the AllowPublicKey field value if set, zero value otherwise.
func (o *Systemuserputpost) GetAllowPublicKey() bool {
	if o == nil || IsNil(o.AllowPublicKey) {
		var ret bool
		return ret
	}
	return *o.AllowPublicKey
}

// GetAllowPublicKeyOk returns a tuple with the AllowPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetAllowPublicKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPublicKey) {
		return nil, false
	}
	return o.AllowPublicKey, true
}

// HasAllowPublicKey returns a boolean if a field has been set.
func (o *Systemuserputpost) HasAllowPublicKey() bool {
	if o != nil && !IsNil(o.AllowPublicKey) {
		return true
	}

	return false
}

// SetAllowPublicKey gets a reference to the given bool and assigns it to the AllowPublicKey field.
func (o *Systemuserputpost) SetAllowPublicKey(v bool) {
	o.AllowPublicKey = &v
}

// GetAlternateEmail returns the AlternateEmail field value if set, zero value otherwise.
func (o *Systemuserputpost) GetAlternateEmail() string {
	if o == nil || IsNil(o.AlternateEmail) {
		var ret string
		return ret
	}
	return *o.AlternateEmail
}

// GetAlternateEmailOk returns a tuple with the AlternateEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetAlternateEmailOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateEmail) {
		return nil, false
	}
	return o.AlternateEmail, true
}

// HasAlternateEmail returns a boolean if a field has been set.
func (o *Systemuserputpost) HasAlternateEmail() bool {
	if o != nil && !IsNil(o.AlternateEmail) {
		return true
	}

	return false
}

// SetAlternateEmail gets a reference to the given string and assigns it to the AlternateEmail field.
func (o *Systemuserputpost) SetAlternateEmail(v string) {
	o.AlternateEmail = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Systemuserputpost) GetAttributes() []SystemuserputAttributesInner {
	if o == nil || IsNil(o.Attributes) {
		var ret []SystemuserputAttributesInner
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetAttributesOk() ([]SystemuserputAttributesInner, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Systemuserputpost) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []SystemuserputAttributesInner and assigns it to the Attributes field.
func (o *Systemuserputpost) SetAttributes(v []SystemuserputAttributesInner) {
	o.Attributes = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *Systemuserputpost) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *Systemuserputpost) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *Systemuserputpost) SetCompany(v string) {
	o.Company = &v
}

// GetCostCenter returns the CostCenter field value if set, zero value otherwise.
func (o *Systemuserputpost) GetCostCenter() string {
	if o == nil || IsNil(o.CostCenter) {
		var ret string
		return ret
	}
	return *o.CostCenter
}

// GetCostCenterOk returns a tuple with the CostCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetCostCenterOk() (*string, bool) {
	if o == nil || IsNil(o.CostCenter) {
		return nil, false
	}
	return o.CostCenter, true
}

// HasCostCenter returns a boolean if a field has been set.
func (o *Systemuserputpost) HasCostCenter() bool {
	if o != nil && !IsNil(o.CostCenter) {
		return true
	}

	return false
}

// SetCostCenter gets a reference to the given string and assigns it to the CostCenter field.
func (o *Systemuserputpost) SetCostCenter(v string) {
	o.CostCenter = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *Systemuserputpost) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *Systemuserputpost) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *Systemuserputpost) SetDepartment(v string) {
	o.Department = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Systemuserputpost) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Systemuserputpost) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Systemuserputpost) SetDescription(v string) {
	o.Description = &v
}

// GetDisableDeviceMaxLoginAttempts returns the DisableDeviceMaxLoginAttempts field value if set, zero value otherwise.
func (o *Systemuserputpost) GetDisableDeviceMaxLoginAttempts() bool {
	if o == nil || IsNil(o.DisableDeviceMaxLoginAttempts) {
		var ret bool
		return ret
	}
	return *o.DisableDeviceMaxLoginAttempts
}

// GetDisableDeviceMaxLoginAttemptsOk returns a tuple with the DisableDeviceMaxLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetDisableDeviceMaxLoginAttemptsOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableDeviceMaxLoginAttempts) {
		return nil, false
	}
	return o.DisableDeviceMaxLoginAttempts, true
}

// HasDisableDeviceMaxLoginAttempts returns a boolean if a field has been set.
func (o *Systemuserputpost) HasDisableDeviceMaxLoginAttempts() bool {
	if o != nil && !IsNil(o.DisableDeviceMaxLoginAttempts) {
		return true
	}

	return false
}

// SetDisableDeviceMaxLoginAttempts gets a reference to the given bool and assigns it to the DisableDeviceMaxLoginAttempts field.
func (o *Systemuserputpost) SetDisableDeviceMaxLoginAttempts(v bool) {
	o.DisableDeviceMaxLoginAttempts = &v
}

// GetDisplayname returns the Displayname field value if set, zero value otherwise.
func (o *Systemuserputpost) GetDisplayname() string {
	if o == nil || IsNil(o.Displayname) {
		var ret string
		return ret
	}
	return *o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetDisplaynameOk() (*string, bool) {
	if o == nil || IsNil(o.Displayname) {
		return nil, false
	}
	return o.Displayname, true
}

// HasDisplayname returns a boolean if a field has been set.
func (o *Systemuserputpost) HasDisplayname() bool {
	if o != nil && !IsNil(o.Displayname) {
		return true
	}

	return false
}

// SetDisplayname gets a reference to the given string and assigns it to the Displayname field.
func (o *Systemuserputpost) SetDisplayname(v string) {
	o.Displayname = &v
}

// GetEmail returns the Email field value
func (o *Systemuserputpost) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Systemuserputpost) SetEmail(v string) {
	o.Email = v
}

// GetEmployeeIdentifier returns the EmployeeIdentifier field value if set, zero value otherwise.
func (o *Systemuserputpost) GetEmployeeIdentifier() string {
	if o == nil || IsNil(o.EmployeeIdentifier) {
		var ret string
		return ret
	}
	return *o.EmployeeIdentifier
}

// GetEmployeeIdentifierOk returns a tuple with the EmployeeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetEmployeeIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeIdentifier) {
		return nil, false
	}
	return o.EmployeeIdentifier, true
}

// HasEmployeeIdentifier returns a boolean if a field has been set.
func (o *Systemuserputpost) HasEmployeeIdentifier() bool {
	if o != nil && !IsNil(o.EmployeeIdentifier) {
		return true
	}

	return false
}

// SetEmployeeIdentifier gets a reference to the given string and assigns it to the EmployeeIdentifier field.
func (o *Systemuserputpost) SetEmployeeIdentifier(v string) {
	o.EmployeeIdentifier = &v
}

// GetEmployeeType returns the EmployeeType field value if set, zero value otherwise.
func (o *Systemuserputpost) GetEmployeeType() string {
	if o == nil || IsNil(o.EmployeeType) {
		var ret string
		return ret
	}
	return *o.EmployeeType
}

// GetEmployeeTypeOk returns a tuple with the EmployeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetEmployeeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeType) {
		return nil, false
	}
	return o.EmployeeType, true
}

// HasEmployeeType returns a boolean if a field has been set.
func (o *Systemuserputpost) HasEmployeeType() bool {
	if o != nil && !IsNil(o.EmployeeType) {
		return true
	}

	return false
}

// SetEmployeeType gets a reference to the given string and assigns it to the EmployeeType field.
func (o *Systemuserputpost) SetEmployeeType(v string) {
	o.EmployeeType = &v
}

// GetEnableManagedUid returns the EnableManagedUid field value if set, zero value otherwise.
func (o *Systemuserputpost) GetEnableManagedUid() bool {
	if o == nil || IsNil(o.EnableManagedUid) {
		var ret bool
		return ret
	}
	return *o.EnableManagedUid
}

// GetEnableManagedUidOk returns a tuple with the EnableManagedUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetEnableManagedUidOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableManagedUid) {
		return nil, false
	}
	return o.EnableManagedUid, true
}

// HasEnableManagedUid returns a boolean if a field has been set.
func (o *Systemuserputpost) HasEnableManagedUid() bool {
	if o != nil && !IsNil(o.EnableManagedUid) {
		return true
	}

	return false
}

// SetEnableManagedUid gets a reference to the given bool and assigns it to the EnableManagedUid field.
func (o *Systemuserputpost) SetEnableManagedUid(v bool) {
	o.EnableManagedUid = &v
}

// GetEnableUserPortalMultifactor returns the EnableUserPortalMultifactor field value if set, zero value otherwise.
func (o *Systemuserputpost) GetEnableUserPortalMultifactor() bool {
	if o == nil || IsNil(o.EnableUserPortalMultifactor) {
		var ret bool
		return ret
	}
	return *o.EnableUserPortalMultifactor
}

// GetEnableUserPortalMultifactorOk returns a tuple with the EnableUserPortalMultifactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetEnableUserPortalMultifactorOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableUserPortalMultifactor) {
		return nil, false
	}
	return o.EnableUserPortalMultifactor, true
}

// HasEnableUserPortalMultifactor returns a boolean if a field has been set.
func (o *Systemuserputpost) HasEnableUserPortalMultifactor() bool {
	if o != nil && !IsNil(o.EnableUserPortalMultifactor) {
		return true
	}

	return false
}

// SetEnableUserPortalMultifactor gets a reference to the given bool and assigns it to the EnableUserPortalMultifactor field.
func (o *Systemuserputpost) SetEnableUserPortalMultifactor(v bool) {
	o.EnableUserPortalMultifactor = &v
}

// GetExternalDn returns the ExternalDn field value if set, zero value otherwise.
func (o *Systemuserputpost) GetExternalDn() string {
	if o == nil || IsNil(o.ExternalDn) {
		var ret string
		return ret
	}
	return *o.ExternalDn
}

// GetExternalDnOk returns a tuple with the ExternalDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetExternalDnOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalDn) {
		return nil, false
	}
	return o.ExternalDn, true
}

// HasExternalDn returns a boolean if a field has been set.
func (o *Systemuserputpost) HasExternalDn() bool {
	if o != nil && !IsNil(o.ExternalDn) {
		return true
	}

	return false
}

// SetExternalDn gets a reference to the given string and assigns it to the ExternalDn field.
func (o *Systemuserputpost) SetExternalDn(v string) {
	o.ExternalDn = &v
}

// GetExternalPasswordExpirationDate returns the ExternalPasswordExpirationDate field value if set, zero value otherwise.
func (o *Systemuserputpost) GetExternalPasswordExpirationDate() time.Time {
	if o == nil || IsNil(o.ExternalPasswordExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExternalPasswordExpirationDate
}

// GetExternalPasswordExpirationDateOk returns a tuple with the ExternalPasswordExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetExternalPasswordExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExternalPasswordExpirationDate) {
		return nil, false
	}
	return o.ExternalPasswordExpirationDate, true
}

// HasExternalPasswordExpirationDate returns a boolean if a field has been set.
func (o *Systemuserputpost) HasExternalPasswordExpirationDate() bool {
	if o != nil && !IsNil(o.ExternalPasswordExpirationDate) {
		return true
	}

	return false
}

// SetExternalPasswordExpirationDate gets a reference to the given time.Time and assigns it to the ExternalPasswordExpirationDate field.
func (o *Systemuserputpost) SetExternalPasswordExpirationDate(v time.Time) {
	o.ExternalPasswordExpirationDate = &v
}

// GetExternalSourceType returns the ExternalSourceType field value if set, zero value otherwise.
func (o *Systemuserputpost) GetExternalSourceType() string {
	if o == nil || IsNil(o.ExternalSourceType) {
		var ret string
		return ret
	}
	return *o.ExternalSourceType
}

// GetExternalSourceTypeOk returns a tuple with the ExternalSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetExternalSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSourceType) {
		return nil, false
	}
	return o.ExternalSourceType, true
}

// HasExternalSourceType returns a boolean if a field has been set.
func (o *Systemuserputpost) HasExternalSourceType() bool {
	if o != nil && !IsNil(o.ExternalSourceType) {
		return true
	}

	return false
}

// SetExternalSourceType gets a reference to the given string and assigns it to the ExternalSourceType field.
func (o *Systemuserputpost) SetExternalSourceType(v string) {
	o.ExternalSourceType = &v
}

// GetExternallyManaged returns the ExternallyManaged field value if set, zero value otherwise.
func (o *Systemuserputpost) GetExternallyManaged() bool {
	if o == nil || IsNil(o.ExternallyManaged) {
		var ret bool
		return ret
	}
	return *o.ExternallyManaged
}

// GetExternallyManagedOk returns a tuple with the ExternallyManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetExternallyManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternallyManaged) {
		return nil, false
	}
	return o.ExternallyManaged, true
}

// HasExternallyManaged returns a boolean if a field has been set.
func (o *Systemuserputpost) HasExternallyManaged() bool {
	if o != nil && !IsNil(o.ExternallyManaged) {
		return true
	}

	return false
}

// SetExternallyManaged gets a reference to the given bool and assigns it to the ExternallyManaged field.
func (o *Systemuserputpost) SetExternallyManaged(v bool) {
	o.ExternallyManaged = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *Systemuserputpost) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *Systemuserputpost) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *Systemuserputpost) SetFirstname(v string) {
	o.Firstname = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *Systemuserputpost) GetJobTitle() string {
	if o == nil || IsNil(o.JobTitle) {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetJobTitleOk() (*string, bool) {
	if o == nil || IsNil(o.JobTitle) {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *Systemuserputpost) HasJobTitle() bool {
	if o != nil && !IsNil(o.JobTitle) {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *Systemuserputpost) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *Systemuserputpost) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *Systemuserputpost) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *Systemuserputpost) SetLastname(v string) {
	o.Lastname = &v
}

// GetLdapBindingUser returns the LdapBindingUser field value if set, zero value otherwise.
func (o *Systemuserputpost) GetLdapBindingUser() bool {
	if o == nil || IsNil(o.LdapBindingUser) {
		var ret bool
		return ret
	}
	return *o.LdapBindingUser
}

// GetLdapBindingUserOk returns a tuple with the LdapBindingUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetLdapBindingUserOk() (*bool, bool) {
	if o == nil || IsNil(o.LdapBindingUser) {
		return nil, false
	}
	return o.LdapBindingUser, true
}

// HasLdapBindingUser returns a boolean if a field has been set.
func (o *Systemuserputpost) HasLdapBindingUser() bool {
	if o != nil && !IsNil(o.LdapBindingUser) {
		return true
	}

	return false
}

// SetLdapBindingUser gets a reference to the given bool and assigns it to the LdapBindingUser field.
func (o *Systemuserputpost) SetLdapBindingUser(v bool) {
	o.LdapBindingUser = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Systemuserputpost) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Systemuserputpost) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Systemuserputpost) SetLocation(v string) {
	o.Location = &v
}

// GetManagedAppleId returns the ManagedAppleId field value if set, zero value otherwise.
func (o *Systemuserputpost) GetManagedAppleId() string {
	if o == nil || IsNil(o.ManagedAppleId) {
		var ret string
		return ret
	}
	return *o.ManagedAppleId
}

// GetManagedAppleIdOk returns a tuple with the ManagedAppleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetManagedAppleIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedAppleId) {
		return nil, false
	}
	return o.ManagedAppleId, true
}

// HasManagedAppleId returns a boolean if a field has been set.
func (o *Systemuserputpost) HasManagedAppleId() bool {
	if o != nil && !IsNil(o.ManagedAppleId) {
		return true
	}

	return false
}

// SetManagedAppleId gets a reference to the given string and assigns it to the ManagedAppleId field.
func (o *Systemuserputpost) SetManagedAppleId(v string) {
	o.ManagedAppleId = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *Systemuserputpost) GetManager() string {
	if o == nil || IsNil(o.Manager) {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetManagerOk() (*string, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *Systemuserputpost) HasManager() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given string and assigns it to the Manager field.
func (o *Systemuserputpost) SetManager(v string) {
	o.Manager = &v
}

// GetMfa returns the Mfa field value if set, zero value otherwise.
func (o *Systemuserputpost) GetMfa() Mfa {
	if o == nil || IsNil(o.Mfa) {
		var ret Mfa
		return ret
	}
	return *o.Mfa
}

// GetMfaOk returns a tuple with the Mfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetMfaOk() (*Mfa, bool) {
	if o == nil || IsNil(o.Mfa) {
		return nil, false
	}
	return o.Mfa, true
}

// HasMfa returns a boolean if a field has been set.
func (o *Systemuserputpost) HasMfa() bool {
	if o != nil && !IsNil(o.Mfa) {
		return true
	}

	return false
}

// SetMfa gets a reference to the given Mfa and assigns it to the Mfa field.
func (o *Systemuserputpost) SetMfa(v Mfa) {
	o.Mfa = &v
}

// GetMiddlename returns the Middlename field value if set, zero value otherwise.
func (o *Systemuserputpost) GetMiddlename() string {
	if o == nil || IsNil(o.Middlename) {
		var ret string
		return ret
	}
	return *o.Middlename
}

// GetMiddlenameOk returns a tuple with the Middlename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetMiddlenameOk() (*string, bool) {
	if o == nil || IsNil(o.Middlename) {
		return nil, false
	}
	return o.Middlename, true
}

// HasMiddlename returns a boolean if a field has been set.
func (o *Systemuserputpost) HasMiddlename() bool {
	if o != nil && !IsNil(o.Middlename) {
		return true
	}

	return false
}

// SetMiddlename gets a reference to the given string and assigns it to the Middlename field.
func (o *Systemuserputpost) SetMiddlename(v string) {
	o.Middlename = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Systemuserputpost) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Systemuserputpost) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Systemuserputpost) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordNeverExpires returns the PasswordNeverExpires field value if set, zero value otherwise.
func (o *Systemuserputpost) GetPasswordNeverExpires() bool {
	if o == nil || IsNil(o.PasswordNeverExpires) {
		var ret bool
		return ret
	}
	return *o.PasswordNeverExpires
}

// GetPasswordNeverExpiresOk returns a tuple with the PasswordNeverExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetPasswordNeverExpiresOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordNeverExpires) {
		return nil, false
	}
	return o.PasswordNeverExpires, true
}

// HasPasswordNeverExpires returns a boolean if a field has been set.
func (o *Systemuserputpost) HasPasswordNeverExpires() bool {
	if o != nil && !IsNil(o.PasswordNeverExpires) {
		return true
	}

	return false
}

// SetPasswordNeverExpires gets a reference to the given bool and assigns it to the PasswordNeverExpires field.
func (o *Systemuserputpost) SetPasswordNeverExpires(v bool) {
	o.PasswordNeverExpires = &v
}

// GetPasswordlessSudo returns the PasswordlessSudo field value if set, zero value otherwise.
func (o *Systemuserputpost) GetPasswordlessSudo() bool {
	if o == nil || IsNil(o.PasswordlessSudo) {
		var ret bool
		return ret
	}
	return *o.PasswordlessSudo
}

// GetPasswordlessSudoOk returns a tuple with the PasswordlessSudo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetPasswordlessSudoOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordlessSudo) {
		return nil, false
	}
	return o.PasswordlessSudo, true
}

// HasPasswordlessSudo returns a boolean if a field has been set.
func (o *Systemuserputpost) HasPasswordlessSudo() bool {
	if o != nil && !IsNil(o.PasswordlessSudo) {
		return true
	}

	return false
}

// SetPasswordlessSudo gets a reference to the given bool and assigns it to the PasswordlessSudo field.
func (o *Systemuserputpost) SetPasswordlessSudo(v bool) {
	o.PasswordlessSudo = &v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *Systemuserputpost) GetPhoneNumbers() []SystemuserputpostPhoneNumbersInner {
	if o == nil || IsNil(o.PhoneNumbers) {
		var ret []SystemuserputpostPhoneNumbersInner
		return ret
	}
	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetPhoneNumbersOk() ([]SystemuserputpostPhoneNumbersInner, bool) {
	if o == nil || IsNil(o.PhoneNumbers) {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *Systemuserputpost) HasPhoneNumbers() bool {
	if o != nil && !IsNil(o.PhoneNumbers) {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []SystemuserputpostPhoneNumbersInner and assigns it to the PhoneNumbers field.
func (o *Systemuserputpost) SetPhoneNumbers(v []SystemuserputpostPhoneNumbersInner) {
	o.PhoneNumbers = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *Systemuserputpost) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *Systemuserputpost) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *Systemuserputpost) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRecoveryEmail returns the RecoveryEmail field value if set, zero value otherwise.
func (o *Systemuserputpost) GetRecoveryEmail() SystemuserputpostRecoveryEmail {
	if o == nil || IsNil(o.RecoveryEmail) {
		var ret SystemuserputpostRecoveryEmail
		return ret
	}
	return *o.RecoveryEmail
}

// GetRecoveryEmailOk returns a tuple with the RecoveryEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetRecoveryEmailOk() (*SystemuserputpostRecoveryEmail, bool) {
	if o == nil || IsNil(o.RecoveryEmail) {
		return nil, false
	}
	return o.RecoveryEmail, true
}

// HasRecoveryEmail returns a boolean if a field has been set.
func (o *Systemuserputpost) HasRecoveryEmail() bool {
	if o != nil && !IsNil(o.RecoveryEmail) {
		return true
	}

	return false
}

// SetRecoveryEmail gets a reference to the given SystemuserputpostRecoveryEmail and assigns it to the RecoveryEmail field.
func (o *Systemuserputpost) SetRecoveryEmail(v SystemuserputpostRecoveryEmail) {
	o.RecoveryEmail = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Systemuserputpost) GetRelationships() []SystemuserputRelationshipsInner {
	if o == nil || IsNil(o.Relationships) {
		var ret []SystemuserputRelationshipsInner
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetRelationshipsOk() ([]SystemuserputRelationshipsInner, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Systemuserputpost) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []SystemuserputRelationshipsInner and assigns it to the Relationships field.
func (o *Systemuserputpost) SetRelationships(v []SystemuserputRelationshipsInner) {
	o.Relationships = v
}

// GetSambaServiceUser returns the SambaServiceUser field value if set, zero value otherwise.
func (o *Systemuserputpost) GetSambaServiceUser() bool {
	if o == nil || IsNil(o.SambaServiceUser) {
		var ret bool
		return ret
	}
	return *o.SambaServiceUser
}

// GetSambaServiceUserOk returns a tuple with the SambaServiceUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetSambaServiceUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SambaServiceUser) {
		return nil, false
	}
	return o.SambaServiceUser, true
}

// HasSambaServiceUser returns a boolean if a field has been set.
func (o *Systemuserputpost) HasSambaServiceUser() bool {
	if o != nil && !IsNil(o.SambaServiceUser) {
		return true
	}

	return false
}

// SetSambaServiceUser gets a reference to the given bool and assigns it to the SambaServiceUser field.
func (o *Systemuserputpost) SetSambaServiceUser(v bool) {
	o.SambaServiceUser = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Systemuserputpost) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Systemuserputpost) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Systemuserputpost) SetState(v string) {
	o.State = &v
}

// GetSudo returns the Sudo field value if set, zero value otherwise.
func (o *Systemuserputpost) GetSudo() bool {
	if o == nil || IsNil(o.Sudo) {
		var ret bool
		return ret
	}
	return *o.Sudo
}

// GetSudoOk returns a tuple with the Sudo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetSudoOk() (*bool, bool) {
	if o == nil || IsNil(o.Sudo) {
		return nil, false
	}
	return o.Sudo, true
}

// HasSudo returns a boolean if a field has been set.
func (o *Systemuserputpost) HasSudo() bool {
	if o != nil && !IsNil(o.Sudo) {
		return true
	}

	return false
}

// SetSudo gets a reference to the given bool and assigns it to the Sudo field.
func (o *Systemuserputpost) SetSudo(v bool) {
	o.Sudo = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *Systemuserputpost) GetSuspended() bool {
	if o == nil || IsNil(o.Suspended) {
		var ret bool
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetSuspendedOk() (*bool, bool) {
	if o == nil || IsNil(o.Suspended) {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *Systemuserputpost) HasSuspended() bool {
	if o != nil && !IsNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given bool and assigns it to the Suspended field.
func (o *Systemuserputpost) SetSuspended(v bool) {
	o.Suspended = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Systemuserputpost) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Systemuserputpost) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Systemuserputpost) SetTags(v []string) {
	o.Tags = v
}

// GetUnixGuid returns the UnixGuid field value if set, zero value otherwise.
func (o *Systemuserputpost) GetUnixGuid() int32 {
	if o == nil || IsNil(o.UnixGuid) {
		var ret int32
		return ret
	}
	return *o.UnixGuid
}

// GetUnixGuidOk returns a tuple with the UnixGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetUnixGuidOk() (*int32, bool) {
	if o == nil || IsNil(o.UnixGuid) {
		return nil, false
	}
	return o.UnixGuid, true
}

// HasUnixGuid returns a boolean if a field has been set.
func (o *Systemuserputpost) HasUnixGuid() bool {
	if o != nil && !IsNil(o.UnixGuid) {
		return true
	}

	return false
}

// SetUnixGuid gets a reference to the given int32 and assigns it to the UnixGuid field.
func (o *Systemuserputpost) SetUnixGuid(v int32) {
	o.UnixGuid = &v
}

// GetUnixUid returns the UnixUid field value if set, zero value otherwise.
func (o *Systemuserputpost) GetUnixUid() int32 {
	if o == nil || IsNil(o.UnixUid) {
		var ret int32
		return ret
	}
	return *o.UnixUid
}

// GetUnixUidOk returns a tuple with the UnixUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetUnixUidOk() (*int32, bool) {
	if o == nil || IsNil(o.UnixUid) {
		return nil, false
	}
	return o.UnixUid, true
}

// HasUnixUid returns a boolean if a field has been set.
func (o *Systemuserputpost) HasUnixUid() bool {
	if o != nil && !IsNil(o.UnixUid) {
		return true
	}

	return false
}

// SetUnixUid gets a reference to the given int32 and assigns it to the UnixUid field.
func (o *Systemuserputpost) SetUnixUid(v int32) {
	o.UnixUid = &v
}

// GetUsername returns the Username field value
func (o *Systemuserputpost) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Systemuserputpost) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Systemuserputpost) SetUsername(v string) {
	o.Username = v
}

func (o Systemuserputpost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Systemuserputpost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountLocked) {
		toSerialize["account_locked"] = o.AccountLocked
	}
	if !IsNil(o.Activated) {
		toSerialize["activated"] = o.Activated
	}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.AllowPublicKey) {
		toSerialize["allow_public_key"] = o.AllowPublicKey
	}
	if !IsNil(o.AlternateEmail) {
		toSerialize["alternateEmail"] = o.AlternateEmail
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisableDeviceMaxLoginAttempts) {
		toSerialize["disableDeviceMaxLoginAttempts"] = o.DisableDeviceMaxLoginAttempts
	}
	if !IsNil(o.Displayname) {
		toSerialize["displayname"] = o.Displayname
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.EmployeeIdentifier) {
		toSerialize["employeeIdentifier"] = o.EmployeeIdentifier
	}
	if !IsNil(o.EmployeeType) {
		toSerialize["employeeType"] = o.EmployeeType
	}
	if !IsNil(o.EnableManagedUid) {
		toSerialize["enable_managed_uid"] = o.EnableManagedUid
	}
	if !IsNil(o.EnableUserPortalMultifactor) {
		toSerialize["enable_user_portal_multifactor"] = o.EnableUserPortalMultifactor
	}
	if !IsNil(o.ExternalDn) {
		toSerialize["external_dn"] = o.ExternalDn
	}
	if !IsNil(o.ExternalPasswordExpirationDate) {
		toSerialize["external_password_expiration_date"] = o.ExternalPasswordExpirationDate
	}
	if !IsNil(o.ExternalSourceType) {
		toSerialize["external_source_type"] = o.ExternalSourceType
	}
	if !IsNil(o.ExternallyManaged) {
		toSerialize["externally_managed"] = o.ExternallyManaged
	}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.JobTitle) {
		toSerialize["jobTitle"] = o.JobTitle
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.LdapBindingUser) {
		toSerialize["ldap_binding_user"] = o.LdapBindingUser
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.ManagedAppleId) {
		toSerialize["managedAppleId"] = o.ManagedAppleId
	}
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if !IsNil(o.Mfa) {
		toSerialize["mfa"] = o.Mfa
	}
	if !IsNil(o.Middlename) {
		toSerialize["middlename"] = o.Middlename
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PasswordNeverExpires) {
		toSerialize["password_never_expires"] = o.PasswordNeverExpires
	}
	if !IsNil(o.PasswordlessSudo) {
		toSerialize["passwordless_sudo"] = o.PasswordlessSudo
	}
	if !IsNil(o.PhoneNumbers) {
		toSerialize["phoneNumbers"] = o.PhoneNumbers
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if !IsNil(o.RecoveryEmail) {
		toSerialize["recoveryEmail"] = o.RecoveryEmail
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.SambaServiceUser) {
		toSerialize["samba_service_user"] = o.SambaServiceUser
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Sudo) {
		toSerialize["sudo"] = o.Sudo
	}
	if !IsNil(o.Suspended) {
		toSerialize["suspended"] = o.Suspended
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UnixGuid) {
		toSerialize["unix_guid"] = o.UnixGuid
	}
	if !IsNil(o.UnixUid) {
		toSerialize["unix_uid"] = o.UnixUid
	}
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Systemuserputpost) UnmarshalJSON(bytes []byte) (err error) {
	varSystemuserputpost := _Systemuserputpost{}

	if err = json.Unmarshal(bytes, &varSystemuserputpost); err == nil {
		*o = Systemuserputpost(varSystemuserputpost)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_locked")
		delete(additionalProperties, "activated")
		delete(additionalProperties, "addresses")
		delete(additionalProperties, "allow_public_key")
		delete(additionalProperties, "alternateEmail")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "company")
		delete(additionalProperties, "costCenter")
		delete(additionalProperties, "department")
		delete(additionalProperties, "description")
		delete(additionalProperties, "disableDeviceMaxLoginAttempts")
		delete(additionalProperties, "displayname")
		delete(additionalProperties, "email")
		delete(additionalProperties, "employeeIdentifier")
		delete(additionalProperties, "employeeType")
		delete(additionalProperties, "enable_managed_uid")
		delete(additionalProperties, "enable_user_portal_multifactor")
		delete(additionalProperties, "external_dn")
		delete(additionalProperties, "external_password_expiration_date")
		delete(additionalProperties, "external_source_type")
		delete(additionalProperties, "externally_managed")
		delete(additionalProperties, "firstname")
		delete(additionalProperties, "jobTitle")
		delete(additionalProperties, "lastname")
		delete(additionalProperties, "ldap_binding_user")
		delete(additionalProperties, "location")
		delete(additionalProperties, "managedAppleId")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "mfa")
		delete(additionalProperties, "middlename")
		delete(additionalProperties, "password")
		delete(additionalProperties, "password_never_expires")
		delete(additionalProperties, "passwordless_sudo")
		delete(additionalProperties, "phoneNumbers")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "recoveryEmail")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "samba_service_user")
		delete(additionalProperties, "state")
		delete(additionalProperties, "sudo")
		delete(additionalProperties, "suspended")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "unix_guid")
		delete(additionalProperties, "unix_uid")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemuserputpost struct {
	value *Systemuserputpost
	isSet bool
}

func (v NullableSystemuserputpost) Get() *Systemuserputpost {
	return v.value
}

func (v *NullableSystemuserputpost) Set(val *Systemuserputpost) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemuserputpost) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemuserputpost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemuserputpost(val *Systemuserputpost) *NullableSystemuserputpost {
	return &NullableSystemuserputpost{value: val, isSet: true}
}

func (v NullableSystemuserputpost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemuserputpost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



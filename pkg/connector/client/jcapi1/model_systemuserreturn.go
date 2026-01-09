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

// checks if the Systemuserreturn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Systemuserreturn{}

// Systemuserreturn struct for Systemuserreturn
type Systemuserreturn struct {
	Id *string `json:"_id,omitempty"`
	AccountLocked *bool `json:"account_locked,omitempty"`
	AccountLockedDate NullableString `json:"account_locked_date,omitempty"`
	Activated *bool `json:"activated,omitempty"`
	Addresses []SystemuserreturnAddressesInner `json:"addresses,omitempty"`
	AllowPublicKey *bool `json:"allow_public_key,omitempty"`
	AlternateEmail *string `json:"alternateEmail,omitempty"`
	Attributes []SystemuserputAttributesInner `json:"attributes,omitempty"`
	BadLoginAttempts *int32 `json:"badLoginAttempts,omitempty"`
	Company *string `json:"company,omitempty"`
	CostCenter *string `json:"costCenter,omitempty"`
	Created *string `json:"created,omitempty"`
	CreationSource *string `json:"creationSource,omitempty"`
	Department *string `json:"department,omitempty"`
	Description *string `json:"description,omitempty"`
	DisableDeviceMaxLoginAttempts *bool `json:"disableDeviceMaxLoginAttempts,omitempty"`
	Displayname *string `json:"displayname,omitempty"`
	Email *string `json:"email,omitempty"`
	// Must be unique per user. 
	EmployeeIdentifier *string `json:"employeeIdentifier,omitempty"`
	EmployeeType *string `json:"employeeType,omitempty"`
	EnableManagedUid *bool `json:"enable_managed_uid,omitempty"`
	EnableUserPortalMultifactor *bool `json:"enable_user_portal_multifactor,omitempty"`
	ExternalDn *string `json:"external_dn,omitempty"`
	ExternalPasswordExpirationDate *string `json:"external_password_expiration_date,omitempty"`
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
	MfaEnrollment *MfaEnrollment `json:"mfaEnrollment,omitempty"`
	Middlename *string `json:"middlename,omitempty"`
	Organization *string `json:"organization,omitempty"`
	PasswordDate NullableString `json:"password_date,omitempty"`
	PasswordExpirationDate NullableString `json:"password_expiration_date,omitempty"`
	PasswordExpired *bool `json:"password_expired,omitempty"`
	PasswordNeverExpires *bool `json:"password_never_expires,omitempty"`
	PasswordlessSudo *bool `json:"passwordless_sudo,omitempty"`
	PhoneNumbers []SystemuserreturnPhoneNumbersInner `json:"phoneNumbers,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
	RecoveryEmail *SystemuserreturnRecoveryEmail `json:"recoveryEmail,omitempty"`
	Relationships []SystemuserputRelationshipsInner `json:"relationships,omitempty"`
	SambaServiceUser *bool `json:"samba_service_user,omitempty"`
	SshKeys []Sshkeylist `json:"ssh_keys,omitempty"`
	State *string `json:"state,omitempty"`
	Sudo *bool `json:"sudo,omitempty"`
	Suspended *bool `json:"suspended,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TotpEnabled *bool `json:"totp_enabled,omitempty"`
	UnixGuid *int32 `json:"unix_guid,omitempty"`
	UnixUid *int32 `json:"unix_uid,omitempty"`
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Systemuserreturn Systemuserreturn

// NewSystemuserreturn instantiates a new Systemuserreturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemuserreturn() *Systemuserreturn {
	this := Systemuserreturn{}
	return &this
}

// NewSystemuserreturnWithDefaults instantiates a new Systemuserreturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemuserreturnWithDefaults() *Systemuserreturn {
	this := Systemuserreturn{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Systemuserreturn) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Systemuserreturn) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Systemuserreturn) SetId(v string) {
	o.Id = &v
}

// GetAccountLocked returns the AccountLocked field value if set, zero value otherwise.
func (o *Systemuserreturn) GetAccountLocked() bool {
	if o == nil || IsNil(o.AccountLocked) {
		var ret bool
		return ret
	}
	return *o.AccountLocked
}

// GetAccountLockedOk returns a tuple with the AccountLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetAccountLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccountLocked) {
		return nil, false
	}
	return o.AccountLocked, true
}

// HasAccountLocked returns a boolean if a field has been set.
func (o *Systemuserreturn) HasAccountLocked() bool {
	if o != nil && !IsNil(o.AccountLocked) {
		return true
	}

	return false
}

// SetAccountLocked gets a reference to the given bool and assigns it to the AccountLocked field.
func (o *Systemuserreturn) SetAccountLocked(v bool) {
	o.AccountLocked = &v
}

// GetAccountLockedDate returns the AccountLockedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Systemuserreturn) GetAccountLockedDate() string {
	if o == nil || IsNil(o.AccountLockedDate.Get()) {
		var ret string
		return ret
	}
	return *o.AccountLockedDate.Get()
}

// GetAccountLockedDateOk returns a tuple with the AccountLockedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Systemuserreturn) GetAccountLockedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountLockedDate.Get(), o.AccountLockedDate.IsSet()
}

// HasAccountLockedDate returns a boolean if a field has been set.
func (o *Systemuserreturn) HasAccountLockedDate() bool {
	if o != nil && o.AccountLockedDate.IsSet() {
		return true
	}

	return false
}

// SetAccountLockedDate gets a reference to the given NullableString and assigns it to the AccountLockedDate field.
func (o *Systemuserreturn) SetAccountLockedDate(v string) {
	o.AccountLockedDate.Set(&v)
}
// SetAccountLockedDateNil sets the value for AccountLockedDate to be an explicit nil
func (o *Systemuserreturn) SetAccountLockedDateNil() {
	o.AccountLockedDate.Set(nil)
}

// UnsetAccountLockedDate ensures that no value is present for AccountLockedDate, not even an explicit nil
func (o *Systemuserreturn) UnsetAccountLockedDate() {
	o.AccountLockedDate.Unset()
}

// GetActivated returns the Activated field value if set, zero value otherwise.
func (o *Systemuserreturn) GetActivated() bool {
	if o == nil || IsNil(o.Activated) {
		var ret bool
		return ret
	}
	return *o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetActivatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Activated) {
		return nil, false
	}
	return o.Activated, true
}

// HasActivated returns a boolean if a field has been set.
func (o *Systemuserreturn) HasActivated() bool {
	if o != nil && !IsNil(o.Activated) {
		return true
	}

	return false
}

// SetActivated gets a reference to the given bool and assigns it to the Activated field.
func (o *Systemuserreturn) SetActivated(v bool) {
	o.Activated = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *Systemuserreturn) GetAddresses() []SystemuserreturnAddressesInner {
	if o == nil || IsNil(o.Addresses) {
		var ret []SystemuserreturnAddressesInner
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetAddressesOk() ([]SystemuserreturnAddressesInner, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *Systemuserreturn) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []SystemuserreturnAddressesInner and assigns it to the Addresses field.
func (o *Systemuserreturn) SetAddresses(v []SystemuserreturnAddressesInner) {
	o.Addresses = v
}

// GetAllowPublicKey returns the AllowPublicKey field value if set, zero value otherwise.
func (o *Systemuserreturn) GetAllowPublicKey() bool {
	if o == nil || IsNil(o.AllowPublicKey) {
		var ret bool
		return ret
	}
	return *o.AllowPublicKey
}

// GetAllowPublicKeyOk returns a tuple with the AllowPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetAllowPublicKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPublicKey) {
		return nil, false
	}
	return o.AllowPublicKey, true
}

// HasAllowPublicKey returns a boolean if a field has been set.
func (o *Systemuserreturn) HasAllowPublicKey() bool {
	if o != nil && !IsNil(o.AllowPublicKey) {
		return true
	}

	return false
}

// SetAllowPublicKey gets a reference to the given bool and assigns it to the AllowPublicKey field.
func (o *Systemuserreturn) SetAllowPublicKey(v bool) {
	o.AllowPublicKey = &v
}

// GetAlternateEmail returns the AlternateEmail field value if set, zero value otherwise.
func (o *Systemuserreturn) GetAlternateEmail() string {
	if o == nil || IsNil(o.AlternateEmail) {
		var ret string
		return ret
	}
	return *o.AlternateEmail
}

// GetAlternateEmailOk returns a tuple with the AlternateEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetAlternateEmailOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateEmail) {
		return nil, false
	}
	return o.AlternateEmail, true
}

// HasAlternateEmail returns a boolean if a field has been set.
func (o *Systemuserreturn) HasAlternateEmail() bool {
	if o != nil && !IsNil(o.AlternateEmail) {
		return true
	}

	return false
}

// SetAlternateEmail gets a reference to the given string and assigns it to the AlternateEmail field.
func (o *Systemuserreturn) SetAlternateEmail(v string) {
	o.AlternateEmail = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Systemuserreturn) GetAttributes() []SystemuserputAttributesInner {
	if o == nil || IsNil(o.Attributes) {
		var ret []SystemuserputAttributesInner
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetAttributesOk() ([]SystemuserputAttributesInner, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Systemuserreturn) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []SystemuserputAttributesInner and assigns it to the Attributes field.
func (o *Systemuserreturn) SetAttributes(v []SystemuserputAttributesInner) {
	o.Attributes = v
}

// GetBadLoginAttempts returns the BadLoginAttempts field value if set, zero value otherwise.
func (o *Systemuserreturn) GetBadLoginAttempts() int32 {
	if o == nil || IsNil(o.BadLoginAttempts) {
		var ret int32
		return ret
	}
	return *o.BadLoginAttempts
}

// GetBadLoginAttemptsOk returns a tuple with the BadLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetBadLoginAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.BadLoginAttempts) {
		return nil, false
	}
	return o.BadLoginAttempts, true
}

// HasBadLoginAttempts returns a boolean if a field has been set.
func (o *Systemuserreturn) HasBadLoginAttempts() bool {
	if o != nil && !IsNil(o.BadLoginAttempts) {
		return true
	}

	return false
}

// SetBadLoginAttempts gets a reference to the given int32 and assigns it to the BadLoginAttempts field.
func (o *Systemuserreturn) SetBadLoginAttempts(v int32) {
	o.BadLoginAttempts = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *Systemuserreturn) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *Systemuserreturn) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *Systemuserreturn) SetCompany(v string) {
	o.Company = &v
}

// GetCostCenter returns the CostCenter field value if set, zero value otherwise.
func (o *Systemuserreturn) GetCostCenter() string {
	if o == nil || IsNil(o.CostCenter) {
		var ret string
		return ret
	}
	return *o.CostCenter
}

// GetCostCenterOk returns a tuple with the CostCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetCostCenterOk() (*string, bool) {
	if o == nil || IsNil(o.CostCenter) {
		return nil, false
	}
	return o.CostCenter, true
}

// HasCostCenter returns a boolean if a field has been set.
func (o *Systemuserreturn) HasCostCenter() bool {
	if o != nil && !IsNil(o.CostCenter) {
		return true
	}

	return false
}

// SetCostCenter gets a reference to the given string and assigns it to the CostCenter field.
func (o *Systemuserreturn) SetCostCenter(v string) {
	o.CostCenter = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Systemuserreturn) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Systemuserreturn) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Systemuserreturn) SetCreated(v string) {
	o.Created = &v
}

// GetCreationSource returns the CreationSource field value if set, zero value otherwise.
func (o *Systemuserreturn) GetCreationSource() string {
	if o == nil || IsNil(o.CreationSource) {
		var ret string
		return ret
	}
	return *o.CreationSource
}

// GetCreationSourceOk returns a tuple with the CreationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetCreationSourceOk() (*string, bool) {
	if o == nil || IsNil(o.CreationSource) {
		return nil, false
	}
	return o.CreationSource, true
}

// HasCreationSource returns a boolean if a field has been set.
func (o *Systemuserreturn) HasCreationSource() bool {
	if o != nil && !IsNil(o.CreationSource) {
		return true
	}

	return false
}

// SetCreationSource gets a reference to the given string and assigns it to the CreationSource field.
func (o *Systemuserreturn) SetCreationSource(v string) {
	o.CreationSource = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *Systemuserreturn) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *Systemuserreturn) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *Systemuserreturn) SetDepartment(v string) {
	o.Department = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Systemuserreturn) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Systemuserreturn) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Systemuserreturn) SetDescription(v string) {
	o.Description = &v
}

// GetDisableDeviceMaxLoginAttempts returns the DisableDeviceMaxLoginAttempts field value if set, zero value otherwise.
func (o *Systemuserreturn) GetDisableDeviceMaxLoginAttempts() bool {
	if o == nil || IsNil(o.DisableDeviceMaxLoginAttempts) {
		var ret bool
		return ret
	}
	return *o.DisableDeviceMaxLoginAttempts
}

// GetDisableDeviceMaxLoginAttemptsOk returns a tuple with the DisableDeviceMaxLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetDisableDeviceMaxLoginAttemptsOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableDeviceMaxLoginAttempts) {
		return nil, false
	}
	return o.DisableDeviceMaxLoginAttempts, true
}

// HasDisableDeviceMaxLoginAttempts returns a boolean if a field has been set.
func (o *Systemuserreturn) HasDisableDeviceMaxLoginAttempts() bool {
	if o != nil && !IsNil(o.DisableDeviceMaxLoginAttempts) {
		return true
	}

	return false
}

// SetDisableDeviceMaxLoginAttempts gets a reference to the given bool and assigns it to the DisableDeviceMaxLoginAttempts field.
func (o *Systemuserreturn) SetDisableDeviceMaxLoginAttempts(v bool) {
	o.DisableDeviceMaxLoginAttempts = &v
}

// GetDisplayname returns the Displayname field value if set, zero value otherwise.
func (o *Systemuserreturn) GetDisplayname() string {
	if o == nil || IsNil(o.Displayname) {
		var ret string
		return ret
	}
	return *o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetDisplaynameOk() (*string, bool) {
	if o == nil || IsNil(o.Displayname) {
		return nil, false
	}
	return o.Displayname, true
}

// HasDisplayname returns a boolean if a field has been set.
func (o *Systemuserreturn) HasDisplayname() bool {
	if o != nil && !IsNil(o.Displayname) {
		return true
	}

	return false
}

// SetDisplayname gets a reference to the given string and assigns it to the Displayname field.
func (o *Systemuserreturn) SetDisplayname(v string) {
	o.Displayname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Systemuserreturn) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Systemuserreturn) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Systemuserreturn) SetEmail(v string) {
	o.Email = &v
}

// GetEmployeeIdentifier returns the EmployeeIdentifier field value if set, zero value otherwise.
func (o *Systemuserreturn) GetEmployeeIdentifier() string {
	if o == nil || IsNil(o.EmployeeIdentifier) {
		var ret string
		return ret
	}
	return *o.EmployeeIdentifier
}

// GetEmployeeIdentifierOk returns a tuple with the EmployeeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetEmployeeIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeIdentifier) {
		return nil, false
	}
	return o.EmployeeIdentifier, true
}

// HasEmployeeIdentifier returns a boolean if a field has been set.
func (o *Systemuserreturn) HasEmployeeIdentifier() bool {
	if o != nil && !IsNil(o.EmployeeIdentifier) {
		return true
	}

	return false
}

// SetEmployeeIdentifier gets a reference to the given string and assigns it to the EmployeeIdentifier field.
func (o *Systemuserreturn) SetEmployeeIdentifier(v string) {
	o.EmployeeIdentifier = &v
}

// GetEmployeeType returns the EmployeeType field value if set, zero value otherwise.
func (o *Systemuserreturn) GetEmployeeType() string {
	if o == nil || IsNil(o.EmployeeType) {
		var ret string
		return ret
	}
	return *o.EmployeeType
}

// GetEmployeeTypeOk returns a tuple with the EmployeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetEmployeeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeType) {
		return nil, false
	}
	return o.EmployeeType, true
}

// HasEmployeeType returns a boolean if a field has been set.
func (o *Systemuserreturn) HasEmployeeType() bool {
	if o != nil && !IsNil(o.EmployeeType) {
		return true
	}

	return false
}

// SetEmployeeType gets a reference to the given string and assigns it to the EmployeeType field.
func (o *Systemuserreturn) SetEmployeeType(v string) {
	o.EmployeeType = &v
}

// GetEnableManagedUid returns the EnableManagedUid field value if set, zero value otherwise.
func (o *Systemuserreturn) GetEnableManagedUid() bool {
	if o == nil || IsNil(o.EnableManagedUid) {
		var ret bool
		return ret
	}
	return *o.EnableManagedUid
}

// GetEnableManagedUidOk returns a tuple with the EnableManagedUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetEnableManagedUidOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableManagedUid) {
		return nil, false
	}
	return o.EnableManagedUid, true
}

// HasEnableManagedUid returns a boolean if a field has been set.
func (o *Systemuserreturn) HasEnableManagedUid() bool {
	if o != nil && !IsNil(o.EnableManagedUid) {
		return true
	}

	return false
}

// SetEnableManagedUid gets a reference to the given bool and assigns it to the EnableManagedUid field.
func (o *Systemuserreturn) SetEnableManagedUid(v bool) {
	o.EnableManagedUid = &v
}

// GetEnableUserPortalMultifactor returns the EnableUserPortalMultifactor field value if set, zero value otherwise.
func (o *Systemuserreturn) GetEnableUserPortalMultifactor() bool {
	if o == nil || IsNil(o.EnableUserPortalMultifactor) {
		var ret bool
		return ret
	}
	return *o.EnableUserPortalMultifactor
}

// GetEnableUserPortalMultifactorOk returns a tuple with the EnableUserPortalMultifactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetEnableUserPortalMultifactorOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableUserPortalMultifactor) {
		return nil, false
	}
	return o.EnableUserPortalMultifactor, true
}

// HasEnableUserPortalMultifactor returns a boolean if a field has been set.
func (o *Systemuserreturn) HasEnableUserPortalMultifactor() bool {
	if o != nil && !IsNil(o.EnableUserPortalMultifactor) {
		return true
	}

	return false
}

// SetEnableUserPortalMultifactor gets a reference to the given bool and assigns it to the EnableUserPortalMultifactor field.
func (o *Systemuserreturn) SetEnableUserPortalMultifactor(v bool) {
	o.EnableUserPortalMultifactor = &v
}

// GetExternalDn returns the ExternalDn field value if set, zero value otherwise.
func (o *Systemuserreturn) GetExternalDn() string {
	if o == nil || IsNil(o.ExternalDn) {
		var ret string
		return ret
	}
	return *o.ExternalDn
}

// GetExternalDnOk returns a tuple with the ExternalDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetExternalDnOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalDn) {
		return nil, false
	}
	return o.ExternalDn, true
}

// HasExternalDn returns a boolean if a field has been set.
func (o *Systemuserreturn) HasExternalDn() bool {
	if o != nil && !IsNil(o.ExternalDn) {
		return true
	}

	return false
}

// SetExternalDn gets a reference to the given string and assigns it to the ExternalDn field.
func (o *Systemuserreturn) SetExternalDn(v string) {
	o.ExternalDn = &v
}

// GetExternalPasswordExpirationDate returns the ExternalPasswordExpirationDate field value if set, zero value otherwise.
func (o *Systemuserreturn) GetExternalPasswordExpirationDate() string {
	if o == nil || IsNil(o.ExternalPasswordExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExternalPasswordExpirationDate
}

// GetExternalPasswordExpirationDateOk returns a tuple with the ExternalPasswordExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetExternalPasswordExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPasswordExpirationDate) {
		return nil, false
	}
	return o.ExternalPasswordExpirationDate, true
}

// HasExternalPasswordExpirationDate returns a boolean if a field has been set.
func (o *Systemuserreturn) HasExternalPasswordExpirationDate() bool {
	if o != nil && !IsNil(o.ExternalPasswordExpirationDate) {
		return true
	}

	return false
}

// SetExternalPasswordExpirationDate gets a reference to the given string and assigns it to the ExternalPasswordExpirationDate field.
func (o *Systemuserreturn) SetExternalPasswordExpirationDate(v string) {
	o.ExternalPasswordExpirationDate = &v
}

// GetExternalSourceType returns the ExternalSourceType field value if set, zero value otherwise.
func (o *Systemuserreturn) GetExternalSourceType() string {
	if o == nil || IsNil(o.ExternalSourceType) {
		var ret string
		return ret
	}
	return *o.ExternalSourceType
}

// GetExternalSourceTypeOk returns a tuple with the ExternalSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetExternalSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSourceType) {
		return nil, false
	}
	return o.ExternalSourceType, true
}

// HasExternalSourceType returns a boolean if a field has been set.
func (o *Systemuserreturn) HasExternalSourceType() bool {
	if o != nil && !IsNil(o.ExternalSourceType) {
		return true
	}

	return false
}

// SetExternalSourceType gets a reference to the given string and assigns it to the ExternalSourceType field.
func (o *Systemuserreturn) SetExternalSourceType(v string) {
	o.ExternalSourceType = &v
}

// GetExternallyManaged returns the ExternallyManaged field value if set, zero value otherwise.
func (o *Systemuserreturn) GetExternallyManaged() bool {
	if o == nil || IsNil(o.ExternallyManaged) {
		var ret bool
		return ret
	}
	return *o.ExternallyManaged
}

// GetExternallyManagedOk returns a tuple with the ExternallyManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetExternallyManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternallyManaged) {
		return nil, false
	}
	return o.ExternallyManaged, true
}

// HasExternallyManaged returns a boolean if a field has been set.
func (o *Systemuserreturn) HasExternallyManaged() bool {
	if o != nil && !IsNil(o.ExternallyManaged) {
		return true
	}

	return false
}

// SetExternallyManaged gets a reference to the given bool and assigns it to the ExternallyManaged field.
func (o *Systemuserreturn) SetExternallyManaged(v bool) {
	o.ExternallyManaged = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *Systemuserreturn) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *Systemuserreturn) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *Systemuserreturn) SetFirstname(v string) {
	o.Firstname = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *Systemuserreturn) GetJobTitle() string {
	if o == nil || IsNil(o.JobTitle) {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetJobTitleOk() (*string, bool) {
	if o == nil || IsNil(o.JobTitle) {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *Systemuserreturn) HasJobTitle() bool {
	if o != nil && !IsNil(o.JobTitle) {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *Systemuserreturn) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *Systemuserreturn) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *Systemuserreturn) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *Systemuserreturn) SetLastname(v string) {
	o.Lastname = &v
}

// GetLdapBindingUser returns the LdapBindingUser field value if set, zero value otherwise.
func (o *Systemuserreturn) GetLdapBindingUser() bool {
	if o == nil || IsNil(o.LdapBindingUser) {
		var ret bool
		return ret
	}
	return *o.LdapBindingUser
}

// GetLdapBindingUserOk returns a tuple with the LdapBindingUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetLdapBindingUserOk() (*bool, bool) {
	if o == nil || IsNil(o.LdapBindingUser) {
		return nil, false
	}
	return o.LdapBindingUser, true
}

// HasLdapBindingUser returns a boolean if a field has been set.
func (o *Systemuserreturn) HasLdapBindingUser() bool {
	if o != nil && !IsNil(o.LdapBindingUser) {
		return true
	}

	return false
}

// SetLdapBindingUser gets a reference to the given bool and assigns it to the LdapBindingUser field.
func (o *Systemuserreturn) SetLdapBindingUser(v bool) {
	o.LdapBindingUser = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Systemuserreturn) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Systemuserreturn) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Systemuserreturn) SetLocation(v string) {
	o.Location = &v
}

// GetManagedAppleId returns the ManagedAppleId field value if set, zero value otherwise.
func (o *Systemuserreturn) GetManagedAppleId() string {
	if o == nil || IsNil(o.ManagedAppleId) {
		var ret string
		return ret
	}
	return *o.ManagedAppleId
}

// GetManagedAppleIdOk returns a tuple with the ManagedAppleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetManagedAppleIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedAppleId) {
		return nil, false
	}
	return o.ManagedAppleId, true
}

// HasManagedAppleId returns a boolean if a field has been set.
func (o *Systemuserreturn) HasManagedAppleId() bool {
	if o != nil && !IsNil(o.ManagedAppleId) {
		return true
	}

	return false
}

// SetManagedAppleId gets a reference to the given string and assigns it to the ManagedAppleId field.
func (o *Systemuserreturn) SetManagedAppleId(v string) {
	o.ManagedAppleId = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *Systemuserreturn) GetManager() string {
	if o == nil || IsNil(o.Manager) {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetManagerOk() (*string, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *Systemuserreturn) HasManager() bool {
	if o != nil && !IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given string and assigns it to the Manager field.
func (o *Systemuserreturn) SetManager(v string) {
	o.Manager = &v
}

// GetMfa returns the Mfa field value if set, zero value otherwise.
func (o *Systemuserreturn) GetMfa() Mfa {
	if o == nil || IsNil(o.Mfa) {
		var ret Mfa
		return ret
	}
	return *o.Mfa
}

// GetMfaOk returns a tuple with the Mfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetMfaOk() (*Mfa, bool) {
	if o == nil || IsNil(o.Mfa) {
		return nil, false
	}
	return o.Mfa, true
}

// HasMfa returns a boolean if a field has been set.
func (o *Systemuserreturn) HasMfa() bool {
	if o != nil && !IsNil(o.Mfa) {
		return true
	}

	return false
}

// SetMfa gets a reference to the given Mfa and assigns it to the Mfa field.
func (o *Systemuserreturn) SetMfa(v Mfa) {
	o.Mfa = &v
}

// GetMfaEnrollment returns the MfaEnrollment field value if set, zero value otherwise.
func (o *Systemuserreturn) GetMfaEnrollment() MfaEnrollment {
	if o == nil || IsNil(o.MfaEnrollment) {
		var ret MfaEnrollment
		return ret
	}
	return *o.MfaEnrollment
}

// GetMfaEnrollmentOk returns a tuple with the MfaEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetMfaEnrollmentOk() (*MfaEnrollment, bool) {
	if o == nil || IsNil(o.MfaEnrollment) {
		return nil, false
	}
	return o.MfaEnrollment, true
}

// HasMfaEnrollment returns a boolean if a field has been set.
func (o *Systemuserreturn) HasMfaEnrollment() bool {
	if o != nil && !IsNil(o.MfaEnrollment) {
		return true
	}

	return false
}

// SetMfaEnrollment gets a reference to the given MfaEnrollment and assigns it to the MfaEnrollment field.
func (o *Systemuserreturn) SetMfaEnrollment(v MfaEnrollment) {
	o.MfaEnrollment = &v
}

// GetMiddlename returns the Middlename field value if set, zero value otherwise.
func (o *Systemuserreturn) GetMiddlename() string {
	if o == nil || IsNil(o.Middlename) {
		var ret string
		return ret
	}
	return *o.Middlename
}

// GetMiddlenameOk returns a tuple with the Middlename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetMiddlenameOk() (*string, bool) {
	if o == nil || IsNil(o.Middlename) {
		return nil, false
	}
	return o.Middlename, true
}

// HasMiddlename returns a boolean if a field has been set.
func (o *Systemuserreturn) HasMiddlename() bool {
	if o != nil && !IsNil(o.Middlename) {
		return true
	}

	return false
}

// SetMiddlename gets a reference to the given string and assigns it to the Middlename field.
func (o *Systemuserreturn) SetMiddlename(v string) {
	o.Middlename = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Systemuserreturn) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Systemuserreturn) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *Systemuserreturn) SetOrganization(v string) {
	o.Organization = &v
}

// GetPasswordDate returns the PasswordDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Systemuserreturn) GetPasswordDate() string {
	if o == nil || IsNil(o.PasswordDate.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordDate.Get()
}

// GetPasswordDateOk returns a tuple with the PasswordDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Systemuserreturn) GetPasswordDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordDate.Get(), o.PasswordDate.IsSet()
}

// HasPasswordDate returns a boolean if a field has been set.
func (o *Systemuserreturn) HasPasswordDate() bool {
	if o != nil && o.PasswordDate.IsSet() {
		return true
	}

	return false
}

// SetPasswordDate gets a reference to the given NullableString and assigns it to the PasswordDate field.
func (o *Systemuserreturn) SetPasswordDate(v string) {
	o.PasswordDate.Set(&v)
}
// SetPasswordDateNil sets the value for PasswordDate to be an explicit nil
func (o *Systemuserreturn) SetPasswordDateNil() {
	o.PasswordDate.Set(nil)
}

// UnsetPasswordDate ensures that no value is present for PasswordDate, not even an explicit nil
func (o *Systemuserreturn) UnsetPasswordDate() {
	o.PasswordDate.Unset()
}

// GetPasswordExpirationDate returns the PasswordExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Systemuserreturn) GetPasswordExpirationDate() string {
	if o == nil || IsNil(o.PasswordExpirationDate.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordExpirationDate.Get()
}

// GetPasswordExpirationDateOk returns a tuple with the PasswordExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Systemuserreturn) GetPasswordExpirationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordExpirationDate.Get(), o.PasswordExpirationDate.IsSet()
}

// HasPasswordExpirationDate returns a boolean if a field has been set.
func (o *Systemuserreturn) HasPasswordExpirationDate() bool {
	if o != nil && o.PasswordExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetPasswordExpirationDate gets a reference to the given NullableString and assigns it to the PasswordExpirationDate field.
func (o *Systemuserreturn) SetPasswordExpirationDate(v string) {
	o.PasswordExpirationDate.Set(&v)
}
// SetPasswordExpirationDateNil sets the value for PasswordExpirationDate to be an explicit nil
func (o *Systemuserreturn) SetPasswordExpirationDateNil() {
	o.PasswordExpirationDate.Set(nil)
}

// UnsetPasswordExpirationDate ensures that no value is present for PasswordExpirationDate, not even an explicit nil
func (o *Systemuserreturn) UnsetPasswordExpirationDate() {
	o.PasswordExpirationDate.Unset()
}

// GetPasswordExpired returns the PasswordExpired field value if set, zero value otherwise.
func (o *Systemuserreturn) GetPasswordExpired() bool {
	if o == nil || IsNil(o.PasswordExpired) {
		var ret bool
		return ret
	}
	return *o.PasswordExpired
}

// GetPasswordExpiredOk returns a tuple with the PasswordExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetPasswordExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordExpired) {
		return nil, false
	}
	return o.PasswordExpired, true
}

// HasPasswordExpired returns a boolean if a field has been set.
func (o *Systemuserreturn) HasPasswordExpired() bool {
	if o != nil && !IsNil(o.PasswordExpired) {
		return true
	}

	return false
}

// SetPasswordExpired gets a reference to the given bool and assigns it to the PasswordExpired field.
func (o *Systemuserreturn) SetPasswordExpired(v bool) {
	o.PasswordExpired = &v
}

// GetPasswordNeverExpires returns the PasswordNeverExpires field value if set, zero value otherwise.
func (o *Systemuserreturn) GetPasswordNeverExpires() bool {
	if o == nil || IsNil(o.PasswordNeverExpires) {
		var ret bool
		return ret
	}
	return *o.PasswordNeverExpires
}

// GetPasswordNeverExpiresOk returns a tuple with the PasswordNeverExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetPasswordNeverExpiresOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordNeverExpires) {
		return nil, false
	}
	return o.PasswordNeverExpires, true
}

// HasPasswordNeverExpires returns a boolean if a field has been set.
func (o *Systemuserreturn) HasPasswordNeverExpires() bool {
	if o != nil && !IsNil(o.PasswordNeverExpires) {
		return true
	}

	return false
}

// SetPasswordNeverExpires gets a reference to the given bool and assigns it to the PasswordNeverExpires field.
func (o *Systemuserreturn) SetPasswordNeverExpires(v bool) {
	o.PasswordNeverExpires = &v
}

// GetPasswordlessSudo returns the PasswordlessSudo field value if set, zero value otherwise.
func (o *Systemuserreturn) GetPasswordlessSudo() bool {
	if o == nil || IsNil(o.PasswordlessSudo) {
		var ret bool
		return ret
	}
	return *o.PasswordlessSudo
}

// GetPasswordlessSudoOk returns a tuple with the PasswordlessSudo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetPasswordlessSudoOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordlessSudo) {
		return nil, false
	}
	return o.PasswordlessSudo, true
}

// HasPasswordlessSudo returns a boolean if a field has been set.
func (o *Systemuserreturn) HasPasswordlessSudo() bool {
	if o != nil && !IsNil(o.PasswordlessSudo) {
		return true
	}

	return false
}

// SetPasswordlessSudo gets a reference to the given bool and assigns it to the PasswordlessSudo field.
func (o *Systemuserreturn) SetPasswordlessSudo(v bool) {
	o.PasswordlessSudo = &v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *Systemuserreturn) GetPhoneNumbers() []SystemuserreturnPhoneNumbersInner {
	if o == nil || IsNil(o.PhoneNumbers) {
		var ret []SystemuserreturnPhoneNumbersInner
		return ret
	}
	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetPhoneNumbersOk() ([]SystemuserreturnPhoneNumbersInner, bool) {
	if o == nil || IsNil(o.PhoneNumbers) {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *Systemuserreturn) HasPhoneNumbers() bool {
	if o != nil && !IsNil(o.PhoneNumbers) {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []SystemuserreturnPhoneNumbersInner and assigns it to the PhoneNumbers field.
func (o *Systemuserreturn) SetPhoneNumbers(v []SystemuserreturnPhoneNumbersInner) {
	o.PhoneNumbers = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *Systemuserreturn) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *Systemuserreturn) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *Systemuserreturn) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRecoveryEmail returns the RecoveryEmail field value if set, zero value otherwise.
func (o *Systemuserreturn) GetRecoveryEmail() SystemuserreturnRecoveryEmail {
	if o == nil || IsNil(o.RecoveryEmail) {
		var ret SystemuserreturnRecoveryEmail
		return ret
	}
	return *o.RecoveryEmail
}

// GetRecoveryEmailOk returns a tuple with the RecoveryEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetRecoveryEmailOk() (*SystemuserreturnRecoveryEmail, bool) {
	if o == nil || IsNil(o.RecoveryEmail) {
		return nil, false
	}
	return o.RecoveryEmail, true
}

// HasRecoveryEmail returns a boolean if a field has been set.
func (o *Systemuserreturn) HasRecoveryEmail() bool {
	if o != nil && !IsNil(o.RecoveryEmail) {
		return true
	}

	return false
}

// SetRecoveryEmail gets a reference to the given SystemuserreturnRecoveryEmail and assigns it to the RecoveryEmail field.
func (o *Systemuserreturn) SetRecoveryEmail(v SystemuserreturnRecoveryEmail) {
	o.RecoveryEmail = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Systemuserreturn) GetRelationships() []SystemuserputRelationshipsInner {
	if o == nil || IsNil(o.Relationships) {
		var ret []SystemuserputRelationshipsInner
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetRelationshipsOk() ([]SystemuserputRelationshipsInner, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Systemuserreturn) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []SystemuserputRelationshipsInner and assigns it to the Relationships field.
func (o *Systemuserreturn) SetRelationships(v []SystemuserputRelationshipsInner) {
	o.Relationships = v
}

// GetSambaServiceUser returns the SambaServiceUser field value if set, zero value otherwise.
func (o *Systemuserreturn) GetSambaServiceUser() bool {
	if o == nil || IsNil(o.SambaServiceUser) {
		var ret bool
		return ret
	}
	return *o.SambaServiceUser
}

// GetSambaServiceUserOk returns a tuple with the SambaServiceUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetSambaServiceUserOk() (*bool, bool) {
	if o == nil || IsNil(o.SambaServiceUser) {
		return nil, false
	}
	return o.SambaServiceUser, true
}

// HasSambaServiceUser returns a boolean if a field has been set.
func (o *Systemuserreturn) HasSambaServiceUser() bool {
	if o != nil && !IsNil(o.SambaServiceUser) {
		return true
	}

	return false
}

// SetSambaServiceUser gets a reference to the given bool and assigns it to the SambaServiceUser field.
func (o *Systemuserreturn) SetSambaServiceUser(v bool) {
	o.SambaServiceUser = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *Systemuserreturn) GetSshKeys() []Sshkeylist {
	if o == nil || IsNil(o.SshKeys) {
		var ret []Sshkeylist
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetSshKeysOk() ([]Sshkeylist, bool) {
	if o == nil || IsNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *Systemuserreturn) HasSshKeys() bool {
	if o != nil && !IsNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []Sshkeylist and assigns it to the SshKeys field.
func (o *Systemuserreturn) SetSshKeys(v []Sshkeylist) {
	o.SshKeys = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Systemuserreturn) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Systemuserreturn) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Systemuserreturn) SetState(v string) {
	o.State = &v
}

// GetSudo returns the Sudo field value if set, zero value otherwise.
func (o *Systemuserreturn) GetSudo() bool {
	if o == nil || IsNil(o.Sudo) {
		var ret bool
		return ret
	}
	return *o.Sudo
}

// GetSudoOk returns a tuple with the Sudo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetSudoOk() (*bool, bool) {
	if o == nil || IsNil(o.Sudo) {
		return nil, false
	}
	return o.Sudo, true
}

// HasSudo returns a boolean if a field has been set.
func (o *Systemuserreturn) HasSudo() bool {
	if o != nil && !IsNil(o.Sudo) {
		return true
	}

	return false
}

// SetSudo gets a reference to the given bool and assigns it to the Sudo field.
func (o *Systemuserreturn) SetSudo(v bool) {
	o.Sudo = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *Systemuserreturn) GetSuspended() bool {
	if o == nil || IsNil(o.Suspended) {
		var ret bool
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetSuspendedOk() (*bool, bool) {
	if o == nil || IsNil(o.Suspended) {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *Systemuserreturn) HasSuspended() bool {
	if o != nil && !IsNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given bool and assigns it to the Suspended field.
func (o *Systemuserreturn) SetSuspended(v bool) {
	o.Suspended = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Systemuserreturn) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Systemuserreturn) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Systemuserreturn) SetTags(v []string) {
	o.Tags = v
}

// GetTotpEnabled returns the TotpEnabled field value if set, zero value otherwise.
func (o *Systemuserreturn) GetTotpEnabled() bool {
	if o == nil || IsNil(o.TotpEnabled) {
		var ret bool
		return ret
	}
	return *o.TotpEnabled
}

// GetTotpEnabledOk returns a tuple with the TotpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetTotpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TotpEnabled) {
		return nil, false
	}
	return o.TotpEnabled, true
}

// HasTotpEnabled returns a boolean if a field has been set.
func (o *Systemuserreturn) HasTotpEnabled() bool {
	if o != nil && !IsNil(o.TotpEnabled) {
		return true
	}

	return false
}

// SetTotpEnabled gets a reference to the given bool and assigns it to the TotpEnabled field.
func (o *Systemuserreturn) SetTotpEnabled(v bool) {
	o.TotpEnabled = &v
}

// GetUnixGuid returns the UnixGuid field value if set, zero value otherwise.
func (o *Systemuserreturn) GetUnixGuid() int32 {
	if o == nil || IsNil(o.UnixGuid) {
		var ret int32
		return ret
	}
	return *o.UnixGuid
}

// GetUnixGuidOk returns a tuple with the UnixGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetUnixGuidOk() (*int32, bool) {
	if o == nil || IsNil(o.UnixGuid) {
		return nil, false
	}
	return o.UnixGuid, true
}

// HasUnixGuid returns a boolean if a field has been set.
func (o *Systemuserreturn) HasUnixGuid() bool {
	if o != nil && !IsNil(o.UnixGuid) {
		return true
	}

	return false
}

// SetUnixGuid gets a reference to the given int32 and assigns it to the UnixGuid field.
func (o *Systemuserreturn) SetUnixGuid(v int32) {
	o.UnixGuid = &v
}

// GetUnixUid returns the UnixUid field value if set, zero value otherwise.
func (o *Systemuserreturn) GetUnixUid() int32 {
	if o == nil || IsNil(o.UnixUid) {
		var ret int32
		return ret
	}
	return *o.UnixUid
}

// GetUnixUidOk returns a tuple with the UnixUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetUnixUidOk() (*int32, bool) {
	if o == nil || IsNil(o.UnixUid) {
		return nil, false
	}
	return o.UnixUid, true
}

// HasUnixUid returns a boolean if a field has been set.
func (o *Systemuserreturn) HasUnixUid() bool {
	if o != nil && !IsNil(o.UnixUid) {
		return true
	}

	return false
}

// SetUnixUid gets a reference to the given int32 and assigns it to the UnixUid field.
func (o *Systemuserreturn) SetUnixUid(v int32) {
	o.UnixUid = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Systemuserreturn) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemuserreturn) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Systemuserreturn) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Systemuserreturn) SetUsername(v string) {
	o.Username = &v
}

func (o Systemuserreturn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Systemuserreturn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.AccountLocked) {
		toSerialize["account_locked"] = o.AccountLocked
	}
	if o.AccountLockedDate.IsSet() {
		toSerialize["account_locked_date"] = o.AccountLockedDate.Get()
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
	if !IsNil(o.BadLoginAttempts) {
		toSerialize["badLoginAttempts"] = o.BadLoginAttempts
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.CostCenter) {
		toSerialize["costCenter"] = o.CostCenter
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreationSource) {
		toSerialize["creationSource"] = o.CreationSource
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
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
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
	if !IsNil(o.MfaEnrollment) {
		toSerialize["mfaEnrollment"] = o.MfaEnrollment
	}
	if !IsNil(o.Middlename) {
		toSerialize["middlename"] = o.Middlename
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if o.PasswordDate.IsSet() {
		toSerialize["password_date"] = o.PasswordDate.Get()
	}
	if o.PasswordExpirationDate.IsSet() {
		toSerialize["password_expiration_date"] = o.PasswordExpirationDate.Get()
	}
	if !IsNil(o.PasswordExpired) {
		toSerialize["password_expired"] = o.PasswordExpired
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
	if !IsNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
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
	if !IsNil(o.TotpEnabled) {
		toSerialize["totp_enabled"] = o.TotpEnabled
	}
	if !IsNil(o.UnixGuid) {
		toSerialize["unix_guid"] = o.UnixGuid
	}
	if !IsNil(o.UnixUid) {
		toSerialize["unix_uid"] = o.UnixUid
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Systemuserreturn) UnmarshalJSON(bytes []byte) (err error) {
	varSystemuserreturn := _Systemuserreturn{}

	if err = json.Unmarshal(bytes, &varSystemuserreturn); err == nil {
		*o = Systemuserreturn(varSystemuserreturn)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "account_locked")
		delete(additionalProperties, "account_locked_date")
		delete(additionalProperties, "activated")
		delete(additionalProperties, "addresses")
		delete(additionalProperties, "allow_public_key")
		delete(additionalProperties, "alternateEmail")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "badLoginAttempts")
		delete(additionalProperties, "company")
		delete(additionalProperties, "costCenter")
		delete(additionalProperties, "created")
		delete(additionalProperties, "creationSource")
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
		delete(additionalProperties, "mfaEnrollment")
		delete(additionalProperties, "middlename")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "password_date")
		delete(additionalProperties, "password_expiration_date")
		delete(additionalProperties, "password_expired")
		delete(additionalProperties, "password_never_expires")
		delete(additionalProperties, "passwordless_sudo")
		delete(additionalProperties, "phoneNumbers")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "recoveryEmail")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "samba_service_user")
		delete(additionalProperties, "ssh_keys")
		delete(additionalProperties, "state")
		delete(additionalProperties, "sudo")
		delete(additionalProperties, "suspended")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "totp_enabled")
		delete(additionalProperties, "unix_guid")
		delete(additionalProperties, "unix_uid")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemuserreturn struct {
	value *Systemuserreturn
	isSet bool
}

func (v NullableSystemuserreturn) Get() *Systemuserreturn {
	return v.value
}

func (v *NullableSystemuserreturn) Set(val *Systemuserreturn) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemuserreturn) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemuserreturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemuserreturn(val *Systemuserreturn) *NullableSystemuserreturn {
	return &NullableSystemuserreturn{value: val, isSet: true}
}

func (v NullableSystemuserreturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemuserreturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



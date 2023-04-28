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

// checks if the Organizationsettingsput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organizationsettingsput{}

// Organizationsettingsput struct for Organizationsettingsput
type Organizationsettingsput struct {
	ContactEmail *string `json:"contactEmail,omitempty"`
	ContactName *string `json:"contactName,omitempty"`
	DeviceIdentificationEnabled *bool `json:"deviceIdentificationEnabled,omitempty"`
	DisableGoogleLogin *bool `json:"disableGoogleLogin,omitempty"`
	DisableLdap *bool `json:"disableLdap,omitempty"`
	DisableUM *bool `json:"disableUM,omitempty"`
	DuplicateLDAPGroups *bool `json:"duplicateLDAPGroups,omitempty"`
	EmailDisclaimer *string `json:"emailDisclaimer,omitempty"`
	EnableManagedUID *bool `json:"enableManagedUID,omitempty"`
	Features *OrganizationsettingsFeatures `json:"features,omitempty"`
	// Object containing Optimizely experimentIds and states corresponding to them
	GrowthData map[string]interface{} `json:"growthData,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Name *string `json:"name,omitempty"`
	NewSystemUserStateDefaults *OrganizationsettingsputNewSystemUserStateDefaults `json:"newSystemUserStateDefaults,omitempty"`
	PasswordCompliance *string `json:"passwordCompliance,omitempty"`
	PasswordPolicy *OrganizationsettingsputPasswordPolicy `json:"passwordPolicy,omitempty"`
	ShowIntro *bool `json:"showIntro,omitempty"`
	SystemUserPasswordExpirationInDays *int32 `json:"systemUserPasswordExpirationInDays,omitempty"`
	SystemUsersCanEdit *bool `json:"systemUsersCanEdit,omitempty"`
	SystemUsersCap *int32 `json:"systemUsersCap,omitempty"`
	TrustedAppConfig *TrustedappConfigPut `json:"trustedAppConfig,omitempty"`
	UserPortal *OrganizationsettingsUserPortal `json:"userPortal,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Organizationsettingsput Organizationsettingsput

// NewOrganizationsettingsput instantiates a new Organizationsettingsput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsettingsput() *Organizationsettingsput {
	this := Organizationsettingsput{}
	return &this
}

// NewOrganizationsettingsputWithDefaults instantiates a new Organizationsettingsput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsettingsputWithDefaults() *Organizationsettingsput {
	this := Organizationsettingsput{}
	return &this
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *Organizationsettingsput) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetContactName() string {
	if o == nil || IsNil(o.ContactName) {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactName) {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasContactName() bool {
	if o != nil && !IsNil(o.ContactName) {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *Organizationsettingsput) SetContactName(v string) {
	o.ContactName = &v
}

// GetDeviceIdentificationEnabled returns the DeviceIdentificationEnabled field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetDeviceIdentificationEnabled() bool {
	if o == nil || IsNil(o.DeviceIdentificationEnabled) {
		var ret bool
		return ret
	}
	return *o.DeviceIdentificationEnabled
}

// GetDeviceIdentificationEnabledOk returns a tuple with the DeviceIdentificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetDeviceIdentificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DeviceIdentificationEnabled) {
		return nil, false
	}
	return o.DeviceIdentificationEnabled, true
}

// HasDeviceIdentificationEnabled returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasDeviceIdentificationEnabled() bool {
	if o != nil && !IsNil(o.DeviceIdentificationEnabled) {
		return true
	}

	return false
}

// SetDeviceIdentificationEnabled gets a reference to the given bool and assigns it to the DeviceIdentificationEnabled field.
func (o *Organizationsettingsput) SetDeviceIdentificationEnabled(v bool) {
	o.DeviceIdentificationEnabled = &v
}

// GetDisableGoogleLogin returns the DisableGoogleLogin field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetDisableGoogleLogin() bool {
	if o == nil || IsNil(o.DisableGoogleLogin) {
		var ret bool
		return ret
	}
	return *o.DisableGoogleLogin
}

// GetDisableGoogleLoginOk returns a tuple with the DisableGoogleLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetDisableGoogleLoginOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableGoogleLogin) {
		return nil, false
	}
	return o.DisableGoogleLogin, true
}

// HasDisableGoogleLogin returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasDisableGoogleLogin() bool {
	if o != nil && !IsNil(o.DisableGoogleLogin) {
		return true
	}

	return false
}

// SetDisableGoogleLogin gets a reference to the given bool and assigns it to the DisableGoogleLogin field.
func (o *Organizationsettingsput) SetDisableGoogleLogin(v bool) {
	o.DisableGoogleLogin = &v
}

// GetDisableLdap returns the DisableLdap field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetDisableLdap() bool {
	if o == nil || IsNil(o.DisableLdap) {
		var ret bool
		return ret
	}
	return *o.DisableLdap
}

// GetDisableLdapOk returns a tuple with the DisableLdap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetDisableLdapOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableLdap) {
		return nil, false
	}
	return o.DisableLdap, true
}

// HasDisableLdap returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasDisableLdap() bool {
	if o != nil && !IsNil(o.DisableLdap) {
		return true
	}

	return false
}

// SetDisableLdap gets a reference to the given bool and assigns it to the DisableLdap field.
func (o *Organizationsettingsput) SetDisableLdap(v bool) {
	o.DisableLdap = &v
}

// GetDisableUM returns the DisableUM field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetDisableUM() bool {
	if o == nil || IsNil(o.DisableUM) {
		var ret bool
		return ret
	}
	return *o.DisableUM
}

// GetDisableUMOk returns a tuple with the DisableUM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetDisableUMOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableUM) {
		return nil, false
	}
	return o.DisableUM, true
}

// HasDisableUM returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasDisableUM() bool {
	if o != nil && !IsNil(o.DisableUM) {
		return true
	}

	return false
}

// SetDisableUM gets a reference to the given bool and assigns it to the DisableUM field.
func (o *Organizationsettingsput) SetDisableUM(v bool) {
	o.DisableUM = &v
}

// GetDuplicateLDAPGroups returns the DuplicateLDAPGroups field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetDuplicateLDAPGroups() bool {
	if o == nil || IsNil(o.DuplicateLDAPGroups) {
		var ret bool
		return ret
	}
	return *o.DuplicateLDAPGroups
}

// GetDuplicateLDAPGroupsOk returns a tuple with the DuplicateLDAPGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetDuplicateLDAPGroupsOk() (*bool, bool) {
	if o == nil || IsNil(o.DuplicateLDAPGroups) {
		return nil, false
	}
	return o.DuplicateLDAPGroups, true
}

// HasDuplicateLDAPGroups returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasDuplicateLDAPGroups() bool {
	if o != nil && !IsNil(o.DuplicateLDAPGroups) {
		return true
	}

	return false
}

// SetDuplicateLDAPGroups gets a reference to the given bool and assigns it to the DuplicateLDAPGroups field.
func (o *Organizationsettingsput) SetDuplicateLDAPGroups(v bool) {
	o.DuplicateLDAPGroups = &v
}

// GetEmailDisclaimer returns the EmailDisclaimer field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetEmailDisclaimer() string {
	if o == nil || IsNil(o.EmailDisclaimer) {
		var ret string
		return ret
	}
	return *o.EmailDisclaimer
}

// GetEmailDisclaimerOk returns a tuple with the EmailDisclaimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetEmailDisclaimerOk() (*string, bool) {
	if o == nil || IsNil(o.EmailDisclaimer) {
		return nil, false
	}
	return o.EmailDisclaimer, true
}

// HasEmailDisclaimer returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasEmailDisclaimer() bool {
	if o != nil && !IsNil(o.EmailDisclaimer) {
		return true
	}

	return false
}

// SetEmailDisclaimer gets a reference to the given string and assigns it to the EmailDisclaimer field.
func (o *Organizationsettingsput) SetEmailDisclaimer(v string) {
	o.EmailDisclaimer = &v
}

// GetEnableManagedUID returns the EnableManagedUID field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetEnableManagedUID() bool {
	if o == nil || IsNil(o.EnableManagedUID) {
		var ret bool
		return ret
	}
	return *o.EnableManagedUID
}

// GetEnableManagedUIDOk returns a tuple with the EnableManagedUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetEnableManagedUIDOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableManagedUID) {
		return nil, false
	}
	return o.EnableManagedUID, true
}

// HasEnableManagedUID returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasEnableManagedUID() bool {
	if o != nil && !IsNil(o.EnableManagedUID) {
		return true
	}

	return false
}

// SetEnableManagedUID gets a reference to the given bool and assigns it to the EnableManagedUID field.
func (o *Organizationsettingsput) SetEnableManagedUID(v bool) {
	o.EnableManagedUID = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetFeatures() OrganizationsettingsFeatures {
	if o == nil || IsNil(o.Features) {
		var ret OrganizationsettingsFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetFeaturesOk() (*OrganizationsettingsFeatures, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given OrganizationsettingsFeatures and assigns it to the Features field.
func (o *Organizationsettingsput) SetFeatures(v OrganizationsettingsFeatures) {
	o.Features = &v
}

// GetGrowthData returns the GrowthData field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetGrowthData() map[string]interface{} {
	if o == nil || IsNil(o.GrowthData) {
		var ret map[string]interface{}
		return ret
	}
	return o.GrowthData
}

// GetGrowthDataOk returns a tuple with the GrowthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetGrowthDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.GrowthData) {
		return map[string]interface{}{}, false
	}
	return o.GrowthData, true
}

// HasGrowthData returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasGrowthData() bool {
	if o != nil && !IsNil(o.GrowthData) {
		return true
	}

	return false
}

// SetGrowthData gets a reference to the given map[string]interface{} and assigns it to the GrowthData field.
func (o *Organizationsettingsput) SetGrowthData(v map[string]interface{}) {
	o.GrowthData = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *Organizationsettingsput) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Organizationsettingsput) SetName(v string) {
	o.Name = &v
}

// GetNewSystemUserStateDefaults returns the NewSystemUserStateDefaults field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetNewSystemUserStateDefaults() OrganizationsettingsputNewSystemUserStateDefaults {
	if o == nil || IsNil(o.NewSystemUserStateDefaults) {
		var ret OrganizationsettingsputNewSystemUserStateDefaults
		return ret
	}
	return *o.NewSystemUserStateDefaults
}

// GetNewSystemUserStateDefaultsOk returns a tuple with the NewSystemUserStateDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetNewSystemUserStateDefaultsOk() (*OrganizationsettingsputNewSystemUserStateDefaults, bool) {
	if o == nil || IsNil(o.NewSystemUserStateDefaults) {
		return nil, false
	}
	return o.NewSystemUserStateDefaults, true
}

// HasNewSystemUserStateDefaults returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasNewSystemUserStateDefaults() bool {
	if o != nil && !IsNil(o.NewSystemUserStateDefaults) {
		return true
	}

	return false
}

// SetNewSystemUserStateDefaults gets a reference to the given OrganizationsettingsputNewSystemUserStateDefaults and assigns it to the NewSystemUserStateDefaults field.
func (o *Organizationsettingsput) SetNewSystemUserStateDefaults(v OrganizationsettingsputNewSystemUserStateDefaults) {
	o.NewSystemUserStateDefaults = &v
}

// GetPasswordCompliance returns the PasswordCompliance field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetPasswordCompliance() string {
	if o == nil || IsNil(o.PasswordCompliance) {
		var ret string
		return ret
	}
	return *o.PasswordCompliance
}

// GetPasswordComplianceOk returns a tuple with the PasswordCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetPasswordComplianceOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordCompliance) {
		return nil, false
	}
	return o.PasswordCompliance, true
}

// HasPasswordCompliance returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasPasswordCompliance() bool {
	if o != nil && !IsNil(o.PasswordCompliance) {
		return true
	}

	return false
}

// SetPasswordCompliance gets a reference to the given string and assigns it to the PasswordCompliance field.
func (o *Organizationsettingsput) SetPasswordCompliance(v string) {
	o.PasswordCompliance = &v
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetPasswordPolicy() OrganizationsettingsputPasswordPolicy {
	if o == nil || IsNil(o.PasswordPolicy) {
		var ret OrganizationsettingsputPasswordPolicy
		return ret
	}
	return *o.PasswordPolicy
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetPasswordPolicyOk() (*OrganizationsettingsputPasswordPolicy, bool) {
	if o == nil || IsNil(o.PasswordPolicy) {
		return nil, false
	}
	return o.PasswordPolicy, true
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasPasswordPolicy() bool {
	if o != nil && !IsNil(o.PasswordPolicy) {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given OrganizationsettingsputPasswordPolicy and assigns it to the PasswordPolicy field.
func (o *Organizationsettingsput) SetPasswordPolicy(v OrganizationsettingsputPasswordPolicy) {
	o.PasswordPolicy = &v
}

// GetShowIntro returns the ShowIntro field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetShowIntro() bool {
	if o == nil || IsNil(o.ShowIntro) {
		var ret bool
		return ret
	}
	return *o.ShowIntro
}

// GetShowIntroOk returns a tuple with the ShowIntro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetShowIntroOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowIntro) {
		return nil, false
	}
	return o.ShowIntro, true
}

// HasShowIntro returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasShowIntro() bool {
	if o != nil && !IsNil(o.ShowIntro) {
		return true
	}

	return false
}

// SetShowIntro gets a reference to the given bool and assigns it to the ShowIntro field.
func (o *Organizationsettingsput) SetShowIntro(v bool) {
	o.ShowIntro = &v
}

// GetSystemUserPasswordExpirationInDays returns the SystemUserPasswordExpirationInDays field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetSystemUserPasswordExpirationInDays() int32 {
	if o == nil || IsNil(o.SystemUserPasswordExpirationInDays) {
		var ret int32
		return ret
	}
	return *o.SystemUserPasswordExpirationInDays
}

// GetSystemUserPasswordExpirationInDaysOk returns a tuple with the SystemUserPasswordExpirationInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetSystemUserPasswordExpirationInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.SystemUserPasswordExpirationInDays) {
		return nil, false
	}
	return o.SystemUserPasswordExpirationInDays, true
}

// HasSystemUserPasswordExpirationInDays returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasSystemUserPasswordExpirationInDays() bool {
	if o != nil && !IsNil(o.SystemUserPasswordExpirationInDays) {
		return true
	}

	return false
}

// SetSystemUserPasswordExpirationInDays gets a reference to the given int32 and assigns it to the SystemUserPasswordExpirationInDays field.
func (o *Organizationsettingsput) SetSystemUserPasswordExpirationInDays(v int32) {
	o.SystemUserPasswordExpirationInDays = &v
}

// GetSystemUsersCanEdit returns the SystemUsersCanEdit field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetSystemUsersCanEdit() bool {
	if o == nil || IsNil(o.SystemUsersCanEdit) {
		var ret bool
		return ret
	}
	return *o.SystemUsersCanEdit
}

// GetSystemUsersCanEditOk returns a tuple with the SystemUsersCanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetSystemUsersCanEditOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemUsersCanEdit) {
		return nil, false
	}
	return o.SystemUsersCanEdit, true
}

// HasSystemUsersCanEdit returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasSystemUsersCanEdit() bool {
	if o != nil && !IsNil(o.SystemUsersCanEdit) {
		return true
	}

	return false
}

// SetSystemUsersCanEdit gets a reference to the given bool and assigns it to the SystemUsersCanEdit field.
func (o *Organizationsettingsput) SetSystemUsersCanEdit(v bool) {
	o.SystemUsersCanEdit = &v
}

// GetSystemUsersCap returns the SystemUsersCap field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetSystemUsersCap() int32 {
	if o == nil || IsNil(o.SystemUsersCap) {
		var ret int32
		return ret
	}
	return *o.SystemUsersCap
}

// GetSystemUsersCapOk returns a tuple with the SystemUsersCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetSystemUsersCapOk() (*int32, bool) {
	if o == nil || IsNil(o.SystemUsersCap) {
		return nil, false
	}
	return o.SystemUsersCap, true
}

// HasSystemUsersCap returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasSystemUsersCap() bool {
	if o != nil && !IsNil(o.SystemUsersCap) {
		return true
	}

	return false
}

// SetSystemUsersCap gets a reference to the given int32 and assigns it to the SystemUsersCap field.
func (o *Organizationsettingsput) SetSystemUsersCap(v int32) {
	o.SystemUsersCap = &v
}

// GetTrustedAppConfig returns the TrustedAppConfig field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetTrustedAppConfig() TrustedappConfigPut {
	if o == nil || IsNil(o.TrustedAppConfig) {
		var ret TrustedappConfigPut
		return ret
	}
	return *o.TrustedAppConfig
}

// GetTrustedAppConfigOk returns a tuple with the TrustedAppConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetTrustedAppConfigOk() (*TrustedappConfigPut, bool) {
	if o == nil || IsNil(o.TrustedAppConfig) {
		return nil, false
	}
	return o.TrustedAppConfig, true
}

// HasTrustedAppConfig returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasTrustedAppConfig() bool {
	if o != nil && !IsNil(o.TrustedAppConfig) {
		return true
	}

	return false
}

// SetTrustedAppConfig gets a reference to the given TrustedappConfigPut and assigns it to the TrustedAppConfig field.
func (o *Organizationsettingsput) SetTrustedAppConfig(v TrustedappConfigPut) {
	o.TrustedAppConfig = &v
}

// GetUserPortal returns the UserPortal field value if set, zero value otherwise.
func (o *Organizationsettingsput) GetUserPortal() OrganizationsettingsUserPortal {
	if o == nil || IsNil(o.UserPortal) {
		var ret OrganizationsettingsUserPortal
		return ret
	}
	return *o.UserPortal
}

// GetUserPortalOk returns a tuple with the UserPortal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettingsput) GetUserPortalOk() (*OrganizationsettingsUserPortal, bool) {
	if o == nil || IsNil(o.UserPortal) {
		return nil, false
	}
	return o.UserPortal, true
}

// HasUserPortal returns a boolean if a field has been set.
func (o *Organizationsettingsput) HasUserPortal() bool {
	if o != nil && !IsNil(o.UserPortal) {
		return true
	}

	return false
}

// SetUserPortal gets a reference to the given OrganizationsettingsUserPortal and assigns it to the UserPortal field.
func (o *Organizationsettingsput) SetUserPortal(v OrganizationsettingsUserPortal) {
	o.UserPortal = &v
}

func (o Organizationsettingsput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organizationsettingsput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactEmail) {
		toSerialize["contactEmail"] = o.ContactEmail
	}
	if !IsNil(o.ContactName) {
		toSerialize["contactName"] = o.ContactName
	}
	if !IsNil(o.DeviceIdentificationEnabled) {
		toSerialize["deviceIdentificationEnabled"] = o.DeviceIdentificationEnabled
	}
	if !IsNil(o.DisableGoogleLogin) {
		toSerialize["disableGoogleLogin"] = o.DisableGoogleLogin
	}
	if !IsNil(o.DisableLdap) {
		toSerialize["disableLdap"] = o.DisableLdap
	}
	if !IsNil(o.DisableUM) {
		toSerialize["disableUM"] = o.DisableUM
	}
	if !IsNil(o.DuplicateLDAPGroups) {
		toSerialize["duplicateLDAPGroups"] = o.DuplicateLDAPGroups
	}
	if !IsNil(o.EmailDisclaimer) {
		toSerialize["emailDisclaimer"] = o.EmailDisclaimer
	}
	if !IsNil(o.EnableManagedUID) {
		toSerialize["enableManagedUID"] = o.EnableManagedUID
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.GrowthData) {
		toSerialize["growthData"] = o.GrowthData
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NewSystemUserStateDefaults) {
		toSerialize["newSystemUserStateDefaults"] = o.NewSystemUserStateDefaults
	}
	if !IsNil(o.PasswordCompliance) {
		toSerialize["passwordCompliance"] = o.PasswordCompliance
	}
	if !IsNil(o.PasswordPolicy) {
		toSerialize["passwordPolicy"] = o.PasswordPolicy
	}
	if !IsNil(o.ShowIntro) {
		toSerialize["showIntro"] = o.ShowIntro
	}
	if !IsNil(o.SystemUserPasswordExpirationInDays) {
		toSerialize["systemUserPasswordExpirationInDays"] = o.SystemUserPasswordExpirationInDays
	}
	if !IsNil(o.SystemUsersCanEdit) {
		toSerialize["systemUsersCanEdit"] = o.SystemUsersCanEdit
	}
	if !IsNil(o.SystemUsersCap) {
		toSerialize["systemUsersCap"] = o.SystemUsersCap
	}
	if !IsNil(o.TrustedAppConfig) {
		toSerialize["trustedAppConfig"] = o.TrustedAppConfig
	}
	if !IsNil(o.UserPortal) {
		toSerialize["userPortal"] = o.UserPortal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Organizationsettingsput) UnmarshalJSON(bytes []byte) (err error) {
	varOrganizationsettingsput := _Organizationsettingsput{}

	if err = json.Unmarshal(bytes, &varOrganizationsettingsput); err == nil {
		*o = Organizationsettingsput(varOrganizationsettingsput)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "contactName")
		delete(additionalProperties, "deviceIdentificationEnabled")
		delete(additionalProperties, "disableGoogleLogin")
		delete(additionalProperties, "disableLdap")
		delete(additionalProperties, "disableUM")
		delete(additionalProperties, "duplicateLDAPGroups")
		delete(additionalProperties, "emailDisclaimer")
		delete(additionalProperties, "enableManagedUID")
		delete(additionalProperties, "features")
		delete(additionalProperties, "growthData")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "newSystemUserStateDefaults")
		delete(additionalProperties, "passwordCompliance")
		delete(additionalProperties, "passwordPolicy")
		delete(additionalProperties, "showIntro")
		delete(additionalProperties, "systemUserPasswordExpirationInDays")
		delete(additionalProperties, "systemUsersCanEdit")
		delete(additionalProperties, "systemUsersCap")
		delete(additionalProperties, "trustedAppConfig")
		delete(additionalProperties, "userPortal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationsettingsput struct {
	value *Organizationsettingsput
	isSet bool
}

func (v NullableOrganizationsettingsput) Get() *Organizationsettingsput {
	return v.value
}

func (v *NullableOrganizationsettingsput) Set(val *Organizationsettingsput) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsettingsput) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsettingsput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsettingsput(val *Organizationsettingsput) *NullableOrganizationsettingsput {
	return &NullableOrganizationsettingsput{value: val, isSet: true}
}

func (v NullableOrganizationsettingsput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsettingsput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



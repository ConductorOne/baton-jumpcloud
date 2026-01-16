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

// checks if the Organizationsettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organizationsettings{}

// Organizationsettings 
type Organizationsettings struct {
	AgentVersion *string `json:"agentVersion,omitempty"`
	BetaFeatures map[string]interface{} `json:"betaFeatures,omitempty"`
	ContactEmail *string `json:"contactEmail,omitempty"`
	ContactName *string `json:"contactName,omitempty"`
	DeviceIdentificationEnabled *bool `json:"deviceIdentificationEnabled,omitempty"`
	DisableCommandRunner *bool `json:"disableCommandRunner,omitempty"`
	DisableGoogleLogin *bool `json:"disableGoogleLogin,omitempty"`
	DisableLdap *bool `json:"disableLdap,omitempty"`
	DisableUM *bool `json:"disableUM,omitempty"`
	DuplicateLDAPGroups *bool `json:"duplicateLDAPGroups,omitempty"`
	EmailDisclaimer *string `json:"emailDisclaimer,omitempty"`
	EnableGoogleApps *bool `json:"enableGoogleApps,omitempty"`
	EnableManagedUID *bool `json:"enableManagedUID,omitempty"`
	EnableO365 *bool `json:"enableO365,omitempty"`
	EnableUserPortalAgentInstall *bool `json:"enableUserPortalAgentInstall,omitempty"`
	Features *OrganizationsettingsFeatures `json:"features,omitempty"`
	// Object containing Optimizely experimentIds and states corresponding to them
	GrowthData map[string]interface{} `json:"growthData,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Name *string `json:"name,omitempty"`
	NewSystemUserStateDefaults *OrganizationsettingsNewSystemUserStateDefaults `json:"newSystemUserStateDefaults,omitempty"`
	PasswordCompliance *string `json:"passwordCompliance,omitempty"`
	PasswordPolicy *OrganizationsettingsPasswordPolicy `json:"passwordPolicy,omitempty"`
	PendingDelete *bool `json:"pendingDelete,omitempty"`
	ShowIntro *bool `json:"showIntro,omitempty"`
	SystemUserPasswordExpirationInDays *int32 `json:"systemUserPasswordExpirationInDays,omitempty"`
	SystemUsersCanEdit *bool `json:"systemUsersCanEdit,omitempty"`
	SystemUsersCap *int32 `json:"systemUsersCap,omitempty"`
	TrustedAppConfig *TrustedappConfigGet `json:"trustedAppConfig,omitempty"`
	UserPortal *OrganizationsettingsUserPortal `json:"userPortal,omitempty"`
	WindowsMDM *OrganizationsettingsWindowsMDM `json:"windowsMDM,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Organizationsettings Organizationsettings

// NewOrganizationsettings instantiates a new Organizationsettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsettings() *Organizationsettings {
	this := Organizationsettings{}
	return &this
}

// NewOrganizationsettingsWithDefaults instantiates a new Organizationsettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsettingsWithDefaults() *Organizationsettings {
	this := Organizationsettings{}
	return &this
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *Organizationsettings) GetAgentVersion() string {
	if o == nil || IsNil(o.AgentVersion) {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetAgentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AgentVersion) {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *Organizationsettings) HasAgentVersion() bool {
	if o != nil && !IsNil(o.AgentVersion) {
		return true
	}

	return false
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *Organizationsettings) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetBetaFeatures returns the BetaFeatures field value if set, zero value otherwise.
func (o *Organizationsettings) GetBetaFeatures() map[string]interface{} {
	if o == nil || IsNil(o.BetaFeatures) {
		var ret map[string]interface{}
		return ret
	}
	return o.BetaFeatures
}

// GetBetaFeaturesOk returns a tuple with the BetaFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetBetaFeaturesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.BetaFeatures) {
		return map[string]interface{}{}, false
	}
	return o.BetaFeatures, true
}

// HasBetaFeatures returns a boolean if a field has been set.
func (o *Organizationsettings) HasBetaFeatures() bool {
	if o != nil && !IsNil(o.BetaFeatures) {
		return true
	}

	return false
}

// SetBetaFeatures gets a reference to the given map[string]interface{} and assigns it to the BetaFeatures field.
func (o *Organizationsettings) SetBetaFeatures(v map[string]interface{}) {
	o.BetaFeatures = v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *Organizationsettings) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *Organizationsettings) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *Organizationsettings) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *Organizationsettings) GetContactName() string {
	if o == nil || IsNil(o.ContactName) {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactName) {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *Organizationsettings) HasContactName() bool {
	if o != nil && !IsNil(o.ContactName) {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *Organizationsettings) SetContactName(v string) {
	o.ContactName = &v
}

// GetDeviceIdentificationEnabled returns the DeviceIdentificationEnabled field value if set, zero value otherwise.
func (o *Organizationsettings) GetDeviceIdentificationEnabled() bool {
	if o == nil || IsNil(o.DeviceIdentificationEnabled) {
		var ret bool
		return ret
	}
	return *o.DeviceIdentificationEnabled
}

// GetDeviceIdentificationEnabledOk returns a tuple with the DeviceIdentificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetDeviceIdentificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DeviceIdentificationEnabled) {
		return nil, false
	}
	return o.DeviceIdentificationEnabled, true
}

// HasDeviceIdentificationEnabled returns a boolean if a field has been set.
func (o *Organizationsettings) HasDeviceIdentificationEnabled() bool {
	if o != nil && !IsNil(o.DeviceIdentificationEnabled) {
		return true
	}

	return false
}

// SetDeviceIdentificationEnabled gets a reference to the given bool and assigns it to the DeviceIdentificationEnabled field.
func (o *Organizationsettings) SetDeviceIdentificationEnabled(v bool) {
	o.DeviceIdentificationEnabled = &v
}

// GetDisableCommandRunner returns the DisableCommandRunner field value if set, zero value otherwise.
func (o *Organizationsettings) GetDisableCommandRunner() bool {
	if o == nil || IsNil(o.DisableCommandRunner) {
		var ret bool
		return ret
	}
	return *o.DisableCommandRunner
}

// GetDisableCommandRunnerOk returns a tuple with the DisableCommandRunner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetDisableCommandRunnerOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableCommandRunner) {
		return nil, false
	}
	return o.DisableCommandRunner, true
}

// HasDisableCommandRunner returns a boolean if a field has been set.
func (o *Organizationsettings) HasDisableCommandRunner() bool {
	if o != nil && !IsNil(o.DisableCommandRunner) {
		return true
	}

	return false
}

// SetDisableCommandRunner gets a reference to the given bool and assigns it to the DisableCommandRunner field.
func (o *Organizationsettings) SetDisableCommandRunner(v bool) {
	o.DisableCommandRunner = &v
}

// GetDisableGoogleLogin returns the DisableGoogleLogin field value if set, zero value otherwise.
func (o *Organizationsettings) GetDisableGoogleLogin() bool {
	if o == nil || IsNil(o.DisableGoogleLogin) {
		var ret bool
		return ret
	}
	return *o.DisableGoogleLogin
}

// GetDisableGoogleLoginOk returns a tuple with the DisableGoogleLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetDisableGoogleLoginOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableGoogleLogin) {
		return nil, false
	}
	return o.DisableGoogleLogin, true
}

// HasDisableGoogleLogin returns a boolean if a field has been set.
func (o *Organizationsettings) HasDisableGoogleLogin() bool {
	if o != nil && !IsNil(o.DisableGoogleLogin) {
		return true
	}

	return false
}

// SetDisableGoogleLogin gets a reference to the given bool and assigns it to the DisableGoogleLogin field.
func (o *Organizationsettings) SetDisableGoogleLogin(v bool) {
	o.DisableGoogleLogin = &v
}

// GetDisableLdap returns the DisableLdap field value if set, zero value otherwise.
func (o *Organizationsettings) GetDisableLdap() bool {
	if o == nil || IsNil(o.DisableLdap) {
		var ret bool
		return ret
	}
	return *o.DisableLdap
}

// GetDisableLdapOk returns a tuple with the DisableLdap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetDisableLdapOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableLdap) {
		return nil, false
	}
	return o.DisableLdap, true
}

// HasDisableLdap returns a boolean if a field has been set.
func (o *Organizationsettings) HasDisableLdap() bool {
	if o != nil && !IsNil(o.DisableLdap) {
		return true
	}

	return false
}

// SetDisableLdap gets a reference to the given bool and assigns it to the DisableLdap field.
func (o *Organizationsettings) SetDisableLdap(v bool) {
	o.DisableLdap = &v
}

// GetDisableUM returns the DisableUM field value if set, zero value otherwise.
func (o *Organizationsettings) GetDisableUM() bool {
	if o == nil || IsNil(o.DisableUM) {
		var ret bool
		return ret
	}
	return *o.DisableUM
}

// GetDisableUMOk returns a tuple with the DisableUM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetDisableUMOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableUM) {
		return nil, false
	}
	return o.DisableUM, true
}

// HasDisableUM returns a boolean if a field has been set.
func (o *Organizationsettings) HasDisableUM() bool {
	if o != nil && !IsNil(o.DisableUM) {
		return true
	}

	return false
}

// SetDisableUM gets a reference to the given bool and assigns it to the DisableUM field.
func (o *Organizationsettings) SetDisableUM(v bool) {
	o.DisableUM = &v
}

// GetDuplicateLDAPGroups returns the DuplicateLDAPGroups field value if set, zero value otherwise.
func (o *Organizationsettings) GetDuplicateLDAPGroups() bool {
	if o == nil || IsNil(o.DuplicateLDAPGroups) {
		var ret bool
		return ret
	}
	return *o.DuplicateLDAPGroups
}

// GetDuplicateLDAPGroupsOk returns a tuple with the DuplicateLDAPGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetDuplicateLDAPGroupsOk() (*bool, bool) {
	if o == nil || IsNil(o.DuplicateLDAPGroups) {
		return nil, false
	}
	return o.DuplicateLDAPGroups, true
}

// HasDuplicateLDAPGroups returns a boolean if a field has been set.
func (o *Organizationsettings) HasDuplicateLDAPGroups() bool {
	if o != nil && !IsNil(o.DuplicateLDAPGroups) {
		return true
	}

	return false
}

// SetDuplicateLDAPGroups gets a reference to the given bool and assigns it to the DuplicateLDAPGroups field.
func (o *Organizationsettings) SetDuplicateLDAPGroups(v bool) {
	o.DuplicateLDAPGroups = &v
}

// GetEmailDisclaimer returns the EmailDisclaimer field value if set, zero value otherwise.
func (o *Organizationsettings) GetEmailDisclaimer() string {
	if o == nil || IsNil(o.EmailDisclaimer) {
		var ret string
		return ret
	}
	return *o.EmailDisclaimer
}

// GetEmailDisclaimerOk returns a tuple with the EmailDisclaimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetEmailDisclaimerOk() (*string, bool) {
	if o == nil || IsNil(o.EmailDisclaimer) {
		return nil, false
	}
	return o.EmailDisclaimer, true
}

// HasEmailDisclaimer returns a boolean if a field has been set.
func (o *Organizationsettings) HasEmailDisclaimer() bool {
	if o != nil && !IsNil(o.EmailDisclaimer) {
		return true
	}

	return false
}

// SetEmailDisclaimer gets a reference to the given string and assigns it to the EmailDisclaimer field.
func (o *Organizationsettings) SetEmailDisclaimer(v string) {
	o.EmailDisclaimer = &v
}

// GetEnableGoogleApps returns the EnableGoogleApps field value if set, zero value otherwise.
func (o *Organizationsettings) GetEnableGoogleApps() bool {
	if o == nil || IsNil(o.EnableGoogleApps) {
		var ret bool
		return ret
	}
	return *o.EnableGoogleApps
}

// GetEnableGoogleAppsOk returns a tuple with the EnableGoogleApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetEnableGoogleAppsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableGoogleApps) {
		return nil, false
	}
	return o.EnableGoogleApps, true
}

// HasEnableGoogleApps returns a boolean if a field has been set.
func (o *Organizationsettings) HasEnableGoogleApps() bool {
	if o != nil && !IsNil(o.EnableGoogleApps) {
		return true
	}

	return false
}

// SetEnableGoogleApps gets a reference to the given bool and assigns it to the EnableGoogleApps field.
func (o *Organizationsettings) SetEnableGoogleApps(v bool) {
	o.EnableGoogleApps = &v
}

// GetEnableManagedUID returns the EnableManagedUID field value if set, zero value otherwise.
func (o *Organizationsettings) GetEnableManagedUID() bool {
	if o == nil || IsNil(o.EnableManagedUID) {
		var ret bool
		return ret
	}
	return *o.EnableManagedUID
}

// GetEnableManagedUIDOk returns a tuple with the EnableManagedUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetEnableManagedUIDOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableManagedUID) {
		return nil, false
	}
	return o.EnableManagedUID, true
}

// HasEnableManagedUID returns a boolean if a field has been set.
func (o *Organizationsettings) HasEnableManagedUID() bool {
	if o != nil && !IsNil(o.EnableManagedUID) {
		return true
	}

	return false
}

// SetEnableManagedUID gets a reference to the given bool and assigns it to the EnableManagedUID field.
func (o *Organizationsettings) SetEnableManagedUID(v bool) {
	o.EnableManagedUID = &v
}

// GetEnableO365 returns the EnableO365 field value if set, zero value otherwise.
func (o *Organizationsettings) GetEnableO365() bool {
	if o == nil || IsNil(o.EnableO365) {
		var ret bool
		return ret
	}
	return *o.EnableO365
}

// GetEnableO365Ok returns a tuple with the EnableO365 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetEnableO365Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableO365) {
		return nil, false
	}
	return o.EnableO365, true
}

// HasEnableO365 returns a boolean if a field has been set.
func (o *Organizationsettings) HasEnableO365() bool {
	if o != nil && !IsNil(o.EnableO365) {
		return true
	}

	return false
}

// SetEnableO365 gets a reference to the given bool and assigns it to the EnableO365 field.
func (o *Organizationsettings) SetEnableO365(v bool) {
	o.EnableO365 = &v
}

// GetEnableUserPortalAgentInstall returns the EnableUserPortalAgentInstall field value if set, zero value otherwise.
func (o *Organizationsettings) GetEnableUserPortalAgentInstall() bool {
	if o == nil || IsNil(o.EnableUserPortalAgentInstall) {
		var ret bool
		return ret
	}
	return *o.EnableUserPortalAgentInstall
}

// GetEnableUserPortalAgentInstallOk returns a tuple with the EnableUserPortalAgentInstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetEnableUserPortalAgentInstallOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableUserPortalAgentInstall) {
		return nil, false
	}
	return o.EnableUserPortalAgentInstall, true
}

// HasEnableUserPortalAgentInstall returns a boolean if a field has been set.
func (o *Organizationsettings) HasEnableUserPortalAgentInstall() bool {
	if o != nil && !IsNil(o.EnableUserPortalAgentInstall) {
		return true
	}

	return false
}

// SetEnableUserPortalAgentInstall gets a reference to the given bool and assigns it to the EnableUserPortalAgentInstall field.
func (o *Organizationsettings) SetEnableUserPortalAgentInstall(v bool) {
	o.EnableUserPortalAgentInstall = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Organizationsettings) GetFeatures() OrganizationsettingsFeatures {
	if o == nil || IsNil(o.Features) {
		var ret OrganizationsettingsFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetFeaturesOk() (*OrganizationsettingsFeatures, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Organizationsettings) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given OrganizationsettingsFeatures and assigns it to the Features field.
func (o *Organizationsettings) SetFeatures(v OrganizationsettingsFeatures) {
	o.Features = &v
}

// GetGrowthData returns the GrowthData field value if set, zero value otherwise.
func (o *Organizationsettings) GetGrowthData() map[string]interface{} {
	if o == nil || IsNil(o.GrowthData) {
		var ret map[string]interface{}
		return ret
	}
	return o.GrowthData
}

// GetGrowthDataOk returns a tuple with the GrowthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetGrowthDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.GrowthData) {
		return map[string]interface{}{}, false
	}
	return o.GrowthData, true
}

// HasGrowthData returns a boolean if a field has been set.
func (o *Organizationsettings) HasGrowthData() bool {
	if o != nil && !IsNil(o.GrowthData) {
		return true
	}

	return false
}

// SetGrowthData gets a reference to the given map[string]interface{} and assigns it to the GrowthData field.
func (o *Organizationsettings) SetGrowthData(v map[string]interface{}) {
	o.GrowthData = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *Organizationsettings) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *Organizationsettings) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *Organizationsettings) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Organizationsettings) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Organizationsettings) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Organizationsettings) SetName(v string) {
	o.Name = &v
}

// GetNewSystemUserStateDefaults returns the NewSystemUserStateDefaults field value if set, zero value otherwise.
func (o *Organizationsettings) GetNewSystemUserStateDefaults() OrganizationsettingsNewSystemUserStateDefaults {
	if o == nil || IsNil(o.NewSystemUserStateDefaults) {
		var ret OrganizationsettingsNewSystemUserStateDefaults
		return ret
	}
	return *o.NewSystemUserStateDefaults
}

// GetNewSystemUserStateDefaultsOk returns a tuple with the NewSystemUserStateDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetNewSystemUserStateDefaultsOk() (*OrganizationsettingsNewSystemUserStateDefaults, bool) {
	if o == nil || IsNil(o.NewSystemUserStateDefaults) {
		return nil, false
	}
	return o.NewSystemUserStateDefaults, true
}

// HasNewSystemUserStateDefaults returns a boolean if a field has been set.
func (o *Organizationsettings) HasNewSystemUserStateDefaults() bool {
	if o != nil && !IsNil(o.NewSystemUserStateDefaults) {
		return true
	}

	return false
}

// SetNewSystemUserStateDefaults gets a reference to the given OrganizationsettingsNewSystemUserStateDefaults and assigns it to the NewSystemUserStateDefaults field.
func (o *Organizationsettings) SetNewSystemUserStateDefaults(v OrganizationsettingsNewSystemUserStateDefaults) {
	o.NewSystemUserStateDefaults = &v
}

// GetPasswordCompliance returns the PasswordCompliance field value if set, zero value otherwise.
func (o *Organizationsettings) GetPasswordCompliance() string {
	if o == nil || IsNil(o.PasswordCompliance) {
		var ret string
		return ret
	}
	return *o.PasswordCompliance
}

// GetPasswordComplianceOk returns a tuple with the PasswordCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetPasswordComplianceOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordCompliance) {
		return nil, false
	}
	return o.PasswordCompliance, true
}

// HasPasswordCompliance returns a boolean if a field has been set.
func (o *Organizationsettings) HasPasswordCompliance() bool {
	if o != nil && !IsNil(o.PasswordCompliance) {
		return true
	}

	return false
}

// SetPasswordCompliance gets a reference to the given string and assigns it to the PasswordCompliance field.
func (o *Organizationsettings) SetPasswordCompliance(v string) {
	o.PasswordCompliance = &v
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise.
func (o *Organizationsettings) GetPasswordPolicy() OrganizationsettingsPasswordPolicy {
	if o == nil || IsNil(o.PasswordPolicy) {
		var ret OrganizationsettingsPasswordPolicy
		return ret
	}
	return *o.PasswordPolicy
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetPasswordPolicyOk() (*OrganizationsettingsPasswordPolicy, bool) {
	if o == nil || IsNil(o.PasswordPolicy) {
		return nil, false
	}
	return o.PasswordPolicy, true
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *Organizationsettings) HasPasswordPolicy() bool {
	if o != nil && !IsNil(o.PasswordPolicy) {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given OrganizationsettingsPasswordPolicy and assigns it to the PasswordPolicy field.
func (o *Organizationsettings) SetPasswordPolicy(v OrganizationsettingsPasswordPolicy) {
	o.PasswordPolicy = &v
}

// GetPendingDelete returns the PendingDelete field value if set, zero value otherwise.
func (o *Organizationsettings) GetPendingDelete() bool {
	if o == nil || IsNil(o.PendingDelete) {
		var ret bool
		return ret
	}
	return *o.PendingDelete
}

// GetPendingDeleteOk returns a tuple with the PendingDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetPendingDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.PendingDelete) {
		return nil, false
	}
	return o.PendingDelete, true
}

// HasPendingDelete returns a boolean if a field has been set.
func (o *Organizationsettings) HasPendingDelete() bool {
	if o != nil && !IsNil(o.PendingDelete) {
		return true
	}

	return false
}

// SetPendingDelete gets a reference to the given bool and assigns it to the PendingDelete field.
func (o *Organizationsettings) SetPendingDelete(v bool) {
	o.PendingDelete = &v
}

// GetShowIntro returns the ShowIntro field value if set, zero value otherwise.
func (o *Organizationsettings) GetShowIntro() bool {
	if o == nil || IsNil(o.ShowIntro) {
		var ret bool
		return ret
	}
	return *o.ShowIntro
}

// GetShowIntroOk returns a tuple with the ShowIntro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetShowIntroOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowIntro) {
		return nil, false
	}
	return o.ShowIntro, true
}

// HasShowIntro returns a boolean if a field has been set.
func (o *Organizationsettings) HasShowIntro() bool {
	if o != nil && !IsNil(o.ShowIntro) {
		return true
	}

	return false
}

// SetShowIntro gets a reference to the given bool and assigns it to the ShowIntro field.
func (o *Organizationsettings) SetShowIntro(v bool) {
	o.ShowIntro = &v
}

// GetSystemUserPasswordExpirationInDays returns the SystemUserPasswordExpirationInDays field value if set, zero value otherwise.
func (o *Organizationsettings) GetSystemUserPasswordExpirationInDays() int32 {
	if o == nil || IsNil(o.SystemUserPasswordExpirationInDays) {
		var ret int32
		return ret
	}
	return *o.SystemUserPasswordExpirationInDays
}

// GetSystemUserPasswordExpirationInDaysOk returns a tuple with the SystemUserPasswordExpirationInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetSystemUserPasswordExpirationInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.SystemUserPasswordExpirationInDays) {
		return nil, false
	}
	return o.SystemUserPasswordExpirationInDays, true
}

// HasSystemUserPasswordExpirationInDays returns a boolean if a field has been set.
func (o *Organizationsettings) HasSystemUserPasswordExpirationInDays() bool {
	if o != nil && !IsNil(o.SystemUserPasswordExpirationInDays) {
		return true
	}

	return false
}

// SetSystemUserPasswordExpirationInDays gets a reference to the given int32 and assigns it to the SystemUserPasswordExpirationInDays field.
func (o *Organizationsettings) SetSystemUserPasswordExpirationInDays(v int32) {
	o.SystemUserPasswordExpirationInDays = &v
}

// GetSystemUsersCanEdit returns the SystemUsersCanEdit field value if set, zero value otherwise.
func (o *Organizationsettings) GetSystemUsersCanEdit() bool {
	if o == nil || IsNil(o.SystemUsersCanEdit) {
		var ret bool
		return ret
	}
	return *o.SystemUsersCanEdit
}

// GetSystemUsersCanEditOk returns a tuple with the SystemUsersCanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetSystemUsersCanEditOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemUsersCanEdit) {
		return nil, false
	}
	return o.SystemUsersCanEdit, true
}

// HasSystemUsersCanEdit returns a boolean if a field has been set.
func (o *Organizationsettings) HasSystemUsersCanEdit() bool {
	if o != nil && !IsNil(o.SystemUsersCanEdit) {
		return true
	}

	return false
}

// SetSystemUsersCanEdit gets a reference to the given bool and assigns it to the SystemUsersCanEdit field.
func (o *Organizationsettings) SetSystemUsersCanEdit(v bool) {
	o.SystemUsersCanEdit = &v
}

// GetSystemUsersCap returns the SystemUsersCap field value if set, zero value otherwise.
func (o *Organizationsettings) GetSystemUsersCap() int32 {
	if o == nil || IsNil(o.SystemUsersCap) {
		var ret int32
		return ret
	}
	return *o.SystemUsersCap
}

// GetSystemUsersCapOk returns a tuple with the SystemUsersCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetSystemUsersCapOk() (*int32, bool) {
	if o == nil || IsNil(o.SystemUsersCap) {
		return nil, false
	}
	return o.SystemUsersCap, true
}

// HasSystemUsersCap returns a boolean if a field has been set.
func (o *Organizationsettings) HasSystemUsersCap() bool {
	if o != nil && !IsNil(o.SystemUsersCap) {
		return true
	}

	return false
}

// SetSystemUsersCap gets a reference to the given int32 and assigns it to the SystemUsersCap field.
func (o *Organizationsettings) SetSystemUsersCap(v int32) {
	o.SystemUsersCap = &v
}

// GetTrustedAppConfig returns the TrustedAppConfig field value if set, zero value otherwise.
func (o *Organizationsettings) GetTrustedAppConfig() TrustedappConfigGet {
	if o == nil || IsNil(o.TrustedAppConfig) {
		var ret TrustedappConfigGet
		return ret
	}
	return *o.TrustedAppConfig
}

// GetTrustedAppConfigOk returns a tuple with the TrustedAppConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetTrustedAppConfigOk() (*TrustedappConfigGet, bool) {
	if o == nil || IsNil(o.TrustedAppConfig) {
		return nil, false
	}
	return o.TrustedAppConfig, true
}

// HasTrustedAppConfig returns a boolean if a field has been set.
func (o *Organizationsettings) HasTrustedAppConfig() bool {
	if o != nil && !IsNil(o.TrustedAppConfig) {
		return true
	}

	return false
}

// SetTrustedAppConfig gets a reference to the given TrustedappConfigGet and assigns it to the TrustedAppConfig field.
func (o *Organizationsettings) SetTrustedAppConfig(v TrustedappConfigGet) {
	o.TrustedAppConfig = &v
}

// GetUserPortal returns the UserPortal field value if set, zero value otherwise.
func (o *Organizationsettings) GetUserPortal() OrganizationsettingsUserPortal {
	if o == nil || IsNil(o.UserPortal) {
		var ret OrganizationsettingsUserPortal
		return ret
	}
	return *o.UserPortal
}

// GetUserPortalOk returns a tuple with the UserPortal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetUserPortalOk() (*OrganizationsettingsUserPortal, bool) {
	if o == nil || IsNil(o.UserPortal) {
		return nil, false
	}
	return o.UserPortal, true
}

// HasUserPortal returns a boolean if a field has been set.
func (o *Organizationsettings) HasUserPortal() bool {
	if o != nil && !IsNil(o.UserPortal) {
		return true
	}

	return false
}

// SetUserPortal gets a reference to the given OrganizationsettingsUserPortal and assigns it to the UserPortal field.
func (o *Organizationsettings) SetUserPortal(v OrganizationsettingsUserPortal) {
	o.UserPortal = &v
}

// GetWindowsMDM returns the WindowsMDM field value if set, zero value otherwise.
func (o *Organizationsettings) GetWindowsMDM() OrganizationsettingsWindowsMDM {
	if o == nil || IsNil(o.WindowsMDM) {
		var ret OrganizationsettingsWindowsMDM
		return ret
	}
	return *o.WindowsMDM
}

// GetWindowsMDMOk returns a tuple with the WindowsMDM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organizationsettings) GetWindowsMDMOk() (*OrganizationsettingsWindowsMDM, bool) {
	if o == nil || IsNil(o.WindowsMDM) {
		return nil, false
	}
	return o.WindowsMDM, true
}

// HasWindowsMDM returns a boolean if a field has been set.
func (o *Organizationsettings) HasWindowsMDM() bool {
	if o != nil && !IsNil(o.WindowsMDM) {
		return true
	}

	return false
}

// SetWindowsMDM gets a reference to the given OrganizationsettingsWindowsMDM and assigns it to the WindowsMDM field.
func (o *Organizationsettings) SetWindowsMDM(v OrganizationsettingsWindowsMDM) {
	o.WindowsMDM = &v
}

func (o Organizationsettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organizationsettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentVersion) {
		toSerialize["agentVersion"] = o.AgentVersion
	}
	if !IsNil(o.BetaFeatures) {
		toSerialize["betaFeatures"] = o.BetaFeatures
	}
	if !IsNil(o.ContactEmail) {
		toSerialize["contactEmail"] = o.ContactEmail
	}
	if !IsNil(o.ContactName) {
		toSerialize["contactName"] = o.ContactName
	}
	if !IsNil(o.DeviceIdentificationEnabled) {
		toSerialize["deviceIdentificationEnabled"] = o.DeviceIdentificationEnabled
	}
	if !IsNil(o.DisableCommandRunner) {
		toSerialize["disableCommandRunner"] = o.DisableCommandRunner
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
	if !IsNil(o.EnableGoogleApps) {
		toSerialize["enableGoogleApps"] = o.EnableGoogleApps
	}
	if !IsNil(o.EnableManagedUID) {
		toSerialize["enableManagedUID"] = o.EnableManagedUID
	}
	if !IsNil(o.EnableO365) {
		toSerialize["enableO365"] = o.EnableO365
	}
	if !IsNil(o.EnableUserPortalAgentInstall) {
		toSerialize["enableUserPortalAgentInstall"] = o.EnableUserPortalAgentInstall
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
	if !IsNil(o.PendingDelete) {
		toSerialize["pendingDelete"] = o.PendingDelete
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
	if !IsNil(o.WindowsMDM) {
		toSerialize["windowsMDM"] = o.WindowsMDM
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Organizationsettings) UnmarshalJSON(bytes []byte) (err error) {
	varOrganizationsettings := _Organizationsettings{}

	if err = json.Unmarshal(bytes, &varOrganizationsettings); err == nil {
		*o = Organizationsettings(varOrganizationsettings)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "agentVersion")
		delete(additionalProperties, "betaFeatures")
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "contactName")
		delete(additionalProperties, "deviceIdentificationEnabled")
		delete(additionalProperties, "disableCommandRunner")
		delete(additionalProperties, "disableGoogleLogin")
		delete(additionalProperties, "disableLdap")
		delete(additionalProperties, "disableUM")
		delete(additionalProperties, "duplicateLDAPGroups")
		delete(additionalProperties, "emailDisclaimer")
		delete(additionalProperties, "enableGoogleApps")
		delete(additionalProperties, "enableManagedUID")
		delete(additionalProperties, "enableO365")
		delete(additionalProperties, "enableUserPortalAgentInstall")
		delete(additionalProperties, "features")
		delete(additionalProperties, "growthData")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "newSystemUserStateDefaults")
		delete(additionalProperties, "passwordCompliance")
		delete(additionalProperties, "passwordPolicy")
		delete(additionalProperties, "pendingDelete")
		delete(additionalProperties, "showIntro")
		delete(additionalProperties, "systemUserPasswordExpirationInDays")
		delete(additionalProperties, "systemUsersCanEdit")
		delete(additionalProperties, "systemUsersCap")
		delete(additionalProperties, "trustedAppConfig")
		delete(additionalProperties, "userPortal")
		delete(additionalProperties, "windowsMDM")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationsettings struct {
	value *Organizationsettings
	isSet bool
}

func (v NullableOrganizationsettings) Get() *Organizationsettings {
	return v.value
}

func (v *NullableOrganizationsettings) Set(val *Organizationsettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsettings(val *Organizationsettings) *NullableOrganizationsettings {
	return &NullableOrganizationsettings{value: val, isSet: true}
}

func (v NullableOrganizationsettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



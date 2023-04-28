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

// checks if the OrganizationsettingsputPasswordPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsettingsputPasswordPolicy{}

// OrganizationsettingsputPasswordPolicy struct for OrganizationsettingsputPasswordPolicy
type OrganizationsettingsputPasswordPolicy struct {
	AllowUsernameSubstring *bool `json:"allowUsernameSubstring,omitempty"`
	// Deprecated field used for the legacy grace period feature.
	DaysAfterExpirationToSelfRecover *int32 `json:"daysAfterExpirationToSelfRecover,omitempty"`
	DaysBeforeExpirationToForceReset *int32 `json:"daysBeforeExpirationToForceReset,omitempty"`
	EffectiveDate *string `json:"effectiveDate,omitempty"`
	EnableDaysAfterExpirationToSelfRecover *bool `json:"enableDaysAfterExpirationToSelfRecover,omitempty"`
	EnableDaysBeforeExpirationToForceReset *bool `json:"enableDaysBeforeExpirationToForceReset,omitempty"`
	EnableLockoutTimeInSeconds *bool `json:"enableLockoutTimeInSeconds,omitempty"`
	EnableMaxHistory *bool `json:"enableMaxHistory,omitempty"`
	EnableMaxLoginAttempts *bool `json:"enableMaxLoginAttempts,omitempty"`
	EnableMinChangePeriodInDays *bool `json:"enableMinChangePeriodInDays,omitempty"`
	EnableMinLength *bool `json:"enableMinLength,omitempty"`
	EnablePasswordExpirationInDays *bool `json:"enablePasswordExpirationInDays,omitempty"`
	GracePeriodDate *string `json:"gracePeriodDate,omitempty"`
	LockoutTimeInSeconds *int32 `json:"lockoutTimeInSeconds,omitempty"`
	MaxHistory *int32 `json:"maxHistory,omitempty"`
	MaxLoginAttempts *int32 `json:"maxLoginAttempts,omitempty"`
	MinChangePeriodInDays *int32 `json:"minChangePeriodInDays,omitempty"`
	MinLength *int32 `json:"minLength,omitempty"`
	NeedsLowercase *bool `json:"needsLowercase,omitempty"`
	NeedsNumeric *bool `json:"needsNumeric,omitempty"`
	NeedsSymbolic *bool `json:"needsSymbolic,omitempty"`
	NeedsUppercase *bool `json:"needsUppercase,omitempty"`
	PasswordExpirationInDays *int32 `json:"passwordExpirationInDays,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationsettingsputPasswordPolicy OrganizationsettingsputPasswordPolicy

// NewOrganizationsettingsputPasswordPolicy instantiates a new OrganizationsettingsputPasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsettingsputPasswordPolicy() *OrganizationsettingsputPasswordPolicy {
	this := OrganizationsettingsputPasswordPolicy{}
	return &this
}

// NewOrganizationsettingsputPasswordPolicyWithDefaults instantiates a new OrganizationsettingsputPasswordPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsettingsputPasswordPolicyWithDefaults() *OrganizationsettingsputPasswordPolicy {
	this := OrganizationsettingsputPasswordPolicy{}
	return &this
}

// GetAllowUsernameSubstring returns the AllowUsernameSubstring field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetAllowUsernameSubstring() bool {
	if o == nil || IsNil(o.AllowUsernameSubstring) {
		var ret bool
		return ret
	}
	return *o.AllowUsernameSubstring
}

// GetAllowUsernameSubstringOk returns a tuple with the AllowUsernameSubstring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetAllowUsernameSubstringOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUsernameSubstring) {
		return nil, false
	}
	return o.AllowUsernameSubstring, true
}

// HasAllowUsernameSubstring returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasAllowUsernameSubstring() bool {
	if o != nil && !IsNil(o.AllowUsernameSubstring) {
		return true
	}

	return false
}

// SetAllowUsernameSubstring gets a reference to the given bool and assigns it to the AllowUsernameSubstring field.
func (o *OrganizationsettingsputPasswordPolicy) SetAllowUsernameSubstring(v bool) {
	o.AllowUsernameSubstring = &v
}

// GetDaysAfterExpirationToSelfRecover returns the DaysAfterExpirationToSelfRecover field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetDaysAfterExpirationToSelfRecover() int32 {
	if o == nil || IsNil(o.DaysAfterExpirationToSelfRecover) {
		var ret int32
		return ret
	}
	return *o.DaysAfterExpirationToSelfRecover
}

// GetDaysAfterExpirationToSelfRecoverOk returns a tuple with the DaysAfterExpirationToSelfRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetDaysAfterExpirationToSelfRecoverOk() (*int32, bool) {
	if o == nil || IsNil(o.DaysAfterExpirationToSelfRecover) {
		return nil, false
	}
	return o.DaysAfterExpirationToSelfRecover, true
}

// HasDaysAfterExpirationToSelfRecover returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasDaysAfterExpirationToSelfRecover() bool {
	if o != nil && !IsNil(o.DaysAfterExpirationToSelfRecover) {
		return true
	}

	return false
}

// SetDaysAfterExpirationToSelfRecover gets a reference to the given int32 and assigns it to the DaysAfterExpirationToSelfRecover field.
func (o *OrganizationsettingsputPasswordPolicy) SetDaysAfterExpirationToSelfRecover(v int32) {
	o.DaysAfterExpirationToSelfRecover = &v
}

// GetDaysBeforeExpirationToForceReset returns the DaysBeforeExpirationToForceReset field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetDaysBeforeExpirationToForceReset() int32 {
	if o == nil || IsNil(o.DaysBeforeExpirationToForceReset) {
		var ret int32
		return ret
	}
	return *o.DaysBeforeExpirationToForceReset
}

// GetDaysBeforeExpirationToForceResetOk returns a tuple with the DaysBeforeExpirationToForceReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetDaysBeforeExpirationToForceResetOk() (*int32, bool) {
	if o == nil || IsNil(o.DaysBeforeExpirationToForceReset) {
		return nil, false
	}
	return o.DaysBeforeExpirationToForceReset, true
}

// HasDaysBeforeExpirationToForceReset returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasDaysBeforeExpirationToForceReset() bool {
	if o != nil && !IsNil(o.DaysBeforeExpirationToForceReset) {
		return true
	}

	return false
}

// SetDaysBeforeExpirationToForceReset gets a reference to the given int32 and assigns it to the DaysBeforeExpirationToForceReset field.
func (o *OrganizationsettingsputPasswordPolicy) SetDaysBeforeExpirationToForceReset(v int32) {
	o.DaysBeforeExpirationToForceReset = &v
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetEffectiveDate() string {
	if o == nil || IsNil(o.EffectiveDate) {
		var ret string
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetEffectiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveDate) {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasEffectiveDate() bool {
	if o != nil && !IsNil(o.EffectiveDate) {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given string and assigns it to the EffectiveDate field.
func (o *OrganizationsettingsputPasswordPolicy) SetEffectiveDate(v string) {
	o.EffectiveDate = &v
}

// GetEnableDaysAfterExpirationToSelfRecover returns the EnableDaysAfterExpirationToSelfRecover field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableDaysAfterExpirationToSelfRecover() bool {
	if o == nil || IsNil(o.EnableDaysAfterExpirationToSelfRecover) {
		var ret bool
		return ret
	}
	return *o.EnableDaysAfterExpirationToSelfRecover
}

// GetEnableDaysAfterExpirationToSelfRecoverOk returns a tuple with the EnableDaysAfterExpirationToSelfRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableDaysAfterExpirationToSelfRecoverOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDaysAfterExpirationToSelfRecover) {
		return nil, false
	}
	return o.EnableDaysAfterExpirationToSelfRecover, true
}

// HasEnableDaysAfterExpirationToSelfRecover returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasEnableDaysAfterExpirationToSelfRecover() bool {
	if o != nil && !IsNil(o.EnableDaysAfterExpirationToSelfRecover) {
		return true
	}

	return false
}

// SetEnableDaysAfterExpirationToSelfRecover gets a reference to the given bool and assigns it to the EnableDaysAfterExpirationToSelfRecover field.
func (o *OrganizationsettingsputPasswordPolicy) SetEnableDaysAfterExpirationToSelfRecover(v bool) {
	o.EnableDaysAfterExpirationToSelfRecover = &v
}

// GetEnableDaysBeforeExpirationToForceReset returns the EnableDaysBeforeExpirationToForceReset field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableDaysBeforeExpirationToForceReset() bool {
	if o == nil || IsNil(o.EnableDaysBeforeExpirationToForceReset) {
		var ret bool
		return ret
	}
	return *o.EnableDaysBeforeExpirationToForceReset
}

// GetEnableDaysBeforeExpirationToForceResetOk returns a tuple with the EnableDaysBeforeExpirationToForceReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableDaysBeforeExpirationToForceResetOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableDaysBeforeExpirationToForceReset) {
		return nil, false
	}
	return o.EnableDaysBeforeExpirationToForceReset, true
}

// HasEnableDaysBeforeExpirationToForceReset returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasEnableDaysBeforeExpirationToForceReset() bool {
	if o != nil && !IsNil(o.EnableDaysBeforeExpirationToForceReset) {
		return true
	}

	return false
}

// SetEnableDaysBeforeExpirationToForceReset gets a reference to the given bool and assigns it to the EnableDaysBeforeExpirationToForceReset field.
func (o *OrganizationsettingsputPasswordPolicy) SetEnableDaysBeforeExpirationToForceReset(v bool) {
	o.EnableDaysBeforeExpirationToForceReset = &v
}

// GetEnableLockoutTimeInSeconds returns the EnableLockoutTimeInSeconds field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableLockoutTimeInSeconds() bool {
	if o == nil || IsNil(o.EnableLockoutTimeInSeconds) {
		var ret bool
		return ret
	}
	return *o.EnableLockoutTimeInSeconds
}

// GetEnableLockoutTimeInSecondsOk returns a tuple with the EnableLockoutTimeInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableLockoutTimeInSecondsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableLockoutTimeInSeconds) {
		return nil, false
	}
	return o.EnableLockoutTimeInSeconds, true
}

// HasEnableLockoutTimeInSeconds returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasEnableLockoutTimeInSeconds() bool {
	if o != nil && !IsNil(o.EnableLockoutTimeInSeconds) {
		return true
	}

	return false
}

// SetEnableLockoutTimeInSeconds gets a reference to the given bool and assigns it to the EnableLockoutTimeInSeconds field.
func (o *OrganizationsettingsputPasswordPolicy) SetEnableLockoutTimeInSeconds(v bool) {
	o.EnableLockoutTimeInSeconds = &v
}

// GetEnableMaxHistory returns the EnableMaxHistory field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableMaxHistory() bool {
	if o == nil || IsNil(o.EnableMaxHistory) {
		var ret bool
		return ret
	}
	return *o.EnableMaxHistory
}

// GetEnableMaxHistoryOk returns a tuple with the EnableMaxHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableMaxHistoryOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMaxHistory) {
		return nil, false
	}
	return o.EnableMaxHistory, true
}

// HasEnableMaxHistory returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasEnableMaxHistory() bool {
	if o != nil && !IsNil(o.EnableMaxHistory) {
		return true
	}

	return false
}

// SetEnableMaxHistory gets a reference to the given bool and assigns it to the EnableMaxHistory field.
func (o *OrganizationsettingsputPasswordPolicy) SetEnableMaxHistory(v bool) {
	o.EnableMaxHistory = &v
}

// GetEnableMaxLoginAttempts returns the EnableMaxLoginAttempts field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableMaxLoginAttempts() bool {
	if o == nil || IsNil(o.EnableMaxLoginAttempts) {
		var ret bool
		return ret
	}
	return *o.EnableMaxLoginAttempts
}

// GetEnableMaxLoginAttemptsOk returns a tuple with the EnableMaxLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableMaxLoginAttemptsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMaxLoginAttempts) {
		return nil, false
	}
	return o.EnableMaxLoginAttempts, true
}

// HasEnableMaxLoginAttempts returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasEnableMaxLoginAttempts() bool {
	if o != nil && !IsNil(o.EnableMaxLoginAttempts) {
		return true
	}

	return false
}

// SetEnableMaxLoginAttempts gets a reference to the given bool and assigns it to the EnableMaxLoginAttempts field.
func (o *OrganizationsettingsputPasswordPolicy) SetEnableMaxLoginAttempts(v bool) {
	o.EnableMaxLoginAttempts = &v
}

// GetEnableMinChangePeriodInDays returns the EnableMinChangePeriodInDays field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableMinChangePeriodInDays() bool {
	if o == nil || IsNil(o.EnableMinChangePeriodInDays) {
		var ret bool
		return ret
	}
	return *o.EnableMinChangePeriodInDays
}

// GetEnableMinChangePeriodInDaysOk returns a tuple with the EnableMinChangePeriodInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableMinChangePeriodInDaysOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMinChangePeriodInDays) {
		return nil, false
	}
	return o.EnableMinChangePeriodInDays, true
}

// HasEnableMinChangePeriodInDays returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasEnableMinChangePeriodInDays() bool {
	if o != nil && !IsNil(o.EnableMinChangePeriodInDays) {
		return true
	}

	return false
}

// SetEnableMinChangePeriodInDays gets a reference to the given bool and assigns it to the EnableMinChangePeriodInDays field.
func (o *OrganizationsettingsputPasswordPolicy) SetEnableMinChangePeriodInDays(v bool) {
	o.EnableMinChangePeriodInDays = &v
}

// GetEnableMinLength returns the EnableMinLength field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableMinLength() bool {
	if o == nil || IsNil(o.EnableMinLength) {
		var ret bool
		return ret
	}
	return *o.EnableMinLength
}

// GetEnableMinLengthOk returns a tuple with the EnableMinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetEnableMinLengthOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMinLength) {
		return nil, false
	}
	return o.EnableMinLength, true
}

// HasEnableMinLength returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasEnableMinLength() bool {
	if o != nil && !IsNil(o.EnableMinLength) {
		return true
	}

	return false
}

// SetEnableMinLength gets a reference to the given bool and assigns it to the EnableMinLength field.
func (o *OrganizationsettingsputPasswordPolicy) SetEnableMinLength(v bool) {
	o.EnableMinLength = &v
}

// GetEnablePasswordExpirationInDays returns the EnablePasswordExpirationInDays field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetEnablePasswordExpirationInDays() bool {
	if o == nil || IsNil(o.EnablePasswordExpirationInDays) {
		var ret bool
		return ret
	}
	return *o.EnablePasswordExpirationInDays
}

// GetEnablePasswordExpirationInDaysOk returns a tuple with the EnablePasswordExpirationInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetEnablePasswordExpirationInDaysOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePasswordExpirationInDays) {
		return nil, false
	}
	return o.EnablePasswordExpirationInDays, true
}

// HasEnablePasswordExpirationInDays returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasEnablePasswordExpirationInDays() bool {
	if o != nil && !IsNil(o.EnablePasswordExpirationInDays) {
		return true
	}

	return false
}

// SetEnablePasswordExpirationInDays gets a reference to the given bool and assigns it to the EnablePasswordExpirationInDays field.
func (o *OrganizationsettingsputPasswordPolicy) SetEnablePasswordExpirationInDays(v bool) {
	o.EnablePasswordExpirationInDays = &v
}

// GetGracePeriodDate returns the GracePeriodDate field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetGracePeriodDate() string {
	if o == nil || IsNil(o.GracePeriodDate) {
		var ret string
		return ret
	}
	return *o.GracePeriodDate
}

// GetGracePeriodDateOk returns a tuple with the GracePeriodDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetGracePeriodDateOk() (*string, bool) {
	if o == nil || IsNil(o.GracePeriodDate) {
		return nil, false
	}
	return o.GracePeriodDate, true
}

// HasGracePeriodDate returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasGracePeriodDate() bool {
	if o != nil && !IsNil(o.GracePeriodDate) {
		return true
	}

	return false
}

// SetGracePeriodDate gets a reference to the given string and assigns it to the GracePeriodDate field.
func (o *OrganizationsettingsputPasswordPolicy) SetGracePeriodDate(v string) {
	o.GracePeriodDate = &v
}

// GetLockoutTimeInSeconds returns the LockoutTimeInSeconds field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetLockoutTimeInSeconds() int32 {
	if o == nil || IsNil(o.LockoutTimeInSeconds) {
		var ret int32
		return ret
	}
	return *o.LockoutTimeInSeconds
}

// GetLockoutTimeInSecondsOk returns a tuple with the LockoutTimeInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetLockoutTimeInSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.LockoutTimeInSeconds) {
		return nil, false
	}
	return o.LockoutTimeInSeconds, true
}

// HasLockoutTimeInSeconds returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasLockoutTimeInSeconds() bool {
	if o != nil && !IsNil(o.LockoutTimeInSeconds) {
		return true
	}

	return false
}

// SetLockoutTimeInSeconds gets a reference to the given int32 and assigns it to the LockoutTimeInSeconds field.
func (o *OrganizationsettingsputPasswordPolicy) SetLockoutTimeInSeconds(v int32) {
	o.LockoutTimeInSeconds = &v
}

// GetMaxHistory returns the MaxHistory field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetMaxHistory() int32 {
	if o == nil || IsNil(o.MaxHistory) {
		var ret int32
		return ret
	}
	return *o.MaxHistory
}

// GetMaxHistoryOk returns a tuple with the MaxHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetMaxHistoryOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxHistory) {
		return nil, false
	}
	return o.MaxHistory, true
}

// HasMaxHistory returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasMaxHistory() bool {
	if o != nil && !IsNil(o.MaxHistory) {
		return true
	}

	return false
}

// SetMaxHistory gets a reference to the given int32 and assigns it to the MaxHistory field.
func (o *OrganizationsettingsputPasswordPolicy) SetMaxHistory(v int32) {
	o.MaxHistory = &v
}

// GetMaxLoginAttempts returns the MaxLoginAttempts field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetMaxLoginAttempts() int32 {
	if o == nil || IsNil(o.MaxLoginAttempts) {
		var ret int32
		return ret
	}
	return *o.MaxLoginAttempts
}

// GetMaxLoginAttemptsOk returns a tuple with the MaxLoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetMaxLoginAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLoginAttempts) {
		return nil, false
	}
	return o.MaxLoginAttempts, true
}

// HasMaxLoginAttempts returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasMaxLoginAttempts() bool {
	if o != nil && !IsNil(o.MaxLoginAttempts) {
		return true
	}

	return false
}

// SetMaxLoginAttempts gets a reference to the given int32 and assigns it to the MaxLoginAttempts field.
func (o *OrganizationsettingsputPasswordPolicy) SetMaxLoginAttempts(v int32) {
	o.MaxLoginAttempts = &v
}

// GetMinChangePeriodInDays returns the MinChangePeriodInDays field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetMinChangePeriodInDays() int32 {
	if o == nil || IsNil(o.MinChangePeriodInDays) {
		var ret int32
		return ret
	}
	return *o.MinChangePeriodInDays
}

// GetMinChangePeriodInDaysOk returns a tuple with the MinChangePeriodInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetMinChangePeriodInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.MinChangePeriodInDays) {
		return nil, false
	}
	return o.MinChangePeriodInDays, true
}

// HasMinChangePeriodInDays returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasMinChangePeriodInDays() bool {
	if o != nil && !IsNil(o.MinChangePeriodInDays) {
		return true
	}

	return false
}

// SetMinChangePeriodInDays gets a reference to the given int32 and assigns it to the MinChangePeriodInDays field.
func (o *OrganizationsettingsputPasswordPolicy) SetMinChangePeriodInDays(v int32) {
	o.MinChangePeriodInDays = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasMinLength() bool {
	if o != nil && !IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *OrganizationsettingsputPasswordPolicy) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetNeedsLowercase returns the NeedsLowercase field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetNeedsLowercase() bool {
	if o == nil || IsNil(o.NeedsLowercase) {
		var ret bool
		return ret
	}
	return *o.NeedsLowercase
}

// GetNeedsLowercaseOk returns a tuple with the NeedsLowercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetNeedsLowercaseOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsLowercase) {
		return nil, false
	}
	return o.NeedsLowercase, true
}

// HasNeedsLowercase returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasNeedsLowercase() bool {
	if o != nil && !IsNil(o.NeedsLowercase) {
		return true
	}

	return false
}

// SetNeedsLowercase gets a reference to the given bool and assigns it to the NeedsLowercase field.
func (o *OrganizationsettingsputPasswordPolicy) SetNeedsLowercase(v bool) {
	o.NeedsLowercase = &v
}

// GetNeedsNumeric returns the NeedsNumeric field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetNeedsNumeric() bool {
	if o == nil || IsNil(o.NeedsNumeric) {
		var ret bool
		return ret
	}
	return *o.NeedsNumeric
}

// GetNeedsNumericOk returns a tuple with the NeedsNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetNeedsNumericOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsNumeric) {
		return nil, false
	}
	return o.NeedsNumeric, true
}

// HasNeedsNumeric returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasNeedsNumeric() bool {
	if o != nil && !IsNil(o.NeedsNumeric) {
		return true
	}

	return false
}

// SetNeedsNumeric gets a reference to the given bool and assigns it to the NeedsNumeric field.
func (o *OrganizationsettingsputPasswordPolicy) SetNeedsNumeric(v bool) {
	o.NeedsNumeric = &v
}

// GetNeedsSymbolic returns the NeedsSymbolic field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetNeedsSymbolic() bool {
	if o == nil || IsNil(o.NeedsSymbolic) {
		var ret bool
		return ret
	}
	return *o.NeedsSymbolic
}

// GetNeedsSymbolicOk returns a tuple with the NeedsSymbolic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetNeedsSymbolicOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsSymbolic) {
		return nil, false
	}
	return o.NeedsSymbolic, true
}

// HasNeedsSymbolic returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasNeedsSymbolic() bool {
	if o != nil && !IsNil(o.NeedsSymbolic) {
		return true
	}

	return false
}

// SetNeedsSymbolic gets a reference to the given bool and assigns it to the NeedsSymbolic field.
func (o *OrganizationsettingsputPasswordPolicy) SetNeedsSymbolic(v bool) {
	o.NeedsSymbolic = &v
}

// GetNeedsUppercase returns the NeedsUppercase field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetNeedsUppercase() bool {
	if o == nil || IsNil(o.NeedsUppercase) {
		var ret bool
		return ret
	}
	return *o.NeedsUppercase
}

// GetNeedsUppercaseOk returns a tuple with the NeedsUppercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetNeedsUppercaseOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsUppercase) {
		return nil, false
	}
	return o.NeedsUppercase, true
}

// HasNeedsUppercase returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasNeedsUppercase() bool {
	if o != nil && !IsNil(o.NeedsUppercase) {
		return true
	}

	return false
}

// SetNeedsUppercase gets a reference to the given bool and assigns it to the NeedsUppercase field.
func (o *OrganizationsettingsputPasswordPolicy) SetNeedsUppercase(v bool) {
	o.NeedsUppercase = &v
}

// GetPasswordExpirationInDays returns the PasswordExpirationInDays field value if set, zero value otherwise.
func (o *OrganizationsettingsputPasswordPolicy) GetPasswordExpirationInDays() int32 {
	if o == nil || IsNil(o.PasswordExpirationInDays) {
		var ret int32
		return ret
	}
	return *o.PasswordExpirationInDays
}

// GetPasswordExpirationInDaysOk returns a tuple with the PasswordExpirationInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsettingsputPasswordPolicy) GetPasswordExpirationInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.PasswordExpirationInDays) {
		return nil, false
	}
	return o.PasswordExpirationInDays, true
}

// HasPasswordExpirationInDays returns a boolean if a field has been set.
func (o *OrganizationsettingsputPasswordPolicy) HasPasswordExpirationInDays() bool {
	if o != nil && !IsNil(o.PasswordExpirationInDays) {
		return true
	}

	return false
}

// SetPasswordExpirationInDays gets a reference to the given int32 and assigns it to the PasswordExpirationInDays field.
func (o *OrganizationsettingsputPasswordPolicy) SetPasswordExpirationInDays(v int32) {
	o.PasswordExpirationInDays = &v
}

func (o OrganizationsettingsputPasswordPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsettingsputPasswordPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowUsernameSubstring) {
		toSerialize["allowUsernameSubstring"] = o.AllowUsernameSubstring
	}
	if !IsNil(o.DaysAfterExpirationToSelfRecover) {
		toSerialize["daysAfterExpirationToSelfRecover"] = o.DaysAfterExpirationToSelfRecover
	}
	if !IsNil(o.DaysBeforeExpirationToForceReset) {
		toSerialize["daysBeforeExpirationToForceReset"] = o.DaysBeforeExpirationToForceReset
	}
	if !IsNil(o.EffectiveDate) {
		toSerialize["effectiveDate"] = o.EffectiveDate
	}
	if !IsNil(o.EnableDaysAfterExpirationToSelfRecover) {
		toSerialize["enableDaysAfterExpirationToSelfRecover"] = o.EnableDaysAfterExpirationToSelfRecover
	}
	if !IsNil(o.EnableDaysBeforeExpirationToForceReset) {
		toSerialize["enableDaysBeforeExpirationToForceReset"] = o.EnableDaysBeforeExpirationToForceReset
	}
	if !IsNil(o.EnableLockoutTimeInSeconds) {
		toSerialize["enableLockoutTimeInSeconds"] = o.EnableLockoutTimeInSeconds
	}
	if !IsNil(o.EnableMaxHistory) {
		toSerialize["enableMaxHistory"] = o.EnableMaxHistory
	}
	if !IsNil(o.EnableMaxLoginAttempts) {
		toSerialize["enableMaxLoginAttempts"] = o.EnableMaxLoginAttempts
	}
	if !IsNil(o.EnableMinChangePeriodInDays) {
		toSerialize["enableMinChangePeriodInDays"] = o.EnableMinChangePeriodInDays
	}
	if !IsNil(o.EnableMinLength) {
		toSerialize["enableMinLength"] = o.EnableMinLength
	}
	if !IsNil(o.EnablePasswordExpirationInDays) {
		toSerialize["enablePasswordExpirationInDays"] = o.EnablePasswordExpirationInDays
	}
	if !IsNil(o.GracePeriodDate) {
		toSerialize["gracePeriodDate"] = o.GracePeriodDate
	}
	if !IsNil(o.LockoutTimeInSeconds) {
		toSerialize["lockoutTimeInSeconds"] = o.LockoutTimeInSeconds
	}
	if !IsNil(o.MaxHistory) {
		toSerialize["maxHistory"] = o.MaxHistory
	}
	if !IsNil(o.MaxLoginAttempts) {
		toSerialize["maxLoginAttempts"] = o.MaxLoginAttempts
	}
	if !IsNil(o.MinChangePeriodInDays) {
		toSerialize["minChangePeriodInDays"] = o.MinChangePeriodInDays
	}
	if !IsNil(o.MinLength) {
		toSerialize["minLength"] = o.MinLength
	}
	if !IsNil(o.NeedsLowercase) {
		toSerialize["needsLowercase"] = o.NeedsLowercase
	}
	if !IsNil(o.NeedsNumeric) {
		toSerialize["needsNumeric"] = o.NeedsNumeric
	}
	if !IsNil(o.NeedsSymbolic) {
		toSerialize["needsSymbolic"] = o.NeedsSymbolic
	}
	if !IsNil(o.NeedsUppercase) {
		toSerialize["needsUppercase"] = o.NeedsUppercase
	}
	if !IsNil(o.PasswordExpirationInDays) {
		toSerialize["passwordExpirationInDays"] = o.PasswordExpirationInDays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationsettingsputPasswordPolicy) UnmarshalJSON(bytes []byte) (err error) {
	varOrganizationsettingsputPasswordPolicy := _OrganizationsettingsputPasswordPolicy{}

	if err = json.Unmarshal(bytes, &varOrganizationsettingsputPasswordPolicy); err == nil {
		*o = OrganizationsettingsputPasswordPolicy(varOrganizationsettingsputPasswordPolicy)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allowUsernameSubstring")
		delete(additionalProperties, "daysAfterExpirationToSelfRecover")
		delete(additionalProperties, "daysBeforeExpirationToForceReset")
		delete(additionalProperties, "effectiveDate")
		delete(additionalProperties, "enableDaysAfterExpirationToSelfRecover")
		delete(additionalProperties, "enableDaysBeforeExpirationToForceReset")
		delete(additionalProperties, "enableLockoutTimeInSeconds")
		delete(additionalProperties, "enableMaxHistory")
		delete(additionalProperties, "enableMaxLoginAttempts")
		delete(additionalProperties, "enableMinChangePeriodInDays")
		delete(additionalProperties, "enableMinLength")
		delete(additionalProperties, "enablePasswordExpirationInDays")
		delete(additionalProperties, "gracePeriodDate")
		delete(additionalProperties, "lockoutTimeInSeconds")
		delete(additionalProperties, "maxHistory")
		delete(additionalProperties, "maxLoginAttempts")
		delete(additionalProperties, "minChangePeriodInDays")
		delete(additionalProperties, "minLength")
		delete(additionalProperties, "needsLowercase")
		delete(additionalProperties, "needsNumeric")
		delete(additionalProperties, "needsSymbolic")
		delete(additionalProperties, "needsUppercase")
		delete(additionalProperties, "passwordExpirationInDays")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationsettingsputPasswordPolicy struct {
	value *OrganizationsettingsputPasswordPolicy
	isSet bool
}

func (v NullableOrganizationsettingsputPasswordPolicy) Get() *OrganizationsettingsputPasswordPolicy {
	return v.value
}

func (v *NullableOrganizationsettingsputPasswordPolicy) Set(val *OrganizationsettingsputPasswordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsettingsputPasswordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsettingsputPasswordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsettingsputPasswordPolicy(val *OrganizationsettingsputPasswordPolicy) *NullableOrganizationsettingsputPasswordPolicy {
	return &NullableOrganizationsettingsputPasswordPolicy{value: val, isSet: true}
}

func (v NullableOrganizationsettingsputPasswordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsettingsputPasswordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



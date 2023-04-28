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

// checks if the Applicationtemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Applicationtemplate{}

// Applicationtemplate struct for Applicationtemplate
type Applicationtemplate struct {
	Id *string `json:"_id,omitempty"`
	Active *bool `json:"active,omitempty"`
	Beta *bool `json:"beta,omitempty"`
	Color *string `json:"color,omitempty"`
	Config *ApplicationConfig `json:"config,omitempty"`
	DisplayLabel *string `json:"displayLabel,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	IsConfigured *bool `json:"isConfigured,omitempty"`
	Jit *ApplicationtemplateJit `json:"jit,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
	LearnMore *string `json:"learnMore,omitempty"`
	Logo *ApplicationtemplateLogo `json:"logo,omitempty"`
	Name *string `json:"name,omitempty"`
	Oidc *ApplicationtemplateOidc `json:"oidc,omitempty"`
	Provision *ApplicationtemplateProvision `json:"provision,omitempty"`
	Sso *Sso `json:"sso,omitempty"`
	SsoUrl *string `json:"ssoUrl,omitempty"`
	Status *string `json:"status,omitempty"`
	Test *string `json:"test,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Applicationtemplate Applicationtemplate

// NewApplicationtemplate instantiates a new Applicationtemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationtemplate() *Applicationtemplate {
	this := Applicationtemplate{}
	return &this
}

// NewApplicationtemplateWithDefaults instantiates a new Applicationtemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationtemplateWithDefaults() *Applicationtemplate {
	this := Applicationtemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Applicationtemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Applicationtemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Applicationtemplate) SetId(v string) {
	o.Id = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Applicationtemplate) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Applicationtemplate) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Applicationtemplate) SetActive(v bool) {
	o.Active = &v
}

// GetBeta returns the Beta field value if set, zero value otherwise.
func (o *Applicationtemplate) GetBeta() bool {
	if o == nil || IsNil(o.Beta) {
		var ret bool
		return ret
	}
	return *o.Beta
}

// GetBetaOk returns a tuple with the Beta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetBetaOk() (*bool, bool) {
	if o == nil || IsNil(o.Beta) {
		return nil, false
	}
	return o.Beta, true
}

// HasBeta returns a boolean if a field has been set.
func (o *Applicationtemplate) HasBeta() bool {
	if o != nil && !IsNil(o.Beta) {
		return true
	}

	return false
}

// SetBeta gets a reference to the given bool and assigns it to the Beta field.
func (o *Applicationtemplate) SetBeta(v bool) {
	o.Beta = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Applicationtemplate) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Applicationtemplate) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Applicationtemplate) SetColor(v string) {
	o.Color = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Applicationtemplate) GetConfig() ApplicationConfig {
	if o == nil || IsNil(o.Config) {
		var ret ApplicationConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetConfigOk() (*ApplicationConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Applicationtemplate) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ApplicationConfig and assigns it to the Config field.
func (o *Applicationtemplate) SetConfig(v ApplicationConfig) {
	o.Config = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *Applicationtemplate) GetDisplayLabel() string {
	if o == nil || IsNil(o.DisplayLabel) {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetDisplayLabelOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayLabel) {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *Applicationtemplate) HasDisplayLabel() bool {
	if o != nil && !IsNil(o.DisplayLabel) {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *Applicationtemplate) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Applicationtemplate) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Applicationtemplate) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Applicationtemplate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIsConfigured returns the IsConfigured field value if set, zero value otherwise.
func (o *Applicationtemplate) GetIsConfigured() bool {
	if o == nil || IsNil(o.IsConfigured) {
		var ret bool
		return ret
	}
	return *o.IsConfigured
}

// GetIsConfiguredOk returns a tuple with the IsConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetIsConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConfigured) {
		return nil, false
	}
	return o.IsConfigured, true
}

// HasIsConfigured returns a boolean if a field has been set.
func (o *Applicationtemplate) HasIsConfigured() bool {
	if o != nil && !IsNil(o.IsConfigured) {
		return true
	}

	return false
}

// SetIsConfigured gets a reference to the given bool and assigns it to the IsConfigured field.
func (o *Applicationtemplate) SetIsConfigured(v bool) {
	o.IsConfigured = &v
}

// GetJit returns the Jit field value if set, zero value otherwise.
func (o *Applicationtemplate) GetJit() ApplicationtemplateJit {
	if o == nil || IsNil(o.Jit) {
		var ret ApplicationtemplateJit
		return ret
	}
	return *o.Jit
}

// GetJitOk returns a tuple with the Jit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetJitOk() (*ApplicationtemplateJit, bool) {
	if o == nil || IsNil(o.Jit) {
		return nil, false
	}
	return o.Jit, true
}

// HasJit returns a boolean if a field has been set.
func (o *Applicationtemplate) HasJit() bool {
	if o != nil && !IsNil(o.Jit) {
		return true
	}

	return false
}

// SetJit gets a reference to the given ApplicationtemplateJit and assigns it to the Jit field.
func (o *Applicationtemplate) SetJit(v ApplicationtemplateJit) {
	o.Jit = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *Applicationtemplate) GetKeywords() []string {
	if o == nil || IsNil(o.Keywords) {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *Applicationtemplate) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []string and assigns it to the Keywords field.
func (o *Applicationtemplate) SetKeywords(v []string) {
	o.Keywords = v
}

// GetLearnMore returns the LearnMore field value if set, zero value otherwise.
func (o *Applicationtemplate) GetLearnMore() string {
	if o == nil || IsNil(o.LearnMore) {
		var ret string
		return ret
	}
	return *o.LearnMore
}

// GetLearnMoreOk returns a tuple with the LearnMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetLearnMoreOk() (*string, bool) {
	if o == nil || IsNil(o.LearnMore) {
		return nil, false
	}
	return o.LearnMore, true
}

// HasLearnMore returns a boolean if a field has been set.
func (o *Applicationtemplate) HasLearnMore() bool {
	if o != nil && !IsNil(o.LearnMore) {
		return true
	}

	return false
}

// SetLearnMore gets a reference to the given string and assigns it to the LearnMore field.
func (o *Applicationtemplate) SetLearnMore(v string) {
	o.LearnMore = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *Applicationtemplate) GetLogo() ApplicationtemplateLogo {
	if o == nil || IsNil(o.Logo) {
		var ret ApplicationtemplateLogo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetLogoOk() (*ApplicationtemplateLogo, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *Applicationtemplate) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given ApplicationtemplateLogo and assigns it to the Logo field.
func (o *Applicationtemplate) SetLogo(v ApplicationtemplateLogo) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Applicationtemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Applicationtemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Applicationtemplate) SetName(v string) {
	o.Name = &v
}

// GetOidc returns the Oidc field value if set, zero value otherwise.
func (o *Applicationtemplate) GetOidc() ApplicationtemplateOidc {
	if o == nil || IsNil(o.Oidc) {
		var ret ApplicationtemplateOidc
		return ret
	}
	return *o.Oidc
}

// GetOidcOk returns a tuple with the Oidc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetOidcOk() (*ApplicationtemplateOidc, bool) {
	if o == nil || IsNil(o.Oidc) {
		return nil, false
	}
	return o.Oidc, true
}

// HasOidc returns a boolean if a field has been set.
func (o *Applicationtemplate) HasOidc() bool {
	if o != nil && !IsNil(o.Oidc) {
		return true
	}

	return false
}

// SetOidc gets a reference to the given ApplicationtemplateOidc and assigns it to the Oidc field.
func (o *Applicationtemplate) SetOidc(v ApplicationtemplateOidc) {
	o.Oidc = &v
}

// GetProvision returns the Provision field value if set, zero value otherwise.
func (o *Applicationtemplate) GetProvision() ApplicationtemplateProvision {
	if o == nil || IsNil(o.Provision) {
		var ret ApplicationtemplateProvision
		return ret
	}
	return *o.Provision
}

// GetProvisionOk returns a tuple with the Provision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetProvisionOk() (*ApplicationtemplateProvision, bool) {
	if o == nil || IsNil(o.Provision) {
		return nil, false
	}
	return o.Provision, true
}

// HasProvision returns a boolean if a field has been set.
func (o *Applicationtemplate) HasProvision() bool {
	if o != nil && !IsNil(o.Provision) {
		return true
	}

	return false
}

// SetProvision gets a reference to the given ApplicationtemplateProvision and assigns it to the Provision field.
func (o *Applicationtemplate) SetProvision(v ApplicationtemplateProvision) {
	o.Provision = &v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *Applicationtemplate) GetSso() Sso {
	if o == nil || IsNil(o.Sso) {
		var ret Sso
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetSsoOk() (*Sso, bool) {
	if o == nil || IsNil(o.Sso) {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *Applicationtemplate) HasSso() bool {
	if o != nil && !IsNil(o.Sso) {
		return true
	}

	return false
}

// SetSso gets a reference to the given Sso and assigns it to the Sso field.
func (o *Applicationtemplate) SetSso(v Sso) {
	o.Sso = &v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise.
func (o *Applicationtemplate) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl) {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SsoUrl) {
		return nil, false
	}
	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *Applicationtemplate) HasSsoUrl() bool {
	if o != nil && !IsNil(o.SsoUrl) {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *Applicationtemplate) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Applicationtemplate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Applicationtemplate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Applicationtemplate) SetStatus(v string) {
	o.Status = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *Applicationtemplate) GetTest() string {
	if o == nil || IsNil(o.Test) {
		var ret string
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicationtemplate) GetTestOk() (*string, bool) {
	if o == nil || IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *Applicationtemplate) HasTest() bool {
	if o != nil && !IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given string and assigns it to the Test field.
func (o *Applicationtemplate) SetTest(v string) {
	o.Test = &v
}

func (o Applicationtemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Applicationtemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Beta) {
		toSerialize["beta"] = o.Beta
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.DisplayLabel) {
		toSerialize["displayLabel"] = o.DisplayLabel
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.IsConfigured) {
		toSerialize["isConfigured"] = o.IsConfigured
	}
	if !IsNil(o.Jit) {
		toSerialize["jit"] = o.Jit
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.LearnMore) {
		toSerialize["learnMore"] = o.LearnMore
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Oidc) {
		toSerialize["oidc"] = o.Oidc
	}
	if !IsNil(o.Provision) {
		toSerialize["provision"] = o.Provision
	}
	if !IsNil(o.Sso) {
		toSerialize["sso"] = o.Sso
	}
	if !IsNil(o.SsoUrl) {
		toSerialize["ssoUrl"] = o.SsoUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Applicationtemplate) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationtemplate := _Applicationtemplate{}

	if err = json.Unmarshal(bytes, &varApplicationtemplate); err == nil {
		*o = Applicationtemplate(varApplicationtemplate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "active")
		delete(additionalProperties, "beta")
		delete(additionalProperties, "color")
		delete(additionalProperties, "config")
		delete(additionalProperties, "displayLabel")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "isConfigured")
		delete(additionalProperties, "jit")
		delete(additionalProperties, "keywords")
		delete(additionalProperties, "learnMore")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "oidc")
		delete(additionalProperties, "provision")
		delete(additionalProperties, "sso")
		delete(additionalProperties, "ssoUrl")
		delete(additionalProperties, "status")
		delete(additionalProperties, "test")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationtemplate struct {
	value *Applicationtemplate
	isSet bool
}

func (v NullableApplicationtemplate) Get() *Applicationtemplate {
	return v.value
}

func (v *NullableApplicationtemplate) Set(val *Applicationtemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationtemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationtemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationtemplate(val *Applicationtemplate) *NullableApplicationtemplate {
	return &NullableApplicationtemplate{value: val, isSet: true}
}

func (v NullableApplicationtemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationtemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



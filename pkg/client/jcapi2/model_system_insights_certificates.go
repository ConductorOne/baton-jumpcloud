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

// checks if the SystemInsightsCertificates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInsightsCertificates{}

// SystemInsightsCertificates struct for SystemInsightsCertificates
type SystemInsightsCertificates struct {
	AuthorityKeyId *string `json:"authority_key_id,omitempty"`
	Ca *int32 `json:"ca,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	KeyAlgorithm *string `json:"key_algorithm,omitempty"`
	KeyStrength *string `json:"key_strength,omitempty"`
	KeyUsage *string `json:"key_usage,omitempty"`
	NotValidAfter *string `json:"not_valid_after,omitempty"`
	NotValidBefore *string `json:"not_valid_before,omitempty"`
	Path *string `json:"path,omitempty"`
	SelfSigned *int32 `json:"self_signed,omitempty"`
	Serial *string `json:"serial,omitempty"`
	Sha1 *string `json:"sha1,omitempty"`
	Sid *string `json:"sid,omitempty"`
	SigningAlgorithm *string `json:"signing_algorithm,omitempty"`
	Store *string `json:"store,omitempty"`
	StoreId *string `json:"store_id,omitempty"`
	StoreLocation *string `json:"store_location,omitempty"`
	Subject *string `json:"subject,omitempty"`
	SubjectKeyId *string `json:"subject_key_id,omitempty"`
	SystemId *string `json:"system_id,omitempty"`
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemInsightsCertificates SystemInsightsCertificates

// NewSystemInsightsCertificates instantiates a new SystemInsightsCertificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInsightsCertificates() *SystemInsightsCertificates {
	this := SystemInsightsCertificates{}
	return &this
}

// NewSystemInsightsCertificatesWithDefaults instantiates a new SystemInsightsCertificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInsightsCertificatesWithDefaults() *SystemInsightsCertificates {
	this := SystemInsightsCertificates{}
	return &this
}

// GetAuthorityKeyId returns the AuthorityKeyId field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetAuthorityKeyId() string {
	if o == nil || IsNil(o.AuthorityKeyId) {
		var ret string
		return ret
	}
	return *o.AuthorityKeyId
}

// GetAuthorityKeyIdOk returns a tuple with the AuthorityKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetAuthorityKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorityKeyId) {
		return nil, false
	}
	return o.AuthorityKeyId, true
}

// HasAuthorityKeyId returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasAuthorityKeyId() bool {
	if o != nil && !IsNil(o.AuthorityKeyId) {
		return true
	}

	return false
}

// SetAuthorityKeyId gets a reference to the given string and assigns it to the AuthorityKeyId field.
func (o *SystemInsightsCertificates) SetAuthorityKeyId(v string) {
	o.AuthorityKeyId = &v
}

// GetCa returns the Ca field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetCa() int32 {
	if o == nil || IsNil(o.Ca) {
		var ret int32
		return ret
	}
	return *o.Ca
}

// GetCaOk returns a tuple with the Ca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetCaOk() (*int32, bool) {
	if o == nil || IsNil(o.Ca) {
		return nil, false
	}
	return o.Ca, true
}

// HasCa returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasCa() bool {
	if o != nil && !IsNil(o.Ca) {
		return true
	}

	return false
}

// SetCa gets a reference to the given int32 and assigns it to the Ca field.
func (o *SystemInsightsCertificates) SetCa(v int32) {
	o.Ca = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetCommonName() string {
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *SystemInsightsCertificates) SetCommonName(v string) {
	o.CommonName = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SystemInsightsCertificates) SetIssuer(v string) {
	o.Issuer = &v
}

// GetKeyAlgorithm returns the KeyAlgorithm field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetKeyAlgorithm() string {
	if o == nil || IsNil(o.KeyAlgorithm) {
		var ret string
		return ret
	}
	return *o.KeyAlgorithm
}

// GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetKeyAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.KeyAlgorithm) {
		return nil, false
	}
	return o.KeyAlgorithm, true
}

// HasKeyAlgorithm returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasKeyAlgorithm() bool {
	if o != nil && !IsNil(o.KeyAlgorithm) {
		return true
	}

	return false
}

// SetKeyAlgorithm gets a reference to the given string and assigns it to the KeyAlgorithm field.
func (o *SystemInsightsCertificates) SetKeyAlgorithm(v string) {
	o.KeyAlgorithm = &v
}

// GetKeyStrength returns the KeyStrength field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetKeyStrength() string {
	if o == nil || IsNil(o.KeyStrength) {
		var ret string
		return ret
	}
	return *o.KeyStrength
}

// GetKeyStrengthOk returns a tuple with the KeyStrength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetKeyStrengthOk() (*string, bool) {
	if o == nil || IsNil(o.KeyStrength) {
		return nil, false
	}
	return o.KeyStrength, true
}

// HasKeyStrength returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasKeyStrength() bool {
	if o != nil && !IsNil(o.KeyStrength) {
		return true
	}

	return false
}

// SetKeyStrength gets a reference to the given string and assigns it to the KeyStrength field.
func (o *SystemInsightsCertificates) SetKeyStrength(v string) {
	o.KeyStrength = &v
}

// GetKeyUsage returns the KeyUsage field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetKeyUsage() string {
	if o == nil || IsNil(o.KeyUsage) {
		var ret string
		return ret
	}
	return *o.KeyUsage
}

// GetKeyUsageOk returns a tuple with the KeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetKeyUsageOk() (*string, bool) {
	if o == nil || IsNil(o.KeyUsage) {
		return nil, false
	}
	return o.KeyUsage, true
}

// HasKeyUsage returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasKeyUsage() bool {
	if o != nil && !IsNil(o.KeyUsage) {
		return true
	}

	return false
}

// SetKeyUsage gets a reference to the given string and assigns it to the KeyUsage field.
func (o *SystemInsightsCertificates) SetKeyUsage(v string) {
	o.KeyUsage = &v
}

// GetNotValidAfter returns the NotValidAfter field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetNotValidAfter() string {
	if o == nil || IsNil(o.NotValidAfter) {
		var ret string
		return ret
	}
	return *o.NotValidAfter
}

// GetNotValidAfterOk returns a tuple with the NotValidAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetNotValidAfterOk() (*string, bool) {
	if o == nil || IsNil(o.NotValidAfter) {
		return nil, false
	}
	return o.NotValidAfter, true
}

// HasNotValidAfter returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasNotValidAfter() bool {
	if o != nil && !IsNil(o.NotValidAfter) {
		return true
	}

	return false
}

// SetNotValidAfter gets a reference to the given string and assigns it to the NotValidAfter field.
func (o *SystemInsightsCertificates) SetNotValidAfter(v string) {
	o.NotValidAfter = &v
}

// GetNotValidBefore returns the NotValidBefore field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetNotValidBefore() string {
	if o == nil || IsNil(o.NotValidBefore) {
		var ret string
		return ret
	}
	return *o.NotValidBefore
}

// GetNotValidBeforeOk returns a tuple with the NotValidBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetNotValidBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.NotValidBefore) {
		return nil, false
	}
	return o.NotValidBefore, true
}

// HasNotValidBefore returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasNotValidBefore() bool {
	if o != nil && !IsNil(o.NotValidBefore) {
		return true
	}

	return false
}

// SetNotValidBefore gets a reference to the given string and assigns it to the NotValidBefore field.
func (o *SystemInsightsCertificates) SetNotValidBefore(v string) {
	o.NotValidBefore = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SystemInsightsCertificates) SetPath(v string) {
	o.Path = &v
}

// GetSelfSigned returns the SelfSigned field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetSelfSigned() int32 {
	if o == nil || IsNil(o.SelfSigned) {
		var ret int32
		return ret
	}
	return *o.SelfSigned
}

// GetSelfSignedOk returns a tuple with the SelfSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetSelfSignedOk() (*int32, bool) {
	if o == nil || IsNil(o.SelfSigned) {
		return nil, false
	}
	return o.SelfSigned, true
}

// HasSelfSigned returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasSelfSigned() bool {
	if o != nil && !IsNil(o.SelfSigned) {
		return true
	}

	return false
}

// SetSelfSigned gets a reference to the given int32 and assigns it to the SelfSigned field.
func (o *SystemInsightsCertificates) SetSelfSigned(v int32) {
	o.SelfSigned = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *SystemInsightsCertificates) SetSerial(v string) {
	o.Serial = &v
}

// GetSha1 returns the Sha1 field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetSha1() string {
	if o == nil || IsNil(o.Sha1) {
		var ret string
		return ret
	}
	return *o.Sha1
}

// GetSha1Ok returns a tuple with the Sha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetSha1Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha1) {
		return nil, false
	}
	return o.Sha1, true
}

// HasSha1 returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasSha1() bool {
	if o != nil && !IsNil(o.Sha1) {
		return true
	}

	return false
}

// SetSha1 gets a reference to the given string and assigns it to the Sha1 field.
func (o *SystemInsightsCertificates) SetSha1(v string) {
	o.Sha1 = &v
}

// GetSid returns the Sid field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetSid() string {
	if o == nil || IsNil(o.Sid) {
		var ret string
		return ret
	}
	return *o.Sid
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetSidOk() (*string, bool) {
	if o == nil || IsNil(o.Sid) {
		return nil, false
	}
	return o.Sid, true
}

// HasSid returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasSid() bool {
	if o != nil && !IsNil(o.Sid) {
		return true
	}

	return false
}

// SetSid gets a reference to the given string and assigns it to the Sid field.
func (o *SystemInsightsCertificates) SetSid(v string) {
	o.Sid = &v
}

// GetSigningAlgorithm returns the SigningAlgorithm field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetSigningAlgorithm() string {
	if o == nil || IsNil(o.SigningAlgorithm) {
		var ret string
		return ret
	}
	return *o.SigningAlgorithm
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SigningAlgorithm) {
		return nil, false
	}
	return o.SigningAlgorithm, true
}

// HasSigningAlgorithm returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasSigningAlgorithm() bool {
	if o != nil && !IsNil(o.SigningAlgorithm) {
		return true
	}

	return false
}

// SetSigningAlgorithm gets a reference to the given string and assigns it to the SigningAlgorithm field.
func (o *SystemInsightsCertificates) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetStore() string {
	if o == nil || IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetStoreOk() (*string, bool) {
	if o == nil || IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasStore() bool {
	if o != nil && !IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *SystemInsightsCertificates) SetStore(v string) {
	o.Store = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *SystemInsightsCertificates) SetStoreId(v string) {
	o.StoreId = &v
}

// GetStoreLocation returns the StoreLocation field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetStoreLocation() string {
	if o == nil || IsNil(o.StoreLocation) {
		var ret string
		return ret
	}
	return *o.StoreLocation
}

// GetStoreLocationOk returns a tuple with the StoreLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetStoreLocationOk() (*string, bool) {
	if o == nil || IsNil(o.StoreLocation) {
		return nil, false
	}
	return o.StoreLocation, true
}

// HasStoreLocation returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasStoreLocation() bool {
	if o != nil && !IsNil(o.StoreLocation) {
		return true
	}

	return false
}

// SetStoreLocation gets a reference to the given string and assigns it to the StoreLocation field.
func (o *SystemInsightsCertificates) SetStoreLocation(v string) {
	o.StoreLocation = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SystemInsightsCertificates) SetSubject(v string) {
	o.Subject = &v
}

// GetSubjectKeyId returns the SubjectKeyId field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetSubjectKeyId() string {
	if o == nil || IsNil(o.SubjectKeyId) {
		var ret string
		return ret
	}
	return *o.SubjectKeyId
}

// GetSubjectKeyIdOk returns a tuple with the SubjectKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetSubjectKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectKeyId) {
		return nil, false
	}
	return o.SubjectKeyId, true
}

// HasSubjectKeyId returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasSubjectKeyId() bool {
	if o != nil && !IsNil(o.SubjectKeyId) {
		return true
	}

	return false
}

// SetSubjectKeyId gets a reference to the given string and assigns it to the SubjectKeyId field.
func (o *SystemInsightsCertificates) SetSubjectKeyId(v string) {
	o.SubjectKeyId = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *SystemInsightsCertificates) SetSystemId(v string) {
	o.SystemId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SystemInsightsCertificates) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsCertificates) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SystemInsightsCertificates) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SystemInsightsCertificates) SetUsername(v string) {
	o.Username = &v
}

func (o SystemInsightsCertificates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInsightsCertificates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorityKeyId) {
		toSerialize["authority_key_id"] = o.AuthorityKeyId
	}
	if !IsNil(o.Ca) {
		toSerialize["ca"] = o.Ca
	}
	if !IsNil(o.CommonName) {
		toSerialize["common_name"] = o.CommonName
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.KeyAlgorithm) {
		toSerialize["key_algorithm"] = o.KeyAlgorithm
	}
	if !IsNil(o.KeyStrength) {
		toSerialize["key_strength"] = o.KeyStrength
	}
	if !IsNil(o.KeyUsage) {
		toSerialize["key_usage"] = o.KeyUsage
	}
	if !IsNil(o.NotValidAfter) {
		toSerialize["not_valid_after"] = o.NotValidAfter
	}
	if !IsNil(o.NotValidBefore) {
		toSerialize["not_valid_before"] = o.NotValidBefore
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.SelfSigned) {
		toSerialize["self_signed"] = o.SelfSigned
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Sha1) {
		toSerialize["sha1"] = o.Sha1
	}
	if !IsNil(o.Sid) {
		toSerialize["sid"] = o.Sid
	}
	if !IsNil(o.SigningAlgorithm) {
		toSerialize["signing_algorithm"] = o.SigningAlgorithm
	}
	if !IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	if !IsNil(o.StoreLocation) {
		toSerialize["store_location"] = o.StoreLocation
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.SubjectKeyId) {
		toSerialize["subject_key_id"] = o.SubjectKeyId
	}
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemInsightsCertificates) UnmarshalJSON(bytes []byte) (err error) {
	varSystemInsightsCertificates := _SystemInsightsCertificates{}

	if err = json.Unmarshal(bytes, &varSystemInsightsCertificates); err == nil {
		*o = SystemInsightsCertificates(varSystemInsightsCertificates)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "authority_key_id")
		delete(additionalProperties, "ca")
		delete(additionalProperties, "common_name")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "key_algorithm")
		delete(additionalProperties, "key_strength")
		delete(additionalProperties, "key_usage")
		delete(additionalProperties, "not_valid_after")
		delete(additionalProperties, "not_valid_before")
		delete(additionalProperties, "path")
		delete(additionalProperties, "self_signed")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "sha1")
		delete(additionalProperties, "sid")
		delete(additionalProperties, "signing_algorithm")
		delete(additionalProperties, "store")
		delete(additionalProperties, "store_id")
		delete(additionalProperties, "store_location")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "subject_key_id")
		delete(additionalProperties, "system_id")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemInsightsCertificates struct {
	value *SystemInsightsCertificates
	isSet bool
}

func (v NullableSystemInsightsCertificates) Get() *SystemInsightsCertificates {
	return v.value
}

func (v *NullableSystemInsightsCertificates) Set(val *SystemInsightsCertificates) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInsightsCertificates) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInsightsCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInsightsCertificates(val *SystemInsightsCertificates) *NullableSystemInsightsCertificates {
	return &NullableSystemInsightsCertificates{value: val, isSet: true}
}

func (v NullableSystemInsightsCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInsightsCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



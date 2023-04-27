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

// checks if the RadiusServersPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RadiusServersPutRequest{}

// RadiusServersPutRequest struct for RadiusServersPutRequest
type RadiusServersPutRequest struct {
	CaCert *string `json:"caCert,omitempty"`
	DeviceCertEnabled *bool `json:"deviceCertEnabled,omitempty"`
	Mfa *string `json:"mfa,omitempty"`
	Name string `json:"name"`
	NetworkSourceIp string `json:"networkSourceIp"`
	SharedSecret string `json:"sharedSecret"`
	Tags []string `json:"tags,omitempty"`
	UserCertEnabled *bool `json:"userCertEnabled,omitempty"`
	UserLockoutAction *string `json:"userLockoutAction,omitempty"`
	UserPasswordEnabled *bool `json:"userPasswordEnabled,omitempty"`
	UserPasswordExpirationAction *string `json:"userPasswordExpirationAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RadiusServersPutRequest RadiusServersPutRequest

// NewRadiusServersPutRequest instantiates a new RadiusServersPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusServersPutRequest(name string, networkSourceIp string, sharedSecret string) *RadiusServersPutRequest {
	this := RadiusServersPutRequest{}
	this.Name = name
	this.NetworkSourceIp = networkSourceIp
	this.SharedSecret = sharedSecret
	return &this
}

// NewRadiusServersPutRequestWithDefaults instantiates a new RadiusServersPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusServersPutRequestWithDefaults() *RadiusServersPutRequest {
	this := RadiusServersPutRequest{}
	return &this
}

// GetCaCert returns the CaCert field value if set, zero value otherwise.
func (o *RadiusServersPutRequest) GetCaCert() string {
	if o == nil || IsNil(o.CaCert) {
		var ret string
		return ret
	}
	return *o.CaCert
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetCaCertOk() (*string, bool) {
	if o == nil || IsNil(o.CaCert) {
		return nil, false
	}
	return o.CaCert, true
}

// HasCaCert returns a boolean if a field has been set.
func (o *RadiusServersPutRequest) HasCaCert() bool {
	if o != nil && !IsNil(o.CaCert) {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given string and assigns it to the CaCert field.
func (o *RadiusServersPutRequest) SetCaCert(v string) {
	o.CaCert = &v
}

// GetDeviceCertEnabled returns the DeviceCertEnabled field value if set, zero value otherwise.
func (o *RadiusServersPutRequest) GetDeviceCertEnabled() bool {
	if o == nil || IsNil(o.DeviceCertEnabled) {
		var ret bool
		return ret
	}
	return *o.DeviceCertEnabled
}

// GetDeviceCertEnabledOk returns a tuple with the DeviceCertEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetDeviceCertEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DeviceCertEnabled) {
		return nil, false
	}
	return o.DeviceCertEnabled, true
}

// HasDeviceCertEnabled returns a boolean if a field has been set.
func (o *RadiusServersPutRequest) HasDeviceCertEnabled() bool {
	if o != nil && !IsNil(o.DeviceCertEnabled) {
		return true
	}

	return false
}

// SetDeviceCertEnabled gets a reference to the given bool and assigns it to the DeviceCertEnabled field.
func (o *RadiusServersPutRequest) SetDeviceCertEnabled(v bool) {
	o.DeviceCertEnabled = &v
}

// GetMfa returns the Mfa field value if set, zero value otherwise.
func (o *RadiusServersPutRequest) GetMfa() string {
	if o == nil || IsNil(o.Mfa) {
		var ret string
		return ret
	}
	return *o.Mfa
}

// GetMfaOk returns a tuple with the Mfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetMfaOk() (*string, bool) {
	if o == nil || IsNil(o.Mfa) {
		return nil, false
	}
	return o.Mfa, true
}

// HasMfa returns a boolean if a field has been set.
func (o *RadiusServersPutRequest) HasMfa() bool {
	if o != nil && !IsNil(o.Mfa) {
		return true
	}

	return false
}

// SetMfa gets a reference to the given string and assigns it to the Mfa field.
func (o *RadiusServersPutRequest) SetMfa(v string) {
	o.Mfa = &v
}

// GetName returns the Name field value
func (o *RadiusServersPutRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RadiusServersPutRequest) SetName(v string) {
	o.Name = v
}

// GetNetworkSourceIp returns the NetworkSourceIp field value
func (o *RadiusServersPutRequest) GetNetworkSourceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkSourceIp
}

// GetNetworkSourceIpOk returns a tuple with the NetworkSourceIp field value
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetNetworkSourceIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkSourceIp, true
}

// SetNetworkSourceIp sets field value
func (o *RadiusServersPutRequest) SetNetworkSourceIp(v string) {
	o.NetworkSourceIp = v
}

// GetSharedSecret returns the SharedSecret field value
func (o *RadiusServersPutRequest) GetSharedSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetSharedSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SharedSecret, true
}

// SetSharedSecret sets field value
func (o *RadiusServersPutRequest) SetSharedSecret(v string) {
	o.SharedSecret = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RadiusServersPutRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RadiusServersPutRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RadiusServersPutRequest) SetTags(v []string) {
	o.Tags = v
}

// GetUserCertEnabled returns the UserCertEnabled field value if set, zero value otherwise.
func (o *RadiusServersPutRequest) GetUserCertEnabled() bool {
	if o == nil || IsNil(o.UserCertEnabled) {
		var ret bool
		return ret
	}
	return *o.UserCertEnabled
}

// GetUserCertEnabledOk returns a tuple with the UserCertEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetUserCertEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UserCertEnabled) {
		return nil, false
	}
	return o.UserCertEnabled, true
}

// HasUserCertEnabled returns a boolean if a field has been set.
func (o *RadiusServersPutRequest) HasUserCertEnabled() bool {
	if o != nil && !IsNil(o.UserCertEnabled) {
		return true
	}

	return false
}

// SetUserCertEnabled gets a reference to the given bool and assigns it to the UserCertEnabled field.
func (o *RadiusServersPutRequest) SetUserCertEnabled(v bool) {
	o.UserCertEnabled = &v
}

// GetUserLockoutAction returns the UserLockoutAction field value if set, zero value otherwise.
func (o *RadiusServersPutRequest) GetUserLockoutAction() string {
	if o == nil || IsNil(o.UserLockoutAction) {
		var ret string
		return ret
	}
	return *o.UserLockoutAction
}

// GetUserLockoutActionOk returns a tuple with the UserLockoutAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetUserLockoutActionOk() (*string, bool) {
	if o == nil || IsNil(o.UserLockoutAction) {
		return nil, false
	}
	return o.UserLockoutAction, true
}

// HasUserLockoutAction returns a boolean if a field has been set.
func (o *RadiusServersPutRequest) HasUserLockoutAction() bool {
	if o != nil && !IsNil(o.UserLockoutAction) {
		return true
	}

	return false
}

// SetUserLockoutAction gets a reference to the given string and assigns it to the UserLockoutAction field.
func (o *RadiusServersPutRequest) SetUserLockoutAction(v string) {
	o.UserLockoutAction = &v
}

// GetUserPasswordEnabled returns the UserPasswordEnabled field value if set, zero value otherwise.
func (o *RadiusServersPutRequest) GetUserPasswordEnabled() bool {
	if o == nil || IsNil(o.UserPasswordEnabled) {
		var ret bool
		return ret
	}
	return *o.UserPasswordEnabled
}

// GetUserPasswordEnabledOk returns a tuple with the UserPasswordEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetUserPasswordEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UserPasswordEnabled) {
		return nil, false
	}
	return o.UserPasswordEnabled, true
}

// HasUserPasswordEnabled returns a boolean if a field has been set.
func (o *RadiusServersPutRequest) HasUserPasswordEnabled() bool {
	if o != nil && !IsNil(o.UserPasswordEnabled) {
		return true
	}

	return false
}

// SetUserPasswordEnabled gets a reference to the given bool and assigns it to the UserPasswordEnabled field.
func (o *RadiusServersPutRequest) SetUserPasswordEnabled(v bool) {
	o.UserPasswordEnabled = &v
}

// GetUserPasswordExpirationAction returns the UserPasswordExpirationAction field value if set, zero value otherwise.
func (o *RadiusServersPutRequest) GetUserPasswordExpirationAction() string {
	if o == nil || IsNil(o.UserPasswordExpirationAction) {
		var ret string
		return ret
	}
	return *o.UserPasswordExpirationAction
}

// GetUserPasswordExpirationActionOk returns a tuple with the UserPasswordExpirationAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusServersPutRequest) GetUserPasswordExpirationActionOk() (*string, bool) {
	if o == nil || IsNil(o.UserPasswordExpirationAction) {
		return nil, false
	}
	return o.UserPasswordExpirationAction, true
}

// HasUserPasswordExpirationAction returns a boolean if a field has been set.
func (o *RadiusServersPutRequest) HasUserPasswordExpirationAction() bool {
	if o != nil && !IsNil(o.UserPasswordExpirationAction) {
		return true
	}

	return false
}

// SetUserPasswordExpirationAction gets a reference to the given string and assigns it to the UserPasswordExpirationAction field.
func (o *RadiusServersPutRequest) SetUserPasswordExpirationAction(v string) {
	o.UserPasswordExpirationAction = &v
}

func (o RadiusServersPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RadiusServersPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaCert) {
		toSerialize["caCert"] = o.CaCert
	}
	if !IsNil(o.DeviceCertEnabled) {
		toSerialize["deviceCertEnabled"] = o.DeviceCertEnabled
	}
	if !IsNil(o.Mfa) {
		toSerialize["mfa"] = o.Mfa
	}
	toSerialize["name"] = o.Name
	toSerialize["networkSourceIp"] = o.NetworkSourceIp
	toSerialize["sharedSecret"] = o.SharedSecret
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UserCertEnabled) {
		toSerialize["userCertEnabled"] = o.UserCertEnabled
	}
	if !IsNil(o.UserLockoutAction) {
		toSerialize["userLockoutAction"] = o.UserLockoutAction
	}
	if !IsNil(o.UserPasswordEnabled) {
		toSerialize["userPasswordEnabled"] = o.UserPasswordEnabled
	}
	if !IsNil(o.UserPasswordExpirationAction) {
		toSerialize["userPasswordExpirationAction"] = o.UserPasswordExpirationAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RadiusServersPutRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRadiusServersPutRequest := _RadiusServersPutRequest{}

	if err = json.Unmarshal(bytes, &varRadiusServersPutRequest); err == nil {
		*o = RadiusServersPutRequest(varRadiusServersPutRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "caCert")
		delete(additionalProperties, "deviceCertEnabled")
		delete(additionalProperties, "mfa")
		delete(additionalProperties, "name")
		delete(additionalProperties, "networkSourceIp")
		delete(additionalProperties, "sharedSecret")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "userCertEnabled")
		delete(additionalProperties, "userLockoutAction")
		delete(additionalProperties, "userPasswordEnabled")
		delete(additionalProperties, "userPasswordExpirationAction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRadiusServersPutRequest struct {
	value *RadiusServersPutRequest
	isSet bool
}

func (v NullableRadiusServersPutRequest) Get() *RadiusServersPutRequest {
	return v.value
}

func (v *NullableRadiusServersPutRequest) Set(val *RadiusServersPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusServersPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusServersPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusServersPutRequest(val *RadiusServersPutRequest) *NullableRadiusServersPutRequest {
	return &NullableRadiusServersPutRequest{value: val, isSet: true}
}

func (v NullableRadiusServersPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusServersPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



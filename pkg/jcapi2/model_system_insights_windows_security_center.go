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

// checks if the SystemInsightsWindowsSecurityCenter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInsightsWindowsSecurityCenter{}

// SystemInsightsWindowsSecurityCenter struct for SystemInsightsWindowsSecurityCenter
type SystemInsightsWindowsSecurityCenter struct {
	Antispyware *string `json:"antispyware,omitempty"`
	Antivirus *string `json:"antivirus,omitempty"`
	Autoupdate *string `json:"autoupdate,omitempty"`
	CollectionTime *string `json:"collection_time,omitempty"`
	Firewall *string `json:"firewall,omitempty"`
	InternetSettings *string `json:"internet_settings,omitempty"`
	SystemId *string `json:"system_id,omitempty"`
	UserAccountControl *string `json:"user_account_control,omitempty"`
	WindowsSecurityCenterService *string `json:"windows_security_center_service,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemInsightsWindowsSecurityCenter SystemInsightsWindowsSecurityCenter

// NewSystemInsightsWindowsSecurityCenter instantiates a new SystemInsightsWindowsSecurityCenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInsightsWindowsSecurityCenter() *SystemInsightsWindowsSecurityCenter {
	this := SystemInsightsWindowsSecurityCenter{}
	return &this
}

// NewSystemInsightsWindowsSecurityCenterWithDefaults instantiates a new SystemInsightsWindowsSecurityCenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInsightsWindowsSecurityCenterWithDefaults() *SystemInsightsWindowsSecurityCenter {
	this := SystemInsightsWindowsSecurityCenter{}
	return &this
}

// GetAntispyware returns the Antispyware field value if set, zero value otherwise.
func (o *SystemInsightsWindowsSecurityCenter) GetAntispyware() string {
	if o == nil || IsNil(o.Antispyware) {
		var ret string
		return ret
	}
	return *o.Antispyware
}

// GetAntispywareOk returns a tuple with the Antispyware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsWindowsSecurityCenter) GetAntispywareOk() (*string, bool) {
	if o == nil || IsNil(o.Antispyware) {
		return nil, false
	}
	return o.Antispyware, true
}

// HasAntispyware returns a boolean if a field has been set.
func (o *SystemInsightsWindowsSecurityCenter) HasAntispyware() bool {
	if o != nil && !IsNil(o.Antispyware) {
		return true
	}

	return false
}

// SetAntispyware gets a reference to the given string and assigns it to the Antispyware field.
func (o *SystemInsightsWindowsSecurityCenter) SetAntispyware(v string) {
	o.Antispyware = &v
}

// GetAntivirus returns the Antivirus field value if set, zero value otherwise.
func (o *SystemInsightsWindowsSecurityCenter) GetAntivirus() string {
	if o == nil || IsNil(o.Antivirus) {
		var ret string
		return ret
	}
	return *o.Antivirus
}

// GetAntivirusOk returns a tuple with the Antivirus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsWindowsSecurityCenter) GetAntivirusOk() (*string, bool) {
	if o == nil || IsNil(o.Antivirus) {
		return nil, false
	}
	return o.Antivirus, true
}

// HasAntivirus returns a boolean if a field has been set.
func (o *SystemInsightsWindowsSecurityCenter) HasAntivirus() bool {
	if o != nil && !IsNil(o.Antivirus) {
		return true
	}

	return false
}

// SetAntivirus gets a reference to the given string and assigns it to the Antivirus field.
func (o *SystemInsightsWindowsSecurityCenter) SetAntivirus(v string) {
	o.Antivirus = &v
}

// GetAutoupdate returns the Autoupdate field value if set, zero value otherwise.
func (o *SystemInsightsWindowsSecurityCenter) GetAutoupdate() string {
	if o == nil || IsNil(o.Autoupdate) {
		var ret string
		return ret
	}
	return *o.Autoupdate
}

// GetAutoupdateOk returns a tuple with the Autoupdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsWindowsSecurityCenter) GetAutoupdateOk() (*string, bool) {
	if o == nil || IsNil(o.Autoupdate) {
		return nil, false
	}
	return o.Autoupdate, true
}

// HasAutoupdate returns a boolean if a field has been set.
func (o *SystemInsightsWindowsSecurityCenter) HasAutoupdate() bool {
	if o != nil && !IsNil(o.Autoupdate) {
		return true
	}

	return false
}

// SetAutoupdate gets a reference to the given string and assigns it to the Autoupdate field.
func (o *SystemInsightsWindowsSecurityCenter) SetAutoupdate(v string) {
	o.Autoupdate = &v
}

// GetCollectionTime returns the CollectionTime field value if set, zero value otherwise.
func (o *SystemInsightsWindowsSecurityCenter) GetCollectionTime() string {
	if o == nil || IsNil(o.CollectionTime) {
		var ret string
		return ret
	}
	return *o.CollectionTime
}

// GetCollectionTimeOk returns a tuple with the CollectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsWindowsSecurityCenter) GetCollectionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionTime) {
		return nil, false
	}
	return o.CollectionTime, true
}

// HasCollectionTime returns a boolean if a field has been set.
func (o *SystemInsightsWindowsSecurityCenter) HasCollectionTime() bool {
	if o != nil && !IsNil(o.CollectionTime) {
		return true
	}

	return false
}

// SetCollectionTime gets a reference to the given string and assigns it to the CollectionTime field.
func (o *SystemInsightsWindowsSecurityCenter) SetCollectionTime(v string) {
	o.CollectionTime = &v
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *SystemInsightsWindowsSecurityCenter) GetFirewall() string {
	if o == nil || IsNil(o.Firewall) {
		var ret string
		return ret
	}
	return *o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsWindowsSecurityCenter) GetFirewallOk() (*string, bool) {
	if o == nil || IsNil(o.Firewall) {
		return nil, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *SystemInsightsWindowsSecurityCenter) HasFirewall() bool {
	if o != nil && !IsNil(o.Firewall) {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given string and assigns it to the Firewall field.
func (o *SystemInsightsWindowsSecurityCenter) SetFirewall(v string) {
	o.Firewall = &v
}

// GetInternetSettings returns the InternetSettings field value if set, zero value otherwise.
func (o *SystemInsightsWindowsSecurityCenter) GetInternetSettings() string {
	if o == nil || IsNil(o.InternetSettings) {
		var ret string
		return ret
	}
	return *o.InternetSettings
}

// GetInternetSettingsOk returns a tuple with the InternetSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsWindowsSecurityCenter) GetInternetSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.InternetSettings) {
		return nil, false
	}
	return o.InternetSettings, true
}

// HasInternetSettings returns a boolean if a field has been set.
func (o *SystemInsightsWindowsSecurityCenter) HasInternetSettings() bool {
	if o != nil && !IsNil(o.InternetSettings) {
		return true
	}

	return false
}

// SetInternetSettings gets a reference to the given string and assigns it to the InternetSettings field.
func (o *SystemInsightsWindowsSecurityCenter) SetInternetSettings(v string) {
	o.InternetSettings = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *SystemInsightsWindowsSecurityCenter) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsWindowsSecurityCenter) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *SystemInsightsWindowsSecurityCenter) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *SystemInsightsWindowsSecurityCenter) SetSystemId(v string) {
	o.SystemId = &v
}

// GetUserAccountControl returns the UserAccountControl field value if set, zero value otherwise.
func (o *SystemInsightsWindowsSecurityCenter) GetUserAccountControl() string {
	if o == nil || IsNil(o.UserAccountControl) {
		var ret string
		return ret
	}
	return *o.UserAccountControl
}

// GetUserAccountControlOk returns a tuple with the UserAccountControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsWindowsSecurityCenter) GetUserAccountControlOk() (*string, bool) {
	if o == nil || IsNil(o.UserAccountControl) {
		return nil, false
	}
	return o.UserAccountControl, true
}

// HasUserAccountControl returns a boolean if a field has been set.
func (o *SystemInsightsWindowsSecurityCenter) HasUserAccountControl() bool {
	if o != nil && !IsNil(o.UserAccountControl) {
		return true
	}

	return false
}

// SetUserAccountControl gets a reference to the given string and assigns it to the UserAccountControl field.
func (o *SystemInsightsWindowsSecurityCenter) SetUserAccountControl(v string) {
	o.UserAccountControl = &v
}

// GetWindowsSecurityCenterService returns the WindowsSecurityCenterService field value if set, zero value otherwise.
func (o *SystemInsightsWindowsSecurityCenter) GetWindowsSecurityCenterService() string {
	if o == nil || IsNil(o.WindowsSecurityCenterService) {
		var ret string
		return ret
	}
	return *o.WindowsSecurityCenterService
}

// GetWindowsSecurityCenterServiceOk returns a tuple with the WindowsSecurityCenterService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsWindowsSecurityCenter) GetWindowsSecurityCenterServiceOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsSecurityCenterService) {
		return nil, false
	}
	return o.WindowsSecurityCenterService, true
}

// HasWindowsSecurityCenterService returns a boolean if a field has been set.
func (o *SystemInsightsWindowsSecurityCenter) HasWindowsSecurityCenterService() bool {
	if o != nil && !IsNil(o.WindowsSecurityCenterService) {
		return true
	}

	return false
}

// SetWindowsSecurityCenterService gets a reference to the given string and assigns it to the WindowsSecurityCenterService field.
func (o *SystemInsightsWindowsSecurityCenter) SetWindowsSecurityCenterService(v string) {
	o.WindowsSecurityCenterService = &v
}

func (o SystemInsightsWindowsSecurityCenter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInsightsWindowsSecurityCenter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Antispyware) {
		toSerialize["antispyware"] = o.Antispyware
	}
	if !IsNil(o.Antivirus) {
		toSerialize["antivirus"] = o.Antivirus
	}
	if !IsNil(o.Autoupdate) {
		toSerialize["autoupdate"] = o.Autoupdate
	}
	if !IsNil(o.CollectionTime) {
		toSerialize["collection_time"] = o.CollectionTime
	}
	if !IsNil(o.Firewall) {
		toSerialize["firewall"] = o.Firewall
	}
	if !IsNil(o.InternetSettings) {
		toSerialize["internet_settings"] = o.InternetSettings
	}
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}
	if !IsNil(o.UserAccountControl) {
		toSerialize["user_account_control"] = o.UserAccountControl
	}
	if !IsNil(o.WindowsSecurityCenterService) {
		toSerialize["windows_security_center_service"] = o.WindowsSecurityCenterService
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemInsightsWindowsSecurityCenter) UnmarshalJSON(bytes []byte) (err error) {
	varSystemInsightsWindowsSecurityCenter := _SystemInsightsWindowsSecurityCenter{}

	if err = json.Unmarshal(bytes, &varSystemInsightsWindowsSecurityCenter); err == nil {
		*o = SystemInsightsWindowsSecurityCenter(varSystemInsightsWindowsSecurityCenter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "antispyware")
		delete(additionalProperties, "antivirus")
		delete(additionalProperties, "autoupdate")
		delete(additionalProperties, "collection_time")
		delete(additionalProperties, "firewall")
		delete(additionalProperties, "internet_settings")
		delete(additionalProperties, "system_id")
		delete(additionalProperties, "user_account_control")
		delete(additionalProperties, "windows_security_center_service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemInsightsWindowsSecurityCenter struct {
	value *SystemInsightsWindowsSecurityCenter
	isSet bool
}

func (v NullableSystemInsightsWindowsSecurityCenter) Get() *SystemInsightsWindowsSecurityCenter {
	return v.value
}

func (v *NullableSystemInsightsWindowsSecurityCenter) Set(val *SystemInsightsWindowsSecurityCenter) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInsightsWindowsSecurityCenter) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInsightsWindowsSecurityCenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInsightsWindowsSecurityCenter(val *SystemInsightsWindowsSecurityCenter) *NullableSystemInsightsWindowsSecurityCenter {
	return &NullableSystemInsightsWindowsSecurityCenter{value: val, isSet: true}
}

func (v NullableSystemInsightsWindowsSecurityCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInsightsWindowsSecurityCenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



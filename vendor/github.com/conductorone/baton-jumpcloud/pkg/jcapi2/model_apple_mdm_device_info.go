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

// checks if the AppleMdmDeviceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppleMdmDeviceInfo{}

// AppleMdmDeviceInfo Apple MDM device information
type AppleMdmDeviceInfo struct {
	ActivationLockAllowedWhileSupervised *bool `json:"activationLockAllowedWhileSupervised,omitempty"`
	AvailableDeviceCapacity *float32 `json:"availableDeviceCapacity,omitempty"`
	DeviceCapacity *float32 `json:"deviceCapacity,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	Iccid *string `json:"iccid,omitempty"`
	Imei *string `json:"imei,omitempty"`
	IsSupervised *bool `json:"isSupervised,omitempty"`
	ModelName *string `json:"modelName,omitempty"`
	SecondIccid *string `json:"secondIccid,omitempty"`
	SecondImei *string `json:"secondImei,omitempty"`
	SecondSubscriberCarrierNetwork *string `json:"secondSubscriberCarrierNetwork,omitempty"`
	SubscriberCarrierNetwork *string `json:"subscriberCarrierNetwork,omitempty"`
	WifiMac *string `json:"wifiMac,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppleMdmDeviceInfo AppleMdmDeviceInfo

// NewAppleMdmDeviceInfo instantiates a new AppleMdmDeviceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppleMdmDeviceInfo() *AppleMdmDeviceInfo {
	this := AppleMdmDeviceInfo{}
	return &this
}

// NewAppleMdmDeviceInfoWithDefaults instantiates a new AppleMdmDeviceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleMdmDeviceInfoWithDefaults() *AppleMdmDeviceInfo {
	this := AppleMdmDeviceInfo{}
	return &this
}

// GetActivationLockAllowedWhileSupervised returns the ActivationLockAllowedWhileSupervised field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetActivationLockAllowedWhileSupervised() bool {
	if o == nil || IsNil(o.ActivationLockAllowedWhileSupervised) {
		var ret bool
		return ret
	}
	return *o.ActivationLockAllowedWhileSupervised
}

// GetActivationLockAllowedWhileSupervisedOk returns a tuple with the ActivationLockAllowedWhileSupervised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetActivationLockAllowedWhileSupervisedOk() (*bool, bool) {
	if o == nil || IsNil(o.ActivationLockAllowedWhileSupervised) {
		return nil, false
	}
	return o.ActivationLockAllowedWhileSupervised, true
}

// HasActivationLockAllowedWhileSupervised returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasActivationLockAllowedWhileSupervised() bool {
	if o != nil && !IsNil(o.ActivationLockAllowedWhileSupervised) {
		return true
	}

	return false
}

// SetActivationLockAllowedWhileSupervised gets a reference to the given bool and assigns it to the ActivationLockAllowedWhileSupervised field.
func (o *AppleMdmDeviceInfo) SetActivationLockAllowedWhileSupervised(v bool) {
	o.ActivationLockAllowedWhileSupervised = &v
}

// GetAvailableDeviceCapacity returns the AvailableDeviceCapacity field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetAvailableDeviceCapacity() float32 {
	if o == nil || IsNil(o.AvailableDeviceCapacity) {
		var ret float32
		return ret
	}
	return *o.AvailableDeviceCapacity
}

// GetAvailableDeviceCapacityOk returns a tuple with the AvailableDeviceCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetAvailableDeviceCapacityOk() (*float32, bool) {
	if o == nil || IsNil(o.AvailableDeviceCapacity) {
		return nil, false
	}
	return o.AvailableDeviceCapacity, true
}

// HasAvailableDeviceCapacity returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasAvailableDeviceCapacity() bool {
	if o != nil && !IsNil(o.AvailableDeviceCapacity) {
		return true
	}

	return false
}

// SetAvailableDeviceCapacity gets a reference to the given float32 and assigns it to the AvailableDeviceCapacity field.
func (o *AppleMdmDeviceInfo) SetAvailableDeviceCapacity(v float32) {
	o.AvailableDeviceCapacity = &v
}

// GetDeviceCapacity returns the DeviceCapacity field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetDeviceCapacity() float32 {
	if o == nil || IsNil(o.DeviceCapacity) {
		var ret float32
		return ret
	}
	return *o.DeviceCapacity
}

// GetDeviceCapacityOk returns a tuple with the DeviceCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetDeviceCapacityOk() (*float32, bool) {
	if o == nil || IsNil(o.DeviceCapacity) {
		return nil, false
	}
	return o.DeviceCapacity, true
}

// HasDeviceCapacity returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasDeviceCapacity() bool {
	if o != nil && !IsNil(o.DeviceCapacity) {
		return true
	}

	return false
}

// SetDeviceCapacity gets a reference to the given float32 and assigns it to the DeviceCapacity field.
func (o *AppleMdmDeviceInfo) SetDeviceCapacity(v float32) {
	o.DeviceCapacity = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *AppleMdmDeviceInfo) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetIccid returns the Iccid field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetIccid() string {
	if o == nil || IsNil(o.Iccid) {
		var ret string
		return ret
	}
	return *o.Iccid
}

// GetIccidOk returns a tuple with the Iccid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetIccidOk() (*string, bool) {
	if o == nil || IsNil(o.Iccid) {
		return nil, false
	}
	return o.Iccid, true
}

// HasIccid returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasIccid() bool {
	if o != nil && !IsNil(o.Iccid) {
		return true
	}

	return false
}

// SetIccid gets a reference to the given string and assigns it to the Iccid field.
func (o *AppleMdmDeviceInfo) SetIccid(v string) {
	o.Iccid = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetImei() string {
	if o == nil || IsNil(o.Imei) {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetImeiOk() (*string, bool) {
	if o == nil || IsNil(o.Imei) {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasImei() bool {
	if o != nil && !IsNil(o.Imei) {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *AppleMdmDeviceInfo) SetImei(v string) {
	o.Imei = &v
}

// GetIsSupervised returns the IsSupervised field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetIsSupervised() bool {
	if o == nil || IsNil(o.IsSupervised) {
		var ret bool
		return ret
	}
	return *o.IsSupervised
}

// GetIsSupervisedOk returns a tuple with the IsSupervised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetIsSupervisedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSupervised) {
		return nil, false
	}
	return o.IsSupervised, true
}

// HasIsSupervised returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasIsSupervised() bool {
	if o != nil && !IsNil(o.IsSupervised) {
		return true
	}

	return false
}

// SetIsSupervised gets a reference to the given bool and assigns it to the IsSupervised field.
func (o *AppleMdmDeviceInfo) SetIsSupervised(v bool) {
	o.IsSupervised = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetModelName() string {
	if o == nil || IsNil(o.ModelName) {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetModelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelName) {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasModelName() bool {
	if o != nil && !IsNil(o.ModelName) {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *AppleMdmDeviceInfo) SetModelName(v string) {
	o.ModelName = &v
}

// GetSecondIccid returns the SecondIccid field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetSecondIccid() string {
	if o == nil || IsNil(o.SecondIccid) {
		var ret string
		return ret
	}
	return *o.SecondIccid
}

// GetSecondIccidOk returns a tuple with the SecondIccid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetSecondIccidOk() (*string, bool) {
	if o == nil || IsNil(o.SecondIccid) {
		return nil, false
	}
	return o.SecondIccid, true
}

// HasSecondIccid returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasSecondIccid() bool {
	if o != nil && !IsNil(o.SecondIccid) {
		return true
	}

	return false
}

// SetSecondIccid gets a reference to the given string and assigns it to the SecondIccid field.
func (o *AppleMdmDeviceInfo) SetSecondIccid(v string) {
	o.SecondIccid = &v
}

// GetSecondImei returns the SecondImei field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetSecondImei() string {
	if o == nil || IsNil(o.SecondImei) {
		var ret string
		return ret
	}
	return *o.SecondImei
}

// GetSecondImeiOk returns a tuple with the SecondImei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetSecondImeiOk() (*string, bool) {
	if o == nil || IsNil(o.SecondImei) {
		return nil, false
	}
	return o.SecondImei, true
}

// HasSecondImei returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasSecondImei() bool {
	if o != nil && !IsNil(o.SecondImei) {
		return true
	}

	return false
}

// SetSecondImei gets a reference to the given string and assigns it to the SecondImei field.
func (o *AppleMdmDeviceInfo) SetSecondImei(v string) {
	o.SecondImei = &v
}

// GetSecondSubscriberCarrierNetwork returns the SecondSubscriberCarrierNetwork field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetSecondSubscriberCarrierNetwork() string {
	if o == nil || IsNil(o.SecondSubscriberCarrierNetwork) {
		var ret string
		return ret
	}
	return *o.SecondSubscriberCarrierNetwork
}

// GetSecondSubscriberCarrierNetworkOk returns a tuple with the SecondSubscriberCarrierNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetSecondSubscriberCarrierNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.SecondSubscriberCarrierNetwork) {
		return nil, false
	}
	return o.SecondSubscriberCarrierNetwork, true
}

// HasSecondSubscriberCarrierNetwork returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasSecondSubscriberCarrierNetwork() bool {
	if o != nil && !IsNil(o.SecondSubscriberCarrierNetwork) {
		return true
	}

	return false
}

// SetSecondSubscriberCarrierNetwork gets a reference to the given string and assigns it to the SecondSubscriberCarrierNetwork field.
func (o *AppleMdmDeviceInfo) SetSecondSubscriberCarrierNetwork(v string) {
	o.SecondSubscriberCarrierNetwork = &v
}

// GetSubscriberCarrierNetwork returns the SubscriberCarrierNetwork field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetSubscriberCarrierNetwork() string {
	if o == nil || IsNil(o.SubscriberCarrierNetwork) {
		var ret string
		return ret
	}
	return *o.SubscriberCarrierNetwork
}

// GetSubscriberCarrierNetworkOk returns a tuple with the SubscriberCarrierNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetSubscriberCarrierNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriberCarrierNetwork) {
		return nil, false
	}
	return o.SubscriberCarrierNetwork, true
}

// HasSubscriberCarrierNetwork returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasSubscriberCarrierNetwork() bool {
	if o != nil && !IsNil(o.SubscriberCarrierNetwork) {
		return true
	}

	return false
}

// SetSubscriberCarrierNetwork gets a reference to the given string and assigns it to the SubscriberCarrierNetwork field.
func (o *AppleMdmDeviceInfo) SetSubscriberCarrierNetwork(v string) {
	o.SubscriberCarrierNetwork = &v
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *AppleMdmDeviceInfo) GetWifiMac() string {
	if o == nil || IsNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleMdmDeviceInfo) GetWifiMacOk() (*string, bool) {
	if o == nil || IsNil(o.WifiMac) {
		return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *AppleMdmDeviceInfo) HasWifiMac() bool {
	if o != nil && !IsNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *AppleMdmDeviceInfo) SetWifiMac(v string) {
	o.WifiMac = &v
}

func (o AppleMdmDeviceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppleMdmDeviceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivationLockAllowedWhileSupervised) {
		toSerialize["activationLockAllowedWhileSupervised"] = o.ActivationLockAllowedWhileSupervised
	}
	if !IsNil(o.AvailableDeviceCapacity) {
		toSerialize["availableDeviceCapacity"] = o.AvailableDeviceCapacity
	}
	if !IsNil(o.DeviceCapacity) {
		toSerialize["deviceCapacity"] = o.DeviceCapacity
	}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !IsNil(o.Iccid) {
		toSerialize["iccid"] = o.Iccid
	}
	if !IsNil(o.Imei) {
		toSerialize["imei"] = o.Imei
	}
	if !IsNil(o.IsSupervised) {
		toSerialize["isSupervised"] = o.IsSupervised
	}
	if !IsNil(o.ModelName) {
		toSerialize["modelName"] = o.ModelName
	}
	if !IsNil(o.SecondIccid) {
		toSerialize["secondIccid"] = o.SecondIccid
	}
	if !IsNil(o.SecondImei) {
		toSerialize["secondImei"] = o.SecondImei
	}
	if !IsNil(o.SecondSubscriberCarrierNetwork) {
		toSerialize["secondSubscriberCarrierNetwork"] = o.SecondSubscriberCarrierNetwork
	}
	if !IsNil(o.SubscriberCarrierNetwork) {
		toSerialize["subscriberCarrierNetwork"] = o.SubscriberCarrierNetwork
	}
	if !IsNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppleMdmDeviceInfo) UnmarshalJSON(bytes []byte) (err error) {
	varAppleMdmDeviceInfo := _AppleMdmDeviceInfo{}

	if err = json.Unmarshal(bytes, &varAppleMdmDeviceInfo); err == nil {
		*o = AppleMdmDeviceInfo(varAppleMdmDeviceInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "activationLockAllowedWhileSupervised")
		delete(additionalProperties, "availableDeviceCapacity")
		delete(additionalProperties, "deviceCapacity")
		delete(additionalProperties, "deviceName")
		delete(additionalProperties, "iccid")
		delete(additionalProperties, "imei")
		delete(additionalProperties, "isSupervised")
		delete(additionalProperties, "modelName")
		delete(additionalProperties, "secondIccid")
		delete(additionalProperties, "secondImei")
		delete(additionalProperties, "secondSubscriberCarrierNetwork")
		delete(additionalProperties, "subscriberCarrierNetwork")
		delete(additionalProperties, "wifiMac")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppleMdmDeviceInfo struct {
	value *AppleMdmDeviceInfo
	isSet bool
}

func (v NullableAppleMdmDeviceInfo) Get() *AppleMdmDeviceInfo {
	return v.value
}

func (v *NullableAppleMdmDeviceInfo) Set(val *AppleMdmDeviceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAppleMdmDeviceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAppleMdmDeviceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppleMdmDeviceInfo(val *AppleMdmDeviceInfo) *NullableAppleMdmDeviceInfo {
	return &NullableAppleMdmDeviceInfo{value: val, isSet: true}
}

func (v NullableAppleMdmDeviceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppleMdmDeviceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



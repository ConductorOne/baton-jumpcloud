/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/ConductorOne/baton-jumpcloud/pkg/jcapi

import (
	"encoding/json"
)

// checks if the SystemInsightsSystemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInsightsSystemInfo{}

// SystemInsightsSystemInfo struct for SystemInsightsSystemInfo
type SystemInsightsSystemInfo struct {
	CollectionTime *string `json:"collection_time,omitempty"`
	ComputerName *string `json:"computer_name,omitempty"`
	CpuBrand *string `json:"cpu_brand,omitempty"`
	CpuLogicalCores *int32 `json:"cpu_logical_cores,omitempty"`
	CpuMicrocode *string `json:"cpu_microcode,omitempty"`
	CpuPhysicalCores *int32 `json:"cpu_physical_cores,omitempty"`
	CpuSubtype *string `json:"cpu_subtype,omitempty"`
	CpuType *string `json:"cpu_type,omitempty"`
	HardwareModel *string `json:"hardware_model,omitempty"`
	HardwareSerial *string `json:"hardware_serial,omitempty"`
	HardwareVendor *string `json:"hardware_vendor,omitempty"`
	HardwareVersion *string `json:"hardware_version,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	LocalHostname *string `json:"local_hostname,omitempty"`
	PhysicalMemory *string `json:"physical_memory,omitempty"`
	SystemId *string `json:"system_id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemInsightsSystemInfo SystemInsightsSystemInfo

// NewSystemInsightsSystemInfo instantiates a new SystemInsightsSystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInsightsSystemInfo() *SystemInsightsSystemInfo {
	this := SystemInsightsSystemInfo{}
	return &this
}

// NewSystemInsightsSystemInfoWithDefaults instantiates a new SystemInsightsSystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInsightsSystemInfoWithDefaults() *SystemInsightsSystemInfo {
	this := SystemInsightsSystemInfo{}
	return &this
}

// GetCollectionTime returns the CollectionTime field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetCollectionTime() string {
	if o == nil || IsNil(o.CollectionTime) {
		var ret string
		return ret
	}
	return *o.CollectionTime
}

// GetCollectionTimeOk returns a tuple with the CollectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetCollectionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionTime) {
		return nil, false
	}
	return o.CollectionTime, true
}

// HasCollectionTime returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasCollectionTime() bool {
	if o != nil && !IsNil(o.CollectionTime) {
		return true
	}

	return false
}

// SetCollectionTime gets a reference to the given string and assigns it to the CollectionTime field.
func (o *SystemInsightsSystemInfo) SetCollectionTime(v string) {
	o.CollectionTime = &v
}

// GetComputerName returns the ComputerName field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetComputerName() string {
	if o == nil || IsNil(o.ComputerName) {
		var ret string
		return ret
	}
	return *o.ComputerName
}

// GetComputerNameOk returns a tuple with the ComputerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetComputerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ComputerName) {
		return nil, false
	}
	return o.ComputerName, true
}

// HasComputerName returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasComputerName() bool {
	if o != nil && !IsNil(o.ComputerName) {
		return true
	}

	return false
}

// SetComputerName gets a reference to the given string and assigns it to the ComputerName field.
func (o *SystemInsightsSystemInfo) SetComputerName(v string) {
	o.ComputerName = &v
}

// GetCpuBrand returns the CpuBrand field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetCpuBrand() string {
	if o == nil || IsNil(o.CpuBrand) {
		var ret string
		return ret
	}
	return *o.CpuBrand
}

// GetCpuBrandOk returns a tuple with the CpuBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetCpuBrandOk() (*string, bool) {
	if o == nil || IsNil(o.CpuBrand) {
		return nil, false
	}
	return o.CpuBrand, true
}

// HasCpuBrand returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasCpuBrand() bool {
	if o != nil && !IsNil(o.CpuBrand) {
		return true
	}

	return false
}

// SetCpuBrand gets a reference to the given string and assigns it to the CpuBrand field.
func (o *SystemInsightsSystemInfo) SetCpuBrand(v string) {
	o.CpuBrand = &v
}

// GetCpuLogicalCores returns the CpuLogicalCores field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetCpuLogicalCores() int32 {
	if o == nil || IsNil(o.CpuLogicalCores) {
		var ret int32
		return ret
	}
	return *o.CpuLogicalCores
}

// GetCpuLogicalCoresOk returns a tuple with the CpuLogicalCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetCpuLogicalCoresOk() (*int32, bool) {
	if o == nil || IsNil(o.CpuLogicalCores) {
		return nil, false
	}
	return o.CpuLogicalCores, true
}

// HasCpuLogicalCores returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasCpuLogicalCores() bool {
	if o != nil && !IsNil(o.CpuLogicalCores) {
		return true
	}

	return false
}

// SetCpuLogicalCores gets a reference to the given int32 and assigns it to the CpuLogicalCores field.
func (o *SystemInsightsSystemInfo) SetCpuLogicalCores(v int32) {
	o.CpuLogicalCores = &v
}

// GetCpuMicrocode returns the CpuMicrocode field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetCpuMicrocode() string {
	if o == nil || IsNil(o.CpuMicrocode) {
		var ret string
		return ret
	}
	return *o.CpuMicrocode
}

// GetCpuMicrocodeOk returns a tuple with the CpuMicrocode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetCpuMicrocodeOk() (*string, bool) {
	if o == nil || IsNil(o.CpuMicrocode) {
		return nil, false
	}
	return o.CpuMicrocode, true
}

// HasCpuMicrocode returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasCpuMicrocode() bool {
	if o != nil && !IsNil(o.CpuMicrocode) {
		return true
	}

	return false
}

// SetCpuMicrocode gets a reference to the given string and assigns it to the CpuMicrocode field.
func (o *SystemInsightsSystemInfo) SetCpuMicrocode(v string) {
	o.CpuMicrocode = &v
}

// GetCpuPhysicalCores returns the CpuPhysicalCores field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetCpuPhysicalCores() int32 {
	if o == nil || IsNil(o.CpuPhysicalCores) {
		var ret int32
		return ret
	}
	return *o.CpuPhysicalCores
}

// GetCpuPhysicalCoresOk returns a tuple with the CpuPhysicalCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetCpuPhysicalCoresOk() (*int32, bool) {
	if o == nil || IsNil(o.CpuPhysicalCores) {
		return nil, false
	}
	return o.CpuPhysicalCores, true
}

// HasCpuPhysicalCores returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasCpuPhysicalCores() bool {
	if o != nil && !IsNil(o.CpuPhysicalCores) {
		return true
	}

	return false
}

// SetCpuPhysicalCores gets a reference to the given int32 and assigns it to the CpuPhysicalCores field.
func (o *SystemInsightsSystemInfo) SetCpuPhysicalCores(v int32) {
	o.CpuPhysicalCores = &v
}

// GetCpuSubtype returns the CpuSubtype field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetCpuSubtype() string {
	if o == nil || IsNil(o.CpuSubtype) {
		var ret string
		return ret
	}
	return *o.CpuSubtype
}

// GetCpuSubtypeOk returns a tuple with the CpuSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetCpuSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.CpuSubtype) {
		return nil, false
	}
	return o.CpuSubtype, true
}

// HasCpuSubtype returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasCpuSubtype() bool {
	if o != nil && !IsNil(o.CpuSubtype) {
		return true
	}

	return false
}

// SetCpuSubtype gets a reference to the given string and assigns it to the CpuSubtype field.
func (o *SystemInsightsSystemInfo) SetCpuSubtype(v string) {
	o.CpuSubtype = &v
}

// GetCpuType returns the CpuType field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetCpuType() string {
	if o == nil || IsNil(o.CpuType) {
		var ret string
		return ret
	}
	return *o.CpuType
}

// GetCpuTypeOk returns a tuple with the CpuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetCpuTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CpuType) {
		return nil, false
	}
	return o.CpuType, true
}

// HasCpuType returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasCpuType() bool {
	if o != nil && !IsNil(o.CpuType) {
		return true
	}

	return false
}

// SetCpuType gets a reference to the given string and assigns it to the CpuType field.
func (o *SystemInsightsSystemInfo) SetCpuType(v string) {
	o.CpuType = &v
}

// GetHardwareModel returns the HardwareModel field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetHardwareModel() string {
	if o == nil || IsNil(o.HardwareModel) {
		var ret string
		return ret
	}
	return *o.HardwareModel
}

// GetHardwareModelOk returns a tuple with the HardwareModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetHardwareModelOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareModel) {
		return nil, false
	}
	return o.HardwareModel, true
}

// HasHardwareModel returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasHardwareModel() bool {
	if o != nil && !IsNil(o.HardwareModel) {
		return true
	}

	return false
}

// SetHardwareModel gets a reference to the given string and assigns it to the HardwareModel field.
func (o *SystemInsightsSystemInfo) SetHardwareModel(v string) {
	o.HardwareModel = &v
}

// GetHardwareSerial returns the HardwareSerial field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetHardwareSerial() string {
	if o == nil || IsNil(o.HardwareSerial) {
		var ret string
		return ret
	}
	return *o.HardwareSerial
}

// GetHardwareSerialOk returns a tuple with the HardwareSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetHardwareSerialOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareSerial) {
		return nil, false
	}
	return o.HardwareSerial, true
}

// HasHardwareSerial returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasHardwareSerial() bool {
	if o != nil && !IsNil(o.HardwareSerial) {
		return true
	}

	return false
}

// SetHardwareSerial gets a reference to the given string and assigns it to the HardwareSerial field.
func (o *SystemInsightsSystemInfo) SetHardwareSerial(v string) {
	o.HardwareSerial = &v
}

// GetHardwareVendor returns the HardwareVendor field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetHardwareVendor() string {
	if o == nil || IsNil(o.HardwareVendor) {
		var ret string
		return ret
	}
	return *o.HardwareVendor
}

// GetHardwareVendorOk returns a tuple with the HardwareVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetHardwareVendorOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareVendor) {
		return nil, false
	}
	return o.HardwareVendor, true
}

// HasHardwareVendor returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasHardwareVendor() bool {
	if o != nil && !IsNil(o.HardwareVendor) {
		return true
	}

	return false
}

// SetHardwareVendor gets a reference to the given string and assigns it to the HardwareVendor field.
func (o *SystemInsightsSystemInfo) SetHardwareVendor(v string) {
	o.HardwareVendor = &v
}

// GetHardwareVersion returns the HardwareVersion field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetHardwareVersion() string {
	if o == nil || IsNil(o.HardwareVersion) {
		var ret string
		return ret
	}
	return *o.HardwareVersion
}

// GetHardwareVersionOk returns a tuple with the HardwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetHardwareVersionOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareVersion) {
		return nil, false
	}
	return o.HardwareVersion, true
}

// HasHardwareVersion returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasHardwareVersion() bool {
	if o != nil && !IsNil(o.HardwareVersion) {
		return true
	}

	return false
}

// SetHardwareVersion gets a reference to the given string and assigns it to the HardwareVersion field.
func (o *SystemInsightsSystemInfo) SetHardwareVersion(v string) {
	o.HardwareVersion = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SystemInsightsSystemInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetLocalHostname returns the LocalHostname field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetLocalHostname() string {
	if o == nil || IsNil(o.LocalHostname) {
		var ret string
		return ret
	}
	return *o.LocalHostname
}

// GetLocalHostnameOk returns a tuple with the LocalHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetLocalHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.LocalHostname) {
		return nil, false
	}
	return o.LocalHostname, true
}

// HasLocalHostname returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasLocalHostname() bool {
	if o != nil && !IsNil(o.LocalHostname) {
		return true
	}

	return false
}

// SetLocalHostname gets a reference to the given string and assigns it to the LocalHostname field.
func (o *SystemInsightsSystemInfo) SetLocalHostname(v string) {
	o.LocalHostname = &v
}

// GetPhysicalMemory returns the PhysicalMemory field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetPhysicalMemory() string {
	if o == nil || IsNil(o.PhysicalMemory) {
		var ret string
		return ret
	}
	return *o.PhysicalMemory
}

// GetPhysicalMemoryOk returns a tuple with the PhysicalMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetPhysicalMemoryOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalMemory) {
		return nil, false
	}
	return o.PhysicalMemory, true
}

// HasPhysicalMemory returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasPhysicalMemory() bool {
	if o != nil && !IsNil(o.PhysicalMemory) {
		return true
	}

	return false
}

// SetPhysicalMemory gets a reference to the given string and assigns it to the PhysicalMemory field.
func (o *SystemInsightsSystemInfo) SetPhysicalMemory(v string) {
	o.PhysicalMemory = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *SystemInsightsSystemInfo) SetSystemId(v string) {
	o.SystemId = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SystemInsightsSystemInfo) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsSystemInfo) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SystemInsightsSystemInfo) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SystemInsightsSystemInfo) SetUuid(v string) {
	o.Uuid = &v
}

func (o SystemInsightsSystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInsightsSystemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionTime) {
		toSerialize["collection_time"] = o.CollectionTime
	}
	if !IsNil(o.ComputerName) {
		toSerialize["computer_name"] = o.ComputerName
	}
	if !IsNil(o.CpuBrand) {
		toSerialize["cpu_brand"] = o.CpuBrand
	}
	if !IsNil(o.CpuLogicalCores) {
		toSerialize["cpu_logical_cores"] = o.CpuLogicalCores
	}
	if !IsNil(o.CpuMicrocode) {
		toSerialize["cpu_microcode"] = o.CpuMicrocode
	}
	if !IsNil(o.CpuPhysicalCores) {
		toSerialize["cpu_physical_cores"] = o.CpuPhysicalCores
	}
	if !IsNil(o.CpuSubtype) {
		toSerialize["cpu_subtype"] = o.CpuSubtype
	}
	if !IsNil(o.CpuType) {
		toSerialize["cpu_type"] = o.CpuType
	}
	if !IsNil(o.HardwareModel) {
		toSerialize["hardware_model"] = o.HardwareModel
	}
	if !IsNil(o.HardwareSerial) {
		toSerialize["hardware_serial"] = o.HardwareSerial
	}
	if !IsNil(o.HardwareVendor) {
		toSerialize["hardware_vendor"] = o.HardwareVendor
	}
	if !IsNil(o.HardwareVersion) {
		toSerialize["hardware_version"] = o.HardwareVersion
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.LocalHostname) {
		toSerialize["local_hostname"] = o.LocalHostname
	}
	if !IsNil(o.PhysicalMemory) {
		toSerialize["physical_memory"] = o.PhysicalMemory
	}
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemInsightsSystemInfo) UnmarshalJSON(bytes []byte) (err error) {
	varSystemInsightsSystemInfo := _SystemInsightsSystemInfo{}

	if err = json.Unmarshal(bytes, &varSystemInsightsSystemInfo); err == nil {
		*o = SystemInsightsSystemInfo(varSystemInsightsSystemInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "collection_time")
		delete(additionalProperties, "computer_name")
		delete(additionalProperties, "cpu_brand")
		delete(additionalProperties, "cpu_logical_cores")
		delete(additionalProperties, "cpu_microcode")
		delete(additionalProperties, "cpu_physical_cores")
		delete(additionalProperties, "cpu_subtype")
		delete(additionalProperties, "cpu_type")
		delete(additionalProperties, "hardware_model")
		delete(additionalProperties, "hardware_serial")
		delete(additionalProperties, "hardware_vendor")
		delete(additionalProperties, "hardware_version")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "local_hostname")
		delete(additionalProperties, "physical_memory")
		delete(additionalProperties, "system_id")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemInsightsSystemInfo struct {
	value *SystemInsightsSystemInfo
	isSet bool
}

func (v NullableSystemInsightsSystemInfo) Get() *SystemInsightsSystemInfo {
	return v.value
}

func (v *NullableSystemInsightsSystemInfo) Set(val *SystemInsightsSystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInsightsSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInsightsSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInsightsSystemInfo(val *SystemInsightsSystemInfo) *NullableSystemInsightsSystemInfo {
	return &NullableSystemInsightsSystemInfo{value: val, isSet: true}
}

func (v NullableSystemInsightsSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInsightsSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



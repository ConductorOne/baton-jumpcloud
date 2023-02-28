/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SoftwareAppSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwareAppSettings{}

// SoftwareAppSettings struct for SoftwareAppSettings
type SoftwareAppSettings struct {
	AllowUpdateDelay *bool `json:"allowUpdateDelay,omitempty"`
	AppleVpp *SoftwareAppAppleVpp `json:"appleVpp,omitempty"`
	// The manifest asset kind (ex: software).
	AssetKind *string `json:"assetKind,omitempty"`
	// The incremental size to use for summing the package as it is downloaded.
	AssetSha256Size *int32 `json:"assetSha256Size,omitempty"`
	// The array of checksums, one each for the hash size up to the total size of the package.
	AssetSha256Strings []string `json:"assetSha256Strings,omitempty"`
	AutoUpdate *bool `json:"autoUpdate,omitempty"`
	// The software app description.
	Description *string `json:"description,omitempty"`
	// State of Install or Uninstall
	DesiredState *string `json:"desiredState,omitempty"`
	// ID of the Enterprise with which this app is associated
	EnterpriseObjectId *string `json:"enterpriseObjectId,omitempty"`
	GoogleAndroid *SoftwareAppGoogleAndroid `json:"googleAndroid,omitempty"`
	// Repository where the app is located within the package manager
	Location *string `json:"location,omitempty"`
	// ID of the repository where the app is located within the package manager
	LocationObjectId *string `json:"locationObjectId,omitempty"`
	PackageId *string `json:"packageId,omitempty"`
	// The package manifest kind (ex: software-package).
	PackageKind *string `json:"packageKind,omitempty"`
	// App store serving the app: APPLE_VPP, CHOCOLATEY, etc.
	PackageManager *string `json:"packageManager,omitempty"`
	// The package manifest subtitle.
	PackageSubtitle *string `json:"packageSubtitle,omitempty"`
	// The package manifest version.
	PackageVersion *string `json:"packageVersion,omitempty"`
}

// NewSoftwareAppSettings instantiates a new SoftwareAppSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareAppSettings() *SoftwareAppSettings {
	this := SoftwareAppSettings{}
	var allowUpdateDelay bool = false
	this.AllowUpdateDelay = &allowUpdateDelay
	var autoUpdate bool = false
	this.AutoUpdate = &autoUpdate
	return &this
}

// NewSoftwareAppSettingsWithDefaults instantiates a new SoftwareAppSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareAppSettingsWithDefaults() *SoftwareAppSettings {
	this := SoftwareAppSettings{}
	var allowUpdateDelay bool = false
	this.AllowUpdateDelay = &allowUpdateDelay
	var autoUpdate bool = false
	this.AutoUpdate = &autoUpdate
	return &this
}

// GetAllowUpdateDelay returns the AllowUpdateDelay field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetAllowUpdateDelay() bool {
	if o == nil || IsNil(o.AllowUpdateDelay) {
		var ret bool
		return ret
	}
	return *o.AllowUpdateDelay
}

// GetAllowUpdateDelayOk returns a tuple with the AllowUpdateDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetAllowUpdateDelayOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUpdateDelay) {
		return nil, false
	}
	return o.AllowUpdateDelay, true
}

// HasAllowUpdateDelay returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasAllowUpdateDelay() bool {
	if o != nil && !IsNil(o.AllowUpdateDelay) {
		return true
	}

	return false
}

// SetAllowUpdateDelay gets a reference to the given bool and assigns it to the AllowUpdateDelay field.
func (o *SoftwareAppSettings) SetAllowUpdateDelay(v bool) {
	o.AllowUpdateDelay = &v
}

// GetAppleVpp returns the AppleVpp field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetAppleVpp() SoftwareAppAppleVpp {
	if o == nil || IsNil(o.AppleVpp) {
		var ret SoftwareAppAppleVpp
		return ret
	}
	return *o.AppleVpp
}

// GetAppleVppOk returns a tuple with the AppleVpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetAppleVppOk() (*SoftwareAppAppleVpp, bool) {
	if o == nil || IsNil(o.AppleVpp) {
		return nil, false
	}
	return o.AppleVpp, true
}

// HasAppleVpp returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasAppleVpp() bool {
	if o != nil && !IsNil(o.AppleVpp) {
		return true
	}

	return false
}

// SetAppleVpp gets a reference to the given SoftwareAppAppleVpp and assigns it to the AppleVpp field.
func (o *SoftwareAppSettings) SetAppleVpp(v SoftwareAppAppleVpp) {
	o.AppleVpp = &v
}

// GetAssetKind returns the AssetKind field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetAssetKind() string {
	if o == nil || IsNil(o.AssetKind) {
		var ret string
		return ret
	}
	return *o.AssetKind
}

// GetAssetKindOk returns a tuple with the AssetKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetAssetKindOk() (*string, bool) {
	if o == nil || IsNil(o.AssetKind) {
		return nil, false
	}
	return o.AssetKind, true
}

// HasAssetKind returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasAssetKind() bool {
	if o != nil && !IsNil(o.AssetKind) {
		return true
	}

	return false
}

// SetAssetKind gets a reference to the given string and assigns it to the AssetKind field.
func (o *SoftwareAppSettings) SetAssetKind(v string) {
	o.AssetKind = &v
}

// GetAssetSha256Size returns the AssetSha256Size field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetAssetSha256Size() int32 {
	if o == nil || IsNil(o.AssetSha256Size) {
		var ret int32
		return ret
	}
	return *o.AssetSha256Size
}

// GetAssetSha256SizeOk returns a tuple with the AssetSha256Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetAssetSha256SizeOk() (*int32, bool) {
	if o == nil || IsNil(o.AssetSha256Size) {
		return nil, false
	}
	return o.AssetSha256Size, true
}

// HasAssetSha256Size returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasAssetSha256Size() bool {
	if o != nil && !IsNil(o.AssetSha256Size) {
		return true
	}

	return false
}

// SetAssetSha256Size gets a reference to the given int32 and assigns it to the AssetSha256Size field.
func (o *SoftwareAppSettings) SetAssetSha256Size(v int32) {
	o.AssetSha256Size = &v
}

// GetAssetSha256Strings returns the AssetSha256Strings field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetAssetSha256Strings() []string {
	if o == nil || IsNil(o.AssetSha256Strings) {
		var ret []string
		return ret
	}
	return o.AssetSha256Strings
}

// GetAssetSha256StringsOk returns a tuple with the AssetSha256Strings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetAssetSha256StringsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssetSha256Strings) {
		return nil, false
	}
	return o.AssetSha256Strings, true
}

// HasAssetSha256Strings returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasAssetSha256Strings() bool {
	if o != nil && !IsNil(o.AssetSha256Strings) {
		return true
	}

	return false
}

// SetAssetSha256Strings gets a reference to the given []string and assigns it to the AssetSha256Strings field.
func (o *SoftwareAppSettings) SetAssetSha256Strings(v []string) {
	o.AssetSha256Strings = v
}

// GetAutoUpdate returns the AutoUpdate field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetAutoUpdate() bool {
	if o == nil || IsNil(o.AutoUpdate) {
		var ret bool
		return ret
	}
	return *o.AutoUpdate
}

// GetAutoUpdateOk returns a tuple with the AutoUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetAutoUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoUpdate) {
		return nil, false
	}
	return o.AutoUpdate, true
}

// HasAutoUpdate returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasAutoUpdate() bool {
	if o != nil && !IsNil(o.AutoUpdate) {
		return true
	}

	return false
}

// SetAutoUpdate gets a reference to the given bool and assigns it to the AutoUpdate field.
func (o *SoftwareAppSettings) SetAutoUpdate(v bool) {
	o.AutoUpdate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SoftwareAppSettings) SetDescription(v string) {
	o.Description = &v
}

// GetDesiredState returns the DesiredState field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetDesiredState() string {
	if o == nil || IsNil(o.DesiredState) {
		var ret string
		return ret
	}
	return *o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetDesiredStateOk() (*string, bool) {
	if o == nil || IsNil(o.DesiredState) {
		return nil, false
	}
	return o.DesiredState, true
}

// HasDesiredState returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasDesiredState() bool {
	if o != nil && !IsNil(o.DesiredState) {
		return true
	}

	return false
}

// SetDesiredState gets a reference to the given string and assigns it to the DesiredState field.
func (o *SoftwareAppSettings) SetDesiredState(v string) {
	o.DesiredState = &v
}

// GetEnterpriseObjectId returns the EnterpriseObjectId field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetEnterpriseObjectId() string {
	if o == nil || IsNil(o.EnterpriseObjectId) {
		var ret string
		return ret
	}
	return *o.EnterpriseObjectId
}

// GetEnterpriseObjectIdOk returns a tuple with the EnterpriseObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetEnterpriseObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseObjectId) {
		return nil, false
	}
	return o.EnterpriseObjectId, true
}

// HasEnterpriseObjectId returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasEnterpriseObjectId() bool {
	if o != nil && !IsNil(o.EnterpriseObjectId) {
		return true
	}

	return false
}

// SetEnterpriseObjectId gets a reference to the given string and assigns it to the EnterpriseObjectId field.
func (o *SoftwareAppSettings) SetEnterpriseObjectId(v string) {
	o.EnterpriseObjectId = &v
}

// GetGoogleAndroid returns the GoogleAndroid field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetGoogleAndroid() SoftwareAppGoogleAndroid {
	if o == nil || IsNil(o.GoogleAndroid) {
		var ret SoftwareAppGoogleAndroid
		return ret
	}
	return *o.GoogleAndroid
}

// GetGoogleAndroidOk returns a tuple with the GoogleAndroid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetGoogleAndroidOk() (*SoftwareAppGoogleAndroid, bool) {
	if o == nil || IsNil(o.GoogleAndroid) {
		return nil, false
	}
	return o.GoogleAndroid, true
}

// HasGoogleAndroid returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasGoogleAndroid() bool {
	if o != nil && !IsNil(o.GoogleAndroid) {
		return true
	}

	return false
}

// SetGoogleAndroid gets a reference to the given SoftwareAppGoogleAndroid and assigns it to the GoogleAndroid field.
func (o *SoftwareAppSettings) SetGoogleAndroid(v SoftwareAppGoogleAndroid) {
	o.GoogleAndroid = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SoftwareAppSettings) SetLocation(v string) {
	o.Location = &v
}

// GetLocationObjectId returns the LocationObjectId field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetLocationObjectId() string {
	if o == nil || IsNil(o.LocationObjectId) {
		var ret string
		return ret
	}
	return *o.LocationObjectId
}

// GetLocationObjectIdOk returns a tuple with the LocationObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetLocationObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocationObjectId) {
		return nil, false
	}
	return o.LocationObjectId, true
}

// HasLocationObjectId returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasLocationObjectId() bool {
	if o != nil && !IsNil(o.LocationObjectId) {
		return true
	}

	return false
}

// SetLocationObjectId gets a reference to the given string and assigns it to the LocationObjectId field.
func (o *SoftwareAppSettings) SetLocationObjectId(v string) {
	o.LocationObjectId = &v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetPackageId() string {
	if o == nil || IsNil(o.PackageId) {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PackageId) {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasPackageId() bool {
	if o != nil && !IsNil(o.PackageId) {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *SoftwareAppSettings) SetPackageId(v string) {
	o.PackageId = &v
}

// GetPackageKind returns the PackageKind field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetPackageKind() string {
	if o == nil || IsNil(o.PackageKind) {
		var ret string
		return ret
	}
	return *o.PackageKind
}

// GetPackageKindOk returns a tuple with the PackageKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetPackageKindOk() (*string, bool) {
	if o == nil || IsNil(o.PackageKind) {
		return nil, false
	}
	return o.PackageKind, true
}

// HasPackageKind returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasPackageKind() bool {
	if o != nil && !IsNil(o.PackageKind) {
		return true
	}

	return false
}

// SetPackageKind gets a reference to the given string and assigns it to the PackageKind field.
func (o *SoftwareAppSettings) SetPackageKind(v string) {
	o.PackageKind = &v
}

// GetPackageManager returns the PackageManager field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetPackageManager() string {
	if o == nil || IsNil(o.PackageManager) {
		var ret string
		return ret
	}
	return *o.PackageManager
}

// GetPackageManagerOk returns a tuple with the PackageManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetPackageManagerOk() (*string, bool) {
	if o == nil || IsNil(o.PackageManager) {
		return nil, false
	}
	return o.PackageManager, true
}

// HasPackageManager returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasPackageManager() bool {
	if o != nil && !IsNil(o.PackageManager) {
		return true
	}

	return false
}

// SetPackageManager gets a reference to the given string and assigns it to the PackageManager field.
func (o *SoftwareAppSettings) SetPackageManager(v string) {
	o.PackageManager = &v
}

// GetPackageSubtitle returns the PackageSubtitle field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetPackageSubtitle() string {
	if o == nil || IsNil(o.PackageSubtitle) {
		var ret string
		return ret
	}
	return *o.PackageSubtitle
}

// GetPackageSubtitleOk returns a tuple with the PackageSubtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetPackageSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.PackageSubtitle) {
		return nil, false
	}
	return o.PackageSubtitle, true
}

// HasPackageSubtitle returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasPackageSubtitle() bool {
	if o != nil && !IsNil(o.PackageSubtitle) {
		return true
	}

	return false
}

// SetPackageSubtitle gets a reference to the given string and assigns it to the PackageSubtitle field.
func (o *SoftwareAppSettings) SetPackageSubtitle(v string) {
	o.PackageSubtitle = &v
}

// GetPackageVersion returns the PackageVersion field value if set, zero value otherwise.
func (o *SoftwareAppSettings) GetPackageVersion() string {
	if o == nil || IsNil(o.PackageVersion) {
		var ret string
		return ret
	}
	return *o.PackageVersion
}

// GetPackageVersionOk returns a tuple with the PackageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppSettings) GetPackageVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PackageVersion) {
		return nil, false
	}
	return o.PackageVersion, true
}

// HasPackageVersion returns a boolean if a field has been set.
func (o *SoftwareAppSettings) HasPackageVersion() bool {
	if o != nil && !IsNil(o.PackageVersion) {
		return true
	}

	return false
}

// SetPackageVersion gets a reference to the given string and assigns it to the PackageVersion field.
func (o *SoftwareAppSettings) SetPackageVersion(v string) {
	o.PackageVersion = &v
}

func (o SoftwareAppSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwareAppSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowUpdateDelay) {
		toSerialize["allowUpdateDelay"] = o.AllowUpdateDelay
	}
	if !IsNil(o.AppleVpp) {
		toSerialize["appleVpp"] = o.AppleVpp
	}
	if !IsNil(o.AssetKind) {
		toSerialize["assetKind"] = o.AssetKind
	}
	if !IsNil(o.AssetSha256Size) {
		toSerialize["assetSha256Size"] = o.AssetSha256Size
	}
	if !IsNil(o.AssetSha256Strings) {
		toSerialize["assetSha256Strings"] = o.AssetSha256Strings
	}
	if !IsNil(o.AutoUpdate) {
		toSerialize["autoUpdate"] = o.AutoUpdate
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DesiredState) {
		toSerialize["desiredState"] = o.DesiredState
	}
	if !IsNil(o.EnterpriseObjectId) {
		toSerialize["enterpriseObjectId"] = o.EnterpriseObjectId
	}
	if !IsNil(o.GoogleAndroid) {
		toSerialize["googleAndroid"] = o.GoogleAndroid
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.LocationObjectId) {
		toSerialize["locationObjectId"] = o.LocationObjectId
	}
	if !IsNil(o.PackageId) {
		toSerialize["packageId"] = o.PackageId
	}
	if !IsNil(o.PackageKind) {
		toSerialize["packageKind"] = o.PackageKind
	}
	if !IsNil(o.PackageManager) {
		toSerialize["packageManager"] = o.PackageManager
	}
	if !IsNil(o.PackageSubtitle) {
		toSerialize["packageSubtitle"] = o.PackageSubtitle
	}
	if !IsNil(o.PackageVersion) {
		toSerialize["packageVersion"] = o.PackageVersion
	}
	return toSerialize, nil
}

type NullableSoftwareAppSettings struct {
	value *SoftwareAppSettings
	isSet bool
}

func (v NullableSoftwareAppSettings) Get() *SoftwareAppSettings {
	return v.value
}

func (v *NullableSoftwareAppSettings) Set(val *SoftwareAppSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareAppSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareAppSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareAppSettings(val *SoftwareAppSettings) *NullableSoftwareAppSettings {
	return &NullableSoftwareAppSettings{value: val, isSet: true}
}

func (v NullableSoftwareAppSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareAppSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



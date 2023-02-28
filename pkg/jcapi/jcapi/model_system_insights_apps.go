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

// checks if the SystemInsightsApps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInsightsApps{}

// SystemInsightsApps struct for SystemInsightsApps
type SystemInsightsApps struct {
	ApplescriptEnabled *string `json:"applescript_enabled,omitempty"`
	BundleExecutable *string `json:"bundle_executable,omitempty"`
	BundleIdentifier *string `json:"bundle_identifier,omitempty"`
	BundleName *string `json:"bundle_name,omitempty"`
	BundlePackageType *string `json:"bundle_package_type,omitempty"`
	BundleShortVersion *string `json:"bundle_short_version,omitempty"`
	BundleVersion *string `json:"bundle_version,omitempty"`
	Category *string `json:"category,omitempty"`
	CollectionTime *string `json:"collection_time,omitempty"`
	Compiler *string `json:"compiler,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	DevelopmentRegion *string `json:"development_region,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Element *string `json:"element,omitempty"`
	Environment *string `json:"environment,omitempty"`
	InfoString *string `json:"info_string,omitempty"`
	LastOpenedTime *float32 `json:"last_opened_time,omitempty"`
	MinimumSystemVersion *string `json:"minimum_system_version,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
	SystemId *string `json:"system_id,omitempty"`
}

// NewSystemInsightsApps instantiates a new SystemInsightsApps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInsightsApps() *SystemInsightsApps {
	this := SystemInsightsApps{}
	return &this
}

// NewSystemInsightsAppsWithDefaults instantiates a new SystemInsightsApps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInsightsAppsWithDefaults() *SystemInsightsApps {
	this := SystemInsightsApps{}
	return &this
}

// GetApplescriptEnabled returns the ApplescriptEnabled field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetApplescriptEnabled() string {
	if o == nil || IsNil(o.ApplescriptEnabled) {
		var ret string
		return ret
	}
	return *o.ApplescriptEnabled
}

// GetApplescriptEnabledOk returns a tuple with the ApplescriptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetApplescriptEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.ApplescriptEnabled) {
		return nil, false
	}
	return o.ApplescriptEnabled, true
}

// HasApplescriptEnabled returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasApplescriptEnabled() bool {
	if o != nil && !IsNil(o.ApplescriptEnabled) {
		return true
	}

	return false
}

// SetApplescriptEnabled gets a reference to the given string and assigns it to the ApplescriptEnabled field.
func (o *SystemInsightsApps) SetApplescriptEnabled(v string) {
	o.ApplescriptEnabled = &v
}

// GetBundleExecutable returns the BundleExecutable field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetBundleExecutable() string {
	if o == nil || IsNil(o.BundleExecutable) {
		var ret string
		return ret
	}
	return *o.BundleExecutable
}

// GetBundleExecutableOk returns a tuple with the BundleExecutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetBundleExecutableOk() (*string, bool) {
	if o == nil || IsNil(o.BundleExecutable) {
		return nil, false
	}
	return o.BundleExecutable, true
}

// HasBundleExecutable returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasBundleExecutable() bool {
	if o != nil && !IsNil(o.BundleExecutable) {
		return true
	}

	return false
}

// SetBundleExecutable gets a reference to the given string and assigns it to the BundleExecutable field.
func (o *SystemInsightsApps) SetBundleExecutable(v string) {
	o.BundleExecutable = &v
}

// GetBundleIdentifier returns the BundleIdentifier field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetBundleIdentifier() string {
	if o == nil || IsNil(o.BundleIdentifier) {
		var ret string
		return ret
	}
	return *o.BundleIdentifier
}

// GetBundleIdentifierOk returns a tuple with the BundleIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetBundleIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.BundleIdentifier) {
		return nil, false
	}
	return o.BundleIdentifier, true
}

// HasBundleIdentifier returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasBundleIdentifier() bool {
	if o != nil && !IsNil(o.BundleIdentifier) {
		return true
	}

	return false
}

// SetBundleIdentifier gets a reference to the given string and assigns it to the BundleIdentifier field.
func (o *SystemInsightsApps) SetBundleIdentifier(v string) {
	o.BundleIdentifier = &v
}

// GetBundleName returns the BundleName field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetBundleName() string {
	if o == nil || IsNil(o.BundleName) {
		var ret string
		return ret
	}
	return *o.BundleName
}

// GetBundleNameOk returns a tuple with the BundleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetBundleNameOk() (*string, bool) {
	if o == nil || IsNil(o.BundleName) {
		return nil, false
	}
	return o.BundleName, true
}

// HasBundleName returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasBundleName() bool {
	if o != nil && !IsNil(o.BundleName) {
		return true
	}

	return false
}

// SetBundleName gets a reference to the given string and assigns it to the BundleName field.
func (o *SystemInsightsApps) SetBundleName(v string) {
	o.BundleName = &v
}

// GetBundlePackageType returns the BundlePackageType field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetBundlePackageType() string {
	if o == nil || IsNil(o.BundlePackageType) {
		var ret string
		return ret
	}
	return *o.BundlePackageType
}

// GetBundlePackageTypeOk returns a tuple with the BundlePackageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetBundlePackageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BundlePackageType) {
		return nil, false
	}
	return o.BundlePackageType, true
}

// HasBundlePackageType returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasBundlePackageType() bool {
	if o != nil && !IsNil(o.BundlePackageType) {
		return true
	}

	return false
}

// SetBundlePackageType gets a reference to the given string and assigns it to the BundlePackageType field.
func (o *SystemInsightsApps) SetBundlePackageType(v string) {
	o.BundlePackageType = &v
}

// GetBundleShortVersion returns the BundleShortVersion field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetBundleShortVersion() string {
	if o == nil || IsNil(o.BundleShortVersion) {
		var ret string
		return ret
	}
	return *o.BundleShortVersion
}

// GetBundleShortVersionOk returns a tuple with the BundleShortVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetBundleShortVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BundleShortVersion) {
		return nil, false
	}
	return o.BundleShortVersion, true
}

// HasBundleShortVersion returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasBundleShortVersion() bool {
	if o != nil && !IsNil(o.BundleShortVersion) {
		return true
	}

	return false
}

// SetBundleShortVersion gets a reference to the given string and assigns it to the BundleShortVersion field.
func (o *SystemInsightsApps) SetBundleShortVersion(v string) {
	o.BundleShortVersion = &v
}

// GetBundleVersion returns the BundleVersion field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetBundleVersion() string {
	if o == nil || IsNil(o.BundleVersion) {
		var ret string
		return ret
	}
	return *o.BundleVersion
}

// GetBundleVersionOk returns a tuple with the BundleVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetBundleVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BundleVersion) {
		return nil, false
	}
	return o.BundleVersion, true
}

// HasBundleVersion returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasBundleVersion() bool {
	if o != nil && !IsNil(o.BundleVersion) {
		return true
	}

	return false
}

// SetBundleVersion gets a reference to the given string and assigns it to the BundleVersion field.
func (o *SystemInsightsApps) SetBundleVersion(v string) {
	o.BundleVersion = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SystemInsightsApps) SetCategory(v string) {
	o.Category = &v
}

// GetCollectionTime returns the CollectionTime field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetCollectionTime() string {
	if o == nil || IsNil(o.CollectionTime) {
		var ret string
		return ret
	}
	return *o.CollectionTime
}

// GetCollectionTimeOk returns a tuple with the CollectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetCollectionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionTime) {
		return nil, false
	}
	return o.CollectionTime, true
}

// HasCollectionTime returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasCollectionTime() bool {
	if o != nil && !IsNil(o.CollectionTime) {
		return true
	}

	return false
}

// SetCollectionTime gets a reference to the given string and assigns it to the CollectionTime field.
func (o *SystemInsightsApps) SetCollectionTime(v string) {
	o.CollectionTime = &v
}

// GetCompiler returns the Compiler field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetCompiler() string {
	if o == nil || IsNil(o.Compiler) {
		var ret string
		return ret
	}
	return *o.Compiler
}

// GetCompilerOk returns a tuple with the Compiler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetCompilerOk() (*string, bool) {
	if o == nil || IsNil(o.Compiler) {
		return nil, false
	}
	return o.Compiler, true
}

// HasCompiler returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasCompiler() bool {
	if o != nil && !IsNil(o.Compiler) {
		return true
	}

	return false
}

// SetCompiler gets a reference to the given string and assigns it to the Compiler field.
func (o *SystemInsightsApps) SetCompiler(v string) {
	o.Compiler = &v
}

// GetCopyright returns the Copyright field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetCopyright() string {
	if o == nil || IsNil(o.Copyright) {
		var ret string
		return ret
	}
	return *o.Copyright
}

// GetCopyrightOk returns a tuple with the Copyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetCopyrightOk() (*string, bool) {
	if o == nil || IsNil(o.Copyright) {
		return nil, false
	}
	return o.Copyright, true
}

// HasCopyright returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasCopyright() bool {
	if o != nil && !IsNil(o.Copyright) {
		return true
	}

	return false
}

// SetCopyright gets a reference to the given string and assigns it to the Copyright field.
func (o *SystemInsightsApps) SetCopyright(v string) {
	o.Copyright = &v
}

// GetDevelopmentRegion returns the DevelopmentRegion field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetDevelopmentRegion() string {
	if o == nil || IsNil(o.DevelopmentRegion) {
		var ret string
		return ret
	}
	return *o.DevelopmentRegion
}

// GetDevelopmentRegionOk returns a tuple with the DevelopmentRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetDevelopmentRegionOk() (*string, bool) {
	if o == nil || IsNil(o.DevelopmentRegion) {
		return nil, false
	}
	return o.DevelopmentRegion, true
}

// HasDevelopmentRegion returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasDevelopmentRegion() bool {
	if o != nil && !IsNil(o.DevelopmentRegion) {
		return true
	}

	return false
}

// SetDevelopmentRegion gets a reference to the given string and assigns it to the DevelopmentRegion field.
func (o *SystemInsightsApps) SetDevelopmentRegion(v string) {
	o.DevelopmentRegion = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SystemInsightsApps) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetElement() string {
	if o == nil || IsNil(o.Element) {
		var ret string
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetElementOk() (*string, bool) {
	if o == nil || IsNil(o.Element) {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasElement() bool {
	if o != nil && !IsNil(o.Element) {
		return true
	}

	return false
}

// SetElement gets a reference to the given string and assigns it to the Element field.
func (o *SystemInsightsApps) SetElement(v string) {
	o.Element = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *SystemInsightsApps) SetEnvironment(v string) {
	o.Environment = &v
}

// GetInfoString returns the InfoString field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetInfoString() string {
	if o == nil || IsNil(o.InfoString) {
		var ret string
		return ret
	}
	return *o.InfoString
}

// GetInfoStringOk returns a tuple with the InfoString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetInfoStringOk() (*string, bool) {
	if o == nil || IsNil(o.InfoString) {
		return nil, false
	}
	return o.InfoString, true
}

// HasInfoString returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasInfoString() bool {
	if o != nil && !IsNil(o.InfoString) {
		return true
	}

	return false
}

// SetInfoString gets a reference to the given string and assigns it to the InfoString field.
func (o *SystemInsightsApps) SetInfoString(v string) {
	o.InfoString = &v
}

// GetLastOpenedTime returns the LastOpenedTime field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetLastOpenedTime() float32 {
	if o == nil || IsNil(o.LastOpenedTime) {
		var ret float32
		return ret
	}
	return *o.LastOpenedTime
}

// GetLastOpenedTimeOk returns a tuple with the LastOpenedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetLastOpenedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.LastOpenedTime) {
		return nil, false
	}
	return o.LastOpenedTime, true
}

// HasLastOpenedTime returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasLastOpenedTime() bool {
	if o != nil && !IsNil(o.LastOpenedTime) {
		return true
	}

	return false
}

// SetLastOpenedTime gets a reference to the given float32 and assigns it to the LastOpenedTime field.
func (o *SystemInsightsApps) SetLastOpenedTime(v float32) {
	o.LastOpenedTime = &v
}

// GetMinimumSystemVersion returns the MinimumSystemVersion field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetMinimumSystemVersion() string {
	if o == nil || IsNil(o.MinimumSystemVersion) {
		var ret string
		return ret
	}
	return *o.MinimumSystemVersion
}

// GetMinimumSystemVersionOk returns a tuple with the MinimumSystemVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetMinimumSystemVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumSystemVersion) {
		return nil, false
	}
	return o.MinimumSystemVersion, true
}

// HasMinimumSystemVersion returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasMinimumSystemVersion() bool {
	if o != nil && !IsNil(o.MinimumSystemVersion) {
		return true
	}

	return false
}

// SetMinimumSystemVersion gets a reference to the given string and assigns it to the MinimumSystemVersion field.
func (o *SystemInsightsApps) SetMinimumSystemVersion(v string) {
	o.MinimumSystemVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SystemInsightsApps) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SystemInsightsApps) SetPath(v string) {
	o.Path = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *SystemInsightsApps) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsApps) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *SystemInsightsApps) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *SystemInsightsApps) SetSystemId(v string) {
	o.SystemId = &v
}

func (o SystemInsightsApps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInsightsApps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplescriptEnabled) {
		toSerialize["applescript_enabled"] = o.ApplescriptEnabled
	}
	if !IsNil(o.BundleExecutable) {
		toSerialize["bundle_executable"] = o.BundleExecutable
	}
	if !IsNil(o.BundleIdentifier) {
		toSerialize["bundle_identifier"] = o.BundleIdentifier
	}
	if !IsNil(o.BundleName) {
		toSerialize["bundle_name"] = o.BundleName
	}
	if !IsNil(o.BundlePackageType) {
		toSerialize["bundle_package_type"] = o.BundlePackageType
	}
	if !IsNil(o.BundleShortVersion) {
		toSerialize["bundle_short_version"] = o.BundleShortVersion
	}
	if !IsNil(o.BundleVersion) {
		toSerialize["bundle_version"] = o.BundleVersion
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.CollectionTime) {
		toSerialize["collection_time"] = o.CollectionTime
	}
	if !IsNil(o.Compiler) {
		toSerialize["compiler"] = o.Compiler
	}
	if !IsNil(o.Copyright) {
		toSerialize["copyright"] = o.Copyright
	}
	if !IsNil(o.DevelopmentRegion) {
		toSerialize["development_region"] = o.DevelopmentRegion
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Element) {
		toSerialize["element"] = o.Element
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.InfoString) {
		toSerialize["info_string"] = o.InfoString
	}
	if !IsNil(o.LastOpenedTime) {
		toSerialize["last_opened_time"] = o.LastOpenedTime
	}
	if !IsNil(o.MinimumSystemVersion) {
		toSerialize["minimum_system_version"] = o.MinimumSystemVersion
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}
	return toSerialize, nil
}

type NullableSystemInsightsApps struct {
	value *SystemInsightsApps
	isSet bool
}

func (v NullableSystemInsightsApps) Get() *SystemInsightsApps {
	return v.value
}

func (v *NullableSystemInsightsApps) Set(val *SystemInsightsApps) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInsightsApps) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInsightsApps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInsightsApps(val *SystemInsightsApps) *NullableSystemInsightsApps {
	return &NullableSystemInsightsApps{value: val, isSet: true}
}

func (v NullableSystemInsightsApps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInsightsApps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



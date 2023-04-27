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

// checks if the SystemInsightsLinuxPackages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInsightsLinuxPackages{}

// SystemInsightsLinuxPackages struct for SystemInsightsLinuxPackages
type SystemInsightsLinuxPackages struct {
	Arch *string `json:"arch,omitempty"`
	InstallTime *int32 `json:"install_time,omitempty"`
	MaintainerOrVendor *string `json:"maintainer_or_vendor,omitempty"`
	MountNamespaceId *string `json:"mount_namespace_id,omitempty"`
	Name *string `json:"name,omitempty"`
	PackageFormat *string `json:"package_format,omitempty"`
	PackageGroupOrSection *string `json:"package_group_or_section,omitempty"`
	PidWithNamespace *int32 `json:"pid_with_namespace,omitempty"`
	ReleaseOrRevision *string `json:"release_or_revision,omitempty"`
	Size *string `json:"size,omitempty"`
	SystemId *string `json:"system_id,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemInsightsLinuxPackages SystemInsightsLinuxPackages

// NewSystemInsightsLinuxPackages instantiates a new SystemInsightsLinuxPackages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInsightsLinuxPackages() *SystemInsightsLinuxPackages {
	this := SystemInsightsLinuxPackages{}
	return &this
}

// NewSystemInsightsLinuxPackagesWithDefaults instantiates a new SystemInsightsLinuxPackages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInsightsLinuxPackagesWithDefaults() *SystemInsightsLinuxPackages {
	this := SystemInsightsLinuxPackages{}
	return &this
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetArch() string {
	if o == nil || IsNil(o.Arch) {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetArchOk() (*string, bool) {
	if o == nil || IsNil(o.Arch) {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasArch() bool {
	if o != nil && !IsNil(o.Arch) {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *SystemInsightsLinuxPackages) SetArch(v string) {
	o.Arch = &v
}

// GetInstallTime returns the InstallTime field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetInstallTime() int32 {
	if o == nil || IsNil(o.InstallTime) {
		var ret int32
		return ret
	}
	return *o.InstallTime
}

// GetInstallTimeOk returns a tuple with the InstallTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetInstallTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.InstallTime) {
		return nil, false
	}
	return o.InstallTime, true
}

// HasInstallTime returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasInstallTime() bool {
	if o != nil && !IsNil(o.InstallTime) {
		return true
	}

	return false
}

// SetInstallTime gets a reference to the given int32 and assigns it to the InstallTime field.
func (o *SystemInsightsLinuxPackages) SetInstallTime(v int32) {
	o.InstallTime = &v
}

// GetMaintainerOrVendor returns the MaintainerOrVendor field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetMaintainerOrVendor() string {
	if o == nil || IsNil(o.MaintainerOrVendor) {
		var ret string
		return ret
	}
	return *o.MaintainerOrVendor
}

// GetMaintainerOrVendorOk returns a tuple with the MaintainerOrVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetMaintainerOrVendorOk() (*string, bool) {
	if o == nil || IsNil(o.MaintainerOrVendor) {
		return nil, false
	}
	return o.MaintainerOrVendor, true
}

// HasMaintainerOrVendor returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasMaintainerOrVendor() bool {
	if o != nil && !IsNil(o.MaintainerOrVendor) {
		return true
	}

	return false
}

// SetMaintainerOrVendor gets a reference to the given string and assigns it to the MaintainerOrVendor field.
func (o *SystemInsightsLinuxPackages) SetMaintainerOrVendor(v string) {
	o.MaintainerOrVendor = &v
}

// GetMountNamespaceId returns the MountNamespaceId field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetMountNamespaceId() string {
	if o == nil || IsNil(o.MountNamespaceId) {
		var ret string
		return ret
	}
	return *o.MountNamespaceId
}

// GetMountNamespaceIdOk returns a tuple with the MountNamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetMountNamespaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MountNamespaceId) {
		return nil, false
	}
	return o.MountNamespaceId, true
}

// HasMountNamespaceId returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasMountNamespaceId() bool {
	if o != nil && !IsNil(o.MountNamespaceId) {
		return true
	}

	return false
}

// SetMountNamespaceId gets a reference to the given string and assigns it to the MountNamespaceId field.
func (o *SystemInsightsLinuxPackages) SetMountNamespaceId(v string) {
	o.MountNamespaceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SystemInsightsLinuxPackages) SetName(v string) {
	o.Name = &v
}

// GetPackageFormat returns the PackageFormat field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetPackageFormat() string {
	if o == nil || IsNil(o.PackageFormat) {
		var ret string
		return ret
	}
	return *o.PackageFormat
}

// GetPackageFormatOk returns a tuple with the PackageFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetPackageFormatOk() (*string, bool) {
	if o == nil || IsNil(o.PackageFormat) {
		return nil, false
	}
	return o.PackageFormat, true
}

// HasPackageFormat returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasPackageFormat() bool {
	if o != nil && !IsNil(o.PackageFormat) {
		return true
	}

	return false
}

// SetPackageFormat gets a reference to the given string and assigns it to the PackageFormat field.
func (o *SystemInsightsLinuxPackages) SetPackageFormat(v string) {
	o.PackageFormat = &v
}

// GetPackageGroupOrSection returns the PackageGroupOrSection field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetPackageGroupOrSection() string {
	if o == nil || IsNil(o.PackageGroupOrSection) {
		var ret string
		return ret
	}
	return *o.PackageGroupOrSection
}

// GetPackageGroupOrSectionOk returns a tuple with the PackageGroupOrSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetPackageGroupOrSectionOk() (*string, bool) {
	if o == nil || IsNil(o.PackageGroupOrSection) {
		return nil, false
	}
	return o.PackageGroupOrSection, true
}

// HasPackageGroupOrSection returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasPackageGroupOrSection() bool {
	if o != nil && !IsNil(o.PackageGroupOrSection) {
		return true
	}

	return false
}

// SetPackageGroupOrSection gets a reference to the given string and assigns it to the PackageGroupOrSection field.
func (o *SystemInsightsLinuxPackages) SetPackageGroupOrSection(v string) {
	o.PackageGroupOrSection = &v
}

// GetPidWithNamespace returns the PidWithNamespace field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetPidWithNamespace() int32 {
	if o == nil || IsNil(o.PidWithNamespace) {
		var ret int32
		return ret
	}
	return *o.PidWithNamespace
}

// GetPidWithNamespaceOk returns a tuple with the PidWithNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetPidWithNamespaceOk() (*int32, bool) {
	if o == nil || IsNil(o.PidWithNamespace) {
		return nil, false
	}
	return o.PidWithNamespace, true
}

// HasPidWithNamespace returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasPidWithNamespace() bool {
	if o != nil && !IsNil(o.PidWithNamespace) {
		return true
	}

	return false
}

// SetPidWithNamespace gets a reference to the given int32 and assigns it to the PidWithNamespace field.
func (o *SystemInsightsLinuxPackages) SetPidWithNamespace(v int32) {
	o.PidWithNamespace = &v
}

// GetReleaseOrRevision returns the ReleaseOrRevision field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetReleaseOrRevision() string {
	if o == nil || IsNil(o.ReleaseOrRevision) {
		var ret string
		return ret
	}
	return *o.ReleaseOrRevision
}

// GetReleaseOrRevisionOk returns a tuple with the ReleaseOrRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetReleaseOrRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseOrRevision) {
		return nil, false
	}
	return o.ReleaseOrRevision, true
}

// HasReleaseOrRevision returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasReleaseOrRevision() bool {
	if o != nil && !IsNil(o.ReleaseOrRevision) {
		return true
	}

	return false
}

// SetReleaseOrRevision gets a reference to the given string and assigns it to the ReleaseOrRevision field.
func (o *SystemInsightsLinuxPackages) SetReleaseOrRevision(v string) {
	o.ReleaseOrRevision = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *SystemInsightsLinuxPackages) SetSize(v string) {
	o.Size = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *SystemInsightsLinuxPackages) SetSystemId(v string) {
	o.SystemId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SystemInsightsLinuxPackages) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsLinuxPackages) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SystemInsightsLinuxPackages) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SystemInsightsLinuxPackages) SetVersion(v string) {
	o.Version = &v
}

func (o SystemInsightsLinuxPackages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInsightsLinuxPackages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Arch) {
		toSerialize["arch"] = o.Arch
	}
	if !IsNil(o.InstallTime) {
		toSerialize["install_time"] = o.InstallTime
	}
	if !IsNil(o.MaintainerOrVendor) {
		toSerialize["maintainer_or_vendor"] = o.MaintainerOrVendor
	}
	if !IsNil(o.MountNamespaceId) {
		toSerialize["mount_namespace_id"] = o.MountNamespaceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PackageFormat) {
		toSerialize["package_format"] = o.PackageFormat
	}
	if !IsNil(o.PackageGroupOrSection) {
		toSerialize["package_group_or_section"] = o.PackageGroupOrSection
	}
	if !IsNil(o.PidWithNamespace) {
		toSerialize["pid_with_namespace"] = o.PidWithNamespace
	}
	if !IsNil(o.ReleaseOrRevision) {
		toSerialize["release_or_revision"] = o.ReleaseOrRevision
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemInsightsLinuxPackages) UnmarshalJSON(bytes []byte) (err error) {
	varSystemInsightsLinuxPackages := _SystemInsightsLinuxPackages{}

	if err = json.Unmarshal(bytes, &varSystemInsightsLinuxPackages); err == nil {
		*o = SystemInsightsLinuxPackages(varSystemInsightsLinuxPackages)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "arch")
		delete(additionalProperties, "install_time")
		delete(additionalProperties, "maintainer_or_vendor")
		delete(additionalProperties, "mount_namespace_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "package_format")
		delete(additionalProperties, "package_group_or_section")
		delete(additionalProperties, "pid_with_namespace")
		delete(additionalProperties, "release_or_revision")
		delete(additionalProperties, "size")
		delete(additionalProperties, "system_id")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemInsightsLinuxPackages struct {
	value *SystemInsightsLinuxPackages
	isSet bool
}

func (v NullableSystemInsightsLinuxPackages) Get() *SystemInsightsLinuxPackages {
	return v.value
}

func (v *NullableSystemInsightsLinuxPackages) Set(val *SystemInsightsLinuxPackages) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInsightsLinuxPackages) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInsightsLinuxPackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInsightsLinuxPackages(val *SystemInsightsLinuxPackages) *NullableSystemInsightsLinuxPackages {
	return &NullableSystemInsightsLinuxPackages{value: val, isSet: true}
}

func (v NullableSystemInsightsLinuxPackages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInsightsLinuxPackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



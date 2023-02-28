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

// checks if the SoftwareAppGoogleAndroid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwareAppGoogleAndroid{}

// SoftwareAppGoogleAndroid googleAndroid is an optional attribute, it will only be present on apps with a 'setting' 'package_manager' type of 'GOOGLE_ANDROID'.
type SoftwareAppGoogleAndroid struct {
	// Whether this app is free, free with in-app purchases, or paid.
	AppPricing *string `json:"appPricing,omitempty"`
	// Latest version currently available for this app.
	AppVersion *string `json:"appVersion,omitempty"`
	// The name of the author of this app.
	Author *string `json:"author,omitempty"`
	// Controls the auto-update mode for the app.
	AutoUpdateMode *string `json:"autoUpdateMode,omitempty"`
	// The app category (e.g. COMMUNICATION, SOCIAL, etc.).
	Category *string `json:"category,omitempty"`
	// The content rating for this app.
	ContentRating *string `json:"contentRating,omitempty"`
	// The display mode of the web app.
	DisplayMode *string `json:"displayMode,omitempty"`
	// How and to whom the package is made available.
	DistributionChannel *string `json:"distributionChannel,omitempty"`
	// Full app description, if available.
	FullDescription *string `json:"fullDescription,omitempty"`
	// A link to an image that can be used as an icon for the app.
	IconUrl *string `json:"iconUrl,omitempty"`
	// The type of installation to perform for an app.
	InstallType *string `json:"installType,omitempty"`
	// The managed configurations template for the app.
	ManagedConfigurationTemplateId *string `json:"managedConfigurationTemplateId,omitempty"`
	// Indicates whether this app has managed properties or not.
	ManagedProperties *bool `json:"managedProperties,omitempty"`
	// The minimum Android SDK necessary to run the app.
	MinSdkVersion *int32 `json:"minSdkVersion,omitempty"`
	// The name of the app in the form enterprises/{enterprise}/applications/{packageName}.
	Name *string `json:"name,omitempty"`
	PermissionGrants []SoftwareAppPermissionGrants `json:"permissionGrants,omitempty"`
	// The policy for granting permission requests to apps.
	RuntimePermission *string `json:"runtimePermission,omitempty"`
	// The start URL, i.e. the URL that should load when the user opens the application. Applicable only for webapps.
	StartUrl *string `json:"startUrl,omitempty"`
	// Type of this android application.
	Type *string `json:"type,omitempty"`
	// The approximate time (within 7 days) the app was last published.
	UpdateTime *string `json:"updateTime,omitempty"`
	// The current version of the web app.
	VersionCode *int32 `json:"versionCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareAppGoogleAndroid SoftwareAppGoogleAndroid

// NewSoftwareAppGoogleAndroid instantiates a new SoftwareAppGoogleAndroid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareAppGoogleAndroid() *SoftwareAppGoogleAndroid {
	this := SoftwareAppGoogleAndroid{}
	return &this
}

// NewSoftwareAppGoogleAndroidWithDefaults instantiates a new SoftwareAppGoogleAndroid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareAppGoogleAndroidWithDefaults() *SoftwareAppGoogleAndroid {
	this := SoftwareAppGoogleAndroid{}
	return &this
}

// GetAppPricing returns the AppPricing field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetAppPricing() string {
	if o == nil || IsNil(o.AppPricing) {
		var ret string
		return ret
	}
	return *o.AppPricing
}

// GetAppPricingOk returns a tuple with the AppPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetAppPricingOk() (*string, bool) {
	if o == nil || IsNil(o.AppPricing) {
		return nil, false
	}
	return o.AppPricing, true
}

// HasAppPricing returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasAppPricing() bool {
	if o != nil && !IsNil(o.AppPricing) {
		return true
	}

	return false
}

// SetAppPricing gets a reference to the given string and assigns it to the AppPricing field.
func (o *SoftwareAppGoogleAndroid) SetAppPricing(v string) {
	o.AppPricing = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *SoftwareAppGoogleAndroid) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *SoftwareAppGoogleAndroid) SetAuthor(v string) {
	o.Author = &v
}

// GetAutoUpdateMode returns the AutoUpdateMode field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetAutoUpdateMode() string {
	if o == nil || IsNil(o.AutoUpdateMode) {
		var ret string
		return ret
	}
	return *o.AutoUpdateMode
}

// GetAutoUpdateModeOk returns a tuple with the AutoUpdateMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetAutoUpdateModeOk() (*string, bool) {
	if o == nil || IsNil(o.AutoUpdateMode) {
		return nil, false
	}
	return o.AutoUpdateMode, true
}

// HasAutoUpdateMode returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasAutoUpdateMode() bool {
	if o != nil && !IsNil(o.AutoUpdateMode) {
		return true
	}

	return false
}

// SetAutoUpdateMode gets a reference to the given string and assigns it to the AutoUpdateMode field.
func (o *SoftwareAppGoogleAndroid) SetAutoUpdateMode(v string) {
	o.AutoUpdateMode = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SoftwareAppGoogleAndroid) SetCategory(v string) {
	o.Category = &v
}

// GetContentRating returns the ContentRating field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetContentRating() string {
	if o == nil || IsNil(o.ContentRating) {
		var ret string
		return ret
	}
	return *o.ContentRating
}

// GetContentRatingOk returns a tuple with the ContentRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetContentRatingOk() (*string, bool) {
	if o == nil || IsNil(o.ContentRating) {
		return nil, false
	}
	return o.ContentRating, true
}

// HasContentRating returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasContentRating() bool {
	if o != nil && !IsNil(o.ContentRating) {
		return true
	}

	return false
}

// SetContentRating gets a reference to the given string and assigns it to the ContentRating field.
func (o *SoftwareAppGoogleAndroid) SetContentRating(v string) {
	o.ContentRating = &v
}

// GetDisplayMode returns the DisplayMode field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetDisplayMode() string {
	if o == nil || IsNil(o.DisplayMode) {
		var ret string
		return ret
	}
	return *o.DisplayMode
}

// GetDisplayModeOk returns a tuple with the DisplayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetDisplayModeOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayMode) {
		return nil, false
	}
	return o.DisplayMode, true
}

// HasDisplayMode returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasDisplayMode() bool {
	if o != nil && !IsNil(o.DisplayMode) {
		return true
	}

	return false
}

// SetDisplayMode gets a reference to the given string and assigns it to the DisplayMode field.
func (o *SoftwareAppGoogleAndroid) SetDisplayMode(v string) {
	o.DisplayMode = &v
}

// GetDistributionChannel returns the DistributionChannel field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetDistributionChannel() string {
	if o == nil || IsNil(o.DistributionChannel) {
		var ret string
		return ret
	}
	return *o.DistributionChannel
}

// GetDistributionChannelOk returns a tuple with the DistributionChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetDistributionChannelOk() (*string, bool) {
	if o == nil || IsNil(o.DistributionChannel) {
		return nil, false
	}
	return o.DistributionChannel, true
}

// HasDistributionChannel returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasDistributionChannel() bool {
	if o != nil && !IsNil(o.DistributionChannel) {
		return true
	}

	return false
}

// SetDistributionChannel gets a reference to the given string and assigns it to the DistributionChannel field.
func (o *SoftwareAppGoogleAndroid) SetDistributionChannel(v string) {
	o.DistributionChannel = &v
}

// GetFullDescription returns the FullDescription field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetFullDescription() string {
	if o == nil || IsNil(o.FullDescription) {
		var ret string
		return ret
	}
	return *o.FullDescription
}

// GetFullDescriptionOk returns a tuple with the FullDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetFullDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.FullDescription) {
		return nil, false
	}
	return o.FullDescription, true
}

// HasFullDescription returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasFullDescription() bool {
	if o != nil && !IsNil(o.FullDescription) {
		return true
	}

	return false
}

// SetFullDescription gets a reference to the given string and assigns it to the FullDescription field.
func (o *SoftwareAppGoogleAndroid) SetFullDescription(v string) {
	o.FullDescription = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *SoftwareAppGoogleAndroid) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetInstallType returns the InstallType field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetInstallType() string {
	if o == nil || IsNil(o.InstallType) {
		var ret string
		return ret
	}
	return *o.InstallType
}

// GetInstallTypeOk returns a tuple with the InstallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetInstallTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstallType) {
		return nil, false
	}
	return o.InstallType, true
}

// HasInstallType returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasInstallType() bool {
	if o != nil && !IsNil(o.InstallType) {
		return true
	}

	return false
}

// SetInstallType gets a reference to the given string and assigns it to the InstallType field.
func (o *SoftwareAppGoogleAndroid) SetInstallType(v string) {
	o.InstallType = &v
}

// GetManagedConfigurationTemplateId returns the ManagedConfigurationTemplateId field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetManagedConfigurationTemplateId() string {
	if o == nil || IsNil(o.ManagedConfigurationTemplateId) {
		var ret string
		return ret
	}
	return *o.ManagedConfigurationTemplateId
}

// GetManagedConfigurationTemplateIdOk returns a tuple with the ManagedConfigurationTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetManagedConfigurationTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedConfigurationTemplateId) {
		return nil, false
	}
	return o.ManagedConfigurationTemplateId, true
}

// HasManagedConfigurationTemplateId returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasManagedConfigurationTemplateId() bool {
	if o != nil && !IsNil(o.ManagedConfigurationTemplateId) {
		return true
	}

	return false
}

// SetManagedConfigurationTemplateId gets a reference to the given string and assigns it to the ManagedConfigurationTemplateId field.
func (o *SoftwareAppGoogleAndroid) SetManagedConfigurationTemplateId(v string) {
	o.ManagedConfigurationTemplateId = &v
}

// GetManagedProperties returns the ManagedProperties field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetManagedProperties() bool {
	if o == nil || IsNil(o.ManagedProperties) {
		var ret bool
		return ret
	}
	return *o.ManagedProperties
}

// GetManagedPropertiesOk returns a tuple with the ManagedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetManagedPropertiesOk() (*bool, bool) {
	if o == nil || IsNil(o.ManagedProperties) {
		return nil, false
	}
	return o.ManagedProperties, true
}

// HasManagedProperties returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasManagedProperties() bool {
	if o != nil && !IsNil(o.ManagedProperties) {
		return true
	}

	return false
}

// SetManagedProperties gets a reference to the given bool and assigns it to the ManagedProperties field.
func (o *SoftwareAppGoogleAndroid) SetManagedProperties(v bool) {
	o.ManagedProperties = &v
}

// GetMinSdkVersion returns the MinSdkVersion field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetMinSdkVersion() int32 {
	if o == nil || IsNil(o.MinSdkVersion) {
		var ret int32
		return ret
	}
	return *o.MinSdkVersion
}

// GetMinSdkVersionOk returns a tuple with the MinSdkVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetMinSdkVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSdkVersion) {
		return nil, false
	}
	return o.MinSdkVersion, true
}

// HasMinSdkVersion returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasMinSdkVersion() bool {
	if o != nil && !IsNil(o.MinSdkVersion) {
		return true
	}

	return false
}

// SetMinSdkVersion gets a reference to the given int32 and assigns it to the MinSdkVersion field.
func (o *SoftwareAppGoogleAndroid) SetMinSdkVersion(v int32) {
	o.MinSdkVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoftwareAppGoogleAndroid) SetName(v string) {
	o.Name = &v
}

// GetPermissionGrants returns the PermissionGrants field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetPermissionGrants() []SoftwareAppPermissionGrants {
	if o == nil || IsNil(o.PermissionGrants) {
		var ret []SoftwareAppPermissionGrants
		return ret
	}
	return o.PermissionGrants
}

// GetPermissionGrantsOk returns a tuple with the PermissionGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetPermissionGrantsOk() ([]SoftwareAppPermissionGrants, bool) {
	if o == nil || IsNil(o.PermissionGrants) {
		return nil, false
	}
	return o.PermissionGrants, true
}

// HasPermissionGrants returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasPermissionGrants() bool {
	if o != nil && !IsNil(o.PermissionGrants) {
		return true
	}

	return false
}

// SetPermissionGrants gets a reference to the given []SoftwareAppPermissionGrants and assigns it to the PermissionGrants field.
func (o *SoftwareAppGoogleAndroid) SetPermissionGrants(v []SoftwareAppPermissionGrants) {
	o.PermissionGrants = v
}

// GetRuntimePermission returns the RuntimePermission field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetRuntimePermission() string {
	if o == nil || IsNil(o.RuntimePermission) {
		var ret string
		return ret
	}
	return *o.RuntimePermission
}

// GetRuntimePermissionOk returns a tuple with the RuntimePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetRuntimePermissionOk() (*string, bool) {
	if o == nil || IsNil(o.RuntimePermission) {
		return nil, false
	}
	return o.RuntimePermission, true
}

// HasRuntimePermission returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasRuntimePermission() bool {
	if o != nil && !IsNil(o.RuntimePermission) {
		return true
	}

	return false
}

// SetRuntimePermission gets a reference to the given string and assigns it to the RuntimePermission field.
func (o *SoftwareAppGoogleAndroid) SetRuntimePermission(v string) {
	o.RuntimePermission = &v
}

// GetStartUrl returns the StartUrl field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetStartUrl() string {
	if o == nil || IsNil(o.StartUrl) {
		var ret string
		return ret
	}
	return *o.StartUrl
}

// GetStartUrlOk returns a tuple with the StartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetStartUrlOk() (*string, bool) {
	if o == nil || IsNil(o.StartUrl) {
		return nil, false
	}
	return o.StartUrl, true
}

// HasStartUrl returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasStartUrl() bool {
	if o != nil && !IsNil(o.StartUrl) {
		return true
	}

	return false
}

// SetStartUrl gets a reference to the given string and assigns it to the StartUrl field.
func (o *SoftwareAppGoogleAndroid) SetStartUrl(v string) {
	o.StartUrl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SoftwareAppGoogleAndroid) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *SoftwareAppGoogleAndroid) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetVersionCode returns the VersionCode field value if set, zero value otherwise.
func (o *SoftwareAppGoogleAndroid) GetVersionCode() int32 {
	if o == nil || IsNil(o.VersionCode) {
		var ret int32
		return ret
	}
	return *o.VersionCode
}

// GetVersionCodeOk returns a tuple with the VersionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareAppGoogleAndroid) GetVersionCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.VersionCode) {
		return nil, false
	}
	return o.VersionCode, true
}

// HasVersionCode returns a boolean if a field has been set.
func (o *SoftwareAppGoogleAndroid) HasVersionCode() bool {
	if o != nil && !IsNil(o.VersionCode) {
		return true
	}

	return false
}

// SetVersionCode gets a reference to the given int32 and assigns it to the VersionCode field.
func (o *SoftwareAppGoogleAndroid) SetVersionCode(v int32) {
	o.VersionCode = &v
}

func (o SoftwareAppGoogleAndroid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwareAppGoogleAndroid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppPricing) {
		toSerialize["appPricing"] = o.AppPricing
	}
	if !IsNil(o.AppVersion) {
		toSerialize["appVersion"] = o.AppVersion
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.AutoUpdateMode) {
		toSerialize["autoUpdateMode"] = o.AutoUpdateMode
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.ContentRating) {
		toSerialize["contentRating"] = o.ContentRating
	}
	if !IsNil(o.DisplayMode) {
		toSerialize["displayMode"] = o.DisplayMode
	}
	if !IsNil(o.DistributionChannel) {
		toSerialize["distributionChannel"] = o.DistributionChannel
	}
	if !IsNil(o.FullDescription) {
		toSerialize["fullDescription"] = o.FullDescription
	}
	if !IsNil(o.IconUrl) {
		toSerialize["iconUrl"] = o.IconUrl
	}
	if !IsNil(o.InstallType) {
		toSerialize["installType"] = o.InstallType
	}
	if !IsNil(o.ManagedConfigurationTemplateId) {
		toSerialize["managedConfigurationTemplateId"] = o.ManagedConfigurationTemplateId
	}
	if !IsNil(o.ManagedProperties) {
		toSerialize["managedProperties"] = o.ManagedProperties
	}
	if !IsNil(o.MinSdkVersion) {
		toSerialize["minSdkVersion"] = o.MinSdkVersion
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PermissionGrants) {
		toSerialize["permissionGrants"] = o.PermissionGrants
	}
	if !IsNil(o.RuntimePermission) {
		toSerialize["runtimePermission"] = o.RuntimePermission
	}
	if !IsNil(o.StartUrl) {
		toSerialize["startUrl"] = o.StartUrl
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.VersionCode) {
		toSerialize["versionCode"] = o.VersionCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SoftwareAppGoogleAndroid) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwareAppGoogleAndroid := _SoftwareAppGoogleAndroid{}

	if err = json.Unmarshal(bytes, &varSoftwareAppGoogleAndroid); err == nil {
		*o = SoftwareAppGoogleAndroid(varSoftwareAppGoogleAndroid)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appPricing")
		delete(additionalProperties, "appVersion")
		delete(additionalProperties, "author")
		delete(additionalProperties, "autoUpdateMode")
		delete(additionalProperties, "category")
		delete(additionalProperties, "contentRating")
		delete(additionalProperties, "displayMode")
		delete(additionalProperties, "distributionChannel")
		delete(additionalProperties, "fullDescription")
		delete(additionalProperties, "iconUrl")
		delete(additionalProperties, "installType")
		delete(additionalProperties, "managedConfigurationTemplateId")
		delete(additionalProperties, "managedProperties")
		delete(additionalProperties, "minSdkVersion")
		delete(additionalProperties, "name")
		delete(additionalProperties, "permissionGrants")
		delete(additionalProperties, "runtimePermission")
		delete(additionalProperties, "startUrl")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "versionCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwareAppGoogleAndroid struct {
	value *SoftwareAppGoogleAndroid
	isSet bool
}

func (v NullableSoftwareAppGoogleAndroid) Get() *SoftwareAppGoogleAndroid {
	return v.value
}

func (v *NullableSoftwareAppGoogleAndroid) Set(val *SoftwareAppGoogleAndroid) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareAppGoogleAndroid) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareAppGoogleAndroid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareAppGoogleAndroid(val *SoftwareAppGoogleAndroid) *NullableSoftwareAppGoogleAndroid {
	return &NullableSoftwareAppGoogleAndroid{value: val, isSet: true}
}

func (v NullableSoftwareAppGoogleAndroid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareAppGoogleAndroid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



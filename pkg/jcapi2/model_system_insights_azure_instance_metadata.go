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

// checks if the SystemInsightsAzureInstanceMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInsightsAzureInstanceMetadata{}

// SystemInsightsAzureInstanceMetadata struct for SystemInsightsAzureInstanceMetadata
type SystemInsightsAzureInstanceMetadata struct {
	CollectionTime *string `json:"collection_time,omitempty"`
	Location *string `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	Offer *string `json:"offer,omitempty"`
	OsType *string `json:"os_type,omitempty"`
	PlacementGroupId *string `json:"placement_group_id,omitempty"`
	PlatformFaultDomain *string `json:"platform_fault_domain,omitempty"`
	PlatformUpdateDomain *string `json:"platform_update_domain,omitempty"`
	Publisher *string `json:"publisher,omitempty"`
	ResourceGroupName *string `json:"resource_group_name,omitempty"`
	Sku *string `json:"sku,omitempty"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	SystemId *string `json:"system_id,omitempty"`
	Version *string `json:"version,omitempty"`
	VmId *string `json:"vm_id,omitempty"`
	VmScaleSetName *string `json:"vm_scale_set_name,omitempty"`
	VmSize *string `json:"vm_size,omitempty"`
	Zone *string `json:"zone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemInsightsAzureInstanceMetadata SystemInsightsAzureInstanceMetadata

// NewSystemInsightsAzureInstanceMetadata instantiates a new SystemInsightsAzureInstanceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInsightsAzureInstanceMetadata() *SystemInsightsAzureInstanceMetadata {
	this := SystemInsightsAzureInstanceMetadata{}
	return &this
}

// NewSystemInsightsAzureInstanceMetadataWithDefaults instantiates a new SystemInsightsAzureInstanceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInsightsAzureInstanceMetadataWithDefaults() *SystemInsightsAzureInstanceMetadata {
	this := SystemInsightsAzureInstanceMetadata{}
	return &this
}

// GetCollectionTime returns the CollectionTime field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetCollectionTime() string {
	if o == nil || IsNil(o.CollectionTime) {
		var ret string
		return ret
	}
	return *o.CollectionTime
}

// GetCollectionTimeOk returns a tuple with the CollectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetCollectionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionTime) {
		return nil, false
	}
	return o.CollectionTime, true
}

// HasCollectionTime returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasCollectionTime() bool {
	if o != nil && !IsNil(o.CollectionTime) {
		return true
	}

	return false
}

// SetCollectionTime gets a reference to the given string and assigns it to the CollectionTime field.
func (o *SystemInsightsAzureInstanceMetadata) SetCollectionTime(v string) {
	o.CollectionTime = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SystemInsightsAzureInstanceMetadata) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SystemInsightsAzureInstanceMetadata) SetName(v string) {
	o.Name = &v
}

// GetOffer returns the Offer field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetOffer() string {
	if o == nil || IsNil(o.Offer) {
		var ret string
		return ret
	}
	return *o.Offer
}

// GetOfferOk returns a tuple with the Offer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetOfferOk() (*string, bool) {
	if o == nil || IsNil(o.Offer) {
		return nil, false
	}
	return o.Offer, true
}

// HasOffer returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasOffer() bool {
	if o != nil && !IsNil(o.Offer) {
		return true
	}

	return false
}

// SetOffer gets a reference to the given string and assigns it to the Offer field.
func (o *SystemInsightsAzureInstanceMetadata) SetOffer(v string) {
	o.Offer = &v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetOsType() string {
	if o == nil || IsNil(o.OsType) {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetOsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OsType) {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasOsType() bool {
	if o != nil && !IsNil(o.OsType) {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *SystemInsightsAzureInstanceMetadata) SetOsType(v string) {
	o.OsType = &v
}

// GetPlacementGroupId returns the PlacementGroupId field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetPlacementGroupId() string {
	if o == nil || IsNil(o.PlacementGroupId) {
		var ret string
		return ret
	}
	return *o.PlacementGroupId
}

// GetPlacementGroupIdOk returns a tuple with the PlacementGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetPlacementGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlacementGroupId) {
		return nil, false
	}
	return o.PlacementGroupId, true
}

// HasPlacementGroupId returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasPlacementGroupId() bool {
	if o != nil && !IsNil(o.PlacementGroupId) {
		return true
	}

	return false
}

// SetPlacementGroupId gets a reference to the given string and assigns it to the PlacementGroupId field.
func (o *SystemInsightsAzureInstanceMetadata) SetPlacementGroupId(v string) {
	o.PlacementGroupId = &v
}

// GetPlatformFaultDomain returns the PlatformFaultDomain field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetPlatformFaultDomain() string {
	if o == nil || IsNil(o.PlatformFaultDomain) {
		var ret string
		return ret
	}
	return *o.PlatformFaultDomain
}

// GetPlatformFaultDomainOk returns a tuple with the PlatformFaultDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetPlatformFaultDomainOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformFaultDomain) {
		return nil, false
	}
	return o.PlatformFaultDomain, true
}

// HasPlatformFaultDomain returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasPlatformFaultDomain() bool {
	if o != nil && !IsNil(o.PlatformFaultDomain) {
		return true
	}

	return false
}

// SetPlatformFaultDomain gets a reference to the given string and assigns it to the PlatformFaultDomain field.
func (o *SystemInsightsAzureInstanceMetadata) SetPlatformFaultDomain(v string) {
	o.PlatformFaultDomain = &v
}

// GetPlatformUpdateDomain returns the PlatformUpdateDomain field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetPlatformUpdateDomain() string {
	if o == nil || IsNil(o.PlatformUpdateDomain) {
		var ret string
		return ret
	}
	return *o.PlatformUpdateDomain
}

// GetPlatformUpdateDomainOk returns a tuple with the PlatformUpdateDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetPlatformUpdateDomainOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformUpdateDomain) {
		return nil, false
	}
	return o.PlatformUpdateDomain, true
}

// HasPlatformUpdateDomain returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasPlatformUpdateDomain() bool {
	if o != nil && !IsNil(o.PlatformUpdateDomain) {
		return true
	}

	return false
}

// SetPlatformUpdateDomain gets a reference to the given string and assigns it to the PlatformUpdateDomain field.
func (o *SystemInsightsAzureInstanceMetadata) SetPlatformUpdateDomain(v string) {
	o.PlatformUpdateDomain = &v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetPublisher() string {
	if o == nil || IsNil(o.Publisher) {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetPublisherOk() (*string, bool) {
	if o == nil || IsNil(o.Publisher) {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasPublisher() bool {
	if o != nil && !IsNil(o.Publisher) {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *SystemInsightsAzureInstanceMetadata) SetPublisher(v string) {
	o.Publisher = &v
}

// GetResourceGroupName returns the ResourceGroupName field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetResourceGroupName() string {
	if o == nil || IsNil(o.ResourceGroupName) {
		var ret string
		return ret
	}
	return *o.ResourceGroupName
}

// GetResourceGroupNameOk returns a tuple with the ResourceGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetResourceGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceGroupName) {
		return nil, false
	}
	return o.ResourceGroupName, true
}

// HasResourceGroupName returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasResourceGroupName() bool {
	if o != nil && !IsNil(o.ResourceGroupName) {
		return true
	}

	return false
}

// SetResourceGroupName gets a reference to the given string and assigns it to the ResourceGroupName field.
func (o *SystemInsightsAzureInstanceMetadata) SetResourceGroupName(v string) {
	o.ResourceGroupName = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *SystemInsightsAzureInstanceMetadata) SetSku(v string) {
	o.Sku = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SystemInsightsAzureInstanceMetadata) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *SystemInsightsAzureInstanceMetadata) SetSystemId(v string) {
	o.SystemId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SystemInsightsAzureInstanceMetadata) SetVersion(v string) {
	o.Version = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetVmId() string {
	if o == nil || IsNil(o.VmId) {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetVmIdOk() (*string, bool) {
	if o == nil || IsNil(o.VmId) {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasVmId() bool {
	if o != nil && !IsNil(o.VmId) {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *SystemInsightsAzureInstanceMetadata) SetVmId(v string) {
	o.VmId = &v
}

// GetVmScaleSetName returns the VmScaleSetName field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetVmScaleSetName() string {
	if o == nil || IsNil(o.VmScaleSetName) {
		var ret string
		return ret
	}
	return *o.VmScaleSetName
}

// GetVmScaleSetNameOk returns a tuple with the VmScaleSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetVmScaleSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.VmScaleSetName) {
		return nil, false
	}
	return o.VmScaleSetName, true
}

// HasVmScaleSetName returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasVmScaleSetName() bool {
	if o != nil && !IsNil(o.VmScaleSetName) {
		return true
	}

	return false
}

// SetVmScaleSetName gets a reference to the given string and assigns it to the VmScaleSetName field.
func (o *SystemInsightsAzureInstanceMetadata) SetVmScaleSetName(v string) {
	o.VmScaleSetName = &v
}

// GetVmSize returns the VmSize field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetVmSize() string {
	if o == nil || IsNil(o.VmSize) {
		var ret string
		return ret
	}
	return *o.VmSize
}

// GetVmSizeOk returns a tuple with the VmSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetVmSizeOk() (*string, bool) {
	if o == nil || IsNil(o.VmSize) {
		return nil, false
	}
	return o.VmSize, true
}

// HasVmSize returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasVmSize() bool {
	if o != nil && !IsNil(o.VmSize) {
		return true
	}

	return false
}

// SetVmSize gets a reference to the given string and assigns it to the VmSize field.
func (o *SystemInsightsAzureInstanceMetadata) SetVmSize(v string) {
	o.VmSize = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *SystemInsightsAzureInstanceMetadata) GetZone() string {
	if o == nil || IsNil(o.Zone) {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsAzureInstanceMetadata) GetZoneOk() (*string, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *SystemInsightsAzureInstanceMetadata) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *SystemInsightsAzureInstanceMetadata) SetZone(v string) {
	o.Zone = &v
}

func (o SystemInsightsAzureInstanceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInsightsAzureInstanceMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionTime) {
		toSerialize["collection_time"] = o.CollectionTime
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Offer) {
		toSerialize["offer"] = o.Offer
	}
	if !IsNil(o.OsType) {
		toSerialize["os_type"] = o.OsType
	}
	if !IsNil(o.PlacementGroupId) {
		toSerialize["placement_group_id"] = o.PlacementGroupId
	}
	if !IsNil(o.PlatformFaultDomain) {
		toSerialize["platform_fault_domain"] = o.PlatformFaultDomain
	}
	if !IsNil(o.PlatformUpdateDomain) {
		toSerialize["platform_update_domain"] = o.PlatformUpdateDomain
	}
	if !IsNil(o.Publisher) {
		toSerialize["publisher"] = o.Publisher
	}
	if !IsNil(o.ResourceGroupName) {
		toSerialize["resource_group_name"] = o.ResourceGroupName
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VmId) {
		toSerialize["vm_id"] = o.VmId
	}
	if !IsNil(o.VmScaleSetName) {
		toSerialize["vm_scale_set_name"] = o.VmScaleSetName
	}
	if !IsNil(o.VmSize) {
		toSerialize["vm_size"] = o.VmSize
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemInsightsAzureInstanceMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varSystemInsightsAzureInstanceMetadata := _SystemInsightsAzureInstanceMetadata{}

	if err = json.Unmarshal(bytes, &varSystemInsightsAzureInstanceMetadata); err == nil {
		*o = SystemInsightsAzureInstanceMetadata(varSystemInsightsAzureInstanceMetadata)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "collection_time")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "offer")
		delete(additionalProperties, "os_type")
		delete(additionalProperties, "placement_group_id")
		delete(additionalProperties, "platform_fault_domain")
		delete(additionalProperties, "platform_update_domain")
		delete(additionalProperties, "publisher")
		delete(additionalProperties, "resource_group_name")
		delete(additionalProperties, "sku")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "system_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "vm_id")
		delete(additionalProperties, "vm_scale_set_name")
		delete(additionalProperties, "vm_size")
		delete(additionalProperties, "zone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemInsightsAzureInstanceMetadata struct {
	value *SystemInsightsAzureInstanceMetadata
	isSet bool
}

func (v NullableSystemInsightsAzureInstanceMetadata) Get() *SystemInsightsAzureInstanceMetadata {
	return v.value
}

func (v *NullableSystemInsightsAzureInstanceMetadata) Set(val *SystemInsightsAzureInstanceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInsightsAzureInstanceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInsightsAzureInstanceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInsightsAzureInstanceMetadata(val *SystemInsightsAzureInstanceMetadata) *NullableSystemInsightsAzureInstanceMetadata {
	return &NullableSystemInsightsAzureInstanceMetadata{value: val, isSet: true}
}

func (v NullableSystemInsightsAzureInstanceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInsightsAzureInstanceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



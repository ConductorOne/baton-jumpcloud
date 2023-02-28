/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jcapi

import (
	"encoding/json"
)

// checks if the SystemInsightsInterfaceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInsightsInterfaceDetails{}

// SystemInsightsInterfaceDetails struct for SystemInsightsInterfaceDetails
type SystemInsightsInterfaceDetails struct {
	Collisions *string `json:"collisions,omitempty"`
	ConnectionId *string `json:"connection_id,omitempty"`
	ConnectionStatus *string `json:"connection_status,omitempty"`
	Description *string `json:"description,omitempty"`
	DhcpEnabled *int32 `json:"dhcp_enabled,omitempty"`
	DhcpLeaseExpires *string `json:"dhcp_lease_expires,omitempty"`
	DhcpLeaseObtained *string `json:"dhcp_lease_obtained,omitempty"`
	DhcpServer *string `json:"dhcp_server,omitempty"`
	DnsDomain *string `json:"dns_domain,omitempty"`
	DnsDomainSuffixSearchOrder *string `json:"dns_domain_suffix_search_order,omitempty"`
	DnsHostName *string `json:"dns_host_name,omitempty"`
	DnsServerSearchOrder *string `json:"dns_server_search_order,omitempty"`
	Enabled *int32 `json:"enabled,omitempty"`
	Flags *int32 `json:"flags,omitempty"`
	FriendlyName *string `json:"friendly_name,omitempty"`
	Ibytes *string `json:"ibytes,omitempty"`
	Idrops *string `json:"idrops,omitempty"`
	Ierrors *string `json:"ierrors,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Ipackets *string `json:"ipackets,omitempty"`
	LastChange *string `json:"last_change,omitempty"`
	LinkSpeed *string `json:"link_speed,omitempty"`
	Mac *string `json:"mac,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	Metric *int32 `json:"metric,omitempty"`
	Mtu *int32 `json:"mtu,omitempty"`
	Obytes *string `json:"obytes,omitempty"`
	Odrops *string `json:"odrops,omitempty"`
	Oerrors *string `json:"oerrors,omitempty"`
	Opackets *string `json:"opackets,omitempty"`
	PciSlot *string `json:"pci_slot,omitempty"`
	PhysicalAdapter *int32 `json:"physical_adapter,omitempty"`
	Service *string `json:"service,omitempty"`
	Speed *int32 `json:"speed,omitempty"`
	SystemId *string `json:"system_id,omitempty"`
	Type *int32 `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemInsightsInterfaceDetails SystemInsightsInterfaceDetails

// NewSystemInsightsInterfaceDetails instantiates a new SystemInsightsInterfaceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInsightsInterfaceDetails() *SystemInsightsInterfaceDetails {
	this := SystemInsightsInterfaceDetails{}
	return &this
}

// NewSystemInsightsInterfaceDetailsWithDefaults instantiates a new SystemInsightsInterfaceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInsightsInterfaceDetailsWithDefaults() *SystemInsightsInterfaceDetails {
	this := SystemInsightsInterfaceDetails{}
	return &this
}

// GetCollisions returns the Collisions field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetCollisions() string {
	if o == nil || IsNil(o.Collisions) {
		var ret string
		return ret
	}
	return *o.Collisions
}

// GetCollisionsOk returns a tuple with the Collisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetCollisionsOk() (*string, bool) {
	if o == nil || IsNil(o.Collisions) {
		return nil, false
	}
	return o.Collisions, true
}

// HasCollisions returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasCollisions() bool {
	if o != nil && !IsNil(o.Collisions) {
		return true
	}

	return false
}

// SetCollisions gets a reference to the given string and assigns it to the Collisions field.
func (o *SystemInsightsInterfaceDetails) SetCollisions(v string) {
	o.Collisions = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *SystemInsightsInterfaceDetails) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetConnectionStatus() string {
	if o == nil || IsNil(o.ConnectionStatus) {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetConnectionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionStatus) {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasConnectionStatus() bool {
	if o != nil && !IsNil(o.ConnectionStatus) {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *SystemInsightsInterfaceDetails) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SystemInsightsInterfaceDetails) SetDescription(v string) {
	o.Description = &v
}

// GetDhcpEnabled returns the DhcpEnabled field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetDhcpEnabled() int32 {
	if o == nil || IsNil(o.DhcpEnabled) {
		var ret int32
		return ret
	}
	return *o.DhcpEnabled
}

// GetDhcpEnabledOk returns a tuple with the DhcpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetDhcpEnabledOk() (*int32, bool) {
	if o == nil || IsNil(o.DhcpEnabled) {
		return nil, false
	}
	return o.DhcpEnabled, true
}

// HasDhcpEnabled returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasDhcpEnabled() bool {
	if o != nil && !IsNil(o.DhcpEnabled) {
		return true
	}

	return false
}

// SetDhcpEnabled gets a reference to the given int32 and assigns it to the DhcpEnabled field.
func (o *SystemInsightsInterfaceDetails) SetDhcpEnabled(v int32) {
	o.DhcpEnabled = &v
}

// GetDhcpLeaseExpires returns the DhcpLeaseExpires field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetDhcpLeaseExpires() string {
	if o == nil || IsNil(o.DhcpLeaseExpires) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseExpires
}

// GetDhcpLeaseExpiresOk returns a tuple with the DhcpLeaseExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetDhcpLeaseExpiresOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpLeaseExpires) {
		return nil, false
	}
	return o.DhcpLeaseExpires, true
}

// HasDhcpLeaseExpires returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasDhcpLeaseExpires() bool {
	if o != nil && !IsNil(o.DhcpLeaseExpires) {
		return true
	}

	return false
}

// SetDhcpLeaseExpires gets a reference to the given string and assigns it to the DhcpLeaseExpires field.
func (o *SystemInsightsInterfaceDetails) SetDhcpLeaseExpires(v string) {
	o.DhcpLeaseExpires = &v
}

// GetDhcpLeaseObtained returns the DhcpLeaseObtained field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetDhcpLeaseObtained() string {
	if o == nil || IsNil(o.DhcpLeaseObtained) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseObtained
}

// GetDhcpLeaseObtainedOk returns a tuple with the DhcpLeaseObtained field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetDhcpLeaseObtainedOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpLeaseObtained) {
		return nil, false
	}
	return o.DhcpLeaseObtained, true
}

// HasDhcpLeaseObtained returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasDhcpLeaseObtained() bool {
	if o != nil && !IsNil(o.DhcpLeaseObtained) {
		return true
	}

	return false
}

// SetDhcpLeaseObtained gets a reference to the given string and assigns it to the DhcpLeaseObtained field.
func (o *SystemInsightsInterfaceDetails) SetDhcpLeaseObtained(v string) {
	o.DhcpLeaseObtained = &v
}

// GetDhcpServer returns the DhcpServer field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetDhcpServer() string {
	if o == nil || IsNil(o.DhcpServer) {
		var ret string
		return ret
	}
	return *o.DhcpServer
}

// GetDhcpServerOk returns a tuple with the DhcpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetDhcpServerOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpServer) {
		return nil, false
	}
	return o.DhcpServer, true
}

// HasDhcpServer returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasDhcpServer() bool {
	if o != nil && !IsNil(o.DhcpServer) {
		return true
	}

	return false
}

// SetDhcpServer gets a reference to the given string and assigns it to the DhcpServer field.
func (o *SystemInsightsInterfaceDetails) SetDhcpServer(v string) {
	o.DhcpServer = &v
}

// GetDnsDomain returns the DnsDomain field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetDnsDomain() string {
	if o == nil || IsNil(o.DnsDomain) {
		var ret string
		return ret
	}
	return *o.DnsDomain
}

// GetDnsDomainOk returns a tuple with the DnsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetDnsDomainOk() (*string, bool) {
	if o == nil || IsNil(o.DnsDomain) {
		return nil, false
	}
	return o.DnsDomain, true
}

// HasDnsDomain returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasDnsDomain() bool {
	if o != nil && !IsNil(o.DnsDomain) {
		return true
	}

	return false
}

// SetDnsDomain gets a reference to the given string and assigns it to the DnsDomain field.
func (o *SystemInsightsInterfaceDetails) SetDnsDomain(v string) {
	o.DnsDomain = &v
}

// GetDnsDomainSuffixSearchOrder returns the DnsDomainSuffixSearchOrder field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetDnsDomainSuffixSearchOrder() string {
	if o == nil || IsNil(o.DnsDomainSuffixSearchOrder) {
		var ret string
		return ret
	}
	return *o.DnsDomainSuffixSearchOrder
}

// GetDnsDomainSuffixSearchOrderOk returns a tuple with the DnsDomainSuffixSearchOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetDnsDomainSuffixSearchOrderOk() (*string, bool) {
	if o == nil || IsNil(o.DnsDomainSuffixSearchOrder) {
		return nil, false
	}
	return o.DnsDomainSuffixSearchOrder, true
}

// HasDnsDomainSuffixSearchOrder returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasDnsDomainSuffixSearchOrder() bool {
	if o != nil && !IsNil(o.DnsDomainSuffixSearchOrder) {
		return true
	}

	return false
}

// SetDnsDomainSuffixSearchOrder gets a reference to the given string and assigns it to the DnsDomainSuffixSearchOrder field.
func (o *SystemInsightsInterfaceDetails) SetDnsDomainSuffixSearchOrder(v string) {
	o.DnsDomainSuffixSearchOrder = &v
}

// GetDnsHostName returns the DnsHostName field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetDnsHostName() string {
	if o == nil || IsNil(o.DnsHostName) {
		var ret string
		return ret
	}
	return *o.DnsHostName
}

// GetDnsHostNameOk returns a tuple with the DnsHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetDnsHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.DnsHostName) {
		return nil, false
	}
	return o.DnsHostName, true
}

// HasDnsHostName returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasDnsHostName() bool {
	if o != nil && !IsNil(o.DnsHostName) {
		return true
	}

	return false
}

// SetDnsHostName gets a reference to the given string and assigns it to the DnsHostName field.
func (o *SystemInsightsInterfaceDetails) SetDnsHostName(v string) {
	o.DnsHostName = &v
}

// GetDnsServerSearchOrder returns the DnsServerSearchOrder field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetDnsServerSearchOrder() string {
	if o == nil || IsNil(o.DnsServerSearchOrder) {
		var ret string
		return ret
	}
	return *o.DnsServerSearchOrder
}

// GetDnsServerSearchOrderOk returns a tuple with the DnsServerSearchOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetDnsServerSearchOrderOk() (*string, bool) {
	if o == nil || IsNil(o.DnsServerSearchOrder) {
		return nil, false
	}
	return o.DnsServerSearchOrder, true
}

// HasDnsServerSearchOrder returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasDnsServerSearchOrder() bool {
	if o != nil && !IsNil(o.DnsServerSearchOrder) {
		return true
	}

	return false
}

// SetDnsServerSearchOrder gets a reference to the given string and assigns it to the DnsServerSearchOrder field.
func (o *SystemInsightsInterfaceDetails) SetDnsServerSearchOrder(v string) {
	o.DnsServerSearchOrder = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetEnabled() int32 {
	if o == nil || IsNil(o.Enabled) {
		var ret int32
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetEnabledOk() (*int32, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given int32 and assigns it to the Enabled field.
func (o *SystemInsightsInterfaceDetails) SetEnabled(v int32) {
	o.Enabled = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetFlags() int32 {
	if o == nil || IsNil(o.Flags) {
		var ret int32
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int32 and assigns it to the Flags field.
func (o *SystemInsightsInterfaceDetails) SetFlags(v int32) {
	o.Flags = &v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName) {
		var ret string
		return ret
	}
	return *o.FriendlyName
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetFriendlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FriendlyName) {
		return nil, false
	}
	return o.FriendlyName, true
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasFriendlyName() bool {
	if o != nil && !IsNil(o.FriendlyName) {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given string and assigns it to the FriendlyName field.
func (o *SystemInsightsInterfaceDetails) SetFriendlyName(v string) {
	o.FriendlyName = &v
}

// GetIbytes returns the Ibytes field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetIbytes() string {
	if o == nil || IsNil(o.Ibytes) {
		var ret string
		return ret
	}
	return *o.Ibytes
}

// GetIbytesOk returns a tuple with the Ibytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetIbytesOk() (*string, bool) {
	if o == nil || IsNil(o.Ibytes) {
		return nil, false
	}
	return o.Ibytes, true
}

// HasIbytes returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasIbytes() bool {
	if o != nil && !IsNil(o.Ibytes) {
		return true
	}

	return false
}

// SetIbytes gets a reference to the given string and assigns it to the Ibytes field.
func (o *SystemInsightsInterfaceDetails) SetIbytes(v string) {
	o.Ibytes = &v
}

// GetIdrops returns the Idrops field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetIdrops() string {
	if o == nil || IsNil(o.Idrops) {
		var ret string
		return ret
	}
	return *o.Idrops
}

// GetIdropsOk returns a tuple with the Idrops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetIdropsOk() (*string, bool) {
	if o == nil || IsNil(o.Idrops) {
		return nil, false
	}
	return o.Idrops, true
}

// HasIdrops returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasIdrops() bool {
	if o != nil && !IsNil(o.Idrops) {
		return true
	}

	return false
}

// SetIdrops gets a reference to the given string and assigns it to the Idrops field.
func (o *SystemInsightsInterfaceDetails) SetIdrops(v string) {
	o.Idrops = &v
}

// GetIerrors returns the Ierrors field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetIerrors() string {
	if o == nil || IsNil(o.Ierrors) {
		var ret string
		return ret
	}
	return *o.Ierrors
}

// GetIerrorsOk returns a tuple with the Ierrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetIerrorsOk() (*string, bool) {
	if o == nil || IsNil(o.Ierrors) {
		return nil, false
	}
	return o.Ierrors, true
}

// HasIerrors returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasIerrors() bool {
	if o != nil && !IsNil(o.Ierrors) {
		return true
	}

	return false
}

// SetIerrors gets a reference to the given string and assigns it to the Ierrors field.
func (o *SystemInsightsInterfaceDetails) SetIerrors(v string) {
	o.Ierrors = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *SystemInsightsInterfaceDetails) SetInterface(v string) {
	o.Interface = &v
}

// GetIpackets returns the Ipackets field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetIpackets() string {
	if o == nil || IsNil(o.Ipackets) {
		var ret string
		return ret
	}
	return *o.Ipackets
}

// GetIpacketsOk returns a tuple with the Ipackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetIpacketsOk() (*string, bool) {
	if o == nil || IsNil(o.Ipackets) {
		return nil, false
	}
	return o.Ipackets, true
}

// HasIpackets returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasIpackets() bool {
	if o != nil && !IsNil(o.Ipackets) {
		return true
	}

	return false
}

// SetIpackets gets a reference to the given string and assigns it to the Ipackets field.
func (o *SystemInsightsInterfaceDetails) SetIpackets(v string) {
	o.Ipackets = &v
}

// GetLastChange returns the LastChange field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetLastChange() string {
	if o == nil || IsNil(o.LastChange) {
		var ret string
		return ret
	}
	return *o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetLastChangeOk() (*string, bool) {
	if o == nil || IsNil(o.LastChange) {
		return nil, false
	}
	return o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasLastChange() bool {
	if o != nil && !IsNil(o.LastChange) {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given string and assigns it to the LastChange field.
func (o *SystemInsightsInterfaceDetails) SetLastChange(v string) {
	o.LastChange = &v
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetLinkSpeed() string {
	if o == nil || IsNil(o.LinkSpeed) {
		var ret string
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetLinkSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.LinkSpeed) {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasLinkSpeed() bool {
	if o != nil && !IsNil(o.LinkSpeed) {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given string and assigns it to the LinkSpeed field.
func (o *SystemInsightsInterfaceDetails) SetLinkSpeed(v string) {
	o.LinkSpeed = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *SystemInsightsInterfaceDetails) SetMac(v string) {
	o.Mac = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetManufacturer() string {
	if o == nil || IsNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetManufacturerOk() (*string, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *SystemInsightsInterfaceDetails) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetMetric() int32 {
	if o == nil || IsNil(o.Metric) {
		var ret int32
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetMetricOk() (*int32, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given int32 and assigns it to the Metric field.
func (o *SystemInsightsInterfaceDetails) SetMetric(v int32) {
	o.Metric = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu) {
		var ret int32
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetMtuOk() (*int32, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int32 and assigns it to the Mtu field.
func (o *SystemInsightsInterfaceDetails) SetMtu(v int32) {
	o.Mtu = &v
}

// GetObytes returns the Obytes field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetObytes() string {
	if o == nil || IsNil(o.Obytes) {
		var ret string
		return ret
	}
	return *o.Obytes
}

// GetObytesOk returns a tuple with the Obytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetObytesOk() (*string, bool) {
	if o == nil || IsNil(o.Obytes) {
		return nil, false
	}
	return o.Obytes, true
}

// HasObytes returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasObytes() bool {
	if o != nil && !IsNil(o.Obytes) {
		return true
	}

	return false
}

// SetObytes gets a reference to the given string and assigns it to the Obytes field.
func (o *SystemInsightsInterfaceDetails) SetObytes(v string) {
	o.Obytes = &v
}

// GetOdrops returns the Odrops field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetOdrops() string {
	if o == nil || IsNil(o.Odrops) {
		var ret string
		return ret
	}
	return *o.Odrops
}

// GetOdropsOk returns a tuple with the Odrops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetOdropsOk() (*string, bool) {
	if o == nil || IsNil(o.Odrops) {
		return nil, false
	}
	return o.Odrops, true
}

// HasOdrops returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasOdrops() bool {
	if o != nil && !IsNil(o.Odrops) {
		return true
	}

	return false
}

// SetOdrops gets a reference to the given string and assigns it to the Odrops field.
func (o *SystemInsightsInterfaceDetails) SetOdrops(v string) {
	o.Odrops = &v
}

// GetOerrors returns the Oerrors field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetOerrors() string {
	if o == nil || IsNil(o.Oerrors) {
		var ret string
		return ret
	}
	return *o.Oerrors
}

// GetOerrorsOk returns a tuple with the Oerrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetOerrorsOk() (*string, bool) {
	if o == nil || IsNil(o.Oerrors) {
		return nil, false
	}
	return o.Oerrors, true
}

// HasOerrors returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasOerrors() bool {
	if o != nil && !IsNil(o.Oerrors) {
		return true
	}

	return false
}

// SetOerrors gets a reference to the given string and assigns it to the Oerrors field.
func (o *SystemInsightsInterfaceDetails) SetOerrors(v string) {
	o.Oerrors = &v
}

// GetOpackets returns the Opackets field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetOpackets() string {
	if o == nil || IsNil(o.Opackets) {
		var ret string
		return ret
	}
	return *o.Opackets
}

// GetOpacketsOk returns a tuple with the Opackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetOpacketsOk() (*string, bool) {
	if o == nil || IsNil(o.Opackets) {
		return nil, false
	}
	return o.Opackets, true
}

// HasOpackets returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasOpackets() bool {
	if o != nil && !IsNil(o.Opackets) {
		return true
	}

	return false
}

// SetOpackets gets a reference to the given string and assigns it to the Opackets field.
func (o *SystemInsightsInterfaceDetails) SetOpackets(v string) {
	o.Opackets = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetPciSlot() string {
	if o == nil || IsNil(o.PciSlot) {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetPciSlotOk() (*string, bool) {
	if o == nil || IsNil(o.PciSlot) {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasPciSlot() bool {
	if o != nil && !IsNil(o.PciSlot) {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *SystemInsightsInterfaceDetails) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetPhysicalAdapter returns the PhysicalAdapter field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetPhysicalAdapter() int32 {
	if o == nil || IsNil(o.PhysicalAdapter) {
		var ret int32
		return ret
	}
	return *o.PhysicalAdapter
}

// GetPhysicalAdapterOk returns a tuple with the PhysicalAdapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetPhysicalAdapterOk() (*int32, bool) {
	if o == nil || IsNil(o.PhysicalAdapter) {
		return nil, false
	}
	return o.PhysicalAdapter, true
}

// HasPhysicalAdapter returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasPhysicalAdapter() bool {
	if o != nil && !IsNil(o.PhysicalAdapter) {
		return true
	}

	return false
}

// SetPhysicalAdapter gets a reference to the given int32 and assigns it to the PhysicalAdapter field.
func (o *SystemInsightsInterfaceDetails) SetPhysicalAdapter(v int32) {
	o.PhysicalAdapter = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SystemInsightsInterfaceDetails) SetService(v string) {
	o.Service = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *SystemInsightsInterfaceDetails) SetSpeed(v int32) {
	o.Speed = &v
}

// GetSystemId returns the SystemId field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetSystemId() string {
	if o == nil || IsNil(o.SystemId) {
		var ret string
		return ret
	}
	return *o.SystemId
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SystemId) {
		return nil, false
	}
	return o.SystemId, true
}

// HasSystemId returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasSystemId() bool {
	if o != nil && !IsNil(o.SystemId) {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given string and assigns it to the SystemId field.
func (o *SystemInsightsInterfaceDetails) SetSystemId(v string) {
	o.SystemId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SystemInsightsInterfaceDetails) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInsightsInterfaceDetails) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SystemInsightsInterfaceDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *SystemInsightsInterfaceDetails) SetType(v int32) {
	o.Type = &v
}

func (o SystemInsightsInterfaceDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInsightsInterfaceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collisions) {
		toSerialize["collisions"] = o.Collisions
	}
	if !IsNil(o.ConnectionId) {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if !IsNil(o.ConnectionStatus) {
		toSerialize["connection_status"] = o.ConnectionStatus
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DhcpEnabled) {
		toSerialize["dhcp_enabled"] = o.DhcpEnabled
	}
	if !IsNil(o.DhcpLeaseExpires) {
		toSerialize["dhcp_lease_expires"] = o.DhcpLeaseExpires
	}
	if !IsNil(o.DhcpLeaseObtained) {
		toSerialize["dhcp_lease_obtained"] = o.DhcpLeaseObtained
	}
	if !IsNil(o.DhcpServer) {
		toSerialize["dhcp_server"] = o.DhcpServer
	}
	if !IsNil(o.DnsDomain) {
		toSerialize["dns_domain"] = o.DnsDomain
	}
	if !IsNil(o.DnsDomainSuffixSearchOrder) {
		toSerialize["dns_domain_suffix_search_order"] = o.DnsDomainSuffixSearchOrder
	}
	if !IsNil(o.DnsHostName) {
		toSerialize["dns_host_name"] = o.DnsHostName
	}
	if !IsNil(o.DnsServerSearchOrder) {
		toSerialize["dns_server_search_order"] = o.DnsServerSearchOrder
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.FriendlyName) {
		toSerialize["friendly_name"] = o.FriendlyName
	}
	if !IsNil(o.Ibytes) {
		toSerialize["ibytes"] = o.Ibytes
	}
	if !IsNil(o.Idrops) {
		toSerialize["idrops"] = o.Idrops
	}
	if !IsNil(o.Ierrors) {
		toSerialize["ierrors"] = o.Ierrors
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.Ipackets) {
		toSerialize["ipackets"] = o.Ipackets
	}
	if !IsNil(o.LastChange) {
		toSerialize["last_change"] = o.LastChange
	}
	if !IsNil(o.LinkSpeed) {
		toSerialize["link_speed"] = o.LinkSpeed
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	if !IsNil(o.Obytes) {
		toSerialize["obytes"] = o.Obytes
	}
	if !IsNil(o.Odrops) {
		toSerialize["odrops"] = o.Odrops
	}
	if !IsNil(o.Oerrors) {
		toSerialize["oerrors"] = o.Oerrors
	}
	if !IsNil(o.Opackets) {
		toSerialize["opackets"] = o.Opackets
	}
	if !IsNil(o.PciSlot) {
		toSerialize["pci_slot"] = o.PciSlot
	}
	if !IsNil(o.PhysicalAdapter) {
		toSerialize["physical_adapter"] = o.PhysicalAdapter
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.SystemId) {
		toSerialize["system_id"] = o.SystemId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemInsightsInterfaceDetails) UnmarshalJSON(bytes []byte) (err error) {
	varSystemInsightsInterfaceDetails := _SystemInsightsInterfaceDetails{}

	if err = json.Unmarshal(bytes, &varSystemInsightsInterfaceDetails); err == nil {
		*o = SystemInsightsInterfaceDetails(varSystemInsightsInterfaceDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "collisions")
		delete(additionalProperties, "connection_id")
		delete(additionalProperties, "connection_status")
		delete(additionalProperties, "description")
		delete(additionalProperties, "dhcp_enabled")
		delete(additionalProperties, "dhcp_lease_expires")
		delete(additionalProperties, "dhcp_lease_obtained")
		delete(additionalProperties, "dhcp_server")
		delete(additionalProperties, "dns_domain")
		delete(additionalProperties, "dns_domain_suffix_search_order")
		delete(additionalProperties, "dns_host_name")
		delete(additionalProperties, "dns_server_search_order")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "flags")
		delete(additionalProperties, "friendly_name")
		delete(additionalProperties, "ibytes")
		delete(additionalProperties, "idrops")
		delete(additionalProperties, "ierrors")
		delete(additionalProperties, "interface")
		delete(additionalProperties, "ipackets")
		delete(additionalProperties, "last_change")
		delete(additionalProperties, "link_speed")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "metric")
		delete(additionalProperties, "mtu")
		delete(additionalProperties, "obytes")
		delete(additionalProperties, "odrops")
		delete(additionalProperties, "oerrors")
		delete(additionalProperties, "opackets")
		delete(additionalProperties, "pci_slot")
		delete(additionalProperties, "physical_adapter")
		delete(additionalProperties, "service")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "system_id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemInsightsInterfaceDetails struct {
	value *SystemInsightsInterfaceDetails
	isSet bool
}

func (v NullableSystemInsightsInterfaceDetails) Get() *SystemInsightsInterfaceDetails {
	return v.value
}

func (v *NullableSystemInsightsInterfaceDetails) Set(val *SystemInsightsInterfaceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInsightsInterfaceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInsightsInterfaceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInsightsInterfaceDetails(val *SystemInsightsInterfaceDetails) *NullableSystemInsightsInterfaceDetails {
	return &NullableSystemInsightsInterfaceDetails{value: val, isSet: true}
}

func (v NullableSystemInsightsInterfaceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInsightsInterfaceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



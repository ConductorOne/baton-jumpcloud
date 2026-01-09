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

// checks if the CloudInsightsEventsListEventDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudInsightsEventsListEventDetails{}

// CloudInsightsEventsListEventDetails Properties of a Cloud Insights trail GET list response event detail object
type CloudInsightsEventsListEventDetails struct {
	Addendum *CloudInsightsEventsListEventDetailsAddendum `json:"addendum,omitempty"`
	AdditionalEventData *CloudInsightsEventsListEventDetailsAdditionalEventData `json:"additionalEventData,omitempty"`
	Annotation *string `json:"annotation,omitempty"`
	ApiVersion *string `json:"apiVersion,omitempty"`
	AwsRegion *string `json:"awsRegion,omitempty"`
	EdgeDeviceDetails *CloudInsightsEventsListEventDetailsEdgeDeviceDetails `json:"edgeDeviceDetails,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	EventCategory *string `json:"eventCategory,omitempty"`
	EventId *string `json:"eventId,omitempty"`
	EventName *string `json:"eventName,omitempty"`
	EventSource *string `json:"eventSource,omitempty"`
	EventTime *CloudInsightsEventsListEventDetailsEventTime `json:"eventTime,omitempty"`
	EventType *string `json:"eventType,omitempty"`
	EventVersion *string `json:"eventVersion,omitempty"`
	ManagementEvent *bool `json:"managementEvent,omitempty"`
	ReadOnly *bool `json:"readOnly,omitempty"`
	RecipientAccountId *string `json:"recipientAccountId,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
	RequestParameters *CloudInsightsEventsListEventDetailsRequestParameters `json:"requestParameters,omitempty"`
	Resources []CloudInsightsEventsListEventDetailsResourcesInner `json:"resources,omitempty"`
	ResponseElements *CloudInsightsEventsListEventDetailsResponseElements `json:"responseElements,omitempty"`
	ServiceEventDetails *CloudInsightsEventsListEventDetailsServiceEventDetails `json:"serviceEventDetails,omitempty"`
	SessionCredentialFromConsole *string `json:"sessionCredentialFromConsole,omitempty"`
	SharedEventID *string `json:"sharedEventID,omitempty"`
	SourceIPAddress *string `json:"sourceIPAddress,omitempty"`
	TlsDetails *CloudInsightsEventsListEventDetailsTlsDetails `json:"tlsDetails,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
	UserIdentity *CloudInsightsEventsListEventDetailsUserIdentity `json:"userIdentity,omitempty"`
	VpcEndpointId *string `json:"vpcEndpointId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudInsightsEventsListEventDetails CloudInsightsEventsListEventDetails

// NewCloudInsightsEventsListEventDetails instantiates a new CloudInsightsEventsListEventDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudInsightsEventsListEventDetails() *CloudInsightsEventsListEventDetails {
	this := CloudInsightsEventsListEventDetails{}
	return &this
}

// NewCloudInsightsEventsListEventDetailsWithDefaults instantiates a new CloudInsightsEventsListEventDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudInsightsEventsListEventDetailsWithDefaults() *CloudInsightsEventsListEventDetails {
	this := CloudInsightsEventsListEventDetails{}
	return &this
}

// GetAddendum returns the Addendum field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetAddendum() CloudInsightsEventsListEventDetailsAddendum {
	if o == nil || IsNil(o.Addendum) {
		var ret CloudInsightsEventsListEventDetailsAddendum
		return ret
	}
	return *o.Addendum
}

// GetAddendumOk returns a tuple with the Addendum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetAddendumOk() (*CloudInsightsEventsListEventDetailsAddendum, bool) {
	if o == nil || IsNil(o.Addendum) {
		return nil, false
	}
	return o.Addendum, true
}

// HasAddendum returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasAddendum() bool {
	if o != nil && !IsNil(o.Addendum) {
		return true
	}

	return false
}

// SetAddendum gets a reference to the given CloudInsightsEventsListEventDetailsAddendum and assigns it to the Addendum field.
func (o *CloudInsightsEventsListEventDetails) SetAddendum(v CloudInsightsEventsListEventDetailsAddendum) {
	o.Addendum = &v
}

// GetAdditionalEventData returns the AdditionalEventData field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetAdditionalEventData() CloudInsightsEventsListEventDetailsAdditionalEventData {
	if o == nil || IsNil(o.AdditionalEventData) {
		var ret CloudInsightsEventsListEventDetailsAdditionalEventData
		return ret
	}
	return *o.AdditionalEventData
}

// GetAdditionalEventDataOk returns a tuple with the AdditionalEventData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetAdditionalEventDataOk() (*CloudInsightsEventsListEventDetailsAdditionalEventData, bool) {
	if o == nil || IsNil(o.AdditionalEventData) {
		return nil, false
	}
	return o.AdditionalEventData, true
}

// HasAdditionalEventData returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasAdditionalEventData() bool {
	if o != nil && !IsNil(o.AdditionalEventData) {
		return true
	}

	return false
}

// SetAdditionalEventData gets a reference to the given CloudInsightsEventsListEventDetailsAdditionalEventData and assigns it to the AdditionalEventData field.
func (o *CloudInsightsEventsListEventDetails) SetAdditionalEventData(v CloudInsightsEventsListEventDetailsAdditionalEventData) {
	o.AdditionalEventData = &v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetAnnotation() string {
	if o == nil || IsNil(o.Annotation) {
		var ret string
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetAnnotationOk() (*string, bool) {
	if o == nil || IsNil(o.Annotation) {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasAnnotation() bool {
	if o != nil && !IsNil(o.Annotation) {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given string and assigns it to the Annotation field.
func (o *CloudInsightsEventsListEventDetails) SetAnnotation(v string) {
	o.Annotation = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CloudInsightsEventsListEventDetails) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetAwsRegion() string {
	if o == nil || IsNil(o.AwsRegion) {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetAwsRegionOk() (*string, bool) {
	if o == nil || IsNil(o.AwsRegion) {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasAwsRegion() bool {
	if o != nil && !IsNil(o.AwsRegion) {
		return true
	}

	return false
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *CloudInsightsEventsListEventDetails) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetEdgeDeviceDetails returns the EdgeDeviceDetails field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetEdgeDeviceDetails() CloudInsightsEventsListEventDetailsEdgeDeviceDetails {
	if o == nil || IsNil(o.EdgeDeviceDetails) {
		var ret CloudInsightsEventsListEventDetailsEdgeDeviceDetails
		return ret
	}
	return *o.EdgeDeviceDetails
}

// GetEdgeDeviceDetailsOk returns a tuple with the EdgeDeviceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetEdgeDeviceDetailsOk() (*CloudInsightsEventsListEventDetailsEdgeDeviceDetails, bool) {
	if o == nil || IsNil(o.EdgeDeviceDetails) {
		return nil, false
	}
	return o.EdgeDeviceDetails, true
}

// HasEdgeDeviceDetails returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasEdgeDeviceDetails() bool {
	if o != nil && !IsNil(o.EdgeDeviceDetails) {
		return true
	}

	return false
}

// SetEdgeDeviceDetails gets a reference to the given CloudInsightsEventsListEventDetailsEdgeDeviceDetails and assigns it to the EdgeDeviceDetails field.
func (o *CloudInsightsEventsListEventDetails) SetEdgeDeviceDetails(v CloudInsightsEventsListEventDetailsEdgeDeviceDetails) {
	o.EdgeDeviceDetails = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CloudInsightsEventsListEventDetails) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CloudInsightsEventsListEventDetails) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetEventCategory returns the EventCategory field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetEventCategory() string {
	if o == nil || IsNil(o.EventCategory) {
		var ret string
		return ret
	}
	return *o.EventCategory
}

// GetEventCategoryOk returns a tuple with the EventCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetEventCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.EventCategory) {
		return nil, false
	}
	return o.EventCategory, true
}

// HasEventCategory returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasEventCategory() bool {
	if o != nil && !IsNil(o.EventCategory) {
		return true
	}

	return false
}

// SetEventCategory gets a reference to the given string and assigns it to the EventCategory field.
func (o *CloudInsightsEventsListEventDetails) SetEventCategory(v string) {
	o.EventCategory = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *CloudInsightsEventsListEventDetails) SetEventId(v string) {
	o.EventId = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *CloudInsightsEventsListEventDetails) SetEventName(v string) {
	o.EventName = &v
}

// GetEventSource returns the EventSource field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetEventSource() string {
	if o == nil || IsNil(o.EventSource) {
		var ret string
		return ret
	}
	return *o.EventSource
}

// GetEventSourceOk returns a tuple with the EventSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetEventSourceOk() (*string, bool) {
	if o == nil || IsNil(o.EventSource) {
		return nil, false
	}
	return o.EventSource, true
}

// HasEventSource returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasEventSource() bool {
	if o != nil && !IsNil(o.EventSource) {
		return true
	}

	return false
}

// SetEventSource gets a reference to the given string and assigns it to the EventSource field.
func (o *CloudInsightsEventsListEventDetails) SetEventSource(v string) {
	o.EventSource = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetEventTime() CloudInsightsEventsListEventDetailsEventTime {
	if o == nil || IsNil(o.EventTime) {
		var ret CloudInsightsEventsListEventDetailsEventTime
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetEventTimeOk() (*CloudInsightsEventsListEventDetailsEventTime, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given CloudInsightsEventsListEventDetailsEventTime and assigns it to the EventTime field.
func (o *CloudInsightsEventsListEventDetails) SetEventTime(v CloudInsightsEventsListEventDetailsEventTime) {
	o.EventTime = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CloudInsightsEventsListEventDetails) SetEventType(v string) {
	o.EventType = &v
}

// GetEventVersion returns the EventVersion field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetEventVersion() string {
	if o == nil || IsNil(o.EventVersion) {
		var ret string
		return ret
	}
	return *o.EventVersion
}

// GetEventVersionOk returns a tuple with the EventVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetEventVersionOk() (*string, bool) {
	if o == nil || IsNil(o.EventVersion) {
		return nil, false
	}
	return o.EventVersion, true
}

// HasEventVersion returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasEventVersion() bool {
	if o != nil && !IsNil(o.EventVersion) {
		return true
	}

	return false
}

// SetEventVersion gets a reference to the given string and assigns it to the EventVersion field.
func (o *CloudInsightsEventsListEventDetails) SetEventVersion(v string) {
	o.EventVersion = &v
}

// GetManagementEvent returns the ManagementEvent field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetManagementEvent() bool {
	if o == nil || IsNil(o.ManagementEvent) {
		var ret bool
		return ret
	}
	return *o.ManagementEvent
}

// GetManagementEventOk returns a tuple with the ManagementEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetManagementEventOk() (*bool, bool) {
	if o == nil || IsNil(o.ManagementEvent) {
		return nil, false
	}
	return o.ManagementEvent, true
}

// HasManagementEvent returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasManagementEvent() bool {
	if o != nil && !IsNil(o.ManagementEvent) {
		return true
	}

	return false
}

// SetManagementEvent gets a reference to the given bool and assigns it to the ManagementEvent field.
func (o *CloudInsightsEventsListEventDetails) SetManagementEvent(v bool) {
	o.ManagementEvent = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *CloudInsightsEventsListEventDetails) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRecipientAccountId returns the RecipientAccountId field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetRecipientAccountId() string {
	if o == nil || IsNil(o.RecipientAccountId) {
		var ret string
		return ret
	}
	return *o.RecipientAccountId
}

// GetRecipientAccountIdOk returns a tuple with the RecipientAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetRecipientAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientAccountId) {
		return nil, false
	}
	return o.RecipientAccountId, true
}

// HasRecipientAccountId returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasRecipientAccountId() bool {
	if o != nil && !IsNil(o.RecipientAccountId) {
		return true
	}

	return false
}

// SetRecipientAccountId gets a reference to the given string and assigns it to the RecipientAccountId field.
func (o *CloudInsightsEventsListEventDetails) SetRecipientAccountId(v string) {
	o.RecipientAccountId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CloudInsightsEventsListEventDetails) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRequestParameters returns the RequestParameters field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetRequestParameters() CloudInsightsEventsListEventDetailsRequestParameters {
	if o == nil || IsNil(o.RequestParameters) {
		var ret CloudInsightsEventsListEventDetailsRequestParameters
		return ret
	}
	return *o.RequestParameters
}

// GetRequestParametersOk returns a tuple with the RequestParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetRequestParametersOk() (*CloudInsightsEventsListEventDetailsRequestParameters, bool) {
	if o == nil || IsNil(o.RequestParameters) {
		return nil, false
	}
	return o.RequestParameters, true
}

// HasRequestParameters returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasRequestParameters() bool {
	if o != nil && !IsNil(o.RequestParameters) {
		return true
	}

	return false
}

// SetRequestParameters gets a reference to the given CloudInsightsEventsListEventDetailsRequestParameters and assigns it to the RequestParameters field.
func (o *CloudInsightsEventsListEventDetails) SetRequestParameters(v CloudInsightsEventsListEventDetailsRequestParameters) {
	o.RequestParameters = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetResources() []CloudInsightsEventsListEventDetailsResourcesInner {
	if o == nil || IsNil(o.Resources) {
		var ret []CloudInsightsEventsListEventDetailsResourcesInner
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetResourcesOk() ([]CloudInsightsEventsListEventDetailsResourcesInner, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []CloudInsightsEventsListEventDetailsResourcesInner and assigns it to the Resources field.
func (o *CloudInsightsEventsListEventDetails) SetResources(v []CloudInsightsEventsListEventDetailsResourcesInner) {
	o.Resources = v
}

// GetResponseElements returns the ResponseElements field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetResponseElements() CloudInsightsEventsListEventDetailsResponseElements {
	if o == nil || IsNil(o.ResponseElements) {
		var ret CloudInsightsEventsListEventDetailsResponseElements
		return ret
	}
	return *o.ResponseElements
}

// GetResponseElementsOk returns a tuple with the ResponseElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetResponseElementsOk() (*CloudInsightsEventsListEventDetailsResponseElements, bool) {
	if o == nil || IsNil(o.ResponseElements) {
		return nil, false
	}
	return o.ResponseElements, true
}

// HasResponseElements returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasResponseElements() bool {
	if o != nil && !IsNil(o.ResponseElements) {
		return true
	}

	return false
}

// SetResponseElements gets a reference to the given CloudInsightsEventsListEventDetailsResponseElements and assigns it to the ResponseElements field.
func (o *CloudInsightsEventsListEventDetails) SetResponseElements(v CloudInsightsEventsListEventDetailsResponseElements) {
	o.ResponseElements = &v
}

// GetServiceEventDetails returns the ServiceEventDetails field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetServiceEventDetails() CloudInsightsEventsListEventDetailsServiceEventDetails {
	if o == nil || IsNil(o.ServiceEventDetails) {
		var ret CloudInsightsEventsListEventDetailsServiceEventDetails
		return ret
	}
	return *o.ServiceEventDetails
}

// GetServiceEventDetailsOk returns a tuple with the ServiceEventDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetServiceEventDetailsOk() (*CloudInsightsEventsListEventDetailsServiceEventDetails, bool) {
	if o == nil || IsNil(o.ServiceEventDetails) {
		return nil, false
	}
	return o.ServiceEventDetails, true
}

// HasServiceEventDetails returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasServiceEventDetails() bool {
	if o != nil && !IsNil(o.ServiceEventDetails) {
		return true
	}

	return false
}

// SetServiceEventDetails gets a reference to the given CloudInsightsEventsListEventDetailsServiceEventDetails and assigns it to the ServiceEventDetails field.
func (o *CloudInsightsEventsListEventDetails) SetServiceEventDetails(v CloudInsightsEventsListEventDetailsServiceEventDetails) {
	o.ServiceEventDetails = &v
}

// GetSessionCredentialFromConsole returns the SessionCredentialFromConsole field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetSessionCredentialFromConsole() string {
	if o == nil || IsNil(o.SessionCredentialFromConsole) {
		var ret string
		return ret
	}
	return *o.SessionCredentialFromConsole
}

// GetSessionCredentialFromConsoleOk returns a tuple with the SessionCredentialFromConsole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetSessionCredentialFromConsoleOk() (*string, bool) {
	if o == nil || IsNil(o.SessionCredentialFromConsole) {
		return nil, false
	}
	return o.SessionCredentialFromConsole, true
}

// HasSessionCredentialFromConsole returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasSessionCredentialFromConsole() bool {
	if o != nil && !IsNil(o.SessionCredentialFromConsole) {
		return true
	}

	return false
}

// SetSessionCredentialFromConsole gets a reference to the given string and assigns it to the SessionCredentialFromConsole field.
func (o *CloudInsightsEventsListEventDetails) SetSessionCredentialFromConsole(v string) {
	o.SessionCredentialFromConsole = &v
}

// GetSharedEventID returns the SharedEventID field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetSharedEventID() string {
	if o == nil || IsNil(o.SharedEventID) {
		var ret string
		return ret
	}
	return *o.SharedEventID
}

// GetSharedEventIDOk returns a tuple with the SharedEventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetSharedEventIDOk() (*string, bool) {
	if o == nil || IsNil(o.SharedEventID) {
		return nil, false
	}
	return o.SharedEventID, true
}

// HasSharedEventID returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasSharedEventID() bool {
	if o != nil && !IsNil(o.SharedEventID) {
		return true
	}

	return false
}

// SetSharedEventID gets a reference to the given string and assigns it to the SharedEventID field.
func (o *CloudInsightsEventsListEventDetails) SetSharedEventID(v string) {
	o.SharedEventID = &v
}

// GetSourceIPAddress returns the SourceIPAddress field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetSourceIPAddress() string {
	if o == nil || IsNil(o.SourceIPAddress) {
		var ret string
		return ret
	}
	return *o.SourceIPAddress
}

// GetSourceIPAddressOk returns a tuple with the SourceIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetSourceIPAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SourceIPAddress) {
		return nil, false
	}
	return o.SourceIPAddress, true
}

// HasSourceIPAddress returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasSourceIPAddress() bool {
	if o != nil && !IsNil(o.SourceIPAddress) {
		return true
	}

	return false
}

// SetSourceIPAddress gets a reference to the given string and assigns it to the SourceIPAddress field.
func (o *CloudInsightsEventsListEventDetails) SetSourceIPAddress(v string) {
	o.SourceIPAddress = &v
}

// GetTlsDetails returns the TlsDetails field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetTlsDetails() CloudInsightsEventsListEventDetailsTlsDetails {
	if o == nil || IsNil(o.TlsDetails) {
		var ret CloudInsightsEventsListEventDetailsTlsDetails
		return ret
	}
	return *o.TlsDetails
}

// GetTlsDetailsOk returns a tuple with the TlsDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetTlsDetailsOk() (*CloudInsightsEventsListEventDetailsTlsDetails, bool) {
	if o == nil || IsNil(o.TlsDetails) {
		return nil, false
	}
	return o.TlsDetails, true
}

// HasTlsDetails returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasTlsDetails() bool {
	if o != nil && !IsNil(o.TlsDetails) {
		return true
	}

	return false
}

// SetTlsDetails gets a reference to the given CloudInsightsEventsListEventDetailsTlsDetails and assigns it to the TlsDetails field.
func (o *CloudInsightsEventsListEventDetails) SetTlsDetails(v CloudInsightsEventsListEventDetailsTlsDetails) {
	o.TlsDetails = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *CloudInsightsEventsListEventDetails) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetUserIdentity returns the UserIdentity field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetUserIdentity() CloudInsightsEventsListEventDetailsUserIdentity {
	if o == nil || IsNil(o.UserIdentity) {
		var ret CloudInsightsEventsListEventDetailsUserIdentity
		return ret
	}
	return *o.UserIdentity
}

// GetUserIdentityOk returns a tuple with the UserIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetUserIdentityOk() (*CloudInsightsEventsListEventDetailsUserIdentity, bool) {
	if o == nil || IsNil(o.UserIdentity) {
		return nil, false
	}
	return o.UserIdentity, true
}

// HasUserIdentity returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasUserIdentity() bool {
	if o != nil && !IsNil(o.UserIdentity) {
		return true
	}

	return false
}

// SetUserIdentity gets a reference to the given CloudInsightsEventsListEventDetailsUserIdentity and assigns it to the UserIdentity field.
func (o *CloudInsightsEventsListEventDetails) SetUserIdentity(v CloudInsightsEventsListEventDetailsUserIdentity) {
	o.UserIdentity = &v
}

// GetVpcEndpointId returns the VpcEndpointId field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetails) GetVpcEndpointId() string {
	if o == nil || IsNil(o.VpcEndpointId) {
		var ret string
		return ret
	}
	return *o.VpcEndpointId
}

// GetVpcEndpointIdOk returns a tuple with the VpcEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetails) GetVpcEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.VpcEndpointId) {
		return nil, false
	}
	return o.VpcEndpointId, true
}

// HasVpcEndpointId returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetails) HasVpcEndpointId() bool {
	if o != nil && !IsNil(o.VpcEndpointId) {
		return true
	}

	return false
}

// SetVpcEndpointId gets a reference to the given string and assigns it to the VpcEndpointId field.
func (o *CloudInsightsEventsListEventDetails) SetVpcEndpointId(v string) {
	o.VpcEndpointId = &v
}

func (o CloudInsightsEventsListEventDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudInsightsEventsListEventDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addendum) {
		toSerialize["addendum"] = o.Addendum
	}
	if !IsNil(o.AdditionalEventData) {
		toSerialize["additionalEventData"] = o.AdditionalEventData
	}
	if !IsNil(o.Annotation) {
		toSerialize["annotation"] = o.Annotation
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.AwsRegion) {
		toSerialize["awsRegion"] = o.AwsRegion
	}
	if !IsNil(o.EdgeDeviceDetails) {
		toSerialize["edgeDeviceDetails"] = o.EdgeDeviceDetails
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.EventCategory) {
		toSerialize["eventCategory"] = o.EventCategory
	}
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.EventSource) {
		toSerialize["eventSource"] = o.EventSource
	}
	if !IsNil(o.EventTime) {
		toSerialize["eventTime"] = o.EventTime
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.EventVersion) {
		toSerialize["eventVersion"] = o.EventVersion
	}
	if !IsNil(o.ManagementEvent) {
		toSerialize["managementEvent"] = o.ManagementEvent
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.RecipientAccountId) {
		toSerialize["recipientAccountId"] = o.RecipientAccountId
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.RequestParameters) {
		toSerialize["requestParameters"] = o.RequestParameters
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.ResponseElements) {
		toSerialize["responseElements"] = o.ResponseElements
	}
	if !IsNil(o.ServiceEventDetails) {
		toSerialize["serviceEventDetails"] = o.ServiceEventDetails
	}
	if !IsNil(o.SessionCredentialFromConsole) {
		toSerialize["sessionCredentialFromConsole"] = o.SessionCredentialFromConsole
	}
	if !IsNil(o.SharedEventID) {
		toSerialize["sharedEventID"] = o.SharedEventID
	}
	if !IsNil(o.SourceIPAddress) {
		toSerialize["sourceIPAddress"] = o.SourceIPAddress
	}
	if !IsNil(o.TlsDetails) {
		toSerialize["tlsDetails"] = o.TlsDetails
	}
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}
	if !IsNil(o.UserIdentity) {
		toSerialize["userIdentity"] = o.UserIdentity
	}
	if !IsNil(o.VpcEndpointId) {
		toSerialize["vpcEndpointId"] = o.VpcEndpointId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudInsightsEventsListEventDetails) UnmarshalJSON(bytes []byte) (err error) {
	varCloudInsightsEventsListEventDetails := _CloudInsightsEventsListEventDetails{}

	if err = json.Unmarshal(bytes, &varCloudInsightsEventsListEventDetails); err == nil {
		*o = CloudInsightsEventsListEventDetails(varCloudInsightsEventsListEventDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "addendum")
		delete(additionalProperties, "additionalEventData")
		delete(additionalProperties, "annotation")
		delete(additionalProperties, "apiVersion")
		delete(additionalProperties, "awsRegion")
		delete(additionalProperties, "edgeDeviceDetails")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "eventCategory")
		delete(additionalProperties, "eventId")
		delete(additionalProperties, "eventName")
		delete(additionalProperties, "eventSource")
		delete(additionalProperties, "eventTime")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "eventVersion")
		delete(additionalProperties, "managementEvent")
		delete(additionalProperties, "readOnly")
		delete(additionalProperties, "recipientAccountId")
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "requestParameters")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "responseElements")
		delete(additionalProperties, "serviceEventDetails")
		delete(additionalProperties, "sessionCredentialFromConsole")
		delete(additionalProperties, "sharedEventID")
		delete(additionalProperties, "sourceIPAddress")
		delete(additionalProperties, "tlsDetails")
		delete(additionalProperties, "userAgent")
		delete(additionalProperties, "userIdentity")
		delete(additionalProperties, "vpcEndpointId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudInsightsEventsListEventDetails struct {
	value *CloudInsightsEventsListEventDetails
	isSet bool
}

func (v NullableCloudInsightsEventsListEventDetails) Get() *CloudInsightsEventsListEventDetails {
	return v.value
}

func (v *NullableCloudInsightsEventsListEventDetails) Set(val *CloudInsightsEventsListEventDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudInsightsEventsListEventDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudInsightsEventsListEventDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudInsightsEventsListEventDetails(val *CloudInsightsEventsListEventDetails) *NullableCloudInsightsEventsListEventDetails {
	return &NullableCloudInsightsEventsListEventDetails{value: val, isSet: true}
}

func (v NullableCloudInsightsEventsListEventDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudInsightsEventsListEventDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



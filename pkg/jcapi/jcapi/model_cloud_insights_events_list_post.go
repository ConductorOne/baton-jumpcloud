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

// checks if the CloudInsightsEventsListPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudInsightsEventsListPost{}

// CloudInsightsEventsListPost struct for CloudInsightsEventsListPost
type CloudInsightsEventsListPost struct {
	AwsAccounts []string `json:"awsAccounts,omitempty"`
	EndTime *CloudInsightsEventsDistinctPostEndTime `json:"endTime,omitempty"`
	Fields []string `json:"fields,omitempty"`
	Limit *float32 `json:"limit,omitempty"`
	NpeFilter *bool `json:"npeFilter,omitempty"`
	SearchAfter map[string]interface{} `json:"searchAfter,omitempty"`
	SearchTerm *CloudInsightsEventsSearchTerm `json:"searchTerm,omitempty"`
	Service *string `json:"service,omitempty"`
	Skip *float32 `json:"skip,omitempty"`
	Sort map[string]interface{} `json:"sort,omitempty"`
	StartTime *CloudInsightsEventsDistinctPostEndTime `json:"startTime,omitempty"`
}

// NewCloudInsightsEventsListPost instantiates a new CloudInsightsEventsListPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudInsightsEventsListPost() *CloudInsightsEventsListPost {
	this := CloudInsightsEventsListPost{}
	return &this
}

// NewCloudInsightsEventsListPostWithDefaults instantiates a new CloudInsightsEventsListPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudInsightsEventsListPostWithDefaults() *CloudInsightsEventsListPost {
	this := CloudInsightsEventsListPost{}
	return &this
}

// GetAwsAccounts returns the AwsAccounts field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetAwsAccounts() []string {
	if o == nil || IsNil(o.AwsAccounts) {
		var ret []string
		return ret
	}
	return o.AwsAccounts
}

// GetAwsAccountsOk returns a tuple with the AwsAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetAwsAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwsAccounts) {
		return nil, false
	}
	return o.AwsAccounts, true
}

// HasAwsAccounts returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasAwsAccounts() bool {
	if o != nil && !IsNil(o.AwsAccounts) {
		return true
	}

	return false
}

// SetAwsAccounts gets a reference to the given []string and assigns it to the AwsAccounts field.
func (o *CloudInsightsEventsListPost) SetAwsAccounts(v []string) {
	o.AwsAccounts = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetEndTime() CloudInsightsEventsDistinctPostEndTime {
	if o == nil || IsNil(o.EndTime) {
		var ret CloudInsightsEventsDistinctPostEndTime
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetEndTimeOk() (*CloudInsightsEventsDistinctPostEndTime, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given CloudInsightsEventsDistinctPostEndTime and assigns it to the EndTime field.
func (o *CloudInsightsEventsListPost) SetEndTime(v CloudInsightsEventsDistinctPostEndTime) {
	o.EndTime = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetFields() []string {
	if o == nil || IsNil(o.Fields) {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *CloudInsightsEventsListPost) SetFields(v []string) {
	o.Fields = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetLimit() float32 {
	if o == nil || IsNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *CloudInsightsEventsListPost) SetLimit(v float32) {
	o.Limit = &v
}

// GetNpeFilter returns the NpeFilter field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetNpeFilter() bool {
	if o == nil || IsNil(o.NpeFilter) {
		var ret bool
		return ret
	}
	return *o.NpeFilter
}

// GetNpeFilterOk returns a tuple with the NpeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetNpeFilterOk() (*bool, bool) {
	if o == nil || IsNil(o.NpeFilter) {
		return nil, false
	}
	return o.NpeFilter, true
}

// HasNpeFilter returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasNpeFilter() bool {
	if o != nil && !IsNil(o.NpeFilter) {
		return true
	}

	return false
}

// SetNpeFilter gets a reference to the given bool and assigns it to the NpeFilter field.
func (o *CloudInsightsEventsListPost) SetNpeFilter(v bool) {
	o.NpeFilter = &v
}

// GetSearchAfter returns the SearchAfter field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetSearchAfter() map[string]interface{} {
	if o == nil || IsNil(o.SearchAfter) {
		var ret map[string]interface{}
		return ret
	}
	return o.SearchAfter
}

// GetSearchAfterOk returns a tuple with the SearchAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetSearchAfterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SearchAfter) {
		return map[string]interface{}{}, false
	}
	return o.SearchAfter, true
}

// HasSearchAfter returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasSearchAfter() bool {
	if o != nil && !IsNil(o.SearchAfter) {
		return true
	}

	return false
}

// SetSearchAfter gets a reference to the given map[string]interface{} and assigns it to the SearchAfter field.
func (o *CloudInsightsEventsListPost) SetSearchAfter(v map[string]interface{}) {
	o.SearchAfter = v
}

// GetSearchTerm returns the SearchTerm field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetSearchTerm() CloudInsightsEventsSearchTerm {
	if o == nil || IsNil(o.SearchTerm) {
		var ret CloudInsightsEventsSearchTerm
		return ret
	}
	return *o.SearchTerm
}

// GetSearchTermOk returns a tuple with the SearchTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetSearchTermOk() (*CloudInsightsEventsSearchTerm, bool) {
	if o == nil || IsNil(o.SearchTerm) {
		return nil, false
	}
	return o.SearchTerm, true
}

// HasSearchTerm returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasSearchTerm() bool {
	if o != nil && !IsNil(o.SearchTerm) {
		return true
	}

	return false
}

// SetSearchTerm gets a reference to the given CloudInsightsEventsSearchTerm and assigns it to the SearchTerm field.
func (o *CloudInsightsEventsListPost) SetSearchTerm(v CloudInsightsEventsSearchTerm) {
	o.SearchTerm = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *CloudInsightsEventsListPost) SetService(v string) {
	o.Service = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetSkip() float32 {
	if o == nil || IsNil(o.Skip) {
		var ret float32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetSkipOk() (*float32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given float32 and assigns it to the Skip field.
func (o *CloudInsightsEventsListPost) SetSkip(v float32) {
	o.Skip = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetSort() map[string]interface{} {
	if o == nil || IsNil(o.Sort) {
		var ret map[string]interface{}
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetSortOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Sort) {
		return map[string]interface{}{}, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given map[string]interface{} and assigns it to the Sort field.
func (o *CloudInsightsEventsListPost) SetSort(v map[string]interface{}) {
	o.Sort = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CloudInsightsEventsListPost) GetStartTime() CloudInsightsEventsDistinctPostEndTime {
	if o == nil || IsNil(o.StartTime) {
		var ret CloudInsightsEventsDistinctPostEndTime
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListPost) GetStartTimeOk() (*CloudInsightsEventsDistinctPostEndTime, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CloudInsightsEventsListPost) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given CloudInsightsEventsDistinctPostEndTime and assigns it to the StartTime field.
func (o *CloudInsightsEventsListPost) SetStartTime(v CloudInsightsEventsDistinctPostEndTime) {
	o.StartTime = &v
}

func (o CloudInsightsEventsListPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudInsightsEventsListPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsAccounts) {
		toSerialize["awsAccounts"] = o.AwsAccounts
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.NpeFilter) {
		toSerialize["npeFilter"] = o.NpeFilter
	}
	if !IsNil(o.SearchAfter) {
		toSerialize["searchAfter"] = o.SearchAfter
	}
	if !IsNil(o.SearchTerm) {
		toSerialize["searchTerm"] = o.SearchTerm
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableCloudInsightsEventsListPost struct {
	value *CloudInsightsEventsListPost
	isSet bool
}

func (v NullableCloudInsightsEventsListPost) Get() *CloudInsightsEventsListPost {
	return v.value
}

func (v *NullableCloudInsightsEventsListPost) Set(val *CloudInsightsEventsListPost) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudInsightsEventsListPost) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudInsightsEventsListPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudInsightsEventsListPost(val *CloudInsightsEventsListPost) *NullableCloudInsightsEventsListPost {
	return &NullableCloudInsightsEventsListPost{value: val, isSet: true}
}

func (v NullableCloudInsightsEventsListPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudInsightsEventsListPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



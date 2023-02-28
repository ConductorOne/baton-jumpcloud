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

// checks if the CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer{}

// CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer struct for CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer
type CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer struct {
	AccountId *string `json:"accountId,omitempty"`
	Arn *string `json:"arn,omitempty"`
	PrincipleId *string `json:"principleId,omitempty"`
	Type *string `json:"type,omitempty"`
	UserName *string `json:"userName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer

// NewCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer instantiates a new CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer() *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer {
	this := CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer{}
	return &this
}

// NewCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuerWithDefaults instantiates a new CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuerWithDefaults() *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer {
	this := CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) SetAccountId(v string) {
	o.AccountId = &v
}

// GetArn returns the Arn field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetArn() string {
	if o == nil || IsNil(o.Arn) {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetArnOk() (*string, bool) {
	if o == nil || IsNil(o.Arn) {
		return nil, false
	}
	return o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) HasArn() bool {
	if o != nil && !IsNil(o.Arn) {
		return true
	}

	return false
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) SetArn(v string) {
	o.Arn = &v
}

// GetPrincipleId returns the PrincipleId field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetPrincipleId() string {
	if o == nil || IsNil(o.PrincipleId) {
		var ret string
		return ret
	}
	return *o.PrincipleId
}

// GetPrincipleIdOk returns a tuple with the PrincipleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetPrincipleIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrincipleId) {
		return nil, false
	}
	return o.PrincipleId, true
}

// HasPrincipleId returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) HasPrincipleId() bool {
	if o != nil && !IsNil(o.PrincipleId) {
		return true
	}

	return false
}

// SetPrincipleId gets a reference to the given string and assigns it to the PrincipleId field.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) SetPrincipleId(v string) {
	o.PrincipleId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) SetType(v string) {
	o.Type = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) SetUserName(v string) {
	o.UserName = &v
}

func (o CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Arn) {
		toSerialize["arn"] = o.Arn
	}
	if !IsNil(o.PrincipleId) {
		toSerialize["principleId"] = o.PrincipleId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) UnmarshalJSON(bytes []byte) (err error) {
	varCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer := _CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer{}

	if err = json.Unmarshal(bytes, &varCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer); err == nil {
		*o = CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer(varCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "arn")
		delete(additionalProperties, "principleId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer struct {
	value *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer
	isSet bool
}

func (v NullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) Get() *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer {
	return v.value
}

func (v *NullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) Set(val *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer(val *CloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) *NullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer {
	return &NullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer{value: val, isSet: true}
}

func (v NullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudInsightsEventsListEventDetailsUserIdentitySessionContextSessionIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



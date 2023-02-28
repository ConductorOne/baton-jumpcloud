/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jcapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SystemsApiService SystemsApi service
type SystemsApiService service

type SystemsApiGraphSystemAssociationsListRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	targets *[]string
	limit *int32
	skip *int32
	date *string
	authorization *string
	xOrgId *string
}

// Targets which a \&quot;system\&quot; can be associated to.
func (r SystemsApiGraphSystemAssociationsListRequest) Targets(targets []string) SystemsApiGraphSystemAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r SystemsApiGraphSystemAssociationsListRequest) Limit(limit int32) SystemsApiGraphSystemAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r SystemsApiGraphSystemAssociationsListRequest) Skip(skip int32) SystemsApiGraphSystemAssociationsListRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r SystemsApiGraphSystemAssociationsListRequest) Date(date string) SystemsApiGraphSystemAssociationsListRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r SystemsApiGraphSystemAssociationsListRequest) Authorization(authorization string) SystemsApiGraphSystemAssociationsListRequest {
	r.authorization = &authorization
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiGraphSystemAssociationsListRequest) XOrgId(xOrgId string) SystemsApiGraphSystemAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemsApiGraphSystemAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphSystemAssociationsListExecute(r)
}

/*
GraphSystemAssociationsList List the associations of a System

This endpoint returns the _direct_ associations of a System.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Systems and Users.


#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/associations?targets=user \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return SystemsApiGraphSystemAssociationsListRequest
*/
func (a *SystemsApiService) GraphSystemAssociationsList(ctx context.Context, systemId string) SystemsApiGraphSystemAssociationsListRequest {
	return SystemsApiGraphSystemAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *SystemsApiService) GraphSystemAssociationsListExecute(r SystemsApiGraphSystemAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.GraphSystemAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.targets == nil {
		return localVarReturnValue, nil, reportError("targets is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "targets", r.targets, "csv")
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Date", r.date, "")
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SystemsApiGraphSystemAssociationsPostRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	date *string
	authorization *string
	xOrgId *string
	body *GraphOperationSystem
}

// Current date header for the System Context API
func (r SystemsApiGraphSystemAssociationsPostRequest) Date(date string) SystemsApiGraphSystemAssociationsPostRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r SystemsApiGraphSystemAssociationsPostRequest) Authorization(authorization string) SystemsApiGraphSystemAssociationsPostRequest {
	r.authorization = &authorization
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiGraphSystemAssociationsPostRequest) XOrgId(xOrgId string) SystemsApiGraphSystemAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemsApiGraphSystemAssociationsPostRequest) Body(body GraphOperationSystem) SystemsApiGraphSystemAssociationsPostRequest {
	r.body = &body
	return r
}

func (r SystemsApiGraphSystemAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphSystemAssociationsPostExecute(r)
}

/*
GraphSystemAssociationsPost Manage associations of a System

This endpoint allows you to manage the _direct_ associations of a System.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Systems and Users.


#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/systems/{System_ID}/associations \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "attributes": {
      "sudo": {
        "enabled": true,
        "withoutPassword": false
      }
    },
    "op": "add",
    "type": "user",
    "id": "UserID"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return SystemsApiGraphSystemAssociationsPostRequest
*/
func (a *SystemsApiService) GraphSystemAssociationsPost(ctx context.Context, systemId string) SystemsApiGraphSystemAssociationsPostRequest {
	return SystemsApiGraphSystemAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
func (a *SystemsApiService) GraphSystemAssociationsPostExecute(r SystemsApiGraphSystemAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.GraphSystemAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Date", r.date, "")
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SystemsApiGraphSystemMemberOfRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	filter *[]string
	limit *int32
	skip *int32
	date *string
	authorization *string
	sort *[]string
	xOrgId *string
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SystemsApiGraphSystemMemberOfRequest) Filter(filter []string) SystemsApiGraphSystemMemberOfRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r SystemsApiGraphSystemMemberOfRequest) Limit(limit int32) SystemsApiGraphSystemMemberOfRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r SystemsApiGraphSystemMemberOfRequest) Skip(skip int32) SystemsApiGraphSystemMemberOfRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r SystemsApiGraphSystemMemberOfRequest) Date(date string) SystemsApiGraphSystemMemberOfRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r SystemsApiGraphSystemMemberOfRequest) Authorization(authorization string) SystemsApiGraphSystemMemberOfRequest {
	r.authorization = &authorization
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r SystemsApiGraphSystemMemberOfRequest) Sort(sort []string) SystemsApiGraphSystemMemberOfRequest {
	r.sort = &sort
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiGraphSystemMemberOfRequest) XOrgId(xOrgId string) SystemsApiGraphSystemMemberOfRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemsApiGraphSystemMemberOfRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemMemberOfExecute(r)
}

/*
GraphSystemMemberOf List the parent Groups of a System

This endpoint returns all the System Groups a System is a member of.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/memberof \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return SystemsApiGraphSystemMemberOfRequest
*/
func (a *SystemsApiService) GraphSystemMemberOf(ctx context.Context, systemId string) SystemsApiGraphSystemMemberOfRequest {
	return SystemsApiGraphSystemMemberOfRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *SystemsApiService) GraphSystemMemberOfExecute(r SystemsApiGraphSystemMemberOfRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.GraphSystemMemberOf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/memberof"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Date", r.date, "")
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SystemsApiGraphSystemTraverseCommandRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r SystemsApiGraphSystemTraverseCommandRequest) Limit(limit int32) SystemsApiGraphSystemTraverseCommandRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiGraphSystemTraverseCommandRequest) XOrgId(xOrgId string) SystemsApiGraphSystemTraverseCommandRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemsApiGraphSystemTraverseCommandRequest) Skip(skip int32) SystemsApiGraphSystemTraverseCommandRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SystemsApiGraphSystemTraverseCommandRequest) Filter(filter []string) SystemsApiGraphSystemTraverseCommandRequest {
	r.filter = &filter
	return r
}

func (r SystemsApiGraphSystemTraverseCommandRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemTraverseCommandExecute(r)
}

/*
GraphSystemTraverseCommand List the Commands bound to a System

This endpoint will return all Commands bound to a System, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System to the corresponding Command; this array represents all grouping and/or associations that would have to be removed to deprovision the Command from this System.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/commands \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return SystemsApiGraphSystemTraverseCommandRequest
*/
func (a *SystemsApiService) GraphSystemTraverseCommand(ctx context.Context, systemId string) SystemsApiGraphSystemTraverseCommandRequest {
	return SystemsApiGraphSystemTraverseCommandRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *SystemsApiService) GraphSystemTraverseCommandExecute(r SystemsApiGraphSystemTraverseCommandRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.GraphSystemTraverseCommand")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/commands"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SystemsApiGraphSystemTraversePolicyRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r SystemsApiGraphSystemTraversePolicyRequest) Limit(limit int32) SystemsApiGraphSystemTraversePolicyRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiGraphSystemTraversePolicyRequest) XOrgId(xOrgId string) SystemsApiGraphSystemTraversePolicyRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemsApiGraphSystemTraversePolicyRequest) Skip(skip int32) SystemsApiGraphSystemTraversePolicyRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SystemsApiGraphSystemTraversePolicyRequest) Filter(filter []string) SystemsApiGraphSystemTraversePolicyRequest {
	r.filter = &filter
	return r
}

func (r SystemsApiGraphSystemTraversePolicyRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemTraversePolicyExecute(r)
}

/*
GraphSystemTraversePolicy List the Policies bound to a System

This endpoint will return all Policies bound to a System, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System to the corresponding Policy; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy from this System.

See `/members` and `/associations` endpoints to manage those collections.

This endpoint is not yet public as we have finish the code.

##### Sample Request

```
curl -X GET https://console.jumpcloud.com/api/v2/{System_ID}/policies \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return SystemsApiGraphSystemTraversePolicyRequest
*/
func (a *SystemsApiService) GraphSystemTraversePolicy(ctx context.Context, systemId string) SystemsApiGraphSystemTraversePolicyRequest {
	return SystemsApiGraphSystemTraversePolicyRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *SystemsApiService) GraphSystemTraversePolicyExecute(r SystemsApiGraphSystemTraversePolicyRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.GraphSystemTraversePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/policies"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SystemsApiGraphSystemTraversePolicyGroupRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	date *string
	authorization *string
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r SystemsApiGraphSystemTraversePolicyGroupRequest) Limit(limit int32) SystemsApiGraphSystemTraversePolicyGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiGraphSystemTraversePolicyGroupRequest) XOrgId(xOrgId string) SystemsApiGraphSystemTraversePolicyGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemsApiGraphSystemTraversePolicyGroupRequest) Skip(skip int32) SystemsApiGraphSystemTraversePolicyGroupRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r SystemsApiGraphSystemTraversePolicyGroupRequest) Date(date string) SystemsApiGraphSystemTraversePolicyGroupRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r SystemsApiGraphSystemTraversePolicyGroupRequest) Authorization(authorization string) SystemsApiGraphSystemTraversePolicyGroupRequest {
	r.authorization = &authorization
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SystemsApiGraphSystemTraversePolicyGroupRequest) Filter(filter []string) SystemsApiGraphSystemTraversePolicyGroupRequest {
	r.filter = &filter
	return r
}

func (r SystemsApiGraphSystemTraversePolicyGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemTraversePolicyGroupExecute(r)
}

/*
GraphSystemTraversePolicyGroup List the Policy Groups bound to a System

This endpoint will return all Policy Groups bound to a System, either directly or indirectly essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System to the corresponding Policy Group; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy Group from this System.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/policygroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return SystemsApiGraphSystemTraversePolicyGroupRequest
*/
func (a *SystemsApiService) GraphSystemTraversePolicyGroup(ctx context.Context, systemId string) SystemsApiGraphSystemTraversePolicyGroupRequest {
	return SystemsApiGraphSystemTraversePolicyGroupRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *SystemsApiService) GraphSystemTraversePolicyGroupExecute(r SystemsApiGraphSystemTraversePolicyGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.GraphSystemTraversePolicyGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/policygroups"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Date", r.date, "")
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SystemsApiGraphSystemTraverseUserRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	date *string
	authorization *string
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r SystemsApiGraphSystemTraverseUserRequest) Limit(limit int32) SystemsApiGraphSystemTraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiGraphSystemTraverseUserRequest) XOrgId(xOrgId string) SystemsApiGraphSystemTraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemsApiGraphSystemTraverseUserRequest) Skip(skip int32) SystemsApiGraphSystemTraverseUserRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r SystemsApiGraphSystemTraverseUserRequest) Date(date string) SystemsApiGraphSystemTraverseUserRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r SystemsApiGraphSystemTraverseUserRequest) Authorization(authorization string) SystemsApiGraphSystemTraverseUserRequest {
	r.authorization = &authorization
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SystemsApiGraphSystemTraverseUserRequest) Filter(filter []string) SystemsApiGraphSystemTraverseUserRequest {
	r.filter = &filter
	return r
}

func (r SystemsApiGraphSystemTraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemTraverseUserExecute(r)
}

/*
GraphSystemTraverseUser List the Users bound to a System

This endpoint will return all Users bound to a System, either directly or indirectly essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this System.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/users \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return SystemsApiGraphSystemTraverseUserRequest
*/
func (a *SystemsApiService) GraphSystemTraverseUser(ctx context.Context, systemId string) SystemsApiGraphSystemTraverseUserRequest {
	return SystemsApiGraphSystemTraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *SystemsApiService) GraphSystemTraverseUserExecute(r SystemsApiGraphSystemTraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.GraphSystemTraverseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Date", r.date, "")
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SystemsApiGraphSystemTraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	date *string
	authorization *string
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r SystemsApiGraphSystemTraverseUserGroupRequest) Limit(limit int32) SystemsApiGraphSystemTraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiGraphSystemTraverseUserGroupRequest) XOrgId(xOrgId string) SystemsApiGraphSystemTraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemsApiGraphSystemTraverseUserGroupRequest) Skip(skip int32) SystemsApiGraphSystemTraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r SystemsApiGraphSystemTraverseUserGroupRequest) Date(date string) SystemsApiGraphSystemTraverseUserGroupRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r SystemsApiGraphSystemTraverseUserGroupRequest) Authorization(authorization string) SystemsApiGraphSystemTraverseUserGroupRequest {
	r.authorization = &authorization
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SystemsApiGraphSystemTraverseUserGroupRequest) Filter(filter []string) SystemsApiGraphSystemTraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r SystemsApiGraphSystemTraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemTraverseUserGroupExecute(r)
}

/*
GraphSystemTraverseUserGroup List the User Groups bound to a System

This endpoint will return all User Groups bound to a System, either directly or indirectly essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this System.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/usergroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return SystemsApiGraphSystemTraverseUserGroupRequest
*/
func (a *SystemsApiService) GraphSystemTraverseUserGroup(ctx context.Context, systemId string) SystemsApiGraphSystemTraverseUserGroupRequest {
	return SystemsApiGraphSystemTraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *SystemsApiService) GraphSystemTraverseUserGroupExecute(r SystemsApiGraphSystemTraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.GraphSystemTraverseUserGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Date", r.date, "")
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SystemsApiSystemsGetFDEKeyRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	xOrgId *string
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiSystemsGetFDEKeyRequest) XOrgId(xOrgId string) SystemsApiSystemsGetFDEKeyRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemsApiSystemsGetFDEKeyRequest) Execute() (*Systemfdekey, *http.Response, error) {
	return r.ApiService.SystemsGetFDEKeyExecute(r)
}

/*
SystemsGetFDEKey Get System FDE Key

This endpoint will return the current (latest) fde key saved for a system.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId
 @return SystemsApiSystemsGetFDEKeyRequest
*/
func (a *SystemsApiService) SystemsGetFDEKey(ctx context.Context, systemId string) SystemsApiSystemsGetFDEKeyRequest {
	return SystemsApiSystemsGetFDEKeyRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return Systemfdekey
func (a *SystemsApiService) SystemsGetFDEKeyExecute(r SystemsApiSystemsGetFDEKeyRequest) (*Systemfdekey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Systemfdekey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.SystemsGetFDEKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/fdekey"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SystemsApiSystemsListSoftwareAppsWithStatusesRequest struct {
	ctx context.Context
	ApiService *SystemsApiService
	systemId string
	xOrgId *string
	filter *[]string
	limit *int32
	skip *int32
	sort *[]string
}

// Organization identifier that can be obtained from console settings.
func (r SystemsApiSystemsListSoftwareAppsWithStatusesRequest) XOrgId(xOrgId string) SystemsApiSystemsListSoftwareAppsWithStatusesRequest {
	r.xOrgId = &xOrgId
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SystemsApiSystemsListSoftwareAppsWithStatusesRequest) Filter(filter []string) SystemsApiSystemsListSoftwareAppsWithStatusesRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r SystemsApiSystemsListSoftwareAppsWithStatusesRequest) Limit(limit int32) SystemsApiSystemsListSoftwareAppsWithStatusesRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r SystemsApiSystemsListSoftwareAppsWithStatusesRequest) Skip(skip int32) SystemsApiSystemsListSoftwareAppsWithStatusesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r SystemsApiSystemsListSoftwareAppsWithStatusesRequest) Sort(sort []string) SystemsApiSystemsListSoftwareAppsWithStatusesRequest {
	r.sort = &sort
	return r
}

func (r SystemsApiSystemsListSoftwareAppsWithStatusesRequest) Execute() ([]SoftwareAppWithStatus, *http.Response, error) {
	return r.ApiService.SystemsListSoftwareAppsWithStatusesExecute(r)
}

/*
SystemsListSoftwareAppsWithStatuses List the associated Software Application Statuses of a System

This endpoint returns all the statuses of the associated Software Applications from the provided JumpCloud system ID.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systems/{system_id}/softwareappstatuses \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return SystemsApiSystemsListSoftwareAppsWithStatusesRequest
*/
func (a *SystemsApiService) SystemsListSoftwareAppsWithStatuses(ctx context.Context, systemId string) SystemsApiSystemsListSoftwareAppsWithStatusesRequest {
	return SystemsApiSystemsListSoftwareAppsWithStatusesRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []SoftwareAppWithStatus
func (a *SystemsApiService) SystemsListSoftwareAppsWithStatusesExecute(r SystemsApiSystemsListSoftwareAppsWithStatusesRequest) ([]SoftwareAppWithStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SoftwareAppWithStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemsApiService.SystemsListSoftwareAppsWithStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/softwareappstatuses"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["x-api-key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

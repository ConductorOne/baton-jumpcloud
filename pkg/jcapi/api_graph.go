/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github.com/ConductorOne/baton-jumpcloud/pkg/jcapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// GraphApiService GraphApi service
type GraphApiService service

type GraphApiGraphActiveDirectoryAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	activedirectoryId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;active_directory\&quot; can be associated to.
func (r GraphApiGraphActiveDirectoryAssociationsListRequest) Targets(targets []string) GraphApiGraphActiveDirectoryAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphActiveDirectoryAssociationsListRequest) Limit(limit int32) GraphApiGraphActiveDirectoryAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphActiveDirectoryAssociationsListRequest) Skip(skip int32) GraphApiGraphActiveDirectoryAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphActiveDirectoryAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphActiveDirectoryAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphActiveDirectoryAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphActiveDirectoryAssociationsListExecute(r)
}

/*
GraphActiveDirectoryAssociationsList List the associations of an Active Directory instance

This endpoint returns the direct associations of this Active Directory instance.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Active Directory and Users.


#### Sample Request
```
curl -X GET 'https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID}/associations?targets=user \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param activedirectoryId
 @return GraphApiGraphActiveDirectoryAssociationsListRequest
*/
func (a *GraphApiService) GraphActiveDirectoryAssociationsList(ctx context.Context, activedirectoryId string) GraphApiGraphActiveDirectoryAssociationsListRequest {
	return GraphApiGraphActiveDirectoryAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		activedirectoryId: activedirectoryId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphActiveDirectoryAssociationsListExecute(r GraphApiGraphActiveDirectoryAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphActiveDirectoryAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectories/{activedirectory_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"activedirectory_id"+"}", url.PathEscape(parameterValueToString(r.activedirectoryId, "activedirectoryId")), -1)

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

type GraphApiGraphActiveDirectoryAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	activedirectoryId string
	xOrgId *string
	body *GraphOperationActiveDirectory
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphActiveDirectoryAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphActiveDirectoryAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphActiveDirectoryAssociationsPostRequest) Body(body GraphOperationActiveDirectory) GraphApiGraphActiveDirectoryAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphActiveDirectoryAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphActiveDirectoryAssociationsPostExecute(r)
}

/*
GraphActiveDirectoryAssociationsPost Manage the associations of an Active Directory instance

This endpoint allows you to manage the _direct_ associations of an Active Directory instance.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Active Directory and Users.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/activedirectories/{AD_Instance_ID}/associations \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "user",
    "id": "{User_ID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param activedirectoryId
 @return GraphApiGraphActiveDirectoryAssociationsPostRequest
*/
func (a *GraphApiService) GraphActiveDirectoryAssociationsPost(ctx context.Context, activedirectoryId string) GraphApiGraphActiveDirectoryAssociationsPostRequest {
	return GraphApiGraphActiveDirectoryAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		activedirectoryId: activedirectoryId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphActiveDirectoryAssociationsPostExecute(r GraphApiGraphActiveDirectoryAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphActiveDirectoryAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectories/{activedirectory_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"activedirectory_id"+"}", url.PathEscape(parameterValueToString(r.activedirectoryId, "activedirectoryId")), -1)

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

type GraphApiGraphActiveDirectoryTraverseUserRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	activedirectoryId string
	filter *[]string
	limit *int32
	xOrgId *string
	skip *int32
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphActiveDirectoryTraverseUserRequest) Filter(filter []string) GraphApiGraphActiveDirectoryTraverseUserRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphActiveDirectoryTraverseUserRequest) Limit(limit int32) GraphApiGraphActiveDirectoryTraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphActiveDirectoryTraverseUserRequest) XOrgId(xOrgId string) GraphApiGraphActiveDirectoryTraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphActiveDirectoryTraverseUserRequest) Skip(skip int32) GraphApiGraphActiveDirectoryTraverseUserRequest {
	r.skip = &skip
	return r
}

func (r GraphApiGraphActiveDirectoryTraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphActiveDirectoryTraverseUserExecute(r)
}

/*
GraphActiveDirectoryTraverseUser List the Users bound to an Active Directory instance

This endpoint will return all Users bound to an Active Directory instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Active Directory instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Active Directory instance.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID}/users \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param activedirectoryId ObjectID of the Active Directory instance.
 @return GraphApiGraphActiveDirectoryTraverseUserRequest
*/
func (a *GraphApiService) GraphActiveDirectoryTraverseUser(ctx context.Context, activedirectoryId string) GraphApiGraphActiveDirectoryTraverseUserRequest {
	return GraphApiGraphActiveDirectoryTraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		activedirectoryId: activedirectoryId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphActiveDirectoryTraverseUserExecute(r GraphApiGraphActiveDirectoryTraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphActiveDirectoryTraverseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectories/{activedirectory_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"activedirectory_id"+"}", url.PathEscape(parameterValueToString(r.activedirectoryId, "activedirectoryId")), -1)

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

type GraphApiGraphActiveDirectoryTraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	activedirectoryId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphActiveDirectoryTraverseUserGroupRequest) Limit(limit int32) GraphApiGraphActiveDirectoryTraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphActiveDirectoryTraverseUserGroupRequest) XOrgId(xOrgId string) GraphApiGraphActiveDirectoryTraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphActiveDirectoryTraverseUserGroupRequest) Skip(skip int32) GraphApiGraphActiveDirectoryTraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphActiveDirectoryTraverseUserGroupRequest) Filter(filter []string) GraphApiGraphActiveDirectoryTraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphActiveDirectoryTraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphActiveDirectoryTraverseUserGroupExecute(r)
}

/*
GraphActiveDirectoryTraverseUserGroup List the User Groups bound to an Active Directory instance

This endpoint will return all Users Groups bound to an Active Directory instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Active Directory instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Active Directory instance.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/activedirectories/{ActiveDirectory_ID}/usergroups \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param activedirectoryId ObjectID of the Active Directory instance.
 @return GraphApiGraphActiveDirectoryTraverseUserGroupRequest
*/
func (a *GraphApiService) GraphActiveDirectoryTraverseUserGroup(ctx context.Context, activedirectoryId string) GraphApiGraphActiveDirectoryTraverseUserGroupRequest {
	return GraphApiGraphActiveDirectoryTraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		activedirectoryId: activedirectoryId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphActiveDirectoryTraverseUserGroupExecute(r GraphApiGraphActiveDirectoryTraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphActiveDirectoryTraverseUserGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/activedirectories/{activedirectory_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"activedirectory_id"+"}", url.PathEscape(parameterValueToString(r.activedirectoryId, "activedirectoryId")), -1)

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

type GraphApiGraphApplicationAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	applicationId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;application\&quot; can be associated to.
func (r GraphApiGraphApplicationAssociationsListRequest) Targets(targets []string) GraphApiGraphApplicationAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphApplicationAssociationsListRequest) Limit(limit int32) GraphApiGraphApplicationAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphApplicationAssociationsListRequest) Skip(skip int32) GraphApiGraphApplicationAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphApplicationAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphApplicationAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphApplicationAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphApplicationAssociationsListExecute(r)
}

/*
GraphApplicationAssociationsList List the associations of an Application

This endpoint will return the _direct_ associations of an Application. A direct association can be a non-homogeneous relationship between 2 different objects, for example Applications and User Groups.


#### Sample Request
```
curl -X GET 'https://console.jumpcloud.com/api/v2/applications/{Application_ID}/associations?targets=user_group \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId ObjectID of the Application.
 @return GraphApiGraphApplicationAssociationsListRequest
*/
func (a *GraphApiService) GraphApplicationAssociationsList(ctx context.Context, applicationId string) GraphApiGraphApplicationAssociationsListRequest {
	return GraphApiGraphApplicationAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphApplicationAssociationsListExecute(r GraphApiGraphApplicationAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphApplicationAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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

type GraphApiGraphApplicationAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	applicationId string
	xOrgId *string
	body *GraphOperationApplication
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphApplicationAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphApplicationAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphApplicationAssociationsPostRequest) Body(body GraphOperationApplication) GraphApiGraphApplicationAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphApplicationAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphApplicationAssociationsPostExecute(r)
}

/*
GraphApplicationAssociationsPost Manage the associations of an Application

This endpoint allows you to manage the _direct_ associations of an Application. A direct association can be a non-homogeneous relationship between 2 different objects, for example Application and User Groups.

#### Sample Request
```
curl -X POST 'https://console.jumpcloud.com/api/v2/applications/{Application_ID}/associations' \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "user_group",
    "id": "{Group_ID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId ObjectID of the Application.
 @return GraphApiGraphApplicationAssociationsPostRequest
*/
func (a *GraphApiService) GraphApplicationAssociationsPost(ctx context.Context, applicationId string) GraphApiGraphApplicationAssociationsPostRequest {
	return GraphApiGraphApplicationAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphApplicationAssociationsPostExecute(r GraphApiGraphApplicationAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphApplicationAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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

type GraphApiGraphApplicationTraverseUserRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	applicationId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphApplicationTraverseUserRequest) Limit(limit int32) GraphApiGraphApplicationTraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphApplicationTraverseUserRequest) XOrgId(xOrgId string) GraphApiGraphApplicationTraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphApplicationTraverseUserRequest) Skip(skip int32) GraphApiGraphApplicationTraverseUserRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphApplicationTraverseUserRequest) Filter(filter []string) GraphApiGraphApplicationTraverseUserRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphApplicationTraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphApplicationTraverseUserExecute(r)
}

/*
GraphApplicationTraverseUser List the Users bound to an Application

This endpoint will return all Users bound to an Application, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Application to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Application.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/applications/{Application_ID}/users \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId ObjectID of the Application.
 @return GraphApiGraphApplicationTraverseUserRequest
*/
func (a *GraphApiService) GraphApplicationTraverseUser(ctx context.Context, applicationId string) GraphApiGraphApplicationTraverseUserRequest {
	return GraphApiGraphApplicationTraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphApplicationTraverseUserExecute(r GraphApiGraphApplicationTraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphApplicationTraverseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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

type GraphApiGraphApplicationTraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	applicationId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphApplicationTraverseUserGroupRequest) Limit(limit int32) GraphApiGraphApplicationTraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphApplicationTraverseUserGroupRequest) XOrgId(xOrgId string) GraphApiGraphApplicationTraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphApplicationTraverseUserGroupRequest) Skip(skip int32) GraphApiGraphApplicationTraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphApplicationTraverseUserGroupRequest) Filter(filter []string) GraphApiGraphApplicationTraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphApplicationTraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphApplicationTraverseUserGroupExecute(r)
}

/*
GraphApplicationTraverseUserGroup List the User Groups bound to an Application

This endpoint will return all Users Groups bound to an Application, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates  each path from this Application to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Application.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/applications/{Application_ID}/usergroups \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId ObjectID of the Application.
 @return GraphApiGraphApplicationTraverseUserGroupRequest
*/
func (a *GraphApiService) GraphApplicationTraverseUserGroup(ctx context.Context, applicationId string) GraphApiGraphApplicationTraverseUserGroupRequest {
	return GraphApiGraphApplicationTraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphApplicationTraverseUserGroupExecute(r GraphApiGraphApplicationTraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphApplicationTraverseUserGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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

type GraphApiGraphCommandAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	commandId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;command\&quot; can be associated to.
func (r GraphApiGraphCommandAssociationsListRequest) Targets(targets []string) GraphApiGraphCommandAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphCommandAssociationsListRequest) Limit(limit int32) GraphApiGraphCommandAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphCommandAssociationsListRequest) Skip(skip int32) GraphApiGraphCommandAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphCommandAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphCommandAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphCommandAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphCommandAssociationsListExecute(r)
}

/*
GraphCommandAssociationsList List the associations of a Command

This endpoint will return the _direct_ associations of this Command.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Commands and User Groups.


#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/commands/{Command_ID}/associations?targets=system_group \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commandId ObjectID of the Command.
 @return GraphApiGraphCommandAssociationsListRequest
*/
func (a *GraphApiService) GraphCommandAssociationsList(ctx context.Context, commandId string) GraphApiGraphCommandAssociationsListRequest {
	return GraphApiGraphCommandAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		commandId: commandId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphCommandAssociationsListExecute(r GraphApiGraphCommandAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphCommandAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commands/{command_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"command_id"+"}", url.PathEscape(parameterValueToString(r.commandId, "commandId")), -1)

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

type GraphApiGraphCommandAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	commandId string
	xOrgId *string
	body *GraphOperationCommand
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphCommandAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphCommandAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphCommandAssociationsPostRequest) Body(body GraphOperationCommand) GraphApiGraphCommandAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphCommandAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphCommandAssociationsPostExecute(r)
}

/*
GraphCommandAssociationsPost Manage the associations of a Command

This endpoint will allow you to manage the _direct_ associations of this Command.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Commands and User Groups.


#### Sample Request
```
 curl -X POST https://console.jumpcloud.com/api/v2/commands/{Command_ID}/associations \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "system_group",
    "id": "Group_ID"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commandId ObjectID of the Command.
 @return GraphApiGraphCommandAssociationsPostRequest
*/
func (a *GraphApiService) GraphCommandAssociationsPost(ctx context.Context, commandId string) GraphApiGraphCommandAssociationsPostRequest {
	return GraphApiGraphCommandAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		commandId: commandId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphCommandAssociationsPostExecute(r GraphApiGraphCommandAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphCommandAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commands/{command_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"command_id"+"}", url.PathEscape(parameterValueToString(r.commandId, "commandId")), -1)

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

type GraphApiGraphCommandTraverseSystemRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	commandId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphCommandTraverseSystemRequest) Limit(limit int32) GraphApiGraphCommandTraverseSystemRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphCommandTraverseSystemRequest) XOrgId(xOrgId string) GraphApiGraphCommandTraverseSystemRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphCommandTraverseSystemRequest) Skip(skip int32) GraphApiGraphCommandTraverseSystemRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphCommandTraverseSystemRequest) Filter(filter []string) GraphApiGraphCommandTraverseSystemRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphCommandTraverseSystemRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphCommandTraverseSystemExecute(r)
}

/*
GraphCommandTraverseSystem List the Systems bound to a Command

This endpoint will return all Systems bound to a Command, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Command to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Command.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/commands/{Command_ID}/systems \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commandId ObjectID of the Command.
 @return GraphApiGraphCommandTraverseSystemRequest
*/
func (a *GraphApiService) GraphCommandTraverseSystem(ctx context.Context, commandId string) GraphApiGraphCommandTraverseSystemRequest {
	return GraphApiGraphCommandTraverseSystemRequest{
		ApiService: a,
		ctx: ctx,
		commandId: commandId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphCommandTraverseSystemExecute(r GraphApiGraphCommandTraverseSystemRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphCommandTraverseSystem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commands/{command_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"command_id"+"}", url.PathEscape(parameterValueToString(r.commandId, "commandId")), -1)

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

type GraphApiGraphCommandTraverseSystemGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	commandId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphCommandTraverseSystemGroupRequest) Limit(limit int32) GraphApiGraphCommandTraverseSystemGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphCommandTraverseSystemGroupRequest) XOrgId(xOrgId string) GraphApiGraphCommandTraverseSystemGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphCommandTraverseSystemGroupRequest) Skip(skip int32) GraphApiGraphCommandTraverseSystemGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphCommandTraverseSystemGroupRequest) Filter(filter []string) GraphApiGraphCommandTraverseSystemGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphCommandTraverseSystemGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphCommandTraverseSystemGroupExecute(r)
}

/*
GraphCommandTraverseSystemGroup List the System Groups bound to a Command

This endpoint will return all System Groups bound to a Command, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Command to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Command.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/commands/{Command_ID}/systemgroups \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commandId ObjectID of the Command.
 @return GraphApiGraphCommandTraverseSystemGroupRequest
*/
func (a *GraphApiService) GraphCommandTraverseSystemGroup(ctx context.Context, commandId string) GraphApiGraphCommandTraverseSystemGroupRequest {
	return GraphApiGraphCommandTraverseSystemGroupRequest{
		ApiService: a,
		ctx: ctx,
		commandId: commandId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphCommandTraverseSystemGroupExecute(r GraphApiGraphCommandTraverseSystemGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphCommandTraverseSystemGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/commands/{command_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"command_id"+"}", url.PathEscape(parameterValueToString(r.commandId, "commandId")), -1)

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

type GraphApiGraphGSuiteAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	gsuiteId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;g_suite\&quot; can be associated to.
func (r GraphApiGraphGSuiteAssociationsListRequest) Targets(targets []string) GraphApiGraphGSuiteAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphGSuiteAssociationsListRequest) Limit(limit int32) GraphApiGraphGSuiteAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphGSuiteAssociationsListRequest) Skip(skip int32) GraphApiGraphGSuiteAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphGSuiteAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphGSuiteAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphGSuiteAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphGSuiteAssociationsListExecute(r)
}

/*
GraphGSuiteAssociationsList List the associations of a G Suite instance

This endpoint returns the _direct_ associations of this G Suite instance.

A direct association can be a non-homogeneous relationship between 2 different objects, for example G Suite and Users.


#### Sample Request
```
curl -X GET 'https://console.jumpcloud.com/api/v2/gsuites/{Gsuite_ID}/associations?targets=user_group \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gsuiteId ObjectID of the G Suite instance.
 @return GraphApiGraphGSuiteAssociationsListRequest
*/
func (a *GraphApiService) GraphGSuiteAssociationsList(ctx context.Context, gsuiteId string) GraphApiGraphGSuiteAssociationsListRequest {
	return GraphApiGraphGSuiteAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		gsuiteId: gsuiteId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphGSuiteAssociationsListExecute(r GraphApiGraphGSuiteAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphGSuiteAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gsuites/{gsuite_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"gsuite_id"+"}", url.PathEscape(parameterValueToString(r.gsuiteId, "gsuiteId")), -1)

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

type GraphApiGraphGSuiteAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	gsuiteId string
	xOrgId *string
	body *GraphOperationGSuite
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphGSuiteAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphGSuiteAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphGSuiteAssociationsPostRequest) Body(body GraphOperationGSuite) GraphApiGraphGSuiteAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphGSuiteAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphGSuiteAssociationsPostExecute(r)
}

/*
GraphGSuiteAssociationsPost Manage the associations of a G Suite instance

This endpoint returns the _direct_ associations of this G Suite instance.

A direct association can be a non-homogeneous relationship between 2 different objects, for example G Suite and Users.


#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/gsuites/{Gsuite_ID}/associations \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "user_group",
    "id": "{Group_ID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gsuiteId ObjectID of the G Suite instance.
 @return GraphApiGraphGSuiteAssociationsPostRequest
*/
func (a *GraphApiService) GraphGSuiteAssociationsPost(ctx context.Context, gsuiteId string) GraphApiGraphGSuiteAssociationsPostRequest {
	return GraphApiGraphGSuiteAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		gsuiteId: gsuiteId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphGSuiteAssociationsPostExecute(r GraphApiGraphGSuiteAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphGSuiteAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gsuites/{gsuite_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"gsuite_id"+"}", url.PathEscape(parameterValueToString(r.gsuiteId, "gsuiteId")), -1)

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

type GraphApiGraphGSuiteTraverseUserRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	gsuiteId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphGSuiteTraverseUserRequest) Limit(limit int32) GraphApiGraphGSuiteTraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphGSuiteTraverseUserRequest) XOrgId(xOrgId string) GraphApiGraphGSuiteTraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphGSuiteTraverseUserRequest) Skip(skip int32) GraphApiGraphGSuiteTraverseUserRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphGSuiteTraverseUserRequest) Filter(filter []string) GraphApiGraphGSuiteTraverseUserRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphGSuiteTraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphGSuiteTraverseUserExecute(r)
}

/*
GraphGSuiteTraverseUser List the Users bound to a G Suite instance

This endpoint will return all Users bound to a G Suite instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this G Suite instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this G Suite instance.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
  curl -X GET https://console.jumpcloud.com/api/v2/gsuites/{Gsuite_ID}/users \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gsuiteId ObjectID of the G Suite instance.
 @return GraphApiGraphGSuiteTraverseUserRequest
*/
func (a *GraphApiService) GraphGSuiteTraverseUser(ctx context.Context, gsuiteId string) GraphApiGraphGSuiteTraverseUserRequest {
	return GraphApiGraphGSuiteTraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		gsuiteId: gsuiteId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphGSuiteTraverseUserExecute(r GraphApiGraphGSuiteTraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphGSuiteTraverseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gsuites/{gsuite_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"gsuite_id"+"}", url.PathEscape(parameterValueToString(r.gsuiteId, "gsuiteId")), -1)

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

type GraphApiGraphGSuiteTraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	gsuiteId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphGSuiteTraverseUserGroupRequest) Limit(limit int32) GraphApiGraphGSuiteTraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphGSuiteTraverseUserGroupRequest) XOrgId(xOrgId string) GraphApiGraphGSuiteTraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphGSuiteTraverseUserGroupRequest) Skip(skip int32) GraphApiGraphGSuiteTraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphGSuiteTraverseUserGroupRequest) Filter(filter []string) GraphApiGraphGSuiteTraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphGSuiteTraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphGSuiteTraverseUserGroupExecute(r)
}

/*
GraphGSuiteTraverseUserGroup List the User Groups bound to a G Suite instance

This endpoint will return all User Groups bound to an G Suite instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this G Suite instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this G Suite instance.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
  curl -X GET https://console.jumpcloud.com/api/v2/gsuites/{GSuite_ID}/usergroups \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gsuiteId ObjectID of the G Suite instance.
 @return GraphApiGraphGSuiteTraverseUserGroupRequest
*/
func (a *GraphApiService) GraphGSuiteTraverseUserGroup(ctx context.Context, gsuiteId string) GraphApiGraphGSuiteTraverseUserGroupRequest {
	return GraphApiGraphGSuiteTraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		gsuiteId: gsuiteId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphGSuiteTraverseUserGroupExecute(r GraphApiGraphGSuiteTraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphGSuiteTraverseUserGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gsuites/{gsuite_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"gsuite_id"+"}", url.PathEscape(parameterValueToString(r.gsuiteId, "gsuiteId")), -1)

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

type GraphApiGraphLdapServerAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	ldapserverId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;ldap_server\&quot; can be associated to.
func (r GraphApiGraphLdapServerAssociationsListRequest) Targets(targets []string) GraphApiGraphLdapServerAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphLdapServerAssociationsListRequest) Limit(limit int32) GraphApiGraphLdapServerAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphLdapServerAssociationsListRequest) Skip(skip int32) GraphApiGraphLdapServerAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphLdapServerAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphLdapServerAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphLdapServerAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphLdapServerAssociationsListExecute(r)
}

/*
GraphLdapServerAssociationsList List the associations of a LDAP Server

This endpoint returns the _direct_ associations of this LDAP Server.

A direct association can be a non-homogeneous relationship between 2 different objects, for example LDAP and Users.

#### Sample Request

```
 curl -X GET 'https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/associations?targets=user_group \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ldapserverId ObjectID of the LDAP Server.
 @return GraphApiGraphLdapServerAssociationsListRequest
*/
func (a *GraphApiService) GraphLdapServerAssociationsList(ctx context.Context, ldapserverId string) GraphApiGraphLdapServerAssociationsListRequest {
	return GraphApiGraphLdapServerAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		ldapserverId: ldapserverId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphLdapServerAssociationsListExecute(r GraphApiGraphLdapServerAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphLdapServerAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ldapservers/{ldapserver_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", url.PathEscape(parameterValueToString(r.ldapserverId, "ldapserverId")), -1)

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

type GraphApiGraphLdapServerAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	ldapserverId string
	xOrgId *string
	body *GraphOperationLdapServer
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphLdapServerAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphLdapServerAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphLdapServerAssociationsPostRequest) Body(body GraphOperationLdapServer) GraphApiGraphLdapServerAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphLdapServerAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphLdapServerAssociationsPostExecute(r)
}

/*
GraphLdapServerAssociationsPost Manage the associations of a LDAP Server

This endpoint allows you to manage the _direct_ associations of a LDAP Server.

A direct association can be a non-homogeneous relationship between 2 different objects, for example LDAP and Users.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/associations \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "user",
    "id": "{User_ID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ldapserverId ObjectID of the LDAP Server.
 @return GraphApiGraphLdapServerAssociationsPostRequest
*/
func (a *GraphApiService) GraphLdapServerAssociationsPost(ctx context.Context, ldapserverId string) GraphApiGraphLdapServerAssociationsPostRequest {
	return GraphApiGraphLdapServerAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		ldapserverId: ldapserverId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphLdapServerAssociationsPostExecute(r GraphApiGraphLdapServerAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphLdapServerAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ldapservers/{ldapserver_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", url.PathEscape(parameterValueToString(r.ldapserverId, "ldapserverId")), -1)

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

type GraphApiGraphLdapServerTraverseUserRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	ldapserverId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphLdapServerTraverseUserRequest) Limit(limit int32) GraphApiGraphLdapServerTraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphLdapServerTraverseUserRequest) XOrgId(xOrgId string) GraphApiGraphLdapServerTraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphLdapServerTraverseUserRequest) Skip(skip int32) GraphApiGraphLdapServerTraverseUserRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphLdapServerTraverseUserRequest) Filter(filter []string) GraphApiGraphLdapServerTraverseUserRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphLdapServerTraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphLdapServerTraverseUserExecute(r)
}

/*
GraphLdapServerTraverseUser List the Users bound to a LDAP Server

This endpoint will return all Users bound to an LDAP Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this LDAP server instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this LDAP server instance.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/users \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ldapserverId ObjectID of the LDAP Server.
 @return GraphApiGraphLdapServerTraverseUserRequest
*/
func (a *GraphApiService) GraphLdapServerTraverseUser(ctx context.Context, ldapserverId string) GraphApiGraphLdapServerTraverseUserRequest {
	return GraphApiGraphLdapServerTraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		ldapserverId: ldapserverId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphLdapServerTraverseUserExecute(r GraphApiGraphLdapServerTraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphLdapServerTraverseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ldapservers/{ldapserver_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", url.PathEscape(parameterValueToString(r.ldapserverId, "ldapserverId")), -1)

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

type GraphApiGraphLdapServerTraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	ldapserverId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphLdapServerTraverseUserGroupRequest) Limit(limit int32) GraphApiGraphLdapServerTraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphLdapServerTraverseUserGroupRequest) XOrgId(xOrgId string) GraphApiGraphLdapServerTraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphLdapServerTraverseUserGroupRequest) Skip(skip int32) GraphApiGraphLdapServerTraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphLdapServerTraverseUserGroupRequest) Filter(filter []string) GraphApiGraphLdapServerTraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphLdapServerTraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphLdapServerTraverseUserGroupExecute(r)
}

/*
GraphLdapServerTraverseUserGroup List the User Groups bound to a LDAP Server

This endpoint will return all Users Groups bound to a LDAP Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this LDAP server instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this LDAP server instance.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/usergroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ldapserverId ObjectID of the LDAP Server.
 @return GraphApiGraphLdapServerTraverseUserGroupRequest
*/
func (a *GraphApiService) GraphLdapServerTraverseUserGroup(ctx context.Context, ldapserverId string) GraphApiGraphLdapServerTraverseUserGroupRequest {
	return GraphApiGraphLdapServerTraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		ldapserverId: ldapserverId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphLdapServerTraverseUserGroupExecute(r GraphApiGraphLdapServerTraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphLdapServerTraverseUserGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ldapservers/{ldapserver_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", url.PathEscape(parameterValueToString(r.ldapserverId, "ldapserverId")), -1)

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

type GraphApiGraphOffice365AssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	office365Id string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;office_365\&quot; can be associated to.
func (r GraphApiGraphOffice365AssociationsListRequest) Targets(targets []string) GraphApiGraphOffice365AssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphOffice365AssociationsListRequest) Limit(limit int32) GraphApiGraphOffice365AssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphOffice365AssociationsListRequest) Skip(skip int32) GraphApiGraphOffice365AssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphOffice365AssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphOffice365AssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphOffice365AssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphOffice365AssociationsListExecute(r)
}

/*
GraphOffice365AssociationsList List the associations of an Office 365 instance

This endpoint returns _direct_ associations of an Office 365 instance.


A direct association can be a non-homogeneous relationship between 2 different objects, for example Office 365 and Users.

#### Sample Request
```
curl -X GET 'https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID}/associations?targets=user_group' \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id ObjectID of the Office 365 instance.
 @return GraphApiGraphOffice365AssociationsListRequest
*/
func (a *GraphApiService) GraphOffice365AssociationsList(ctx context.Context, office365Id string) GraphApiGraphOffice365AssociationsListRequest {
	return GraphApiGraphOffice365AssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphOffice365AssociationsListExecute(r GraphApiGraphOffice365AssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphOffice365AssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", url.PathEscape(parameterValueToString(r.office365Id, "office365Id")), -1)

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

type GraphApiGraphOffice365AssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	office365Id string
	xOrgId *string
	body *GraphOperationOffice365
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphOffice365AssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphOffice365AssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphOffice365AssociationsPostRequest) Body(body GraphOperationOffice365) GraphApiGraphOffice365AssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphOffice365AssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphOffice365AssociationsPostExecute(r)
}

/*
GraphOffice365AssociationsPost Manage the associations of an Office 365 instance

This endpoint allows you to manage the _direct_ associations of a Office 365 instance.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Office 365 and Users.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID}/associations \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "user_group",
    "id": "{Group_ID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id ObjectID of the Office 365 instance.
 @return GraphApiGraphOffice365AssociationsPostRequest
*/
func (a *GraphApiService) GraphOffice365AssociationsPost(ctx context.Context, office365Id string) GraphApiGraphOffice365AssociationsPostRequest {
	return GraphApiGraphOffice365AssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphOffice365AssociationsPostExecute(r GraphApiGraphOffice365AssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphOffice365AssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", url.PathEscape(parameterValueToString(r.office365Id, "office365Id")), -1)

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

type GraphApiGraphOffice365TraverseUserRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	office365Id string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphOffice365TraverseUserRequest) Limit(limit int32) GraphApiGraphOffice365TraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphOffice365TraverseUserRequest) XOrgId(xOrgId string) GraphApiGraphOffice365TraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphOffice365TraverseUserRequest) Skip(skip int32) GraphApiGraphOffice365TraverseUserRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphOffice365TraverseUserRequest) Filter(filter []string) GraphApiGraphOffice365TraverseUserRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphOffice365TraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphOffice365TraverseUserExecute(r)
}

/*
GraphOffice365TraverseUser List the Users bound to an Office 365 instance

This endpoint will return all Users bound to an Office 365 instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Office 365 instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this Office 365 instance.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID}/users \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id ObjectID of the Office 365 suite.
 @return GraphApiGraphOffice365TraverseUserRequest
*/
func (a *GraphApiService) GraphOffice365TraverseUser(ctx context.Context, office365Id string) GraphApiGraphOffice365TraverseUserRequest {
	return GraphApiGraphOffice365TraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphOffice365TraverseUserExecute(r GraphApiGraphOffice365TraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphOffice365TraverseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", url.PathEscape(parameterValueToString(r.office365Id, "office365Id")), -1)

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

type GraphApiGraphOffice365TraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	office365Id string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphOffice365TraverseUserGroupRequest) Limit(limit int32) GraphApiGraphOffice365TraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphOffice365TraverseUserGroupRequest) XOrgId(xOrgId string) GraphApiGraphOffice365TraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphOffice365TraverseUserGroupRequest) Skip(skip int32) GraphApiGraphOffice365TraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphOffice365TraverseUserGroupRequest) Filter(filter []string) GraphApiGraphOffice365TraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphOffice365TraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphOffice365TraverseUserGroupExecute(r)
}

/*
GraphOffice365TraverseUserGroup List the User Groups bound to an Office 365 instance

This endpoint will return all Users Groups bound to an Office 365 instance, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Office 365 instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Office 365 instance.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
  curl -X GET https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID/usergroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id ObjectID of the Office 365 suite.
 @return GraphApiGraphOffice365TraverseUserGroupRequest
*/
func (a *GraphApiService) GraphOffice365TraverseUserGroup(ctx context.Context, office365Id string) GraphApiGraphOffice365TraverseUserGroupRequest {
	return GraphApiGraphOffice365TraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphOffice365TraverseUserGroupExecute(r GraphApiGraphOffice365TraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphOffice365TraverseUserGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", url.PathEscape(parameterValueToString(r.office365Id, "office365Id")), -1)

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

type GraphApiGraphPolicyAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	policyId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;policy\&quot; can be associated to.
func (r GraphApiGraphPolicyAssociationsListRequest) Targets(targets []string) GraphApiGraphPolicyAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphPolicyAssociationsListRequest) Limit(limit int32) GraphApiGraphPolicyAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphPolicyAssociationsListRequest) Skip(skip int32) GraphApiGraphPolicyAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphPolicyAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphPolicyAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphPolicyAssociationsListExecute(r)
}

/*
GraphPolicyAssociationsList List the associations of a Policy

This endpoint returns the _direct_ associations of a Policy.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Policies and Systems.

#### Sample Request
```
curl -X GET 'https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/associations?targets=system_group \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId ObjectID of the Policy.
 @return GraphApiGraphPolicyAssociationsListRequest
*/
func (a *GraphApiService) GraphPolicyAssociationsList(ctx context.Context, policyId string) GraphApiGraphPolicyAssociationsListRequest {
	return GraphApiGraphPolicyAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphPolicyAssociationsListExecute(r GraphApiGraphPolicyAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{policy_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterValueToString(r.policyId, "policyId")), -1)

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

type GraphApiGraphPolicyAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	policyId string
	xOrgId *string
	body *GraphOperationPolicy
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphPolicyAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphPolicyAssociationsPostRequest) Body(body GraphOperationPolicy) GraphApiGraphPolicyAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphPolicyAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphPolicyAssociationsPostExecute(r)
}

/*
GraphPolicyAssociationsPost Manage the associations of a Policy

This endpoint allows you to manage the _direct_ associations of a Policy.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Policies and Systems.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/associations/ \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "system_group",
    "id": "{Group_ID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId ObjectID of the Policy.
 @return GraphApiGraphPolicyAssociationsPostRequest
*/
func (a *GraphApiService) GraphPolicyAssociationsPost(ctx context.Context, policyId string) GraphApiGraphPolicyAssociationsPostRequest {
	return GraphApiGraphPolicyAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphPolicyAssociationsPostExecute(r GraphApiGraphPolicyAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{policy_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterValueToString(r.policyId, "policyId")), -1)

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

type GraphApiGraphPolicyGroupAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;policy_group\&quot; can be associated to.
func (r GraphApiGraphPolicyGroupAssociationsListRequest) Targets(targets []string) GraphApiGraphPolicyGroupAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphPolicyGroupAssociationsListRequest) Limit(limit int32) GraphApiGraphPolicyGroupAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphPolicyGroupAssociationsListRequest) Skip(skip int32) GraphApiGraphPolicyGroupAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyGroupAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphPolicyGroupAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphPolicyGroupAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphPolicyGroupAssociationsListExecute(r)
}

/*
GraphPolicyGroupAssociationsList List the associations of a Policy Group.

This endpoint returns the _direct_ associations of this Policy Group.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Policy Groups and Policies.


#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/associations?targets=system \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the Policy Group.
 @return GraphApiGraphPolicyGroupAssociationsListRequest
*/
func (a *GraphApiService) GraphPolicyGroupAssociationsList(ctx context.Context, groupId string) GraphApiGraphPolicyGroupAssociationsListRequest {
	return GraphApiGraphPolicyGroupAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphPolicyGroupAssociationsListExecute(r GraphApiGraphPolicyGroupAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyGroupAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policygroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphPolicyGroupAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	xOrgId *string
	body *GraphOperationPolicyGroup
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyGroupAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphPolicyGroupAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphPolicyGroupAssociationsPostRequest) Body(body GraphOperationPolicyGroup) GraphApiGraphPolicyGroupAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphPolicyGroupAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphPolicyGroupAssociationsPostExecute(r)
}

/*
GraphPolicyGroupAssociationsPost Manage the associations of a Policy Group

This endpoint manages the _direct_ associations of this Policy Group.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Policy Groups and Policies.


#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/associations \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "system",
    "id": "{SystemID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the Policy Group.
 @return GraphApiGraphPolicyGroupAssociationsPostRequest
*/
func (a *GraphApiService) GraphPolicyGroupAssociationsPost(ctx context.Context, groupId string) GraphApiGraphPolicyGroupAssociationsPostRequest {
	return GraphApiGraphPolicyGroupAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphPolicyGroupAssociationsPostExecute(r GraphApiGraphPolicyGroupAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyGroupAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policygroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphPolicyGroupMembersListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	skip *int32
	xOrgId *string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphPolicyGroupMembersListRequest) Limit(limit int32) GraphApiGraphPolicyGroupMembersListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphPolicyGroupMembersListRequest) Skip(skip int32) GraphApiGraphPolicyGroupMembersListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyGroupMembersListRequest) XOrgId(xOrgId string) GraphApiGraphPolicyGroupMembersListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphPolicyGroupMembersListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphPolicyGroupMembersListExecute(r)
}

/*
GraphPolicyGroupMembersList List the members of a Policy Group

This endpoint returns the Policy members of a Policy Group.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/members \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the Policy Group.
 @return GraphApiGraphPolicyGroupMembersListRequest
*/
func (a *GraphApiService) GraphPolicyGroupMembersList(ctx context.Context, groupId string) GraphApiGraphPolicyGroupMembersListRequest {
	return GraphApiGraphPolicyGroupMembersListRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphPolicyGroupMembersListExecute(r GraphApiGraphPolicyGroupMembersListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyGroupMembersList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policygroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type GraphApiGraphPolicyGroupMembersPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	xOrgId *string
	body *GraphOperationPolicyGroupMember
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyGroupMembersPostRequest) XOrgId(xOrgId string) GraphApiGraphPolicyGroupMembersPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphPolicyGroupMembersPostRequest) Body(body GraphOperationPolicyGroupMember) GraphApiGraphPolicyGroupMembersPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphPolicyGroupMembersPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphPolicyGroupMembersPostExecute(r)
}

/*
GraphPolicyGroupMembersPost Manage the members of a Policy Group

This endpoint allows you to manage the Policy members of a Policy Group.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/members \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "policy",
    "id": "{Policy_ID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the Policy Group.
 @return GraphApiGraphPolicyGroupMembersPostRequest
*/
func (a *GraphApiService) GraphPolicyGroupMembersPost(ctx context.Context, groupId string) GraphApiGraphPolicyGroupMembersPostRequest {
	return GraphApiGraphPolicyGroupMembersPostRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphPolicyGroupMembersPostExecute(r GraphApiGraphPolicyGroupMembersPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyGroupMembersPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policygroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphPolicyGroupMembershipRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	filter *[]string
	limit *int32
	skip *int32
	sort *[]string
	xOrgId *string
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphPolicyGroupMembershipRequest) Filter(filter []string) GraphApiGraphPolicyGroupMembershipRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphPolicyGroupMembershipRequest) Limit(limit int32) GraphApiGraphPolicyGroupMembershipRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphPolicyGroupMembershipRequest) Skip(skip int32) GraphApiGraphPolicyGroupMembershipRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r GraphApiGraphPolicyGroupMembershipRequest) Sort(sort []string) GraphApiGraphPolicyGroupMembershipRequest {
	r.sort = &sort
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyGroupMembershipRequest) XOrgId(xOrgId string) GraphApiGraphPolicyGroupMembershipRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphPolicyGroupMembershipRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphPolicyGroupMembershipExecute(r)
}

/*
GraphPolicyGroupMembership List the Policy Group's membership

This endpoint returns all Policy members that are a member of this Policy Group.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/membership \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the Policy Group.
 @return GraphApiGraphPolicyGroupMembershipRequest
*/
func (a *GraphApiService) GraphPolicyGroupMembership(ctx context.Context, groupId string) GraphApiGraphPolicyGroupMembershipRequest {
	return GraphApiGraphPolicyGroupMembershipRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphPolicyGroupMembershipExecute(r GraphApiGraphPolicyGroupMembershipRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyGroupMembership")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policygroups/{group_id}/membership"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphPolicyGroupTraverseSystemRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphPolicyGroupTraverseSystemRequest) Limit(limit int32) GraphApiGraphPolicyGroupTraverseSystemRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyGroupTraverseSystemRequest) XOrgId(xOrgId string) GraphApiGraphPolicyGroupTraverseSystemRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphPolicyGroupTraverseSystemRequest) Skip(skip int32) GraphApiGraphPolicyGroupTraverseSystemRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphPolicyGroupTraverseSystemRequest) Filter(filter []string) GraphApiGraphPolicyGroupTraverseSystemRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphPolicyGroupTraverseSystemRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphPolicyGroupTraverseSystemExecute(r)
}

/*
GraphPolicyGroupTraverseSystem List the Systems bound to a Policy Group

This endpoint will return all Systems bound to a Policy Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Policy Group to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Policy Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/systems \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the Policy Group.
 @return GraphApiGraphPolicyGroupTraverseSystemRequest
*/
func (a *GraphApiService) GraphPolicyGroupTraverseSystem(ctx context.Context, groupId string) GraphApiGraphPolicyGroupTraverseSystemRequest {
	return GraphApiGraphPolicyGroupTraverseSystemRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphPolicyGroupTraverseSystemExecute(r GraphApiGraphPolicyGroupTraverseSystemRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyGroupTraverseSystem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policygroups/{group_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphPolicyGroupTraverseSystemGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphPolicyGroupTraverseSystemGroupRequest) Limit(limit int32) GraphApiGraphPolicyGroupTraverseSystemGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyGroupTraverseSystemGroupRequest) XOrgId(xOrgId string) GraphApiGraphPolicyGroupTraverseSystemGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphPolicyGroupTraverseSystemGroupRequest) Skip(skip int32) GraphApiGraphPolicyGroupTraverseSystemGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphPolicyGroupTraverseSystemGroupRequest) Filter(filter []string) GraphApiGraphPolicyGroupTraverseSystemGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphPolicyGroupTraverseSystemGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphPolicyGroupTraverseSystemGroupExecute(r)
}

/*
GraphPolicyGroupTraverseSystemGroup List the System Groups bound to Policy Groups

This endpoint will return all System Groups bound to a Policy Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Policy Group to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Policy Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/policygroups/{GroupID}/systemgroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the Policy Group.
 @return GraphApiGraphPolicyGroupTraverseSystemGroupRequest
*/
func (a *GraphApiService) GraphPolicyGroupTraverseSystemGroup(ctx context.Context, groupId string) GraphApiGraphPolicyGroupTraverseSystemGroupRequest {
	return GraphApiGraphPolicyGroupTraverseSystemGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphPolicyGroupTraverseSystemGroupExecute(r GraphApiGraphPolicyGroupTraverseSystemGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyGroupTraverseSystemGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policygroups/{group_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphPolicyMemberOfRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	policyId string
	filter *[]string
	limit *int32
	skip *int32
	date *string
	authorization *string
	sort *[]string
	xOrgId *string
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphPolicyMemberOfRequest) Filter(filter []string) GraphApiGraphPolicyMemberOfRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphPolicyMemberOfRequest) Limit(limit int32) GraphApiGraphPolicyMemberOfRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphPolicyMemberOfRequest) Skip(skip int32) GraphApiGraphPolicyMemberOfRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r GraphApiGraphPolicyMemberOfRequest) Date(date string) GraphApiGraphPolicyMemberOfRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r GraphApiGraphPolicyMemberOfRequest) Authorization(authorization string) GraphApiGraphPolicyMemberOfRequest {
	r.authorization = &authorization
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r GraphApiGraphPolicyMemberOfRequest) Sort(sort []string) GraphApiGraphPolicyMemberOfRequest {
	r.sort = &sort
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyMemberOfRequest) XOrgId(xOrgId string) GraphApiGraphPolicyMemberOfRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphPolicyMemberOfRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphPolicyMemberOfExecute(r)
}

/*
GraphPolicyMemberOf List the parent Groups of a Policy

This endpoint returns all the Policy Groups a Policy is a member of.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/memberof \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId ObjectID of the Policy.
 @return GraphApiGraphPolicyMemberOfRequest
*/
func (a *GraphApiService) GraphPolicyMemberOf(ctx context.Context, policyId string) GraphApiGraphPolicyMemberOfRequest {
	return GraphApiGraphPolicyMemberOfRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphPolicyMemberOfExecute(r GraphApiGraphPolicyMemberOfRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyMemberOf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{policy_id}/memberof"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterValueToString(r.policyId, "policyId")), -1)

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

type GraphApiGraphPolicyTraverseSystemRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	policyId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphPolicyTraverseSystemRequest) Limit(limit int32) GraphApiGraphPolicyTraverseSystemRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyTraverseSystemRequest) XOrgId(xOrgId string) GraphApiGraphPolicyTraverseSystemRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphPolicyTraverseSystemRequest) Skip(skip int32) GraphApiGraphPolicyTraverseSystemRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphPolicyTraverseSystemRequest) Filter(filter []string) GraphApiGraphPolicyTraverseSystemRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphPolicyTraverseSystemRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphPolicyTraverseSystemExecute(r)
}

/*
GraphPolicyTraverseSystem List the Systems bound to a Policy

This endpoint will return all Systems bound to a Policy, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Policy to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Policy.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/systems \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId ObjectID of the Command.
 @return GraphApiGraphPolicyTraverseSystemRequest
*/
func (a *GraphApiService) GraphPolicyTraverseSystem(ctx context.Context, policyId string) GraphApiGraphPolicyTraverseSystemRequest {
	return GraphApiGraphPolicyTraverseSystemRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphPolicyTraverseSystemExecute(r GraphApiGraphPolicyTraverseSystemRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyTraverseSystem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{policy_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterValueToString(r.policyId, "policyId")), -1)

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

type GraphApiGraphPolicyTraverseSystemGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	policyId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphPolicyTraverseSystemGroupRequest) Limit(limit int32) GraphApiGraphPolicyTraverseSystemGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphPolicyTraverseSystemGroupRequest) XOrgId(xOrgId string) GraphApiGraphPolicyTraverseSystemGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphPolicyTraverseSystemGroupRequest) Skip(skip int32) GraphApiGraphPolicyTraverseSystemGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphPolicyTraverseSystemGroupRequest) Filter(filter []string) GraphApiGraphPolicyTraverseSystemGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphPolicyTraverseSystemGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphPolicyTraverseSystemGroupExecute(r)
}

/*
GraphPolicyTraverseSystemGroup List the System Groups bound to a Policy

This endpoint will return all Systems Groups bound to a Policy, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Policy to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Policy.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET  https://console.jumpcloud.com/api/v2/policies/{Policy_ID}/systemgroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyId ObjectID of the Command.
 @return GraphApiGraphPolicyTraverseSystemGroupRequest
*/
func (a *GraphApiService) GraphPolicyTraverseSystemGroup(ctx context.Context, policyId string) GraphApiGraphPolicyTraverseSystemGroupRequest {
	return GraphApiGraphPolicyTraverseSystemGroupRequest{
		ApiService: a,
		ctx: ctx,
		policyId: policyId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphPolicyTraverseSystemGroupExecute(r GraphApiGraphPolicyTraverseSystemGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphPolicyTraverseSystemGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/{policy_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"policy_id"+"}", url.PathEscape(parameterValueToString(r.policyId, "policyId")), -1)

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

type GraphApiGraphRadiusServerAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	radiusserverId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;radius_server\&quot; can be associated to.
func (r GraphApiGraphRadiusServerAssociationsListRequest) Targets(targets []string) GraphApiGraphRadiusServerAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphRadiusServerAssociationsListRequest) Limit(limit int32) GraphApiGraphRadiusServerAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphRadiusServerAssociationsListRequest) Skip(skip int32) GraphApiGraphRadiusServerAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphRadiusServerAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphRadiusServerAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphRadiusServerAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphRadiusServerAssociationsListExecute(r)
}

/*
GraphRadiusServerAssociationsList List the associations of a RADIUS  Server

This endpoint returns the _direct_ associations of a Radius Server.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Radius Servers and Users.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/radiusservers/{RADIUS_ID}/associations?targets=user_group \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param radiusserverId ObjectID of the Radius Server.
 @return GraphApiGraphRadiusServerAssociationsListRequest
*/
func (a *GraphApiService) GraphRadiusServerAssociationsList(ctx context.Context, radiusserverId string) GraphApiGraphRadiusServerAssociationsListRequest {
	return GraphApiGraphRadiusServerAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		radiusserverId: radiusserverId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphRadiusServerAssociationsListExecute(r GraphApiGraphRadiusServerAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphRadiusServerAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/radiusservers/{radiusserver_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"radiusserver_id"+"}", url.PathEscape(parameterValueToString(r.radiusserverId, "radiusserverId")), -1)

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

type GraphApiGraphRadiusServerAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	radiusserverId string
	xOrgId *string
	body *GraphOperationRadiusServer
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphRadiusServerAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphRadiusServerAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphRadiusServerAssociationsPostRequest) Body(body GraphOperationRadiusServer) GraphApiGraphRadiusServerAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphRadiusServerAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphRadiusServerAssociationsPostExecute(r)
}

/*
GraphRadiusServerAssociationsPost Manage the associations of a RADIUS Server

This endpoint allows you to manage the _direct_ associations of a Radius Server.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Radius Servers and Users.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/radiusservers/{RADIUS_ID}/associations \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
	
"type":"user", 
"id":"{USER_ID}", 
"op":"add"
	
}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param radiusserverId ObjectID of the Radius Server.
 @return GraphApiGraphRadiusServerAssociationsPostRequest
*/
func (a *GraphApiService) GraphRadiusServerAssociationsPost(ctx context.Context, radiusserverId string) GraphApiGraphRadiusServerAssociationsPostRequest {
	return GraphApiGraphRadiusServerAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		radiusserverId: radiusserverId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphRadiusServerAssociationsPostExecute(r GraphApiGraphRadiusServerAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphRadiusServerAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/radiusservers/{radiusserver_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"radiusserver_id"+"}", url.PathEscape(parameterValueToString(r.radiusserverId, "radiusserverId")), -1)

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

type GraphApiGraphRadiusServerTraverseUserRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	radiusserverId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphRadiusServerTraverseUserRequest) Limit(limit int32) GraphApiGraphRadiusServerTraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphRadiusServerTraverseUserRequest) XOrgId(xOrgId string) GraphApiGraphRadiusServerTraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphRadiusServerTraverseUserRequest) Skip(skip int32) GraphApiGraphRadiusServerTraverseUserRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphRadiusServerTraverseUserRequest) Filter(filter []string) GraphApiGraphRadiusServerTraverseUserRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphRadiusServerTraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphRadiusServerTraverseUserExecute(r)
}

/*
GraphRadiusServerTraverseUser List the Users bound to a RADIUS  Server

This endpoint will return all Users bound to a RADIUS Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this RADIUS server instance to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this RADIUS server instance.

See `/members` and `/associations` endpoints to manage those collections.


#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/ldapservers/{LDAP_ID}/users \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
  ```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param radiusserverId ObjectID of the Radius Server.
 @return GraphApiGraphRadiusServerTraverseUserRequest
*/
func (a *GraphApiService) GraphRadiusServerTraverseUser(ctx context.Context, radiusserverId string) GraphApiGraphRadiusServerTraverseUserRequest {
	return GraphApiGraphRadiusServerTraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		radiusserverId: radiusserverId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphRadiusServerTraverseUserExecute(r GraphApiGraphRadiusServerTraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphRadiusServerTraverseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/radiusservers/{radiusserver_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"radiusserver_id"+"}", url.PathEscape(parameterValueToString(r.radiusserverId, "radiusserverId")), -1)

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

type GraphApiGraphRadiusServerTraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	radiusserverId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphRadiusServerTraverseUserGroupRequest) Limit(limit int32) GraphApiGraphRadiusServerTraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphRadiusServerTraverseUserGroupRequest) XOrgId(xOrgId string) GraphApiGraphRadiusServerTraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphRadiusServerTraverseUserGroupRequest) Skip(skip int32) GraphApiGraphRadiusServerTraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphRadiusServerTraverseUserGroupRequest) Filter(filter []string) GraphApiGraphRadiusServerTraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphRadiusServerTraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphRadiusServerTraverseUserGroupExecute(r)
}

/*
GraphRadiusServerTraverseUserGroup List the User Groups bound to a RADIUS  Server

This endpoint will return all Users Groups bound to a RADIUS Server, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this RADIUS server instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this RADIUS server instance.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/radiusservers/{RADIUS_ID}/usergroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param radiusserverId ObjectID of the Radius Server.
 @return GraphApiGraphRadiusServerTraverseUserGroupRequest
*/
func (a *GraphApiService) GraphRadiusServerTraverseUserGroup(ctx context.Context, radiusserverId string) GraphApiGraphRadiusServerTraverseUserGroupRequest {
	return GraphApiGraphRadiusServerTraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		radiusserverId: radiusserverId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphRadiusServerTraverseUserGroupExecute(r GraphApiGraphRadiusServerTraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphRadiusServerTraverseUserGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/radiusservers/{radiusserver_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"radiusserver_id"+"}", url.PathEscape(parameterValueToString(r.radiusserverId, "radiusserverId")), -1)

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

type GraphApiGraphSoftwareappsAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	softwareAppId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;software_app\&quot; can be associated to.
func (r GraphApiGraphSoftwareappsAssociationsListRequest) Targets(targets []string) GraphApiGraphSoftwareappsAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSoftwareappsAssociationsListRequest) Limit(limit int32) GraphApiGraphSoftwareappsAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSoftwareappsAssociationsListRequest) Skip(skip int32) GraphApiGraphSoftwareappsAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSoftwareappsAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphSoftwareappsAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSoftwareappsAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphSoftwareappsAssociationsListExecute(r)
}

/*
GraphSoftwareappsAssociationsList List the associations of a Software Application

This endpoint will return the _direct_ associations of a Software Application. A direct association can be a non-homogeneous relationship between 2 different objects, for example Software Application and System Groups.


#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/associations?targets=system_group \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softwareAppId ObjectID of the Software App.
 @return GraphApiGraphSoftwareappsAssociationsListRequest
*/
func (a *GraphApiService) GraphSoftwareappsAssociationsList(ctx context.Context, softwareAppId string) GraphApiGraphSoftwareappsAssociationsListRequest {
	return GraphApiGraphSoftwareappsAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphSoftwareappsAssociationsListExecute(r GraphApiGraphSoftwareappsAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSoftwareappsAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{software_app_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", url.PathEscape(parameterValueToString(r.softwareAppId, "softwareAppId")), -1)

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

type GraphApiGraphSoftwareappsAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	softwareAppId string
	xOrgId *string
	body *GraphOperationSoftwareApp
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSoftwareappsAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphSoftwareappsAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSoftwareappsAssociationsPostRequest) Body(body GraphOperationSoftwareApp) GraphApiGraphSoftwareappsAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphSoftwareappsAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphSoftwareappsAssociationsPostExecute(r)
}

/*
GraphSoftwareappsAssociationsPost Manage the associations of a software application.

This endpoint allows you to associate or disassociate a software application to a system or system group.

#### Sample Request
```
$ curl -X POST https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/associations \
-H 'Accept: application/json' \
-H 'Content-Type: application/json' \
-H 'x-api-key: {API_KEY}' \
-d '{
  "id": "<object_id>",
  "op": "add",
  "type": "system"
 }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softwareAppId ObjectID of the Software App.
 @return GraphApiGraphSoftwareappsAssociationsPostRequest
*/
func (a *GraphApiService) GraphSoftwareappsAssociationsPost(ctx context.Context, softwareAppId string) GraphApiGraphSoftwareappsAssociationsPostRequest {
	return GraphApiGraphSoftwareappsAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphSoftwareappsAssociationsPostExecute(r GraphApiGraphSoftwareappsAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSoftwareappsAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{software_app_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", url.PathEscape(parameterValueToString(r.softwareAppId, "softwareAppId")), -1)

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

type GraphApiGraphSoftwareappsTraverseSystemRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	softwareAppId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSoftwareappsTraverseSystemRequest) Limit(limit int32) GraphApiGraphSoftwareappsTraverseSystemRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSoftwareappsTraverseSystemRequest) XOrgId(xOrgId string) GraphApiGraphSoftwareappsTraverseSystemRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSoftwareappsTraverseSystemRequest) Skip(skip int32) GraphApiGraphSoftwareappsTraverseSystemRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSoftwareappsTraverseSystemRequest) Filter(filter []string) GraphApiGraphSoftwareappsTraverseSystemRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSoftwareappsTraverseSystemRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSoftwareappsTraverseSystemExecute(r)
}

/*
GraphSoftwareappsTraverseSystem List the Systems bound to a Software App.

This endpoint will return all Systems bound to a Software App, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Software App to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this Software App.

See `/associations` endpoint to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/systems \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softwareAppId ObjectID of the Software App.
 @return GraphApiGraphSoftwareappsTraverseSystemRequest
*/
func (a *GraphApiService) GraphSoftwareappsTraverseSystem(ctx context.Context, softwareAppId string) GraphApiGraphSoftwareappsTraverseSystemRequest {
	return GraphApiGraphSoftwareappsTraverseSystemRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSoftwareappsTraverseSystemExecute(r GraphApiGraphSoftwareappsTraverseSystemRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSoftwareappsTraverseSystem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{software_app_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", url.PathEscape(parameterValueToString(r.softwareAppId, "softwareAppId")), -1)

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

type GraphApiGraphSoftwareappsTraverseSystemGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	softwareAppId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSoftwareappsTraverseSystemGroupRequest) Limit(limit int32) GraphApiGraphSoftwareappsTraverseSystemGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSoftwareappsTraverseSystemGroupRequest) XOrgId(xOrgId string) GraphApiGraphSoftwareappsTraverseSystemGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSoftwareappsTraverseSystemGroupRequest) Skip(skip int32) GraphApiGraphSoftwareappsTraverseSystemGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSoftwareappsTraverseSystemGroupRequest) Filter(filter []string) GraphApiGraphSoftwareappsTraverseSystemGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSoftwareappsTraverseSystemGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSoftwareappsTraverseSystemGroupExecute(r)
}

/*
GraphSoftwareappsTraverseSystemGroup List the System Groups bound to a Software App.

This endpoint will return all Systems Groups bound to a Software App, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this Software App to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this Software App.

See `/associations` endpoint to manage those collections.

#### Sample Request
```
curl -X GET  https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/systemgroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softwareAppId ObjectID of the Software App.
 @return GraphApiGraphSoftwareappsTraverseSystemGroupRequest
*/
func (a *GraphApiService) GraphSoftwareappsTraverseSystemGroup(ctx context.Context, softwareAppId string) GraphApiGraphSoftwareappsTraverseSystemGroupRequest {
	return GraphApiGraphSoftwareappsTraverseSystemGroupRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSoftwareappsTraverseSystemGroupExecute(r GraphApiGraphSoftwareappsTraverseSystemGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSoftwareappsTraverseSystemGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{software_app_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", url.PathEscape(parameterValueToString(r.softwareAppId, "softwareAppId")), -1)

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

type GraphApiGraphSystemAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	systemId string
	targets *[]string
	limit *int32
	skip *int32
	date *string
	authorization *string
	xOrgId *string
}

// Targets which a \&quot;system\&quot; can be associated to.
func (r GraphApiGraphSystemAssociationsListRequest) Targets(targets []string) GraphApiGraphSystemAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemAssociationsListRequest) Limit(limit int32) GraphApiGraphSystemAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemAssociationsListRequest) Skip(skip int32) GraphApiGraphSystemAssociationsListRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r GraphApiGraphSystemAssociationsListRequest) Date(date string) GraphApiGraphSystemAssociationsListRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r GraphApiGraphSystemAssociationsListRequest) Authorization(authorization string) GraphApiGraphSystemAssociationsListRequest {
	r.authorization = &authorization
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphSystemAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSystemAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
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
 @return GraphApiGraphSystemAssociationsListRequest
*/
func (a *GraphApiService) GraphSystemAssociationsList(ctx context.Context, systemId string) GraphApiGraphSystemAssociationsListRequest {
	return GraphApiGraphSystemAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphSystemAssociationsListExecute(r GraphApiGraphSystemAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemAssociationsList")
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

type GraphApiGraphSystemAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	systemId string
	date *string
	authorization *string
	xOrgId *string
	body *GraphOperationSystem
}

// Current date header for the System Context API
func (r GraphApiGraphSystemAssociationsPostRequest) Date(date string) GraphApiGraphSystemAssociationsPostRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r GraphApiGraphSystemAssociationsPostRequest) Authorization(authorization string) GraphApiGraphSystemAssociationsPostRequest {
	r.authorization = &authorization
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphSystemAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSystemAssociationsPostRequest) Body(body GraphOperationSystem) GraphApiGraphSystemAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphSystemAssociationsPostRequest) Execute() (*http.Response, error) {
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
 @return GraphApiGraphSystemAssociationsPostRequest
*/
func (a *GraphApiService) GraphSystemAssociationsPost(ctx context.Context, systemId string) GraphApiGraphSystemAssociationsPostRequest {
	return GraphApiGraphSystemAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphSystemAssociationsPostExecute(r GraphApiGraphSystemAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemAssociationsPost")
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

type GraphApiGraphSystemGroupAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;system_group\&quot; can be associated to.
func (r GraphApiGraphSystemGroupAssociationsListRequest) Targets(targets []string) GraphApiGraphSystemGroupAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemGroupAssociationsListRequest) Limit(limit int32) GraphApiGraphSystemGroupAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemGroupAssociationsListRequest) Skip(skip int32) GraphApiGraphSystemGroupAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSystemGroupAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphSystemGroupAssociationsListExecute(r)
}

/*
GraphSystemGroupAssociationsList List the associations of a System Group

This endpoint returns the _direct_ associations of a System Group.

A direct association can be a non-homogeneous relationship between 2 different objects, for example System Groups and Users.


#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/associations?targets=user \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupAssociationsListRequest
*/
func (a *GraphApiService) GraphSystemGroupAssociationsList(ctx context.Context, groupId string) GraphApiGraphSystemGroupAssociationsListRequest {
	return GraphApiGraphSystemGroupAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphSystemGroupAssociationsListExecute(r GraphApiGraphSystemGroupAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphSystemGroupAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	xOrgId *string
	body *GraphOperationSystemGroup
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSystemGroupAssociationsPostRequest) Body(body GraphOperationSystemGroup) GraphApiGraphSystemGroupAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphSystemGroupAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphSystemGroupAssociationsPostExecute(r)
}

/*
GraphSystemGroupAssociationsPost Manage the associations of a System Group

This endpoint allows you to manage the _direct_ associations of a System Group.

A direct association can be a non-homogeneous relationship between 2 different objects, for example System Groups and Users.


#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/associations \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "user",
    "id": "{UserID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupAssociationsPostRequest
*/
func (a *GraphApiService) GraphSystemGroupAssociationsPost(ctx context.Context, groupId string) GraphApiGraphSystemGroupAssociationsPostRequest {
	return GraphApiGraphSystemGroupAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphSystemGroupAssociationsPostExecute(r GraphApiGraphSystemGroupAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphSystemGroupMembersListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	skip *int32
	xOrgId *string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemGroupMembersListRequest) Limit(limit int32) GraphApiGraphSystemGroupMembersListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemGroupMembersListRequest) Skip(skip int32) GraphApiGraphSystemGroupMembersListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupMembersListRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupMembersListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSystemGroupMembersListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphSystemGroupMembersListExecute(r)
}

/*
GraphSystemGroupMembersList List the members of a System Group

This endpoint returns the system members of a System Group.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{Group_ID}/members \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupMembersListRequest
*/
func (a *GraphApiService) GraphSystemGroupMembersList(ctx context.Context, groupId string) GraphApiGraphSystemGroupMembersListRequest {
	return GraphApiGraphSystemGroupMembersListRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphSystemGroupMembersListExecute(r GraphApiGraphSystemGroupMembersListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupMembersList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type GraphApiGraphSystemGroupMembersPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	date *string
	authorization *string
	xOrgId *string
	body *GraphOperationSystemGroupMember
}

// Current date header for the System Context API
func (r GraphApiGraphSystemGroupMembersPostRequest) Date(date string) GraphApiGraphSystemGroupMembersPostRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r GraphApiGraphSystemGroupMembersPostRequest) Authorization(authorization string) GraphApiGraphSystemGroupMembersPostRequest {
	r.authorization = &authorization
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupMembersPostRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupMembersPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSystemGroupMembersPostRequest) Body(body GraphOperationSystemGroupMember) GraphApiGraphSystemGroupMembersPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphSystemGroupMembersPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphSystemGroupMembersPostExecute(r)
}

/*
GraphSystemGroupMembersPost Manage the members of a System Group

This endpoint allows you to manage the system members of a System Group.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/systemgroups/{Group_ID}/members \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "system",
    "id": "{System_ID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupMembersPostRequest
*/
func (a *GraphApiService) GraphSystemGroupMembersPost(ctx context.Context, groupId string) GraphApiGraphSystemGroupMembersPostRequest {
	return GraphApiGraphSystemGroupMembersPostRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphSystemGroupMembersPostExecute(r GraphApiGraphSystemGroupMembersPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupMembersPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphSystemGroupMembershipRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemGroupMembershipRequest) Limit(limit int32) GraphApiGraphSystemGroupMembershipRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemGroupMembershipRequest) Skip(skip int32) GraphApiGraphSystemGroupMembershipRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r GraphApiGraphSystemGroupMembershipRequest) Sort(sort []string) GraphApiGraphSystemGroupMembershipRequest {
	r.sort = &sort
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemGroupMembershipRequest) Filter(filter []string) GraphApiGraphSystemGroupMembershipRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupMembershipRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupMembershipRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSystemGroupMembershipRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemGroupMembershipExecute(r)
}

/*
GraphSystemGroupMembership List the System Group's membership

This endpoint returns all Systems that are a member of this System Group.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{Group_ID/membership \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupMembershipRequest
*/
func (a *GraphApiService) GraphSystemGroupMembership(ctx context.Context, groupId string) GraphApiGraphSystemGroupMembershipRequest {
	return GraphApiGraphSystemGroupMembershipRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemGroupMembershipExecute(r GraphApiGraphSystemGroupMembershipRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupMembership")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/membership"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
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

type GraphApiGraphSystemGroupTraverseCommandRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemGroupTraverseCommandRequest) Limit(limit int32) GraphApiGraphSystemGroupTraverseCommandRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupTraverseCommandRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupTraverseCommandRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemGroupTraverseCommandRequest) Skip(skip int32) GraphApiGraphSystemGroupTraverseCommandRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemGroupTraverseCommandRequest) Filter(filter []string) GraphApiGraphSystemGroupTraverseCommandRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemGroupTraverseCommandRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemGroupTraverseCommandExecute(r)
}

/*
GraphSystemGroupTraverseCommand List the Commands bound to a System Group

This endpoint will return all Commands bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the group's type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System Group to the corresponding Command; this array represents all grouping and/or associations that would have to be removed to deprovision the Command from this System Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/commands \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupTraverseCommandRequest
*/
func (a *GraphApiService) GraphSystemGroupTraverseCommand(ctx context.Context, groupId string) GraphApiGraphSystemGroupTraverseCommandRequest {
	return GraphApiGraphSystemGroupTraverseCommandRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemGroupTraverseCommandExecute(r GraphApiGraphSystemGroupTraverseCommandRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupTraverseCommand")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/commands"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphSystemGroupTraversePolicyRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemGroupTraversePolicyRequest) Limit(limit int32) GraphApiGraphSystemGroupTraversePolicyRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupTraversePolicyRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupTraversePolicyRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemGroupTraversePolicyRequest) Skip(skip int32) GraphApiGraphSystemGroupTraversePolicyRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemGroupTraversePolicyRequest) Filter(filter []string) GraphApiGraphSystemGroupTraversePolicyRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemGroupTraversePolicyRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemGroupTraversePolicyExecute(r)
}

/*
GraphSystemGroupTraversePolicy List the Policies bound to a System Group

This endpoint will return all Policies bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System Group to the corresponding Policy; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy from this System Group.

See `/members` and `/associations` endpoints to manage those collections.

This endpoint is not public yet as we haven't finished the code.

##### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/policies \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupTraversePolicyRequest
*/
func (a *GraphApiService) GraphSystemGroupTraversePolicy(ctx context.Context, groupId string) GraphApiGraphSystemGroupTraversePolicyRequest {
	return GraphApiGraphSystemGroupTraversePolicyRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemGroupTraversePolicyExecute(r GraphApiGraphSystemGroupTraversePolicyRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupTraversePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/policies"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphSystemGroupTraversePolicyGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemGroupTraversePolicyGroupRequest) Limit(limit int32) GraphApiGraphSystemGroupTraversePolicyGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupTraversePolicyGroupRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupTraversePolicyGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemGroupTraversePolicyGroupRequest) Skip(skip int32) GraphApiGraphSystemGroupTraversePolicyGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemGroupTraversePolicyGroupRequest) Filter(filter []string) GraphApiGraphSystemGroupTraversePolicyGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemGroupTraversePolicyGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemGroupTraversePolicyGroupExecute(r)
}

/*
GraphSystemGroupTraversePolicyGroup List the Policy Groups bound to a System Group

This endpoint will return all Policy Groups bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System Group to the corresponding Policy Group; this array represents all grouping and/or associations that would have to be removed to deprovision the Policy Group from this System Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/policygroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupTraversePolicyGroupRequest
*/
func (a *GraphApiService) GraphSystemGroupTraversePolicyGroup(ctx context.Context, groupId string) GraphApiGraphSystemGroupTraversePolicyGroupRequest {
	return GraphApiGraphSystemGroupTraversePolicyGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemGroupTraversePolicyGroupExecute(r GraphApiGraphSystemGroupTraversePolicyGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupTraversePolicyGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/policygroups"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphSystemGroupTraverseUserRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemGroupTraverseUserRequest) Limit(limit int32) GraphApiGraphSystemGroupTraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupTraverseUserRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupTraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemGroupTraverseUserRequest) Skip(skip int32) GraphApiGraphSystemGroupTraverseUserRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemGroupTraverseUserRequest) Filter(filter []string) GraphApiGraphSystemGroupTraverseUserRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemGroupTraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemGroupTraverseUserExecute(r)
}

/*
GraphSystemGroupTraverseUser List the Users bound to a System Group

This endpoint will return all Users bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System Group to the corresponding User; this array represents all grouping and/or associations that would have to be removed to deprovision the User from this System Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/users \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupTraverseUserRequest
*/
func (a *GraphApiService) GraphSystemGroupTraverseUser(ctx context.Context, groupId string) GraphApiGraphSystemGroupTraverseUserRequest {
	return GraphApiGraphSystemGroupTraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemGroupTraverseUserExecute(r GraphApiGraphSystemGroupTraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupTraverseUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphSystemGroupTraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemGroupTraverseUserGroupRequest) Limit(limit int32) GraphApiGraphSystemGroupTraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemGroupTraverseUserGroupRequest) XOrgId(xOrgId string) GraphApiGraphSystemGroupTraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemGroupTraverseUserGroupRequest) Skip(skip int32) GraphApiGraphSystemGroupTraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemGroupTraverseUserGroupRequest) Filter(filter []string) GraphApiGraphSystemGroupTraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemGroupTraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphSystemGroupTraverseUserGroupExecute(r)
}

/*
GraphSystemGroupTraverseUserGroup List the User Groups bound to a System Group

This endpoint will return all User Groups bound to a System Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this System Group to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this System Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/systemgroups/{GroupID}/usergroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the System Group.
 @return GraphApiGraphSystemGroupTraverseUserGroupRequest
*/
func (a *GraphApiService) GraphSystemGroupTraverseUserGroup(ctx context.Context, groupId string) GraphApiGraphSystemGroupTraverseUserGroupRequest {
	return GraphApiGraphSystemGroupTraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemGroupTraverseUserGroupExecute(r GraphApiGraphSystemGroupTraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemGroupTraverseUserGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systemgroups/{group_id}/usergroups"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphSystemMemberOfRequest struct {
	ctx context.Context
	ApiService *GraphApiService
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
func (r GraphApiGraphSystemMemberOfRequest) Filter(filter []string) GraphApiGraphSystemMemberOfRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemMemberOfRequest) Limit(limit int32) GraphApiGraphSystemMemberOfRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemMemberOfRequest) Skip(skip int32) GraphApiGraphSystemMemberOfRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r GraphApiGraphSystemMemberOfRequest) Date(date string) GraphApiGraphSystemMemberOfRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r GraphApiGraphSystemMemberOfRequest) Authorization(authorization string) GraphApiGraphSystemMemberOfRequest {
	r.authorization = &authorization
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r GraphApiGraphSystemMemberOfRequest) Sort(sort []string) GraphApiGraphSystemMemberOfRequest {
	r.sort = &sort
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemMemberOfRequest) XOrgId(xOrgId string) GraphApiGraphSystemMemberOfRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphSystemMemberOfRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return GraphApiGraphSystemMemberOfRequest
*/
func (a *GraphApiService) GraphSystemMemberOf(ctx context.Context, systemId string) GraphApiGraphSystemMemberOfRequest {
	return GraphApiGraphSystemMemberOfRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemMemberOfExecute(r GraphApiGraphSystemMemberOfRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemMemberOf")
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

type GraphApiGraphSystemTraverseCommandRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemTraverseCommandRequest) Limit(limit int32) GraphApiGraphSystemTraverseCommandRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemTraverseCommandRequest) XOrgId(xOrgId string) GraphApiGraphSystemTraverseCommandRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemTraverseCommandRequest) Skip(skip int32) GraphApiGraphSystemTraverseCommandRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemTraverseCommandRequest) Filter(filter []string) GraphApiGraphSystemTraverseCommandRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemTraverseCommandRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return GraphApiGraphSystemTraverseCommandRequest
*/
func (a *GraphApiService) GraphSystemTraverseCommand(ctx context.Context, systemId string) GraphApiGraphSystemTraverseCommandRequest {
	return GraphApiGraphSystemTraverseCommandRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemTraverseCommandExecute(r GraphApiGraphSystemTraverseCommandRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemTraverseCommand")
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

type GraphApiGraphSystemTraversePolicyRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemTraversePolicyRequest) Limit(limit int32) GraphApiGraphSystemTraversePolicyRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemTraversePolicyRequest) XOrgId(xOrgId string) GraphApiGraphSystemTraversePolicyRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemTraversePolicyRequest) Skip(skip int32) GraphApiGraphSystemTraversePolicyRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemTraversePolicyRequest) Filter(filter []string) GraphApiGraphSystemTraversePolicyRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemTraversePolicyRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return GraphApiGraphSystemTraversePolicyRequest
*/
func (a *GraphApiService) GraphSystemTraversePolicy(ctx context.Context, systemId string) GraphApiGraphSystemTraversePolicyRequest {
	return GraphApiGraphSystemTraversePolicyRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemTraversePolicyExecute(r GraphApiGraphSystemTraversePolicyRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemTraversePolicy")
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

type GraphApiGraphSystemTraversePolicyGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	date *string
	authorization *string
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemTraversePolicyGroupRequest) Limit(limit int32) GraphApiGraphSystemTraversePolicyGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemTraversePolicyGroupRequest) XOrgId(xOrgId string) GraphApiGraphSystemTraversePolicyGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemTraversePolicyGroupRequest) Skip(skip int32) GraphApiGraphSystemTraversePolicyGroupRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r GraphApiGraphSystemTraversePolicyGroupRequest) Date(date string) GraphApiGraphSystemTraversePolicyGroupRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r GraphApiGraphSystemTraversePolicyGroupRequest) Authorization(authorization string) GraphApiGraphSystemTraversePolicyGroupRequest {
	r.authorization = &authorization
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemTraversePolicyGroupRequest) Filter(filter []string) GraphApiGraphSystemTraversePolicyGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemTraversePolicyGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return GraphApiGraphSystemTraversePolicyGroupRequest
*/
func (a *GraphApiService) GraphSystemTraversePolicyGroup(ctx context.Context, systemId string) GraphApiGraphSystemTraversePolicyGroupRequest {
	return GraphApiGraphSystemTraversePolicyGroupRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemTraversePolicyGroupExecute(r GraphApiGraphSystemTraversePolicyGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemTraversePolicyGroup")
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

type GraphApiGraphSystemTraverseUserRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	date *string
	authorization *string
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemTraverseUserRequest) Limit(limit int32) GraphApiGraphSystemTraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemTraverseUserRequest) XOrgId(xOrgId string) GraphApiGraphSystemTraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemTraverseUserRequest) Skip(skip int32) GraphApiGraphSystemTraverseUserRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r GraphApiGraphSystemTraverseUserRequest) Date(date string) GraphApiGraphSystemTraverseUserRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r GraphApiGraphSystemTraverseUserRequest) Authorization(authorization string) GraphApiGraphSystemTraverseUserRequest {
	r.authorization = &authorization
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemTraverseUserRequest) Filter(filter []string) GraphApiGraphSystemTraverseUserRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemTraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return GraphApiGraphSystemTraverseUserRequest
*/
func (a *GraphApiService) GraphSystemTraverseUser(ctx context.Context, systemId string) GraphApiGraphSystemTraverseUserRequest {
	return GraphApiGraphSystemTraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemTraverseUserExecute(r GraphApiGraphSystemTraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemTraverseUser")
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

type GraphApiGraphSystemTraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	systemId string
	limit *int32
	xOrgId *string
	skip *int32
	date *string
	authorization *string
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphSystemTraverseUserGroupRequest) Limit(limit int32) GraphApiGraphSystemTraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphSystemTraverseUserGroupRequest) XOrgId(xOrgId string) GraphApiGraphSystemTraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphSystemTraverseUserGroupRequest) Skip(skip int32) GraphApiGraphSystemTraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// Current date header for the System Context API
func (r GraphApiGraphSystemTraverseUserGroupRequest) Date(date string) GraphApiGraphSystemTraverseUserGroupRequest {
	r.date = &date
	return r
}

// Authorization header for the System Context API
func (r GraphApiGraphSystemTraverseUserGroupRequest) Authorization(authorization string) GraphApiGraphSystemTraverseUserGroupRequest {
	r.authorization = &authorization
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphSystemTraverseUserGroupRequest) Filter(filter []string) GraphApiGraphSystemTraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphSystemTraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return GraphApiGraphSystemTraverseUserGroupRequest
*/
func (a *GraphApiService) GraphSystemTraverseUserGroup(ctx context.Context, systemId string) GraphApiGraphSystemTraverseUserGroupRequest {
	return GraphApiGraphSystemTraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphSystemTraverseUserGroupExecute(r GraphApiGraphSystemTraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphSystemTraverseUserGroup")
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

type GraphApiGraphUserAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;user\&quot; can be associated to.
func (r GraphApiGraphUserAssociationsListRequest) Targets(targets []string) GraphApiGraphUserAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserAssociationsListRequest) Limit(limit int32) GraphApiGraphUserAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserAssociationsListRequest) Skip(skip int32) GraphApiGraphUserAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphUserAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphUserAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphUserAssociationsListExecute(r)
}

/*
GraphUserAssociationsList List the associations of a User

This endpoint returns the _direct_ associations of a User.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Users and Systems.


#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/associations?targets=system_group \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserAssociationsListRequest
*/
func (a *GraphApiService) GraphUserAssociationsList(ctx context.Context, userId string) GraphApiGraphUserAssociationsListRequest {
	return GraphApiGraphUserAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphUserAssociationsListExecute(r GraphApiGraphUserAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	xOrgId *string
	body *GraphOperationUser
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphUserAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphUserAssociationsPostRequest) Body(body GraphOperationUser) GraphApiGraphUserAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphUserAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphUserAssociationsPostExecute(r)
}

/*
GraphUserAssociationsPost Manage the associations of a User

This endpoint allows you to manage the _direct_ associations of a User.

A direct association can be a non-homogeneous relationship between 2 different objects, for example Users and Systems.


#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/users/{UserID}/associations \
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
    "type": "system_group",
    "id": "{GroupID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserAssociationsPostRequest
*/
func (a *GraphApiService) GraphUserAssociationsPost(ctx context.Context, userId string) GraphApiGraphUserAssociationsPostRequest {
	return GraphApiGraphUserAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphUserAssociationsPostExecute(r GraphApiGraphUserAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserGroupAssociationsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;user_group\&quot; can be associated to.
func (r GraphApiGraphUserGroupAssociationsListRequest) Targets(targets []string) GraphApiGraphUserGroupAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupAssociationsListRequest) Limit(limit int32) GraphApiGraphUserGroupAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupAssociationsListRequest) Skip(skip int32) GraphApiGraphUserGroupAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupAssociationsListRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphUserGroupAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphUserGroupAssociationsListExecute(r)
}

/*
GraphUserGroupAssociationsList List the associations of a User Group.

This endpoint returns the _direct_ associations of this User Group.

A direct association can be a non-homogeneous relationship between 2 different objects, for example User Groups and Users.


#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/associations?targets=system \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupAssociationsListRequest
*/
func (a *GraphApiService) GraphUserGroupAssociationsList(ctx context.Context, groupId string) GraphApiGraphUserGroupAssociationsListRequest {
	return GraphApiGraphUserGroupAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphUserGroupAssociationsListExecute(r GraphApiGraphUserGroupAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupAssociationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupAssociationsPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	xOrgId *string
	body *GraphOperationUserGroup
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupAssociationsPostRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphUserGroupAssociationsPostRequest) Body(body GraphOperationUserGroup) GraphApiGraphUserGroupAssociationsPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphUserGroupAssociationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphUserGroupAssociationsPostExecute(r)
}

/*
GraphUserGroupAssociationsPost Manage the associations of a User Group

This endpoint manages the _direct_ associations of this User Group.

A direct association can be a non-homogeneous relationship between 2 different objects, for example User Groups and Users.


#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/associations \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "system",
    "id": "{SystemID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupAssociationsPostRequest
*/
func (a *GraphApiService) GraphUserGroupAssociationsPost(ctx context.Context, groupId string) GraphApiGraphUserGroupAssociationsPostRequest {
	return GraphApiGraphUserGroupAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphUserGroupAssociationsPostExecute(r GraphApiGraphUserGroupAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupAssociationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupMembersListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	skip *int32
	xOrgId *string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupMembersListRequest) Limit(limit int32) GraphApiGraphUserGroupMembersListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupMembersListRequest) Skip(skip int32) GraphApiGraphUserGroupMembersListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupMembersListRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupMembersListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphUserGroupMembersListRequest) Execute() ([]GraphConnection, *http.Response, error) {
	return r.ApiService.GraphUserGroupMembersListExecute(r)
}

/*
GraphUserGroupMembersList List the members of a User Group

This endpoint returns the user members of a User Group.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/members \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupMembersListRequest
*/
func (a *GraphApiService) GraphUserGroupMembersList(ctx context.Context, groupId string) GraphApiGraphUserGroupMembersListRequest {
	return GraphApiGraphUserGroupMembersListRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *GraphApiService) GraphUserGroupMembersListExecute(r GraphApiGraphUserGroupMembersListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupMembersList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type GraphApiGraphUserGroupMembersPostRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	xOrgId *string
	body *GraphOperationUserGroupMember
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupMembersPostRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupMembersPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphUserGroupMembersPostRequest) Body(body GraphOperationUserGroupMember) GraphApiGraphUserGroupMembersPostRequest {
	r.body = &body
	return r
}

func (r GraphApiGraphUserGroupMembersPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.GraphUserGroupMembersPostExecute(r)
}

/*
GraphUserGroupMembersPost Manage the members of a User Group

This endpoint allows you to manage the user members of a User Group.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/members \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "op": "add",
    "type": "user",
    "id": "{User_ID}"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupMembersPostRequest
*/
func (a *GraphApiService) GraphUserGroupMembersPost(ctx context.Context, groupId string) GraphApiGraphUserGroupMembersPostRequest {
	return GraphApiGraphUserGroupMembersPostRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GraphApiService) GraphUserGroupMembersPostExecute(r GraphApiGraphUserGroupMembersPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupMembersPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupMembershipRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	filter *[]string
	limit *int32
	skip *int32
	sort *[]string
	xOrgId *string
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupMembershipRequest) Filter(filter []string) GraphApiGraphUserGroupMembershipRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupMembershipRequest) Limit(limit int32) GraphApiGraphUserGroupMembershipRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupMembershipRequest) Skip(skip int32) GraphApiGraphUserGroupMembershipRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r GraphApiGraphUserGroupMembershipRequest) Sort(sort []string) GraphApiGraphUserGroupMembershipRequest {
	r.sort = &sort
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupMembershipRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupMembershipRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphUserGroupMembershipRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupMembershipExecute(r)
}

/*
GraphUserGroupMembership List the User Group's membership

This endpoint returns all users members that are a member of this User Group.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/membership \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupMembershipRequest
*/
func (a *GraphApiService) GraphUserGroupMembership(ctx context.Context, groupId string) GraphApiGraphUserGroupMembershipRequest {
	return GraphApiGraphUserGroupMembershipRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupMembershipExecute(r GraphApiGraphUserGroupMembershipRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupMembership")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/membership"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupTraverseActiveDirectoryRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupTraverseActiveDirectoryRequest) Limit(limit int32) GraphApiGraphUserGroupTraverseActiveDirectoryRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupTraverseActiveDirectoryRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupTraverseActiveDirectoryRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupTraverseActiveDirectoryRequest) Skip(skip int32) GraphApiGraphUserGroupTraverseActiveDirectoryRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupTraverseActiveDirectoryRequest) Filter(filter []string) GraphApiGraphUserGroupTraverseActiveDirectoryRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserGroupTraverseActiveDirectoryRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupTraverseActiveDirectoryExecute(r)
}

/*
GraphUserGroupTraverseActiveDirectory List the Active Directories bound to a User Group

This endpoint will return all Active Directory Instances bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User Group to the corresponding Active Directory; this array represents all grouping and/or associations that would have to be removed to deprovision the Active Directory from this User Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/activedirectories \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupTraverseActiveDirectoryRequest
*/
func (a *GraphApiService) GraphUserGroupTraverseActiveDirectory(ctx context.Context, groupId string) GraphApiGraphUserGroupTraverseActiveDirectoryRequest {
	return GraphApiGraphUserGroupTraverseActiveDirectoryRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupTraverseActiveDirectoryExecute(r GraphApiGraphUserGroupTraverseActiveDirectoryRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupTraverseActiveDirectory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/activedirectories"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupTraverseApplicationRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupTraverseApplicationRequest) Limit(limit int32) GraphApiGraphUserGroupTraverseApplicationRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupTraverseApplicationRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupTraverseApplicationRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupTraverseApplicationRequest) Skip(skip int32) GraphApiGraphUserGroupTraverseApplicationRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupTraverseApplicationRequest) Filter(filter []string) GraphApiGraphUserGroupTraverseApplicationRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserGroupTraverseApplicationRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupTraverseApplicationExecute(r)
}

/*
GraphUserGroupTraverseApplication List the Applications bound to a User Group

This endpoint will return all Applications bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User Group to the corresponding Application; this array represents all grouping and/or associations that would have to be removed to deprovision the Application from this User Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/applications \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupTraverseApplicationRequest
*/
func (a *GraphApiService) GraphUserGroupTraverseApplication(ctx context.Context, groupId string) GraphApiGraphUserGroupTraverseApplicationRequest {
	return GraphApiGraphUserGroupTraverseApplicationRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupTraverseApplicationExecute(r GraphApiGraphUserGroupTraverseApplicationRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupTraverseApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupTraverseDirectoryRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupTraverseDirectoryRequest) Limit(limit int32) GraphApiGraphUserGroupTraverseDirectoryRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupTraverseDirectoryRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupTraverseDirectoryRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupTraverseDirectoryRequest) Skip(skip int32) GraphApiGraphUserGroupTraverseDirectoryRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupTraverseDirectoryRequest) Filter(filter []string) GraphApiGraphUserGroupTraverseDirectoryRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserGroupTraverseDirectoryRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupTraverseDirectoryExecute(r)
}

/*
GraphUserGroupTraverseDirectory List the Directories bound to a User Group

This endpoint will return all Directories bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User Group to the corresponding Directory; this array represents all grouping and/or associations that would have to be removed to deprovision the Directories from this User Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/directories \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupTraverseDirectoryRequest
*/
func (a *GraphApiService) GraphUserGroupTraverseDirectory(ctx context.Context, groupId string) GraphApiGraphUserGroupTraverseDirectoryRequest {
	return GraphApiGraphUserGroupTraverseDirectoryRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupTraverseDirectoryExecute(r GraphApiGraphUserGroupTraverseDirectoryRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupTraverseDirectory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/directories"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupTraverseGSuiteRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupTraverseGSuiteRequest) Limit(limit int32) GraphApiGraphUserGroupTraverseGSuiteRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupTraverseGSuiteRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupTraverseGSuiteRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupTraverseGSuiteRequest) Skip(skip int32) GraphApiGraphUserGroupTraverseGSuiteRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupTraverseGSuiteRequest) Filter(filter []string) GraphApiGraphUserGroupTraverseGSuiteRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserGroupTraverseGSuiteRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupTraverseGSuiteExecute(r)
}

/*
GraphUserGroupTraverseGSuite List the G Suite instances bound to a User Group

This endpoint will return all G Suite Instances bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User Group to the corresponding G Suite instance; this array represents all grouping and/or associations that would have to be removed to deprovision the G Suite instance from this User Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID/gsuites \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupTraverseGSuiteRequest
*/
func (a *GraphApiService) GraphUserGroupTraverseGSuite(ctx context.Context, groupId string) GraphApiGraphUserGroupTraverseGSuiteRequest {
	return GraphApiGraphUserGroupTraverseGSuiteRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupTraverseGSuiteExecute(r GraphApiGraphUserGroupTraverseGSuiteRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupTraverseGSuite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/gsuites"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupTraverseLdapServerRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupTraverseLdapServerRequest) Limit(limit int32) GraphApiGraphUserGroupTraverseLdapServerRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupTraverseLdapServerRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupTraverseLdapServerRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupTraverseLdapServerRequest) Skip(skip int32) GraphApiGraphUserGroupTraverseLdapServerRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupTraverseLdapServerRequest) Filter(filter []string) GraphApiGraphUserGroupTraverseLdapServerRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserGroupTraverseLdapServerRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupTraverseLdapServerExecute(r)
}

/*
GraphUserGroupTraverseLdapServer List the LDAP Servers bound to a User Group

This endpoint will return all LDAP Servers bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User Group to the corresponding LDAP Server; this array represents all grouping and/or associations that would have to be removed to deprovision the LDAP Server from this User Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/ldapservers \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupTraverseLdapServerRequest
*/
func (a *GraphApiService) GraphUserGroupTraverseLdapServer(ctx context.Context, groupId string) GraphApiGraphUserGroupTraverseLdapServerRequest {
	return GraphApiGraphUserGroupTraverseLdapServerRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupTraverseLdapServerExecute(r GraphApiGraphUserGroupTraverseLdapServerRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupTraverseLdapServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/ldapservers"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupTraverseOffice365Request struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupTraverseOffice365Request) Limit(limit int32) GraphApiGraphUserGroupTraverseOffice365Request {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupTraverseOffice365Request) XOrgId(xOrgId string) GraphApiGraphUserGroupTraverseOffice365Request {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupTraverseOffice365Request) Skip(skip int32) GraphApiGraphUserGroupTraverseOffice365Request {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupTraverseOffice365Request) Filter(filter []string) GraphApiGraphUserGroupTraverseOffice365Request {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserGroupTraverseOffice365Request) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupTraverseOffice365Execute(r)
}

/*
GraphUserGroupTraverseOffice365 List the Office 365 instances bound to a User Group

This endpoint will return all Office 365 instances bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User Group to the corresponding Office 365 instance; this array represents all grouping and/or associations that would have to be removed to deprovision the Office 365 instance from this User Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/office365s \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupTraverseOffice365Request
*/
func (a *GraphApiService) GraphUserGroupTraverseOffice365(ctx context.Context, groupId string) GraphApiGraphUserGroupTraverseOffice365Request {
	return GraphApiGraphUserGroupTraverseOffice365Request{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupTraverseOffice365Execute(r GraphApiGraphUserGroupTraverseOffice365Request) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupTraverseOffice365")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/office365s"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupTraverseRadiusServerRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupTraverseRadiusServerRequest) Limit(limit int32) GraphApiGraphUserGroupTraverseRadiusServerRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupTraverseRadiusServerRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupTraverseRadiusServerRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupTraverseRadiusServerRequest) Skip(skip int32) GraphApiGraphUserGroupTraverseRadiusServerRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupTraverseRadiusServerRequest) Filter(filter []string) GraphApiGraphUserGroupTraverseRadiusServerRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserGroupTraverseRadiusServerRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupTraverseRadiusServerExecute(r)
}

/*
GraphUserGroupTraverseRadiusServer List the RADIUS Servers bound to a User Group

This endpoint will return all RADIUS servers bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User Group to the corresponding RADIUS Server; this array represents all grouping and/or associations that would have to be removed to deprovision the RADIUS Server from this User Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/radiusservers \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupTraverseRadiusServerRequest
*/
func (a *GraphApiService) GraphUserGroupTraverseRadiusServer(ctx context.Context, groupId string) GraphApiGraphUserGroupTraverseRadiusServerRequest {
	return GraphApiGraphUserGroupTraverseRadiusServerRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupTraverseRadiusServerExecute(r GraphApiGraphUserGroupTraverseRadiusServerRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupTraverseRadiusServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/radiusservers"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupTraverseSystemRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupTraverseSystemRequest) Limit(limit int32) GraphApiGraphUserGroupTraverseSystemRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupTraverseSystemRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupTraverseSystemRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupTraverseSystemRequest) Skip(skip int32) GraphApiGraphUserGroupTraverseSystemRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupTraverseSystemRequest) Filter(filter []string) GraphApiGraphUserGroupTraverseSystemRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserGroupTraverseSystemRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupTraverseSystemExecute(r)
}

/*
GraphUserGroupTraverseSystem List the Systems bound to a User Group

This endpoint will return all Systems bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User Group to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this User Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/systems \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupTraverseSystemRequest
*/
func (a *GraphApiService) GraphUserGroupTraverseSystem(ctx context.Context, groupId string) GraphApiGraphUserGroupTraverseSystemRequest {
	return GraphApiGraphUserGroupTraverseSystemRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupTraverseSystemExecute(r GraphApiGraphUserGroupTraverseSystemRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupTraverseSystem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserGroupTraverseSystemGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	groupId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserGroupTraverseSystemGroupRequest) Limit(limit int32) GraphApiGraphUserGroupTraverseSystemGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserGroupTraverseSystemGroupRequest) XOrgId(xOrgId string) GraphApiGraphUserGroupTraverseSystemGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserGroupTraverseSystemGroupRequest) Skip(skip int32) GraphApiGraphUserGroupTraverseSystemGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserGroupTraverseSystemGroupRequest) Filter(filter []string) GraphApiGraphUserGroupTraverseSystemGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserGroupTraverseSystemGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserGroupTraverseSystemGroupExecute(r)
}

/*
GraphUserGroupTraverseSystemGroup List the System Groups bound to User Groups

This endpoint will return all System Groups bound to a User Group, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User Group to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this User Group.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/usergroups/{GroupID}/systemgroups \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId ObjectID of the User Group.
 @return GraphApiGraphUserGroupTraverseSystemGroupRequest
*/
func (a *GraphApiService) GraphUserGroupTraverseSystemGroup(ctx context.Context, groupId string) GraphApiGraphUserGroupTraverseSystemGroupRequest {
	return GraphApiGraphUserGroupTraverseSystemGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserGroupTraverseSystemGroupExecute(r GraphApiGraphUserGroupTraverseSystemGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserGroupTraverseSystemGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usergroups/{group_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type GraphApiGraphUserMemberOfRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	filter *[]string
	limit *int32
	skip *int32
	sort *[]string
	xOrgId *string
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserMemberOfRequest) Filter(filter []string) GraphApiGraphUserMemberOfRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserMemberOfRequest) Limit(limit int32) GraphApiGraphUserMemberOfRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserMemberOfRequest) Skip(skip int32) GraphApiGraphUserMemberOfRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r GraphApiGraphUserMemberOfRequest) Sort(sort []string) GraphApiGraphUserMemberOfRequest {
	r.sort = &sort
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserMemberOfRequest) XOrgId(xOrgId string) GraphApiGraphUserMemberOfRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiGraphUserMemberOfRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserMemberOfExecute(r)
}

/*
GraphUserMemberOf List the parent Groups of a User

This endpoint returns all the User Groups a User is a member of.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/memberof \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserMemberOfRequest
*/
func (a *GraphApiService) GraphUserMemberOf(ctx context.Context, userId string) GraphApiGraphUserMemberOfRequest {
	return GraphApiGraphUserMemberOfRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserMemberOfExecute(r GraphApiGraphUserMemberOfRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserMemberOf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/memberof"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserTraverseActiveDirectoryRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	filter *[]string
	limit *int32
	xOrgId *string
	skip *int32
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserTraverseActiveDirectoryRequest) Filter(filter []string) GraphApiGraphUserTraverseActiveDirectoryRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserTraverseActiveDirectoryRequest) Limit(limit int32) GraphApiGraphUserTraverseActiveDirectoryRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserTraverseActiveDirectoryRequest) XOrgId(xOrgId string) GraphApiGraphUserTraverseActiveDirectoryRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserTraverseActiveDirectoryRequest) Skip(skip int32) GraphApiGraphUserTraverseActiveDirectoryRequest {
	r.skip = &skip
	return r
}

func (r GraphApiGraphUserTraverseActiveDirectoryRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserTraverseActiveDirectoryExecute(r)
}

/*
GraphUserTraverseActiveDirectory List the Active Directory instances bound to a User

This endpoint will return all Active Directory Instances bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User to the corresponding Active Directory instance; this array represents all grouping and/or associations that would have to be removed to deprovision the Active Directory instance from this User.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/activedirectories \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserTraverseActiveDirectoryRequest
*/
func (a *GraphApiService) GraphUserTraverseActiveDirectory(ctx context.Context, userId string) GraphApiGraphUserTraverseActiveDirectoryRequest {
	return GraphApiGraphUserTraverseActiveDirectoryRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserTraverseActiveDirectoryExecute(r GraphApiGraphUserTraverseActiveDirectoryRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserTraverseActiveDirectory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/activedirectories"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserTraverseApplicationRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserTraverseApplicationRequest) Limit(limit int32) GraphApiGraphUserTraverseApplicationRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserTraverseApplicationRequest) XOrgId(xOrgId string) GraphApiGraphUserTraverseApplicationRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserTraverseApplicationRequest) Skip(skip int32) GraphApiGraphUserTraverseApplicationRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserTraverseApplicationRequest) Filter(filter []string) GraphApiGraphUserTraverseApplicationRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserTraverseApplicationRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserTraverseApplicationExecute(r)
}

/*
GraphUserTraverseApplication List the Applications bound to a User

This endpoint will return all Applications bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User to the corresponding Application; this array represents all grouping and/or associations that would have to be removed to deprovision the Application from this User.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/applications \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserTraverseApplicationRequest
*/
func (a *GraphApiService) GraphUserTraverseApplication(ctx context.Context, userId string) GraphApiGraphUserTraverseApplicationRequest {
	return GraphApiGraphUserTraverseApplicationRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserTraverseApplicationExecute(r GraphApiGraphUserTraverseApplicationRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserTraverseApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserTraverseDirectoryRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserTraverseDirectoryRequest) Limit(limit int32) GraphApiGraphUserTraverseDirectoryRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserTraverseDirectoryRequest) XOrgId(xOrgId string) GraphApiGraphUserTraverseDirectoryRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserTraverseDirectoryRequest) Skip(skip int32) GraphApiGraphUserTraverseDirectoryRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserTraverseDirectoryRequest) Filter(filter []string) GraphApiGraphUserTraverseDirectoryRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserTraverseDirectoryRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserTraverseDirectoryExecute(r)
}

/*
GraphUserTraverseDirectory List the Directories bound to a User

This endpoint will return all Directories bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User to the corresponding Directory; this array represents all grouping and/or associations that would have to be removed to deprovision the Directory from this User.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/directories \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserTraverseDirectoryRequest
*/
func (a *GraphApiService) GraphUserTraverseDirectory(ctx context.Context, userId string) GraphApiGraphUserTraverseDirectoryRequest {
	return GraphApiGraphUserTraverseDirectoryRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserTraverseDirectoryExecute(r GraphApiGraphUserTraverseDirectoryRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserTraverseDirectory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/directories"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserTraverseGSuiteRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserTraverseGSuiteRequest) Limit(limit int32) GraphApiGraphUserTraverseGSuiteRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserTraverseGSuiteRequest) XOrgId(xOrgId string) GraphApiGraphUserTraverseGSuiteRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserTraverseGSuiteRequest) Skip(skip int32) GraphApiGraphUserTraverseGSuiteRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserTraverseGSuiteRequest) Filter(filter []string) GraphApiGraphUserTraverseGSuiteRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserTraverseGSuiteRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserTraverseGSuiteExecute(r)
}

/*
GraphUserTraverseGSuite List the G Suite instances bound to a User

This endpoint will return all G-Suite Instances bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User to the corresponding G Suite instance; this array represents all grouping and/or associations that would have to be removed to deprovision the G Suite instance from this User.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/gsuites \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserTraverseGSuiteRequest
*/
func (a *GraphApiService) GraphUserTraverseGSuite(ctx context.Context, userId string) GraphApiGraphUserTraverseGSuiteRequest {
	return GraphApiGraphUserTraverseGSuiteRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserTraverseGSuiteExecute(r GraphApiGraphUserTraverseGSuiteRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserTraverseGSuite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/gsuites"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserTraverseLdapServerRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserTraverseLdapServerRequest) Limit(limit int32) GraphApiGraphUserTraverseLdapServerRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserTraverseLdapServerRequest) XOrgId(xOrgId string) GraphApiGraphUserTraverseLdapServerRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserTraverseLdapServerRequest) Skip(skip int32) GraphApiGraphUserTraverseLdapServerRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserTraverseLdapServerRequest) Filter(filter []string) GraphApiGraphUserTraverseLdapServerRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserTraverseLdapServerRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserTraverseLdapServerExecute(r)
}

/*
GraphUserTraverseLdapServer List the LDAP servers bound to a User

This endpoint will return all LDAP Servers bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User to the corresponding LDAP Server; this array represents all grouping and/or associations that would have to be removed to deprovision the LDAP Server from this User.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/ldapservers \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserTraverseLdapServerRequest
*/
func (a *GraphApiService) GraphUserTraverseLdapServer(ctx context.Context, userId string) GraphApiGraphUserTraverseLdapServerRequest {
	return GraphApiGraphUserTraverseLdapServerRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserTraverseLdapServerExecute(r GraphApiGraphUserTraverseLdapServerRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserTraverseLdapServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/ldapservers"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserTraverseOffice365Request struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserTraverseOffice365Request) Limit(limit int32) GraphApiGraphUserTraverseOffice365Request {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserTraverseOffice365Request) XOrgId(xOrgId string) GraphApiGraphUserTraverseOffice365Request {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserTraverseOffice365Request) Skip(skip int32) GraphApiGraphUserTraverseOffice365Request {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserTraverseOffice365Request) Filter(filter []string) GraphApiGraphUserTraverseOffice365Request {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserTraverseOffice365Request) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserTraverseOffice365Execute(r)
}

/*
GraphUserTraverseOffice365 List the Office 365 instances bound to a User

This endpoint will return all Office 365 Instances bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User to the corresponding Office 365 instance; this array represents all grouping and/or associations that would have to be removed to deprovision the Office 365 instance from this User.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/office365s \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserTraverseOffice365Request
*/
func (a *GraphApiService) GraphUserTraverseOffice365(ctx context.Context, userId string) GraphApiGraphUserTraverseOffice365Request {
	return GraphApiGraphUserTraverseOffice365Request{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserTraverseOffice365Execute(r GraphApiGraphUserTraverseOffice365Request) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserTraverseOffice365")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/office365s"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserTraverseRadiusServerRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserTraverseRadiusServerRequest) Limit(limit int32) GraphApiGraphUserTraverseRadiusServerRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserTraverseRadiusServerRequest) XOrgId(xOrgId string) GraphApiGraphUserTraverseRadiusServerRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserTraverseRadiusServerRequest) Skip(skip int32) GraphApiGraphUserTraverseRadiusServerRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserTraverseRadiusServerRequest) Filter(filter []string) GraphApiGraphUserTraverseRadiusServerRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserTraverseRadiusServerRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserTraverseRadiusServerExecute(r)
}

/*
GraphUserTraverseRadiusServer List the RADIUS Servers bound to a User

This endpoint will return all RADIUS Servers bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User to the corresponding RADIUS Server; this array represents all grouping and/or associations that would have to be removed to deprovision the RADIUS Server from this User.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/radiusservers \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserTraverseRadiusServerRequest
*/
func (a *GraphApiService) GraphUserTraverseRadiusServer(ctx context.Context, userId string) GraphApiGraphUserTraverseRadiusServerRequest {
	return GraphApiGraphUserTraverseRadiusServerRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserTraverseRadiusServerExecute(r GraphApiGraphUserTraverseRadiusServerRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserTraverseRadiusServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/radiusservers"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserTraverseSystemRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserTraverseSystemRequest) Limit(limit int32) GraphApiGraphUserTraverseSystemRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserTraverseSystemRequest) XOrgId(xOrgId string) GraphApiGraphUserTraverseSystemRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserTraverseSystemRequest) Skip(skip int32) GraphApiGraphUserTraverseSystemRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserTraverseSystemRequest) Filter(filter []string) GraphApiGraphUserTraverseSystemRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserTraverseSystemRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserTraverseSystemExecute(r)
}

/*
GraphUserTraverseSystem List the Systems bound to a User

This endpoint will return all Systems bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this User.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/systems\
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserTraverseSystemRequest
*/
func (a *GraphApiService) GraphUserTraverseSystem(ctx context.Context, userId string) GraphApiGraphUserTraverseSystemRequest {
	return GraphApiGraphUserTraverseSystemRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserTraverseSystemExecute(r GraphApiGraphUserTraverseSystemRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserTraverseSystem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiGraphUserTraverseSystemGroupRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	userId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r GraphApiGraphUserTraverseSystemGroupRequest) Limit(limit int32) GraphApiGraphUserTraverseSystemGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiGraphUserTraverseSystemGroupRequest) XOrgId(xOrgId string) GraphApiGraphUserTraverseSystemGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r GraphApiGraphUserTraverseSystemGroupRequest) Skip(skip int32) GraphApiGraphUserTraverseSystemGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiGraphUserTraverseSystemGroupRequest) Filter(filter []string) GraphApiGraphUserTraverseSystemGroupRequest {
	r.filter = &filter
	return r
}

func (r GraphApiGraphUserTraverseSystemGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
	return r.ApiService.GraphUserTraverseSystemGroupExecute(r)
}

/*
GraphUserTraverseSystemGroup List the System Groups bound to a User

This endpoint will return all System Groups bound to a User, either directly or indirectly, essentially traversing the JumpCloud Graph for your Organization.

Each element will contain the type, id, attributes and paths.

The `attributes` object is a key/value hash of compiled graph attributes for all paths followed.

The `paths` array enumerates each path from this User to the corresponding System Group; this array represents all grouping and/or associations that would have to be removed to deprovision the System Group from this User.

See `/members` and `/associations` endpoints to manage those collections.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/users/{UserID}/systemgroups\
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ObjectID of the User.
 @return GraphApiGraphUserTraverseSystemGroupRequest
*/
func (a *GraphApiService) GraphUserTraverseSystemGroup(ctx context.Context, userId string) GraphApiGraphUserTraverseSystemGroupRequest {
	return GraphApiGraphUserTraverseSystemGroupRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *GraphApiService) GraphUserTraverseSystemGroupExecute(r GraphApiGraphUserTraverseSystemGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.GraphUserTraverseSystemGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}/systemgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type GraphApiPolicystatusesSystemsListRequest struct {
	ctx context.Context
	ApiService *GraphApiService
	systemId string
	fields *[]string
	filter *[]string
	limit *int32
	skip *int32
	sort *[]string
	xOrgId *string
}

// The comma separated fields included in the returned records. If omitted, the default list of fields will be returned. 
func (r GraphApiPolicystatusesSystemsListRequest) Fields(fields []string) GraphApiPolicystatusesSystemsListRequest {
	r.fields = &fields
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r GraphApiPolicystatusesSystemsListRequest) Filter(filter []string) GraphApiPolicystatusesSystemsListRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r GraphApiPolicystatusesSystemsListRequest) Limit(limit int32) GraphApiPolicystatusesSystemsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r GraphApiPolicystatusesSystemsListRequest) Skip(skip int32) GraphApiPolicystatusesSystemsListRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r GraphApiPolicystatusesSystemsListRequest) Sort(sort []string) GraphApiPolicystatusesSystemsListRequest {
	r.sort = &sort
	return r
}

// Organization identifier that can be obtained from console settings.
func (r GraphApiPolicystatusesSystemsListRequest) XOrgId(xOrgId string) GraphApiPolicystatusesSystemsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r GraphApiPolicystatusesSystemsListRequest) Execute() ([]PolicyResult, *http.Response, error) {
	return r.ApiService.PolicystatusesSystemsListExecute(r)
}

/*
PolicystatusesSystemsList List the policy statuses for a system

This endpoint returns the policy results for a particular system.

##### Sample Request

```
curl -X GET https://console.jumpcloud.com/api/v2/systems/{System_ID}/policystatuses \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param systemId ObjectID of the System.
 @return GraphApiPolicystatusesSystemsListRequest
*/
func (a *GraphApiService) PolicystatusesSystemsList(ctx context.Context, systemId string) GraphApiPolicystatusesSystemsListRequest {
	return GraphApiPolicystatusesSystemsListRequest{
		ApiService: a,
		ctx: ctx,
		systemId: systemId,
	}
}

// Execute executes the request
//  @return []PolicyResult
func (a *GraphApiService) PolicystatusesSystemsListExecute(r GraphApiPolicystatusesSystemsListRequest) ([]PolicyResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PolicyResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GraphApiService.PolicystatusesSystemsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systems/{system_id}/policystatuses"
	localVarPath = strings.Replace(localVarPath, "{"+"system_id"+"}", url.PathEscape(parameterValueToString(r.systemId, "systemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
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

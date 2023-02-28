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


// Office365ApiService Office365Api service
type Office365ApiService service

type Office365ApiGraphOffice365AssociationsListRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;office_365\&quot; can be associated to.
func (r Office365ApiGraphOffice365AssociationsListRequest) Targets(targets []string) Office365ApiGraphOffice365AssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r Office365ApiGraphOffice365AssociationsListRequest) Limit(limit int32) Office365ApiGraphOffice365AssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r Office365ApiGraphOffice365AssociationsListRequest) Skip(skip int32) Office365ApiGraphOffice365AssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r Office365ApiGraphOffice365AssociationsListRequest) XOrgId(xOrgId string) Office365ApiGraphOffice365AssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r Office365ApiGraphOffice365AssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
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
 @return Office365ApiGraphOffice365AssociationsListRequest
*/
func (a *Office365ApiService) GraphOffice365AssociationsList(ctx context.Context, office365Id string) Office365ApiGraphOffice365AssociationsListRequest {
	return Office365ApiGraphOffice365AssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *Office365ApiService) GraphOffice365AssociationsListExecute(r Office365ApiGraphOffice365AssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.GraphOffice365AssociationsList")
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

type Office365ApiGraphOffice365AssociationsPostRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	xOrgId *string
	body *GraphOperationOffice365
}

// Organization identifier that can be obtained from console settings.
func (r Office365ApiGraphOffice365AssociationsPostRequest) XOrgId(xOrgId string) Office365ApiGraphOffice365AssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r Office365ApiGraphOffice365AssociationsPostRequest) Body(body GraphOperationOffice365) Office365ApiGraphOffice365AssociationsPostRequest {
	r.body = &body
	return r
}

func (r Office365ApiGraphOffice365AssociationsPostRequest) Execute() (*http.Response, error) {
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
 @return Office365ApiGraphOffice365AssociationsPostRequest
*/
func (a *Office365ApiService) GraphOffice365AssociationsPost(ctx context.Context, office365Id string) Office365ApiGraphOffice365AssociationsPostRequest {
	return Office365ApiGraphOffice365AssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
func (a *Office365ApiService) GraphOffice365AssociationsPostExecute(r Office365ApiGraphOffice365AssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.GraphOffice365AssociationsPost")
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

type Office365ApiGraphOffice365TraverseUserRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r Office365ApiGraphOffice365TraverseUserRequest) Limit(limit int32) Office365ApiGraphOffice365TraverseUserRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r Office365ApiGraphOffice365TraverseUserRequest) XOrgId(xOrgId string) Office365ApiGraphOffice365TraverseUserRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r Office365ApiGraphOffice365TraverseUserRequest) Skip(skip int32) Office365ApiGraphOffice365TraverseUserRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r Office365ApiGraphOffice365TraverseUserRequest) Filter(filter []string) Office365ApiGraphOffice365TraverseUserRequest {
	r.filter = &filter
	return r
}

func (r Office365ApiGraphOffice365TraverseUserRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return Office365ApiGraphOffice365TraverseUserRequest
*/
func (a *Office365ApiService) GraphOffice365TraverseUser(ctx context.Context, office365Id string) Office365ApiGraphOffice365TraverseUserRequest {
	return Office365ApiGraphOffice365TraverseUserRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *Office365ApiService) GraphOffice365TraverseUserExecute(r Office365ApiGraphOffice365TraverseUserRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.GraphOffice365TraverseUser")
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

type Office365ApiGraphOffice365TraverseUserGroupRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r Office365ApiGraphOffice365TraverseUserGroupRequest) Limit(limit int32) Office365ApiGraphOffice365TraverseUserGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r Office365ApiGraphOffice365TraverseUserGroupRequest) XOrgId(xOrgId string) Office365ApiGraphOffice365TraverseUserGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r Office365ApiGraphOffice365TraverseUserGroupRequest) Skip(skip int32) Office365ApiGraphOffice365TraverseUserGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r Office365ApiGraphOffice365TraverseUserGroupRequest) Filter(filter []string) Office365ApiGraphOffice365TraverseUserGroupRequest {
	r.filter = &filter
	return r
}

func (r Office365ApiGraphOffice365TraverseUserGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return Office365ApiGraphOffice365TraverseUserGroupRequest
*/
func (a *Office365ApiService) GraphOffice365TraverseUserGroup(ctx context.Context, office365Id string) Office365ApiGraphOffice365TraverseUserGroupRequest {
	return Office365ApiGraphOffice365TraverseUserGroupRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *Office365ApiService) GraphOffice365TraverseUserGroupExecute(r Office365ApiGraphOffice365TraverseUserGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.GraphOffice365TraverseUserGroup")
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

type Office365ApiOffice365sGetRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	xOrgId *string
}

// Organization identifier that can be obtained from console settings.
func (r Office365ApiOffice365sGetRequest) XOrgId(xOrgId string) Office365ApiOffice365sGetRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r Office365ApiOffice365sGetRequest) Execute() (*Office365, *http.Response, error) {
	return r.ApiService.Office365sGetExecute(r)
}

/*
Office365sGet Get Office 365 instance

This endpoint returns a specific Office 365 instance.

#####

Sample Request

```
curl -X GET https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID} \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'

```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id ObjectID of the Office 365 instance.
 @return Office365ApiOffice365sGetRequest
*/
func (a *Office365ApiService) Office365sGet(ctx context.Context, office365Id string) Office365ApiOffice365sGetRequest {
	return Office365ApiOffice365sGetRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return Office365
func (a *Office365ApiService) Office365sGetExecute(r Office365ApiOffice365sGetRequest) (*Office365, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Office365
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.Office365sGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", url.PathEscape(parameterValueToString(r.office365Id, "office365Id")), -1)

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

type Office365ApiOffice365sListImportUsersRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	consistencyLevel *string
	top *int32
	skipToken *string
	filter *string
	search *string
	orderby *string
	count *bool
}

// Defines the consistency header for O365 requests. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#request-headers
func (r Office365ApiOffice365sListImportUsersRequest) ConsistencyLevel(consistencyLevel string) Office365ApiOffice365sListImportUsersRequest {
	r.consistencyLevel = &consistencyLevel
	return r
}

// Office 365 API maximum number of results per page. See https://docs.microsoft.com/en-us/graph/paging.
func (r Office365ApiOffice365sListImportUsersRequest) Top(top int32) Office365ApiOffice365sListImportUsersRequest {
	r.top = &top
	return r
}

// Office 365 API token used to access the next page of results. See https://docs.microsoft.com/en-us/graph/paging.
func (r Office365ApiOffice365sListImportUsersRequest) SkipToken(skipToken string) Office365ApiOffice365sListImportUsersRequest {
	r.skipToken = &skipToken
	return r
}

// Office 365 API filter parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters.
func (r Office365ApiOffice365sListImportUsersRequest) Filter(filter string) Office365ApiOffice365sListImportUsersRequest {
	r.filter = &filter
	return r
}

// Office 365 API search parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters.
func (r Office365ApiOffice365sListImportUsersRequest) Search(search string) Office365ApiOffice365sListImportUsersRequest {
	r.search = &search
	return r
}

// Office 365 API orderby parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters.
func (r Office365ApiOffice365sListImportUsersRequest) Orderby(orderby string) Office365ApiOffice365sListImportUsersRequest {
	r.orderby = &orderby
	return r
}

// Office 365 API count parameter. See https://docs.microsoft.com/en-us/graph/api/user-list?view&#x3D;graph-rest-1.0&amp;tabs&#x3D;http#optional-query-parameters.
func (r Office365ApiOffice365sListImportUsersRequest) Count(count bool) Office365ApiOffice365sListImportUsersRequest {
	r.count = &count
	return r
}

func (r Office365ApiOffice365sListImportUsersRequest) Execute() (*Office365sListImportUsers200Response, *http.Response, error) {
	return r.ApiService.Office365sListImportUsersExecute(r)
}

/*
Office365sListImportUsers Get a list of users to import from an Office 365 instance

Lists Office 365 users available for import.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id
 @return Office365ApiOffice365sListImportUsersRequest
*/
func (a *Office365ApiService) Office365sListImportUsers(ctx context.Context, office365Id string) Office365ApiOffice365sListImportUsersRequest {
	return Office365ApiOffice365sListImportUsersRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return Office365sListImportUsers200Response
func (a *Office365ApiService) Office365sListImportUsersExecute(r Office365ApiOffice365sListImportUsersRequest) (*Office365sListImportUsers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Office365sListImportUsers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.Office365sListImportUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}/import/users"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", url.PathEscape(parameterValueToString(r.office365Id, "office365Id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.top != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "top", r.top, "")
	}
	if r.skipToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipToken", r.skipToken, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.orderby != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderby", r.orderby, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
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
	if r.consistencyLevel != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "ConsistencyLevel", r.consistencyLevel, "")
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

type Office365ApiOffice365sPatchRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	xOrgId *string
	body *Office365
}

// Organization identifier that can be obtained from console settings.
func (r Office365ApiOffice365sPatchRequest) XOrgId(xOrgId string) Office365ApiOffice365sPatchRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r Office365ApiOffice365sPatchRequest) Body(body Office365) Office365ApiOffice365sPatchRequest {
	r.body = &body
	return r
}

func (r Office365ApiOffice365sPatchRequest) Execute() (*Office365, *http.Response, error) {
	return r.ApiService.Office365sPatchExecute(r)
}

/*
Office365sPatch Update existing Office 365 instance.

This endpoint allows updating some attributes of an Office 365 instance.

#####

Sample Request

```
curl -X PATCH https://console.jumpcloud.com/api/v2/office365s/{OFFICE365_ID} \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "userLockoutAction": "maintain",
    "userPasswordExpirationAction": "suspend"
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id ObjectID of the Office 365 instance.
 @return Office365ApiOffice365sPatchRequest
*/
func (a *Office365ApiService) Office365sPatch(ctx context.Context, office365Id string) Office365ApiOffice365sPatchRequest {
	return Office365ApiOffice365sPatchRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return Office365
func (a *Office365ApiService) Office365sPatchExecute(r Office365ApiOffice365sPatchRequest) (*Office365, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Office365
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.Office365sPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}"
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type Office365ApiTranslationRulesOffice365DeleteRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	id string
}

func (r Office365ApiTranslationRulesOffice365DeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.TranslationRulesOffice365DeleteExecute(r)
}

/*
TranslationRulesOffice365Delete Deletes a Office 365 translation rule

This endpoint allows you to delete a translation rule for a specific Office 365 instance. These rules specify how JumpCloud attributes translate to [Microsoft Graph](https://developer.microsoft.com/en-us/graph) attributes.

#### Sample Request

```
curl -X DELETE https://console.jumpcloud.com/api/v2/office365s/{office365_id}/translationrules/{id} \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
  ```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id
 @param id
 @return Office365ApiTranslationRulesOffice365DeleteRequest
*/
func (a *Office365ApiService) TranslationRulesOffice365Delete(ctx context.Context, office365Id string, id string) Office365ApiTranslationRulesOffice365DeleteRequest {
	return Office365ApiTranslationRulesOffice365DeleteRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
		id: id,
	}
}

// Execute executes the request
func (a *Office365ApiService) TranslationRulesOffice365DeleteExecute(r Office365ApiTranslationRulesOffice365DeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.TranslationRulesOffice365Delete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}/translationrules/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", url.PathEscape(parameterValueToString(r.office365Id, "office365Id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type Office365ApiTranslationRulesOffice365GetRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	id string
}

func (r Office365ApiTranslationRulesOffice365GetRequest) Execute() (*Office365TranslationRule, *http.Response, error) {
	return r.ApiService.TranslationRulesOffice365GetExecute(r)
}

/*
TranslationRulesOffice365Get Gets a specific Office 365 translation rule

This endpoint returns a specific translation rule for a specific Office 365 instance. These rules specify how JumpCloud attributes translate to [Microsoft Graph](https://developer.microsoft.com/en-us/graph) attributes.

###### Sample Request

```
  curl -X GET https://console.jumpcloud.com/api/v2/office365s/{office365_id}/translationrules/{id} \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
  ```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id
 @param id
 @return Office365ApiTranslationRulesOffice365GetRequest
*/
func (a *Office365ApiService) TranslationRulesOffice365Get(ctx context.Context, office365Id string, id string) Office365ApiTranslationRulesOffice365GetRequest {
	return Office365ApiTranslationRulesOffice365GetRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
		id: id,
	}
}

// Execute executes the request
//  @return Office365TranslationRule
func (a *Office365ApiService) TranslationRulesOffice365GetExecute(r Office365ApiTranslationRulesOffice365GetRequest) (*Office365TranslationRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Office365TranslationRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.TranslationRulesOffice365Get")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}/translationrules/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", url.PathEscape(parameterValueToString(r.office365Id, "office365Id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type Office365ApiTranslationRulesOffice365ListRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	fields *[]string
	filter *[]string
	limit *int32
	skip *int32
	sort *[]string
}

// The comma separated fields included in the returned records. If omitted, the default list of fields will be returned. 
func (r Office365ApiTranslationRulesOffice365ListRequest) Fields(fields []string) Office365ApiTranslationRulesOffice365ListRequest {
	r.fields = &fields
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r Office365ApiTranslationRulesOffice365ListRequest) Filter(filter []string) Office365ApiTranslationRulesOffice365ListRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r Office365ApiTranslationRulesOffice365ListRequest) Limit(limit int32) Office365ApiTranslationRulesOffice365ListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r Office365ApiTranslationRulesOffice365ListRequest) Skip(skip int32) Office365ApiTranslationRulesOffice365ListRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r Office365ApiTranslationRulesOffice365ListRequest) Sort(sort []string) Office365ApiTranslationRulesOffice365ListRequest {
	r.sort = &sort
	return r
}

func (r Office365ApiTranslationRulesOffice365ListRequest) Execute() ([]Office365TranslationRule, *http.Response, error) {
	return r.ApiService.TranslationRulesOffice365ListExecute(r)
}

/*
TranslationRulesOffice365List List all the Office 365 Translation Rules

This endpoint returns all translation rules for a specific Office 365 instance. These rules specify how JumpCloud attributes translate to [Microsoft Graph](https://developer.microsoft.com/en-us/graph) attributes.

##### Sample Request

```
 curl -X GET  https://console.jumpcloud.com/api/v2/office365s/{office365_id}/translationrules \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}'
  ```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id
 @return Office365ApiTranslationRulesOffice365ListRequest
*/
func (a *Office365ApiService) TranslationRulesOffice365List(ctx context.Context, office365Id string) Office365ApiTranslationRulesOffice365ListRequest {
	return Office365ApiTranslationRulesOffice365ListRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return []Office365TranslationRule
func (a *Office365ApiService) TranslationRulesOffice365ListExecute(r Office365ApiTranslationRulesOffice365ListRequest) ([]Office365TranslationRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Office365TranslationRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.TranslationRulesOffice365List")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}/translationrules"
	localVarPath = strings.Replace(localVarPath, "{"+"office365_id"+"}", url.PathEscape(parameterValueToString(r.office365Id, "office365Id")), -1)

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

type Office365ApiTranslationRulesOffice365PostRequest struct {
	ctx context.Context
	ApiService *Office365ApiService
	office365Id string
	body *Office365TranslationRuleRequest
}

func (r Office365ApiTranslationRulesOffice365PostRequest) Body(body Office365TranslationRuleRequest) Office365ApiTranslationRulesOffice365PostRequest {
	r.body = &body
	return r
}

func (r Office365ApiTranslationRulesOffice365PostRequest) Execute() (*Office365TranslationRule, *http.Response, error) {
	return r.ApiService.TranslationRulesOffice365PostExecute(r)
}

/*
TranslationRulesOffice365Post Create a new Office 365 Translation Rule

This endpoint allows you to create a translation rule for a specific Office 365 instance. These rules specify how JumpCloud attributes translate to [Microsoft Graph](https://developer.microsoft.com/en-us/graph) attributes.

##### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/office365s/{office365_id}/translationrules \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    {Translation Rule Parameters}
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param office365Id
 @return Office365ApiTranslationRulesOffice365PostRequest
*/
func (a *Office365ApiService) TranslationRulesOffice365Post(ctx context.Context, office365Id string) Office365ApiTranslationRulesOffice365PostRequest {
	return Office365ApiTranslationRulesOffice365PostRequest{
		ApiService: a,
		ctx: ctx,
		office365Id: office365Id,
	}
}

// Execute executes the request
//  @return Office365TranslationRule
func (a *Office365ApiService) TranslationRulesOffice365PostExecute(r Office365ApiTranslationRulesOffice365PostRequest) (*Office365TranslationRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Office365TranslationRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Office365ApiService.TranslationRulesOffice365Post")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/office365s/{office365_id}/translationrules"
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

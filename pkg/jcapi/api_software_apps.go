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


// SoftwareAppsApiService SoftwareAppsApi service
type SoftwareAppsApiService service

type SoftwareAppsApiGraphSoftwareappsAssociationsListRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	softwareAppId string
	targets *[]string
	limit *int32
	skip *int32
	xOrgId *string
}

// Targets which a \&quot;software_app\&quot; can be associated to.
func (r SoftwareAppsApiGraphSoftwareappsAssociationsListRequest) Targets(targets []string) SoftwareAppsApiGraphSoftwareappsAssociationsListRequest {
	r.targets = &targets
	return r
}

// The number of records to return at once. Limited to 100.
func (r SoftwareAppsApiGraphSoftwareappsAssociationsListRequest) Limit(limit int32) SoftwareAppsApiGraphSoftwareappsAssociationsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r SoftwareAppsApiGraphSoftwareappsAssociationsListRequest) Skip(skip int32) SoftwareAppsApiGraphSoftwareappsAssociationsListRequest {
	r.skip = &skip
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiGraphSoftwareappsAssociationsListRequest) XOrgId(xOrgId string) SoftwareAppsApiGraphSoftwareappsAssociationsListRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SoftwareAppsApiGraphSoftwareappsAssociationsListRequest) Execute() ([]GraphConnection, *http.Response, error) {
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
 @return SoftwareAppsApiGraphSoftwareappsAssociationsListRequest
*/
func (a *SoftwareAppsApiService) GraphSoftwareappsAssociationsList(ctx context.Context, softwareAppId string) SoftwareAppsApiGraphSoftwareappsAssociationsListRequest {
	return SoftwareAppsApiGraphSoftwareappsAssociationsListRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
//  @return []GraphConnection
func (a *SoftwareAppsApiService) GraphSoftwareappsAssociationsListExecute(r SoftwareAppsApiGraphSoftwareappsAssociationsListRequest) ([]GraphConnection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.GraphSoftwareappsAssociationsList")
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

type SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	softwareAppId string
	xOrgId *string
	body *GraphOperationSoftwareApp
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest) XOrgId(xOrgId string) SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest) Body(body GraphOperationSoftwareApp) SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest {
	r.body = &body
	return r
}

func (r SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest) Execute() (*http.Response, error) {
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
 @return SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest
*/
func (a *SoftwareAppsApiService) GraphSoftwareappsAssociationsPost(ctx context.Context, softwareAppId string) SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest {
	return SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
func (a *SoftwareAppsApiService) GraphSoftwareappsAssociationsPostExecute(r SoftwareAppsApiGraphSoftwareappsAssociationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.GraphSoftwareappsAssociationsPost")
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

type SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	softwareAppId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest) Limit(limit int32) SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest) XOrgId(xOrgId string) SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest) Skip(skip int32) SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest) Filter(filter []string) SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest {
	r.filter = &filter
	return r
}

func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest
*/
func (a *SoftwareAppsApiService) GraphSoftwareappsTraverseSystem(ctx context.Context, softwareAppId string) SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest {
	return SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *SoftwareAppsApiService) GraphSoftwareappsTraverseSystemExecute(r SoftwareAppsApiGraphSoftwareappsTraverseSystemRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.GraphSoftwareappsTraverseSystem")
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

type SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	softwareAppId string
	limit *int32
	xOrgId *string
	skip *int32
	filter *[]string
}

// The number of records to return at once. Limited to 100.
func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest) Limit(limit int32) SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest {
	r.limit = &limit
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest) XOrgId(xOrgId string) SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest) Skip(skip int32) SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest {
	r.skip = &skip
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest) Filter(filter []string) SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest {
	r.filter = &filter
	return r
}

func (r SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest) Execute() ([]GraphObjectWithPaths, *http.Response, error) {
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
 @return SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest
*/
func (a *SoftwareAppsApiService) GraphSoftwareappsTraverseSystemGroup(ctx context.Context, softwareAppId string) SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest {
	return SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
//  @return []GraphObjectWithPaths
func (a *SoftwareAppsApiService) GraphSoftwareappsTraverseSystemGroupExecute(r SoftwareAppsApiGraphSoftwareappsTraverseSystemGroupRequest) ([]GraphObjectWithPaths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GraphObjectWithPaths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.GraphSoftwareappsTraverseSystemGroup")
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

type SoftwareAppsApiSoftwareAppStatusesListRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	softwareAppId string
	xOrgId *string
	filter *[]string
	limit *int32
	skip *int32
	sort *[]string
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiSoftwareAppStatusesListRequest) XOrgId(xOrgId string) SoftwareAppsApiSoftwareAppStatusesListRequest {
	r.xOrgId = &xOrgId
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SoftwareAppsApiSoftwareAppStatusesListRequest) Filter(filter []string) SoftwareAppsApiSoftwareAppStatusesListRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r SoftwareAppsApiSoftwareAppStatusesListRequest) Limit(limit int32) SoftwareAppsApiSoftwareAppStatusesListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r SoftwareAppsApiSoftwareAppStatusesListRequest) Skip(skip int32) SoftwareAppsApiSoftwareAppStatusesListRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r SoftwareAppsApiSoftwareAppStatusesListRequest) Sort(sort []string) SoftwareAppsApiSoftwareAppStatusesListRequest {
	r.sort = &sort
	return r
}

func (r SoftwareAppsApiSoftwareAppStatusesListRequest) Execute() ([]SoftwareAppStatus, *http.Response, error) {
	return r.ApiService.SoftwareAppStatusesListExecute(r)
}

/*
SoftwareAppStatusesList Get the status of the provided Software Application

This endpoint allows you to get the status of the provided Software Application on associated JumpCloud systems.

#### Sample Request
```
$ curl -X GET https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/statuses \
-H 'Accept: application/json' \
-H 'Content-Type: application/json' \
-H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softwareAppId ObjectID of the Software App.
 @return SoftwareAppsApiSoftwareAppStatusesListRequest
*/
func (a *SoftwareAppsApiService) SoftwareAppStatusesList(ctx context.Context, softwareAppId string) SoftwareAppsApiSoftwareAppStatusesListRequest {
	return SoftwareAppsApiSoftwareAppStatusesListRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
//  @return []SoftwareAppStatus
func (a *SoftwareAppsApiService) SoftwareAppStatusesListExecute(r SoftwareAppsApiSoftwareAppStatusesListRequest) ([]SoftwareAppStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SoftwareAppStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.SoftwareAppStatusesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{software_app_id}/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", url.PathEscape(parameterValueToString(r.softwareAppId, "softwareAppId")), -1)

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

type SoftwareAppsApiSoftwareAppsDeleteRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	id string
	xOrgId *string
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiSoftwareAppsDeleteRequest) XOrgId(xOrgId string) SoftwareAppsApiSoftwareAppsDeleteRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SoftwareAppsApiSoftwareAppsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.SoftwareAppsDeleteExecute(r)
}

/*
SoftwareAppsDelete Delete a configured Software Application

Removes a Software Application configuration.

Warning: This is a destructive operation and will unmanage the application on all affected systems.

#### Sample Request
```
curl -X DELETE https://console.jumpcloud.com/api/v2/softwareapps/{id} \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return SoftwareAppsApiSoftwareAppsDeleteRequest
*/
func (a *SoftwareAppsApiService) SoftwareAppsDelete(ctx context.Context, id string) SoftwareAppsApiSoftwareAppsDeleteRequest {
	return SoftwareAppsApiSoftwareAppsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *SoftwareAppsApiService) SoftwareAppsDeleteExecute(r SoftwareAppsApiSoftwareAppsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.SoftwareAppsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{id}"
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SoftwareAppsApiSoftwareAppsGetRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	id string
	xOrgId *string
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiSoftwareAppsGetRequest) XOrgId(xOrgId string) SoftwareAppsApiSoftwareAppsGetRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SoftwareAppsApiSoftwareAppsGetRequest) Execute() (*SoftwareApp, *http.Response, error) {
	return r.ApiService.SoftwareAppsGetExecute(r)
}

/*
SoftwareAppsGet Retrieve a configured Software Application.

Retrieves a Software Application.
The optional isConfigEnabled and appConfiguration apple_vpp attributes are populated in this response.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/softwareapps/{id} \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return SoftwareAppsApiSoftwareAppsGetRequest
*/
func (a *SoftwareAppsApiService) SoftwareAppsGet(ctx context.Context, id string) SoftwareAppsApiSoftwareAppsGetRequest {
	return SoftwareAppsApiSoftwareAppsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SoftwareApp
func (a *SoftwareAppsApiService) SoftwareAppsGetExecute(r SoftwareAppsApiSoftwareAppsGetRequest) (*SoftwareApp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SoftwareApp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.SoftwareAppsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{id}"
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

type SoftwareAppsApiSoftwareAppsListRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	xOrgId *string
	filter *[]string
	limit *int32
	skip *int32
	sort *[]string
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiSoftwareAppsListRequest) XOrgId(xOrgId string) SoftwareAppsApiSoftwareAppsListRequest {
	r.xOrgId = &xOrgId
	return r
}

// A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60;
func (r SoftwareAppsApiSoftwareAppsListRequest) Filter(filter []string) SoftwareAppsApiSoftwareAppsListRequest {
	r.filter = &filter
	return r
}

// The number of records to return at once. Limited to 100.
func (r SoftwareAppsApiSoftwareAppsListRequest) Limit(limit int32) SoftwareAppsApiSoftwareAppsListRequest {
	r.limit = &limit
	return r
}

// The offset into the records to return.
func (r SoftwareAppsApiSoftwareAppsListRequest) Skip(skip int32) SoftwareAppsApiSoftwareAppsListRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
func (r SoftwareAppsApiSoftwareAppsListRequest) Sort(sort []string) SoftwareAppsApiSoftwareAppsListRequest {
	r.sort = &sort
	return r
}

func (r SoftwareAppsApiSoftwareAppsListRequest) Execute() ([]SoftwareApp, *http.Response, error) {
	return r.ApiService.SoftwareAppsListExecute(r)
}

/*
SoftwareAppsList Get all configured Software Applications.

This endpoint allows you to get all configured Software Applications that will be managed by JumpCloud on associated JumpCloud systems.
The optional isConfigEnabled and appConfiguration apple_vpp attributes are not included in the response.

#### Sample Request
```
$ curl -X GET https://console.jumpcloud.com/api/v2/softwareapps \
-H 'Accept: application/json' \
-H 'Content-Type: application/json' \
-H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SoftwareAppsApiSoftwareAppsListRequest
*/
func (a *SoftwareAppsApiService) SoftwareAppsList(ctx context.Context) SoftwareAppsApiSoftwareAppsListRequest {
	return SoftwareAppsApiSoftwareAppsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SoftwareApp
func (a *SoftwareAppsApiService) SoftwareAppsListExecute(r SoftwareAppsApiSoftwareAppsListRequest) ([]SoftwareApp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SoftwareApp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.SoftwareAppsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps"

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

type SoftwareAppsApiSoftwareAppsPostRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	xOrgId *string
	body *SoftwareApp
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiSoftwareAppsPostRequest) XOrgId(xOrgId string) SoftwareAppsApiSoftwareAppsPostRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SoftwareAppsApiSoftwareAppsPostRequest) Body(body SoftwareApp) SoftwareAppsApiSoftwareAppsPostRequest {
	r.body = &body
	return r
}

func (r SoftwareAppsApiSoftwareAppsPostRequest) Execute() (*SoftwareApp, *http.Response, error) {
	return r.ApiService.SoftwareAppsPostExecute(r)
}

/*
SoftwareAppsPost Create a Software Application that will be managed by JumpCloud.

This endpoint allows you to create a Software Application that will be managed by JumpCloud on associated JumpCloud systems.
The optional isConfigEnabled and appConfiguration apple_vpp attributes are not included in the response.

#### Sample Request
```
$ curl -X POST https://console.jumpcloud.com/api/v2/softwareapps \
-H 'Accept: application/json' \
-H 'Content-Type: application/json' \
-H 'x-api-key: {API_KEY}' \
-d '{
  "displayName": "Adobe Reader",
  "settings": [{"packageId": "adobereader"}]
}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SoftwareAppsApiSoftwareAppsPostRequest
*/
func (a *SoftwareAppsApiService) SoftwareAppsPost(ctx context.Context) SoftwareAppsApiSoftwareAppsPostRequest {
	return SoftwareAppsApiSoftwareAppsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SoftwareApp
func (a *SoftwareAppsApiService) SoftwareAppsPostExecute(r SoftwareAppsApiSoftwareAppsPostRequest) (*SoftwareApp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SoftwareApp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.SoftwareAppsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps"

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

type SoftwareAppsApiSoftwareAppsReclaimLicensesRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	softwareAppId string
}

func (r SoftwareAppsApiSoftwareAppsReclaimLicensesRequest) Execute() (*SoftwareAppReclaimLicenses, *http.Response, error) {
	return r.ApiService.SoftwareAppsReclaimLicensesExecute(r)
}

/*
SoftwareAppsReclaimLicenses Reclaim Licenses for a Software Application.

This endpoint allows you to reclaim the licenses from a software app associated with devices that are deleted.
#### Sample Request
```
$ curl -X POST https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/reclaim-licenses \
-H 'Accept: application/json' \
-H 'Content-Type: application/json' \
-H 'x-api-key: {API_KEY}' \
-d '{}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softwareAppId
 @return SoftwareAppsApiSoftwareAppsReclaimLicensesRequest
*/
func (a *SoftwareAppsApiService) SoftwareAppsReclaimLicenses(ctx context.Context, softwareAppId string) SoftwareAppsApiSoftwareAppsReclaimLicensesRequest {
	return SoftwareAppsApiSoftwareAppsReclaimLicensesRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
//  @return SoftwareAppReclaimLicenses
func (a *SoftwareAppsApiService) SoftwareAppsReclaimLicensesExecute(r SoftwareAppsApiSoftwareAppsReclaimLicensesRequest) (*SoftwareAppReclaimLicenses, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SoftwareAppReclaimLicenses
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.SoftwareAppsReclaimLicenses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{software_app_id}/reclaim-licenses"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", url.PathEscape(parameterValueToString(r.softwareAppId, "softwareAppId")), -1)

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

type SoftwareAppsApiSoftwareAppsRetryInstallationRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	softwareAppId string
	body *SoftwareAppsRetryInstallationRequest
}

func (r SoftwareAppsApiSoftwareAppsRetryInstallationRequest) Body(body SoftwareAppsRetryInstallationRequest) SoftwareAppsApiSoftwareAppsRetryInstallationRequest {
	r.body = &body
	return r
}

func (r SoftwareAppsApiSoftwareAppsRetryInstallationRequest) Execute() (*http.Response, error) {
	return r.ApiService.SoftwareAppsRetryInstallationExecute(r)
}

/*
SoftwareAppsRetryInstallation Retry Installation for a Software Application

This endpoints initiates an installation retry of an Apple VPP App for the provided system IDs
#### Sample Request
```
$ curl -X POST https://console.jumpcloud.com/api/v2/softwareapps/{software_app_id}/retry-installation \
-H 'Accept: application/json' \
-H 'Content-Type: application/json' \
-H 'x-api-key: {API_KEY}' \
-d '{"system_ids": "{<system_id_1>, <system_id_2>, ...}"}'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param softwareAppId
 @return SoftwareAppsApiSoftwareAppsRetryInstallationRequest
*/
func (a *SoftwareAppsApiService) SoftwareAppsRetryInstallation(ctx context.Context, softwareAppId string) SoftwareAppsApiSoftwareAppsRetryInstallationRequest {
	return SoftwareAppsApiSoftwareAppsRetryInstallationRequest{
		ApiService: a,
		ctx: ctx,
		softwareAppId: softwareAppId,
	}
}

// Execute executes the request
func (a *SoftwareAppsApiService) SoftwareAppsRetryInstallationExecute(r SoftwareAppsApiSoftwareAppsRetryInstallationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.SoftwareAppsRetryInstallation")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{software_app_id}/retry-installation"
	localVarPath = strings.Replace(localVarPath, "{"+"software_app_id"+"}", url.PathEscape(parameterValueToString(r.softwareAppId, "softwareAppId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type SoftwareAppsApiSoftwareAppsUpdateRequest struct {
	ctx context.Context
	ApiService *SoftwareAppsApiService
	id string
	xOrgId *string
	body *SoftwareApp
}

// Organization identifier that can be obtained from console settings.
func (r SoftwareAppsApiSoftwareAppsUpdateRequest) XOrgId(xOrgId string) SoftwareAppsApiSoftwareAppsUpdateRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SoftwareAppsApiSoftwareAppsUpdateRequest) Body(body SoftwareApp) SoftwareAppsApiSoftwareAppsUpdateRequest {
	r.body = &body
	return r
}

func (r SoftwareAppsApiSoftwareAppsUpdateRequest) Execute() (*SoftwareApp, *http.Response, error) {
	return r.ApiService.SoftwareAppsUpdateExecute(r)
}

/*
SoftwareAppsUpdate Update a Software Application Configuration.

This endpoint updates a specific Software Application configuration for the organization.
displayName can be changed alone if no settings are provided.
If a setting is provided, it should include all its information since this endpoint will update all the settings' fields.
The optional isConfigEnabled and appConfiguration apple_vpp attributes are not included in the response.

#### Sample Request - displayName only
```
 curl -X PUT https://console.jumpcloud.com/api/v2/softwareapps/{id} \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "displayName": "My Software App"
  }'
```

#### Sample Request - all attributes
```
 curl -X PUT https://console.jumpcloud.com/api/v2/softwareapps/{id} \
  -H 'accept: application/json' \
  -H 'content-type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{
    "displayName": "My Software App",
    "settings": [
      {
        "packageId": "123456",
        "autoUpdate": false,
        "allowUpdateDelay": false,
        "packageManager": "APPLE_VPP",
        "locationObjectId": "123456789012123456789012",
        "location": "123456",
        "desiredState": "Install",
        "appleVpp": {
          "appConfiguration": "<?xml version=\"1.0\" encoding=\"UTF-8\"?><!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\"><plist version=\"1.0\"><dict><key>MyKey</key><string>My String</string></dict></plist>",
          "assignedLicenses": 20,
          "availableLicenses": 10,
          "details": {},
          "isConfigEnabled": true,
          "supportedDeviceFamilies": [
            "IPAD",
            "MAC"
          ],
          "totalLicenses": 30
        },
        "packageSubtitle": "My package subtitle",
        "packageVersion": "1.2.3",
        "packageKind": "software-package",
        "assetKind": "software",
        "assetSha256Size": 256,
        "assetSha256Strings": [
          "a123b123c123d123"
        ],
        "description": "My app description"
      }
    ]
  }'
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return SoftwareAppsApiSoftwareAppsUpdateRequest
*/
func (a *SoftwareAppsApiService) SoftwareAppsUpdate(ctx context.Context, id string) SoftwareAppsApiSoftwareAppsUpdateRequest {
	return SoftwareAppsApiSoftwareAppsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SoftwareApp
func (a *SoftwareAppsApiService) SoftwareAppsUpdateExecute(r SoftwareAppsApiSoftwareAppsUpdateRequest) (*SoftwareApp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SoftwareApp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SoftwareAppsApiService.SoftwareAppsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/softwareapps/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

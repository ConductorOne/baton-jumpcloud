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
)


// SystemInsightsApiService SystemInsightsApi service
type SystemInsightsApiService service

type SystemInsightsApiSysteminsightsListAlfRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	filter *[]string
	skip *int32
	sort *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListAlfRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListAlfRequest {
	r.xOrgId = &xOrgId
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListAlfRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListAlfRequest {
	r.filter = &filter
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListAlfRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListAlfRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListAlfRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListAlfRequest {
	r.sort = &sort
	return r
}

func (r SystemInsightsApiSysteminsightsListAlfRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListAlfRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListAlfRequest) Execute() ([]SystemInsightsAlf, *http.Response, error) {
	return r.ApiService.SysteminsightsListAlfExecute(r)
}

/*
SysteminsightsListAlf List System Insights ALF

Valid filter fields are `system_id` and `global_state`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListAlfRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAlf(ctx context.Context) SystemInsightsApiSysteminsightsListAlfRequest {
	return SystemInsightsApiSysteminsightsListAlfRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAlf
func (a *SystemInsightsApiService) SysteminsightsListAlfExecute(r SystemInsightsApiSysteminsightsListAlfRequest) ([]SystemInsightsAlf, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsAlf
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListAlf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/alf"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListAlfExceptionsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	filter *[]string
	skip *int32
	sort *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListAlfExceptionsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListAlfExceptionsRequest {
	r.xOrgId = &xOrgId
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListAlfExceptionsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListAlfExceptionsRequest {
	r.filter = &filter
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListAlfExceptionsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListAlfExceptionsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListAlfExceptionsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListAlfExceptionsRequest {
	r.sort = &sort
	return r
}

func (r SystemInsightsApiSysteminsightsListAlfExceptionsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListAlfExceptionsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListAlfExceptionsRequest) Execute() ([]SystemInsightsAlfExceptions, *http.Response, error) {
	return r.ApiService.SysteminsightsListAlfExceptionsExecute(r)
}

/*
SysteminsightsListAlfExceptions List System Insights ALF Exceptions

Valid filter fields are `system_id` and `state`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListAlfExceptionsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAlfExceptions(ctx context.Context) SystemInsightsApiSysteminsightsListAlfExceptionsRequest {
	return SystemInsightsApiSysteminsightsListAlfExceptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAlfExceptions
func (a *SystemInsightsApiService) SysteminsightsListAlfExceptionsExecute(r SystemInsightsApiSysteminsightsListAlfExceptionsRequest) ([]SystemInsightsAlfExceptions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsAlfExceptions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListAlfExceptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/alf_exceptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	filter *[]string
	skip *int32
	sort *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest {
	r.xOrgId = &xOrgId
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest {
	r.filter = &filter
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest {
	r.sort = &sort
	return r
}

func (r SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest) Execute() ([]SystemInsightsAlfExplicitAuths, *http.Response, error) {
	return r.ApiService.SysteminsightsListAlfExplicitAuthsExecute(r)
}

/*
SysteminsightsListAlfExplicitAuths List System Insights ALF Explicit Authentications

Valid filter fields are `system_id` and `process`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAlfExplicitAuths(ctx context.Context) SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest {
	return SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAlfExplicitAuths
func (a *SystemInsightsApiService) SysteminsightsListAlfExplicitAuthsExecute(r SystemInsightsApiSysteminsightsListAlfExplicitAuthsRequest) ([]SystemInsightsAlfExplicitAuths, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsAlfExplicitAuths
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListAlfExplicitAuths")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/alf_explicit_auths"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListAppcompatShimsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListAppcompatShimsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListAppcompatShimsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListAppcompatShimsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListAppcompatShimsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListAppcompatShimsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListAppcompatShimsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListAppcompatShimsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListAppcompatShimsRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListAppcompatShimsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListAppcompatShimsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListAppcompatShimsRequest) Execute() ([]SystemInsightsAppcompatShims, *http.Response, error) {
	return r.ApiService.SysteminsightsListAppcompatShimsExecute(r)
}

/*
SysteminsightsListAppcompatShims List System Insights Application Compatibility Shims

Valid filter fields are `system_id` and `enabled`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListAppcompatShimsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAppcompatShims(ctx context.Context) SystemInsightsApiSysteminsightsListAppcompatShimsRequest {
	return SystemInsightsApiSysteminsightsListAppcompatShimsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAppcompatShims
func (a *SystemInsightsApiService) SysteminsightsListAppcompatShimsExecute(r SystemInsightsApiSysteminsightsListAppcompatShimsRequest) ([]SystemInsightsAppcompatShims, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsAppcompatShims
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListAppcompatShims")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/appcompat_shims"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListAppsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListAppsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListAppsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListAppsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListAppsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListAppsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListAppsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListAppsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListAppsRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListAppsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListAppsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListAppsRequest) Execute() ([]SystemInsightsApps, *http.Response, error) {
	return r.ApiService.SysteminsightsListAppsExecute(r)
}

/*
SysteminsightsListApps List System Insights Apps

Lists all apps for macOS devices. For Windows devices, use [List System Insights Programs](#operation/systeminsights_list_programs).

Valid filter fields are `system_id` and `bundle_name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListAppsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListApps(ctx context.Context) SystemInsightsApiSysteminsightsListAppsRequest {
	return SystemInsightsApiSysteminsightsListAppsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsApps
func (a *SystemInsightsApiService) SysteminsightsListAppsExecute(r SystemInsightsApiSysteminsightsListAppsRequest) ([]SystemInsightsApps, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsApps
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListApps")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/apps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListAuthorizedKeysRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListAuthorizedKeysRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListAuthorizedKeysRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListAuthorizedKeysRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListAuthorizedKeysRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListAuthorizedKeysRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListAuthorizedKeysRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListAuthorizedKeysRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListAuthorizedKeysRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListAuthorizedKeysRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListAuthorizedKeysRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListAuthorizedKeysRequest) Execute() ([]SystemInsightsAuthorizedKeys, *http.Response, error) {
	return r.ApiService.SysteminsightsListAuthorizedKeysExecute(r)
}

/*
SysteminsightsListAuthorizedKeys List System Insights Authorized Keys

Valid filter fields are `system_id` and `uid`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListAuthorizedKeysRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAuthorizedKeys(ctx context.Context) SystemInsightsApiSysteminsightsListAuthorizedKeysRequest {
	return SystemInsightsApiSysteminsightsListAuthorizedKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAuthorizedKeys
func (a *SystemInsightsApiService) SysteminsightsListAuthorizedKeysExecute(r SystemInsightsApiSysteminsightsListAuthorizedKeysRequest) ([]SystemInsightsAuthorizedKeys, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsAuthorizedKeys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListAuthorizedKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/authorized_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest) Execute() ([]SystemInsightsAzureInstanceMetadata, *http.Response, error) {
	return r.ApiService.SysteminsightsListAzureInstanceMetadataExecute(r)
}

/*
SysteminsightsListAzureInstanceMetadata List System Insights Azure Instance Metadata

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAzureInstanceMetadata(ctx context.Context) SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest {
	return SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAzureInstanceMetadata
func (a *SystemInsightsApiService) SysteminsightsListAzureInstanceMetadataExecute(r SystemInsightsApiSysteminsightsListAzureInstanceMetadataRequest) ([]SystemInsightsAzureInstanceMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsAzureInstanceMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListAzureInstanceMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/azure_instance_metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
			var v string
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

type SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest) Execute() ([]SystemInsightsAzureInstanceTags, *http.Response, error) {
	return r.ApiService.SysteminsightsListAzureInstanceTagsExecute(r)
}

/*
SysteminsightsListAzureInstanceTags List System Insights Azure Instance Tags

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAzureInstanceTags(ctx context.Context) SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest {
	return SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAzureInstanceTags
func (a *SystemInsightsApiService) SysteminsightsListAzureInstanceTagsExecute(r SystemInsightsApiSysteminsightsListAzureInstanceTagsRequest) ([]SystemInsightsAzureInstanceTags, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsAzureInstanceTags
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListAzureInstanceTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/azure_instance_tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
			var v string
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

type SystemInsightsApiSysteminsightsListBatteryRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListBatteryRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListBatteryRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListBatteryRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListBatteryRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListBatteryRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListBatteryRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListBatteryRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListBatteryRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListBatteryRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListBatteryRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListBatteryRequest) Execute() ([]SystemInsightsBattery, *http.Response, error) {
	return r.ApiService.SysteminsightsListBatteryExecute(r)
}

/*
SysteminsightsListBattery List System Insights Battery

Valid filter fields are `system_id` and `health`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListBatteryRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListBattery(ctx context.Context) SystemInsightsApiSysteminsightsListBatteryRequest {
	return SystemInsightsApiSysteminsightsListBatteryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsBattery
func (a *SystemInsightsApiService) SysteminsightsListBatteryExecute(r SystemInsightsApiSysteminsightsListBatteryRequest) ([]SystemInsightsBattery, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsBattery
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListBattery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/battery"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListBitlockerInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListBitlockerInfoRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListBitlockerInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListBitlockerInfoRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListBitlockerInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListBitlockerInfoRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListBitlockerInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListBitlockerInfoRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListBitlockerInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListBitlockerInfoRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListBitlockerInfoRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListBitlockerInfoRequest) Execute() ([]SystemInsightsBitlockerInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListBitlockerInfoExecute(r)
}

/*
SysteminsightsListBitlockerInfo List System Insights Bitlocker Info

Valid filter fields are `system_id` and `protection_status`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListBitlockerInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListBitlockerInfo(ctx context.Context) SystemInsightsApiSysteminsightsListBitlockerInfoRequest {
	return SystemInsightsApiSysteminsightsListBitlockerInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsBitlockerInfo
func (a *SystemInsightsApiService) SysteminsightsListBitlockerInfoExecute(r SystemInsightsApiSysteminsightsListBitlockerInfoRequest) ([]SystemInsightsBitlockerInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsBitlockerInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListBitlockerInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/bitlocker_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListBrowserPluginsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListBrowserPluginsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListBrowserPluginsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListBrowserPluginsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListBrowserPluginsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListBrowserPluginsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListBrowserPluginsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListBrowserPluginsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListBrowserPluginsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListBrowserPluginsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListBrowserPluginsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListBrowserPluginsRequest) Execute() ([]SystemInsightsBrowserPlugins, *http.Response, error) {
	return r.ApiService.SysteminsightsListBrowserPluginsExecute(r)
}

/*
SysteminsightsListBrowserPlugins List System Insights Browser Plugins

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListBrowserPluginsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListBrowserPlugins(ctx context.Context) SystemInsightsApiSysteminsightsListBrowserPluginsRequest {
	return SystemInsightsApiSysteminsightsListBrowserPluginsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsBrowserPlugins
func (a *SystemInsightsApiService) SysteminsightsListBrowserPluginsExecute(r SystemInsightsApiSysteminsightsListBrowserPluginsRequest) ([]SystemInsightsBrowserPlugins, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsBrowserPlugins
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListBrowserPlugins")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/browser_plugins"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListCertificatesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListCertificatesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListCertificatesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListCertificatesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListCertificatesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; Note: You can only filter by &#x60;system_id&#x60; and &#x60;common_name&#x60; 
func (r SystemInsightsApiSysteminsightsListCertificatesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListCertificatesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListCertificatesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListCertificatesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListCertificatesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListCertificatesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListCertificatesRequest) Execute() ([]SystemInsightsCertificates, *http.Response, error) {
	return r.ApiService.SysteminsightsListCertificatesExecute(r)
}

/*
SysteminsightsListCertificates List System Insights Certificates

Valid filter fields are `system_id` and `common_name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListCertificatesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListCertificates(ctx context.Context) SystemInsightsApiSysteminsightsListCertificatesRequest {
	return SystemInsightsApiSysteminsightsListCertificatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsCertificates
func (a *SystemInsightsApiService) SysteminsightsListCertificatesExecute(r SystemInsightsApiSysteminsightsListCertificatesRequest) ([]SystemInsightsCertificates, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsCertificates
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListCertificates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/certificates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListChassisInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListChassisInfoRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListChassisInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListChassisInfoRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListChassisInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListChassisInfoRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListChassisInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListChassisInfoRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListChassisInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListChassisInfoRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListChassisInfoRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListChassisInfoRequest) Execute() ([]SystemInsightsChassisInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListChassisInfoExecute(r)
}

/*
SysteminsightsListChassisInfo List System Insights Chassis Info

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListChassisInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListChassisInfo(ctx context.Context) SystemInsightsApiSysteminsightsListChassisInfoRequest {
	return SystemInsightsApiSysteminsightsListChassisInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsChassisInfo
func (a *SystemInsightsApiService) SysteminsightsListChassisInfoExecute(r SystemInsightsApiSysteminsightsListChassisInfoRequest) ([]SystemInsightsChassisInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsChassisInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListChassisInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/chassis_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
			var v string
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

type SystemInsightsApiSysteminsightsListChromeExtensionsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListChromeExtensionsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListChromeExtensionsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListChromeExtensionsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListChromeExtensionsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListChromeExtensionsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListChromeExtensionsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListChromeExtensionsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListChromeExtensionsRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListChromeExtensionsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListChromeExtensionsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListChromeExtensionsRequest) Execute() ([]SystemInsightsChromeExtensions, *http.Response, error) {
	return r.ApiService.SysteminsightsListChromeExtensionsExecute(r)
}

/*
SysteminsightsListChromeExtensions List System Insights Chrome Extensions

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListChromeExtensionsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListChromeExtensions(ctx context.Context) SystemInsightsApiSysteminsightsListChromeExtensionsRequest {
	return SystemInsightsApiSysteminsightsListChromeExtensionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsChromeExtensions
func (a *SystemInsightsApiService) SysteminsightsListChromeExtensionsExecute(r SystemInsightsApiSysteminsightsListChromeExtensionsRequest) ([]SystemInsightsChromeExtensions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsChromeExtensions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListChromeExtensions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/chrome_extensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListConnectivityRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListConnectivityRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListConnectivityRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListConnectivityRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListConnectivityRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListConnectivityRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListConnectivityRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListConnectivityRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListConnectivityRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListConnectivityRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListConnectivityRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListConnectivityRequest) Execute() ([]SystemInsightsConnectivity, *http.Response, error) {
	return r.ApiService.SysteminsightsListConnectivityExecute(r)
}

/*
SysteminsightsListConnectivity List System Insights Connectivity

The only valid filter field is `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListConnectivityRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListConnectivity(ctx context.Context) SystemInsightsApiSysteminsightsListConnectivityRequest {
	return SystemInsightsApiSysteminsightsListConnectivityRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsConnectivity
func (a *SystemInsightsApiService) SysteminsightsListConnectivityExecute(r SystemInsightsApiSysteminsightsListConnectivityRequest) ([]SystemInsightsConnectivity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsConnectivity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListConnectivity")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/connectivity"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListCrashesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListCrashesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListCrashesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListCrashesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListCrashesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListCrashesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListCrashesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListCrashesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListCrashesRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListCrashesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListCrashesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListCrashesRequest) Execute() ([]SystemInsightsCrashes, *http.Response, error) {
	return r.ApiService.SysteminsightsListCrashesExecute(r)
}

/*
SysteminsightsListCrashes List System Insights Crashes

Valid filter fields are `system_id` and `identifier`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListCrashesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListCrashes(ctx context.Context) SystemInsightsApiSysteminsightsListCrashesRequest {
	return SystemInsightsApiSysteminsightsListCrashesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsCrashes
func (a *SystemInsightsApiService) SysteminsightsListCrashesExecute(r SystemInsightsApiSysteminsightsListCrashesRequest) ([]SystemInsightsCrashes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsCrashes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListCrashes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/crashes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListCupsDestinationsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListCupsDestinationsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListCupsDestinationsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListCupsDestinationsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListCupsDestinationsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListCupsDestinationsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListCupsDestinationsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListCupsDestinationsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListCupsDestinationsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListCupsDestinationsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListCupsDestinationsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListCupsDestinationsRequest) Execute() ([]SystemInsightsCupsDestinations, *http.Response, error) {
	return r.ApiService.SysteminsightsListCupsDestinationsExecute(r)
}

/*
SysteminsightsListCupsDestinations List System Insights CUPS Destinations

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListCupsDestinationsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListCupsDestinations(ctx context.Context) SystemInsightsApiSysteminsightsListCupsDestinationsRequest {
	return SystemInsightsApiSysteminsightsListCupsDestinationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsCupsDestinations
func (a *SystemInsightsApiService) SysteminsightsListCupsDestinationsExecute(r SystemInsightsApiSysteminsightsListCupsDestinationsRequest) ([]SystemInsightsCupsDestinations, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsCupsDestinations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListCupsDestinations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/cups_destinations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListDiskEncryptionRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListDiskEncryptionRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListDiskEncryptionRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListDiskEncryptionRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListDiskEncryptionRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListDiskEncryptionRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListDiskEncryptionRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListDiskEncryptionRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListDiskEncryptionRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListDiskEncryptionRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListDiskEncryptionRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListDiskEncryptionRequest) Execute() ([]SystemInsightsDiskEncryption, *http.Response, error) {
	return r.ApiService.SysteminsightsListDiskEncryptionExecute(r)
}

/*
SysteminsightsListDiskEncryption List System Insights Disk Encryption

Valid filter fields are `system_id` and `encryption_status`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListDiskEncryptionRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListDiskEncryption(ctx context.Context) SystemInsightsApiSysteminsightsListDiskEncryptionRequest {
	return SystemInsightsApiSysteminsightsListDiskEncryptionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsDiskEncryption
func (a *SystemInsightsApiService) SysteminsightsListDiskEncryptionExecute(r SystemInsightsApiSysteminsightsListDiskEncryptionRequest) ([]SystemInsightsDiskEncryption, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsDiskEncryption
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListDiskEncryption")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/disk_encryption"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListDiskInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListDiskInfoRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListDiskInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListDiskInfoRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListDiskInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListDiskInfoRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListDiskInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListDiskInfoRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListDiskInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListDiskInfoRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListDiskInfoRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListDiskInfoRequest) Execute() ([]SystemInsightsDiskInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListDiskInfoExecute(r)
}

/*
SysteminsightsListDiskInfo List System Insights Disk Info

Valid filter fields are `system_id` and `disk_index`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListDiskInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListDiskInfo(ctx context.Context) SystemInsightsApiSysteminsightsListDiskInfoRequest {
	return SystemInsightsApiSysteminsightsListDiskInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsDiskInfo
func (a *SystemInsightsApiService) SysteminsightsListDiskInfoExecute(r SystemInsightsApiSysteminsightsListDiskInfoRequest) ([]SystemInsightsDiskInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsDiskInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListDiskInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/disk_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListDnsResolversRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListDnsResolversRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListDnsResolversRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListDnsResolversRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListDnsResolversRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListDnsResolversRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListDnsResolversRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListDnsResolversRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListDnsResolversRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListDnsResolversRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListDnsResolversRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListDnsResolversRequest) Execute() ([]SystemInsightsDnsResolvers, *http.Response, error) {
	return r.ApiService.SysteminsightsListDnsResolversExecute(r)
}

/*
SysteminsightsListDnsResolvers List System Insights DNS Resolvers

Valid filter fields are `system_id` and `type`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListDnsResolversRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListDnsResolvers(ctx context.Context) SystemInsightsApiSysteminsightsListDnsResolversRequest {
	return SystemInsightsApiSysteminsightsListDnsResolversRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsDnsResolvers
func (a *SystemInsightsApiService) SysteminsightsListDnsResolversExecute(r SystemInsightsApiSysteminsightsListDnsResolversRequest) ([]SystemInsightsDnsResolvers, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsDnsResolvers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListDnsResolvers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/dns_resolvers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListEtcHostsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListEtcHostsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListEtcHostsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListEtcHostsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListEtcHostsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListEtcHostsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListEtcHostsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListEtcHostsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListEtcHostsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListEtcHostsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListEtcHostsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListEtcHostsRequest) Execute() ([]SystemInsightsEtcHosts, *http.Response, error) {
	return r.ApiService.SysteminsightsListEtcHostsExecute(r)
}

/*
SysteminsightsListEtcHosts List System Insights Etc Hosts

Valid filter fields are `system_id` and `address`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListEtcHostsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListEtcHosts(ctx context.Context) SystemInsightsApiSysteminsightsListEtcHostsRequest {
	return SystemInsightsApiSysteminsightsListEtcHostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsEtcHosts
func (a *SystemInsightsApiService) SysteminsightsListEtcHostsExecute(r SystemInsightsApiSysteminsightsListEtcHostsRequest) ([]SystemInsightsEtcHosts, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsEtcHosts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListEtcHosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/etc_hosts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListFirefoxAddonsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListFirefoxAddonsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListFirefoxAddonsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListFirefoxAddonsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListFirefoxAddonsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListFirefoxAddonsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListFirefoxAddonsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListFirefoxAddonsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListFirefoxAddonsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListFirefoxAddonsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListFirefoxAddonsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListFirefoxAddonsRequest) Execute() ([]SystemInsightsFirefoxAddons, *http.Response, error) {
	return r.ApiService.SysteminsightsListFirefoxAddonsExecute(r)
}

/*
SysteminsightsListFirefoxAddons List System Insights Firefox Addons

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListFirefoxAddonsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListFirefoxAddons(ctx context.Context) SystemInsightsApiSysteminsightsListFirefoxAddonsRequest {
	return SystemInsightsApiSysteminsightsListFirefoxAddonsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsFirefoxAddons
func (a *SystemInsightsApiService) SysteminsightsListFirefoxAddonsExecute(r SystemInsightsApiSysteminsightsListFirefoxAddonsRequest) ([]SystemInsightsFirefoxAddons, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsFirefoxAddons
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListFirefoxAddons")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/firefox_addons"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListGroupsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListGroupsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListGroupsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListGroupsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListGroupsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListGroupsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListGroupsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListGroupsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListGroupsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListGroupsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListGroupsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListGroupsRequest) Execute() ([]SystemInsightsGroups, *http.Response, error) {
	return r.ApiService.SysteminsightsListGroupsExecute(r)
}

/*
SysteminsightsListGroups List System Insights Groups

Valid filter fields are `system_id` and `groupname`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListGroupsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListGroups(ctx context.Context) SystemInsightsApiSysteminsightsListGroupsRequest {
	return SystemInsightsApiSysteminsightsListGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsGroups
func (a *SystemInsightsApiService) SysteminsightsListGroupsExecute(r SystemInsightsApiSysteminsightsListGroupsRequest) ([]SystemInsightsGroups, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsGroups
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListIeExtensionsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListIeExtensionsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListIeExtensionsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListIeExtensionsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListIeExtensionsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListIeExtensionsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListIeExtensionsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListIeExtensionsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListIeExtensionsRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListIeExtensionsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListIeExtensionsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListIeExtensionsRequest) Execute() ([]SystemInsightsIeExtensions, *http.Response, error) {
	return r.ApiService.SysteminsightsListIeExtensionsExecute(r)
}

/*
SysteminsightsListIeExtensions List System Insights IE Extensions

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListIeExtensionsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListIeExtensions(ctx context.Context) SystemInsightsApiSysteminsightsListIeExtensionsRequest {
	return SystemInsightsApiSysteminsightsListIeExtensionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsIeExtensions
func (a *SystemInsightsApiService) SysteminsightsListIeExtensionsExecute(r SystemInsightsApiSysteminsightsListIeExtensionsRequest) ([]SystemInsightsIeExtensions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsIeExtensions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListIeExtensions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/ie_extensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListInterfaceAddressesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListInterfaceAddressesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListInterfaceAddressesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListInterfaceAddressesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListInterfaceAddressesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListInterfaceAddressesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListInterfaceAddressesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListInterfaceAddressesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListInterfaceAddressesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListInterfaceAddressesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListInterfaceAddressesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListInterfaceAddressesRequest) Execute() ([]SystemInsightsInterfaceAddresses, *http.Response, error) {
	return r.ApiService.SysteminsightsListInterfaceAddressesExecute(r)
}

/*
SysteminsightsListInterfaceAddresses List System Insights Interface Addresses

Valid filter fields are `system_id` and `address`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListInterfaceAddressesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListInterfaceAddresses(ctx context.Context) SystemInsightsApiSysteminsightsListInterfaceAddressesRequest {
	return SystemInsightsApiSysteminsightsListInterfaceAddressesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsInterfaceAddresses
func (a *SystemInsightsApiService) SysteminsightsListInterfaceAddressesExecute(r SystemInsightsApiSysteminsightsListInterfaceAddressesRequest) ([]SystemInsightsInterfaceAddresses, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsInterfaceAddresses
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListInterfaceAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/interface_addresses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListInterfaceDetailsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListInterfaceDetailsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListInterfaceDetailsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListInterfaceDetailsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListInterfaceDetailsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListInterfaceDetailsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListInterfaceDetailsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListInterfaceDetailsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListInterfaceDetailsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListInterfaceDetailsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListInterfaceDetailsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListInterfaceDetailsRequest) Execute() ([]SystemInsightsInterfaceDetails, *http.Response, error) {
	return r.ApiService.SysteminsightsListInterfaceDetailsExecute(r)
}

/*
SysteminsightsListInterfaceDetails List System Insights Interface Details

Valid filter fields are `system_id` and `interface`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListInterfaceDetailsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListInterfaceDetails(ctx context.Context) SystemInsightsApiSysteminsightsListInterfaceDetailsRequest {
	return SystemInsightsApiSysteminsightsListInterfaceDetailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsInterfaceDetails
func (a *SystemInsightsApiService) SysteminsightsListInterfaceDetailsExecute(r SystemInsightsApiSysteminsightsListInterfaceDetailsRequest) ([]SystemInsightsInterfaceDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsInterfaceDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListInterfaceDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/interface_details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListKernelInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListKernelInfoRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListKernelInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListKernelInfoRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListKernelInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListKernelInfoRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListKernelInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListKernelInfoRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListKernelInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListKernelInfoRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListKernelInfoRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListKernelInfoRequest) Execute() ([]SystemInsightsKernelInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListKernelInfoExecute(r)
}

/*
SysteminsightsListKernelInfo List System Insights Kernel Info

Valid filter fields are `system_id` and `version`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListKernelInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListKernelInfo(ctx context.Context) SystemInsightsApiSysteminsightsListKernelInfoRequest {
	return SystemInsightsApiSysteminsightsListKernelInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsKernelInfo
func (a *SystemInsightsApiService) SysteminsightsListKernelInfoExecute(r SystemInsightsApiSysteminsightsListKernelInfoRequest) ([]SystemInsightsKernelInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsKernelInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListKernelInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/kernel_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListLaunchdRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListLaunchdRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListLaunchdRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListLaunchdRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListLaunchdRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListLaunchdRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListLaunchdRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListLaunchdRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListLaunchdRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListLaunchdRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListLaunchdRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListLaunchdRequest) Execute() ([]SystemInsightsLaunchd, *http.Response, error) {
	return r.ApiService.SysteminsightsListLaunchdExecute(r)
}

/*
SysteminsightsListLaunchd List System Insights Launchd

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListLaunchdRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListLaunchd(ctx context.Context) SystemInsightsApiSysteminsightsListLaunchdRequest {
	return SystemInsightsApiSysteminsightsListLaunchdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsLaunchd
func (a *SystemInsightsApiService) SysteminsightsListLaunchdExecute(r SystemInsightsApiSysteminsightsListLaunchdRequest) ([]SystemInsightsLaunchd, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsLaunchd
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListLaunchd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/launchd"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListLinuxPackagesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListLinuxPackagesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListLinuxPackagesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListLinuxPackagesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListLinuxPackagesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListLinuxPackagesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListLinuxPackagesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListLinuxPackagesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListLinuxPackagesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListLinuxPackagesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListLinuxPackagesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListLinuxPackagesRequest) Execute() ([]SystemInsightsLinuxPackages, *http.Response, error) {
	return r.ApiService.SysteminsightsListLinuxPackagesExecute(r)
}

/*
SysteminsightsListLinuxPackages List System Insights Linux Packages

Lists all programs for Linux devices. For macOS devices, use [List System Insights System Apps](#operation/systeminsights_list_apps). For windows devices, use [List System Insights System Apps](#operation/systeminsights_list_programs).

Valid filter fields are `name` and `package_format`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListLinuxPackagesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListLinuxPackages(ctx context.Context) SystemInsightsApiSysteminsightsListLinuxPackagesRequest {
	return SystemInsightsApiSysteminsightsListLinuxPackagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsLinuxPackages
func (a *SystemInsightsApiService) SysteminsightsListLinuxPackagesExecute(r SystemInsightsApiSysteminsightsListLinuxPackagesRequest) ([]SystemInsightsLinuxPackages, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsLinuxPackages
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListLinuxPackages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/linux_packages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListLoggedInUsersRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListLoggedInUsersRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListLoggedInUsersRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListLoggedInUsersRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListLoggedInUsersRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListLoggedInUsersRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListLoggedInUsersRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListLoggedInUsersRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListLoggedInUsersRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListLoggedInUsersRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListLoggedInUsersRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListLoggedInUsersRequest) Execute() ([]SystemInsightsLoggedInUsers, *http.Response, error) {
	return r.ApiService.SysteminsightsListLoggedInUsersExecute(r)
}

/*
SysteminsightsListLoggedInUsers List System Insights Logged-In Users

Valid filter fields are `system_id` and `user`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListLoggedInUsersRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListLoggedInUsers(ctx context.Context) SystemInsightsApiSysteminsightsListLoggedInUsersRequest {
	return SystemInsightsApiSysteminsightsListLoggedInUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsLoggedInUsers
func (a *SystemInsightsApiService) SysteminsightsListLoggedInUsersExecute(r SystemInsightsApiSysteminsightsListLoggedInUsersRequest) ([]SystemInsightsLoggedInUsers, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsLoggedInUsers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListLoggedInUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/logged_in_users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListLogicalDrivesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListLogicalDrivesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListLogicalDrivesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListLogicalDrivesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListLogicalDrivesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListLogicalDrivesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListLogicalDrivesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListLogicalDrivesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListLogicalDrivesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListLogicalDrivesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListLogicalDrivesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListLogicalDrivesRequest) Execute() ([]SystemInsightsLogicalDrives, *http.Response, error) {
	return r.ApiService.SysteminsightsListLogicalDrivesExecute(r)
}

/*
SysteminsightsListLogicalDrives List System Insights Logical Drives

Valid filter fields are `system_id` and `device_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListLogicalDrivesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListLogicalDrives(ctx context.Context) SystemInsightsApiSysteminsightsListLogicalDrivesRequest {
	return SystemInsightsApiSysteminsightsListLogicalDrivesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsLogicalDrives
func (a *SystemInsightsApiService) SysteminsightsListLogicalDrivesExecute(r SystemInsightsApiSysteminsightsListLogicalDrivesRequest) ([]SystemInsightsLogicalDrives, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsLogicalDrives
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListLogicalDrives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/logical_drives"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListManagedPoliciesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListManagedPoliciesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListManagedPoliciesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListManagedPoliciesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListManagedPoliciesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListManagedPoliciesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListManagedPoliciesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListManagedPoliciesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListManagedPoliciesRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListManagedPoliciesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListManagedPoliciesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListManagedPoliciesRequest) Execute() ([]SystemInsightsManagedPolicies, *http.Response, error) {
	return r.ApiService.SysteminsightsListManagedPoliciesExecute(r)
}

/*
SysteminsightsListManagedPolicies List System Insights Managed Policies

Valid filter fields are `system_id` and `domain`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListManagedPoliciesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListManagedPolicies(ctx context.Context) SystemInsightsApiSysteminsightsListManagedPoliciesRequest {
	return SystemInsightsApiSysteminsightsListManagedPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsManagedPolicies
func (a *SystemInsightsApiService) SysteminsightsListManagedPoliciesExecute(r SystemInsightsApiSysteminsightsListManagedPoliciesRequest) ([]SystemInsightsManagedPolicies, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsManagedPolicies
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListManagedPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/managed_policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListMountsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListMountsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListMountsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListMountsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListMountsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListMountsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListMountsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListMountsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListMountsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListMountsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListMountsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListMountsRequest) Execute() ([]SystemInsightsMounts, *http.Response, error) {
	return r.ApiService.SysteminsightsListMountsExecute(r)
}

/*
SysteminsightsListMounts List System Insights Mounts

Valid filter fields are `system_id` and `path`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListMountsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListMounts(ctx context.Context) SystemInsightsApiSysteminsightsListMountsRequest {
	return SystemInsightsApiSysteminsightsListMountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsMounts
func (a *SystemInsightsApiService) SysteminsightsListMountsExecute(r SystemInsightsApiSysteminsightsListMountsRequest) ([]SystemInsightsMounts, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsMounts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListMounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/mounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListOsVersionRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListOsVersionRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListOsVersionRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListOsVersionRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListOsVersionRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListOsVersionRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListOsVersionRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListOsVersionRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListOsVersionRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListOsVersionRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListOsVersionRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListOsVersionRequest) Execute() ([]SystemInsightsOsVersion, *http.Response, error) {
	return r.ApiService.SysteminsightsListOsVersionExecute(r)
}

/*
SysteminsightsListOsVersion List System Insights OS Version

Valid filter fields are `system_id` and `version`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListOsVersionRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListOsVersion(ctx context.Context) SystemInsightsApiSysteminsightsListOsVersionRequest {
	return SystemInsightsApiSysteminsightsListOsVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsOsVersion
func (a *SystemInsightsApiService) SysteminsightsListOsVersionExecute(r SystemInsightsApiSysteminsightsListOsVersionRequest) ([]SystemInsightsOsVersion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsOsVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListOsVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/os_version"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListPatchesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListPatchesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListPatchesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListPatchesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListPatchesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListPatchesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListPatchesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListPatchesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListPatchesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListPatchesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListPatchesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListPatchesRequest) Execute() ([]SystemInsightsPatches, *http.Response, error) {
	return r.ApiService.SysteminsightsListPatchesExecute(r)
}

/*
SysteminsightsListPatches List System Insights Patches

Valid filter fields are `system_id` and `hotfix_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListPatchesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListPatches(ctx context.Context) SystemInsightsApiSysteminsightsListPatchesRequest {
	return SystemInsightsApiSysteminsightsListPatchesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsPatches
func (a *SystemInsightsApiService) SysteminsightsListPatchesExecute(r SystemInsightsApiSysteminsightsListPatchesRequest) ([]SystemInsightsPatches, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsPatches
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListPatches")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/patches"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListProgramsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListProgramsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListProgramsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListProgramsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListProgramsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListProgramsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListProgramsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListProgramsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListProgramsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListProgramsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListProgramsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListProgramsRequest) Execute() ([]SystemInsightsPrograms, *http.Response, error) {
	return r.ApiService.SysteminsightsListProgramsExecute(r)
}

/*
SysteminsightsListPrograms List System Insights Programs

Lists all programs for Windows devices. For macOS devices, use [List System Insights Apps](#operation/systeminsights_list_apps).

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListProgramsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListPrograms(ctx context.Context) SystemInsightsApiSysteminsightsListProgramsRequest {
	return SystemInsightsApiSysteminsightsListProgramsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsPrograms
func (a *SystemInsightsApiService) SysteminsightsListProgramsExecute(r SystemInsightsApiSysteminsightsListProgramsRequest) ([]SystemInsightsPrograms, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsPrograms
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListPrograms")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/programs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListPythonPackagesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListPythonPackagesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListPythonPackagesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListPythonPackagesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListPythonPackagesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListPythonPackagesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListPythonPackagesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListPythonPackagesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListPythonPackagesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListPythonPackagesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListPythonPackagesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListPythonPackagesRequest) Execute() ([]SystemInsightsPythonPackages, *http.Response, error) {
	return r.ApiService.SysteminsightsListPythonPackagesExecute(r)
}

/*
SysteminsightsListPythonPackages List System Insights Python Packages

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListPythonPackagesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListPythonPackages(ctx context.Context) SystemInsightsApiSysteminsightsListPythonPackagesRequest {
	return SystemInsightsApiSysteminsightsListPythonPackagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsPythonPackages
func (a *SystemInsightsApiService) SysteminsightsListPythonPackagesExecute(r SystemInsightsApiSysteminsightsListPythonPackagesRequest) ([]SystemInsightsPythonPackages, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsPythonPackages
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListPythonPackages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/python_packages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListSafariExtensionsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListSafariExtensionsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListSafariExtensionsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListSafariExtensionsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListSafariExtensionsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListSafariExtensionsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListSafariExtensionsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListSafariExtensionsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListSafariExtensionsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListSafariExtensionsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListSafariExtensionsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListSafariExtensionsRequest) Execute() ([]SystemInsightsSafariExtensions, *http.Response, error) {
	return r.ApiService.SysteminsightsListSafariExtensionsExecute(r)
}

/*
SysteminsightsListSafariExtensions List System Insights Safari Extensions

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListSafariExtensionsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSafariExtensions(ctx context.Context) SystemInsightsApiSysteminsightsListSafariExtensionsRequest {
	return SystemInsightsApiSysteminsightsListSafariExtensionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSafariExtensions
func (a *SystemInsightsApiService) SysteminsightsListSafariExtensionsExecute(r SystemInsightsApiSysteminsightsListSafariExtensionsRequest) ([]SystemInsightsSafariExtensions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsSafariExtensions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListSafariExtensions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/safari_extensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListScheduledTasksRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListScheduledTasksRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListScheduledTasksRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListScheduledTasksRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListScheduledTasksRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListScheduledTasksRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListScheduledTasksRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListScheduledTasksRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListScheduledTasksRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListScheduledTasksRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListScheduledTasksRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListScheduledTasksRequest) Execute() ([]SystemInsightsScheduledTasks, *http.Response, error) {
	return r.ApiService.SysteminsightsListScheduledTasksExecute(r)
}

/*
SysteminsightsListScheduledTasks List System Insights Scheduled Tasks

Valid filter fields are `system_id` and `enabled`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListScheduledTasksRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListScheduledTasks(ctx context.Context) SystemInsightsApiSysteminsightsListScheduledTasksRequest {
	return SystemInsightsApiSysteminsightsListScheduledTasksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsScheduledTasks
func (a *SystemInsightsApiService) SysteminsightsListScheduledTasksExecute(r SystemInsightsApiSysteminsightsListScheduledTasksRequest) ([]SystemInsightsScheduledTasks, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsScheduledTasks
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListScheduledTasks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/scheduled_tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListSecurebootRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListSecurebootRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListSecurebootRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListSecurebootRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListSecurebootRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListSecurebootRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListSecurebootRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListSecurebootRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListSecurebootRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListSecurebootRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListSecurebootRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListSecurebootRequest) Execute() ([]SystemInsightsSecureboot, *http.Response, error) {
	return r.ApiService.SysteminsightsListSecurebootExecute(r)
}

/*
SysteminsightsListSecureboot List System Insights Secure Boot

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListSecurebootRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSecureboot(ctx context.Context) SystemInsightsApiSysteminsightsListSecurebootRequest {
	return SystemInsightsApiSysteminsightsListSecurebootRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSecureboot
func (a *SystemInsightsApiService) SysteminsightsListSecurebootExecute(r SystemInsightsApiSysteminsightsListSecurebootRequest) ([]SystemInsightsSecureboot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsSecureboot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListSecureboot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/secureboot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
			var v string
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

type SystemInsightsApiSysteminsightsListServicesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListServicesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListServicesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListServicesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListServicesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListServicesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListServicesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListServicesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListServicesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListServicesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListServicesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListServicesRequest) Execute() ([]SystemInsightsServices, *http.Response, error) {
	return r.ApiService.SysteminsightsListServicesExecute(r)
}

/*
SysteminsightsListServices List System Insights Services

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListServicesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListServices(ctx context.Context) SystemInsightsApiSysteminsightsListServicesRequest {
	return SystemInsightsApiSysteminsightsListServicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsServices
func (a *SystemInsightsApiService) SysteminsightsListServicesExecute(r SystemInsightsApiSysteminsightsListServicesRequest) ([]SystemInsightsServices, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsServices
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/services"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListShadowRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListShadowRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListShadowRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListShadowRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListShadowRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListShadowRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListShadowRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListShadowRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListShadowRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListShadowRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListShadowRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListShadowRequest) Execute() ([]SystemInsightsShadow, *http.Response, error) {
	return r.ApiService.SysteminsightsListShadowExecute(r)
}

/*
SysteminsightsListShadow LIst System Insights Shadow

Valid filter fields are `system_id` and `username`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListShadowRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListShadow(ctx context.Context) SystemInsightsApiSysteminsightsListShadowRequest {
	return SystemInsightsApiSysteminsightsListShadowRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsShadow
func (a *SystemInsightsApiService) SysteminsightsListShadowExecute(r SystemInsightsApiSysteminsightsListShadowRequest) ([]SystemInsightsShadow, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsShadow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListShadow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/shadow"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListSharedFoldersRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListSharedFoldersRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListSharedFoldersRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListSharedFoldersRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListSharedFoldersRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListSharedFoldersRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListSharedFoldersRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListSharedFoldersRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListSharedFoldersRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListSharedFoldersRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListSharedFoldersRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListSharedFoldersRequest) Execute() ([]SystemInsightsSharedFolders, *http.Response, error) {
	return r.ApiService.SysteminsightsListSharedFoldersExecute(r)
}

/*
SysteminsightsListSharedFolders List System Insights Shared Folders

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListSharedFoldersRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSharedFolders(ctx context.Context) SystemInsightsApiSysteminsightsListSharedFoldersRequest {
	return SystemInsightsApiSysteminsightsListSharedFoldersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSharedFolders
func (a *SystemInsightsApiService) SysteminsightsListSharedFoldersExecute(r SystemInsightsApiSysteminsightsListSharedFoldersRequest) ([]SystemInsightsSharedFolders, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsSharedFolders
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListSharedFolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/shared_folders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListSharedResourcesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListSharedResourcesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListSharedResourcesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListSharedResourcesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListSharedResourcesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListSharedResourcesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListSharedResourcesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListSharedResourcesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListSharedResourcesRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListSharedResourcesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListSharedResourcesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListSharedResourcesRequest) Execute() ([]SystemInsightsSharedResources, *http.Response, error) {
	return r.ApiService.SysteminsightsListSharedResourcesExecute(r)
}

/*
SysteminsightsListSharedResources List System Insights Shared Resources

Valid filter fields are `system_id` and `type`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListSharedResourcesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSharedResources(ctx context.Context) SystemInsightsApiSysteminsightsListSharedResourcesRequest {
	return SystemInsightsApiSysteminsightsListSharedResourcesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSharedResources
func (a *SystemInsightsApiService) SysteminsightsListSharedResourcesExecute(r SystemInsightsApiSysteminsightsListSharedResourcesRequest) ([]SystemInsightsSharedResources, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsSharedResources
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListSharedResources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/shared_resources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListSharingPreferencesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListSharingPreferencesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListSharingPreferencesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListSharingPreferencesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListSharingPreferencesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListSharingPreferencesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListSharingPreferencesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListSharingPreferencesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListSharingPreferencesRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListSharingPreferencesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListSharingPreferencesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListSharingPreferencesRequest) Execute() ([]SystemInsightsSharingPreferences, *http.Response, error) {
	return r.ApiService.SysteminsightsListSharingPreferencesExecute(r)
}

/*
SysteminsightsListSharingPreferences List System Insights Sharing Preferences

Only valid filed field is `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListSharingPreferencesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSharingPreferences(ctx context.Context) SystemInsightsApiSysteminsightsListSharingPreferencesRequest {
	return SystemInsightsApiSysteminsightsListSharingPreferencesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSharingPreferences
func (a *SystemInsightsApiService) SysteminsightsListSharingPreferencesExecute(r SystemInsightsApiSysteminsightsListSharingPreferencesRequest) ([]SystemInsightsSharingPreferences, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsSharingPreferences
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListSharingPreferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/sharing_preferences"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListSipConfigRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListSipConfigRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListSipConfigRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListSipConfigRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListSipConfigRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListSipConfigRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListSipConfigRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListSipConfigRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListSipConfigRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListSipConfigRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListSipConfigRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListSipConfigRequest) Execute() ([]SystemInsightsSipConfig, *http.Response, error) {
	return r.ApiService.SysteminsightsListSipConfigExecute(r)
}

/*
SysteminsightsListSipConfig List System Insights SIP Config

Valid filter fields are `system_id` and `enabled`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListSipConfigRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSipConfig(ctx context.Context) SystemInsightsApiSysteminsightsListSipConfigRequest {
	return SystemInsightsApiSysteminsightsListSipConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSipConfig
func (a *SystemInsightsApiService) SysteminsightsListSipConfigExecute(r SystemInsightsApiSysteminsightsListSipConfigRequest) ([]SystemInsightsSipConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsSipConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListSipConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/sip_config"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListStartupItemsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListStartupItemsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListStartupItemsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListStartupItemsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListStartupItemsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListStartupItemsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListStartupItemsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListStartupItemsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListStartupItemsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListStartupItemsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListStartupItemsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListStartupItemsRequest) Execute() ([]SystemInsightsStartupItems, *http.Response, error) {
	return r.ApiService.SysteminsightsListStartupItemsExecute(r)
}

/*
SysteminsightsListStartupItems List System Insights Startup Items

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListStartupItemsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListStartupItems(ctx context.Context) SystemInsightsApiSysteminsightsListStartupItemsRequest {
	return SystemInsightsApiSysteminsightsListStartupItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsStartupItems
func (a *SystemInsightsApiService) SysteminsightsListStartupItemsExecute(r SystemInsightsApiSysteminsightsListStartupItemsRequest) ([]SystemInsightsStartupItems, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsStartupItems
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListStartupItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/startup_items"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListSystemControlsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListSystemControlsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListSystemControlsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListSystemControlsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListSystemControlsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; Note: You can only filter by &#x60;system_id&#x60; and &#x60;name&#x60; 
func (r SystemInsightsApiSysteminsightsListSystemControlsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListSystemControlsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListSystemControlsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListSystemControlsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListSystemControlsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListSystemControlsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListSystemControlsRequest) Execute() ([]SystemInsightsSystemControls, *http.Response, error) {
	return r.ApiService.SysteminsightsListSystemControlsExecute(r)
}

/*
SysteminsightsListSystemControls List System Insights System Control

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListSystemControlsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSystemControls(ctx context.Context) SystemInsightsApiSysteminsightsListSystemControlsRequest {
	return SystemInsightsApiSysteminsightsListSystemControlsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSystemControls
func (a *SystemInsightsApiService) SysteminsightsListSystemControlsExecute(r SystemInsightsApiSysteminsightsListSystemControlsRequest) ([]SystemInsightsSystemControls, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsSystemControls
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListSystemControls")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/system_controls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListSystemInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListSystemInfoRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListSystemInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListSystemInfoRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListSystemInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListSystemInfoRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListSystemInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListSystemInfoRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListSystemInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListSystemInfoRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListSystemInfoRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListSystemInfoRequest) Execute() ([]SystemInsightsSystemInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListSystemInfoExecute(r)
}

/*
SysteminsightsListSystemInfo List System Insights System Info

Valid filter fields are `system_id` and `cpu_subtype`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListSystemInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSystemInfo(ctx context.Context) SystemInsightsApiSysteminsightsListSystemInfoRequest {
	return SystemInsightsApiSysteminsightsListSystemInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSystemInfo
func (a *SystemInsightsApiService) SysteminsightsListSystemInfoExecute(r SystemInsightsApiSysteminsightsListSystemInfoRequest) ([]SystemInsightsSystemInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsSystemInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListSystemInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/system_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListTpmInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListTpmInfoRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListTpmInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListTpmInfoRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListTpmInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListTpmInfoRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListTpmInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListTpmInfoRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListTpmInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListTpmInfoRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListTpmInfoRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListTpmInfoRequest) Execute() ([]SystemInsightsTpmInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListTpmInfoExecute(r)
}

/*
SysteminsightsListTpmInfo List System Insights TPM Info

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListTpmInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListTpmInfo(ctx context.Context) SystemInsightsApiSysteminsightsListTpmInfoRequest {
	return SystemInsightsApiSysteminsightsListTpmInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsTpmInfo
func (a *SystemInsightsApiService) SysteminsightsListTpmInfoExecute(r SystemInsightsApiSysteminsightsListTpmInfoRequest) ([]SystemInsightsTpmInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsTpmInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListTpmInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/tpm_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xOrgId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-org-id", r.xOrgId, "")
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
			var v string
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

type SystemInsightsApiSysteminsightsListUptimeRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListUptimeRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListUptimeRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListUptimeRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListUptimeRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, gte, in. e.g: Filter for single value: &#x60;filter&#x3D;field:gte:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListUptimeRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListUptimeRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListUptimeRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListUptimeRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListUptimeRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListUptimeRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListUptimeRequest) Execute() ([]SystemInsightsUptime, *http.Response, error) {
	return r.ApiService.SysteminsightsListUptimeExecute(r)
}

/*
SysteminsightsListUptime List System Insights Uptime

Valid filter fields are `system_id` and `days`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListUptimeRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUptime(ctx context.Context) SystemInsightsApiSysteminsightsListUptimeRequest {
	return SystemInsightsApiSysteminsightsListUptimeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUptime
func (a *SystemInsightsApiService) SysteminsightsListUptimeExecute(r SystemInsightsApiSysteminsightsListUptimeRequest) ([]SystemInsightsUptime, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsUptime
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListUptime")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/uptime"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListUsbDevicesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListUsbDevicesRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListUsbDevicesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListUsbDevicesRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListUsbDevicesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListUsbDevicesRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListUsbDevicesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListUsbDevicesRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListUsbDevicesRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListUsbDevicesRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListUsbDevicesRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListUsbDevicesRequest) Execute() ([]SystemInsightsUsbDevices, *http.Response, error) {
	return r.ApiService.SysteminsightsListUsbDevicesExecute(r)
}

/*
SysteminsightsListUsbDevices List System Insights USB Devices

Valid filter fields are `system_id` and `model`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListUsbDevicesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUsbDevices(ctx context.Context) SystemInsightsApiSysteminsightsListUsbDevicesRequest {
	return SystemInsightsApiSysteminsightsListUsbDevicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUsbDevices
func (a *SystemInsightsApiService) SysteminsightsListUsbDevicesExecute(r SystemInsightsApiSysteminsightsListUsbDevicesRequest) ([]SystemInsightsUsbDevices, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsUsbDevices
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListUsbDevices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/usb_devices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListUserGroupsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListUserGroupsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListUserGroupsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListUserGroupsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListUserGroupsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListUserGroupsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListUserGroupsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListUserGroupsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListUserGroupsRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListUserGroupsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListUserGroupsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListUserGroupsRequest) Execute() ([]SystemInsightsUserGroups, *http.Response, error) {
	return r.ApiService.SysteminsightsListUserGroupsExecute(r)
}

/*
SysteminsightsListUserGroups List System Insights User Groups

Only valid filter field is `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListUserGroupsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUserGroups(ctx context.Context) SystemInsightsApiSysteminsightsListUserGroupsRequest {
	return SystemInsightsApiSysteminsightsListUserGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUserGroups
func (a *SystemInsightsApiService) SysteminsightsListUserGroupsExecute(r SystemInsightsApiSysteminsightsListUserGroupsRequest) ([]SystemInsightsUserGroups, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsUserGroups
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListUserGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/user_groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListUserSshKeysRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListUserSshKeysRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListUserSshKeysRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListUserSshKeysRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListUserSshKeysRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListUserSshKeysRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListUserSshKeysRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListUserSshKeysRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListUserSshKeysRequest {
	r.filter = &filter
	return r
}

func (r SystemInsightsApiSysteminsightsListUserSshKeysRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListUserSshKeysRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListUserSshKeysRequest) Execute() ([]SystemInsightsUserSshKeys, *http.Response, error) {
	return r.ApiService.SysteminsightsListUserSshKeysExecute(r)
}

/*
SysteminsightsListUserSshKeys List System Insights User SSH Keys

Valid filter fields are `system_id` and `uid`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListUserSshKeysRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUserSshKeys(ctx context.Context) SystemInsightsApiSysteminsightsListUserSshKeysRequest {
	return SystemInsightsApiSysteminsightsListUserSshKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUserSshKeys
func (a *SystemInsightsApiService) SysteminsightsListUserSshKeysExecute(r SystemInsightsApiSysteminsightsListUserSshKeysRequest) ([]SystemInsightsUserSshKeys, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsUserSshKeys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListUserSshKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/user_ssh_keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListUserassistRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListUserassistRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListUserassistRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListUserassistRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListUserassistRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListUserassistRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListUserassistRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListUserassistRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListUserassistRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListUserassistRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListUserassistRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListUserassistRequest) Execute() ([]SystemInsightsUserassist, *http.Response, error) {
	return r.ApiService.SysteminsightsListUserassistExecute(r)
}

/*
SysteminsightsListUserassist List System Insights User Assist

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListUserassistRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUserassist(ctx context.Context) SystemInsightsApiSysteminsightsListUserassistRequest {
	return SystemInsightsApiSysteminsightsListUserassistRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUserassist
func (a *SystemInsightsApiService) SysteminsightsListUserassistExecute(r SystemInsightsApiSysteminsightsListUserassistRequest) ([]SystemInsightsUserassist, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsUserassist
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListUserassist")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/userassist"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
			var v string
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

type SystemInsightsApiSysteminsightsListUsersRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListUsersRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListUsersRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListUsersRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListUsersRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListUsersRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListUsersRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListUsersRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListUsersRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListUsersRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListUsersRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListUsersRequest) Execute() ([]SystemInsightsUsers, *http.Response, error) {
	return r.ApiService.SysteminsightsListUsersExecute(r)
}

/*
SysteminsightsListUsers List System Insights Users

Valid filter fields are `system_id` and `username`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListUsersRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUsers(ctx context.Context) SystemInsightsApiSysteminsightsListUsersRequest {
	return SystemInsightsApiSysteminsightsListUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUsers
func (a *SystemInsightsApiService) SysteminsightsListUsersExecute(r SystemInsightsApiSysteminsightsListUsersRequest) ([]SystemInsightsUsers, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsUsers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListWifiNetworksRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListWifiNetworksRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListWifiNetworksRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListWifiNetworksRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListWifiNetworksRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListWifiNetworksRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListWifiNetworksRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListWifiNetworksRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListWifiNetworksRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListWifiNetworksRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListWifiNetworksRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListWifiNetworksRequest) Execute() ([]SystemInsightsWifiNetworks, *http.Response, error) {
	return r.ApiService.SysteminsightsListWifiNetworksExecute(r)
}

/*
SysteminsightsListWifiNetworks List System Insights WiFi Networks

Valid filter fields are `system_id` and `security_type`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListWifiNetworksRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListWifiNetworks(ctx context.Context) SystemInsightsApiSysteminsightsListWifiNetworksRequest {
	return SystemInsightsApiSysteminsightsListWifiNetworksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsWifiNetworks
func (a *SystemInsightsApiService) SysteminsightsListWifiNetworksExecute(r SystemInsightsApiSysteminsightsListWifiNetworksRequest) ([]SystemInsightsWifiNetworks, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsWifiNetworks
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListWifiNetworks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/wifi_networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListWifiStatusRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListWifiStatusRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListWifiStatusRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListWifiStatusRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListWifiStatusRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListWifiStatusRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListWifiStatusRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListWifiStatusRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListWifiStatusRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListWifiStatusRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListWifiStatusRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListWifiStatusRequest) Execute() ([]SystemInsightsWifiStatus, *http.Response, error) {
	return r.ApiService.SysteminsightsListWifiStatusExecute(r)
}

/*
SysteminsightsListWifiStatus List System Insights WiFi Status

Valid filter fields are `system_id` and `security_type`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListWifiStatusRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListWifiStatus(ctx context.Context) SystemInsightsApiSysteminsightsListWifiStatusRequest {
	return SystemInsightsApiSysteminsightsListWifiStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsWifiStatus
func (a *SystemInsightsApiService) SysteminsightsListWifiStatusExecute(r SystemInsightsApiSysteminsightsListWifiStatusRequest) ([]SystemInsightsWifiStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsWifiStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListWifiStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/wifi_status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest) Execute() ([]SystemInsightsWindowsSecurityCenter, *http.Response, error) {
	return r.ApiService.SysteminsightsListWindowsSecurityCenterExecute(r)
}

/*
SysteminsightsListWindowsSecurityCenter List System Insights Windows Security Center

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListWindowsSecurityCenter(ctx context.Context) SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest {
	return SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsWindowsSecurityCenter
func (a *SystemInsightsApiService) SysteminsightsListWindowsSecurityCenterExecute(r SystemInsightsApiSysteminsightsListWindowsSecurityCenterRequest) ([]SystemInsightsWindowsSecurityCenter, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsWindowsSecurityCenter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListWindowsSecurityCenter")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/windows_security_center"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
			var v string
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

type SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest) Skip(skip int32) SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest) Sort(sort []string) SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest) Filter(filter []string) SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest) XOrgId(xOrgId string) SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest) Limit(limit int32) SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest {
	r.limit = &limit
	return r
}

func (r SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest) Execute() ([]SystemInsightsWindowsSecurityProducts, *http.Response, error) {
	return r.ApiService.SysteminsightsListWindowsSecurityProductsExecute(r)
}

/*
SysteminsightsListWindowsSecurityProducts List System Insights Windows Security Products

Valid filter fields are `system_id` and `state`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListWindowsSecurityProducts(ctx context.Context) SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest {
	return SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsWindowsSecurityProducts
func (a *SystemInsightsApiService) SysteminsightsListWindowsSecurityProductsExecute(r SystemInsightsApiSysteminsightsListWindowsSecurityProductsRequest) ([]SystemInsightsWindowsSecurityProducts, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SystemInsightsWindowsSecurityProducts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemInsightsApiService.SysteminsightsListWindowsSecurityProducts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/systeminsights/windows_security_products"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "csv")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

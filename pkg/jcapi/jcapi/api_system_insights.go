/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SystemInsightsApiService SystemInsightsApi service
type SystemInsightsApiService service

type ApiSysteminsightsListAlfRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	filter *[]string
	skip *int32
	sort *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListAlfRequest) XOrgId(xOrgId string) ApiSysteminsightsListAlfRequest {
	r.xOrgId = &xOrgId
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListAlfRequest) Filter(filter []string) ApiSysteminsightsListAlfRequest {
	r.filter = &filter
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListAlfRequest) Skip(skip int32) ApiSysteminsightsListAlfRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListAlfRequest) Sort(sort []string) ApiSysteminsightsListAlfRequest {
	r.sort = &sort
	return r
}

func (r ApiSysteminsightsListAlfRequest) Limit(limit int32) ApiSysteminsightsListAlfRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListAlfRequest) Execute() ([]SystemInsightsAlf, *http.Response, error) {
	return r.ApiService.SysteminsightsListAlfExecute(r)
}

/*
SysteminsightsListAlf List System Insights ALF

Valid filter fields are `system_id` and `global_state`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListAlfRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAlf(ctx context.Context) ApiSysteminsightsListAlfRequest {
	return ApiSysteminsightsListAlfRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAlf
func (a *SystemInsightsApiService) SysteminsightsListAlfExecute(r ApiSysteminsightsListAlfRequest) ([]SystemInsightsAlf, *http.Response, error) {
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

type ApiSysteminsightsListAlfExceptionsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	filter *[]string
	skip *int32
	sort *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListAlfExceptionsRequest) XOrgId(xOrgId string) ApiSysteminsightsListAlfExceptionsRequest {
	r.xOrgId = &xOrgId
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListAlfExceptionsRequest) Filter(filter []string) ApiSysteminsightsListAlfExceptionsRequest {
	r.filter = &filter
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListAlfExceptionsRequest) Skip(skip int32) ApiSysteminsightsListAlfExceptionsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListAlfExceptionsRequest) Sort(sort []string) ApiSysteminsightsListAlfExceptionsRequest {
	r.sort = &sort
	return r
}

func (r ApiSysteminsightsListAlfExceptionsRequest) Limit(limit int32) ApiSysteminsightsListAlfExceptionsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListAlfExceptionsRequest) Execute() ([]SystemInsightsAlfExceptions, *http.Response, error) {
	return r.ApiService.SysteminsightsListAlfExceptionsExecute(r)
}

/*
SysteminsightsListAlfExceptions List System Insights ALF Exceptions

Valid filter fields are `system_id` and `state`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListAlfExceptionsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAlfExceptions(ctx context.Context) ApiSysteminsightsListAlfExceptionsRequest {
	return ApiSysteminsightsListAlfExceptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAlfExceptions
func (a *SystemInsightsApiService) SysteminsightsListAlfExceptionsExecute(r ApiSysteminsightsListAlfExceptionsRequest) ([]SystemInsightsAlfExceptions, *http.Response, error) {
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

type ApiSysteminsightsListAlfExplicitAuthsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	filter *[]string
	skip *int32
	sort *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListAlfExplicitAuthsRequest) XOrgId(xOrgId string) ApiSysteminsightsListAlfExplicitAuthsRequest {
	r.xOrgId = &xOrgId
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListAlfExplicitAuthsRequest) Filter(filter []string) ApiSysteminsightsListAlfExplicitAuthsRequest {
	r.filter = &filter
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListAlfExplicitAuthsRequest) Skip(skip int32) ApiSysteminsightsListAlfExplicitAuthsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListAlfExplicitAuthsRequest) Sort(sort []string) ApiSysteminsightsListAlfExplicitAuthsRequest {
	r.sort = &sort
	return r
}

func (r ApiSysteminsightsListAlfExplicitAuthsRequest) Limit(limit int32) ApiSysteminsightsListAlfExplicitAuthsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListAlfExplicitAuthsRequest) Execute() ([]SystemInsightsAlfExplicitAuths, *http.Response, error) {
	return r.ApiService.SysteminsightsListAlfExplicitAuthsExecute(r)
}

/*
SysteminsightsListAlfExplicitAuths List System Insights ALF Explicit Authentications

Valid filter fields are `system_id` and `process`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListAlfExplicitAuthsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAlfExplicitAuths(ctx context.Context) ApiSysteminsightsListAlfExplicitAuthsRequest {
	return ApiSysteminsightsListAlfExplicitAuthsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAlfExplicitAuths
func (a *SystemInsightsApiService) SysteminsightsListAlfExplicitAuthsExecute(r ApiSysteminsightsListAlfExplicitAuthsRequest) ([]SystemInsightsAlfExplicitAuths, *http.Response, error) {
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

type ApiSysteminsightsListAppcompatShimsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListAppcompatShimsRequest) XOrgId(xOrgId string) ApiSysteminsightsListAppcompatShimsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListAppcompatShimsRequest) Skip(skip int32) ApiSysteminsightsListAppcompatShimsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListAppcompatShimsRequest) Sort(sort []string) ApiSysteminsightsListAppcompatShimsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListAppcompatShimsRequest) Filter(filter []string) ApiSysteminsightsListAppcompatShimsRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListAppcompatShimsRequest) Limit(limit int32) ApiSysteminsightsListAppcompatShimsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListAppcompatShimsRequest) Execute() ([]SystemInsightsAppcompatShims, *http.Response, error) {
	return r.ApiService.SysteminsightsListAppcompatShimsExecute(r)
}

/*
SysteminsightsListAppcompatShims List System Insights Application Compatibility Shims

Valid filter fields are `system_id` and `enabled`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListAppcompatShimsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAppcompatShims(ctx context.Context) ApiSysteminsightsListAppcompatShimsRequest {
	return ApiSysteminsightsListAppcompatShimsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAppcompatShims
func (a *SystemInsightsApiService) SysteminsightsListAppcompatShimsExecute(r ApiSysteminsightsListAppcompatShimsRequest) ([]SystemInsightsAppcompatShims, *http.Response, error) {
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

type ApiSysteminsightsListAppsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListAppsRequest) XOrgId(xOrgId string) ApiSysteminsightsListAppsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListAppsRequest) Skip(skip int32) ApiSysteminsightsListAppsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListAppsRequest) Sort(sort []string) ApiSysteminsightsListAppsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListAppsRequest) Filter(filter []string) ApiSysteminsightsListAppsRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListAppsRequest) Limit(limit int32) ApiSysteminsightsListAppsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListAppsRequest) Execute() ([]SystemInsightsApps, *http.Response, error) {
	return r.ApiService.SysteminsightsListAppsExecute(r)
}

/*
SysteminsightsListApps List System Insights Apps

Lists all apps for macOS devices. For Windows devices, use [List System Insights Programs](#operation/systeminsights_list_programs).

Valid filter fields are `system_id` and `bundle_name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListAppsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListApps(ctx context.Context) ApiSysteminsightsListAppsRequest {
	return ApiSysteminsightsListAppsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsApps
func (a *SystemInsightsApiService) SysteminsightsListAppsExecute(r ApiSysteminsightsListAppsRequest) ([]SystemInsightsApps, *http.Response, error) {
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

type ApiSysteminsightsListAuthorizedKeysRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListAuthorizedKeysRequest) XOrgId(xOrgId string) ApiSysteminsightsListAuthorizedKeysRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListAuthorizedKeysRequest) Skip(skip int32) ApiSysteminsightsListAuthorizedKeysRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListAuthorizedKeysRequest) Sort(sort []string) ApiSysteminsightsListAuthorizedKeysRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListAuthorizedKeysRequest) Filter(filter []string) ApiSysteminsightsListAuthorizedKeysRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListAuthorizedKeysRequest) Limit(limit int32) ApiSysteminsightsListAuthorizedKeysRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListAuthorizedKeysRequest) Execute() ([]SystemInsightsAuthorizedKeys, *http.Response, error) {
	return r.ApiService.SysteminsightsListAuthorizedKeysExecute(r)
}

/*
SysteminsightsListAuthorizedKeys List System Insights Authorized Keys

Valid filter fields are `system_id` and `uid`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListAuthorizedKeysRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAuthorizedKeys(ctx context.Context) ApiSysteminsightsListAuthorizedKeysRequest {
	return ApiSysteminsightsListAuthorizedKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAuthorizedKeys
func (a *SystemInsightsApiService) SysteminsightsListAuthorizedKeysExecute(r ApiSysteminsightsListAuthorizedKeysRequest) ([]SystemInsightsAuthorizedKeys, *http.Response, error) {
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

type ApiSysteminsightsListAzureInstanceMetadataRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListAzureInstanceMetadataRequest) Skip(skip int32) ApiSysteminsightsListAzureInstanceMetadataRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListAzureInstanceMetadataRequest) Sort(sort []string) ApiSysteminsightsListAzureInstanceMetadataRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListAzureInstanceMetadataRequest) Filter(filter []string) ApiSysteminsightsListAzureInstanceMetadataRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListAzureInstanceMetadataRequest) XOrgId(xOrgId string) ApiSysteminsightsListAzureInstanceMetadataRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListAzureInstanceMetadataRequest) Limit(limit int32) ApiSysteminsightsListAzureInstanceMetadataRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListAzureInstanceMetadataRequest) Execute() ([]SystemInsightsAzureInstanceMetadata, *http.Response, error) {
	return r.ApiService.SysteminsightsListAzureInstanceMetadataExecute(r)
}

/*
SysteminsightsListAzureInstanceMetadata List System Insights Azure Instance Metadata

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListAzureInstanceMetadataRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAzureInstanceMetadata(ctx context.Context) ApiSysteminsightsListAzureInstanceMetadataRequest {
	return ApiSysteminsightsListAzureInstanceMetadataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAzureInstanceMetadata
func (a *SystemInsightsApiService) SysteminsightsListAzureInstanceMetadataExecute(r ApiSysteminsightsListAzureInstanceMetadataRequest) ([]SystemInsightsAzureInstanceMetadata, *http.Response, error) {
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

type ApiSysteminsightsListAzureInstanceTagsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListAzureInstanceTagsRequest) Skip(skip int32) ApiSysteminsightsListAzureInstanceTagsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListAzureInstanceTagsRequest) Sort(sort []string) ApiSysteminsightsListAzureInstanceTagsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListAzureInstanceTagsRequest) Filter(filter []string) ApiSysteminsightsListAzureInstanceTagsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListAzureInstanceTagsRequest) XOrgId(xOrgId string) ApiSysteminsightsListAzureInstanceTagsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListAzureInstanceTagsRequest) Limit(limit int32) ApiSysteminsightsListAzureInstanceTagsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListAzureInstanceTagsRequest) Execute() ([]SystemInsightsAzureInstanceTags, *http.Response, error) {
	return r.ApiService.SysteminsightsListAzureInstanceTagsExecute(r)
}

/*
SysteminsightsListAzureInstanceTags List System Insights Azure Instance Tags

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListAzureInstanceTagsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListAzureInstanceTags(ctx context.Context) ApiSysteminsightsListAzureInstanceTagsRequest {
	return ApiSysteminsightsListAzureInstanceTagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsAzureInstanceTags
func (a *SystemInsightsApiService) SysteminsightsListAzureInstanceTagsExecute(r ApiSysteminsightsListAzureInstanceTagsRequest) ([]SystemInsightsAzureInstanceTags, *http.Response, error) {
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

type ApiSysteminsightsListBatteryRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListBatteryRequest) XOrgId(xOrgId string) ApiSysteminsightsListBatteryRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListBatteryRequest) Skip(skip int32) ApiSysteminsightsListBatteryRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListBatteryRequest) Sort(sort []string) ApiSysteminsightsListBatteryRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListBatteryRequest) Filter(filter []string) ApiSysteminsightsListBatteryRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListBatteryRequest) Limit(limit int32) ApiSysteminsightsListBatteryRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListBatteryRequest) Execute() ([]SystemInsightsBattery, *http.Response, error) {
	return r.ApiService.SysteminsightsListBatteryExecute(r)
}

/*
SysteminsightsListBattery List System Insights Battery

Valid filter fields are `system_id` and `health`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListBatteryRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListBattery(ctx context.Context) ApiSysteminsightsListBatteryRequest {
	return ApiSysteminsightsListBatteryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsBattery
func (a *SystemInsightsApiService) SysteminsightsListBatteryExecute(r ApiSysteminsightsListBatteryRequest) ([]SystemInsightsBattery, *http.Response, error) {
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

type ApiSysteminsightsListBitlockerInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListBitlockerInfoRequest) Skip(skip int32) ApiSysteminsightsListBitlockerInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListBitlockerInfoRequest) Sort(sort []string) ApiSysteminsightsListBitlockerInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListBitlockerInfoRequest) Filter(filter []string) ApiSysteminsightsListBitlockerInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListBitlockerInfoRequest) XOrgId(xOrgId string) ApiSysteminsightsListBitlockerInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListBitlockerInfoRequest) Limit(limit int32) ApiSysteminsightsListBitlockerInfoRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListBitlockerInfoRequest) Execute() ([]SystemInsightsBitlockerInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListBitlockerInfoExecute(r)
}

/*
SysteminsightsListBitlockerInfo List System Insights Bitlocker Info

Valid filter fields are `system_id` and `protection_status`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListBitlockerInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListBitlockerInfo(ctx context.Context) ApiSysteminsightsListBitlockerInfoRequest {
	return ApiSysteminsightsListBitlockerInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsBitlockerInfo
func (a *SystemInsightsApiService) SysteminsightsListBitlockerInfoExecute(r ApiSysteminsightsListBitlockerInfoRequest) ([]SystemInsightsBitlockerInfo, *http.Response, error) {
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

type ApiSysteminsightsListBrowserPluginsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListBrowserPluginsRequest) Skip(skip int32) ApiSysteminsightsListBrowserPluginsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListBrowserPluginsRequest) Sort(sort []string) ApiSysteminsightsListBrowserPluginsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListBrowserPluginsRequest) Filter(filter []string) ApiSysteminsightsListBrowserPluginsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListBrowserPluginsRequest) XOrgId(xOrgId string) ApiSysteminsightsListBrowserPluginsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListBrowserPluginsRequest) Limit(limit int32) ApiSysteminsightsListBrowserPluginsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListBrowserPluginsRequest) Execute() ([]SystemInsightsBrowserPlugins, *http.Response, error) {
	return r.ApiService.SysteminsightsListBrowserPluginsExecute(r)
}

/*
SysteminsightsListBrowserPlugins List System Insights Browser Plugins

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListBrowserPluginsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListBrowserPlugins(ctx context.Context) ApiSysteminsightsListBrowserPluginsRequest {
	return ApiSysteminsightsListBrowserPluginsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsBrowserPlugins
func (a *SystemInsightsApiService) SysteminsightsListBrowserPluginsExecute(r ApiSysteminsightsListBrowserPluginsRequest) ([]SystemInsightsBrowserPlugins, *http.Response, error) {
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

type ApiSysteminsightsListCertificatesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListCertificatesRequest) Skip(skip int32) ApiSysteminsightsListCertificatesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListCertificatesRequest) Sort(sort []string) ApiSysteminsightsListCertificatesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; Note: You can only filter by &#x60;system_id&#x60; and &#x60;common_name&#x60; 
func (r ApiSysteminsightsListCertificatesRequest) Filter(filter []string) ApiSysteminsightsListCertificatesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListCertificatesRequest) XOrgId(xOrgId string) ApiSysteminsightsListCertificatesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListCertificatesRequest) Limit(limit int32) ApiSysteminsightsListCertificatesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListCertificatesRequest) Execute() ([]SystemInsightsCertificates, *http.Response, error) {
	return r.ApiService.SysteminsightsListCertificatesExecute(r)
}

/*
SysteminsightsListCertificates List System Insights Certificates

Valid filter fields are `system_id` and `common_name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListCertificatesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListCertificates(ctx context.Context) ApiSysteminsightsListCertificatesRequest {
	return ApiSysteminsightsListCertificatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsCertificates
func (a *SystemInsightsApiService) SysteminsightsListCertificatesExecute(r ApiSysteminsightsListCertificatesRequest) ([]SystemInsightsCertificates, *http.Response, error) {
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

type ApiSysteminsightsListChassisInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListChassisInfoRequest) Skip(skip int32) ApiSysteminsightsListChassisInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListChassisInfoRequest) Sort(sort []string) ApiSysteminsightsListChassisInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListChassisInfoRequest) Filter(filter []string) ApiSysteminsightsListChassisInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListChassisInfoRequest) XOrgId(xOrgId string) ApiSysteminsightsListChassisInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListChassisInfoRequest) Limit(limit int32) ApiSysteminsightsListChassisInfoRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListChassisInfoRequest) Execute() ([]SystemInsightsChassisInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListChassisInfoExecute(r)
}

/*
SysteminsightsListChassisInfo List System Insights Chassis Info

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListChassisInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListChassisInfo(ctx context.Context) ApiSysteminsightsListChassisInfoRequest {
	return ApiSysteminsightsListChassisInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsChassisInfo
func (a *SystemInsightsApiService) SysteminsightsListChassisInfoExecute(r ApiSysteminsightsListChassisInfoRequest) ([]SystemInsightsChassisInfo, *http.Response, error) {
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

type ApiSysteminsightsListChromeExtensionsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListChromeExtensionsRequest) XOrgId(xOrgId string) ApiSysteminsightsListChromeExtensionsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListChromeExtensionsRequest) Skip(skip int32) ApiSysteminsightsListChromeExtensionsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListChromeExtensionsRequest) Sort(sort []string) ApiSysteminsightsListChromeExtensionsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListChromeExtensionsRequest) Filter(filter []string) ApiSysteminsightsListChromeExtensionsRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListChromeExtensionsRequest) Limit(limit int32) ApiSysteminsightsListChromeExtensionsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListChromeExtensionsRequest) Execute() ([]SystemInsightsChromeExtensions, *http.Response, error) {
	return r.ApiService.SysteminsightsListChromeExtensionsExecute(r)
}

/*
SysteminsightsListChromeExtensions List System Insights Chrome Extensions

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListChromeExtensionsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListChromeExtensions(ctx context.Context) ApiSysteminsightsListChromeExtensionsRequest {
	return ApiSysteminsightsListChromeExtensionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsChromeExtensions
func (a *SystemInsightsApiService) SysteminsightsListChromeExtensionsExecute(r ApiSysteminsightsListChromeExtensionsRequest) ([]SystemInsightsChromeExtensions, *http.Response, error) {
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

type ApiSysteminsightsListConnectivityRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListConnectivityRequest) XOrgId(xOrgId string) ApiSysteminsightsListConnectivityRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListConnectivityRequest) Skip(skip int32) ApiSysteminsightsListConnectivityRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListConnectivityRequest) Sort(sort []string) ApiSysteminsightsListConnectivityRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListConnectivityRequest) Filter(filter []string) ApiSysteminsightsListConnectivityRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListConnectivityRequest) Limit(limit int32) ApiSysteminsightsListConnectivityRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListConnectivityRequest) Execute() ([]SystemInsightsConnectivity, *http.Response, error) {
	return r.ApiService.SysteminsightsListConnectivityExecute(r)
}

/*
SysteminsightsListConnectivity List System Insights Connectivity

The only valid filter field is `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListConnectivityRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListConnectivity(ctx context.Context) ApiSysteminsightsListConnectivityRequest {
	return ApiSysteminsightsListConnectivityRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsConnectivity
func (a *SystemInsightsApiService) SysteminsightsListConnectivityExecute(r ApiSysteminsightsListConnectivityRequest) ([]SystemInsightsConnectivity, *http.Response, error) {
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

type ApiSysteminsightsListCrashesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListCrashesRequest) XOrgId(xOrgId string) ApiSysteminsightsListCrashesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListCrashesRequest) Skip(skip int32) ApiSysteminsightsListCrashesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListCrashesRequest) Sort(sort []string) ApiSysteminsightsListCrashesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListCrashesRequest) Filter(filter []string) ApiSysteminsightsListCrashesRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListCrashesRequest) Limit(limit int32) ApiSysteminsightsListCrashesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListCrashesRequest) Execute() ([]SystemInsightsCrashes, *http.Response, error) {
	return r.ApiService.SysteminsightsListCrashesExecute(r)
}

/*
SysteminsightsListCrashes List System Insights Crashes

Valid filter fields are `system_id` and `identifier`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListCrashesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListCrashes(ctx context.Context) ApiSysteminsightsListCrashesRequest {
	return ApiSysteminsightsListCrashesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsCrashes
func (a *SystemInsightsApiService) SysteminsightsListCrashesExecute(r ApiSysteminsightsListCrashesRequest) ([]SystemInsightsCrashes, *http.Response, error) {
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

type ApiSysteminsightsListCupsDestinationsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListCupsDestinationsRequest) Skip(skip int32) ApiSysteminsightsListCupsDestinationsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListCupsDestinationsRequest) Sort(sort []string) ApiSysteminsightsListCupsDestinationsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListCupsDestinationsRequest) Filter(filter []string) ApiSysteminsightsListCupsDestinationsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListCupsDestinationsRequest) XOrgId(xOrgId string) ApiSysteminsightsListCupsDestinationsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListCupsDestinationsRequest) Limit(limit int32) ApiSysteminsightsListCupsDestinationsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListCupsDestinationsRequest) Execute() ([]SystemInsightsCupsDestinations, *http.Response, error) {
	return r.ApiService.SysteminsightsListCupsDestinationsExecute(r)
}

/*
SysteminsightsListCupsDestinations List System Insights CUPS Destinations

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListCupsDestinationsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListCupsDestinations(ctx context.Context) ApiSysteminsightsListCupsDestinationsRequest {
	return ApiSysteminsightsListCupsDestinationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsCupsDestinations
func (a *SystemInsightsApiService) SysteminsightsListCupsDestinationsExecute(r ApiSysteminsightsListCupsDestinationsRequest) ([]SystemInsightsCupsDestinations, *http.Response, error) {
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

type ApiSysteminsightsListDiskEncryptionRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListDiskEncryptionRequest) Skip(skip int32) ApiSysteminsightsListDiskEncryptionRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListDiskEncryptionRequest) Sort(sort []string) ApiSysteminsightsListDiskEncryptionRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListDiskEncryptionRequest) Filter(filter []string) ApiSysteminsightsListDiskEncryptionRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListDiskEncryptionRequest) XOrgId(xOrgId string) ApiSysteminsightsListDiskEncryptionRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListDiskEncryptionRequest) Limit(limit int32) ApiSysteminsightsListDiskEncryptionRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListDiskEncryptionRequest) Execute() ([]SystemInsightsDiskEncryption, *http.Response, error) {
	return r.ApiService.SysteminsightsListDiskEncryptionExecute(r)
}

/*
SysteminsightsListDiskEncryption List System Insights Disk Encryption

Valid filter fields are `system_id` and `encryption_status`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListDiskEncryptionRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListDiskEncryption(ctx context.Context) ApiSysteminsightsListDiskEncryptionRequest {
	return ApiSysteminsightsListDiskEncryptionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsDiskEncryption
func (a *SystemInsightsApiService) SysteminsightsListDiskEncryptionExecute(r ApiSysteminsightsListDiskEncryptionRequest) ([]SystemInsightsDiskEncryption, *http.Response, error) {
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

type ApiSysteminsightsListDiskInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListDiskInfoRequest) Skip(skip int32) ApiSysteminsightsListDiskInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListDiskInfoRequest) Sort(sort []string) ApiSysteminsightsListDiskInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListDiskInfoRequest) Filter(filter []string) ApiSysteminsightsListDiskInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListDiskInfoRequest) XOrgId(xOrgId string) ApiSysteminsightsListDiskInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListDiskInfoRequest) Limit(limit int32) ApiSysteminsightsListDiskInfoRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListDiskInfoRequest) Execute() ([]SystemInsightsDiskInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListDiskInfoExecute(r)
}

/*
SysteminsightsListDiskInfo List System Insights Disk Info

Valid filter fields are `system_id` and `disk_index`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListDiskInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListDiskInfo(ctx context.Context) ApiSysteminsightsListDiskInfoRequest {
	return ApiSysteminsightsListDiskInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsDiskInfo
func (a *SystemInsightsApiService) SysteminsightsListDiskInfoExecute(r ApiSysteminsightsListDiskInfoRequest) ([]SystemInsightsDiskInfo, *http.Response, error) {
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

type ApiSysteminsightsListDnsResolversRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListDnsResolversRequest) Skip(skip int32) ApiSysteminsightsListDnsResolversRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListDnsResolversRequest) Sort(sort []string) ApiSysteminsightsListDnsResolversRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListDnsResolversRequest) Filter(filter []string) ApiSysteminsightsListDnsResolversRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListDnsResolversRequest) XOrgId(xOrgId string) ApiSysteminsightsListDnsResolversRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListDnsResolversRequest) Limit(limit int32) ApiSysteminsightsListDnsResolversRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListDnsResolversRequest) Execute() ([]SystemInsightsDnsResolvers, *http.Response, error) {
	return r.ApiService.SysteminsightsListDnsResolversExecute(r)
}

/*
SysteminsightsListDnsResolvers List System Insights DNS Resolvers

Valid filter fields are `system_id` and `type`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListDnsResolversRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListDnsResolvers(ctx context.Context) ApiSysteminsightsListDnsResolversRequest {
	return ApiSysteminsightsListDnsResolversRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsDnsResolvers
func (a *SystemInsightsApiService) SysteminsightsListDnsResolversExecute(r ApiSysteminsightsListDnsResolversRequest) ([]SystemInsightsDnsResolvers, *http.Response, error) {
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

type ApiSysteminsightsListEtcHostsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListEtcHostsRequest) Skip(skip int32) ApiSysteminsightsListEtcHostsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListEtcHostsRequest) Sort(sort []string) ApiSysteminsightsListEtcHostsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListEtcHostsRequest) Filter(filter []string) ApiSysteminsightsListEtcHostsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListEtcHostsRequest) XOrgId(xOrgId string) ApiSysteminsightsListEtcHostsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListEtcHostsRequest) Limit(limit int32) ApiSysteminsightsListEtcHostsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListEtcHostsRequest) Execute() ([]SystemInsightsEtcHosts, *http.Response, error) {
	return r.ApiService.SysteminsightsListEtcHostsExecute(r)
}

/*
SysteminsightsListEtcHosts List System Insights Etc Hosts

Valid filter fields are `system_id` and `address`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListEtcHostsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListEtcHosts(ctx context.Context) ApiSysteminsightsListEtcHostsRequest {
	return ApiSysteminsightsListEtcHostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsEtcHosts
func (a *SystemInsightsApiService) SysteminsightsListEtcHostsExecute(r ApiSysteminsightsListEtcHostsRequest) ([]SystemInsightsEtcHosts, *http.Response, error) {
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

type ApiSysteminsightsListFirefoxAddonsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListFirefoxAddonsRequest) Skip(skip int32) ApiSysteminsightsListFirefoxAddonsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListFirefoxAddonsRequest) Sort(sort []string) ApiSysteminsightsListFirefoxAddonsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListFirefoxAddonsRequest) Filter(filter []string) ApiSysteminsightsListFirefoxAddonsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListFirefoxAddonsRequest) XOrgId(xOrgId string) ApiSysteminsightsListFirefoxAddonsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListFirefoxAddonsRequest) Limit(limit int32) ApiSysteminsightsListFirefoxAddonsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListFirefoxAddonsRequest) Execute() ([]SystemInsightsFirefoxAddons, *http.Response, error) {
	return r.ApiService.SysteminsightsListFirefoxAddonsExecute(r)
}

/*
SysteminsightsListFirefoxAddons List System Insights Firefox Addons

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListFirefoxAddonsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListFirefoxAddons(ctx context.Context) ApiSysteminsightsListFirefoxAddonsRequest {
	return ApiSysteminsightsListFirefoxAddonsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsFirefoxAddons
func (a *SystemInsightsApiService) SysteminsightsListFirefoxAddonsExecute(r ApiSysteminsightsListFirefoxAddonsRequest) ([]SystemInsightsFirefoxAddons, *http.Response, error) {
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

type ApiSysteminsightsListGroupsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListGroupsRequest) Skip(skip int32) ApiSysteminsightsListGroupsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListGroupsRequest) Sort(sort []string) ApiSysteminsightsListGroupsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListGroupsRequest) Filter(filter []string) ApiSysteminsightsListGroupsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListGroupsRequest) XOrgId(xOrgId string) ApiSysteminsightsListGroupsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListGroupsRequest) Limit(limit int32) ApiSysteminsightsListGroupsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListGroupsRequest) Execute() ([]SystemInsightsGroups, *http.Response, error) {
	return r.ApiService.SysteminsightsListGroupsExecute(r)
}

/*
SysteminsightsListGroups List System Insights Groups

Valid filter fields are `system_id` and `groupname`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListGroupsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListGroups(ctx context.Context) ApiSysteminsightsListGroupsRequest {
	return ApiSysteminsightsListGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsGroups
func (a *SystemInsightsApiService) SysteminsightsListGroupsExecute(r ApiSysteminsightsListGroupsRequest) ([]SystemInsightsGroups, *http.Response, error) {
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

type ApiSysteminsightsListIeExtensionsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListIeExtensionsRequest) XOrgId(xOrgId string) ApiSysteminsightsListIeExtensionsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListIeExtensionsRequest) Skip(skip int32) ApiSysteminsightsListIeExtensionsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListIeExtensionsRequest) Sort(sort []string) ApiSysteminsightsListIeExtensionsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListIeExtensionsRequest) Filter(filter []string) ApiSysteminsightsListIeExtensionsRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListIeExtensionsRequest) Limit(limit int32) ApiSysteminsightsListIeExtensionsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListIeExtensionsRequest) Execute() ([]SystemInsightsIeExtensions, *http.Response, error) {
	return r.ApiService.SysteminsightsListIeExtensionsExecute(r)
}

/*
SysteminsightsListIeExtensions List System Insights IE Extensions

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListIeExtensionsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListIeExtensions(ctx context.Context) ApiSysteminsightsListIeExtensionsRequest {
	return ApiSysteminsightsListIeExtensionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsIeExtensions
func (a *SystemInsightsApiService) SysteminsightsListIeExtensionsExecute(r ApiSysteminsightsListIeExtensionsRequest) ([]SystemInsightsIeExtensions, *http.Response, error) {
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

type ApiSysteminsightsListInterfaceAddressesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListInterfaceAddressesRequest) Skip(skip int32) ApiSysteminsightsListInterfaceAddressesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListInterfaceAddressesRequest) Sort(sort []string) ApiSysteminsightsListInterfaceAddressesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListInterfaceAddressesRequest) Filter(filter []string) ApiSysteminsightsListInterfaceAddressesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListInterfaceAddressesRequest) XOrgId(xOrgId string) ApiSysteminsightsListInterfaceAddressesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListInterfaceAddressesRequest) Limit(limit int32) ApiSysteminsightsListInterfaceAddressesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListInterfaceAddressesRequest) Execute() ([]SystemInsightsInterfaceAddresses, *http.Response, error) {
	return r.ApiService.SysteminsightsListInterfaceAddressesExecute(r)
}

/*
SysteminsightsListInterfaceAddresses List System Insights Interface Addresses

Valid filter fields are `system_id` and `address`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListInterfaceAddressesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListInterfaceAddresses(ctx context.Context) ApiSysteminsightsListInterfaceAddressesRequest {
	return ApiSysteminsightsListInterfaceAddressesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsInterfaceAddresses
func (a *SystemInsightsApiService) SysteminsightsListInterfaceAddressesExecute(r ApiSysteminsightsListInterfaceAddressesRequest) ([]SystemInsightsInterfaceAddresses, *http.Response, error) {
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

type ApiSysteminsightsListInterfaceDetailsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListInterfaceDetailsRequest) Skip(skip int32) ApiSysteminsightsListInterfaceDetailsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListInterfaceDetailsRequest) Sort(sort []string) ApiSysteminsightsListInterfaceDetailsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListInterfaceDetailsRequest) Filter(filter []string) ApiSysteminsightsListInterfaceDetailsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListInterfaceDetailsRequest) XOrgId(xOrgId string) ApiSysteminsightsListInterfaceDetailsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListInterfaceDetailsRequest) Limit(limit int32) ApiSysteminsightsListInterfaceDetailsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListInterfaceDetailsRequest) Execute() ([]SystemInsightsInterfaceDetails, *http.Response, error) {
	return r.ApiService.SysteminsightsListInterfaceDetailsExecute(r)
}

/*
SysteminsightsListInterfaceDetails List System Insights Interface Details

Valid filter fields are `system_id` and `interface`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListInterfaceDetailsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListInterfaceDetails(ctx context.Context) ApiSysteminsightsListInterfaceDetailsRequest {
	return ApiSysteminsightsListInterfaceDetailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsInterfaceDetails
func (a *SystemInsightsApiService) SysteminsightsListInterfaceDetailsExecute(r ApiSysteminsightsListInterfaceDetailsRequest) ([]SystemInsightsInterfaceDetails, *http.Response, error) {
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

type ApiSysteminsightsListKernelInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListKernelInfoRequest) Skip(skip int32) ApiSysteminsightsListKernelInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListKernelInfoRequest) Sort(sort []string) ApiSysteminsightsListKernelInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListKernelInfoRequest) Filter(filter []string) ApiSysteminsightsListKernelInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListKernelInfoRequest) XOrgId(xOrgId string) ApiSysteminsightsListKernelInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListKernelInfoRequest) Limit(limit int32) ApiSysteminsightsListKernelInfoRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListKernelInfoRequest) Execute() ([]SystemInsightsKernelInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListKernelInfoExecute(r)
}

/*
SysteminsightsListKernelInfo List System Insights Kernel Info

Valid filter fields are `system_id` and `version`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListKernelInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListKernelInfo(ctx context.Context) ApiSysteminsightsListKernelInfoRequest {
	return ApiSysteminsightsListKernelInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsKernelInfo
func (a *SystemInsightsApiService) SysteminsightsListKernelInfoExecute(r ApiSysteminsightsListKernelInfoRequest) ([]SystemInsightsKernelInfo, *http.Response, error) {
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

type ApiSysteminsightsListLaunchdRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListLaunchdRequest) XOrgId(xOrgId string) ApiSysteminsightsListLaunchdRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListLaunchdRequest) Skip(skip int32) ApiSysteminsightsListLaunchdRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListLaunchdRequest) Sort(sort []string) ApiSysteminsightsListLaunchdRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListLaunchdRequest) Filter(filter []string) ApiSysteminsightsListLaunchdRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListLaunchdRequest) Limit(limit int32) ApiSysteminsightsListLaunchdRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListLaunchdRequest) Execute() ([]SystemInsightsLaunchd, *http.Response, error) {
	return r.ApiService.SysteminsightsListLaunchdExecute(r)
}

/*
SysteminsightsListLaunchd List System Insights Launchd

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListLaunchdRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListLaunchd(ctx context.Context) ApiSysteminsightsListLaunchdRequest {
	return ApiSysteminsightsListLaunchdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsLaunchd
func (a *SystemInsightsApiService) SysteminsightsListLaunchdExecute(r ApiSysteminsightsListLaunchdRequest) ([]SystemInsightsLaunchd, *http.Response, error) {
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

type ApiSysteminsightsListLinuxPackagesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListLinuxPackagesRequest) Skip(skip int32) ApiSysteminsightsListLinuxPackagesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListLinuxPackagesRequest) Sort(sort []string) ApiSysteminsightsListLinuxPackagesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListLinuxPackagesRequest) Filter(filter []string) ApiSysteminsightsListLinuxPackagesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListLinuxPackagesRequest) XOrgId(xOrgId string) ApiSysteminsightsListLinuxPackagesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListLinuxPackagesRequest) Limit(limit int32) ApiSysteminsightsListLinuxPackagesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListLinuxPackagesRequest) Execute() ([]SystemInsightsLinuxPackages, *http.Response, error) {
	return r.ApiService.SysteminsightsListLinuxPackagesExecute(r)
}

/*
SysteminsightsListLinuxPackages List System Insights Linux Packages

Lists all programs for Linux devices. For macOS devices, use [List System Insights System Apps](#operation/systeminsights_list_apps). For windows devices, use [List System Insights System Apps](#operation/systeminsights_list_programs).

Valid filter fields are `name` and `package_format`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListLinuxPackagesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListLinuxPackages(ctx context.Context) ApiSysteminsightsListLinuxPackagesRequest {
	return ApiSysteminsightsListLinuxPackagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsLinuxPackages
func (a *SystemInsightsApiService) SysteminsightsListLinuxPackagesExecute(r ApiSysteminsightsListLinuxPackagesRequest) ([]SystemInsightsLinuxPackages, *http.Response, error) {
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

type ApiSysteminsightsListLoggedInUsersRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListLoggedInUsersRequest) XOrgId(xOrgId string) ApiSysteminsightsListLoggedInUsersRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListLoggedInUsersRequest) Skip(skip int32) ApiSysteminsightsListLoggedInUsersRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListLoggedInUsersRequest) Sort(sort []string) ApiSysteminsightsListLoggedInUsersRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListLoggedInUsersRequest) Filter(filter []string) ApiSysteminsightsListLoggedInUsersRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListLoggedInUsersRequest) Limit(limit int32) ApiSysteminsightsListLoggedInUsersRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListLoggedInUsersRequest) Execute() ([]SystemInsightsLoggedInUsers, *http.Response, error) {
	return r.ApiService.SysteminsightsListLoggedInUsersExecute(r)
}

/*
SysteminsightsListLoggedInUsers List System Insights Logged-In Users

Valid filter fields are `system_id` and `user`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListLoggedInUsersRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListLoggedInUsers(ctx context.Context) ApiSysteminsightsListLoggedInUsersRequest {
	return ApiSysteminsightsListLoggedInUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsLoggedInUsers
func (a *SystemInsightsApiService) SysteminsightsListLoggedInUsersExecute(r ApiSysteminsightsListLoggedInUsersRequest) ([]SystemInsightsLoggedInUsers, *http.Response, error) {
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

type ApiSysteminsightsListLogicalDrivesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListLogicalDrivesRequest) Skip(skip int32) ApiSysteminsightsListLogicalDrivesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListLogicalDrivesRequest) Sort(sort []string) ApiSysteminsightsListLogicalDrivesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListLogicalDrivesRequest) Filter(filter []string) ApiSysteminsightsListLogicalDrivesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListLogicalDrivesRequest) XOrgId(xOrgId string) ApiSysteminsightsListLogicalDrivesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListLogicalDrivesRequest) Limit(limit int32) ApiSysteminsightsListLogicalDrivesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListLogicalDrivesRequest) Execute() ([]SystemInsightsLogicalDrives, *http.Response, error) {
	return r.ApiService.SysteminsightsListLogicalDrivesExecute(r)
}

/*
SysteminsightsListLogicalDrives List System Insights Logical Drives

Valid filter fields are `system_id` and `device_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListLogicalDrivesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListLogicalDrives(ctx context.Context) ApiSysteminsightsListLogicalDrivesRequest {
	return ApiSysteminsightsListLogicalDrivesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsLogicalDrives
func (a *SystemInsightsApiService) SysteminsightsListLogicalDrivesExecute(r ApiSysteminsightsListLogicalDrivesRequest) ([]SystemInsightsLogicalDrives, *http.Response, error) {
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

type ApiSysteminsightsListManagedPoliciesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListManagedPoliciesRequest) XOrgId(xOrgId string) ApiSysteminsightsListManagedPoliciesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListManagedPoliciesRequest) Skip(skip int32) ApiSysteminsightsListManagedPoliciesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListManagedPoliciesRequest) Sort(sort []string) ApiSysteminsightsListManagedPoliciesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListManagedPoliciesRequest) Filter(filter []string) ApiSysteminsightsListManagedPoliciesRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListManagedPoliciesRequest) Limit(limit int32) ApiSysteminsightsListManagedPoliciesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListManagedPoliciesRequest) Execute() ([]SystemInsightsManagedPolicies, *http.Response, error) {
	return r.ApiService.SysteminsightsListManagedPoliciesExecute(r)
}

/*
SysteminsightsListManagedPolicies List System Insights Managed Policies

Valid filter fields are `system_id` and `domain`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListManagedPoliciesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListManagedPolicies(ctx context.Context) ApiSysteminsightsListManagedPoliciesRequest {
	return ApiSysteminsightsListManagedPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsManagedPolicies
func (a *SystemInsightsApiService) SysteminsightsListManagedPoliciesExecute(r ApiSysteminsightsListManagedPoliciesRequest) ([]SystemInsightsManagedPolicies, *http.Response, error) {
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

type ApiSysteminsightsListMountsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListMountsRequest) Skip(skip int32) ApiSysteminsightsListMountsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListMountsRequest) Sort(sort []string) ApiSysteminsightsListMountsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListMountsRequest) Filter(filter []string) ApiSysteminsightsListMountsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListMountsRequest) XOrgId(xOrgId string) ApiSysteminsightsListMountsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListMountsRequest) Limit(limit int32) ApiSysteminsightsListMountsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListMountsRequest) Execute() ([]SystemInsightsMounts, *http.Response, error) {
	return r.ApiService.SysteminsightsListMountsExecute(r)
}

/*
SysteminsightsListMounts List System Insights Mounts

Valid filter fields are `system_id` and `path`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListMountsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListMounts(ctx context.Context) ApiSysteminsightsListMountsRequest {
	return ApiSysteminsightsListMountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsMounts
func (a *SystemInsightsApiService) SysteminsightsListMountsExecute(r ApiSysteminsightsListMountsRequest) ([]SystemInsightsMounts, *http.Response, error) {
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

type ApiSysteminsightsListOsVersionRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListOsVersionRequest) Skip(skip int32) ApiSysteminsightsListOsVersionRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListOsVersionRequest) Sort(sort []string) ApiSysteminsightsListOsVersionRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListOsVersionRequest) Filter(filter []string) ApiSysteminsightsListOsVersionRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListOsVersionRequest) XOrgId(xOrgId string) ApiSysteminsightsListOsVersionRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListOsVersionRequest) Limit(limit int32) ApiSysteminsightsListOsVersionRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListOsVersionRequest) Execute() ([]SystemInsightsOsVersion, *http.Response, error) {
	return r.ApiService.SysteminsightsListOsVersionExecute(r)
}

/*
SysteminsightsListOsVersion List System Insights OS Version

Valid filter fields are `system_id` and `version`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListOsVersionRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListOsVersion(ctx context.Context) ApiSysteminsightsListOsVersionRequest {
	return ApiSysteminsightsListOsVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsOsVersion
func (a *SystemInsightsApiService) SysteminsightsListOsVersionExecute(r ApiSysteminsightsListOsVersionRequest) ([]SystemInsightsOsVersion, *http.Response, error) {
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

type ApiSysteminsightsListPatchesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListPatchesRequest) Skip(skip int32) ApiSysteminsightsListPatchesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListPatchesRequest) Sort(sort []string) ApiSysteminsightsListPatchesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListPatchesRequest) Filter(filter []string) ApiSysteminsightsListPatchesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListPatchesRequest) XOrgId(xOrgId string) ApiSysteminsightsListPatchesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListPatchesRequest) Limit(limit int32) ApiSysteminsightsListPatchesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListPatchesRequest) Execute() ([]SystemInsightsPatches, *http.Response, error) {
	return r.ApiService.SysteminsightsListPatchesExecute(r)
}

/*
SysteminsightsListPatches List System Insights Patches

Valid filter fields are `system_id` and `hotfix_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListPatchesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListPatches(ctx context.Context) ApiSysteminsightsListPatchesRequest {
	return ApiSysteminsightsListPatchesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsPatches
func (a *SystemInsightsApiService) SysteminsightsListPatchesExecute(r ApiSysteminsightsListPatchesRequest) ([]SystemInsightsPatches, *http.Response, error) {
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

type ApiSysteminsightsListProgramsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListProgramsRequest) Skip(skip int32) ApiSysteminsightsListProgramsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListProgramsRequest) Sort(sort []string) ApiSysteminsightsListProgramsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListProgramsRequest) Filter(filter []string) ApiSysteminsightsListProgramsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListProgramsRequest) XOrgId(xOrgId string) ApiSysteminsightsListProgramsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListProgramsRequest) Limit(limit int32) ApiSysteminsightsListProgramsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListProgramsRequest) Execute() ([]SystemInsightsPrograms, *http.Response, error) {
	return r.ApiService.SysteminsightsListProgramsExecute(r)
}

/*
SysteminsightsListPrograms List System Insights Programs

Lists all programs for Windows devices. For macOS devices, use [List System Insights Apps](#operation/systeminsights_list_apps).

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListProgramsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListPrograms(ctx context.Context) ApiSysteminsightsListProgramsRequest {
	return ApiSysteminsightsListProgramsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsPrograms
func (a *SystemInsightsApiService) SysteminsightsListProgramsExecute(r ApiSysteminsightsListProgramsRequest) ([]SystemInsightsPrograms, *http.Response, error) {
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

type ApiSysteminsightsListPythonPackagesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListPythonPackagesRequest) Skip(skip int32) ApiSysteminsightsListPythonPackagesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListPythonPackagesRequest) Sort(sort []string) ApiSysteminsightsListPythonPackagesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListPythonPackagesRequest) Filter(filter []string) ApiSysteminsightsListPythonPackagesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListPythonPackagesRequest) XOrgId(xOrgId string) ApiSysteminsightsListPythonPackagesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListPythonPackagesRequest) Limit(limit int32) ApiSysteminsightsListPythonPackagesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListPythonPackagesRequest) Execute() ([]SystemInsightsPythonPackages, *http.Response, error) {
	return r.ApiService.SysteminsightsListPythonPackagesExecute(r)
}

/*
SysteminsightsListPythonPackages List System Insights Python Packages

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListPythonPackagesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListPythonPackages(ctx context.Context) ApiSysteminsightsListPythonPackagesRequest {
	return ApiSysteminsightsListPythonPackagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsPythonPackages
func (a *SystemInsightsApiService) SysteminsightsListPythonPackagesExecute(r ApiSysteminsightsListPythonPackagesRequest) ([]SystemInsightsPythonPackages, *http.Response, error) {
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

type ApiSysteminsightsListSafariExtensionsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListSafariExtensionsRequest) Skip(skip int32) ApiSysteminsightsListSafariExtensionsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListSafariExtensionsRequest) Sort(sort []string) ApiSysteminsightsListSafariExtensionsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListSafariExtensionsRequest) Filter(filter []string) ApiSysteminsightsListSafariExtensionsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListSafariExtensionsRequest) XOrgId(xOrgId string) ApiSysteminsightsListSafariExtensionsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListSafariExtensionsRequest) Limit(limit int32) ApiSysteminsightsListSafariExtensionsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListSafariExtensionsRequest) Execute() ([]SystemInsightsSafariExtensions, *http.Response, error) {
	return r.ApiService.SysteminsightsListSafariExtensionsExecute(r)
}

/*
SysteminsightsListSafariExtensions List System Insights Safari Extensions

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListSafariExtensionsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSafariExtensions(ctx context.Context) ApiSysteminsightsListSafariExtensionsRequest {
	return ApiSysteminsightsListSafariExtensionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSafariExtensions
func (a *SystemInsightsApiService) SysteminsightsListSafariExtensionsExecute(r ApiSysteminsightsListSafariExtensionsRequest) ([]SystemInsightsSafariExtensions, *http.Response, error) {
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

type ApiSysteminsightsListScheduledTasksRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListScheduledTasksRequest) Skip(skip int32) ApiSysteminsightsListScheduledTasksRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListScheduledTasksRequest) Sort(sort []string) ApiSysteminsightsListScheduledTasksRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListScheduledTasksRequest) Filter(filter []string) ApiSysteminsightsListScheduledTasksRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListScheduledTasksRequest) XOrgId(xOrgId string) ApiSysteminsightsListScheduledTasksRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListScheduledTasksRequest) Limit(limit int32) ApiSysteminsightsListScheduledTasksRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListScheduledTasksRequest) Execute() ([]SystemInsightsScheduledTasks, *http.Response, error) {
	return r.ApiService.SysteminsightsListScheduledTasksExecute(r)
}

/*
SysteminsightsListScheduledTasks List System Insights Scheduled Tasks

Valid filter fields are `system_id` and `enabled`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListScheduledTasksRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListScheduledTasks(ctx context.Context) ApiSysteminsightsListScheduledTasksRequest {
	return ApiSysteminsightsListScheduledTasksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsScheduledTasks
func (a *SystemInsightsApiService) SysteminsightsListScheduledTasksExecute(r ApiSysteminsightsListScheduledTasksRequest) ([]SystemInsightsScheduledTasks, *http.Response, error) {
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

type ApiSysteminsightsListSecurebootRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListSecurebootRequest) Skip(skip int32) ApiSysteminsightsListSecurebootRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListSecurebootRequest) Sort(sort []string) ApiSysteminsightsListSecurebootRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListSecurebootRequest) Filter(filter []string) ApiSysteminsightsListSecurebootRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListSecurebootRequest) XOrgId(xOrgId string) ApiSysteminsightsListSecurebootRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListSecurebootRequest) Limit(limit int32) ApiSysteminsightsListSecurebootRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListSecurebootRequest) Execute() ([]SystemInsightsSecureboot, *http.Response, error) {
	return r.ApiService.SysteminsightsListSecurebootExecute(r)
}

/*
SysteminsightsListSecureboot List System Insights Secure Boot

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListSecurebootRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSecureboot(ctx context.Context) ApiSysteminsightsListSecurebootRequest {
	return ApiSysteminsightsListSecurebootRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSecureboot
func (a *SystemInsightsApiService) SysteminsightsListSecurebootExecute(r ApiSysteminsightsListSecurebootRequest) ([]SystemInsightsSecureboot, *http.Response, error) {
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

type ApiSysteminsightsListServicesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListServicesRequest) Skip(skip int32) ApiSysteminsightsListServicesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListServicesRequest) Sort(sort []string) ApiSysteminsightsListServicesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListServicesRequest) Filter(filter []string) ApiSysteminsightsListServicesRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListServicesRequest) XOrgId(xOrgId string) ApiSysteminsightsListServicesRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListServicesRequest) Limit(limit int32) ApiSysteminsightsListServicesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListServicesRequest) Execute() ([]SystemInsightsServices, *http.Response, error) {
	return r.ApiService.SysteminsightsListServicesExecute(r)
}

/*
SysteminsightsListServices List System Insights Services

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListServicesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListServices(ctx context.Context) ApiSysteminsightsListServicesRequest {
	return ApiSysteminsightsListServicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsServices
func (a *SystemInsightsApiService) SysteminsightsListServicesExecute(r ApiSysteminsightsListServicesRequest) ([]SystemInsightsServices, *http.Response, error) {
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

type ApiSysteminsightsListShadowRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListShadowRequest) XOrgId(xOrgId string) ApiSysteminsightsListShadowRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListShadowRequest) Skip(skip int32) ApiSysteminsightsListShadowRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListShadowRequest) Sort(sort []string) ApiSysteminsightsListShadowRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListShadowRequest) Filter(filter []string) ApiSysteminsightsListShadowRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListShadowRequest) Limit(limit int32) ApiSysteminsightsListShadowRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListShadowRequest) Execute() ([]SystemInsightsShadow, *http.Response, error) {
	return r.ApiService.SysteminsightsListShadowExecute(r)
}

/*
SysteminsightsListShadow LIst System Insights Shadow

Valid filter fields are `system_id` and `username`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListShadowRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListShadow(ctx context.Context) ApiSysteminsightsListShadowRequest {
	return ApiSysteminsightsListShadowRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsShadow
func (a *SystemInsightsApiService) SysteminsightsListShadowExecute(r ApiSysteminsightsListShadowRequest) ([]SystemInsightsShadow, *http.Response, error) {
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

type ApiSysteminsightsListSharedFoldersRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListSharedFoldersRequest) XOrgId(xOrgId string) ApiSysteminsightsListSharedFoldersRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListSharedFoldersRequest) Skip(skip int32) ApiSysteminsightsListSharedFoldersRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListSharedFoldersRequest) Sort(sort []string) ApiSysteminsightsListSharedFoldersRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListSharedFoldersRequest) Filter(filter []string) ApiSysteminsightsListSharedFoldersRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListSharedFoldersRequest) Limit(limit int32) ApiSysteminsightsListSharedFoldersRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListSharedFoldersRequest) Execute() ([]SystemInsightsSharedFolders, *http.Response, error) {
	return r.ApiService.SysteminsightsListSharedFoldersExecute(r)
}

/*
SysteminsightsListSharedFolders List System Insights Shared Folders

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListSharedFoldersRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSharedFolders(ctx context.Context) ApiSysteminsightsListSharedFoldersRequest {
	return ApiSysteminsightsListSharedFoldersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSharedFolders
func (a *SystemInsightsApiService) SysteminsightsListSharedFoldersExecute(r ApiSysteminsightsListSharedFoldersRequest) ([]SystemInsightsSharedFolders, *http.Response, error) {
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

type ApiSysteminsightsListSharedResourcesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListSharedResourcesRequest) XOrgId(xOrgId string) ApiSysteminsightsListSharedResourcesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListSharedResourcesRequest) Skip(skip int32) ApiSysteminsightsListSharedResourcesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListSharedResourcesRequest) Sort(sort []string) ApiSysteminsightsListSharedResourcesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListSharedResourcesRequest) Filter(filter []string) ApiSysteminsightsListSharedResourcesRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListSharedResourcesRequest) Limit(limit int32) ApiSysteminsightsListSharedResourcesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListSharedResourcesRequest) Execute() ([]SystemInsightsSharedResources, *http.Response, error) {
	return r.ApiService.SysteminsightsListSharedResourcesExecute(r)
}

/*
SysteminsightsListSharedResources List System Insights Shared Resources

Valid filter fields are `system_id` and `type`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListSharedResourcesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSharedResources(ctx context.Context) ApiSysteminsightsListSharedResourcesRequest {
	return ApiSysteminsightsListSharedResourcesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSharedResources
func (a *SystemInsightsApiService) SysteminsightsListSharedResourcesExecute(r ApiSysteminsightsListSharedResourcesRequest) ([]SystemInsightsSharedResources, *http.Response, error) {
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

type ApiSysteminsightsListSharingPreferencesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListSharingPreferencesRequest) XOrgId(xOrgId string) ApiSysteminsightsListSharingPreferencesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListSharingPreferencesRequest) Skip(skip int32) ApiSysteminsightsListSharingPreferencesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListSharingPreferencesRequest) Sort(sort []string) ApiSysteminsightsListSharingPreferencesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListSharingPreferencesRequest) Filter(filter []string) ApiSysteminsightsListSharingPreferencesRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListSharingPreferencesRequest) Limit(limit int32) ApiSysteminsightsListSharingPreferencesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListSharingPreferencesRequest) Execute() ([]SystemInsightsSharingPreferences, *http.Response, error) {
	return r.ApiService.SysteminsightsListSharingPreferencesExecute(r)
}

/*
SysteminsightsListSharingPreferences List System Insights Sharing Preferences

Only valid filed field is `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListSharingPreferencesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSharingPreferences(ctx context.Context) ApiSysteminsightsListSharingPreferencesRequest {
	return ApiSysteminsightsListSharingPreferencesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSharingPreferences
func (a *SystemInsightsApiService) SysteminsightsListSharingPreferencesExecute(r ApiSysteminsightsListSharingPreferencesRequest) ([]SystemInsightsSharingPreferences, *http.Response, error) {
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

type ApiSysteminsightsListSipConfigRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListSipConfigRequest) XOrgId(xOrgId string) ApiSysteminsightsListSipConfigRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListSipConfigRequest) Skip(skip int32) ApiSysteminsightsListSipConfigRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListSipConfigRequest) Sort(sort []string) ApiSysteminsightsListSipConfigRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListSipConfigRequest) Filter(filter []string) ApiSysteminsightsListSipConfigRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListSipConfigRequest) Limit(limit int32) ApiSysteminsightsListSipConfigRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListSipConfigRequest) Execute() ([]SystemInsightsSipConfig, *http.Response, error) {
	return r.ApiService.SysteminsightsListSipConfigExecute(r)
}

/*
SysteminsightsListSipConfig List System Insights SIP Config

Valid filter fields are `system_id` and `enabled`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListSipConfigRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSipConfig(ctx context.Context) ApiSysteminsightsListSipConfigRequest {
	return ApiSysteminsightsListSipConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSipConfig
func (a *SystemInsightsApiService) SysteminsightsListSipConfigExecute(r ApiSysteminsightsListSipConfigRequest) ([]SystemInsightsSipConfig, *http.Response, error) {
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

type ApiSysteminsightsListStartupItemsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListStartupItemsRequest) Skip(skip int32) ApiSysteminsightsListStartupItemsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListStartupItemsRequest) Sort(sort []string) ApiSysteminsightsListStartupItemsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListStartupItemsRequest) Filter(filter []string) ApiSysteminsightsListStartupItemsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListStartupItemsRequest) XOrgId(xOrgId string) ApiSysteminsightsListStartupItemsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListStartupItemsRequest) Limit(limit int32) ApiSysteminsightsListStartupItemsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListStartupItemsRequest) Execute() ([]SystemInsightsStartupItems, *http.Response, error) {
	return r.ApiService.SysteminsightsListStartupItemsExecute(r)
}

/*
SysteminsightsListStartupItems List System Insights Startup Items

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListStartupItemsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListStartupItems(ctx context.Context) ApiSysteminsightsListStartupItemsRequest {
	return ApiSysteminsightsListStartupItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsStartupItems
func (a *SystemInsightsApiService) SysteminsightsListStartupItemsExecute(r ApiSysteminsightsListStartupItemsRequest) ([]SystemInsightsStartupItems, *http.Response, error) {
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

type ApiSysteminsightsListSystemControlsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListSystemControlsRequest) Skip(skip int32) ApiSysteminsightsListSystemControlsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListSystemControlsRequest) Sort(sort []string) ApiSysteminsightsListSystemControlsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; Note: You can only filter by &#x60;system_id&#x60; and &#x60;name&#x60; 
func (r ApiSysteminsightsListSystemControlsRequest) Filter(filter []string) ApiSysteminsightsListSystemControlsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListSystemControlsRequest) XOrgId(xOrgId string) ApiSysteminsightsListSystemControlsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListSystemControlsRequest) Limit(limit int32) ApiSysteminsightsListSystemControlsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListSystemControlsRequest) Execute() ([]SystemInsightsSystemControls, *http.Response, error) {
	return r.ApiService.SysteminsightsListSystemControlsExecute(r)
}

/*
SysteminsightsListSystemControls List System Insights System Control

Valid filter fields are `system_id` and `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListSystemControlsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSystemControls(ctx context.Context) ApiSysteminsightsListSystemControlsRequest {
	return ApiSysteminsightsListSystemControlsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSystemControls
func (a *SystemInsightsApiService) SysteminsightsListSystemControlsExecute(r ApiSysteminsightsListSystemControlsRequest) ([]SystemInsightsSystemControls, *http.Response, error) {
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

type ApiSysteminsightsListSystemInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListSystemInfoRequest) Skip(skip int32) ApiSysteminsightsListSystemInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListSystemInfoRequest) Sort(sort []string) ApiSysteminsightsListSystemInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListSystemInfoRequest) Filter(filter []string) ApiSysteminsightsListSystemInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListSystemInfoRequest) XOrgId(xOrgId string) ApiSysteminsightsListSystemInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListSystemInfoRequest) Limit(limit int32) ApiSysteminsightsListSystemInfoRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListSystemInfoRequest) Execute() ([]SystemInsightsSystemInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListSystemInfoExecute(r)
}

/*
SysteminsightsListSystemInfo List System Insights System Info

Valid filter fields are `system_id` and `cpu_subtype`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListSystemInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListSystemInfo(ctx context.Context) ApiSysteminsightsListSystemInfoRequest {
	return ApiSysteminsightsListSystemInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsSystemInfo
func (a *SystemInsightsApiService) SysteminsightsListSystemInfoExecute(r ApiSysteminsightsListSystemInfoRequest) ([]SystemInsightsSystemInfo, *http.Response, error) {
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

type ApiSysteminsightsListTpmInfoRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListTpmInfoRequest) Skip(skip int32) ApiSysteminsightsListTpmInfoRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListTpmInfoRequest) Sort(sort []string) ApiSysteminsightsListTpmInfoRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListTpmInfoRequest) Filter(filter []string) ApiSysteminsightsListTpmInfoRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListTpmInfoRequest) XOrgId(xOrgId string) ApiSysteminsightsListTpmInfoRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListTpmInfoRequest) Limit(limit int32) ApiSysteminsightsListTpmInfoRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListTpmInfoRequest) Execute() ([]SystemInsightsTpmInfo, *http.Response, error) {
	return r.ApiService.SysteminsightsListTpmInfoExecute(r)
}

/*
SysteminsightsListTpmInfo List System Insights TPM Info

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListTpmInfoRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListTpmInfo(ctx context.Context) ApiSysteminsightsListTpmInfoRequest {
	return ApiSysteminsightsListTpmInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsTpmInfo
func (a *SystemInsightsApiService) SysteminsightsListTpmInfoExecute(r ApiSysteminsightsListTpmInfoRequest) ([]SystemInsightsTpmInfo, *http.Response, error) {
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

type ApiSysteminsightsListUptimeRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListUptimeRequest) Skip(skip int32) ApiSysteminsightsListUptimeRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListUptimeRequest) Sort(sort []string) ApiSysteminsightsListUptimeRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, gte, in. e.g: Filter for single value: &#x60;filter&#x3D;field:gte:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListUptimeRequest) Filter(filter []string) ApiSysteminsightsListUptimeRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListUptimeRequest) XOrgId(xOrgId string) ApiSysteminsightsListUptimeRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListUptimeRequest) Limit(limit int32) ApiSysteminsightsListUptimeRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListUptimeRequest) Execute() ([]SystemInsightsUptime, *http.Response, error) {
	return r.ApiService.SysteminsightsListUptimeExecute(r)
}

/*
SysteminsightsListUptime List System Insights Uptime

Valid filter fields are `system_id` and `days`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListUptimeRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUptime(ctx context.Context) ApiSysteminsightsListUptimeRequest {
	return ApiSysteminsightsListUptimeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUptime
func (a *SystemInsightsApiService) SysteminsightsListUptimeExecute(r ApiSysteminsightsListUptimeRequest) ([]SystemInsightsUptime, *http.Response, error) {
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

type ApiSysteminsightsListUsbDevicesRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListUsbDevicesRequest) XOrgId(xOrgId string) ApiSysteminsightsListUsbDevicesRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListUsbDevicesRequest) Skip(skip int32) ApiSysteminsightsListUsbDevicesRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListUsbDevicesRequest) Sort(sort []string) ApiSysteminsightsListUsbDevicesRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListUsbDevicesRequest) Filter(filter []string) ApiSysteminsightsListUsbDevicesRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListUsbDevicesRequest) Limit(limit int32) ApiSysteminsightsListUsbDevicesRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListUsbDevicesRequest) Execute() ([]SystemInsightsUsbDevices, *http.Response, error) {
	return r.ApiService.SysteminsightsListUsbDevicesExecute(r)
}

/*
SysteminsightsListUsbDevices List System Insights USB Devices

Valid filter fields are `system_id` and `model`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListUsbDevicesRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUsbDevices(ctx context.Context) ApiSysteminsightsListUsbDevicesRequest {
	return ApiSysteminsightsListUsbDevicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUsbDevices
func (a *SystemInsightsApiService) SysteminsightsListUsbDevicesExecute(r ApiSysteminsightsListUsbDevicesRequest) ([]SystemInsightsUsbDevices, *http.Response, error) {
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

type ApiSysteminsightsListUserGroupsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListUserGroupsRequest) XOrgId(xOrgId string) ApiSysteminsightsListUserGroupsRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListUserGroupsRequest) Skip(skip int32) ApiSysteminsightsListUserGroupsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListUserGroupsRequest) Sort(sort []string) ApiSysteminsightsListUserGroupsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListUserGroupsRequest) Filter(filter []string) ApiSysteminsightsListUserGroupsRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListUserGroupsRequest) Limit(limit int32) ApiSysteminsightsListUserGroupsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListUserGroupsRequest) Execute() ([]SystemInsightsUserGroups, *http.Response, error) {
	return r.ApiService.SysteminsightsListUserGroupsExecute(r)
}

/*
SysteminsightsListUserGroups List System Insights User Groups

Only valid filter field is `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListUserGroupsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUserGroups(ctx context.Context) ApiSysteminsightsListUserGroupsRequest {
	return ApiSysteminsightsListUserGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUserGroups
func (a *SystemInsightsApiService) SysteminsightsListUserGroupsExecute(r ApiSysteminsightsListUserGroupsRequest) ([]SystemInsightsUserGroups, *http.Response, error) {
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

type ApiSysteminsightsListUserSshKeysRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	xOrgId *string
	skip *int32
	sort *[]string
	filter *[]string
	limit *int32
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListUserSshKeysRequest) XOrgId(xOrgId string) ApiSysteminsightsListUserSshKeysRequest {
	r.xOrgId = &xOrgId
	return r
}

// The offset into the records to return.
func (r ApiSysteminsightsListUserSshKeysRequest) Skip(skip int32) ApiSysteminsightsListUserSshKeysRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListUserSshKeysRequest) Sort(sort []string) ApiSysteminsightsListUserSshKeysRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListUserSshKeysRequest) Filter(filter []string) ApiSysteminsightsListUserSshKeysRequest {
	r.filter = &filter
	return r
}

func (r ApiSysteminsightsListUserSshKeysRequest) Limit(limit int32) ApiSysteminsightsListUserSshKeysRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListUserSshKeysRequest) Execute() ([]SystemInsightsUserSshKeys, *http.Response, error) {
	return r.ApiService.SysteminsightsListUserSshKeysExecute(r)
}

/*
SysteminsightsListUserSshKeys List System Insights User SSH Keys

Valid filter fields are `system_id` and `uid`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListUserSshKeysRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUserSshKeys(ctx context.Context) ApiSysteminsightsListUserSshKeysRequest {
	return ApiSysteminsightsListUserSshKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUserSshKeys
func (a *SystemInsightsApiService) SysteminsightsListUserSshKeysExecute(r ApiSysteminsightsListUserSshKeysRequest) ([]SystemInsightsUserSshKeys, *http.Response, error) {
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

type ApiSysteminsightsListUserassistRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListUserassistRequest) Skip(skip int32) ApiSysteminsightsListUserassistRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListUserassistRequest) Sort(sort []string) ApiSysteminsightsListUserassistRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListUserassistRequest) Filter(filter []string) ApiSysteminsightsListUserassistRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListUserassistRequest) XOrgId(xOrgId string) ApiSysteminsightsListUserassistRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListUserassistRequest) Limit(limit int32) ApiSysteminsightsListUserassistRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListUserassistRequest) Execute() ([]SystemInsightsUserassist, *http.Response, error) {
	return r.ApiService.SysteminsightsListUserassistExecute(r)
}

/*
SysteminsightsListUserassist List System Insights User Assist

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListUserassistRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUserassist(ctx context.Context) ApiSysteminsightsListUserassistRequest {
	return ApiSysteminsightsListUserassistRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUserassist
func (a *SystemInsightsApiService) SysteminsightsListUserassistExecute(r ApiSysteminsightsListUserassistRequest) ([]SystemInsightsUserassist, *http.Response, error) {
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

type ApiSysteminsightsListUsersRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListUsersRequest) Skip(skip int32) ApiSysteminsightsListUsersRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListUsersRequest) Sort(sort []string) ApiSysteminsightsListUsersRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListUsersRequest) Filter(filter []string) ApiSysteminsightsListUsersRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListUsersRequest) XOrgId(xOrgId string) ApiSysteminsightsListUsersRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListUsersRequest) Limit(limit int32) ApiSysteminsightsListUsersRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListUsersRequest) Execute() ([]SystemInsightsUsers, *http.Response, error) {
	return r.ApiService.SysteminsightsListUsersExecute(r)
}

/*
SysteminsightsListUsers List System Insights Users

Valid filter fields are `system_id` and `username`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListUsersRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListUsers(ctx context.Context) ApiSysteminsightsListUsersRequest {
	return ApiSysteminsightsListUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsUsers
func (a *SystemInsightsApiService) SysteminsightsListUsersExecute(r ApiSysteminsightsListUsersRequest) ([]SystemInsightsUsers, *http.Response, error) {
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

type ApiSysteminsightsListWifiNetworksRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListWifiNetworksRequest) Skip(skip int32) ApiSysteminsightsListWifiNetworksRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListWifiNetworksRequest) Sort(sort []string) ApiSysteminsightsListWifiNetworksRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListWifiNetworksRequest) Filter(filter []string) ApiSysteminsightsListWifiNetworksRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListWifiNetworksRequest) XOrgId(xOrgId string) ApiSysteminsightsListWifiNetworksRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListWifiNetworksRequest) Limit(limit int32) ApiSysteminsightsListWifiNetworksRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListWifiNetworksRequest) Execute() ([]SystemInsightsWifiNetworks, *http.Response, error) {
	return r.ApiService.SysteminsightsListWifiNetworksExecute(r)
}

/*
SysteminsightsListWifiNetworks List System Insights WiFi Networks

Valid filter fields are `system_id` and `security_type`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListWifiNetworksRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListWifiNetworks(ctx context.Context) ApiSysteminsightsListWifiNetworksRequest {
	return ApiSysteminsightsListWifiNetworksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsWifiNetworks
func (a *SystemInsightsApiService) SysteminsightsListWifiNetworksExecute(r ApiSysteminsightsListWifiNetworksRequest) ([]SystemInsightsWifiNetworks, *http.Response, error) {
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

type ApiSysteminsightsListWifiStatusRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListWifiStatusRequest) Skip(skip int32) ApiSysteminsightsListWifiStatusRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListWifiStatusRequest) Sort(sort []string) ApiSysteminsightsListWifiStatusRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListWifiStatusRequest) Filter(filter []string) ApiSysteminsightsListWifiStatusRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListWifiStatusRequest) XOrgId(xOrgId string) ApiSysteminsightsListWifiStatusRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListWifiStatusRequest) Limit(limit int32) ApiSysteminsightsListWifiStatusRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListWifiStatusRequest) Execute() ([]SystemInsightsWifiStatus, *http.Response, error) {
	return r.ApiService.SysteminsightsListWifiStatusExecute(r)
}

/*
SysteminsightsListWifiStatus List System Insights WiFi Status

Valid filter fields are `system_id` and `security_type`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListWifiStatusRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListWifiStatus(ctx context.Context) ApiSysteminsightsListWifiStatusRequest {
	return ApiSysteminsightsListWifiStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsWifiStatus
func (a *SystemInsightsApiService) SysteminsightsListWifiStatusExecute(r ApiSysteminsightsListWifiStatusRequest) ([]SystemInsightsWifiStatus, *http.Response, error) {
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

type ApiSysteminsightsListWindowsSecurityCenterRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListWindowsSecurityCenterRequest) Skip(skip int32) ApiSysteminsightsListWindowsSecurityCenterRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListWindowsSecurityCenterRequest) Sort(sort []string) ApiSysteminsightsListWindowsSecurityCenterRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListWindowsSecurityCenterRequest) Filter(filter []string) ApiSysteminsightsListWindowsSecurityCenterRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListWindowsSecurityCenterRequest) XOrgId(xOrgId string) ApiSysteminsightsListWindowsSecurityCenterRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListWindowsSecurityCenterRequest) Limit(limit int32) ApiSysteminsightsListWindowsSecurityCenterRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListWindowsSecurityCenterRequest) Execute() ([]SystemInsightsWindowsSecurityCenter, *http.Response, error) {
	return r.ApiService.SysteminsightsListWindowsSecurityCenterExecute(r)
}

/*
SysteminsightsListWindowsSecurityCenter List System Insights Windows Security Center

Valid filter fields are `system_id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListWindowsSecurityCenterRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListWindowsSecurityCenter(ctx context.Context) ApiSysteminsightsListWindowsSecurityCenterRequest {
	return ApiSysteminsightsListWindowsSecurityCenterRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsWindowsSecurityCenter
func (a *SystemInsightsApiService) SysteminsightsListWindowsSecurityCenterExecute(r ApiSysteminsightsListWindowsSecurityCenterRequest) ([]SystemInsightsWindowsSecurityCenter, *http.Response, error) {
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

type ApiSysteminsightsListWindowsSecurityProductsRequest struct {
	ctx context.Context
	ApiService *SystemInsightsApiService
	skip *int32
	sort *[]string
	filter *[]string
	xOrgId *string
	limit *int32
}

// The offset into the records to return.
func (r ApiSysteminsightsListWindowsSecurityProductsRequest) Skip(skip int32) ApiSysteminsightsListWindowsSecurityProductsRequest {
	r.skip = &skip
	return r
}

// The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. e.g: Sort by single field: &#x60;sort&#x3D;field&#x60; Sort descending by single field: &#x60;sort&#x3D;-field&#x60; Sort by multiple fields: &#x60;sort&#x3D;field1,-field2,field3&#x60; 
func (r ApiSysteminsightsListWindowsSecurityProductsRequest) Sort(sort []string) ApiSysteminsightsListWindowsSecurityProductsRequest {
	r.sort = &sort
	return r
}

// Supported operators are: eq, in. e.g: Filter for single value: &#x60;filter&#x3D;field:eq:value&#x60; Filter for any value in a list: (note \&quot;pipe\&quot; character: &#x60;|&#x60; separating values) &#x60;filter&#x3D;field:in:value1|value2|value3&#x60; 
func (r ApiSysteminsightsListWindowsSecurityProductsRequest) Filter(filter []string) ApiSysteminsightsListWindowsSecurityProductsRequest {
	r.filter = &filter
	return r
}

// Organization identifier that can be obtained from console settings.
func (r ApiSysteminsightsListWindowsSecurityProductsRequest) XOrgId(xOrgId string) ApiSysteminsightsListWindowsSecurityProductsRequest {
	r.xOrgId = &xOrgId
	return r
}

func (r ApiSysteminsightsListWindowsSecurityProductsRequest) Limit(limit int32) ApiSysteminsightsListWindowsSecurityProductsRequest {
	r.limit = &limit
	return r
}

func (r ApiSysteminsightsListWindowsSecurityProductsRequest) Execute() ([]SystemInsightsWindowsSecurityProducts, *http.Response, error) {
	return r.ApiService.SysteminsightsListWindowsSecurityProductsExecute(r)
}

/*
SysteminsightsListWindowsSecurityProducts List System Insights Windows Security Products

Valid filter fields are `system_id` and `state`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSysteminsightsListWindowsSecurityProductsRequest
*/
func (a *SystemInsightsApiService) SysteminsightsListWindowsSecurityProducts(ctx context.Context) ApiSysteminsightsListWindowsSecurityProductsRequest {
	return ApiSysteminsightsListWindowsSecurityProductsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SystemInsightsWindowsSecurityProducts
func (a *SystemInsightsApiService) SysteminsightsListWindowsSecurityProductsExecute(r ApiSysteminsightsListWindowsSecurityProductsRequest) ([]SystemInsightsWindowsSecurityProducts, *http.Response, error) {
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

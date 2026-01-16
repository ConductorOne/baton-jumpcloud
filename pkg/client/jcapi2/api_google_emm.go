/*
JumpCloud API

# Overview  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.  ## API Best Practices  Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.  Covered topics include: 1. Important Considerations 2. Supported HTTP Request Methods 3. Response codes 4. API Key rotation 5. Paginating 6. Error handling 7. Retry rates  [JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)  # Directory Objects  This API offers the ability to interact with some of our core features; otherwise known as Directory Objects. The Directory Objects are:  * Commands * Policies * Policy Groups * Applications * Systems * Users * User Groups * System Groups * Radius Servers * Directories: Office 365, LDAP,G-Suite, Active Directory * Duo accounts and applications.  The Directory Object is an important concept to understand in order to successfully use JumpCloud API.  ## JumpCloud Graph  We've also introduced the concept of the JumpCloud Graph along with  Directory Objects. The Graph is a powerful aspect of our platform which will enable you to associate objects with each other, or establish membership for certain objects to become members of other objects.  Specific `GET` endpoints will allow you to traverse the JumpCloud Graph to return all indirect and directly bound objects in your organization.  | ![alt text](https://s3.amazonaws.com/jumpcloud-kb/Knowledge+Base+Photos/API+Docs/jumpcloud_graph.png \"JumpCloud Graph Model Example\") | |:--:| | **This diagram highlights our association and membership model as it relates to Directory Objects.** |  # API Key  ## Access Your API Key  To locate your API Key:  1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/). 2. Go to the username drop down located in the top-right of the Console. 3. Retrieve your API key from API Settings.  ## API Key Considerations  This API key is associated to the currently logged in administrator. Other admins will have different API keys.  **WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.  You can also reset your API key in the same location in the JumpCloud Admin Console.  ## Recycling or Resetting Your API Key  In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.  Your API key will be passed in as a header with the header name \"x-api-key\".  ```bash curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/v2/systemgroups\" ```  # System Context  * [Introduction](#introduction) * [Supported endpoints](#supported-endpoints) * [Response codes](#response-codes) * [Authentication](#authentication) * [Additional examples](#additional-examples) * [Third party](#third-party)  ## Introduction  JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.  **Notes:**   * The following documentation applies to Linux Operating Systems only.  * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.  ## Supported Endpoints  JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.  * A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,   * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put) * A system may delete itself from your JumpCloud organization   * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete) * A system may fetch its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)   * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)   * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser) * A system may alter its direct resource associations under v2 (Groups)   * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost) * A system may alter its System Group associations   * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)     * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected  ## Response Codes  If endpoints other than those described above are called using the System Context API, the server will return a `401` response.  ## Authentication  To allow for secure access to our APIs, you must authenticate each API request. JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests. The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API. To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.  Here is a breakdown of the example script with explanations.  First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.  ```bash #!/bin/bash conf=\"`cat /opt/jc/jcagent.conf`\" regex=\"systemKey\\\":\\\"(\\w+)\\\"\"  if [[ $conf =~ $regex ]] ; then   systemKey=\"${BASH_REMATCH[1]}\" fi ```  Then, the script retrieves the current date in the correct format.  ```bash now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`; ```  Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.  ```bash signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" ```  The next step is to calculate and apply the signature. This is a two-step process:  1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key`` 2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``  The combined steps above result in:  ```bash signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ; ```  Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.  ```bash curl -iq \\   -H \"Accept: application/json\" \\   -H \"Content-Type: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Input Data  All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.  The following example demonstrates how to update the `displayName` of the system.  ```bash signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\" signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;  curl -iq \\   -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\   -X \"PUT\" \\   -H \"Content-Type: application/json\" \\   -H \"Accept: application/json\" \\   -H \"Date: ${now}\" \\   -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\   --url https://console.jumpcloud.com/api/systems/${systemKey} ```  ### Output Data  All results will be formatted as JSON.  Here is an abbreviated example of response output:  ```json {   \"_id\": \"525ee96f52e144993e000015\",   \"agentServer\": \"lappy386\",   \"agentVersion\": \"0.9.42\",   \"arch\": \"x86_64\",   \"connectionKey\": \"127.0.0.1_51812\",   \"displayName\": \"ubuntu-1204\",   \"firstContact\": \"2013-10-16T19:30:55.611Z\",   \"hostname\": \"ubuntu-1204\"   ... ```  ## Additional Examples  ### Signing Authentication Example  This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.  [SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)  ### Shutdown Hook  This example demonstrates how to make an authenticated request on system shutdown. Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.  [Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.  After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.  1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`. 2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.  ## Third Party  ### Chef Cookbooks  [https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)  [https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)  # Multi-Tenant Portal Headers  Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.  The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.  **NOTE** Single Tenant Admins do not need to provide this header when making an API request.  ## Header Value  `x-org-id`  ## API Response Codes  * `400` Malformed ID. * `400` x-org-id and Organization path ID do not match. * `401` ID not included for multi-tenant admin * `403` ID included on unsupported route. * `404` Organization ID Not Found.  ```bash curl -X GET https://console.jumpcloud.com/api/v2/directories \\   -H 'accept: application/json' \\   -H 'content-type: application/json' \\   -H 'x-api-key: {API_KEY}' \\   -H 'x-org-id: {ORG_ID}'  ```  ## To Obtain an Individual Organization ID via the UI  As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.  1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal. 2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access. 3. You will automatically be routed to that Organization's Admin Console. 4. Go to Settings in the sub-tenant's primary navigation. 5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.  ## To Obtain All Organization IDs via the API  * You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.  ```bash curl -X GET \\   https://console.jumpcloud.com/api/organizations/ \\   -H 'Accept: application/json' \\   -H 'Content-Type: application/json' \\   -H 'x-api-key: {API_KEY}' ```  # SDKs  You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:  * [Python](https://github.com/TheJumpCloud/jcapi-python) * [Go](https://github.com/TheJumpCloud/jcapi-go) * [Ruby](https://github.com/TheJumpCloud/jcapi-ruby) * [Java](https://github.com/TheJumpCloud/jcapi-java) 

API version: 2.0
Contact: support@jumpcloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jcapi2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// GoogleEMMApiService GoogleEMMApi service
type GoogleEMMApiService service

type GoogleEMMApiDevicesGetDeviceRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	deviceId string
}

func (r GoogleEMMApiDevicesGetDeviceRequest) Execute() (*JumpcloudGoogleEmmDevice, *http.Response, error) {
	return r.ApiService.DevicesGetDeviceExecute(r)
}

/*
DevicesGetDevice Get device

Gets a Google EMM enrolled device details.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId} \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return GoogleEMMApiDevicesGetDeviceRequest
*/
func (a *GoogleEMMApiService) DevicesGetDevice(ctx context.Context, deviceId string) GoogleEMMApiDevicesGetDeviceRequest {
	return GoogleEMMApiDevicesGetDeviceRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return JumpcloudGoogleEmmDevice
func (a *GoogleEMMApiService) DevicesGetDeviceExecute(r GoogleEMMApiDevicesGetDeviceRequest) (*JumpcloudGoogleEmmDevice, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JumpcloudGoogleEmmDevice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.DevicesGetDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/devices/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

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

type GoogleEMMApiDevicesGetDeviceAndroidPolicyRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	deviceId string
}

func (r GoogleEMMApiDevicesGetDeviceAndroidPolicyRequest) Execute() (*JumpcloudGoogleEmmDeviceAndroidPolicy, *http.Response, error) {
	return r.ApiService.DevicesGetDeviceAndroidPolicyExecute(r)
}

/*
DevicesGetDeviceAndroidPolicy Get the policy JSON of a device

Gets an android JSON policy for a Google EMM enrolled device.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId}/policy_results \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return GoogleEMMApiDevicesGetDeviceAndroidPolicyRequest
*/
func (a *GoogleEMMApiService) DevicesGetDeviceAndroidPolicy(ctx context.Context, deviceId string) GoogleEMMApiDevicesGetDeviceAndroidPolicyRequest {
	return GoogleEMMApiDevicesGetDeviceAndroidPolicyRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return JumpcloudGoogleEmmDeviceAndroidPolicy
func (a *GoogleEMMApiService) DevicesGetDeviceAndroidPolicyExecute(r GoogleEMMApiDevicesGetDeviceAndroidPolicyRequest) (*JumpcloudGoogleEmmDeviceAndroidPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JumpcloudGoogleEmmDeviceAndroidPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.DevicesGetDeviceAndroidPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/devices/{deviceId}/policy_results"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

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

type GoogleEMMApiDevicesLockDeviceRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	deviceId string
	body *map[string]interface{}
}

func (r GoogleEMMApiDevicesLockDeviceRequest) Body(body map[string]interface{}) GoogleEMMApiDevicesLockDeviceRequest {
	r.body = &body
	return r
}

func (r GoogleEMMApiDevicesLockDeviceRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DevicesLockDeviceExecute(r)
}

/*
DevicesLockDevice Lock device

Locks a Google EMM enrolled device, as if the lock screen timeout had expired.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId}/lock \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return GoogleEMMApiDevicesLockDeviceRequest
*/
func (a *GoogleEMMApiService) DevicesLockDevice(ctx context.Context, deviceId string) GoogleEMMApiDevicesLockDeviceRequest {
	return GoogleEMMApiDevicesLockDeviceRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *GoogleEMMApiService) DevicesLockDeviceExecute(r GoogleEMMApiDevicesLockDeviceRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.DevicesLockDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/devices/{deviceId}/lock"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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

type GoogleEMMApiDevicesRelinquishOwnershipRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	deviceId string
	body *map[string]interface{}
}

func (r GoogleEMMApiDevicesRelinquishOwnershipRequest) Body(body map[string]interface{}) GoogleEMMApiDevicesRelinquishOwnershipRequest {
	r.body = &body
	return r
}

func (r GoogleEMMApiDevicesRelinquishOwnershipRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DevicesRelinquishOwnershipExecute(r)
}

/*
DevicesRelinquishOwnership Relinquish Ownership of COPE device

Removes the work profile and all policies from a company-owned Android 8.0+ device, relinquishing the device for personal use. Apps and data associated with the personal profile(s) are preserved.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId}/reliquinshownership \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return GoogleEMMApiDevicesRelinquishOwnershipRequest
*/
func (a *GoogleEMMApiService) DevicesRelinquishOwnership(ctx context.Context, deviceId string) GoogleEMMApiDevicesRelinquishOwnershipRequest {
	return GoogleEMMApiDevicesRelinquishOwnershipRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *GoogleEMMApiService) DevicesRelinquishOwnershipExecute(r GoogleEMMApiDevicesRelinquishOwnershipRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.DevicesRelinquishOwnership")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/devices/{deviceId}/relinquishownership"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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

type GoogleEMMApiDevicesResetPasswordRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	deviceId string
	body *DevicesResetPasswordRequest
}

func (r GoogleEMMApiDevicesResetPasswordRequest) Body(body DevicesResetPasswordRequest) GoogleEMMApiDevicesResetPasswordRequest {
	r.body = &body
	return r
}

func (r GoogleEMMApiDevicesResetPasswordRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DevicesResetPasswordExecute(r)
}

/*
DevicesResetPassword Reset Password of a device

Reset the user's password of a Google EMM enrolled device.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/google-emm/devices/{deviceId}/resetpassword \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{ 'new_password' : 'string' }' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return GoogleEMMApiDevicesResetPasswordRequest
*/
func (a *GoogleEMMApiService) DevicesResetPassword(ctx context.Context, deviceId string) GoogleEMMApiDevicesResetPasswordRequest {
	return GoogleEMMApiDevicesResetPasswordRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *GoogleEMMApiService) DevicesResetPasswordExecute(r GoogleEMMApiDevicesResetPasswordRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.DevicesResetPassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/devices/{deviceId}/resetpassword"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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

type GoogleEMMApiEnterprisesCreateEnterpriseRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	body *JumpcloudGoogleEmmCreateEnterpriseRequest
}

func (r GoogleEMMApiEnterprisesCreateEnterpriseRequest) Body(body JumpcloudGoogleEmmCreateEnterpriseRequest) GoogleEMMApiEnterprisesCreateEnterpriseRequest {
	r.body = &body
	return r
}

func (r GoogleEMMApiEnterprisesCreateEnterpriseRequest) Execute() (*JumpcloudGoogleEmmEnterprise, *http.Response, error) {
	return r.ApiService.EnterprisesCreateEnterpriseExecute(r)
}

/*
EnterprisesCreateEnterprise Create a Google Enterprise

Creates a Google EMM enterprise.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/google-emm/enterprises \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{ 'signupUrlName': 'string', 'enrollmentToken': 'string' }' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GoogleEMMApiEnterprisesCreateEnterpriseRequest
*/
func (a *GoogleEMMApiService) EnterprisesCreateEnterprise(ctx context.Context) GoogleEMMApiEnterprisesCreateEnterpriseRequest {
	return GoogleEMMApiEnterprisesCreateEnterpriseRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JumpcloudGoogleEmmEnterprise
func (a *GoogleEMMApiService) EnterprisesCreateEnterpriseExecute(r GoogleEMMApiEnterprisesCreateEnterpriseRequest) (*JumpcloudGoogleEmmEnterprise, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JumpcloudGoogleEmmEnterprise
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.EnterprisesCreateEnterprise")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/enterprises"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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

type GoogleEMMApiEnterprisesDeleteEnterpriseRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	enterpriseId string
}

func (r GoogleEMMApiEnterprisesDeleteEnterpriseRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.EnterprisesDeleteEnterpriseExecute(r)
}

/*
EnterprisesDeleteEnterprise Delete a Google Enterprise

Removes a Google EMM enterprise.

 Warning: This is a destructive operation and will remove all data associated with Google EMM enterprise from JumpCloud including devices and applications associated with the given enterprise.

#### Sample Request
```
curl -X DELETE https://console.jumpcloud.com/api/v2/google-emm/devices/{enterpriseId} \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param enterpriseId
 @return GoogleEMMApiEnterprisesDeleteEnterpriseRequest
*/
func (a *GoogleEMMApiService) EnterprisesDeleteEnterprise(ctx context.Context, enterpriseId string) GoogleEMMApiEnterprisesDeleteEnterpriseRequest {
	return GoogleEMMApiEnterprisesDeleteEnterpriseRequest{
		ApiService: a,
		ctx: ctx,
		enterpriseId: enterpriseId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *GoogleEMMApiService) EnterprisesDeleteEnterpriseExecute(r GoogleEMMApiEnterprisesDeleteEnterpriseRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.EnterprisesDeleteEnterprise")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/enterprises/{enterpriseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"enterpriseId"+"}", url.PathEscape(parameterValueToString(r.enterpriseId, "enterpriseId")), -1)

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

type GoogleEMMApiEnterprisesGetConnectionStatusRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	enterpriseId string
}

func (r GoogleEMMApiEnterprisesGetConnectionStatusRequest) Execute() (*JumpcloudGoogleEmmConnectionStatus, *http.Response, error) {
	return r.ApiService.EnterprisesGetConnectionStatusExecute(r)
}

/*
EnterprisesGetConnectionStatus Test connection with Google

Gives a connection status between JumpCloud and Google.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/google-emm/devices/{enterpriseId}/connection-status \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param enterpriseId
 @return GoogleEMMApiEnterprisesGetConnectionStatusRequest
*/
func (a *GoogleEMMApiService) EnterprisesGetConnectionStatus(ctx context.Context, enterpriseId string) GoogleEMMApiEnterprisesGetConnectionStatusRequest {
	return GoogleEMMApiEnterprisesGetConnectionStatusRequest{
		ApiService: a,
		ctx: ctx,
		enterpriseId: enterpriseId,
	}
}

// Execute executes the request
//  @return JumpcloudGoogleEmmConnectionStatus
func (a *GoogleEMMApiService) EnterprisesGetConnectionStatusExecute(r GoogleEMMApiEnterprisesGetConnectionStatusRequest) (*JumpcloudGoogleEmmConnectionStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JumpcloudGoogleEmmConnectionStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.EnterprisesGetConnectionStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/enterprises/{enterpriseId}/connection-status"
	localVarPath = strings.Replace(localVarPath, "{"+"enterpriseId"+"}", url.PathEscape(parameterValueToString(r.enterpriseId, "enterpriseId")), -1)

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

type GoogleEMMApiEnterprisesListEnterprisesRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	queryParametersLimit *string
	queryParametersSkip *string
	queryParametersFields *[]string
}

func (r GoogleEMMApiEnterprisesListEnterprisesRequest) QueryParametersLimit(queryParametersLimit string) GoogleEMMApiEnterprisesListEnterprisesRequest {
	r.queryParametersLimit = &queryParametersLimit
	return r
}

func (r GoogleEMMApiEnterprisesListEnterprisesRequest) QueryParametersSkip(queryParametersSkip string) GoogleEMMApiEnterprisesListEnterprisesRequest {
	r.queryParametersSkip = &queryParametersSkip
	return r
}

func (r GoogleEMMApiEnterprisesListEnterprisesRequest) QueryParametersFields(queryParametersFields []string) GoogleEMMApiEnterprisesListEnterprisesRequest {
	r.queryParametersFields = &queryParametersFields
	return r
}

func (r GoogleEMMApiEnterprisesListEnterprisesRequest) Execute() (*JumpcloudGoogleEmmListEnterprisesResponse, *http.Response, error) {
	return r.ApiService.EnterprisesListEnterprisesExecute(r)
}

/*
EnterprisesListEnterprises List Google Enterprises

Lists all Google EMM enterprises. An empty list indicates that the Organization is not configured with a Google EMM enterprise yet. 

 Note: Currently only one Google Enterprise per Organization is supported.

#### Sample Request
```
curl -X GET https://console.jumpcloud.com/api/v2/google-emm/enterprises \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GoogleEMMApiEnterprisesListEnterprisesRequest
*/
func (a *GoogleEMMApiService) EnterprisesListEnterprises(ctx context.Context) GoogleEMMApiEnterprisesListEnterprisesRequest {
	return GoogleEMMApiEnterprisesListEnterprisesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JumpcloudGoogleEmmListEnterprisesResponse
func (a *GoogleEMMApiService) EnterprisesListEnterprisesExecute(r GoogleEMMApiEnterprisesListEnterprisesRequest) (*JumpcloudGoogleEmmListEnterprisesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JumpcloudGoogleEmmListEnterprisesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.EnterprisesListEnterprises")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/enterprises"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.queryParametersLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "queryParameters.limit", r.queryParametersLimit, "")
	}
	if r.queryParametersSkip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "queryParameters.skip", r.queryParametersSkip, "")
	}
	if r.queryParametersFields != nil {
		t := *r.queryParametersFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "queryParameters.fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "queryParameters.fields", t, "multi")
		}
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

type GoogleEMMApiEnterprisesPatchEnterpriseRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	enterpriseId string
	body *EnterprisesPatchEnterpriseRequest
}

func (r GoogleEMMApiEnterprisesPatchEnterpriseRequest) Body(body EnterprisesPatchEnterpriseRequest) GoogleEMMApiEnterprisesPatchEnterpriseRequest {
	r.body = &body
	return r
}

func (r GoogleEMMApiEnterprisesPatchEnterpriseRequest) Execute() (*JumpcloudGoogleEmmEnterprise, *http.Response, error) {
	return r.ApiService.EnterprisesPatchEnterpriseExecute(r)
}

/*
EnterprisesPatchEnterprise Update a Google Enterprise

Updates a Google EMM enterprise details.

#### Sample Request
```
curl -X PATCH https://console.jumpcloud.com/api/v2/google-emm/enterprises \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
  -d '{ 'allowDeviceEnrollment': true, 'deviceGroupId': 'string' }' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param enterpriseId
 @return GoogleEMMApiEnterprisesPatchEnterpriseRequest
*/
func (a *GoogleEMMApiService) EnterprisesPatchEnterprise(ctx context.Context, enterpriseId string) GoogleEMMApiEnterprisesPatchEnterpriseRequest {
	return GoogleEMMApiEnterprisesPatchEnterpriseRequest{
		ApiService: a,
		ctx: ctx,
		enterpriseId: enterpriseId,
	}
}

// Execute executes the request
//  @return JumpcloudGoogleEmmEnterprise
func (a *GoogleEMMApiService) EnterprisesPatchEnterpriseExecute(r GoogleEMMApiEnterprisesPatchEnterpriseRequest) (*JumpcloudGoogleEmmEnterprise, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JumpcloudGoogleEmmEnterprise
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.EnterprisesPatchEnterprise")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/enterprises/{enterpriseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"enterpriseId"+"}", url.PathEscape(parameterValueToString(r.enterpriseId, "enterpriseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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

type GoogleEMMApiSignupURLsCreateRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
}

func (r GoogleEMMApiSignupURLsCreateRequest) Execute() (*JumpcloudGoogleEmmSignupURL, *http.Response, error) {
	return r.ApiService.SignupURLsCreateExecute(r)
}

/*
SignupURLsCreate Get a Signup URL to enroll Google enterprise

Creates a Google EMM enterprise signup URL.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/google-emm/signup-urls \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GoogleEMMApiSignupURLsCreateRequest
*/
func (a *GoogleEMMApiService) SignupURLsCreate(ctx context.Context) GoogleEMMApiSignupURLsCreateRequest {
	return GoogleEMMApiSignupURLsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JumpcloudGoogleEmmSignupURL
func (a *GoogleEMMApiService) SignupURLsCreateExecute(r GoogleEMMApiSignupURLsCreateRequest) (*JumpcloudGoogleEmmSignupURL, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JumpcloudGoogleEmmSignupURL
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.SignupURLsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/signup-urls"

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

type GoogleEMMApiWebTokensCreateWebTokenRequest struct {
	ctx context.Context
	ApiService *GoogleEMMApiService
	body *JumpcloudGoogleEmmCreateWebTokenRequest
}

func (r GoogleEMMApiWebTokensCreateWebTokenRequest) Body(body JumpcloudGoogleEmmCreateWebTokenRequest) GoogleEMMApiWebTokensCreateWebTokenRequest {
	r.body = &body
	return r
}

func (r GoogleEMMApiWebTokensCreateWebTokenRequest) Execute() (*JumpcloudGoogleEmmWebToken, *http.Response, error) {
	return r.ApiService.WebTokensCreateWebTokenExecute(r)
}

/*
WebTokensCreateWebToken Get a web token to render Google Play iFrame

Creates a web token to access an embeddable managed Google Play web UI for a given Google EMM enterprise.

#### Sample Request
```
curl -X POST https://console.jumpcloud.com/api/v2/google-emm/web-tokens \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: {API_KEY}' \
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GoogleEMMApiWebTokensCreateWebTokenRequest
*/
func (a *GoogleEMMApiService) WebTokensCreateWebToken(ctx context.Context) GoogleEMMApiWebTokensCreateWebTokenRequest {
	return GoogleEMMApiWebTokensCreateWebTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JumpcloudGoogleEmmWebToken
func (a *GoogleEMMApiService) WebTokensCreateWebTokenExecute(r GoogleEMMApiWebTokensCreateWebTokenRequest) (*JumpcloudGoogleEmmWebToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JumpcloudGoogleEmmWebToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleEMMApiService.WebTokensCreateWebToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google-emm/web-tokens"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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

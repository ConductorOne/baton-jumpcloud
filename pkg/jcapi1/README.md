# Go API client for jcapi1

# Overview

JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, and system users.

## API Best Practices

Read the linked Help Article below for guidance on retrying failed requests to JumpCloud's REST API, as well as best practices for structuring subsequent retry requests. Customizing retry mechanisms based on these recommendations will increase the reliability and dependability of your API calls.

Covered topics include:
1. Important Considerations
2. Supported HTTP Request Methods
3. Response codes
4. API Key rotation
5. Paginating
6. Error handling
7. Retry rates

[JumpCloud Help Center - API Best Practices](https://support.jumpcloud.com/support/s/article/JumpCloud-API-Best-Practices)

# API Key

## Access Your API Key

To locate your API Key:

1. Log into the [JumpCloud Admin Console](https://console.jumpcloud.com/).
2. Go to the username drop down located in the top-right of the Console.
3. Retrieve your API key from API Settings.

## API Key Considerations

This API key is associated to the currently logged in administrator. Other admins will have different API keys.

**WARNING** Please keep this API key secret, as it grants full access to any data accessible via your JumpCloud console account.

You can also reset your API key in the same location in the JumpCloud Admin Console.

## Recycling or Resetting Your API Key

In order to revoke access with the current API key, simply reset your API key. This will render all calls using the previous API key inaccessible.

Your API key will be passed in as a header with the header name \"x-api-key\".

```bash
curl -H \"x-api-key: [YOUR_API_KEY_HERE]\" \"https://console.jumpcloud.com/api/systemusers\"
```

# System Context

* [Introduction](#introduction)
* [Supported endpoints](#supported-endpoints)
* [Response codes](#response-codes)
* [Authentication](#authentication)
* [Additional examples](#additional-examples)
* [Third party](#third-party)

## Introduction

JumpCloud System Context Authorization is an alternative way to authenticate with a subset of JumpCloud's REST APIs. Using this method, a system can manage its information and resource associations, allowing modern auto provisioning environments to scale as needed.

**Notes:**

 * The following documentation applies to Linux Operating Systems only.
 * Systems that have been automatically enrolled using Apple's Device Enrollment Program (DEP) or systems enrolled using the User Portal install are not eligible to use the System Context API to prevent unauthorized access to system groups and resources. If a script that utilizes the System Context API is invoked on a system enrolled in this way, it will display an error.

## Supported Endpoints

JumpCloud System Context Authorization can be used in conjunction with Systems endpoints found in the V1 API and certain System Group endpoints found in the v2 API.

* A system may fetch, alter, and delete metadata about itself, including manipulating a system's Group and Systemuser associations,
  * `/api/systems/{system_id}` | [`GET`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_get) [`PUT`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_put)
* A system may delete itself from your JumpCloud organization
  * `/api/systems/{system_id}` | [`DELETE`](https://docs.jumpcloud.com/api/1.0/index.html#operation/systems_delete)
* A system may fetch its direct resource associations under v2 (Groups)
  * `/api/v2/systems/{system_id}/memberof` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembership)
  * `/api/v2/systems/{system_id}/associations` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsList)
  * `/api/v2/systems/{system_id}/users` | [`GET`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemTraverseUser)
* A system may alter its direct resource associations under v2 (Groups)
  * `/api/v2/systems/{system_id}/associations` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemAssociationsPost)
* A system may alter its System Group associations
  * `/api/v2/systemgroups/{group_id}/members` | [`POST`](https://docs.jumpcloud.com/api/2.0/index.html#operation/graph_systemGroupMembersPost)
    * _NOTE_ If a system attempts to alter the system group membership of a different system the request will be rejected

## Response Codes

If endpoints other than those described above are called using the System Context API, the server will return a `401` response.

## Authentication

To allow for secure access to our APIs, you must authenticate each API request.
JumpCloud System Context Authorization uses [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures-00) to authenticate API requests.
The HTTP Signatures sent with each request are similar to the signatures used by the Amazon Web Services REST API.
To help with the request-signing process, we have provided an [example bash script](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh). This example API request simply requests the entire system record. You must be root, or have permissions to access the contents of the `/opt/jc` directory to generate a signature.

Here is a breakdown of the example script with explanations.

First, the script extracts the systemKey from the JSON formatted `/opt/jc/jcagent.conf` file.

```bash
#!/bin/bash
conf=\"`cat /opt/jc/jcagent.conf`\"
regex=\"systemKey\\\":\\\"(\\w+)\\\"\"

if [[ $conf =~ $regex ]] ; then
  systemKey=\"${BASH_REMATCH[1]}\"
fi
```

Then, the script retrieves the current date in the correct format.

```bash
now=`date -u \"+%a, %d %h %Y %H:%M:%S GMT\"`;
```

Next, we build a signing string to demonstrate the expected signature format. The signed string must consist of the [request-line](https://tools.ietf.org/html/rfc2616#page-35) and the date header, separated by a newline character.

```bash
signstr=\"GET /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\"
```

The next step is to calculate and apply the signature. This is a two-step process:

1. Create a signature from the signing string using the JumpCloud Agent private key: ``printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key``
2. Then Base64-encode the signature string and trim off the newline characters: ``| openssl enc -e -a | tr -d '\\n'``

The combined steps above result in:

```bash
signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;
```

Finally, we make sure the API call sending the signature has the same Authorization and Date header values, HTTP method, and URL that were used in the signing string.

```bash
curl -iq \\
  -H \"Accept: application/json\" \\
  -H \"Content-Type: application/json\" \\
  -H \"Date: ${now}\" \\
  -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\
  --url https://console.jumpcloud.com/api/systems/${systemKey}
```

### Input Data

All PUT and POST methods should use the HTTP Content-Type header with a value of 'application/json'. PUT methods are used for updating a record. POST methods are used to create a record.

The following example demonstrates how to update the `displayName` of the system.

```bash
signstr=\"PUT /api/systems/${systemKey} HTTP/1.1\\ndate: ${now}\"
signature=`printf \"$signstr\" | openssl dgst -sha256 -sign /opt/jc/client.key | openssl enc -e -a | tr -d '\\n'` ;

curl -iq \\
  -d \"{\\\"displayName\\\" : \\\"updated-system-name-1\\\"}\" \\
  -X \"PUT\" \\
  -H \"Content-Type: application/json\" \\
  -H \"Accept: application/json\" \\
  -H \"Date: ${now}\" \\
  -H \"Authorization: Signature keyId=\\\"system/${systemKey}\\\",headers=\\\"request-line date\\\",algorithm=\\\"rsa-sha256\\\",signature=\\\"${signature}\\\"\" \\
  --url https://console.jumpcloud.com/api/systems/${systemKey}
```

### Output Data

All results will be formatted as JSON.

Here is an abbreviated example of response output:

```json
{
  \"_id\": \"525ee96f52e144993e000015\",
  \"agentServer\": \"lappy386\",
  \"agentVersion\": \"0.9.42\",
  \"arch\": \"x86_64\",
  \"connectionKey\": \"127.0.0.1_51812\",
  \"displayName\": \"ubuntu-1204\",
  \"firstContact\": \"2013-10-16T19:30:55.611Z\",
  \"hostname\": \"ubuntu-1204\"
  ...
```

## Additional Examples

### Signing Authentication Example

This example demonstrates how to make an authenticated request to fetch the JumpCloud record for this system.

[SigningExample.sh](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/shell/SigningExample.sh)

### Shutdown Hook

This example demonstrates how to make an authenticated request on system shutdown.
Using an init.d script registered at run level 0, you can call the System Context API as the system is shutting down.

[Instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) is an example of an init.d script that only runs at system shutdown.

After customizing the [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) script, you should install it on the system(s) running the JumpCloud agent.

1. Copy the modified [instance-shutdown-initd](https://github.com/TheJumpCloud/SystemContextAPI/blob/master/examples/instance-shutdown-initd) to `/etc/init.d/instance-shutdown`.
2. On Ubuntu systems, run `update-rc.d instance-shutdown defaults`. On RedHat/CentOS systems, run `chkconfig --add instance-shutdown`.

## Third Party

### Chef Cookbooks

[https://github.com/nshenry03/jumpcloud](https://github.com/nshenry03/jumpcloud)

[https://github.com/cjs226/jumpcloud](https://github.com/cjs226/jumpcloud)

# Multi-Tenant Portal Headers

Multi-Tenant Organization API Headers are available for JumpCloud Admins to use when making API requests from Organizations that have multiple managed organizations.

The `x-org-id` is a required header for all multi-tenant admins when making API requests to JumpCloud. This header will define to which organization you would like to make the request.

**NOTE** Single Tenant Admins do not need to provide this header when making an API request.

## Header Value

`x-org-id`

## API Response Codes

* `400` Malformed ID.
* `400` x-org-id and Organization path ID do not match.
* `401` ID not included for multi-tenant admin
* `403` ID included on unsupported route.
* `404` Organization ID Not Found.

```bash
curl -X GET https://console.jumpcloud.com/api/v2/directories \\
  -H 'accept: application/json' \\
  -H 'content-type: application/json' \\
  -H 'x-api-key: {API_KEY}' \\
  -H 'x-org-id: {ORG_ID}'

```

## To Obtain an Individual Organization ID via the UI

As a prerequisite, your Primary Organization will need to be setup for Multi-Tenancy. This provides access to the Multi-Tenant Organization Admin Portal.

1. Log into JumpCloud [Admin Console](https://console.jumpcloud.com). If you are a multi-tenant Admin, you will automatically be routed to the Multi-Tenant Admin Portal.
2. From the Multi-Tenant Portal's primary navigation bar, select the Organization you'd like to access.
3. You will automatically be routed to that Organization's Admin Console.
4. Go to Settings in the sub-tenant's primary navigation.
5. You can obtain your Organization ID below your Organization's Contact Information on the Settings page.

## To Obtain All Organization IDs via the API

* You can make an API request to this endpoint using the API key of your Primary Organization.  `https://console.jumpcloud.com/api/organizations/` This will return all your managed organizations.

```bash
curl -X GET \\
  https://console.jumpcloud.com/api/organizations/ \\
  -H 'Accept: application/json' \\
  -H 'Content-Type: application/json' \\
  -H 'x-api-key: {API_KEY}'
```

# SDKs

You can find language specific SDKs that can help you kickstart your Integration with JumpCloud in the following GitHub repositories:

* [Python](https://github.com/TheJumpCloud/jcapi-python)
* [Go](https://github.com/TheJumpCloud/jcapi-go)
* [Ruby](https://github.com/TheJumpCloud/jcapi-ruby)
* [Java](https://github.com/TheJumpCloud/jcapi-java)


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://support.jumpcloud.com/support/s/](https://support.jumpcloud.com/support/s/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import jcapi1 "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), jcapi1.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), jcapi1.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), jcapi1.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), jcapi1.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://console.jumpcloud.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationTemplatesApi* | [**ApplicationTemplatesGet**](docs/ApplicationTemplatesApi.md#applicationtemplatesget) | **Get** /application-templates/{id} | Get an Application Template
*ApplicationTemplatesApi* | [**ApplicationTemplatesList**](docs/ApplicationTemplatesApi.md#applicationtemplateslist) | **Get** /application-templates | List Application Templates
*ApplicationsApi* | [**ApplicationsDelete**](docs/ApplicationsApi.md#applicationsdelete) | **Delete** /applications/{id} | Delete an Application
*ApplicationsApi* | [**ApplicationsGet**](docs/ApplicationsApi.md#applicationsget) | **Get** /applications/{id} | Get an Application
*ApplicationsApi* | [**ApplicationsList**](docs/ApplicationsApi.md#applicationslist) | **Get** /applications | Applications
*ApplicationsApi* | [**ApplicationsPost**](docs/ApplicationsApi.md#applicationspost) | **Post** /applications | Create an Application
*ApplicationsApi* | [**ApplicationsPut**](docs/ApplicationsApi.md#applicationsput) | **Put** /applications/{id} | Update an Application
*CommandResultsApi* | [**CommandResultsDelete**](docs/CommandResultsApi.md#commandresultsdelete) | **Delete** /commandresults/{id} | Delete a Command result
*CommandResultsApi* | [**CommandResultsGet**](docs/CommandResultsApi.md#commandresultsget) | **Get** /commandresults/{id} | List an individual Command result
*CommandResultsApi* | [**CommandResultsList**](docs/CommandResultsApi.md#commandresultslist) | **Get** /commandresults | List all Command Results
*CommandTriggersApi* | [**CommandTriggerWebhookPost**](docs/CommandTriggersApi.md#commandtriggerwebhookpost) | **Post** /command/trigger/{triggername} | Launch a command via a Trigger
*CommandsApi* | [**CommandFileGet**](docs/CommandsApi.md#commandfileget) | **Get** /files/command/{id} | Get a Command File
*CommandsApi* | [**CommandsDelete**](docs/CommandsApi.md#commandsdelete) | **Delete** /commands/{id} | Delete a Command
*CommandsApi* | [**CommandsGet**](docs/CommandsApi.md#commandsget) | **Get** /commands/{id} | List an individual Command
*CommandsApi* | [**CommandsGetResults**](docs/CommandsApi.md#commandsgetresults) | **Get** /commands/{id}/results | Get results for a specific command
*CommandsApi* | [**CommandsList**](docs/CommandsApi.md#commandslist) | **Get** /commands | List All Commands
*CommandsApi* | [**CommandsPost**](docs/CommandsApi.md#commandspost) | **Post** /commands | Create A Command
*CommandsApi* | [**CommandsPut**](docs/CommandsApi.md#commandsput) | **Put** /commands/{id} | Update a Command
*ManagedServiceProviderApi* | [**AdminTotpresetBegin**](docs/ManagedServiceProviderApi.md#admintotpresetbegin) | **Post** /users/resettotp/{id} | Administrator TOTP Reset Initiation
*ManagedServiceProviderApi* | [**OrganizationList**](docs/ManagedServiceProviderApi.md#organizationlist) | **Get** /organizations | Get Organization Details
*ManagedServiceProviderApi* | [**UsersPut**](docs/ManagedServiceProviderApi.md#usersput) | **Put** /users/{id} | Update a user
*ManagedServiceProviderApi* | [**UsersReactivateGet**](docs/ManagedServiceProviderApi.md#usersreactivateget) | **Get** /users/reactivate/{id} | Administrator Password Reset Initiation
*OrganizationsApi* | [**OrganizationList**](docs/OrganizationsApi.md#organizationlist) | **Get** /organizations | Get Organization Details
*OrganizationsApi* | [**OrganizationPut**](docs/OrganizationsApi.md#organizationput) | **Put** /organizations/{id} | Update an Organization
*OrganizationsApi* | [**OrganizationsGet**](docs/OrganizationsApi.md#organizationsget) | **Get** /organizations/{id} | Get an Organization
*RadiusServersApi* | [**RadiusServersDelete**](docs/RadiusServersApi.md#radiusserversdelete) | **Delete** /radiusservers/{id} | Delete Radius Server
*RadiusServersApi* | [**RadiusServersGet**](docs/RadiusServersApi.md#radiusserversget) | **Get** /radiusservers/{id} | Get Radius Server
*RadiusServersApi* | [**RadiusServersList**](docs/RadiusServersApi.md#radiusserverslist) | **Get** /radiusservers | List Radius Servers
*RadiusServersApi* | [**RadiusServersPost**](docs/RadiusServersApi.md#radiusserverspost) | **Post** /radiusservers | Create a Radius Server
*RadiusServersApi* | [**RadiusServersPut**](docs/RadiusServersApi.md#radiusserversput) | **Put** /radiusservers/{id} | Update Radius Servers
*SearchApi* | [**SearchCommandresultsPost**](docs/SearchApi.md#searchcommandresultspost) | **Post** /search/commandresults | Search Commands Results
*SearchApi* | [**SearchCommandsPost**](docs/SearchApi.md#searchcommandspost) | **Post** /search/commands | Search Commands
*SearchApi* | [**SearchOrganizationsPost**](docs/SearchApi.md#searchorganizationspost) | **Post** /search/organizations | Search Organizations
*SearchApi* | [**SearchSystemsPost**](docs/SearchApi.md#searchsystemspost) | **Post** /search/systems | Search Systems
*SearchApi* | [**SearchSystemusersPost**](docs/SearchApi.md#searchsystemuserspost) | **Post** /search/systemusers | Search System Users
*SystemsApi* | [**SystemsCommandBuiltinErase**](docs/SystemsApi.md#systemscommandbuiltinerase) | **Post** /systems/{system_id}/command/builtin/erase | Erase a System
*SystemsApi* | [**SystemsCommandBuiltinLock**](docs/SystemsApi.md#systemscommandbuiltinlock) | **Post** /systems/{system_id}/command/builtin/lock | Lock a System
*SystemsApi* | [**SystemsCommandBuiltinRestart**](docs/SystemsApi.md#systemscommandbuiltinrestart) | **Post** /systems/{system_id}/command/builtin/restart | Restart a System
*SystemsApi* | [**SystemsCommandBuiltinShutdown**](docs/SystemsApi.md#systemscommandbuiltinshutdown) | **Post** /systems/{system_id}/command/builtin/shutdown | Shutdown a System
*SystemsApi* | [**SystemsDelete**](docs/SystemsApi.md#systemsdelete) | **Delete** /systems/{id} | Delete a System
*SystemsApi* | [**SystemsGet**](docs/SystemsApi.md#systemsget) | **Get** /systems/{id} | List an individual system
*SystemsApi* | [**SystemsList**](docs/SystemsApi.md#systemslist) | **Get** /systems | List All Systems
*SystemsApi* | [**SystemsPut**](docs/SystemsApi.md#systemsput) | **Put** /systems/{id} | Update a system
*SystemusersApi* | [**SshkeyDelete**](docs/SystemusersApi.md#sshkeydelete) | **Delete** /systemusers/{systemuser_id}/sshkeys/{id} | Delete a system user&#39;s Public SSH Keys
*SystemusersApi* | [**SshkeyList**](docs/SystemusersApi.md#sshkeylist) | **Get** /systemusers/{id}/sshkeys | List a system user&#39;s public SSH keys
*SystemusersApi* | [**SshkeyPost**](docs/SystemusersApi.md#sshkeypost) | **Post** /systemusers/{id}/sshkeys | Create a system user&#39;s Public SSH Key
*SystemusersApi* | [**SystemusersDelete**](docs/SystemusersApi.md#systemusersdelete) | **Delete** /systemusers/{id} | Delete a system user
*SystemusersApi* | [**SystemusersExpire**](docs/SystemusersApi.md#systemusersexpire) | **Post** /systemusers/{id}/expire | Expire a system user&#39;s password
*SystemusersApi* | [**SystemusersGet**](docs/SystemusersApi.md#systemusersget) | **Get** /systemusers/{id} | List a system user
*SystemusersApi* | [**SystemusersList**](docs/SystemusersApi.md#systemuserslist) | **Get** /systemusers | List all system users
*SystemusersApi* | [**SystemusersMfasync**](docs/SystemusersApi.md#systemusersmfasync) | **Post** /systemusers/{id}/mfasync | Sync a systemuser&#39;s mfa enrollment status
*SystemusersApi* | [**SystemusersPost**](docs/SystemusersApi.md#systemuserspost) | **Post** /systemusers | Create a system user
*SystemusersApi* | [**SystemusersPut**](docs/SystemusersApi.md#systemusersput) | **Put** /systemusers/{id} | Update a system user
*SystemusersApi* | [**SystemusersResetmfa**](docs/SystemusersApi.md#systemusersresetmfa) | **Post** /systemusers/{id}/resetmfa | Reset a system user&#39;s MFA token
*SystemusersApi* | [**SystemusersStateActivate**](docs/SystemusersApi.md#systemusersstateactivate) | **Post** /systemusers/{id}/state/activate | Activate System User
*SystemusersApi* | [**SystemusersUnlock**](docs/SystemusersApi.md#systemusersunlock) | **Post** /systemusers/{id}/unlock | Unlock a system user
*UsersApi* | [**AdminTotpresetBegin**](docs/UsersApi.md#admintotpresetbegin) | **Post** /users/resettotp/{id} | Administrator TOTP Reset Initiation
*UsersApi* | [**UsersPut**](docs/UsersApi.md#usersput) | **Put** /users/{id} | Update a user
*UsersApi* | [**UsersReactivateGet**](docs/UsersApi.md#usersreactivateget) | **Get** /users/reactivate/{id} | Administrator Password Reset Initiation


## Documentation For Models

 - [Application](docs/Application.md)
 - [ApplicationConfig](docs/ApplicationConfig.md)
 - [ApplicationConfigAcsUrl](docs/ApplicationConfigAcsUrl.md)
 - [ApplicationConfigAcsUrlTooltip](docs/ApplicationConfigAcsUrlTooltip.md)
 - [ApplicationConfigAcsUrlTooltipVariables](docs/ApplicationConfigAcsUrlTooltipVariables.md)
 - [ApplicationConfigConstantAttributes](docs/ApplicationConfigConstantAttributes.md)
 - [ApplicationConfigConstantAttributesValueInner](docs/ApplicationConfigConstantAttributesValueInner.md)
 - [ApplicationConfigDatabaseAttributes](docs/ApplicationConfigDatabaseAttributes.md)
 - [ApplicationLogo](docs/ApplicationLogo.md)
 - [Applicationslist](docs/Applicationslist.md)
 - [Applicationtemplate](docs/Applicationtemplate.md)
 - [ApplicationtemplateJit](docs/ApplicationtemplateJit.md)
 - [ApplicationtemplateLogo](docs/ApplicationtemplateLogo.md)
 - [ApplicationtemplateOidc](docs/ApplicationtemplateOidc.md)
 - [ApplicationtemplateProvision](docs/ApplicationtemplateProvision.md)
 - [Applicationtemplateslist](docs/Applicationtemplateslist.md)
 - [Command](docs/Command.md)
 - [Commandfilereturn](docs/Commandfilereturn.md)
 - [CommandfilereturnResultsInner](docs/CommandfilereturnResultsInner.md)
 - [Commandresult](docs/Commandresult.md)
 - [CommandresultResponse](docs/CommandresultResponse.md)
 - [CommandresultResponseData](docs/CommandresultResponseData.md)
 - [Commandresultslist](docs/Commandresultslist.md)
 - [CommandresultslistResultsInner](docs/CommandresultslistResultsInner.md)
 - [Commandslist](docs/Commandslist.md)
 - [CommandslistResultsInner](docs/CommandslistResultsInner.md)
 - [Error](docs/Error.md)
 - [ErrorDetails](docs/ErrorDetails.md)
 - [ErrorDetailsAllOf](docs/ErrorDetailsAllOf.md)
 - [Fde](docs/Fde.md)
 - [Mfa](docs/Mfa.md)
 - [MfaEnrollment](docs/MfaEnrollment.md)
 - [MfaEnrollmentStatus](docs/MfaEnrollmentStatus.md)
 - [Organization](docs/Organization.md)
 - [OrganizationPutRequest](docs/OrganizationPutRequest.md)
 - [Organizationentitlement](docs/Organizationentitlement.md)
 - [OrganizationentitlementEntitlementProductsInner](docs/OrganizationentitlementEntitlementProductsInner.md)
 - [Organizationsettings](docs/Organizationsettings.md)
 - [OrganizationsettingsFeatures](docs/OrganizationsettingsFeatures.md)
 - [OrganizationsettingsFeaturesDirectoryInsights](docs/OrganizationsettingsFeaturesDirectoryInsights.md)
 - [OrganizationsettingsFeaturesDirectoryInsightsPremium](docs/OrganizationsettingsFeaturesDirectoryInsightsPremium.md)
 - [OrganizationsettingsFeaturesSystemInsights](docs/OrganizationsettingsFeaturesSystemInsights.md)
 - [OrganizationsettingsNewSystemUserStateDefaults](docs/OrganizationsettingsNewSystemUserStateDefaults.md)
 - [OrganizationsettingsPasswordPolicy](docs/OrganizationsettingsPasswordPolicy.md)
 - [OrganizationsettingsUserPortal](docs/OrganizationsettingsUserPortal.md)
 - [OrganizationsettingsWindowsMDM](docs/OrganizationsettingsWindowsMDM.md)
 - [Organizationsettingsput](docs/Organizationsettingsput.md)
 - [OrganizationsettingsputNewSystemUserStateDefaults](docs/OrganizationsettingsputNewSystemUserStateDefaults.md)
 - [OrganizationsettingsputPasswordPolicy](docs/OrganizationsettingsputPasswordPolicy.md)
 - [Organizationslist](docs/Organizationslist.md)
 - [OrganizationslistResultsInner](docs/OrganizationslistResultsInner.md)
 - [RadiusServersPutRequest](docs/RadiusServersPutRequest.md)
 - [Radiusserver](docs/Radiusserver.md)
 - [Radiusserverpost](docs/Radiusserverpost.md)
 - [Radiusserverput](docs/Radiusserverput.md)
 - [Radiusserverslist](docs/Radiusserverslist.md)
 - [Search](docs/Search.md)
 - [Sshkeylist](docs/Sshkeylist.md)
 - [Sshkeypost](docs/Sshkeypost.md)
 - [Sso](docs/Sso.md)
 - [System](docs/System.md)
 - [SystemBuiltInCommandsInner](docs/SystemBuiltInCommandsInner.md)
 - [SystemDomainInfo](docs/SystemDomainInfo.md)
 - [SystemMdm](docs/SystemMdm.md)
 - [SystemMdmInternal](docs/SystemMdmInternal.md)
 - [SystemNetworkInterfacesInner](docs/SystemNetworkInterfacesInner.md)
 - [SystemOsVersionDetail](docs/SystemOsVersionDetail.md)
 - [SystemProvisionMetadata](docs/SystemProvisionMetadata.md)
 - [SystemProvisionMetadataProvisioner](docs/SystemProvisionMetadataProvisioner.md)
 - [SystemServiceAccountState](docs/SystemServiceAccountState.md)
 - [SystemSshdParamsInner](docs/SystemSshdParamsInner.md)
 - [SystemSystemInsights](docs/SystemSystemInsights.md)
 - [SystemUserMetricsInner](docs/SystemUserMetricsInner.md)
 - [Systemput](docs/Systemput.md)
 - [SystemputAgentBoundMessagesInner](docs/SystemputAgentBoundMessagesInner.md)
 - [Systemslist](docs/Systemslist.md)
 - [Systemuserput](docs/Systemuserput.md)
 - [SystemuserputAddressesInner](docs/SystemuserputAddressesInner.md)
 - [SystemuserputAttributesInner](docs/SystemuserputAttributesInner.md)
 - [SystemuserputPhoneNumbersInner](docs/SystemuserputPhoneNumbersInner.md)
 - [SystemuserputRelationshipsInner](docs/SystemuserputRelationshipsInner.md)
 - [Systemuserputpost](docs/Systemuserputpost.md)
 - [SystemuserputpostAddressesInner](docs/SystemuserputpostAddressesInner.md)
 - [SystemuserputpostPhoneNumbersInner](docs/SystemuserputpostPhoneNumbersInner.md)
 - [SystemuserputpostRecoveryEmail](docs/SystemuserputpostRecoveryEmail.md)
 - [Systemuserreturn](docs/Systemuserreturn.md)
 - [SystemuserreturnAddressesInner](docs/SystemuserreturnAddressesInner.md)
 - [SystemuserreturnPhoneNumbersInner](docs/SystemuserreturnPhoneNumbersInner.md)
 - [SystemuserreturnRecoveryEmail](docs/SystemuserreturnRecoveryEmail.md)
 - [SystemusersResetmfaRequest](docs/SystemusersResetmfaRequest.md)
 - [SystemusersStateActivateRequest](docs/SystemusersStateActivateRequest.md)
 - [Systemuserslist](docs/Systemuserslist.md)
 - [Triggerreturn](docs/Triggerreturn.md)
 - [TrustedappConfigGet](docs/TrustedappConfigGet.md)
 - [TrustedappConfigGetTrustedAppsInner](docs/TrustedappConfigGetTrustedAppsInner.md)
 - [TrustedappConfigPut](docs/TrustedappConfigPut.md)
 - [Userput](docs/Userput.md)
 - [Userreturn](docs/Userreturn.md)
 - [UserreturnGrowthData](docs/UserreturnGrowthData.md)


## Documentation For Authorization



### x-api-key

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: x-api-key and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@jumpcloud.com


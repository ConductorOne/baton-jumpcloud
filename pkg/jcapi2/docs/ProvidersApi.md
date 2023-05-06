# \ProvidersApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutotaskCreateConfiguration**](ProvidersApi.md#AutotaskCreateConfiguration) | **Post** /providers/{provider_id}/integrations/autotask | Creates a new Autotask integration for the provider
[**AutotaskDeleteConfiguration**](ProvidersApi.md#AutotaskDeleteConfiguration) | **Delete** /integrations/autotask/{UUID} | Delete Autotask Integration
[**AutotaskGetConfiguration**](ProvidersApi.md#AutotaskGetConfiguration) | **Get** /integrations/autotask/{UUID} | Retrieve Autotask Integration Configuration
[**AutotaskPatchMappings**](ProvidersApi.md#AutotaskPatchMappings) | **Patch** /integrations/autotask/{UUID}/mappings | Create, edit, and/or delete Autotask Mappings
[**AutotaskPatchSettings**](ProvidersApi.md#AutotaskPatchSettings) | **Patch** /integrations/autotask/{UUID}/settings | Create, edit, and/or delete Autotask Integration settings
[**AutotaskRetrieveAllAlertConfigurationOptions**](ProvidersApi.md#AutotaskRetrieveAllAlertConfigurationOptions) | **Get** /providers/{provider_id}/integrations/autotask/alerts/configuration/options | Get all Autotask ticketing alert configuration options for a provider
[**AutotaskRetrieveAllAlertConfigurations**](ProvidersApi.md#AutotaskRetrieveAllAlertConfigurations) | **Get** /providers/{provider_id}/integrations/autotask/alerts/configuration | Get all Autotask ticketing alert configurations for a provider
[**AutotaskRetrieveCompanies**](ProvidersApi.md#AutotaskRetrieveCompanies) | **Get** /integrations/autotask/{UUID}/companies | Retrieve Autotask Companies
[**AutotaskRetrieveCompanyTypes**](ProvidersApi.md#AutotaskRetrieveCompanyTypes) | **Get** /integrations/autotask/{UUID}/companytypes | Retrieve Autotask Company Types
[**AutotaskRetrieveContracts**](ProvidersApi.md#AutotaskRetrieveContracts) | **Get** /integrations/autotask/{UUID}/contracts | Retrieve Autotask Contracts
[**AutotaskRetrieveContractsFields**](ProvidersApi.md#AutotaskRetrieveContractsFields) | **Get** /integrations/autotask/{UUID}/contracts/fields | Retrieve Autotask Contract Fields
[**AutotaskRetrieveMappings**](ProvidersApi.md#AutotaskRetrieveMappings) | **Get** /integrations/autotask/{UUID}/mappings | Retrieve Autotask mappings
[**AutotaskRetrieveServices**](ProvidersApi.md#AutotaskRetrieveServices) | **Get** /integrations/autotask/{UUID}/contracts/services | Retrieve Autotask Contract Services
[**AutotaskRetrieveSettings**](ProvidersApi.md#AutotaskRetrieveSettings) | **Get** /integrations/autotask/{UUID}/settings | Retrieve Autotask Integration settings
[**AutotaskUpdateAlertConfiguration**](ProvidersApi.md#AutotaskUpdateAlertConfiguration) | **Put** /providers/{provider_id}/integrations/autotask/alerts/{alert_UUID}/configuration | Update an Autotask ticketing alert&#39;s configuration
[**AutotaskUpdateConfiguration**](ProvidersApi.md#AutotaskUpdateConfiguration) | **Patch** /integrations/autotask/{UUID} | Update Autotask Integration configuration
[**ConnectwiseCreateConfiguration**](ProvidersApi.md#ConnectwiseCreateConfiguration) | **Post** /providers/{provider_id}/integrations/connectwise | Creates a new ConnectWise integration for the provider
[**ConnectwiseDeleteConfiguration**](ProvidersApi.md#ConnectwiseDeleteConfiguration) | **Delete** /integrations/connectwise/{UUID} | Delete ConnectWise Integration
[**ConnectwiseGetConfiguration**](ProvidersApi.md#ConnectwiseGetConfiguration) | **Get** /integrations/connectwise/{UUID} | Retrieve ConnectWise Integration Configuration
[**ConnectwisePatchMappings**](ProvidersApi.md#ConnectwisePatchMappings) | **Patch** /integrations/connectwise/{UUID}/mappings | Create, edit, and/or delete ConnectWise Mappings
[**ConnectwisePatchSettings**](ProvidersApi.md#ConnectwisePatchSettings) | **Patch** /integrations/connectwise/{UUID}/settings | Create, edit, and/or delete ConnectWise Integration settings
[**ConnectwiseRetrieveAdditions**](ProvidersApi.md#ConnectwiseRetrieveAdditions) | **Get** /integrations/connectwise/{UUID}/agreements/{agreement_ID}/additions | Retrieve ConnectWise Additions
[**ConnectwiseRetrieveAgreements**](ProvidersApi.md#ConnectwiseRetrieveAgreements) | **Get** /integrations/connectwise/{UUID}/agreements | Retrieve ConnectWise Agreements
[**ConnectwiseRetrieveAllAlertConfigurationOptions**](ProvidersApi.md#ConnectwiseRetrieveAllAlertConfigurationOptions) | **Get** /providers/{provider_id}/integrations/connectwise/alerts/configuration/options | Get all ConnectWise ticketing alert configuration options for a provider
[**ConnectwiseRetrieveAllAlertConfigurations**](ProvidersApi.md#ConnectwiseRetrieveAllAlertConfigurations) | **Get** /providers/{provider_id}/integrations/connectwise/alerts/configuration | Get all ConnectWise ticketing alert configurations for a provider
[**ConnectwiseRetrieveCompanies**](ProvidersApi.md#ConnectwiseRetrieveCompanies) | **Get** /integrations/connectwise/{UUID}/companies | Retrieve ConnectWise Companies
[**ConnectwiseRetrieveCompanyTypes**](ProvidersApi.md#ConnectwiseRetrieveCompanyTypes) | **Get** /integrations/connectwise/{UUID}/companytypes | Retrieve ConnectWise Company Types
[**ConnectwiseRetrieveMappings**](ProvidersApi.md#ConnectwiseRetrieveMappings) | **Get** /integrations/connectwise/{UUID}/mappings | Retrieve ConnectWise mappings
[**ConnectwiseRetrieveSettings**](ProvidersApi.md#ConnectwiseRetrieveSettings) | **Get** /integrations/connectwise/{UUID}/settings | Retrieve ConnectWise Integration settings
[**ConnectwiseUpdateAlertConfiguration**](ProvidersApi.md#ConnectwiseUpdateAlertConfiguration) | **Put** /providers/{provider_id}/integrations/connectwise/alerts/{alert_UUID}/configuration | Update a ConnectWise ticketing alert&#39;s configuration
[**ConnectwiseUpdateConfiguration**](ProvidersApi.md#ConnectwiseUpdateConfiguration) | **Patch** /integrations/connectwise/{UUID} | Update ConnectWise Integration configuration
[**MtpIntegrationRetrieveAlerts**](ProvidersApi.md#MtpIntegrationRetrieveAlerts) | **Get** /providers/{provider_id}/integrations/ticketing/alerts | Get all ticketing alerts available for a provider&#39;s ticketing integration.
[**MtpIntegrationRetrieveSyncErrors**](ProvidersApi.md#MtpIntegrationRetrieveSyncErrors) | **Get** /integrations/{integration_type}/{UUID}/errors | Retrieve Recent Integration Sync Errors
[**PolicyGroupTemplatesDelete**](ProvidersApi.md#PolicyGroupTemplatesDelete) | **Delete** /providers/{provider_id}/policygrouptemplates/{id} | Deletes policy group template.
[**PolicyGroupTemplatesGet**](ProvidersApi.md#PolicyGroupTemplatesGet) | **Get** /providers/{provider_id}/policygrouptemplates/{id} | Gets a provider&#39;s policy group template.
[**PolicyGroupTemplatesGetConfiguredPolicyTemplate**](ProvidersApi.md#PolicyGroupTemplatesGetConfiguredPolicyTemplate) | **Get** /providers/{provider_id}/configuredpolicytemplates/{id} | Retrieve a configured policy template by id.
[**PolicyGroupTemplatesList**](ProvidersApi.md#PolicyGroupTemplatesList) | **Get** /providers/{provider_id}/policygrouptemplates | List a provider&#39;s policy group templates.
[**PolicyGroupTemplatesListConfiguredPolicyTemplates**](ProvidersApi.md#PolicyGroupTemplatesListConfiguredPolicyTemplates) | **Get** /providers/{provider_id}/configuredpolicytemplates | List a provider&#39;s configured policy templates.
[**PolicyGroupTemplatesListMembers**](ProvidersApi.md#PolicyGroupTemplatesListMembers) | **Get** /providers/{provider_id}/policygrouptemplates/{id}/members | Gets the list of members from a policy group template.
[**ProviderOrganizationsCreateOrg**](ProvidersApi.md#ProviderOrganizationsCreateOrg) | **Post** /providers/{provider_id}/organizations | Create Provider Organization
[**ProviderOrganizationsUpdateOrg**](ProvidersApi.md#ProviderOrganizationsUpdateOrg) | **Put** /providers/{provider_id}/organizations/{id} | Update Provider Organization
[**ProvidersGetProvider**](ProvidersApi.md#ProvidersGetProvider) | **Get** /providers/{provider_id} | Retrieve Provider
[**ProvidersListAdministrators**](ProvidersApi.md#ProvidersListAdministrators) | **Get** /providers/{provider_id}/administrators | List Provider Administrators
[**ProvidersListOrganizations**](ProvidersApi.md#ProvidersListOrganizations) | **Get** /providers/{provider_id}/organizations | List Provider Organizations
[**ProvidersPostAdmins**](ProvidersApi.md#ProvidersPostAdmins) | **Post** /providers/{provider_id}/administrators | Create a new Provider Administrator
[**ProvidersRemoveAdministrator**](ProvidersApi.md#ProvidersRemoveAdministrator) | **Delete** /providers/{provider_id}/administrators/{id} | Delete Provider Administrator
[**ProvidersRetrieveIntegrations**](ProvidersApi.md#ProvidersRetrieveIntegrations) | **Get** /providers/{provider_id}/integrations | Retrieve Integrations for Provider
[**ProvidersRetrieveInvoice**](ProvidersApi.md#ProvidersRetrieveInvoice) | **Get** /providers/{provider_id}/invoices/{ID} | Download a provider&#39;s invoice.
[**ProvidersRetrieveInvoices**](ProvidersApi.md#ProvidersRetrieveInvoices) | **Get** /providers/{provider_id}/invoices | List a provider&#39;s invoices.
[**SyncroCreateConfiguration**](ProvidersApi.md#SyncroCreateConfiguration) | **Post** /providers/{provider_id}/integrations/syncro | Creates a new Syncro integration for the provider
[**SyncroDeleteConfiguration**](ProvidersApi.md#SyncroDeleteConfiguration) | **Delete** /integrations/syncro/{UUID} | Delete Syncro Integration
[**SyncroGetConfiguration**](ProvidersApi.md#SyncroGetConfiguration) | **Get** /integrations/syncro/{UUID} | Retrieve Syncro Integration Configuration
[**SyncroPatchMappings**](ProvidersApi.md#SyncroPatchMappings) | **Patch** /integrations/syncro/{UUID}/mappings | Create, edit, and/or delete Syncro Mappings
[**SyncroPatchSettings**](ProvidersApi.md#SyncroPatchSettings) | **Patch** /integrations/syncro/{UUID}/settings | Create, edit, and/or delete Syncro Integration settings
[**SyncroRetrieveAllAlertConfigurationOptions**](ProvidersApi.md#SyncroRetrieveAllAlertConfigurationOptions) | **Get** /providers/{provider_id}/integrations/syncro/alerts/configuration/options | Get all Syncro ticketing alert configuration options for a provider
[**SyncroRetrieveAllAlertConfigurations**](ProvidersApi.md#SyncroRetrieveAllAlertConfigurations) | **Get** /providers/{provider_id}/integrations/syncro/alerts/configuration | Get all Syncro ticketing alert configurations for a provider
[**SyncroRetrieveBillingMappingConfigurationOptions**](ProvidersApi.md#SyncroRetrieveBillingMappingConfigurationOptions) | **Get** /integrations/syncro/{UUID}/billing_mapping_configuration_options | Retrieve Syncro billing mappings dependencies
[**SyncroRetrieveCompanies**](ProvidersApi.md#SyncroRetrieveCompanies) | **Get** /integrations/syncro/{UUID}/companies | Retrieve Syncro Companies
[**SyncroRetrieveMappings**](ProvidersApi.md#SyncroRetrieveMappings) | **Get** /integrations/syncro/{UUID}/mappings | Retrieve Syncro mappings
[**SyncroRetrieveSettings**](ProvidersApi.md#SyncroRetrieveSettings) | **Get** /integrations/syncro/{UUID}/settings | Retrieve Syncro Integration settings
[**SyncroUpdateAlertConfiguration**](ProvidersApi.md#SyncroUpdateAlertConfiguration) | **Put** /providers/{provider_id}/integrations/syncro/alerts/{alert_UUID}/configuration | Update a Syncro ticketing alert&#39;s configuration
[**SyncroUpdateConfiguration**](ProvidersApi.md#SyncroUpdateConfiguration) | **Patch** /integrations/syncro/{UUID} | Update Syncro Integration configuration



## AutotaskCreateConfiguration

> AutotaskCreateConfiguration201Response AutotaskCreateConfiguration(ctx, providerId).Body(body).Execute()

Creates a new Autotask integration for the provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    body := *openapiclient.NewAutotaskIntegrationReq("Secret_example", "Username_example") // AutotaskIntegrationReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskCreateConfiguration(context.Background(), providerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskCreateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskCreateConfiguration`: AutotaskCreateConfiguration201Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskCreateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskCreateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AutotaskIntegrationReq**](AutotaskIntegrationReq.md) |  | 

### Return type

[**AutotaskCreateConfiguration201Response**](AutotaskCreateConfiguration201Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskDeleteConfiguration

> AutotaskDeleteConfiguration(ctx, uUID).Execute()

Delete Autotask Integration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProvidersApi.AutotaskDeleteConfiguration(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskDeleteConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskDeleteConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskGetConfiguration

> AutotaskIntegration AutotaskGetConfiguration(ctx, uUID).Execute()

Retrieve Autotask Integration Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskGetConfiguration(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskGetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskGetConfiguration`: AutotaskIntegration
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskGetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutotaskIntegration**](AutotaskIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskPatchMappings

> AutotaskMappingResponse AutotaskPatchMappings(ctx, uUID).Body(body).Execute()

Create, edit, and/or delete Autotask Mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    body := *openapiclient.NewAutotaskMappingRequest() // AutotaskMappingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskPatchMappings(context.Background(), uUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskPatchMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskPatchMappings`: AutotaskMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskPatchMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskPatchMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AutotaskMappingRequest**](AutotaskMappingRequest.md) |  | 

### Return type

[**AutotaskMappingResponse**](AutotaskMappingResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskPatchSettings

> AutotaskSettings AutotaskPatchSettings(ctx, uUID).Body(body).Execute()

Create, edit, and/or delete Autotask Integration settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    body := *openapiclient.NewAutotaskSettingsPatchReq() // AutotaskSettingsPatchReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskPatchSettings(context.Background(), uUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskPatchSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskPatchSettings`: AutotaskSettings
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskPatchSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskPatchSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AutotaskSettingsPatchReq**](AutotaskSettingsPatchReq.md) |  | 

### Return type

[**AutotaskSettings**](AutotaskSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskRetrieveAllAlertConfigurationOptions

> AutotaskTicketingAlertConfigurationOptions AutotaskRetrieveAllAlertConfigurationOptions(ctx, providerId).Execute()

Get all Autotask ticketing alert configuration options for a provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskRetrieveAllAlertConfigurationOptions(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskRetrieveAllAlertConfigurationOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskRetrieveAllAlertConfigurationOptions`: AutotaskTicketingAlertConfigurationOptions
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskRetrieveAllAlertConfigurationOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskRetrieveAllAlertConfigurationOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutotaskTicketingAlertConfigurationOptions**](AutotaskTicketingAlertConfigurationOptions.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskRetrieveAllAlertConfigurations

> AutotaskTicketingAlertConfigurationList AutotaskRetrieveAllAlertConfigurations(ctx, providerId).Execute()

Get all Autotask ticketing alert configurations for a provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskRetrieveAllAlertConfigurations(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskRetrieveAllAlertConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskRetrieveAllAlertConfigurations`: AutotaskTicketingAlertConfigurationList
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskRetrieveAllAlertConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskRetrieveAllAlertConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutotaskTicketingAlertConfigurationList**](AutotaskTicketingAlertConfigurationList.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskRetrieveCompanies

> AutotaskCompanyResp AutotaskRetrieveCompanies(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve Autotask Companies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskRetrieveCompanies(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskRetrieveCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskRetrieveCompanies`: AutotaskCompanyResp
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskRetrieveCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskRetrieveCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**AutotaskCompanyResp**](AutotaskCompanyResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskRetrieveCompanyTypes

> AutotaskCompanyTypeResp AutotaskRetrieveCompanyTypes(ctx, uUID).Execute()

Retrieve Autotask Company Types



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskRetrieveCompanyTypes(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskRetrieveCompanyTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskRetrieveCompanyTypes`: AutotaskCompanyTypeResp
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskRetrieveCompanyTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskRetrieveCompanyTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutotaskCompanyTypeResp**](AutotaskCompanyTypeResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskRetrieveContracts

> AutotaskRetrieveContracts200Response AutotaskRetrieveContracts(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve Autotask Contracts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskRetrieveContracts(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskRetrieveContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskRetrieveContracts`: AutotaskRetrieveContracts200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskRetrieveContracts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskRetrieveContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**AutotaskRetrieveContracts200Response**](AutotaskRetrieveContracts200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskRetrieveContractsFields

> AutotaskRetrieveContractsFields200Response AutotaskRetrieveContractsFields(ctx, uUID).Execute()

Retrieve Autotask Contract Fields



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskRetrieveContractsFields(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskRetrieveContractsFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskRetrieveContractsFields`: AutotaskRetrieveContractsFields200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskRetrieveContractsFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskRetrieveContractsFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutotaskRetrieveContractsFields200Response**](AutotaskRetrieveContractsFields200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskRetrieveMappings

> AutotaskRetrieveMappings200Response AutotaskRetrieveMappings(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve Autotask mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskRetrieveMappings(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskRetrieveMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskRetrieveMappings`: AutotaskRetrieveMappings200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskRetrieveMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskRetrieveMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**AutotaskRetrieveMappings200Response**](AutotaskRetrieveMappings200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskRetrieveServices

> AutotaskRetrieveServices200Response AutotaskRetrieveServices(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve Autotask Contract Services



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskRetrieveServices(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskRetrieveServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskRetrieveServices`: AutotaskRetrieveServices200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskRetrieveServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskRetrieveServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**AutotaskRetrieveServices200Response**](AutotaskRetrieveServices200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskRetrieveSettings

> AutotaskSettings AutotaskRetrieveSettings(ctx, uUID).Execute()

Retrieve Autotask Integration settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskRetrieveSettings(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskRetrieveSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskRetrieveSettings`: AutotaskSettings
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskRetrieveSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskRetrieveSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutotaskSettings**](AutotaskSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskUpdateAlertConfiguration

> AutotaskTicketingAlertConfiguration AutotaskUpdateAlertConfiguration(ctx, providerId, alertUUID).Body(body).Execute()

Update an Autotask ticketing alert's configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    alertUUID := "alertUUID_example" // string | 
    body := *openapiclient.NewAutotaskTicketingAlertConfigurationRequest("Destination_example", int32(123), *openapiclient.NewAutotaskTicketingAlertConfigurationPriority(), false, *openapiclient.NewAutotaskTicketingAlertConfigurationPriority()) // AutotaskTicketingAlertConfigurationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskUpdateAlertConfiguration(context.Background(), providerId, alertUUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskUpdateAlertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskUpdateAlertConfiguration`: AutotaskTicketingAlertConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskUpdateAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**alertUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskUpdateAlertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AutotaskTicketingAlertConfigurationRequest**](AutotaskTicketingAlertConfigurationRequest.md) |  | 

### Return type

[**AutotaskTicketingAlertConfiguration**](AutotaskTicketingAlertConfiguration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutotaskUpdateConfiguration

> AutotaskIntegration AutotaskUpdateConfiguration(ctx, uUID).Body(body).Execute()

Update Autotask Integration configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    body := *openapiclient.NewAutotaskIntegrationPatchReq() // AutotaskIntegrationPatchReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.AutotaskUpdateConfiguration(context.Background(), uUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.AutotaskUpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutotaskUpdateConfiguration`: AutotaskIntegration
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.AutotaskUpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutotaskUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AutotaskIntegrationPatchReq**](AutotaskIntegrationPatchReq.md) |  | 

### Return type

[**AutotaskIntegration**](AutotaskIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseCreateConfiguration

> AutotaskCreateConfiguration201Response ConnectwiseCreateConfiguration(ctx, providerId).Body(body).Execute()

Creates a new ConnectWise integration for the provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    body := *openapiclient.NewConnectwiseIntegrationReq("CompanyId_example", "PrivateKey_example", "PublicKey_example", "Url_example") // ConnectwiseIntegrationReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseCreateConfiguration(context.Background(), providerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseCreateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseCreateConfiguration`: AutotaskCreateConfiguration201Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseCreateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseCreateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConnectwiseIntegrationReq**](ConnectwiseIntegrationReq.md) |  | 

### Return type

[**AutotaskCreateConfiguration201Response**](AutotaskCreateConfiguration201Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseDeleteConfiguration

> ConnectwiseDeleteConfiguration(ctx, uUID).Execute()

Delete ConnectWise Integration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProvidersApi.ConnectwiseDeleteConfiguration(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseDeleteConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseDeleteConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseGetConfiguration

> ConnectwiseIntegration ConnectwiseGetConfiguration(ctx, uUID).Execute()

Retrieve ConnectWise Integration Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseGetConfiguration(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseGetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseGetConfiguration`: ConnectwiseIntegration
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseGetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectwiseIntegration**](ConnectwiseIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwisePatchMappings

> ConnectWiseMappingRequest ConnectwisePatchMappings(ctx, uUID).Body(body).Execute()

Create, edit, and/or delete ConnectWise Mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    body := *openapiclient.NewConnectWiseMappingRequest() // ConnectWiseMappingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwisePatchMappings(context.Background(), uUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwisePatchMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwisePatchMappings`: ConnectWiseMappingRequest
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwisePatchMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwisePatchMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConnectWiseMappingRequest**](ConnectWiseMappingRequest.md) |  | 

### Return type

[**ConnectWiseMappingRequest**](ConnectWiseMappingRequest.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwisePatchSettings

> ConnectWiseSettings ConnectwisePatchSettings(ctx, uUID).Body(body).Execute()

Create, edit, and/or delete ConnectWise Integration settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    body := *openapiclient.NewConnectWiseSettingsPatchReq() // ConnectWiseSettingsPatchReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwisePatchSettings(context.Background(), uUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwisePatchSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwisePatchSettings`: ConnectWiseSettings
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwisePatchSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwisePatchSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConnectWiseSettingsPatchReq**](ConnectWiseSettingsPatchReq.md) |  | 

### Return type

[**ConnectWiseSettings**](ConnectWiseSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseRetrieveAdditions

> ConnectwiseRetrieveAdditions200Response ConnectwiseRetrieveAdditions(ctx, uUID, agreementID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve ConnectWise Additions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    agreementID := "agreementID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseRetrieveAdditions(context.Background(), uUID, agreementID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseRetrieveAdditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseRetrieveAdditions`: ConnectwiseRetrieveAdditions200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseRetrieveAdditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseRetrieveAdditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**ConnectwiseRetrieveAdditions200Response**](ConnectwiseRetrieveAdditions200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseRetrieveAgreements

> ConnectwiseRetrieveAgreements200Response ConnectwiseRetrieveAgreements(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve ConnectWise Agreements



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseRetrieveAgreements(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseRetrieveAgreements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseRetrieveAgreements`: ConnectwiseRetrieveAgreements200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseRetrieveAgreements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseRetrieveAgreementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**ConnectwiseRetrieveAgreements200Response**](ConnectwiseRetrieveAgreements200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseRetrieveAllAlertConfigurationOptions

> ConnectWiseTicketingAlertConfigurationOptions ConnectwiseRetrieveAllAlertConfigurationOptions(ctx, providerId).Execute()

Get all ConnectWise ticketing alert configuration options for a provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseRetrieveAllAlertConfigurationOptions(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseRetrieveAllAlertConfigurationOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseRetrieveAllAlertConfigurationOptions`: ConnectWiseTicketingAlertConfigurationOptions
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseRetrieveAllAlertConfigurationOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseRetrieveAllAlertConfigurationOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectWiseTicketingAlertConfigurationOptions**](ConnectWiseTicketingAlertConfigurationOptions.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseRetrieveAllAlertConfigurations

> ConnectWiseTicketingAlertConfigurationList ConnectwiseRetrieveAllAlertConfigurations(ctx, providerId).Execute()

Get all ConnectWise ticketing alert configurations for a provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseRetrieveAllAlertConfigurations(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseRetrieveAllAlertConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseRetrieveAllAlertConfigurations`: ConnectWiseTicketingAlertConfigurationList
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseRetrieveAllAlertConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseRetrieveAllAlertConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectWiseTicketingAlertConfigurationList**](ConnectWiseTicketingAlertConfigurationList.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseRetrieveCompanies

> ConnectwiseCompanyResp ConnectwiseRetrieveCompanies(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve ConnectWise Companies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseRetrieveCompanies(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseRetrieveCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseRetrieveCompanies`: ConnectwiseCompanyResp
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseRetrieveCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseRetrieveCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**ConnectwiseCompanyResp**](ConnectwiseCompanyResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseRetrieveCompanyTypes

> ConnectwiseCompanyTypeResp ConnectwiseRetrieveCompanyTypes(ctx, uUID).Execute()

Retrieve ConnectWise Company Types



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseRetrieveCompanyTypes(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseRetrieveCompanyTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseRetrieveCompanyTypes`: ConnectwiseCompanyTypeResp
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseRetrieveCompanyTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseRetrieveCompanyTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectwiseCompanyTypeResp**](ConnectwiseCompanyTypeResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseRetrieveMappings

> ConnectwiseRetrieveMappings200Response ConnectwiseRetrieveMappings(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve ConnectWise mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseRetrieveMappings(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseRetrieveMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseRetrieveMappings`: ConnectwiseRetrieveMappings200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseRetrieveMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseRetrieveMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**ConnectwiseRetrieveMappings200Response**](ConnectwiseRetrieveMappings200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseRetrieveSettings

> ConnectWiseSettings ConnectwiseRetrieveSettings(ctx, uUID).Execute()

Retrieve ConnectWise Integration settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseRetrieveSettings(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseRetrieveSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseRetrieveSettings`: ConnectWiseSettings
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseRetrieveSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseRetrieveSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectWiseSettings**](ConnectWiseSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseUpdateAlertConfiguration

> ConnectWiseTicketingAlertConfiguration ConnectwiseUpdateAlertConfiguration(ctx, providerId, alertUUID).Body(body).Execute()

Update a ConnectWise ticketing alert's configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    alertUUID := "alertUUID_example" // string | 
    body := *openapiclient.NewConnectWiseTicketingAlertConfigurationRequest(false) // ConnectWiseTicketingAlertConfigurationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseUpdateAlertConfiguration(context.Background(), providerId, alertUUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseUpdateAlertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseUpdateAlertConfiguration`: ConnectWiseTicketingAlertConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseUpdateAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**alertUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseUpdateAlertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ConnectWiseTicketingAlertConfigurationRequest**](ConnectWiseTicketingAlertConfigurationRequest.md) |  | 

### Return type

[**ConnectWiseTicketingAlertConfiguration**](ConnectWiseTicketingAlertConfiguration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectwiseUpdateConfiguration

> ConnectwiseIntegration ConnectwiseUpdateConfiguration(ctx, uUID).Body(body).Execute()

Update ConnectWise Integration configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    body := *openapiclient.NewConnectwiseIntegrationPatchReq() // ConnectwiseIntegrationPatchReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ConnectwiseUpdateConfiguration(context.Background(), uUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ConnectwiseUpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectwiseUpdateConfiguration`: ConnectwiseIntegration
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ConnectwiseUpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectwiseUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConnectwiseIntegrationPatchReq**](ConnectwiseIntegrationPatchReq.md) |  | 

### Return type

[**ConnectwiseIntegration**](ConnectwiseIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MtpIntegrationRetrieveAlerts

> TicketingIntegrationAlertsResp MtpIntegrationRetrieveAlerts(ctx, providerId).Execute()

Get all ticketing alerts available for a provider's ticketing integration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.MtpIntegrationRetrieveAlerts(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.MtpIntegrationRetrieveAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MtpIntegrationRetrieveAlerts`: TicketingIntegrationAlertsResp
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.MtpIntegrationRetrieveAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMtpIntegrationRetrieveAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TicketingIntegrationAlertsResp**](TicketingIntegrationAlertsResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MtpIntegrationRetrieveSyncErrors

> IntegrationSyncErrorResp MtpIntegrationRetrieveSyncErrors(ctx, uUID, integrationType).Execute()

Retrieve Recent Integration Sync Errors



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    integrationType := "integrationType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.MtpIntegrationRetrieveSyncErrors(context.Background(), uUID, integrationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.MtpIntegrationRetrieveSyncErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MtpIntegrationRetrieveSyncErrors`: IntegrationSyncErrorResp
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.MtpIntegrationRetrieveSyncErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 
**integrationType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMtpIntegrationRetrieveSyncErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IntegrationSyncErrorResp**](IntegrationSyncErrorResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesDelete

> PolicyGroupTemplatesDelete(ctx, providerId, id).Execute()

Deletes policy group template.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProvidersApi.PolicyGroupTemplatesDelete(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.PolicyGroupTemplatesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesGet

> PolicyGroupTemplate PolicyGroupTemplatesGet(ctx, providerId, id).Execute()

Gets a provider's policy group template.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.PolicyGroupTemplatesGet(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.PolicyGroupTemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesGet`: PolicyGroupTemplate
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.PolicyGroupTemplatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PolicyGroupTemplate**](PolicyGroupTemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesGetConfiguredPolicyTemplate

> ConfiguredPolicyTemplate PolicyGroupTemplatesGetConfiguredPolicyTemplate(ctx, providerId, id).Execute()

Retrieve a configured policy template by id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesGetConfiguredPolicyTemplate`: ConfiguredPolicyTemplate
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesGetConfiguredPolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConfiguredPolicyTemplate**](ConfiguredPolicyTemplate.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesList

> PolicyGroupTemplates PolicyGroupTemplatesList(ctx, providerId).Fields(fields).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()

List a provider's policy group templates.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.PolicyGroupTemplatesList(context.Background(), providerId).Fields(fields).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.PolicyGroupTemplatesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesList`: PolicyGroupTemplates
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.PolicyGroupTemplatesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]

### Return type

[**PolicyGroupTemplates**](PolicyGroupTemplates.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesListConfiguredPolicyTemplates

> PolicyGroupTemplatesListConfiguredPolicyTemplates200Response PolicyGroupTemplatesListConfiguredPolicyTemplates(ctx, providerId).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()

List a provider's configured policy templates.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.PolicyGroupTemplatesListConfiguredPolicyTemplates(context.Background(), providerId).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.PolicyGroupTemplatesListConfiguredPolicyTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesListConfiguredPolicyTemplates`: PolicyGroupTemplatesListConfiguredPolicyTemplates200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.PolicyGroupTemplatesListConfiguredPolicyTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesListConfiguredPolicyTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]

### Return type

[**PolicyGroupTemplatesListConfiguredPolicyTemplates200Response**](PolicyGroupTemplatesListConfiguredPolicyTemplates200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PolicyGroupTemplatesListMembers

> PolicyGroupTemplateMembers PolicyGroupTemplatesListMembers(ctx, providerId, id).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()

Gets the list of members from a policy group template.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.PolicyGroupTemplatesListMembers(context.Background(), providerId, id).Skip(skip).Sort(sort).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.PolicyGroupTemplatesListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyGroupTemplatesListMembers`: PolicyGroupTemplateMembers
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.PolicyGroupTemplatesListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyGroupTemplatesListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]

### Return type

[**PolicyGroupTemplateMembers**](PolicyGroupTemplateMembers.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderOrganizationsCreateOrg

> Organization ProviderOrganizationsCreateOrg(ctx, providerId).Body(body).Execute()

Create Provider Organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    body := *openapiclient.NewCreateOrganization() // CreateOrganization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProviderOrganizationsCreateOrg(context.Background(), providerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProviderOrganizationsCreateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderOrganizationsCreateOrg`: Organization
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProviderOrganizationsCreateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderOrganizationsCreateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrganization**](CreateOrganization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderOrganizationsUpdateOrg

> Organization ProviderOrganizationsUpdateOrg(ctx, providerId, id).Body(body).Execute()

Update Provider Organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 
    body := *openapiclient.NewOrganization() // Organization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProviderOrganizationsUpdateOrg(context.Background(), providerId, id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProviderOrganizationsUpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProviderOrganizationsUpdateOrg`: Organization
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProviderOrganizationsUpdateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderOrganizationsUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Organization**](Organization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersGetProvider

> Provider ProvidersGetProvider(ctx, providerId).Fields(fields).Execute()

Retrieve Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersGetProvider(context.Background(), providerId).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersGetProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersGetProvider`: Provider
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersGetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]

### Return type

[**Provider**](Provider.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersListAdministrators

> ProvidersListAdministrators200Response ProvidersListAdministrators(ctx, providerId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).SortIgnoreCase(sortIgnoreCase).Execute()

List Provider Administrators



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    sortIgnoreCase := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersListAdministrators(context.Background(), providerId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).SortIgnoreCase(sortIgnoreCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersListAdministrators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersListAdministrators`: ProvidersListAdministrators200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersListAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersListAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **sortIgnoreCase** | **[]string** | The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**ProvidersListAdministrators200Response**](ProvidersListAdministrators200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersListOrganizations

> ProvidersListOrganizations200Response ProvidersListOrganizations(ctx, providerId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).SortIgnoreCase(sortIgnoreCase).Execute()

List Provider Organizations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    sortIgnoreCase := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersListOrganizations(context.Background(), providerId).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).SortIgnoreCase(sortIgnoreCase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersListOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersListOrganizations`: ProvidersListOrganizations200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersListOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersListOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **sortIgnoreCase** | **[]string** | The comma separated fields used to sort the collection, ignoring case. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**ProvidersListOrganizations200Response**](ProvidersListOrganizations200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersPostAdmins

> Administrator ProvidersPostAdmins(ctx, providerId).Body(body).Execute()

Create a new Provider Administrator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    body := *openapiclient.NewProviderAdminReq("Email_example") // ProviderAdminReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersPostAdmins(context.Background(), providerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersPostAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersPostAdmins`: Administrator
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersPostAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersPostAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProviderAdminReq**](ProviderAdminReq.md) |  | 

### Return type

[**Administrator**](Administrator.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersRemoveAdministrator

> ProvidersRemoveAdministrator(ctx, providerId, id).Execute()

Delete Provider Administrator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProvidersApi.ProvidersRemoveAdministrator(context.Background(), providerId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersRemoveAdministrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersRemoveAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersRetrieveIntegrations

> IntegrationsResponse ProvidersRetrieveIntegrations(ctx, providerId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve Integrations for Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersRetrieveIntegrations(context.Background(), providerId).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersRetrieveIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersRetrieveIntegrations`: IntegrationsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersRetrieveIntegrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersRetrieveIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**IntegrationsResponse**](IntegrationsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersRetrieveInvoice

> *os.File ProvidersRetrieveInvoice(ctx, providerId, iD).Execute()

Download a provider's invoice.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    iD := "iD_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersRetrieveInvoice(context.Background(), providerId, iD).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersRetrieveInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersRetrieveInvoice`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersRetrieveInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**iD** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersRetrieveInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersRetrieveInvoices

> ProviderInvoiceResponse ProvidersRetrieveInvoices(ctx, providerId).Skip(skip).Sort(sort).Limit(limit).Execute()

List a provider's invoices.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.ProvidersRetrieveInvoices(context.Background(), providerId).Skip(skip).Sort(sort).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.ProvidersRetrieveInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvidersRetrieveInvoices`: ProviderInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.ProvidersRetrieveInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersRetrieveInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]

### Return type

[**ProviderInvoiceResponse**](ProviderInvoiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroCreateConfiguration

> AutotaskCreateConfiguration201Response SyncroCreateConfiguration(ctx, providerId).Body(body).Execute()

Creates a new Syncro integration for the provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    body := *openapiclient.NewSyncroIntegrationReq("ApiToken_example", "Subdomain_example") // SyncroIntegrationReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroCreateConfiguration(context.Background(), providerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroCreateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroCreateConfiguration`: AutotaskCreateConfiguration201Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroCreateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroCreateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyncroIntegrationReq**](SyncroIntegrationReq.md) |  | 

### Return type

[**AutotaskCreateConfiguration201Response**](AutotaskCreateConfiguration201Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroDeleteConfiguration

> SyncroDeleteConfiguration(ctx, uUID).Execute()

Delete Syncro Integration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProvidersApi.SyncroDeleteConfiguration(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroDeleteConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroDeleteConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroGetConfiguration

> SyncroIntegration SyncroGetConfiguration(ctx, uUID).Execute()

Retrieve Syncro Integration Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroGetConfiguration(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroGetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroGetConfiguration`: SyncroIntegration
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroGetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncroIntegration**](SyncroIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroPatchMappings

> SyncroMappingRequest SyncroPatchMappings(ctx, uUID).Body(body).Execute()

Create, edit, and/or delete Syncro Mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    body := *openapiclient.NewSyncroMappingRequest() // SyncroMappingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroPatchMappings(context.Background(), uUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroPatchMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroPatchMappings`: SyncroMappingRequest
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroPatchMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroPatchMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyncroMappingRequest**](SyncroMappingRequest.md) |  | 

### Return type

[**SyncroMappingRequest**](SyncroMappingRequest.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroPatchSettings

> SyncroSettings SyncroPatchSettings(ctx, uUID).Body(body).Execute()

Create, edit, and/or delete Syncro Integration settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    body := *openapiclient.NewSyncroSettingsPatchReq() // SyncroSettingsPatchReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroPatchSettings(context.Background(), uUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroPatchSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroPatchSettings`: SyncroSettings
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroPatchSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroPatchSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyncroSettingsPatchReq**](SyncroSettingsPatchReq.md) |  | 

### Return type

[**SyncroSettings**](SyncroSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroRetrieveAllAlertConfigurationOptions

> SyncroTicketingAlertConfigurationOptions SyncroRetrieveAllAlertConfigurationOptions(ctx, providerId).Execute()

Get all Syncro ticketing alert configuration options for a provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroRetrieveAllAlertConfigurationOptions(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroRetrieveAllAlertConfigurationOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroRetrieveAllAlertConfigurationOptions`: SyncroTicketingAlertConfigurationOptions
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroRetrieveAllAlertConfigurationOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroRetrieveAllAlertConfigurationOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncroTicketingAlertConfigurationOptions**](SyncroTicketingAlertConfigurationOptions.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroRetrieveAllAlertConfigurations

> SyncroTicketingAlertConfigurationList SyncroRetrieveAllAlertConfigurations(ctx, providerId).Execute()

Get all Syncro ticketing alert configurations for a provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroRetrieveAllAlertConfigurations(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroRetrieveAllAlertConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroRetrieveAllAlertConfigurations`: SyncroTicketingAlertConfigurationList
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroRetrieveAllAlertConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroRetrieveAllAlertConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncroTicketingAlertConfigurationList**](SyncroTicketingAlertConfigurationList.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroRetrieveBillingMappingConfigurationOptions

> SyncroBillingMappingConfigurationOptionsResp SyncroRetrieveBillingMappingConfigurationOptions(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve Syncro billing mappings dependencies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroRetrieveBillingMappingConfigurationOptions(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroRetrieveBillingMappingConfigurationOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroRetrieveBillingMappingConfigurationOptions`: SyncroBillingMappingConfigurationOptionsResp
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroRetrieveBillingMappingConfigurationOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroRetrieveBillingMappingConfigurationOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**SyncroBillingMappingConfigurationOptionsResp**](SyncroBillingMappingConfigurationOptionsResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroRetrieveCompanies

> SyncroCompanyResp SyncroRetrieveCompanies(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve Syncro Companies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroRetrieveCompanies(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroRetrieveCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroRetrieveCompanies`: SyncroCompanyResp
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroRetrieveCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroRetrieveCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**SyncroCompanyResp**](SyncroCompanyResp.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroRetrieveMappings

> SyncroRetrieveMappings200Response SyncroRetrieveMappings(ctx, uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()

Retrieve Syncro mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    fields := []string{"Inner_example"} // []string | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  (optional) (default to [])
    filter := []string{"Inner_example"} // []string | A filter to apply to the query.  **Filter structure**: `<field>:<operator>:<value>`.  **field** = Populate with a valid field from an endpoint response.  **operator** =  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** = Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** `GET /api/v2/groups?filter=name:eq:Test+Group` (optional) (default to [])
    limit := int32(56) // int32 | The number of records to return at once. Limited to 100. (optional) (default to 10)
    skip := int32(56) // int32 | The offset into the records to return. (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | The comma separated fields used to sort the collection. Default sort is ascending, prefix with `-` to sort descending.  (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroRetrieveMappings(context.Background(), uUID).Fields(fields).Filter(filter).Limit(limit).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroRetrieveMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroRetrieveMappings`: SyncroRetrieveMappings200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroRetrieveMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroRetrieveMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | The comma separated fields included in the returned records. If omitted, the default list of fields will be returned.  | [default to []]
 **filter** | **[]string** | A filter to apply to the query.  **Filter structure**: &#x60;&lt;field&gt;:&lt;operator&gt;:&lt;value&gt;&#x60;.  **field** &#x3D; Populate with a valid field from an endpoint response.  **operator** &#x3D;  Supported operators are: eq, ne, gt, ge, lt, le, between, search, in. _Note: v1 operators differ from v2 operators._  **value** &#x3D; Populate with the value you want to search for. Is case sensitive. Supports wild cards.  **EX:** &#x60;GET /api/v2/groups?filter&#x3D;name:eq:Test+Group&#x60; | [default to []]
 **limit** | **int32** | The number of records to return at once. Limited to 100. | [default to 10]
 **skip** | **int32** | The offset into the records to return. | [default to 0]
 **sort** | **[]string** | The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to []]

### Return type

[**SyncroRetrieveMappings200Response**](SyncroRetrieveMappings200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroRetrieveSettings

> SyncroSettings SyncroRetrieveSettings(ctx, uUID).Execute()

Retrieve Syncro Integration settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroRetrieveSettings(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroRetrieveSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroRetrieveSettings`: SyncroSettings
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroRetrieveSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroRetrieveSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncroSettings**](SyncroSettings.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroUpdateAlertConfiguration

> SyncroTicketingAlertConfiguration SyncroUpdateAlertConfiguration(ctx, providerId, alertUUID).Body(body).Execute()

Update a Syncro ticketing alert's configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    providerId := "providerId_example" // string | 
    alertUUID := "alertUUID_example" // string | 
    body := *openapiclient.NewSyncroTicketingAlertConfigurationRequest("ProblemType_example", false) // SyncroTicketingAlertConfigurationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroUpdateAlertConfiguration(context.Background(), providerId, alertUUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroUpdateAlertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroUpdateAlertConfiguration`: SyncroTicketingAlertConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroUpdateAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** |  | 
**alertUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroUpdateAlertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SyncroTicketingAlertConfigurationRequest**](SyncroTicketingAlertConfigurationRequest.md) |  | 

### Return type

[**SyncroTicketingAlertConfiguration**](SyncroTicketingAlertConfiguration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncroUpdateConfiguration

> SyncroIntegration SyncroUpdateConfiguration(ctx, uUID).Body(body).Execute()

Update Syncro Integration configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func main() {
    uUID := "uUID_example" // string | 
    body := *openapiclient.NewSyncroIntegrationPatchReq() // SyncroIntegrationPatchReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidersApi.SyncroUpdateConfiguration(context.Background(), uUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidersApi.SyncroUpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncroUpdateConfiguration`: SyncroIntegration
    fmt.Fprintf(os.Stdout, "Response from `ProvidersApi.SyncroUpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncroUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyncroIntegrationPatchReq**](SyncroIntegrationPatchReq.md) |  | 

### Return type

[**SyncroIntegration**](SyncroIntegration.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


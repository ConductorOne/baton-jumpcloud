# \CommandTriggersApi

All URIs are relative to *https://console.jumpcloud.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandTriggerWebhookPost**](CommandTriggersApi.md#CommandTriggerWebhookPost) | **Post** /command/trigger/{triggername} | Launch a command via a Trigger



## CommandTriggerWebhookPost

> Triggerreturn CommandTriggerWebhookPost(ctx, triggername).XOrgId(xOrgId).Body(body).Execute()

Launch a command via a Trigger



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
)

func main() {
    triggername := "triggername_example" // string | 
    xOrgId := "xOrgId_example" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommandTriggersApi.CommandTriggerWebhookPost(context.Background(), triggername).XOrgId(xOrgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommandTriggersApi.CommandTriggerWebhookPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommandTriggerWebhookPost`: Triggerreturn
    fmt.Fprintf(os.Stdout, "Response from `CommandTriggersApi.CommandTriggerWebhookPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggername** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommandTriggerWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xOrgId** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**Triggerreturn**](Triggerreturn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


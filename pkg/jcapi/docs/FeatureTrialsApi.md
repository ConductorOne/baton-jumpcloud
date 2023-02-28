# \FeatureTrialsApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FeatureTrialsGetFeatureTrials**](FeatureTrialsApi.md#FeatureTrialsGetFeatureTrials) | **Get** /featureTrials/{feature_code} | Check current feature trial usage for a specific feature



## FeatureTrialsGetFeatureTrials

> FeatureTrialData FeatureTrialsGetFeatureTrials(ctx, featureCode).Execute()

Check current feature trial usage for a specific feature



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func main() {
    featureCode := "featureCode_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureTrialsApi.FeatureTrialsGetFeatureTrials(context.Background(), featureCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureTrialsApi.FeatureTrialsGetFeatureTrials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FeatureTrialsGetFeatureTrials`: FeatureTrialData
    fmt.Fprintf(os.Stdout, "Response from `FeatureTrialsApi.FeatureTrialsGetFeatureTrials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFeatureTrialsGetFeatureTrialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeatureTrialData**](FeatureTrialData.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


/*
JumpCloud API

Testing PolicyGroupTemplatesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package jcapi2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
)

func Test_jcapi2_PolicyGroupTemplatesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PolicyGroupTemplatesApiService PolicyGroupTemplatesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string
		var id string

		resp, httpRes, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesGet(context.Background(), providerId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyGroupTemplatesApiService PolicyGroupTemplatesGetConfiguredPolicyTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string
		var id string

		resp, httpRes, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesGetConfiguredPolicyTemplate(context.Background(), providerId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyGroupTemplatesApiService PolicyGroupTemplatesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string

		resp, httpRes, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesList(context.Background(), providerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyGroupTemplatesApiService PolicyGroupTemplatesListConfiguredPolicyTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string

		resp, httpRes, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesListConfiguredPolicyTemplates(context.Background(), providerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyGroupTemplatesApiService PolicyGroupTemplatesListMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string
		var id string

		resp, httpRes, err := apiClient.PolicyGroupTemplatesApi.PolicyGroupTemplatesListMembers(context.Background(), providerId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
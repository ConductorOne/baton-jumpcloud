/*
JumpCloud API

Testing SystemsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package jcapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/ConductorOne/baton-jumpcloud/pkg/jcapi"
)

func Test_jcapi_SystemsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SystemsApiService GraphSystemAssociationsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemsApi.GraphSystemAssociationsList(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GraphSystemAssociationsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		httpRes, err := apiClient.SystemsApi.GraphSystemAssociationsPost(context.Background(), systemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GraphSystemMemberOf", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemsApi.GraphSystemMemberOf(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GraphSystemTraverseCommand", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemsApi.GraphSystemTraverseCommand(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GraphSystemTraversePolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemsApi.GraphSystemTraversePolicy(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GraphSystemTraversePolicyGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemsApi.GraphSystemTraversePolicyGroup(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GraphSystemTraverseUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemsApi.GraphSystemTraverseUser(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService GraphSystemTraverseUserGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemsApi.GraphSystemTraverseUserGroup(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService SystemsGetFDEKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemsApi.SystemsGetFDEKey(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemsApiService SystemsListSoftwareAppsWithStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemsApi.SystemsListSoftwareAppsWithStatuses(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
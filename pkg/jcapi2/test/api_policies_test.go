/*
JumpCloud API

Testing PoliciesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package jcapi2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
)

func Test_jcapi2_PoliciesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PoliciesApiService GraphPolicyAssociationsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PoliciesApi.GraphPolicyAssociationsList(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService GraphPolicyAssociationsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId string

		httpRes, err := apiClient.PoliciesApi.GraphPolicyAssociationsPost(context.Background(), policyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService GraphPolicyMemberOf", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PoliciesApi.GraphPolicyMemberOf(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService GraphPolicyTraverseSystem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PoliciesApi.GraphPolicyTraverseSystem(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService GraphPolicyTraverseSystemGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PoliciesApi.GraphPolicyTraverseSystemGroup(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PoliciesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.PoliciesApi.PoliciesDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PoliciesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.PoliciesApi.PoliciesGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PoliciesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PoliciesApi.PoliciesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PoliciesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PoliciesApi.PoliciesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PoliciesPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.PoliciesApi.PoliciesPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PolicyresultsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.PoliciesApi.PolicyresultsGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PolicyresultsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PoliciesApi.PolicyresultsList(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PolicyresultsOrgList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PoliciesApi.PolicyresultsOrgList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PolicystatusesPoliciesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.PoliciesApi.PolicystatusesPoliciesList(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PolicystatusesSystemsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.PoliciesApi.PolicystatusesSystemsList(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PolicytemplatesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.PoliciesApi.PolicytemplatesGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesApiService PolicytemplatesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PoliciesApi.PolicytemplatesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

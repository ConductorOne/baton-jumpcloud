/*
JumpCloud API

Testing SystemGroupsApiService

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

func Test_jcapi2_SystemGroupsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SystemGroupsApiService GraphSystemGroupAssociationsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SystemGroupsApi.GraphSystemGroupAssociationsList(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GraphSystemGroupAssociationsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		httpRes, err := apiClient.SystemGroupsApi.GraphSystemGroupAssociationsPost(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GraphSystemGroupMembersList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SystemGroupsApi.GraphSystemGroupMembersList(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GraphSystemGroupMembersPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		httpRes, err := apiClient.SystemGroupsApi.GraphSystemGroupMembersPost(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GraphSystemGroupMembership", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SystemGroupsApi.GraphSystemGroupMembership(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GraphSystemGroupTraversePolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SystemGroupsApi.GraphSystemGroupTraversePolicy(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GraphSystemGroupTraversePolicyGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SystemGroupsApi.GraphSystemGroupTraversePolicyGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GraphSystemGroupTraverseUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SystemGroupsApi.GraphSystemGroupTraverseUser(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GraphSystemGroupTraverseUserGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SystemGroupsApi.GraphSystemGroupTraverseUserGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GroupsSystemDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SystemGroupsApi.GroupsSystemDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GroupsSystemGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SystemGroupsApi.GroupsSystemGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GroupsSystemList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemGroupsApi.GroupsSystemList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GroupsSystemPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SystemGroupsApi.GroupsSystemPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GroupsSystemPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SystemGroupsApi.GroupsSystemPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GroupsSystemSuggestionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SystemGroupsApi.GroupsSystemSuggestionsGet(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemGroupsApiService GroupsSystemSuggestionsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SystemGroupsApi.GroupsSystemSuggestionsPost(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

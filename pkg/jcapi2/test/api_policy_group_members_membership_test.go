/*
JumpCloud API

Testing PolicyGroupMembersMembershipApiService

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

func Test_jcapi2_PolicyGroupMembersMembershipApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PolicyGroupMembersMembershipApiService GraphPolicyGroupMembersList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.PolicyGroupMembersMembershipApi.GraphPolicyGroupMembersList(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyGroupMembersMembershipApiService GraphPolicyGroupMembersPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		httpRes, err := apiClient.PolicyGroupMembersMembershipApi.GraphPolicyGroupMembersPost(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyGroupMembersMembershipApiService GraphPolicyGroupMembership", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.PolicyGroupMembersMembershipApi.GraphPolicyGroupMembership(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
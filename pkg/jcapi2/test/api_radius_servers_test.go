/*
JumpCloud API

Testing RADIUSServersApiService

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

func Test_jcapi2_RADIUSServersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RADIUSServersApiService GraphRadiusServerAssociationsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var radiusserverId string

		resp, httpRes, err := apiClient.RADIUSServersApi.GraphRadiusServerAssociationsList(context.Background(), radiusserverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RADIUSServersApiService GraphRadiusServerAssociationsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var radiusserverId string

		httpRes, err := apiClient.RADIUSServersApi.GraphRadiusServerAssociationsPost(context.Background(), radiusserverId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RADIUSServersApiService GraphRadiusServerTraverseUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var radiusserverId string

		resp, httpRes, err := apiClient.RADIUSServersApi.GraphRadiusServerTraverseUser(context.Background(), radiusserverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RADIUSServersApiService GraphRadiusServerTraverseUserGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var radiusserverId string

		resp, httpRes, err := apiClient.RADIUSServersApi.GraphRadiusServerTraverseUserGroup(context.Background(), radiusserverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
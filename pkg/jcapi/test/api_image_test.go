/*
JumpCloud API

Testing ImageApiService

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

func Test_jcapi_ImageApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImageApiService ApplicationsDeleteLogo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		httpRes, err := apiClient.ImageApi.ApplicationsDeleteLogo(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
/*
JumpCloud API

Testing ImageApiService

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

func Test_jcapi2_ImageApiService(t *testing.T) {

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

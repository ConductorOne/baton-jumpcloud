/*
JumpCloud API

Testing FeatureTrialsApiService

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

func Test_jcapi2_FeatureTrialsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FeatureTrialsApiService FeatureTrialsGetFeatureTrials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var featureCode string

		resp, httpRes, err := apiClient.FeatureTrialsApi.FeatureTrialsGetFeatureTrials(context.Background(), featureCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

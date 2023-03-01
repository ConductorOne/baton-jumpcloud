package connector

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
	"github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

func New(ctx context.Context, apiKey string, orgId string) (*Jumpcloud, error) {
	httpClient, err := uhttp.NewClient(ctx, uhttp.WithLogger(true, nil))
	if err != nil {
		return nil, err
	}

	cc1 := jcapi1.NewConfiguration()
	cc1.HTTPClient = httpClient
	cc1.UserAgent = "baton-jumpcloud/0.1.0"

	cc2 := jcapi2.NewConfiguration()
	cc2.HTTPClient = httpClient
	cc2.UserAgent = "baton-jumpcloud/0.1.0"

	if orgId != "" {
		// optional, only needed by API keys linked to mutli-tenant admins
		cc1.AddDefaultHeader("x-org-id", orgId)
		cc2.AddDefaultHeader("x-org-id", orgId)
	}

	return &Jumpcloud{
		_client1: jcapi1.NewAPIClient(cc1),
		_client2: jcapi2.NewAPIClient(cc2),
		apiKey:   apiKey,
	}, nil
}

func (jc *Jumpcloud) client1(ctx context.Context) (context.Context, *jcapi1.APIClient) {
	return context.WithValue(ctx, jcapi1.ContextAPIKeys, map[string]jcapi1.APIKey{
		"x-api-key": {
			Key: jc.apiKey,
		},
	}), jc._client1
}

func (jc *Jumpcloud) client2(ctx context.Context) (context.Context, *jcapi2.APIClient) {
	return context.WithValue(ctx, jcapi2.ContextAPIKeys, map[string]jcapi2.APIKey{
		"x-api-key": {
			Key: jc.apiKey,
		},
	}), jc._client2
}

type jc1Func func(ctx context.Context) (context.Context, *jcapi1.APIClient)
type jc2Func func(ctx context.Context) (context.Context, *jcapi2.APIClient)

type Jumpcloud struct {
	_client1 *jcapi1.APIClient
	_client2 *jcapi2.APIClient
	apiKey   string
}

var (
	// resourceTypeRole = &v2.ResourceType{
	// 	Id:          "role",
	// 	DisplayName: "Role",
	// 	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_ROLE},
	// }
	resourceTypeUser = &v2.ResourceType{
		Id:          "user",
		DisplayName: "User",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_USER},
	}
	resourceTypeGroup = &v2.ResourceType{
		Id:          "group",
		DisplayName: "Group",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_GROUP},
	}
	resourceTypeApp = &v2.ResourceType{
		Id:          "app",
		DisplayName: "App",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_APP},
	}
)

func (c *Jumpcloud) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "JumpCloud",
	}, nil
}

func (c *Jumpcloud) Validate(ctx context.Context) (annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	ctx, client := c.client2(ctx)

	_, resp, err := client.DirectoriesApi.DirectoriesList(ctx).Limit(1).Execute()
	if err != nil {
		l.Error("DirectoriesList for Validate Failed", zap.Error(err))
		return nil, fmt.Errorf("jumpcloud-connector: failed to verify api key: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("jumpcloud-connector verify returned non-200: '%d'", resp.StatusCode)
		l.Error("Invalid Status Code from Verity", zap.Error(err))
		return nil, err
	}

	return nil, nil
}

func (s *Jumpcloud) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{
		newUserBuilder(s.client1, s.client2),
		newGroupBuilder(s.client1, s.client2),
	}
}

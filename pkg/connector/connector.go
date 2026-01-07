package connector

import (
	"context"
	"fmt"
	"net/http"

	cfg "github.com/conductorone/baton-jumpcloud/pkg/config"
	"github.com/conductorone/baton-jumpcloud/pkg/jcapi1"
	"github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/field"
	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type Connector struct {
	_client1 *jcapi1.APIClient
	_client2 *jcapi2.APIClient
	ext      *ExtensionClient
	apiKey   string
}

// Option is a function that configures a Connector.
type Option func(*Connector) error

// WithAPIKey configures the connector with API key authentication.
func WithAPIKey(ctx context.Context, apiKey string, orgId string) Option {
	return func(c *Connector) error {
		httpClient, err := uhttp.NewClient(ctx, uhttp.WithLogger(true, nil), uhttp.WithUserAgent("baton-jumpcloud/0.1.0"))
		if err != nil {
			return err
		}

		cc1 := jcapi1.NewConfiguration()
		cc1.HTTPClient = httpClient
		cc1.UserAgent = "baton-jumpcloud/0.1.0"

		cc2 := jcapi2.NewConfiguration()
		cc2.HTTPClient = httpClient
		cc2.UserAgent = "baton-jumpcloud/0.1.0"

		if orgId != "" {
			// optional, only needed by API keys linked to multi-tenant admins
			cc1.AddDefaultHeader("x-org-id", orgId)
			cc2.AddDefaultHeader("x-org-id", orgId)
		}

		c._client1 = jcapi1.NewAPIClient(cc1)
		c._client2 = jcapi2.NewAPIClient(cc2)

		baseHttpClient, err := uhttp.NewBaseHttpClientWithContext(ctx, httpClient)
		if err != nil {
			return err
		}

		c.ext = &ExtensionClient{
			client: baseHttpClient,
			apiKey: apiKey,
			orgId:  orgId,
		}
		c.apiKey = apiKey

		return nil
	}
}

func NewLambdaConnector(ctx context.Context, jumpcloudCfg *cfg.Jumpcloud, cliOpts *cli.ConnectorOpts) (connectorbuilder.ConnectorBuilderV2, []connectorbuilder.Opt, error) {
	if err := field.Validate(cfg.ConfigurationSchema, jumpcloudCfg); err != nil {
		return nil, nil, err
	}

	l := ctxzap.Extract(ctx)

	opts := WithAPIKey(
		ctx,
		jumpcloudCfg.ApiKey,
		jumpcloudCfg.OrgId,
	)

	cb, err := New(ctx, opts)
	if err != nil {
		l.Error("error creating connector", zap.Error(err))
		return nil, nil, err
	}

	return cb, nil, nil
}

// New returns a new instance of the connector.
func New(ctx context.Context, opts ...Option) (*Connector, error) {
	c := &Connector{}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, fmt.Errorf("failed to apply connector option during initialization: %w", err)
		}
	}

	if c._client1 == nil || c._client2 == nil {
		return nil, uhttp.WrapErrors(
			codes.InvalidArgument,
			"connector initialization failed: API client not configured",
			fmt.Errorf("API clients not properly initialized during connector setup"),
		)
	}

	return c, nil
}

func (jc *Connector) client1(ctx context.Context) (context.Context, *jcapi1.APIClient) {
	return context.WithValue(ctx, jcapi1.ContextAPIKeys, map[string]jcapi1.APIKey{
		"x-api-key": {
			Key: jc.apiKey,
		},
	}), jc._client1
}

func (jc *Connector) client2(ctx context.Context) (context.Context, *jcapi2.APIClient) {
	return context.WithValue(ctx, jcapi2.ContextAPIKeys, map[string]jcapi2.APIKey{
		"x-api-key": {
			Key: jc.apiKey,
		},
	}), jc._client2
}

type jc1Func func(ctx context.Context) (context.Context, *jcapi1.APIClient)
type jc2Func func(ctx context.Context) (context.Context, *jcapi2.APIClient)

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced from the upstream service.
func (c *Connector) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncerV2 {
	return []connectorbuilder.ResourceSyncerV2{
		newUserBuilder(c.client1, c.client2, c.ext),
		newGroupBuilder(c.client1, c.client2),
		newRoleBuilder(c.client1, c.ext),
		newAppBuilder(c.client1, c.client2, c.ext),
	}
}

// Metadata returns metadata about the connector.
func (c *Connector) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "JumpCloud",
		Description: "The JumpCloud connector syncs users, groups, roles, and apps.",
	}, nil
}

// Validate is called to ensure that the connector is properly configured. It should exercise any API credentials
// to be sure that they are valid.
func (c *Connector) Validate(ctx context.Context) (annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	ctx, client := c.client2(ctx)

	_, resp, err := client.DirectoriesApi.DirectoriesList(ctx).Limit(1).Execute()
	if err != nil {
		l.Error("DirectoriesList for Validate Failed", zap.Error(err))
		return nil, wrapSDKError(err, resp, "failed to verify api key")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		code := httpStatusToGRPCCode(resp.StatusCode)
		err := uhttp.WrapErrors(
			code,
			fmt.Sprintf("validate: unexpected status code %d", resp.StatusCode),
			nil,
		)
		l.Error("jumpcloud-connector: Invalid Status Code from Validate", zap.Error(err))
		return nil, err
	}

	return nil, nil
}

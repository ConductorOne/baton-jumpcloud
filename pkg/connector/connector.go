package connector

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-jumpcloud/pkg/client"
	cfg "github.com/conductorone/baton-jumpcloud/pkg/config"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type Connector struct {
	client *client.Client
	// skipRoleResourceType reports whether role is excluded from the sync filter.
	// Named for the skip condition so the zero value is safe: a Connector built
	// without the option syncs everything.
	skipRoleResourceType bool
}

// Option is a function that configures a Connector.
type Option func(*Connector) error

// WithAPIKey configures the connector with API key authentication.
func WithAPIKey(ctx context.Context, apiKey string, orgId string, baseURL string) Option {
	return func(c *Connector) error {
		client, err := client.NewClient(ctx, apiKey, orgId, baseURL)
		if err != nil {
			return err
		}

		c.client = client

		return nil
	}
}

// WithSkipRoleResourceType configures whether the role resource type is excluded from
// the sync filter. The user builder emits cross-type role grants, so when role is
// excluded those grants must be suppressed rather than left pointing at an unsynced type.
func WithSkipRoleResourceType(skip bool) Option {
	return func(c *Connector) error {
		c.skipRoleResourceType = skip
		return nil
	}
}

func NewLambdaConnector(ctx context.Context, jumpcloudCfg *cfg.Jumpcloud, cliOpts *cli.ConnectorOpts) (connectorbuilder.ConnectorBuilderV2, []connectorbuilder.Opt, error) {
	l := ctxzap.Extract(ctx)

	opts := WithAPIKey(
		ctx,
		jumpcloudCfg.ApiKey,
		jumpcloudCfg.OrgId,
		jumpcloudCfg.BaseUrl,
	)

	// nil opts means no filter, so nothing is skipped.
	skipRoleResourceType := cliOpts != nil && !cliOpts.WillSyncResourceType(RoleResourceTypeID)

	cb, err := New(ctx, opts, WithSkipRoleResourceType(skipRoleResourceType))
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

	if c.client == nil {
		return nil, uhttp.WrapErrors(
			codes.InvalidArgument,
			"connector initialization failed: API client not configured",
			nil,
		)
	}

	return c, nil
}

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced from the upstream service.
func (c *Connector) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncerV2 {
	return []connectorbuilder.ResourceSyncerV2{
		newUserBuilder(c.client, c.skipRoleResourceType),
		newGroupBuilder(c.client),
		newRoleBuilder(),
		newAppBuilder(c.client),
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

	options := &client.Options{}
	_, err := c.client.ListDirectories(ctx, options.WithLimit(1))
	if err != nil {
		l.Error("DirectoriesList for Validate Failed", zap.Error(err))
		return nil, err
	}

	return nil, nil
}

package main

import (
	"context"

	cfg "github.com/conductorone/baton-jumpcloud/pkg/config"
	"github.com/conductorone/baton-jumpcloud/pkg/connector"
	"github.com/conductorone/baton-sdk/pkg/config"
)

var version = "dev"

func main() {
	ctx := context.Background()
	config.RunConnector(ctx, "baton-jumpcloud", version, cfg.ConfigurationSchema, connector.NewLambdaConnector)
}

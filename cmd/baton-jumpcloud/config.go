package main

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/spf13/cobra"
)

// config defines the external configuration required for the connector to run.
type config struct {
	cli.BaseConfig `mapstructure:",squash"` // Puts the base config options in the same place as the connector options
	ApiKey         string                   `mapstructure:"api-key"`
	OrgId          string
}

// validateConfig is run after the configuration is loaded, and should return an error if it isn't valid.
func validateConfig(ctx context.Context, cfg *config) error {
	if cfg.ApiKey == "" {
		return fmt.Errorf("api key is missing")
	}
	return nil
}

func cmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("api-key", "", "The API key for the JumpCloud account.  ($BATON_API_KEY)")
	cmd.PersistentFlags().String("org-id", "", "The Org ID for the JumpCloud account (optional). ($BATON_ORG_ID)")
}

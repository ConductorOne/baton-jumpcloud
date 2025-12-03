package config

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	ApiKeyField = field.StringField(
		"api-key",
		field.WithDisplayName("API Key"),
		field.WithDescription("The API key for the JumpCloud account."),
		field.WithIsSecret(true),
		field.WithRequired(true),
	)

	OrgIdField = field.StringField(
		"org-id",
		field.WithDisplayName("Organization ID"),
		field.WithDescription("The Org ID for the JumpCloud account (optional, only needed by API keys linked to multi-tenant admins)."),
	)

	// ConfigurationFields defines the external configuration required for the
	// connector to run.
	ConfigurationFields = []field.SchemaField{
		ApiKeyField,
		OrgIdField,
	}
)

//go:generate go run ./gen

var ConfigurationSchema = field.NewConfiguration(
	ConfigurationFields,
	field.WithConnectorDisplayName("JumpCloud"),
	field.WithHelpUrl("/docs/baton/jumpcloud"),
	field.WithIconUrl("/static/app-icons/jumpcloud.svg"),
)

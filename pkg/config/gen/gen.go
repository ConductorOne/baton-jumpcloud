package main

import (
	cfg "github.com/conductorone/baton-jumpcloud/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("jumpcloud", cfg.ConfigurationSchema)
}

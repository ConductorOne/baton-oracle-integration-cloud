package main

import (
	cfg "github.com/conductorone/baton-oracle-integration-cloud/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("oracle-integration-cloud", cfg.Config)
}

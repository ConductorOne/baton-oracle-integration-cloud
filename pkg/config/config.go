package config

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	ConfigurationFields = []field.SchemaField{}

	ConfigurationSchema = field.Configuration{
		Fields: ConfigurationFields,
	}
)

var Config = field.NewConfiguration(ConfigurationFields)

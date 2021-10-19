package datadog

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type datadogConfig struct {
	APIKey *string `cty:"api_key"`
	AppKey *string `cty:"app_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"app_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &datadogConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) datadogConfig {
	if connection == nil || connection.Config == nil {
		return datadogConfig{}
	}
	config, _ := connection.Config.(datadogConfig)
	return config
}

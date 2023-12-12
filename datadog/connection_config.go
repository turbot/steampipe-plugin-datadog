package datadog

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type datadogConfig struct {
	APIKey *string `hcl:"api_key"`
	AppKey *string `hcl:"app_key"`
	// By default it is https://api.datadoghq.com/
	// If working with "EU" version of Datadog, use https://api.datadoghq.eu/
	ApiURL *string `hcl:"api_url"`
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

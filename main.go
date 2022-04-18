package main

import (
	"github.com/turbot/steampipe-plugin-datadog/datadog"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: datadog.Plugin})
}

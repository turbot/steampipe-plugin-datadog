package main

import (
	"github.com/turbot/steampipe-plugin-datadog/datadog"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: datadog.Plugin})
}

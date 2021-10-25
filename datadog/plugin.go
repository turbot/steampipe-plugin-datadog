package datadog

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-datadog",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromJSONTag(),
		TableMap: map[string]*plugin.Table{
			"datadog_integration_aws": tableDatadogIntegrationAws(ctx),
			"datadog_dashboard":       tableDatadogDashboard(ctx),
			"datadog_monitor":         tableDatadogMonitor(ctx),
			"datadog_permission":      tableDatadogPermission(ctx),
			"datadog_role":            tableDatadogRole(ctx),
			"datadog_user":            tableDatadogUser(ctx),
		},
	}
	return p
}

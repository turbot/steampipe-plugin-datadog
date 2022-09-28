package datadog

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-datadog",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"datadog_dashboard":                  tableDatadogDashboard(ctx),
			"datadog_integration_aws":            tableDatadogIntegrationAws(ctx),
			"datadog_log_event":                  tableDatadogLogEvent(ctx),
			"datadog_logs_metric":                tableDatadogLogsMetric(ctx),
			"datadog_monitor":                    tableDatadogMonitor(ctx),
			"datadog_permission":                 tableDatadogPermission(ctx),
			"datadog_role":                       tableDatadogRole(ctx),
			"datadog_security_monitoring_rule":   tableDatadogSecurityMonitoringRule(ctx),
			"datadog_security_monitoring_signal": tableDatadogSecurityMonitoringSignal(ctx),
			"datadog_user":                       tableDatadogUser(ctx),
		},
	}
	return p
}

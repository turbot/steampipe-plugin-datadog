package datadog

import (
	"context"
	"strings"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableDatadogSecurityMonitoringRule(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_security_monitoring_rule",
		Description: "Security monitoring rules define conditional logic that is applied to all ingested logs and cloud configurations.",
		Get: &plugin.GetConfig{
			Hydrate:    getSecurityMonitoringRule,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listSecurityMonitoringRules,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the rule."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the rule."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Description: "When the rule was created, timestamp in milliseconds."},
			{Name: "creation_author_id", Type: proto.ColumnType_STRING, Description: "User ID of the user who created the rule."},
			{Name: "has_extended_title", Type: proto.ColumnType_BOOL, Description: "Whether the notifications include the triggering group-by values in their title."},
			{Name: "is_default", Type: proto.ColumnType_BOOL, Description: "Whether the rule is included by default."},
			{Name: "is_deleted", Type: proto.ColumnType_BOOL, Description: "Whether the rule has been deleted."},
			{Name: "is_enabled", Type: proto.ColumnType_BOOL, Description: "Whether the rule is enabled."},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "Message for generated signals."},
			{Name: "queries", Type: proto.ColumnType_STRING, Description: "Queries for selecting logs which are part of the rule."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The security monitoring rule type."},
			{Name: "update_author_id", Type: proto.ColumnType_STRING, Description: "User ID of the user who updated the rule."},
			{Name: "version", Type: proto.ColumnType_INT, Description: "The version of the rule."},

			// JSON columns
			{Name: "cases", Type: proto.ColumnType_JSON, Description: "Cases for generating signals."},
			{Name: "filters", Type: proto.ColumnType_JSON, Description: "Additional queries to filter matched events before they are processed."},
			{Name: "options", Type: proto.ColumnType_JSON, Description: "Additional options for security monitoring rules."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags for generated signals."},
		},
	}
}

func listSecurityMonitoringRules(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, _, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_security_monitoring_rule.listSecurityMonitoringRules", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/SecurityMonitoringApi.md#ListSecurityMonitoringRules
	opts := datadog.ListSecurityMonitoringRulesOptionalParameters{
		PageSize:   datadog.PtrInt64(100),
		PageNumber: datadog.PtrInt64(0),
	}

	count := int64(0)
	for {
		resp, _, err := apiClient.SecurityMonitoringApi.ListSecurityMonitoringRules(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("datadog_security_monitoring_rule.listSecurityMonitoringRules", "query_error", err)
			return nil, err
		}

		for _, securityMonitoringRule := range resp.GetData() {
			count++
			d.StreamListItem(ctx, securityMonitoringRule)
			// Check if context has been cancelled or if the limit has been hit (if specified)
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if count >= resp.Meta.Page.GetTotalCount() {
			return nil, nil
		}
		opts.WithPageNumber(*opts.PageNumber + 1)
	}

}

func getSecurityMonitoringRule(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var ruleID string

	if h.Item != nil {
		ruleID = *h.Item.(datadog.SecurityMonitoringRuleResponse).Id
	} else {
		ruleID = d.KeyColumnQualString("id")
	}

	if strings.TrimSpace(ruleID) == "" {
		return nil, nil
	}

	ctx, apiClient, _, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_security_monitoring_rule.getRole", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/SecurityMonitoringApi.md#getsecuritymonitoringrule
	resp, _, err := apiClient.SecurityMonitoringApi.GetSecurityMonitoringRule(ctx, ruleID)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_security_monitoring_rule.getRole", "query_error", err)
		if err.Error() == "404 Not Found" {
			return nil, nil
		}
		return nil, err
	}

	return resp, nil
}

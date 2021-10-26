package datadog

import (
	"context"
	"strings"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDatadogDashboard(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_dashboard",
		Description: "A dashboard is Datadog’s tool for visually tracking, analyzing, and displaying key performance metrics.",
		Get: &plugin.GetConfig{
			Hydrate:    getDashboard,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listDashboards,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Dashboard identifier."},
			{Name: "author_handle", Type: proto.ColumnType_STRING, Description: "Identifier of the dashboard author."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the dashboard."},
			{Name: "layout_type", Type: proto.ColumnType_STRING, Description: "Layout type of the dashboard. Can be on of \"free\" or \"ordered\"."},

			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description").Transform(valueFromNullable), Description: "Description of the dashboard."},
			{Name: "is_read_only", Type: proto.ColumnType_BOOL, Description: "Indicates if the dashboard is read-only. If True, only the author and admins can make changes to it."},
			{Name: "modified_at", Type: proto.ColumnType_TIMESTAMP, Description: "Modification time of the dashboard."},
			{Name: "reflow_type", Type: proto.ColumnType_STRING, Hydrate: getDashboard, Description: "DashboardReflowType Reflow type for a **new dashboard layout** dashboard. If set to 'fixed', the dashboard expects all widgets to have a layout, and if it's set to 'auto', widgets should not have layouts."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the dashboard."},

			// JSON columns
			{Name: "restricted_roles", Type: proto.ColumnType_JSON, Hydrate: getDashboard, Description: "A list of role identifiers. Only the author and users associated with at least one of these roles can edit this dashboard. Overrides the `is_read_only` property if both are present."},
			{Name: "template_variable_presets", Type: proto.ColumnType_JSON, Hydrate: getDashboard, Description: "List of template variables saved views."},
			{Name: "template_variables", Type: proto.ColumnType_JSON, Hydrate: getDashboard, Description: "List of template variables for this dashboard."},
			{Name: "widgets", Type: proto.ColumnType_JSON, Hydrate: getDashboard, Description: "List of widgets to display on the dashboard."},

			// common fields
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the dashboard."},
		},
	}
}

func listDashboards(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_dashboard.listDashboards", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v1/datadog/docs/DashboardsApi.md#listdashboards
	resp, _, err := apiClient.DashboardsApi.ListDashboards(ctx, datadog.ListDashboardsOptionalParameters{})
	if err != nil {
		plugin.Logger(ctx).Error("datadog_dashboard.listDashboards", "query_error", err)
	}

	for _, dashboard := range resp.GetDashboards() {
		d.StreamListItem(ctx, dashboard)
		// Check if context has been cancelled or if the limit has been hit (if specified)
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getDashboard(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var dasboardID string
	if h.Item != nil {
		dasboardID = *h.Item.(datadog.DashboardSummaryDefinition).Id
	} else {
		dasboardID = d.KeyColumnQuals["id"].GetStringValue()
	}

	if strings.TrimSpace(dasboardID) == "" {
		return nil, nil
	}

	ctx, apiClient, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_dashboard.getDashboard", "connection_error", err)
		return nil, err
	}

	resp, _, err := apiClient.DashboardsApi.GetDashboard(ctx, dasboardID)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_dashboard.getDashboard", "query_error", err)
		if err.Error() == "404 Not Found" {
			return nil, nil
		}
		return nil, err
	}

	return resp, nil
}

package datadog

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDatadogDashboard(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_dashboard",
		Description: "Users in Datadog.",
		List: &plugin.ListConfig{
			Hydrate: listDashboards,
			// KeyColumns: plugin.KeyColumnSlice{
			// 	{Name: "status", Require: plugin.Optional},
			// },
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Dashboard identifier."},
			{Name: "author_handle", Type: proto.ColumnType_STRING, Description: "Identifier of the dashboard author."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Creation date of the dashboard."},

			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromJSONTag().Transform(valueFromNullableString), Description: "Description of the dashboard."},
			{Name: "is_read_only", Type: proto.ColumnType_BOOL, Description: "Indicates if the dashboard is read-only. If True, only the author and admins can make changes to it."},
			{Name: "layout_type", Type: proto.ColumnType_STRING, Description: "Creation date of the dashboard."},
			{Name: "modified_at", Type: proto.ColumnType_TIMESTAMP, Description: "Modification time of the dashboard."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the dashboard."},

			// common fields
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the dashboard."},
		},
	}
}

func listDashboards(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_dashboard.listDashboards", "connection_error", err)
		return nil, err
	}

	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/UsersApi.md#listusers
	opts := datadog.ListDashboardsOptionalParameters{}

	resp, _, err := apiClient.DashboardsApi.ListDashboards(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_dashboard.listDashboards", "query_error", err)
	}

	for _, dashboard := range resp.GetDashboards() {
		d.StreamListItem(ctx, dashboard)
		// Check if context has been cancelled or if the limit has been hit (if specified)
		// if there is a limit, it will return the number of rows required to reach this limit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

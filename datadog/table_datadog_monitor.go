package datadog

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDatadogMonitor(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_monitor",
		Description: "Datadog monitor resource.",
		List: &plugin.ListConfig{
			Hydrate: listMonitors,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "name", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the monitor."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the monitor."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created"), Description: "Timestamp of the monitor creation."},
			{Name: "creator_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Creator.Email"), Description: "Email of the creator."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the monitor. For more information about `type`, see the [monitor options](https://docs.datadoghq.com/monitors/guide/monitor_api_options/) docs."},

			// Other useful columns
			{Name: "message", Type: proto.ColumnType_STRING, Description: "Timestamp of the monitor creation."},
			{Name: "modified_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Modified"), Description: "Last timestamp when the monitor was edited."},
			{Name: "multi", Type: proto.ColumnType_BOOL, Description: "Whether or not the monitor is broken down on different groups."},
			{Name: "overall_state", Type: proto.ColumnType_STRING, Description: "Current state of the monitor."},
			{Name: "priority", Type: proto.ColumnType_INT, Transform: transform.FromJSONTag().Transform(valueFromNullable), Description: "Integer from 1 (high) to 5 (low) indicating alert severity."},
			{Name: "query", Type: proto.ColumnType_STRING, Description: "The monitor query."},

			// JSON fields
			{Name: "options", Type: proto.ColumnType_JSON, Description: "A list of role identifiers that can be pulled from the Roles API. Cannot be used with `locked` option."},
			{Name: "restricted_roles", Type: proto.ColumnType_JSON, Description: "Relationships of the user object returned by the API."},
			{Name: "group_states", Type: proto.ColumnType_JSON, Transform: transform.FromField("State.Groups"), Description: "Dictionary where the keys are groups (comma separated lists of tags) and the values are the list of groups your monitor is broken down on."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags associated to monitor."},
		},
	}
}

func listMonitors(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_monitor.listMonitors", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v1/datadog/docs/MonitorsApi.md#listmonitors
	// page := int64(0) // int64 | The page to start paginating from. If this argument is not specified, the request returns all monitors without pagination.
	opts := datadog.ListMonitorsOptionalParameters{}

	name := d.KeyColumnQualString("name")
	if name != "" {
		opts.WithName(name)
	}

	resp, _, err := apiClient.MonitorsApi.ListMonitors(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_monitor.listMonitors", "query_error", err)
	}

	for _, monitor := range resp {
		d.StreamListItem(ctx, monitor)
		// Check if context has been cancelled or if the limit has been hit (if specified)
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

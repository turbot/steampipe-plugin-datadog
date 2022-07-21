package datadog

import (
	"context"
	"fmt"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableDatadogHost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_host",
		Description: "A host is any piece of infrastructure that runs an instance of the Datadog Agent such as a bare metal instance, a VM or a container.",
		List: &plugin.ListConfig{
			Hydrate: listHosts,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "name", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the host."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the host."},
			{Name: "aws_name", Type: proto.ColumnType_STRING, Description: "AWS name of the host."},
		},
	}
}

func listHosts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_host.listHosts", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v1/datadog/docs/MonitorsApi.md#listmonitors
	// page := int64(0) // int64 | The page to start paginating from. If this argument is not specified, the request returns all monitors without pagination.
	opts := datadog.ListHostsOptionalParameters{}

	name := d.KeyColumnQualString("name")
	if name != "" {
		opts.WithFilter(fmt.Sprintf("exact:host:%s", name))
	}

	resp, _, err := apiClient.HostsApi.ListHosts(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_monitor.listMonitors", "query_error", err)
		return nil, err
	}

	for _, host := range *resp.HostList {
		d.StreamListItem(ctx, host)
		// Check if context has been cancelled or if the limit has been hit (if specified)
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

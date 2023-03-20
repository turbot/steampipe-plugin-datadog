package datadog

import (
	"context"
	"fmt"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDatadogHost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_host",
		Description: "A host is any piece of infrastructure that runs an instance of the Datadog Agent such as a bare metal instance or a VM.",
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
			{Name: "up", Type: proto.ColumnType_BOOL, Description: "Whether the expected metrics for the host are being received or not."},

			// Other useful columns
			{Name: "aws_name", Type: proto.ColumnType_STRING, Description: "AWS name of the host."},
			{Name: "last_reported_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromGo().Transform(transform.UnixToTimestamp), Description: "Last time the host reported a metric data point."},
			{Name: "is_muted", Type: proto.ColumnType_BOOL, Description: "Whether or not the host is muted."},
			{Name: "mute_timeout", Type: proto.ColumnType_INT, Description: "The timeout of the mute applied to the host."},

			// JSON columns
			{Name: "aliases", Type: proto.ColumnType_JSON, Description: "An array of aliases that the host is known by such as AWS EC2 instance name, AWS internal IP address etc."},
			{Name: "apps", Type: proto.ColumnType_JSON, Description: "An array containing host apps such as system services, containers and more."},
			{Name: "meta", Type: proto.ColumnType_JSON, Description: "An object containing host metadata such as operating system and version."},
			{Name: "metrics", Type: proto.ColumnType_JSON, Description: "An object containing host metrics such as CPU, iowait and load."},
			{Name: "sources", Type: proto.ColumnType_JSON, Description: "An array containing the sources of the host metrics."},
			{Name: "tags_by_source", Type: proto.ColumnType_JSON, Description: "An object containing tags for each data source such as AWS, Datadog Agent etc."},
		},
	}
}

func listHosts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_host.listHosts", "connection_error", err)
		return nil, err
	}

	opts := datadog.ListHostsOptionalParameters{}

	name := d.KeyColumnQualString("name")
	hostNameFilter := fmt.Sprintf("exact:host:%s", name)
	if name != "" {
		opts.WithFilter(hostNameFilter)
	}

	opts.WithCount(1000)
	opts.WithIncludeHostsMetadata(true)
	opts.WithIncludeMutedHostsData(true)

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < 1000 {
			opts.WithCount(int64(*limit))
		}
	}

	count := int64(0)

	for {
		resp, _, err := apiClient.HostsApi.ListHosts(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("datadog_host.listHosts", "query_error", err)
			return nil, err
		}

		for _, host := range *resp.HostList {
			count++
			d.StreamListItem(ctx, host)
			// Check if context has been cancelled or if the limit has been hit (if specified)
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Break loop if no results are returned
		if len(*resp.HostList) == 0 {
			return nil, nil
		}
		// Set the start point for the next page
		opts.WithStart(count)
	} 
}

package datadog

import (
	"context"
	"fmt"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableDatadogHost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_host",
		Description: "A host is any piece of infrastructure that runs an instance of the Datadog Agent such as a bare metal instance or a VM.",
		// There is no Datadog hosts endpoint for getting a single host so instead we use a filter to perform an exact match
		Get: &plugin.GetConfig{
			Hydrate:    listHosts,
			KeyColumns: plugin.SingleColumn("name"),
		},
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
			{Name: "mute_timeout", Type: proto.ColumnType_BOOL, Description: "The timeout of the mute applied to the host."},

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

	paging := true
	count := int64(0)

	for paging {
		resp, _, err := apiClient.HostsApi.ListHosts(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("datadog_monitor.listMonitors", "query_error", err)
			return nil, err
		}

		for _, host := range *resp.HostList {
			count += resp.GetTotalReturned()
			d.StreamListItem(ctx, host)
			// Check if context has been cancelled or if the limit has been hit (if specified)
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Break loop if using a filter, which is guaranteed to be an exact match
		if name != "" {
			return nil, nil
		}
		// Break loop if host list is less than the maximum page size of 1000 items
		if count > 1000 {
			return nil, nil
		}
		opts.WithFrom(count)
	}

	return nil, nil
}

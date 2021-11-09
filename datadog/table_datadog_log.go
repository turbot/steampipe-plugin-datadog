package datadog

import (
	"context"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDatadogLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_log",
		Description: "Log-based metrics are a cost-efficient way to summarize log data from the entire ingest stream.",
		List: &plugin.ListConfig{
			Hydrate: listLogs,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "query", Require: plugin.Optional},
				{Name: "timestamp", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique ID of the Log."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Search query following logs syntax."},
			{Name: "host", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Host"), Description: "Name of the machine from where the logs are being sent."},
			{Name: "message", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Message"), Description: "The message https://docs.datadoghq.com/logs/log_collection/#reserved-attributes of your log. By default, Datadog ingests the value of the message attribute as the body of the log entry. That value is then highlighted and displayed in the Logstream, where it is indexed for full text search."},
			{Name: "service", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Service"), Description: "The name of the application or service generating the log events. It is used to switch from Logs to APM, so make sure you define the same value when you use both products."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Status"), Description: "Status of the message associated with your log."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.Timestamp"), Description: "Timestamp of log."},

			// JSON columns
			{Name: "attributes", Type: proto.ColumnType_JSON, Transform: transform.FromField("Attributes.Attributes"), Description: "JSON object of attributes from your log."},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Attributes.Tags"), Description: "Array of tags associated with your log."},
		},
	}
}

func listLogs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, _, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_log.listLogs", "connection_error", err)
		return nil, err
	}

	sort := datadog.LogsSort("timestamp")
	opts := *datadog.NewListLogsGetOptionalParameters()
	opts.WithSort(sort)
	opts.WithPageLimit(100)

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < int64(*opts.PageLimit) {
			opts.WithPageLimit(int32(*limit))
		}
	}

	query := d.KeyColumnQualString("query")
	if query != "" {
		opts.WithFilterQuery(query)
	}

	quals := d.Quals
	if quals["timestamp"] != nil {
		opts.WithFilterTo(time.Now())
		for _, q := range quals["timestamp"].Quals {
			timestamp := q.Value.GetTimestampValue().AsTime()
			switch q.Operator {
			case "=":
				opts.WithFilterFrom(timestamp)
				opts.WithFilterTo(timestamp)
			case ">=", ">":
				opts.WithFilterFrom(timestamp)
			case "<", "<=":
				opts.WithFilterTo(timestamp)
			}
		}
	}

	for {
		// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/LogsApi.md#listlogsget
		resp, _, err := apiClient.LogsApi.ListLogsGet(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("datadog_log.listLogs", "query_error", err)
		}

		for _, log := range resp.GetData() {
			d.StreamListItem(ctx, log)
			// Check if context has been cancelled or if the limit has been hit (if specified)
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if resp.HasLinks() {
			if resp.Links.HasNext() {
				opts.WithPageCursor(*resp.Meta.GetPage().After)
			} else {
				break
			}
		} else {
			break
		}
	}

	return nil, nil
}

// Example
// https://docs.datadoghq.com/logs/explorer/search_syntax/
// select * from datadog_log where query = '@detail.eventSource:s3.amazonaws.com' and timestamp >= (current_date - interval '2' day)

// By default API too pulls events only for last 15 minutes

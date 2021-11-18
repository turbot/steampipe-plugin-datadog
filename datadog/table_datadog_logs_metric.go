package datadog

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDatadogLogsMetric(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_logs_metric",
		Description: "Log-based metrics are a cost-efficient way to summarize log data from the entire ingest stream.",
		Get: &plugin.GetConfig{
			Hydrate:    getLogsMetric,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listLogsMetrics,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The name of the log-based metric."},
			{Name: "compute_aggregation_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Compute.AggregationType"), Description: "The type of aggregation to used for computing metric. Can be one of \"count\", \"distribution\"."},
			{Name: "compute_path", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Compute.Path"), Description: "The path to the value the log-based metric will aggregate on (only used if the aggregation type is a \"distribution\")."},
			{Name: "filter_query", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Filter.Query"), Description: "The search query - following the log search syntax to filter logs."},

			// JSON columns
			{Name: "group_by", Type: proto.ColumnType_JSON, Transform: transform.FromField("Attributes.GroupBy"), Description: "List of rules for the group by."},
		},
	}
}

func listLogsMetrics(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, _, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_logs_metric.listLogsMetrics", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/LogsMetricsApi.md#ListLogsMetrics
	resp, _, err := apiClient.LogsMetricsApi.ListLogsMetrics(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_logs_metric.listLogsMetrics", "query_error", err)
		return nil, err
	}

	for _, logMetric := range resp.GetData() {
		d.StreamListItem(ctx, logMetric)
		// Check if context has been cancelled or if the limit has been hit (if specified)
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getLogsMetric(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	metricID := d.KeyColumnQualString("id")
	ctx, apiClient, _, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_logs_metric.getLogsMetric", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/LogsMetricsApi.md#GetLogsMetric
	resp, _, err := apiClient.LogsMetricsApi.GetLogsMetric(ctx, metricID)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_logs_metric.getLogsMetric", "query_error", err)
		if err.Error() == "404 Not Found" {
			return nil, nil
		}
		return nil, err
	}

	return resp.GetData(), nil
}

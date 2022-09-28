package datadog

import (
	"context"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDatadogSecurityMonitoringSignal(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_security_monitoring_signal",
		Description: "Signals are threats detected based on a security monitoring rule.",
		List: &plugin.ListConfig{
			Hydrate: listSecurityMonitoringSignals,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "timestamp", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
				{Name: "filter_query", Require: plugin.Optional},
			},
		},

		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique ID of the security signal."},
			{Name: "message", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Message"), Description: "The message in the security signal defined by the rule that generated the signal."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.Timestamp"), Description: "The timestamp of the security signal."},
			{Name: "filter_query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("filter_query"), Description: "The search query for security signals. For more information refer https://docs.datadoghq.com/security_platform/explorer/"},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Attributes").Transform(signalTitle), Description: "Title of the security signal"},

			// JSON columns
			{Name: "attributes", Type: proto.ColumnType_JSON, Transform: transform.FromField("Attributes.Attributes"), Description: "A JSON object of attributes in the security signal."},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Attributes.Tags"), Description: "An array of tags associated with the security signal."},
		},
	}
}

func listSecurityMonitoringSignals(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, _, configuration, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_security_monitoring_signal.listSecurityMonitoringSignals", "connection_error", err)
		return nil, err
	}
	filterQuery := d.KeyColumnQualString("filter_query")
	filterFrom := time.Now().AddDate(0, 0, -1) // By default list signals for last one day
	filterTo := time.Now()
	pageLimit := int32(50)

	// Update page limit if data is requested with a limit clause in query
	limit := d.QueryContext.Limit
	if limit != nil && int32(*limit) < pageLimit {
		pageLimit = int32(*limit)
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/SecurityMonitoringApi.md#listsecuritymonitoringsignals
	opts := datadog.ListSecurityMonitoringSignalsOptionalParameters{
		FilterFrom: &filterFrom,
		FilterTo:   &filterTo,
		PageLimit:  &pageLimit,
	}

	// Sort results with timestamp
	opts.WithSort(datadog.SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING)

	if filterQuery != "" {
		opts.WithFilterQuery(filterQuery)
	}

	quals := d.Quals

	if quals["timestamp"] != nil {
		for _, q := range quals["timestamp"].Quals {
			// Subtracted 1 minute to FilterFrom time and Added 1 minute to FilterTo time to miss any results due to time conersions in steampipe
			switch q.Operator {
			case "=":
				opts.WithFilterFrom(q.Value.GetTimestampValue().AsTime().Add(-60 * time.Second))
				opts.WithFilterTo(q.Value.GetTimestampValue().AsTime().Add(60 * time.Second))
			case ">=", ">":
				opts.WithFilterFrom(q.Value.GetTimestampValue().AsTime().Add(-60 * time.Second))
			case "<", "<=":
				opts.WithFilterTo(q.Value.GetTimestampValue().AsTime().Add(60 * time.Second))
			}
		}
	}

	if opts.FilterTo == nil {
		opts.WithFilterTo(time.Now())
	}

	configuration.SetUnstableOperationEnabled("ListSecurityMonitoringSignals", true)
	apiClient := datadog.NewAPIClient(configuration)

	for {
		resp, _, err := apiClient.SecurityMonitoringApi.ListSecurityMonitoringSignals(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("datadog_security_monitoring_signal.listSecurityMonitoringSignals", "query_error", err)
			return nil, err
		}

		for _, securityMonitoringSignal := range resp.GetData() {
			d.StreamListItem(ctx, securityMonitoringSignal)
			// Check if context has been cancelled or if the limit has been hit (if specified)
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if meta, ok := resp.GetMetaOk(); ok {
			if page, pageOk := meta.GetPageOk(); pageOk {
				if page.HasAfter() {
					opts.WithPageCursor(page.GetAfter())
				}
			}
		} else {
			break
		}
	}
	return nil, nil
}

//// TRANSFORM FUNCTION

func signalTitle(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}

	attributes := d.Value.(*map[string]interface{})
	data := *attributes

	if data[d.ColumnName] != nil {
		return data[d.ColumnName], nil
	}
	return nil, nil
}

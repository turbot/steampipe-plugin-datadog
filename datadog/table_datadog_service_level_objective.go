package datadog

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDatadogServiceLevelObjective(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_service_level_objective",
		Description: "A SLO provides a target percentage of a specific metric over a certain period of time.",
		List: &plugin.ListConfig{
			Hydrate: listSLOs,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "name", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the SLO."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the SLO."},
			{Name: "creator_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Creator.Email"), Description: "Email of the creator."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(convertDatetime), Description: "Timestamp of the SLO creation."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the SLO. For more information about type, see https://docs.datadoghq.com/monitors/service_level_objectives/."},

			// Other useful columns
			{Name: "modified_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("ModifiedAt").Transform(convertDatetime), Description: "Last timestamp when the monitor was edited."},
			{Name: "monitor_ids", Type: proto.ColumnType_STRING, Description: "The monitor based SLO have monitors assiciated with them. Shows list of associated monitors"},
			{Name: "query", Type: proto.ColumnType_STRING, Description: "The Metric based SLOs use querys to detiremine state. Shows associated query"},

			// JSON columns
			{Name: "monitor_tags", Type: proto.ColumnType_JSON, Description: "If monitors are associated with SLO have tags they will show here."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags associated to SLO."},
			{Name: "thresholds", Type: proto.ColumnType_JSON, Description: "Thresholds that are set for the SLOs."},
		},
	}
}

func listSLOs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_slo.listSLOs", "connection_error", err)
		return nil, err
	}

	opts := datadog.ListSLOsOptionalParameters{}

	id := d.KeyColumnQualString("id")
	if id != "" {
		opts.WithIds(id)
	}

	resp, _, err := apiClient.ServiceLevelObjectivesApi.ListSLOs(ctx, opts)

	if err != nil {
		plugin.Logger(ctx).Error("datadog_slo.listSLOs", "query_error", err)
		return nil, err
	}

	for _, slo := range resp.GetData() {
		d.StreamListItem(ctx, slo)
		// Check if context has been cancelled or if the limit has been hit (if specified)
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

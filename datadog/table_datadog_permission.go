package datadog

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDatadogPermission(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_permission",
		Description: "Permissions provide the base level of access for roles.",
		List: &plugin.ListConfig{
			Hydrate: listPermissions,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Name"), Description: "Name of the permission."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id of the permission."},
			{Name: "restricted", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Attributes.Restricted"), Description: "Whether or not the permission is restricted."},
			{Name: "group_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.GroupName"), Description: "Name of the permission group."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.Created"), Description: "Creation time of the permission."},

			// Other columns
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Description"), Description: "Description of the permission."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.DisplayName"), Description: "Displayed name for the permission."},
			{Name: "display_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.DisplayType"), Description: "Displayed type the permission."},
		},
	}
}

func listPermissions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, _, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_permission.listPermissions", "connection_error", err)
		return nil, err
	}

	// https: //github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/RolesApi.md#ListPermissions
	resp, _, err := apiClient.RolesApi.ListPermissions(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_permission.listPermissions", "query_error", err)
	}

	for _, permission := range resp.GetData() {
		d.StreamListItem(ctx, permission)
		// Check if context has been cancelled or if the limit has been hit (if specified)
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

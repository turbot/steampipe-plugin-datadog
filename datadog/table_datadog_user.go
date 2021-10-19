package datadog

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDatadogUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_user",
		Description: "Users in Datadog.",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "status", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Email"), Description: "Email of the user."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id of the user."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Name").Transform(ValueFromNullableStrint), Description: "Name of the user."},
			{Name: "handle", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Handle"), Description: "Handle of the user."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.CreatedAt"), Description: "Creation time of the user."},

			// Other useful columns
			{Name: "disabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Attributes.Disabled"), Description: "Indicates if the user is disabled."},
			{Name: "icon", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Icon"), Description: "URL of the user's icon."},
			{Name: "modified_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.ModifiedAt"), Description: "Time that the user was last modified."},
			{Name: "service_account", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Attributes.ServiceAccount"), Description: "Indicates if the user is a service account."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Status"), Description: "Status of the user."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Title").Transform(ValueFromNullableStrint), Description: "Title of the user."},
			{Name: "verified", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Attributes.Verified"), Description: "Indicates the verification status of the user."},

			// JSON fields
			{Name: "roles", Type: proto.ColumnType_JSON, Transform: transform.FromField("Relationships.Roles"), Description: "A list containing type and ID of a role attached to user."},
			{Name: "relationships", Type: proto.ColumnType_JSON, Description: "Relationships of the user object returned by the API."},
		},
	}
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_user.listUsers", "connection_error", err)
		return nil, err
	}

	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/UsersApi.md#listusers
	opts := datadog.ListUsersOptionalParameters{
		PageSize:   datadog.PtrInt64(int64(100)),
		PageNumber: datadog.PtrInt64(int64(0)),
		Sort:       datadog.PtrString("name"),
		// Filter:     &filter, // Need to explore this field
	}

	fiterStatus := d.KeyColumnQualString("status")
	if fiterStatus != "" {
		opts.FilterStatus = datadog.PtrString(fiterStatus)
	}

	paging := true
	count := int64(0)

	for paging {
		resp, _, err := apiClient.UsersApi.ListUsers(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("datadog_user.listUsers", "query_error", err)
		}

		for _, user := range resp.GetData() {
			count++
			d.StreamListItem(ctx, user)
			// Check if context has been cancelled or if the limit has been hit (if specified)
			// if there is a limit, it will return the number of rows required to reach this limit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if count >= resp.Meta.Page.GetTotalCount() {
			paging = false
		}
		opts.PageNumber = datadog.PtrInt64(*opts.PageNumber + 1)
	}

	return nil, nil
}

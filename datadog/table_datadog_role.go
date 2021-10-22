package datadog

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDatadogRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_role",
		Description: "Datadog role resource.",
		List: &plugin.ListConfig{
			Hydrate: listRoles,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Name"), Description: "Name of the role."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Id of the role."},
			{Name: "user_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Attributes.UserCount"), Description: "Number of users associated with the role."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.CreatedAt"), Description: "Creation time of the role."},
			{Name: "modified_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.ModifiedAt"), Description: "Time of last role modification."},

			// JSON fields
			{Name: "users", Type: proto.ColumnType_JSON, Hydrate: listRoleUsers, Transform: transform.From(userList), Description: "Set of objects containing the permission ID and the name of the permissions granted to this role."},
			{Name: "permissions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Relationships.Permissions.Data"), Description: "List of users emails attached to role."},
		},
	}
}

func listRoles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_role.listRoles", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/RolesApi.md#listroles
	opts := datadog.ListRolesOptionalParameters{
		PageSize:   datadog.PtrInt64(int64(100)),
		PageNumber: datadog.PtrInt64(int64(0)),
	}

	// fiterStatus := d.KeyColumnQualString("status")
	// if fiterStatus != "" {
	// 	opts.FilterStatus = datadog.PtrString(fiterStatus)
	// }

	paging := true
	count := int64(0)

	for paging {
		resp, _, err := apiClient.RolesApi.ListRoles(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("datadog_role.listRoles", "query_error", err)
		}

		for _, role := range resp.GetData() {
			count++
			d.StreamListItem(ctx, role)
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

func listRoleUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	role := h.Item.(datadog.Role)
	ctx, apiClient, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_role.listRoleUsers", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/RolesApi.md#listroles
	opts := datadog.ListRoleUsersOptionalParameters{
		PageSize:   datadog.PtrInt64(int64(100)),
		PageNumber: datadog.PtrInt64(int64(0)),
	}

	paging := true
	count := int64(0)
	var users []datadog.User

	for paging {
		resp, _, err := apiClient.RolesApi.ListRoleUsers(ctx, *role.Id, opts)
		if err != nil {
			plugin.Logger(ctx).Error("datadog_role.listRoleUsers", "query_error", err)
		}

		noOfUsers := len(resp.GetData())
		users = append(users, resp.GetData()...)
		count += int64(noOfUsers)

		if count >= resp.Meta.Page.GetTotalCount() {
			paging = false
		}
		opts.PageNumber = datadog.PtrInt64(*opts.PageNumber + 1)
	}

	return users, nil
}

//// TRANSFORM FUNCTION

func userList(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	users := d.HydrateItem.([]datadog.User)

	var user_emails []string

	for _, user := range users {
		user_emails = append(user_emails, *user.Attributes.Email)
	}

	return user_emails, nil
}

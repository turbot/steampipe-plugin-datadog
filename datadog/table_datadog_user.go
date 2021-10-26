package datadog

import (
	"context"
	"strings"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDatadogUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_user",
		Description: "Datadog dashboard resource.",
		Get: &plugin.GetConfig{
			Hydrate:    getUser,
			KeyColumns: plugin.SingleColumn("id"),
		},
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
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Name").Transform(valueFromNullable), Description: "Name of the user."},
			{Name: "handle", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Handle"), Description: "Handle of the user."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.CreatedAt"), Description: "Creation time of the user."},

			// Other useful columns
			{Name: "disabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Attributes.Disabled"), Description: "Indicates if the user is disabled."},
			{Name: "icon", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Icon"), Description: "URL of the user's icon."},
			{Name: "modified_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.ModifiedAt"), Description: "Time that the user was last modified."},
			{Name: "service_account", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Attributes.ServiceAccount"), Description: "Indicates if the user is a service account."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Status"), Description: "Status of the user."},
			{Name: "verified", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Attributes.Verified"), Description: "Indicates the verification status of the user."},

			// JSON fields
			{Name: "role_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Relationships.Roles.Data").Transform(roleList), Description: "A list containing id of roles attached to user."},
			{Name: "relationships", Type: proto.ColumnType_JSON, Description: "Relationships of the user object returned by the API."},

			// common fields
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Title").Transform(valueFromNullable), Description: "Title of the user."},
		},
	}
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, _, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_user.listUsers", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/UsersApi.md#listusers
	opts := datadog.ListUsersOptionalParameters{
		PageSize:   datadog.PtrInt64(int64(100)),
		PageNumber: datadog.PtrInt64(int64(0)),
		Sort:       datadog.PtrString("name"),
		// Filter:     &filter, //TODO Need to explore this field
	}

	fiterStatus := d.KeyColumnQualString("status")
	if fiterStatus != "" {
		opts.WithFilterStatus(fiterStatus)
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
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Break loop if using filter
		if resp.Meta.Page.HasTotalFilteredCount() {
			if count >= resp.Meta.Page.GetTotalFilteredCount() {
				return nil, nil
			}
		}
		// Break loop if not using filter
		if count >= resp.Meta.Page.GetTotalCount() {
			return nil, nil
		}
		opts.WithPageNumber(*opts.PageNumber + 1)
	}

	return nil, nil
}

func getUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var userID string

	if h.Item != nil {
		userID = *h.Item.(datadog.User).Id
	} else {
		userID = d.KeyColumnQualString("id")
	}

	if strings.TrimSpace(userID) == "" {
		return nil, nil
	}

	ctx, apiClient, _, err := connectV2(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_role.getUser", "connection_error", err)
		return nil, err
	}

	// https: //github.com/DataDog/datadog-api-client-go/blob/master/api/v2/datadog/docs/UsersApi.md#GetUser
	resp, _, err := apiClient.UsersApi.GetUser(ctx, userID)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_role.getUser", "query_error", err)
		if err.Error() == "404 Not Found" {
			return nil, nil
		}
		return nil, err
	}

	return resp.GetData(), nil
}

//// TRANSFORM FUNCTION

func roleList(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	roles := d.Value.(*[]datadog.RelationshipToRoleData)

	var roleIds []string

	for _, role := range *roles {
		roleIds = append(roleIds, *role.Id)
	}

	return roleIds, nil
}

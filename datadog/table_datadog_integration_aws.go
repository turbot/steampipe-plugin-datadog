package datadog

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableDatadogIntegrationAws(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_integration_aws",
		Description: "Datadog AWS integration resource.",
		List: &plugin.ListConfig{
			Hydrate: listAWSIntegrations,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "account_id", Require: plugin.Optional},
				{Name: "role_name", Require: plugin.Optional},
				{Name: "access_key_id", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "account_id", Type: proto.ColumnType_STRING, Description: "Your AWS Account ID without dashes."},
			{Name: "role_name", Type: proto.ColumnType_STRING, Description: "Your Datadog role delegation name."},
			{Name: "access_key_id", Type: proto.ColumnType_STRING, Description: "Your AWS access key ID. Only required if your AWS account is a GovCloud or China account."},
			{Name: "secret_access_key", Type: proto.ColumnType_STRING, Description: "Your AWS secret access key. Only required if your AWS account is a GovCloud or China account."},

			// Other useful columns
			{Name: "cspm_resource_collection_enabled", Type: proto.ColumnType_BOOL, Description: "Whether Datadog collects cloud security posture management resources from your AWS account. This includes additional resources not covered under the general `resource_collection`."},
			{Name: "metrics_collection_enabled", Type: proto.ColumnType_BOOL, Description: "Whether Datadog collects metrics for this AWS account."},
			{Name: "resource_collection_enabled", Type: proto.ColumnType_BOOL, Description: "Whether Datadog collects a standard set of resources from your AWS account."},

			// JSON fields
			{Name: "account_specific_namespace_rules", Type: proto.ColumnType_JSON, Description: "An object, that enables or disables metric collection for specific AWS namespaces for this AWS account only."},
			{Name: "excluded_regions", Type: proto.ColumnType_JSON, Description: "An array of AWS regions to exclude from metrics collection."},
			{Name: "filter_tags", Type: proto.ColumnType_JSON, Description: "The array of EC2 tags (in the form `key:value`) defines a filter that Datadog uses when collecting metrics from EC2. Wildcards, such as `?` (for single characters) and `*` (for multiple characters) can also be used. Only hosts that match one of the defined tags will be imported into Datadog. The rest will be ignored. Host matching a given tag can also be excluded by adding `!` before the tag. For example, `env:production,instance-type:c1.*,!region:us-east-1`"},
			{Name: "host_tags", Type: proto.ColumnType_JSON, Description: "Array of tags (in the form `key:value`) to add to all hosts and metrics reporting through this integration."},
		},
	}
}

func listAWSIntegrations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_integration_aws.listAWSIntegrations", "connection_error", err)
		return nil, err
	}

	// https://github.com/DataDog/datadog-api-client-go/blob/master/api/v1/datadog/docs/AWSIntegrationApi.md#ListAWSAccounts
	opts := datadog.ListAWSAccountsOptionalParameters{}
	if d.KeyColumnQualString("account_id") != "" {
		opts.WithAccountId(d.KeyColumnQualString("account_id"))
	}
	if d.KeyColumnQualString("role_name") != "" {
		opts.WithRoleName(d.KeyColumnQualString("role_name"))
	}
	if d.KeyColumnQualString("access_key_id") != "" {
		opts.WithAccessKeyId(d.KeyColumnQualString("access_key_id"))
	}

	// Paging not supported by this API as of Date 10-25-2021
	resp, _, err := apiClient.AWSIntegrationApi.ListAWSAccounts(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_integration_aws.listAWSIntegrations", "query_error", err)
	}

	for _, account := range resp.GetAccounts() {
		d.StreamListItem(ctx, account)
		// Check if context has been cancelled or if the limit has been hit (if specified)
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

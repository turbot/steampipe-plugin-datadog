package datadog

import (
	"context"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/pkg/errors"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (context.Context, error) {

	// Load connection from cache, which preserves throttling protection etc
	// Not sure if we should cache this --  as we are mod0fying the context in this function
	// cacheKey := "datadog_connect"
	// if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
	// 	return cachedData.(context.Context), nil
	// }

	// Default to the env var settings
	apiKey := os.Getenv("DD_CLIENT_API_KEY")
	AppKey := os.Getenv("DD_CLIENT_APP_KEY")

	// Prefer config settings
	config := GetConfig(d.Connection)

	if config.APIKey != nil {
		apiKey = *config.APIKey
	}
	if config.AppKey != nil {
		AppKey = *config.AppKey
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, errors.New("api_key must be configured")
	}

	if AppKey == "" {
		return nil, errors.New("app_key must be configured")
	}

	ctx = context.WithValue(ctx, datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {Key: apiKey},
			"appKeyAuth": {Key: AppKey},
		},
	)

	ctx = context.WithValue(
		ctx,
		datadog.ContextServerVariables,
		map[string]string{"basePath": "v2"},
	)

	return ctx, nil
}

//// TRANSFORM FUNCTIONS

func ValueFromNullableStrint(_ context.Context, d *transform.TransformData) (interface{}, error) {
	nullableString := d.Value.(datadog.NullableString)
	if nullableString.IsSet() {
		return nullableString.Get(), nil
	}

	return nil, nil
}

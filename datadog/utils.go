package datadog

import (
	"context"
	"os"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	datadogV2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/pkg/errors"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func connectV1(ctx context.Context, d *plugin.QueryData) (context.Context, error) {

	// Load connection from cache, which preserves throttling protection etc
	// Not sure if we should cache this --  as we are modifying the context in this function
	// cacheKey := "datadog_connect"
	// if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
	// 	return cachedData.(context.Context), nil
	// }

	// Default to the env var settings
	apiKey := os.Getenv("DD_CLIENT_API_KEY")
	appKey := os.Getenv("DD_CLIENT_APP_KEY")

	// Prefer config settings
	config := GetConfig(d.Connection)

	if config.APIKey != nil {
		apiKey = *config.APIKey
	}
	if config.AppKey != nil {
		appKey = *config.AppKey
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, errors.New("api_key must be configured")
	}

	if appKey == "" {
		return nil, errors.New("app_key must be configured")
	}

	ctx = context.WithValue(ctx, datadogV1.ContextAPIKeys,
		map[string]datadogV1.APIKey{
			"apiKeyAuth": {Key: apiKey},
			"appKeyAuth": {Key: appKey},
		},
	)

	ctx = context.WithValue(
		ctx,
		datadogV1.ContextServerVariables,
		map[string]string{"basePath": "v2"},
	)

	return ctx, nil
}

func connectV2(ctx context.Context, d *plugin.QueryData) (context.Context, error) {

	// Load connection from cache, which preserves throttling protection etc
	// Not sure if we should cache this --  as we are modifying the context in this function
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

	ctx = context.WithValue(ctx, datadogV2.ContextAPIKeys,
		map[string]datadogV2.APIKey{
			"apiKeyAuth": {Key: apiKey},
			"appKeyAuth": {Key: AppKey},
		},
	)

	ctx = context.WithValue(
		ctx,
		datadogV2.ContextServerVariables,
		map[string]string{"basePath": "v2"},
	)

	return ctx, nil
}

//// TRANSFORM FUNCTIONS

func valueFromNullable(_ context.Context, d *transform.TransformData) (interface{}, error) {
	switch item := d.Value.(type) {
	// datadogV1
	case datadogV1.NullableString:
		if item.IsSet() {
			return item.Get(), nil
		}
	case datadogV1.NullableTime:
		if item.IsSet() {
			return item.Get(), nil
		}
	case datadogV1.NullableInt:
		if item.IsSet() {
			return item.Get(), nil
		}
	case datadogV1.NullableInt32:
		if item.IsSet() {
			return item.Get(), nil
		}
	case datadogV1.NullableInt64:
		if item.IsSet() {
			return item.Get(), nil
		}
	// datadogV2
	case datadogV2.NullableString:
		if item.IsSet() {
			return item.Get(), nil
		}
	case datadogV2.NullableTime:
		if item.IsSet() {
			return item.Get(), nil
		}
	case datadogV2.NullableInt:
		if item.IsSet() {
			return item.Get(), nil
		}
	case datadogV2.NullableInt32:
		if item.IsSet() {
			return item.Get(), nil
		}
	case datadogV2.NullableInt64:
		if item.IsSet() {
			return item.Get(), nil
		}
	}
	return nil, nil
}

package datadog

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	datadogV2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"

	"github.com/pkg/errors"
	"github.com/turbot/steampipe-plugin-datadog/internal/transport"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func connectV1(ctx context.Context, d *plugin.QueryData) (context.Context, *datadogV1.APIClient, error) {

	// Load connection from cache, which preserves throttling protection etc
	// Not sure if we should cache this --  as we are modifying the context in this function
	// cacheKey := "datadog_connect"
	// if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
	// 	return cachedData.(context.Context), nil
	// }

	// Default to the env var settings
	apiKey := os.Getenv("DD_CLIENT_API_KEY")
	appKey := os.Getenv("DD_CLIENT_APP_KEY")
	apiURL := "https://api.datadoghq.com/"

	// Prefer config settings
	config := GetConfig(d.Connection)

	if config.APIKey != nil {
		apiKey = *config.APIKey
	}
	if config.AppKey != nil {
		appKey = *config.AppKey
	}
	if config.ApiURL != nil {
		apiURL = *config.ApiURL
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, nil, errors.New("api_key must be configured")
	}

	if appKey == "" {
		return nil, nil, errors.New("app_key must be configured")
	}

	ctx = context.WithValue(ctx, datadogV1.ContextAPIKeys,
		map[string]datadogV1.APIKey{
			"apiKeyAuth": {Key: apiKey},
			"appKeyAuth": {Key: appKey},
		},
	)

	if apiURL != "" {
		parsedAPIURL, parseErr := url.Parse(apiURL)
		if parseErr != nil {
			return nil, nil, fmt.Errorf(`invalid API URL : %v`, parseErr)
		}
		if parsedAPIURL.Host == "" || parsedAPIURL.Scheme == "" {
			return nil, nil, fmt.Errorf(`missing protocol or host : %v`, apiURL)
		}

		strings.Split(parsedAPIURL.Host, "/")
		ctx = context.WithValue(ctx,
			datadogV1.ContextServerVariables,
			map[string]string{
				"name":     parsedAPIURL.Host,
				"protocol": parsedAPIURL.Scheme,
			})
	}

	ctx = context.WithValue(
		ctx,
		datadogV1.ContextServerVariables,
		map[string]string{"basePath": "v2"},
	)

	// Modify default client for retry handling
	httpClientV1 := http.DefaultClient
	ctOptions := transport.CustomTransportOptions{}
	timeout := time.Duration(int64(60)) * time.Second
	ctOptions.Timeout = &timeout
	httpClientV1.Transport = transport.NewCustomTransport(httpClientV1.Transport, ctOptions)

	configuration := datadogV1.NewConfiguration()
	configuration.HTTPClient = httpClientV1
	configuration.UserAgent = "Steampipe"
	apiClient := datadogV1.NewAPIClient(configuration)

	return ctx, apiClient, nil
}

func connectV2(ctx context.Context, d *plugin.QueryData) (context.Context, *datadogV2.APIClient, *datadogV2.Configuration, error) {

	// Load connection from cache, which preserves throttling protection etc
	// Not sure if we should cache this --  as we are modifying the context in this function
	// cacheKey := "datadog_connect"
	// if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
	// 	return cachedData.(context.Context), nil
	// }

	// Default to the env var settings
	apiKey := os.Getenv("DD_CLIENT_API_KEY")
	AppKey := os.Getenv("DD_CLIENT_APP_KEY")
	apiURL := "https://api.datadoghq.com/"

	// Prefer config settings
	config := GetConfig(d.Connection)

	if config.APIKey != nil {
		apiKey = *config.APIKey
	}
	if config.AppKey != nil {
		AppKey = *config.AppKey
	}
	if config.ApiURL != nil {
		apiURL = *config.ApiURL
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, nil, nil, errors.New("api_key must be configured")
	}

	if AppKey == "" {
		return nil, nil, nil, errors.New("app_key must be configured")
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

	if apiURL != "" {
		parsedAPIURL, parseErr := url.Parse(apiURL)
		if parseErr != nil {
			return nil, nil, nil, fmt.Errorf(`invalid API URL : %v`, parseErr)
		}
		if parsedAPIURL.Host == "" || parsedAPIURL.Scheme == "" {
			return nil, nil, nil, fmt.Errorf(`missing protocol or host : %v`, apiURL)
		}

		strings.Split(parsedAPIURL.Host, "/")
		ctx = context.WithValue(ctx,
			datadogV2.ContextServerVariables,
			map[string]string{
				"name":     parsedAPIURL.Host,
				"protocol": parsedAPIURL.Scheme,
			})
	}

	// Modify default client for rety handling
	httpClientV2 := http.DefaultClient
	ctOptions := transport.CustomTransportOptions{}
	timeout := time.Duration(int64(60)) * time.Second
	ctOptions.Timeout = &timeout
	httpClientV2.Transport = transport.NewCustomTransport(httpClientV2.Transport, ctOptions)

	configuration := datadogV2.NewConfiguration()
	configuration.HTTPClient = httpClientV2
	configuration.UserAgent = "Steampipe"
	apiClient := datadogV2.NewAPIClient(configuration)

	return ctx, apiClient, configuration, nil
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

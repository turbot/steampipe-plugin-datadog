package datadog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDatadogServiceLevelObjective(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "datadog_service_level_objective",
		Description: "An SLO(Service Level Objective) provides a target percentage of a specific metric over a certain period of time.",
		Get: &plugin.GetConfig{
			Hydrate:    getSLO,
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: listSLOs,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the SLO.", Transform: transform.FromField("Attributes.Name").Transform(transform.NullIfZeroValue)},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromGo(), Description: "ID of the SLO."},
			{Name: "creator_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.Creator.Email").Transform(transform.NullIfZeroValue), Description: "Email of the creator."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.CreatedAt").Transform(transform.NullIfZeroValue).Transform(convertDatetime), Description: "Timestamp of the SLO creation."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Attributes.SLOType").Transform(transform.NullIfZeroValue), Description: "The type of the SLO. For more information about type, see https://docs.datadoghq.com/monitors/service_level_objectives/."},

			// Other useful columns
			{Name: "modified_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Attributes.ModifiedAt").Transform(transform.NullIfZeroValue).Transform(convertDatetime), Description: "Last timestamp when the monitor was edited."},

			// JSON columns
			{Name: "configured_alert_ids", Type: proto.ColumnType_JSON, Hydrate: getSLO, Description: "Get the IDs of SLO monitors that reference this SLO."},
			{Name: "description", Type: proto.ColumnType_JSON, Description: "Description of the SLO.", Transform: transform.FromField("Attributes.Description").Transform(transform.NullIfZeroValue)},
			{Name: "groups", Type: proto.ColumnType_JSON, Description: "A list of (up to 20) monitor groups that narrow the scope of a monitor service level objective.", Transform: transform.FromField("Attributes.Groups").Transform(transform.NullIfZeroValue)},
			{Name: "monitor_ids", Type: proto.ColumnType_JSON, Description: "A list of monitor ids that defines the scope of a monitor service level objective.", Transform: transform.FromField("Attributes.MonitorIDs").Transform(transform.NullIfZeroValue)},
			{Name: "query", Type: proto.ColumnType_JSON, Description: "The Metric based SLOs use queries to determine the state. Shows associated query.", Transform: transform.FromField("Attributes.Query").Transform(transform.NullIfZeroValue)},
			{Name: "monitor_tags", Type: proto.ColumnType_JSON, Description: "If monitors that are associated with SLO have tags they will show here.", Transform: transform.FromField("Attributes.MonitorTags").Transform(transform.NullIfZeroValue)},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "Tags associated with SLO.", Transform: transform.FromField("Attributes.AllTags").Transform(transform.NullIfZeroValue)},
			{Name: "thresholds", Type: proto.ColumnType_JSON, Description: "Thresholds that are set for the SLOs.", Transform: transform.FromField("Attributes.Thresholds").Transform(transform.NullIfZeroValue)},
		},
	}
}

// Using `apiClient.ServiceLevelObjectivesApi.SearchSLO` doesn't populate the response and returns empty rows.
// Additionally, `apiClient.ServiceLevelObjectivesApi.ListSLOs(ctx, opts)` does not return SLOs of the "By Time Slice" type.
// Therefore, we opted for a raw API call to list all the SLOs (as is done in the Datadog console). This approach also addresses the issue: https://github.com/turbot/steampipe-plugin-datadog/issues/63.

func listSLOs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	ctx, apiClient, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_service_level_objective.listSLOs", "connection_error", err)
		return nil, err
	}

	steampipeConfig := GetConfig(d.Connection)

	keys := ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey)
	apiKey := keys["apiKeyAuth"].Key
	appKey := keys["appKeyAuth"].Key

	pageNumber := 0
	for {
		params := url.Values{}
		params.Add("query", "")
		params.Add("sort", "")
		params.Add("include_facets", "false")
		params.Add("include_permissions", "true")
		params.Add("page[size]", "1")
		params.Add("page[number]", fmt.Sprint(pageNumber))

		fullUrl := fmt.Sprintf("%s?%s", *steampipeConfig.ApiURL+"api/v1/slo/search", params.Encode())

		buildReq, err := http.NewRequest("GET", fullUrl, nil)
		if err != nil {
			plugin.Logger(ctx).Error("Error creating the request:", err)
			return nil, err
		}

		// Set headers
		buildReq.Header.Set("Content-Type", "application/json")
		buildReq.Header.Set("Accept", "*/*")
		buildReq.Header.Set("DD-API-KEY", apiKey)
		buildReq.Header.Set("DD-APPLICATION-KEY", appKey)

		response, err := SearchSLO(apiClient, buildReq)
		if err != nil {
			plugin.Logger(ctx).Error("datadog_service_level_objective.listSLOs", "api_error", err)
			return nil, err
		}

		if response == nil || response.Data.Attributes.SLOs == nil {
			break
		}

		for _, slo := range response.Data.Attributes.SLOs {
			if slo.Data != nil {
				d.StreamListItem(ctx, slo.Data)

				// Check if context has been cancelled or if the limit has been hit (if specified)
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}

		// Check if we should continue pagination
		if response.Meta.Pagination == nil || response.Meta.Pagination.LastNumber == nil || response.Meta.Pagination.LastNumber == nil || *response.Meta.Pagination.LastNumber == pageNumber {
			break
		}
		pageNumber++
	}

	return nil, nil
}

func SearchSLO(client *datadog.APIClient, req *http.Request) (*ApiResponse, error) {
	resp, err := client.CallAPI(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &ApiResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func getSLO(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var sloID string
	if h.Item != nil {
		sloID = h.Item.(*SLOData).ID
	} else {
		sloID = d.EqualsQuals["id"].GetStringValue()
	}

	// Empty value check
	if sloID == "" {
		return nil, nil
	}

	withConfiguredAlertIds := true
	opts := datadog.GetSLOOptionalParameters{
		WithConfiguredAlertIds: &withConfiguredAlertIds,
	}

	ctx, apiClient, err := connectV1(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("datadog_service_level_objective.getSLO", "connection_error", err)
		return nil, err
	}

	resp, _, err := apiClient.ServiceLevelObjectivesApi.GetSLO(ctx, sloID, opts)
	if err != nil {
		if err.Error() == "404 Not Found" {
			return nil, nil
		}
		plugin.Logger(ctx).Error("datadog_service_level_objective.getSLO", "query_error", err)
		return nil, err
	}

	return resp.GetData(), nil
}

// Root structure to represent the API response
type ApiResponse struct {
	Data  DataResponse  `json:"data"`
	Meta  MetaResponse  `json:"meta"`
	Links LinksResponse `json:"links"`
}

type DataResponse struct {
	Type       string        `json:"type"`
	Attributes SLOAttributes `json:"attributes"`
}

type SLOAttributes struct {
	SLOs []SLOResponse `json:"slos"`
}

type SLOResponse struct {
	Data *SLOData `json:"data"`
}

type SLOData struct {
	Type       string                `json:"type"`
	Attributes *SLOAttributesDetails `json:"attributes"`
	ID         string                `json:"id"`
}

type SLOAttributesDetails struct {
	Timeframe        *string         `json:"timeframe"`
	WarningThreshold *float64        `json:"warning_threshold"`
	Groups           interface{}     `json:"groups"`
	Thresholds       []Threshold     `json:"thresholds"`
	CreatedAt        *int64          `json:"created_at"`
	Status           SLOStatus       `json:"status"`
	TeamTags         []string        `json:"team_tags"`
	ModifiedAt       *int64          `json:"modified_at"`
	OverallStatus    []OverallStatus `json:"overall_status"`
	AllTags          []string        `json:"all_tags"`
	Query            *Query          `json:"query,omitempty"`
	RequestContext   RequestContext  `json:"request_context"`
	MonitorIDs       []int           `json:"monitor_ids,omitempty"`
	MonitorTags      []int           `json:"monitor_tags,omitempty"`
	ServiceTags      []string        `json:"service_tags"`
	Name             *string         `json:"name"`
	Description      interface{}     `json:"description"`
	SLOType          *string         `json:"slo_type"`
	TargetThreshold  *float64        `json:"target_threshold"`
	Creator          Creator         `json:"creator"`
	EnvTags          []string        `json:"env_tags"`
}

type Threshold struct {
	TargetDisplay  *string     `json:"target_display"`
	Timeframe      *string     `json:"timeframe"`
	WarningDisplay interface{} `json:"warning_display"`
	Warning        interface{} `json:"warning"`
	Target         *float64    `json:"target"`
}

type SLOStatus struct {
	CalculationError        interface{}     `json:"calculation_error"`
	IndexedAt               *int64          `json:"indexed_at"`
	SpanPrecision           *int            `json:"span_precision"`
	RawErrorBudgetRemaining BudgetRemaining `json:"raw_error_budget_remaining"`
	ErrorBudgetRemaining    *float64        `json:"error_budget_remaining"`
	State                   *string         `json:"state"`
	SLI                     *float64        `json:"sli"`
}

type BudgetRemaining struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type OverallStatus struct {
	Status                  float64         `json:"status"`
	Timeframe               string          `json:"timeframe"`
	IndexedAt               int64           `json:"indexed_at"`
	SpanPrecision           int             `json:"span_precision"`
	RawErrorBudgetRemaining BudgetRemaining `json:"raw_error_budget_remaining"`
	ErrorBudgetRemaining    float64         `json:"error_budget_remaining"`
	Error                   interface{}     `json:"error"`
	State                   string          `json:"state"`
	Target                  float64         `json:"target"`
}

type Query struct {
	Metrics     interface{} `json:"metrics"`
	Denominator string      `json:"denominator"`
	Numerator   string      `json:"numerator"`
}

type RequestContext struct {
	UserHasWriteAccessForSLO bool `json:"user_has_write_access_for_slo"`
}

type Creator struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type MetaResponse struct {
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	Number      *int    `json:"number"`
	FirstNumber *int    `json:"first_number"`
	PrevNumber  *int    `json:"prev_number"`
	NextNumber  *int    `json:"next_number"`
	LastNumber  *int    `json:"last_number"`
	Size        *int    `json:"size"`
	Type        *string `json:"type"`
	Total       *int    `json:"total"`
}

type LinksResponse struct {
	Self  string `json:"self"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
	First string `json:"first"`
}

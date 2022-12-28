## v0.3.0 [2022-12-28]

_What's new?_

- New tables added
  - [datadog_service_level_objective](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_service_level_objective) ([#25](https://github.com/turbot/steampipe-plugin-datadog/pull/25)) (Thanks [@whume](https://github.com/whume) for the contribution!)

## v0.2.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#21](https://github.com/turbot/steampipe-plugin-datadog/pull/21))
- Recompiled plugin with Go version `1.19`. ([#21](https://github.com/turbot/steampipe-plugin-datadog/pull/21))

## v0.1.0 [2022-04-19]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk-v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) ([#13](https://github.com/turbot/steampipe-plugin-datadog/pull/13))
- Recompiled plugin with Go version 1.18 ([#13](https://github.com/turbot/steampipe-plugin-datadog/pull/13))
- Added support for native Linux ARM and Mac M1 builds. ([#17](https://github.com/turbot/steampipe-plugin-datadog/pull/17))

## v0.0.3 [2021-12-15]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk-v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#10](https://github.com/turbot/steampipe-plugin-datadog/pull/10))
- Recompiled plugin with Go version 1.17 ([#10](https://github.com/turbot/steampipe-plugin-datadog/pull/10))

## v0.0.2 [2021-11-18]

_Bug fixes_

- Fixed: When setting the `api_url` config argument to "https://api.datadoghq.eu/", connections should now work properly ([#8](https://github.com/turbot/steampipe-plugin-datadog/pull/8))
- Fixed: All tables' list calls now handle API errors more effectively ([#8](https://github.com/turbot/steampipe-plugin-datadog/pull/8))

## v0.0.1 [2021-11-14]

_What's new?_

- New tables added
  - [datadog_dashboard](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_dashboard)
  - [datadog_integration_aws](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_integration_aws)
  - [datadog_log_event](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_log_event)
  - [datadog_logs_metric](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_logs_metric)
  - [datadog_monitor](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_monitor)
  - [datadog_permission](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_permission)
  - [datadog_role](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_role)
  - [datadog_security_monitoring_rule](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_security_monitoring_rule)
  - [datadog_security_monitoring_signal](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_security_monitoring_signal)

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#61](https://github.com/turbot/steampipe-plugin-datadog/pull/61))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#61](https://github.com/turbot/steampipe-plugin-datadog/pull/61))

## v0.8.0 [2024-02-06]

_Bug fixes_

- Fixed pagination in the `datadog_monitor` table to correctly return data instead of an error. ([#48](https://github.com/turbot/steampipe-plugin-datadog/pull/48)) (Thanks [@mdb](https://github.com/mdb) for the contribution!)

## v0.7.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#44](https://github.com/turbot/steampipe-plugin-datadog/pull/44))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#44](https://github.com/turbot/steampipe-plugin-datadog/pull/44))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-datadog/blob/main/docs/LICENSE). ([#44](https://github.com/turbot/steampipe-plugin-datadog/pull/44))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#43](https://github.com/turbot/steampipe-plugin-datadog/pull/43))

## v0.6.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#37](https://github.com/turbot/steampipe-plugin-datadog/pull/37))

## v0.6.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#34](https://github.com/turbot/steampipe-plugin-datadog/pull/34))
- Recompiled plugin with Go version `1.21`. ([#34](https://github.com/turbot/steampipe-plugin-datadog/pull/34))

## v0.5.0 [2023-03-23]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#28](https://github.com/turbot/steampipe-plugin-datadog/pull/28))

## v0.4.0 [2023-01-05]

_What's new?_

- New tables added
  - [datadog_host](https://hub.steampipe.io/plugins/turbot/datadog/tables/datadog_host) ([#24](https://github.com/turbot/steampipe-plugin-datadog/pull/24)) (Thanks [@marcus-crane](https://github.com/marcus-crane) for the contribution!)

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

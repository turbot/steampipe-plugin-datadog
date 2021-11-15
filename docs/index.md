---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/datadog.svg"
brand_color: "#632CA6"
display_name: "Datadog"
short_name: "datadog"
description: "Steampipe plugin for querying dashboards, users, roles and more from Datadog."
og_description: "Query Datadog with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/datadog-social-graphic.png"
---

# Datadog + Steampipe

[Datadog](https://www.datadoghq.com/) is the essential monitoring and security platform for cloud applications.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  name,
  type,
  overall_state
from
  datadog_monitor
where
  overall_state in ('Alert', 'Warn');
```

```
+-------------------------------+---------------+---------------+
| name                          | type          | overall_state |
+-------------------------------+---------------+---------------+
| Spend Alert                   | query alert   | Alert         |
| [Auto] Clock in sync with NTP | service check | Warn          |
+-------------------------------+---------------+---------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/datadog/tables)**

## Get started

### Install

Download and install the latest Datadog plugin:

```bash
steampipe plugin install datadog
```

### Configuration

Installing the latest datadog plugin will create a config file (`~/.steampipe/config/datadog.spc`) with a single connection named `datadog`:

```hcl
connection "datadog" {
  plugin = "datadog"

  # Get your API key from https://app.datadoghq.com/organization-settings/api-keys
  # Steampipe will resolve the API key in below order:
  #   1. The "api_key" specified here in the config
  #   2. The `DD_CLIENT_API_KEY` environment variable
  # api_key = "1a2345bc6d78e9d98fa7bcd6e5ef56a7"

  # Get your application key from https://app.datadoghq.com/organization-settings/application-keys
  # Steampipe will resolve the application key in below order:
  #   1. The "app_key" specified here in the config
  #   2. The `DD_CLIENT_APP_KEY` environment variable
  # app_key = "b1cf234c0ed4c567890b524a3b42f1bd91c111a1"

  # The API URL. By default it is pointed to "https://api.datadoghq.com/"
  # If working with the EU version of Datadog, use "https://api.datadoghq.eu/"
  # Please note that this URL must not end with the /api/ path.
  # api_url = "https://api.datadoghq.com/"
}
```

- `api_key` (required) - [API keys](https://docs.datadoghq.com/account_management/api-app-keys/#api-keys) are unique to an organization. An API key is required by the Datadog Agent to submit metrics and events to Datadog. [Get an API key](https://app.datadoghq.com/organization-settings/api-keys). May alternatively be set via the `DD_CLIENT_API_KEY` environment variable.

- `app_key` (required) - [Application keys](https://docs.datadoghq.com/account_management/api-app-keys/#application-keys) in conjunction with organization’s API key, give users access to Datadog’s programmatic API. Application keys are associated with the user account that created them and have the permissions and capabilities of the user who created them. [Get an application key](https://app.datadoghq.com/organization-settings/application-keys). May alternatively be set via the `DD_CLIENT_APP_KEY` environment variable.

- `api_url` (optional) - The API URL used for all requests. Defaults to "https://api.datadoghq.com/". If working with the EU version, this should be changed to "https://api.datadoghq.eu/".

## Get Involved

- Open source: https://github.com/turbot/steampipe-plugin-datadog
- Community: [Slack Channel](https://steampipe.io/community/join)

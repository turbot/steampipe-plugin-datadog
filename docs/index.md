---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/datadog.svg"
brand_color: "#FF9900"
display_name: "Datadog"
short_name: "datadog"
description: "Steampipe plugin for querying instances, buckets, databases and more from AWS."
og_description: "Query AWS with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/aws-social-graphic.png"
---

# WIP

# Datadog + Steampipe

[Datadog](https://www.datadoghq.com/) is the essential monitoring and security platform for cloud applications.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  email,
  id,
  name,
  status,
  disabled
from
  datadog_user;
```

```
+---------------------+--------------------------------------+-----------------------+---------+----------+
| email               | id                                   | name                  | status  | disabled |
+---------------------+--------------------------------------+-----------------------+---------+----------+
| subhajit@turbot.com | 0f550528-30e1-11ec-9255-da7ad0900002 | Subhajit Kumar Mondal | Active  | true     |
| lalit@turbot.com    | 45c9984e-30a0-11ec-9224-da7ad0900002 | Lalit Bhardwaj        | Active  | true     |
| subham@turbot.com   | 252dcb01-30e1-11ec-9255-da7ad0900002 | <null>                | Pending | false    |
+---------------------+--------------------------------------+-----------------------+---------+----------+
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
  plugin    = "datadog"

  # Authentication - API Key
  # Steampipe will resolve the API key in below order:
  #  2. The region specified here in the config for "api_key"
  #  2. The `DD_CLIENT_API_KEY` environment variable
  # api_key   = "1a2345bc6d78e9d98fa7bcd6e5ef56a7"

  # Authentication - APP Key
  # Steampipe will resolve the API key in below order:
  #  2. The region specified here in the config for "app_key"
  #  2. The `DD_CLIENT_APP_KEY` environment variable
  # app_key   = "b1cf234c0ed4c567890b524a3b42f1bd91c111a1"
}
```

- `api_key` (required) - [API](https://docs.datadoghq.com/account_management/api-app-keys/#api-keys) keys are unique to an organization. An API key is required by the Datadog Agent to submit metrics and events to Datadog.

- `app_key` (required) - [Application](https://docs.datadoghq.com/account_management/api-app-keys/#application-keys) keys, in conjunction with organization’s API key, give users access to Datadog’s programmatic API. Application keys are associated with the user account that created them and have the permissions and capabilities of the user who created them.

## Get Involved

- Open source: https://github.com/turbot/steampipe-plugin-datadog
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)

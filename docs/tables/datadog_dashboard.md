---
title: "Steampipe Table: datadog_dashboard - Query Datadog Dashboards using SQL"
description: "Allows users to query Datadog Dashboards, providing insights into the state, configuration, and management of dashboards in a Datadog account."
---

# Table: datadog_dashboard - Query Datadog Dashboards using SQL

Datadog Dashboards are an essential feature of the Datadog monitoring service, allowing users to visualize, analyze, and correlate data from different sources in one place. Dashboards can be customized to display data from various infrastructure components, applications, and services, providing a unified view of the system's performance and health. They are instrumental in identifying patterns, troubleshooting issues, and making data-driven decisions.

## Table Usage Guide

The `datadog_dashboard` table provides insights into the configuration and state of Dashboards within a Datadog account. As a DevOps engineer, use this table to explore dashboard-specific details, including its layout, title, widgets, and associated metadata. Utilize it to manage and monitor your dashboards, ensuring optimal system performance and proactive issue resolution.

## Examples

### Basic info
Explore the characteristics of your Datadog dashboards, such as layout type and accessibility settings. This can help you understand the structure and restrictions of your dashboards, enhancing your data visualization management.

```sql+postgres
select
  id,
  author_handle,
  layout_type,
  url,
  is_read_only,
  created_at,
  jsonb_pretty(restricted_roles) as restricted_roles
from
  datadog_dashboard;
```

```sql+sqlite
select
  id,
  author_handle,
  layout_type,
  url,
  is_read_only,
  created_at,
  restricted_roles
from
  datadog_dashboard;
```

### List dashboards with restricted editing access
Discover the segments that have limited edit access on dashboards. This can be useful in managing user permissions and maintaining security protocols within your organization.

```sql+postgres
select
  dashboard.id,
  title as dashboard_title,
  dr.users as role_users
from
  datadog_dashboard as dashboard,
  jsonb_array_elements_text(restricted_roles) as role,
  datadog_role as dr
where
  dr.id = role;
```

```sql+sqlite
select
  dashboard.id,
  title as dashboard_title,
  dr.users as role_users
from
  datadog_dashboard as dashboard,
  json_each(restricted_roles) as role,
  datadog_role as dr
where
  dr.id = role.value;
```

### List read-only dashboards (only dashboard author and admins can make changes to it)
Identify instances where dashboards have been set to read-only, allowing only the author and admins to make changes. This can be useful for maintaining control over dashboard configurations and preventing unauthorized modifications.

```sql+postgres
select
  id,
  title,
  is_read_only
from
  datadog_dashboard
where
  is_read_only;
```

```sql+sqlite
select
  id,
  title,
  is_read_only
from
  datadog_dashboard
where
  is_read_only;
```
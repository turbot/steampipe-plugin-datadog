---
title: "Steampipe Table: datadog_monitor - Query Datadog Monitors using SQL"
description: "Allows users to query Monitors in Datadog, specifically the monitor details, providing insights into monitor configurations and status."
---

# Table: datadog_monitor - Query Datadog Monitors using SQL

Datadog is a monitoring and analytics platform for developers, IT operations teams and business users. It brings together data from servers, containers, databases, and third-party services to make your stack entirely observable. Monitors in Datadog provide alerts and notifications based on the metrics and events collected from these systems.

## Table Usage Guide

The `datadog_monitor` table provides insights into Monitors within Datadog. As a DevOps engineer, explore monitor details through this table, including type, query, message, and options. Utilize it to uncover information about monitors, such as their configuration, status, and alert conditions.

## Examples

### Basic info
Analyze the settings to understand the overall state and priority of each monitor in your Datadog account, along with their creators and associated messages. This can be useful in assessing the health and urgency of different monitors, and identifying any potential issues or concerns.

```sql+postgres
select
  name,
  id,
  creator_email,
  overall_state,
  priority,
  query,
  message
from
  datadog_monitor;
```

```sql+sqlite
select
  name,
  id,
  creator_email,
  overall_state,
  priority,
  query,
  message
from
  datadog_monitor;
```

### List monitors in "Alert" and "Warn" state
Explore which monitors are in a state of alert or warning to identify potential issues and take necessary action in a timely manner. This can help in maintaining system health and preventing unexpected failures.

```sql+postgres
select
  name,
  type,
  created_at,
  message,
  overall_state
from
  datadog_monitor
where
  overall_state in ('Alert', 'Warn');
```

```sql+sqlite
select
  name,
  type,
  created_at,
  message,
  overall_state
from
  datadog_monitor
where
  overall_state in ('Alert', 'Warn');
```

### List monitors in "Alert" state with an "aws" tag
Explore which monitors are in an alert state and are tagged with 'aws'. This can be useful for quickly identifying potential issues within your 'aws' resources.

```sql+postgres
select
  name,
  type,
  created_at,
  overall_state,
  message,
  tags
from
  datadog_monitor
where
  overall_state in ('Alert') and
  tags @> '["aws"]'::jsonb;
```

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```
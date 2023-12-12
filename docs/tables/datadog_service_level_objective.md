---
title: "Steampipe Table: datadog_service_level_objective - Query Datadog Service Level Objectives using SQL"
description: "Allows users to query Service Level Objectives in Datadog, specifically the metrics related to the performance and availability of services."
---

# Table: datadog_service_level_objective - Query Datadog Service Level Objectives using SQL

A Service Level Objective (SLO) in Datadog is a target that a service aims to achieve over a given period. It quantifies the long-term performance of a service in terms of its availability or response time. SLOs are typically used to set expectations for service performance and to measure whether those expectations are being met.

## Table Usage Guide

The `datadog_service_level_objective` table provides insights into Service Level Objectives within Datadog. As an SRE or DevOps engineer, explore SLO-specific details through this table, including thresholds, timeframes, and associated metadata. Utilize it to uncover information about SLOs, such as those that are not meeting their targets, the historical performance of SLOs, and the verification of SLO configurations.

## Examples

### Basic info
Explore which service level objectives have been set up in your Datadog account and when they were created. This can assist in understanding what performance metrics are being monitored and by whom, aiding in the overall management and optimization of your services.

```sql+postgres
select
  name,
  type,
  thresholds,
  created_at,
  creator_email
from
  datadog_service_level_objective;
```

```sql+sqlite
select
  name,
  type,
  thresholds,
  created_at,
  creator_email
from
  datadog_service_level_objective;
```

### List metric type SLOs
Explore which service level objectives (SLOs) are based on metrics in your Datadog account. This can help you assess the performance of specific services or systems over time.

```sql+postgres
select
  name,
  type,
  created_at,
  monitor_ids
from
  datadog_service_level_objective
where
  type = 'metric';
```

```sql+sqlite
select
  name,
  type,
  created_at,
  monitor_ids
from
  datadog_service_level_objective
where
  type = 'metric';
```

### List SLOs that are type monitor and have thresholds set to 2.5 9's over 7 days
Identify service level objectives (SLOs) that are classified as 'monitor' type and have specific thresholds set. This can be useful in managing and monitoring system performance over a week.

```sql+postgres
select
  name,
  type,
  thresholds,
  created_at
from
  datadog_service_level_objective
where
  type = 'monitor'
  and thresholds @> '[{"target":99.5,"target_display":"99.5","timeframe":"7d"}]'::jsonb;
```

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```
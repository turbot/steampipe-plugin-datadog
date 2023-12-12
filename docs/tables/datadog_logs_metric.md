---
title: "Steampipe Table: datadog_logs_metric - Query Datadog Logs Metrics using SQL"
description: "Allows users to query Datadog Logs Metrics, specifically providing insights into log-based metrics, their respective query definitions, and related metadata."
---

# Table: datadog_logs_metric - Query Datadog Logs Metrics using SQL

Datadog Logs Metrics is a feature within the Datadog Log Management service that allows you to generate metrics from logs and analyze data in the form of counts, distributions, or gauges over a given period. This feature provides a way to measure high volumes of log data, and track application performance, system behavior, and key business metrics. It is a useful tool for monitoring, alerting, and conducting historical analysis.

## Table Usage Guide

The `datadog_logs_metric` table provides insights into log-based metrics within Datadog Log Management. As a DevOps engineer or system administrator, explore detailed information about these metrics, including their types, query definitions, and associated metadata. Utilize it to monitor and analyze system performance, application behavior, and business metrics, and to create alerts based on these metrics.

## Examples

### Basic info
Explore the configuration of your log-based metrics in Datadog to understand the aggregation types and paths. This can help in identifying any issues or potential improvements in your log management strategy.

```sql+postgres
select
  id,
  compute_aggregation_type,
  compute_path,
  filter_query,
  jsonb_pretty(group_by) as group_by
from
  datadog_logs_metric;
```

```sql+sqlite
select
  id,
  compute_aggregation_type,
  compute_path,
  filter_query,
  group_by
from
  datadog_logs_metric;
```

### Get count of metrics by compute_aggregation_type
Explore the distribution of metrics based on their aggregation types to gain insights into the different computation methods used in your Datadog logs.

```sql+postgres
select
  count(*),
  compute_aggregation_type
from
  datadog_logs_metric
group by
  compute_aggregation_type;
```

```sql+sqlite
select
  count(*),
  compute_aggregation_type
from
  datadog_logs_metric
group by
  compute_aggregation_type;
```

### Get details of filter_query and group_by clause for a specific log metric
Explore the specific log metrics to understand the grouping and filtering details. This can be beneficial in analyzing how data is categorized and segmented for a particular metric.

```sql+postgres
select
  filter_query,
  jsonb_pretty(group_by) as group_by
from
  datadog_logs_metric
where
  id = 's3_bucket_by_region';
```

```sql+sqlite
select
  filter_query,
  group_by
from
  datadog_logs_metric
where
  id = 's3_bucket_by_region';
```